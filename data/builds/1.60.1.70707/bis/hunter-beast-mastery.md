# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 102.3. Weights run: 1.7s. Verify run: 1.8s. 233 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.427 ± 0.014, crit=0.922 ± 0.016 per rating point (14 rating = 1%, 12.906 per %), hit=0.966 ± 0.037 per rating point (10 rating = 1%, 9.657 per %), melee_haste=8.423 ± 0.604

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.08 DPS) | yes | Red Winter Hat (21524, -2.26 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.9 ranged_attack_power points (0.52 DPS) | yes | Tarnished Locket (279870, -0.45 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.30 DPS) | yes | Slime-encrusted Pads (6461, -1.01 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.6 ranged_attack_power points (1.56 DPS) | yes | Sentry Cloak (2059, -0.52 DPS) [world_drop]; Bristlebark Cape (14571, -0.52 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.64 DPS, sim-verified) [dungeon] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 37.7 ranged_attack_power points (2.86 DPS) | yes | Brawler's Leather Armor (252490, -0.93 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -1.04 DPS) [crafted]; Dark Leather Tunic (2317, -1.30 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.30 DPS) | yes | Wolf Bracers (4794, -0.26 DPS) [vendor]; Bravo's Armbands (270015, -0.26 DPS) [quest]; Bristlebark Bindings (14569, -0.52 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 20.6 ranged_attack_power points (1.56 DPS) | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Forest Leather Gloves (3058, -0.52 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.52 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.37 DPS) | yes | Deviate Scale Belt (6468, -0.07 DPS) [crafted]; Dusty Belt (279897, -0.07 DPS) [quest]; Dark Leather Belt (4249, -0.33 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 30.8 ranged_attack_power points (2.34 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.26 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.4 ranged_attack_power points (2.08 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.52 DPS) [dungeon]; Gray Leather Boots (2315, -0.78 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.9 ranged_attack_power points (0.52 DPS) | yes | - |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.4 ranged_attack_power points (0.26 DPS) | yes | Ring of the Shadow (1462, -0.46 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 33.4 ranged_attack_power points (2.54 DPS) | yes | Impaling Harpoon (5200, -0.20 DPS) [dungeon]; Lupine Axe (1220, -0.98 DPS) [world]; Piercing Axe (6094, -0.98 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.65 DPS) | yes | Lil Timmy's Peashooter (13136, -1.18 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.31 DPS) [crafted]; Venomstrike (6469, -3.46 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 233, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 131.4. Weights run: 1.7s. Verify run: 1.7s. 440 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.228 ± 0.013, crit=1.063 ± 0.018 per rating point (14 rating = 1%, 14.883 per %), hit=1.077 ± 0.045 per rating point (10 rating = 1%, 10.766 per %), melee_haste=9.392 ± 0.790

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 32.3 ranged_attack_power points (2.38 DPS) | yes | Brawler's Leather Hood (252504, -0.52 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.72 DPS) [world]; Resilient Cap (14401, -1.91 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.9 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 35.5 ranged_attack_power points (2.62 DPS) | yes | Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.95 DPS) [crafted]; Insignia Mantle (4721, -0.95 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 25.8 ranged_attack_power points (1.91 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Fenrus' Hide (6340, -0.48 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.48 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 45.2 ranged_attack_power points (3.34 DPS) | yes | Tunic of Westfall (2041, -0.77 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -1.19 DPS) [dungeon]; Green Leather Armor (4255, -1.43 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.4 ranged_attack_power points (1.43 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.24 DPS) [world]; Forest Leather Bracers (3202, -0.24 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.1 ranged_attack_power points (2.15 DPS) | yes | Pilferer's Gloves (7358, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.72 DPS) [dungeon]; Insignia Gloves (6408, -0.72 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (131.4 DPS) | yes | Belt of Vindication (3562, -0.48 DPS) [quest]; Green Leather Belt (4257, -0.72 DPS) [crafted]; Skulker's Leather Belt (252520, -1.33 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 45.2 ranged_attack_power points (3.34 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.24 DPS) [crafted]; Insignia Leggings (4054, -1.19 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 25.8 ranged_attack_power points (1.91 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, -0.48 DPS) [dungeon]; Footpads of the Fang (10411, -0.48 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.4 ranged_attack_power points (1.43 DPS) | yes | Defias Renegade Ring (1076, -0.24 DPS) [dungeon]; Seal of Wrynn (2933, -0.72 DPS) [quest]; Thunderbrow Ring (13097, -0.72 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 16.1 ranged_attack_power points (1.19 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Seal of Wrynn (2933, -0.48 DPS) [quest]; Thunderbrow Ring (13097, -0.48 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 19.4 ranged_attack_power points (1.43 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 16.9 ranged_attack_power points (1.25 DPS) | yes | The Butcher (8226, -0.58 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 251.8 ranged_attack_power points (18.60 DPS) | yes | Ironweaver (13137, -0.70 DPS) [world_drop]; Crystalpine Stinger (13037, -1.21 DPS) [world_drop]; Chesterfall Musket (7729, -1.51 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 440, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 174.6. Weights run: 1.9s. Verify run: 1.9s. 696 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.220 ± 0.037, crit=1.634 ± 0.060 per rating point (14 rating = 1%, 22.875 per %), hit=1.126 ± 0.067 per rating point (10 rating = 1%, 11.255 per %), melee_haste=11.458 ± 1.115

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 41.9 ranged_attack_power points (3.08 DPS) | yes | Nightscape Headband (8176, -0.24 DPS) [crafted]; Guard's Chain Helm (250499, -0.24 DPS) [crafted]; Skullsplitter Helm (1624, -0.47 DPS) [world] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.9 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest]; Amberglow Talisman (10824, -0.71 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 35.4 ranged_attack_power points (2.61 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.8 ranged_attack_power points (1.90 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 77.3 ranged_attack_power points (5.69 DPS) | yes | Wolffear Harness (13110, -2.08 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.13 DPS) [crafted]; Tough Scorpid Breastplate (8203, -2.13 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 25.8 ranged_attack_power points (1.90 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.42 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 42.9 ranged_attack_power points (3.16 DPS) | yes | Dragonscale Gauntlets (8347, -0.05 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.79 DPS) [crafted]; Scarlet Gauntlets (10331, -0.79 DPS) [dungeon] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 32.2 ranged_attack_power points (2.37 DPS) | yes | Blackforge Girdle (6425, -0.24 DPS) [dungeon]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 67.6 ranged_attack_power points (4.98 DPS) | yes | Triprunner Dungarees (9624, -1.20 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.66 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.66 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 41.9 ranged_attack_power points (3.08 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.47 DPS) [dungeon]; Dusky Boots (7390, -0.47 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.2 ranged_attack_power points (2.37 DPS) | yes | Disengagement Ring (276202, -0.47 DPS) [vendor]; Field Researcher's Loop (281634, -0.71 DPS) [quest]; Mark of Kern (2262, -0.90 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 29.0 ranged_attack_power points (2.13 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Field Researcher's Loop (281634, -0.47 DPS) [quest]; Mark of Kern (2262, -0.66 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (174.6 DPS) | yes | Headhunting Spear (1522, -1.21 DPS) [world]; Armor Piercer (6679, -1.21 DPS) [dungeon]; Gut Ripper (2164, -4.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (25.76 DPS) | yes | Swiftwind (13038, -2.21 DPS) [world_drop]; Baelog's Shortbow (9400, -2.46 DPS) [dungeon]; The Silencer (13138, -3.72 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 696, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 220.9. Weights run: 2.0s. Verify run: 2.5s. 885 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.088 ± 0.033, crit=1.665 ± 0.060 per rating point (14 rating = 1%, 23.310 per %), hit=1.447 ± 0.079 per rating point (10 rating = 1%, 14.472 per %), melee_haste=8.414 ± 1.127

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 52.5 ranged_attack_power points (4.03 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Lordrec Helmet (10741, -0.24 DPS) [quest]; Sprightring Helm (17776, -0.47 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.1 ranged_attack_power points (3.08 DPS) | yes | Kaleidoscope Chain (13084, -1.62 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -2.16 DPS) [world]; Erudite's Amulet (277204, -2.61 DPS) [quest] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 37.1 ranged_attack_power points (2.85 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Nightscape Shoulders (8192, -0.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 46.3 ranged_attack_power points (3.56 DPS) | yes | Blackflame Cape (13109, -0.71 DPS) [world_drop]; Serpentskin Cloak (8259, -0.95 DPS) [world_drop]; Nightscape Cloak (8195, -1.19 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 74.1 ranged_attack_power points (5.69 DPS) | yes | Blazewind Breastplate (11193, -0.24 DPS) [quest]; Knight's Chain Armor (220828, -0.34 DPS) [vendor]; Charred Leather Tunic (19127, -1.42 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.3 ranged_attack_power points (3.56 DPS) | yes | Bloodlust Bracelets (14807, -0.95 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.95 DPS) [crafted]; Pridelord Bands (14672, -1.19 DPS) [world_drop] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 49.4 ranged_attack_power points (3.79 DPS) | yes | Gloves of Holy Might (867, -0.47 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.47 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.47 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 46.3 ranged_attack_power points (3.56 DPS) | yes | Stalker's Mail Belt (252588, -0.24 DPS) [crafted]; Seared Mail Girdle (19125, -0.71 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.61 DPS, sim-verified) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 66.5 ranged_attack_power points (5.11 DPS) | yes | Keeper's Woolies (14668, -0.60 DPS) [world_drop]; Oilskin Leggings (9414, -0.84 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.28 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 61.8 ranged_attack_power points (4.74 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.47 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -1.19 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.2 ranged_attack_power points (3.32 DPS) | yes | Falcon's Hook (7552, -1.19 DPS) [dungeon]; Disengagement Ring (276202, -1.42 DPS) [vendor]; Field Researcher's Loop (281634, -1.66 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.9 ranged_attack_power points (2.37 DPS) | yes | Falcon's Hook (7552, -0.24 DPS) [dungeon]; Disengagement Ring (276202, -0.47 DPS) [vendor]; Field Researcher's Loop (281634, -0.71 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (220.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (220.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (220.9 DPS) | yes | Grimlok's Charge (9416, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -2.39 DPS, sim-verified) [world_drop] |
| off_hand | Widowmaker (4091) (or Blue Glittering Axe (7942), Dazzling Mithril Rapier (7944), Nordic Longshank (9401), Speedsteel Rapier (13034), Mithril Stiletto (250611)) | World drop [world_drop] | 24.7 ranged_attack_power points (1.90 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Dazzling Mithril Rapier (7944, +0.00 DPS) [crafted]; Nordic Longshank (9401, +0.00 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (220.9 DPS) | yes | Gryphonwing Long Bow (13022, -0.71 DPS) [dungeon]; Houndmaster's Bow (11628, -2.88 DPS) [dungeon]; Dark Iron Rifle (16004, -5.49 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Widowmaker; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 885, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 5420001505001251-0053542001000000-000000000000000000)

Set DPS (verified): 288.6. Weights run: 2.0s. Verify run: 7.0s. 1755 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.079 ± 0.037, crit=2.066 ± 0.076 per rating point (14 rating = 1%, 28.923 per %), hit=2.107 ± 0.109 per rating point (10 rating = 1%, 21.070 per %), melee_haste=15.588 ± 1.691

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (288.6 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -9.63 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.0 ranged_attack_power points (3.04 DPS) | yes | Sunprism Pendant (20649, -0.23 DPS) [quest]; Kaleidoscope Chain (13084, -2.10 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.13 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 72.0 ranged_attack_power points (5.46 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.2 ranged_attack_power points (5.02 DPS) | yes | Dark Phantom Cape (13122, -1.07 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.22 DPS) [world_drop]; Deep River Cloak (15789, -2.45 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 163.9 ranged_attack_power points (12.44 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.30 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.30 DPS) [vendor]; Shadowcraft Tunic (16721, -4.88 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 46.2 ranged_attack_power points (3.51 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -8.08 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (288.6 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -5.48 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (288.6 DPS) | yes | Wicked Leather Belt (15088, +0.00 DPS) [crafted]; Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Sagebrush Girdle (17778, -7.01 DPS, sim-verified) [quest] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (288.6 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -9.76 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (288.6 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -8.19 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.1 ranged_attack_power points (3.27 DPS) | yes | Falcon's Hook (7552, -1.17 DPS) [dungeon]; Disengagement Ring (276202, -1.40 DPS) [vendor]; Field Researcher's Loop (281634, -1.64 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.8 ranged_attack_power points (2.34 DPS) | yes | Falcon's Hook (7552, -0.23 DPS) [dungeon]; Disengagement Ring (276202, -0.47 DPS) [vendor]; Field Researcher's Loop (281634, -0.70 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (288.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (288.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (288.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -9.63 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (288.6 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.64 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 5420001505001251-0053542001000000-000000000000000000)

Set DPS (verified): 625.7. Weights run: 1.8s. Verify run: 6.4s. 1755 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.259 ± 0.036, crit=2.297 ± 0.077 per rating point (14 rating = 1%, 32.163 per %), hit=3.335 ± 0.169 per rating point (10 rating = 1%, 33.352 per %), melee_haste=17.483 ± 2.271

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (625.7 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -12.33 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.4 ranged_attack_power points (5.66 DPS) | yes | Sunprism Pendant (20649, -0.44 DPS) [quest]; Kaleidoscope Chain (13084, -3.92 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.06 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 78.4 ranged_attack_power points (10.46 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 68.9 ranged_attack_power points (9.20 DPS) | yes | Dark Phantom Cape (13122, -3.41 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -3.98 DPS) [world_drop]; Deep River Cloak (15789, -4.41 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 186.1 ranged_attack_power points (24.85 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.76 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.76 DPS) [vendor]; Shadowcraft Tunic (16721, -12.04 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 48.9 ranged_attack_power points (6.53 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -5.00 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 68.4 ranged_attack_power points (9.14 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -6.89 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 48.9 ranged_attack_power points (6.53 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Knight-Captain's Chain Girdle (16400, -0.44 DPS) [pvp]; Wicked Leather Belt (15088, -5.91 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (625.7 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -8.42 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (625.7 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -3.97 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.6 ranged_attack_power points (6.09 DPS) | yes | Falcon's Hook (7552, -2.18 DPS) [dungeon]; Disengagement Ring (276202, -2.61 DPS) [vendor]; Field Researcher's Loop (281634, -3.05 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 32.6 ranged_attack_power points (4.35 DPS) | yes | Falcon's Hook (7552, -0.44 DPS) [dungeon]; Disengagement Ring (276202, -0.87 DPS) [vendor]; Field Researcher's Loop (281634, -1.31 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (625.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (625.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (625.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -20.41 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (625.7 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -17.64 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 104.5. Weights run: 1.7s. Verify run: 2.9s. 229 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.427 ± 0.014, crit=0.922 ± 0.016 per rating point (14 rating = 1%, 12.906 per %), hit=0.966 ± 0.037 per rating point (10 rating = 1%, 9.657 per %), melee_haste=8.423 ± 0.604

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.08 DPS) | yes | Red Winter Hat (21524, -2.20 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.9 ranged_attack_power points (0.52 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.30 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.6 ranged_attack_power points (1.56 DPS) | yes | Cape of the Brotherhood (5193, -0.26 DPS) [dungeon]; Sentry Cloak (2059, -0.52 DPS) [world_drop]; Bristlebark Cape (14571, -0.52 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 24.0 ranged_attack_power points (1.82 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.26 DPS) [crafted]; Prospector's Chestpiece (14562, -0.26 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.30 DPS) | yes | Wolf Bracers (4794, -0.26 DPS) [vendor]; Bristlebark Bindings (14569, -0.52 DPS) [world_drop]; Spare Part Bindings (279875, -0.52 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (104.5 DPS) | yes | Forest Leather Gloves (3058, -0.52 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.52 DPS) [crafted]; Serpent Gloves (5970, -0.61 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.37 DPS) | yes | Deviate Scale Belt (6468, -0.07 DPS) [crafted]; Guardsman Belt (3429, -0.33 DPS) [world]; Dark Leather Belt (4249, -0.33 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 30.8 ranged_attack_power points (2.34 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.26 DPS) [world]; Brawler's Leather Pants (252500, -0.47 DPS, sim-verified) [crafted] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (104.5 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.26 DPS) [crafted] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 10.3 ranged_attack_power points (0.78 DPS) | yes | The 1 Ring (8350, -0.52 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.9 ranged_attack_power points (0.52 DPS) | yes | The 1 Ring (8350, -0.66 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 33.4 ranged_attack_power points (2.54 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.72 DPS) [quest]; Lupine Axe (1220, -0.98 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.65 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.31 DPS) [crafted]; Venomstrike (6469, -3.46 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Bounty Hunter's Ring; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 229, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 132.4. Weights run: 1.7s. Verify run: 1.7s. 437 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.228 ± 0.013, crit=1.063 ± 0.018 per rating point (14 rating = 1%, 14.883 per %), hit=1.077 ± 0.045 per rating point (10 rating = 1%, 10.766 per %), melee_haste=9.392 ± 0.790

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 32.3 ranged_attack_power points (2.38 DPS) | yes | Brawler's Leather Hood (252504, -0.52 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.72 DPS) [world]; Resilient Cap (14401, -1.91 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.9 ranged_attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 35.5 ranged_attack_power points (2.62 DPS) | yes | Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.95 DPS) [crafted]; Insignia Mantle (4721, -0.95 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 25.8 ranged_attack_power points (1.91 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Fenrus' Hide (6340, -0.48 DPS) [dungeon]; Swiftrunner Cape (6745, -0.48 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 45.2 ranged_attack_power points (3.34 DPS) | yes | Panther Armor (6670, -1.19 DPS) [quest]; Garb of Florid Feathers (273026, -1.19 DPS) [dungeon]; Green Leather Armor (4255, -1.43 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.4 ranged_attack_power points (1.43 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.24 DPS) [world_drop]; Loamflake Bracers (15462, -0.24 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.1 ranged_attack_power points (2.15 DPS) | yes | Braced Handguards (6784, -0.48 DPS) [quest]; Pilferer's Gloves (7358, -0.71 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.72 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 29.1 ranged_attack_power points (2.15 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Deftkin Belt (16659, -0.31 DPS) [quest]; Larcenist's Belt (279846, -0.72 DPS) [quest] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 45.2 ranged_attack_power points (3.34 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.24 DPS) [crafted]; Insignia Leggings (4054, -1.19 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 25.8 ranged_attack_power points (1.91 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.24 DPS) [quest] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.4 ranged_attack_power points (1.43 DPS) | yes | Defias Renegade Ring (1076, -0.24 DPS) [dungeon]; Band of the Fist (17694, -0.48 DPS) [quest]; Thunderbrow Ring (13097, -0.72 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 16.1 ranged_attack_power points (1.19 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Band of the Fist (17694, -0.24 DPS) [quest]; Thunderbrow Ring (13097, -0.48 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 19.4 ranged_attack_power points (1.43 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 16.9 ranged_attack_power points (1.25 DPS) | yes | Wingblade (6504, -1.34 DPS, sim-verified) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 251.8 ranged_attack_power points (18.60 DPS) | yes | Ironweaver (13137, -1.00 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.21 DPS) [world_drop]; Chesterfall Musket (7729, -1.51 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 437, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 175.0. Weights run: 1.9s. Verify run: 1.9s. 684 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.220 ± 0.037, crit=1.634 ± 0.060 per rating point (14 rating = 1%, 22.875 per %), hit=1.126 ± 0.067 per rating point (10 rating = 1%, 11.255 per %), melee_haste=11.458 ± 1.115

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 41.9 ranged_attack_power points (3.08 DPS) | yes | Nightscape Headband (8176, -0.24 DPS) [crafted]; Guard's Chain Helm (250499, -0.24 DPS) [crafted]; Skullsplitter Helm (1624, -0.47 DPS) [world] |
| neck | Ethereal Talisman (4430) (or Kaleidoscope Chain (13084)) | The Crown of Will [quest] | 12.9 ranged_attack_power points (0.95 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.06 DPS) [world]; Erudite's Amulet (277204, -0.47 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 35.4 ranged_attack_power points (2.61 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 25.8 ranged_attack_power points (1.90 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 77.3 ranged_attack_power points (5.69 DPS) | yes | Wolffear Harness (13110, -1.11 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.13 DPS) [crafted]; Tough Scorpid Breastplate (8203, -2.13 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 25.8 ranged_attack_power points (1.90 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.24 DPS) [crafted]; Branded Leather Bracers (19508, -0.42 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 42.9 ranged_attack_power points (3.16 DPS) | yes | Dragonscale Gauntlets (8347, -0.05 DPS) [crafted]; Tough Scorpid Gloves (8204, -0.79 DPS) [crafted]; Scarlet Gauntlets (10331, -0.79 DPS) [dungeon] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 32.2 ranged_attack_power points (2.37 DPS) | yes | Blackforge Girdle (6425, -0.24 DPS) [dungeon]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.24 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 67.6 ranged_attack_power points (4.98 DPS) | yes | Triprunner Dungarees (9624, -0.71 DPS) [quest]; Petrolspill Leggings (9509, -1.66 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.66 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 41.9 ranged_attack_power points (3.08 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.47 DPS) [dungeon]; Dusky Boots (7390, -0.47 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.2 ranged_attack_power points (2.37 DPS) | yes | Disengagement Ring (276202, -0.47 DPS) [vendor]; Field Researcher's Loop (281634, -0.71 DPS) [quest]; Mark of Kern (2262, -0.90 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 29.0 ranged_attack_power points (2.13 DPS) | yes | Disengagement Ring (276202, +0.00 DPS, sim-verified) [vendor]; Field Researcher's Loop (281634, -0.47 DPS) [quest]; Mark of Kern (2262, -0.66 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (175.0 DPS) | yes | Headhunting Spear (1522, -1.21 DPS) [world]; Armor Piercer (6679, -1.21 DPS) [dungeon]; Gut Ripper (2164, -4.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (25.76 DPS) | yes | Swiftwind (13038, -2.21 DPS) [world_drop]; Baelog's Shortbow (9400, -2.46 DPS) [dungeon]; The Silencer (13138, -2.81 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Ethereal Talisman; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 230.0. Weights run: 2.0s. Verify run: 2.8s. 874 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.088 ± 0.033, crit=1.665 ± 0.060 per rating point (14 rating = 1%, 23.310 per %), hit=1.447 ± 0.079 per rating point (10 rating = 1%, 14.472 per %), melee_haste=8.414 ± 1.127

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 52.5 ranged_attack_power points (4.03 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.47 DPS) [quest]; Tough Scorpid Helm (8208, -0.71 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.1 ranged_attack_power points (3.08 DPS) | yes | Woven Ivy Necklace (19159, -1.17 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -2.13 DPS) [quest]; Kaleidoscope Chain (13084, -2.13 DPS) [world_drop] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 37.1 ranged_attack_power points (2.85 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Nightscape Shoulders (8192, -0.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 46.3 ranged_attack_power points (3.56 DPS) | yes | Blackflame Cape (13109, -0.75 DPS, sim-verified) [world_drop]; Serpentskin Cloak (8259, -0.95 DPS) [world_drop]; Nightscape Cloak (8195, -1.19 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 74.1 ranged_attack_power points (5.69 DPS) | yes | Blazewind Breastplate (11193, -0.24 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.34 DPS) [vendor]; Charred Leather Tunic (19127, -1.42 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.3 ranged_attack_power points (3.56 DPS) | yes | Bloodlust Bracelets (14807, -0.95 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.95 DPS) [crafted]; Pridelord Bands (14672, -1.19 DPS) [world_drop] |
| hands | First Sergeant's Chain Gauntlets (220830) | PvP rank 9 · First Sergeant · Horde [vendor] | 49.4 ranged_attack_power points (3.79 DPS) | yes | Gloves of Holy Might (867, -0.47 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -0.47 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.47 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 46.3 ranged_attack_power points (3.56 DPS) | yes | Stalker's Mail Belt (252588, -0.24 DPS) [crafted]; Seared Mail Girdle (19125, -0.71 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.99 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 66.5 ranged_attack_power points (5.11 DPS) | yes | Keeper's Woolies (14668, -0.60 DPS) [world_drop]; Oilskin Leggings (9414, -0.84 DPS) [dungeon]; Basilisk Hide Pants (1718, -1.66 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 61.8 ranged_attack_power points (4.74 DPS) | yes | Fleetfoot Greaves (11627, -0.24 DPS) [dungeon]; Elven Chain Boots (13125, -0.47 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -1.19 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.2 ranged_attack_power points (3.32 DPS) | yes | Falcon's Hook (7552, -1.19 DPS) [dungeon]; Disengagement Ring (276202, -1.42 DPS) [vendor]; White Bone Band (11862, -1.48 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.9 ranged_attack_power points (2.37 DPS) | yes | Falcon's Hook (7552, -0.24 DPS) [dungeon]; Disengagement Ring (276202, -0.47 DPS) [vendor]; White Bone Band (11862, -0.53 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -7.94 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -0.67 DPS, sim-verified) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Grimlok's Charge (9416, +0.00 DPS) [dungeon]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Widowmaker (4091) | World drop [world_drop] | sim-verified (230.0 DPS) | yes | White Bone Shredder (11863, -0.24 DPS) [quest]; Force of the Hippogryph (9684, -2.41 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.71 DPS) [dungeon]; Dark Iron Rifle (16004, -2.61 DPS) [crafted]; Bow of Searing Arrows (2825, -7.76 DPS, sim-verified) [world_drop] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: First Sergeant's Chain Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Widowmaker; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 874, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5420001505001251-0053542001000000-000000000000000000)

Set DPS (verified): 292.7. Weights run: 2.0s. Verify run: 6.5s. 1752 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.079 ± 0.037, crit=2.066 ± 0.076 per rating point (14 rating = 1%, 28.923 per %), hit=2.107 ± 0.109 per rating point (10 rating = 1%, 21.070 per %), melee_haste=15.588 ± 1.691

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (292.7 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.58 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.0 ranged_attack_power points (3.04 DPS) | yes | Woven Ivy Necklace (19159, -0.94 DPS) [quest]; Sunprism Pendant (20649, -1.06 DPS, sim-verified) [quest]; Ethereal Talisman (4430, -2.10 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 72.0 ranged_attack_power points (5.46 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 66.2 ranged_attack_power points (5.02 DPS) | yes | Dark Phantom Cape (13122, -2.00 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.22 DPS) [world_drop]; Deep River Cloak (15789, -2.45 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 163.9 ranged_attack_power points (12.44 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.30 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.30 DPS) [vendor]; Shadowcraft Tunic (16721, -7.12 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 46.2 ranged_attack_power points (3.51 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -2.61 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 64.7 ranged_attack_power points (4.91 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -4.64 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 46.2 ranged_attack_power points (3.51 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Legionnaire's Chain Girdle (16529, -0.23 DPS) [pvp]; Wicked Leather Belt (15088, -3.47 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (292.7 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -5.69 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (292.7 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.03 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.1 ranged_attack_power points (3.27 DPS) | yes | Falcon's Hook (7552, -1.17 DPS) [dungeon]; Disengagement Ring (276202, -1.40 DPS) [vendor]; White Bone Band (11862, -1.45 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.8 ranged_attack_power points (2.34 DPS) | yes | Falcon's Hook (7552, -0.23 DPS) [dungeon]; Disengagement Ring (276202, -0.47 DPS) [vendor]; White Bone Band (11862, -0.52 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (292.7 DPS) | yes | Frozen Heart of the Mountain (249469, -7.21 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (292.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (292.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -10.49 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (292.7 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -12.58 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1752, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5420001505001251-0053542001000000-000000000000000000)

Set DPS (verified): 640.2. Weights run: 1.8s. Verify run: 6.3s. 1752 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.259 ± 0.036, crit=2.297 ± 0.077 per rating point (14 rating = 1%, 32.163 per %), hit=3.335 ± 0.169 per rating point (10 rating = 1%, 33.352 per %), melee_haste=17.483 ± 2.271

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (640.2 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -13.31 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.4 ranged_attack_power points (5.66 DPS) | yes | Sunprism Pendant (20649, -0.44 DPS) [quest]; Woven Ivy Necklace (19159, -1.74 DPS) [quest]; Ethereal Talisman (4430, -3.92 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 78.4 ranged_attack_power points (10.46 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 68.9 ranged_attack_power points (9.20 DPS) | yes | Dark Phantom Cape (13122, -3.43 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -3.98 DPS) [world_drop]; Deep River Cloak (15789, -4.41 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 186.1 ranged_attack_power points (24.85 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.76 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.76 DPS) [vendor]; Shadowcraft Tunic (16721, -11.77 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 48.9 ranged_attack_power points (6.53 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -5.01 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 68.4 ranged_attack_power points (9.14 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.17 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 48.9 ranged_attack_power points (6.53 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Legionnaire's Chain Girdle (16529, -0.44 DPS) [pvp]; Wicked Leather Belt (15088, -5.48 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (640.2 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -8.80 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (640.2 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -5.83 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.6 ranged_attack_power points (6.09 DPS) | yes | Falcon's Hook (7552, -2.18 DPS) [dungeon]; Disengagement Ring (276202, -2.61 DPS) [vendor]; White Bone Band (11862, -2.89 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 32.6 ranged_attack_power points (4.35 DPS) | yes | Falcon's Hook (7552, -0.44 DPS) [dungeon]; Disengagement Ring (276202, -0.87 DPS) [vendor]; White Bone Band (11862, -1.15 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (640.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (640.2 DPS) | yes | Molten Heart of the Mountain (249470, -3.34 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (640.2 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -19.13 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (640.2 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -18.79 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1752, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

