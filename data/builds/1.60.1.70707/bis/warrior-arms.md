# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 25.5. Weights run: 1.5s. Verify run: 2.4s. 300 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.209 ± 0.030, crit=0.828 ± 0.027 per rating point (14 rating = 1%, 11.592 per %), hit=1.232 ± 0.083 per rating point (10 rating = 1%, 12.321 per %), melee_haste=8.730 ± 0.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.28 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.4 attack_power points (0.09 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.00 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.00 DPS) [crafted] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 8.4 attack_power points (0.30 DPS) | yes | Lambent Scale Cloak (4706, -0.01 DPS) [world_drop]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (25.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -1.01 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.04 DPS) [quest]; Cryptwalker Bracers (280095, -0.07 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.18 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Cobrahn's Grasp (6460, -1.37 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.15 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.2 attack_power points (0.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; Ring of the Moon (12052, -0.04 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 302.8 attack_power points (10.71 DPS) | yes | Duskbringer (2205, -0.30 DPS) [dungeon]; Night Reaver (1318, -0.50 DPS) [dungeon]; Living Root (6631, -0.94 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.8 attack_power points (0.17 DPS) | yes | Dwarven Fishing Pole (3567, -0.03 DPS) [quest]; Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 73.3. Weights run: 1.5s. Verify run: 1.6s. 558 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.258 ± 0.064, crit=1.785 ± 0.070 per rating point (14 rating = 1%, 24.983 per %), hit=2.019 ± 0.197 per rating point (10 rating = 1%, 20.187 per %), melee_haste=10.573 ± 1.434

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 38.3 attack_power points (1.47 DPS) | yes | Tusken Helm (6686, -0.47 DPS) [dungeon]; Slayer's Skullcap (14753, -0.52 DPS) [world_drop]; Guard's Chain Helm (250499, -1.02 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.0 attack_power points (0.66 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.42 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.8 attack_power points (1.34 DPS) | yes | Barbaric Iron Shoulders (7913, -0.36 DPS) [crafted]; Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Shoulders (5964, -0.52 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.8 attack_power points (0.84 DPS) | yes | Sergeant Major's Cape (16315, -0.18 DPS) [pvp]; Fenrus' Hide (6340, -0.32 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.33 DPS, sim-verified) [world_drop] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Brawler's Leather Tunic (252508, -0.13 DPS) [crafted]; Avenger's Armor (1488, -0.14 DPS) [dungeon]; Dusky Leather Armor (7374, -0.78 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.6 attack_power points (0.83 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Barbaric Bracers (18948, -0.17 DPS) [crafted]; Pugilist Bracers (4438, -0.21 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.6 attack_power points (1.06 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.10 DPS) [world_drop]; Fletcher's Gloves (7348, -0.10 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Warden's Leather Belt (252460, -0.31 DPS) [crafted]; Skulker's Leather Belt (252520, -0.33 DPS, sim-verified) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 37.8 attack_power points (1.45 DPS) | yes | Brawler's Leather Pants (252500, -0.21 DPS) [crafted]; Trapper's Leather Pants (252501, -0.21 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.37 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.1 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Alacritous Treads (277234, -0.11 DPS) [quest]; Veteran's Boots (250503, -0.19 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.8 attack_power points (0.88 DPS) | yes | Seal of Wrynn (2933, -0.38 DPS) [quest]; Defias Renegade Ring (1076, -0.44 DPS) [dungeon]; Monkey Ring (6748, -0.44 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 13.6 attack_power points (0.52 DPS) | yes | Seal of Wrynn (2933, -0.03 DPS) [quest]; Defias Renegade Ring (1076, -0.09 DPS) [dungeon]; Monkey Ring (6748, -0.09 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (73.3 DPS) | yes | Corpsemaker (6687, -0.54 DPS) [dungeon]; Felbough Arm (273037, -0.89 DPS) [world]; Viscous Hammer (13045, -20.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 13.6 attack_power points (0.52 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.17 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Ring of Precision; main_hand: Manual Crowd Pummeler; ranged: Golemsight Long Gun

No-known-source sample (15 of 558, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 96.2. Weights run: 1.7s. Verify run: 1.7s. 788 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.862 ± 0.063, crit=1.809 ± 0.078 per rating point (14 rating = 1%, 25.325 per %), hit=1.807 ± 0.208 per rating point (10 rating = 1%, 18.071 per %), melee_haste=13.533 ± 1.711

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 46.3 attack_power points (2.38 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.60 DPS) [crafted]; Hawkeye's Helm (14591, -0.92 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 attack_power points (0.79 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.18 DPS) [world]; River Pride Choker (13087, -0.38 DPS) [world_drop]; Gazlowe's Charm (13088, -0.38 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.5 attack_power points (1.57 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.38 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.44 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 23.2 attack_power points (1.19 DPS) | yes | Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.31 DPS) [quest]; Parachute Cloak (10518, -0.43 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 44.7 attack_power points (2.30 DPS) | yes | Kolkar Marauder Chain (6773, -0.19 DPS) [quest]; Carapace of Tuten'kash (10775, -0.50 DPS) [dungeon]; Wolffear Harness (13110, -0.67 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.4 attack_power points (1.10 DPS) | yes | Branded Leather Bracers (19508, -0.07 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.12 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 45.3 attack_power points (2.33 DPS) | yes | Scarlet Gauntlets (10331, -0.14 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.34 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.44 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.8 attack_power points (1.79 DPS) | yes | Officer's Belt (250556, -0.18 DPS) [crafted]; Prowler's Leather Belt (252459, -0.29 DPS) [crafted]; Skulker's Leather Belt (252520, -0.31 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.16 DPS) | yes | Symbolic Legplates (14829, -0.04 DPS) [world_drop]; Triprunner Dungarees (9624, -0.13 DPS) [quest]; Basilisk Hide Pants (1718, -0.15 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.0 attack_power points (2.01 DPS) | yes | Blackforge Greaves (6423, -0.12 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.8 attack_power points (1.17 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.14 DPS) [dungeon]; Assault Band (13095, -0.14 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (96.2 DPS) | yes | Primitive Fishing Pole (276203, -1.09 DPS) [vendor]; Bonebiter (6830, -1.23 DPS) [quest]; Fiery War Axe (870, -3.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (96.2 DPS) | yes | Master Hunter's Rifle (17687, -0.11 DPS) [quest]; The Silencer (13138, -0.18 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.05 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 788, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 179.8. Weights run: 1.8s. Verify run: 2.1s. 1011 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.211 ± 0.063, crit=1.487 ± 0.085 per rating point (14 rating = 1%, 20.812 per %), hit=2.163 ± 0.204 per rating point (10 rating = 1%, 21.632 per %), melee_haste=12.823 ± 1.747

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 40.8 attack_power points (3.31 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Sunscale Helmet (14849, -0.13 DPS) [world_drop]; Chromite Barbute (8142, -0.18 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (2.09 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.12 DPS) [world]; Kaleidoscope Chain (13084, -1.25 DPS, sim-verified) [world_drop]; River Pride Choker (13087, -1.44 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.7 attack_power points (2.90 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.3 attack_power points (1.57 DPS) | yes | Dark Phantom Cape (13122, -0.09 DPS) [world_drop]; Bloodlust Cape (14801, -0.10 DPS) [world_drop]; Pridelord Cape (14673, -0.19 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 50.8 attack_power points (4.13 DPS) | yes | Warforged Chestplate (11195, -0.23 DPS) [quest]; Warbear Harness (15064, -0.57 DPS) [crafted]; Valorous Chestguard (8274, -0.88 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.3 attack_power points (2.38 DPS) | yes | Bracers of the Stone Princess (17714, -0.10 DPS) [dungeon]; Arena Bands (18711, -0.10 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 42.9 attack_power points (3.48 DPS) | yes | Gloves of Holy Might (867, -0.17 DPS) [world_drop]; Stonerender Gauntlets (17007, -0.24 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.32 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.5 attack_power points (3.46 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Belt of the Gladiator (13134, -0.53 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 56.9 attack_power points (4.62 DPS) | yes | Gryphon Rider's Leggings (9652, -0.71 DPS) [quest]; Centurion Legplates (10740, -0.71 DPS) [quest]; Serpentskin Leggings (8262, -0.94 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.3 attack_power points (3.19 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.39 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.62 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.25 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.62 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.25 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (179.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (179.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (179.8 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Nightblade (1982, -5.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (179.8 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.25 DPS) [world_drop]; Dark Iron Rifle (16004, -1.85 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 1011, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 241.7. Weights run: 1.8s. Verify run: 6.9s. 2022 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.515 ± 0.081, crit=2.164 ± 0.115 per rating point (14 rating = 1%, 30.294 per %), hit=3.604 ± 0.283 per rating point (10 rating = 1%, 36.045 per %), melee_haste=23.267 ± 2.700

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 168.7 attack_power points (13.98 DPS) | yes | Stalwart Helm (250599, -4.05 DPS, sim-verified) [crafted]; Lieutenant Commander's Plate Helm (23314, -5.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.7 attack_power points (2.46 DPS) | yes | Sunprism Pendant (20649, -1.15 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.30 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.47 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 58.0 attack_power points (4.81 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 42.7 attack_power points (3.54 DPS) | yes | Sergeant Major's Cape (16337, -1.21 DPS) [pvp]; Fel Cape (279269, -1.36 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.66 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 123.6 attack_power points (10.24 DPS) | yes | Dawn Armor (252483, -0.10 DPS) [crafted]; Field Marshal's Plate Armor (231536, -3.32 DPS) [pvp]; Knight-Captain's Plate Hauberk (227046, -4.25 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 32.0 attack_power points (2.65 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.08 DPS) [crafted]; Deepfury Bracers (13120, -0.11 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (241.7 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.66 DPS) [world_drop]; Timbermaw Brawlers (19049, -4.33 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 63.2 attack_power points (5.24 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -1.34 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.77 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (241.7 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.23 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.1 attack_power points (4.15 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Shadowcraft Boots (16711, -0.62 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.2 attack_power points (1.76 DPS) | yes | Mark of Kern (2262, -0.10 DPS) [dungeon]; Blackstone Ring (17713, -0.10 DPS) [dungeon]; Ring of Living Stone (18400, -0.10 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.5 attack_power points (1.70 DPS) | yes | Mark of Kern (2262, -0.05 DPS) [dungeon]; Blackstone Ring (17713, -0.05 DPS) [dungeon]; Ring of Living Stone (18400, -0.05 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (241.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (241.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (241.7 DPS) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (241.7 DPS) | yes | Skull Splitting Crossbow (13039, -0.10 DPS) [world_drop]; Guttbuster (13139, -0.26 DPS) [world_drop]; Dark Iron Rifle (16004, -1.77 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2022, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (human, 02325213032515201-05050020000000000-200000000000000000)

Set DPS (verified): 726.2. Weights run: 1.6s. Verify run: 6.3s. 2022 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.666 ± 0.118, crit=2.121 ± 0.165 per rating point (14 rating = 1%, 29.700 per %), hit=5.107 ± 0.434 per rating point (10 rating = 1%, 51.073 per %), melee_haste=16.310 ± 3.610

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 201.1 attack_power points (34.37 DPS) | yes | Stalwart Helm (250599, -10.34 DPS, sim-verified) [crafted]; Lieutenant Commander's Plate Helm (23314, -12.68 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.7 attack_power points (5.58 DPS) | yes | Sunprism Pendant (20649, -2.16 DPS) [quest]; Kaleidoscope Chain (13084, -2.94 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.53 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 74.1 attack_power points (12.66 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -1.20 DPS) [pvp]; Dark Warder's Pauldrons (22241, -9.16 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.0 attack_power points (7.69 DPS) | yes | Sergeant Major's Cape (16337, -2.40 DPS) [pvp]; Fel Cape (279269, -2.61 DPS) [crafted]; Dark Phantom Cape (13122, -3.42 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 142.4 attack_power points (24.34 DPS) | yes | Timbermaw Tunic (252484, -1.73 DPS) [crafted]; Field Marshal's Plate Armor (231536, -9.26 DPS) [pvp]; Shadowcraft Tunic (16721, -9.95 DPS) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 35.2 attack_power points (6.02 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.17 DPS) [crafted]; Deepfury Bracers (13120, -0.24 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (726.2 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.13 DPS) [crafted]; Timbermaw Brawlers (19049, -21.90 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 69.5 attack_power points (11.88 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -3.03 DPS) [pvp]; Prowler's Leather Waistguard (252473, -3.77 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (726.2 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -19.66 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.3 attack_power points (8.77 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.75 DPS) [crafted]; Shadowcraft Boots (16711, -3.83 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 23.3 attack_power points (3.99 DPS) | yes | Ring of Living Stone (18400, -0.23 DPS) [quest]; Falcon's Hook (7552, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.39 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.6 attack_power points (3.86 DPS) | yes | Ring of Living Stone (18400, -0.10 DPS) [quest]; Falcon's Hook (7552, -0.17 DPS) [dungeon]; Ring of the Underwood (2951, -0.26 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (726.2 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (726.2 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (726.2 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -30.06 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (726.2 DPS) | yes | Skull Splitting Crossbow (13039, -0.47 DPS) [world_drop]; Guttbuster (13139, -0.58 DPS) [world_drop]; Dark Iron Rifle (16004, -4.36 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2022, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 25.5. Weights run: 1.5s. Verify run: 2.4s. 283 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.209 ± 0.030, crit=0.828 ± 0.027 per rating point (14 rating = 1%, 11.592 per %), hit=1.232 ± 0.083 per rating point (10 rating = 1%, 12.321 per %), melee_haste=8.730 ± 0.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.28 DPS) [crafted]; Brawler's Leather Hood (252504, -0.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 2.4 attack_power points (0.09 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 6.0 attack_power points (0.21 DPS) | yes | Rough Bronze Shoulders (3480, -0.00 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.00 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 8.4 attack_power points (0.30 DPS) | yes | Lambent Scale Cloak (4706, -0.01 DPS) [world_drop]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (25.5 DPS) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.78 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Forest Leather Bracers (3202, -0.14 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -0.94 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.18 DPS) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Cobrahn's Grasp (6460, -1.09 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (25.5 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.90 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.0 attack_power points (0.57 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; The 1 Ring (8350, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.21 DPS) | yes | Bounty Hunter's Ring (5351, -0.08 DPS) [quest]; The 1 Ring (8350, -0.10 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.83 DPS) | yes | Forsaken Greataxe (251533, +0.00 DPS, sim-verified) [quest]; Smite's Mighty Hammer (7230, -0.12 DPS) [dungeon]; Living Root (6631, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.8 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.03 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.03 DPS) [crafted]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 283, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 73.0. Weights run: 1.5s. Verify run: 1.6s. 541 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.258 ± 0.064, crit=1.785 ± 0.070 per rating point (14 rating = 1%, 24.983 per %), hit=2.019 ± 0.197 per rating point (10 rating = 1%, 20.187 per %), melee_haste=10.573 ± 1.434

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 38.3 attack_power points (1.47 DPS) | yes | Tusken Helm (6686, -0.47 DPS) [dungeon]; Slayer's Skullcap (14753, -0.52 DPS) [world_drop]; Guard's Chain Helm (250499, -0.96 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 17.0 attack_power points (0.66 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.19 DPS) [world]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.8 attack_power points (1.34 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Shoulders (5964, -0.52 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.62 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 21.8 attack_power points (0.84 DPS) | yes | Fenrus' Hide (6340, -0.32 DPS) [dungeon]; Swiftrunner Cape (6745, -0.32 DPS) [quest]; Tigerstrike Mantle (13108, -0.35 DPS, sim-verified) [world_drop] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Brawler's Leather Tunic (252508, -0.13 DPS) [crafted]; Avenger's Armor (1488, -0.14 DPS) [dungeon]; Dusky Leather Armor (7374, -0.62 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.6 attack_power points (0.83 DPS) | yes | Yorgen Bracers (13012, -0.11 DPS) [world_drop]; Barbaric Bracers (18948, -0.17 DPS) [crafted]; Pugilist Bracers (4438, -0.21 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.6 attack_power points (1.06 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Mail Combat Gauntlets (4075, -0.10 DPS) [world_drop]; Fletcher's Gloves (7348, -0.10 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 33.6 attack_power points (1.29 DPS) | yes | Skulker's Leather Belt (252520, -0.05 DPS) [crafted]; Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Warden's Leather Belt (252460, -0.31 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 37.8 attack_power points (1.45 DPS) | yes | Brawler's Leather Pants (252500, -0.21 DPS) [crafted]; Trapper's Leather Pants (252501, -0.21 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.35 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.1 attack_power points (0.93 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Veteran's Boots (250503, -0.19 DPS) [crafted]; Insignia Boots (4055, -0.23 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.8 attack_power points (0.88 DPS) | yes | Ring of Precision (1491, -0.35 DPS) [dungeon]; Defias Renegade Ring (1076, -0.44 DPS) [dungeon]; Monkey Ring (6748, -0.44 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 17.0 attack_power points (0.66 DPS) | yes | Defias Renegade Ring (1076, -0.22 DPS) [dungeon]; Monkey Ring (6748, -0.22 DPS) [quest]; Ring of Precision (1491, -0.35 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (73.0 DPS) | yes | Corpsemaker (6687, -0.54 DPS) [dungeon]; Felbough Arm (273037, -0.89 DPS) [world]; Viscous Hammer (13045, -20.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 13.6 attack_power points (0.52 DPS) | yes | Double-barreled Shotgun (2098, -0.03 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.17 DPS) [world_drop]; Silver Star (3463, -0.17 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler; ranged: Golemsight Long Gun

No-known-source sample (15 of 541, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 94.9. Weights run: 1.7s. Verify run: 1.7s. 768 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.862 ± 0.063, crit=1.809 ± 0.078 per rating point (14 rating = 1%, 25.325 per %), hit=1.807 ± 0.208 per rating point (10 rating = 1%, 18.071 per %), melee_haste=13.533 ± 1.711

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 46.3 attack_power points (2.38 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.60 DPS) [crafted]; Hawkeye's Helm (14591, -0.92 DPS) [world_drop] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 17.4 attack_power points (0.90 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.28 DPS) [world]; River Pride Choker (13087, -0.49 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.5 attack_power points (1.57 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.38 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.44 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 19.0 attack_power points (0.98 DPS) | yes | Imperial Cloak (6432, -0.21 DPS) [dungeon]; Parachute Cloak (10518, -0.21 DPS) [crafted]; Tigerstrike Mantle (13108, -0.21 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 44.7 attack_power points (2.30 DPS) | yes | Kolkar Marauder Chain (6773, -0.19 DPS) [quest]; Carapace of Tuten'kash (10775, -0.50 DPS) [dungeon]; Wolffear Harness (13110, -0.67 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.4 attack_power points (1.10 DPS) | yes | Branded Leather Bracers (19508, -0.07 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.12 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 45.3 attack_power points (2.33 DPS) | yes | Scarlet Gauntlets (10331, -0.14 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.34 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.44 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.8 attack_power points (1.79 DPS) | yes | Officer's Belt (250556, -0.18 DPS) [crafted]; Prowler's Leather Belt (252459, -0.29 DPS) [crafted]; Skulker's Leather Belt (252520, -0.31 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.16 DPS) | yes | Symbolic Legplates (14829, -0.04 DPS) [world_drop]; Triprunner Dungarees (9624, -0.13 DPS) [quest]; Basilisk Hide Pants (1718, -0.15 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.0 attack_power points (2.01 DPS) | yes | Blackforge Greaves (6423, -0.12 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.23 DPS) [crafted] |
| finger1 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.8 attack_power points (1.17 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.14 DPS) [dungeon]; Assault Band (13095, -0.14 DPS) [world_drop] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.6 attack_power points (1.16 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Mark of Kern (2262, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (94.9 DPS) | yes | Primitive Fishing Pole (276203, -1.09 DPS) [vendor]; Staff of Jordan (873, -2.14 DPS) [world_drop]; Fiery War Axe (870, -3.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (94.9 DPS) | yes | Master Hunter's Rifle (17687, -0.11 DPS) [quest]; The Silencer (13138, -0.18 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.17 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Falcon's Hook; finger2: Ring of the Underwood; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 768, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 183.6. Weights run: 1.8s. Verify run: 2.1s. 990 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.211 ± 0.063, crit=1.487 ± 0.085 per rating point (14 rating = 1%, 20.812 per %), hit=2.163 ± 0.204 per rating point (10 rating = 1%, 21.632 per %), melee_haste=12.823 ± 1.747

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 40.8 attack_power points (3.31 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Sunscale Helmet (14849, -0.13 DPS) [world_drop]; Chromite Barbute (8142, -0.18 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (2.09 DPS) | yes | Woven Ivy Necklace (19159, -0.23 DPS) [quest]; Ethereal Talisman (4430, -0.89 DPS) [quest]; Kaleidoscope Chain (13084, -1.05 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.7 attack_power points (2.90 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.16 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.2 attack_power points (1.48 DPS) | yes | Bloodlust Cape (14801, -0.01 DPS) [world_drop]; Pridelord Cape (14673, -0.10 DPS) [world_drop]; Blackflame Cape (13109, -0.30 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 10 · Stone Guard · Horde [vendor] | 50.8 attack_power points (4.13 DPS) | yes | Warforged Chestplate (11195, -0.23 DPS) [quest]; Warbear Harness (15064, -0.57 DPS) [crafted]; Valorous Chestguard (8274, -0.88 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.3 attack_power points (2.38 DPS) | yes | Bracers of the Stone Princess (17714, -0.10 DPS) [dungeon]; Arena Bands (18711, -0.10 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 42.9 attack_power points (3.48 DPS) | yes | Gloves of Holy Might (867, -0.17 DPS) [world_drop]; Stonerender Gauntlets (17007, -0.24 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.32 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.5 attack_power points (3.46 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Belt of the Gladiator (13134, -0.53 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 10 · Stone Guard · Horde [vendor] | 56.9 attack_power points (4.62 DPS) | yes | Serpentskin Leggings (8262, -0.94 DPS) [world_drop]; Golem Shard Leggings (13074, -1.05 DPS) [world_drop]; Scarlet Leggings (10330, -1.21 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.3 attack_power points (3.19 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.39 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.95 DPS) | yes | Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop]; Thunderbrow Ring (13097, -0.35 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.62 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (183.6 DPS) | yes | Molten Heart of the Mountain (249470, -4.26 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (183.6 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (183.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Nightblade (1982, -3.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (183.6 DPS) | yes | Houndmaster's Bow (11628, -0.16 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.25 DPS) [world_drop]; Dark Iron Rifle (16004, -1.88 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 990, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 248.0. Weights run: 1.8s. Verify run: 6.7s. 2005 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.515 ± 0.081, crit=2.164 ± 0.115 per rating point (14 rating = 1%, 30.294 per %), hit=3.604 ± 0.283 per rating point (10 rating = 1%, 36.045 per %), melee_haste=23.267 ± 2.700

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 168.7 attack_power points (13.98 DPS) | yes | Champion's Plate Helm (227043, -5.00 DPS) [pvp]; Blood Guard's Plate Helm (220803, -6.33 DPS) [vendor]; Stalwart Helm (250599, -7.03 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.7 attack_power points (2.46 DPS) | yes | Woven Ivy Necklace (19159, -0.34 DPS) [quest]; Sunprism Pendant (20649, -0.95 DPS) [quest]; Ethereal Talisman (4430, -1.13 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 58.0 attack_power points (4.81 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -1.86 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 42.7 attack_power points (3.54 DPS) | yes | Fel Cape (279269, -1.28 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.66 DPS) [world_drop]; Pridelord Cape (14673, -1.91 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 123.6 attack_power points (10.24 DPS) | yes | Dawn Armor (252483, -3.00 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -3.32 DPS) [pvp]; Legionnaire's Plate Hauberk (227049, -4.25 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 32.0 attack_power points (2.65 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.08 DPS) [crafted]; Deepfury Bracers (13120, -0.11 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (248.0 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.66 DPS) [world_drop]; Timbermaw Brawlers (19049, -5.03 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 63.2 attack_power points (5.24 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -1.34 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.55 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (248.0 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.79 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.1 attack_power points (4.15 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Shadowcraft Boots (16711, -0.62 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.99 DPS) | yes | Thunderbrow Ring (13097, -0.29 DPS) [world_drop]; Blackstone Ring (17713, -0.33 DPS) [dungeon]; Ring of Living Stone (18400, -0.33 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.2 attack_power points (1.76 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Blackstone Ring (17713, -0.10 DPS) [dungeon]; Ring of Living Stone (18400, -0.10 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (248.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (248.0 DPS) | yes | Molten Heart of the Mountain (249470, -3.14 DPS, sim-verified) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (248.0 DPS) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (248.0 DPS) | yes | Skull Splitting Crossbow (13039, -0.10 DPS) [world_drop]; Guttbuster (13139, -0.26 DPS) [world_drop]; Dark Iron Rifle (16004, -1.84 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Masons Fraternity Ring; main_hand: Forest Defender's Axe; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2005, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (orc, 02325213032515201-05050020000000000-200000000000000000)

Set DPS (verified): 738.1. Weights run: 1.6s. Verify run: 6.4s. 2005 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.666 ± 0.118, crit=2.121 ± 0.165 per rating point (14 rating = 1%, 29.700 per %), hit=5.107 ± 0.434 per rating point (10 rating = 1%, 51.073 per %), melee_haste=16.310 ± 3.610

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 201.1 attack_power points (34.37 DPS) | yes | Stalwart Helm (250599, -12.52 DPS, sim-verified) [crafted]; Champion's Plate Helm (227043, -12.68 DPS) [pvp]; Blood Guard's Plate Helm (220803, -15.68 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.7 attack_power points (5.58 DPS) | yes | Woven Ivy Necklace (19159, -0.76 DPS) [quest]; Sunprism Pendant (20649, -2.16 DPS) [quest]; Ethereal Talisman (4430, -2.56 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 74.1 attack_power points (12.66 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Champion's Plate Shoulders (227042, -1.20 DPS) [pvp]; Dark Warder's Pauldrons (22241, -7.84 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.0 attack_power points (7.69 DPS) | yes | Fel Cape (279269, -2.61 DPS) [crafted]; Dark Phantom Cape (13122, -3.42 DPS) [world_drop]; Pridelord Cape (14673, -4.00 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 142.4 attack_power points (24.34 DPS) | yes | Timbermaw Tunic (252484, -1.73 DPS) [crafted]; Warlord's Plate Armor (231530, -9.26 DPS) [pvp]; Shadowcraft Tunic (16721, -9.95 DPS) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 35.2 attack_power points (6.02 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.17 DPS) [crafted]; Deepfury Bracers (13120, -0.24 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (738.1 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Storm Gauntlets (12632, -1.13 DPS) [crafted]; Timbermaw Brawlers (19049, -23.40 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 69.5 attack_power points (11.88 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -3.03 DPS) [pvp]; Prowler's Leather Waistguard (252473, -4.14 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (738.1 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -21.80 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.3 attack_power points (8.77 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.17 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.75 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (4.10 DPS) | yes | Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Ring of Living Stone (18400, -0.34 DPS) [quest]; Falcon's Hook (7552, -0.41 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 23.3 attack_power points (3.99 DPS) | yes | Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Ring of Living Stone (18400, -0.23 DPS) [quest]; Falcon's Hook (7552, -0.30 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (738.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (738.1 DPS) | yes | Molten Heart of the Mountain (249470, -10.69 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (738.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -34.45 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (738.1 DPS) | yes | Dark Iron Rifle (16004, +0.00 DPS) [crafted]; Skull Splitting Crossbow (13039, -0.47 DPS) [world_drop]; Guttbuster (13139, -0.58 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Masons Fraternity Ring; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2005, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

