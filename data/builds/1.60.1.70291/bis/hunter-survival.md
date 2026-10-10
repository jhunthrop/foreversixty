# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.5. Weights run: 2.9s. Verify run: 3.8s. 221 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.152 ± 0.014, strength=1.210 ± 0.002, crit=0.853 ± 0.016 per rating point (14 rating = 1%, 11.936 per %), hit=1.143 ± 0.038 per rating point (10 rating = 1%, 11.434 per %), melee_haste=not significant (2.968 ± 1.264)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.2 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.9 attack_power points (0.54 DPS) | yes | Erudite's Amulet (277204, -0.35 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.8 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.45 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.9 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 23.7 attack_power points (1.00 DPS) | yes | Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Prospector's Chestpiece (14562, -0.35 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.36 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 11.0 attack_power points (0.46 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.10 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.5 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.03 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Brawler's Leather Belt (252428, -0.19 DPS) [crafted]; Dusty Belt (279897, -0.31 DPS) [quest]; Deviate Scale Belt (6468, -2.70 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.18 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.09 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.4 attack_power points (0.57 DPS) | yes | Demon Band (12054, -0.36 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; The 1 Ring (8350, -0.42 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.9 attack_power points (0.54 DPS) | yes | Demon Band (12054, -0.31 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; The 1 Ring (8350, -0.40 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.49 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.9 attack_power points (9.98 DPS) | yes | Assassin's Blade (1935, -0.38 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.6 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Owlsight Rifle (15205, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 221, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 97.0. Weights run: 3.2s. Verify run: 2.9s. 395 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.866 ± 0.012, strength=1.210 ± 0.002, crit=0.928 ± 0.017 per rating point (14 rating = 1%, 12.996 per %), hit=1.427 ± 0.049 per rating point (10 rating = 1%, 14.269 per %), melee_haste=8.219 ± 1.242

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.7 attack_power points (0.93 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.21 DPS) [crafted] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 14.9 attack_power points (0.75 DPS) | yes | Ghostshard Talisman (7731, -0.05 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Wolfpack Medallion (5754, -0.37 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.6 attack_power points (1.33 DPS) | yes | Mantle of Thieves (2264, -0.40 DPS) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.7 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp]; Cloak of Night (4447, -0.28 DPS) [world] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 26.1 attack_power points (1.31 DPS) | yes | Brawler's Leather Tunic (252508, -0.20 DPS) [crafted]; Tunic of Westfall (2041, -0.28 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 16.0 attack_power points (0.80 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.24 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.7 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Gloves of the Fang (10413, -0.18 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 24.1 attack_power points (1.21 DPS) | yes | Highlander's Chain Girdle (20090, -0.00 DPS) [rep]; Highlander's Leather Girdle (20117, -0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.10 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 26.1 attack_power points (1.31 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -0.01 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.04 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.6 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.6 attack_power points (1.08 DPS) | yes | Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Monkey Ring (6748, -0.43 DPS) [quest]; Ring of Precision (1491, -0.52 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.5 attack_power points (0.92 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.27 DPS) [quest]; Ring of Precision (1491, -0.36 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.11 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (97.0 DPS) | yes | Shoni's Disarming Tool (9608, -4.78 DPS) [quest]; Satyr's Rod (15962, -15.94 DPS) [world_drop]; Swinetusk Shank (6691, -16.40 DPS, sim-verified) [dungeon] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.60 DPS) | yes | Golemsight Long Gun (273029, -0.04 DPS) [dungeon]; Silver Star (3463, -0.13 DPS) [quest]; Double-barreled Shotgun (2098, -0.14 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Scorn's Focal Dagger; ranged: Glass Shooter

No-known-source sample (15 of 395, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 141.9. Weights run: 3.3s. Verify run: 2.9s. 628 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.860 ± 0.011, strength=1.210 ± 0.001, crit=0.936 ± 0.016 per rating point (14 rating = 1%, 13.109 per %), hit=1.457 ± 0.073 per rating point (10 rating = 1%, 14.573 per %), melee_haste=9.860 ± 1.098

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.8 attack_power points (2.22 DPS) | yes | Barbaric Iron Helm (7915, -0.40 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.56 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.63 DPS) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 20.5 attack_power points (1.35 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.03 DPS) [quest]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.54 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.5 attack_power points (2.14 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.51 DPS) [dungeon]; Nightscape Shoulders (8192, -0.79 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.4 attack_power points (1.54 DPS) | yes | Sergeant Major's Cape (16336, -0.33 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Parachute Cloak (10518, -0.56 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.32 DPS) | yes | Hawkeye's Bracers (14590, -0.26 DPS) [world_drop]; Ravager's Armguards (14770, -0.27 DPS) [world_drop]; Dusky Bracers (7378, -0.34 DPS) [crafted] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 33.1 attack_power points (2.18 DPS) | yes | Gloves of Holy Might (867, -0.00 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.07 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.24 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.97 DPS) | yes | Ogron's Sash (13117, -0.16 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.39 DPS) [rep]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 39.1 attack_power points (2.57 DPS) | yes | Triprunner Dungarees (9624, -0.13 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.84 DPS) [crafted]; Hawkeye's Breeches (14595, -0.86 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.5 attack_power points (1.94 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.6 attack_power points (1.62 DPS) | yes | Ring of the Underwood (2951, -0.23 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Mark of Kern (2262, -0.30 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.6 attack_power points (1.42 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.34 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.6 attack_power points (29.45 DPS) | yes | Vanquisher's Sword (10823, +0.00 DPS) [quest]; Shoni's Disarming Tool (9608, -14.66 DPS) [quest]; Stonecloth Branch (15963, -29.21 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.06 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 628, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 182.3. Weights run: 3.2s. Verify run: 3.4s. 787 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.781 ± 0.010, strength=1.210 ± 0.001, crit=1.016 ± 0.017 per rating point (14 rating = 1%, 14.218 per %), hit=1.478 ± 0.081 per rating point (10 rating = 1%, 14.778 per %), melee_haste=9.527 ± 1.231

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 42.7 attack_power points (3.11 DPS) | yes | Bloomsprout Headpiece (17767, -0.49 DPS) [dungeon]; White Bandit Mask (10008, -0.71 DPS) [crafted]; Embrace of the Lycan (9479, -1.60 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.2 attack_power points (2.13 DPS) | yes | Sentinel's Medallion (19539, -0.57 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.67 DPS) [quest] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 37.4 attack_power points (2.72 DPS) | yes | Sunburn Spaulders (274751, -0.42 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.50 DPS) [crafted]; Failed Flying Experiment (9647, -0.54 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (182.3 DPS) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Dark Hooded Cape (5257, -0.30 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 45.4 attack_power points (3.31 DPS) | yes | Blazewind Breastplate (11193, -0.06 DPS) [quest]; Fungus Shroud Armor (17742, -0.06 DPS) [dungeon]; Quillward Harness (10583, -0.13 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.6 attack_power points (2.30 DPS) | yes | Bracers of the Stone Princess (17714, -0.26 DPS) [dungeon]; Arena Bands (18711, -0.26 DPS) [world]; Skulker's Leather Bracers (252540, -0.51 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 51.3 attack_power points (3.74 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.60 DPS, sim-verified) [crafted]; Gloves of Holy Might (867, -1.24 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -1.33 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 39.7 attack_power points (2.89 DPS) | yes | Girdle of Beastial Fury (11686, -0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 44.9 attack_power points (3.27 DPS) | yes | Serpentskin Leggings (8262, -0.05 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.42 DPS) [vendor]; Basilisk Hide Pants (1718, -0.54 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 37.5 attack_power points (2.73 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.16 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 28.1 attack_power points (2.05 DPS) | yes | Ironspine's Eye (7686, -0.53 DPS) [dungeon]; Ring of the Underwood (2951, -0.58 DPS) [world_drop]; Blackstone Ring (17713, -0.59 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.9 attack_power points (1.82 DPS) | yes | Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.34 DPS) [world_drop]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -0.95 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (+30.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.23 DPS) [dungeon]; Shoni's Disarming Tool (9608, -21.98 DPS) [quest]; Grizzle's Skinner (11702, -30.74 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.51 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop]; Dark Iron Rifle (16004, -1.46 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; legs: Gryphon Rider's Leggings; feet: Sandstalker Ankleguards; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 787, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 289.3. Weights run: 3.1s. Verify run: 11.1s. 1685 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.196 ± 0.025, strength=1.210 ± 0.002, crit=2.432 ± 0.049 per rating point (14 rating = 1%, 34.047 per %), hit=3.471 ± 0.225 per rating point (10 rating = 1%, 34.709 per %), melee_haste=not significant (12.617 ± 3.461)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -5.82 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Medallion of the Dawn (22659, -0.05 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -3.85 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 62.7 attack_power points (4.83 DPS) | yes | Cape of the Black Baron (13340, -0.75 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.37 DPS) [rep]; Windshear Cape (20691, -1.55 DPS) [world] |
| chest | Beastmaster's Tunic (226886) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Chain Armor (227089, +0.00 DPS) [vendor]; Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Dawn Armor (252483, -5.04 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -6.80 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Beaststalker's Gloves (16676, -3.36 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -4.68 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 179.7 attack_power points (13.84 DPS) | yes | Marshal's Chain Legplates (231558, -3.89 DPS) [vendor]; Sentinel's Leather Pants (237818, -4.03 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -4.52 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -6.30 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.52 DPS) [dungeon]; Cutthroat's Signet (272408, -1.69 DPS) [vendor]; Naglering (11669, -6.98 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.32 DPS) [dungeon]; Cutthroat's Signet (272408, -1.49 DPS) [vendor]; Naglering (11669, -5.97 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -6.32 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -0.71 DPS) [dungeon]; Hand of Justice (11815, -0.87 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -16.67 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.8 attack_power points (64.60 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -7.13 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -19.97 DPS) [dungeon] |
| ranged | Dark Iron Rifle (16004) | Engineering [crafted] | sim-verified (289.3 DPS) | yes | Precisely Calibrated Boomstick (2100, +0.00 DPS) [world_drop]; Satyr's Bow (18323, +0.00 DPS) [dungeon]; The Purifier (22656, +0.00 DPS) [quest] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Beastmaster's Tunic; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Dark Iron Rifle

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 765.3. Weights run: 2.7s. Verify run: 8.8s. 1685 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.408 ± 0.029, strength=1.210 ± 0.002, crit=2.999 ± 0.059 per rating point (14 rating = 1%, 41.984 per %), hit=4.997 ± 0.425 per rating point (10 rating = 1%, 49.969 per %), melee_haste=not significant (15.586 ± 5.412)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 113.5 attack_power points (18.33 DPS) | yes | Beastmaster's Cap (226887, +0.00 DPS) [quest]; Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.97 DPS) [quest]; Medallion of the Dawn (22659, -1.29 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 78.0 attack_power points (12.59 DPS) | yes | Cape of the Black Baron (13340, -3.53 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.52 DPS) [vendor]; Stalwart Cloak (272415, -4.52 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Chain Armor (231563, -6.51 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -6.80 DPS) [vendor]; Tunic of Undead Slaying (23089, -24.69 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.49 DPS) [dungeon]; Bracers of Subterfuge (22668, -1.75 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 103.0 attack_power points (16.63 DPS) | yes | Marksman's Girdle (22232, -0.40 DPS) [dungeon]; Highlander's Chain Girdle (20043, -4.37 DPS) [rep]; Highlander's Leather Girdle (20045, -4.37 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 218.2 attack_power points (35.23 DPS) | yes | Sentinel's Leather Pants (237818, -11.18 DPS) [vendor]; Marshal's Chain Legplates (231558, -12.51 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -12.80 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.53 DPS) [dungeon]; Cutthroat's Signet (272408, -3.92 DPS) [vendor]; Naglering (11669, -11.82 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.50 DPS) [dungeon]; Cutthroat's Signet (272408, -3.89 DPS) [vendor]; Naglering (11669, -14.02 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (+11.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -6.30 DPS) [crafted]; Counterattack Lodestone (18537, -10.00 DPS) [dungeon]; Hand of Justice (11815, -10.33 DPS) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -5.16 DPS, sim-verified) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -62.36 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 840.7 attack_power points (135.73 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -10.21 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -42.18 DPS) [dungeon] |
| ranged | Dark Iron Rifle (16004) | Engineering [crafted] | sim-verified (765.3 DPS) | yes | Precisely Calibrated Boomstick (2100, +0.00 DPS) [world_drop]; Satyr's Bow (18323, +0.00 DPS) [dungeon]; The Purifier (22656, +0.00 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Blackhand's Breadth; trinket2: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Dark Iron Rifle

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.9. Weights run: 2.9s. Verify run: 3.7s. 210 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.152 ± 0.014, strength=1.210 ± 0.002, crit=0.853 ± 0.016 per rating point (14 rating = 1%, 11.936 per %), hit=1.143 ± 0.038 per rating point (10 rating = 1%, 11.434 per %), melee_haste=not significant (2.968 ± 1.264)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.2 attack_power points (0.73 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.9 attack_power points (0.54 DPS) | yes | Erudite's Amulet (277204, -0.37 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.8 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.46 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.9 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.17 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (73.9 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.94 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.8 attack_power points (0.45 DPS) | yes | Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.18 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.8 attack_power points (0.75 DPS) | yes | Bristlebark Gloves (14572, -0.18 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.18 DPS) [crafted]; Serpent Gloves (5970, -0.20 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.76 DPS) | yes | Brawler's Leather Belt (252428, -0.19 DPS) [crafted]; Murloc Scale Belt (5780, -0.38 DPS) [crafted]; Deviate Scale Belt (6468, -2.60 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.08 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.99 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.4 attack_power points (0.57 DPS) | yes | Bounty Hunter's Ring (5351, -0.29 DPS) [quest]; Demon Band (12054, -0.36 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.9 attack_power points (0.54 DPS) | yes | Bounty Hunter's Ring (5351, -0.28 DPS, sim-verified) [quest]; Demon Band (12054, -0.34 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.49 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; The Axe of Severing (23171, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.9 attack_power points (9.98 DPS) | yes | Wingblade (6504, -6.39 DPS, sim-verified) [quest]; Tork Wrench (11855, -9.88 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.6 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Privateer Musket (5309, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 210, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 99.2. Weights run: 3.2s. Verify run: 2.2s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.866 ± 0.012, strength=1.210 ± 0.002, crit=0.928 ± 0.017 per rating point (14 rating = 1%, 12.996 per %), hit=1.427 ± 0.049 per rating point (10 rating = 1%, 14.269 per %), melee_haste=8.219 ± 1.242

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.7 attack_power points (0.93 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.21 DPS) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 14.9 attack_power points (0.75 DPS) | yes | Ghostshard Talisman (7731, -0.05 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Wolfpack Medallion (5754, -0.37 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.6 attack_power points (1.33 DPS) | yes | Mantle of Thieves (2264, -0.40 DPS) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.7 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 26.1 attack_power points (1.31 DPS) | yes | Brawler's Leather Tunic (252508, -0.20 DPS) [crafted]; Panther Armor (6670, -0.35 DPS) [quest]; Brawler's Leather Armor (252490, -0.35 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 16.0 attack_power points (0.80 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.24 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.7 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Gloves of the Fang (10413, -0.18 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 24.1 attack_power points (1.21 DPS) | yes | Defiler's Chain Girdle (20152, -0.00 DPS) [rep]; Defiler's Leather Girdle (20191, -0.00 DPS) [rep]; Prowler's Leather Belt (252459, -0.10 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Burrower [dungeon] | 26.1 attack_power points (1.31 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -0.01 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.04 DPS) [crafted] |
| feet | Vorrel's Boots (7751) | Vorrel's Revenge [quest] | 18.7 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Feet of the Lynx (1121, -0.28 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.6 attack_power points (1.08 DPS) | yes | Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Monkey Ring (6748, -0.43 DPS) [quest]; Band of the Fist (17694, -0.47 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.5 attack_power points (0.92 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.27 DPS) [quest]; Band of the Fist (17694, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.11 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | sim-verified (99.2 DPS) | yes | Tork Wrench (11855, -15.97 DPS) [quest]; Satyr's Rod (15962, -16.00 DPS) [world_drop]; Swinetusk Shank (6691, -17.87 DPS, sim-verified) [dungeon] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.60 DPS) | yes | Golemsight Long Gun (273029, -0.04 DPS) [dungeon]; Silver Star (3463, -0.13 DPS) [quest]; Double-barreled Shotgun (2098, -0.14 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Vorrel's Boots; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Serrated Raptor Claw; ranged: Glass Shooter

No-known-source sample (15 of 380, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 142.3. Weights run: 3.3s. Verify run: 2.8s. 593 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.860 ± 0.011, strength=1.210 ± 0.001, crit=0.936 ± 0.016 per rating point (14 rating = 1%, 13.109 per %), hit=1.457 ± 0.073 per rating point (10 rating = 1%, 14.573 per %), melee_haste=9.860 ± 1.098

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.8 attack_power points (2.22 DPS) | yes | Barbaric Iron Helm (7915, -0.40 DPS) [crafted]; Hawkeye's Helm (14591, -0.56 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.63 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 20.5 attack_power points (1.35 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.03 DPS) [quest]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon]; Ethereal Talisman (4430, -0.46 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.5 attack_power points (2.14 DPS) | yes | Forest Tracker Epaulets (2278, -0.43 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.51 DPS) [dungeon]; Nightscape Shoulders (8192, -0.79 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.4 attack_power points (1.54 DPS) | yes | First Sergeant's Cloak (16340, -0.33 DPS) [pvp]; Hawkeye's Cloak (14593, -0.45 DPS) [world_drop]; Parachute Cloak (10518, -0.56 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (142.3 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Kolkar Marauder Chain (6773, -0.11 DPS) [quest]; Tough Scorpid Breastplate (8203, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.32 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.46 DPS, sim-verified) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 33.1 attack_power points (2.18 DPS) | yes | Gloves of Holy Might (867, -0.00 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.07 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.24 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.97 DPS) | yes | Ogron's Sash (13117, -0.16 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.39 DPS) [rep]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 39.1 attack_power points (2.57 DPS) | yes | Triprunner Dungarees (9624, -0.13 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.84 DPS) [crafted]; Hawkeye's Breeches (14595, -0.86 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.5 attack_power points (1.94 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.6 attack_power points (1.62 DPS) | yes | Ring of the Underwood (2951, -0.23 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Mark of Kern (2262, -0.30 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.6 attack_power points (1.42 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.81 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.6 attack_power points (29.45 DPS) | yes | Vanquisher's Sword (10823, -0.60 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -29.21 DPS) [world_drop]; Tork Wrench (11855, -29.29 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.06 DPS) [world_drop]; Monolithic Bow (9426, -0.08 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 593, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 184.5. Weights run: 3.2s. Verify run: 3.3s. 745 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.781 ± 0.010, strength=1.210 ± 0.001, crit=1.016 ± 0.017 per rating point (14 rating = 1%, 14.218 per %), hit=1.478 ± 0.081 per rating point (10 rating = 1%, 14.778 per %), melee_haste=9.527 ± 1.231

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Bloomsprout Headpiece (17767, +0.00 DPS) [dungeon]; Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.2 attack_power points (2.13 DPS) | yes | Woven Ivy Necklace (19159, -0.43 DPS) [quest]; Scout's Medallion (19535, -0.57 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.67 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 31.6 attack_power points (2.30 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Failed Flying Experiment (9647, -0.12 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (184.5 DPS) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Dark Hooded Cape (5257, -0.30 DPS) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 45.4 attack_power points (3.31 DPS) | yes | Blazewind Breastplate (11193, -0.06 DPS) [quest]; Fungus Shroud Armor (17742, -0.06 DPS) [dungeon]; Quillward Harness (10583, -0.13 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.6 attack_power points (2.30 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 51.3 attack_power points (3.74 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.93 DPS, sim-verified) [crafted]; Gloves of Holy Might (867, -1.24 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -1.33 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 39.7 attack_power points (2.89 DPS) | yes | Girdle of Beastial Fury (11686, -0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Basilisk Hide Pants (1718, -0.13 DPS) [world_drop]; Triprunner Dungarees (9624, -0.25 DPS) [quest]; Serpentskin Leggings (8262, -2.18 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 37.5 attack_power points (2.73 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.16 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 28.1 attack_power points (2.05 DPS) | yes | White Bone Band (11862, -0.30 DPS) [quest]; Ironspine's Eye (7686, -0.53 DPS) [dungeon]; Ring of the Underwood (2951, -0.58 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.9 attack_power points (1.82 DPS) | yes | White Bone Band (11862, -0.07 DPS) [quest]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.34 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, -2.53 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.84 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodrazor (809, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (+31.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of Celebras (17738, -3.23 DPS) [dungeon]; White Bone Shredder (11863, -5.40 DPS) [quest]; Grizzle's Skinner (11702, -31.94 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.51 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop]; Dark Iron Rifle (16004, -1.18 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; legs: Stone Guard's Chain Legplates; feet: Sandstalker Ankleguards; finger1: Legionnaire's Band; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Shadowblade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 745, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 296.5. Weights run: 3.1s. Verify run: 10.7s. 1664 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.196 ± 0.025, strength=1.210 ± 0.002, crit=2.432 ± 0.049 per rating point (14 rating = 1%, 34.047 per %), hit=3.471 ± 0.225 per rating point (10 rating = 1%, 34.709 per %), melee_haste=not significant (12.617 ± 3.461)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -4.73 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Medallion of the Dawn (22659, -0.05 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -1.90 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 62.7 attack_power points (4.83 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.37 DPS) [rep]; Windshear Cape (20691, -1.55 DPS) [world] |
| chest | Beastmaster's Tunic (226886) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Chain Armor (227083, +0.00 DPS) [vendor]; Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Dawn Armor (252483, -4.44 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -6.63 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -3.56 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Beaststalker's Belt (16680, -4.48 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 179.7 attack_power points (13.84 DPS) | yes | Outrider's Chain Leggings (22673, -1.46 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -3.89 DPS) [vendor]; Sentinel's Leather Pants (237818, -4.03 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.59 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.52 DPS) [dungeon]; Cutthroat's Signet (272408, -1.69 DPS) [vendor]; Naglering (11669, -5.49 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.32 DPS) [dungeon]; Cutthroat's Signet (272408, -1.49 DPS) [vendor]; Naglering (11669, -4.89 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (296.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.70 DPS) [crafted]; Counterattack Lodestone (18537, -3.41 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -13.95 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.8 attack_power points (64.60 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -5.43 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -19.97 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; The Purifier (22656, -0.56 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.81 DPS) [world_drop] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Beastmaster's Tunic; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1664, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 776.3. Weights run: 2.7s. Verify run: 8.4s. 1664 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.408 ± 0.029, strength=1.210 ± 0.002, crit=2.999 ± 0.059 per rating point (14 rating = 1%, 41.984 per %), hit=4.997 ± 0.425 per rating point (10 rating = 1%, 49.969 per %), melee_haste=not significant (15.586 ± 5.412)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 113.5 attack_power points (18.33 DPS) | yes | Beastmaster's Cap (226887, +0.00 DPS) [quest]; Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Mark of Fordring (15411, -0.97 DPS) [quest]; Medallion of the Dawn (22659, -1.29 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 78.0 attack_power points (12.59 DPS) | yes | Cape of the Black Baron (13340, -3.53 DPS) [dungeon]; Arcanoweave Cloak (272411, -4.52 DPS) [vendor]; Stalwart Cloak (272415, -4.52 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Chain Armor (231566, -6.51 DPS) [vendor]; Legionnaire's Chain Armor (227083, -6.80 DPS) [vendor]; Tunic of Undead Slaying (23089, -25.69 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.49 DPS) [dungeon]; Bracers of Subterfuge (22668, -1.75 DPS) [quest] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 103.0 attack_power points (16.63 DPS) | yes | Marksman's Girdle (22232, -0.40 DPS) [dungeon]; Defiler's Chain Girdle (20150, -4.37 DPS) [rep]; Defiler's Leather Girdle (20190, -4.37 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 218.2 attack_power points (35.23 DPS) | yes | Outrider's Chain Leggings (22673, -4.45 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -11.18 DPS) [vendor]; General's Chain Legplates (231567, -12.51 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.53 DPS) [dungeon]; Cutthroat's Signet (272408, -3.92 DPS) [vendor]; Naglering (11669, -15.38 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -3.50 DPS) [dungeon]; Cutthroat's Signet (272408, -3.89 DPS) [vendor]; Naglering (11669, -13.24 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -21.45 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (776.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.17 DPS) [crafted]; Counterattack Lodestone (18537, -8.88 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -67.30 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 840.7 attack_power points (135.73 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -11.59 DPS, sim-verified) [crafted]; Dal'Rend's Tribal Guardian (12939, -42.18 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Purifier (22656, -2.46 DPS) [quest]; Precisely Calibrated Boomstick (2100, -3.79 DPS) [world_drop]; Dark Iron Rifle (16004, -5.68 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: Satyr's Bow

No-known-source sample (15 of 1664, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

