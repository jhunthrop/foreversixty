# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 54.2. Weights run: 2.8s. Verify run: 1.4s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.184 ± 0.006, strength=1.000 ± 0.001, crit=0.655 ± 0.013 per rating point (14 rating = 1%, 9.174 per %), hit=1.029 ± 0.034 per rating point (10 rating = 1%, 10.292 per %), melee_haste=not significant (2.963 ± 0.972)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 9.5 attack_power points (0.41 DPS) | yes | Defender's Leather Hood (252447, -0.06 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 7.1 attack_power points (0.31 DPS) | yes | Erudite's Amulet (277204, -0.10 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.9 attack_power points (0.26 DPS) | yes | Slime-encrusted Pads (6461, -0.27 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 7.1 attack_power points (0.31 DPS) | yes | Catacomb Cloak (279899, -0.05 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.07 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 13.3 attack_power points (0.58 DPS) | yes | Tunic of Westfall (2041, -0.01 DPS) [quest]; Defender's Leather Armor (252434, -0.12 DPS) [crafted]; Prospector's Chestpiece (14562, -0.18 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.7 attack_power points (0.29 DPS) | yes | Forest Leather Bracers (3202, -0.04 DPS) [world_drop]; Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 11.1 attack_power points (0.48 DPS) | yes | Bristlebark Gloves (14572, -0.10 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Fletcher's Gloves (7348, -0.20 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.52 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 16.7 attack_power points (0.72 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.13 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.5 attack_power points (0.54 DPS) | yes | Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Blackened Defias Boots (10402, -0.23 DPS) [dungeon]; Footpads of the Fang (10411, -0.23 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 attack_power points (0.38 DPS) | yes | Demon Band (12054, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; The 1 Ring (8350, -0.28 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 7.1 attack_power points (0.31 DPS) | yes | Demon Band (12054, -0.13 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; The 1 Ring (8350, -0.21 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.79 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.9 attack_power points (10.14 DPS) | yes | Cruel Barb (5191, +0.00 DPS) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.7 attack_power points (0.21 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 77.1. Weights run: 3.1s. Verify run: 1.6s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.201 ± 0.006, strength=1.000 ± 0.001, crit=0.739 ± 0.013 per rating point (14 rating = 1%, 10.347 per %), hit=1.129 ± 0.040 per rating point (10 rating = 1%, 11.293 per %), melee_haste=9.119 ± 0.816

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 12.0 attack_power points (0.63 DPS) | yes | Defender's Leather Helm (252455, -0.00 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS) [world]; Brawler's Leather Hood (252504, -0.13 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.73 DPS) | yes | Sentinel's Medallion (19541, -0.19 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.27 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 18.2 attack_power points (0.95 DPS) | yes | Mantle of Thieves (2264, -0.17 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.38 DPS) [crafted]; Bristlebark Amice (14573, -0.42 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.4 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.14 DPS) [pvp] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 16.8 attack_power points (0.88 DPS) | yes | Raptorbane Armor (3566, -0.04 DPS) [quest]; Brawler's Leather Tunic (252508, -0.06 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 11.2 attack_power points (0.59 DPS) | yes | Cultist's Armguards (270032, -0.06 DPS) [quest]; Jurassic Wristguards (6198, -0.10 DPS) [world]; Barbaric Bracers (18948, -0.13 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.84 DPS) | yes | Insignia Gloves (6408, -0.09 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.26 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.31 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.38 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.36 DPS) | yes | Petrolspill Leggings (9509, -0.48 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.48 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.59 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.6 attack_power points (0.66 DPS) | yes | Disjointed Shoes (277226, -0.03 DPS) [quest]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.16 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 14.8 attack_power points (0.77 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Insurgent's Band (272067, -0.30 DPS) [vendor]; Monkey Ring (6748, -0.33 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 13.2 attack_power points (0.69 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Insurgent's Band (272067, -0.22 DPS) [vendor]; Monkey Ring (6748, -0.25 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (17.81 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (77.1 DPS) | yes | Shoni's Disarming Tool (9608, -4.99 DPS) [quest]; Satyr's Rod (15962, -16.68 DPS) [world_drop]; Swinetusk Shank (6691, -17.19 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.47 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.16 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 114.9. Weights run: 3.2s. Verify run: 1.6s. 597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.302 ± 0.006, strength=1.000 ± 0.001, crit=0.688 ± 0.012 per rating point (14 rating = 1%, 9.626 per %), hit=1.192 ± 0.056 per rating point (10 rating = 1%, 11.923 per %), melee_haste=7.138 ± 0.826

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.3 attack_power points (1.76 DPS) | yes | Raging Berserker's Helm (7719, -0.19 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.32 DPS) [crafted]; Hawkeye's Helm (14591, -0.49 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.39 DPS) | yes | Sentinel's Medallion (19540, -0.40 DPS) [rep]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.3 attack_power points (1.83 DPS) | yes | Forest Tracker Epaulets (2278, -0.50 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.58 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.0 attack_power points (1.18 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.43 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 32.7 attack_power points (2.28 DPS) | yes | Wolffear Harness (13110, -0.74 DPS) [world_drop]; Kolkar Marauder Chain (6773, -0.80 DPS, sim-verified) [quest]; Tough Scorpid Breastplate (8203, -0.92 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.39 DPS) | yes | Ravager's Armguards (14770, -0.52 DPS, sim-verified) [world_drop]; Hawkeye's Bracers (14590, -0.57 DPS) [world_drop]; Dusky Bracers (7378, -0.67 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (2.23 DPS) | yes | Gloves of Holy Might (867, -0.17 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.49 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.69 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.09 DPS) | yes | Highlander's Chain Girdle (20090, -0.43 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.65 DPS) [world_drop]; Blackened Defias Belt (10403, -0.84 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.3 attack_power points (1.90 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.09 DPS) [dungeon]; Scarlet Leggings (10330, -0.44 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 22.0 attack_power points (1.53 DPS) | yes | Skulker's Leather Shoes (252531, -0.05 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.13 DPS) [crafted]; Imperial Leather Boots (6431, -0.19 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.39 DPS) | yes | Protector's Band (19515, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.39 DPS) | yes | Protector's Band (19515, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (114.9 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.07 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (30.76 DPS) | yes | Jhordy's Misplaced Screwdriver (274753, +0.00 DPS) [vendor]; Shoni's Disarming Tool (9608, -15.12 DPS) [quest]; Stonecloth Branch (15963, -30.55 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (114.9 DPS) | yes | Monolithic Bow (9426, -0.28 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest]; Bow of Searing Arrows (2825, -0.99 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 155.9. Weights run: 3.2s. Verify run: 1.8s. 754 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.322 ± 0.006, strength=1.000 ± 0.001, crit=0.750 ± 0.013 per rating point (14 rating = 1%, 10.506 per %), hit=1.177 ± 0.065 per rating point (10 rating = 1%, 11.774 per %), melee_haste=12.370 ± 1.096

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (3.07 DPS) | yes | Bloomsprout Headpiece (17767, -0.31 DPS) [dungeon]; Knight-Lieutenant's Chain Helmet (220822, -0.64 DPS) [vendor]; White Bandit Mask (10008, -1.11 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 22.2 attack_power points (1.70 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.17 DPS) [quest]; Sentinel's Medallion (19539, -0.48 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 27.7 attack_power points (2.12 DPS) | yes | Sunburn Spaulders (274751, -0.09 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.32 DPS) [crafted]; Failed Flying Experiment (9647, -0.34 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.5 attack_power points (1.88 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Dark Hooded Cape (5257, -0.56 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 34.8 attack_power points (2.67 DPS) | yes | Blazewind Breastplate (11193, -0.11 DPS) [quest]; Quillward Harness (10583, -0.13 DPS) [dungeon]; Fungus Shroud Armor (17742, -0.13 DPS) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.15 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Deepfury Bracers (13120, -0.32 DPS) [world_drop]; Branded Leather Bracers (19508, -0.61 DPS) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 39.1 attack_power points (3.00 DPS) | yes | Gauntlets of Divinity (7724, -0.54 DPS) [dungeon]; Gloves of Holy Might (867, -0.66 DPS) [world_drop]; Rockgrip Gauntlets (17736, -0.85 DPS) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.91 DPS) | yes | Substandard Belt Chain (274757, -0.40 DPS) [vendor]; Skulker's Leather Waistguard (252474, -0.57 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.57 DPS) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 34.8 attack_power points (2.67 DPS) | yes | Serpentskin Leggings (8262, -0.05 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.45 DPS) [vendor]; Basilisk Hide Pants (1718, -0.54 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 28.5 attack_power points (2.18 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 31.8 attack_power points (2.44 DPS) | yes | Mark of Kern (2262, -0.90 DPS) [dungeon]; Assault Band (13095, -0.90 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.02 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 21.9 attack_power points (1.68 DPS) | yes | Mark of Kern (2262, -0.15 DPS) [dungeon]; Assault Band (13095, -0.15 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.26 DPS) [quest] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (155.9 DPS) | yes | Molten Heart of the Mountain (249470, -2.82 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (155.9 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (155.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Barman Shanker (12791, -7.17 DPS, sim-verified) [dungeon] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (40.36 DPS) | yes | Claw of Celebras (17738, -3.39 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.13 DPS) [quest]; Grizzle's Skinner (11702, -33.31 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (155.9 DPS) | yes | Stinging Bow (10624, -0.12 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.12 DPS) [world_drop]; Dark Iron Rifle (16004, -1.49 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Blackveil Cape; chest: Warbear Harness; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 228.9. Weights run: 3.0s. Verify run: 1.8s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.700 ± 0.019, strength=1.000 ± 0.002, crit=1.993 ± 0.043 per rating point (14 rating = 1%, 27.904 per %), hit=3.148 ± 0.223 per rating point (10 rating = 1%, 31.484 per %), melee_haste=not significant (10.478 ± 3.308)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -3.87 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 57.0 attack_power points (4.68 DPS) | yes | Beads of Ogre Might (22150, -0.12 DPS) [quest]; Mark of Fordring (15411, -0.25 DPS) [quest]; Medallion of the Dawn (22659, -0.42 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -2.99 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.5 attack_power points (4.89 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -1.40 DPS) [rep]; Windshear Cape (20691, -2.13 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Armor (231563, -0.94 DPS) [vendor]; Obsidian Mail Tunic (22191, -1.51 DPS) [crafted]; Tunic of Undead Slaying (23089, -7.14 DPS, sim-verified) [world] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.01 DPS) [rep]; Bracers of the Eclipse (18375, -0.20 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -0.99 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 78.3 attack_power points (6.43 DPS) | yes | Marshal's Chain Grips (231560, -0.11 DPS) [pvp]; Marshal's Chain Vices (231578, -1.20 DPS) [vendor]; Stormshroud Gloves (21278, -1.55 DPS) [crafted] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 71.0 attack_power points (5.83 DPS) | yes | Marksman's Girdle (22232, -0.31 DPS) [dungeon]; Dense Timbermaw Belt (227807, -0.57 DPS) [vendor]; Highlander's Chain Girdle (20043, -0.75 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 146.8 attack_power points (12.06 DPS) | yes | Marshal's Chain Legplates (231558, -2.96 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -3.70 DPS) [vendor]; Sentinel's Leather Pants (237818, -3.70 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -2.54 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.51 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Naglering (11669, -3.83 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.42 DPS) [dungeon]; Cutthroat's Signet (272408, -0.56 DPS) [vendor]; Naglering (11669, -2.50 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -7.41 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 831.6 attack_power points (68.32 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -5.88 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -20.72 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.71 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Slashclaw Bracers; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 624.7. Weights run: 3.3s. Verify run: 1.8s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.834 ± 0.023, strength=1.000 ± 0.002, crit=2.534 ± 0.050 per rating point (14 rating = 1%, 35.482 per %), hit=4.375 ± 0.391 per rating point (10 rating = 1%, 43.747 per %), melee_haste=not significant (10.243 ± 4.792)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 123.0 attack_power points (21.23 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Outlaw's Collar (279253, +0.00 DPS) [crafted]; Field Marshal's Chain Greathelm (231562, -0.40 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 71.3 attack_power points (12.30 DPS) | yes | Beads of Ogre Might (22150, -0.61 DPS) [quest]; Mark of Fordring (15411, -1.69 DPS) [quest]; Medallion of the Dawn (22659, -2.03 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 111.5 attack_power points (19.24 DPS) | yes | Field Marshal's Chain Pauldrons (231557, -0.50 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -3.46 DPS) [vendor]; Wyrmhide Spaulders (12082, -5.14 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 71.7 attack_power points (12.38 DPS) | yes | Cape of the Black Baron (13340, -4.18 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.83 DPS) [vendor]; Stalwart Cloak (272415, -4.83 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (624.7 DPS) | yes | Field Marshal's Chain Armor (231563, -4.56 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -4.97 DPS) [vendor]; Tunic of Undead Slaying (23089, -18.05 DPS, sim-verified) [world] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-verified (624.7 DPS) | yes | Blackmist Armguards (12966, -1.35 DPS) [dungeon]; Forest Stalker's Bracers (19587, -1.85 DPS) [rep]; Wristwraps of Undead Slaying (23093, -4.93 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Maxwell's Mission [quest] | 104.0 attack_power points (17.95 DPS) | yes | Marshal's Chain Grips (231560, -3.14 DPS) [pvp]; Stormshroud Gloves (21278, -4.28 DPS) [crafted]; Marshal's Chain Vices (231578, -5.18 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 85.3 attack_power points (14.72 DPS) | yes | Marksman's Girdle (22232, -2.63 DPS, sim-verified) [dungeon]; Highlander's Chain Girdle (20043, -2.72 DPS) [rep]; Highlander's Leather Girdle (20045, -2.72 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 178.9 attack_power points (30.88 DPS) | yes | Marshal's Chain Legplates (231558, -10.05 DPS) [vendor]; Sentinel's Leather Pants (237818, -10.08 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -10.46 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 80.4 attack_power points (13.88 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (624.7 DPS) | yes | Tarnished Elven Ring (18500, -4.14 DPS) [dungeon]; Cutthroat's Signet (272408, -4.45 DPS) [vendor]; Naglering (11669, -8.11 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (624.7 DPS) | yes | Tarnished Elven Ring (18500, -0.95 DPS) [dungeon]; Cutthroat's Signet (272408, -1.27 DPS) [vendor]; Naglering (11669, -8.44 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (624.7 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -3.49 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (624.7 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -4.94 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (624.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -29.52 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 832.8 attack_power points (143.75 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -10.75 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -43.73 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (624.7 DPS) | yes | Blackcrow (12651, -0.76 DPS) [dungeon]; The Purifier (22656, -2.38 DPS) [quest]; Dark Iron Rifle (16004, -3.58 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Slashclaw Bracers; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 54.5. Weights run: 2.8s. Verify run: 1.4s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.184 ± 0.006, strength=1.000 ± 0.001, crit=0.655 ± 0.013 per rating point (14 rating = 1%, 9.174 per %), hit=1.029 ± 0.034 per rating point (10 rating = 1%, 10.292 per %), melee_haste=not significant (2.963 ± 0.972)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 9.5 attack_power points (0.41 DPS) | yes | Defender's Leather Hood (252447, -0.06 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 7.1 attack_power points (0.31 DPS) | yes | Erudite's Amulet (277204, -0.10 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.9 attack_power points (0.26 DPS) | yes | Slime-encrusted Pads (6461, -0.27 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 7.1 attack_power points (0.31 DPS) | yes | Catacomb Cloak (279899, -0.05 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.07 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 13.3 attack_power points (0.58 DPS) | yes | Defender's Leather Armor (252434, -0.12 DPS) [crafted]; Prospector's Chestpiece (14562, -0.18 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.21 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.9 attack_power points (0.26 DPS) | yes | Bristlebark Bindings (14569, -0.02 DPS) [world_drop]; Wolf Bracers (4794, -0.05 DPS) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 11.1 attack_power points (0.48 DPS) | yes | Bristlebark Gloves (14572, -0.10 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Fletcher's Gloves (7348, -0.21 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.78 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Deviate Scale Belt (6468, -0.45 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.52 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 16.7 attack_power points (0.72 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.13 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.5 attack_power points (0.54 DPS) | yes | Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Blackened Defias Boots (10402, -0.23 DPS) [dungeon]; Footpads of the Fang (10411, -0.23 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 attack_power points (0.38 DPS) | yes | Demon Band (12054, -0.21 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.22 DPS) [quest]; Loop of Sacrifice (281673, -0.25 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 7.1 attack_power points (0.31 DPS) | yes | Demon Band (12054, -0.13 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Loop of Sacrifice (281673, -0.18 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.79 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.9 attack_power points (10.14 DPS) | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -10.05 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.7 attack_power points (0.21 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 77.6. Weights run: 3.1s. Verify run: 1.5s. 352 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.201 ± 0.006, strength=1.000 ± 0.001, crit=0.739 ± 0.013 per rating point (14 rating = 1%, 10.347 per %), hit=1.129 ± 0.040 per rating point (10 rating = 1%, 11.293 per %), melee_haste=9.119 ± 0.816

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 12.0 attack_power points (0.63 DPS) | yes | Defender's Leather Helm (252455, -0.00 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS) [world]; Brawler's Leather Hood (252504, -0.13 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.73 DPS) | yes | Scout's Medallion (19537, -0.18 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.27 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 18.2 attack_power points (0.95 DPS) | yes | Mantle of Thieves (2264, -0.26 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.38 DPS) [crafted]; Bristlebark Amice (14573, -0.42 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 11.4 attack_power points (0.60 DPS) | yes | Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Wildhunter Cloak (16658, -0.07 DPS) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 16.8 attack_power points (0.88 DPS) | yes | Brawler's Leather Tunic (252508, -0.06 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Defender's Leather Tunic (252450, -0.19 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 11.2 attack_power points (0.59 DPS) | yes | Cultist's Armguards (270032, -0.06 DPS) [quest]; Jurassic Wristguards (6198, -0.10 DPS) [world]; Barbaric Bracers (18948, -0.13 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.84 DPS) | yes | Insignia Gloves (6408, -0.09 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.20 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.26 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.31 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.38 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.36 DPS) | yes | Petrolspill Leggings (9509, -0.48 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.48 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.82 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 12.6 attack_power points (0.66 DPS) | yes | Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.16 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.16 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 14.8 attack_power points (0.77 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Insurgent's Band (272067, -0.30 DPS) [vendor]; Band of the Fist (17694, -0.31 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 13.2 attack_power points (0.69 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Insurgent's Band (272067, -0.22 DPS) [vendor]; Band of the Fist (17694, -0.23 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (17.81 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (77.6 DPS) | yes | Tork Wrench (11855, -16.64 DPS) [quest]; Satyr's Rod (15962, -16.68 DPS) [world_drop]; Swinetusk Shank (6691, -17.21 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.47 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.16 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 115.5. Weights run: 3.2s. Verify run: 1.6s. 563 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.302 ± 0.006, strength=1.000 ± 0.001, crit=0.688 ± 0.012 per rating point (14 rating = 1%, 9.626 per %), hit=1.192 ± 0.056 per rating point (10 rating = 1%, 11.923 per %), melee_haste=7.138 ± 0.826

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.3 attack_power points (1.76 DPS) | yes | Raging Berserker's Helm (7719, -0.19 DPS) [dungeon]; Barbaric Iron Helm (7915, -0.32 DPS) [crafted]; Hawkeye's Helm (14591, -0.49 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.39 DPS) | yes | Scout's Medallion (19536, -0.38 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.3 attack_power points (1.83 DPS) | yes | Forest Tracker Epaulets (2278, -0.50 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.58 DPS) [dungeon]; Nightscape Shoulders (8192, -0.84 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.0 attack_power points (1.18 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Parachute Cloak (10518, -0.46 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 32.7 attack_power points (2.28 DPS) | yes | Wolffear Harness (13110, -0.74 DPS) [world_drop]; Kolkar Marauder Chain (6773, -0.86 DPS, sim-verified) [quest]; Tough Scorpid Breastplate (8203, -0.92 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.39 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.48 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (2.23 DPS) | yes | Gloves of Holy Might (867, -0.17 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.49 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.69 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.09 DPS) | yes | Defiler's Chain Girdle (20152, -0.43 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.65 DPS) [world_drop]; Blackened Defias Belt (10403, -0.84 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.3 attack_power points (1.90 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.09 DPS) [dungeon]; Scarlet Leggings (10330, -0.44 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 22.0 attack_power points (1.53 DPS) | yes | Skulker's Leather Shoes (252531, -0.05 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.13 DPS) [crafted]; Imperial Leather Boots (6431, -0.19 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.39 DPS) | yes | Legionnaire's Band (19512, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.39 DPS) | yes | Legionnaire's Band (19512, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.37 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (30.76 DPS) | yes | Jhordy's Misplaced Screwdriver (274753, +0.00 DPS) [vendor]; Stonecloth Branch (15963, -30.55 DPS) [world_drop]; Tork Wrench (11855, -30.62 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | Monolithic Bow (9426, -0.28 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest]; Bow of Searing Arrows (2825, -1.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 162.2. Weights run: 3.2s. Verify run: 1.8s. 713 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.322 ± 0.006, strength=1.000 ± 0.001, crit=0.750 ± 0.013 per rating point (14 rating = 1%, 10.506 per %), hit=1.177 ± 0.065 per rating point (10 rating = 1%, 11.774 per %), melee_haste=12.370 ± 1.096

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (3.07 DPS) | yes | Blood Guard's Chain Helmet (220821, -0.64 DPS) [vendor]; Bloomsprout Headpiece (17767, -0.83 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.11 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 22.2 attack_power points (1.70 DPS) | yes | Woven Ivy Necklace (19159, -0.33 DPS) [quest]; Scout's Medallion (19535, -0.48 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.81 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.5 attack_power points (2.03 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.23 DPS) [crafted]; Failed Flying Experiment (9647, -0.25 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.5 attack_power points (1.88 DPS) | yes | Blisterbane Wrap (12552, -0.36 DPS) [dungeon]; Dark Phantom Cape (13122, -0.36 DPS) [world_drop]; Dark Hooded Cape (5257, -0.56 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 34.8 attack_power points (2.67 DPS) | yes | Blazewind Breastplate (11193, -0.11 DPS) [quest]; Quillward Harness (10583, -0.13 DPS) [dungeon]; Fungus Shroud Armor (17742, -0.13 DPS) [dungeon] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (2.15 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 39.1 attack_power points (3.00 DPS) | yes | Gloves of Holy Might (867, -0.66 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.76 DPS, sim-verified) [dungeon]; Rockgrip Gauntlets (17736, -0.85 DPS) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.91 DPS) | yes | Substandard Belt Chain (274757, -0.40 DPS) [vendor]; Skulker's Leather Waistguard (252474, -0.57 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.57 DPS) [rep] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | sim-verified (162.2 DPS) | yes | Basilisk Hide Pants (1718, -0.10 DPS) [world_drop]; Triprunner Dungarees (9624, -0.17 DPS) [quest]; Serpentskin Leggings (8262, -1.70 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 28.5 attack_power points (2.18 DPS) | yes | Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.55 DPS, sim-verified) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 31.8 attack_power points (2.44 DPS) | yes | Legionnaire's Band (19511, -0.76 DPS) [rep]; Mark of Kern (2262, -0.90 DPS) [dungeon]; Assault Band (13095, -0.90 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.84 DPS) | yes | Legionnaire's Band (19511, -0.16 DPS) [rep]; Mark of Kern (2262, -0.31 DPS) [dungeon]; Assault Band (13095, -0.31 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -3.04 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -3.32 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Hanzo Sword (8190, -7.57 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (40.36 DPS) | yes | Claw of Celebras (17738, -3.39 DPS) [dungeon]; White Bone Shredder (11863, -5.99 DPS) [quest]; Grizzle's Skinner (11702, -35.47 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.12 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.12 DPS) [world_drop]; Dark Iron Rifle (16004, -1.91 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Chain Legplates; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 231.3. Weights run: 3.0s. Verify run: 1.8s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.700 ± 0.019, strength=1.000 ± 0.002, crit=1.993 ± 0.043 per rating point (14 rating = 1%, 27.904 per %), hit=3.148 ± 0.223 per rating point (10 rating = 1%, 31.484 per %), melee_haste=not significant (10.478 ± 3.308)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -3.66 DPS, sim-verified) [dungeon] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 57.0 attack_power points (4.68 DPS) | yes | Beads of Ogre Might (22150, -0.12 DPS) [quest]; Mark of Fordring (15411, -0.25 DPS) [quest]; Medallion of the Dawn (22659, -0.42 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -2.70 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.5 attack_power points (4.89 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.40 DPS) [rep]; Windshear Cape (20691, -2.13 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Armor (231566, -0.94 DPS) [vendor]; Obsidian Mail Tunic (22191, -1.51 DPS) [crafted]; Tunic of Undead Slaying (23089, -7.50 DPS, sim-verified) [world] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wristwraps of Undead Slaying (23093, +0.00 DPS) [world]; Forest Stalker's Bracers (19587, -0.01 DPS) [rep]; Bracers of the Eclipse (18375, -0.20 DPS) [dungeon] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 78.3 attack_power points (6.43 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; General's Chain Grips (231569, -0.11 DPS) [vendor]; General's Chain Gloves (16571, -1.20 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 71.0 attack_power points (5.83 DPS) | yes | Marksman's Girdle (22232, -0.31 DPS) [dungeon]; Dense Timbermaw Belt (227807, -0.57 DPS) [vendor]; Defiler's Chain Girdle (20150, -0.75 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 146.8 attack_power points (12.06 DPS) | yes | Outrider's Chain Leggings (22673, -0.97 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -2.96 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -3.70 DPS) [vendor] |
| feet | Bloodmail Boots (14616) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Windreaver Greaves (13967, -3.25 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.51 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Naglering (11669, -3.88 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.42 DPS) [dungeon]; Cutthroat's Signet (272408, -0.56 DPS) [vendor]; Naglering (11669, -3.34 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -7.65 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 831.6 attack_power points (68.32 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -5.75 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -20.72 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Blackcrow (12651, -0.26 DPS) [dungeon]; The Purifier (22656, -0.71 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Slashclaw Bracers; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Bloodmail Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Second Wind; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 627.9. Weights run: 3.3s. Verify run: 1.8s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.834 ± 0.023, strength=1.000 ± 0.002, crit=2.534 ± 0.050 per rating point (14 rating = 1%, 35.482 per %), hit=4.375 ± 0.391 per rating point (10 rating = 1%, 43.747 per %), melee_haste=not significant (10.243 ± 4.792)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 123.0 attack_power points (21.23 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted]; Warlord's Chain Greathelm (231568, -0.40 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 71.3 attack_power points (12.30 DPS) | yes | Beads of Ogre Might (22150, -0.61 DPS) [quest]; Mark of Fordring (15411, -1.69 DPS) [quest]; Medallion of the Dawn (22659, -2.03 DPS) [quest] |
| shoulder | Truestrike Shoulders (12927) | Blackrock Spire: Pyroguard Emberseer [dungeon] | 111.5 attack_power points (19.24 DPS) | yes | Warlord's Chain Pauldrons (231565, -0.50 DPS) [vendor]; Warlord's Chain Shoulders (231572, -3.46 DPS) [pvp]; Wyrmhide Spaulders (12082, -5.12 DPS, sim-verified) [quest] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 71.7 attack_power points (12.38 DPS) | yes | Cape of the Black Baron (13340, -4.18 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.83 DPS) [vendor]; Stalwart Cloak (272415, -4.83 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (627.9 DPS) | yes | Warlord's Chain Armor (231566, -4.56 DPS) [vendor]; Legionnaire's Chain Armor (227083, -4.97 DPS) [vendor]; Tunic of Undead Slaying (23089, -22.13 DPS, sim-verified) [world] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-verified (627.9 DPS) | yes | Blackmist Armguards (12966, -1.35 DPS) [dungeon]; Forest Stalker's Bracers (19587, -1.85 DPS) [rep]; Wristwraps of Undead Slaying (23093, -5.16 DPS, sim-verified) [world] |
| hands | Voone's Vice Grips (13963) | Warlord's Command [quest] | 104.0 attack_power points (17.95 DPS) | yes | General's Chain Grips (231569, -3.14 DPS) [vendor]; Stormshroud Gloves (21278, -4.28 DPS) [crafted]; General's Chain Gloves (16571, -5.18 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 85.3 attack_power points (14.72 DPS) | yes | Marksman's Girdle (22232, -0.52 DPS) [dungeon]; Defiler's Chain Girdle (20150, -2.72 DPS) [rep]; Defiler's Leather Girdle (20190, -2.72 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 178.9 attack_power points (30.88 DPS) | yes | Outrider's Chain Leggings (22673, -5.89 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -10.05 DPS) [vendor]; Sentinel's Leather Pants (237818, -10.08 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 80.4 attack_power points (13.88 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -1.93 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (627.9 DPS) | yes | Tarnished Elven Ring (18500, -4.14 DPS) [dungeon]; Cutthroat's Signet (272408, -4.45 DPS) [vendor]; Naglering (11669, -12.09 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (627.9 DPS) | yes | Tarnished Elven Ring (18500, -0.95 DPS) [dungeon]; Cutthroat's Signet (272408, -1.27 DPS) [vendor]; Naglering (11669, -9.63 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (627.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (627.9 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -0.29 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.74 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (627.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -27.60 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 832.8 attack_power points (143.75 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -11.84 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -43.73 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (627.9 DPS) | yes | Blackcrow (12651, -0.76 DPS) [dungeon]; The Purifier (22656, -2.38 DPS) [quest]; Dark Iron Rifle (16004, -3.44 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Truestrike Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Slashclaw Bracers; hands: Voone's Vice Grips; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

