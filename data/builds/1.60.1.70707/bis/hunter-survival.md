# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 69.6. Weights run: 3.3s. Verify run: 5.6s. 266 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.893 ± 0.019, strength=1.215 ± 0.002, crit=1.358 ± 0.020 per rating point (14 rating = 1%, 19.009 per %), hit=1.926 ± 0.052 per rating point (10 rating = 1%, 19.261 per %), melee_haste=6.834 ± 1.185, intellect=1.237 ± 0.030

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 23.1 attack_power points (1.09 DPS) | yes | Defender's Leather Hood (252447, -0.63 DPS) [crafted]; Trapper's Leather Hood (252505, -0.71 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.80 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 5.8 attack_power points (0.27 DPS) | yes | Scholarly Pendant (277203, -0.25 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.44 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 14.5 attack_power points (0.68 DPS) | yes | Magician's Mantle (12998, -0.27 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.45 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 17.4 attack_power points (0.82 DPS) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Sentry Cloak (2059, -0.27 DPS) [world_drop]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 31.8 attack_power points (1.50 DPS) | yes | Brawler's Leather Armor (252490, -0.25 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.37 DPS) [crafted]; Prospector's Chestpiece (14562, -0.57 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 14.5 attack_power points (0.68 DPS) | yes | Bravo's Armbands (270015, -0.10 DPS) [quest]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Bristlebark Bindings (14569, -0.16 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.6 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Fletcher's Gloves (7348, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.12 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (69.6 DPS) | yes | Brawler's Leather Belt (252428, -0.07 DPS) [crafted]; Trapper's Leather Belt (252429, -0.13 DPS) [crafted]; Deviate Scale Belt (6468, -2.94 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, -2.19 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (69.6 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Trapper's Leather Boots (252440, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.40 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 13.2 attack_power points (0.62 DPS) | yes | Demon Band (12054, -0.39 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.45 DPS) [world_drop]; Minor Channeling Ring (1449, -0.51 DPS) [quest] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 5.3 attack_power points (0.25 DPS) | yes | Demon Band (12054, -0.02 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; Minor Channeling Ring (1449, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.3 attack_power points (11.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Sword of the Fallen (286977, +0.00 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 232.6 attack_power points (10.95 DPS) | yes | Cruel Barb (5191, -0.62 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.6 attack_power points (0.54 DPS) | yes | Deadly Blunderbuss (4369, -0.25 DPS, sim-verified) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop]; Owlsight Rifle (15205, -0.27 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Butcher's Cleaver; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0000000000000000-500230231050000000)

Set DPS (verified): 100.4. Weights run: 3.4s. Verify run: 2.1s. 473 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.429 ± 0.015, strength=1.213 ± 0.002, crit=1.441 ± 0.021 per rating point (14 rating = 1%, 20.168 per %), hit=1.997 ± 0.061 per rating point (10 rating = 1%, 19.971 per %), melee_haste=8.976 ± 1.092, intellect=2.455 ± 0.048

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.3 attack_power points (2.21 DPS) | yes | Resilient Cap (14401, -0.42 DPS) [world_drop]; Dokebi Hat (14584, -0.42 DPS) [world_drop]; Nightsky Cowl (4039, -0.55 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.7 attack_power points (1.51 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.95 DPS) [world_drop]; Erudite's Amulet (277204, -1.23 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (1.85 DPS) | yes | Feline Mantle (3748, -0.47 DPS) [dungeon]; Mantle of Thieves (2264, -0.48 DPS) [dungeon]; Death Speaker Mantle (6685, -1.52 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 24.4 attack_power points (1.37 DPS) | yes | Hawkeye's Cloak (14593, -0.21 DPS) [world_drop]; Tigerstrike Mantle (13108, -0.28 DPS) [world_drop]; Resilient Cape (14400, -0.54 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 58.7 attack_power points (3.31 DPS) | yes | Green Leather Armor (4255, -0.47 DPS, sim-verified) [crafted]; Spirewind Fetter (9406, -1.17 DPS) [dungeon]; Resilient Tunic (14398, -1.24 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 23.3 attack_power points (1.31 DPS) | yes | Hawkeye's Bracers (14590, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.48 DPS) [world_drop]; Technician's Bracers (270042, -0.49 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 29.4 attack_power points (1.65 DPS) | yes | Insignia Gloves (6408, +0.00 DPS, sim-verified) [world_drop]; Resilient Handgrips (14403, -0.41 DPS) [world_drop]; Cartilage Shapers (273054, -0.42 DPS) [world] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 36.6 attack_power points (2.06 DPS) | yes | Prowler's Leather Belt (252459, -0.07 DPS) [crafted]; Green Leather Belt (4257, -0.41 DPS) [crafted]; Skulker's Leather Belt (252520, -0.42 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.5 attack_power points (3.13 DPS) | yes | Trapper's Leather Legguards (252517, -0.66 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -0.93 DPS) [crafted]; Petrolspill Leggings (9509, -1.21 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.1 attack_power points (1.30 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Insignia Boots (4055, -0.21 DPS) [world_drop]; Trapper's Leather Boots (252440, -0.49 DPS, sim-verified) [crafted] |
| finger1 | Seal of Wrynn (2933) | An Audience with the King [quest] | 20.7 attack_power points (1.17 DPS) | yes | Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Ring of Precision (1491, -0.35 DPS) [dungeon]; Azora's Will (4999, -0.48 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.6 attack_power points (1.10 DPS) | yes | Thunderbrow Ring (13097, +0.00 DPS, sim-verified) [world_drop]; Ring of Precision (1491, -0.28 DPS) [dungeon]; Azora's Will (4999, -0.41 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 326.9 attack_power points (18.42 DPS) | yes | Combatant Claymore (2877, +0.00 DPS) [world_drop]; Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Zealot Blade (13033) | World drop [world_drop] | 322.3 attack_power points (18.16 DPS) | yes | Bloody Brass Knuckles (7683, +0.00 DPS, sim-verified) [dungeon]; Shoni's Disarming Tool (9608, -5.08 DPS) [quest] |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | 19.5 attack_power points (1.10 DPS) | yes | Golemsight Long Gun (273029, -0.28 DPS) [dungeon]; Double-barreled Shotgun (2098, -0.49 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.55 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Seal of Wrynn; finger2: Lavishly Jeweled Ring; main_hand: Pronged Reaver; off_hand: Zealot Blade; ranged: Crystalpine Stinger

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 160.6. Weights run: 3.5s. Verify run: 4.5s. 728 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.155 ± 0.012, strength=1.210 ± 0.002, crit=1.467 ± 0.020 per rating point (14 rating = 1%, 20.544 per %), hit=1.728 ± 0.083 per rating point (10 rating = 1%, 17.279 per %), melee_haste=not significant (4.040 ± 1.074), intellect=3.383 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-verified (160.6 DPS) | yes | Thinking Cap (2624, +0.00 DPS) [world]; Enduring Cap (3020, +0.00 DPS) [world_drop]; Papal Fez (9431, +0.00 DPS) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, +0.00 DPS) [world]; Necklace of Calisea (1714, -0.26 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop]; Inquisitor's Shawl (19507, +0.00 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 40.5 attack_power points (3.16 DPS) | yes | Blackforge Cape (6424, -1.05 DPS) [dungeon]; Glyphed Cloak (4732, -1.31 DPS) [world_drop]; Resilient Cape (14400, -1.57 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Garb of Florid Feathers (273026, +0.00 DPS) [dungeon]; Red Mageweave Vest (10007, -0.63 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.01 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (160.6 DPS) | yes | Tough Scorpid Bracers (8205, +0.00 DPS) [crafted]; Windchaser Cuffs (14429, +0.00 DPS) [world_drop]; Warden's Wristbands (14600, +0.00 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Aurora Gloves (4042, +0.00 DPS) [world_drop]; Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stalker's Leather Gloves (252526, +0.00 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Razzeric's Customized Seatbelt (6726, +0.00 DPS) [quest]; Deathmage Sash (10771, +0.00 DPS) [dungeon]; Saltscale Girdle (273023, +0.00 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.10 DPS) [crafted]; Misplaced Pantaloons (276201, -0.10 DPS) [vendor] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | sim-verified (160.6 DPS) | yes | Briar Tredders (10582, +0.00 DPS) [dungeon]; Prowler's Leather Shoes (252465, +0.00 DPS) [crafted]; Stalker's Mail Boots (252562, +0.00 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Ogremind Ring (1993, -0.02 DPS) [world_drop]; Voodoo Band (1996, -0.02 DPS) [world] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | sim-verified (160.6 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Ogremind Ring (1993, +0.00 DPS) [world_drop]; Voodoo Band (1996, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (34.48 DPS) | yes | Shoni's Disarming Tool (9608, -16.37 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (160.6 DPS) | yes | Shrapnel Blaster (4127, +0.00 DPS) [quest]; Skystriker Bow (13020, +0.00 DPS) [world_drop]; Crystalpine Stinger (13037, +0.00 DPS) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Kaleidoscope Chain; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 728, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band

### Band 50 (dwarf, 0000000000000000-3210000000000000-500230231051222151)

Set DPS (verified): 216.6. Weights run: 3.5s. Verify run: 4.6s. 917 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.216 ± 0.013, strength=1.210 ± 0.002, crit=1.691 ± 0.023 per rating point (14 rating = 1%, 23.671 per %), hit=2.005 ± 0.099 per rating point (10 rating = 1%, 20.045 per %), melee_haste=10.180 ± 1.294, intellect=3.745 ± 0.052

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-verified (216.6 DPS) | yes | Winged Helm (13112, +0.00 DPS) [world_drop]; Eye of Theradras (17715, +0.00 DPS) [dungeon]; Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Horizon Choker (13085, +0.00 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Glowing Eye of Mordresh (10769, -0.40 DPS) [dungeon] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | sim-verified (216.6 DPS) | yes | Imperial Red Mantle (8250, +0.00 DPS) [world_drop]; Ironfeather Shoulders (15067, +0.00 DPS) [crafted]; Rotgrip Mantle (17732, +0.00 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (216.6 DPS) | yes | Imperial Red Cloak (8248, +0.00 DPS) [world_drop]; Keeper's Cloak (14665, +0.00 DPS) [world_drop]; Sergeant Major's Cape (16336, +0.00 DPS) [pvp] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | 94.4 attack_power points (7.75 DPS) | yes | Knight's Mail Armor (223078, -0.92 DPS) [vendor]; Relentless Chain (17777, -1.14 DPS) [quest]; Garb of Florid Feathers (273026, -1.50 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | sim-verified (216.6 DPS) | yes | Forgotten Wraps (9433, +0.00 DPS) [world_drop]; Bracers of the Stone Princess (17714, +0.00 DPS) [dungeon]; Prowler's Leather Bracers (252539, +0.00 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (216.6 DPS) | yes | Hibernal Gloves (8110, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Stalker's Mail Gauntlets (252584, +0.00 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | sim-verified (216.6 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Sagebrush Girdle (17778, +0.00 DPS) [quest]; Stalker's Mail Belt (252588, +0.00 DPS) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 92.1 attack_power points (7.56 DPS) | yes | Imperial Red Pants (8251, -0.22 DPS) [world_drop]; Knight's Mail Legplates (223074, -0.74 DPS) [vendor]; Searingscale Leggings (11749, -1.20 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | sim-verified (216.6 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Rancor Boots (11865, +0.00 DPS) [quest]; Stalker's Mail Sabatons (252576, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (216.6 DPS) | yes | Blackstone Ring (17713, +0.00 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon]; Voodoo Band (1996, -0.39 DPS) [world] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | sim-verified (216.6 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Voodoo Band (1996, +0.00 DPS) [world]; Blackstone Ring (17713, +0.00 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (216.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (216.6 DPS) | yes | - |
| main_hand | Thrash Blade (17705) | Corruption of Earth and Seed [quest] | sim-verified (216.6 DPS) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (43.21 DPS) | yes | Claw of Celebras (17738, -3.63 DPS) [dungeon]; Shoni's Disarming Tool (9608, -24.16 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skystriker Bow (13020, -0.10 DPS) [world_drop]; Crystalpine Stinger (13037, -0.59 DPS) [world_drop]; Guttbuster (13139, -0.79 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thrash Blade; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 917, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60 (dwarf, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 312.9. Weights run: 2.5s. Verify run: 7.6s. 1788 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.106 ± 0.019, strength=1.210 ± 0.002, crit=2.225 ± 0.037 per rating point (14 rating = 1%, 31.149 per %), hit=2.718 ± 0.156 per rating point (10 rating = 1%, 27.181 per %), melee_haste=25.007 ± 2.090, intellect=5.483 ± 0.081

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 283.7 attack_power points (25.59 DPS) | yes | Field Marshal's Chain Greathelm (231562, -8.40 DPS) [vendor]; Magister's Crown (16686, -9.45 DPS, sim-verified) [dungeon]; Field Marshal's Chain Helm (231580, -10.39 DPS) [pvp] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 76.8 attack_power points (6.92 DPS) | yes | Opaline Medallion (16623, -0.49 DPS) [quest]; Archlight Talisman (15856, -1.98 DPS) [quest]; Archmagister's Faceted Pendant (273049, -2.87 DPS) [world] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | 120.6 attack_power points (10.88 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Shroud of the Nathrezim (18720, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Sergeant Major's Cape (16337) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 70.4 attack_power points (6.35 DPS) | yes | Keeper's Cloak (14665, -1.40 DPS) [world_drop]; Cape of the Black Baron (13340, -1.70 DPS) [dungeon]; Imperial Red Cloak (8248, -2.39 DPS, sim-verified) [world_drop] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (312.9 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Blue Dragonscale Breastplate (15048, -5.81 DPS, sim-verified) [crafted] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 93.2 attack_power points (8.41 DPS) | yes | Marshal's Chain Bracers (16461, -1.15 DPS) [pvp]; Hero's Bracers (8302, -1.71 DPS) [dungeon]; Tranquil Wristguards (279256, -1.72 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (312.9 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Mourning Handguards (279257, +0.00 DPS) [crafted]; Savage Gladiator Grips (11730, -4.70 DPS, sim-verified) [dungeon] |
| waist | Elunarian Belt (14465) | World drop [world_drop] | 116.9 attack_power points (10.54 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Magister's Belt (16685, -0.16 DPS) [dungeon]; Devout Belt (16696, -0.65 DPS) [dungeon] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (312.9 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -5.50 DPS, sim-verified) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 90.6 attack_power points (8.17 DPS) | yes | Rancor Boots (11865, +0.00 DPS) [quest]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 65.8 attack_power points (5.93 DPS) | yes | Band of the Hierophant (13096, -0.99 DPS) [world_drop]; Blackstone Ring (17713, -1.68 DPS) [dungeon]; Voodoo Band (1996, -2.47 DPS) [world] |
| finger2 | Frigid Ring (18679) (or Band of the Hierophant (13096)) | Princess Tempestria [world] | 54.8 attack_power points (4.94 DPS) | yes | Band of the Hierophant (13096, +0.00 DPS) [world_drop]; Blackstone Ring (17713, -0.69 DPS) [dungeon]; Voodoo Band (1996, -1.48 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (312.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (312.9 DPS) | yes | Mark of the Chosen (17774, -2.05 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (312.9 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -4.19 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 618.0 attack_power points (55.73 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Bagh Nakh (250618, -4.72 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -12.24 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (312.9 DPS) | yes | Skystriker Bow (13020, +0.00 DPS) [world_drop]; Crystalpine Stinger (13037, -0.28 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -0.36 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Horizon Choker; shoulder: Magister's Mantle; back: Sergeant Major's Cape; chest: Savage Gladiator Chain; wrist: Magiskull Cuffs; hands: Devilsaur Gauntlets; waist: Elunarian Belt; legs: Devilsaur Leggings; feet: Timmy's Galoshes; finger1: Glowing Crystal Ring; finger2: Frigid Ring; main_hand: Teebu's Blazing Longsword; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1788, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60, raid preset (dwarf, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 771.5. Weights run: 2.1s. Verify run: 6.4s. 1788 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.239 ± 0.020, strength=1.210 ± 0.002, crit=2.556 ± 0.040 per rating point (14 rating = 1%, 35.783 per %), hit=3.285 ± 0.272 per rating point (10 rating = 1%, 32.847 per %), melee_haste=20.170 ± 3.223, intellect=1.939 ± 0.018

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 156.5 attack_power points (28.13 DPS) | yes | Field Marshal's Chain Greathelm (231562, -0.29 DPS) [vendor]; Field Marshal's Chain Helm (231580, -3.83 DPS) [pvp]; Mask of the Unforgiven (13404, -16.91 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (6.32 DPS) | yes | Horizon Choker (13085, -1.44 DPS) [world_drop]; Sunprism Pendant (20649, -1.49 DPS) [quest]; Opaline Medallion (16623, -1.79 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 66.8 attack_power points (12.01 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.6 attack_power points (9.63 DPS) | yes | Sergeant Major's Cape (16337, -1.88 DPS) [pvp]; Fel Cape (279269, -3.20 DPS) [crafted]; Dark Phantom Cape (13122, -3.60 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.5 attack_power points (27.23 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -2.93 DPS) [vendor]; Savage Gladiator Chain (11726, -6.78 DPS, sim-verified) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 38.8 attack_power points (6.97 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.06 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (771.5 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -9.24 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 56.9 attack_power points (10.22 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.02 DPS) [crafted]; Stalker's Mail Belt (252588, -1.10 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (771.5 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Knight's Chain Legplates (220832, -11.76 DPS, sim-verified) [vendor] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 57.5 attack_power points (10.33 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 52.8 attack_power points (9.50 DPS) | yes | Masons Fraternity Ring (9533, -3.87 DPS) [quest]; Ring of the Underwood (2951, -5.04 DPS) [world_drop]; Falcon's Hook (7552, -5.23 DPS) [dungeon] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 35.8 attack_power points (6.43 DPS) | yes | Masons Fraternity Ring (9533, -0.80 DPS) [quest]; Ring of the Underwood (2951, -1.97 DPS) [world_drop]; Falcon's Hook (7552, -2.16 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (771.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (771.5 DPS) | yes | Mark of the Chosen (17774, -5.18 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (771.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -33.10 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 623.2 attack_power points (112.02 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Bagh Nakh (250618, -7.93 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -25.34 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (771.5 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.48 DPS) [world_drop]; Eaglehorn Long Bow (13023, -3.09 DPS) [world_drop]; Guttbuster (13139, -3.24 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Sagebrush Girdle; legs: Devilsaur Leggings; feet: Timmy's Galoshes; finger1: Blackstone Ring; finger2: Band of the Penitent; main_hand: Teebu's Blazing Longsword; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1788, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 70.4. Weights run: 3.3s. Verify run: 5.4s. 262 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.893 ± 0.019, strength=1.215 ± 0.002, crit=1.358 ± 0.020 per rating point (14 rating = 1%, 19.009 per %), hit=1.926 ± 0.052 per rating point (10 rating = 1%, 19.261 per %), melee_haste=6.834 ± 1.185, intellect=1.237 ± 0.030

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 23.1 attack_power points (1.09 DPS) | yes | Defender's Leather Hood (252447, -0.63 DPS) [crafted]; Trapper's Leather Hood (252505, -0.74 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.80 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 5.8 attack_power points (0.27 DPS) | yes | Roadwatcher's Confidence (281265, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 14.5 attack_power points (0.68 DPS) | yes | Magician's Mantle (12998, -0.29 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.45 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 17.4 attack_power points (0.82 DPS) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Sentry Cloak (2059, -0.27 DPS) [world_drop]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (70.4 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.23 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 14.5 attack_power points (0.68 DPS) | yes | Wolf Bracers (4794, -0.14 DPS) [vendor]; Bristlebark Bindings (14569, -0.16 DPS) [world_drop]; Spare Part Bindings (279875, -0.27 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 22.2 attack_power points (1.05 DPS) | yes | Fletcher's Gloves (7348, -0.15 DPS) [crafted]; Serpent Gloves (5970, -0.23 DPS) [dungeon]; Bristlebark Gloves (14572, -0.27 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (70.4 DPS) | yes | Brawler's Leather Belt (252428, -0.07 DPS) [crafted]; Trapper's Leather Belt (252429, -0.13 DPS) [crafted]; Deviate Scale Belt (6468, -2.95 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.4 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, -2.26 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (70.4 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Trapper's Leather Boots (252440, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.39 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 13.2 attack_power points (0.62 DPS) | yes | Bounty Hunter's Ring (5351, -0.21 DPS) [quest]; The 1 Ring (8350, -0.37 DPS) [world]; Demon Band (12054, -0.39 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 9.8 attack_power points (0.46 DPS) | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world]; Demon Band (12054, -0.23 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.3 attack_power points (11.22 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Wingblade (6504, +0.00 DPS, sim-verified) [quest] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 232.6 attack_power points (10.95 DPS) | yes | Cruel Barb (5191, -0.32 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.6 attack_power points (0.54 DPS) | yes | Deadly Blunderbuss (4369, -0.27 DPS, sim-verified) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop]; Privateer Musket (5309, -0.27 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Butcher's Cleaver; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band

### Band 30 (troll, 0000000000000000-0000000000000000-500230231050000000)

Set DPS (verified): 100.8. Weights run: 3.4s. Verify run: 2.0s. 470 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.429 ± 0.015, strength=1.213 ± 0.002, crit=1.441 ± 0.021 per rating point (14 rating = 1%, 20.168 per %), hit=1.997 ± 0.061 per rating point (10 rating = 1%, 19.971 per %), melee_haste=8.976 ± 1.092, intellect=2.455 ± 0.048

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 39.3 attack_power points (2.21 DPS) | yes | Resilient Cap (14401, -0.42 DPS) [world_drop]; Dokebi Hat (14584, -0.42 DPS) [world_drop]; Nightsky Cowl (4039, -0.55 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 26.7 attack_power points (1.51 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.95 DPS) [world_drop]; Erudite's Amulet (277204, -1.23 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (1.85 DPS) | yes | Feline Mantle (3748, -0.47 DPS) [dungeon]; Mantle of Thieves (2264, -0.48 DPS) [dungeon]; Death Speaker Mantle (6685, -1.57 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.6 attack_power points (1.16 DPS) | yes | Tigerstrike Mantle (13108, -0.07 DPS) [world_drop]; Resilient Cape (14400, -0.33 DPS) [world_drop]; Swiftrunner Cape (6745, -0.34 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 58.7 attack_power points (3.31 DPS) | yes | Green Leather Armor (4255, -0.54 DPS, sim-verified) [crafted]; Spirewind Fetter (9406, -1.17 DPS) [dungeon]; Resilient Tunic (14398, -1.24 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 23.3 attack_power points (1.31 DPS) | yes | Hawkeye's Bracers (14590, +0.00 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.48 DPS) [world_drop]; Technician's Bracers (270042, -0.49 DPS) [quest] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 29.4 attack_power points (1.65 DPS) | yes | Gloves of Kapelan (6744, -0.35 DPS) [quest]; Insignia Gloves (6408, -0.35 DPS) [world_drop]; Resilient Handgrips (14403, -0.41 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 36.6 attack_power points (2.06 DPS) | yes | Prowler's Leather Belt (252459, -0.07 DPS) [crafted]; Green Leather Belt (4257, -0.41 DPS) [crafted]; Skulker's Leather Belt (252520, -0.42 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 55.5 attack_power points (3.13 DPS) | yes | Trapper's Leather Legguards (252517, -0.62 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, -0.93 DPS) [crafted]; Petrolspill Leggings (9509, -1.21 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.1 attack_power points (1.30 DPS) | yes | Acidic Walkers (9454, -0.19 DPS) [dungeon]; Insignia Boots (4055, -0.21 DPS) [world_drop]; Trapper's Leather Boots (252440, -0.50 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.6 attack_power points (1.10 DPS) | yes | Loop of Sacrifice (281673, -0.21 DPS) [quest]; Ring of Precision (1491, -0.28 DPS) [dungeon]; Band of the Fist (17694, -0.28 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 17.0 attack_power points (0.96 DPS) | yes | Ring of Precision (1491, -0.14 DPS) [dungeon]; Band of the Fist (17694, -0.14 DPS) [quest]; Loop of Sacrifice (281673, -0.60 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | 326.9 attack_power points (18.42 DPS) | yes | Combatant Claymore (2877, +0.00 DPS) [world_drop]; Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Zealot Blade (13033) | World drop [world_drop] | 322.3 attack_power points (18.16 DPS) | yes | Bloody Brass Knuckles (7683, +0.00 DPS) [dungeon] |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | 19.5 attack_power points (1.10 DPS) | yes | Golemsight Long Gun (273029, -0.28 DPS) [dungeon]; Double-barreled Shotgun (2098, -0.49 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.55 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Zealot Blade; ranged: Crystalpine Stinger

No-known-source sample (15 of 470, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 157.2. Weights run: 3.5s. Verify run: 4.1s. 716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.155 ± 0.012, strength=1.210 ± 0.002, crit=1.467 ± 0.020 per rating point (14 rating = 1%, 20.544 per %), hit=1.728 ± 0.083 per rating point (10 rating = 1%, 17.279 per %), melee_haste=not significant (4.040 ± 1.074), intellect=3.383 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-verified (157.2 DPS) | yes | Thinking Cap (2624, +0.00 DPS) [world]; Enduring Cap (3020, +0.00 DPS) [world_drop]; Papal Fez (9431, +0.00 DPS) [dungeon] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | sim-verified (157.2 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, +0.00 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop]; Inquisitor's Shawl (19507, +0.00 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Glyphed Cloak (4732, +0.00 DPS) [world_drop]; Blackforge Cape (6424, +0.00 DPS) [dungeon]; Resilient Cape (14400, +0.00 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Garb of Florid Feathers (273026, +0.00 DPS) [dungeon]; Red Mageweave Vest (10007, -0.63 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.01 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | sim-verified (157.2 DPS) | yes | Tough Scorpid Bracers (8205, +0.00 DPS) [crafted]; Windchaser Cuffs (14429, +0.00 DPS) [world_drop]; Warden's Wristbands (14600, +0.00 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Aurora Gloves (4042, +0.00 DPS) [world_drop]; Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stalker's Leather Gloves (252526, +0.00 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Razzeric's Customized Seatbelt (6726, +0.00 DPS) [quest]; Deathmage Sash (10771, +0.00 DPS) [dungeon]; Saltscale Girdle (273023, +0.00 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.10 DPS) [crafted]; Misplaced Pantaloons (276201, -0.10 DPS) [vendor] |
| feet | Blackforge Greaves (6423) | Uldaman: Ancient Treasure [dungeon] | sim-verified (157.2 DPS) | yes | Briar Tredders (10582, +0.00 DPS) [dungeon]; Prowler's Leather Shoes (252465, +0.00 DPS) [crafted]; Stalker's Mail Boots (252562, +0.00 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Ogremind Ring (1993, -0.02 DPS) [world_drop]; Voodoo Band (1996, -0.02 DPS) [world] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | sim-verified (157.2 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Ogremind Ring (1993, +0.00 DPS) [world_drop]; Voodoo Band (1996, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fiery War Axe (870, +0.00 DPS) [world_drop]; Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (34.48 DPS) | yes | - |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-verified (157.2 DPS) | yes | Skystriker Bow (13020, +0.00 DPS) [world_drop]; Crystalpine Stinger (13037, +0.00 DPS) [world_drop]; The Silencer (13138, -0.08 DPS) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Ethereal Talisman; chest: Icemail Jerkin; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Blackforge Greaves; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 0000000000000000-3210000000000000-500230231051222151)

Set DPS (verified): 217.8. Weights run: 3.5s. Verify run: 3.9s. 906 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.216 ± 0.013, strength=1.210 ± 0.002, crit=1.691 ± 0.023 per rating point (14 rating = 1%, 23.671 per %), hit=2.005 ± 0.099 per rating point (10 rating = 1%, 20.045 per %), melee_haste=10.180 ± 1.294, intellect=3.745 ± 0.052

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 94.8 attack_power points (7.78 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Tough Scorpid Helm (8208, -0.93 DPS) [crafted]; Winged Helm (13112, -2.09 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (217.8 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Glowing Eye of Mordresh (10769, -0.40 DPS) [dungeon]; Horizon Choker (13085, -2.50 DPS, sim-verified) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 74.9 attack_power points (6.15 DPS) | yes | Rotgrip Mantle (17732, -0.61 DPS) [dungeon]; Imperial Red Mantle (8250, -0.81 DPS) [world_drop]; Sheepshear Mantle (13115, -1.23 DPS) [world_drop] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 41.2 attack_power points (3.38 DPS) | yes | Keeper's Cloak (14665, -0.31 DPS) [world_drop]; Big Voodoo Cloak (8216, -0.61 DPS) [crafted]; Highborne Cloak (14450, -0.61 DPS) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) | PvP rank 10 · Stone Guard · Horde [vendor] | 94.4 attack_power points (7.75 DPS) | yes | Relentless Chain (17777, -1.14 DPS) [quest]; Garb of Florid Feathers (273026, -1.50 DPS) [dungeon]; Stone Guard's Mail Armor (220826, -1.74 DPS, sim-verified) [vendor] |
| wrist | Forgotten Wraps (9433) | World drop [world_drop] | 51.6 attack_power points (4.24 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Prowler's Leather Bracers (252539, -0.53 DPS) [crafted]; Shizzle's Nozzle Wiper (11917, -0.55 DPS) [quest] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | 64.7 attack_power points (5.31 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.44 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.73 DPS) [crafted]; Hibernal Gloves (8110, -0.82 DPS, sim-verified) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 78.2 attack_power points (6.42 DPS) | yes | Stalker's Mail Belt (252588, -0.80 DPS) [crafted]; Serenity Belt (13144, -1.19 DPS) [world_drop]; Dawnspire Cord (12466, -2.78 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 92.1 attack_power points (7.56 DPS) | yes | Stone Guard's Mail Legplates (220834, -0.74 DPS) [vendor]; Searingscale Leggings (11749, -1.20 DPS) [dungeon]; Imperial Red Pants (8251, -1.79 DPS, sim-verified) [world_drop] |
| feet | Rancor Boots (11865) | Ribbly Screwspigot [quest] | 67.8 attack_power points (5.56 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Stalker's Mail Sabatons (252576, -0.43 DPS) [crafted]; Greaves of Withering Despair (22240, -0.84 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.0 attack_power points (3.29 DPS) | yes | Lavishly Jeweled Ring (1156, -1.08 DPS) [dungeon]; Ogremind Ring (1993, -1.14 DPS) [world_drop]; Voodoo Band (1996, -1.14 DPS) [world] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.0 attack_power points (2.55 DPS) | yes | Ogremind Ring (1993, -0.39 DPS) [world_drop]; Voodoo Band (1996, -0.39 DPS) [world]; Lavishly Jeweled Ring (1156, -2.21 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -4.10 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thrash Blade (17705, +0.00 DPS, sim-verified) [quest] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (43.21 DPS) | yes | Julie's Dagger (6660, +0.00 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.63 DPS) [dungeon]; White Bone Shredder (11863, -5.83 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalpine Stinger (13037, -0.59 DPS) [world_drop]; Guttbuster (13139, -0.79 DPS) [world_drop]; Skystriker Bow (13020, -1.65 DPS, sim-verified) [world_drop] |

**New at 50:** head: Eye of Theradras; neck: Skibi's Pendant; shoulder: Ironfeather Shoulders; back: Imperial Red Cloak; chest: Stone Guard's Chain Armor; wrist: Forgotten Wraps; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Rancor Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Dawn's Edge; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 906, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 314.0. Weights run: 2.5s. Verify run: 7.0s. 1785 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.106 ± 0.019, strength=1.210 ± 0.002, crit=2.225 ± 0.037 per rating point (14 rating = 1%, 31.149 per %), hit=2.718 ± 0.156 per rating point (10 rating = 1%, 27.181 per %), melee_haste=25.007 ± 2.090, intellect=5.483 ± 0.081

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 283.7 attack_power points (25.59 DPS) | yes | Warlord's Chain Greathelm (231568, -8.40 DPS) [vendor]; Magister's Crown (16686, -9.73 DPS, sim-verified) [dungeon]; Warlord's Chain Helmet (16566, -10.39 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 76.8 attack_power points (6.92 DPS) | yes | Archlight Talisman (15856, -1.98 DPS) [quest]; Archmagister's Faceted Pendant (273049, -2.87 DPS) [world]; Glowing Eye of Mordresh (10769, -2.97 DPS) [dungeon] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | 120.6 attack_power points (10.88 DPS) | yes | Shroud of the Nathrezim (18720, +0.00 DPS, sim-verified) [dungeon]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Imperial Red Cloak (8248) | World drop [world_drop] | 60.3 attack_power points (5.44 DPS) | yes | Keeper's Cloak (14665, -0.49 DPS) [world_drop]; Cape of the Black Baron (13340, -0.79 DPS) [dungeon]; Big Voodoo Cloak (8216, -0.99 DPS) [crafted] |
| chest | Savage Gladiator Chain (11726) | Blackrock Depths: Gorosh the Dervish [dungeon] | sim-verified (314.0 DPS) | yes | Warlord's Chain Chestpiece (16565, +0.00 DPS) [vendor]; Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Blue Dragonscale Breastplate (15048, -6.19 DPS, sim-verified) [crafted] |
| wrist | Magiskull Cuffs (13107) | World drop [world_drop] | 93.2 attack_power points (8.41 DPS) | yes | General's Chain Wristguards (16570, -1.15 DPS) [pvp]; Hero's Bracers (8302, -1.71 DPS) [dungeon]; Tranquil Wristguards (279256, -1.72 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (314.0 DPS) | yes | General's Chain Grips (231569, +0.00 DPS) [vendor]; Mourning Handguards (279257, +0.00 DPS) [crafted]; Savage Gladiator Grips (11730, -6.22 DPS, sim-verified) [dungeon] |
| waist | Elunarian Belt (14465) | World drop [world_drop] | 116.9 attack_power points (10.54 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Magister's Belt (16685, -0.16 DPS) [dungeon]; Devout Belt (16696, -0.65 DPS) [dungeon] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (314.0 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; Red Dragonscale Leggings (252603, +0.00 DPS) [crafted]; Savage Gladiator Leggings (11728, -7.64 DPS, sim-verified) [dungeon] |
| feet | Savage Gladiator Greaves (11731) | Blackrock Depths: Anub'shiah [dungeon] | sim-verified (314.0 DPS) | yes | Timmy's Galoshes (13402, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 65.8 attack_power points (5.93 DPS) | yes | Band of the Hierophant (13096, -0.99 DPS) [world_drop]; Chemist's Ring (15702, -0.99 DPS) [quest]; Blackstone Ring (17713, -1.68 DPS) [dungeon] |
| finger2 | Frigid Ring (18679) (or Band of the Hierophant (13096), Chemist's Ring (15702)) | Princess Tempestria [world] | 54.8 attack_power points (4.94 DPS) | yes | Band of the Hierophant (13096, +0.00 DPS) [world_drop]; Chemist's Ring (15702, +0.00 DPS) [quest]; Blackstone Ring (17713, -0.69 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (314.0 DPS) | yes | Molten Heart of the Mountain (249470, -5.82 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (314.0 DPS) | yes | Molten Heart of the Mountain (249470, -1.94 DPS, sim-verified) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (314.0 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -4.81 DPS, sim-verified) [crafted] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 618.0 attack_power points (55.73 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Bagh Nakh (250618, -5.55 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -12.24 DPS) [dungeon] |
| ranged | Skystriker Bow (13020) | World drop [world_drop] | sim-verified (314.0 DPS) | yes | Blackcrow (12651, +0.00 DPS, sim-verified) [dungeon]; Crystalpine Stinger (13037, -0.68 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -0.76 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Horizon Choker; shoulder: Magister's Mantle; chest: Savage Gladiator Chain; wrist: Magiskull Cuffs; hands: Devilsaur Gauntlets; waist: Elunarian Belt; legs: Devilsaur Leggings; feet: Savage Gladiator Greaves; finger1: Glowing Crystal Ring; finger2: Frigid Ring; trinket2: Frozen Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Heartseeker; ranged: Skystriker Bow

No-known-source sample (15 of 1785, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 0000000000000000-3250050010000000-500230231051222151)

Set DPS (verified): 781.8. Weights run: 2.1s. Verify run: 6.3s. 1785 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, agility=2.239 ± 0.020, strength=1.210 ± 0.002, crit=2.556 ± 0.040 per rating point (14 rating = 1%, 35.783 per %), hit=3.285 ± 0.272 per rating point (10 rating = 1%, 32.847 per %), melee_haste=20.170 ± 3.223, intellect=1.939 ± 0.018

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 156.5 attack_power points (28.13 DPS) | yes | Warlord's Chain Greathelm (231568, -0.29 DPS) [vendor]; Warlord's Chain Helmet (16566, -3.83 DPS) [vendor]; Mask of the Unforgiven (13404, -21.45 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (6.32 DPS) | yes | Woven Ivy Necklace (19159, -1.39 DPS) [quest]; Horizon Choker (13085, -1.44 DPS) [world_drop]; Sunprism Pendant (20649, -1.49 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 66.8 attack_power points (12.01 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.6 attack_power points (9.63 DPS) | yes | Dark Phantom Cape (13122, -3.60 DPS) [world_drop]; Cloak of Flames (3475, -3.73 DPS) [world_drop]; Fel Cape (279269, -5.72 DPS, sim-verified) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.5 attack_power points (27.23 DPS) | yes | Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -2.93 DPS) [vendor]; Savage Gladiator Chain (11726, -6.80 DPS, sim-verified) [dungeon] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 38.8 attack_power points (6.97 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.06 DPS) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (781.8 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Mourning Handguards (279257, -13.52 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 56.9 attack_power points (10.22 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.02 DPS) [crafted]; Stalker's Mail Belt (252588, -1.10 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (781.8 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Stone Guard's Chain Legplates (220833, -15.70 DPS, sim-verified) [vendor] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 57.5 attack_power points (10.33 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.10 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 52.8 attack_power points (9.50 DPS) | yes | Masons Fraternity Ring (9533, -3.87 DPS) [quest]; Ring of the Underwood (2951, -5.04 DPS) [world_drop]; White Bone Band (11862, -5.19 DPS) [quest] |
| finger2 | Band of the Penitent (13217) | Houses of the Holy [quest] | 35.8 attack_power points (6.43 DPS) | yes | Masons Fraternity Ring (9533, -0.80 DPS) [quest]; Ring of the Underwood (2951, -1.97 DPS) [world_drop]; White Bone Band (11862, -2.12 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (781.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (781.8 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (781.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Sword of Zeal (6622, -23.52 DPS, sim-verified) [world_drop] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 623.2 attack_power points (112.02 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Bagh Nakh (250618, -8.56 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -25.34 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (781.8 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.48 DPS) [world_drop]; Eaglehorn Long Bow (13023, -3.09 DPS) [world_drop]; Guttbuster (13139, -3.24 DPS) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; legs: Devilsaur Leggings; feet: Timmy's Galoshes; finger2: Band of the Penitent; trinket2: Frozen Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1785, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

