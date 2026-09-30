# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 57.3. Weights run: 1.4s. Verify run: 1.3s. 218 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.041 ± 0.008, strength=1.000 ± 0.001, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.602 per %), hit=0.039 ± 0.002 per rating point (10 rating = 1%, 0.386 per %), melee_haste=not significant (0.818 ± 0.763)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.34 DPS) | yes | Defender's Leather Hood (252447, +0.00 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.21 DPS) | yes | Forest Leather Mantle (4709, -0.25 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.25 DPS) | yes | Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Catacomb Cloak (279899, -0.17 DPS, sim-verified) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 attack_power points (0.50 DPS) | yes | Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.16 DPS) [world_drop]; Tunic of Westfall (2041, -0.19 DPS, sim-verified) [quest] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.2 attack_power points (0.25 DPS) | yes | Bristlebark Bindings (14569, -0.04 DPS) [world_drop]; Wolf Bracers (4794, -0.08 DPS) [vendor]; Forest Leather Bracers (3202, -0.22 DPS, sim-verified) [world_drop] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.41 DPS) | yes | Brawler's Leather Gloves (252494, -0.08 DPS) [crafted]; Gold-flecked Gloves (5195, -0.13 DPS) [dungeon]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.73 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.48 DPS) [world]; Deviate Scale Belt (6468, -0.55 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.4 attack_power points (0.62 DPS) | yes | Trapper's Leather Pants (252501, +0.43 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.46 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.2 attack_power points (0.33 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 attack_power points (0.25 DPS) | yes | Demon Band (12054, -0.15 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.05 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 attack_power points (9.43 DPS) | yes | Cruel Barb (5191, +0.26 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.08 DPS) [crafted]; Light Bow (4576, -0.08 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 30 (dwarf, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 77.4. Weights run: 1.5s. Verify run: 1.4s. 364 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.050 ± 0.013, strength=1.000 ± 0.001, crit=0.193 ± 0.007 per rating point (14 rating = 1%, 2.706 per %), hit=0.059 ± 0.002 per rating point (10 rating = 1%, 0.586 per %), melee_haste=not significant (0.559 ± 0.829)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.54 DPS) | yes | Brawler's Leather Helm (252512, +0.00 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.63 DPS) | yes | Sentinel's Medallion (19541, -0.26 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Sentinel's Medallion (20444, -0.35 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 attack_power points (0.75 DPS) | yes | Mantle of Thieves (2264, -0.18 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Bristlebark Amice (14573, -0.33 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.4 attack_power points (0.47 DPS) | yes | Wolfmaster Cape (6314, -0.05 DPS, sim-verified) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.73 DPS) | yes | Dusky Leather Armor (7374, +0.00 DPS, sim-verified) [crafted]; Brawler's Leather Tunic (252508, -0.07 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 attack_power points (0.47 DPS) | yes | Cultist's Armguards (270032, +0.00 DPS, sim-verified) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.73 DPS) | yes | Insignia Gloves (6408, -0.12 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.21 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.09 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.18 DPS) | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.47 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.54 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.16 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 attack_power points (0.61 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Monkey Ring (6748, -0.28 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.56 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS, sim-verified) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Protector's Band (20439, -0.19 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (15.43 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (77.4 DPS) | yes | Shoni's Disarming Tool (9608, -4.32 DPS) [quest]; Swinetusk Shank (6691, -12.72 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -14.46 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.41 DPS) | yes | Double-barreled Shotgun (2098, -0.14 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 364, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (dwarf, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 112.4. Weights run: 1.5s. Verify run: 1.3s. 593 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.166 ± 0.016, strength=1.000 ± 0.001, crit=0.232 ± 0.008 per rating point (14 rating = 1%, 3.251 per %), hit=0.068 ± 0.003 per rating point (10 rating = 1%, 0.675 per %), melee_haste=not significant (3.438 ± 1.346)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 23.8 attack_power points (1.24 DPS) | yes | Hawkeye's Helm (14591, -0.36 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.37 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.39 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Medallion (19540, -0.06 DPS) [rep]; Sentinel's Medallion (19541, -0.24 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.12 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.8 attack_power points (1.29 DPS) | yes | Flintrock Shoulders (7755, -0.42 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.61 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Cloak (14593, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Dark Hooded Cape (5257, -2.11 DPS, sim-verified) [world] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Wolffear Harness (13110, -0.12 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted]; Quillward Harness (10583, -1.83 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Hawkeye's Bracers (14590, -0.47 DPS) [world_drop]; Cultist's Armguards (270032, -0.52 DPS) [quest]; Ravager's Armguards (14770, -0.64 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.66 DPS) | yes | Gloves of Holy Might (867, -0.45 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.59 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.59 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.56 DPS) | yes | Highlander's Leather Girdle (20117, -0.31 DPS) [rep]; Highlander's Chain Girdle (20090, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.55 DPS) [world_drop] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.35 DPS) | yes | Triprunner Dungarees (9624, -0.10 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.61 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | World drop [world_drop] | 20.7 attack_power points (1.07 DPS) | yes | Skulker's Leather Shoes (252531, +0.15 DPS, sim-verified) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [world_drop] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.04 DPS) | yes | Ironspine's Eye (7686, -0.29 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Mark of Kern (2262, -1.92 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19515) | Silverwing Sentinels [rep] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Mark of Kern (2262, -1.44 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (24.64 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (22.96 DPS) | yes | Shoni's Disarming Tool (9608, -11.29 DPS) [quest]; Stonecloth Branch (15963, -22.80 DPS) [world_drop]; Dazzling Longsword (869, -44.81 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Monolithic Bow (9426, -0.23 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.26 DPS) [vendor]; Bow of Searing Arrows (2825, -1.12 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; feet: Blackforge Greaves; finger1: Assault Band; finger2: Protector's Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 593, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 138.0. Weights run: 1.5s. Verify run: 1.7s. 756 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.196 ± 0.022, strength=1.000 ± 0.001, crit=0.308 ± 0.011 per rating point (14 rating = 1%, 4.313 per %), hit=0.091 ± 0.004 per rating point (10 rating = 1%, 0.907 per %), melee_haste=not significant (2.108 ± 0.907)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.04 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.20 DPS) [dungeon]; White Bandit Mask (10008, -0.81 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.6 attack_power points (1.05 DPS) | yes | Sentinel's Medallion (19539, -0.32 DPS) [rep]; Ghostshard Talisman (7731, -0.33 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -2.27 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.2 attack_power points (1.29 DPS) | yes | Skulker's Leather Shoulder (252535, +0.08 DPS, sim-verified) [crafted]; Failed Flying Experiment (9647, -0.16 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.17 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.7 attack_power points (1.16 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS, sim-verified) [dungeon]; Dark Phantom Cape (13122, -0.25 DPS) [world_drop]; Duskbat Drape (19982, -0.31 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mixologist's Tunic (12793, -0.07 DPS) [dungeon]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Grizzled Pelt (22274, -3.34 DPS, sim-verified) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.43 DPS) | yes | Branded Leather Bracers (19508, -0.41 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.52 DPS) [crafted]; Deepfury Bracers (13120, -0.58 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.3 attack_power points (1.86 DPS) | yes | Rockgrip Gauntlets (17736, -0.42 DPS) [dungeon]; Skulker's Leather Gauntlets (252548, -0.54 DPS) [crafted]; Gauntlets of Divinity (7724, -0.90 DPS, sim-verified) [dungeon] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Leather Girdle (20116, -0.05 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.11 DPS) [crafted]; Girdle of Beastial Fury (11686, -3.13 DPS, sim-verified) [dungeon] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 32.9 attack_power points (1.68 DPS) | yes | Serpentskin Leggings (8262, -0.05 DPS, sim-verified) [world_drop]; Ferine Leggings (6690, -0.35 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.40 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.6 attack_power points (1.36 DPS) | yes | Sandstalker Ankleguards (12470, +0.32 DPS, sim-verified) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.9 attack_power points (1.07 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.8 attack_power points (1.06 DPS) | yes | Assault Band (13095, +0.21 DPS, sim-verified) [world_drop]; Mark of Kern (2262, -0.04 DPS) [dungeon]; Protector's Band (19515, -0.16 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Shadowblade (2163, -3.46 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (26.89 DPS) | yes | Claw of Celebras (17738, -2.26 DPS) [dungeon]; Grizzle's Skinner (11702, -12.90 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -15.41 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Dark Iron Rifle (16004, -1.57 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Bracers of the Stone Princess; hands: Raider Gloves; waist: Substandard Belt Chain; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 756, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 258.1. Weights run: 1.5s. Verify run: 1.6s. 1617 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.205 ± 0.024, strength=1.000 ± 0.001, crit=0.457 ± 0.018 per rating point (14 rating = 1%, 6.402 per %), hit=0.119 ± 0.006 per rating point (10 rating = 1%, 1.190 per %), melee_haste=not significant (0.393 ± 1.319)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Visor (239532) | Leonid Barthalomew the Revered [vendor] | 120.2 attack_power points (6.06 DPS) | yes | Dawnstalker Headpiece (239540, -2.44 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -2.98 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, -6.60 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 32.9 attack_power points (1.66 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Will of the Martyr (17044, -0.15 DPS) [quest]; Imperial Jewel (11933, -4.10 DPS, sim-verified) [dungeon] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | 88.5 attack_power points (4.46 DPS) | yes | Highlander's Leather Shoulders (20059, -1.86 DPS) [rep]; Dawnstalker Spaulders (239542, -1.98 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, -2.27 DPS, sim-verified) [vendor] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 40.0 attack_power points (2.02 DPS) | yes | Cape of the Black Baron (13340, +0.12 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.55 DPS) [vendor]; Windshear Cape (20691, -0.70 DPS) [world] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (258.1 DPS) | yes | Field Marshal's Chain Armor (231563, -2.06 DPS) [vendor]; Dawnstalker Tunic (239543, -2.20 DPS) [vendor]; Tunic of Undead Slaying (23089, -17.43 DPS, sim-verified) [world] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (258.1 DPS) | yes | Dawnstalker Wristguards (239544, -0.64 DPS) [vendor]; Windtalker's Wristguards (19582, -0.81 DPS) [rep]; Wristwraps of Undead Slaying (23093, -8.42 DPS, sim-verified) [world] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | 93.4 attack_power points (4.70 DPS) | yes | Dawnstalker Handguards (239539, -1.72 DPS, sim-verified) [vendor]; Marshal's Chain Grips (231560, -2.16 DPS) [pvp]; Cadaverous Gloves (14640, -2.49 DPS) [dungeon] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 67.8 attack_power points (3.41 DPS) | yes | Dawnstalker Girdle (239538, -0.91 DPS) [vendor]; Ferocity of the Timbermaw (227805, -1.07 DPS) [vendor]; Dense Timbermaw Belt (227807, -3.76 DPS, sim-verified) [vendor] |
| legs | Dawnstalker Leggings (239533) | Leonid Barthalomew the Revered [vendor] | 119.0 attack_power points (6.00 DPS) | yes | Dawnstalker Legguards (239541, -2.50 DPS) [vendor]; Sentinel's Chain Leggings (237819, -3.17 DPS) [vendor]; Marshal's Chain Legplates (231558, -6.45 DPS, sim-verified) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | 99.8 attack_power points (5.03 DPS) | yes | Scalegut Treaders (275618, -2.63 DPS) [crafted]; Dawnstalker Boots (239537, -2.80 DPS) [vendor]; Marshal's Chain Sabatons (231561, -3.22 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (258.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.12 DPS) [quest]; Naglering (11669, -6.34 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (258.1 DPS) | yes | Don Julio's Band (19325, -0.08 DPS) [rep]; Blackstone Ring (17713, -0.20 DPS) [dungeon]; Naglering (11669, -2.94 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (258.1 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -6.82 DPS, sim-verified) [crafted] |
| trinket2 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-verified (258.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.98 DPS, sim-verified) [crafted]; Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (258.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Lobotomizer (19324, -11.03 DPS, sim-verified) [rep] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.2 attack_power points (41.68 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Core Hound Tooth (18805, -39.09 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (258.1 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS, sim-verified) [crafted]; Malgen's Long Bow (22318, -0.10 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.25 DPS) [world_drop] |

**New at 60:** head: Dawnstalker Visor; neck: Amulet of the Darkmoon; shoulder: Dawnstalker Pauldrons; back: Cloak of the Honor Guard; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Leggings; feet: Dawnstalker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Blue Dragon; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1617, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 57.2. Weights run: 1.4s. Verify run: 1.3s. 214 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.041 ± 0.008, strength=1.000 ± 0.001, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.602 per %), hit=0.039 ± 0.002 per rating point (10 rating = 1%, 0.386 per %), melee_haste=not significant (0.818 ± 0.763)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.34 DPS) | yes | Defender's Leather Hood (252447, -0.05 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.21 DPS) | yes | Forest Leather Mantle (4709, -0.22 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.25 DPS) | yes | Catacomb Cloak (279899, +0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 attack_power points (0.50 DPS) | yes | Defender's Leather Armor (252434, -0.16 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.16 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.17 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 attack_power points (0.21 DPS) | yes | Bristlebark Bindings (14569, +0.00 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -0.04 DPS) [vendor]; Ratchet Wristwraps (274742, -0.08 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.41 DPS) | yes | Brawler's Leather Gloves (252494, -0.08 DPS) [crafted]; Gold-flecked Gloves (5195, -0.13 DPS) [dungeon]; Bristlebark Gloves (14572, -0.13 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (0.73 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.48 DPS) [world]; Deviate Scale Belt (6468, -0.53 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.4 attack_power points (0.62 DPS) | yes | Trapper's Leather Pants (252501, +0.45 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.46 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.2 attack_power points (0.33 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest]; Loop of Sacrifice (281673, -0.21 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 attack_power points (0.25 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS) [quest]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Demon Band (12054, -0.16 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.05 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 attack_power points (9.43 DPS) | yes | Cruel Barb (5191, +0.05 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -9.35 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, +0.00 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.08 DPS) [crafted]; Light Bow (4576, -0.08 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 214, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6255 Fishing Pole (JEFFTEST); 6478 Rat Stompers

### Band 30 (troll, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 77.2. Weights run: 1.5s. Verify run: 1.4s. 361 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.050 ± 0.013, strength=1.000 ± 0.001, crit=0.193 ± 0.007 per rating point (14 rating = 1%, 2.706 per %), hit=0.059 ± 0.002 per rating point (10 rating = 1%, 0.586 per %), melee_haste=not significant (0.559 ± 0.829)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.54 DPS) | yes | Brawler's Leather Helm (252512, +0.00 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.63 DPS) | yes | Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Scout's Medallion (19537, -0.27 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 attack_power points (0.75 DPS) | yes | Barbaric Shoulders (5964, -0.29 DPS) [crafted]; Mantle of Thieves (2264, -0.30 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.33 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.4 attack_power points (0.47 DPS) | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Wolfmaster Cape (6314, -0.02 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.7 attack_power points (0.67 DPS) | yes | Brawler's Leather Tunic (252508, +0.05 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Defender's Leather Tunic (252450, -0.11 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 attack_power points (0.47 DPS) | yes | Cultist's Armguards (270032, +0.00 DPS, sim-verified) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.73 DPS) | yes | Insignia Gloves (6408, -0.11 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.21 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.09 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon]; Deftkin Belt (16659, -0.35 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.18 DPS) | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.04 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.4 attack_power points (0.52 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS, sim-verified) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 attack_power points (0.61 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.3 attack_power points (0.56 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS, sim-verified) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (15.43 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS, sim-verified) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (77.2 DPS) | yes | Swinetusk Shank (6691, -12.91 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -14.42 DPS) [quest]; Satyr's Rod (15962, -14.46 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.41 DPS) | yes | Double-barreled Shotgun (2098, -0.14 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 361, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 110.3. Weights run: 1.5s. Verify run: 1.3s. 581 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.166 ± 0.016, strength=1.000 ± 0.001, crit=0.232 ± 0.008 per rating point (14 rating = 1%, 3.251 per %), hit=0.068 ± 0.003 per rating point (10 rating = 1%, 0.675 per %), melee_haste=not significant (3.438 ± 1.346)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 23.8 attack_power points (1.24 DPS) | yes | Hawkeye's Helm (14591, -0.36 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.37 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.39 DPS) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Scout's Medallion (19536, -0.06 DPS) [rep]; Ethereal Talisman (4430, -0.23 DPS) [quest]; Zealous Shadowshard Pendant (17772, -1.11 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.8 attack_power points (1.29 DPS) | yes | Flintrock Shoulders (7755, -0.42 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.61 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.7 attack_power points (0.81 DPS) | yes | Hawkeye's Cloak (14593, +0.96 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.29 DPS) [dungeon]; Wildhunter Cloak (16658, -0.29 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Wolffear Harness (13110, -0.12 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted]; Quillward Harness (10583, -1.84 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Hawkeye's Bracers (14590, -0.47 DPS) [world_drop]; Cultist's Armguards (270032, -0.52 DPS) [quest]; Ravager's Armguards (14770, -0.62 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.66 DPS) | yes | Gloves of Holy Might (867, -0.45 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.56 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.59 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.56 DPS) | yes | Defiler's Leather Girdle (20191, -0.31 DPS) [rep]; Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.55 DPS) [world_drop] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.35 DPS) | yes | Triprunner Dungarees (9624, -0.10 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.40 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | World drop [world_drop] | 20.7 attack_power points (1.07 DPS) | yes | Skulker's Leather Shoes (252531, +0.59 DPS, sim-verified) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [world_drop] |
| finger1 | Assault Band (13095) (or Mark of Kern (2262)) | World drop [world_drop] | 20.0 attack_power points (1.04 DPS) | yes | Ironspine's Eye (7686, -0.29 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Mark of Kern (2262, -1.57 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19512) | Warsong Outriders [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironspine's Eye (7686, -0.15 DPS) [dungeon]; Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Mark of Kern (2262, -1.46 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (24.64 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (22.96 DPS) | yes | Stonecloth Branch (15963, -22.80 DPS) [world_drop]; Tork Wrench (11855, -22.85 DPS) [quest]; Dazzling Longsword (869, -44.32 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Monolithic Bow (9426, -0.23 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.26 DPS) [vendor]; Bow of Searing Arrows (2825, -1.11 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; feet: Blackforge Greaves; finger1: Assault Band; finger2: Legionnaire's Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 581, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 143.0. Weights run: 1.5s. Verify run: 1.7s. 745 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.196 ± 0.022, strength=1.000 ± 0.001, crit=0.308 ± 0.011 per rating point (14 rating = 1%, 4.313 per %), hit=0.091 ± 0.004 per rating point (10 rating = 1%, 0.907 per %), melee_haste=not significant (2.108 ± 0.907)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.04 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Bloomsprout Headpiece (17767, -0.20 DPS) [dungeon]; White Bandit Mask (10008, -0.81 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.6 attack_power points (1.05 DPS) | yes | Woven Ivy Necklace (19159, -0.19 DPS) [quest]; Scout's Medallion (19535, -0.32 DPS) [rep]; Zealous Shadowshard Pendant (17772, -2.46 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.2 attack_power points (1.29 DPS) | yes | Failed Flying Experiment (9647, -0.16 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.17 DPS) [crafted]; Skulker's Leather Shoulder (252535, -0.55 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.7 attack_power points (1.16 DPS) | yes | Blisterbane Wrap (12552, +0.00 DPS, sim-verified) [dungeon]; Dark Phantom Cape (13122, -0.25 DPS) [world_drop]; Duskbat Drape (19982, -0.31 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mixologist's Tunic (12793, -0.07 DPS) [dungeon]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Grizzled Pelt (22274, -3.37 DPS, sim-verified) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | 28.0 attack_power points (1.43 DPS) | yes | Branded Leather Bracers (19508, -0.41 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.52 DPS) [crafted]; Deepfury Bracers (13120, -1.46 DPS, sim-verified) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.3 attack_power points (1.86 DPS) | yes | Rockgrip Gauntlets (17736, -0.42 DPS) [dungeon]; Skulker's Leather Gauntlets (252548, -0.54 DPS) [crafted]; Gauntlets of Divinity (7724, -0.92 DPS, sim-verified) [dungeon] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Leather Girdle (20192, -0.05 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.11 DPS) [crafted]; Girdle of Beastial Fury (11686, -3.15 DPS, sim-verified) [dungeon] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 32.1 attack_power points (1.64 DPS) | yes | Ferine Leggings (6690, +0.90 DPS, sim-verified) [dungeon]; Basilisk Hide Pants (1718, -0.36 DPS) [world_drop]; Triprunner Dungarees (9624, -0.39 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.6 attack_power points (1.36 DPS) | yes | Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon]; Sandstalker Ankleguards (12470, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.23 DPS) | yes | Legionnaire's Band (19511, -0.17 DPS) [rep]; Mark of Kern (2262, -0.20 DPS) [dungeon]; Assault Band (13095, -0.20 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.9 attack_power points (1.07 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Legionnaire's Band (19511, -1.20 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 attack_power points (0.00 DPS) | yes | Blessed Prayer Beads (19990, -1.26 DPS, sim-verified) [quest] |
| main_hand | Hanzo Sword (8190) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Shadowblade (2163, -5.57 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (26.89 DPS) | yes | Claw of Celebras (17738, -2.26 DPS) [dungeon]; White Bone Shredder (11863, -4.04 DPS) [quest]; Grizzle's Skinner (11702, -14.42 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Dark Iron Rifle (16004, -1.58 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Bracers of the Stone Princess; hands: Raider Gloves; waist: Substandard Belt Chain; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hanzo Sword; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 745, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 262.9. Weights run: 1.5s. Verify run: 1.6s. 1605 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.205 ± 0.024, strength=1.000 ± 0.001, crit=0.457 ± 0.018 per rating point (14 rating = 1%, 6.402 per %), hit=0.119 ± 0.006 per rating point (10 rating = 1%, 1.190 per %), melee_haste=not significant (0.393 ± 1.319)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Visor (239532) | Leonid Barthalomew the Revered [vendor] | 120.2 attack_power points (6.06 DPS) | yes | Dawnstalker Headpiece (239540, -2.44 DPS) [vendor]; Champion's Chain Greathelm (227080, -2.98 DPS) [vendor]; Warlord's Chain Greathelm (231568, -6.06 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 32.9 attack_power points (1.66 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Will of the Martyr (17044, -0.15 DPS) [quest]; Imperial Jewel (11933, -4.11 DPS, sim-verified) [dungeon] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | 88.5 attack_power points (4.46 DPS) | yes | Defiler's Leather Shoulders (20194, -1.86 DPS) [rep]; Dawnstalker Spaulders (239542, -1.98 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -3.55 DPS, sim-verified) [vendor] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 40.0 attack_power points (2.02 DPS) | yes | Cape of the Black Baron (13340, +0.13 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -0.55 DPS) [vendor]; Windshear Cape (20691, -0.70 DPS) [world] |
| chest | Dawnstalker Breastplate (239529) | Leonid Barthalomew the Revered [vendor] | sim-verified (262.9 DPS) | yes | Warlord's Chain Armor (231566, -2.06 DPS) [vendor]; Dawnstalker Tunic (239543, -2.20 DPS) [vendor]; Tunic of Undead Slaying (23089, -17.73 DPS, sim-verified) [world] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (262.9 DPS) | yes | Dawnstalker Wristguards (239544, -0.64 DPS) [vendor]; Windtalker's Wristguards (19582, -0.81 DPS) [rep]; Wristwraps of Undead Slaying (23093, -8.50 DPS, sim-verified) [world] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | 93.4 attack_power points (4.70 DPS) | yes | Dawnstalker Handguards (239539, -2.28 DPS) [vendor]; Cadaverous Gloves (14640, -2.49 DPS) [dungeon]; General's Chain Grips (231569, -2.95 DPS, sim-verified) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 67.8 attack_power points (3.41 DPS) | yes | Dawnstalker Girdle (239538, -0.91 DPS) [vendor]; Ferocity of the Timbermaw (227805, -1.07 DPS) [vendor]; Dense Timbermaw Belt (227807, -3.86 DPS, sim-verified) [vendor] |
| legs | Dawnstalker Leggings (239533) | Leonid Barthalomew the Revered [vendor] | 119.0 attack_power points (6.00 DPS) | yes | Dawnstalker Legguards (239541, -2.50 DPS) [vendor]; Sentinel's Chain Leggings (237819, -3.17 DPS) [vendor]; General's Chain Legplates (231567, -5.92 DPS, sim-verified) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | 99.8 attack_power points (5.03 DPS) | yes | General's Chain Sabatons (231564, -2.11 DPS) [pvp]; Dawnstalker Boots (239537, -2.80 DPS) [vendor]; Scalegut Treaders (275618, -8.30 DPS, sim-verified) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (21205) | The Changing of Paths - Protector No More [quest] | sim-verified (262.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (234202, +0.00 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (21204, -0.12 DPS) [quest]; Naglering (11669, -5.38 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (262.9 DPS) | yes | White Bone Band (11862, -0.06 DPS) [quest]; Don Julio's Band (19325, -0.08 DPS) [rep]; Naglering (11669, -3.02 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (262.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (262.9 DPS) | yes | Counterattack Lodestone (18537, -1.05 DPS) [dungeon]; Hand of Justice (11815, -1.15 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.54 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (262.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; The Lobotomizer (19324, -8.60 DPS, sim-verified) [rep] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.2 attack_power points (41.68 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Core Hound Tooth (18805, -42.34 DPS, sim-verified) [world_drop] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (262.9 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS, sim-verified) [crafted]; Malgen's Long Bow (22318, -0.10 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.25 DPS) [world_drop] |

**New at 60:** head: Dawnstalker Visor; neck: Amulet of the Darkmoon; shoulder: Dawnstalker Pauldrons; back: Deathguard's Cloak; chest: Dawnstalker Breastplate; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Dawnstalker Leggings; feet: Dawnstalker Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1605, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

