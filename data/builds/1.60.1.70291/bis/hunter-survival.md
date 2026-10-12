# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.5. Weights run: 1.5s. Verify run: 4.2s. 212 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.045 ± 0.022, strength=1.210 ± 0.003, crit=0.748 ± 0.023 per rating point (14 rating = 1%, 10.475 per %), hit=1.112 ± 0.058 per rating point (10 rating = 1%, 11.118 per %), melee_haste=8.258 ± 1.893

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.4 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.26 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.54 DPS) | yes | Erudite's Amulet (277204, -0.35 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.2 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.45 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.3 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.16 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.5 attack_power points (0.98 DPS) | yes | Brawler's Leather Armor (252490, -0.09 DPS) [crafted]; Prospector's Chestpiece (14562, -0.34 DPS) [world_drop]; Defender's Leather Armor (252434, -0.34 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.6 attack_power points (0.46 DPS) | yes | Forest Leather Bracers (3202, -0.02 DPS) [world_drop]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.5 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.03 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.22 DPS) [crafted]; Dusty Belt (279897, -0.34 DPS) [quest]; Deviate Scale Belt (6468, -2.70 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.18 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.09 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 13.0 attack_power points (0.57 DPS) | yes | Demon Band (12054, -0.36 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; The 1 Ring (8350, -0.43 DPS) [world] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.3 attack_power points (0.54 DPS) | yes | Demon Band (12054, -0.31 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.88 DPS) | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (10.34 DPS) | yes | Cruel Barb (5191, -0.92 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.2 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Owlsight Rifle (15205, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 212, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 98.6. Weights run: 1.4s. Verify run: 2.6s. 380 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.818 ± 0.019, strength=1.210 ± 0.002, crit=0.862 ± 0.027 per rating point (14 rating = 1%, 12.067 per %), hit=1.190 ± 0.071 per rating point (10 rating = 1%, 11.903 per %), melee_haste=not significant (5.793 ± 1.801)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.2 attack_power points (0.93 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.19 DPS) [crafted] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 14.5 attack_power points (0.75 DPS) | yes | Ghostshard Talisman (7731, -0.03 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop]; Wolfpack Medallion (5754, -0.37 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.1 attack_power points (1.33 DPS) | yes | Mantle of Thieves (2264, -0.33 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.4 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp]; Cloak of Night (4447, -0.28 DPS) [world] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.5 attack_power points (1.30 DPS) | yes | Brawler's Leather Tunic (252508, -0.19 DPS) [crafted]; Tunic of Westfall (2041, -0.28 DPS) [quest]; Brawler's Leather Armor (252490, -0.34 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.8 attack_power points (0.81 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.4 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.23 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.02 DPS) [crafted]; Prowler's Leather Belt (252459, -0.11 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.33 DPS) | yes | Troll's Bane Leggings (13114, -0.03 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.06 DPS) [crafted]; Petrolspill Leggings (9509, -1.26 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.2 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.2 attack_power points (1.09 DPS) | yes | Thunderbrow Ring (13097, -0.31 DPS) [world_drop]; Monkey Ring (6748, -0.43 DPS) [quest]; Ring of Precision (1491, -0.53 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.2 attack_power points (0.93 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.28 DPS) [quest]; Ring of Precision (1491, -0.37 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.49 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (98.6 DPS) | yes | Shoni's Disarming Tool (9608, -4.79 DPS) [quest]; Swinetusk Shank (6691, -16.98 DPS, sim-verified) [dungeon] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.61 DPS) | yes | Golemsight Long Gun (273029, -0.06 DPS) [dungeon]; Silver Star (3463, -0.15 DPS) [quest]; Double-barreled Shotgun (2098, -0.15 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Glass Shooter

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 141.9. Weights run: 1.5s. Verify run: 3.1s. 594 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.834 ± 0.016, strength=1.210 ± 0.002, crit=0.886 ± 0.026 per rating point (14 rating = 1%, 12.408 per %), hit=1.534 ± 0.101 per rating point (10 rating = 1%, 15.343 per %), melee_haste=7.052 ± 1.532

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.5 attack_power points (2.38 DPS) | yes | Barbaric Iron Helm (7915, -0.55 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.60 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.69 DPS) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 20.2 attack_power points (1.43 DPS) | yes | Ghostshard Talisman (7731, -0.44 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Wolfpack Medallion (5754, -0.91 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.2 attack_power points (2.29 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.55 DPS) [dungeon]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.2 attack_power points (1.65 DPS) | yes | Sergeant Major's Cape (16336, -0.35 DPS) [pvp]; Hawkeye's Cloak (14593, -0.48 DPS) [world_drop]; Parachute Cloak (10518, -0.61 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 31.2 attack_power points (2.22 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Nightscape Tunic (8175, -0.26 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.26 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.42 DPS) | yes | Hawkeye's Bracers (14590, -0.30 DPS) [world_drop]; Ravager's Armguards (14770, -0.30 DPS) [world_drop]; Dusky Bracers (7378, -0.38 DPS) [crafted] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 32.9 attack_power points (2.34 DPS) | yes | Gloves of Holy Might (867, -0.03 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.26 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.13 DPS) | yes | Ogron's Sash (13117, -0.18 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.43 DPS) [rep]; Skulker's Leather Belt (252520, -0.44 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.5 attack_power points (2.74 DPS) | yes | Triprunner Dungarees (9624, -0.13 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.88 DPS) [crafted]; Ferine Leggings (6690, -0.89 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.2 attack_power points (2.08 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.22 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 24.4 attack_power points (1.73 DPS) | yes | Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.30 DPS) [dungeon]; Mark of Kern (2262, -0.31 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.3 attack_power points (1.52 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Falcon's Hook (7552, -0.09 DPS) [dungeon]; Mark of Kern (2262, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.26 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.4 attack_power points (31.82 DPS) | yes | Vanquisher's Sword (10823, +0.00 DPS) [quest]; Shoni's Disarming Tool (9608, -15.84 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (141.9 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.09 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 594, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 50 (dwarf, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 189.6. Weights run: 1.5s. Verify run: 6.3s. 745 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.754 ± 0.016, strength=1.210 ± 0.002, crit=0.958 ± 0.027 per rating point (14 rating = 1%, 13.414 per %), hit=1.272 ± 0.115 per rating point (10 rating = 1%, 12.722 per %), melee_haste=9.759 ± 1.824

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 41.5 attack_power points (3.25 DPS) | yes | Helm of Fire (8348, -0.91 DPS) [crafted]; Scorpashi Skullcap (14658, -1.04 DPS) [world_drop]; White Bandit Mask (10008, -5.71 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.9 attack_power points (2.26 DPS) | yes | Sentinel's Medallion (19540, -0.75 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 36.2 attack_power points (2.84 DPS) | yes | Skulker's Leather Shoulder (252535, -0.47 DPS) [crafted]; Failed Flying Experiment (9647, -0.52 DPS) [quest]; Sunburn Spaulders (274751, -4.07 DPS, sim-verified) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.3 attack_power points (2.06 DPS) | yes | Dark Hooded Cape (5257, -0.31 DPS) [world]; Blackflame Cape (13109, -0.41 DPS) [world_drop]; Pridelord Cape (14673, -0.54 DPS) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-verified (189.6 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS) [quest]; Charred Leather Tunic (19127, -0.64 DPS) [quest]; Warbear Harness (15064, -4.75 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.2 attack_power points (2.44 DPS) | yes | Bracers of the Stone Princess (17714, -0.25 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.54 DPS) [crafted]; Arena Bands (18711, -0.71 DPS, sim-verified) [world] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (189.6 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Raider Gloves (272100, -0.74 DPS, sim-verified) [vendor] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 39.3 attack_power points (3.08 DPS) | yes | Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.46 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (189.6 DPS) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Basilisk Hide Pants (1718, -0.09 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -4.61 DPS, sim-verified) [quest] |
| feet | Sergeant Major's Chain Sabatons (220837) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (189.6 DPS) | yes | Albino Crocscale Boots (17728, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Sandstalker Ankleguards (12470, -1.65 DPS, sim-verified) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 27.9 attack_power points (2.18 DPS) | yes | Ironspine's Eye (7686, -0.57 DPS) [dungeon]; Mark of Kern (2262, -0.62 DPS) [dungeon]; Blackstone Ring (17713, -0.62 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.6 attack_power points (1.92 DPS) | yes | Ironspine's Eye (7686, -0.31 DPS) [dungeon]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (189.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (189.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (189.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thrash Blade (17705, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 523.4 attack_power points (41.01 DPS) | yes | Julie's Dagger (6660, -1.08 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.23 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.40 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (189.6 DPS) | yes | Guttbuster (13139, -0.54 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.60 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Substandard Belt Chain; legs: Knight's Chain Legplates; feet: Sergeant Major's Chain Sabatons; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Cestus; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 745, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60 (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 288.8. Weights run: 1.4s. Verify run: 12.5s. 1373 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.955 ± 0.026, strength=1.210 ± 0.002, crit=1.792 ± 0.051 per rating point (14 rating = 1%, 25.083 per %), hit=2.617 ± 0.222 per rating point (10 rating = 1%, 26.173 per %), melee_haste=13.942 ± 3.152

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (288.8 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -5.22 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 50.2 attack_power points (3.49 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -0.08 DPS) [quest]; Skibi's Pendant (13089, -1.30 DPS) [world_drop] |
| shoulder | Beastmaster's Mantle (226884) | Anthion's Parting Words [quest] | sim-verified (288.8 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Magister's Mantle (16689, -3.57 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 54.2 attack_power points (3.77 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -0.72 DPS) [rep]; Shifting Cloak (18511, -1.46 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 123.6 attack_power points (8.60 DPS) | yes | Field Marshal's Chain Armor (231563, -0.52 DPS) [vendor]; Obsidian Mail Tunic (22191, -1.57 DPS) [crafted]; Knight-Captain's Chain Armor (227089, -1.60 DPS) [vendor] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (288.8 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -5.91 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (288.8 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -3.79 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (288.8 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -3.81 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 144.8 attack_power points (10.07 DPS) | yes | Marshal's Chain Legplates (231558, -1.99 DPS) [vendor]; Sentinel's Leather Pants (237818, -2.91 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -3.07 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (288.8 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.00 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 73.1 attack_power points (5.08 DPS) | yes | Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Protector's Band (19514, -2.58 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 67.3 attack_power points (4.68 DPS) | yes | Cutthroat's Signet (272408, -0.95 DPS) [vendor]; Protector's Band (19514, -2.17 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (288.8 DPS) | yes | Frozen Heart of the Mountain (249469, -8.29 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (288.8 DPS) | yes | Talisman of Ascendance (22678, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (288.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -7.45 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 836.6 attack_power points (58.20 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -4.20 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -24.66 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (288.8 DPS) | yes | The Purifier (22656, -0.16 DPS) [quest]; Bloodseeker (19107, -0.28 DPS) [quest]; Dark Iron Rifle (16004, -1.99 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Beastmaster's Mantle; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1373, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 766.2. Weights run: 1.4s. Verify run: 11.0s. 1373 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.195 ± 0.030, strength=1.210 ± 0.003, crit=2.564 ± 0.061 per rating point (14 rating = 1%, 35.892 per %), hit=3.860 ± 0.439 per rating point (10 rating = 1%, 38.603 per %), melee_haste=not significant (20.989 ± 5.891)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 99.9 attack_power points (14.13 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Beastmaster's Cap (226887, -4.67 DPS, sim-verified) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 62.6 attack_power points (8.85 DPS) | yes | Medallion of the Dawn (22659, -0.38 DPS) [quest]; Amulet of the Darkmoon (19491, -1.24 DPS) [quest]; Skibi's Pendant (13089, -3.96 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 95.2 attack_power points (13.46 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.6 attack_power points (9.42 DPS) | yes | Cloak of the Honor Guard (20073, -3.06 DPS) [rep]; Stalwart Cloak (272415, -3.96 DPS) [vendor]; Cape of the Black Baron (13340, -6.92 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.7 attack_power points (22.02 DPS) | yes | Field Marshal's Chain Armor (231563, -3.50 DPS) [vendor]; Timbermaw Tunic (252484, -5.71 DPS) [crafted]; Knight-Captain's Chain Armor (227089, -10.44 DPS, sim-verified) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 55.0 attack_power points (7.78 DPS) | yes | Bracers of Subterfuge (22668, -1.41 DPS) [quest]; Marshal's Chain Bracers (16461, -1.57 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.65 DPS) [pvp] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 74.5 attack_power points (10.54 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, -0.27 DPS) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 88.5 attack_power points (12.51 DPS) | yes | Highlander's Leather Girdle (20045, -2.63 DPS) [rep]; Ferocity of the Timbermaw (227805, -2.79 DPS) [vendor]; Highlander's Chain Girdle (20043, -6.61 DPS, sim-verified) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 187.2 attack_power points (26.48 DPS) | yes | Sentinel's Leather Pants (237818, -7.94 DPS) [vendor]; Marshal's Chain Legplates (231558, -7.95 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -8.85 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 62.7 attack_power points (8.87 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.62 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 91.3 attack_power points (12.91 DPS) | yes | Cutthroat's Signet (272408, -3.10 DPS) [vendor]; Protector's Band (19514, -7.44 DPS) [rep]; Channeler's Ring (272406, -7.45 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 90.5 attack_power points (12.80 DPS) | yes | Cutthroat's Signet (272408, -4.36 DPS, sim-verified) [vendor]; Protector's Band (19514, -7.33 DPS) [rep]; Channeler's Ring (272406, -7.34 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (766.2 DPS) | yes | Frozen Heart of the Mountain (249469, -23.47 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (766.2 DPS) | yes | Talisman of Ascendance (22678, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (766.2 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -49.40 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.8 attack_power points (118.63 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -14.43 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -50.44 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (766.2 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.73 DPS) [world_drop]; Bloodseeker (19107, -1.53 DPS) [quest]; Dark Iron Rifle (16004, -7.21 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1373, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 73.9. Weights run: 1.5s. Verify run: 4.2s. 201 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.045 ± 0.022, strength=1.210 ± 0.003, crit=0.748 ± 0.023 per rating point (14 rating = 1%, 10.475 per %), hit=1.112 ± 0.058 per rating point (10 rating = 1%, 11.118 per %), melee_haste=8.258 ± 1.893

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.4 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.3 attack_power points (0.54 DPS) | yes | Erudite's Amulet (277204, -0.37 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.2 attack_power points (0.45 DPS) | yes | Slime-encrusted Pads (6461, -0.46 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.3 attack_power points (0.54 DPS) | yes | Cape of the Brotherhood (5193, -0.09 DPS) [dungeon]; Dark Leather Cloak (2316, -0.16 DPS) [crafted]; Bristlebark Cape (14571, -0.18 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (73.9 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.94 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.2 attack_power points (0.45 DPS) | yes | Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Wolf Bracers (4794, -0.09 DPS) [vendor]; Ratchet Wristwraps (274742, -0.18 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.1 attack_power points (0.75 DPS) | yes | Bristlebark Gloves (14572, -0.18 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.18 DPS) [crafted]; Serpent Gloves (5970, -0.21 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.22 DPS) [crafted]; Murloc Scale Belt (5780, -0.41 DPS) [crafted]; Deviate Scale Belt (6468, -2.60 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (73.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.08 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (73.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.99 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.0 attack_power points (0.57 DPS) | yes | Bounty Hunter's Ring (5351, -0.30 DPS) [quest]; Demon Band (12054, -0.36 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.3 attack_power points (0.54 DPS) | yes | Bounty Hunter's Ring (5351, -0.28 DPS, sim-verified) [quest]; Demon Band (12054, -0.32 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 attack_power points (10.88 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 236.7 attack_power points (10.34 DPS) | yes | Wingblade (6504, -6.39 DPS, sim-verified) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.2 attack_power points (0.36 DPS) | yes | Deadly Blunderbuss (4369, -0.18 DPS) [crafted]; Light Bow (4576, -0.18 DPS) [world_drop]; Privateer Musket (5309, -0.18 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 201, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 209611 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 0000000000000000-0000000000000000-500230131051000000)

Set DPS (verified): 99.8. Weights run: 1.4s. Verify run: 2.6s. 365 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.818 ± 0.019, strength=1.210 ± 0.002, crit=0.862 ± 0.027 per rating point (14 rating = 1%, 12.067 per %), hit=1.190 ± 0.071 per rating point (10 rating = 1%, 11.903 per %), melee_haste=not significant (5.793 ± 1.801)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 18.2 attack_power points (0.93 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.19 DPS) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 14.5 attack_power points (0.75 DPS) | yes | Ghostshard Talisman (7731, -0.03 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop]; Wolfpack Medallion (5754, -0.37 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.1 attack_power points (1.33 DPS) | yes | Mantle of Thieves (2264, -0.29 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.56 DPS) [crafted]; Bristlebark Amice (14573, -0.59 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 16.4 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Cloak of Night (4447, -0.28 DPS) [world]; Swiftrunner Cape (6745, -0.28 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 25.5 attack_power points (1.30 DPS) | yes | Brawler's Leather Tunic (252508, -0.19 DPS) [crafted]; Panther Armor (6670, -0.34 DPS) [quest]; Brawler's Leather Armor (252490, -0.34 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.8 attack_power points (0.81 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.25 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.4 attack_power points (0.99 DPS) | yes | Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.12 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.23 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.02 DPS) [crafted]; Prowler's Leather Belt (252459, -0.11 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.33 DPS) | yes | Troll's Bane Leggings (13114, -0.03 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.06 DPS) [crafted]; Petrolspill Leggings (9509, -1.09 DPS, sim-verified) [dungeon] |
| feet | Vorrel's Boots (7751) | Vorrel's Revenge [quest] | 18.2 attack_power points (0.93 DPS) | yes | Feet of the Lynx (1121, -0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.19 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.2 attack_power points (1.09 DPS) | yes | Thunderbrow Ring (13097, -0.31 DPS) [world_drop]; Monkey Ring (6748, -0.43 DPS) [quest]; Band of the Fist (17694, -0.47 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.2 attack_power points (0.93 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.28 DPS) [quest]; Band of the Fist (17694, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (17.49 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Morbid Dawn (7689, +0.00 DPS) [dungeon] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | sim-verified (99.8 DPS) | yes | Swinetusk Shank (6691, -17.42 DPS, sim-verified) [dungeon] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.61 DPS) | yes | Golemsight Long Gun (273029, -0.06 DPS) [dungeon]; Silver Star (3463, -0.15 DPS) [quest]; Double-barreled Shotgun (2098, -0.15 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Vorrel's Boots; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Pronged Reaver; off_hand: Serrated Raptor Claw; ranged: Glass Shooter

No-known-source sample (15 of 365, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 0000000000000000-0000000000000000-500230131051120151)

Set DPS (verified): 142.3. Weights run: 1.5s. Verify run: 3.1s. 560 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.834 ± 0.016, strength=1.210 ± 0.002, crit=0.886 ± 0.026 per rating point (14 rating = 1%, 12.408 per %), hit=1.534 ± 0.101 per rating point (10 rating = 1%, 15.343 per %), melee_haste=7.052 ± 1.532

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 33.5 attack_power points (2.38 DPS) | yes | Barbaric Iron Helm (7915, -0.43 DPS) [crafted]; Hawkeye's Helm (14591, -0.60 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.69 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 20.2 attack_power points (1.43 DPS) | yes | Ghostshard Talisman (7731, -0.44 DPS) [dungeon]; Ethereal Talisman (4430, -0.48 DPS) [quest]; Kaleidoscope Chain (13084, -0.57 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 32.2 attack_power points (2.29 DPS) | yes | Forest Tracker Epaulets (2278, -0.42 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.55 DPS) [dungeon]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 23.2 attack_power points (1.65 DPS) | yes | Hawkeye's Cloak (14593, -0.45 DPS, sim-verified) [world_drop]; Imperial Cloak (6432, -0.61 DPS) [dungeon]; Parachute Cloak (10518, -0.61 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 31.2 attack_power points (2.22 DPS) | yes | Kolkar Marauder Chain (6773, -0.10 DPS) [quest]; Nightscape Tunic (8175, -0.26 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.26 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.42 DPS) | yes | Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.49 DPS, sim-verified) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 32.9 attack_power points (2.34 DPS) | yes | Gloves of Holy Might (867, -0.03 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.06 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.26 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.13 DPS) | yes | Ogron's Sash (13117, -0.18 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.43 DPS) [rep]; Skulker's Leather Belt (252520, -0.44 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 38.5 attack_power points (2.74 DPS) | yes | Triprunner Dungarees (9624, -0.13 DPS) [quest]; Veteran's Silvered Chain Leggings (250523, -0.88 DPS) [crafted]; Ferine Leggings (6690, -0.89 DPS) [dungeon] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 29.2 attack_power points (2.08 DPS) | yes | Skulker's Leather Shoes (252531, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.22 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 24.4 attack_power points (1.73 DPS) | yes | Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.30 DPS) [dungeon]; Mark of Kern (2262, -0.31 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 21.3 attack_power points (1.52 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Falcon's Hook (7552, -0.09 DPS) [dungeon]; Mark of Kern (2262, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (142.3 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.11 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 447.4 attack_power points (31.82 DPS) | yes | Vanquisher's Sword (10823, -0.46 DPS, sim-verified) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (142.3 DPS) | yes | Swiftwind (13038, -0.08 DPS) [world_drop]; Monolithic Bow (9426, -0.09 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Scarlet Gauntlets; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Legionnaire's Band; finger2: Ironspine's Eye; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 560, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 0000000000000000-3250000000000000-500230131051120151)

Set DPS (verified): 189.2. Weights run: 1.5s. Verify run: 3.7s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.754 ± 0.016, strength=1.210 ± 0.002, crit=0.958 ± 0.027 per rating point (14 rating = 1%, 13.414 per %), hit=1.272 ± 0.115 per rating point (10 rating = 1%, 12.722 per %), melee_haste=9.759 ± 1.824

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 32.6 attack_power points (2.55 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.22 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.9 attack_power points (2.26 DPS) | yes | Scout's Medallion (19535, -0.61 DPS) [rep]; Woven Ivy Necklace (19159, -0.83 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 31.3 attack_power points (2.45 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.09 DPS) [crafted]; Failed Flying Experiment (9647, -0.13 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.3 attack_power points (2.06 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Pridelord Cape (14673, -0.54 DPS) [world_drop]; Dark Hooded Cape (5257, -0.65 DPS, sim-verified) [world] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 44.9 attack_power points (3.52 DPS) | yes | Blazewind Breastplate (11193, -0.07 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.40 DPS) [vendor]; Charred Leather Tunic (19127, -1.04 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 31.2 attack_power points (2.44 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Arena Bands (18711, -1.11 DPS, sim-verified) [world] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 50.7 attack_power points (3.97 DPS) | yes | Gloves of Holy Might (867, -1.35 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -1.41 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -2.11 DPS, sim-verified) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 39.3 attack_power points (3.08 DPS) | yes | Skulker's Leather Waistguard (252474, -0.02 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.46 DPS) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 43.8 attack_power points (3.43 DPS) | yes | Stone Guard's Chain Legplates (220833, -0.46 DPS) [vendor]; Basilisk Hide Pants (1718, -0.55 DPS) [world_drop]; Triprunner Dungarees (9624, -0.67 DPS) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 37.1 attack_power points (2.91 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Albino Crocscale Boots (17728, -0.16 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.16 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 27.9 attack_power points (2.18 DPS) | yes | White Bone Band (11862, -0.30 DPS) [quest]; Ironspine's Eye (7686, -0.57 DPS) [dungeon]; Blackstone Ring (17713, -0.62 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.6 attack_power points (1.92 DPS) | yes | Ironspine's Eye (7686, -0.31 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon]; White Bone Band (11862, -0.80 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (189.2 DPS) | yes | Molten Heart of the Mountain (249470, -5.12 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (189.2 DPS) | yes | Molten Heart of the Mountain (249470, -2.19 DPS, sim-verified) [crafted] |
| main_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | sim-verified (189.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thrash Blade (17705, +0.00 DPS, sim-verified) [quest]; Darkspear Raider's Reaper (272080, +0.00 DPS) [vendor] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 523.4 attack_power points (41.01 DPS) | yes | Julie's Dagger (6660, +0.00 DPS) [world_drop]; Claw of Celebras (17738, -3.23 DPS) [dungeon]; White Bone Shredder (11863, -5.59 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (189.2 DPS) | yes | Guttbuster (13139, -0.54 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.82 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; legs: Serpentskin Leggings; feet: Sandstalker Ankleguards; finger1: Legionnaire's Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Cestus; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 294.1. Weights run: 1.4s. Verify run: 12.5s. 1355 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=1.955 ± 0.026, strength=1.210 ± 0.002, crit=1.792 ± 0.051 per rating point (14 rating = 1%, 25.083 per %), hit=2.617 ± 0.222 per rating point (10 rating = 1%, 26.173 per %), melee_haste=13.942 ± 3.152

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (294.1 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -6.84 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 50.2 attack_power points (3.49 DPS) | yes | Amulet of the Darkmoon (19491, -0.06 DPS) [quest]; Medallion of the Dawn (22659, -0.08 DPS) [quest]; Skibi's Pendant (13089, -1.30 DPS) [world_drop] |
| shoulder | Beastmaster's Mantle (226884) | Anthion's Parting Words [quest] | sim-verified (294.1 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Magister's Mantle (16689, -3.12 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 54.2 attack_power points (3.77 DPS) | yes | Cape of the Black Baron (13340, -0.34 DPS) [dungeon]; Deathguard's Cloak (20068, -0.72 DPS) [rep]; Shifting Cloak (18511, -1.46 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 123.6 attack_power points (8.60 DPS) | yes | Warlord's Chain Armor (231566, -0.52 DPS) [vendor]; Obsidian Mail Tunic (22191, -1.31 DPS, sim-verified) [crafted]; Legionnaire's Chain Armor (227083, -1.60 DPS) [vendor] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (294.1 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -6.57 DPS, sim-verified) [dungeon] |
| hands | Beastmaster's Gauntlets (226883) | Mokvar [vendor] | sim-verified (294.1 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Beaststalker's Gloves (16676, -3.57 DPS, sim-verified) [dungeon] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (294.1 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -3.78 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 144.8 attack_power points (10.07 DPS) | yes | Outrider's Chain Leggings (22673, -1.78 DPS, sim-verified) [rep]; General's Chain Legplates (231567, -1.99 DPS) [vendor]; Sentinel's Leather Pants (237818, -2.91 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (294.1 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -5.60 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 73.1 attack_power points (5.08 DPS) | yes | Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Legionnaire's Band (19510, -2.58 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 67.3 attack_power points (4.68 DPS) | yes | Cutthroat's Signet (272408, -0.95 DPS) [vendor]; Legionnaire's Band (19510, -2.17 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (294.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (294.1 DPS) | yes | Frozen Heart of the Mountain (249469, -2.56 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.71 DPS, sim-verified) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (294.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -7.48 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 836.6 attack_power points (58.20 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -5.02 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -24.66 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (294.1 DPS) | yes | The Purifier (22656, -0.16 DPS) [quest]; Bloodseeker (19107, -0.28 DPS) [quest]; Dark Iron Rifle (16004, -2.50 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Beads of Ogre Might; shoulder: Beastmaster's Mantle; back: Howler's Furs; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beastmaster's Gauntlets; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1355, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 0000000000000000-3250050000500000-500230131051120151)

Set DPS (verified): 776.8. Weights run: 1.4s. Verify run: 10.7s. 1355 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.195 ± 0.030, strength=1.210 ± 0.003, crit=2.564 ± 0.061 per rating point (14 rating = 1%, 35.892 per %), hit=3.860 ± 0.439 per rating point (10 rating = 1%, 38.603 per %), melee_haste=not significant (20.989 ± 5.891)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 99.9 attack_power points (14.13 DPS) | yes | Beastmaster's Cap (226887, +0.00 DPS) [quest]; Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 62.6 attack_power points (8.85 DPS) | yes | Medallion of the Dawn (22659, -0.38 DPS) [quest]; Amulet of the Darkmoon (19491, -1.24 DPS) [quest]; Skibi's Pendant (13089, -3.96 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 95.2 attack_power points (13.46 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.6 attack_power points (9.42 DPS) | yes | Cape of the Black Baron (13340, -1.93 DPS) [dungeon]; Deathguard's Cloak (20068, -3.06 DPS) [rep]; Stalwart Cloak (272415, -3.96 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.7 attack_power points (22.02 DPS) | yes | Warlord's Chain Armor (231566, -3.50 DPS) [vendor]; Legionnaire's Chain Armor (227083, -4.40 DPS) [vendor]; Timbermaw Tunic (252484, -5.16 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 55.0 attack_power points (7.78 DPS) | yes | Bracers of Subterfuge (22668, -1.41 DPS) [quest]; General's Chain Wristguards (16570, -1.57 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.65 DPS) [pvp] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 74.5 attack_power points (10.54 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 88.5 attack_power points (12.51 DPS) | yes | Defiler's Chain Girdle (20150, -2.63 DPS) [rep]; Defiler's Leather Girdle (20190, -2.63 DPS) [rep]; Ferocity of the Timbermaw (227805, -2.79 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 187.2 attack_power points (26.48 DPS) | yes | Sentinel's Leather Pants (237818, -7.94 DPS) [vendor]; General's Chain Legplates (231567, -7.95 DPS) [vendor]; Outrider's Chain Leggings (22673, -8.52 DPS, sim-verified) [rep] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 62.7 attack_power points (8.87 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.72 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 91.3 attack_power points (12.91 DPS) | yes | Cutthroat's Signet (272408, -3.10 DPS) [vendor]; Legionnaire's Band (19510, -7.44 DPS) [rep]; Channeler's Ring (272406, -7.45 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 90.5 attack_power points (12.80 DPS) | yes | Cutthroat's Signet (272408, -6.40 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -7.33 DPS) [rep]; Channeler's Ring (272406, -7.34 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (776.8 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (776.8 DPS) | yes | Frozen Heart of the Mountain (249469, -4.85 DPS) [crafted]; Molten Heart of the Mountain (249470, -11.63 DPS, sim-verified) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (776.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -45.91 DPS, sim-verified) [crafted] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 838.8 attack_power points (118.63 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -12.33 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -50.44 DPS) [dungeon] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (776.8 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.73 DPS) [world_drop]; Bloodseeker (19107, -1.53 DPS) [quest]; Dark Iron Rifle (16004, -8.52 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Belt of Preserved Heads; legs: Sentinel's Chain Leggings; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1355, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

