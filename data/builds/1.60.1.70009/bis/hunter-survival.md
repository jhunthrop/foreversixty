# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 68.5. Weights run: 3.6s. Verify run: 4.7s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.479 ± 0.008, strength=1.210 ± 0.002, crit=0.833 ± 0.016 per rating point (14 rating = 1%, 11.663 per %), hit=1.179 ± 0.043 per rating point (10 rating = 1%, 11.791 per %), melee_haste=not significant (1.128 ± 1.063)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 11.8 attack_power points (0.50 DPS) | yes | Defender's Leather Hood (252447, -0.09 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.9 attack_power points (0.37 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.4 attack_power points (0.31 DPS) | yes | Slime-encrusted Pads (6461, -0.33 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 8.9 attack_power points (0.37 DPS) | yes | Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.08 DPS) [crafted]; Grave Shroud (279865, -0.10 DPS) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 16.4 attack_power points (0.69 DPS) | yes | Tunic of Westfall (2041, -0.01 DPS) [quest]; Defender's Leather Armor (252434, -0.15 DPS) [crafted]; Prospector's Chestpiece (14562, -0.21 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 8.3 attack_power points (0.35 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Wolf Bracers (4794, -0.10 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.5 DPS) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.21 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.75 DPS) | yes | Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Dusty Belt (279897, -0.44 DPS) [quest]; Deviate Scale Belt (6468, -2.82 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.29 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (68.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.19 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.8 attack_power points (0.45 DPS) | yes | Demon Band (12054, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; The 1 Ring (8350, -0.34 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.9 attack_power points (0.37 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.43 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 235.5 attack_power points (9.86 DPS) | yes | Cruel Barb (5191, -0.77 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.9 attack_power points (0.25 DPS) | yes | Fine Longbow (11304, -0.08 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 95.3. Weights run: 3.9s. Verify run: 2.7s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.506 ± 0.008, strength=1.210 ± 0.002, crit=0.943 ± 0.017 per rating point (14 rating = 1%, 13.195 per %), hit=1.456 ± 0.052 per rating point (10 rating = 1%, 14.559 per %), melee_haste=7.018 ± 1.198

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.1 attack_power points (0.76 DPS) | yes | Defender's Leather Helm (252455, -0.03 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS) [world]; Brawler's Leather Hood (252504, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.70 DPS) | yes | Sentinel's Medallion (19541, -0.10 DPS) [rep]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.6 attack_power points (1.14 DPS) | yes | Mantle of Thieves (2264, -0.27 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.45 DPS) [crafted]; Bristlebark Amice (14573, -0.50 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.2 attack_power points (0.71 DPS) | yes | Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.17 DPS) [pvp]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.1 attack_power points (1.06 DPS) | yes | Brawler's Leather Tunic (252508, -0.09 DPS) [crafted]; Brawler's Leather Armor (252490, -0.23 DPS) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (0.70 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.15 DPS) [crafted]; Cultist's Armguards (270032, -0.19 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.5 attack_power points (0.88 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.21 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted]; Prowler's Leather Belt (252459, -0.20 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.31 DPS) | yes | Petrolspill Leggings (9509, -0.25 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.25 DPS) [world_drop]; Brawler's Leather Legguards (252516, -1.11 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.7 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.4 attack_power points (0.92 DPS) | yes | Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Monkey Ring (6748, -0.39 DPS) [quest]; Ring of Precision (1491, -0.47 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.3 attack_power points (0.82 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Monkey Ring (6748, -0.29 DPS) [quest]; Ring of Precision (1491, -0.36 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.16 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (95.3 DPS) | yes | Shoni's Disarming Tool (9608, -4.79 DPS) [quest]; Satyr's Rod (15962, -16.01 DPS) [world_drop]; Swinetusk Shank (6691, -20.17 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.45 DPS) | yes | Double-barreled Shotgun (2098, -0.04 DPS) [world_drop]; Silver Star (3463, -0.07 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.15 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 140.7. Weights run: 4.1s. Verify run: 3.3s. 597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.638 ± 0.008, strength=1.210 ± 0.001, crit=0.899 ± 0.015 per rating point (14 rating = 1%, 12.582 per %), hit=1.491 ± 0.081 per rating point (10 rating = 1%, 14.909 per %), melee_haste=8.343 ± 1.157

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 31.3 attack_power points (2.09 DPS) | yes | Raging Berserker's Helm (7719, -0.20 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.38 DPS) [crafted]; Hawkeye's Helm (14591, -0.56 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.33 DPS) | yes | Sentinel's Medallion (19540, -0.13 DPS) [rep]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.0 attack_power points (2.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.51 DPS) [dungeon]; Nightscape Shoulders (8192, -0.80 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 21.2 attack_power points (1.41 DPS) | yes | Sergeant Major's Cape (16336, -0.28 DPS) [pvp]; Hawkeye's Cloak (14593, -0.41 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.52 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 40.8 attack_power points (2.72 DPS) | yes | Kolkar Marauder Chain (6773, -0.80 DPS, sim-verified) [quest]; Wolffear Harness (13110, -0.86 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -1.08 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.33 DPS) | yes | Ravager's Armguards (14770, -0.33 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.36 DPS) [world_drop]; Dusky Bracers (7378, -0.46 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 32.6 attack_power points (2.17 DPS) | yes | Gauntlets of Divinity (7724, -0.04 DPS) [dungeon]; Scarlet Gauntlets (10331, -0.11 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.35 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.00 DPS) | yes | Ogron's Sash (13117, -0.29 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.40 DPS) [rep]; Skulker's Leather Belt (252520, -0.53 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.4 attack_power points (2.29 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Ferine Leggings (6690, -0.56 DPS) [dungeon]; Scarlet Leggings (10330, -0.60 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 27.3 attack_power points (1.82 DPS) | yes | Skulker's Leather Shoes (252531, -0.05 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.17 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 22.8 attack_power points (1.52 DPS) | yes | Assault Band (13095, -0.19 DPS) [world_drop]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Ring of the Underwood (2951, -0.27 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.33 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.03 DPS) [dungeon]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (140.7 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.30 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 445.6 attack_power points (29.67 DPS) | yes | Vanquisher's Sword (10823, +0.00 DPS) [quest]; Shoni's Disarming Tool (9608, -14.71 DPS) [quest]; Stonecloth Branch (15963, -29.43 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (140.7 DPS) | yes | Monolithic Bow (9426, -0.12 DPS) [dungeon]; Swiftwind (13038, -0.17 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.98 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 184.5. Weights run: 4.1s. Verify run: 3.8s. 754 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.681 ± 0.009, strength=1.210 ± 0.001, crit=1.016 ± 0.017 per rating point (14 rating = 1%, 14.221 per %), hit=1.478 ± 0.086 per rating point (10 rating = 1%, 14.777 per %), melee_haste=8.729 ± 1.242

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloomsprout Headpiece (17767, -0.38 DPS) [dungeon]; White Bandit Mask (10008, -0.69 DPS) [crafted]; Embrace of the Lycan (9479, -1.95 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.9 attack_power points (2.06 DPS) | yes | Sentinel's Medallion (19539, -0.57 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.58 DPS) [quest] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 36.1 attack_power points (2.67 DPS) | yes | Sunburn Spaulders (274751, -0.41 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.49 DPS) [crafted]; Failed Flying Experiment (9647, -0.53 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.8 attack_power points (2.28 DPS) | yes | Blisterbane Wrap (12552, -0.41 DPS) [dungeon]; Dark Phantom Cape (13122, -0.41 DPS) [world_drop]; Dark Hooded Cape (5257, -0.68 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 43.6 attack_power points (3.22 DPS) | yes | Blazewind Breastplate (11193, -0.09 DPS) [quest]; Fungus Shroud Armor (17742, -0.11 DPS) [dungeon]; Quillward Harness (10583, -0.14 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 30.1 attack_power points (2.22 DPS) | yes | Bracers of the Stone Princess (17714, -0.15 DPS) [dungeon]; Arena Bands (18711, -0.15 DPS) [world]; Skulker's Leather Bracers (252540, -0.48 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 49.1 attack_power points (3.63 DPS) | yes | Gloves of Holy Might (867, -1.10 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.21 DPS, sim-verified) [crafted]; Prowler's Leather Gauntlets (252547, -1.26 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 39.7 attack_power points (2.93 DPS) | yes | Substandard Belt Chain (274757, -0.11 DPS) [vendor]; Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.19 DPS) [crafted] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 43.4 attack_power points (3.21 DPS) | yes | Serpentskin Leggings (8262, -0.05 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.41 DPS) [vendor]; Basilisk Hide Pants (1718, -0.60 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 35.8 attack_power points (2.65 DPS) | yes | Skulker's Leather Boots (252469, -0.05 DPS) [crafted]; Prowler's Leather Boots (252468, -0.12 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.8 attack_power points (2.57 DPS) | yes | Masons Fraternity Ring (9533, -0.83 DPS) [quest]; Mark of Kern (2262, -1.09 DPS) [dungeon]; Assault Band (13095, -1.09 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 27.2 attack_power points (2.01 DPS) | yes | Masons Fraternity Ring (9533, -0.27 DPS) [quest]; Mark of Kern (2262, -0.53 DPS) [dungeon]; Assault Band (13095, -0.53 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (+31.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.27 DPS) [dungeon]; Shoni's Disarming Tool (9608, -22.31 DPS) [quest]; Grizzle's Skinner (11702, -31.20 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.44 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.44 DPS) [world_drop]; Dark Iron Rifle (16004, -1.46 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 301.3. Weights run: 3.9s. Verify run: 13.0s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.154 ± 0.025, strength=1.210 ± 0.002, crit=2.357 ± 0.049 per rating point (14 rating = 1%, 32.993 per %), hit=3.225 ± 0.253 per rating point (10 rating = 1%, 32.252 per %), melee_haste=not significant (11.626 ± 3.665)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -7.63 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.6 attack_power points (5.16 DPS) | yes | Medallion of the Dawn (22659, -0.60 DPS) [quest]; Beads of Ogre Might (22150, -0.66 DPS) [quest]; Mark of Fordring (15411, -2.52 DPS, sim-verified) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.1 attack_power points (7.28 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.21 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (4.81 DPS) | yes | Cloak of the Honor Guard (20073, -1.24 DPS) [rep]; Windshear Cape (20691, -1.46 DPS) [world]; Cape of the Black Baron (13340, -1.98 DPS, sim-verified) [dungeon] |
| chest | Beastmaster's Tunic (226886) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS) [vendor]; Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Dawn Armor (252483, -5.28 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -6.10 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Beaststalker's Gloves (16676, -3.88 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -4.62 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 173.6 attack_power points (13.87 DPS) | yes | Marshal's Chain Legplates (231558, -3.69 DPS) [vendor]; Sentinel's Leather Pants (237818, -3.95 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -4.41 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.71 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (301.3 DPS) | yes | Tarnished Elven Ring (18500, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.50 DPS) [vendor]; The Postmaster's Seal (13392, -3.04 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.52 DPS) [dungeon]; Cutthroat's Signet (272408, -0.69 DPS) [vendor]; Naglering (11669, -7.18 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Burst of Knowledge (11832, -9.98 DPS, sim-verified) [dungeon] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -14.55 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.4 attack_power points (66.97 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -6.50 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -20.69 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.46 DPS) [quest]; Dark Iron Rifle (16004, -3.55 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Beastmaster's Tunic; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 833.0. Weights run: 3.3s. Verify run: 11.9s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.364 ± 0.026, strength=1.210 ± 0.002, crit=2.937 ± 0.054 per rating point (14 rating = 1%, 41.125 per %), hit=4.751 ± 0.460 per rating point (10 rating = 1%, 47.513 per %), melee_haste=not significant (14.618 ± 5.778)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (833.0 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -9.46 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 83.0 attack_power points (14.47 DPS) | yes | Beads of Ogre Might (22150, -2.00 DPS) [quest]; Mark of Fordring (15411, -2.76 DPS) [quest]; Medallion of the Dawn (22659, -3.11 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 119.0 attack_power points (20.75 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS, sim-verified) [vendor]; Field Marshal's Chain Spaulders (16468, -1.75 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.5 attack_power points (13.16 DPS) | yes | Cape of the Black Baron (13340, -3.50 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.88 DPS) [vendor]; Stalwart Cloak (272415, -4.88 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (833.0 DPS) | yes | Field Marshal's Chain Armor (231563, -6.45 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -6.86 DPS) [vendor]; Tunic of Undead Slaying (23089, -28.39 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (833.0 DPS) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Slashclaw Bracers (13211, -5.79 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 116.3 attack_power points (20.28 DPS) | yes | Marshal's Chain Grips (231560, -3.04 DPS) [pvp]; Marshal's Chain Vices (231578, -4.45 DPS) [vendor]; Stormshroud Gloves (21278, -4.82 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 99.9 attack_power points (17.42 DPS) | yes | Marksman's Girdle (22232, -0.48 DPS) [dungeon]; Highlander's Chain Girdle (20043, -4.32 DPS) [rep]; Highlander's Leather Girdle (20045, -4.32 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 212.5 attack_power points (37.05 DPS) | yes | Sentinel's Leather Pants (237818, -11.58 DPS) [vendor]; Marshal's Chain Legplates (231558, -12.80 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -13.20 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (833.0 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -4.30 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (833.0 DPS) | yes | Tarnished Elven Ring (18500, -3.78 DPS) [dungeon]; Cutthroat's Signet (272408, -4.19 DPS) [vendor]; Naglering (11669, -14.83 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (833.0 DPS) | yes | Tarnished Elven Ring (18500, -1.24 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Naglering (11669, -13.01 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (833.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (833.0 DPS) | yes | Frozen Heart of the Mountain (249469, -6.88 DPS) [crafted]; Counterattack Lodestone (18537, -10.50 DPS) [dungeon]; Hand of Justice (11815, -10.85 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (833.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -75.45 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 840.3 attack_power points (146.49 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -13.07 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -45.48 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (833.0 DPS) | yes | Blackcrow (12651, -0.83 DPS) [dungeon]; The Purifier (22656, -2.35 DPS) [quest]; Dark Iron Rifle (16004, -6.48 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 68.7. Weights run: 3.6s. Verify run: 4.5s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.479 ± 0.008, strength=1.210 ± 0.002, crit=0.833 ± 0.016 per rating point (14 rating = 1%, 11.663 per %), hit=1.179 ± 0.043 per rating point (10 rating = 1%, 11.791 per %), melee_haste=not significant (1.128 ± 1.063)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 11.8 attack_power points (0.50 DPS) | yes | Defender's Leather Hood (252447, -0.09 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.9 attack_power points (0.37 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.4 attack_power points (0.31 DPS) | yes | Slime-encrusted Pads (6461, -0.31 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 8.9 attack_power points (0.37 DPS) | yes | Cape of the Brotherhood (5193, -0.06 DPS) [dungeon]; Dark Leather Cloak (2316, -0.08 DPS) [crafted]; Grave Shroud (279865, -0.10 DPS) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 16.4 attack_power points (0.69 DPS) | yes | Defender's Leather Armor (252434, -0.15 DPS) [crafted]; Prospector's Chestpiece (14562, -0.21 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.25 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 7.4 attack_power points (0.31 DPS) | yes | Bristlebark Bindings (14569, -0.02 DPS) [world_drop]; Wolf Bracers (4794, -0.06 DPS) [vendor]; Ratchet Wristwraps (274742, -0.12 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.7 DPS) | yes | Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.13 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.75 DPS) | yes | Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Ruffian Belt (5975, -0.45 DPS) [world]; Deviate Scale Belt (6468, -2.75 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.22 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (68.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.11 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.8 attack_power points (0.45 DPS) | yes | Demon Band (12054, -0.25 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.26 DPS) [quest]; Loop of Sacrifice (281673, -0.30 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.9 attack_power points (0.37 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.19 DPS) [quest]; Loop of Sacrifice (281673, -0.22 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.43 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 235.5 attack_power points (9.86 DPS) | yes | Cruel Barb (5191, -0.62 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -9.76 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.9 attack_power points (0.25 DPS) | yes | Fine Longbow (11304, -0.08 DPS) [vendor]; Deadly Blunderbuss (4369, -0.12 DPS) [crafted]; Light Bow (4576, -0.12 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 96.0. Weights run: 3.9s. Verify run: 2.7s. 352 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.506 ± 0.008, strength=1.210 ± 0.002, crit=0.943 ± 0.017 per rating point (14 rating = 1%, 13.195 per %), hit=1.456 ± 0.052 per rating point (10 rating = 1%, 14.559 per %), melee_haste=7.018 ± 1.198

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.1 attack_power points (0.76 DPS) | yes | Defender's Leather Helm (252455, -0.03 DPS) [crafted]; Tribal Worg Helm (6204, -0.15 DPS) [world]; Brawler's Leather Hood (252504, -0.15 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.70 DPS) | yes | Scout's Medallion (19537, -0.10 DPS) [rep]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.6 attack_power points (1.14 DPS) | yes | Mantle of Thieves (2264, -0.22 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.45 DPS) [crafted]; Bristlebark Amice (14573, -0.50 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.2 attack_power points (0.71 DPS) | yes | Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon]; Wildhunter Cloak (16658, -0.21 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.1 attack_power points (1.06 DPS) | yes | Brawler's Leather Tunic (252508, -0.09 DPS) [crafted]; Brawler's Leather Armor (252490, -0.23 DPS) [crafted]; Defender's Leather Tunic (252450, -0.24 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (0.70 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.15 DPS) [crafted]; Cultist's Armguards (270032, -0.19 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.5 attack_power points (0.88 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.21 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted]; Prowler's Leather Belt (252459, -0.20 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.31 DPS) | yes | Petrolspill Leggings (9509, -0.25 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.25 DPS) [world_drop]; Brawler's Leather Legguards (252516, -1.20 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.7 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.4 attack_power points (0.92 DPS) | yes | Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Band of the Fist (17694, -0.38 DPS) [quest]; Monkey Ring (6748, -0.39 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.3 attack_power points (0.82 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Band of the Fist (17694, -0.27 DPS) [quest]; Monkey Ring (6748, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.16 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (96.0 DPS) | yes | Tork Wrench (11855, -15.96 DPS) [quest]; Satyr's Rod (15962, -16.01 DPS) [world_drop]; Swinetusk Shank (6691, -20.37 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.45 DPS) | yes | Double-barreled Shotgun (2098, -0.04 DPS) [world_drop]; Silver Star (3463, -0.07 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.15 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 141.2. Weights run: 4.1s. Verify run: 3.2s. 563 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.638 ± 0.008, strength=1.210 ± 0.001, crit=0.899 ± 0.015 per rating point (14 rating = 1%, 12.582 per %), hit=1.491 ± 0.081 per rating point (10 rating = 1%, 14.909 per %), melee_haste=8.343 ± 1.157

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 31.3 attack_power points (2.09 DPS) | yes | Raging Berserker's Helm (7719, -0.20 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.38 DPS) [crafted]; Hawkeye's Helm (14591, -0.56 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.33 DPS) | yes | Scout's Medallion (19536, -0.13 DPS) [rep]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.0 attack_power points (2.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.51 DPS) [dungeon]; Nightscape Shoulders (8192, -0.80 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 21.2 attack_power points (1.41 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.41 DPS) [world_drop]; Parachute Cloak (10518, -0.54 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 40.8 attack_power points (2.72 DPS) | yes | Wolffear Harness (13110, -0.86 DPS) [world_drop]; Kolkar Marauder Chain (6773, -0.89 DPS, sim-verified) [quest]; Tough Scorpid Breastplate (8203, -1.08 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.33 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.33 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 32.6 attack_power points (2.17 DPS) | yes | Scarlet Gauntlets (10331, -0.11 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.35 DPS) [crafted]; Gauntlets of Divinity (7724, -0.49 DPS, sim-verified) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.00 DPS) | yes | Ogron's Sash (13117, -0.29 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.40 DPS) [rep]; Skulker's Leather Belt (252520, -0.53 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.4 attack_power points (2.29 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Ferine Leggings (6690, -0.56 DPS) [dungeon]; Scarlet Leggings (10330, -0.60 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 27.3 attack_power points (1.82 DPS) | yes | Skulker's Leather Shoes (252531, -0.05 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.17 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 22.8 attack_power points (1.52 DPS) | yes | Assault Band (13095, -0.19 DPS) [world_drop]; Ironspine's Eye (7686, -0.21 DPS) [dungeon]; Ring of the Underwood (2951, -0.27 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.33 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.03 DPS) [dungeon]; Ring of the Underwood (2951, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (141.2 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.25 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 445.6 attack_power points (29.67 DPS) | yes | Vanquisher's Sword (10823, +0.00 DPS) [quest]; Stonecloth Branch (15963, -29.43 DPS) [world_drop]; Tork Wrench (11855, -29.51 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (141.2 DPS) | yes | Monolithic Bow (9426, -0.12 DPS) [dungeon]; Swiftwind (13038, -0.17 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.99 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 186.6. Weights run: 4.1s. Verify run: 3.6s. 713 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.681 ± 0.009, strength=1.210 ± 0.001, crit=1.016 ± 0.017 per rating point (14 rating = 1%, 14.221 per %), hit=1.478 ± 0.086 per rating point (10 rating = 1%, 14.777 per %), melee_haste=8.729 ± 1.242

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 41.7 attack_power points (3.08 DPS) | yes | Blood Guard's Chain Helmet (220821, -0.04 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.42 DPS) [dungeon]; White Bandit Mask (10008, -0.73 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.9 attack_power points (2.06 DPS) | yes | Woven Ivy Necklace (19159, -0.41 DPS) [quest]; Scout's Medallion (19535, -0.57 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.58 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.5 attack_power points (2.25 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Failed Flying Experiment (9647, -0.12 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 30.8 attack_power points (2.28 DPS) | yes | Blisterbane Wrap (12552, -0.41 DPS) [dungeon]; Dark Phantom Cape (13122, -0.41 DPS) [world_drop]; Dark Hooded Cape (5257, -0.68 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 43.6 attack_power points (3.22 DPS) | yes | Blazewind Breastplate (11193, -0.09 DPS) [quest]; Fungus Shroud Armor (17742, -0.11 DPS) [dungeon]; Quillward Harness (10583, -0.14 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 30.1 attack_power points (2.22 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 49.1 attack_power points (3.63 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.82 DPS, sim-verified) [crafted]; Gloves of Holy Might (867, -1.10 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -1.26 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 39.7 attack_power points (2.93 DPS) | yes | Substandard Belt Chain (274757, -0.11 DPS) [vendor]; Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.19 DPS) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Basilisk Hide Pants (1718, -0.18 DPS) [world_drop]; Triprunner Dungarees (9624, -0.29 DPS) [quest]; Serpentskin Leggings (8262, -2.17 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 35.8 attack_power points (2.65 DPS) | yes | Skulker's Leather Boots (252469, -0.05 DPS) [crafted]; Prowler's Leather Boots (252468, -0.12 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.8 attack_power points (2.57 DPS) | yes | White Bone Band (11862, -0.80 DPS) [quest]; Masons Fraternity Ring (9533, -0.83 DPS) [quest]; Mark of Kern (2262, -1.09 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 27.2 attack_power points (2.01 DPS) | yes | White Bone Band (11862, -0.24 DPS) [quest]; Masons Fraternity Ring (9533, -0.27 DPS) [quest]; Mark of Kern (2262, -0.53 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.89 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Devilsaur Tooth (19992, -2.05 DPS, sim-verified) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (+31.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.27 DPS) [dungeon]; White Bone Shredder (11863, -5.53 DPS) [quest]; Grizzle's Skinner (11702, -31.64 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.44 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.44 DPS) [world_drop]; Dark Iron Rifle (16004, -1.30 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Chain Legplates; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Shadowblade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 303.3. Weights run: 3.9s. Verify run: 12.5s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.154 ± 0.025, strength=1.210 ± 0.002, crit=2.357 ± 0.049 per rating point (14 rating = 1%, 32.993 per %), hit=3.225 ± 0.253 per rating point (10 rating = 1%, 32.252 per %), melee_haste=not significant (11.626 ± 3.665)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (303.3 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -5.21 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.6 attack_power points (5.16 DPS) | yes | Mark of Fordring (15411, -0.44 DPS) [quest]; Medallion of the Dawn (22659, -0.60 DPS) [quest]; Beads of Ogre Might (22150, -0.66 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 91.1 attack_power points (7.28 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.21 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 60.3 attack_power points (4.81 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.24 DPS) [rep]; Windshear Cape (20691, -1.46 DPS) [world] |
| chest | Beastmaster's Tunic (226886) | Saving the Best for Last [quest] | sim-verified (303.3 DPS) | yes | Legionnaire's Chain Armor (227083, +0.00 DPS) [vendor]; Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Dawn Armor (252483, -3.66 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (303.3 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -5.70 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (303.3 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -4.05 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (303.3 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -4.12 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 173.6 attack_power points (13.87 DPS) | yes | Outrider's Chain Leggings (22673, -2.35 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -3.69 DPS) [vendor]; Sentinel's Leather Pants (237818, -3.95 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (303.3 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.49 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (303.3 DPS) | yes | Tarnished Elven Ring (18500, -1.33 DPS) [dungeon]; Cutthroat's Signet (272408, -1.50 DPS) [vendor]; Naglering (11669, -5.75 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (303.3 DPS) | yes | Tarnished Elven Ring (18500, -0.52 DPS) [dungeon]; Cutthroat's Signet (272408, -0.69 DPS) [vendor]; Naglering (11669, -4.27 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (303.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (303.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (303.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -14.08 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.4 attack_power points (66.97 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -6.73 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -20.69 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (303.3 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.46 DPS) [quest] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Beastmaster's Tunic; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 834.2. Weights run: 3.3s. Verify run: 11.3s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.364 ± 0.026, strength=1.210 ± 0.002, crit=2.937 ± 0.054 per rating point (14 rating = 1%, 41.125 per %), hit=4.751 ± 0.460 per rating point (10 rating = 1%, 47.513 per %), melee_haste=not significant (14.618 ± 5.778)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (834.2 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -7.75 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 83.0 attack_power points (14.47 DPS) | yes | Beads of Ogre Might (22150, -2.00 DPS) [quest]; Mark of Fordring (15411, -2.76 DPS) [quest]; Medallion of the Dawn (22659, -3.11 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 119.0 attack_power points (20.75 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS, sim-verified) [vendor]; Warlord's Chain Shoulders (231572, -1.75 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 75.5 attack_power points (13.16 DPS) | yes | Cape of the Black Baron (13340, -3.50 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.88 DPS) [vendor]; Stalwart Cloak (272415, -4.88 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (834.2 DPS) | yes | Warlord's Chain Armor (231566, -6.45 DPS) [vendor]; Legionnaire's Chain Armor (227083, -6.86 DPS) [vendor]; Tunic of Undead Slaying (23089, -26.01 DPS, sim-verified) [world] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (834.2 DPS) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Slashclaw Bracers (13211, -5.11 DPS, sim-verified) [dungeon] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 116.3 attack_power points (20.28 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; General's Chain Grips (231569, -3.04 DPS) [vendor]; General's Chain Gloves (16571, -4.45 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 99.9 attack_power points (17.42 DPS) | yes | Marksman's Girdle (22232, -0.48 DPS) [dungeon]; Defiler's Chain Girdle (20150, -4.32 DPS) [rep]; Defiler's Leather Girdle (20190, -4.32 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 212.5 attack_power points (37.05 DPS) | yes | Outrider's Chain Leggings (22673, -5.23 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -11.58 DPS) [vendor]; General's Chain Legplates (231567, -12.80 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 94.8 attack_power points (16.53 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -2.64 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (834.2 DPS) | yes | Tarnished Elven Ring (18500, -3.78 DPS) [dungeon]; Cutthroat's Signet (272408, -4.19 DPS) [vendor]; Naglering (11669, -10.30 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (834.2 DPS) | yes | Tarnished Elven Ring (18500, -1.24 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Naglering (11669, -11.88 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (834.2 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -5.20 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (834.2 DPS) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.67 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (834.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -73.80 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 840.3 attack_power points (146.49 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -12.63 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -45.48 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (834.2 DPS) | yes | Blackcrow (12651, -0.83 DPS) [dungeon]; The Purifier (22656, -2.35 DPS) [quest]; Dark Iron Rifle (16004, -3.82 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

