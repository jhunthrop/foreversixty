# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 37.3. Weights run: 1.8s. Verify run: 1.1s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.646 ± 0.017, crit=0.911 ± 0.025 per rating point (14 rating = 1%, 12.750 per %), hit=1.407 ± 0.069 per rating point (10 rating = 1%, 14.066 per %), melee_haste=8.418 ± 0.377

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.57 DPS) [crafted]; Brawler's Leather Hood (252504, -0.62 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 3.9 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.12 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Grave Shroud (279865, -0.03 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Dark Leather Cloak (2316, -0.09 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.84 DPS) | yes | Veteran's Chain Shirt (250488, -0.14 DPS) [crafted]; Defender's Leather Armor (252434, -0.17 DPS) [crafted]; Brawler's Leather Armor (252490, -0.23 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.42 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS) [quest]; Bravo's Armbands (270015, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.17 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.67 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Fletcher's Gloves (7348, -0.14 DPS) [crafted]; Polar Gauntlets (7606, -0.17 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.09 DPS) [dungeon]; Ruffian Belt (5975, -0.25 DPS) [world]; Brawler's Leather Belt (252428, -0.31 DPS) [crafted] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.93 DPS) | yes | Veteran's Chain Leggings (250493, -0.03 DPS) [crafted]; Defender's Leather Pants (252445, -0.06 DPS) [crafted]; Totemic Leather Pants (252446, -0.17 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.2 attack_power points (0.56 DPS) | yes | Veteran's Boots (250503, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.09 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.6 attack_power points (0.45 DPS) | yes | Signet of the Zhevra (285330, -0.28 DPS) [world]; The 1 Ring (8350, -0.33 DPS) [world]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Signet of the Zhevra (285330, -0.20 DPS, sim-verified) [world]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.25 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (15.09 DPS) | yes | Smite's Mighty Hammer (7230, -2.19 DPS, sim-verified) [dungeon]; Duskbringer (2205, -2.70 DPS) [dungeon]; Monstrous Cleaver (279864, -3.07 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 80.9. Weights run: 1.8s. Verify run: 1.4s. 404 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.640 ± 0.017, crit=0.905 ± 0.024 per rating point (14 rating = 1%, 12.670 per %), hit=1.555 ± 0.067 per rating point (10 rating = 1%, 15.554 per %), melee_haste=8.455 ± 0.405

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.37 DPS) | yes | Veteran's Chain Helm (250498, -0.11 DPS) [crafted]; Defender's Leather Helm (252455, -0.11 DPS) [crafted]; Barbaric Iron Helm (7915, -0.12 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.74 DPS) | yes | Kaleidoscope Chain (13084, -0.18 DPS) [world_drop]; River Pride Choker (13087, -0.32 DPS) [world_drop]; Sentinel's Medallion (19541, -0.47 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.0 attack_power points (0.90 DPS) | yes | Barbaric Iron Shoulders (7913, -0.06 DPS) [crafted]; Golden Scale Shoulders (3841, -0.16 DPS) [crafted]; Mail Combat Spaulders (6404, -0.16 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.6 attack_power points (0.56 DPS) | yes | Hawkeye's Cloak (14593, -0.00 DPS) [world_drop]; Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Slayer's Cape (14752, -0.13 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.58 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.32 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.84 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Bands of Serra'kis (6902, -0.21 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.22 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.16 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.18 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.21 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (1.26 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 26.5 attack_power points (1.39 DPS) | yes | Ferine Leggings (6690, +0.00 DPS, sim-verified) [dungeon]; Golden Scale Leggings (3843, -0.24 DPS) [crafted]; Chausses of Westfall (6087, -0.24 DPS) [quest] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (80.9 DPS) | yes | Brawler's Leather Boots (252439, -0.04 DPS) [crafted]; Alacritous Treads (277234, -0.04 DPS) [quest]; Trouncing Boots (4464, -1.38 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.9 attack_power points (0.94 DPS) | yes | Ironspine's Eye (7686, -0.22 DPS) [dungeon]; Tiger Band (6749, -0.31 DPS) [quest]; Silverlaine's Family Seal (6321, -0.42 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.8 attack_power points (0.83 DPS) | yes | Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.31 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.92 DPS) [dungeon]; Viscous Hammer (13045, -18.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 123.5. Weights run: 1.9s. Verify run: 1.3s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.007 ± 0.023, crit=1.425 ± 0.033 per rating point (14 rating = 1%, 19.947 per %), hit=1.885 ± 0.083 per rating point (10 rating = 1%, 18.848 per %), melee_haste=11.050 ± 0.467

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 45.9 attack_power points (2.64 DPS) | yes | White Bandit Mask (10008, -0.74 DPS) [crafted]; Chromite Barbute (8142, -0.81 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -1.03 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.15 DPS) | yes | Kaleidoscope Chain (13084, -0.46 DPS) [world_drop]; Sentinel's Medallion (19540, -0.51 DPS) [rep]; Ghostshard Talisman (7731, -0.52 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Tracker Epaulets (2278, -0.05 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Sunburn Spaulders (274751, -1.29 DPS, sim-verified) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Cloak (14593, -0.29 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.34 DPS) [quest]; Dark Hooded Cape (5257, -1.21 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 35.1 attack_power points (2.02 DPS) | yes | Kolkar Marauder Chain (6773, -0.00 DPS) [quest]; Avenger's Armor (1488, -0.29 DPS) [dungeon]; Jouster's Chestplate (8157, -0.29 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.15 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Pugilist Bracers (4438, -0.23 DPS) [dungeon]; Yorgen Bracers (13012, -0.29 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 39.9 attack_power points (2.29 DPS) | yes | Scarlet Gauntlets (10331, -0.45 DPS, sim-verified) [dungeon]; Gauntlets of Divinity (7724, -0.46 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.46 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.72 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.00 DPS) [rep]; Highlander's Chain Girdle (20089, -0.12 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.41 DPS) | yes | Firemane Leggings (13129, -0.23 DPS) [world_drop]; Symbolic Legplates (14829, -0.34 DPS) [world_drop]; Orcish War Leggings (7929, -0.46 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 33.0 attack_power points (1.90 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Obsidian Greaves (13068, -0.40 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.1 attack_power points (1.38 DPS) | yes | Assault Band (13095, -0.23 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.15 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Ironspine's Eye (7686, -0.17 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -19.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 164.7. Weights run: 1.9s. Verify run: 1.3s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.382 ± 0.031, crit=1.947 ± 0.044 per rating point (14 rating = 1%, 27.258 per %), hit=2.379 ± 0.104 per rating point (10 rating = 1%, 23.791 per %), melee_haste=12.733 ± 1.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 79.6 attack_power points (4.71 DPS) | yes | Raging Berserker's Helm (7719, -1.58 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -1.78 DPS) [dungeon]; Ornate Mithril Helm (7937, -1.80 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.0 attack_power points (1.71 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.53 DPS) [quest]; Sentinel's Medallion (19539, -0.73 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 51.5 attack_power points (3.04 DPS) | yes | Officer's Pauldrons (250576, -0.70 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.83 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.88 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.5 attack_power points (1.92 DPS) | yes | Dark Hooded Cape (5257, -0.59 DPS) [world]; Sergeant Major's Cape (16336, -0.65 DPS) [pvp]; Dark Phantom Cape (13122, -0.70 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 60.3 attack_power points (3.56 DPS) | yes | Mixologist's Tunic (12793, -0.32 DPS) [dungeon]; Warforged Chestplate (11195, -0.44 DPS) [quest]; Warbear Harness (15064, -0.66 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 32.5 attack_power points (1.92 DPS) | yes | Runed Golem Shackles (12550, -0.10 DPS) [dungeon]; Deepfury Bracers (13120, -0.18 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.18 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 62.2 attack_power points (3.68 DPS) | yes | Raider Gloves (272100, -0.58 DPS) [vendor]; Officer's Gloves (250551, -0.86 DPS) [crafted]; Gloves of Holy Might (867, -0.88 DPS) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 51.5 attack_power points (3.04 DPS) | yes | Girdle of Beastial Fury (11686, -0.23 DPS) [dungeon]; Prowler's Leather Waistguard (252473, -0.24 DPS) [crafted]; Highlander's Plate Girdle (20124, -1.47 DPS, sim-verified) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 67.5 attack_power points (3.99 DPS) | yes | Gryphon Rider's Leggings (9652, -0.81 DPS) [quest]; Centurion Legplates (10740, -0.81 DPS) [quest]; Stormshroud Pants (15057, -0.96 DPS, sim-verified) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 48.8 attack_power points (2.88 DPS) | yes | Prowler's Leather Boots (252468, -0.29 DPS) [crafted]; Skulker's Leather Boots (252469, -0.39 DPS) [crafted]; Officer's Sabatons (250561, -0.46 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.8 attack_power points (2.59 DPS) | yes | Thunderbrow Ring (13097, -1.30 DPS) [world_drop]; Ironspine's Eye (7686, -1.33 DPS) [dungeon]; Mark of Kern (2262, -1.41 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 34.4 attack_power points (2.04 DPS) | yes | Thunderbrow Ring (13097, -0.75 DPS) [world_drop]; Ironspine's Eye (7686, -0.78 DPS) [dungeon]; Mark of Kern (2262, -0.85 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (164.7 DPS) | yes | - |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (164.7 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (164.7 DPS) | yes | Warmonger (13052, -0.07 DPS) [world_drop]; Taran Icebreaker (2915, -0.90 DPS) [world_drop]; Blight (7959, -2.36 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Sanctified Orb; main_hand: Thorium Greatmace

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 254.5. Weights run: 1.9s. Verify run: 1.3s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.958 ± 0.046, crit=2.771 ± 0.066 per rating point (14 rating = 1%, 38.791 per %), hit=3.227 ± 0.165 per rating point (10 rating = 1%, 32.272 per %), melee_haste=19.069 ± 1.879

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 181.7 attack_power points (11.43 DPS) | yes | Eye of Rend (12587, -1.98 DPS, sim-verified) [dungeon]; Mask of the Unforgiven (13404, -4.93 DPS) [dungeon]; Knight-Lieutenant's Plate Helm (220804, -5.16 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 64.8 attack_power points (4.08 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Pendant of Celerity (22340, -0.20 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.7 attack_power points (5.77 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.20 DPS) [dungeon]; Highlander's Plate Spaulders (20057, -1.18 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (3.79 DPS) | yes | Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Windshear Cape (20691, -0.84 DPS) [world]; Cloak of the Honor Guard (20073, -1.04 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (254.5 DPS) | yes | Dawn Armor (252483, -0.20 DPS) [crafted]; Savage Gladiator Chain (11726, -1.63 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.98 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (254.5 DPS) | yes | Forest Stalker's Bracers (19587, -0.61 DPS) [rep]; Berserker Bracers (19578, -0.86 DPS) [rep]; Bracers of Undead Slaying (23090, -2.82 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-verified (254.5 DPS) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Raider Gauntlets (272095, -0.08 DPS) [vendor]; Raider Gloves (272099, -0.32 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 98.2 attack_power points (6.18 DPS) | yes | Ferocity of the Timbermaw (227805, -0.48 DPS) [vendor]; Highlander's Plate Girdle (20041, -1.38 DPS) [rep]; Belt of Preserved Heads (20216, -2.42 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (254.5 DPS) | yes | Titanic Leggings (22385, -0.57 DPS) [crafted]; Sentinel's Plate Legguards (237825, -0.57 DPS) [vendor]; Cloudkeeper Legplates (14554, -3.64 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 71.4 attack_power points (4.49 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Lamellar Sabatons (227146, -0.39 DPS) [vendor]; Drudge Boots (21532, -0.60 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (254.5 DPS) | yes | Tarnished Elven Ring (18500, -1.60 DPS) [dungeon]; Cutthroat's Signet (272408, -1.72 DPS) [vendor]; Naglering (11669, -4.12 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (254.5 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.49 DPS) [vendor]; Naglering (11669, -2.58 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (254.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (254.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -1.43 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (254.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -5.45 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-verified (254.5 DPS) | yes | Libram of Fervor (23203, -1.35 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 54003000000000000-0530000000000000-05205331001330320)

Set DPS (verified): 538.5. Weights run: 2.3s. Verify run: 1.6s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.653 ± 0.038, crit=2.329 ± 0.053 per rating point (14 rating = 1%, 32.600 per %), hit=3.643 ± 0.230 per rating point (10 rating = 1%, 36.426 per %), melee_haste=13.590 ± 1.771

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 177.7 attack_power points (23.27 DPS) | yes | Helm of the Executioner (22411, -9.69 DPS) [dungeon]; Stalwart Helm (250599, -10.27 DPS) [crafted]; Mask of the Unforgiven (13404, -10.47 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 61.2 attack_power points (8.02 DPS) | yes | Beads of Ogre Might (22150, -0.10 DPS) [quest]; Mark of Fordring (15411, -0.34 DPS) [quest]; Medallion of the Dawn (22659, -0.61 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Wyrmhide Spaulders (12082, -0.57 DPS) [quest]; Truestrike Shoulders (12927, -6.19 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 64.4 attack_power points (8.44 DPS) | yes | Cape of the Black Baron (13340, -2.57 DPS) [dungeon]; Windshear Cape (20691, -2.88 DPS) [world]; Cloak of the Honor Guard (20073, -2.90 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.62 DPS) [crafted]; Obsidian Mail Tunic (22191, -3.45 DPS) [crafted]; Breastplate of Undead Slaying (23087, -19.31 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.76 DPS) [rep]; Berserker Bracers (19578, -1.83 DPS) [rep]; Bracers of Undead Slaying (23090, -4.29 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Raider Gauntlets (272095, -1.33 DPS) [vendor]; Marshal's Lamellar Gauntlets (231650, -2.32 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.88 DPS) [vendor]; Marksman's Girdle (22232, -2.73 DPS) [dungeon]; Belt of Preserved Heads (20216, -5.56 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -1.37 DPS) [vendor]; Sentinel's Chain Leggings (237819, -1.57 DPS) [vendor]; Cloudkeeper Legplates (14554, -2.99 DPS, sim-verified) [world_drop] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | 71.1 attack_power points (9.31 DPS) | yes | Knight-Lieutenant's Lamellar Sabatons (227146, -1.59 DPS) [vendor]; Drudge Boots (21532, -1.89 DPS) [quest]; Windreaver Greaves (13967, -3.69 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.12 DPS) [dungeon]; Cutthroat's Signet (272408, -3.33 DPS) [vendor]; Naglering (11669, -6.79 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.65 DPS) [dungeon]; Cutthroat's Signet (272408, -0.87 DPS) [vendor]; Naglering (11669, -3.92 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+11.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -8.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Invocation (249442, +0.00 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 37.9. Weights run: 1.8s. Verify run: 1.1s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.646 ± 0.017, crit=0.911 ± 0.025 per rating point (14 rating = 1%, 12.750 per %), hit=1.407 ± 0.069 per rating point (10 rating = 1%, 14.066 per %), melee_haste=8.418 ± 0.377

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.57 DPS) [crafted]; Brawler's Leather Hood (252504, -0.62 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 3.9 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.12 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Grave Shroud (279865, -0.03 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Subterranean Cape (14149, -0.08 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.84 DPS) | yes | Veteran's Chain Shirt (250488, -0.14 DPS) [crafted]; Defender's Leather Armor (252434, -0.17 DPS) [crafted]; Brawler's Leather Armor (252490, -0.23 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.42 DPS) | yes | Raptorcrest Bracers (270010, -0.17 DPS) [quest]; Bristlebark Bindings (14569, -0.17 DPS) [world_drop]; Runed Copper Bracers (2854, -0.25 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.67 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Fletcher's Gloves (7348, -0.14 DPS) [crafted]; Blackened Defias Gloves (10401, -0.17 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Cobrahn's Grasp (6460, -0.09 DPS) [dungeon]; Ruffian Belt (5975, -0.25 DPS) [world]; Brawler's Leather Belt (252428, -0.31 DPS) [crafted] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 21.2 attack_power points (0.89 DPS) | yes | Defender's Leather Pants (252445, -0.03 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted]; Hulking Leggings (14748, -0.14 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.2 attack_power points (0.56 DPS) | yes | Veteran's Boots (250503, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.09 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.6 attack_power points (0.45 DPS) | yes | Loop of Sacrifice (281673, -0.19 DPS) [quest]; Signet of the Zhevra (285330, -0.28 DPS) [world]; The 1 Ring (8350, -0.33 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Signet of the Zhevra (285330, -0.17 DPS) [world]; The 1 Ring (8350, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (12.88 DPS) | yes | Duskbringer (2205, -0.49 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.65 DPS, sim-verified) [dungeon]; Monstrous Cleaver (279864, -0.86 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 87.1. Weights run: 1.8s. Verify run: 1.4s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.640 ± 0.017, crit=0.905 ± 0.024 per rating point (14 rating = 1%, 12.670 per %), hit=1.555 ± 0.067 per rating point (10 rating = 1%, 15.554 per %), melee_haste=8.455 ± 0.405

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.37 DPS) | yes | Veteran's Chain Helm (250498, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.11 DPS) [crafted]; Barbaric Iron Helm (7915, -0.12 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.74 DPS) | yes | River Pride Choker (13087, -0.32 DPS) [world_drop]; Scout's Medallion (19537, -0.47 DPS) [rep]; Kaleidoscope Chain (13084, -0.50 DPS, sim-verified) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.0 attack_power points (0.90 DPS) | yes | Barbaric Iron Shoulders (7913, -0.06 DPS) [crafted]; Golden Scale Shoulders (3841, -0.16 DPS) [crafted]; Mail Combat Spaulders (6404, -0.16 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.5 attack_power points (0.55 DPS) | yes | Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Wildhunter Cloak (16658, -0.03 DPS) [quest]; Slayer's Cape (14752, -0.13 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.58 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.32 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.84 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Bands of Serra'kis (6902, -0.21 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.22 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.16 DPS) | yes | Warsong Gauntlets (16978, -0.11 DPS) [quest]; The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.18 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (1.26 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 26.5 attack_power points (1.39 DPS) | yes | Ferine Leggings (6690, -0.03 DPS) [dungeon]; Golden Scale Leggings (3843, -0.24 DPS) [crafted]; Slayer's Pants (14757, -0.24 DPS) [world_drop] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (87.1 DPS) | yes | Brawler's Leather Boots (252439, -0.04 DPS) [crafted]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Trouncing Boots (4464, -1.31 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.9 attack_power points (0.94 DPS) | yes | Ironspine's Eye (7686, -0.22 DPS) [dungeon]; Tiger Band (6749, -0.31 DPS) [quest]; Band of the Fist (17694, -0.39 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.8 attack_power points (0.83 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Band of the Fist (17694, -0.28 DPS) [quest]; Ironspine's Eye (7686, -0.45 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.92 DPS) [dungeon]; Viscous Hammer (13045, -21.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 132.0. Weights run: 1.9s. Verify run: 1.3s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.007 ± 0.023, crit=1.425 ± 0.033 per rating point (14 rating = 1%, 19.947 per %), hit=1.885 ± 0.083 per rating point (10 rating = 1%, 18.848 per %), melee_haste=11.050 ± 0.467

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 45.9 attack_power points (2.64 DPS) | yes | White Bandit Mask (10008, -0.74 DPS) [crafted]; Hard Gold Coif (250537, -1.03 DPS) [crafted]; Chromite Barbute (8142, -1.33 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.15 DPS) | yes | Ethereal Talisman (4430, -0.34 DPS) [quest]; Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.32 DPS) | yes | Hard Gold Pauldrons (250539, +0.00 DPS, sim-verified) [crafted]; Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.17 DPS) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | sim-verified (132.0 DPS) | yes | Hawkeye's Cloak (14593, -0.29 DPS) [world_drop]; Wildhunter Cloak (16658, -0.46 DPS) [quest]; Dark Hooded Cape (5257, -1.61 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 35.1 attack_power points (2.02 DPS) | yes | Kolkar Marauder Chain (6773, -0.00 DPS) [quest]; Avenger's Armor (1488, -0.29 DPS) [dungeon]; Jouster's Chestplate (8157, -0.29 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.15 DPS) | yes | Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.64 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 39.9 attack_power points (2.29 DPS) | yes | Scarlet Gauntlets (10331, -0.34 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.46 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.46 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.72 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.00 DPS) [rep]; Defiler's Chain Girdle (20153, -0.12 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.41 DPS) | yes | Symbolic Legplates (14829, -0.34 DPS) [world_drop]; Orcish War Leggings (7929, -0.46 DPS) [crafted]; Firemane Leggings (13129, -0.68 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 33.0 attack_power points (1.90 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Obsidian Greaves (13068, -0.40 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.1 attack_power points (1.38 DPS) | yes | Assault Band (13095, -0.23 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.15 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Ironspine's Eye (7686, -0.17 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -23.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: First Sergeant's Cloak; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 172.6. Weights run: 1.9s. Verify run: 1.2s. 702 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.382 ± 0.031, crit=1.947 ± 0.044 per rating point (14 rating = 1%, 27.258 per %), hit=2.379 ± 0.104 per rating point (10 rating = 1%, 23.791 per %), melee_haste=12.733 ± 1.090

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 55.9 attack_power points (3.30 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -0.37 DPS) [dungeon]; Ornate Mithril Helm (7937, -0.39 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.0 attack_power points (1.71 DPS) | yes | Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.53 DPS) [quest]; Scout's Medallion (19535, -0.73 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 39.7 attack_power points (2.34 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.13 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.5 attack_power points (1.92 DPS) | yes | Dark Hooded Cape (5257, -0.59 DPS) [world]; First Sergeant's Cloak (16340, -0.65 DPS) [pvp]; Dark Phantom Cape (13122, -0.70 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 60.3 attack_power points (3.56 DPS) | yes | Mixologist's Tunic (12793, -0.32 DPS) [dungeon]; Warforged Chestplate (11195, -0.44 DPS) [quest]; Warbear Harness (15064, -0.66 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 32.5 attack_power points (1.92 DPS) | yes | Runed Golem Shackles (12550, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 62.2 attack_power points (3.68 DPS) | yes | Raider Gloves (272100, -0.58 DPS) [vendor]; Officer's Gloves (250551, -0.86 DPS) [crafted]; Gloves of Holy Might (867, -0.88 DPS) [world_drop] |
| waist | Defiler's Plate Girdle (20205) | The Defilers [rep] | 49.3 attack_power points (2.91 DPS) | yes | Girdle of Beastial Fury (11686, -0.10 DPS) [dungeon]; Prowler's Leather Waistguard (252473, -0.11 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.12 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 67.5 attack_power points (3.99 DPS) | yes | Serpentskin Leggings (8262, -0.99 DPS) [world_drop]; Golem Shard Leggings (13074, -1.13 DPS) [world_drop]; Stormshroud Pants (15057, -1.38 DPS, sim-verified) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 48.8 attack_power points (2.88 DPS) | yes | Prowler's Leather Boots (252468, -0.29 DPS) [crafted]; Skulker's Leather Boots (252469, -0.39 DPS) [crafted]; Officer's Sabatons (250561, -0.46 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.8 attack_power points (2.59 DPS) | yes | White Bone Band (11862, -1.17 DPS) [quest]; Thunderbrow Ring (13097, -1.30 DPS) [world_drop]; Ironspine's Eye (7686, -1.33 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 34.4 attack_power points (2.04 DPS) | yes | White Bone Band (11862, -0.62 DPS) [quest]; Thunderbrow Ring (13097, -0.75 DPS) [world_drop]; Ironspine's Eye (7686, -0.78 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (172.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.6 DPS) | yes | Molten Heart of the Mountain (249470, -1.28 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (172.6 DPS) | yes | Warmonger (13052, -0.07 DPS) [world_drop]; Taran Icebreaker (2915, -0.90 DPS) [world_drop]; The Jackhammer (9423, -3.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Defiler's Plate Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 266.5. Weights run: 1.9s. Verify run: 1.3s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.958 ± 0.046, crit=2.771 ± 0.066 per rating point (14 rating = 1%, 38.791 per %), hit=3.227 ± 0.165 per rating point (10 rating = 1%, 32.272 per %), melee_haste=19.069 ± 1.879

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 181.7 attack_power points (11.43 DPS) | yes | Eye of Rend (12587, -1.85 DPS, sim-verified) [dungeon]; Mask of the Unforgiven (13404, -4.93 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -5.16 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 64.8 attack_power points (4.08 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Pendant of Celerity (22340, -0.20 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.7 attack_power points (5.77 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.20 DPS) [dungeon]; Defiler's Plate Spaulders (20212, -1.18 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (3.79 DPS) | yes | Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Windshear Cape (20691, -0.84 DPS) [world]; Deathguard's Cloak (20068, -1.04 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (266.5 DPS) | yes | Dawn Armor (252483, -0.20 DPS) [crafted]; Savage Gladiator Chain (11726, -1.63 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.43 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (266.5 DPS) | yes | Forest Stalker's Bracers (19587, -0.61 DPS) [rep]; Berserker Bracers (19578, -0.86 DPS) [rep]; Bracers of Undead Slaying (23090, -2.23 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-verified (266.5 DPS) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Raider Gauntlets (272095, -0.08 DPS) [vendor]; Raider Gloves (272099, -0.32 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 98.2 attack_power points (6.18 DPS) | yes | Ferocity of the Timbermaw (227805, -0.48 DPS) [vendor]; Defiler's Plate Girdle (20204, -1.38 DPS) [rep]; Belt of Preserved Heads (20216, -3.28 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (266.5 DPS) | yes | Titanic Leggings (22385, -0.57 DPS) [crafted]; Sentinel's Plate Legguards (237825, -0.57 DPS) [vendor]; Cloudkeeper Legplates (14554, -3.83 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 71.4 attack_power points (4.49 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS, sim-verified) [dungeon]; Drudge Boots (21532, -0.60 DPS) [quest]; Battlechaser's Greaves (12555, -0.95 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (266.5 DPS) | yes | Tarnished Elven Ring (18500, -1.60 DPS) [dungeon]; Cutthroat's Signet (272408, -1.72 DPS) [vendor]; Naglering (11669, -3.93 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (266.5 DPS) | yes | Tarnished Elven Ring (18500, -0.37 DPS) [dungeon]; Cutthroat's Signet (272408, -0.49 DPS) [vendor]; Naglering (11669, -2.26 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (266.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Talisman of Ascendance (22678, -2.55 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (266.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -1.68 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (266.5 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -7.27 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-verified (266.5 DPS) | yes | Libram of Fervor (23203, -1.34 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 54003000000000000-0530000000000000-05205331001330320)

Set DPS (verified): 551.4. Weights run: 2.3s. Verify run: 1.6s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.653 ± 0.038, crit=2.329 ± 0.053 per rating point (14 rating = 1%, 32.600 per %), hit=3.643 ± 0.230 per rating point (10 rating = 1%, 36.426 per %), melee_haste=13.590 ± 1.771

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 177.7 attack_power points (23.27 DPS) | yes | Helm of the Executioner (22411, -9.69 DPS) [dungeon]; Mask of the Unforgiven (13404, -9.82 DPS, sim-verified) [dungeon]; Stalwart Helm (250599, -10.27 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 61.2 attack_power points (8.02 DPS) | yes | Beads of Ogre Might (22150, -0.10 DPS) [quest]; Mark of Fordring (15411, -0.34 DPS) [quest]; Medallion of the Dawn (22659, -0.61 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Wyrmhide Spaulders (12082, -0.57 DPS) [quest]; Truestrike Shoulders (12927, -6.38 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 64.4 attack_power points (8.44 DPS) | yes | Cape of the Black Baron (13340, -2.57 DPS) [dungeon]; Windshear Cape (20691, -2.88 DPS) [world]; Deathguard's Cloak (20068, -2.90 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.62 DPS) [crafted]; Obsidian Mail Tunic (22191, -3.45 DPS) [crafted]; Breastplate of Undead Slaying (23087, -18.07 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.76 DPS) [rep]; Berserker Bracers (19578, -1.83 DPS) [rep]; Bracers of Undead Slaying (23090, -3.85 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Raider Gauntlets (272095, -1.33 DPS) [vendor]; Timbermaw Brawlers (19049, -2.41 DPS) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.88 DPS) [vendor]; Marksman's Girdle (22232, -2.73 DPS) [dungeon]; Belt of Preserved Heads (20216, -5.55 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Sentinel's Plate Legguards (237825, -1.37 DPS) [vendor]; Sentinel's Chain Leggings (237819, -1.57 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | 71.1 attack_power points (9.31 DPS) | yes | Drudge Boots (21532, -1.89 DPS) [quest]; Fine Dawn Treaders (227815, -2.16 DPS) [vendor]; Windreaver Greaves (13967, -4.10 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.12 DPS) [dungeon]; Cutthroat's Signet (272408, -3.33 DPS) [vendor]; Naglering (11669, -6.76 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.65 DPS) [dungeon]; Cutthroat's Signet (272408, -0.87 DPS) [vendor]; Naglering (11669, -3.69 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Talisman of Ascendance (22678, -6.04 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -14.07 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Invocation (249442, +0.00 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

