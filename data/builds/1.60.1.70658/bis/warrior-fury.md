# Leveling BiS: Fury

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 36.7. Weights run: 1.8s. Verify run: 2.9s. 293 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.271 ± 0.028, crit=0.759 ± 0.028 per rating point (14 rating = 1%, 10.620 per %), hit=0.982 ± 0.056 per rating point (10 rating = 1%, 9.821 per %), melee_haste=5.774 ± 0.588

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.20 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.44 DPS) [crafted]; Brawler's Leather Hood (252504, -0.59 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.5 attack_power points (0.15 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.4 attack_power points (0.38 DPS) | yes | Rough Bronze Shoulders (3480, -0.02 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.02 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.5 attack_power points (0.51 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.06 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.7 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.64 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.60 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.06 DPS) [quest]; Cryptwalker Bracers (280095, -0.12 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.7 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.99 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.08 DPS) | yes | Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Deviate Scale Belt (6468, -0.34 DPS) [crafted]; Cobrahn's Grasp (6460, -3.21 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.7 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.90 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.98 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Defender's Leather Boots (252441, -0.38 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.48 DPS) | yes | Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.3 attack_power points (0.20 DPS) | yes | Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Ring of the Moon (12052, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (14.98 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Living Root (6631, +0.00 DPS) [dungeon]; Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.1 attack_power points (14.39 DPS) | yes | Diamond Hammer (2194, -0.31 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.1 attack_power points (0.31 DPS) | yes | Dwarven Fishing Pole (3567, -0.07 DPS) [quest]; Fine Longbow (11304, -0.07 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.07 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: The 1 Ring; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers

### Band 30 (human, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 64.4. Weights run: 1.9s. Verify run: 1.7s. 551 eligible items had no known source.

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
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (21.36 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (64.4 DPS) | yes | Shoni's Disarming Tool (9608, -5.79 DPS) [quest]; Swinetusk Shank (6691, -22.93 DPS, sim-verified) [dungeon] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.2 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.13 DPS) [world]; Golemsight Long Gun (273029, -0.22 DPS, sim-verified) [dungeon]; Long Battle Bow (15284, -0.26 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Seal of Wrynn; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 551, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 40 (human, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 124.2. Weights run: 2.0s. Verify run: 1.9s. 781 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.650 ± 0.067, crit=1.518 ± 0.087 per rating point (14 rating = 1%, 21.253 per %), hit=1.823 ± 0.144 per rating point (10 rating = 1%, 18.227 per %), melee_haste=9.276 ± 1.167

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.8 attack_power points (3.82 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Barbaric Iron Helm (7915, -0.96 DPS) [crafted]; Hard Gold Coif (250537, -1.38 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.6 attack_power points (1.27 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.23 DPS) [world]; River Pride Choker (13087, -0.58 DPS) [world_drop]; Gazlowe's Charm (13088, -0.58 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.2 attack_power points (2.46 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.54 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.55 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 21.9 attack_power points (1.91 DPS) | yes | Hawkeye's Cloak (14593, -0.38 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.52 DPS) [quest]; Parachute Cloak (10518, -0.76 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 39.6 attack_power points (3.46 DPS) | yes | Kolkar Marauder Chain (6773, -0.00 DPS) [quest]; Carapace of Tuten'kash (10775, -0.56 DPS) [dungeon]; Jouster's Chestplate (8157, -0.84 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (1.80 DPS) | yes | Branded Leather Bracers (19508, -0.05 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.24 DPS) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.3 attack_power points (3.60 DPS) | yes | Scarlet Gauntlets (10331, -0.07 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.41 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.56 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.9 attack_power points (2.87 DPS) | yes | Officer's Belt (250556, -0.26 DPS) [crafted]; Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.52 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.67 DPS) | yes | Symbolic Legplates (14829, -0.18 DPS) [world_drop]; Firemane Leggings (13129, -0.35 DPS) [world_drop]; Triprunner Dungarees (9624, -0.55 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.6 attack_power points (3.28 DPS) | yes | Blackforge Greaves (6423, -0.27 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.35 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.47 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.0 attack_power points (1.83 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 20.9 attack_power points (1.82 DPS) | yes | Ring of the Underwood (2951, -0.03 DPS) [world_drop]; Mark of Kern (2262, -0.07 DPS) [dungeon]; Assault Band (13095, -0.07 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (124.2 DPS) | yes | Shoni's Disarming Tool (9608, -18.97 DPS) [quest]; Ardent Custodian (868, -39.94 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.21 DPS) [quest]; The Silencer (13138, -0.26 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.12 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 781, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (human, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 172.2. Weights run: 2.2s. Verify run: 2.5s. 1004 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.508 ± 0.067, crit=1.745 ± 0.091 per rating point (14 rating = 1%, 24.436 per %), hit=2.124 ± 0.147 per rating point (10 rating = 1%, 21.242 per %), melee_haste=12.295 ± 1.335

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 44.4 attack_power points (4.44 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.23 DPS) [dungeon]; Sunscale Helmet (14849, -0.34 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.6 attack_power points (2.96 DPS) | yes | Kaleidoscope Chain (13084, -1.56 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.76 DPS) [world]; River Pride Choker (13087, -2.16 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 38.1 attack_power points (3.80 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.20 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.25 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 22.6 attack_power points (2.26 DPS) | yes | Sergeant Major's Cape (16336, -0.16 DPS) [pvp]; Pridelord Cape (14673, -0.30 DPS) [world_drop]; Blackflame Cape (13109, -0.45 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 54.4 attack_power points (5.44 DPS) | yes | Warbear Harness (15064, -0.53 DPS) [crafted]; Warforged Chestplate (11195, -0.64 DPS) [quest]; Blazewind Breastplate (11193, -1.37 DPS) [quest] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 31.0 attack_power points (3.10 DPS) | yes | Deepfury Bracers (13120, -0.04 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.25 DPS) [crafted]; Arena Bands (18711, -0.30 DPS) [world] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 45.6 attack_power points (4.55 DPS) | yes | Gloves of Holy Might (867, -0.11 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.40 DPS) [crafted]; Stonerender Gauntlets (17007, -0.56 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 46.1 attack_power points (4.60 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Officer's Waistguard (250571, -0.30 DPS) [crafted]; Belt of the Gladiator (13134, -1.01 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 63.5 attack_power points (6.34 DPS) | yes | Gryphon Rider's Leggings (9652, -1.09 DPS) [quest]; Centurion Legplates (10740, -1.09 DPS) [quest]; Serpentskin Leggings (8262, -1.34 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 42.6 attack_power points (4.25 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS) [crafted]; Officer's Sabatons (250561, -0.30 DPS) [crafted]; Officer's Boots (250546, -0.60 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.1 attack_power points (2.11 DPS) | yes | Mark of Kern (2262, -0.11 DPS) [dungeon]; Assault Band (13095, -0.11 DPS) [world_drop]; Blackstone Ring (17713, -0.11 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.5 attack_power points (2.05 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop]; Blackstone Ring (17713, -0.05 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.2 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (172.2 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (172.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Flurry Axe (871, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 534.5 attack_power points (53.38 DPS) | yes | Thorium Cestus (250614, -2.51 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.22 DPS) [dungeon]; Shoni's Disarming Tool (9608, -30.94 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (172.2 DPS) | yes | Skull Splitting Crossbow (13039, -0.11 DPS) [world_drop]; Guttbuster (13139, -0.30 DPS) [world_drop]; Dark Iron Rifle (16004, -1.33 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1004, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 247.7. Weights run: 2.1s. Verify run: 7.9s. 2015 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.791 ± 0.086, crit=2.559 ± 0.123 per rating point (14 rating = 1%, 35.825 per %), hit=3.086 ± 0.244 per rating point (10 rating = 1%, 30.859 per %), melee_haste=19.160 ± 1.920

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 169.4 attack_power points (17.12 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -6.13 DPS) [vendor]; Outlaw's Collar (279253, -9.74 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.3 attack_power points (3.36 DPS) | yes | Sunprism Pendant (20649, -1.19 DPS) [quest]; Kaleidoscope Chain (13084, -1.83 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.15 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 54.7 attack_power points (5.53 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.9 attack_power points (4.74 DPS) | yes | Fel Cape (279269, -1.62 DPS, sim-verified) [crafted]; Sergeant Major's Cape (16337, -1.67 DPS) [pvp]; Dark Phantom Cape (13122, -2.02 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 134.1 attack_power points (13.55 DPS) | yes | Dawn Armor (252483, -1.64 DPS, sim-verified) [crafted]; Field Marshal's Plate Armor (231536, -4.16 DPS) [pvp]; Knight-Captain's Plate Hauberk (227046, -5.68 DPS) [pvp] |
| wrist | Arena Wristguards (18709) | Arena Treasure Chest [world] | 35.8 attack_power points (3.62 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.31 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (247.7 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.81 DPS) [world_drop]; Timbermaw Brawlers (19049, -6.33 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 67.1 attack_power points (6.78 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.63 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -1.82 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (247.7 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -6.34 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 52.3 attack_power points (5.29 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.55 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.67 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.1 attack_power points (2.53 DPS) | yes | Ring of the Underwood (2951, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.37 DPS) [world_drop]; Ring of Living Stone (18400, -0.51 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.1 attack_power points (2.24 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Ring of Living Stone (18400, -0.21 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (247.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (247.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (247.7 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 708.6 attack_power points (71.62 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -22.89 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (247.7 DPS) | yes | Guttbuster (13139, -0.48 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.51 DPS) [world_drop]; Dark Iron Rifle (16004, -2.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Arena Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Falcon's Hook; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 2015, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (human, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 755.4. Weights run: 2.3s. Verify run: 8.1s. 2015 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.001 ± 0.116, crit=2.645 ± 0.162 per rating point (14 rating = 1%, 37.029 per %), hit=3.874 ± 0.304 per rating point (10 rating = 1%, 38.740 per %), melee_haste=21.492 ± 2.914

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 191.1 attack_power points (40.43 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -14.63 DPS) [vendor]; Stalwart Helm (250599, -15.84 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.0 attack_power points (7.83 DPS) | yes | Sunprism Pendant (20649, -3.76 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -4.27 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -5.29 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 64.7 attack_power points (13.68 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.0 attack_power points (10.58 DPS) | yes | Fel Cape (279269, -2.75 DPS) [crafted]; Sergeant Major's Cape (16337, -3.47 DPS) [pvp]; Dark Phantom Cape (13122, -4.23 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 149.8 attack_power points (31.68 DPS) | yes | Timbermaw Tunic (252484, -0.88 DPS) [crafted]; Field Marshal's Plate Armor (231536, -10.48 DPS) [pvp]; Knight-Captain's Plate Hauberk (227046, -14.08 DPS) [pvp] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 38.8 attack_power points (8.21 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.38 DPS) [world]; Officer's Wristguards (250581, -0.55 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (755.4 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -1.69 DPS) [world_drop]; Timbermaw Brawlers (19049, -19.42 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 74.2 attack_power points (15.70 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, -4.19 DPS) [crafted]; Prowler's Leather Waistguard (252473, -5.14 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (755.4 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -18.78 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.0 attack_power points (11.42 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.72 DPS) [crafted]; Shadowcraft Boots (16711, -3.75 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.0 attack_power points (5.92 DPS) | yes | Ring of the Underwood (2951, -0.76 DPS) [world_drop]; Thunderbrow Ring (13097, -0.93 DPS) [world_drop]; Ring of Living Stone (18400, -1.27 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.6 attack_power points (5.20 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Ring of Living Stone (18400, -0.55 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (755.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (755.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (755.4 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -6.91 DPS, sim-verified) [world_drop] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 718.2 attack_power points (151.91 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -49.93 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (755.4 DPS) | yes | Guttbuster (13139, -1.14 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.57 DPS) [world_drop]; Dark Iron Rifle (16004, -3.64 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Deepfury Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Falcon's Hook; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 2015, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 00000000000000000-05150000000000000-000000000000000000)

Set DPS (verified): 36.7. Weights run: 1.8s. Verify run: 2.9s. 276 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.271 ± 0.028, crit=0.759 ± 0.028 per rating point (14 rating = 1%, 10.620 per %), hit=0.982 ± 0.056 per rating point (10 rating = 1%, 9.821 per %), melee_haste=5.774 ± 0.588

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (1.20 DPS) | yes | Defender's Leather Hood (252447, -0.25 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.44 DPS) [crafted]; Brawler's Leather Hood (252504, -0.59 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.5 attack_power points (0.15 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.4 attack_power points (0.38 DPS) | yes | Rough Bronze Shoulders (3480, -0.02 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.02 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.5 attack_power points (0.51 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.06 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.7 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -2.31 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.60 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Forest Leather Bracers (3202, -0.22 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.7 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -2.62 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.08 DPS) | yes | Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Deviate Scale Belt (6468, -0.34 DPS) [crafted]; Cobrahn's Grasp (6460, -2.97 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.7 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -2.60 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.98 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.08 DPS) [crafted]; Defender's Leather Boots (252441, -0.38 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.48 DPS) | yes | Bounty Hunter's Ring (5351, -0.25 DPS) [quest]; The 1 Ring (8350, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.36 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS) [quest]; The 1 Ring (8350, -0.16 DPS) [world]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (14.98 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Hammerbone (270018, +0.00 DPS) [quest] |
| off_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 239.1 attack_power points (14.39 DPS) | yes | Diamond Hammer (2194, -0.32 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 5.1 attack_power points (0.31 DPS) | yes | Fine Longbow (11304, -0.07 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.07 DPS) [crafted]; Deadly Blunderbuss (4369, -0.15 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Shadowfang; off_hand: Butcher's Cleaver; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 276, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring

### Band 30 (troll, 00000000000000000-05153105010000000-000000000000000000)

Set DPS (verified): 64.8. Weights run: 1.9s. Verify run: 1.7s. 534 eligible items had no known source.

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
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 345.4 attack_power points (21.36 DPS) | yes | Swinetusk Shank (6691, +0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, +0.00 DPS) [dungeon]; Viscous Hammer (13045, +0.00 DPS) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (64.8 DPS) | yes | Swinetusk Shank (6691, -23.32 DPS, sim-verified) [dungeon] |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.2 attack_power points (0.63 DPS) | yes | Gravespike Repeater (273033, -0.13 DPS) [world]; Golemsight Long Gun (273029, -0.20 DPS, sim-verified) [dungeon]; Long Battle Bow (15284, -0.26 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Double-barreled Shotgun

No-known-source sample (15 of 534, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (troll, 00000000000000000-05153105022011401-000000000000000000)

Set DPS (verified): 123.6. Weights run: 2.0s. Verify run: 1.9s. 761 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.650 ± 0.067, crit=1.518 ± 0.087 per rating point (14 rating = 1%, 21.253 per %), hit=1.823 ± 0.144 per rating point (10 rating = 1%, 18.227 per %), melee_haste=9.276 ± 1.167

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.8 attack_power points (3.82 DPS) | yes | White Bandit Mask (10008, -0.32 DPS) [crafted]; Barbaric Iron Helm (7915, -0.96 DPS) [crafted]; Hard Gold Coif (250537, -1.38 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 16.6 attack_power points (1.45 DPS) | yes | Kaleidoscope Chain (13084, -0.17 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.40 DPS) [world]; River Pride Choker (13087, -0.75 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.2 attack_power points (2.46 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.54 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.55 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.6 attack_power points (1.53 DPS) | yes | Imperial Cloak (6432, -0.38 DPS) [dungeon]; Parachute Cloak (10518, -0.38 DPS) [crafted]; Tigerstrike Mantle (13108, -0.38 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 39.6 attack_power points (3.46 DPS) | yes | Kolkar Marauder Chain (6773, -0.00 DPS) [quest]; Carapace of Tuten'kash (10775, -0.56 DPS) [dungeon]; Jouster's Chestplate (8157, -0.84 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (1.80 DPS) | yes | Branded Leather Bracers (19508, -0.05 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.24 DPS) [world_drop]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 41.3 attack_power points (3.60 DPS) | yes | Scarlet Gauntlets (10331, -0.07 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.41 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.56 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.9 attack_power points (2.87 DPS) | yes | Officer's Belt (250556, -0.26 DPS) [crafted]; Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.52 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.67 DPS) | yes | Symbolic Legplates (14829, -0.18 DPS) [world_drop]; Firemane Leggings (13129, -0.35 DPS) [world_drop]; Triprunner Dungarees (9624, -0.55 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.6 attack_power points (3.28 DPS) | yes | Blackforge Greaves (6423, -0.27 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.35 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.47 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.0 attack_power points (1.83 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.08 DPS) [dungeon]; Assault Band (13095, -0.08 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 20.9 attack_power points (1.82 DPS) | yes | Ring of the Underwood (2951, -0.03 DPS) [world_drop]; Mark of Kern (2262, -0.07 DPS) [dungeon]; Assault Band (13095, -0.07 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (123.6 DPS) | yes | Ardent Custodian (868, -40.54 DPS, sim-verified) [world_drop] |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.21 DPS) [quest]; The Silencer (13138, -0.26 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.16 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Monolithic Bow

No-known-source sample (15 of 761, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 34200000000000000-05153105022011501-000000000000000000)

Set DPS (verified): 177.6. Weights run: 2.2s. Verify run: 2.5s. 983 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.508 ± 0.067, crit=1.745 ± 0.091 per rating point (14 rating = 1%, 24.436 per %), hit=2.124 ± 0.147 per rating point (10 rating = 1%, 21.242 per %), melee_haste=12.295 ± 1.335

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 44.4 attack_power points (4.44 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.23 DPS) [dungeon]; Sunscale Helmet (14849, -0.34 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.6 attack_power points (2.96 DPS) | yes | Woven Ivy Necklace (19159, -0.40 DPS) [quest]; Ethereal Talisman (4430, -1.36 DPS) [quest]; Kaleidoscope Chain (13084, -1.56 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 38.1 attack_power points (3.80 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.20 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.25 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 22.6 attack_power points (2.26 DPS) | yes | Pridelord Cape (14673, -0.30 DPS) [world_drop]; Blackflame Cape (13109, -0.45 DPS) [world_drop]; Bloodlust Cape (14801, -0.46 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 10 · Stone Guard · Horde [vendor] | 54.4 attack_power points (5.44 DPS) | yes | Warbear Harness (15064, -0.53 DPS) [crafted]; Warforged Chestplate (11195, -0.64 DPS) [quest]; Blazewind Breastplate (11193, -1.37 DPS) [quest] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 31.0 attack_power points (3.10 DPS) | yes | Deepfury Bracers (13120, -0.04 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.25 DPS) [crafted]; Arena Bands (18711, -0.30 DPS) [world] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 45.6 attack_power points (4.55 DPS) | yes | Gloves of Holy Might (867, -0.11 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.40 DPS) [crafted]; Stonerender Gauntlets (17007, -0.56 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 46.1 attack_power points (4.60 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Officer's Waistguard (250571, -0.30 DPS) [crafted]; Belt of the Gladiator (13134, -1.01 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 10 · Stone Guard · Horde [vendor] | 63.5 attack_power points (6.34 DPS) | yes | Serpentskin Leggings (8262, -1.33 DPS, sim-verified) [world_drop]; Stormshroud Pants (15057, -1.46 DPS) [crafted]; Golem Shard Leggings (13074, -1.95 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 42.6 attack_power points (4.25 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS) [crafted]; Officer's Sabatons (250561, -0.30 DPS) [crafted]; Officer's Boots (250546, -0.60 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.40 DPS) | yes | Thunderbrow Ring (13097, -0.35 DPS) [world_drop]; Mark of Kern (2262, -0.40 DPS) [dungeon]; Blackstone Ring (17713, -0.40 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.1 attack_power points (2.11 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.11 DPS) [dungeon]; Blackstone Ring (17713, -0.11 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (177.6 DPS) | yes | Molten Heart of the Mountain (249470, -4.42 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (177.6 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (177.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Flurry Axe (871, +0.00 DPS, sim-verified) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 534.5 attack_power points (53.38 DPS) | yes | Thorium Cestus (250614, -3.15 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.22 DPS) [dungeon]; White Bone Shredder (11863, -8.08 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (177.6 DPS) | yes | Skull Splitting Crossbow (13039, -0.11 DPS) [world_drop]; Guttbuster (13139, -0.30 DPS) [world_drop]; Dark Iron Rifle (16004, -2.29 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 983, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 255.8. Weights run: 2.1s. Verify run: 7.7s. 1998 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.791 ± 0.086, crit=2.559 ± 0.123 per rating point (14 rating = 1%, 35.825 per %), hit=3.086 ± 0.244 per rating point (10 rating = 1%, 30.859 per %), melee_haste=19.160 ± 1.920

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 169.4 attack_power points (17.12 DPS) | yes | Champion's Plate Helm (227043, -6.13 DPS) [pvp]; Blood Guard's Plate Helm (220803, -7.75 DPS) [vendor]; Outlaw's Collar (279253, -9.63 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.3 attack_power points (3.36 DPS) | yes | Sunprism Pendant (20649, -1.19 DPS) [quest]; Woven Ivy Necklace (19159, -1.26 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -1.63 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 54.7 attack_power points (5.53 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.9 attack_power points (4.74 DPS) | yes | Dark Phantom Cape (13122, -2.02 DPS) [world_drop]; Pridelord Cape (14673, -2.50 DPS) [world_drop]; Fel Cape (279269, -2.92 DPS, sim-verified) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 134.1 attack_power points (13.55 DPS) | yes | Dawn Armor (252483, -3.92 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -4.16 DPS) [pvp]; Legionnaire's Plate Hauberk (227049, -5.68 DPS) [pvp] |
| wrist | Arena Wristguards (18709) | Arena Treasure Chest [world] | 35.8 attack_power points (3.62 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.10 DPS) [world_drop]; Officer's Wristguards (250581, -0.31 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (255.8 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.81 DPS) [world_drop]; Timbermaw Brawlers (19049, -5.92 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 67.1 attack_power points (6.78 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.82 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.58 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (255.8 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -6.78 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 52.3 attack_power points (5.29 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.67 DPS) [crafted]; Shadowcraft Boots (16711, -2.48 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.1 attack_power points (2.53 DPS) | yes | Falcon's Hook (7552, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.37 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.43 DPS) | yes | Ring of the Underwood (2951, -0.21 DPS) [world_drop]; Thunderbrow Ring (13097, -0.27 DPS) [world_drop]; Falcon's Hook (7552, -2.09 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (255.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (255.8 DPS) | yes | Molten Heart of the Mountain (249470, -2.55 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (255.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -1.48 DPS, sim-verified) [world_drop] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 708.6 attack_power points (71.62 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Persuader (22384, -1.72 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -22.89 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (255.8 DPS) | yes | Guttbuster (13139, -0.48 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.51 DPS) [world_drop]; Dark Iron Rifle (16004, -3.65 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Arena Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: White Bone Band; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 1998, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 34320003002000000-05153105022011501-200000000000000000)

Set DPS (verified): 773.2. Weights run: 2.3s. Verify run: 8.5s. 1998 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.001 ± 0.116, crit=2.645 ± 0.162 per rating point (14 rating = 1%, 37.029 per %), hit=3.874 ± 0.304 per rating point (10 rating = 1%, 38.740 per %), melee_haste=21.492 ± 2.914

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 191.1 attack_power points (40.43 DPS) | yes | Champion's Plate Helm (227043, -14.63 DPS) [pvp]; Blood Guard's Plate Helm (220803, -18.35 DPS) [vendor]; Stalwart Helm (250599, -18.43 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.0 attack_power points (7.83 DPS) | yes | Woven Ivy Necklace (19159, -1.23 DPS) [quest]; Sunprism Pendant (20649, -2.75 DPS) [quest]; Ethereal Talisman (4430, -3.81 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 64.7 attack_power points (13.68 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -7.45 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.0 attack_power points (10.58 DPS) | yes | Fel Cape (279269, -3.98 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -4.23 DPS) [world_drop]; Pridelord Cape (14673, -5.37 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (773.2 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Warlord's Plate Armor (231530, -9.60 DPS) [pvp]; Legionnaire's Plate Hauberk (227049, -13.20 DPS) [pvp] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 38.8 attack_power points (8.21 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Arena Wristguards (18709, -0.38 DPS) [world]; Officer's Wristguards (250581, -0.55 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (773.2 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -1.69 DPS) [world_drop]; Timbermaw Brawlers (19049, -18.49 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 74.2 attack_power points (15.70 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Prowler's Leather Waistguard (252473, -3.88 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -4.19 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (773.2 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -17.82 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 54.0 attack_power points (11.42 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.08 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.72 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.0 attack_power points (5.92 DPS) | yes | Ring of the Underwood (2951, -0.76 DPS) [world_drop]; White Bone Band (11862, -0.85 DPS) [quest]; Thunderbrow Ring (13097, -0.93 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 24.6 attack_power points (5.20 DPS) | yes | Ring of the Underwood (2951, -0.04 DPS) [world_drop]; White Bone Band (11862, -0.13 DPS) [quest]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (773.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (773.2 DPS) | yes | Molten Heart of the Mountain (249470, -10.96 DPS, sim-verified) [crafted] |
| main_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | sim-verified (773.2 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | Greenhammer (279261) | Blacksmithing [crafted] | 718.2 attack_power points (151.91 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -49.93 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (773.2 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Guttbuster (13139, -1.14 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.57 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Deepfury Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; main_hand: Ironwood Blade; off_hand: Greenhammer

No-known-source sample (15 of 1998, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

