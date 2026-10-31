# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 103.3. Weights run: 1.6s. Verify run: 1.6s. 227 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.374 ± 0.014, crit=0.918 ± 0.016 per rating point (14 rating = 1%, 12.847 per %), hit=0.950 ± 0.034 per rating point (10 rating = 1%, 9.504 per %), melee_haste=5.132 ± 0.678

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.0 ranged_attack_power points (2.03 DPS) | yes | Red Winter Hat (21524, -2.09 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.7 ranged_attack_power points (0.51 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 16.9 ranged_attack_power points (1.27 DPS) | yes | Slime-encrusted Pads (6461, -0.40 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.2 ranged_attack_power points (1.52 DPS) | yes | Cape of the Brotherhood (5193, -0.39 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.51 DPS) [world_drop]; Bristlebark Cape (14571, -0.51 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 37.1 ranged_attack_power points (2.79 DPS) | yes | Brawler's Leather Armor (252490, -0.76 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -1.01 DPS) [crafted]; Dark Leather Tunic (2317, -1.27 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.9 ranged_attack_power points (1.27 DPS) | yes | Wolf Bracers (4794, -0.25 DPS) [vendor]; Bravo's Armbands (270015, -0.25 DPS) [quest]; Bristlebark Bindings (14569, -0.51 DPS) [world_drop] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.51 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.51 DPS) [crafted]; Serpent Gloves (5970, -1.95 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.35 DPS) | yes | Dusty Belt (279897, +0.00 DPS, sim-verified) [quest]; Deviate Scale Belt (6468, -0.08 DPS) [crafted]; Dark Leather Belt (4249, -0.34 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 30.4 ranged_attack_power points (2.28 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.25 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.25 DPS) [crafted]; Feet of the Lynx (1121, -1.56 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.7 ranged_attack_power points (0.51 DPS) | yes | - |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.4 ranged_attack_power points (0.25 DPS) | yes | Ring of the Shadow (1462, -0.53 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 33.0 ranged_attack_power points (2.48 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Lupine Axe (1220, -0.96 DPS) [world]; Piercing Axe (6094, -0.96 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.6 ranged_attack_power points (13.49 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.27 DPS) [crafted]; Venomstrike (6469, -3.42 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 227, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 129.3. Weights run: 1.7s. Verify run: 1.7s. 434 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.308 ± 0.017, crit=1.262 ± 0.023 per rating point (14 rating = 1%, 17.670 per %), hit=1.073 ± 0.049 per rating point (10 rating = 1%, 10.727 per %), melee_haste=not significant (-3.863 ± 1.015)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 33.1 ranged_attack_power points (2.36 DPS) | yes | Brawler's Leather Hood (252504, -0.54 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.71 DPS) [world]; Resilient Cap (14401, -1.89 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.2 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.4 ranged_attack_power points (2.60 DPS) | yes | Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.95 DPS) [crafted]; Insignia Mantle (4721, -0.95 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.5 ranged_attack_power points (1.89 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Fenrus' Hide (6340, -0.47 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.47 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 46.3 ranged_attack_power points (3.31 DPS) | yes | Tunic of Westfall (2041, -0.78 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -1.18 DPS) [dungeon]; Green Leather Armor (4255, -1.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.8 ranged_attack_power points (1.42 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.24 DPS) [world]; Forest Leather Bracers (3202, -0.24 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.8 ranged_attack_power points (2.13 DPS) | yes | Pilferer's Gloves (7358, -0.52 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.71 DPS) [dungeon]; Insignia Gloves (6408, -0.71 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (129.3 DPS) | yes | Belt of Vindication (3562, -0.47 DPS) [quest]; Green Leather Belt (4257, -0.71 DPS) [crafted]; Skulker's Leather Belt (252520, -1.52 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 46.3 ranged_attack_power points (3.31 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.24 DPS) [crafted]; Insignia Leggings (4054, -1.18 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 26.5 ranged_attack_power points (1.89 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, -0.47 DPS) [dungeon]; Footpads of the Fang (10411, -0.47 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.8 ranged_attack_power points (1.42 DPS) | yes | Defias Renegade Ring (1076, -0.24 DPS) [dungeon]; Seal of Wrynn (2933, -0.71 DPS) [quest]; Thunderbrow Ring (13097, -0.71 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 16.5 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Seal of Wrynn (2933, -0.47 DPS) [quest]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 19.8 ranged_attack_power points (1.42 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 17.2 ranged_attack_power points (1.23 DPS) | yes | The Butcher (8226, -1.37 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 252.2 ranged_attack_power points (18.02 DPS) | yes | Ironweaver (13137, -0.51 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.18 DPS) [world_drop]; Chesterfall Musket (7729, -1.50 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 167.8. Weights run: 1.7s. Verify run: 1.9s. 690 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.153 ± 0.017, crit=1.420 ± 0.028 per rating point (14 rating = 1%, 19.886 per %), hit=1.294 ± 0.055 per rating point (10 rating = 1%, 12.939 per %), melee_haste=6.810 ± 1.067

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 41.0 ranged_attack_power points (3.10 DPS) | yes | Nightscape Headband (8176, -0.24 DPS) [crafted]; Guard's Chain Helm (250499, -0.24 DPS) [crafted]; Skullsplitter Helm (1624, -0.48 DPS) [world] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.6 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, +0.00 DPS, sim-verified) [world]; Erudite's Amulet (277204, -0.48 DPS) [quest]; Amberglow Talisman (10824, -0.72 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 34.7 ranged_attack_power points (2.62 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.2 ranged_attack_power points (1.91 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 75.7 ranged_attack_power points (5.73 DPS) | yes | Wolffear Harness (13110, -1.39 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.15 DPS) [crafted]; Tough Scorpid Breastplate (8203, -2.15 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 25.2 ranged_attack_power points (1.91 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.40 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 39.9 ranged_attack_power points (3.02 DPS) | yes | Dragonscale Gauntlets (8347, -0.08 DPS) [crafted]; Gauntlets of Divinity (7724, -0.60 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.63 DPS) [crafted] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 31.5 ranged_attack_power points (2.39 DPS) | yes | Blackforge Girdle (6425, -0.24 DPS) [dungeon]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 66.2 ranged_attack_power points (5.01 DPS) | yes | Triprunner Dungarees (9624, -0.83 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.67 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.67 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 41.0 ranged_attack_power points (3.10 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.48 DPS) [dungeon]; Dusky Boots (7390, -0.48 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 31.5 ranged_attack_power points (2.39 DPS) | yes | Disengagement Ring (276202, -0.48 DPS) [vendor]; Field Researcher's Loop (281634, -0.72 DPS) [quest]; Mark of Kern (2262, -0.87 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 28.4 ranged_attack_power points (2.15 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Field Researcher's Loop (281634, -0.48 DPS) [quest]; Mark of Kern (2262, -0.63 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (167.8 DPS) | yes | Headhunting Spear (1522, -1.23 DPS) [world]; Armor Piercer (6679, -1.23 DPS) [dungeon]; Gut Ripper (2164, -4.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (26.48 DPS) | yes | Swiftwind (13038, -2.30 DPS) [world_drop]; Baelog's Shortbow (9400, -2.53 DPS) [dungeon]; The Silencer (13138, -3.02 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 227.7. Weights run: 1.8s. Verify run: 4.0s. 879 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.072 ± 0.017, crit=1.626 ± 0.031 per rating point (14 rating = 1%, 22.759 per %), hit=1.748 ± 0.072 per rating point (10 rating = 1%, 17.484 per %), melee_haste=12.189 ± 1.173

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 52.2 ranged_attack_power points (4.26 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Lordrec Helmet (10741, -0.25 DPS) [quest]; Sprightring Helm (17776, -0.50 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.9 ranged_attack_power points (3.26 DPS) | yes | Archmagister's Faceted Pendant (273049, -2.28 DPS) [world]; Kaleidoscope Chain (13084, -2.38 DPS, sim-verified) [world_drop]; Erudite's Amulet (277204, -2.76 DPS) [quest] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 36.9 ranged_attack_power points (3.01 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.25 DPS) [world_drop]; Nightscape Shoulders (8192, -1.42 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 46.1 ranged_attack_power points (3.76 DPS) | yes | Serpentskin Cloak (8259, -1.00 DPS) [world_drop]; Nightscape Cloak (8195, -1.25 DPS) [crafted]; Blackflame Cape (13109, -1.70 DPS, sim-verified) [world_drop] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-verified (227.7 DPS) | yes | Blazewind Breastplate (11193, +0.00 DPS) [quest]; Charred Leather Tunic (19127, -1.11 DPS) [quest]; Icemail Jerkin (1981, -3.32 DPS, sim-verified) [world_drop] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.1 ranged_attack_power points (3.76 DPS) | yes | Wicked Leather Bracers (15084, -1.00 DPS) [crafted]; Bloodlust Bracelets (14807, -1.11 DPS, sim-verified) [world_drop]; Pridelord Bands (14672, -1.25 DPS) [world_drop] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 49.1 ranged_attack_power points (4.01 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.50 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.50 DPS) [crafted]; Gloves of Holy Might (867, -0.52 DPS) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 46.1 ranged_attack_power points (3.76 DPS) | yes | Stalker's Mail Belt (252588, -0.25 DPS) [crafted]; Seared Mail Girdle (19125, -0.75 DPS) [quest]; Skulker's Leather Waistguard (252474, -4.21 DPS, sim-verified) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 65.8 ranged_attack_power points (5.37 DPS) | yes | Keeper's Woolies (14668, -0.60 DPS) [world_drop]; Oilskin Leggings (9414, -0.86 DPS) [dungeon]; Basilisk Hide Pants (1718, -3.41 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 61.4 ranged_attack_power points (5.02 DPS) | yes | Fleetfoot Greaves (11627, -0.25 DPS) [dungeon]; Elven Chain Boots (13125, -0.50 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -1.25 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.0 ranged_attack_power points (3.51 DPS) | yes | Falcon's Hook (7552, -1.25 DPS) [dungeon]; Disengagement Ring (276202, -1.51 DPS) [vendor]; Field Researcher's Loop (281634, -1.76 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.7 ranged_attack_power points (2.51 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; Field Researcher's Loop (281634, -0.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (227.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (227.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (227.7 DPS) | yes | Steel Spear (250605, -0.21 DPS) [crafted]; Grimlok's Charge (9416, -0.52 DPS) [dungeon]; Hanzo Sword (8190, -6.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (227.7 DPS) | yes | Gryphonwing Long Bow (13022, -0.75 DPS) [dungeon]; Houndmaster's Bow (11628, -3.05 DPS) [dungeon]; Dark Iron Rifle (16004, -6.09 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 284.7. Weights run: 1.7s. Verify run: 6.4s. 1749 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.077 ± 0.019, crit=2.087 ± 0.041 per rating point (14 rating = 1%, 29.218 per %), hit=2.182 ± 0.105 per rating point (10 rating = 1%, 21.823 per %), melee_haste=15.069 ± 1.652

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (284.7 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -10.36 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.0 ranged_attack_power points (3.24 DPS) | yes | Sunprism Pendant (20649, -0.25 DPS) [quest]; Kaleidoscope Chain (13084, -2.24 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.27 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 72.1 ranged_attack_power points (5.83 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS, sim-verified) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.2 ranged_attack_power points (5.36 DPS) | yes | Dark Phantom Cape (13122, -1.62 DPS) [world_drop]; Blackflame Cape (13109, -2.37 DPS) [world_drop]; Deep River Cloak (15789, -2.62 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 164.9 ranged_attack_power points (13.35 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.51 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.51 DPS) [vendor]; Shadowcraft Tunic (16721, -4.96 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) | Stratholme: Fleshflayer Ghoul [dungeon] | sim-verified (284.7 DPS) | yes | Knight-Captain's Chain Armguards (16402, +0.00 DPS) [pvp]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -2.83 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 64.6 ranged_attack_power points (5.23 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -5.17 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 46.2 ranged_attack_power points (3.74 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Knight-Captain's Chain Girdle (16400, -0.25 DPS) [pvp]; Wicked Leather Belt (15088, -4.60 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (284.7 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -6.38 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (284.7 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -3.51 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.1 ranged_attack_power points (3.49 DPS) | yes | Falcon's Hook (7552, -1.25 DPS) [dungeon]; Disengagement Ring (276202, -1.49 DPS) [vendor]; Field Researcher's Loop (281634, -1.74 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.8 ranged_attack_power points (2.49 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; Field Researcher's Loop (281634, -0.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (284.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (284.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (284.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -10.83 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (284.7 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.18 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 629.1. Weights run: 1.7s. Verify run: 6.4s. 1749 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.265 ± 0.021, crit=2.324 ± 0.042 per rating point (14 rating = 1%, 32.532 per %), hit=2.931 ± 0.142 per rating point (10 rating = 1%, 29.307 per %), melee_haste=15.726 ± 2.080

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (629.1 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -13.04 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.4 ranged_attack_power points (6.49 DPS) | yes | Sunprism Pendant (20649, -0.50 DPS) [quest]; Kaleidoscope Chain (13084, -4.49 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.66 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 77.7 ranged_attack_power points (11.88 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 69.0 ranged_attack_power points (10.55 DPS) | yes | Dark Phantom Cape (13122, -3.67 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.56 DPS) [world_drop]; Deep River Cloak (15789, -5.06 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 182.6 ranged_attack_power points (27.93 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.98 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.98 DPS) [vendor]; Shadowcraft Tunic (16721, -13.72 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 49.0 ranged_attack_power points (7.49 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -5.48 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 68.6 ranged_attack_power points (10.48 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -6.52 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 49.0 ranged_attack_power points (7.49 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Knight-Captain's Chain Girdle (16400, -0.50 DPS) [pvp]; Wicked Leather Belt (15088, -5.39 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (629.1 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -8.97 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (629.1 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.43 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.7 ranged_attack_power points (6.99 DPS) | yes | Falcon's Hook (7552, -2.50 DPS) [dungeon]; Disengagement Ring (276202, -3.00 DPS) [vendor]; Field Researcher's Loop (281634, -3.49 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 32.6 ranged_attack_power points (4.99 DPS) | yes | Falcon's Hook (7552, -0.50 DPS) [dungeon]; Disengagement Ring (276202, -1.00 DPS) [vendor]; Field Researcher's Loop (281634, -1.50 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (629.1 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (629.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (629.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -23.19 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (629.1 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -18.36 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 103.6. Weights run: 1.6s. Verify run: 1.5s. 223 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.374 ± 0.014, crit=0.918 ± 0.016 per rating point (14 rating = 1%, 12.847 per %), hit=0.950 ± 0.034 per rating point (10 rating = 1%, 9.504 per %), melee_haste=5.132 ± 0.678

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.0 ranged_attack_power points (2.03 DPS) | yes | Red Winter Hat (21524, -2.14 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.7 ranged_attack_power points (0.51 DPS) | yes | Roadwatcher's Confidence (281265, -0.42 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 16.9 ranged_attack_power points (1.27 DPS) | yes | Slime-encrusted Pads (6461, -0.72 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.2 ranged_attack_power points (1.52 DPS) | yes | Cape of the Brotherhood (5193, -0.35 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.51 DPS) [world_drop]; Bristlebark Cape (14571, -0.51 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 23.6 ranged_attack_power points (1.77 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.25 DPS) [crafted]; Prospector's Chestpiece (14562, -0.25 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.9 ranged_attack_power points (1.27 DPS) | yes | Wolf Bracers (4794, -0.29 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.51 DPS) [world_drop]; Spare Part Bindings (279875, -0.51 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (103.6 DPS) | yes | Forest Leather Gloves (3058, -0.51 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.51 DPS) [crafted]; Serpent Gloves (5970, -1.28 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.35 DPS) | yes | Deviate Scale Belt (6468, -0.08 DPS) [crafted]; Guardsman Belt (3429, -0.34 DPS) [world]; Dark Leather Belt (4249, -0.34 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 30.4 ranged_attack_power points (2.28 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.25 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.0 ranged_attack_power points (2.03 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.51 DPS) [dungeon]; Gray Leather Boots (2315, -0.76 DPS) [crafted] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 10.1 ranged_attack_power points (0.76 DPS) | yes | The 1 Ring (8350, -0.51 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.7 ranged_attack_power points (0.51 DPS) | yes | The 1 Ring (8350, -0.81 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 33.0 ranged_attack_power points (2.48 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.70 DPS) [quest]; Lupine Axe (1220, -0.96 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.6 ranged_attack_power points (13.49 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.27 DPS) [crafted]; Venomstrike (6469, -3.42 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Bounty Hunter's Ring; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 223, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 134.1. Weights run: 1.7s. Verify run: 1.9s. 431 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.308 ± 0.017, crit=1.262 ± 0.023 per rating point (14 rating = 1%, 17.670 per %), hit=1.073 ± 0.049 per rating point (10 rating = 1%, 10.727 per %), melee_haste=not significant (-3.863 ± 1.015)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 33.1 ranged_attack_power points (2.36 DPS) | yes | Brawler's Leather Hood (252504, -0.56 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.71 DPS) [world]; Resilient Cap (14401, -1.89 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.2 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.09 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.4 ranged_attack_power points (2.60 DPS) | yes | Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.95 DPS) [crafted]; Insignia Mantle (4721, -0.95 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.5 ranged_attack_power points (1.89 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Fenrus' Hide (6340, -0.47 DPS) [dungeon]; Swiftrunner Cape (6745, -0.47 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Panther Armor (6670, +0.00 DPS) [quest]; Green Leather Armor (4255, -0.24 DPS) [crafted]; Dusky Leather Armor (7374, -3.24 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.8 ranged_attack_power points (1.42 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.24 DPS) [world_drop]; Loamflake Bracers (15462, -0.24 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.8 ranged_attack_power points (2.13 DPS) | yes | Braced Handguards (6784, -0.47 DPS) [quest]; Pilferer's Gloves (7358, -0.49 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.71 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Deftkin Belt (16659, -0.32 DPS) [quest]; Larcenist's Belt (279846, -0.71 DPS) [quest]; Skulker's Leather Belt (252520, -1.77 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 46.3 ranged_attack_power points (3.31 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.24 DPS) [crafted]; Insignia Leggings (4054, -1.18 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 26.5 ranged_attack_power points (1.89 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.24 DPS) [quest] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.8 ranged_attack_power points (1.42 DPS) | yes | Defias Renegade Ring (1076, -0.24 DPS) [dungeon]; Band of the Fist (17694, -0.47 DPS) [quest]; Thunderbrow Ring (13097, -0.71 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 16.5 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Band of the Fist (17694, -0.24 DPS) [quest]; Thunderbrow Ring (13097, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 19.8 ranged_attack_power points (1.42 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 17.2 ranged_attack_power points (1.23 DPS) | yes | Wingblade (6504, -2.21 DPS, sim-verified) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 252.2 ranged_attack_power points (18.02 DPS) | yes | Ironweaver (13137, -0.71 DPS) [world_drop]; Crystalpine Stinger (13037, -1.18 DPS) [world_drop]; Chesterfall Musket (7729, -1.50 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 168.5. Weights run: 1.7s. Verify run: 1.8s. 678 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.153 ± 0.017, crit=1.420 ± 0.028 per rating point (14 rating = 1%, 19.886 per %), hit=1.294 ± 0.055 per rating point (10 rating = 1%, 12.939 per %), melee_haste=6.810 ± 1.067

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 41.0 ranged_attack_power points (3.10 DPS) | yes | Guard's Chain Helm (250499, -0.24 DPS) [crafted]; Skullsplitter Helm (1624, -0.48 DPS) [world]; Nightscape Headband (8176, -1.62 DPS, sim-verified) [crafted] |
| neck | Ethereal Talisman (4430) (or Kaleidoscope Chain (13084)) | The Crown of Will [quest] | 12.6 ranged_attack_power points (0.95 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 34.7 ranged_attack_power points (2.62 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.2 ranged_attack_power points (1.91 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 75.7 ranged_attack_power points (5.73 DPS) | yes | Wolffear Harness (13110, -1.86 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.15 DPS) [crafted]; Tough Scorpid Breastplate (8203, -2.15 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 25.2 ranged_attack_power points (1.91 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.40 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 39.9 ranged_attack_power points (3.02 DPS) | yes | Dragonscale Gauntlets (8347, -0.08 DPS) [crafted]; Gauntlets of Divinity (7724, -0.60 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.63 DPS) [crafted] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 31.5 ranged_attack_power points (2.39 DPS) | yes | Blackforge Girdle (6425, -0.24 DPS) [dungeon]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 66.2 ranged_attack_power points (5.01 DPS) | yes | Triprunner Dungarees (9624, -1.30 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.67 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.67 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 41.0 ranged_attack_power points (3.10 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.48 DPS) [dungeon]; Dusky Boots (7390, -0.48 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 31.5 ranged_attack_power points (2.39 DPS) | yes | Disengagement Ring (276202, -0.48 DPS) [vendor]; Field Researcher's Loop (281634, -0.72 DPS) [quest]; Mark of Kern (2262, -0.87 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 28.4 ranged_attack_power points (2.15 DPS) | yes | Disengagement Ring (276202, +0.00 DPS, sim-verified) [vendor]; Field Researcher's Loop (281634, -0.48 DPS) [quest]; Mark of Kern (2262, -0.63 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (168.5 DPS) | yes | Headhunting Spear (1522, -1.23 DPS) [world]; Armor Piercer (6679, -1.23 DPS) [dungeon]; Gut Ripper (2164, -4.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (26.48 DPS) | yes | Swiftwind (13038, -2.30 DPS) [world_drop]; Baelog's Shortbow (9400, -2.53 DPS) [dungeon]; The Silencer (13138, -3.28 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Ethereal Talisman; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 678, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 233.8. Weights run: 1.8s. Verify run: 2.5s. 868 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.072 ± 0.017, crit=1.626 ± 0.031 per rating point (14 rating = 1%, 22.759 per %), hit=1.748 ± 0.072 per rating point (10 rating = 1%, 17.484 per %), melee_haste=12.189 ± 1.173

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 52.2 ranged_attack_power points (4.26 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.50 DPS) [quest]; Tough Scorpid Helm (8208, -0.75 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.9 ranged_attack_power points (3.26 DPS) | yes | Woven Ivy Necklace (19159, -1.94 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -2.26 DPS) [quest]; Kaleidoscope Chain (13084, -2.26 DPS) [world_drop] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 36.9 ranged_attack_power points (3.01 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.25 DPS) [world_drop]; Nightscape Shoulders (8192, -1.06 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 46.1 ranged_attack_power points (3.76 DPS) | yes | Serpentskin Cloak (8259, -1.00 DPS) [world_drop]; Nightscape Cloak (8195, -1.25 DPS) [crafted]; Blackflame Cape (13109, -1.78 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 73.7 ranged_attack_power points (6.02 DPS) | yes | Blazewind Breastplate (11193, -0.25 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.40 DPS) [vendor]; Charred Leather Tunic (19127, -1.51 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.1 ranged_attack_power points (3.76 DPS) | yes | Bloodlust Bracelets (14807, -0.91 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -1.00 DPS) [crafted]; Pridelord Bands (14672, -1.25 DPS) [world_drop] |
| hands | First Sergeant's Chain Gauntlets (220830) | PvP rank 9 · First Sergeant · Horde [vendor] | 49.1 ranged_attack_power points (4.01 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.50 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.50 DPS) [crafted]; Gloves of Holy Might (867, -0.52 DPS) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 46.1 ranged_attack_power points (3.76 DPS) | yes | Stalker's Mail Belt (252588, -0.25 DPS) [crafted]; Seared Mail Girdle (19125, -0.75 DPS) [quest]; Skulker's Leather Waistguard (252474, -4.43 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 65.8 ranged_attack_power points (5.37 DPS) | yes | Keeper's Woolies (14668, -0.60 DPS) [world_drop]; Oilskin Leggings (9414, -0.86 DPS) [dungeon]; Basilisk Hide Pants (1718, -3.33 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 61.4 ranged_attack_power points (5.02 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.50 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -1.25 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.0 ranged_attack_power points (3.51 DPS) | yes | Falcon's Hook (7552, -1.25 DPS) [dungeon]; Disengagement Ring (276202, -1.51 DPS) [vendor]; White Bone Band (11862, -1.55 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.7 ranged_attack_power points (2.51 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; White Bone Band (11862, -0.55 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -8.91 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -0.98 DPS, sim-verified) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Widowmaker (4091) | World drop [world_drop] | sim-verified (233.8 DPS) | yes | White Bone Shredder (11863, -0.25 DPS) [quest]; Force of the Hippogryph (9684, -2.94 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.75 DPS) [dungeon]; Houndmaster's Bow (11628, -3.05 DPS) [dungeon]; Dark Iron Rifle (16004, -6.39 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: First Sergeant's Chain Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Widowmaker; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 296.1. Weights run: 1.7s. Verify run: 6.1s. 1746 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.077 ± 0.019, crit=2.087 ± 0.041 per rating point (14 rating = 1%, 29.218 per %), hit=2.182 ± 0.105 per rating point (10 rating = 1%, 21.823 per %), melee_haste=15.069 ± 1.652

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (296.1 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -10.40 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.0 ranged_attack_power points (3.24 DPS) | yes | Sunprism Pendant (20649, -0.25 DPS) [quest]; Woven Ivy Necklace (19159, -1.00 DPS) [quest]; Ethereal Talisman (4430, -2.24 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 72.1 ranged_attack_power points (5.83 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS, sim-verified) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.2 ranged_attack_power points (5.36 DPS) | yes | Dark Phantom Cape (13122, -1.82 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.37 DPS) [world_drop]; Deep River Cloak (15789, -2.62 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 164.9 ranged_attack_power points (13.35 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.51 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.51 DPS) [vendor]; Shadowcraft Tunic (16721, -5.73 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) | Stratholme: Fleshflayer Ghoul [dungeon] | sim-verified (296.1 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -6.50 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (296.1 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Mourning Handguards (279257, -4.94 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (296.1 DPS) | yes | Wicked Leather Belt (15088, +0.00 DPS) [crafted]; General's Chain Girdle (16572, +0.00 DPS) [pvp]; Sagebrush Girdle (17778, -5.47 DPS, sim-verified) [quest] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (296.1 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -10.13 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (296.1 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -6.53 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.1 ranged_attack_power points (3.49 DPS) | yes | Falcon's Hook (7552, -1.25 DPS) [dungeon]; Disengagement Ring (276202, -1.49 DPS) [vendor]; White Bone Band (11862, -1.54 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.8 ranged_attack_power points (2.49 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.50 DPS) [vendor]; White Bone Band (11862, -0.55 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (296.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (296.1 DPS) | yes | Molten Heart of the Mountain (249470, -1.83 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (296.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -11.03 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (296.1 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -11.35 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 647.3. Weights run: 1.7s. Verify run: 6.5s. 1746 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.265 ± 0.021, crit=2.324 ± 0.042 per rating point (14 rating = 1%, 32.532 per %), hit=2.931 ± 0.142 per rating point (10 rating = 1%, 29.307 per %), melee_haste=15.726 ± 2.080

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (647.3 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -13.46 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.4 ranged_attack_power points (6.49 DPS) | yes | Sunprism Pendant (20649, -0.50 DPS) [quest]; Woven Ivy Necklace (19159, -2.00 DPS) [quest]; Ethereal Talisman (4430, -4.49 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 77.7 ranged_attack_power points (11.88 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 69.0 ranged_attack_power points (10.55 DPS) | yes | Dark Phantom Cape (13122, -3.67 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.56 DPS) [world_drop]; Deep River Cloak (15789, -5.06 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 182.6 ranged_attack_power points (27.93 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.98 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.98 DPS) [vendor]; Shadowcraft Tunic (16721, -13.67 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 49.0 ranged_attack_power points (7.49 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -5.50 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 68.6 ranged_attack_power points (10.48 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.13 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 49.0 ranged_attack_power points (7.49 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Legionnaire's Chain Girdle (16529, -0.50 DPS) [pvp]; Wicked Leather Belt (15088, -5.92 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (647.3 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -9.13 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (647.3 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -6.20 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.7 ranged_attack_power points (6.99 DPS) | yes | Falcon's Hook (7552, -2.50 DPS) [dungeon]; Disengagement Ring (276202, -3.00 DPS) [vendor]; White Bone Band (11862, -3.32 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 32.6 ranged_attack_power points (4.99 DPS) | yes | Falcon's Hook (7552, -0.50 DPS) [dungeon]; Disengagement Ring (276202, -1.00 DPS) [vendor]; White Bone Band (11862, -1.32 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (647.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (647.3 DPS) | yes | Molten Heart of the Mountain (249470, -5.05 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (647.3 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -22.05 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (647.3 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -20.34 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

