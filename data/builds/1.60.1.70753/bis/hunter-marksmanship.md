# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 102.7. Weights run: 2.4s. Verify run: 2.4s. 266 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.421 ± 0.015, crit=0.930 ± 0.016 per rating point (14 rating = 1%, 13.025 per %), hit=1.493 ± 0.042 per rating point (10 rating = 1%, 14.926 per %), melee_haste=5.734 ± 0.707, intellect=1.690 ± 0.036, spell_power=0.154 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.03 DPS) | yes | Trapper's Leather Hood (252505, -1.31 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.40 DPS) [crafted]; Totemic Leather Hood (252448, -1.93 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.8 ranged_attack_power points (0.51 DPS) | yes | Tarnished Locket (279870, -0.52 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.27 DPS) | yes | Magician's Mantle (12998, -0.55 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.71 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.22 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.5 ranged_attack_power points (1.52 DPS) | yes | Cape of the Brotherhood (5193, -0.25 DPS) [dungeon]; Sentry Cloak (2059, -0.51 DPS) [world_drop]; Bristlebark Cape (14571, -0.51 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 37.6 ranged_attack_power points (2.79 DPS) | yes | Trapper's Leather Armor (252491, -0.81 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -1.02 DPS) [crafted]; Mystic's Wrap (14369, -1.15 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.27 DPS) | yes | Wolf Bracers (4794, -0.25 DPS) [vendor]; Bristlebark Bindings (14569, -0.51 DPS) [world_drop]; Bravo's Armbands (270015, -0.51 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 21.6 ranged_attack_power points (1.60 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.21 DPS) [crafted]; Riding Gloves (1304, -0.47 DPS) [quest] |
| waist | Trapper's Leather Belt (252429) | Leatherworking [crafted] | 18.8 ranged_attack_power points (1.39 DPS) | yes | Blackened Defias Belt (10403, -0.06 DPS) [dungeon]; Deviate Scale Belt (6468, -0.12 DPS) [crafted]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 37.5 ranged_attack_power points (2.79 DPS) | yes | Brawler's Leather Pants (252500, -0.50 DPS) [crafted]; Bluegill Breeches (3022, -0.76 DPS) [world]; Leggings of the Fang (10410, -0.76 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.4 ranged_attack_power points (2.03 DPS) | yes | Trapper's Leather Boots (252440, -0.42 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.51 DPS) [dungeon]; Footpads of the Fang (10411, -0.51 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 17.0 ranged_attack_power points (1.26 DPS) | yes | Volcanic Rock Ring (12053, -0.88 DPS) [world_drop]; Minor Channeling Ring (1449, -0.95 DPS) [quest]; Furlfeather Ring (286745, -1.25 DPS) [world] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 5.1 ranged_attack_power points (0.38 DPS) | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Volcanic Rock Ring (12053, -0.34 DPS, sim-verified) [world_drop]; Furlfeather Ring (286745, -0.37 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 39.2 ranged_attack_power points (2.91 DPS) | yes | Bronze Dory (250603, -0.46 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.39 DPS) [world]; Piercing Axe (6094, -1.39 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.34 DPS) | yes | Lil Timmy's Peashooter (13136, -0.93 DPS) [world_drop]; Venomstrike (6469, -1.51 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -3.24 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Trapper's Leather Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 137.0. Weights run: 2.5s. Verify run: 2.4s. 473 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.292 ± 0.017, crit=1.236 ± 0.024 per rating point (14 rating = 1%, 17.308 per %), hit=1.729 ± 0.056 per rating point (10 rating = 1%, 17.294 per %), melee_haste=8.326 ± 1.112, intellect=3.323 ± 0.051, spell_power=0.164 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 53.2 ranged_attack_power points (3.75 DPS) | yes | Resilient Cap (14401, -0.71 DPS) [world_drop]; Nightsky Cowl (4039, -0.94 DPS) [world_drop]; Insignia Cap (4052, -1.17 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 31.9 ranged_attack_power points (2.25 DPS) | yes | Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Crystal Starfire Medallion (5003, -1.31 DPS) [world_drop]; Erudite's Amulet (277204, -1.79 DPS) [quest] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 37.5 ranged_attack_power points (2.64 DPS) | yes | Forest Tracker Epaulets (2278, -0.09 DPS) [world_drop]; Feline Mantle (3748, -0.26 DPS) [dungeon]; Mantle of Thieves (2264, -0.33 DPS) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 26.5 ranged_attack_power points (1.86 DPS) | yes | Tigerstrike Mantle (13108, -0.01 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Resilient Cape (14400, -0.46 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 79.5 ranged_attack_power points (5.60 DPS) | yes | Green Leather Armor (4255, -1.80 DPS, sim-verified) [crafted]; Resilient Tunic (14398, -2.10 DPS) [world_drop]; Resilient Robe (14405, -2.10 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 26.6 ranged_attack_power points (1.87 DPS) | yes | Nightsky Wristbands (6407, -0.47 DPS) [world_drop]; Insignia Bracers (6410, -0.48 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.48 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 39.8 ranged_attack_power points (2.80 DPS) | yes | Cartilage Shapers (273054, -0.71 DPS) [world]; Resilient Handgrips (14403, -0.74 DPS, sim-verified) [world_drop]; Pilferer's Gloves (7358, -0.95 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 49.6 ranged_attack_power points (3.49 DPS) | yes | Green Leather Belt (4257, -0.98 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -1.16 DPS) [crafted]; Silver-lined Belt (13011, -1.38 DPS) [world_drop] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 65.9 ranged_attack_power points (4.64 DPS) | yes | Trapper's Leather Legguards (252517, -0.42 DPS, sim-verified) [crafted]; Petrolspill Leggings (9509, -1.39 DPS) [dungeon]; Troll's Bane Leggings (13114, -1.39 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 29.8 ranged_attack_power points (2.10 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Feet of the Lynx (1121, -0.24 DPS) [world_drop]; Insignia Boots (4055, -0.24 DPS) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 26.5 ranged_attack_power points (1.87 DPS) | yes | Ring of Precision (1491, -0.48 DPS) [dungeon]; Snake Hoop (6750, -0.63 DPS) [quest]; Azora's Will (4999, -0.70 DPS) [world_drop] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 23.2 ranged_attack_power points (1.63 DPS) | yes | Ring of Precision (1491, -0.24 DPS) [dungeon]; Snake Hoop (6750, -0.39 DPS) [quest]; Azora's Will (4999, -0.46 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | 50.2 ranged_attack_power points (3.54 DPS) | yes | Impaling Harpoon (5200, -0.28 DPS) [dungeon]; Kam's Walking Stick (2280, -0.52 DPS) [dungeon]; Glimmering Staff (249392, -0.96 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-verified (137.0 DPS) | yes | Ironweaver (13137, -0.47 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.71 DPS) [world_drop]; Golemsight Long Gun (273029, -2.19 DPS, sim-verified) [dungeon] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Seal of Wrynn; main_hand: Felbough Arm; ranged: Crystalpine Stinger

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 184.2. Weights run: 2.6s. Verify run: 2.8s. 728 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.129 ± 0.017, crit=1.367 ± 0.027 per rating point (14 rating = 1%, 19.137 per %), hit=1.736 ± 0.062 per rating point (10 rating = 1%, 17.355 per %), melee_haste=11.350 ± 1.054, intellect=2.918 ± 0.054, spell_power=0.159 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 50.9 ranged_attack_power points (3.87 DPS) | yes | Warden's Wizard Hat (14604, -0.11 DPS) [world_drop]; Enduring Cap (3020, -0.32 DPS) [world_drop]; Thinking Cap (2624, -0.77 DPS, sim-verified) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 29.5 ranged_attack_power points (2.25 DPS) | yes | Glowing Eye of Mordresh (10769, -0.38 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop]; Necklace of Calisea (1714, -0.69 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | 47.3 ranged_attack_power points (3.60 DPS) | yes | Sheepshear Mantle (13115, -0.05 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.63 DPS) [dungeon]; Nightscape Shoulders (8192, -0.98 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 36.3 ranged_attack_power points (2.76 DPS) | yes | Imperial Cloak (6432, -0.86 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.86 DPS) [world_drop]; Parachute Cloak (10518, -1.74 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 92.5 ranged_attack_power points (7.04 DPS) | yes | Tough Scorpid Breastplate (8203, -1.91 DPS) [crafted]; Wolffear Harness (13110, -2.99 DPS) [world_drop]; Garb of Florid Feathers (273026, -3.44 DPS, sim-verified) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 39.4 ranged_attack_power points (3.00 DPS) | yes | Warden's Wristbands (14600, -0.51 DPS) [world_drop]; Windchaser Cuffs (14429, -0.52 DPS) [world_drop]; Green Leather Bracers (4259, -0.93 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | 57.6 ranged_attack_power points (4.38 DPS) | yes | Aurora Gloves (4042, -0.70 DPS) [world_drop]; Stalker's Leather Gloves (252526, -1.14 DPS, sim-verified) [crafted]; Mistscape Gloves (6428, -1.43 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 45.7 ranged_attack_power points (3.48 DPS) | yes | Deathmage Sash (10771, -0.06 DPS) [dungeon]; Green Leather Belt (4257, -0.71 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.81 DPS) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 65.7 ranged_attack_power points (5.00 DPS) | yes | Triprunner Dungarees (9624, -0.71 DPS) [quest]; Ferine Leggings (6690, -0.94 DPS, sim-verified) [dungeon]; Trapper's Leather Legguards (252517, -1.30 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 54.8 ranged_attack_power points (4.17 DPS) | yes | Swampwalker Boots (2276, -1.08 DPS) [world_drop]; Worn Running Boots (9398, -1.08 DPS) [dungeon]; Briar Tredders (10582, -1.35 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 31.3 ranged_attack_power points (2.38 DPS) | yes | Disengagement Ring (276202, -0.48 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Field Researcher's Loop (281634, -0.71 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 28.2 ranged_attack_power points (2.14 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Field Researcher's Loop (281634, -0.48 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (184.2 DPS) | yes | Steel Spear (250605, -2.57 DPS) [crafted]; Felbough Arm (273037, -2.72 DPS) [world]; Gut Ripper (2164, -10.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (184.2 DPS) | yes | Skystriker Bow (13020, -2.27 DPS) [world_drop]; Swiftwind (13038, -2.33 DPS) [world_drop]; The Silencer (13138, -4.06 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; shoulder: Windchaser Amice; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 728, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 242.1. Weights run: 2.7s. Verify run: 2.5s. 917 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.042 ± 0.016, crit=1.548 ± 0.030 per rating point (14 rating = 1%, 21.678 per %), hit=2.378 ± 0.075 per rating point (10 rating = 1%, 23.776 per %), melee_haste=11.371 ± 1.061, intellect=2.996 ± 0.058, spell_power=0.151 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 84.5 ranged_attack_power points (6.99 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.49 DPS) [dungeon]; Khan's Helmet (14785, -1.26 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 41.9 ranged_attack_power points (3.47 DPS) | yes | Skibi's Pendant (13089, -0.20 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.99 DPS) [world]; Glowing Eye of Mordresh (10769, -1.40 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 62.2 ranged_attack_power points (5.14 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, -0.08 DPS) [vendor]; Imperial Red Mantle (8250, -0.42 DPS) [world_drop]; Tough Scorpid Shoulders (8207, -1.24 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 45.6 ranged_attack_power points (3.77 DPS) | yes | Sergeant Major's Cape (16336, -0.78 DPS) [pvp]; Serpentskin Cloak (8259, -1.01 DPS) [world_drop]; Blackflame Cape (13109, -1.37 DPS, sim-verified) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | 97.3 ranged_attack_power points (8.04 DPS) | yes | Icemail Jerkin (1981, -0.04 DPS) [world_drop]; Garb of Florid Feathers (273026, -2.06 DPS) [dungeon]; Blazewind Breastplate (11193, -2.26 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 45.6 ranged_attack_power points (3.77 DPS) | yes | Stalker's Mail Bracers (252570, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.05 DPS) [world_drop]; Bracers of the Stone Princess (17714, -0.22 DPS) [dungeon] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | 69.6 ranged_attack_power points (5.75 DPS) | yes | Tough Scorpid Gloves (8204, -1.01 DPS) [crafted]; Hibernal Gloves (8110, -1.25 DPS, sim-verified) [world_drop]; Stalker's Leather Gloves (252526, -1.50 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 81.6 ranged_attack_power points (6.74 DPS) | yes | Stalker's Mail Belt (252588, -1.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.96 DPS) [dungeon]; Highborne Cord (14454, -2.51 DPS) [world_drop] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 94.2 ranged_attack_power points (7.79 DPS) | yes | Tough Scorpid Leggings (8206, -1.04 DPS) [crafted]; Searingscale Leggings (11749, -1.30 DPS) [dungeon]; Imperial Red Pants (8251, -1.48 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 75.8 ranged_attack_power points (6.26 DPS) | yes | Stalker's Mail Sabatons (252576, -0.77 DPS) [crafted]; Albino Crocscale Boots (17728, -1.23 DPS) [dungeon]; Rancor Boots (11865, -1.60 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.8 ranged_attack_power points (3.62 DPS) | yes | Ring of the Underwood (2951, -1.10 DPS) [world_drop]; Falcon's Hook (7552, -1.36 DPS) [dungeon]; Disengagement Ring (276202, -1.61 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.6 ranged_attack_power points (3.52 DPS) | yes | Falcon's Hook (7552, -1.26 DPS) [dungeon]; Ring of the Underwood (2951, -1.31 DPS, sim-verified) [world_drop]; Disengagement Ring (276202, -1.51 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (242.1 DPS) | yes | Uther's Strength (11302, -1.69 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (242.1 DPS) | yes | Uther's Strength (11302, -2.07 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (242.1 DPS) | yes | Glowing Brightwood Staff (812, -0.72 DPS) [world_drop]; Spellshifter Rod (9527, -1.53 DPS) [quest]; Hanzo Sword (8190, -13.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (242.1 DPS) | yes | Megashot Rifle (17717, -2.64 DPS) [dungeon]; Dark Iron Rifle (16004, -3.13 DPS) [crafted]; Gryphonwing Long Bow (13022, -3.80 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 917, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 302.8. Weights run: 2.7s. Verify run: 5.8s. 1788 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.065 ± 0.019, crit=2.035 ± 0.041 per rating point (14 rating = 1%, 28.485 per %), hit=2.901 ± 0.107 per rating point (10 rating = 1%, 29.007 per %), melee_haste=12.653 ± 1.505, intellect=3.168 ± 0.081, spell_power=0.153 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 213.2 ranged_attack_power points (18.27 DPS) | yes | Field Marshal's Chain Helm (231580, -3.64 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -5.02 DPS) [vendor]; Outlaw's Collar (279253, -10.61 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 44.3 ranged_attack_power points (3.80 DPS) | yes | Opaline Medallion (16623, -0.27 DPS) [quest]; Skibi's Pendant (13089, -0.39 DPS) [world_drop]; Sunprism Pendant (20649, -0.65 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 80.2 ranged_attack_power points (6.87 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.0 ranged_attack_power points (5.65 DPS) | yes | Sergeant Major's Cape (16337, -1.48 DPS, sim-verified) [pvp]; Dark Phantom Cape (13122, -1.71 DPS) [world_drop]; Cerise Drape (15804, -2.45 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 170.9 ranged_attack_power points (14.64 DPS) | yes | Field Marshal's Chain Breastplate (16466, -0.02 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -0.02 DPS) [vendor]; Knight's Chain Armor (220828, -6.76 DPS, sim-verified) [vendor] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 57.3 ranged_attack_power points (4.91 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.45 DPS) [pvp]; Magiskull Cuffs (13107, -1.39 DPS, sim-verified) [world_drop] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 108.7 ranged_attack_power points (9.31 DPS) | yes | Marshal's Chain Vices (231578, -0.27 DPS) [vendor]; Marshal's Chain Grips (231560, -1.28 DPS) [pvp]; Hero's Gauntlets (8305, -2.97 DPS, sim-verified) [dungeon] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 84.0 ranged_attack_power points (7.19 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Knight-Captain's Chain Girdle (16400, -1.08 DPS) [pvp]; Stalker's Mail Belt (252588, -1.37 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | 111.4 ranged_attack_power points (9.54 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -1.83 DPS, sim-verified) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 83.8 ranged_attack_power points (7.18 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 49.0 ranged_attack_power points (4.20 DPS) | yes | Glowing Crystal Ring (18402, -0.94 DPS) [quest]; Band of the Hierophant (13096, -1.48 DPS) [world_drop]; Frigid Ring (18679, -1.48 DPS) [world] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.9 ranged_attack_power points (3.68 DPS) | yes | Glowing Crystal Ring (18402, -0.42 DPS) [quest]; Band of the Hierophant (13096, -0.96 DPS) [world_drop]; Frigid Ring (18679, -0.96 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (302.8 DPS) | yes | Uther's Strength (11302, -2.16 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (302.8 DPS) | yes | Uther's Strength (11302, -1.48 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (302.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -12.95 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (302.8 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -3.90 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Timmy's Galoshes; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1788, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 664.1. Weights run: 2.7s. Verify run: 5.9s. 1788 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.231 ± 0.020, crit=2.259 ± 0.042 per rating point (14 rating = 1%, 31.625 per %), hit=3.754 ± 0.143 per rating point (10 rating = 1%, 37.542 per %), melee_haste=18.099 ± 1.896, intellect=2.656 ± 0.054, spell_power=0.502 ± 0.002

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 200.4 ranged_attack_power points (32.34 DPS) | yes | Field Marshal's Chain Helm (231580, -4.37 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -7.42 DPS) [vendor]; Outlaw's Collar (279253, -12.07 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.0 ranged_attack_power points (6.78 DPS) | yes | Sunprism Pendant (20649, -0.52 DPS) [quest]; Horizon Choker (13085, -0.78 DPS) [world_drop]; Opaline Medallion (16623, -1.21 DPS) [quest] |
| shoulder | Spaulders of the Unseen (13116) | World drop [world_drop] | 81.7 ranged_attack_power points (13.19 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, +0.00 DPS) [dungeon]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 68.5 ranged_attack_power points (11.05 DPS) | yes | Sergeant Major's Cape (16337, -3.45 DPS) [pvp]; Dark Phantom Cape (13122, -3.89 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.79 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 188.7 ranged_attack_power points (30.46 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.48 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.48 DPS) [vendor]; Icemail Jerkin (1981, -14.41 DPS, sim-verified) [world_drop] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 57.0 ranged_attack_power points (9.20 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.33 DPS) [pvp]; Beaststalker's Bindings (16681, -1.38 DPS) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 105.0 ranged_attack_power points (16.95 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -1.37 DPS) [pvp]; Hero's Gauntlets (8305, -5.41 DPS, sim-verified) [dungeon] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 80.3 ranged_attack_power points (12.96 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Stalker's Mail Belt (252588, -1.38 DPS) [crafted]; Knight-Captain's Chain Girdle (16400, -1.81 DPS) [pvp] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | 110.6 ranged_attack_power points (17.84 DPS) | yes | Knight's Chain Legplates (220832, +0.00 DPS) [vendor]; Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 81.5 ranged_attack_power points (13.15 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.5 ranged_attack_power points (9.29 DPS) | yes | Ring of the Underwood (2951, -4.07 DPS) [world_drop]; Glowing Crystal Ring (18402, -4.14 DPS) [quest]; Band of the Penitent (13217, -4.18 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.2 ranged_attack_power points (7.30 DPS) | yes | Glowing Crystal Ring (18402, -2.16 DPS) [quest]; Band of the Penitent (13217, -2.20 DPS) [quest]; Ring of the Underwood (2951, -2.69 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (664.1 DPS) | yes | Uther's Strength (11302, -4.97 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (664.1 DPS) | yes | Uther's Strength (11302, -2.34 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (664.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -25.64 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (664.1 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -11.08 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Skibi's Pendant; shoulder: Spaulders of the Unseen; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Timmy's Galoshes; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1788, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 104.0. Weights run: 2.4s. Verify run: 2.4s. 262 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.421 ± 0.015, crit=0.930 ± 0.016 per rating point (14 rating = 1%, 13.025 per %), hit=1.493 ± 0.042 per rating point (10 rating = 1%, 14.926 per %), melee_haste=5.734 ± 0.707, intellect=1.690 ± 0.036, spell_power=0.154 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.03 DPS) | yes | Trapper's Leather Hood (252505, -0.98 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.40 DPS) [crafted]; Totemic Leather Hood (252448, -1.93 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.8 ranged_attack_power points (0.51 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.27 DPS) | yes | Magician's Mantle (12998, -0.08 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -0.71 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.22 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.5 ranged_attack_power points (1.52 DPS) | yes | Cape of the Brotherhood (5193, -0.25 DPS) [dungeon]; Sentry Cloak (2059, -0.51 DPS) [world_drop]; Bristlebark Cape (14571, -0.51 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 29.0 ranged_attack_power points (2.15 DPS) | yes | Brawler's Leather Armor (252490, -0.38 DPS) [crafted]; Mystic's Wrap (14369, -0.51 DPS) [world_drop]; Mystic's Robe (14371, -0.51 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.27 DPS) | yes | Wolf Bracers (4794, -0.25 DPS) [vendor]; Bristlebark Bindings (14569, -0.51 DPS) [world_drop]; Spare Part Bindings (279875, -0.51 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 21.6 ranged_attack_power points (1.60 DPS) | yes | Gloves of the Fang (10413, -0.08 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.21 DPS) [crafted]; Forest Leather Gloves (3058, -0.59 DPS) [world_drop] |
| waist | Trapper's Leather Belt (252429) | Leatherworking [crafted] | 18.8 ranged_attack_power points (1.39 DPS) | yes | Blackened Defias Belt (10403, -0.06 DPS) [dungeon]; Deviate Scale Belt (6468, -0.12 DPS) [crafted]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 37.5 ranged_attack_power points (2.79 DPS) | yes | Leggings of the Fang (10410, -0.36 DPS, sim-verified) [dungeon]; Brawler's Leather Pants (252500, -0.50 DPS) [crafted]; Bluegill Breeches (3022, -0.76 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.4 ranged_attack_power points (2.03 DPS) | yes | Trapper's Leather Boots (252440, -0.26 DPS) [crafted]; Blackened Defias Boots (10402, -0.51 DPS) [dungeon]; Footpads of the Fang (10411, -0.51 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 17.0 ranged_attack_power points (1.26 DPS) | yes | Loop of Sacrifice (281673, -0.63 DPS) [quest]; The 1 Ring (8350, -0.88 DPS) [world]; Volcanic Rock Ring (12053, -0.88 DPS) [world_drop] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 10.3 ranged_attack_power points (0.76 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; The 1 Ring (8350, -0.38 DPS) [world]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 39.2 ranged_attack_power points (2.91 DPS) | yes | Crescent Staff (6505, +0.00 DPS, sim-verified) [quest]; Bronze Dory (250603, -0.44 DPS) [crafted]; Staff of Orgrimmar (15444, -0.64 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.34 DPS) | yes | Lil Timmy's Peashooter (13136, -0.93 DPS) [world_drop]; Venomstrike (6469, -1.02 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -3.24 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Trapper's Leather Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Bounty Hunter's Ring; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 138.1. Weights run: 2.5s. Verify run: 2.4s. 470 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.292 ± 0.017, crit=1.236 ± 0.024 per rating point (14 rating = 1%, 17.308 per %), hit=1.729 ± 0.056 per rating point (10 rating = 1%, 17.294 per %), melee_haste=8.326 ± 1.112, intellect=3.323 ± 0.051, spell_power=0.164 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 53.2 ranged_attack_power points (3.75 DPS) | yes | Resilient Cap (14401, -0.53 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -0.94 DPS) [world_drop]; Insignia Cap (4052, -1.17 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 31.9 ranged_attack_power points (2.25 DPS) | yes | Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Crystal Starfire Medallion (5003, -1.31 DPS) [world_drop]; Erudite's Amulet (277204, -1.79 DPS) [quest] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 37.5 ranged_attack_power points (2.64 DPS) | yes | Forest Tracker Epaulets (2278, -0.09 DPS) [world_drop]; Feline Mantle (3748, -0.26 DPS) [dungeon]; Mantle of Thieves (2264, -0.33 DPS) [dungeon] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.3 ranged_attack_power points (1.86 DPS) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Resilient Cape (14400, -0.45 DPS) [world_drop]; Swiftrunner Cape (6745, -0.46 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 79.5 ranged_attack_power points (5.60 DPS) | yes | Green Leather Armor (4255, -1.67 DPS, sim-verified) [crafted]; Resilient Tunic (14398, -2.10 DPS) [world_drop]; Resilient Robe (14405, -2.10 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 26.6 ranged_attack_power points (1.87 DPS) | yes | Insignia Bracers (6410, -0.48 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.48 DPS) [world_drop]; Nightsky Wristbands (6407, -0.49 DPS, sim-verified) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 39.8 ranged_attack_power points (2.80 DPS) | yes | Cartilage Shapers (273054, -0.71 DPS) [world]; Resilient Handgrips (14403, -0.93 DPS, sim-verified) [world_drop]; Shredder Operating Gloves (16740, -0.93 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 49.6 ranged_attack_power points (3.49 DPS) | yes | Green Leather Belt (4257, -0.96 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -1.16 DPS) [crafted]; Lilac Sash (6780, -1.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 65.9 ranged_attack_power points (4.64 DPS) | yes | Trapper's Leather Legguards (252517, -0.62 DPS, sim-verified) [crafted]; Petrolspill Leggings (9509, -1.39 DPS) [dungeon]; Troll's Bane Leggings (13114, -1.39 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 29.8 ranged_attack_power points (2.10 DPS) | yes | Acidic Walkers (9454, -0.17 DPS) [dungeon]; Insignia Boots (4055, -0.24 DPS) [world_drop]; Warsong Boots (16977, -0.24 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 26.5 ranged_attack_power points (1.87 DPS) | yes | Snake Hoop (6750, -0.63 DPS) [quest]; Azora's Will (4999, -0.70 DPS) [world_drop]; Loop of Sacrifice (281673, -0.70 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.7 ranged_attack_power points (1.39 DPS) | yes | Snake Hoop (6750, -0.15 DPS) [quest]; Azora's Will (4999, -0.22 DPS) [world_drop]; Loop of Sacrifice (281673, -0.22 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Impaling Harpoon (5200, -0.00 DPS) [dungeon]; Kam's Walking Stick (2280, -0.24 DPS) [dungeon]; Felbough Arm (273037, -1.54 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironweaver (13137, -0.47 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.71 DPS) [world_drop]; Golemsight Long Gun (273029, -1.60 DPS, sim-verified) [dungeon] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Ring of Precision; main_hand: Crescent Staff; ranged: Crystalpine Stinger

No-known-source sample (15 of 470, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 183.2. Weights run: 2.6s. Verify run: 2.7s. 716 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.129 ± 0.017, crit=1.367 ± 0.027 per rating point (14 rating = 1%, 19.137 per %), hit=1.736 ± 0.062 per rating point (10 rating = 1%, 17.355 per %), melee_haste=11.350 ± 1.054, intellect=2.918 ± 0.054, spell_power=0.159 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 50.9 ranged_attack_power points (3.87 DPS) | yes | Warden's Wizard Hat (14604, -0.11 DPS) [world_drop]; Enduring Cap (3020, -0.32 DPS) [world_drop]; Thinking Cap (2624, -1.01 DPS, sim-verified) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 29.5 ranged_attack_power points (2.25 DPS) | yes | Glowing Eye of Mordresh (10769, -0.38 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.40 DPS) [world_drop]; Necklace of Calisea (1714, -0.69 DPS) [dungeon] |
| shoulder | Windchaser Amice (14432) | World drop [world_drop] | 47.3 ranged_attack_power points (3.60 DPS) | yes | Sheepshear Mantle (13115, -0.05 DPS) [world_drop]; Inquisitor's Shawl (19507, -0.63 DPS) [dungeon]; Nightscape Shoulders (8192, -0.98 DPS) [crafted] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.0 ranged_attack_power points (1.91 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Blackforge Cape (6424, -0.13 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 92.5 ranged_attack_power points (7.04 DPS) | yes | Tough Scorpid Breastplate (8203, -1.91 DPS) [crafted]; Garb of Florid Feathers (273026, -2.59 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -2.99 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 39.4 ranged_attack_power points (3.00 DPS) | yes | Warden's Wristbands (14600, -0.51 DPS) [world_drop]; Windchaser Cuffs (14429, -0.52 DPS) [world_drop]; Green Leather Bracers (4259, -0.88 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | 57.6 ranged_attack_power points (4.38 DPS) | yes | Stalker's Leather Gloves (252526, -0.44 DPS) [crafted]; Aurora Gloves (4042, -0.70 DPS) [world_drop]; Mistscape Gloves (6428, -1.43 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 45.7 ranged_attack_power points (3.48 DPS) | yes | Deathmage Sash (10771, -0.06 DPS) [dungeon]; Green Leather Belt (4257, -0.71 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.81 DPS) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 65.7 ranged_attack_power points (5.00 DPS) | yes | Ferine Leggings (6690, -0.36 DPS) [dungeon]; Triprunner Dungarees (9624, -0.71 DPS) [quest]; Trapper's Leather Legguards (252517, -1.30 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 54.8 ranged_attack_power points (4.17 DPS) | yes | Swampwalker Boots (2276, -1.08 DPS) [world_drop]; Worn Running Boots (9398, -1.08 DPS) [dungeon]; Briar Tredders (10582, -1.24 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 31.3 ranged_attack_power points (2.38 DPS) | yes | Disengagement Ring (276202, -0.48 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Field Researcher's Loop (281634, -0.71 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 28.2 ranged_attack_power points (2.14 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon]; Field Researcher's Loop (281634, -0.48 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (183.2 DPS) | yes | Steel Spear (250605, -2.57 DPS) [crafted]; Felbough Arm (273037, -2.72 DPS) [world]; Gut Ripper (2164, -10.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (183.2 DPS) | yes | Skystriker Bow (13020, -2.27 DPS) [world_drop]; Swiftwind (13038, -2.33 DPS) [world_drop]; The Silencer (13138, -4.00 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; shoulder: Windchaser Amice; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 249.9. Weights run: 2.7s. Verify run: 2.5s. 906 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.042 ± 0.016, crit=1.548 ± 0.030 per rating point (14 rating = 1%, 21.678 per %), hit=2.378 ± 0.075 per rating point (10 rating = 1%, 23.776 per %), melee_haste=11.371 ± 1.061, intellect=2.996 ± 0.058, spell_power=0.151 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 84.5 ranged_attack_power points (6.99 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.49 DPS) [dungeon]; Khan's Helmet (14785, -1.26 DPS) [world_drop] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 41.9 ranged_attack_power points (3.47 DPS) | yes | Skibi's Pendant (13089, -0.20 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.99 DPS) [world]; Woven Ivy Necklace (19159, -1.20 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 62.2 ranged_attack_power points (5.14 DPS) | yes | Blood Guard's Chain Epaulets (220824, -0.08 DPS) [vendor]; Imperial Red Mantle (8250, -0.42 DPS) [world_drop]; Tough Scorpid Shoulders (8207, -0.86 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 45.6 ranged_attack_power points (3.77 DPS) | yes | Blackflame Cape (13109, -0.79 DPS, sim-verified) [world_drop]; Serpentskin Cloak (8259, -1.01 DPS) [world_drop]; Imperial Red Cloak (8248, -1.05 DPS) [world_drop] |
| chest | Stone Guard's Chain Armor (220827) | PvP rank 10 · Stone Guard · Horde [vendor] | 97.3 ranged_attack_power points (8.04 DPS) | yes | Icemail Jerkin (1981, -0.04 DPS) [world_drop]; Garb of Florid Feathers (273026, -2.06 DPS) [dungeon]; Blazewind Breastplate (11193, -2.26 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 45.6 ranged_attack_power points (3.77 DPS) | yes | Stalker's Mail Bracers (252570, -0.02 DPS) [crafted]; Forgotten Wraps (9433, -0.05 DPS) [world_drop]; Bracers of the Stone Princess (17714, -0.22 DPS) [dungeon] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | 69.6 ranged_attack_power points (5.75 DPS) | yes | Tough Scorpid Gloves (8204, -1.01 DPS) [crafted]; Hibernal Gloves (8110, -1.40 DPS, sim-verified) [world_drop]; Stalker's Leather Gloves (252526, -1.50 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 81.6 ranged_attack_power points (6.74 DPS) | yes | Stalker's Mail Belt (252588, -0.79 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -1.96 DPS) [dungeon]; Highborne Cord (14454, -2.51 DPS) [world_drop] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 94.2 ranged_attack_power points (7.79 DPS) | yes | Tough Scorpid Leggings (8206, -1.04 DPS) [crafted]; Imperial Red Pants (8251, -1.10 DPS, sim-verified) [world_drop]; Searingscale Leggings (11749, -1.30 DPS) [dungeon] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 75.8 ranged_attack_power points (6.26 DPS) | yes | Stalker's Mail Sabatons (252576, -0.77 DPS) [crafted]; Rancor Boots (11865, -1.17 DPS, sim-verified) [quest]; Albino Crocscale Boots (17728, -1.23 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.8 ranged_attack_power points (3.62 DPS) | yes | Ring of the Underwood (2951, -1.10 DPS) [world_drop]; Falcon's Hook (7552, -1.36 DPS) [dungeon]; Disengagement Ring (276202, -1.61 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.6 ranged_attack_power points (3.52 DPS) | yes | Falcon's Hook (7552, -1.26 DPS) [dungeon]; Ring of the Underwood (2951, -1.31 DPS, sim-verified) [world_drop]; Disengagement Ring (276202, -1.51 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (249.9 DPS) | yes | Frozen Heart of the Mountain (249469, -7.83 DPS, sim-verified) [crafted]; Uther's Strength (11302, -8.24 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (249.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (249.9 DPS) | yes | Glowing Brightwood Staff (812, -0.72 DPS) [world_drop]; Spellshifter Rod (9527, -1.53 DPS) [quest]; Hanzo Sword (8190, -12.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (249.9 DPS) | yes | Megashot Rifle (17717, -2.64 DPS) [dungeon]; Gryphonwing Long Bow (13022, -3.12 DPS, sim-verified) [dungeon]; Dark Iron Rifle (16004, -3.13 DPS) [crafted] |

**New at 50:** head: Tough Scorpid Helm; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 906, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 311.8. Weights run: 2.7s. Verify run: 5.8s. 1785 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.065 ± 0.019, crit=2.035 ± 0.041 per rating point (14 rating = 1%, 28.485 per %), hit=2.901 ± 0.107 per rating point (10 rating = 1%, 29.007 per %), melee_haste=12.653 ± 1.505, intellect=3.168 ± 0.081, spell_power=0.153 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 213.2 ranged_attack_power points (18.27 DPS) | yes | Warlord's Chain Helmet (16566, -3.64 DPS) [vendor]; Warlord's Chain Helm (231571, -3.64 DPS) [vendor]; Outlaw's Collar (279253, -11.15 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 44.3 ranged_attack_power points (3.80 DPS) | yes | Skibi's Pendant (13089, -0.39 DPS) [world_drop]; Sunprism Pendant (20649, -0.65 DPS) [quest]; Archlight Talisman (15856, -0.93 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 80.2 ranged_attack_power points (6.87 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.0 ranged_attack_power points (5.65 DPS) | yes | Dark Phantom Cape (13122, -1.32 DPS, sim-verified) [world_drop]; Cerise Drape (15804, -2.45 DPS) [quest]; Blackflame Cape (13109, -2.50 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 170.9 ranged_attack_power points (14.64 DPS) | yes | Warlord's Chain Chestpiece (16565, -0.02 DPS) [vendor]; Warlord's Chain Hauberk (231573, -0.02 DPS) [vendor]; Stone Guard's Chain Armor (220827, -7.23 DPS, sim-verified) [vendor] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 57.3 ranged_attack_power points (4.91 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Magiskull Cuffs (13107, -0.30 DPS) [world_drop]; Legionnaire's Chain Bracers (16517, -0.45 DPS) [pvp] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 108.7 ranged_attack_power points (9.31 DPS) | yes | General's Chain Gloves (16571, -0.27 DPS) [vendor]; General's Chain Vices (231575, -0.27 DPS) [vendor]; Hero's Gauntlets (8305, -2.97 DPS, sim-verified) [dungeon] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 84.0 ranged_attack_power points (7.19 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Stalker's Mail Belt (252588, -0.81 DPS) [crafted]; Legionnaire's Chain Girdle (16529, -1.08 DPS) [pvp] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | 111.4 ranged_attack_power points (9.54 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -1.62 DPS, sim-verified) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 83.8 ranged_attack_power points (7.18 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -0.56 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 49.0 ranged_attack_power points (4.20 DPS) | yes | Glowing Crystal Ring (18402, -0.94 DPS) [quest]; Band of the Hierophant (13096, -1.48 DPS) [world_drop]; Frigid Ring (18679, -1.48 DPS) [world] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.9 ranged_attack_power points (3.68 DPS) | yes | Glowing Crystal Ring (18402, -0.42 DPS) [quest]; Band of the Hierophant (13096, -0.96 DPS) [world_drop]; Frigid Ring (18679, -0.96 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (311.8 DPS) | yes | Uther's Strength (11302, -8.86 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (311.8 DPS) | yes | Molten Heart of the Mountain (249470, -1.26 DPS, sim-verified) [crafted]; Uther's Strength (11302, -2.16 DPS) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (311.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -12.14 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (311.8 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -4.87 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Timmy's Galoshes; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1785, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 681.1. Weights run: 2.7s. Verify run: 5.9s. 1785 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.231 ± 0.020, crit=2.259 ± 0.042 per rating point (14 rating = 1%, 31.625 per %), hit=3.754 ± 0.143 per rating point (10 rating = 1%, 37.542 per %), melee_haste=18.099 ± 1.896, intellect=2.656 ± 0.054, spell_power=0.502 ± 0.002

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 200.4 ranged_attack_power points (32.34 DPS) | yes | Warlord's Chain Helmet (16566, -4.37 DPS) [vendor]; Warlord's Chain Helm (231571, -4.37 DPS) [vendor]; Outlaw's Collar (279253, -12.58 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.0 ranged_attack_power points (6.78 DPS) | yes | Sunprism Pendant (20649, -0.52 DPS) [quest]; Horizon Choker (13085, -0.78 DPS) [world_drop]; Archlight Talisman (15856, -1.52 DPS) [quest] |
| shoulder | Spaulders of the Unseen (13116) | World drop [world_drop] | 81.7 ranged_attack_power points (13.19 DPS) | yes | Shadowcraft Spaulders (16708, +0.00 DPS) [dungeon]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 68.5 ranged_attack_power points (11.05 DPS) | yes | Dark Phantom Cape (13122, -3.89 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.79 DPS) [world_drop]; Cloak of Flames (3475, -4.99 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 188.7 ranged_attack_power points (30.46 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.48 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.48 DPS) [vendor]; Icemail Jerkin (1981, -14.67 DPS, sim-verified) [world_drop] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 57.0 ranged_attack_power points (9.20 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.33 DPS) [pvp]; Beaststalker's Bindings (16681, -1.38 DPS) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 105.0 ranged_attack_power points (16.95 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Hero's Gauntlets (8305, -5.42 DPS, sim-verified) [dungeon] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 80.3 ranged_attack_power points (12.96 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Stalker's Mail Belt (252588, -1.38 DPS) [crafted]; Legionnaire's Chain Girdle (16529, -1.81 DPS) [pvp] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | 110.6 ranged_attack_power points (17.84 DPS) | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS) [vendor]; General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 81.5 ranged_attack_power points (13.15 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.38 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.5 ranged_attack_power points (9.29 DPS) | yes | Ring of the Underwood (2951, -4.07 DPS) [world_drop]; Glowing Crystal Ring (18402, -4.14 DPS) [quest]; Band of the Penitent (13217, -4.18 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.2 ranged_attack_power points (7.30 DPS) | yes | Glowing Crystal Ring (18402, -2.16 DPS) [quest]; Band of the Penitent (13217, -2.20 DPS) [quest]; Ring of the Underwood (2951, -2.64 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (681.1 DPS) | yes | Uther's Strength (11302, -17.31 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (681.1 DPS) | yes | Uther's Strength (11302, -4.97 DPS) [world_drop]; Molten Heart of the Mountain (249470, -5.02 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (681.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -23.64 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (681.1 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -9.13 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Skibi's Pendant; shoulder: Spaulders of the Unseen; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Timmy's Galoshes; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1785, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

