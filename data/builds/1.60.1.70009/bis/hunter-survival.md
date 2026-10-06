# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 57.3. Weights run: 3.2s. Verify run: 1.8s. 220 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.037 ± 0.006, strength=1.000 ± 0.001, crit=0.119 ± 0.003 per rating point (14 rating = 1%, 1.665 per %), hit=0.039 ± 0.001 per rating point (10 rating = 1%, 0.389 per %), melee_haste=not significant (0.776 ± 0.540)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.33 DPS) | yes | Defender's Leather Hood (252447, -0.01 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.08 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.21 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.25 DPS) | yes | Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Catacomb Cloak (279899, -0.17 DPS, sim-verified) [quest] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 attack_power points (0.49 DPS) | yes | Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.16 DPS) [world_drop]; Tunic of Westfall (2041, -0.19 DPS, sim-verified) [quest] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.1 attack_power points (0.25 DPS) | yes | Bristlebark Bindings (14569, -0.04 DPS) [world_drop]; Wolf Bracers (4794, -0.08 DPS) [vendor]; Forest Leather Bracers (3202, -0.22 DPS, sim-verified) [world_drop] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.41 DPS) | yes | Bristlebark Gloves (14572, -0.08 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.08 DPS) [crafted]; Gold-flecked Gloves (5195, -0.13 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.73 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.48 DPS) [world]; Deviate Scale Belt (6468, -0.55 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 attack_power points (0.62 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.46 DPS) | yes | Brawler's Leather Boots (252439, -0.04 DPS) [crafted]; Blackened Defias Boots (10402, -0.20 DPS) [dungeon]; Footpads of the Fang (10411, -0.20 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.1 attack_power points (0.33 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 attack_power points (0.25 DPS) | yes | Demon Band (12054, -0.09 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.17 DPS) [dungeon]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.05 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 attack_power points (9.43 DPS) | yes | Cruel Barb (5191, +0.00 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.08 DPS) [crafted]; Light Bow (4576, -0.08 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 220, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 85.1. Weights run: 3.4s. Verify run: 1.9s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.044 ± 0.008, strength=1.000 ± 0.001, crit=0.190 ± 0.005 per rating point (14 rating = 1%, 2.660 per %), hit=0.058 ± 0.001 per rating point (10 rating = 1%, 0.578 per %), melee_haste=not significant (0.523 ± 0.580)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.54 DPS) | yes | Brawler's Leather Helm (252512, -0.07 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.63 DPS) | yes | Sentinel's Medallion (19541, -0.26 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.5 attack_power points (0.75 DPS) | yes | Mantle of Thieves (2264, -0.27 DPS) [dungeon]; Barbaric Shoulders (5964, -0.28 DPS) [crafted]; Bristlebark Amice (14573, -0.33 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.3 attack_power points (0.47 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.72 DPS) | yes | Dusky Leather Armor (7374, -0.06 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.07 DPS) [crafted]; Brawler's Leather Armor (252490, -0.17 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 attack_power points (0.46 DPS) | yes | Cultist's Armguards (270032, -0.01 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.09 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.72 DPS) | yes | Insignia Gloves (6408, -0.12 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.21 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.09 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.18 DPS) | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.47 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.54 DPS) | yes | Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.08 DPS) [crafted]; Insignia Boots (4055, -0.17 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.4 attack_power points (0.61 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Monkey Ring (6748, -0.28 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.56 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Monkey Ring (6748, -0.22 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (15.42 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (85.1 DPS) | yes | Shoni's Disarming Tool (9608, -4.32 DPS) [quest]; Satyr's Rod (15962, -14.45 DPS) [world_drop]; Swinetusk Shank (6691, -20.46 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.41 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 118.6. Weights run: 3.4s. Verify run: 1.8s. 597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.178 ± 0.012, strength=1.000 ± 0.001, crit=0.234 ± 0.005 per rating point (14 rating = 1%, 3.282 per %), hit=0.068 ± 0.002 per rating point (10 rating = 1%, 0.677 per %), melee_haste=not significant (3.021 ± 0.977)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 24.0 attack_power points (1.24 DPS) | yes | Hawkeye's Helm (14591, -0.36 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.38 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.40 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.04 DPS) | yes | Sentinel's Medallion (19540, -0.37 DPS) [rep]; Ghostshard Talisman (7731, -0.47 DPS, sim-verified) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.0 attack_power points (1.30 DPS) | yes | Flintrock Shoulders (7755, -0.42 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.55 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.62 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.8 attack_power points (0.82 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.30 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 30.4 attack_power points (1.58 DPS) | yes | Wolffear Harness (13110, -0.54 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -0.66 DPS) [crafted]; Kolkar Marauder Chain (6773, -0.86 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Hawkeye's Bracers (14590, -0.46 DPS) [world_drop]; Cultist's Armguards (270032, -0.52 DPS) [quest]; Ravager's Armguards (14770, -0.58 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.66 DPS) | yes | Gloves of Holy Might (867, -0.45 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.46 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.58 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.56 DPS) | yes | Highlander's Chain Girdle (20090, -0.47 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.54 DPS) [world_drop]; Blackened Defias Belt (10403, -0.62 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.35 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.96 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 20.8 attack_power points (1.08 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.04 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (24.65 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (22.96 DPS) | yes | Dazzling Longsword (869, -10.61 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -11.29 DPS) [quest]; Stonecloth Branch (15963, -22.81 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (118.6 DPS) | yes | Monolithic Bow (9426, -0.23 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.26 DPS) [vendor]; Bow of Searing Arrows (2825, -1.10 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Highlander's Leather Girdle; feet: Blackforge Greaves; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 597, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 162.2. Weights run: 3.4s. Verify run: 2.1s. 754 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.195 ± 0.016, strength=1.000 ± 0.001, crit=0.317 ± 0.008 per rating point (14 rating = 1%, 4.442 per %), hit=0.089 ± 0.002 per rating point (10 rating = 1%, 0.885 per %), melee_haste=2.785 ± 0.650

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.04 DPS) | yes | Bloomsprout Headpiece (17767, -0.20 DPS) [dungeon]; White Bandit Mask (10008, -0.81 DPS) [crafted]; Knight-Lieutenant's Chain Helmet (220822, -0.84 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.5 attack_power points (1.05 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.03 DPS) [quest]; Sentinel's Medallion (19539, -0.32 DPS) [rep]; Ghostshard Talisman (7731, -0.33 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.1 attack_power points (1.29 DPS) | yes | Skulker's Leather Shoulder (252535, -0.15 DPS) [crafted]; Failed Flying Experiment (9647, -0.16 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.17 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.7 attack_power points (1.16 DPS) | yes | Blisterbane Wrap (12552, -0.25 DPS) [dungeon]; Dark Phantom Cape (13122, -0.25 DPS) [world_drop]; Dark Hooded Cape (5257, -0.35 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.5 attack_power points (1.66 DPS) | yes | Mixologist's Tunic (12793, -0.07 DPS) [dungeon]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Blazewind Breastplate (11193, -0.10 DPS) [quest] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.43 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Deepfury Bracers (13120, -0.31 DPS) [world_drop]; Branded Leather Bracers (19508, -0.41 DPS) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.3 attack_power points (1.85 DPS) | yes | Rockgrip Gauntlets (17736, -0.42 DPS) [dungeon]; Skulker's Leather Gauntlets (252548, -0.54 DPS) [crafted]; Gauntlets of Divinity (7724, -1.04 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (1.94 DPS) | yes | Substandard Belt Chain (274757, -0.36 DPS) [vendor]; Highlander's Leather Girdle (20116, -0.41 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.47 DPS) [crafted] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 32.9 attack_power points (1.68 DPS) | yes | Serpentskin Leggings (8262, -0.04 DPS) [world_drop]; Ferine Leggings (6690, -0.35 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.40 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.5 attack_power points (1.36 DPS) | yes | Sandstalker Ankleguards (12470, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.9 attack_power points (1.07 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 20.8 attack_power points (1.06 DPS) | yes | Mark of Kern (2262, -0.04 DPS) [dungeon]; Assault Band (13095, -0.04 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (162.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (162.2 DPS) | yes | Molten Heart of the Mountain (249470, -0.73 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (162.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Hanzo Sword (8190, -13.77 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (26.91 DPS) | yes | Claw of Celebras (17738, -2.26 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.42 DPS) [quest]; Grizzle's Skinner (11702, -19.65 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (162.2 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Dark Iron Rifle (16004, -1.61 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 754, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 238.0. Weights run: 3.4s. Verify run: 2.1s. 1670 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.209 ± 0.020, strength=1.000 ± 0.001, crit=0.475 ± 0.013 per rating point (14 rating = 1%, 6.656 per %), hit=0.117 ± 0.004 per rating point (10 rating = 1%, 1.172 per %), melee_haste=not significant (1.596 ± 0.951)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crown of Tyranny (13359) | Stratholme: Balnazzar [dungeon] | 40.5 attack_power points (2.04 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.78 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 33.0 attack_power points (1.66 DPS) | yes | Imperial Jewel (11933, -0.05 DPS) [dungeon]; Medallion of the Dawn (22659, -0.12 DPS) [quest]; Will of the Martyr (17044, -0.15 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 51.8 attack_power points (2.61 DPS) | yes | Highlander's Lizardhide Shoulders (20060, +0.00 DPS, sim-verified) [rep]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Lieutenant Commander's Chain Pauldrons (227084, -0.26 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 40.0 attack_power points (2.02 DPS) | yes | Cape of the Black Baron (13340, -0.10 DPS) [dungeon]; Howler's Furs (272414, -0.55 DPS) [vendor]; Windshear Cape (20691, -0.70 DPS) [world] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (238.0 DPS) | yes | Field Marshal's Chain Armor (231563, -0.07 DPS) [vendor]; Cadaverous Armor (14637, -0.25 DPS) [dungeon]; Tunic of Undead Slaying (23089, -10.17 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | sim-verified (238.0 DPS) | yes | Bracers of the Eclipse (18375, -0.10 DPS) [dungeon]; Forest Stalker's Bracers (19587, -0.20 DPS) [rep]; Wristwraps of Undead Slaying (23093, -6.67 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (2.22 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Raider Gloves (272099, -0.02 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.20 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (3.23 DPS) | yes | Chiselbrand Girdle (12634, -1.01 DPS) [dungeon]; Heavy Timbermaw Belt (19043, -1.11 DPS) [crafted]; Ferocity of the Timbermaw (227805, -1.28 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 56.8 attack_power points (2.87 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Black Dragonscale Leggings (15052, -0.14 DPS) [crafted]; Knight-Captain's Chain Legplates (227085, -2.43 DPS, sim-verified) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 47.7 attack_power points (2.40 DPS) | yes | Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Knight-Lieutenant's Chain Sabatons (227088, -0.18 DPS) [vendor]; Pads of the Dread Wolf (13210, -0.39 DPS) [dungeon] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (238.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.12 DPS) [quest]; Blackstone Ring (17713, -0.21 DPS) [dungeon]; Naglering (11669, -3.09 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (238.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.04 DPS) [quest]; Blackstone Ring (17713, -0.13 DPS) [dungeon]; Naglering (11669, -5.41 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (238.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (238.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Devilsaur Eye (19991, -4.95 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (238.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -26.26 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.2 attack_power points (41.73 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -3.21 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -12.50 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (238.0 DPS) | yes | Malgen's Long Bow (22318, -0.10 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.25 DPS) [world_drop]; Dark Iron Rifle (16004, -2.28 DPS, sim-verified) [crafted] |

**New at 60:** head: Crown of Tyranny; neck: Amulet of the Darkmoon; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Cadaverous Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Second Wind; trinket2: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1670, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-00000000000000000-500230100000000000)

Set DPS (verified): 57.2. Weights run: 3.2s. Verify run: 1.7s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.037 ± 0.006, strength=1.000 ± 0.001, crit=0.119 ± 0.003 per rating point (14 rating = 1%, 1.665 per %), hit=0.039 ± 0.001 per rating point (10 rating = 1%, 0.389 per %), melee_haste=not significant (0.776 ± 0.540)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.3 attack_power points (0.33 DPS) | yes | Defender's Leather Hood (252447, -0.01 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.08 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.2 attack_power points (0.21 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.2 attack_power points (0.25 DPS) | yes | Catacomb Cloak (279899, -0.01 DPS) [quest]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.3 attack_power points (0.49 DPS) | yes | Defender's Leather Armor (252434, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.16 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.17 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.2 attack_power points (0.21 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.04 DPS) [vendor]; Ratchet Wristwraps (274742, -0.08 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.2 attack_power points (0.41 DPS) | yes | Bristlebark Gloves (14572, -0.08 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.08 DPS) [crafted]; Gold-flecked Gloves (5195, -0.13 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.73 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Ruffian Belt (5975, -0.48 DPS) [world]; Deviate Scale Belt (6468, -0.53 DPS, sim-verified) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.3 attack_power points (0.62 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Leggings of the Fang (10410, -0.04 DPS) [dungeon]; Defender's Leather Pants (252445, -0.09 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.3 attack_power points (0.46 DPS) | yes | Brawler's Leather Boots (252439, -0.04 DPS) [crafted]; Blackened Defias Boots (10402, -0.20 DPS) [dungeon]; Footpads of the Fang (10411, -0.20 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.1 attack_power points (0.33 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest]; Loop of Sacrifice (281673, -0.21 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.2 attack_power points (0.25 DPS) | yes | Demon Band (12054, -0.09 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.13 DPS) [quest]; Loop of Sacrifice (281673, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.05 DPS) | yes | Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 233.6 attack_power points (9.43 DPS) | yes | Cruel Barb (5191, +0.00 DPS) [dungeon]; Tork Wrench (11855, -9.35 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.1 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Deadly Blunderbuss (4369, -0.08 DPS) [crafted]; Light Bow (4576, -0.08 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-00000000000000000-500230131051000000)

Set DPS (verified): 84.9. Weights run: 3.4s. Verify run: 1.8s. 352 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.044 ± 0.008, strength=1.000 ± 0.001, crit=0.190 ± 0.005 per rating point (14 rating = 1%, 2.660 per %), hit=0.058 ± 0.001 per rating point (10 rating = 1%, 0.578 per %), melee_haste=not significant (0.523 ± 0.580)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.54 DPS) | yes | Brawler's Leather Helm (252512, -0.07 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.14 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.14 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.63 DPS) | yes | Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Scout's Medallion (19537, -0.27 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.5 attack_power points (0.75 DPS) | yes | Barbaric Shoulders (5964, -0.28 DPS) [crafted]; Mantle of Thieves (2264, -0.30 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.33 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.3 attack_power points (0.47 DPS) | yes | Wolfmaster Cape (6314, -0.01 DPS) [dungeon]; Wildhunter Cloak (16658, -0.01 DPS) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.6 attack_power points (0.66 DPS) | yes | Brawler's Leather Tunic (252508, -0.01 DPS) [crafted]; Brawler's Leather Armor (252490, -0.10 DPS) [crafted]; Defender's Leather Tunic (252450, -0.11 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.3 attack_power points (0.46 DPS) | yes | Cultist's Armguards (270032, -0.01 DPS) [quest]; Jurassic Wristguards (6198, -0.09 DPS) [world]; Barbaric Bracers (18948, -0.09 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.72 DPS) | yes | Insignia Gloves (6408, -0.12 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.21 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.09 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.27 DPS) [dungeon]; Deftkin Belt (16659, -0.35 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.18 DPS) | yes | Brawler's Leather Pants (252500, -0.48 DPS) [crafted]; Trapper's Leather Pants (252501, -0.48 DPS) [crafted]; Brawler's Leather Legguards (252516, -1.04 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.4 attack_power points (0.51 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.14 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.14 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.4 attack_power points (0.61 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Insurgent's Band (272067, -0.20 DPS) [vendor]; Band of the Fist (17694, -0.24 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.3 attack_power points (0.56 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Band of the Fist (17694, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 340.4 attack_power points (15.42 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (84.9 DPS) | yes | Tork Wrench (11855, -14.40 DPS) [quest]; Satyr's Rod (15962, -14.45 DPS) [world_drop]; Swinetusk Shank (6691, -20.54 DPS, sim-verified) [dungeon] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.41 DPS) | yes | Double-barreled Shotgun (2098, -0.13 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.22 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 352, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-00000000000000000-500230131051120151)

Set DPS (verified): 119.0. Weights run: 3.4s. Verify run: 1.8s. 563 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.178 ± 0.012, strength=1.000 ± 0.001, crit=0.234 ± 0.005 per rating point (14 rating = 1%, 3.282 per %), hit=0.068 ± 0.002 per rating point (10 rating = 1%, 0.677 per %), melee_haste=not significant (3.021 ± 0.977)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 24.0 attack_power points (1.24 DPS) | yes | Hawkeye's Helm (14591, -0.36 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.37 DPS, sim-verified) [crafted]; Raging Berserker's Helm (7719, -0.40 DPS) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.04 DPS) | yes | Scout's Medallion (19536, -0.37 DPS) [rep]; Ghostshard Talisman (7731, -0.47 DPS, sim-verified) [dungeon]; Ethereal Talisman (4430, -0.53 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.0 attack_power points (1.30 DPS) | yes | Flintrock Shoulders (7755, -0.42 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.55 DPS, sim-verified) [world_drop]; Barbaric Iron Shoulders (7913, -0.62 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 15.8 attack_power points (0.82 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Wildhunter Cloak (16658, -0.30 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 30.4 attack_power points (1.58 DPS) | yes | Wolffear Harness (13110, -0.54 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -0.66 DPS) [crafted]; Kolkar Marauder Chain (6773, -0.83 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Ravager's Armguards (14770, -0.58 DPS, sim-verified) [world_drop] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 attack_power points (1.66 DPS) | yes | Gloves of Holy Might (867, -0.45 DPS) [world_drop]; Scarlet Gauntlets (10331, -0.46 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.58 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.56 DPS) | yes | Defiler's Chain Girdle (20152, -0.47 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.54 DPS) [world_drop]; Blackened Defias Belt (10403, -0.62 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.35 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Scarlet Leggings (10330, -0.26 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.04 DPS, sim-verified) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 20.8 attack_power points (1.08 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.04 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.04 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Ironspine's Eye (7686, -0.28 DPS) [dungeon]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (24.65 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (22.96 DPS) | yes | Dazzling Longsword (869, -10.56 DPS, sim-verified) [world_drop]; Stonecloth Branch (15963, -22.81 DPS) [world_drop]; Tork Wrench (11855, -22.86 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (119.0 DPS) | yes | Monolithic Bow (9426, -0.23 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.26 DPS) [vendor]; Bow of Searing Arrows (2825, -1.10 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gauntlets of Divinity; waist: Defiler's Leather Girdle; feet: Blackforge Greaves; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-32005000000000000-500230131051120151)

Set DPS (verified): 166.5. Weights run: 3.4s. Verify run: 2.0s. 713 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.195 ± 0.016, strength=1.000 ± 0.001, crit=0.317 ± 0.008 per rating point (14 rating = 1%, 4.442 per %), hit=0.089 ± 0.002 per rating point (10 rating = 1%, 0.885 per %), melee_haste=2.785 ± 0.650

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.04 DPS) | yes | White Bandit Mask (10008, -0.81 DPS) [crafted]; Blood Guard's Chain Helmet (220821, -0.84 DPS) [vendor]; Bloomsprout Headpiece (17767, -1.11 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 20.5 attack_power points (1.05 DPS) | yes | Woven Ivy Necklace (19159, -0.19 DPS) [quest]; Scout's Medallion (19535, -0.32 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.94 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.1 attack_power points (1.29 DPS) | yes | Skulker's Leather Shoulder (252535, -0.15 DPS) [crafted]; Failed Flying Experiment (9647, -0.16 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.17 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.7 attack_power points (1.16 DPS) | yes | Blisterbane Wrap (12552, -0.25 DPS) [dungeon]; Dark Phantom Cape (13122, -0.25 DPS) [world_drop]; Dark Hooded Cape (5257, -0.35 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 32.5 attack_power points (1.66 DPS) | yes | Mixologist's Tunic (12793, -0.07 DPS) [dungeon]; Quillward Harness (10583, -0.09 DPS) [dungeon]; Blazewind Breastplate (11193, -0.10 DPS) [quest] |
| wrist | Arena Bands (18711) (or Bracers of the Stone Princess (17714)) | Arena Treasure Chest [world] | 28.0 attack_power points (1.43 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19583, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 36.3 attack_power points (1.85 DPS) | yes | Rockgrip Gauntlets (17736, -0.42 DPS) [dungeon]; Skulker's Leather Gauntlets (252548, -0.54 DPS) [crafted]; Gauntlets of Divinity (7724, -1.07 DPS, sim-verified) [dungeon] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (1.94 DPS) | yes | Substandard Belt Chain (274757, -0.36 DPS) [vendor]; Defiler's Leather Girdle (20192, -0.41 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.47 DPS) [crafted] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 32.1 attack_power points (1.64 DPS) | yes | Ferine Leggings (6690, +0.00 DPS, sim-verified) [dungeon]; Basilisk Hide Pants (1718, -0.36 DPS) [world_drop]; Triprunner Dungarees (9624, -0.39 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 26.5 attack_power points (1.36 DPS) | yes | Sandstalker Ankleguards (12470, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Albino Crocscale Boots (17728, -0.13 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.23 DPS) | yes | Legionnaire's Band (19511, -0.17 DPS) [rep]; Mark of Kern (2262, -0.20 DPS) [dungeon]; Assault Band (13095, -0.20 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.9 attack_power points (1.07 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Legionnaire's Band (19511, -1.62 DPS, sim-verified) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (166.5 DPS) | yes | Frozen Heart of the Mountain (249469, -2.14 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (166.5 DPS) | yes | Frozen Heart of the Mountain (249469, -2.85 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (166.5 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor]; Hanzo Sword (8190, -15.23 DPS, sim-verified) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (26.91 DPS) | yes | Claw of Celebras (17738, -2.26 DPS) [dungeon]; White Bone Shredder (11863, -4.04 DPS) [quest]; Grizzle's Skinner (11702, -20.66 DPS, sim-verified) [dungeon] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (166.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.01 DPS) [world_drop]; Dark Iron Rifle (16004, -1.61 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; back: Blackveil Cape; chest: Warbear Harness; wrist: Arena Bands; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Bloodrazor; off_hand: Thorium Cestus; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 713, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-32005500005000000-500230131051120151)

Set DPS (verified): 234.2. Weights run: 3.4s. Verify run: 2.0s. 1650 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.209 ± 0.020, strength=1.000 ± 0.001, crit=0.475 ± 0.013 per rating point (14 rating = 1%, 6.656 per %), hit=0.117 ± 0.004 per rating point (10 rating = 1%, 1.172 per %), melee_haste=not significant (1.596 ± 0.951)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crown of Tyranny (13359) | Stratholme: Balnazzar [dungeon] | 40.5 attack_power points (2.04 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.18 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 33.0 attack_power points (1.66 DPS) | yes | Imperial Jewel (11933, -0.05 DPS) [dungeon]; Medallion of the Dawn (22659, -0.12 DPS) [quest]; Will of the Martyr (17044, -0.15 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 51.8 attack_power points (2.61 DPS) | yes | Defiler's Lizardhide Shoulders (20175, +0.00 DPS, sim-verified) [rep]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Champion's Chain Pauldrons (227078, -0.26 DPS) [pvp] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 40.0 attack_power points (2.02 DPS) | yes | Cape of the Black Baron (13340, -0.10 DPS) [dungeon]; Howler's Furs (272414, -0.55 DPS) [vendor]; Windshear Cape (20691, -0.70 DPS) [world] |
| chest | Obsidian Mail Tunic (22191) | Blacksmithing [crafted] | sim-verified (234.2 DPS) | yes | Warlord's Chain Armor (231566, -0.07 DPS) [vendor]; Cadaverous Armor (14637, -0.25 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.99 DPS, sim-verified) [world] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | sim-verified (234.2 DPS) | yes | Bracers of the Eclipse (18375, -0.10 DPS) [dungeon]; Forest Stalker's Bracers (19587, -0.20 DPS) [rep]; Wristwraps of Undead Slaying (23093, -6.57 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (2.22 DPS) | yes | General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor]; Skul's Fingerbone Claws (13395, -0.20 DPS) [dungeon] |
| waist | Dense Timbermaw Belt (227807) | Meilosh [vendor] | 64.0 attack_power points (3.23 DPS) | yes | Ferocity of the Timbermaw (227805, -0.88 DPS) [vendor]; Chiselbrand Girdle (12634, -1.01 DPS) [dungeon]; Heavy Timbermaw Belt (19043, -1.11 DPS) [crafted] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 56.8 attack_power points (2.87 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -0.12 DPS) [vendor]; Black Dragonscale Leggings (15052, -7.91 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 47.7 attack_power points (2.40 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; Blood Guard's Chain Sabatons (227082, -0.18 DPS) [vendor]; Pads of the Dread Wolf (13210, -1.11 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (234.2 DPS) | yes | Don Julio's Band (19325, -0.07 DPS) [rep]; Signet Ring of the Bronze Dragonflight (21201, -0.12 DPS) [quest]; Naglering (11669, -3.03 DPS, sim-verified) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | sim-verified (234.2 DPS) | yes | Don Julio's Band (19325, -0.01 DPS) [rep]; Signet Ring of the Bronze Dragonflight (21201, -0.05 DPS) [quest]; Naglering (11669, -2.46 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (234.2 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (234.2 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -3.69 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (234.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -26.97 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 827.2 attack_power points (41.73 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -3.28 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -12.50 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (234.2 DPS) | yes | Malgen's Long Bow (22318, -0.10 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.25 DPS) [world_drop]; Dark Iron Rifle (16004, -2.26 DPS, sim-verified) [crafted] |

**New at 60:** head: Crown of Tyranny; neck: Amulet of the Darkmoon; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Obsidian Mail Tunic; wrist: Windtalker's Wristguards; hands: Cadaverous Gloves; waist: Dense Timbermaw Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Second Wind; trinket2: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Riphook

No-known-source sample (15 of 1650, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

