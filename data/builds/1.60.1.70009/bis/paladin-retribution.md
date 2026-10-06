# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 35.5. Weights run: 2.3s. Verify run: 1.1s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.110 ± 0.010, crit=0.165 ± 0.003 per rating point (14 rating = 1%, 2.308 per %), hit=0.169 ± 0.002 per rating point (10 rating = 1%, 1.692 per %), melee_haste=1.212 ± 0.069

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.7 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.73 DPS) | yes | Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.37 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.59 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.66 DPS) | yes | Cobrahn's Grasp (6460, -0.13 DPS) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.81 DPS) | yes | Veteran's Chain Leggings (250493, -0.13 DPS) [crafted]; Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Signet of the Zhevra (285330, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Ring of the Moon (12052, -0.22 DPS) [world_drop]; The 1 Ring (8350, -0.23 DPS, sim-verified) [world]; Signet of the Zhevra (285330, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (13.14 DPS) | yes | Smite's Mighty Hammer (7230, -2.33 DPS, sim-verified) [dungeon]; Duskbringer (2205, -2.35 DPS) [dungeon]; Monstrous Cleaver (279864, -2.67 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-05025331200000000)

Set DPS (verified): 72.6. Weights run: 2.3s. Verify run: 1.3s. 404 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.177 ± 0.016, crit=0.227 ± 0.005 per rating point (14 rating = 1%, 3.176 per %), hit=0.229 ± 0.003 per rating point (10 rating = 1%, 2.293 per %), melee_haste=1.741 ± 0.098

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.11 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.60 DPS) | yes | Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop]; Sentinel's Medallion (19541, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.60 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.04 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.09 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.43 DPS) | yes | Sergeant Major's Cape (16315, -0.06 DPS) [pvp]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.28 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.26 DPS) [crafted]; Hard Gold Cuirass (250533, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.68 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Patterned Bronze Bracers (2868, -0.26 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.94 DPS) | yes | The Frozen Clutch (23170, -0.09 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop] |
| waist | Girdle of Golem Strength (9405) (or Highlander's Plate Girdle (20126)) | World drop [world_drop] | 24.0 attack_power points (1.03 DPS) | yes | Highlander's Plate Girdle (20126, +0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.11 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Chausses of Westfall (6087, -0.17 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.33 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (72.6 DPS) | yes | Glimmering Mail Greaves (4073, -0.09 DPS) [world_drop]; Slayer's Slippers (14756, -0.09 DPS) [world_drop]; Trouncing Boots (4464, -1.08 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (0.71 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 13.1 attack_power points (0.56 DPS) | yes | Tiger Band (6749, -0.05 DPS) [quest]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.46 DPS) [dungeon]; Viscous Hammer (13045, -16.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-05025331211330200)

Set DPS (verified): 103.7. Weights run: 2.3s. Verify run: 1.3s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.334 ± 0.028, crit=0.459 ± 0.008 per rating point (14 rating = 1%, 6.429 per %), hit=0.318 ± 0.004 per rating point (10 rating = 1%, 3.183 per %), melee_haste=2.750 ± 0.171

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 32.4 attack_power points (1.50 DPS) | yes | Icemetal Barbute (10763, -0.20 DPS) [dungeon]; Hard Gold Coif (250537, -0.20 DPS) [crafted]; Chromite Barbute (8142, -0.67 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.92 DPS) | yes | Kaleidoscope Chain (13084, -0.49 DPS) [world_drop]; Ghostshard Talisman (7731, -0.50 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -0.55 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.02 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.18 DPS) [dungeon]; Chromite Pauldrons (8144, -1.26 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.0 attack_power points (0.65 DPS) | yes | Wolfmaster Cape (6314, -0.19 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop]; Dark Hooded Cape (5257, -0.43 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.3 attack_power points (1.40 DPS) | yes | Avenger's Armor (1488, -0.02 DPS) [dungeon]; Jouster's Chestplate (8157, -0.02 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.92 DPS) | yes | Pugilist Bracers (4438, -0.18 DPS) [dungeon]; Ravager's Armguards (14770, -0.22 DPS) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.48 DPS) | yes | Truesilver Gauntlets (7938, -0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -0.22 DPS) [dungeon]; Gloves of Holy Might (867, -0.26 DPS) [world_drop] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.39 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.00 DPS) [rep]; Highlander's Chain Girdle (20090, -0.28 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.94 DPS) | yes | Firemane Leggings (13129, -0.18 DPS) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.46 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 28.3 attack_power points (1.31 DPS) | yes | Prowler's Leather Shoes (252465, -0.18 DPS) [crafted]; Obsidian Greaves (13068, -0.29 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.29 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.92 DPS) | yes | Protector's Band (19515, -0.06 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.92 DPS) | yes | Protector's Band (19515, -0.06 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (103.7 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -6.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 54000000000000000-0000000000000000-05025331211330210)

Set DPS (verified): 121.7. Weights run: 2.5s. Verify run: 1.5s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=0.416 ± 0.035, crit=0.608 ± 0.012 per rating point (14 rating = 1%, 8.507 per %), hit=0.405 ± 0.005 per rating point (10 rating = 1%, 4.055 per %), melee_haste=3.743 ± 0.452

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | sim-verified (121.7 DPS) | yes | Sunscale Helmet (14849, -0.16 DPS) [world_drop]; Raging Berserker's Helm (7719, -0.19 DPS) [dungeon]; Embrace of the Lycan (9479, -1.89 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.93 DPS) | yes | Skibi's Pendant (13089, -0.17 DPS) [world_drop]; Ghostshard Talisman (7731, -0.28 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.45 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 32.7 attack_power points (1.53 DPS) | yes | Officer's Pauldrons (250576, -0.04 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.14 DPS) [world_drop]; Earthslag Shoulders (11632, -0.19 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 19.8 attack_power points (0.92 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Cape (16336, -0.19 DPS) [pvp]; Dark Hooded Cape (5257, -0.32 DPS) [world] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 52.8 attack_power points (2.46 DPS) | yes | Mixologist's Tunic (12793, -0.40 DPS) [dungeon]; Valorous Chestguard (8274, -0.41 DPS) [world_drop]; Coldmetal Guard (274758, -0.51 DPS) [vendor] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 30.8 attack_power points (1.44 DPS) | yes | Bracers of the Stone Princess (17714, -0.13 DPS) [dungeon]; Arena Bands (18711, -0.13 DPS) [world]; Officer's Wristguards (250581, -0.19 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 52.6 attack_power points (2.45 DPS) | yes | Truesilver Gauntlets (7938, -0.81 DPS) [crafted]; Maddening Gauntlets (11867, -0.82 DPS) [quest]; Officer's Gloves (250551, -0.96 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 47.6 attack_power points (2.22 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.37 DPS) [dungeon]; Belt of the Gladiator (13134, -0.37 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.55 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 48.4 attack_power points (2.26 DPS) | yes | Scarlet Leggings (10330, -0.10 DPS) [dungeon]; Silvershell Leggings (10633, -0.21 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.21 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 36.2 attack_power points (1.69 DPS) | yes | Prowler's Leather Boots (252468, -0.14 DPS) [crafted]; Officer's Sabatons (250561, -0.18 DPS) [crafted]; Officer's Boots (250546, -0.22 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.7 attack_power points (1.20 DPS) | yes | Mark of Kern (2262, -0.27 DPS) [dungeon]; Assault Band (13095, -0.27 DPS) [world_drop]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.1 attack_power points (1.12 DPS) | yes | Assault Band (13095, -0.19 DPS) [world_drop]; Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Mark of Kern (2262, -0.70 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Drakefang Butcher (12463, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -3.79 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Plate Helm; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Sanctified Orb; main_hand: Nightblade

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 54000000000000000-3232000000000000-05025331211330210)

Set DPS (verified): 194.6. Weights run: 2.5s. Verify run: 1.3s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=0.550 ± 0.046, crit=0.805 ± 0.018 per rating point (14 rating = 1%, 11.273 per %), hit=0.489 ± 0.006 per rating point (10 rating = 1%, 4.887 per %), melee_haste=5.292 ± 0.781

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 71.9 attack_power points (3.43 DPS) | yes | Field Marshal's Lamellar Headguard (231648, -0.58 DPS) [vendor]; Field Marshal's Lamellar Faceguard (16474, -0.69 DPS) [vendor]; Embrace of the Lycan (9479, -3.00 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 35.3 attack_power points (1.68 DPS) | yes | Amulet of the Darkmoon (19491, -0.13 DPS) [quest]; Imperial Jewel (11933, -0.16 DPS) [dungeon]; Will of the Martyr (17044, -0.25 DPS) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 49.0 attack_power points (2.33 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS, sim-verified) [rep]; Field Marshal's Lamellar Shoulders (231651, -0.13 DPS) [vendor]; Black Dragonscale Shoulders (15051, -0.43 DPS) [crafted] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 37.4 attack_power points (1.78 DPS) | yes | Cloak of the Honor Guard (20073, -0.03 DPS) [rep]; Howler's Furs (272414, -0.21 DPS) [vendor]; Shadewood Cloak (18328, -0.42 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (194.6 DPS) | yes | Obsidian Mail Tunic (22191, -0.31 DPS) [crafted]; Cadaverous Armor (14637, -0.56 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.43 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-verified (194.6 DPS) | yes | Windtalker's Wristguards (19582, -0.39 DPS) [rep]; Gordok Bracers of Power (18533, -0.42 DPS) [dungeon]; Bracers of Undead Slaying (23090, -4.20 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (194.6 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.45 DPS) [vendor]; Timbermaw Brawlers (19049, -0.55 DPS) [crafted]; Razor Gauntlets (18326, -4.42 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 70.7 attack_power points (3.37 DPS) | yes | Ferocity of the Timbermaw (227805, -0.20 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.32 DPS) [vendor]; Heavy Obsidian Belt (22197, -0.75 DPS) [crafted] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (194.6 DPS) | yes | Sentinel's Plate Legguards (237825, -0.01 DPS) [vendor]; Cloudkeeper Legplates (14554, -1.02 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.4 attack_power points (2.02 DPS) | yes | Clutchlord's Stompers (275627, -0.03 DPS) [crafted]; Pads of the Dread Wolf (13210, -0.11 DPS) [dungeon]; Marshal's Lamellar Sabatons (231652, -0.13 DPS) [vendor] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (194.6 DPS) | yes | Band of the Ogre King (18522, -0.08 DPS) [dungeon]; Myrmidon's Signet (2246, -0.31 DPS) [world_drop]; Naglering (11669, -3.17 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (194.6 DPS) | yes | Band of the Ogre King (18522, -0.06 DPS) [dungeon]; Myrmidon's Signet (2246, -0.30 DPS) [world_drop]; Naglering (11669, -3.53 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (194.6 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.74 DPS) [crafted] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (194.6 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Second Wind (11819, -1.56 DPS, sim-verified) [dungeon] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (194.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -5.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Highlander's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Scalegut Treaders; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Draconic Infused Emblem; main_hand: Blackblade of Shahram

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 33.1. Weights run: 2.3s. Verify run: 1.0s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.110 ± 0.010, crit=0.165 ± 0.003 per rating point (14 rating = 1%, 2.308 per %), hit=0.169 ± 0.002 per rating point (10 rating = 1%, 1.692 per %), melee_haste=1.212 ± 0.069

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.7 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.07 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.73 DPS) | yes | Veteran's Chain Shirt (250488, -0.20 DPS) [crafted]; Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.37 DPS) | yes | Raptorcrest Bracers (270010, -0.17 DPS, sim-verified) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.22 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.59 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.66 DPS) | yes | Cobrahn's Grasp (6460, -0.13 DPS) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.5 attack_power points (0.68 DPS) | yes | Defender's Leather Pants (252445, -0.00 DPS) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.5 attack_power points (0.39 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.31 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.22 DPS) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (11.22 DPS) | yes | Smite's Mighty Hammer (7230, -0.29 DPS) [dungeon]; Duskbringer (2205, -0.43 DPS) [dungeon]; Monstrous Cleaver (279864, -0.75 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-05025331200000000)

Set DPS (verified): 72.3. Weights run: 2.3s. Verify run: 1.3s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.177 ± 0.016, crit=0.227 ± 0.005 per rating point (14 rating = 1%, 3.176 per %), hit=0.229 ± 0.003 per rating point (10 rating = 1%, 2.293 per %), melee_haste=1.741 ± 0.098

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.11 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.60 DPS) | yes | Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop]; Scout's Medallion (19537, -0.54 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.60 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.04 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.43 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.28 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.26 DPS) [crafted]; Hard Gold Cuirass (250533, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.68 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.25 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.94 DPS) | yes | Warsong Gauntlets (16978, -0.09 DPS) [quest]; The Frozen Clutch (23170, -0.09 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.17 DPS) [world] |
| waist | Girdle of Golem Strength (9405) (or Defiler's Plate Girdle (20207)) | World drop [world_drop] | 24.0 attack_power points (1.03 DPS) | yes | Defiler's Plate Girdle (20207, +0.00 DPS) [rep]; Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.11 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.33 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (72.3 DPS) | yes | Glimmering Mail Greaves (4073, -0.09 DPS) [world_drop]; Slayer's Slippers (14756, -0.09 DPS) [world_drop]; Trouncing Boots (4464, -1.08 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (0.71 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 13.1 attack_power points (0.56 DPS) | yes | Tiger Band (6749, -0.05 DPS) [quest]; Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.46 DPS) [dungeon]; Viscous Hammer (13045, -16.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Girdle of Golem Strength; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-05025331211330200)

Set DPS (verified): 103.3. Weights run: 2.3s. Verify run: 1.3s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.334 ± 0.028, crit=0.459 ± 0.008 per rating point (14 rating = 1%, 6.429 per %), hit=0.318 ± 0.004 per rating point (10 rating = 1%, 3.183 per %), melee_haste=2.750 ± 0.171

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 32.4 attack_power points (1.50 DPS) | yes | Icemetal Barbute (10763, -0.20 DPS) [dungeon]; Hard Gold Coif (250537, -0.20 DPS) [crafted]; Chromite Barbute (8142, -0.74 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.92 DPS) | yes | Ethereal Talisman (4430, -0.40 DPS) [quest]; Kaleidoscope Chain (13084, -0.49 DPS) [world_drop]; Ghostshard Talisman (7731, -0.50 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.02 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.18 DPS) [dungeon]; Chromite Pauldrons (8144, -1.25 DPS, sim-verified) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.0 attack_power points (0.65 DPS) | yes | Wolfmaster Cape (6314, -0.19 DPS) [dungeon]; Wildhunter Cloak (16658, -0.19 DPS) [quest]; Dark Hooded Cape (5257, -0.71 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 30.3 attack_power points (1.40 DPS) | yes | Avenger's Armor (1488, -0.02 DPS) [dungeon]; Jouster's Chestplate (8157, -0.02 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.11 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.92 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.48 DPS) | yes | Truesilver Gauntlets (7938, -0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -0.22 DPS) [dungeon]; Gloves of Holy Might (867, -0.26 DPS) [world_drop] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.39 DPS) | yes | Boar Champion's Belt (10768, -0.00 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.18 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.94 DPS) | yes | Firemane Leggings (13129, -0.18 DPS) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.46 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 28.3 attack_power points (1.31 DPS) | yes | Prowler's Leather Shoes (252465, -0.18 DPS) [crafted]; Obsidian Greaves (13068, -0.29 DPS) [world_drop]; Skirmisher's Mail Boots (252564, -0.29 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.92 DPS) | yes | Legionnaire's Band (19512, -0.06 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.92 DPS) | yes | Legionnaire's Band (19512, -0.06 DPS) [rep]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (103.3 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -13.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 54000000000000000-0000000000000000-05025331211330210)

Set DPS (verified): 118.1. Weights run: 2.5s. Verify run: 1.3s. 702 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=0.416 ± 0.035, crit=0.608 ± 0.012 per rating point (14 rating = 1%, 8.507 per %), hit=0.405 ± 0.005 per rating point (10 rating = 1%, 4.055 per %), melee_haste=3.743 ± 0.452

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (2.31 DPS) | yes | Blood Guard's Plate Helm (220803, -0.39 DPS) [vendor]; Sunscale Helmet (14849, -0.56 DPS) [world_drop]; Raging Berserker's Helm (7719, -0.58 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.93 DPS) | yes | Woven Ivy Necklace (19159, -0.14 DPS) [quest]; Skibi's Pendant (13089, -0.17 DPS) [world_drop]; Ghostshard Talisman (7731, -0.28 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 31.9 attack_power points (1.49 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.10 DPS) [world_drop]; Earthslag Shoulders (11632, -0.16 DPS) [dungeon] |
| back | Bloodlust Cape (14801) | World drop [world_drop] | 19.8 attack_power points (0.92 DPS) | yes | Blackveil Cape (11626, +0.00 DPS, sim-verified) [dungeon]; First Sergeant's Cloak (16340, -0.19 DPS) [pvp]; Dark Hooded Cape (5257, -0.32 DPS) [world] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 52.8 attack_power points (2.46 DPS) | yes | Mixologist's Tunic (12793, -0.40 DPS) [dungeon]; Valorous Chestguard (8274, -0.41 DPS) [world_drop]; Coldmetal Guard (274758, -0.51 DPS) [vendor] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 30.8 attack_power points (1.44 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 52.6 attack_power points (2.45 DPS) | yes | Truesilver Gauntlets (7938, -0.81 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.84 DPS) [crafted]; Officer's Gloves (250551, -0.89 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 47.6 attack_power points (2.22 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.37 DPS) [dungeon]; Belt of the Gladiator (13134, -0.37 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.55 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 48.4 attack_power points (2.26 DPS) | yes | Scarlet Leggings (10330, -0.10 DPS) [dungeon]; Silvershell Leggings (10633, -0.21 DPS) [dungeon]; Elemental Rockridge Leggings (17711, -0.21 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 36.2 attack_power points (1.69 DPS) | yes | Prowler's Leather Boots (252468, -0.14 DPS) [crafted]; Officer's Sabatons (250561, -0.18 DPS) [crafted]; Officer's Boots (250546, -0.22 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 25.7 attack_power points (1.20 DPS) | yes | White Bone Band (11862, -0.08 DPS) [quest]; Mark of Kern (2262, -0.27 DPS) [dungeon]; Assault Band (13095, -0.27 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.1 attack_power points (1.12 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; White Bone Band (11862, -1.45 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (118.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (118.1 DPS) | yes | Molten Heart of the Mountain (249470, -1.20 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (118.1 DPS) | yes | Taran Icebreaker (2915, -0.71 DPS) [world_drop]; Drakefang Butcher (12463, -1.55 DPS) [dungeon]; Blight (7959, -1.79 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Bloodlust Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 54000000000000000-3232000000000000-05025331211330210)

Set DPS (verified): 187.1. Weights run: 2.5s. Verify run: 1.3s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=0.550 ± 0.046, crit=0.805 ± 0.018 per rating point (14 rating = 1%, 11.273 per %), hit=0.489 ± 0.006 per rating point (10 rating = 1%, 4.887 per %), melee_haste=5.292 ± 0.781

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 71.9 attack_power points (3.43 DPS) | yes | Warbear Helm (252485, -1.17 DPS) [crafted]; Soulforge Greathelm (226976, -1.21 DPS) [vendor]; Embrace of the Lycan (9479, -7.39 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 35.3 attack_power points (1.68 DPS) | yes | Amulet of the Darkmoon (19491, -0.13 DPS) [quest]; Imperial Jewel (11933, -0.16 DPS) [dungeon]; Will of the Martyr (17044, -0.25 DPS) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 49.0 attack_power points (2.33 DPS) | yes | Defiler's Leather Shoulders (20194, -0.43 DPS) [rep]; Darkspear Spaulders (272108, -0.54 DPS) [vendor]; Black Dragonscale Shoulders (15051, -1.59 DPS, sim-verified) [crafted] |
| back | Shroud of Domination (22337) | Blackrock Spire: Lord Valthalak [dungeon] | 37.4 attack_power points (1.78 DPS) | yes | Deathguard's Cloak (20068, -0.03 DPS) [rep]; Howler's Furs (272414, -0.21 DPS) [vendor]; Shadewood Cloak (18328, -0.42 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (187.1 DPS) | yes | Obsidian Mail Tunic (22191, -0.31 DPS) [crafted]; Cadaverous Armor (14637, -0.56 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.16 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-verified (187.1 DPS) | yes | Windtalker's Wristguards (19582, -0.39 DPS) [rep]; Gordok Bracers of Power (18533, -0.42 DPS) [dungeon]; Bracers of Undead Slaying (23090, -3.95 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (187.1 DPS) | yes | Radiant Gloves of the Dawn (227817, -0.45 DPS) [vendor]; Timbermaw Brawlers (19049, -0.55 DPS) [crafted]; Razor Gauntlets (18326, -4.19 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 70.7 attack_power points (3.37 DPS) | yes | Ferocity of the Timbermaw (227805, -0.20 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.32 DPS) [vendor]; Heavy Obsidian Belt (22197, -0.75 DPS) [crafted] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (187.1 DPS) | yes | Sentinel's Plate Legguards (237825, -0.01 DPS) [vendor]; Outrider's Plate Legguards (22651, -0.44 DPS) [rep]; Cloudkeeper Legplates (14554, -5.31 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 42.4 attack_power points (2.02 DPS) | yes | Clutchlord's Stompers (275627, -0.03 DPS) [crafted]; Pads of the Dread Wolf (13210, -0.11 DPS) [dungeon]; Goregasher Stompers (275624, -0.21 DPS) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (187.1 DPS) | yes | Band of the Ogre King (18522, -0.08 DPS) [dungeon]; Myrmidon's Signet (2246, -0.31 DPS) [world_drop]; Naglering (11669, -3.02 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (187.1 DPS) | yes | Band of the Ogre King (18522, -0.06 DPS) [dungeon]; Myrmidon's Signet (2246, -0.30 DPS) [world_drop]; Naglering (11669, -5.47 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (187.1 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.74 DPS) [crafted] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (187.1 DPS) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (187.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -3.34 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Shroud of Domination; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Scalegut Treaders; finger1: Legionnaire's Band; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Draconic Infused Emblem; main_hand: The Unstoppable Force

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

