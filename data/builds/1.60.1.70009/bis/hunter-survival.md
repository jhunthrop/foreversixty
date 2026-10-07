# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 62.4. Weights run: 3.7s. Verify run: 2.0s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.033 ± 0.001, strength=1.000 ± 0.001, crit=0.120 ± 0.003 per rating point (14 rating = 1%, 1.675 per %), hit=0.040 ± 0.001 per rating point (10 rating = 1%, 0.395 per %), melee_haste=not significant (1.750 ± 0.546)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.35 DPS) | yes | Defender's Leather Hood (252447, -0.01 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 attack_power points (0.26 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.22 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.26 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.2 attack_power points (0.52 DPS) | yes | Tunic of Westfall (2041, -0.04 DPS) [quest]; Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.17 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.1 attack_power points (0.26 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.04 DPS) [world_drop]; Wolf Bracers (4794, -0.08 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.43 DPS) | yes | Brawler's Leather Gloves (252494, -0.09 DPS) [crafted]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon]; Bristlebark Gloves (14572, -0.16 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Brawler's Leather Belt (252428, -0.42 DPS) [crafted]; Ruffian Belt (5975, -0.51 DPS) [world]; Deviate Scale Belt (6468, -0.65 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 attack_power points (0.65 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.48 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.1 attack_power points (0.34 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; The 1 Ring (8350, -0.26 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 attack_power points (0.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; The 1 Ring (8350, -0.18 DPS) [world]; Demon Band (12054, -0.19 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.53 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 attack_power points (9.88 DPS) | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 93.3. Weights run: 3.8s. Verify run: 2.1s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.053 ± 0.002, strength=1.000 ± 0.001, crit=0.200 ± 0.005 per rating point (14 rating = 1%, 2.801 per %), hit=0.060 ± 0.001 per rating point (10 rating = 1%, 0.596 per %), melee_haste=not significant (1.421 ± 0.553)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Helm (252512, -0.07 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Kaleidoscope Chain (13084, -0.27 DPS) [world_drop]; Sentinel's Medallion (19541, -0.30 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 attack_power points (0.79 DPS) | yes | Mantle of Thieves (2264, -0.29 DPS) [dungeon]; Barbaric Shoulders (5964, -0.30 DPS) [crafted]; Bristlebark Amice (14573, -0.34 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.4 attack_power points (0.49 DPS) | yes | Wolfmaster Cape (6314, -0.02 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.76 DPS) | yes | Dusky Leather Armor (7374, -0.06 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.07 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 attack_power points (0.49 DPS) | yes | Cultist's Armguards (270032, -0.02 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.76 DPS) | yes | Insignia Gloves (6408, -0.13 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.22 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.14 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.40 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Brawler's Leather Pants (252500, -0.50 DPS) [crafted]; Trapper's Leather Pants (252501, -0.50 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.24 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.17 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 attack_power points (0.64 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Insurgent's Band (272067, -0.21 DPS) [vendor]; Monkey Ring (6748, -0.29 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.58 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.16 DPS) [vendor]; Monkey Ring (6748, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (16.14 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (93.3 DPS) | yes | Shoni's Disarming Tool (9608, -4.52 DPS) [quest]; Satyr's Rod (15962, -15.13 DPS) [world_drop]; Swinetusk Shank (6691, -27.20 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.18 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 129.1. Weights run: 4.0s. Verify run: 2.2s. 597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.178 ± 0.003, strength=1.000 ± 0.001, crit=0.260 ± 0.006 per rating point (14 rating = 1%, 3.643 per %), hit=0.074 ± 0.002 per rating point (10 rating = 1%, 0.738 per %), melee_haste=not significant (1.140 ± 0.585)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 24.0 attack_power points (1.25 DPS) | yes | Hawkeye's Helm (14591, -0.37 DPS) [world_drop]; Raging Berserker's Helm (7719, -0.38 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.39 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.05 DPS) | yes | Sentinel's Medallion (19540, -0.37 DPS) [rep]; Ghostshard Talisman (7731, -0.53 DPS, sim-verified) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.0 attack_power points (1.31 DPS) | yes | Flintrock Shoulders (7755, -0.43 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.62 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.62 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.8 attack_power points (0.83 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.30 DPS) [quest]; Sergeant Major's Cape (16336, -1.82 DPS, sim-verified) [pvp] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 30.4 attack_power points (1.59 DPS) | yes | Wolffear Harness (13110, -0.54 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -0.67 DPS) [crafted]; Kolkar Marauder Chain (6773, -0.87 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.05 DPS) | yes | Hawkeye's Bracers (14590, -0.47 DPS) [world_drop]; Cultist's Armguards (270032, -0.52 DPS) [quest]; Ravager's Armguards (14770, -0.66 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.68 DPS) | yes | Gloves of Holy Might (867, -0.44 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.45 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.59 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.57 DPS) | yes | Highlander's Chain Girdle (20090, -0.53 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.54 DPS) [world_drop]; Blackened Defias Belt (10403, -0.63 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.36 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -6.60 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 20.8 attack_power points (1.09 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.05 DPS) | yes | Protector's Band (19515, -0.14 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.05 DPS) | yes | Protector's Band (19515, -0.14 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (129.1 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -9.77 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (23.14 DPS) | yes | Shoni's Disarming Tool (9608, -11.38 DPS) [quest]; Dazzling Longsword (869, -20.43 DPS, sim-verified) [world_drop]; Stonecloth Branch (15963, -22.99 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (129.1 DPS) | yes | Monolithic Bow (9426, -0.23 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.26 DPS) [vendor]; Bow of Searing Arrows (2825, -1.24 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; feet: Blackforge Greaves; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-3200500000000000-500230131051120151)

Set DPS (verified): 174.2. Weights run: 4.1s. Verify run: 2.4s. 754 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.207 ± 0.004, strength=1.000 ± 0.001, crit=0.362 ± 0.009 per rating point (14 rating = 1%, 5.069 per %), hit=0.098 ± 0.003 per rating point (10 rating = 1%, 0.976 per %), melee_haste=not significant (0.226 ± 0.731)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.08 DPS) | yes | Bloomsprout Headpiece (17767, -0.21 DPS) [dungeon]; Knight-Lieutenant's Chain Helmet (220822, -0.81 DPS) [vendor]; White Bandit Mask (10008, -0.82 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.7 attack_power points (1.08 DPS) | yes | Sentinel's Medallion (19539, -0.32 DPS) [rep]; Ghostshard Talisman (7731, -0.35 DPS) [dungeon]; Zealous Shadowshard Pendant (17772, -2.47 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.3 attack_power points (1.32 DPS) | yes | Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted]; Skulker's Leather Shoulder (252535, -1.27 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.9 attack_power points (1.19 DPS) | yes | Blisterbane Wrap (12552, -0.25 DPS) [dungeon]; Dark Hooded Cape (5257, -0.36 DPS) [world]; Dark Phantom Cape (13122, -1.40 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.7 attack_power points (1.70 DPS) | yes | Mixologist's Tunic (12793, -0.08 DPS) [dungeon]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Blazewind Breastplate (11193, -0.10 DPS) [quest] |
| wrist | Arena Bands (18711) | Arena Treasure Chest [world] | 28.0 attack_power points (1.46 DPS) | yes | Deepfury Bracers (13120, -0.31 DPS) [world_drop]; Branded Leather Bracers (19508, -0.42 DPS) [dungeon]; Bracers of the Stone Princess (17714, -4.04 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.5 attack_power points (1.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.44 DPS) [dungeon]; Skulker's Leather Gauntlets (252548, -0.55 DPS) [crafted]; Gauntlets of Divinity (7724, -0.96 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (1.98 DPS) | yes | Substandard Belt Chain (274757, -0.36 DPS) [vendor]; Highlander's Leather Girdle (20116, -0.42 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.47 DPS) [crafted] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 33.1 attack_power points (1.72 DPS) | yes | Serpentskin Leggings (8262, -0.04 DPS) [world_drop]; Ferine Leggings (6690, -0.37 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.40 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.7 attack_power points (1.39 DPS) | yes | Sandstalker Ankleguards (12470, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.09 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.9 attack_power points (1.09 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Assault Band (13095, -0.04 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (174.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (174.2 DPS) | yes | Molten Heart of the Mountain (249470, -0.68 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (174.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Hanzo Sword (8190, -15.91 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (27.41 DPS) | yes | Claw of Celebras (17738, -2.31 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.71 DPS) [quest]; Grizzle's Skinner (11702, -32.39 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (174.2 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Dark Iron Rifle (16004, -1.76 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-3200550000500000-500230131051120151)

Set DPS (verified): 251.7. Weights run: 4.1s. Verify run: 2.4s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.263 ± 0.007, strength=1.000 ± 0.001, crit=0.577 ± 0.015 per rating point (14 rating = 1%, 8.084 per %), hit=0.148 ± 0.005 per rating point (10 rating = 1%, 1.481 per %), melee_haste=not significant (3.456 ± 1.135)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crown of Tyranny (13359) | Stratholme: Balnazzar [dungeon] | 40.6 attack_power points (2.09 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -2.37 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 34.0 attack_power points (1.75 DPS) | yes | Medallion of the Dawn (22659, -0.10 DPS) [quest]; Imperial Jewel (11933, -0.10 DPS) [dungeon]; Will of the Martyr (17044, -0.21 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 52.7 attack_power points (2.71 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Lieutenant Commander's Chain Pauldrons (227084, -0.22 DPS) [pvp]; Highlander's Lizardhide Shoulders (20060, -3.76 DPS, sim-verified) [rep] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 40.3 attack_power points (2.08 DPS) | yes | Cape of the Black Baron (13340, -0.07 DPS) [dungeon]; Howler's Furs (272414, -0.56 DPS) [vendor]; Windshear Cape (20691, -0.69 DPS) [world] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (251.7 DPS) | yes | Field Marshal's Chain Armor (231563, -0.03 DPS) [vendor]; Cadaverous Armor (14637, -0.31 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.97 DPS, sim-verified) [world] |
| wrist | Wristwraps of Undead Slaying (23093) | Bone Witch [world] | sim-verified (251.7 DPS) | yes | Bracers of the Eclipse (18375, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Windtalker's Wristguards (19582, -6.56 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 45.1 attack_power points (2.32 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Cadaverous Gloves (14640, -0.06 DPS) [dungeon]; Knight-Lieutenant's Chain Grips (227087, -0.24 DPS) [vendor] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (3.29 DPS) | yes | Chiselbrand Girdle (12634, -1.03 DPS) [dungeon]; Highlander's Chain Girdle (20043, -1.13 DPS) [rep]; Ferocity of the Timbermaw (227805, -1.34 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 61.9 attack_power points (3.18 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Devilsaur Leggings (15062, -0.40 DPS) [crafted]; Knight-Captain's Chain Legplates (227085, -2.35 DPS, sim-verified) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 48.1 attack_power points (2.48 DPS) | yes | Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Knight-Lieutenant's Chain Sabatons (227088, -0.18 DPS) [vendor]; Pads of the Dread Wolf (13210, -0.42 DPS) [dungeon] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (251.7 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.09 DPS) [quest]; Blackstone Ring (17713, -0.23 DPS) [dungeon]; Naglering (11669, -3.41 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (251.7 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.07 DPS) [quest]; Blackstone Ring (17713, -0.21 DPS) [dungeon]; Naglering (11669, -5.43 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (251.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -4.41 DPS, sim-verified) [quest] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (251.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (251.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -25.95 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.7 attack_power points (42.61 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -2.61 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -12.78 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (251.7 DPS) | yes | Malgen's Long Bow (22318, -0.10 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.22 DPS) [world_drop]; Dark Iron Rifle (16004, -2.43 DPS, sim-verified) [crafted] |

**New at 60:** head: Crown of Tyranny; neck: Amulet of the Darkmoon; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Obsidian Mail Tunic; wrist: Wristwraps of Undead Slaying; hands: Raider Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Devilsaur Eye; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 62.8. Weights run: 3.7s. Verify run: 1.9s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.033 ± 0.001, strength=1.000 ± 0.001, crit=0.120 ± 0.003 per rating point (14 rating = 1%, 1.675 per %), hit=0.040 ± 0.001 per rating point (10 rating = 1%, 0.395 per %), melee_haste=not significant (1.750 ± 0.546)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.35 DPS) | yes | Defender's Leather Hood (252447, -0.01 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 attack_power points (0.26 DPS) | yes | Erudite's Amulet (277204, -0.09 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.22 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.26 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.2 attack_power points (0.52 DPS) | yes | Prospector's Chestpiece (14562, -0.17 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.17 DPS) [crafted]; Defender's Leather Armor (252434, -0.35 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 attack_power points (0.22 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.04 DPS) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.43 DPS) | yes | Bristlebark Gloves (14572, -0.09 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.09 DPS) [crafted]; Gold-flecked Gloves (5195, -0.14 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Brawler's Leather Belt (252428, -0.42 DPS) [crafted]; Ruffian Belt (5975, -0.51 DPS) [world]; Deviate Scale Belt (6468, -0.66 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 attack_power points (0.65 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.48 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.21 DPS) [dungeon]; Footpads of the Fang (10411, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.1 attack_power points (0.34 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.21 DPS) [quest]; Loop of Sacrifice (281673, -0.22 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 attack_power points (0.26 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS) [quest]; Loop of Sacrifice (281673, -0.14 DPS) [quest]; Demon Band (12054, -0.20 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.53 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 attack_power points (9.88 DPS) | yes | Cruel Barb (5191, +0.00 DPS) [dungeon]; Tork Wrench (11855, -9.79 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted]; Light Bow (4576, -0.09 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 93.6. Weights run: 3.8s. Verify run: 2.1s. 352 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.053 ± 0.002, strength=1.000 ± 0.001, crit=0.200 ± 0.005 per rating point (14 rating = 1%, 2.801 per %), hit=0.060 ± 0.001 per rating point (10 rating = 1%, 0.596 per %), melee_haste=not significant (1.421 ± 0.553)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.57 DPS) | yes | Brawler's Leather Helm (252512, -0.07 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.66 DPS) | yes | Kaleidoscope Chain (13084, -0.27 DPS) [world_drop]; Scout's Medallion (19537, -0.29 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.6 attack_power points (0.79 DPS) | yes | Barbaric Shoulders (5964, -0.30 DPS) [crafted]; Bristlebark Amice (14573, -0.34 DPS) [world_drop]; Mantle of Thieves (2264, -0.43 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.4 attack_power points (0.49 DPS) | yes | Wolfmaster Cape (6314, -0.02 DPS) [dungeon]; Wildhunter Cloak (16658, -0.02 DPS) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.7 attack_power points (0.70 DPS) | yes | Brawler's Leather Tunic (252508, -0.02 DPS) [crafted]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Defender's Leather Tunic (252450, -0.11 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 attack_power points (0.49 DPS) | yes | Cultist's Armguards (270032, -0.02 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.10 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.76 DPS) | yes | Insignia Gloves (6408, -0.13 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.22 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.14 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.37 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.23 DPS) | yes | Brawler's Leather Pants (252500, -0.50 DPS) [crafted]; Trapper's Leather Pants (252501, -0.50 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.39 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.4 attack_power points (0.54 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.5 attack_power points (0.64 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Insurgent's Band (272067, -0.21 DPS) [vendor]; Band of the Fist (17694, -0.25 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.3 attack_power points (0.58 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.16 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (16.14 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (93.6 DPS) | yes | Tork Wrench (11855, -15.08 DPS) [quest]; Satyr's Rod (15962, -15.13 DPS) [world_drop]; Swinetusk Shank (6691, -27.69 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.43 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.18 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.23 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 126.7. Weights run: 4.0s. Verify run: 2.2s. 563 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.178 ± 0.003, strength=1.000 ± 0.001, crit=0.260 ± 0.006 per rating point (14 rating = 1%, 3.643 per %), hit=0.074 ± 0.002 per rating point (10 rating = 1%, 0.738 per %), melee_haste=not significant (1.140 ± 0.585)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 24.0 attack_power points (1.25 DPS) | yes | Barbaric Iron Helm (7915, -0.23 DPS) [crafted]; Hawkeye's Helm (14591, -0.37 DPS) [world_drop]; Raging Berserker's Helm (7719, -0.38 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.05 DPS) | yes | Scout's Medallion (19536, -0.37 DPS) [rep]; Ghostshard Talisman (7731, -0.52 DPS, sim-verified) [dungeon]; Ethereal Talisman (4430, -0.54 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.0 attack_power points (1.31 DPS) | yes | Flintrock Shoulders (7755, -0.43 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.61 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.62 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.8 attack_power points (0.83 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Wildhunter Cloak (16658, -0.30 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 30.4 attack_power points (1.59 DPS) | yes | Wolffear Harness (13110, -0.54 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -0.67 DPS) [crafted]; Kolkar Marauder Chain (6773, -0.84 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.05 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.72 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.68 DPS) | yes | Gloves of Holy Might (867, -0.44 DPS) [world_drop]; Skulker's Leather Gloves (252525, -0.59 DPS) [crafted]; Scarlet Gauntlets (10331, -0.63 DPS, sim-verified) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.57 DPS) | yes | Defiler's Chain Girdle (20152, -0.52 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.54 DPS) [world_drop]; Blackened Defias Belt (10403, -0.63 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.36 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.94 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 20.8 attack_power points (1.09 DPS) | yes | Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Skulker's Leather Shoes (252531, -1.53 DPS, sim-verified) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.05 DPS) | yes | Legionnaire's Band (19512, -0.14 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.05 DPS) | yes | Legionnaire's Band (19512, -0.14 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (126.7 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -8.32 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (23.14 DPS) | yes | Dazzling Longsword (869, -18.67 DPS, sim-verified) [world_drop]; Stonecloth Branch (15963, -22.99 DPS) [world_drop]; Tork Wrench (11855, -23.04 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (126.7 DPS) | yes | Monolithic Bow (9426, -0.23 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.26 DPS) [vendor]; Bow of Searing Arrows (2825, -1.22 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; feet: Blackforge Greaves; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-3200500000000000-500230131051120151)

Set DPS (verified): 176.9. Weights run: 4.1s. Verify run: 2.5s. 713 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.207 ± 0.004, strength=1.000 ± 0.001, crit=0.362 ± 0.009 per rating point (14 rating = 1%, 5.069 per %), hit=0.098 ± 0.003 per rating point (10 rating = 1%, 0.976 per %), melee_haste=not significant (0.226 ± 0.731)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.08 DPS) | yes | Bloomsprout Headpiece (17767, -0.21 DPS) [dungeon]; Blood Guard's Chain Helmet (220821, -0.81 DPS) [vendor]; White Bandit Mask (10008, -0.82 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.7 attack_power points (1.08 DPS) | yes | Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Scout's Medallion (19535, -0.32 DPS) [rep]; Zealous Shadowshard Pendant (17772, -8.27 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.3 attack_power points (1.32 DPS) | yes | Skulker's Leather Shoulder (252535, -0.16 DPS) [crafted]; Failed Flying Experiment (9647, -0.17 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.18 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.9 attack_power points (1.19 DPS) | yes | Blisterbane Wrap (12552, -0.25 DPS) [dungeon]; Dark Hooded Cape (5257, -0.36 DPS) [world]; Dark Phantom Cape (13122, -0.66 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.7 attack_power points (1.70 DPS) | yes | Mixologist's Tunic (12793, -0.08 DPS) [dungeon]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Blazewind Breastplate (11193, -0.10 DPS) [quest] |
| wrist | Bracers of the Stone Princess (17714) | Maraudon: Princess Theradras [dungeon] | sim-verified (176.9 DPS) | yes | Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -1.98 DPS, sim-verified) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.5 attack_power points (1.90 DPS) | yes | Rockgrip Gauntlets (17736, -0.44 DPS) [dungeon]; Skulker's Leather Gauntlets (252548, -0.55 DPS) [crafted]; Gauntlets of Divinity (7724, -1.40 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (1.98 DPS) | yes | Substandard Belt Chain (274757, -0.36 DPS) [vendor]; Defiler's Leather Girdle (20192, -0.42 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.47 DPS) [crafted] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 32.3 attack_power points (1.68 DPS) | yes | Ferine Leggings (6690, -0.33 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.36 DPS) [world_drop]; Triprunner Dungarees (9624, -0.39 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.7 attack_power points (1.39 DPS) | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.25 DPS) | yes | Legionnaire's Band (19511, -0.16 DPS) [rep]; Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.0 attack_power points (1.09 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Legionnaire's Band (19511, -1.41 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.18 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Hanzo Sword (8190, -16.80 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (27.41 DPS) | yes | Claw of Celebras (17738, -2.31 DPS) [dungeon]; White Bone Shredder (11863, -4.11 DPS) [quest]; Grizzle's Skinner (11702, -29.79 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Dark Iron Rifle (16004, -1.74 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Bracers of the Stone Princess; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-3200550000500000-500230131051120151)

Set DPS (verified): 258.3. Weights run: 4.1s. Verify run: 2.5s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.263 ± 0.007, strength=1.000 ± 0.001, crit=0.577 ± 0.015 per rating point (14 rating = 1%, 8.084 per %), hit=0.148 ± 0.005 per rating point (10 rating = 1%, 1.481 per %), melee_haste=not significant (3.456 ± 1.135)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+7.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Crown of Tyranny (13359, -7.58 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 34.0 attack_power points (1.75 DPS) | yes | Medallion of the Dawn (22659, -0.10 DPS) [quest]; Imperial Jewel (11933, -0.10 DPS) [dungeon]; Will of the Martyr (17044, -0.21 DPS) [quest] |
| shoulder | Defiler's Lizardhide Shoulders (20175) | The Defilers [rep] | sim-verified (+10.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Chain Pauldrons (227078, +0.00 DPS) [pvp]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Defiler's Leather Shoulders (20194, -10.81 DPS, sim-verified) [rep] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 40.3 attack_power points (2.08 DPS) | yes | Cape of the Black Baron (13340, -0.07 DPS) [dungeon]; Howler's Furs (272414, -0.56 DPS) [vendor]; Windshear Cape (20691, -0.69 DPS) [world] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Armor (231566, -0.03 DPS) [vendor]; Cadaverous Armor (14637, -0.31 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.50 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wristwraps of Undead Slaying (23093, +0.00 DPS) [world]; Bracers of the Eclipse (18375, -0.07 DPS) [dungeon]; Forest Stalker's Bracers (19587, -0.15 DPS) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 45.1 attack_power points (2.32 DPS) | yes | General's Chain Grips (231569, +0.00 DPS) [vendor]; Blood Guard's Chain Grips (227081, -0.24 DPS) [vendor]; Cadaverous Gloves (14640, -1.73 DPS, sim-verified) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (3.29 DPS) | yes | Ferocity of the Timbermaw (227805, -0.85 DPS) [vendor]; Chiselbrand Girdle (12634, -1.03 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.13 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 61.9 attack_power points (3.18 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -0.21 DPS) [vendor]; Devilsaur Leggings (15062, -6.85 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 48.1 attack_power points (2.48 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; Blood Guard's Chain Sabatons (227082, -0.18 DPS) [vendor]; Pads of the Dread Wolf (13210, -0.42 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.09 DPS) [quest]; White Bone Band (11862, -0.10 DPS) [quest]; Naglering (11669, -2.89 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.07 DPS) [quest]; White Bone Band (11862, -0.08 DPS) [quest]; Naglering (11669, -7.43 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (-5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -3.21 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -14.36 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.7 attack_power points (42.61 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -3.65 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -12.78 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.10 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.22 DPS) [world_drop]; Dark Iron Rifle (16004, -2.36 DPS, sim-verified) [crafted] |

**New at 60:** neck: Amulet of the Darkmoon; shoulder: Defiler's Lizardhide Shoulders; back: Deathguard's Cloak; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Raider Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Legionnaire's Band; finger2: Don Julio's Band; trinket1: Second Wind; trinket2: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

