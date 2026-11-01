# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 36.9. Weights run: 2.5s. Verify run: 3.9s. 300 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.271 ± 0.028, crit=0.759 ± 0.028 per rating point (14 rating = 1%, 10.620 per %), hit=0.982 ± 0.056 per rating point (10 rating = 1%, 9.821 per %), melee_haste=5.774 ± 0.588

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.20 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.44 DPS) [crafted]; Brawler's Leather Hood (252504, -0.59 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.5 attack_power points (0.15 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.4 attack_power points (0.38 DPS) | yes | Rough Bronze Shoulders (3480, -0.02 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.02 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.5 attack_power points (0.51 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.06 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.44 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.60 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.06 DPS) [quest]; Cryptwalker Bracers (280095, -0.12 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.77 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.08 DPS) | yes | Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Deviate Scale Belt (6468, -0.34 DPS) [crafted]; Cobrahn's Grasp (6460, -3.04 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.75 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.98 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Defender's Leather Boots (252441, -0.38 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.48 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.3 attack_power points (0.20 DPS) | yes | Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Ring of the Moon (12052, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.0 attack_power points (14.39 DPS) | yes | Rockslicer (872, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Diamond Hammer (2194, -9.32 DPS, sim-verified) [world_drop] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.9 DPS) | yes | Diamond Hammer (2194, -0.37 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.1 attack_power points (0.31 DPS) | yes | Dwarven Fishing Pole (3567, -0.07 DPS) [quest]; Fine Longbow (11304, -0.07 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.07 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: The 1 Ring; main_hand: Butcher's Cleaver; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 64.4. Weights run: 2.7s. Verify run: 2.2s. 558 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.389 ± 0.038, crit=1.048 ± 0.044 per rating point (14 rating = 1%, 14.666 per %), hit=1.363 ± 0.083 per rating point (10 rating = 1%, 13.634 per %), melee_haste=7.721 ± 0.909

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 30.5 attack_power points (1.89 DPS) | yes | Tusken Helm (6686, -0.28 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.40 DPS) [crafted]; Defender's Leather Helm (252455, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.6 attack_power points (0.84 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; River Pride Choker (13087, -0.34 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 25.3 attack_power points (1.56 DPS) | yes | Barbaric Iron Shoulders (7913, -0.31 DPS) [crafted]; Barbaric Shoulders (5964, -0.52 DPS) [crafted]; Bristlebark Amice (14573, -0.68 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.7 attack_power points (0.97 DPS) | yes | Sergeant Major's Cape (16315, -0.13 DPS) [pvp]; Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.35 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.86 DPS) | yes | Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.29 DPS, sim-verified) [crafted]; Slayer's Surcoat (14751, -0.36 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 16.3 attack_power points (1.01 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Pugilist Bracers (4438, -0.02 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.3 attack_power points (1.38 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.02 DPS) [world]; Mail Combat Gauntlets (4075, -0.05 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 28.3 attack_power points (1.75 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted]; Girdle of Golem Strength (9405, -0.27 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 31.7 attack_power points (1.96 DPS) | yes | Brawler's Leather Legguards (252516, -0.33 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.35 DPS) [dungeon]; Glimmering Mail Legguards (6386, -0.37 DPS) [world_drop] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 attack_power points (1.06 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Alacritous Treads (277234, -0.01 DPS) [quest]; Veteran's Boots (250503, -0.10 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.2 attack_power points (1.25 DPS) | yes | Silverlaine's Family Seal (6321, -0.63 DPS) [dungeon]; Tiger Band (6749, -0.63 DPS) [quest]; Bonebinder's Signet (273032, -0.63 DPS) [world] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 10.2 attack_power points (0.63 DPS) | yes | Silverlaine's Family Seal (6321, -0.01 DPS) [dungeon]; Tiger Band (6749, -0.01 DPS) [quest]; Bonebinder's Signet (273032, -0.01 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 330.8 attack_power points (20.46 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (64.4 DPS) | yes | Shoni's Disarming Tool (9608, -1.91 DPS) [quest]; Swinetusk Shank (6691, -22.93 DPS, sim-verified) [dungeon] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.2 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.13 DPS) [world]; Golemsight Long Gun (273029, -0.22 DPS, sim-verified) [dungeon]; Long Battle Bow (15284, -0.26 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 558, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 40 (human, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 125.0. Weights run: 3.0s. Verify run: 2.8s. 788 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.755 ± 0.070, crit=1.717 ± 0.093 per rating point (14 rating = 1%, 24.039 per %), hit=1.425 ± 0.147 per rating point (10 rating = 1%, 14.248 per %), melee_haste=10.697 ± 1.274

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 45.1 attack_power points (3.88 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Barbaric Iron Helm (7915, -0.97 DPS) [crafted]; Hard Gold Coif (250537, -1.47 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.0 attack_power points (1.29 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.26 DPS) [world]; River Pride Choker (13087, -0.60 DPS) [world_drop]; Gazlowe's Charm (13088, -0.60 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.3 attack_power points (2.52 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.58 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.63 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 22.5 attack_power points (1.94 DPS) | yes | Hawkeye's Cloak (14593, -0.37 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.52 DPS) [quest]; Parachute Cloak (10518, -0.73 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 42.1 attack_power points (3.63 DPS) | yes | Kolkar Marauder Chain (6773, -0.16 DPS) [quest]; Carapace of Tuten'kash (10775, -0.70 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -1.00 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.0 attack_power points (1.81 DPS) | yes | Branded Leather Bracers (19508, -0.09 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.21 DPS) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 44.0 attack_power points (3.79 DPS) | yes | Scarlet Gauntlets (10331, -0.21 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.56 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.71 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 33.8 attack_power points (2.91 DPS) | yes | Officer's Belt (250556, -0.28 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted]; Skulker's Leather Belt (252520, -0.52 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.62 DPS) | yes | Symbolic Legplates (14829, -0.13 DPS) [world_drop]; Firemane Leggings (13129, -0.34 DPS) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 38.3 attack_power points (3.30 DPS) | yes | Blackforge Greaves (6423, -0.24 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.43 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 21.8 attack_power points (1.88 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.15 DPS) [dungeon]; Assault Band (13095, -0.15 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 21.5 attack_power points (1.86 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (125.0 DPS) | yes | Shoni's Disarming Tool (9608, -13.32 DPS) [quest]; Ardent Custodian (868, -40.42 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.19 DPS) [quest]; The Silencer (13138, -0.28 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.15 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 788, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 184.0. Weights run: 3.3s. Verify run: 3.7s. 1011 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.601 ± 0.088, crit=1.844 ± 0.121 per rating point (14 rating = 1%, 25.814 per %), hit=3.013 ± 0.195 per rating point (10 rating = 1%, 30.130 per %), melee_haste=13.962 ± 1.827

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 45.8 attack_power points (4.27 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.24 DPS) [dungeon]; Sunscale Helmet (14849, -0.39 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.8 attack_power points (2.87 DPS) | yes | Kaleidoscope Chain (13084, -1.57 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -1.75 DPS) [world]; River Pride Choker (13087, -2.13 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 38.8 attack_power points (3.62 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.19 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.22 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 24.0 attack_power points (2.24 DPS) | yes | Sergeant Major's Cape (16336, -0.22 DPS) [pvp]; Pridelord Cape (14673, -0.34 DPS) [world_drop]; Blackflame Cape (13109, -0.45 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 55.8 attack_power points (5.20 DPS) | yes | Warbear Harness (15064, -0.47 DPS) [crafted]; Warforged Chestplate (11195, -0.73 DPS) [quest]; Blazewind Breastplate (11193, -1.21 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.0 attack_power points (2.98 DPS) | yes | Officer's Wristguards (250581, -0.04 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.26 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 46.4 attack_power points (4.32 DPS) | yes | Gloves of Holy Might (867, -0.06 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.37 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -0.56 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 47.2 attack_power points (4.40 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Officer's Waistguard (250571, -0.30 DPS) [crafted]; Belt of the Gladiator (13134, -1.04 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 65.8 attack_power points (6.13 DPS) | yes | Centurion Legplates (10740, -1.10 DPS) [quest]; Gryphon Rider's Leggings (9652, -1.25 DPS, sim-verified) [quest]; Stormshroud Pants (15057, -1.32 DPS) [crafted] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 43.6 attack_power points (4.06 DPS) | yes | Skulker's Leather Boots (252469, -0.07 DPS) [crafted]; Officer's Sabatons (250561, -0.30 DPS) [crafted]; Officer's Boots (250546, -0.60 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.4 attack_power points (2.09 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Ring of the Underwood (2951, -0.22 DPS) [world_drop]; Blackstone Ring (17713, -0.23 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (1.94 DPS) | yes | Falcon's Hook (7552, -0.04 DPS) [dungeon]; Ring of the Underwood (2951, -0.07 DPS) [world_drop]; Blackstone Ring (17713, -0.07 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (184.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (184.0 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (184.0 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 535.8 attack_power points (49.92 DPS) | yes | Thorium Cestus (250614, -4.23 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -4.99 DPS) [dungeon]; Shoni's Disarming Tool (9608, -23.15 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (184.0 DPS) | yes | Skull Splitting Crossbow (13039, -0.23 DPS) [world_drop]; Guttbuster (13139, -0.34 DPS) [world_drop]; Dark Iron Rifle (16004, -2.88 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Knight's Plate Hauberk; wrist: Deepfury Bracers; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Molten Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1011, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 267.6. Weights run: 3.3s. Verify run: 11.4s. 2022 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.663 ± 0.121, crit=2.375 ± 0.173 per rating point (14 rating = 1%, 33.251 per %), hit=3.941 ± 0.305 per rating point (10 rating = 1%, 39.415 per %), melee_haste=17.795 ± 2.708

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 181.3 attack_power points (17.19 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -6.32 DPS) [vendor]; Stalwart Helm (250599, -11.05 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (3.00 DPS) | yes | Sunprism Pendant (20649, -1.35 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.61 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.86 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 61.8 attack_power points (5.86 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 44.9 attack_power points (4.26 DPS) | yes | Sergeant Major's Cape (16337, -1.48 DPS) [pvp]; Fel Cape (279269, -1.55 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.90 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.2 attack_power points (12.72 DPS) | yes | Timbermaw Tunic (252484, -0.47 DPS) [crafted]; Field Marshal's Plate Armor (231536, -4.33 DPS) [pvp]; Knight-Captain's Plate Hauberk (227046, -5.59 DPS) [pvp] |
| wrist | Arena Wristguards (18709) | Arena Treasure Chest [world] | 33.3 attack_power points (3.15 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.03 DPS) [world_drop]; Tranquil Wristguards (279256, -0.12 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (267.6 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.76 DPS) [world_drop]; Timbermaw Brawlers (19049, -8.70 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 65.3 attack_power points (6.19 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -1.64 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.93 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (267.6 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -9.61 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.3 attack_power points (4.86 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.66 DPS) [crafted]; Shadowcraft Boots (16711, -1.91 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 23.3 attack_power points (2.21 DPS) | yes | Falcon's Hook (7552, -0.22 DPS) [dungeon]; Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Ring of Living Stone (18400, -0.31 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.0 attack_power points (1.99 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Ring of the Underwood (2951, -0.03 DPS) [world_drop]; Ring of Living Stone (18400, -0.09 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (267.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (267.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (267.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -2.91 DPS, sim-verified) [world_drop] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 716.2 attack_power points (67.90 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Persuader (22384, -3.00 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -22.18 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (267.6 DPS) | yes | Skull Splitting Crossbow (13039, -0.31 DPS) [world_drop]; Guttbuster (13139, -0.38 DPS) [world_drop]; Dark Iron Rifle (16004, -3.65 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Arena Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 2022, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 764.5. Weights run: 3.6s. Verify run: 12.2s. 2022 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.881 ± 0.145, crit=2.517 ± 0.203 per rating point (14 rating = 1%, 35.245 per %), hit=4.521 ± 0.359 per rating point (10 rating = 1%, 45.210 per %), melee_haste=25.375 ± 3.494

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 200.5 attack_power points (40.02 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -14.74 DPS) [vendor]; Stalwart Helm (250599, -16.77 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.4 attack_power points (7.08 DPS) | yes | Sunprism Pendant (20649, -2.57 DPS) [quest]; Kaleidoscope Chain (13084, -3.82 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.68 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 69.9 attack_power points (13.95 DPS) | yes | Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -8.46 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 48.2 attack_power points (9.62 DPS) | yes | Sergeant Major's Cape (16337, -3.11 DPS) [pvp]; Dark Phantom Cape (13122, -3.99 DPS) [world_drop]; Fel Cape (279269, -5.79 DPS, sim-verified) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 150.0 attack_power points (29.95 DPS) | yes | Timbermaw Tunic (252484, -1.67 DPS) [crafted]; Field Marshal's Plate Armor (231536, -10.63 DPS) [pvp]; Shadowcraft Tunic (16721, -12.97 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.0 attack_power points (7.39 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.30 DPS) [crafted]; Arena Wristguards (18709, -0.35 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (764.5 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -1.60 DPS) [world_drop]; Timbermaw Brawlers (19049, -26.65 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 72.5 attack_power points (14.48 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -3.88 DPS) [pvp]; Prowler's Leather Waistguard (252473, -5.62 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (764.5 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -25.31 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 53.1 attack_power points (10.59 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.75 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.3 attack_power points (5.26 DPS) | yes | Thunderbrow Ring (13097, -0.62 DPS) [world_drop]; Ring of the Underwood (2951, -0.62 DPS) [world_drop]; Ring of Living Stone (18400, -0.86 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.5 attack_power points (4.70 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Ring of Living Stone (18400, -0.30 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (764.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (764.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (764.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -5.68 DPS, sim-verified) [world_drop] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 723.7 attack_power points (144.46 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Persuader (22384, -4.78 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -48.21 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (764.5 DPS) | yes | Guttbuster (13139, -0.94 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.14 DPS) [world_drop]; Dark Iron Rifle (16004, -5.03 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Falcon's Hook; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 2022, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 37.1. Weights run: 2.5s. Verify run: 4.1s. 283 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.271 ± 0.028, crit=0.759 ± 0.028 per rating point (14 rating = 1%, 10.620 per %), hit=0.982 ± 0.056 per rating point (10 rating = 1%, 9.821 per %), melee_haste=5.774 ± 0.588

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.20 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.44 DPS) [crafted]; Brawler's Leather Hood (252504, -0.59 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.5 attack_power points (0.15 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.4 attack_power points (0.38 DPS) | yes | Rough Bronze Shoulders (3480, -0.02 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.02 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.5 attack_power points (0.51 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.06 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.1 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.47 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.60 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Forest Leather Bracers (3202, -0.22 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.1 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.84 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.08 DPS) | yes | Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Deviate Scale Belt (6468, -0.34 DPS) [crafted]; Cobrahn's Grasp (6460, -3.16 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.1 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.76 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.98 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Defender's Leather Boots (252441, -0.38 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.48 DPS) | yes | Bounty Hunter's Ring (5351, -0.25 DPS) [quest]; The 1 Ring (8350, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.36 DPS) | yes | The 1 Ring (8350, -0.16 DPS) [world]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Bounty Hunter's Ring (5351, -0.28 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.0 attack_power points (14.39 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.1 DPS) | yes | Diamond Hammer (2194, +0.00 DPS) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.1 attack_power points (0.31 DPS) | yes | Fine Longbow (11304, -0.07 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.07 DPS) [crafted]; Deadly Blunderbuss (4369, -0.15 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Butcher's Cleaver; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 64.8. Weights run: 2.7s. Verify run: 2.3s. 541 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.389 ± 0.038, crit=1.048 ± 0.044 per rating point (14 rating = 1%, 14.666 per %), hit=1.363 ± 0.083 per rating point (10 rating = 1%, 13.634 per %), melee_haste=7.721 ± 0.909

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 30.5 attack_power points (1.89 DPS) | yes | Tusken Helm (6686, -0.28 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.40 DPS) [crafted]; Defender's Leather Helm (252455, -0.40 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.6 attack_power points (0.84 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; River Pride Choker (13087, -0.34 DPS) [world_drop]; Erudite's Amulet (277204, -0.67 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 25.3 attack_power points (1.56 DPS) | yes | Barbaric Iron Shoulders (7913, -0.31 DPS) [crafted]; Barbaric Shoulders (5964, -0.52 DPS) [crafted]; Bristlebark Amice (14573, -0.68 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.7 attack_power points (0.97 DPS) | yes | Tigerstrike Mantle (13108, -0.21 DPS, sim-verified) [world_drop]; Wildhunter Cloak (16658, -0.35 DPS) [quest]; Construct Cloak (279848, -0.35 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.86 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.10 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.36 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 16.3 attack_power points (1.01 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Pugilist Bracers (4438, -0.02 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.3 attack_power points (1.38 DPS) | yes | Gauntlets of Ogre Strength (3341, -0.02 DPS) [world]; Mail Combat Gauntlets (4075, -0.05 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 28.3 attack_power points (1.75 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted]; Girdle of Golem Strength (9405, -0.27 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 31.7 attack_power points (1.96 DPS) | yes | Brawler's Leather Legguards (252516, -0.31 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.35 DPS) [dungeon]; Glimmering Mail Legguards (6386, -0.37 DPS) [world_drop] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 attack_power points (1.06 DPS) | yes | Brawler's Leather Boots (252439, -0.01 DPS) [crafted]; Veteran's Boots (250503, -0.10 DPS) [crafted]; Hard Gold Boots (250534, -0.19 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.2 attack_power points (1.25 DPS) | yes | Silverlaine's Family Seal (6321, -0.63 DPS) [dungeon]; Tiger Band (6749, -0.63 DPS) [quest]; Bonebinder's Signet (273032, -0.63 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 13.6 attack_power points (0.84 DPS) | yes | Silverlaine's Family Seal (6321, -0.22 DPS) [dungeon]; Tiger Band (6749, -0.22 DPS) [quest]; Bonebinder's Signet (273032, -0.22 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 330.8 attack_power points (20.46 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (64.8 DPS) | yes | Swinetusk Shank (6691, -23.32 DPS, sim-verified) [dungeon] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.2 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.13 DPS) [world]; Golemsight Long Gun (273029, -0.20 DPS, sim-verified) [dungeon]; Long Battle Bow (15284, -0.26 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 541, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (troll, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 124.5. Weights run: 3.0s. Verify run: 2.9s. 768 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.755 ± 0.070, crit=1.717 ± 0.093 per rating point (14 rating = 1%, 24.039 per %), hit=1.425 ± 0.147 per rating point (10 rating = 1%, 14.248 per %), melee_haste=10.697 ± 1.274

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 45.1 attack_power points (3.88 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Barbaric Iron Helm (7915, -0.97 DPS) [crafted]; Hard Gold Coif (250537, -1.47 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.0 attack_power points (1.47 DPS) | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.43 DPS) [world]; River Pride Choker (13087, -0.78 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.3 attack_power points (2.52 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.58 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.63 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 18.3 attack_power points (1.57 DPS) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 42.1 attack_power points (3.63 DPS) | yes | Kolkar Marauder Chain (6773, -0.16 DPS) [quest]; Carapace of Tuten'kash (10775, -0.70 DPS) [dungeon]; Veteran's Silvered Chain Shirt (250518, -1.00 DPS) [crafted] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.0 attack_power points (1.81 DPS) | yes | Branded Leather Bracers (19508, -0.09 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.21 DPS) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 44.0 attack_power points (3.79 DPS) | yes | Scarlet Gauntlets (10331, -0.21 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.56 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.71 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 33.8 attack_power points (2.91 DPS) | yes | Officer's Belt (250556, -0.28 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted]; Skulker's Leather Belt (252520, -0.52 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.62 DPS) | yes | Symbolic Legplates (14829, -0.13 DPS) [world_drop]; Firemane Leggings (13129, -0.34 DPS) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 38.3 attack_power points (3.30 DPS) | yes | Blackforge Greaves (6423, -0.24 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.43 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 21.8 attack_power points (1.88 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.15 DPS) [dungeon]; Assault Band (13095, -0.15 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 21.5 attack_power points (1.86 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (124.5 DPS) | yes | Ardent Custodian (868, -40.25 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.19 DPS) [quest]; The Silencer (13138, -0.28 DPS) [world_drop]; Bow of Searing Arrows (2825, -0.98 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 768, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 191.1. Weights run: 3.3s. Verify run: 3.9s. 990 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.601 ± 0.088, crit=1.844 ± 0.121 per rating point (14 rating = 1%, 25.814 per %), hit=3.013 ± 0.195 per rating point (10 rating = 1%, 30.130 per %), melee_haste=13.962 ± 1.827

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 45.8 attack_power points (4.27 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.24 DPS) [dungeon]; Sunscale Helmet (14849, -0.39 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.8 attack_power points (2.87 DPS) | yes | Woven Ivy Necklace (19159, -0.41 DPS) [quest]; Ethereal Talisman (4430, -1.34 DPS) [quest]; Kaleidoscope Chain (13084, -1.53 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 38.8 attack_power points (3.62 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.19 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.22 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 24.0 attack_power points (2.24 DPS) | yes | Pridelord Cape (14673, -0.34 DPS) [world_drop]; Blackflame Cape (13109, -0.45 DPS) [world_drop]; Bloodlust Cape (14801, -0.56 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 10 · Stone Guard · Horde [vendor] | 55.8 attack_power points (5.20 DPS) | yes | Warbear Harness (15064, -0.47 DPS) [crafted]; Warforged Chestplate (11195, -0.73 DPS) [quest]; Blazewind Breastplate (11193, -1.21 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.0 attack_power points (2.98 DPS) | yes | Officer's Wristguards (250581, -0.04 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.26 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 46.4 attack_power points (4.32 DPS) | yes | Gloves of Holy Might (867, -0.06 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.37 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -0.56 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 47.2 attack_power points (4.40 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Officer's Waistguard (250571, -0.30 DPS) [crafted]; Belt of the Gladiator (13134, -1.04 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 10 · Stone Guard · Horde [vendor] | 65.8 attack_power points (6.13 DPS) | yes | Serpentskin Leggings (8262, -1.32 DPS) [world_drop]; Stormshroud Pants (15057, -1.51 DPS, sim-verified) [crafted]; Golem Shard Leggings (13074, -2.03 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 43.6 attack_power points (4.06 DPS) | yes | Skulker's Leather Boots (252469, -0.07 DPS) [crafted]; Officer's Sabatons (250561, -0.30 DPS) [crafted]; Officer's Boots (250546, -0.60 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.24 DPS) | yes | Thunderbrow Ring (13097, -0.30 DPS) [world_drop]; Falcon's Hook (7552, -0.33 DPS) [dungeon]; Ring of the Underwood (2951, -0.37 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.4 attack_power points (2.09 DPS) | yes | Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon]; Ring of the Underwood (2951, -0.22 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (191.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (191.1 DPS) | yes | Molten Heart of the Mountain (249470, -2.11 DPS, sim-verified) [crafted] |
| main_hand | Flurry Axe (871) | World drop [world_drop] | sim-verified (191.1 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Hammer of the Northern Wind (810, -1.11 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 535.8 attack_power points (49.92 DPS) | yes | Thorium Cestus (250614, -2.29 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -4.99 DPS) [dungeon]; White Bone Shredder (11863, -7.60 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (191.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.23 DPS) [world_drop]; Guttbuster (13139, -0.34 DPS) [world_drop]; Dark Iron Rifle (16004, -2.34 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Deepfury Bracers; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Flurry Axe; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 990, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 275.3. Weights run: 3.3s. Verify run: 11.5s. 2005 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.663 ± 0.121, crit=2.375 ± 0.173 per rating point (14 rating = 1%, 33.251 per %), hit=3.941 ± 0.305 per rating point (10 rating = 1%, 39.415 per %), melee_haste=17.795 ± 2.708

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 181.3 attack_power points (17.19 DPS) | yes | Champion's Plate Helm (227043, -6.32 DPS) [pvp]; Blood Guard's Plate Helm (220803, -7.84 DPS) [vendor]; Stalwart Helm (250599, -8.50 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (3.00 DPS) | yes | Woven Ivy Necklace (19159, -0.44 DPS) [quest]; Sunprism Pendant (20649, -1.11 DPS) [quest]; Ethereal Talisman (4430, -1.42 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 61.8 attack_power points (5.86 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 44.9 attack_power points (4.26 DPS) | yes | Fel Cape (279269, -1.47 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.90 DPS) [world_drop]; Pridelord Cape (14673, -2.27 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (275.3 DPS) | yes | Dawn Armor (252483, -3.84 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -3.85 DPS) [pvp]; Legionnaire's Plate Hauberk (227049, -5.11 DPS) [pvp] |
| wrist | Arena Wristguards (18709) | Arena Treasure Chest [world] | 33.3 attack_power points (3.15 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.03 DPS) [world_drop]; Tranquil Wristguards (279256, -0.12 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.76 DPS) [world_drop]; Timbermaw Brawlers (19049, -4.98 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 65.3 attack_power points (6.19 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -1.64 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.86 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -5.00 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.3 attack_power points (4.86 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.66 DPS) [crafted]; Shadowcraft Boots (16711, -2.68 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.28 DPS) | yes | Thunderbrow Ring (13097, -0.29 DPS) [world_drop]; Falcon's Hook (7552, -0.29 DPS) [dungeon]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 23.3 attack_power points (2.21 DPS) | yes | Thunderbrow Ring (13097, -0.22 DPS) [world_drop]; Falcon's Hook (7552, -0.22 DPS) [dungeon]; Ring of the Underwood (2951, -0.25 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.69 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Annihilator (12798, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 716.2 attack_power points (67.90 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -22.18 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.31 DPS) [world_drop]; Guttbuster (13139, -0.38 DPS) [world_drop]; Dark Iron Rifle (16004, -2.28 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Arena Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 2005, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 778.4. Weights run: 3.6s. Verify run: 12.4s. 2005 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.881 ± 0.145, crit=2.517 ± 0.203 per rating point (14 rating = 1%, 35.245 per %), hit=4.521 ± 0.359 per rating point (10 rating = 1%, 45.210 per %), melee_haste=25.375 ± 3.494

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 200.5 attack_power points (40.02 DPS) | yes | Champion's Plate Helm (227043, -14.74 DPS) [pvp]; Stalwart Helm (250599, -18.09 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -18.26 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.4 attack_power points (7.08 DPS) | yes | Woven Ivy Necklace (19159, -1.06 DPS) [quest]; Sunprism Pendant (20649, -2.57 DPS) [quest]; Ethereal Talisman (4430, -3.38 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 69.9 attack_power points (13.95 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 48.2 attack_power points (9.62 DPS) | yes | Fel Cape (279269, -2.59 DPS) [crafted]; Dark Phantom Cape (13122, -3.99 DPS) [world_drop]; Pridelord Cape (14673, -4.93 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 150.0 attack_power points (29.95 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -10.63 DPS) [pvp]; Shadowcraft Tunic (16721, -12.97 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.0 attack_power points (7.39 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.30 DPS) [crafted]; Arena Wristguards (18709, -0.35 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (778.4 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -1.60 DPS) [world_drop]; Timbermaw Brawlers (19049, -16.30 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 72.5 attack_power points (14.48 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -3.88 DPS) [pvp]; Prowler's Leather Waistguard (252473, -5.42 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (778.4 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -16.70 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 53.1 attack_power points (10.59 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.75 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.3 attack_power points (5.26 DPS) | yes | Falcon's Hook (7552, -0.56 DPS) [dungeon]; Thunderbrow Ring (13097, -0.62 DPS) [world_drop]; Ring of the Underwood (2951, -0.62 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (4.79 DPS) | yes | Falcon's Hook (7552, -0.09 DPS) [dungeon]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Ring of the Underwood (2951, -0.16 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (778.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (778.4 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (778.4 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 723.7 attack_power points (144.46 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -48.21 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (778.4 DPS) | yes | Guttbuster (13139, -0.94 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.14 DPS) [world_drop]; Dark Iron Rifle (16004, -5.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 2005, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

