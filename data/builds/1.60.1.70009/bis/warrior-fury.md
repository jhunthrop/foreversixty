# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 33.3. Weights run: 2.1s. Verify run: 1.2s. 302 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.589 ± 0.017, crit=0.842 ± 0.025 per rating point (14 rating = 1%, 11.784 per %), hit=1.062 ± 0.049 per rating point (10 rating = 1%, 10.620 per %), melee_haste=7.191 ± 0.471

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.23 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.76 DPS) [crafted]; Brawler's Leather Hood (252504, -0.82 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 3.5 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.06 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.32 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.16 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Grave Shroud (279865, -0.04 DPS) [quest]; Catacomb Cloak (279899, -0.11 DPS) [quest]; Dark Leather Cloak (2316, -0.12 DPS) [crafted] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Veteran's Chain Shirt (250488, -0.22 DPS, sim-verified) [crafted]; Defender's Leather Armor (252434, -0.23 DPS) [crafted]; Brawler's Leather Armor (252490, -0.31 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.54 DPS) | yes | Cryptwalker Bracers (280095, -0.11 DPS) [quest]; Bravo's Armbands (270015, -0.20 DPS) [quest]; Bristlebark Bindings (14569, -0.23 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.86 DPS) | yes | Gold-flecked Gloves (5195, -0.11 DPS) [dungeon]; Polar Gauntlets (7606, -0.21 DPS) [quest]; Blackened Defias Gloves (10401, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.96 DPS) | yes | Cobrahn's Grasp (6460, -0.12 DPS) [dungeon]; Ruffian Belt (5975, -0.32 DPS) [world]; Brawler's Leather Belt (252428, -0.41 DPS) [crafted] |
| legs | Chausses of Westfall (6087) | The Defias Brotherhood [quest] | 22.0 attack_power points (1.18 DPS) | yes | Veteran's Chain Leggings (250493, -0.06 DPS) [crafted]; Defender's Leather Pants (252445, -0.09 DPS) [crafted]; Totemic Leather Pants (252446, -0.21 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 12.9 attack_power points (0.69 DPS) | yes | Veteran's Boots (250503, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.12 DPS) [world_drop]; Defender's Leather Boots (252441, -0.16 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.4 attack_power points (0.56 DPS) | yes | Signet of the Zhevra (285330, -0.37 DPS) [world]; The 1 Ring (8350, -0.42 DPS) [world]; Ring of the Moon (12052, -0.45 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Signet of the Zhevra (285330, -0.29 DPS, sim-verified) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.32 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.34 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.7 attack_power points (12.74 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop]; Redbeard Crest (12997, -12.10 DPS) [world_drop]; Furen's Favor (6970, -12.42 DPS) [quest] |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567), Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.21 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.09 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Chausses of Westfall; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 61.7. Weights run: 2.3s. Verify run: 1.3s. 490 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.797 ± 0.025, crit=1.138 ± 0.036 per rating point (14 rating = 1%, 15.933 per %), hit=1.509 ± 0.069 per rating point (10 rating = 1%, 15.086 per %), melee_haste=10.645 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.48 DPS) | yes | Barbaric Iron Helm (7915, -0.05 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS) [crafted]; Defender's Leather Helm (252455, -0.11 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.80 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.34 DPS) [world_drop]; Sentinel's Medallion (19541, -0.43 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 18.8 attack_power points (1.07 DPS) | yes | Barbaric Iron Shoulders (7913, -0.11 DPS) [crafted]; Golden Scale Shoulders (3841, -0.27 DPS) [crafted]; Mail Combat Spaulders (6404, -0.27 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.6 attack_power points (0.66 DPS) | yes | Sergeant Major's Cape (16315, -0.02 DPS) [pvp]; Wolfmaster Cape (6314, -0.09 DPS) [dungeon]; Slayer's Cape (14752, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.70 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.30 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.91 DPS) | yes | Yorgen Bracers (13012, -0.09 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop]; Bands of Serra'kis (6902, -0.23 DPS) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.25 DPS) | yes | The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.16 DPS) [world_drop]; Insignia Gloves (6408, -0.18 DPS) [world_drop] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 24.8 attack_power points (1.41 DPS) | yes | Girdle of Golem Strength (9405, -0.04 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.04 DPS) [rep]; Highlander's Chain Girdle (20090, -0.04 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 27.6 attack_power points (1.57 DPS) | yes | Ferine Leggings (6690, -0.09 DPS) [dungeon]; Golden Scale Leggings (3843, -0.32 DPS) [crafted]; Chausses of Westfall (6087, -0.32 DPS) [quest] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.6 attack_power points (1.11 DPS) | yes | Brawler's Leather Boots (252439, -0.32 DPS) [crafted]; Alacritous Treads (277234, -0.32 DPS) [quest]; Hard Gold Boots (250534, -0.49 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 18.4 attack_power points (1.04 DPS) | yes | Ironspine's Eye (7686, -0.18 DPS) [dungeon]; Tiger Band (6749, -0.36 DPS) [quest]; Silverlaine's Family Seal (6321, -0.48 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.8 attack_power points (0.95 DPS) | yes | Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Tiger Band (6749, -0.27 DPS) [quest]; Silverlaine's Family Seal (6321, -0.39 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (19.62 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (19.02 DPS) | yes | Royal Diplomatic Scepter (9457, -0.87 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -6.26 DPS) [quest]; Slayer's Shield (15892, -18.21 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.51 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; Long Battle Bow (15284, -0.17 DPS) [world_drop]; Precision Bow (8183, -0.25 DPS) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 490, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow

### Band 40 (human, 00000000000000000-05153105022011500-000000000000000000)

Set DPS (verified): 114.0. Weights run: 2.4s. Verify run: 1.3s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.615 ± 0.063, crit=2.307 ± 0.090 per rating point (14 rating = 1%, 32.304 per %), hit=2.575 ± 0.138 per rating point (10 rating = 1%, 25.751 per %), melee_haste=11.387 ± 1.289

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 58.3 attack_power points (3.61 DPS) | yes | Chromite Barbute (8142, -0.77 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.15 DPS) [crafted]; Barbaric Iron Helm (7915, -1.60 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.24 DPS) | yes | Sentinel's Medallion (19540, -0.14 DPS) [rep]; Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Ghostshard Talisman (7731, -0.37 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.8 attack_power points (1.84 DPS) | yes | Forest Tracker Epaulets (2278, -0.12 DPS) [world_drop]; Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.48 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 24.2 attack_power points (1.50 DPS) | yes | Sergeant Major's Cape (16336, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.42 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.52 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 46.7 attack_power points (2.89 DPS) | yes | Kolkar Marauder Chain (6773, -0.46 DPS) [quest]; Carapace of Tuten'kash (10775, -0.85 DPS) [dungeon]; Jouster's Chestplate (8157, -1.03 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.5 attack_power points (1.27 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.17 DPS) [world_drop]; Yorgen Bracers (13012, -0.22 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 52.3 attack_power points (3.24 DPS) | yes | Dragonscale Gauntlets (8347, -0.64 DPS) [crafted]; Scarlet Gauntlets (10331, -0.75 DPS) [dungeon]; Plated Fist of Hakoo (13071, -1.00 DPS) [world_drop] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 40.3 attack_power points (2.50 DPS) | yes | Ogron's Sash (13117, -0.48 DPS) [world_drop]; Boar Champion's Belt (10768, -0.64 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.64 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.60 DPS) | yes | Symbolic Legplates (14829, -0.14 DPS) [world_drop]; Firemane Leggings (13129, -0.25 DPS) [world_drop]; Triprunner Dungarees (9624, -0.43 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.3 attack_power points (2.31 DPS) | yes | Blackforge Greaves (6423, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.34 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.9 attack_power points (1.79 DPS) | yes | Thunderbrow Ring (13097, -0.50 DPS) [world_drop]; Falcon's Hook (7552, -0.52 DPS) [dungeon]; Ring of the Underwood (2951, -0.54 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.5 attack_power points (1.40 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ring of the Underwood (2951, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (114.0 DPS) | yes | Shoni's Disarming Tool (9608, -13.96 DPS) [quest]; Skullance Shield (13081, -26.48 DPS) [world_drop]; Ardent Custodian (868, -39.66 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.15 DPS) [quest]; The Silencer (13138, -0.18 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.94 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 185.1. Weights run: 2.5s. Verify run: 1.6s. 867 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.303 ± 0.057, crit=1.861 ± 0.081 per rating point (14 rating = 1%, 26.056 per %), hit=1.651 ± 0.141 per rating point (10 rating = 1%, 16.512 per %), melee_haste=11.121 ± 1.183

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 78.6 attack_power points (7.49 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.13 DPS, sim-verified) [vendor]; Raging Berserker's Helm (7719, -2.53 DPS) [dungeon]; Embrace of the Lycan (9479, -2.91 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.9 attack_power points (2.57 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.66 DPS) [quest]; Sentinel's Medallion (19539, -1.08 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 49.5 attack_power points (4.72 DPS) | yes | Officer's Pauldrons (250576, -1.25 DPS) [crafted]; Wyrmslayer Spaulders (13066, -1.44 DPS) [world_drop]; Knight-Lieutenant's Plate Pauldrons (220795, -1.88 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.2 attack_power points (2.88 DPS) | yes | Dark Hooded Cape (5257, -0.88 DPS) [world]; Sergeant Major's Cape (16336, -0.99 DPS) [pvp]; Dark Phantom Cape (13122, -1.02 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 56.1 attack_power points (5.34 DPS) | yes | Mixologist's Tunic (12793, -0.55 DPS) [dungeon]; Warforged Chestplate (11195, -0.77 DPS) [quest]; Warbear Harness (15064, -1.01 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.8 attack_power points (2.84 DPS) | yes | Runed Golem Shackles (12550, -0.17 DPS) [dungeon]; Bracers of the Stone Princess (17714, -0.17 DPS) [dungeon]; Arena Bands (18711, -0.17 DPS) [world] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 57.0 attack_power points (5.43 DPS) | yes | Gloves of Holy Might (867, -1.05 DPS) [world_drop]; Officer's Gloves (250551, -1.27 DPS) [crafted]; Raider Gloves (272100, -1.83 DPS, sim-verified) [vendor] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 48.1 attack_power points (4.58 DPS) | yes | Highlander's Chain Girdle (20088, -0.19 DPS) [rep]; Highlander's Leather Girdle (20115, -0.19 DPS) [rep]; Highlander's Plate Girdle (20124, -0.19 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 63.1 attack_power points (6.01 DPS) | yes | Gryphon Rider's Leggings (9652, -1.29 DPS) [quest]; Centurion Legplates (10740, -1.29 DPS) [quest]; Stormshroud Pants (15057, -1.87 DPS, sim-verified) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 44.9 attack_power points (4.28 DPS) | yes | Prowler's Leather Boots (252468, -0.44 DPS) [crafted]; Skulker's Leather Boots (252469, -0.57 DPS) [crafted]; Officer's Sabatons (250561, -0.69 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.5 attack_power points (3.48 DPS) | yes | Mark of Kern (2262, -1.57 DPS) [dungeon]; Assault Band (13095, -1.57 DPS) [world_drop]; Thunderbrow Ring (13097, -1.58 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.7 attack_power points (3.02 DPS) | yes | Mark of Kern (2262, -1.12 DPS) [dungeon]; Assault Band (13095, -1.12 DPS) [world_drop]; Thunderbrow Ring (13097, -1.13 DPS) [world_drop] |
| trinket1 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (185.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (185.1 DPS) | yes | Molten Heart of the Mountain (249470, -2.48 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (185.1 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hanzo Sword (8190, -1.55 DPS, sim-verified) [world_drop] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.94 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -5.99 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.53 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (185.1 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -2.26 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Diamond Flask; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 867, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 294.0. Weights run: 2.4s. Verify run: 1.7s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.316 ± 0.120, crit=3.309 ± 0.172 per rating point (14 rating = 1%, 46.322 per %), hit=3.538 ± 0.321 per rating point (10 rating = 1%, 35.377 per %), melee_haste=21.075 ± 2.823

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 199.4 attack_power points (20.38 DPS) | yes | Lieutenant Commander's Plate Helm (227044, -7.74 DPS) [vendor]; Eye of Rend (12587, -8.37 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (294.0 DPS) | yes | Medallion of the Dawn (22659, -0.20 DPS) [quest]; Pendant of Celerity (22340, -0.23 DPS) [dungeon]; Rage of Mugamba (19577, -4.52 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 108.9 attack_power points (11.12 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Plate Shoulderguards (231537, -0.04 DPS) [pvp]; Truestrike Shoulders (12927, -1.44 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 63.4 attack_power points (6.48 DPS) | yes | Windshear Cape (20691, -1.29 DPS) [world]; Fel Cape (279269, -1.74 DPS) [crafted]; Cape of the Black Baron (13340, -1.81 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (294.0 DPS) | yes | Timbermaw Tunic (252484, -1.37 DPS) [crafted]; Savage Gladiator Chain (11726, -3.42 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -14.89 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (294.0 DPS) | yes | Forest Stalker's Bracers (19587, -1.60 DPS) [rep]; Vambraces of the Sadist (13400, -2.18 DPS) [dungeon]; Bracers of Undead Slaying (23090, -7.60 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-verified (294.0 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.54 DPS) [pvp]; Raider Gloves (272099, -0.72 DPS) [vendor]; Razor Gauntlets (18326, -5.73 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 100.3 attack_power points (10.25 DPS) | yes | Ferocity of the Timbermaw (227805, -0.91 DPS) [vendor]; Marksman's Girdle (22232, -1.67 DPS) [dungeon]; Belt of Preserved Heads (20216, -4.49 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (294.0 DPS) | yes | Sentinel's Plate Legguards (237825, -2.77 DPS) [vendor]; Titanic Leggings (22385, -3.27 DPS) [crafted]; Cloudkeeper Legplates (14554, -4.84 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 81.7 attack_power points (8.35 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Boots of Heroism (21995, -0.65 DPS) [quest]; Battleboots of Heroism (226857, -0.65 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (294.0 DPS) | yes | Tarnished Elven Ring (18500, -2.82 DPS) [dungeon]; Cutthroat's Signet (272408, -3.06 DPS) [vendor]; Naglering (11669, -9.88 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (294.0 DPS) | yes | Tarnished Elven Ring (18500, -0.71 DPS) [dungeon]; Cutthroat's Signet (272408, -0.95 DPS) [vendor]; Naglering (11669, -7.09 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (294.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -3.51 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (294.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Diamond Flask (20130, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (294.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Felstriker (12590, -2.41 DPS, sim-verified) [dungeon] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (294.0 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -27.66 DPS) [dungeon]; Skullflame Shield (1168, -99.24 DPS, sim-verified) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (294.0 DPS) | yes | Satyr's Bow (18323, -0.41 DPS) [dungeon]; Blackcrow (12651, -0.77 DPS) [dungeon]; Dark Iron Rifle (16004, -5.06 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

### Band 60, raid preset (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 778.7. Weights run: 2.8s. Verify run: 1.8s. 1943 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.950 ± 0.139, crit=2.785 ± 0.198 per rating point (14 rating = 1%, 38.995 per %), hit=3.817 ± 0.353 per rating point (10 rating = 1%, 38.174 per %), melee_haste=33.827 ± 3.794

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 190.3 attack_power points (44.19 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -16.52 DPS) [vendor]; Mask of the Unforgiven (13404, -19.84 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (778.7 DPS) | yes | Mark of Fordring (15411, -0.56 DPS) [quest]; Medallion of the Dawn (22659, -1.03 DPS) [quest]; Rage of Mugamba (19577, -5.03 DPS, sim-verified) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 100.3 attack_power points (23.30 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, -2.02 DPS) [vendor]; Darkspear Pauldrons (272105, -9.14 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.2 attack_power points (15.36 DPS) | yes | Windshear Cape (20691, -4.86 DPS) [world]; Cloak of the Honor Guard (20073, -5.21 DPS) [rep]; Cape of the Black Baron (13340, -6.73 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (778.7 DPS) | yes | Timbermaw Tunic (252484, -2.14 DPS) [crafted]; Savage Gladiator Chain (11726, -7.32 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -28.36 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (778.7 DPS) | yes | Forest Stalker's Bracers (19587, -4.21 DPS) [rep]; Berserker Bracers (19578, -5.47 DPS) [rep]; Bracers of Undead Slaying (23090, -12.93 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | sim-verified (778.7 DPS) | yes | Marshal's Plate Gauntlets (231541, -3.46 DPS) [pvp]; Stormshroud Gloves (21278, -3.88 DPS) [crafted]; Razor Gauntlets (18326, -14.17 DPS, sim-verified) [dungeon] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 95.4 attack_power points (22.15 DPS) | yes | Radiant Girdle of the Dawn (227814, -0.56 DPS) [vendor]; Ferocity of the Timbermaw (227805, -2.39 DPS) [vendor]; Marksman's Girdle (22232, -3.78 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (778.7 DPS) | yes | Titanic Leggings (22385, -2.10 DPS) [crafted]; Sentinel's Plate Legguards (237825, -3.31 DPS) [vendor]; Cloudkeeper Legplates (14554, -11.96 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 78.2 attack_power points (18.15 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Windreaver Greaves (13967, -0.23 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (778.7 DPS) | yes | Tarnished Elven Ring (18500, -5.98 DPS) [dungeon]; Cutthroat's Signet (272408, -6.43 DPS) [vendor]; Naglering (11669, -17.61 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (778.7 DPS) | yes | Tarnished Elven Ring (18500, -1.36 DPS) [dungeon]; Cutthroat's Signet (272408, -1.81 DPS) [vendor]; Naglering (11669, -12.81 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (778.7 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -3.77 DPS, sim-verified) [dungeon] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (778.7 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -7.40 DPS, sim-verified) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (778.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Quel'Serrar (18348, -23.60 DPS, sim-verified) [quest] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (778.7 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -61.94 DPS) [dungeon]; Skullflame Shield (1168, -218.38 DPS, sim-verified) [world_drop] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (778.7 DPS) | yes | Blackcrow (12651, -0.89 DPS) [dungeon]; The Purifier (22656, -1.17 DPS) [quest]; Dark Iron Rifle (16004, -8.37 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Diamond Flask; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Satyr's Bow

No-known-source sample (15 of 1943, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle

## Horde

### Band 20 (troll, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 33.9. Weights run: 2.1s. Verify run: 1.2s. 272 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.589 ± 0.017, crit=0.842 ± 0.025 per rating point (14 rating = 1%, 11.784 per %), hit=1.062 ± 0.049 per rating point (10 rating = 1%, 10.620 per %), melee_haste=7.191 ± 0.471

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.24 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.76 DPS) [crafted]; Brawler's Leather Hood (252504, -0.82 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 3.5 attack_power points (0.19 DPS) | yes | Erudite's Amulet (277204, -0.06 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.32 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.16 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Grave Shroud (279865, -0.04 DPS) [quest]; Subterranean Cape (14149, -0.11 DPS) [dungeon]; Catacomb Cloak (279899, -0.11 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (1.07 DPS) | yes | Defender's Leather Armor (252434, -0.23 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.25 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.31 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.54 DPS) | yes | Bristlebark Bindings (14569, -0.23 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.24 DPS, sim-verified) [quest]; Runed Copper Bracers (2854, -0.32 DPS) [crafted] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.86 DPS) | yes | Gold-flecked Gloves (5195, -0.11 DPS) [dungeon]; Blackened Defias Gloves (10401, -0.21 DPS) [dungeon]; Fletcher's Gloves (7348, -0.23 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.96 DPS) | yes | Cobrahn's Grasp (6460, -0.18 DPS, sim-verified) [dungeon]; Ruffian Belt (5975, -0.32 DPS) [world]; Brawler's Leather Belt (252428, -0.41 DPS) [crafted] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 20.9 attack_power points (1.12 DPS) | yes | Defender's Leather Pants (252445, -0.03 DPS) [crafted]; Totemic Leather Pants (252446, -0.16 DPS) [crafted]; Hulking Leggings (14748, -0.17 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 12.9 attack_power points (0.69 DPS) | yes | Veteran's Boots (250503, -0.03 DPS) [crafted]; Feet of the Lynx (1121, -0.12 DPS) [world_drop]; Defender's Leather Boots (252441, -0.16 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.4 attack_power points (0.56 DPS) | yes | Loop of Sacrifice (281673, -0.23 DPS) [quest]; Signet of the Zhevra (285330, -0.37 DPS) [world]; The 1 Ring (8350, -0.42 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.43 DPS) | yes | Loop of Sacrifice (281673, -0.11 DPS) [quest]; Signet of the Zhevra (285330, -0.24 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (13.34 DPS) | yes | The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 237.7 attack_power points (12.74 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop]; Redbeard Crest (12997, -12.10 DPS) [world_drop]; Ruga's Bulwark (7120, -12.42 DPS) [quest] |
| ranged | Cracked Blacksmith Hammer (285279) (or Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.21 DPS) | yes | Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.09 DPS) [world_drop]; Heavy Shortbow (3036, -0.11 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 272, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 62.6. Weights run: 2.3s. Verify run: 1.2s. 454 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=0.797 ± 0.025, crit=1.138 ± 0.036 per rating point (14 rating = 1%, 15.933 per %), hit=1.509 ± 0.069 per rating point (10 rating = 1%, 15.086 per %), melee_haste=10.645 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tusken Helm (6686) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 26.0 attack_power points (1.48 DPS) | yes | Barbaric Iron Helm (7915, -0.05 DPS) [crafted]; Veteran's Chain Helm (250498, -0.11 DPS) [crafted]; Defender's Leather Helm (252455, -0.11 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.80 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.34 DPS) [world_drop]; Scout's Medallion (19537, -0.43 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 18.8 attack_power points (1.07 DPS) | yes | Barbaric Iron Shoulders (7913, -0.11 DPS) [crafted]; Golden Scale Shoulders (3841, -0.27 DPS) [crafted]; Mail Combat Spaulders (6404, -0.27 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.6 attack_power points (0.66 DPS) | yes | Wolfmaster Cape (6314, -0.09 DPS) [dungeon]; Wildhunter Cloak (16658, -0.09 DPS) [quest]; Slayer's Cape (14752, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.70 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.30 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.34 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.91 DPS) | yes | Yorgen Bracers (13012, -0.09 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop]; Bands of Serra'kis (6902, -0.23 DPS) [dungeon] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.25 DPS) | yes | Warsong Gauntlets (16978, -0.11 DPS) [quest]; The Frozen Clutch (23170, -0.11 DPS) [dungeon]; Mail Combat Gauntlets (4075, -0.16 DPS) [world_drop] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 24.8 attack_power points (1.41 DPS) | yes | Girdle of Golem Strength (9405, -0.04 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.04 DPS) [rep]; Defiler's Chain Girdle (20152, -0.04 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 27.6 attack_power points (1.57 DPS) | yes | Ferine Leggings (6690, -0.09 DPS) [dungeon]; Golden Scale Leggings (3843, -0.32 DPS) [crafted]; Slayer's Pants (14757, -0.32 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 19.6 attack_power points (1.11 DPS) | yes | Brawler's Leather Boots (252439, -0.32 DPS) [crafted]; Hard Gold Boots (250534, -0.34 DPS, sim-verified) [crafted]; Veteran's Boots (250503, -0.36 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 18.4 attack_power points (1.04 DPS) | yes | Ironspine's Eye (7686, -0.18 DPS) [dungeon]; Tiger Band (6749, -0.36 DPS) [quest]; Band of the Fist (17694, -0.41 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.8 attack_power points (0.95 DPS) | yes | Ironspine's Eye (7686, -0.09 DPS) [dungeon]; Tiger Band (6749, -0.27 DPS) [quest]; Band of the Fist (17694, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (19.62 DPS) | yes | Morbid Dawn (7689, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Ironspine's Fist (7687) | Scarlet Monastery: Ironspine [dungeon] | 334.9 attack_power points (19.02 DPS) | yes | Royal Diplomatic Scepter (9457, -0.80 DPS, sim-verified) [dungeon]; Slayer's Shield (15892, -18.21 DPS) [world_drop]; Shield of Thorsen (13079, -18.23 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.51 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; Long Battle Bow (15284, -0.17 DPS) [world_drop]; Precision Bow (8183, -0.25 DPS) [world_drop] |

**New at 30:** head: Tusken Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Ironspine's Fist; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 454, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (troll, 00000000000000000-05153105022011500-000000000000000000)

Set DPS (verified): 115.3. Weights run: 2.4s. Verify run: 1.3s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.615 ± 0.063, crit=2.307 ± 0.090 per rating point (14 rating = 1%, 32.304 per %), hit=2.575 ± 0.138 per rating point (10 rating = 1%, 25.751 per %), melee_haste=11.387 ± 1.289

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 58.3 attack_power points (3.61 DPS) | yes | Chromite Barbute (8142, -0.84 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.15 DPS) [crafted]; Barbaric Iron Helm (7915, -1.60 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.24 DPS) | yes | Scout's Medallion (19536, -0.14 DPS) [rep]; Ethereal Talisman (4430, -0.22 DPS) [quest]; Kaleidoscope Chain (13084, -0.34 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.8 attack_power points (1.84 DPS) | yes | Forest Tracker Epaulets (2278, -0.12 DPS) [world_drop]; Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.48 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 24.2 attack_power points (1.50 DPS) | yes | First Sergeant's Cloak (16340, -0.15 DPS) [pvp]; Hawkeye's Cloak (14593, -0.42 DPS) [world_drop]; Parachute Cloak (10518, -0.70 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 46.7 attack_power points (2.89 DPS) | yes | Kolkar Marauder Chain (6773, -0.52 DPS, sim-verified) [quest]; Carapace of Tuten'kash (10775, -0.85 DPS) [dungeon]; Jouster's Chestplate (8157, -1.03 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.5 attack_power points (1.27 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 52.3 attack_power points (3.24 DPS) | yes | Dragonscale Gauntlets (8347, -0.56 DPS, sim-verified) [crafted]; Scarlet Gauntlets (10331, -0.75 DPS) [dungeon]; Plated Fist of Hakoo (13071, -1.00 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 40.3 attack_power points (2.50 DPS) | yes | Ogron's Sash (13117, -0.55 DPS, sim-verified) [world_drop]; Boar Champion's Belt (10768, -0.64 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.64 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.60 DPS) | yes | Symbolic Legplates (14829, -0.14 DPS) [world_drop]; Firemane Leggings (13129, -0.25 DPS) [world_drop]; Triprunner Dungarees (9624, -0.43 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.3 attack_power points (2.31 DPS) | yes | Blackforge Greaves (6423, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.34 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.9 attack_power points (1.79 DPS) | yes | Thunderbrow Ring (13097, -0.50 DPS) [world_drop]; Falcon's Hook (7552, -0.52 DPS) [dungeon]; Ring of the Underwood (2951, -0.54 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.5 attack_power points (1.40 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Ring of the Underwood (2951, -0.15 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (115.3 DPS) | yes | Skullance Shield (13081, -26.48 DPS) [world_drop]; Savage Boar's Guard (10767, -26.51 DPS) [dungeon]; Ardent Custodian (868, -39.97 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.15 DPS) [quest]; The Silencer (13138, -0.18 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.06 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 191.6. Weights run: 2.5s. Verify run: 1.7s. 811 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.303 ± 0.057, crit=1.861 ± 0.081 per rating point (14 rating = 1%, 26.056 per %), hit=1.651 ± 0.141 per rating point (10 rating = 1%, 16.512 per %), melee_haste=11.121 ± 1.183

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 78.6 attack_power points (7.49 DPS) | yes | Blood Guard's Plate Helm (220803, -0.95 DPS) [vendor]; Embrace of the Lycan (9479, -2.91 DPS) [dungeon]; Raging Berserker's Helm (7719, -3.99 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.9 attack_power points (2.57 DPS) | yes | Woven Ivy Necklace (19159, -0.31 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.66 DPS) [quest]; Scout's Medallion (19535, -1.08 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 49.5 attack_power points (4.72 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -0.14 DPS) [vendor]; Wyrmslayer Spaulders (13066, -1.44 DPS) [world_drop]; Officer's Pauldrons (250576, -2.04 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.2 attack_power points (2.88 DPS) | yes | Dark Hooded Cape (5257, -0.88 DPS) [world]; First Sergeant's Cloak (16340, -0.99 DPS) [pvp]; Dark Phantom Cape (13122, -1.02 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 56.1 attack_power points (5.34 DPS) | yes | Mixologist's Tunic (12793, -0.55 DPS) [dungeon]; Warforged Chestplate (11195, -0.77 DPS) [quest]; Warbear Harness (15064, -1.01 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.8 attack_power points (2.84 DPS) | yes | Runed Golem Shackles (12550, +0.00 DPS) [dungeon]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 57.0 attack_power points (5.43 DPS) | yes | Raider Gloves (272100, -0.80 DPS) [vendor]; Gloves of Holy Might (867, -1.05 DPS) [world_drop]; Officer's Gloves (250551, -1.27 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193), Defiler's Plate Girdle (20205)) | The Defilers [rep] | 46.1 attack_power points (4.39 DPS) | yes | Defiler's Leather Girdle (20193, +0.00 DPS) [rep]; Defiler's Plate Girdle (20205, +0.00 DPS) [rep]; Girdle of Beastial Fury (11686, -0.01 DPS) [dungeon] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 63.1 attack_power points (6.01 DPS) | yes | Stormshroud Pants (15057, -1.05 DPS) [crafted]; Serpentskin Leggings (8262, -1.55 DPS) [world_drop]; Golem Shard Leggings (13074, -1.82 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths: Anvilrage Overseer [dungeon] | 44.9 attack_power points (4.28 DPS) | yes | Prowler's Leather Boots (252468, -0.44 DPS) [crafted]; Skulker's Leather Boots (252469, -0.57 DPS) [crafted]; Officer's Sabatons (250561, -0.69 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.5 attack_power points (3.48 DPS) | yes | White Bone Band (11862, -1.19 DPS) [quest]; Mark of Kern (2262, -1.57 DPS) [dungeon]; Assault Band (13095, -1.57 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.7 attack_power points (3.02 DPS) | yes | White Bone Band (11862, -0.74 DPS) [quest]; Mark of Kern (2262, -1.12 DPS) [dungeon]; Assault Band (13095, -1.12 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (191.6 DPS) | yes | Frozen Heart of the Mountain (249469, -3.69 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-verified (191.6 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -6.39 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (191.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hookfang Shanker (11635, -1.49 DPS, sim-verified) [dungeon] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (51.94 DPS) | yes | Dawn's Edge (12774, +0.00 DPS) [crafted]; Claw of Celebras (17738, -5.99 DPS) [dungeon]; White Bone Shredder (11863, -8.86 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (191.6 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.17 DPS) [world_drop]; Dark Iron Rifle (16004, -2.27 DPS, sim-verified) [crafted] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Defiler's Chain Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Diamond Flask; main_hand: Hammer of the Northern Wind; off_hand: Shadowblade; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 811, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 298.1. Weights run: 2.4s. Verify run: 1.7s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=2.316 ± 0.120, crit=3.309 ± 0.172 per rating point (14 rating = 1%, 46.322 per %), hit=3.538 ± 0.321 per rating point (10 rating = 1%, 35.377 per %), melee_haste=21.075 ± 2.823

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 199.4 attack_power points (20.38 DPS) | yes | Eye of Rend (12587, -6.75 DPS, sim-verified) [dungeon]; Champion's Plate Helm (227043, -7.74 DPS) [pvp]; Outlaw's Collar (279253, -8.31 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (298.1 DPS) | yes | Medallion of the Dawn (22659, -0.20 DPS) [quest]; Pendant of Celerity (22340, -0.23 DPS) [dungeon]; Rage of Mugamba (19577, -4.15 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 108.9 attack_power points (11.12 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -0.04 DPS) [pvp]; Truestrike Shoulders (12927, -1.44 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 63.4 attack_power points (6.48 DPS) | yes | Windshear Cape (20691, -1.29 DPS) [world]; Cape of the Black Baron (13340, -1.60 DPS, sim-verified) [dungeon]; Fel Cape (279269, -1.74 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (298.1 DPS) | yes | Timbermaw Tunic (252484, -1.37 DPS) [crafted]; Savage Gladiator Chain (11726, -3.42 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -13.03 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (298.1 DPS) | yes | Forest Stalker's Bracers (19587, -1.60 DPS) [rep]; Vambraces of the Sadist (13400, -2.18 DPS) [dungeon]; Bracers of Undead Slaying (23090, -6.77 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-verified (298.1 DPS) | yes | General's Plate Gauntlets (231532, -0.54 DPS) [pvp]; Raider Gloves (272099, -0.72 DPS) [vendor]; Razor Gauntlets (18326, -4.42 DPS, sim-verified) [dungeon] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 100.3 attack_power points (10.25 DPS) | yes | Ferocity of the Timbermaw (227805, -0.91 DPS) [vendor]; Marksman's Girdle (22232, -1.67 DPS) [dungeon]; Belt of Preserved Heads (20216, -1.86 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (298.1 DPS) | yes | Sentinel's Plate Legguards (237825, -2.77 DPS) [vendor]; Titanic Leggings (22385, -3.27 DPS) [crafted]; Cloudkeeper Legplates (14554, -3.75 DPS, sim-verified) [world_drop] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 81.7 attack_power points (8.35 DPS) | yes | Boots of Heroism (21995, +0.00 DPS, sim-verified) [quest]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Battleboots of Heroism (226857, -0.65 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (298.1 DPS) | yes | Tarnished Elven Ring (18500, -2.82 DPS) [dungeon]; Cutthroat's Signet (272408, -3.06 DPS) [vendor]; Naglering (11669, -9.05 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (298.1 DPS) | yes | Tarnished Elven Ring (18500, -0.71 DPS) [dungeon]; Cutthroat's Signet (272408, -0.95 DPS) [vendor]; Naglering (11669, -6.57 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (298.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Diamond Flask (20130, -4.69 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (298.1 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Diamond Flask (20130, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.57 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (298.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Felstriker (12590, -1.94 DPS, sim-verified) [dungeon] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (298.1 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -27.66 DPS) [dungeon]; Skullflame Shield (1168, -101.74 DPS, sim-verified) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (298.1 DPS) | yes | Satyr's Bow (18323, -0.41 DPS) [dungeon]; Blackcrow (12651, -0.77 DPS) [dungeon]; Dark Iron Rifle (16004, -3.95 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 784.3. Weights run: 2.8s. Verify run: 1.9s. 1913 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.950 ± 0.139, crit=2.785 ± 0.198 per rating point (14 rating = 1%, 38.995 per %), hit=3.817 ± 0.353 per rating point (10 rating = 1%, 38.174 per %), melee_haste=33.827 ± 3.794

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 190.3 attack_power points (44.19 DPS) | yes | Champion's Plate Helm (227043, -16.52 DPS) [pvp]; Fury Visor (20521, -17.92 DPS) [quest]; Mask of the Unforgiven (13404, -19.92 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (784.3 DPS) | yes | Mark of Fordring (15411, -0.56 DPS) [quest]; Medallion of the Dawn (22659, -1.03 DPS) [quest]; Rage of Mugamba (19577, -7.67 DPS, sim-verified) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 100.3 attack_power points (23.30 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, -2.02 DPS) [vendor]; Darkspear Pauldrons (272105, -6.37 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.2 attack_power points (15.36 DPS) | yes | Windshear Cape (20691, -4.86 DPS) [world]; Deathguard's Cloak (20068, -5.21 DPS) [rep]; Cape of the Black Baron (13340, -5.97 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (784.3 DPS) | yes | Timbermaw Tunic (252484, -2.14 DPS) [crafted]; Savage Gladiator Chain (11726, -7.32 DPS) [dungeon]; Breastplate of Undead Slaying (23087, -29.01 DPS, sim-verified) [world] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (784.3 DPS) | yes | Forest Stalker's Bracers (19587, -4.21 DPS) [rep]; Berserker Bracers (19578, -5.47 DPS) [rep]; Bracers of Undead Slaying (23090, -16.03 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | sim-verified (784.3 DPS) | yes | General's Plate Gauntlets (231532, -3.46 DPS) [pvp]; Stormshroud Gloves (21278, -3.88 DPS) [crafted]; Razor Gauntlets (18326, -11.37 DPS, sim-verified) [dungeon] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 95.4 attack_power points (22.15 DPS) | yes | Radiant Girdle of the Dawn (227814, -0.56 DPS) [vendor]; Ferocity of the Timbermaw (227805, -2.39 DPS) [vendor]; Marksman's Girdle (22232, -3.78 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (784.3 DPS) | yes | Titanic Leggings (22385, -2.10 DPS) [crafted]; Sentinel's Plate Legguards (237825, -3.31 DPS) [vendor]; Cloudkeeper Legplates (14554, -9.48 DPS, sim-verified) [world_drop] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 78.2 attack_power points (18.15 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Windreaver Greaves (13967, -0.23 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (784.3 DPS) | yes | Tarnished Elven Ring (18500, -5.98 DPS) [dungeon]; Cutthroat's Signet (272408, -6.43 DPS) [vendor]; Naglering (11669, -20.82 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (784.3 DPS) | yes | Tarnished Elven Ring (18500, -1.36 DPS) [dungeon]; Cutthroat's Signet (272408, -1.81 DPS) [vendor]; Naglering (11669, -15.69 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (784.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (784.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Diamond Flask (20130, +0.00 DPS) [quest] |
| main_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | sim-verified (784.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Quel'Serrar (18348, -23.21 DPS, sim-verified) [quest] |
| off_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (784.3 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -61.94 DPS) [dungeon]; Skullflame Shield (1168, -218.39 DPS, sim-verified) [world_drop] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (784.3 DPS) | yes | Blackcrow (12651, -0.89 DPS) [dungeon]; The Purifier (22656, -1.17 DPS) [quest]; Dark Iron Rifle (16004, -10.90 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Boots of Heroism; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Ravencrest's Legacy; off_hand: Shadowsong's Sorrow; ranged: Satyr's Bow

No-known-source sample (15 of 1913, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

