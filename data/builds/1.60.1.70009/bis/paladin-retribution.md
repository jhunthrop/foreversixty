# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 36.3. Weights run: 1.4s. Verify run: 0.8s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.659 ± 0.018, crit=0.931 ± 0.025 per rating point (14 rating = 1%, 13.030 per %), hit=1.308 ± 0.069 per rating point (10 rating = 1%, 13.082 per %), melee_haste=8.619 ± 0.501

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.82 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.55 DPS) [crafted]; Brawler's Leather Hood (252504, -0.61 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.0 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.11 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.33 DPS) | yes | Grave Shroud (279865, -0.03 DPS) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Dark Leather Cloak (2316, -0.08 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Veteran's Chain Shirt (250488, -0.14 DPS) [crafted]; Defender's Leather Armor (252434, -0.17 DPS) [crafted]; Brawler's Leather Armor (252490, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.41 DPS) | yes | Cryptwalker Bracers (280095, -0.08 DPS) [quest]; Bravo's Armbands (270015, -0.14 DPS) [quest]; Bristlebark Bindings (14569, -0.17 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.66 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Fletcher's Gloves (7348, -0.12 DPS) [crafted]; Polar Gauntlets (7606, -0.16 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.74 DPS) | yes | Cobrahn's Grasp (6460, -0.08 DPS) [dungeon]; Ruffian Belt (5975, -0.25 DPS) [world]; Brawler's Leather Belt (252428, -0.30 DPS) [crafted] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.91 DPS) | yes | Veteran's Chain Leggings (250493, -0.03 DPS) [crafted]; Defender's Leather Pants (252445, -0.06 DPS) [crafted]; Totemic Leather Pants (252446, -0.16 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.3 attack_power points (0.55 DPS) | yes | Veteran's Boots (250503, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.08 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.6 attack_power points (0.44 DPS) | yes | Signet of the Zhevra (285330, -0.28 DPS) [world]; The 1 Ring (8350, -0.33 DPS) [world]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.33 DPS) | yes | The 1 Ring (8350, -0.22 DPS) [world]; Signet of the Zhevra (285330, -0.23 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.25 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (14.77 DPS) | yes | Duskbringer (2205, -2.65 DPS) [dungeon]; Monstrous Cleaver (279864, -3.00 DPS) [quest]; Smite's Mighty Hammer (7230, -3.34 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 80.4. Weights run: 1.3s. Verify run: 1.0s. 404 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.640 ± 0.017, crit=0.903 ± 0.024 per rating point (14 rating = 1%, 12.646 per %), hit=1.567 ± 0.066 per rating point (10 rating = 1%, 15.673 per %), melee_haste=8.077 ± 0.510

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.36 DPS) | yes | Veteran's Chain Helm (250498, -0.10 DPS) [crafted]; Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Barbaric Iron Helm (7915, -0.12 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.73 DPS) | yes | Kaleidoscope Chain (13084, -0.18 DPS) [world_drop]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Sentinel's Medallion (19541, -0.46 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.0 attack_power points (0.89 DPS) | yes | Barbaric Iron Shoulders (7913, -0.06 DPS) [crafted]; Golden Scale Shoulders (3841, -0.16 DPS) [crafted]; Mail Combat Spaulders (6404, -0.16 DPS) [world_drop] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 10.6 attack_power points (0.55 DPS) | yes | Hawkeye's Cloak (14593, -0.00 DPS) [world_drop]; Wolfmaster Cape (6314, -0.03 DPS) [dungeon]; Slayer's Cape (14752, -0.13 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.57 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.31 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.84 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Bands of Serra'kis (6902, -0.21 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.22 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.15 DPS) | yes | The Frozen Clutch (23170, -0.10 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.18 DPS) [world_drop]; Bonefist Gauntlets (4465, -0.21 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (1.26 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 26.5 attack_power points (1.39 DPS) | yes | Ferine Leggings (6690, +0.00 DPS, sim-verified) [dungeon]; Golden Scale Leggings (3843, -0.23 DPS) [crafted]; Chausses of Westfall (6087, -0.23 DPS) [quest] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (80.4 DPS) | yes | Brawler's Leather Boots (252439, -0.04 DPS) [crafted]; Alacritous Treads (277234, -0.04 DPS) [quest]; Trouncing Boots (4464, -1.33 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.9 attack_power points (0.94 DPS) | yes | Ironspine's Eye (7686, -0.22 DPS) [dungeon]; Tiger Band (6749, -0.31 DPS) [quest]; Silverlaine's Family Seal (6321, -0.41 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.8 attack_power points (0.83 DPS) | yes | Ironspine's Eye (7686, -0.11 DPS) [dungeon]; Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.31 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -18.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 120.9. Weights run: 1.3s. Verify run: 1.0s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.010 ± 0.023, crit=1.429 ± 0.033 per rating point (14 rating = 1%, 20.009 per %), hit=1.894 ± 0.083 per rating point (10 rating = 1%, 18.936 per %), melee_haste=10.890 ± 0.616

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.0 attack_power points (2.60 DPS) | yes | White Bandit Mask (10008, -0.73 DPS) [crafted]; Chromite Barbute (8142, -0.82 DPS, sim-verified) [dungeon]; Hard Gold Coif (250537, -1.02 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.13 DPS) | yes | Kaleidoscope Chain (13084, -0.45 DPS) [world_drop]; Sentinel's Medallion (19540, -0.50 DPS) [rep]; Ghostshard Talisman (7731, -0.51 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | sim-verified (120.9 DPS) | yes | Forest Tracker Epaulets (2278, -0.05 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Sunburn Spaulders (274751, -1.28 DPS, sim-verified) [vendor] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 18.1 attack_power points (1.02 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.28 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.34 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 35.2 attack_power points (1.99 DPS) | yes | Kolkar Marauder Chain (6773, -0.01 DPS) [quest]; Avenger's Armor (1488, -0.29 DPS) [dungeon]; Jouster's Chestplate (8157, -0.29 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Ravager's Armguards (14770, -0.11 DPS) [world_drop]; Pugilist Bracers (4438, -0.23 DPS) [dungeon]; Yorgen Bracers (13012, -0.28 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 40.0 attack_power points (2.26 DPS) | yes | Scarlet Gauntlets (10331, -0.44 DPS, sim-verified) [dungeon]; Gauntlets of Divinity (7724, -0.45 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.45 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.70 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.00 DPS) [rep]; Highlander's Chain Girdle (20089, -0.11 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.37 DPS) | yes | Firemane Leggings (13129, -0.23 DPS) [world_drop]; Symbolic Legplates (14829, -0.34 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 33.1 attack_power points (1.87 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Blackforge Greaves (6423, -0.28 DPS) [dungeon]; Obsidian Greaves (13068, -0.40 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.1 attack_power points (1.36 DPS) | yes | Assault Band (13095, -0.23 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Ironspine's Eye (7686, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -18.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 159.7. Weights run: 1.4s. Verify run: 0.9s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.351 ± 0.031, crit=1.901 ± 0.044 per rating point (14 rating = 1%, 26.613 per %), hit=2.390 ± 0.103 per rating point (10 rating = 1%, 23.901 per %), melee_haste=14.331 ± 1.115

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 79.1 attack_power points (4.58 DPS) | yes | Embrace of the Lycan (9479, -1.71 DPS) [dungeon]; Ornate Mithril Helm (7937, -1.77 DPS) [crafted]; Raging Berserker's Helm (7719, -1.78 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.6 attack_power points (1.65 DPS) | yes | Sentinel's Medallion (19539, -0.72 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.80 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 50.8 attack_power points (2.94 DPS) | yes | Officer's Pauldrons (250576, -0.66 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.79 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.84 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.1 attack_power points (1.86 DPS) | yes | Dark Hooded Cape (5257, -0.57 DPS) [world]; Sergeant Major's Cape (16336, -0.63 DPS) [pvp]; Dark Phantom Cape (13122, -0.69 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 59.6 attack_power points (3.45 DPS) | yes | Mixologist's Tunic (12793, -0.30 DPS) [dungeon]; Warforged Chestplate (11195, -0.39 DPS) [quest]; Warbear Harness (15064, -0.64 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 32.3 attack_power points (1.87 DPS) | yes | Runed Golem Shackles (12550, -0.09 DPS) [dungeon]; Prowler's Leather Bracers (252539, -0.18 DPS) [crafted]; Deepfury Bracers (13120, -0.19 DPS) [world_drop] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 61.9 attack_power points (3.58 DPS) | yes | Raider Gloves (272100, -0.59 DPS) [vendor]; Officer's Gloves (250551, -0.84 DPS) [crafted]; Gloves of Holy Might (867, -0.89 DPS) [world_drop] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 50.8 attack_power points (2.94 DPS) | yes | Girdle of Beastial Fury (11686, -0.19 DPS) [dungeon]; Prowler's Leather Waistguard (252473, -0.22 DPS) [crafted]; Highlander's Plate Girdle (20124, -3.01 DPS, sim-verified) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 66.5 attack_power points (3.85 DPS) | yes | Centurion Legplates (10740, -0.77 DPS) [quest]; Stormshroud Pants (15057, -0.77 DPS) [crafted]; Gryphon Rider's Leggings (9652, -0.84 DPS, sim-verified) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 48.4 attack_power points (2.80 DPS) | yes | Prowler's Leather Boots (252468, -0.28 DPS) [crafted]; Skulker's Leather Boots (252469, -0.38 DPS) [crafted]; Officer's Sabatons (250561, -0.44 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.9 attack_power points (2.54 DPS) | yes | Thunderbrow Ring (13097, -1.29 DPS) [world_drop]; Ironspine's Eye (7686, -1.33 DPS) [dungeon]; Mark of Kern (2262, -1.38 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 34.2 attack_power points (1.98 DPS) | yes | Thunderbrow Ring (13097, -0.72 DPS) [world_drop]; Ironspine's Eye (7686, -0.76 DPS) [dungeon]; Mark of Kern (2262, -0.82 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (159.7 DPS) | yes | - |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (159.7 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (159.7 DPS) | yes | Warmonger (13052, -0.05 DPS) [world_drop]; Taran Icebreaker (2915, -0.88 DPS) [world_drop]; Blight (7959, -4.49 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Sanctified Orb; main_hand: Thorium Greatmace

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 252.5. Weights run: 1.3s. Verify run: 0.9s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.934 ± 0.046, crit=2.736 ± 0.066 per rating point (14 rating = 1%, 38.297 per %), hit=3.144 ± 0.161 per rating point (10 rating = 1%, 31.436 per %), melee_haste=17.545 ± 1.934

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 179.1 attack_power points (11.11 DPS) | yes | Eye of Rend (12587, -3.30 DPS, sim-verified) [dungeon]; Mask of the Unforgiven (13404, -4.83 DPS) [dungeon]; Knight-Lieutenant's Plate Helm (220804, -5.01 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 64.3 attack_power points (3.99 DPS) | yes | Medallion of the Dawn (22659, -0.12 DPS) [quest]; Pendant of Celerity (22340, -0.24 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.34 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 90.5 attack_power points (5.61 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.23 DPS) [dungeon]; Highlander's Plate Spaulders (20057, -1.12 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.4 attack_power points (3.69 DPS) | yes | Cape of the Black Baron (13340, -0.65 DPS) [dungeon]; Windshear Cape (20691, -0.80 DPS) [world]; Cloak of the Honor Guard (20073, -0.98 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.27 DPS) [crafted]; Savage Gladiator Chain (11726, -1.61 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.88 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.55 DPS) [rep]; Berserker Bracers (19578, -0.77 DPS) [rep]; Bracers of Undead Slaying (23090, -2.17 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, -0.02 DPS) [quest]; Raider Gloves (272099, -0.26 DPS) [vendor]; Razor Gauntlets (18326, -4.51 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 97.7 attack_power points (6.06 DPS) | yes | Ferocity of the Timbermaw (227805, -0.47 DPS) [vendor]; Highlander's Plate Girdle (20041, -1.36 DPS) [rep]; Belt of Preserved Heads (20216, -2.33 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -0.51 DPS) [vendor]; Titanic Leggings (22385, -0.53 DPS) [crafted]; Cloudkeeper Legplates (14554, -3.17 DPS, sim-verified) [world_drop] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (252.5 DPS) | yes | Knight-Lieutenant's Lamellar Sabatons (227146, -0.24 DPS) [vendor]; Drudge Boots (21532, -0.44 DPS) [quest]; Windreaver Greaves (13967, -3.75 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.57 DPS) [dungeon]; Cutthroat's Signet (272408, -1.69 DPS) [vendor]; Naglering (11669, -3.45 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.36 DPS) [dungeon]; Cutthroat's Signet (272408, -0.48 DPS) [vendor]; Naglering (11669, -2.03 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+8.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -3.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, -1.83 DPS, sim-verified) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 514.4. Weights run: 1.5s. Verify run: 1.1s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.747 ± 0.041, crit=2.468 ± 0.058 per rating point (14 rating = 1%, 34.553 per %), hit=3.527 ± 0.235 per rating point (10 rating = 1%, 35.275 per %), melee_haste=23.584 ± 2.419

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 179.3 attack_power points (22.43 DPS) | yes | Helm of the Executioner (22411, -9.75 DPS) [dungeon]; Knight-Lieutenant's Plate Helm (220804, -10.11 DPS) [vendor]; Mask of the Unforgiven (13404, -10.12 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Beads of Ogre Might (22150, -0.16 DPS) [quest]; Medallion of the Dawn (22659, -0.25 DPS) [quest]; Pendant of Celerity (22340, -5.23 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Wyrmhide Spaulders (12082, -1.40 DPS) [quest]; Truestrike Shoulders (12927, -6.10 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 63.3 attack_power points (7.92 DPS) | yes | Cape of the Black Baron (13340, -2.14 DPS) [dungeon]; Windshear Cape (20691, -2.44 DPS) [world]; Cloak of the Honor Guard (20073, -2.57 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.52 DPS) [crafted]; Savage Gladiator Chain (11726, -3.39 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -18.64 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.56 DPS) [rep]; Berserker Bracers (19578, -1.76 DPS) [rep]; Bracers of Undead Slaying (23090, -3.69 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Raider Gauntlets (272095, -0.96 DPS) [vendor]; Marshal's Lamellar Gauntlets (231650, -1.79 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 94.0 attack_power points (11.76 DPS) | yes | Ferocity of the Timbermaw (227805, -0.88 DPS) [vendor]; Marksman's Girdle (22232, -2.75 DPS) [dungeon]; Belt of Preserved Heads (20216, -5.32 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -0.70 DPS) [vendor]; Sentinel's Plate Legguards (237825, -0.92 DPS) [vendor]; Cloudkeeper Legplates (14554, -2.74 DPS, sim-verified) [world_drop] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | 70.8 attack_power points (8.86 DPS) | yes | Knight-Lieutenant's Lamellar Sabatons (227146, -1.23 DPS) [vendor]; Drudge Boots (21532, -1.56 DPS) [quest]; Windreaver Greaves (13967, -3.63 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.05 DPS) [dungeon]; Cutthroat's Signet (272408, -3.27 DPS) [vendor]; Naglering (11669, -6.06 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.66 DPS) [dungeon]; Cutthroat's Signet (272408, -0.87 DPS) [vendor]; Naglering (11669, -3.34 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Maelstrom (19289, -11.23 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Maelstrom (19289, -7.33 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -3.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 35.4. Weights run: 1.4s. Verify run: 0.9s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.659 ± 0.018, crit=0.931 ± 0.025 per rating point (14 rating = 1%, 13.030 per %), hit=1.308 ± 0.069 per rating point (10 rating = 1%, 13.082 per %), melee_haste=8.619 ± 0.501

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.82 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.55 DPS) [crafted]; Brawler's Leather Hood (252504, -0.61 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.0 attack_power points (0.16 DPS) | yes | Erudite's Amulet (277204, -0.05 DPS) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) | Blacksmithing [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.11 DPS) [dungeon]; Rough Bronze Shoulders (3480, -0.71 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.33 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Catacomb Cloak (279899, -0.08 DPS) [quest]; Subterranean Cape (14149, -0.08 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.82 DPS) | yes | Veteran's Chain Shirt (250488, -0.14 DPS) [crafted]; Defender's Leather Armor (252434, -0.17 DPS) [crafted]; Brawler's Leather Armor (252490, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.41 DPS) | yes | Raptorcrest Bracers (270010, -0.16 DPS) [quest]; Bristlebark Bindings (14569, -0.17 DPS) [world_drop]; Runed Copper Bracers (2854, -0.25 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.66 DPS) | yes | Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Fletcher's Gloves (7348, -0.12 DPS) [crafted]; Blackened Defias Gloves (10401, -0.16 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.74 DPS) | yes | Cobrahn's Grasp (6460, -0.08 DPS) [dungeon]; Ruffian Belt (5975, -0.25 DPS) [world]; Brawler's Leather Belt (252428, -0.30 DPS) [crafted] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 21.3 attack_power points (0.88 DPS) | yes | Defender's Leather Pants (252445, -0.03 DPS) [crafted]; Totemic Leather Pants (252446, -0.14 DPS) [crafted]; Hulking Leggings (14748, -0.14 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 13.3 attack_power points (0.55 DPS) | yes | Veteran's Boots (250503, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.08 DPS) [world_drop]; Defender's Leather Boots (252441, -0.14 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.6 attack_power points (0.44 DPS) | yes | Signet of the Zhevra (285330, -0.28 DPS) [world]; The 1 Ring (8350, -0.33 DPS) [world]; Demon Band (12054, -0.45 DPS, sim-verified) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Signet of the Zhevra (285330, -0.08 DPS) [world]; The 1 Ring (8350, -0.14 DPS) [world]; Demon Band (12054, -0.42 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (12.60 DPS) | yes | Smite's Mighty Hammer (7230, -0.45 DPS, sim-verified) [dungeon]; Duskbringer (2205, -0.48 DPS) [dungeon]; Monstrous Cleaver (279864, -0.84 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 87.1. Weights run: 1.3s. Verify run: 1.1s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=0.640 ± 0.017, crit=0.903 ± 0.024 per rating point (14 rating = 1%, 12.646 per %), hit=1.567 ± 0.066 per rating point (10 rating = 1%, 15.673 per %), melee_haste=8.077 ± 0.510

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.36 DPS) | yes | Veteran's Chain Helm (250498, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.10 DPS) [crafted]; Barbaric Iron Helm (7915, -0.12 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.73 DPS) | yes | River Pride Choker (13087, -0.31 DPS) [world_drop]; Scout's Medallion (19537, -0.46 DPS) [rep]; Kaleidoscope Chain (13084, -0.51 DPS, sim-verified) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 17.0 attack_power points (0.89 DPS) | yes | Barbaric Iron Shoulders (7913, -0.06 DPS) [crafted]; Golden Scale Shoulders (3841, -0.16 DPS) [crafted]; Mail Combat Spaulders (6404, -0.16 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.5 attack_power points (0.55 DPS) | yes | Wolfmaster Cape (6314, -0.02 DPS) [dungeon]; Wildhunter Cloak (16658, -0.02 DPS) [quest]; Slayer's Cape (14752, -0.13 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.57 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.31 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.84 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Bands of Serra'kis (6902, -0.21 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.22 DPS) [world_drop] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.15 DPS) | yes | Warsong Gauntlets (16978, -0.10 DPS) [quest]; The Frozen Clutch (23170, -0.10 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.18 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (1.26 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 26.5 attack_power points (1.39 DPS) | yes | Ferine Leggings (6690, -0.02 DPS) [dungeon]; Golden Scale Leggings (3843, -0.23 DPS) [crafted]; Slayer's Pants (14757, -0.23 DPS) [world_drop] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (87.1 DPS) | yes | Brawler's Leather Boots (252439, -0.04 DPS) [crafted]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Trouncing Boots (4464, -1.35 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.9 attack_power points (0.94 DPS) | yes | Ironspine's Eye (7686, -0.22 DPS) [dungeon]; Tiger Band (6749, -0.31 DPS) [quest]; Band of the Fist (17694, -0.39 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.8 attack_power points (0.83 DPS) | yes | Tiger Band (6749, -0.20 DPS) [quest]; Band of the Fist (17694, -0.28 DPS) [quest]; Ironspine's Eye (7686, -0.46 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -21.34 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Veteran's Silvered Chain Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 132.9. Weights run: 1.3s. Verify run: 1.0s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.010 ± 0.023, crit=1.429 ± 0.033 per rating point (14 rating = 1%, 20.009 per %), hit=1.894 ± 0.083 per rating point (10 rating = 1%, 18.936 per %), melee_haste=10.890 ± 0.616

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.0 attack_power points (2.60 DPS) | yes | White Bandit Mask (10008, -0.73 DPS) [crafted]; Hard Gold Coif (250537, -1.02 DPS) [crafted]; Chromite Barbute (8142, -1.28 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.13 DPS) | yes | Ethereal Talisman (4430, -0.34 DPS) [quest]; Ghostshard Talisman (7731, -0.34 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Tracker Epaulets (2278, -0.05 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Sunburn Spaulders (274751, -1.33 DPS, sim-verified) [vendor] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Cloak (14593, -0.28 DPS) [world_drop]; Wildhunter Cloak (16658, -0.46 DPS) [quest]; Dark Hooded Cape (5257, -1.62 DPS, sim-verified) [world] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 35.2 attack_power points (1.99 DPS) | yes | Kolkar Marauder Chain (6773, -0.01 DPS) [quest]; Avenger's Armor (1488, -0.29 DPS) [dungeon]; Jouster's Chestplate (8157, -0.29 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.63 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 40.0 attack_power points (2.26 DPS) | yes | Scarlet Gauntlets (10331, -0.33 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.45 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.45 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.70 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.00 DPS) [rep]; Defiler's Chain Girdle (20153, -0.11 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.37 DPS) | yes | Symbolic Legplates (14829, -0.34 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted]; Firemane Leggings (13129, -0.68 DPS, sim-verified) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 33.1 attack_power points (1.87 DPS) | yes | Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Blackforge Greaves (6423, -0.28 DPS) [dungeon]; Obsidian Greaves (13068, -0.40 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.1 attack_power points (1.36 DPS) | yes | Assault Band (13095, -0.23 DPS) [world_drop]; Thunderbrow Ring (13097, -0.29 DPS) [world_drop]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.13 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Ironspine's Eye (7686, -0.16 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -23.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 164.9. Weights run: 1.4s. Verify run: 0.9s. 702 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.351 ± 0.031, crit=1.901 ± 0.044 per rating point (14 rating = 1%, 26.613 per %), hit=2.390 ± 0.103 per rating point (10 rating = 1%, 23.901 per %), melee_haste=14.331 ± 1.115

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 55.2 attack_power points (3.20 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -0.33 DPS) [dungeon]; Ornate Mithril Helm (7937, -0.38 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.6 attack_power points (1.65 DPS) | yes | Woven Ivy Necklace (19159, -0.19 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.50 DPS) [quest]; Ethereal Talisman (4430, -0.70 DPS) [quest] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 39.4 attack_power points (2.28 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.13 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 32.1 attack_power points (1.86 DPS) | yes | Dark Hooded Cape (5257, -0.57 DPS) [world]; First Sergeant's Cloak (16340, -0.63 DPS) [pvp]; Dark Phantom Cape (13122, -0.69 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 59.6 attack_power points (3.45 DPS) | yes | Mixologist's Tunic (12793, -0.30 DPS) [dungeon]; Warforged Chestplate (11195, -0.39 DPS) [quest]; Warbear Harness (15064, -0.64 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 32.3 attack_power points (1.87 DPS) | yes | Runed Golem Shackles (12550, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 61.9 attack_power points (3.58 DPS) | yes | Raider Gloves (272100, -0.59 DPS) [vendor]; Officer's Gloves (250551, -0.84 DPS) [crafted]; Gloves of Holy Might (867, -0.89 DPS) [world_drop] |
| waist | Defiler's Plate Girdle (20205) | The Defilers [rep] | 48.6 attack_power points (2.81 DPS) | yes | Girdle of Beastial Fury (11686, -0.06 DPS) [dungeon]; Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.12 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 66.5 attack_power points (3.85 DPS) | yes | Serpentskin Leggings (8262, -0.94 DPS) [world_drop]; Golem Shard Leggings (13074, -1.05 DPS) [world_drop]; Stormshroud Pants (15057, -1.42 DPS, sim-verified) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 48.4 attack_power points (2.80 DPS) | yes | Prowler's Leather Boots (252468, -0.28 DPS) [crafted]; Skulker's Leather Boots (252469, -0.38 DPS) [crafted]; Officer's Sabatons (250561, -0.44 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.9 attack_power points (2.54 DPS) | yes | White Bone Band (11862, -1.15 DPS) [quest]; Thunderbrow Ring (13097, -1.29 DPS) [world_drop]; Ironspine's Eye (7686, -1.33 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 34.2 attack_power points (1.98 DPS) | yes | White Bone Band (11862, -0.59 DPS) [quest]; Thunderbrow Ring (13097, -0.72 DPS) [world_drop]; Ironspine's Eye (7686, -0.76 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (164.9 DPS) | yes | Molten Heart of the Mountain (249470, -3.14 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (164.9 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (164.9 DPS) | yes | Warmonger (13052, -0.05 DPS) [world_drop]; Taran Icebreaker (2915, -0.88 DPS) [world_drop]; Blight (7959, -3.65 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Defiler's Plate Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 263.7. Weights run: 1.3s. Verify run: 1.0s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.934 ± 0.046, crit=2.736 ± 0.066 per rating point (14 rating = 1%, 38.297 per %), hit=3.144 ± 0.161 per rating point (10 rating = 1%, 31.436 per %), melee_haste=17.545 ± 1.934

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 179.1 attack_power points (11.11 DPS) | yes | Eye of Rend (12587, -2.83 DPS, sim-verified) [dungeon]; Mask of the Unforgiven (13404, -4.83 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -5.01 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 64.3 attack_power points (3.99 DPS) | yes | Medallion of the Dawn (22659, -0.12 DPS) [quest]; Pendant of Celerity (22340, -0.24 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.34 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 90.5 attack_power points (5.61 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.23 DPS) [dungeon]; Defiler's Plate Spaulders (20212, -1.12 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.4 attack_power points (3.69 DPS) | yes | Cape of the Black Baron (13340, -0.65 DPS) [dungeon]; Windshear Cape (20691, -0.80 DPS) [world]; Deathguard's Cloak (20068, -0.98 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.27 DPS) [crafted]; Savage Gladiator Chain (11726, -1.61 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.55 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.55 DPS) [rep]; Berserker Bracers (19578, -0.77 DPS) [rep]; Bracers of Undead Slaying (23090, -2.13 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, -0.02 DPS) [quest]; Raider Gloves (272099, -0.26 DPS) [vendor]; Razor Gauntlets (18326, -5.20 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 97.7 attack_power points (6.06 DPS) | yes | Ferocity of the Timbermaw (227805, -0.47 DPS) [vendor]; Defiler's Plate Girdle (20204, -1.36 DPS) [rep]; Belt of Preserved Heads (20216, -2.93 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -0.51 DPS) [vendor]; Titanic Leggings (22385, -0.53 DPS) [crafted]; Cloudkeeper Legplates (14554, -3.53 DPS, sim-verified) [world_drop] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (263.7 DPS) | yes | Drudge Boots (21532, -0.44 DPS) [quest]; Battlechaser's Greaves (12555, -0.79 DPS) [dungeon]; Windreaver Greaves (13967, -4.70 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.57 DPS) [dungeon]; Cutthroat's Signet (272408, -1.69 DPS) [vendor]; Naglering (11669, -3.80 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.36 DPS) [dungeon]; Cutthroat's Signet (272408, -0.48 DPS) [vendor]; Naglering (11669, -2.19 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Talisman of Ascendance (22678, -2.52 DPS, sim-verified) [quest] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -3.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Law (272435, +0.00 DPS) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Draconic Infused Emblem; trinket2: Burst of Knowledge; main_hand: The Unstoppable Force

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 522.1. Weights run: 1.5s. Verify run: 1.1s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.005, agility=1.747 ± 0.041, crit=2.468 ± 0.058 per rating point (14 rating = 1%, 34.553 per %), hit=3.527 ± 0.235 per rating point (10 rating = 1%, 35.275 per %), melee_haste=23.584 ± 2.419

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 179.3 attack_power points (22.43 DPS) | yes | Mask of the Unforgiven (13404, -9.31 DPS, sim-verified) [dungeon]; Helm of the Executioner (22411, -9.75 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -10.11 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 61.5 attack_power points (7.69 DPS) | yes | Mark of Fordring (15411, +0.00 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -0.28 DPS) [quest]; Medallion of the Dawn (22659, -0.37 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (522.1 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Wyrmhide Spaulders (12082, -1.40 DPS) [quest]; Truestrike Shoulders (12927, -5.29 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 63.3 attack_power points (7.92 DPS) | yes | Cape of the Black Baron (13340, -2.14 DPS) [dungeon]; Windshear Cape (20691, -2.44 DPS) [world]; Deathguard's Cloak (20068, -2.57 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -0.52 DPS) [crafted]; Savage Gladiator Chain (11726, -3.39 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -17.30 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.56 DPS) [rep]; Berserker Bracers (19578, -1.76 DPS) [rep]; Bracers of Undead Slaying (23090, -3.98 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razor Gauntlets (18326, +0.00 DPS) [dungeon]; Raider Gauntlets (272095, -0.96 DPS) [vendor]; Raider Gloves (272099, -1.87 DPS) [vendor] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 94.0 attack_power points (11.76 DPS) | yes | Ferocity of the Timbermaw (227805, -0.88 DPS) [vendor]; Marksman's Girdle (22232, -2.75 DPS) [dungeon]; Belt of Preserved Heads (20216, -5.37 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -0.70 DPS) [vendor]; Sentinel's Plate Legguards (237825, -0.92 DPS) [vendor]; Cloudkeeper Legplates (14554, -3.19 DPS, sim-verified) [world_drop] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | 70.8 attack_power points (8.86 DPS) | yes | Drudge Boots (21532, -1.56 DPS) [quest]; Fine Dawn Treaders (227815, -2.04 DPS) [vendor]; Windreaver Greaves (13967, -3.85 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.05 DPS) [dungeon]; Cutthroat's Signet (272408, -3.27 DPS) [vendor]; Naglering (11669, -6.85 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.66 DPS) [dungeon]; Cutthroat's Signet (272408, -0.87 DPS) [vendor]; Naglering (11669, -3.92 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -5.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Draconic Infused Emblem; main_hand: The Unstoppable Force

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

