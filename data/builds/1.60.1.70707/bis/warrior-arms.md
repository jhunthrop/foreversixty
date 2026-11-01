# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 25.5. Weights run: 1.8s. Verify run: 3.6s. 339 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.209 ± 0.030, crit=0.828 ± 0.027 per rating point (14 rating = 1%, 11.592 per %), hit=1.232 ± 0.083 per rating point (10 rating = 1%, 12.321 per %), melee_haste=8.730 ± 0.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.28 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.4 attack_power points (0.09 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS) [quest]; Tarnished Locket (279870, -0.30 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.00 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Dark Leather Cloak (2316, -0.01 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.03 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.07 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (25.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -1.01 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.01 DPS) [quest]; Bristlebark Bindings (14569, -0.08 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.18 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Cobrahn's Grasp (6460, -1.36 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.15 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.2 attack_power points (0.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; Ring of the Moon (12052, -0.04 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 302.8 attack_power points (10.71 DPS) | yes | Duskbringer (2205, -0.35 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.46 DPS) [dungeon]; Sword of the Fallen (286977, -0.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.8 attack_power points (0.17 DPS) | yes | Dwarven Fishing Pole (3567, -0.03 DPS) [quest]; Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 339, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 73.3. Weights run: 1.9s. Verify run: 1.8s. 597 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.258 ± 0.064, crit=1.785 ± 0.070 per rating point (14 rating = 1%, 24.983 per %), hit=2.019 ± 0.197 per rating point (10 rating = 1%, 20.187 per %), melee_haste=10.573 ± 1.434

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 38.3 attack_power points (1.47 DPS) | yes | Tusken Helm (6686, -0.47 DPS) [dungeon]; Slayer's Skullcap (14753, -0.52 DPS) [world_drop]; Guard's Chain Helm (250499, -1.02 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.0 attack_power points (0.66 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.42 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.8 attack_power points (1.34 DPS) | yes | Barbaric Iron Shoulders (7913, -0.36 DPS) [crafted]; Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Shoulders (5964, -0.52 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.8 attack_power points (0.84 DPS) | yes | Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.32 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.33 DPS, sim-verified) [world_drop] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Brawler's Leather Tunic (252508, -0.13 DPS) [crafted]; Avenger's Armor (1488, -0.14 DPS) [dungeon]; Dusky Leather Armor (7374, -0.78 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.6 attack_power points (0.83 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Barbaric Bracers (18948, -0.17 DPS) [crafted]; Pugilist Bracers (4438, -0.21 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.6 attack_power points (1.06 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.10 DPS) [world_drop]; Fletcher's Gloves (7348, -0.10 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Warden's Leather Belt (252460, -0.31 DPS) [crafted]; Skulker's Leather Belt (252520, -0.33 DPS, sim-verified) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 37.8 attack_power points (1.45 DPS) | yes | Brawler's Leather Pants (252500, -0.21 DPS) [crafted]; Trapper's Leather Pants (252501, -0.21 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.37 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.1 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Alacritous Treads (277234, -0.11 DPS) [quest]; Veteran's Boots (250503, -0.19 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.8 attack_power points (0.88 DPS) | yes | Seal of Wrynn (2933, -0.38 DPS) [quest]; Monkey Ring (6748, -0.44 DPS) [quest]; Tiger Band (6749, -0.49 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 13.6 attack_power points (0.52 DPS) | yes | Seal of Wrynn (2933, -0.03 DPS) [quest]; Monkey Ring (6748, -0.09 DPS) [quest]; Tiger Band (6749, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (73.3 DPS) | yes | Felbough Arm (273037, -0.11 DPS) [world]; Corpsemaker (6687, -0.54 DPS) [dungeon]; Viscous Hammer (13045, -20.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 13.6 attack_power points (0.52 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.17 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Manual Crowd Pummeler; ranged: Golemsight Long Gun

No-known-source sample (15 of 597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 96.5. Weights run: 2.1s. Verify run: 2.7s. 826 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.862 ± 0.063, crit=1.809 ± 0.078 per rating point (14 rating = 1%, 25.325 per %), hit=1.807 ± 0.208 per rating point (10 rating = 1%, 18.071 per %), melee_haste=13.533 ± 1.711

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 51.3 attack_power points (2.64 DPS) | yes | Chromite Barbute (8142, -0.26 DPS) [dungeon]; White Bandit Mask (10008, -0.45 DPS) [crafted]; Barbaric Iron Helm (7915, -0.85 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 attack_power points (0.79 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.18 DPS) [world]; River Pride Choker (13087, -0.38 DPS) [world_drop]; Gazlowe's Charm (13088, -0.38 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.5 attack_power points (1.57 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.38 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.44 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 23.2 attack_power points (1.19 DPS) | yes | Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.31 DPS) [quest]; Parachute Cloak (10518, -0.43 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 62.8 attack_power points (3.22 DPS) | yes | Kolkar Marauder Chain (6773, -1.12 DPS) [quest]; Carapace of Tuten'kash (10775, -1.43 DPS) [dungeon]; Wolffear Harness (13110, -1.60 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.4 attack_power points (1.10 DPS) | yes | Branded Leather Bracers (19508, -0.07 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.12 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 42.6 attack_power points (2.19 DPS) | yes | Plated Fist of Hakoo (13071, -0.21 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.30 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.31 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.8 attack_power points (1.79 DPS) | yes | Officer's Belt (250556, -0.18 DPS) [crafted]; Boar Champion's Belt (10768, -0.24 DPS) [dungeon]; Prowler's Leather Belt (252459, -0.29 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.16 DPS) | yes | Symbolic Legplates (14829, -0.04 DPS) [world_drop]; Triprunner Dungarees (9624, -0.13 DPS) [quest]; Basilisk Hide Pants (1718, -0.15 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.0 attack_power points (2.01 DPS) | yes | Blackforge Greaves (6423, -0.12 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.8 attack_power points (1.17 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.14 DPS) [dungeon]; Assault Band (13095, -0.14 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (96.5 DPS) | yes | Fiery War Axe (870, -0.62 DPS) [world_drop]; Primitive Fishing Pole (276203, -1.12 DPS) [vendor]; Bonebiter (6830, -1.26 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (96.5 DPS) | yes | Master Hunter's Rifle (17687, -0.11 DPS) [quest]; The Silencer (13138, -0.18 DPS) [world_drop]; Swiftwind (13038, -0.23 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Scarlet Gauntlets; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 826, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 183.5. Weights run: 2.2s. Verify run: 3.2s. 1049 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.211 ± 0.063, crit=1.487 ± 0.085 per rating point (14 rating = 1%, 20.812 per %), hit=2.163 ± 0.204 per rating point (10 rating = 1%, 21.632 per %), melee_haste=12.823 ± 1.747

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.8 attack_power points (3.80 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.49 DPS) [crafted]; Sunscale Helmet (14849, -0.61 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (2.09 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.12 DPS) [world]; River Pride Choker (13087, -1.44 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.48 DPS, sim-verified) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.7 attack_power points (2.90 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.3 attack_power points (1.57 DPS) | yes | Dark Phantom Cape (13122, -0.09 DPS) [world_drop]; Bloodlust Cape (14801, -0.10 DPS) [world_drop]; Pridelord Cape (14673, -0.19 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 50.8 attack_power points (4.13 DPS) | yes | Icemail Jerkin (1981, -0.01 DPS) [world_drop]; Warforged Chestplate (11195, -0.23 DPS) [quest]; Warbear Harness (15064, -0.57 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.3 attack_power points (2.38 DPS) | yes | Bracers of the Stone Princess (17714, -0.10 DPS) [dungeon]; Arena Bands (18711, -0.10 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 42.9 attack_power points (3.48 DPS) | yes | Stonerender Gauntlets (17007, -0.24 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.32 DPS) [crafted]; Alabaster Plate Gauntlets (8314, -0.52 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.5 attack_power points (3.46 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Belt of the Gladiator (13134, -0.53 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 56.9 attack_power points (4.62 DPS) | yes | Gryphon Rider's Leggings (9652, -0.71 DPS) [quest]; Centurion Legplates (10740, -0.71 DPS) [quest]; Serpentskin Leggings (8262, -0.94 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.3 attack_power points (3.19 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.39 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.6 attack_power points (3.38 DPS) | yes | Assault Band (13095, -1.76 DPS) [world_drop]; Thunderbrow Ring (13097, -1.79 DPS) [world_drop]; Masons Fraternity Ring (9533, -2.00 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.62 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.25 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (183.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (183.5 DPS) | yes | Mark of the Chosen (17774, -1.21 DPS, sim-verified) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (183.5 DPS) | yes | Blight (7959, +0.00 DPS) [crafted]; Warmonger (13052, -0.34 DPS) [world_drop]; Bleakwood Hew (12769, -2.62 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (183.5 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.25 DPS) [world_drop]; Guttbuster (13139, -0.35 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 1049, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 246.3. Weights run: 2.1s. Verify run: 8.6s. 2061 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.709 ± 0.089, crit=2.441 ± 0.128 per rating point (14 rating = 1%, 34.172 per %), hit=3.656 ± 0.310 per rating point (10 rating = 1%, 36.565 per %), melee_haste=16.323 ± 2.850

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 177.5 attack_power points (14.75 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.38 DPS) [vendor]; Mask of the Unforgiven (13404, -6.01 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.2 attack_power points (2.68 DPS) | yes | Sunprism Pendant (20649, -1.15 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.44 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.68 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 59.5 attack_power points (4.94 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.6 attack_power points (3.79 DPS) | yes | Fel Cape (279269, -0.95 DPS) [crafted]; Sergeant Major's Cape (16337, -1.33 DPS) [pvp]; Cloak of Flames (3475, -2.97 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.0 attack_power points (11.13 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -0.81 DPS) [dungeon]; Field Marshal's Plate Armor (231536, -3.65 DPS) [pvp] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 48.2 attack_power points (4.00 DPS) | yes | Marshal's Plate Bracers (16481, -0.18 DPS) [pvp]; Deepfury Bracers (13120, -1.21 DPS) [world_drop]; Arena Wristguards (18709, -1.44 DPS, sim-verified) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (246.3 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Officer's Gloves (250551, -1.23 DPS) [crafted]; Timbermaw Brawlers (19049, -4.01 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 65.9 attack_power points (5.48 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.45 DPS) [crafted]; Brigam Girdle (13142, -2.18 DPS, sim-verified) [dungeon] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (246.3 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.00 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.7 attack_power points (4.29 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.57 DPS) [crafted]; Shadowcraft Boots (16711, -2.05 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 56.6 attack_power points (4.70 DPS) | yes | Masons Fraternity Ring (9533, -2.71 DPS) [quest]; Falcon's Hook (7552, -2.92 DPS) [dungeon]; Thunderbrow Ring (13097, -2.95 DPS) [world_drop] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 34.2 attack_power points (2.84 DPS) | yes | Masons Fraternity Ring (9533, -0.85 DPS) [quest]; Falcon's Hook (7552, -1.06 DPS) [dungeon]; Thunderbrow Ring (13097, -1.08 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (246.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (246.3 DPS) | yes | Mark of the Chosen (17774, -1.43 DPS, sim-verified) [quest] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (246.3 DPS) | yes | Arcanite Champion (12790, +0.00 DPS, sim-verified) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (246.3 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.51 DPS, sim-verified) [world_drop]; Skull Splitting Crossbow (13039, -1.80 DPS) [world_drop]; Guttbuster (13139, -1.83 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Vambraces of the Sadist; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Band of the Penitent; main_hand: Forest Defender's Axe; ranged: Blackcrow

No-known-source sample (15 of 2061, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 60, raid preset (human, 02325213032515201-05050020000000000-200000000000000000)

Set DPS (verified): 729.1. Weights run: 1.9s. Verify run: 6.3s. 2061 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.822 ± 0.112, crit=2.314 ± 0.154 per rating point (14 rating = 1%, 32.394 per %), hit=5.572 ± 0.442 per rating point (10 rating = 1%, 55.717 per %), melee_haste=27.243 ± 4.036

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 215.8 attack_power points (39.74 DPS) | yes | Stalwart Helm (250599, -14.36 DPS) [crafted]; Mask of the Unforgiven (13404, -14.46 DPS, sim-verified) [dungeon]; Lieutenant Commander's Plate Helm (23314, -15.01 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.7 attack_power points (6.39 DPS) | yes | Sunprism Pendant (20649, -2.36 DPS) [quest]; Kaleidoscope Chain (13084, -3.42 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.18 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 79.1 attack_power points (14.56 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -1.71 DPS) [pvp] |
| back | Cloak of Flames (3475) | World drop [world_drop] | 55.7 attack_power points (10.26 DPS) | yes | Cape of the Black Baron (13340, -1.54 DPS) [dungeon]; Fel Cape (279269, -4.29 DPS) [crafted]; Sergeant Major's Cape (16337, -4.33 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.5 attack_power points (28.64 DPS) | yes | Timbermaw Tunic (252484, -3.27 DPS) [crafted]; Savage Gladiator Chain (11726, -5.53 DPS) [dungeon]; Icemail Jerkin (1981, -10.33 DPS) [world_drop] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 47.8 attack_power points (8.80 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -2.15 DPS) [world_drop]; Tranquil Wristguards (279256, -2.32 DPS) [crafted] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 71.7 attack_power points (13.20 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -2.94 DPS) [crafted]; Devilsaur Gauntlets (15063, -4.27 DPS, sim-verified) [crafted] |
| waist | Brigam Girdle (13142) | Blackrock Spire: General Drakkisath [dungeon] | 86.5 attack_power points (15.93 DPS) | yes | Marshal's Plate Girdle (16482, -1.78 DPS) [pvp]; Might of the Timbermaw (19044, -2.73 DPS) [crafted]; Knight-Captain's Plate Girdle (16407, -6.22 DPS) [pvp] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | sim-verified (729.1 DPS) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Knight-Captain's Plate Leggings (23301, -0.59 DPS) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 55.7 attack_power points (10.26 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.14 DPS) [dungeon]; Scalegut Treaders (275618, -0.58 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 75.7 attack_power points (13.94 DPS) | yes | Masons Fraternity Ring (9533, -9.24 DPS) [quest]; Thunderbrow Ring (13097, -9.69 DPS) [world_drop]; Falcon's Hook (7552, -9.71 DPS) [dungeon] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 32.4 attack_power points (5.96 DPS) | yes | Masons Fraternity Ring (9533, -1.27 DPS) [quest]; Thunderbrow Ring (13097, -1.72 DPS) [world_drop]; Falcon's Hook (7552, -1.73 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (729.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (729.1 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (729.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -36.77 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (729.1 DPS) | yes | Precisely Calibrated Boomstick (2100, -7.33 DPS, sim-verified) [world_drop]; Guttbuster (13139, -7.37 DPS) [world_drop]; Skull Splitting Crossbow (13039, -7.47 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cloak of Flames; chest: Dawn Armor; wrist: Vambraces of the Sadist; hands: Timbermaw Brawlers; waist: Brigam Girdle; legs: Warbear Woolies; feet: Greaves of Withering Despair; finger2: Band of the Penitent; trinket1: Molten Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Arcanite Champion; ranged: Blackcrow

No-known-source sample (15 of 2061, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 25.7. Weights run: 1.8s. Verify run: 3.7s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.209 ± 0.030, crit=0.828 ± 0.027 per rating point (14 rating = 1%, 11.592 per %), hit=1.232 ± 0.083 per rating point (10 rating = 1%, 12.321 per %), melee_haste=8.730 ± 0.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.28 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.4 attack_power points (0.09 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.00 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.00 DPS) [crafted] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Dark Leather Cloak (2316, -0.01 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.03 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.07 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (25.7 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -1.04 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Forest Leather Bracers (3202, -0.14 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.7 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.21 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Cobrahn's Grasp (6460, -1.39 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.7 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.18 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; The 1 Ring (8350, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.21 DPS) | yes | The 1 Ring (8350, -0.10 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (10.72 DPS) | yes | Smite's Mighty Hammer (7230, -0.01 DPS) [dungeon]; Hammerbone (270018, -0.10 DPS) [quest]; Duskbringer (2205, -0.30 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.8 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Forsaken Greataxe; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 73.0. Weights run: 1.9s. Verify run: 1.8s. 580 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.258 ± 0.064, crit=1.785 ± 0.070 per rating point (14 rating = 1%, 24.983 per %), hit=2.019 ± 0.197 per rating point (10 rating = 1%, 20.187 per %), melee_haste=10.573 ± 1.434

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 38.3 attack_power points (1.47 DPS) | yes | Tusken Helm (6686, -0.47 DPS) [dungeon]; Slayer's Skullcap (14753, -0.52 DPS) [world_drop]; Guard's Chain Helm (250499, -0.70 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.0 attack_power points (0.66 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.8 attack_power points (1.34 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.51 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.52 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.8 attack_power points (0.84 DPS) | yes | Tigerstrike Mantle (13108, -0.14 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.32 DPS) [dungeon]; Swiftrunner Cape (6745, -0.32 DPS) [quest] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Brawler's Leather Tunic (252508, -0.13 DPS) [crafted]; Avenger's Armor (1488, -0.14 DPS) [dungeon]; Dusky Leather Armor (7374, -0.33 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.6 attack_power points (0.83 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Barbaric Bracers (18948, -0.17 DPS) [crafted]; Pugilist Bracers (4438, -0.21 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.6 attack_power points (1.06 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.10 DPS) [world_drop]; Fletcher's Gloves (7348, -0.10 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Skulker's Leather Belt (252520, -0.05 DPS) [crafted]; Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Warden's Leather Belt (252460, -0.31 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 37.8 attack_power points (1.45 DPS) | yes | Brawler's Leather Legguards (252516, -0.13 DPS) [crafted]; Brawler's Leather Pants (252500, -0.21 DPS) [crafted]; Trapper's Leather Pants (252501, -0.21 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.1 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Veteran's Boots (250503, -0.19 DPS) [crafted]; Insignia Boots (4055, -0.23 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.8 attack_power points (0.88 DPS) | yes | Ring of Precision (1491, -0.35 DPS) [dungeon]; Monkey Ring (6748, -0.44 DPS) [quest]; Tiger Band (6749, -0.49 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 17.0 attack_power points (0.66 DPS) | yes | Ring of Precision (1491, -0.13 DPS) [dungeon]; Monkey Ring (6748, -0.22 DPS) [quest]; Tiger Band (6749, -0.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (73.0 DPS) | yes | Felbough Arm (273037, -0.11 DPS) [world]; Corpsemaker (6687, -0.54 DPS) [dungeon]; Viscous Hammer (13045, -20.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 13.6 attack_power points (0.52 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.17 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler; ranged: Golemsight Long Gun

No-known-source sample (15 of 580, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 95.4. Weights run: 2.1s. Verify run: 2.7s. 806 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.862 ± 0.063, crit=1.809 ± 0.078 per rating point (14 rating = 1%, 25.325 per %), hit=1.807 ± 0.208 per rating point (10 rating = 1%, 18.071 per %), melee_haste=13.533 ± 1.711

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 51.3 attack_power points (2.64 DPS) | yes | Chromite Barbute (8142, -0.26 DPS) [dungeon]; White Bandit Mask (10008, -0.45 DPS) [crafted]; Barbaric Iron Helm (7915, -0.85 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.4 attack_power points (0.90 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.28 DPS) [world]; River Pride Choker (13087, -0.49 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.5 attack_power points (1.57 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.38 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.44 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.0 attack_power points (0.98 DPS) | yes | Imperial Cloak (6432, -0.21 DPS) [dungeon]; Parachute Cloak (10518, -0.21 DPS) [crafted]; Tigerstrike Mantle (13108, -0.21 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 62.8 attack_power points (3.22 DPS) | yes | Kolkar Marauder Chain (6773, -1.12 DPS) [quest]; Carapace of Tuten'kash (10775, -1.43 DPS) [dungeon]; Wolffear Harness (13110, -1.60 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.4 attack_power points (1.10 DPS) | yes | Branded Leather Bracers (19508, -0.07 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.12 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | 42.6 attack_power points (2.19 DPS) | yes | Plated Fist of Hakoo (13071, -0.21 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.30 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.31 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.8 attack_power points (1.79 DPS) | yes | Officer's Belt (250556, -0.18 DPS) [crafted]; Boar Champion's Belt (10768, -0.24 DPS) [dungeon]; Prowler's Leather Belt (252459, -0.29 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.16 DPS) | yes | Symbolic Legplates (14829, -0.04 DPS) [world_drop]; Triprunner Dungarees (9624, -0.13 DPS) [quest]; Basilisk Hide Pants (1718, -0.15 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.0 attack_power points (2.01 DPS) | yes | Blackforge Greaves (6423, -0.12 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.8 attack_power points (1.17 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.14 DPS) [dungeon]; Assault Band (13095, -0.14 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (95.4 DPS) | yes | Fiery War Axe (870, -0.62 DPS) [world_drop]; Primitive Fishing Pole (276203, -1.12 DPS) [vendor]; Thornstone Sledgehammer (1722, -2.43 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (95.4 DPS) | yes | Master Hunter's Rifle (17687, -0.11 DPS) [quest]; The Silencer (13138, -0.18 DPS) [world_drop]; Swiftwind (13038, -0.23 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Scarlet Gauntlets; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 806, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 187.1. Weights run: 2.2s. Verify run: 3.2s. 1028 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.211 ± 0.063, crit=1.487 ± 0.085 per rating point (14 rating = 1%, 20.812 per %), hit=2.163 ± 0.204 per rating point (10 rating = 1%, 21.632 per %), melee_haste=12.823 ± 1.747

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.8 attack_power points (3.80 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.49 DPS) [crafted]; Sunscale Helmet (14849, -0.61 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (2.09 DPS) | yes | Woven Ivy Necklace (19159, -0.23 DPS) [quest]; Ethereal Talisman (4430, -0.89 DPS) [quest]; Kaleidoscope Chain (13084, -1.05 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.7 attack_power points (2.90 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.2 attack_power points (1.48 DPS) | yes | Bloodlust Cape (14801, -0.01 DPS) [world_drop]; Pridelord Cape (14673, -0.10 DPS) [world_drop]; Blackflame Cape (13109, -0.30 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 10 · Stone Guard · Horde [vendor] | 50.8 attack_power points (4.13 DPS) | yes | Icemail Jerkin (1981, +0.00 DPS, sim-verified) [world_drop]; Warforged Chestplate (11195, -0.23 DPS) [quest]; Warbear Harness (15064, -0.57 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.3 attack_power points (2.38 DPS) | yes | Bracers of the Stone Princess (17714, -0.10 DPS) [dungeon]; Arena Bands (18711, -0.10 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 42.9 attack_power points (3.48 DPS) | yes | Stonerender Gauntlets (17007, -0.24 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.32 DPS) [crafted]; Alabaster Plate Gauntlets (8314, -0.52 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.5 attack_power points (3.46 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Belt of the Gladiator (13134, -0.53 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 10 · Stone Guard · Horde [vendor] | 56.9 attack_power points (4.62 DPS) | yes | Golem Shard Leggings (13074, -1.05 DPS) [world_drop]; Serpentskin Leggings (8262, -1.17 DPS, sim-verified) [world_drop]; Scarlet Leggings (10330, -1.21 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.3 attack_power points (3.19 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.39 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.6 attack_power points (3.38 DPS) | yes | Mark of Kern (2262, -1.76 DPS) [dungeon]; Assault Band (13095, -1.76 DPS) [world_drop]; Thunderbrow Ring (13097, -1.79 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.95 DPS) | yes | Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.35 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (187.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (187.1 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (187.1 DPS) | yes | Blight (7959, +0.00 DPS, sim-verified) [crafted]; Warmonger (13052, -0.34 DPS) [world_drop]; Bleakwood Hew (12769, -2.62 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (187.1 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.25 DPS) [world_drop]; Guttbuster (13139, -0.35 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 1028, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 248.6. Weights run: 2.1s. Verify run: 6.9s. 2044 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.709 ± 0.089, crit=2.441 ± 0.128 per rating point (14 rating = 1%, 34.172 per %), hit=3.656 ± 0.310 per rating point (10 rating = 1%, 36.565 per %), melee_haste=16.323 ± 2.850

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 177.5 attack_power points (14.75 DPS) | yes | Champion's Plate Helm (227043, -5.38 DPS) [pvp]; Mask of the Unforgiven (13404, -5.83 DPS) [dungeon]; Stalwart Helm (250599, -6.68 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.2 attack_power points (2.68 DPS) | yes | Woven Ivy Necklace (19159, -0.40 DPS) [quest]; Sunprism Pendant (20649, -0.97 DPS) [quest]; Ethereal Talisman (4430, -1.28 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 59.5 attack_power points (4.94 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Blood Guard's Plate Pauldrons (220796, -0.27 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.6 attack_power points (3.79 DPS) | yes | Cloak of Flames (3475, -0.75 DPS) [world_drop]; Fel Cape (279269, -0.95 DPS) [crafted]; Dark Phantom Cape (13122, -1.66 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (248.6 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -0.55 DPS) [dungeon]; Warlord's Plate Armor (231530, -3.39 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | sim-verified (248.6 DPS) | yes | Vambraces of the Sadist (13400, +0.00 DPS) [dungeon]; General's Plate Armguards (16546, +0.00 DPS) [pvp]; Arena Wristguards (18709, +0.00 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (248.6 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Officer's Gloves (250551, -1.23 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 65.9 attack_power points (5.48 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Brigam Girdle (13142, -0.11 DPS) [dungeon]; Prowler's Leather Waistguard (252473, -1.45 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (248.6 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.7 attack_power points (4.29 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.40 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.57 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | sim-verified (248.6 DPS) | yes | Band of the Penitent (13217, +0.00 DPS) [quest]; Blackstone Ring (17713, +0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.22 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (248.6 DPS) | yes | Band of the Penitent (13217, +0.00 DPS) [quest]; Blackstone Ring (17713, +0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.21 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (248.6 DPS) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.33 DPS) [world_drop]; Guttbuster (13139, -0.35 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: White Bone Band; finger2: Masons Fraternity Ring; main_hand: Forest Defender's Axe; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2044, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (orc, 02325213032515201-05050020000000000-200000000000000000)

Set DPS (verified): 747.7. Weights run: 1.9s. Verify run: 6.7s. 2044 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.822 ± 0.112, crit=2.314 ± 0.154 per rating point (14 rating = 1%, 32.394 per %), hit=5.572 ± 0.442 per rating point (10 rating = 1%, 55.717 per %), melee_haste=27.243 ± 4.036

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 215.8 attack_power points (39.74 DPS) | yes | Mask of the Unforgiven (13404, -11.51 DPS, sim-verified) [dungeon]; Stalwart Helm (250599, -14.36 DPS) [crafted]; Champion's Plate Helm (227043, -15.01 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.7 attack_power points (6.39 DPS) | yes | Woven Ivy Necklace (19159, -0.94 DPS) [quest]; Sunprism Pendant (20649, -2.36 DPS) [quest]; Ethereal Talisman (4430, -3.02 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 79.1 attack_power points (14.56 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Champion's Plate Shoulders (227042, -1.71 DPS) [pvp] |
| back | Cloak of Flames (3475) | World drop [world_drop] | 55.7 attack_power points (10.26 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Fel Cape (279269, -4.29 DPS) [crafted]; Dark Phantom Cape (13122, -5.23 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.5 attack_power points (28.64 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -5.53 DPS) [dungeon]; Icemail Jerkin (1981, -10.33 DPS) [world_drop] |
| wrist | Vambraces of the Sadist (13400) | Stratholme: Timmy the Cruel [dungeon] | 47.8 attack_power points (8.80 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -2.15 DPS) [world_drop]; Tranquil Wristguards (279256, -2.32 DPS) [crafted] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 71.7 attack_power points (13.20 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.08 DPS) [crafted]; Storm Gauntlets (12632, -2.94 DPS) [crafted] |
| waist | Brigam Girdle (13142) | Blackrock Spire: General Drakkisath [dungeon] | 86.5 attack_power points (15.93 DPS) | yes | General's Plate Girdle (16547, -1.78 DPS) [pvp]; Might of the Timbermaw (19044, -2.73 DPS) [crafted]; Legionnaire's Plate Cinch (16511, -6.22 DPS) [pvp] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Legionnaire's Plate Leggings (227048, -0.59 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (747.7 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Scalegut Treaders (275618, -0.44 DPS) [crafted]; Greaves of Withering Despair (22240, -8.25 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 75.7 attack_power points (13.94 DPS) | yes | Masons Fraternity Ring (9533, -9.24 DPS) [quest]; White Bone Band (11862, -9.52 DPS) [quest]; Thunderbrow Ring (13097, -9.69 DPS) [world_drop] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 32.4 attack_power points (5.96 DPS) | yes | Masons Fraternity Ring (9533, -1.27 DPS) [quest]; White Bone Band (11862, -1.55 DPS) [quest]; Thunderbrow Ring (13097, -1.72 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -9.58 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.83 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -34.74 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -6.57 DPS) [world_drop]; Guttbuster (13139, -7.37 DPS) [world_drop]; Skull Splitting Crossbow (13039, -7.47 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cloak of Flames; chest: Dawn Armor; wrist: Vambraces of the Sadist; hands: Timbermaw Brawlers; waist: Brigam Girdle; legs: Warbear Woolies; feet: Shadowcraft Boots; finger2: Band of the Penitent; trinket2: Molten Heart of the Mountain; main_hand: Arcanite Champion; ranged: Blackcrow

No-known-source sample (15 of 2044, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

