# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 67.2. Weights run: 2.8s. Verify run: 3.9s. 227 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.916 ± 0.019, strength=1.215 ± 0.002, crit=1.373 ± 0.021 per rating point (14 rating = 1%, 19.223 per %), hit=1.115 ± 0.041 per rating point (10 rating = 1%, 11.155 per %), melee_haste=7.286 ± 1.179

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 23.3 attack_power points (1.05 DPS) | yes | Defender's Leather Hood (252447, -0.60 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 5.8 attack_power points (0.26 DPS) | yes | Tarnished Locket (279870, -0.28 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 14.6 attack_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 17.5 attack_power points (0.79 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 32.1 attack_power points (1.44 DPS) | yes | Brawler's Leather Armor (252490, -0.25 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.52 DPS) [crafted]; Prospector's Chestpiece (14562, -0.55 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 14.6 attack_power points (0.65 DPS) | yes | Bravo's Armbands (270015, -0.02 DPS) [quest]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.2 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -1.78 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (67.2 DPS) | yes | Brawler's Leather Belt (252428, -0.07 DPS) [crafted]; Dusty Belt (279897, -0.15 DPS) [quest]; Deviate Scale Belt (6468, -2.58 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.2 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.07 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (67.2 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.03 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.8 attack_power points (0.26 DPS) | yes | The 1 Ring (8350, -0.08 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.9 attack_power points (0.22 DPS) | yes | The 1 Ring (8350, -0.03 DPS) [world]; Ring of the Moon (12052, -0.16 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (11.17 DPS) | yes | Night Reaver (1318, +0.00 DPS) [dungeon]; Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.4 attack_power points (10.70 DPS) | yes | Assassin's Blade (1935, -0.26 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.7 attack_power points (0.52 DPS) | yes | Light Bow (4576, -0.26 DPS) [world_drop]; Owlsight Rifle (15205, -0.26 DPS) [quest]; Deadly Blunderbuss (4369, -0.28 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Lavishly Jeweled Ring; finger2: Demon Band; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 227, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230231050000000)

Set DPS (verified): 95.7. Weights run: 2.9s. Verify run: 4.3s. 434 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.454 ± 0.015, strength=1.213 ± 0.002, crit=1.468 ± 0.021 per rating point (14 rating = 1%, 20.555 per %), hit=1.293 ± 0.051 per rating point (10 rating = 1%, 12.928 per %), melee_haste=6.132 ± 1.167

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.5 attack_power points (1.30 DPS) | yes | Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world]; Defender's Leather Helm (252455, -0.53 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.7 attack_power points (0.78 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; Erudite's Amulet (277204, -0.52 DPS) [quest]; River Pride Choker (13087, -0.52 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.1 attack_power points (1.75 DPS) | yes | Mantle of Thieves (2264, -0.49 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.78 DPS) [world_drop]; Barbaric Shoulders (5964, -0.78 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.8 attack_power points (1.10 DPS) | yes | Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Fenrus' Hide (6340, -0.32 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.32 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 34.4 attack_power points (1.82 DPS) | yes | Brawler's Leather Tunic (252508, -0.39 DPS) [crafted]; Tunic of Westfall (2041, -0.41 DPS, sim-verified) [quest]; Brawler's Leather Armor (252490, -0.59 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.6 attack_power points (1.04 DPS) | yes | Insignia Bracers (6410, -0.26 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS) [crafted]; Demonhide Bracers (270033, -0.32 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (95.7 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Cartilage Shapers (273054, +0.00 DPS) [world]; Insignia Gloves (6408, -0.41 DPS, sim-verified) [world_drop] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 29.4 attack_power points (1.55 DPS) | yes | Prowler's Leather Belt (252459, -0.20 DPS) [crafted]; Stalker's Leather Belt (252521, -0.38 DPS) [crafted]; Warden's Leather Belt (252460, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (95.7 DPS) | yes | Dusky Leather Leggings (7373, +0.00 DPS) [crafted]; Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (95.7 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.90 DPS) | yes | Defias Renegade Ring (1076, -0.25 DPS) [dungeon]; Monkey Ring (6748, -0.25 DPS) [quest]; Seal of Wrynn (2933, -0.32 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 14.7 attack_power points (0.78 DPS) | yes | Defias Renegade Ring (1076, -0.13 DPS) [dungeon]; Monkey Ring (6748, -0.13 DPS) [quest]; Seal of Wrynn (2933, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (18.07 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Swinetusk Shank (6691, -2.99 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (95.7 DPS) | yes | Shoni's Disarming Tool (9608, -4.95 DPS) [quest]; Swinetusk Shank (6691, -16.90 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.7 attack_power points (0.78 DPS) | yes | Double-barreled Shotgun (2098, -0.20 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop]; Lil Timmy's Peashooter (13136, -0.26 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Gloves of the Fang; waist: Skulker's Leather Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 149.9. Weights run: 3.0s. Verify run: 2.6s. 690 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.143 ± 0.013, strength=1.210 ± 0.002, crit=1.441 ± 0.021 per rating point (14 rating = 1%, 20.167 per %), hit=1.332 ± 0.079 per rating point (10 rating = 1%, 13.318 per %), melee_haste=not significant (3.107 ± 1.094)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 36.9 attack_power points (2.70 DPS) | yes | Barbaric Iron Helm (7915, -0.49 DPS) [crafted]; Hawkeye's Helm (14591, -0.62 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.66 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.4 attack_power points (0.98 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; River Pride Choker (13087, -0.63 DPS) [world_drop]; Gazlowe's Charm (13088, -0.63 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.6 attack_power points (2.17 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.60 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 20.1 attack_power points (1.47 DPS) | yes | Imperial Cloak (6432, -0.22 DPS) [dungeon]; Parachute Cloak (10518, -0.22 DPS) [crafted]; Hawkeye's Cloak (14593, -1.56 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 51.4 attack_power points (3.77 DPS) | yes | Nightscape Tunic (8175, -1.41 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.41 DPS) [crafted]; Wolffear Harness (13110, -1.48 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.47 DPS) | yes | Hawkeye's Bracers (14590, -0.17 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.21 DPS) [dungeon]; Dusky Bracers (7378, -0.21 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 40.2 attack_power points (2.94 DPS) | yes | Dragonscale Gauntlets (8347, -0.52 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.58 DPS) [crafted]; Scarlet Gauntlets (10331, -0.58 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 30.2 attack_power points (2.21 DPS) | yes | Skulker's Leather Belt (252520, -0.27 DPS) [crafted]; Scorpashi Sash (14652, -0.29 DPS) [world_drop]; Officer's Belt (250556, -0.38 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.0 attack_power points (3.30 DPS) | yes | Triprunner Dungarees (9624, -0.21 DPS) [quest]; Petrolspill Leggings (9509, -1.10 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.10 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 32.3 attack_power points (2.37 DPS) | yes | Skulker's Leather Shoes (252531, -0.02 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.29 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 23.8 attack_power points (1.75 DPS) | yes | Mark of Kern (2262, -0.28 DPS) [dungeon]; Assault Band (13095, -0.28 DPS) [world_drop]; Disengagement Ring (276202, -0.49 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.9 attack_power points (1.68 DPS) | yes | Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop]; Disengagement Ring (276202, -0.42 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (149.9 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -7.28 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (32.40 DPS) | yes | Shoni's Disarming Tool (9608, -15.93 DPS) [quest]; Dazzling Longsword (869, -19.41 DPS, sim-verified) [world_drop] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (149.9 DPS) | yes | The Silencer (13138, -0.07 DPS) [world_drop]; Monolithic Bow (9426, -0.10 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.17 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 0000000000000000-3210000000000000-500230231051222151)

Set DPS (verified): 200.9. Weights run: 3.0s. Verify run: 2.9s. 879 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.201 ± 0.013, strength=1.210 ± 0.002, crit=1.641 ± 0.023 per rating point (14 rating = 1%, 22.969 per %), hit=1.643 ± 0.095 per rating point (10 rating = 1%, 16.428 per %), melee_haste=not significant (4.864 ± 1.317)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 37.5 attack_power points (2.89 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.01 DPS) [crafted]; Knight-Lieutenant's Mail Helmet (223075, -0.10 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.7 attack_power points (2.67 DPS) | yes | Kaleidoscope Chain (13084, -1.62 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.74 DPS) [world]; River Pride Choker (13087, -2.30 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 35.1 attack_power points (2.70 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest]; Knight-Lieutenant's Mail Epaulets (223073, -0.10 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 33.0 attack_power points (2.54 DPS) | yes | Blackflame Cape (13109, -0.51 DPS) [world_drop]; Serpentskin Cloak (8259, -0.68 DPS) [world_drop]; Pridelord Cape (14673, -0.74 DPS) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | 56.0 attack_power points (4.31 DPS) | yes | Warbear Harness (15064, -0.24 DPS) [crafted]; Icemail Jerkin (1981, -0.24 DPS) [world_drop]; Blazewind Breastplate (11193, -2.99 DPS, sim-verified) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.9 attack_power points (2.91 DPS) | yes | Skulker's Leather Bracers (252540, -0.74 DPS) [crafted]; Pridelord Bands (14672, -0.75 DPS) [world_drop]; Arena Bands (18711, -0.76 DPS) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 43.0 attack_power points (3.31 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.10 DPS) [crafted]; Serpentskin Gloves (8260, -0.42 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.48 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 45.3 attack_power points (3.49 DPS) | yes | Prowler's Leather Waistguard (252473, -0.15 DPS) [crafted]; Sagebrush Girdle (17778, -0.95 DPS) [quest]; Pridelord Girdle (14674, -1.08 DPS) [world_drop] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 53.8 attack_power points (4.14 DPS) | yes | Serpentskin Leggings (8262, -0.22 DPS) [world_drop]; Basilisk Hide Pants (1718, -0.58 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -3.11 DPS, sim-verified) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 44.0 attack_power points (3.39 DPS) | yes | Skulker's Leather Boots (252469, -0.16 DPS) [crafted]; Fleetfoot Greaves (11627, -0.17 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.31 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.8 attack_power points (2.37 DPS) | yes | Falcon's Hook (7552, -0.57 DPS) [dungeon]; Mark of Kern (2262, -0.83 DPS) [dungeon]; Blackstone Ring (17713, -0.83 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.4 attack_power points (1.88 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Mark of Kern (2262, -0.34 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (200.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (200.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (200.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -3.39 DPS, sim-verified) [crafted] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (40.53 DPS) | yes | Julie's Dagger (6660, +0.00 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.41 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.23 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (200.9 DPS) | yes | Guttbuster (13139, -0.74 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.02 DPS) [world_drop]; Dark Iron Rifle (16004, -2.65 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thrash Blade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 290.0. Weights run: 2.2s. Verify run: 5.0s. 1749 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.063 ± 0.018, strength=1.210 ± 0.002, crit=2.136 ± 0.037 per rating point (14 rating = 1%, 29.911 per %), hit=2.828 ± 0.159 per rating point (10 rating = 1%, 28.275 per %), melee_haste=22.415 ± 2.111

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+18.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -18.78 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.9 attack_power points (2.79 DPS) | yes | Sunprism Pendant (20649, -0.69 DPS) [quest]; Kaleidoscope Chain (13084, -1.68 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.77 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 51.0 attack_power points (4.33 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.9 attack_power points (4.32 DPS) | yes | Dark Phantom Cape (13122, -1.26 DPS, sim-verified) [world_drop]; Fel Cape (279269, -1.78 DPS) [crafted]; Sergeant Major's Cape (16337, -2.10 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.5 attack_power points (11.40 DPS) | yes | Field Marshal's Chain Armor (231563, -1.00 DPS) [vendor]; Timbermaw Tunic (252484, -1.96 DPS, sim-verified) [crafted]; Knight-Captain's Chain Armor (227089, -1.96 DPS) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 35.8 attack_power points (3.03 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Beaststalker's Bindings (16681, -0.41 DPS) [dungeon] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 57.9 attack_power points (4.91 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -4.58 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 54.3 attack_power points (4.60 DPS) | yes | Skulker's Leather Waistguard (252474, -0.92 DPS) [crafted]; Marshal's Chain Girdle (16464, -0.94 DPS) [pvp]; Heavy Timbermaw Belt (19043, -1.04 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 75.9 attack_power points (6.43 DPS) | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Warbear Woolies (15065, -4.64 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.5 attack_power points (4.62 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.9 attack_power points (2.45 DPS) | yes | Falcon's Hook (7552, -0.57 DPS) [dungeon]; Mark of Kern (2262, -0.75 DPS) [dungeon]; Blackstone Ring (17713, -0.75 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.0 attack_power points (1.95 DPS) | yes | Falcon's Hook (7552, -0.07 DPS) [dungeon]; Mark of Kern (2262, -0.26 DPS) [dungeon]; Blackstone Ring (17713, -0.26 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+52.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -11.58 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -52.01 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -0.70 DPS) [world_drop]; Guttbuster (13139, -0.74 DPS) [world_drop]; Dark Iron Rifle (16004, -2.21 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 687.2. Weights run: 1.8s. Verify run: 4.3s. 1749 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.196 ± 0.020, strength=1.210 ± 0.002, crit=2.447 ± 0.040 per rating point (14 rating = 1%, 34.260 per %), hit=3.527 ± 0.278 per rating point (10 rating = 1%, 35.267 per %), melee_haste=16.406 ± 3.267

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+9.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.80 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.6 attack_power points (5.86 DPS) | yes | Sunprism Pendant (20649, -1.40 DPS) [quest]; Kaleidoscope Chain (13084, -3.55 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.82 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 55.4 attack_power points (9.37 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Dark Warder's Pauldrons (22241, -3.22 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 52.9 attack_power points (8.96 DPS) | yes | Dark Phantom Cape (13122, -3.39 DPS) [world_drop]; Fel Cape (279269, -4.04 DPS, sim-verified) [crafted]; Sergeant Major's Cape (16337, -4.35 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 150.8 attack_power points (25.52 DPS) | yes | Field Marshal's Chain Armor (231563, -3.62 DPS) [vendor]; Knight-Captain's Chain Armor (227089, -4.98 DPS) [vendor]; Timbermaw Tunic (252484, -6.10 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.8 attack_power points (6.39 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.59 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 62.3 attack_power points (10.54 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -16.31 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 56.1 attack_power points (9.51 DPS) | yes | Marshal's Chain Girdle (16464, -1.79 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.84 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.18 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 80.3 attack_power points (13.59 DPS) | yes | Knight-Captain's Chain Legplates (227085, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Warbear Woolies (15065, -16.60 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 56.7 attack_power points (9.60 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.7 attack_power points (5.20 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Mark of Kern (2262, -1.82 DPS) [dungeon]; Blackstone Ring (17713, -1.82 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.4 attack_power points (4.13 DPS) | yes | Falcon's Hook (7552, -0.17 DPS) [dungeon]; Mark of Kern (2262, -0.74 DPS) [dungeon]; Blackstone Ring (17713, -0.74 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+83.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -23.95 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -83.70 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.49 DPS) [world_drop]; Guttbuster (13139, -1.62 DPS) [world_drop]; Dark Iron Rifle (16004, -5.17 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 67.8. Weights run: 2.8s. Verify run: 3.7s. 223 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.916 ± 0.019, strength=1.215 ± 0.002, crit=1.373 ± 0.021 per rating point (14 rating = 1%, 19.223 per %), hit=1.115 ± 0.041 per rating point (10 rating = 1%, 11.155 per %), melee_haste=7.286 ± 1.179

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 23.3 attack_power points (1.05 DPS) | yes | Defender's Leather Hood (252447, -0.60 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 5.8 attack_power points (0.26 DPS) | yes | Roadwatcher's Confidence (281265, -0.27 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 14.6 attack_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 17.5 attack_power points (0.79 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (67.8 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.75 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 14.6 attack_power points (0.65 DPS) | yes | Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 22.4 attack_power points (1.00 DPS) | yes | Serpent Gloves (5970, -0.22 DPS) [dungeon]; Bristlebark Gloves (14572, -0.26 DPS) [world_drop]; Fletcher's Gloves (7348, -0.26 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (67.8 DPS) | yes | Brawler's Leather Belt (252428, -0.07 DPS) [crafted]; Dark Leather Belt (4249, -0.28 DPS) [crafted]; Deviate Scale Belt (6468, -2.45 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.8 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.95 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (67.8 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.89 DPS, sim-verified) [world_drop] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 8.7 attack_power points (0.39 DPS) | yes | Demon Band (12054, -0.17 DPS) [world_drop]; The 1 Ring (8350, -0.21 DPS) [world]; Loop of Sacrifice (281673, -0.23 DPS) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.8 attack_power points (0.26 DPS) | yes | The 1 Ring (8350, -0.08 DPS) [world]; Loop of Sacrifice (281673, -0.10 DPS) [quest]; Demon Band (12054, -0.27 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (11.17 DPS) | yes | Crescent Staff (6505, +0.00 DPS) [quest]; Living Root (6631, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.4 attack_power points (10.70 DPS) | yes | Wingblade (6504, -5.90 DPS, sim-verified) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.7 attack_power points (0.52 DPS) | yes | Light Bow (4576, -0.26 DPS) [world_drop]; Privateer Musket (5309, -0.26 DPS) [quest]; Deadly Blunderbuss (4369, -0.27 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Bounty Hunter's Ring; finger2: Lavishly Jeweled Ring; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 223, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 0000000000000000-0000000000000000-500230231050000000)

Set DPS (verified): 96.0. Weights run: 2.9s. Verify run: 2.5s. 431 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.454 ± 0.015, strength=1.213 ± 0.002, crit=1.468 ± 0.021 per rating point (14 rating = 1%, 20.555 per %), hit=1.293 ± 0.051 per rating point (10 rating = 1%, 12.928 per %), melee_haste=6.132 ± 1.167

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.5 attack_power points (1.30 DPS) | yes | Brawler's Leather Hood (252504, -0.21 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world]; Defender's Leather Helm (252455, -0.53 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.7 attack_power points (0.78 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; Erudite's Amulet (277204, -0.52 DPS) [quest]; River Pride Choker (13087, -0.52 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.1 attack_power points (1.75 DPS) | yes | Mantle of Thieves (2264, -0.31 DPS, sim-verified) [dungeon]; Bristlebark Amice (14573, -0.78 DPS) [world_drop]; Barbaric Shoulders (5964, -0.78 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.8 attack_power points (1.10 DPS) | yes | Tigerstrike Mantle (13108, -0.06 DPS) [world_drop]; Fenrus' Hide (6340, -0.32 DPS) [dungeon]; Swiftrunner Cape (6745, -0.32 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 34.4 attack_power points (1.82 DPS) | yes | Brawler's Leather Tunic (252508, -0.26 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.52 DPS) [quest]; Brawler's Leather Armor (252490, -0.59 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.6 attack_power points (1.04 DPS) | yes | Insignia Bracers (6410, -0.26 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS) [crafted]; Loamflake Bracers (15462, -0.39 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.2 attack_power points (1.23 DPS) | yes | Cartilage Shapers (273054, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.06 DPS) [crafted]; Wolfclaw Gloves (1978, -0.13 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 29.4 attack_power points (1.55 DPS) | yes | Prowler's Leather Belt (252459, -0.20 DPS) [crafted]; Stalker's Leather Belt (252521, -0.38 DPS) [crafted]; Warden's Leather Belt (252460, -0.39 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 34.4 attack_power points (1.82 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.20 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.3 attack_power points (1.23 DPS) | yes | Insignia Boots (4055, -0.19 DPS) [world_drop]; Warsong Boots (16977, -0.19 DPS) [quest]; Brawler's Leather Boots (252439, -0.26 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.1 attack_power points (0.90 DPS) | yes | Band of the Fist (17694, -0.13 DPS) [quest]; Defias Renegade Ring (1076, -0.25 DPS) [dungeon]; Monkey Ring (6748, -0.25 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 14.7 attack_power points (0.78 DPS) | yes | Band of the Fist (17694, -0.00 DPS) [quest]; Defias Renegade Ring (1076, -0.13 DPS) [dungeon]; Monkey Ring (6748, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 341.5 attack_power points (18.07 DPS) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (96.0 DPS) | yes | Swinetusk Shank (6691, -17.70 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.7 attack_power points (0.78 DPS) | yes | Double-barreled Shotgun (2098, -0.20 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop]; Lil Timmy's Peashooter (13136, -0.26 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 147.4. Weights run: 3.0s. Verify run: 2.5s. 678 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.143 ± 0.013, strength=1.210 ± 0.002, crit=1.441 ± 0.021 per rating point (14 rating = 1%, 20.167 per %), hit=1.332 ± 0.079 per rating point (10 rating = 1%, 13.318 per %), melee_haste=not significant (3.107 ± 1.094)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 36.9 attack_power points (2.70 DPS) | yes | Barbaric Iron Helm (7915, -0.53 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.62 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.66 DPS) [world_drop] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 14.6 attack_power points (1.07 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; River Pride Choker (13087, -0.72 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.6 attack_power points (2.17 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.60 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 18.6 attack_power points (1.37 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Parachute Cloak (10518, -0.11 DPS) [crafted]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 51.4 attack_power points (3.77 DPS) | yes | Wolffear Harness (13110, -1.10 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.41 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.41 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.47 DPS) | yes | Hawkeye's Bracers (14590, -0.17 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.21 DPS) [dungeon]; Dusky Bracers (7378, -0.21 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 40.2 attack_power points (2.94 DPS) | yes | Dragonscale Gauntlets (8347, -0.52 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.58 DPS) [crafted]; Scarlet Gauntlets (10331, -0.62 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 30.2 attack_power points (2.21 DPS) | yes | Skulker's Leather Belt (252520, -0.27 DPS) [crafted]; Scorpashi Sash (14652, -0.29 DPS) [world_drop]; Officer's Belt (250556, -0.38 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 45.0 attack_power points (3.30 DPS) | yes | Triprunner Dungarees (9624, -0.21 DPS) [quest]; Petrolspill Leggings (9509, -1.10 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.10 DPS) [world_drop] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | 32.3 attack_power points (2.37 DPS) | yes | Skulker's Leather Shoes (252531, -0.02 DPS) [crafted]; Imperial Leather Boots (6431, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.29 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 23.8 attack_power points (1.75 DPS) | yes | Mark of Kern (2262, -0.28 DPS) [dungeon]; Assault Band (13095, -0.28 DPS) [world_drop]; Disengagement Ring (276202, -0.49 DPS) [vendor] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.9 attack_power points (1.68 DPS) | yes | Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop]; Disengagement Ring (276202, -0.42 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (147.4 DPS) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -5.48 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (32.40 DPS) | yes | Dazzling Longsword (869, -19.39 DPS, sim-verified) [world_drop] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (147.4 DPS) | yes | The Silencer (13138, -0.07 DPS) [world_drop]; Monolithic Bow (9426, -0.10 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.17 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 678, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 0000000000000000-3210000000000000-500230231051222151)

Set DPS (verified): 208.3. Weights run: 3.0s. Verify run: 2.8s. 868 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.201 ± 0.013, strength=1.210 ± 0.002, crit=1.641 ± 0.023 per rating point (14 rating = 1%, 22.969 per %), hit=1.643 ± 0.095 per rating point (10 rating = 1%, 16.428 per %), melee_haste=not significant (4.864 ± 1.317)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 37.5 attack_power points (2.89 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.01 DPS) [crafted]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.7 attack_power points (2.67 DPS) | yes | Woven Ivy Necklace (19159, -0.91 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -1.53 DPS) [quest]; Kaleidoscope Chain (13084, -1.62 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 35.1 attack_power points (2.70 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest]; Blood Guard's Mail Epaulets (220823, -0.10 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 33.0 attack_power points (2.54 DPS) | yes | Serpentskin Cloak (8259, -0.68 DPS) [world_drop]; Pridelord Cape (14673, -0.74 DPS) [world_drop]; Blackflame Cape (13109, -1.10 DPS, sim-verified) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) | PvP rank 10 · Stone Guard · Horde [vendor] | 56.0 attack_power points (4.31 DPS) | yes | Warbear Harness (15064, -0.24 DPS) [crafted]; Icemail Jerkin (1981, -0.24 DPS) [world_drop]; Blazewind Breastplate (11193, -3.30 DPS, sim-verified) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.9 attack_power points (2.91 DPS) | yes | Pridelord Bands (14672, -0.75 DPS) [world_drop]; Arena Bands (18711, -0.76 DPS) [world]; Skulker's Leather Bracers (252540, -0.91 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 43.0 attack_power points (3.31 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.10 DPS) [crafted]; Serpentskin Gloves (8260, -0.42 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.48 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 45.3 attack_power points (3.49 DPS) | yes | Prowler's Leather Waistguard (252473, -0.15 DPS) [crafted]; Sagebrush Girdle (17778, -0.95 DPS) [quest]; Pridelord Girdle (14674, -1.08 DPS) [world_drop] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 53.8 attack_power points (4.14 DPS) | yes | Basilisk Hide Pants (1718, -0.58 DPS) [world_drop]; Stormshroud Pants (15057, -0.60 DPS) [crafted]; Serpentskin Leggings (8262, -3.46 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 44.0 attack_power points (3.39 DPS) | yes | Skulker's Leather Boots (252469, -0.16 DPS) [crafted]; Fleetfoot Greaves (11627, -0.17 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.31 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.8 attack_power points (2.37 DPS) | yes | White Bone Band (11862, -0.52 DPS) [quest]; Falcon's Hook (7552, -0.57 DPS) [dungeon]; Blackstone Ring (17713, -0.83 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.4 attack_power points (1.88 DPS) | yes | White Bone Band (11862, -0.03 DPS) [quest]; Falcon's Hook (7552, -0.08 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (208.3 DPS) | yes | Molten Heart of the Mountain (249470, -4.93 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (208.3 DPS) | yes | Molten Heart of the Mountain (249470, -1.21 DPS, sim-verified) [crafted] |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (208.3 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -4.82 DPS, sim-verified) [crafted] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (40.53 DPS) | yes | Julie's Dagger (6660, +0.00 DPS) [world_drop]; Claw of Celebras (17738, -3.41 DPS) [dungeon]; White Bone Shredder (11863, -5.48 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (208.3 DPS) | yes | Guttbuster (13139, -0.74 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.02 DPS) [world_drop]; Dark Iron Rifle (16004, -2.80 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thrash Blade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 301.1. Weights run: 2.2s. Verify run: 6.2s. 1746 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.063 ± 0.018, strength=1.210 ± 0.002, crit=2.136 ± 0.037 per rating point (14 rating = 1%, 29.911 per %), hit=2.828 ± 0.159 per rating point (10 rating = 1%, 28.275 per %), melee_haste=22.415 ± 2.111

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (301.1 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -15.22 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.9 attack_power points (2.79 DPS) | yes | Woven Ivy Necklace (19159, -0.60 DPS) [quest]; Sunprism Pendant (20649, -0.69 DPS) [quest]; Ethereal Talisman (4430, -1.57 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (301.1 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -2.85 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.9 attack_power points (4.32 DPS) | yes | Fel Cape (279269, -1.78 DPS) [crafted]; Dark Phantom Cape (13122, -2.11 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.22 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.5 attack_power points (11.40 DPS) | yes | Warlord's Chain Armor (231566, -1.00 DPS) [vendor]; Legionnaire's Chain Armor (227083, -1.96 DPS) [vendor]; Timbermaw Tunic (252484, -3.45 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (301.1 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 57.9 attack_power points (4.91 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -6.19 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 54.3 attack_power points (4.60 DPS) | yes | General's Chain Girdle (16572, -0.94 DPS) [pvp]; Heavy Timbermaw Belt (19043, -1.04 DPS) [crafted]; Skulker's Leather Waistguard (252474, -1.23 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 75.9 attack_power points (6.43 DPS) | yes | Legionnaire's Chain Legplates (227079, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; Warbear Woolies (15065, -6.36 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.5 attack_power points (4.62 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.9 attack_power points (2.45 DPS) | yes | Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.57 DPS) [dungeon]; Blackstone Ring (17713, -0.75 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.03 DPS) | yes | Ring of the Underwood (2951, -0.08 DPS) [world_drop]; Falcon's Hook (7552, -0.15 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (301.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (301.1 DPS) | yes | Molten Heart of the Mountain (249470, -2.49 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (301.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -4.73 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (301.1 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -11.58 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -54.91 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (301.1 DPS) | yes | Eaglehorn Long Bow (13023, -0.70 DPS) [world_drop]; Guttbuster (13139, -0.74 DPS) [world_drop]; Dark Iron Rifle (16004, -3.17 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Devout Mantle; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Devout Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: White Bone Band; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 697.7. Weights run: 1.8s. Verify run: 4.3s. 1746 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.196 ± 0.020, strength=1.210 ± 0.002, crit=2.447 ± 0.040 per rating point (14 rating = 1%, 34.260 per %), hit=3.527 ± 0.278 per rating point (10 rating = 1%, 35.267 per %), melee_haste=16.406 ± 3.267

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.72 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.6 attack_power points (5.86 DPS) | yes | Woven Ivy Necklace (19159, -1.28 DPS) [quest]; Sunprism Pendant (20649, -1.40 DPS) [quest]; Ethereal Talisman (4430, -3.35 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 55.4 attack_power points (9.37 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -3.01 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 52.9 attack_power points (8.96 DPS) | yes | Dark Phantom Cape (13122, -3.39 DPS) [world_drop]; Fel Cape (279269, -4.41 DPS, sim-verified) [crafted]; Blackflame Cape (13109, -4.50 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 150.8 attack_power points (25.52 DPS) | yes | Warlord's Chain Armor (231566, -3.62 DPS) [vendor]; Legionnaire's Chain Armor (227083, -4.98 DPS) [vendor]; Timbermaw Tunic (252484, -7.53 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.8 attack_power points (6.39 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.59 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 62.3 attack_power points (10.54 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -14.66 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 56.1 attack_power points (9.51 DPS) | yes | General's Chain Girdle (16572, -1.79 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.84 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.18 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 80.3 attack_power points (13.59 DPS) | yes | Legionnaire's Chain Legplates (227079, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; Warbear Woolies (15065, -14.86 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 56.7 attack_power points (9.60 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.7 attack_power points (5.20 DPS) | yes | White Bone Band (11862, -1.14 DPS) [quest]; Falcon's Hook (7552, -1.24 DPS) [dungeon]; Blackstone Ring (17713, -1.82 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.4 attack_power points (4.13 DPS) | yes | White Bone Band (11862, -0.06 DPS) [quest]; Falcon's Hook (7552, -0.17 DPS) [dungeon]; Blackstone Ring (17713, -0.74 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -5.34 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+82.9 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -23.95 DPS) [dungeon]; Teebu's Blazing Longsword (1728, -82.93 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.49 DPS) [world_drop]; Guttbuster (13139, -1.62 DPS) [world_drop]; Dark Iron Rifle (16004, -5.07 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; trinket2: Molten Heart of the Mountain; main_hand: Ironwood Blade; off_hand: Heartseeker

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

