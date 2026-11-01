# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 67.3. Weights run: 2.4s. Verify run: 3.4s. 233 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.944 ± 0.019, strength=1.215 ± 0.002, crit=1.398 ± 0.021 per rating point (14 rating = 1%, 19.575 per %), hit=1.115 ± 0.041 per rating point (10 rating = 1%, 11.155 per %), melee_haste=7.304 ± 1.179

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 23.6 attack_power points (1.06 DPS) | yes | Defender's Leather Hood (252447, -0.63 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 5.9 attack_power points (0.26 DPS) | yes | Tarnished Locket (279870, -0.28 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 14.7 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.69 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 17.7 attack_power points (0.79 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 32.4 attack_power points (1.45 DPS) | yes | Brawler's Leather Armor (252490, -0.27 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.53 DPS) [crafted]; Prospector's Chestpiece (14562, -0.55 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 14.7 attack_power points (0.66 DPS) | yes | Bravo's Armbands (270015, -0.02 DPS) [quest]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.16 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.3 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -1.74 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (67.3 DPS) | yes | Brawler's Leather Belt (252428, -0.06 DPS) [crafted]; Dusty Belt (279897, -0.15 DPS) [quest]; Deviate Scale Belt (6468, -2.53 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.3 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.01 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (67.3 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.97 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 attack_power points (0.26 DPS) | yes | The 1 Ring (8350, -0.08 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.9 attack_power points (0.22 DPS) | yes | The 1 Ring (8350, -0.03 DPS) [world]; Ring of the Moon (12052, -0.16 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (11.17 DPS) | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.5 attack_power points (10.70 DPS) | yes | Assassin's Blade (1935, -0.23 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.8 attack_power points (0.53 DPS) | yes | Light Bow (4576, -0.26 DPS) [world_drop]; Owlsight Rifle (15205, -0.26 DPS) [quest]; Deadly Blunderbuss (4369, -0.28 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Lavishly Jeweled Ring; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 233, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230231050000000)

Set DPS (verified): 97.4. Weights run: 2.5s. Verify run: 3.5s. 440 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.463 ± 0.015, strength=1.213 ± 0.002, crit=1.486 ± 0.021 per rating point (14 rating = 1%, 20.802 per %), hit=1.269 ± 0.051 per rating point (10 rating = 1%, 12.694 per %), melee_haste=5.483 ± 1.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.6 attack_power points (1.32 DPS) | yes | Brawler's Leather Hood (252504, -0.30 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.40 DPS) [world]; Defender's Leather Helm (252455, -0.54 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.7 attack_power points (0.79 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.31 DPS, sim-verified) [world]; Erudite's Amulet (277204, -0.52 DPS) [quest]; River Pride Choker (13087, -0.53 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.2 attack_power points (1.78 DPS) | yes | Mantle of Thieves (2264, -0.69 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.79 DPS) [world_drop]; Barbaric Shoulders (5964, -0.79 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.9 attack_power points (1.12 DPS) | yes | Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Fenrus' Hide (6340, -0.33 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.33 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 34.5 attack_power points (1.85 DPS) | yes | Brawler's Leather Tunic (252508, -0.40 DPS) [crafted]; Tunic of Westfall (2041, -0.45 DPS, sim-verified) [quest]; Brawler's Leather Armor (252490, -0.60 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.6 attack_power points (1.05 DPS) | yes | Insignia Bracers (6410, -0.26 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS) [crafted]; Demonhide Bracers (270033, -0.33 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (97.4 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Cartilage Shapers (273054, +0.00 DPS) [world]; Insignia Gloves (6408, -1.45 DPS, sim-verified) [world_drop] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 29.4 attack_power points (1.58 DPS) | yes | Prowler's Leather Belt (252459, -0.20 DPS) [crafted]; Stalker's Leather Belt (252521, -0.39 DPS) [crafted]; Warden's Leather Belt (252460, -0.40 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (97.4 DPS) | yes | Dusky Leather Leggings (7373, +0.00 DPS) [crafted]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Petrolspill Leggings (9509, -1.31 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.3 attack_power points (1.25 DPS) | yes | Insignia Boots (4055, -0.20 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.27 DPS) [crafted]; Footpads of the Fang (10411, -0.46 DPS) [dungeon] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.92 DPS) | yes | Defias Renegade Ring (1076, -0.26 DPS) [dungeon]; Monkey Ring (6748, -0.26 DPS) [quest]; Seal of Wrynn (2933, -0.33 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 14.8 attack_power points (0.79 DPS) | yes | Defias Renegade Ring (1076, -0.13 DPS) [dungeon]; Monkey Ring (6748, -0.13 DPS) [quest]; Seal of Wrynn (2933, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (18.33 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Swinetusk Shank (6691, -3.98 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (97.4 DPS) | yes | Shoni's Disarming Tool (9608, -5.03 DPS) [quest]; Swinetusk Shank (6691, -18.03 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.8 attack_power points (0.79 DPS) | yes | Double-barreled Shotgun (2098, -0.20 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop]; Lil Timmy's Peashooter (13136, -0.26 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Gloves of the Fang; waist: Skulker's Leather Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 440, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 151.5. Weights run: 2.5s. Verify run: 2.1s. 696 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.187 ± 0.013, strength=1.210 ± 0.002, crit=1.511 ± 0.021 per rating point (14 rating = 1%, 21.158 per %), hit=1.216 ± 0.082 per rating point (10 rating = 1%, 12.162 per %), melee_haste=not significant (4.113 ± 1.102)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 37.4 attack_power points (2.78 DPS) | yes | Barbaric Iron Helm (7915, -0.51 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.63 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.66 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.6 attack_power points (1.01 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; River Pride Choker (13087, -0.65 DPS) [world_drop]; Gazlowe's Charm (13088, -0.65 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.1 attack_power points (2.24 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.45 DPS) [crafted]; Mantle of Thieves (2264, -0.61 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 20.4 attack_power points (1.52 DPS) | yes | Imperial Cloak (6432, -0.21 DPS) [dungeon]; Parachute Cloak (10518, -0.21 DPS) [crafted]; Hawkeye's Cloak (14593, -1.31 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 52.5 attack_power points (3.90 DPS) | yes | Wolffear Harness (13110, -1.13 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.46 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.46 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.49 DPS) | yes | Hawkeye's Bracers (14590, -0.15 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.19 DPS) [dungeon]; Dusky Bracers (7378, -0.19 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.2 attack_power points (3.06 DPS) | yes | Dragonscale Gauntlets (8347, -0.51 DPS) [crafted]; Scarlet Gauntlets (10331, -0.60 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.62 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 30.6 attack_power points (2.27 DPS) | yes | Skulker's Leather Belt (252520, -0.27 DPS) [crafted]; Scorpashi Sash (14652, -0.29 DPS) [world_drop]; Officer's Belt (250556, -0.40 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.9 attack_power points (3.42 DPS) | yes | Triprunner Dungarees (9624, -0.22 DPS) [quest]; Petrolspill Leggings (9509, -1.14 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.14 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 32.8 attack_power points (2.44 DPS) | yes | Skulker's Leather Shoes (252531, -0.02 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.31 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 24.3 attack_power points (1.81 DPS) | yes | Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop]; Disengagement Ring (276202, -0.51 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.3 attack_power points (1.73 DPS) | yes | Mark of Kern (2262, -0.25 DPS) [dungeon]; Assault Band (13095, -0.25 DPS) [world_drop]; Disengagement Ring (276202, -0.43 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (151.5 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -7.00 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (32.87 DPS) | yes | Shoni's Disarming Tool (9608, -16.16 DPS) [quest]; Dazzling Longsword (869, -18.37 DPS, sim-verified) [world_drop] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (151.5 DPS) | yes | The Silencer (13138, -0.10 DPS) [world_drop]; Monolithic Bow (9426, -0.11 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.21 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 0000000000000000-3210000000000000-500230231051222151)

Set DPS (verified): 205.5. Weights run: 2.5s. Verify run: 2.5s. 885 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.218 ± 0.013, strength=1.210 ± 0.002, crit=1.677 ± 0.023 per rating point (14 rating = 1%, 23.474 per %), hit=1.593 ± 0.091 per rating point (10 rating = 1%, 15.934 per %), melee_haste=5.869 ± 1.248

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 37.7 attack_power points (2.94 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.00 DPS) [crafted]; Knight-Lieutenant's Mail Helmet (223075, -0.07 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.9 attack_power points (2.72 DPS) | yes | Kaleidoscope Chain (13084, -0.74 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -1.79 DPS) [world]; River Pride Choker (13087, -2.34 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 35.3 attack_power points (2.75 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, -0.07 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 33.3 attack_power points (2.60 DPS) | yes | Blackflame Cape (13109, -0.52 DPS) [world_drop]; Serpentskin Cloak (8259, -0.69 DPS) [world_drop]; Pridelord Cape (14673, -0.76 DPS) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | 56.7 attack_power points (4.43 DPS) | yes | Warbear Harness (15064, -0.27 DPS) [crafted]; Icemail Jerkin (1981, -0.27 DPS) [world_drop]; Blazewind Breastplate (11193, -3.42 DPS, sim-verified) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 38.1 attack_power points (2.97 DPS) | yes | Pridelord Bands (14672, -0.77 DPS) [world_drop]; Arena Bands (18711, -0.79 DPS) [world]; Skulker's Leather Bracers (252540, -0.84 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 43.5 attack_power points (3.39 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.12 DPS) [crafted]; Serpentskin Gloves (8260, -0.45 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.51 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 45.6 attack_power points (3.56 DPS) | yes | Prowler's Leather Waistguard (252473, -0.16 DPS) [crafted]; Sagebrush Girdle (17778, -0.96 DPS) [quest]; Pridelord Girdle (14674, -1.10 DPS) [world_drop] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 54.5 attack_power points (4.26 DPS) | yes | Serpentskin Leggings (8262, -0.26 DPS) [world_drop]; Stormshroud Pants (15057, -0.59 DPS) [crafted]; Gryphon Rider's Leggings (9652, -3.56 DPS, sim-verified) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 44.4 attack_power points (3.46 DPS) | yes | Skulker's Leather Boots (252469, -0.17 DPS) [crafted]; Fleetfoot Greaves (11627, -0.17 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.33 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.1 attack_power points (2.42 DPS) | yes | Falcon's Hook (7552, -0.58 DPS) [dungeon]; Mark of Kern (2262, -0.86 DPS) [dungeon]; Blackstone Ring (17713, -0.86 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.6 attack_power points (1.92 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (205.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (205.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (205.5 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -3.65 DPS, sim-verified) [crafted] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (41.09 DPS) | yes | Julie's Dagger (6660, +0.00 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.46 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.55 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (205.5 DPS) | yes | Guttbuster (13139, -0.76 DPS) [world_drop]; Gryphonwing Long Bow (13022, -1.04 DPS) [dungeon]; Dark Iron Rifle (16004, -2.80 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thrash Blade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 885, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 297.4. Weights run: 1.8s. Verify run: 5.8s. 1755 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.105 ± 0.019, strength=1.210 ± 0.002, crit=2.221 ± 0.037 per rating point (14 rating = 1%, 31.093 per %), hit=2.713 ± 0.156 per rating point (10 rating = 1%, 27.127 per %), melee_haste=24.977 ± 2.087

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (297.4 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -14.48 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.4 attack_power points (2.87 DPS) | yes | Sunprism Pendant (20649, -0.70 DPS) [quest]; Kaleidoscope Chain (13084, -1.73 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.84 DPS) [world] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (297.4 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -2.63 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 51.6 attack_power points (4.43 DPS) | yes | Fel Cape (279269, -1.76 DPS) [crafted]; Sergeant Major's Cape (16337, -2.15 DPS) [pvp]; Dark Phantom Cape (13122, -2.16 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 136.1 attack_power points (11.69 DPS) | yes | Field Marshal's Chain Armor (231563, -0.98 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -1.89 DPS) [vendor]; Timbermaw Tunic (252484, -3.67 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (297.4 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 59.1 attack_power points (5.08 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -6.39 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 54.9 attack_power points (4.71 DPS) | yes | Marshal's Chain Girdle (16464, -0.94 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.09 DPS) [crafted]; Skulker's Leather Waistguard (252474, -1.23 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 77.1 attack_power points (6.62 DPS) | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Warbear Woolies (15065, -6.54 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.8 attack_power points (4.71 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 29.5 attack_power points (2.53 DPS) | yes | Falcon's Hook (7552, -0.59 DPS) [dungeon]; Mark of Kern (2262, -0.81 DPS) [dungeon]; Blackstone Ring (17713, -0.81 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.5 attack_power points (2.02 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.30 DPS) [dungeon]; Blackstone Ring (17713, -0.30 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (297.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (297.4 DPS) | yes | Mark of the Chosen (17774, -3.63 DPS, sim-verified) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (297.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -1.70 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (297.4 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -11.85 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -52.95 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (297.4 DPS) | yes | Eaglehorn Long Bow (13023, -0.72 DPS) [world_drop]; Guttbuster (13139, -0.77 DPS) [world_drop]; Dark Iron Rifle (16004, -3.22 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Magister's Mantle; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Magister's Bindings; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 690.7. Weights run: 1.5s. Verify run: 3.9s. 1755 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.238 ± 0.020, strength=1.210 ± 0.002, crit=2.552 ± 0.040 per rating point (14 rating = 1%, 35.734 per %), hit=3.280 ± 0.272 per rating point (10 rating = 1%, 32.801 per %), melee_haste=20.128 ± 3.218

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+9.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.77 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.1 attack_power points (6.02 DPS) | yes | Sunprism Pendant (20649, -1.42 DPS) [quest]; Kaleidoscope Chain (13084, -3.65 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.96 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 55.8 attack_power points (9.55 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.6 attack_power points (9.17 DPS) | yes | Dark Phantom Cape (13122, -3.42 DPS) [world_drop]; Fel Cape (279269, -3.81 DPS, sim-verified) [crafted]; Sergeant Major's Cape (16337, -4.45 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.3 attack_power points (25.91 DPS) | yes | Field Marshal's Chain Armor (231563, -3.38 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -4.55 DPS) [vendor]; Timbermaw Tunic (252484, -7.65 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 38.4 attack_power points (6.57 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.46 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 63.7 attack_power points (10.91 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -13.58 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 56.7 attack_power points (9.71 DPS) | yes | Marshal's Chain Girdle (16464, -1.78 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.86 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.22 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 81.7 attack_power points (13.99 DPS) | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Warbear Woolies (15065, -13.90 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 56.8 attack_power points (9.73 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.3 attack_power points (5.36 DPS) | yes | Falcon's Hook (7552, -1.29 DPS) [dungeon]; Mark of Kern (2262, -1.94 DPS) [dungeon]; Blackstone Ring (17713, -1.94 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.8 attack_power points (4.25 DPS) | yes | Falcon's Hook (7552, -0.18 DPS) [dungeon]; Mark of Kern (2262, -0.82 DPS) [dungeon]; Blackstone Ring (17713, -0.82 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+78.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -24.50 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -78.28 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.53 DPS) [world_drop]; Guttbuster (13139, -1.68 DPS) [world_drop]; Dark Iron Rifle (16004, -5.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 68.1. Weights run: 2.4s. Verify run: 3.3s. 229 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.944 ± 0.019, strength=1.215 ± 0.002, crit=1.398 ± 0.021 per rating point (14 rating = 1%, 19.575 per %), hit=1.115 ± 0.041 per rating point (10 rating = 1%, 11.155 per %), melee_haste=7.304 ± 1.179

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 23.6 attack_power points (1.06 DPS) | yes | Defender's Leather Hood (252447, -0.67 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 5.9 attack_power points (0.26 DPS) | yes | Roadwatcher's Confidence (281265, -0.27 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 14.7 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.70 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 17.7 attack_power points (0.79 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (68.1 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.84 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 14.7 attack_power points (0.66 DPS) | yes | Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.16 DPS) [world_drop]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 22.5 attack_power points (1.01 DPS) | yes | Fletcher's Gloves (7348, -0.13 DPS) [crafted]; Serpent Gloves (5970, -0.22 DPS) [dungeon]; Bristlebark Gloves (14572, -0.26 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (68.1 DPS) | yes | Brawler's Leather Belt (252428, -0.06 DPS) [crafted]; Dark Leather Belt (4249, -0.28 DPS) [crafted]; Deviate Scale Belt (6468, -2.55 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (68.1 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.02 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (68.1 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.98 DPS, sim-verified) [world_drop] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 8.8 attack_power points (0.40 DPS) | yes | Demon Band (12054, -0.18 DPS) [world_drop]; The 1 Ring (8350, -0.21 DPS) [world]; Loop of Sacrifice (281673, -0.23 DPS) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.9 attack_power points (0.26 DPS) | yes | The 1 Ring (8350, -0.08 DPS) [world]; Loop of Sacrifice (281673, -0.10 DPS) [quest]; Demon Band (12054, -0.31 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (11.17 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.5 attack_power points (10.70 DPS) | yes | Wingblade (6504, -5.58 DPS, sim-verified) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.8 attack_power points (0.53 DPS) | yes | Light Bow (4576, -0.26 DPS) [world_drop]; Privateer Musket (5309, -0.26 DPS) [quest]; Deadly Blunderbuss (4369, -0.27 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Bounty Hunter's Ring; finger2: Lavishly Jeweled Ring; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 229, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 0000000000000000-0000000000000000-500230231050000000)

Set DPS (verified): 97.0. Weights run: 2.5s. Verify run: 3.4s. 437 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.463 ± 0.015, strength=1.213 ± 0.002, crit=1.486 ± 0.021 per rating point (14 rating = 1%, 20.802 per %), hit=1.269 ± 0.051 per rating point (10 rating = 1%, 12.694 per %), melee_haste=5.483 ± 1.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.6 attack_power points (1.32 DPS) | yes | Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.40 DPS) [world]; Defender's Leather Helm (252455, -0.54 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.7 attack_power points (0.79 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.15 DPS) [world]; Erudite's Amulet (277204, -0.52 DPS) [quest]; River Pride Choker (13087, -0.53 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.2 attack_power points (1.78 DPS) | yes | Mantle of Thieves (2264, -0.46 DPS) [dungeon]; Bristlebark Amice (14573, -0.79 DPS) [world_drop]; Barbaric Shoulders (5964, -0.79 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.9 attack_power points (1.12 DPS) | yes | Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Fenrus' Hide (6340, -0.33 DPS) [dungeon]; Swiftrunner Cape (6745, -0.33 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 34.5 attack_power points (1.85 DPS) | yes | Brawler's Leather Tunic (252508, -0.40 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.53 DPS) [quest]; Brawler's Leather Armor (252490, -0.60 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.6 attack_power points (1.05 DPS) | yes | Insignia Bracers (6410, -0.26 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS) [crafted]; Loamflake Bracers (15462, -0.39 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (97.0 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Cartilage Shapers (273054, +0.00 DPS) [world]; Insignia Gloves (6408, -1.33 DPS, sim-verified) [world_drop] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 29.4 attack_power points (1.58 DPS) | yes | Prowler's Leather Belt (252459, -0.20 DPS) [crafted]; Stalker's Leather Belt (252521, -0.39 DPS) [crafted]; Warden's Leather Belt (252460, -0.40 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (97.0 DPS) | yes | Dusky Leather Leggings (7373, +0.00 DPS) [crafted]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Petrolspill Leggings (9509, -1.18 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.3 attack_power points (1.25 DPS) | yes | Insignia Boots (4055, -0.20 DPS) [world_drop]; Warsong Boots (16977, -0.20 DPS) [quest]; Brawler's Leather Boots (252439, -0.27 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.92 DPS) | yes | Band of the Fist (17694, -0.13 DPS) [quest]; Defias Renegade Ring (1076, -0.26 DPS) [dungeon]; Monkey Ring (6748, -0.26 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 14.8 attack_power points (0.79 DPS) | yes | Band of the Fist (17694, +0.00 DPS, sim-verified) [quest]; Defias Renegade Ring (1076, -0.13 DPS) [dungeon]; Monkey Ring (6748, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (18.33 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Swinetusk Shank (6691, -3.28 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (97.0 DPS) | yes | Swinetusk Shank (6691, -17.30 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.8 attack_power points (0.79 DPS) | yes | Double-barreled Shotgun (2098, -0.20 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop]; Lil Timmy's Peashooter (13136, -0.26 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; waist: Skulker's Leather Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 437, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 148.7. Weights run: 2.5s. Verify run: 2.2s. 684 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.187 ± 0.013, strength=1.210 ± 0.002, crit=1.511 ± 0.021 per rating point (14 rating = 1%, 21.158 per %), hit=1.216 ± 0.082 per rating point (10 rating = 1%, 12.162 per %), melee_haste=not significant (4.113 ± 1.102)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 37.4 attack_power points (2.78 DPS) | yes | Barbaric Iron Helm (7915, -0.52 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.63 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.66 DPS) [world_drop] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 14.8 attack_power points (1.10 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.21 DPS) [world]; River Pride Choker (13087, -0.74 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.1 attack_power points (2.24 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.45 DPS) [crafted]; Mantle of Thieves (2264, -0.61 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 18.9 attack_power points (1.41 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Parachute Cloak (10518, -0.11 DPS) [crafted]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 52.5 attack_power points (3.90 DPS) | yes | Wolffear Harness (13110, -0.97 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.46 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.46 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.49 DPS) | yes | Hawkeye's Bracers (14590, -0.15 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.19 DPS) [dungeon]; Dusky Bracers (7378, -0.19 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.2 attack_power points (3.06 DPS) | yes | Dragonscale Gauntlets (8347, -0.51 DPS) [crafted]; Scarlet Gauntlets (10331, -0.61 DPS, sim-verified) [dungeon]; Skulker's Leather Gloves (252525, -0.62 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 30.6 attack_power points (2.27 DPS) | yes | Skulker's Leather Belt (252520, -0.27 DPS) [crafted]; Scorpashi Sash (14652, -0.29 DPS) [world_drop]; Officer's Belt (250556, -0.40 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.9 attack_power points (3.42 DPS) | yes | Triprunner Dungarees (9624, -0.22 DPS) [quest]; Petrolspill Leggings (9509, -1.14 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.14 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 32.8 attack_power points (2.44 DPS) | yes | Skulker's Leather Shoes (252531, -0.02 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.31 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 24.3 attack_power points (1.81 DPS) | yes | Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop]; Disengagement Ring (276202, -0.51 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.3 attack_power points (1.73 DPS) | yes | Mark of Kern (2262, -0.25 DPS) [dungeon]; Assault Band (13095, -0.25 DPS) [world_drop]; Disengagement Ring (276202, -0.43 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (148.7 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -5.90 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (32.87 DPS) | yes | Dazzling Longsword (869, -20.50 DPS, sim-verified) [world_drop] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (148.7 DPS) | yes | The Silencer (13138, -0.10 DPS) [world_drop]; Monolithic Bow (9426, -0.11 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.19 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 0000000000000000-3210000000000000-500230231051222151)

Set DPS (verified): 211.4. Weights run: 2.5s. Verify run: 2.6s. 874 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.218 ± 0.013, strength=1.210 ± 0.002, crit=1.677 ± 0.023 per rating point (14 rating = 1%, 23.474 per %), hit=1.593 ± 0.091 per rating point (10 rating = 1%, 15.934 per %), melee_haste=5.869 ± 1.248

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 37.7 attack_power points (2.94 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.00 DPS) [crafted]; Blood Guard's Mail Helmet (220820, -0.07 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.9 attack_power points (2.72 DPS) | yes | Woven Ivy Necklace (19159, -0.88 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -1.56 DPS) [quest]; Kaleidoscope Chain (13084, -1.65 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 35.3 attack_power points (2.75 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Blood Guard's Mail Epaulets (220823, -0.07 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 33.3 attack_power points (2.60 DPS) | yes | Serpentskin Cloak (8259, -0.69 DPS) [world_drop]; Pridelord Cape (14673, -0.76 DPS) [world_drop]; Blackflame Cape (13109, -0.81 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) | PvP rank 10 · Stone Guard · Horde [vendor] | 56.7 attack_power points (4.43 DPS) | yes | Warbear Harness (15064, -0.27 DPS) [crafted]; Icemail Jerkin (1981, -0.27 DPS) [world_drop]; Blazewind Breastplate (11193, -3.06 DPS, sim-verified) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 38.1 attack_power points (2.97 DPS) | yes | Skulker's Leather Bracers (252540, -0.76 DPS) [crafted]; Pridelord Bands (14672, -0.77 DPS) [world_drop]; Arena Bands (18711, -0.79 DPS) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 43.5 attack_power points (3.39 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.12 DPS) [crafted]; Serpentskin Gloves (8260, -0.45 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.51 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 45.6 attack_power points (3.56 DPS) | yes | Prowler's Leather Waistguard (252473, -0.16 DPS) [crafted]; Sagebrush Girdle (17778, -0.96 DPS) [quest]; Pridelord Girdle (14674, -1.10 DPS) [world_drop] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 54.5 attack_power points (4.26 DPS) | yes | Stormshroud Pants (15057, -0.59 DPS) [crafted]; Basilisk Hide Pants (1718, -0.62 DPS) [world_drop]; Serpentskin Leggings (8262, -3.22 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 44.4 attack_power points (3.46 DPS) | yes | Skulker's Leather Boots (252469, -0.17 DPS) [crafted]; Fleetfoot Greaves (11627, -0.17 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.33 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.1 attack_power points (2.42 DPS) | yes | White Bone Band (11862, -0.55 DPS) [quest]; Falcon's Hook (7552, -0.58 DPS) [dungeon]; Blackstone Ring (17713, -0.86 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.6 attack_power points (1.92 DPS) | yes | White Bone Band (11862, -0.05 DPS) [quest]; Falcon's Hook (7552, -0.08 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.42 DPS, sim-verified) [crafted] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -3.37 DPS, sim-verified) [crafted] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (211.4 DPS) | yes | Thorium Cestus (250614, -2.29 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -2.29 DPS) [dungeon]; White Bone Shredder (11863, -4.39 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.76 DPS) [world_drop]; Gryphonwing Long Bow (13022, -1.04 DPS) [dungeon]; Dark Iron Rifle (16004, -2.72 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thrash Blade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 874, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 303.8. Weights run: 1.8s. Verify run: 5.4s. 1752 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.105 ± 0.019, strength=1.210 ± 0.002, crit=2.221 ± 0.037 per rating point (14 rating = 1%, 31.093 per %), hit=2.713 ± 0.156 per rating point (10 rating = 1%, 27.127 per %), melee_haste=24.977 ± 2.087

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (303.8 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -14.27 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.4 attack_power points (2.87 DPS) | yes | Woven Ivy Necklace (19159, -0.62 DPS) [quest]; Sunprism Pendant (20649, -0.70 DPS) [quest]; Ethereal Talisman (4430, -1.63 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (303.8 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -3.38 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 51.6 attack_power points (4.43 DPS) | yes | Fel Cape (279269, -1.76 DPS) [crafted]; Dark Phantom Cape (13122, -2.22 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.26 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 136.1 attack_power points (11.69 DPS) | yes | Warlord's Chain Armor (231566, -0.98 DPS) [vendor]; Legionnaire's Chain Armor (227083, -1.89 DPS) [vendor]; Timbermaw Tunic (252484, -3.64 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (303.8 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 59.1 attack_power points (5.08 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -6.65 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 54.9 attack_power points (4.71 DPS) | yes | General's Chain Girdle (16572, -0.94 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.09 DPS) [crafted]; Skulker's Leather Waistguard (252474, -1.24 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 77.1 attack_power points (6.62 DPS) | yes | Legionnaire's Chain Legplates (227079, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; Warbear Woolies (15065, -6.81 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.8 attack_power points (4.71 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 29.5 attack_power points (2.53 DPS) | yes | Ring of the Underwood (2951, -0.52 DPS) [world_drop]; Falcon's Hook (7552, -0.59 DPS) [dungeon]; Blackstone Ring (17713, -0.81 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.06 DPS) | yes | Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (303.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (303.8 DPS) | yes | Molten Heart of the Mountain (249470, -2.67 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (303.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -4.41 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (303.8 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -11.85 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -57.92 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (303.8 DPS) | yes | Eaglehorn Long Bow (13023, -0.72 DPS) [world_drop]; Guttbuster (13139, -0.77 DPS) [world_drop]; Dark Iron Rifle (16004, -3.24 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Magister's Mantle; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Magister's Bindings; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: White Bone Band; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1752, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 702.2. Weights run: 1.5s. Verify run: 3.9s. 1752 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.238 ± 0.020, strength=1.210 ± 0.002, crit=2.552 ± 0.040 per rating point (14 rating = 1%, 35.734 per %), hit=3.280 ± 0.272 per rating point (10 rating = 1%, 32.801 per %), melee_haste=20.128 ± 3.218

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+9.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.85 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.1 attack_power points (6.02 DPS) | yes | Woven Ivy Necklace (19159, -1.33 DPS) [quest]; Sunprism Pendant (20649, -1.42 DPS) [quest]; Ethereal Talisman (4430, -3.45 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 55.8 attack_power points (9.55 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.6 attack_power points (9.17 DPS) | yes | Dark Phantom Cape (13122, -3.42 DPS) [world_drop]; Fel Cape (279269, -3.95 DPS, sim-verified) [crafted]; Blackflame Cape (13109, -4.57 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.3 attack_power points (25.91 DPS) | yes | Warlord's Chain Armor (231566, -3.38 DPS) [vendor]; Legionnaire's Chain Armor (227083, -4.55 DPS) [vendor]; Timbermaw Tunic (252484, -9.85 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 38.4 attack_power points (6.57 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.46 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 63.7 attack_power points (10.91 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -13.84 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 56.7 attack_power points (9.71 DPS) | yes | General's Chain Girdle (16572, -1.78 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.86 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.22 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 81.7 attack_power points (13.99 DPS) | yes | Legionnaire's Chain Legplates (227079, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; Warbear Woolies (15065, -14.08 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 56.8 attack_power points (9.73 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.3 attack_power points (5.36 DPS) | yes | White Bone Band (11862, -1.25 DPS) [quest]; Falcon's Hook (7552, -1.29 DPS) [dungeon]; Blackstone Ring (17713, -1.94 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.8 attack_power points (4.25 DPS) | yes | White Bone Band (11862, -0.14 DPS) [quest]; Falcon's Hook (7552, -0.18 DPS) [dungeon]; Blackstone Ring (17713, -0.82 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+15.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -8.01 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+78.2 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -24.50 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -78.24 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.53 DPS) [world_drop]; Guttbuster (13139, -1.68 DPS) [world_drop]; Dark Iron Rifle (16004, -5.28 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1752, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

