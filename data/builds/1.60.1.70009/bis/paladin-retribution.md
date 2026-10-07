# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 35.5. Weights run: 2.5s. Verify run: 1.1s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.110 ± 0.002, crit=0.155 ± 0.003 per rating point (14 rating = 1%, 2.172 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.489 per %), melee_haste=1.144 ± 0.061

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 0.7 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.22 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS, sim-verified) [crafted]; Serpent's Shoulders (5404, -0.20 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Catacomb Cloak (279899, -0.07 DPS) [quest]; Dark Leather Cloak (2316, -0.13 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.73 DPS) | yes | Veteran's Chain Shirt (250488, -0.20 DPS) [crafted]; Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.58 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Polar Gauntlets (7606, -0.15 DPS) [quest]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Cobrahn's Grasp (6460, -0.13 DPS) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (0.80 DPS) | yes | Veteran's Chain Leggings (250493, -0.13 DPS) [crafted]; Defender's Leather Pants (252445, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.6 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.4 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Signet of the Zhevra (285330, -0.28 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | The 1 Ring (8350, -0.21 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.22 DPS) [world_drop]; Signet of the Zhevra (285330, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (13.05 DPS) | yes | Duskbringer (2205, -2.34 DPS) [dungeon]; Monstrous Cleaver (279864, -2.65 DPS) [quest]; Smite's Mighty Hammer (7230, -3.54 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 77.8. Weights run: 2.4s. Verify run: 1.4s. 404 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.162 ± 0.003, crit=0.229 ± 0.005 per rating point (14 rating = 1%, 3.206 per %), hit=0.228 ± 0.002 per rating point (10 rating = 1%, 2.281 per %), melee_haste=1.522 ± 0.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.13 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.61 DPS) | yes | Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop]; Sentinel's Medallion (19541, -0.55 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.61 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.04 DPS) [crafted]; Glimmering Mail Pauldrons (6388, -0.09 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 attack_power points (0.43 DPS) | yes | Sergeant Major's Cape (16315, -0.06 DPS) [pvp]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.30 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.26 DPS) [crafted]; Hard Gold Cuirass (250533, -0.35 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.69 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Cultist's Armguards (270032, -0.26 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.95 DPS) | yes | The Frozen Clutch (23170, -0.09 DPS) [dungeon]; Bonefist Gauntlets (4465, -0.17 DPS) [world]; Mail Combat Gauntlets (4075, -0.23 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.04 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.13 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Chausses of Westfall (6087, -0.17 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.31 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (77.8 DPS) | yes | Disjointed Shoes (277226, -0.09 DPS) [quest]; Glimmering Mail Greaves (4073, -0.09 DPS) [world_drop]; Trouncing Boots (4464, -1.29 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (0.71 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 13.0 attack_power points (0.56 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Tiger Band (6749, -0.26 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.48 DPS) [dungeon]; Viscous Hammer (13045, -18.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wolfmaster Cape; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330320)

Set DPS (verified): 116.6. Weights run: 2.4s. Verify run: 1.3s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.444 ± 0.007, crit=0.628 ± 0.010 per rating point (14 rating = 1%, 8.794 per %), hit=0.389 ± 0.004 per rating point (10 rating = 1%, 3.887 per %), melee_haste=3.924 ± 0.173

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 34.8 attack_power points (1.61 DPS) | yes | Icemetal Barbute (10763, -0.31 DPS) [dungeon]; Hard Gold Coif (250537, -0.31 DPS) [crafted]; Chromite Barbute (8142, -0.74 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.93 DPS) | yes | Kaleidoscope Chain (13084, -0.47 DPS) [world_drop]; Ghostshard Talisman (7731, -0.50 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -0.56 DPS) [world_drop] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.02 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.19 DPS) [dungeon]; Chromite Pauldrons (8144, -1.45 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.7 attack_power points (0.68 DPS) | yes | Wolfmaster Cape (6314, -0.22 DPS) [dungeon]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop]; Dark Hooded Cape (5257, -1.17 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.1 attack_power points (1.44 DPS) | yes | Avenger's Armor (1488, -0.05 DPS) [dungeon]; Jouster's Chestplate (8157, -0.05 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.14 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.93 DPS) | yes | Pugilist Bracers (4438, -0.19 DPS) [dungeon]; Ravager's Armguards (14770, -0.20 DPS) [world_drop]; Yorgen Bracers (13012, -0.31 DPS) [world_drop] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.48 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Gloves of Holy Might (867, -0.15 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.16 DPS) [dungeon] |
| waist | Highlander's Plate Girdle (20125) (or Boar Champion's Belt (10768)) | The League of Arathor [rep] | 30.0 attack_power points (1.39 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.00 DPS) [rep]; Scarlet Belt (10329, -0.28 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.95 DPS) | yes | Firemane Leggings (13129, -0.19 DPS) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.43 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.1 attack_power points (1.35 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Obsidian Greaves (13068, -0.30 DPS) [world_drop]; Blackforge Greaves (6423, -0.31 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.93 DPS) | yes | Protector's Band (19515, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.93 DPS) | yes | Protector's Band (19515, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (116.6 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -11.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Highlander's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 54001000000000000-0000000000000000-05025331001330320)

Set DPS (verified): 139.5. Weights run: 2.6s. Verify run: 1.4s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=0.660 ± 0.014, crit=0.931 ± 0.020 per rating point (14 rating = 1%, 13.032 per %), hit=0.511 ± 0.007 per rating point (10 rating = 1%, 5.110 per %), melee_haste=4.845 ± 0.805

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (2.33 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -0.37 DPS) [dungeon]; Sunscale Helmet (14849, -0.49 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.94 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.28 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 37.2 attack_power points (1.75 DPS) | yes | Wyrmslayer Spaulders (13066, -0.26 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.33 DPS) [crafted]; Officer's Pauldrons (250576, -0.66 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.4 attack_power points (1.05 DPS) | yes | Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world]; Bloodlust Cape (14801, -0.71 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Tremors of the Earth [quest] | 52.8 attack_power points (2.48 DPS) | yes | Mixologist's Tunic (12793, -0.28 DPS) [dungeon]; Knight's Plate Hauberk (220794, -0.32 DPS) [vendor]; Valorous Chestguard (8274, -0.41 DPS) [world_drop] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 30.8 attack_power points (1.45 DPS) | yes | Officer's Wristguards (250581, -0.12 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.13 DPS) [dungeon]; Arena Bands (18711, -0.13 DPS) [world] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 55.0 attack_power points (2.58 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.86 DPS) [crafted]; Raider Gloves (272100, -0.87 DPS) [vendor]; Officer's Gloves (250551, -0.96 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 47.6 attack_power points (2.23 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.38 DPS) [dungeon]; Belt of the Gladiator (13134, -0.38 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.42 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 48.4 attack_power points (2.27 DPS) | yes | Scarlet Leggings (10330, -0.10 DPS) [dungeon]; Knight's Plate Leggings (220797, -0.11 DPS) [vendor]; Elemental Rockridge Leggings (17711, -0.21 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 39.4 attack_power points (1.85 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.23 DPS) [crafted]; Officer's Boots (250546, -0.29 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 27.9 attack_power points (1.31 DPS) | yes | Mark of Kern (2262, -0.37 DPS) [dungeon]; Assault Band (13095, -0.37 DPS) [world_drop]; Thunderbrow Ring (13097, -0.39 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 25.1 attack_power points (1.18 DPS) | yes | Assault Band (13095, -0.24 DPS) [world_drop]; Thunderbrow Ring (13097, -0.26 DPS) [world_drop]; Mark of Kern (2262, -0.95 DPS, sim-verified) [dungeon] |
| trinket1 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (139.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (139.5 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (139.5 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Drakefang Butcher (12463, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -5.08 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Knight-Lieutenant's Plate Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Sanctified Orb; trinket2: Frozen Heart of the Mountain; main_hand: Nightblade

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 54003000000000000-3230000000000000-05025331001330320)

Set DPS (verified): 245.1. Weights run: 2.6s. Verify run: 1.4s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=0.862 ± 0.020, crit=1.216 ± 0.028 per rating point (14 rating = 1%, 17.019 per %), hit=0.648 ± 0.009 per rating point (10 rating = 1%, 6.481 per %), melee_haste=7.960 ± 1.283

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulforge Helm (22091) | Saving the Best for Last [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Lamellar Faceguard (16474, +0.00 DPS) [vendor]; Field Marshal's Lamellar Headguard (231648, +0.00 DPS) [vendor]; Lionheart Helm (12640, -5.15 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 41.0 attack_power points (1.96 DPS) | yes | Amulet of the Darkmoon (19491, -0.13 DPS) [quest]; Imperial Jewel (11933, -0.43 DPS) [dungeon]; Beads of Ogre Might (22150, -0.50 DPS) [quest] |
| shoulder | Highlander's Lamellar Spaulders (20058) | The League of Arathor [rep] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Lamellar Shoulders (231651, -0.38 DPS) [vendor]; Highlander's Leather Shoulders (20059, -0.42 DPS) [rep]; Highlander's Plate Spaulders (20057, -4.14 DPS, sim-verified) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 38.3 attack_power points (1.83 DPS) | yes | Shroud of Domination (22337, -0.04 DPS) [dungeon]; Howler's Furs (272414, -0.18 DPS) [vendor]; Cape of the Black Baron (13340, -0.26 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Obsidian Mail Tunic (22191, -0.58 DPS) [crafted]; Cadaverous Armor (14637, -0.99 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.85 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.39 DPS) [rep]; Windtalker's Wristguards (19582, -0.51 DPS) [rep]; Bracers of Undead Slaying (23090, -4.20 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Radiant Gloves of the Dawn (227817, -0.46 DPS) [vendor]; Timbermaw Brawlers (19049, -0.51 DPS) [crafted]; Razor Gauntlets (18326, -4.57 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 76.4 attack_power points (3.66 DPS) | yes | Ferocity of the Timbermaw (227805, -0.22 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.59 DPS) [vendor]; Might of the Timbermaw (19044, -0.87 DPS) [crafted] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Titanic Leggings (22385, -0.19 DPS) [crafted]; Warbear Woolies (15065, -1.09 DPS) [crafted]; Cloudkeeper Legplates (14554, -2.78 DPS, sim-verified) [world_drop] |
| feet | Knight-Lieutenant's Lamellar Sabatons (227146) | Captain Dirgehammer [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Drudge Boots (21532, -0.01 DPS) [quest]; Battlechaser's Greaves (12555, -0.07 DPS) [dungeon]; Scalegut Treaders (275618, -2.76 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Ogre King (18522, -0.42 DPS) [dungeon]; Myrmidon's Signet (2246, -0.55 DPS) [world_drop]; Naglering (11669, -3.90 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Ogre King (18522, -0.24 DPS) [dungeon]; Myrmidon's Signet (2246, -0.38 DPS) [world_drop]; Naglering (11669, -3.32 DPS, sim-verified) [dungeon] |
| trinket1 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Darkmoon Card: Maelstrom (19289, -1.74 DPS, sim-verified) [quest] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -7.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Soulforge Helm; neck: Medallion of the Dawn; shoulder: Highlander's Lamellar Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Knight-Lieutenant's Lamellar Sabatons; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Draconic Infused Emblem; trinket2: Burst of Knowledge; main_hand: Blackblade of Shahram

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 32.2. Weights run: 2.5s. Verify run: 1.2s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.110 ± 0.002, crit=0.155 ± 0.003 per rating point (14 rating = 1%, 2.172 per %), hit=0.149 ± 0.002 per rating point (10 rating = 1%, 1.489 per %), melee_haste=1.144 ± 0.061

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.69 DPS) [crafted]; Brawler's Leather Hood (252504, -0.70 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 0.7 attack_power points (0.02 DPS) | yes | Erudite's Amulet (277204, -0.01 DPS) [quest] |
| shoulder | Silvered Bronze Shoulders (3481) | Blacksmithing [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent's Shoulders (5404, -0.20 DPS) [dungeon]; Rough Bronze Shoulders (3480, -0.59 DPS, sim-verified) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.29 DPS) | yes | Grave Shroud (279865, -0.06 DPS) [quest]; Subterranean Cape (14149, -0.07 DPS) [dungeon]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.73 DPS) | yes | Veteran's Chain Shirt (250488, -0.19 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.21 DPS) [crafted]; Totemic Leather Armor (252435, -0.22 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.36 DPS) | yes | Raptorcrest Bracers (270010, -0.15 DPS) [quest]; Bristlebark Bindings (14569, -0.21 DPS) [world_drop]; Runed Copper Bracers (2854, -0.22 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.58 DPS) | yes | Gold-flecked Gloves (5195, -0.07 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.15 DPS) [dungeon]; Dagmire Gauntlets (6481, -0.21 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.65 DPS) | yes | Cobrahn's Grasp (6460, -0.13 DPS) [dungeon]; Ruffian Belt (5975, -0.22 DPS) [world]; Hulking Belt (14746, -0.29 DPS) [world_drop] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 18.6 attack_power points (0.67 DPS) | yes | Defender's Leather Pants (252445, -0.00 DPS) [crafted]; Totemic Leather Pants (252446, -0.02 DPS) [crafted]; Hulking Leggings (14748, -0.08 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 10.6 attack_power points (0.38 DPS) | yes | Veteran's Boots (250503, -0.00 DPS) [crafted]; Guard's Boots (250504, -0.02 DPS) [crafted]; Defender's Leather Boots (252441, -0.02 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.4 attack_power points (0.31 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Ring of the Moon (12052, -0.23 DPS) [world_drop]; Demon Band (12054, -0.40 DPS, sim-verified) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | The 1 Ring (8350, -0.14 DPS) [world]; Ring of the Moon (12052, -0.15 DPS) [world_drop]; Demon Band (12054, -0.48 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (11.14 DPS) | yes | Smite's Mighty Hammer (7230, -0.28 DPS) [dungeon]; Duskbringer (2205, -0.43 DPS) [dungeon]; Monstrous Cleaver (279864, -0.74 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Silvered Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 77.5. Weights run: 2.4s. Verify run: 1.4s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.162 ± 0.003, crit=0.229 ± 0.005 per rating point (14 rating = 1%, 3.206 per %), hit=0.228 ± 0.002 per rating point (10 rating = 1%, 2.281 per %), melee_haste=1.522 ± 0.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.13 DPS) | yes | Veteran's Chain Helm (250498, -0.09 DPS) [crafted]; Defender's Leather Helm (252455, -0.09 DPS) [crafted]; Crusader's Chain Helm (250502, -0.17 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.61 DPS) | yes | Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.26 DPS) [world_drop]; Scout's Medallion (19537, -0.55 DPS) [rep] |
| shoulder | Golden Scale Shoulders (3841) (or Mail Combat Spaulders (6404)) | Blacksmithing [crafted] | 14.0 attack_power points (0.61 DPS) | yes | Mail Combat Spaulders (6404, +0.00 DPS) [world_drop]; Barbaric Iron Shoulders (7913, -0.04 DPS) [crafted]; Elite Shoulders (4835, -0.09 DPS) [vendor] |
| back | Wildhunter Cloak (16658) (or Wolfmaster Cape (6314)) | The Hunt Completed [quest] | 10.0 attack_power points (0.43 DPS) | yes | Wolfmaster Cape (6314, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.09 DPS) [world_drop]; Slayer's Cape (14752, -0.09 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.30 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.26 DPS) [crafted]; Hard Gold Cuirass (250533, -0.35 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.69 DPS) | yes | Yorgen Bracers (13012, -0.15 DPS) [world_drop]; Bands of Serra'kis (6902, -0.17 DPS) [dungeon]; Grimtoll Wristguards (15459, -0.25 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.95 DPS) | yes | The Frozen Clutch (23170, -0.09 DPS) [dungeon]; Warsong Gauntlets (16978, -0.09 DPS) [quest]; Bonefist Gauntlets (4465, -0.17 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.04 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Girdle of Golem Strength (9405, -0.00 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.00 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.13 DPS) | yes | Golden Scale Leggings (3843, -0.17 DPS) [crafted]; Slayer's Pants (14757, -0.17 DPS) [world_drop]; Veteran's Silvered Chain Leggings (250523, -0.31 DPS, sim-verified) [crafted] |
| feet | Hard Gold Boots (250534) | Blacksmithing [crafted] | sim-verified (77.5 DPS) | yes | Glimmering Mail Greaves (4073, -0.09 DPS) [world_drop]; Slayer's Slippers (14756, -0.09 DPS) [world_drop]; Trouncing Boots (4464, -1.28 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.5 attack_power points (0.71 DPS) | yes | Tiger Band (6749, -0.19 DPS) [quest]; Silverlaine's Family Seal (6321, -0.28 DPS) [dungeon]; Ironspine's Eye (7686, -0.30 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 13.0 attack_power points (0.56 DPS) | yes | Silverlaine's Family Seal (6321, -0.13 DPS) [dungeon]; Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Tiger Band (6749, -0.29 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Cobalt Crusher (7730, -1.48 DPS) [dungeon]; Viscous Hammer (13045, -17.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Golden Scale Shoulders; back: Wildhunter Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Hard Gold Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330320)

Set DPS (verified): 115.9. Weights run: 2.4s. Verify run: 1.3s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.444 ± 0.007, crit=0.628 ± 0.010 per rating point (14 rating = 1%, 8.794 per %), hit=0.389 ± 0.004 per rating point (10 rating = 1%, 3.887 per %), melee_haste=3.924 ± 0.173

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 34.8 attack_power points (1.61 DPS) | yes | Icemetal Barbute (10763, -0.31 DPS) [dungeon]; Hard Gold Coif (250537, -0.31 DPS) [crafted]; Chromite Barbute (8142, -0.74 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.93 DPS) | yes | Ethereal Talisman (4430, -0.38 DPS) [quest]; Kaleidoscope Chain (13084, -0.47 DPS) [world_drop]; Ghostshard Talisman (7731, -0.50 DPS, sim-verified) [dungeon] |
| shoulder | Hard Gold Pauldrons (250539) | Blacksmithing [crafted] | 22.0 attack_power points (1.02 DPS) | yes | Shining Mithril Pauldrons (250541, -0.09 DPS) [crafted]; Imperial Leather Spaulders (4737, -0.19 DPS) [dungeon]; Chromite Pauldrons (8144, -1.44 DPS, sim-verified) [dungeon] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | 14.7 attack_power points (0.68 DPS) | yes | Wolfmaster Cape (6314, -0.22 DPS) [dungeon]; Wildhunter Cloak (16658, -0.22 DPS) [quest]; Dark Hooded Cape (5257, -1.16 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 31.1 attack_power points (1.44 DPS) | yes | Avenger's Armor (1488, -0.05 DPS) [dungeon]; Jouster's Chestplate (8157, -0.05 DPS) [dungeon]; Shining Mithril Breastplate (250540, -0.14 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.93 DPS) | yes | Pugilist Bracers (4438, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp] |
| hands | Truesilver Gauntlets (7938) | Blacksmithing [crafted] | 32.0 attack_power points (1.48 DPS) | yes | Gauntlets of Divinity (7724, -0.00 DPS) [dungeon]; Gloves of Holy Might (867, -0.15 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.16 DPS) [dungeon] |
| waist | Defiler's Plate Girdle (20206) (or Boar Champion's Belt (10768)) | The Defilers [rep] | 30.0 attack_power points (1.39 DPS) | yes | Boar Champion's Belt (10768, +0.00 DPS) [dungeon]; Defiler's Leather Girdle (20192, -0.00 DPS) [rep]; Tharg's Shoelace (9705, -0.19 DPS) [quest] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (1.95 DPS) | yes | Firemane Leggings (13129, -0.19 DPS) [world_drop]; Orcish War Leggings (7929, -0.37 DPS) [crafted]; Symbolic Legplates (14829, -0.43 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 29.1 attack_power points (1.35 DPS) | yes | Prowler's Leather Shoes (252465, +0.00 DPS, sim-verified) [crafted]; Obsidian Greaves (13068, -0.30 DPS) [world_drop]; Blackforge Greaves (6423, -0.31 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (0.93 DPS) | yes | Legionnaire's Band (19512, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.93 DPS) | yes | Legionnaire's Band (19512, -0.02 DPS) [rep]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (115.9 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -18.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Hard Gold Pauldrons; back: First Sergeant's Cloak; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Truesilver Gauntlets; waist: Defiler's Plate Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 54001000000000000-0000000000000000-05025331001330320)

Set DPS (verified): 136.2. Weights run: 2.6s. Verify run: 1.3s. 702 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=0.660 ± 0.014, crit=0.931 ± 0.020 per rating point (14 rating = 1%, 13.032 per %), hit=0.511 ± 0.007 per rating point (10 rating = 1%, 5.110 per %), melee_haste=4.845 ± 0.805

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (2.33 DPS) | yes | Blood Guard's Plate Helm (220803, -0.13 DPS) [vendor]; Raging Berserker's Helm (7719, -0.37 DPS) [dungeon]; Sunscale Helmet (14849, -0.49 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (0.94 DPS) | yes | Skibi's Pendant (13089, -0.02 DPS) [world_drop]; Woven Ivy Necklace (19159, -0.04 DPS) [quest]; Ghostshard Talisman (7731, -0.28 DPS) [dungeon] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 33.9 attack_power points (1.59 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.10 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.4 attack_power points (1.05 DPS) | yes | First Sergeant's Cloak (16340, -0.25 DPS) [pvp]; Dark Hooded Cape (5257, -0.33 DPS) [world]; Bloodlust Cape (14801, -0.78 DPS, sim-verified) [world_drop] |
| chest | Warforged Chestplate (11195) | Broken Alliances [quest] | 52.8 attack_power points (2.48 DPS) | yes | Mixologist's Tunic (12793, -0.28 DPS) [dungeon]; Stone Guard's Plate Armor (220801, -0.32 DPS) [vendor]; Valorous Chestguard (8274, -0.41 DPS) [world_drop] |
| wrist | Runed Golem Shackles (12550) | Blackrock Depths: Anvilrage Overseer [dungeon] | 30.8 attack_power points (1.45 DPS) | yes | Berserker Bracers (19580, +0.00 DPS) [pvp]; Officer's Wristguards (250581, +0.00 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 55.0 attack_power points (2.58 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.86 DPS) [crafted]; Raider Gloves (272100, -0.87 DPS) [vendor]; Officer's Gloves (250551, -0.88 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 47.6 attack_power points (2.23 DPS) | yes | Atal'alarion's Tusk Ring (10798, -0.38 DPS) [dungeon]; Belt of the Gladiator (13134, -0.38 DPS) [world_drop]; Prowler's Leather Waistguard (252473, -0.42 DPS) [crafted] |
| legs | Golem Shard Leggings (13074) | World drop [world_drop] | 48.4 attack_power points (2.27 DPS) | yes | Scarlet Leggings (10330, -0.10 DPS) [dungeon]; Stone Guard's Plate Leggings (220798, -0.11 DPS) [vendor]; Elemental Rockridge Leggings (17711, -0.21 DPS) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 39.4 attack_power points (1.85 DPS) | yes | Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Officer's Sabatons (250561, -0.23 DPS) [crafted]; Officer's Boots (250546, -0.29 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 27.9 attack_power points (1.31 DPS) | yes | White Bone Band (11862, -0.18 DPS) [quest]; Mark of Kern (2262, -0.37 DPS) [dungeon]; Assault Band (13095, -0.37 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 25.1 attack_power points (1.18 DPS) | yes | Mark of Kern (2262, -0.24 DPS) [dungeon]; Assault Band (13095, -0.24 DPS) [world_drop]; White Bone Band (11862, -1.46 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (136.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.2 DPS) | yes | Molten Heart of the Mountain (249470, -1.28 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (136.2 DPS) | yes | Taran Icebreaker (2915, -0.72 DPS) [world_drop]; Drakefang Butcher (12463, -1.56 DPS) [dungeon]; Blight (7959, -2.04 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; shoulder: Officer's Pauldrons; back: Blackveil Cape; chest: Warforged Chestplate; wrist: Runed Golem Shackles; hands: Raider Gauntlets; waist: Girdle of Beastial Fury; legs: Golem Shard Leggings; feet: Battlechaser's Greaves; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 54003000000000000-3230000000000000-05025331001330320)

Set DPS (verified): 224.0. Weights run: 2.6s. Verify run: 1.4s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=0.862 ± 0.020, crit=1.216 ± 0.028 per rating point (14 rating = 1%, 17.019 per %), hit=0.648 ± 0.009 per rating point (10 rating = 1%, 6.481 per %), melee_haste=7.960 ± 1.283

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soulforge Greathelm (226976) | Mokvar [vendor] | sim-verified (224.0 DPS) | yes | Blood Guard's Plate Helm (220803, -0.01 DPS) [vendor]; Warbear Helm (252485, -0.03 DPS) [crafted]; Lionheart Helm (12640, -3.76 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 41.0 attack_power points (1.96 DPS) | yes | Amulet of the Darkmoon (19491, -0.13 DPS) [quest]; Imperial Jewel (11933, -0.43 DPS) [dungeon]; Beads of Ogre Might (22150, -0.50 DPS) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | 54.2 attack_power points (2.60 DPS) | yes | Defiler's Leather Shoulders (20194, -0.42 DPS) [rep]; Darkspear Spaulders (272108, -0.52 DPS) [vendor]; Blood Guard's Plate Pauldrons (220796, -0.62 DPS) [vendor] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 38.3 attack_power points (1.83 DPS) | yes | Shroud of Domination (22337, -0.04 DPS) [dungeon]; Howler's Furs (272414, -0.18 DPS) [vendor]; Cape of the Black Baron (13340, -0.26 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Obsidian Mail Tunic (22191, -0.58 DPS) [crafted]; Cadaverous Armor (14637, -0.99 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -9.62 DPS, sim-verified) [world] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.39 DPS) [rep]; Windtalker's Wristguards (19582, -0.51 DPS) [rep]; Bracers of Undead Slaying (23090, -4.04 DPS, sim-verified) [world] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Radiant Gloves of the Dawn (227817, -0.46 DPS) [vendor]; Timbermaw Brawlers (19049, -0.51 DPS) [crafted]; Razor Gauntlets (18326, -4.35 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 76.4 attack_power points (3.66 DPS) | yes | Ferocity of the Timbermaw (227805, -0.22 DPS) [vendor]; Dense Timbermaw Belt (227807, -0.59 DPS) [vendor]; Might of the Timbermaw (19044, -0.87 DPS) [crafted] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Titanic Leggings (22385, -0.19 DPS) [crafted]; Outrider's Plate Legguards (22651, -0.71 DPS) [rep]; Cloudkeeper Legplates (14554, -4.97 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 44.9 attack_power points (2.15 DPS) | yes | Drudge Boots (21532, -0.08 DPS) [quest]; Battlechaser's Greaves (12555, -0.14 DPS) [dungeon]; Clutchlord's Stompers (275627, -0.15 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Ogre King (18522, -0.42 DPS) [dungeon]; Myrmidon's Signet (2246, -0.55 DPS) [world_drop]; Naglering (11669, -5.48 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Ogre King (18522, -0.24 DPS) [dungeon]; Myrmidon's Signet (2246, -0.38 DPS) [world_drop]; Naglering (11669, -3.19 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Draconic Infused Emblem (22268) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, -1.93 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackblade of Shahram (12592, +0.00 DPS) [dungeon]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Soulforge Greathelm; neck: Medallion of the Dawn; shoulder: Defiler's Plate Spaulders; back: Deathguard's Cloak; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Burst of Knowledge; trinket2: Draconic Infused Emblem; main_hand: The Unstoppable Force

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

