# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 102.3. Weights run: 1.5s. Verify run: 1.3s. 233 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.416 ± 0.015, crit=0.929 ± 0.016 per rating point (14 rating = 1%, 13.006 per %), hit=0.957 ± 0.034 per rating point (10 rating = 1%, 9.570 per %), melee_haste=5.167 ± 0.683

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.3 ranged_attack_power points (2.01 DPS) | yes | Red Winter Hat (21524, -2.09 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.8 ranged_attack_power points (0.50 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.26 DPS) | yes | Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.5 ranged_attack_power points (1.51 DPS) | yes | Cape of the Brotherhood (5193, -0.38 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.50 DPS) [world_drop]; Bristlebark Cape (14571, -0.50 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 37.6 ranged_attack_power points (2.76 DPS) | yes | Brawler's Leather Armor (252490, -0.76 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -1.00 DPS) [crafted]; Dark Leather Tunic (2317, -1.26 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.26 DPS) | yes | Wolf Bracers (4794, -0.25 DPS) [vendor]; Bravo's Armbands (270015, -0.25 DPS) [quest]; Bristlebark Bindings (14569, -0.50 DPS) [world_drop] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.50 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.50 DPS) [crafted]; Serpent Gloves (5970, -1.94 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.32 DPS) | yes | Dusty Belt (279897, +0.00 DPS, sim-verified) [quest]; Deviate Scale Belt (6468, -0.07 DPS) [crafted]; Dark Leather Belt (4249, -0.32 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 30.7 ranged_attack_power points (2.26 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.25 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.25 DPS) [crafted]; Feet of the Lynx (1121, -1.54 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.8 ranged_attack_power points (0.50 DPS) | yes | - |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.4 ranged_attack_power points (0.25 DPS) | yes | Ring of the Shadow (1462, -0.52 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 33.3 ranged_attack_power points (2.45 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Lupine Axe (1220, -0.94 DPS) [world]; Piercing Axe (6094, -0.94 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.21 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.21 DPS) [crafted]; Venomstrike (6469, -3.35 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 233, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 126.6. Weights run: 1.5s. Verify run: 1.6s. 440 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.287 ± 0.016, crit=1.233 ± 0.023 per rating point (14 rating = 1%, 17.263 per %), hit=1.039 ± 0.047 per rating point (10 rating = 1%, 10.389 per %), melee_haste=not significant (-3.739 ± 0.983)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 32.9 ranged_attack_power points (2.35 DPS) | yes | Brawler's Leather Hood (252504, -0.54 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.70 DPS) [world]; Resilient Cap (14401, -1.88 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 ranged_attack_power points (0.94 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.2 ranged_attack_power points (2.58 DPS) | yes | Mantle of Thieves (2264, -0.23 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.94 DPS) [crafted]; Insignia Mantle (4721, -0.94 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.3 ranged_attack_power points (1.88 DPS) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Fenrus' Hide (6340, -0.47 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.47 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 46.0 ranged_attack_power points (3.29 DPS) | yes | Tunic of Westfall (2041, -0.78 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -1.17 DPS) [dungeon]; Green Leather Armor (4255, -1.41 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.7 ranged_attack_power points (1.41 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.23 DPS) [world]; Forest Leather Bracers (3202, -0.23 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.6 ranged_attack_power points (2.11 DPS) | yes | Pilferer's Gloves (7358, -0.51 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.70 DPS) [dungeon]; Insignia Gloves (6408, -0.70 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (126.6 DPS) | yes | Belt of Vindication (3562, -0.47 DPS) [quest]; Green Leather Belt (4257, -0.70 DPS) [crafted]; Skulker's Leather Belt (252520, -1.50 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 46.0 ranged_attack_power points (3.29 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.23 DPS) [crafted]; Insignia Leggings (4054, -1.17 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 26.3 ranged_attack_power points (1.88 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, -0.47 DPS) [dungeon]; Footpads of the Fang (10411, -0.47 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.7 ranged_attack_power points (1.41 DPS) | yes | Defias Renegade Ring (1076, -0.23 DPS) [dungeon]; Seal of Wrynn (2933, -0.70 DPS) [quest]; Thunderbrow Ring (13097, -0.70 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 16.4 ranged_attack_power points (1.17 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Seal of Wrynn (2933, -0.47 DPS) [quest]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 19.7 ranged_attack_power points (1.41 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 17.1 ranged_attack_power points (1.23 DPS) | yes | The Butcher (8226, -1.37 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 252.1 ranged_attack_power points (18.02 DPS) | yes | Ironweaver (13137, -0.51 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.18 DPS) [world_drop]; Chesterfall Musket (7729, -1.49 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 440, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 163.5. Weights run: 1.6s. Verify run: 1.8s. 696 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.130 ± 0.016, crit=1.376 ± 0.027 per rating point (14 rating = 1%, 19.263 per %), hit=1.241 ± 0.053 per rating point (10 rating = 1%, 12.410 per %), melee_haste=6.521 ± 1.021

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 40.7 ranged_attack_power points (3.08 DPS) | yes | Nightscape Headband (8176, -0.24 DPS) [crafted]; Guard's Chain Helm (250499, -0.24 DPS) [crafted]; Skullsplitter Helm (1624, -0.47 DPS) [world] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.5 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, +0.00 DPS, sim-verified) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest]; Amberglow Talisman (10824, -0.71 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 34.4 ranged_attack_power points (2.61 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.0 ranged_attack_power points (1.89 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 75.1 ranged_attack_power points (5.68 DPS) | yes | Wolffear Harness (13110, -1.40 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.13 DPS) [crafted]; Tough Scorpid Breastplate (8203, -2.13 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 25.0 ranged_attack_power points (1.89 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.38 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 39.3 ranged_attack_power points (2.97 DPS) | yes | Dragonscale Gauntlets (8347, -0.09 DPS) [crafted]; Gauntlets of Divinity (7724, -0.55 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.60 DPS) [crafted] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 31.3 ranged_attack_power points (2.37 DPS) | yes | Blackforge Girdle (6425, -0.24 DPS) [dungeon]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 65.7 ranged_attack_power points (4.97 DPS) | yes | Triprunner Dungarees (9624, -0.82 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.66 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.66 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 40.7 ranged_attack_power points (3.08 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.47 DPS) [dungeon]; Dusky Boots (7390, -0.47 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 31.3 ranged_attack_power points (2.37 DPS) | yes | Disengagement Ring (276202, -0.47 DPS) [vendor]; Field Researcher's Loop (281634, -0.71 DPS) [quest]; Mark of Kern (2262, -0.86 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 28.2 ranged_attack_power points (2.13 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Field Researcher's Loop (281634, -0.47 DPS) [quest]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (163.5 DPS) | yes | Headhunting Spear (1522, -1.23 DPS) [world]; Armor Piercer (6679, -1.23 DPS) [dungeon]; Gut Ripper (2164, -4.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (26.48 DPS) | yes | Swiftwind (13038, -2.31 DPS) [world_drop]; Baelog's Shortbow (9400, -2.53 DPS) [dungeon]; The Silencer (13138, -3.04 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 216.2. Weights run: 1.6s. Verify run: 2.4s. 885 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.045 ± 0.016, crit=1.558 ± 0.029 per rating point (14 rating = 1%, 21.806 per %), hit=1.653 ± 0.068 per rating point (10 rating = 1%, 16.535 per %), melee_haste=11.545 ± 1.106

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 51.8 ranged_attack_power points (4.23 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Lordrec Helmet (10741, -0.25 DPS) [quest]; Sprightring Helm (17776, -0.50 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.6 ranged_attack_power points (3.23 DPS) | yes | Kaleidoscope Chain (13084, -1.71 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -2.25 DPS) [world]; Erudite's Amulet (277204, -2.74 DPS) [quest] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 36.5 ranged_attack_power points (2.98 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.25 DPS) [world_drop]; Nightscape Shoulders (8192, -0.96 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 45.7 ranged_attack_power points (3.73 DPS) | yes | Serpentskin Cloak (8259, -0.99 DPS) [world_drop]; Nightscape Cloak (8195, -1.24 DPS) [crafted]; Blackflame Cape (13109, -1.64 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 73.1 ranged_attack_power points (5.97 DPS) | yes | Blazewind Breastplate (11193, -0.25 DPS) [quest]; Knight's Chain Armor (220828, -0.46 DPS) [vendor]; Charred Leather Tunic (19127, -1.49 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 45.7 ranged_attack_power points (3.73 DPS) | yes | Bloodlust Bracelets (14807, -0.84 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.99 DPS) [crafted]; Pridelord Bands (14672, -1.24 DPS) [world_drop] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 48.7 ranged_attack_power points (3.98 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.50 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.50 DPS) [crafted]; Gloves of Holy Might (867, -0.56 DPS) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 45.7 ranged_attack_power points (3.73 DPS) | yes | Stalker's Mail Belt (252588, -0.25 DPS) [crafted]; Seared Mail Girdle (19125, -0.75 DPS) [quest]; Skulker's Leather Waistguard (252474, -3.89 DPS, sim-verified) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 64.4 ranged_attack_power points (5.26 DPS) | yes | Keeper's Woolies (14668, -0.54 DPS) [world_drop]; Oilskin Leggings (9414, -0.79 DPS) [dungeon]; Basilisk Hide Pants (1718, -3.36 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 60.9 ranged_attack_power points (4.97 DPS) | yes | Fleetfoot Greaves (11627, -0.25 DPS) [dungeon]; Elven Chain Boots (13125, -0.50 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -1.24 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.6 ranged_attack_power points (3.48 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Disengagement Ring (276202, -1.49 DPS) [vendor]; Field Researcher's Loop (281634, -1.74 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.4 ranged_attack_power points (2.49 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; Field Researcher's Loop (281634, -0.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (216.2 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (216.2 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (216.2 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -2.63 DPS, sim-verified) [world_drop] |
| off_hand | Widowmaker (4091) (or Blue Glittering Axe (7942), Dazzling Mithril Rapier (7944), Nordic Longshank (9401), Speedsteel Rapier (13034), Mithril Stiletto (250611)) | World drop [world_drop] | 24.4 ranged_attack_power points (1.99 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Dazzling Mithril Rapier (7944, +0.00 DPS) [crafted]; Nordic Longshank (9401, +0.00 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (216.2 DPS) | yes | Gryphonwing Long Bow (13022, -0.74 DPS) [dungeon]; Houndmaster's Bow (11628, -3.02 DPS) [dungeon]; Dark Iron Rifle (16004, -5.93 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Widowmaker; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 885, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 278.8. Weights run: 1.6s. Verify run: 5.9s. 1755 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.062 ± 0.019, crit=2.040 ± 0.040 per rating point (14 rating = 1%, 28.560 per %), hit=2.008 ± 0.095 per rating point (10 rating = 1%, 20.083 per %), melee_haste=12.948 ± 1.586

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (278.8 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -9.71 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.8 ranged_attack_power points (3.22 DPS) | yes | Sunprism Pendant (20649, -0.25 DPS) [quest]; Kaleidoscope Chain (13084, -2.23 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.25 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 71.4 ranged_attack_power points (5.78 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS, sim-verified) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 65.9 ranged_attack_power points (5.34 DPS) | yes | Dark Phantom Cape (13122, -1.11 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.36 DPS) [world_drop]; Deep River Cloak (15789, -2.61 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 161.9 ranged_attack_power points (13.11 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.37 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.37 DPS) [vendor]; Shadowcraft Tunic (16721, -5.02 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) | Stratholme: Fleshflayer Ghoul [dungeon] | sim-verified (278.8 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -2.83 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 64.3 ranged_attack_power points (5.21 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -5.62 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 45.9 ranged_attack_power points (3.72 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Knight-Captain's Chain Girdle (16400, -0.25 DPS) [pvp]; Wicked Leather Belt (15088, -4.51 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (278.8 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -6.56 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (278.8 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -3.03 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.9 ranged_attack_power points (3.47 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Disengagement Ring (276202, -1.49 DPS) [vendor]; Field Researcher's Loop (281634, -1.74 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.6 ranged_attack_power points (2.48 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; Field Researcher's Loop (281634, -0.74 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (278.8 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (278.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (278.8 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -10.52 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (278.8 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -9.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 618.1. Weights run: 1.7s. Verify run: 6.2s. 1755 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.250 ± 0.021, crit=2.280 ± 0.042 per rating point (14 rating = 1%, 31.921 per %), hit=2.868 ± 0.139 per rating point (10 rating = 1%, 28.682 per %), melee_haste=15.329 ± 2.025

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (618.1 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -12.99 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.3 ranged_attack_power points (6.46 DPS) | yes | Sunprism Pendant (20649, -0.50 DPS) [quest]; Kaleidoscope Chain (13084, -4.47 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.63 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 77.2 ranged_attack_power points (11.81 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 68.8 ranged_attack_power points (10.51 DPS) | yes | Dark Phantom Cape (13122, -3.67 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.55 DPS) [world_drop]; Deep River Cloak (15789, -5.05 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 180.9 ranged_attack_power points (27.66 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.88 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.88 DPS) [vendor]; Shadowcraft Tunic (16721, -13.56 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 48.8 ranged_attack_power points (7.46 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -5.48 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 68.3 ranged_attack_power points (10.44 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -6.54 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 48.8 ranged_attack_power points (7.46 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Knight-Captain's Chain Girdle (16400, -0.50 DPS) [pvp]; Wicked Leather Belt (15088, -5.36 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (618.1 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -8.95 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (618.1 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.47 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.5 ranged_attack_power points (6.96 DPS) | yes | Falcon's Hook (7552, -2.49 DPS) [dungeon]; Disengagement Ring (276202, -2.98 DPS) [vendor]; Field Researcher's Loop (281634, -3.48 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 32.5 ranged_attack_power points (4.97 DPS) | yes | Falcon's Hook (7552, -0.50 DPS) [dungeon]; Disengagement Ring (276202, -0.99 DPS) [vendor]; Field Researcher's Loop (281634, -1.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (618.1 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (618.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (618.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -23.01 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (618.1 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -18.34 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 102.6. Weights run: 1.5s. Verify run: 1.3s. 229 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.416 ± 0.015, crit=0.929 ± 0.016 per rating point (14 rating = 1%, 13.006 per %), hit=0.957 ± 0.034 per rating point (10 rating = 1%, 9.570 per %), melee_haste=5.167 ± 0.683

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.3 ranged_attack_power points (2.01 DPS) | yes | Red Winter Hat (21524, -2.14 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.8 ranged_attack_power points (0.50 DPS) | yes | Roadwatcher's Confidence (281265, -0.42 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.26 DPS) | yes | Slime-encrusted Pads (6461, -0.72 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.5 ranged_attack_power points (1.51 DPS) | yes | Cape of the Brotherhood (5193, -0.35 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.50 DPS) [world_drop]; Bristlebark Cape (14571, -0.50 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 23.9 ranged_attack_power points (1.76 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.25 DPS) [crafted]; Prospector's Chestpiece (14562, -0.25 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.26 DPS) | yes | Wolf Bracers (4794, -0.29 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.50 DPS) [world_drop]; Spare Part Bindings (279875, -0.50 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (102.6 DPS) | yes | Forest Leather Gloves (3058, -0.50 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.50 DPS) [crafted]; Serpent Gloves (5970, -1.26 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.32 DPS) | yes | Deviate Scale Belt (6468, -0.07 DPS) [crafted]; Guardsman Belt (3429, -0.32 DPS) [world]; Dark Leather Belt (4249, -0.32 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 30.7 ranged_attack_power points (2.26 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.25 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.3 ranged_attack_power points (2.01 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.50 DPS) [dungeon]; Gray Leather Boots (2315, -0.75 DPS) [crafted] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 10.2 ranged_attack_power points (0.75 DPS) | yes | The 1 Ring (8350, -0.50 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.8 ranged_attack_power points (0.50 DPS) | yes | The 1 Ring (8350, -0.81 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 33.3 ranged_attack_power points (2.45 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.69 DPS) [quest]; Lupine Axe (1220, -0.94 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.21 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.21 DPS) [crafted]; Venomstrike (6469, -3.35 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Bounty Hunter's Ring; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 229, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 131.3. Weights run: 1.5s. Verify run: 1.7s. 437 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.287 ± 0.016, crit=1.233 ± 0.023 per rating point (14 rating = 1%, 17.263 per %), hit=1.039 ± 0.047 per rating point (10 rating = 1%, 10.389 per %), melee_haste=not significant (-3.739 ± 0.983)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 32.9 ranged_attack_power points (2.35 DPS) | yes | Brawler's Leather Hood (252504, -0.55 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.70 DPS) [world]; Resilient Cap (14401, -1.88 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 ranged_attack_power points (0.94 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.2 ranged_attack_power points (2.58 DPS) | yes | Mantle of Thieves (2264, -0.23 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.94 DPS) [crafted]; Insignia Mantle (4721, -0.94 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.3 ranged_attack_power points (1.88 DPS) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Fenrus' Hide (6340, -0.47 DPS) [dungeon]; Swiftrunner Cape (6745, -0.47 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Panther Armor (6670, +0.00 DPS) [quest]; Green Leather Armor (4255, -0.23 DPS) [crafted]; Dusky Leather Armor (7374, -3.17 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.7 ranged_attack_power points (1.41 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.23 DPS) [world_drop]; Loamflake Bracers (15462, -0.23 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.6 ranged_attack_power points (2.11 DPS) | yes | Braced Handguards (6784, -0.47 DPS) [quest]; Pilferer's Gloves (7358, -0.48 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.70 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Deftkin Belt (16659, -0.32 DPS) [quest]; Larcenist's Belt (279846, -0.70 DPS) [quest]; Skulker's Leather Belt (252520, -1.75 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 46.0 ranged_attack_power points (3.29 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.23 DPS) [crafted]; Insignia Leggings (4054, -1.17 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 26.3 ranged_attack_power points (1.88 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.23 DPS) [quest] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.7 ranged_attack_power points (1.41 DPS) | yes | Defias Renegade Ring (1076, -0.23 DPS) [dungeon]; Band of the Fist (17694, -0.47 DPS) [quest]; Thunderbrow Ring (13097, -0.70 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 16.4 ranged_attack_power points (1.17 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Band of the Fist (17694, -0.23 DPS) [quest]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 19.7 ranged_attack_power points (1.41 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 17.1 ranged_attack_power points (1.23 DPS) | yes | Wingblade (6504, -2.19 DPS, sim-verified) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 252.1 ranged_attack_power points (18.02 DPS) | yes | Ironweaver (13137, -0.70 DPS) [world_drop]; Crystalpine Stinger (13037, -1.18 DPS) [world_drop]; Chesterfall Musket (7729, -1.49 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 437, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 164.2. Weights run: 1.6s. Verify run: 1.7s. 684 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.130 ± 0.016, crit=1.376 ± 0.027 per rating point (14 rating = 1%, 19.263 per %), hit=1.241 ± 0.053 per rating point (10 rating = 1%, 12.410 per %), melee_haste=6.521 ± 1.021

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 40.7 ranged_attack_power points (3.08 DPS) | yes | Guard's Chain Helm (250499, -0.24 DPS) [crafted]; Skullsplitter Helm (1624, -0.47 DPS) [world]; Nightscape Headband (8176, -1.59 DPS, sim-verified) [crafted] |
| neck | Ethereal Talisman (4430) (or Kaleidoscope Chain (13084)) | The Crown of Will [quest] | 12.5 ranged_attack_power points (0.95 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.04 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 34.4 ranged_attack_power points (2.61 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.0 ranged_attack_power points (1.89 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 75.1 ranged_attack_power points (5.68 DPS) | yes | Wolffear Harness (13110, -1.85 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.13 DPS) [crafted]; Tough Scorpid Breastplate (8203, -2.13 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 25.0 ranged_attack_power points (1.89 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.38 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 39.3 ranged_attack_power points (2.97 DPS) | yes | Dragonscale Gauntlets (8347, -0.09 DPS) [crafted]; Gauntlets of Divinity (7724, -0.55 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.60 DPS) [crafted] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 31.3 ranged_attack_power points (2.37 DPS) | yes | Blackforge Girdle (6425, -0.24 DPS) [dungeon]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 65.7 ranged_attack_power points (4.97 DPS) | yes | Triprunner Dungarees (9624, -1.28 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.66 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.66 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 40.7 ranged_attack_power points (3.08 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.47 DPS) [dungeon]; Dusky Boots (7390, -0.47 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 31.3 ranged_attack_power points (2.37 DPS) | yes | Disengagement Ring (276202, -0.47 DPS) [vendor]; Field Researcher's Loop (281634, -0.71 DPS) [quest]; Mark of Kern (2262, -0.86 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 28.2 ranged_attack_power points (2.13 DPS) | yes | Disengagement Ring (276202, +0.00 DPS, sim-verified) [vendor]; Field Researcher's Loop (281634, -0.47 DPS) [quest]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (164.2 DPS) | yes | Headhunting Spear (1522, -1.23 DPS) [world]; Armor Piercer (6679, -1.23 DPS) [dungeon]; Gut Ripper (2164, -4.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (26.48 DPS) | yes | Swiftwind (13038, -2.31 DPS) [world_drop]; Baelog's Shortbow (9400, -2.53 DPS) [dungeon]; The Silencer (13138, -3.28 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Ethereal Talisman; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 226.3. Weights run: 1.6s. Verify run: 2.2s. 874 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.045 ± 0.016, crit=1.558 ± 0.029 per rating point (14 rating = 1%, 21.806 per %), hit=1.653 ± 0.068 per rating point (10 rating = 1%, 16.535 per %), melee_haste=11.545 ± 1.106

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 51.8 ranged_attack_power points (4.23 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.50 DPS) [quest]; Tough Scorpid Helm (8208, -0.75 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.6 ranged_attack_power points (3.23 DPS) | yes | Woven Ivy Necklace (19159, -1.90 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -2.24 DPS) [quest]; Kaleidoscope Chain (13084, -2.24 DPS) [world_drop] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 36.5 ranged_attack_power points (2.98 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.25 DPS) [world_drop]; Nightscape Shoulders (8192, -1.03 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 45.7 ranged_attack_power points (3.73 DPS) | yes | Serpentskin Cloak (8259, -0.99 DPS) [world_drop]; Nightscape Cloak (8195, -1.24 DPS) [crafted]; Blackflame Cape (13109, -1.74 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 73.1 ranged_attack_power points (5.97 DPS) | yes | Blazewind Breastplate (11193, -0.25 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.46 DPS) [vendor]; Charred Leather Tunic (19127, -1.49 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 45.7 ranged_attack_power points (3.73 DPS) | yes | Bloodlust Bracelets (14807, -0.92 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.99 DPS) [crafted]; Pridelord Bands (14672, -1.24 DPS) [world_drop] |
| hands | First Sergeant's Chain Gauntlets (220830) | PvP rank 9 · First Sergeant · Horde [vendor] | 48.7 ranged_attack_power points (3.98 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.50 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.50 DPS) [crafted]; Gloves of Holy Might (867, -0.56 DPS) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 45.7 ranged_attack_power points (3.73 DPS) | yes | Stalker's Mail Belt (252588, -0.25 DPS) [crafted]; Seared Mail Girdle (19125, -0.75 DPS) [quest]; Skulker's Leather Waistguard (252474, -4.35 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 64.4 ranged_attack_power points (5.26 DPS) | yes | Keeper's Woolies (14668, -0.54 DPS) [world_drop]; Oilskin Leggings (9414, -0.79 DPS) [dungeon]; Basilisk Hide Pants (1718, -3.23 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 60.9 ranged_attack_power points (4.97 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.50 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -1.24 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.6 ranged_attack_power points (3.48 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Disengagement Ring (276202, -1.49 DPS) [vendor]; White Bone Band (11862, -1.52 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.4 ranged_attack_power points (2.49 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; White Bone Band (11862, -0.53 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -8.85 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -0.92 DPS, sim-verified) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Widowmaker (4091) | World drop [world_drop] | sim-verified (226.3 DPS) | yes | White Bone Shredder (11863, -0.25 DPS) [quest]; Force of the Hippogryph (9684, -2.90 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.74 DPS) [dungeon]; Houndmaster's Bow (11628, -3.02 DPS) [dungeon]; Dark Iron Rifle (16004, -6.06 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: First Sergeant's Chain Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Widowmaker; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 874, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 290.9. Weights run: 1.6s. Verify run: 5.8s. 1752 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.062 ± 0.019, crit=2.040 ± 0.040 per rating point (14 rating = 1%, 28.560 per %), hit=2.008 ± 0.095 per rating point (10 rating = 1%, 20.083 per %), melee_haste=12.948 ± 1.586

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (290.9 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -11.00 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.8 ranged_attack_power points (3.22 DPS) | yes | Sunprism Pendant (20649, -0.25 DPS) [quest]; Woven Ivy Necklace (19159, -0.99 DPS) [quest]; Ethereal Talisman (4430, -2.23 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 71.4 ranged_attack_power points (5.78 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 65.9 ranged_attack_power points (5.34 DPS) | yes | Dark Phantom Cape (13122, -2.00 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.36 DPS) [world_drop]; Deep River Cloak (15789, -2.61 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 161.9 ranged_attack_power points (13.11 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.37 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.37 DPS) [vendor]; Shadowcraft Tunic (16721, -5.81 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) | Stratholme: Fleshflayer Ghoul [dungeon] | sim-verified (290.9 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -7.51 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (290.9 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Mourning Handguards (279257, -5.26 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (290.9 DPS) | yes | Wicked Leather Belt (15088, +0.00 DPS) [crafted]; General's Chain Girdle (16572, +0.00 DPS) [pvp]; Sagebrush Girdle (17778, -6.02 DPS, sim-verified) [quest] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (290.9 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -10.87 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (290.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -7.12 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.9 ranged_attack_power points (3.47 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Disengagement Ring (276202, -1.49 DPS) [vendor]; White Bone Band (11862, -1.53 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.6 ranged_attack_power points (2.48 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; White Bone Band (11862, -0.54 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (290.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (290.9 DPS) | yes | Molten Heart of the Mountain (249470, -2.65 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (290.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -11.87 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (290.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.36 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1752, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 636.3. Weights run: 1.7s. Verify run: 5.9s. 1752 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.250 ± 0.021, crit=2.280 ± 0.042 per rating point (14 rating = 1%, 31.921 per %), hit=2.868 ± 0.139 per rating point (10 rating = 1%, 28.682 per %), melee_haste=15.329 ± 2.025

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (636.3 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -13.42 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.3 ranged_attack_power points (6.46 DPS) | yes | Sunprism Pendant (20649, -0.50 DPS) [quest]; Woven Ivy Necklace (19159, -1.99 DPS) [quest]; Ethereal Talisman (4430, -4.47 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 77.2 ranged_attack_power points (11.81 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 68.8 ranged_attack_power points (10.51 DPS) | yes | Dark Phantom Cape (13122, -3.67 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.55 DPS) [world_drop]; Deep River Cloak (15789, -5.05 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 180.9 ranged_attack_power points (27.66 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.88 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.88 DPS) [vendor]; Shadowcraft Tunic (16721, -13.51 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 48.8 ranged_attack_power points (7.46 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -5.50 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 68.3 ranged_attack_power points (10.44 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.14 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 48.8 ranged_attack_power points (7.46 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Legionnaire's Chain Girdle (16529, -0.50 DPS) [pvp]; Wicked Leather Belt (15088, -5.88 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (636.3 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -9.11 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (636.3 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -6.20 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.5 ranged_attack_power points (6.96 DPS) | yes | Falcon's Hook (7552, -2.49 DPS) [dungeon]; Disengagement Ring (276202, -2.98 DPS) [vendor]; White Bone Band (11862, -3.29 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 32.5 ranged_attack_power points (4.97 DPS) | yes | Falcon's Hook (7552, -0.50 DPS) [dungeon]; Disengagement Ring (276202, -0.99 DPS) [vendor]; White Bone Band (11862, -1.30 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (636.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (636.3 DPS) | yes | Molten Heart of the Mountain (249470, -4.94 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (636.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -21.89 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (636.3 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -20.26 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1752, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

