# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 108.9. Weights run: 2.9s. Verify run: 4.6s. 227 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.112 ± 0.015, crit=0.605 ± 0.016 per rating point (14 rating = 1%, 8.470 per %), hit=1.109 ± 0.048 per rating point (10 rating = 1%, 11.089 per %), melee_haste=7.909 ± 0.726

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.9 ranged_attack_power points (1.72 DPS) | yes | Red Winter Hat (21524, -1.81 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.2 ranged_attack_power points (0.43 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.65 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.7 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.37 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.2 ranged_attack_power points (2.37 DPS) | yes | Brawler's Leather Armor (252490, -0.62 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.86 DPS) [crafted]; Dark Leather Tunic (2317, -1.08 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bravo's Armbands (270015, -0.22 DPS) [quest]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (108.9 DPS) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -0.61 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Dusty Belt (279897, -0.17 DPS) [quest]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 28.0 ranged_attack_power points (1.94 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world]; Brawler's Leather Pants (252500, -0.39 DPS, sim-verified) [crafted] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (108.9 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.22 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.2 ranged_attack_power points (0.43 DPS) | yes | - |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.1 ranged_attack_power points (0.22 DPS) | yes | Ring of the Shadow (1462, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.9 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, -0.20 DPS) [dungeon]; Lupine Axe (1220, -0.85 DPS) [world]; Piercing Axe (6094, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.42 DPS) | yes | Lil Timmy's Peashooter (13136, -0.88 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.00 DPS) [crafted]; Venomstrike (6469, -3.13 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 227, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 137.9. Weights run: 3.0s. Verify run: 2.6s. 434 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.894 ± 0.013, crit=0.674 ± 0.017 per rating point (14 rating = 1%, 9.436 per %), hit=1.207 ± 0.058 per rating point (10 rating = 1%, 12.067 per %), melee_haste=11.026 ± 0.959

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.97 DPS) | yes | Brawler's Leather Hood (252504, -0.42 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 ranged_attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.03 DPS) [world_drop]; Erudite's Amulet (277204, -0.42 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.17 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.39 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.39 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Tunic of Westfall (2041, -0.65 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -0.99 DPS) [dungeon]; Green Leather Armor (4255, -1.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.20 DPS) [world]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Pilferer's Gloves (7358, -0.20 DPS) [crafted]; Serpent Gloves (5970, -0.59 DPS) [dungeon]; Insignia Gloves (6408, -0.59 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS, sim-verified) [crafted]; Belt of Vindication (3562, -0.39 DPS) [quest]; Blackened Defias Belt (10403, -0.55 DPS) [dungeon] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, -0.39 DPS) [dungeon]; Footpads of the Fang (10411, -0.39 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Seal of Wrynn (2933, -0.59 DPS) [quest]; Thunderbrow Ring (13097, -0.59 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) (or Defias Renegade Ring (1076)) | Willix the Importer [quest] | 14.5 ranged_attack_power points (0.99 DPS) | yes | Defias Renegade Ring (1076, +0.00 DPS) [dungeon]; Seal of Wrynn (2933, -0.39 DPS) [quest]; Thunderbrow Ring (13097, -0.39 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Jade Serpentblade (3850) | Blacksmithing [crafted] | 15.6 ranged_attack_power points (1.06 DPS) | yes | The Butcher (8226, -1.46 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 249.8 ranged_attack_power points (17.04 DPS) | yes | Ironweaver (13137, -1.02 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.07 DPS) [world_drop]; Chesterfall Musket (7729, -1.26 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Jade Serpentblade; ranged: Golemsight Long Gun

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 181.8. Weights run: 3.3s. Verify run: 3.1s. 690 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.808 ± 0.013, crit=0.752 ± 0.020 per rating point (14 rating = 1%, 10.534 per %), hit=1.307 ± 0.095 per rating point (10 rating = 1%, 13.070 per %), melee_haste=12.121 ± 1.404

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 ranged_attack_power points (0.79 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Erudite's Amulet (277204, -0.42 DPS) [quest]; Amberglow Talisman (10824, -0.60 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 30.9 ranged_attack_power points (2.03 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.18 DPS) [dungeon]; Flintrock Shoulders (7755, -0.18 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 67.4 ranged_attack_power points (4.42 DPS) | yes | Wolffear Harness (13110, -0.99 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.66 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.66 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.10 DPS) | yes | Gloves of Holy Might (867, -0.10 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.26 DPS) [crafted]; Scarlet Gauntlets (10331, -0.26 DPS) [dungeon] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 28.1 ranged_attack_power points (1.84 DPS) | yes | Blackforge Girdle (6425, -0.18 DPS) [dungeon]; Ogron's Sash (13117, -0.18 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.18 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.0 ranged_attack_power points (3.87 DPS) | yes | Triprunner Dungarees (9624, -0.67 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.1 ranged_attack_power points (1.84 DPS) | yes | Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon]; Assault Band (13095, -0.53 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.3 ranged_attack_power points (1.66 DPS) | yes | Disengagement Ring (276202, +0.00 DPS, sim-verified) [vendor]; Mark of Kern (2262, -0.35 DPS) [dungeon]; Assault Band (13095, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (181.8 DPS) | yes | Hellslayer Battle Axe (13017, -0.82 DPS) [world_drop]; Armor Piercer (6679, -1.03 DPS) [dungeon]; Gut Ripper (2164, -3.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.98 DPS) | yes | Swiftwind (13038, -2.16 DPS) [world_drop]; Baelog's Shortbow (9400, -2.20 DPS) [dungeon]; The Silencer (13138, -2.32 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band; 5008 Quicksilver Ring

### Band 50 (dwarf, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 219.0. Weights run: 3.3s. Verify run: 3.7s. 879 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.784 ± 0.013, crit=0.920 ± 0.024 per rating point (14 rating = 1%, 12.877 per %), hit=1.562 ± 0.105 per rating point (10 rating = 1%, 15.618 per %), melee_haste=10.622 ± 1.555

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.3 ranged_attack_power points (3.21 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Lordrec Helmet (10741, -0.19 DPS) [quest]; Sprightring Helm (17776, -0.38 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.2 ranged_attack_power points (2.46 DPS) | yes | Kaleidoscope Chain (13084, -1.70 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.76 DPS, sim-verified) [world]; Erudite's Amulet (277204, -2.08 DPS) [quest] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 33.4 ranged_attack_power points (2.27 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.19 DPS) [world_drop]; Nightscape Shoulders (8192, -0.62 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.94 DPS) [crafted]; Blackflame Cape (13109, -1.13 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 66.8 ranged_attack_power points (4.54 DPS) | yes | Blazewind Breastplate (11193, -0.19 DPS) [quest]; Knight's Chain Armor (220828, -0.83 DPS) [vendor]; Charred Leather Tunic (19127, -1.13 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.93 DPS) [world]; Bloodlust Bracelets (14807, -1.46 DPS, sim-verified) [world_drop] |
| hands | Sergeant Major's Chain Gauntlets (220829) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 44.5 ranged_attack_power points (3.02 DPS) | yes | Stalker's Mail Gauntlets (252584, -0.38 DPS) [crafted]; Heavy Scorpid Gauntlet (15078, -0.76 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -0.81 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Stalker's Mail Belt (252588, -0.19 DPS) [crafted]; Seared Mail Girdle (19125, -0.57 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.75 DPS, sim-verified) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.5 ranged_attack_power points (3.97 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Knight's Chain Legplates (220832, -0.45 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.7 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, -0.19 DPS) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -0.94 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.0 ranged_attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.94 DPS) [dungeon]; Disengagement Ring (276202, -1.13 DPS) [vendor]; Blackstone Ring (17713, -1.29 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.8 ranged_attack_power points (1.89 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Disengagement Ring (276202, -0.38 DPS) [vendor]; Blackstone Ring (17713, -0.53 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (219.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (219.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (219.0 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -1.99 DPS, sim-verified) [world_drop] |
| off_hand | Widowmaker (4091) (or Blue Glittering Axe (7942), Dazzling Mithril Rapier (7944), Nordic Longshank (9401), Speedsteel Rapier (13034), Mithril Stiletto (250611)) | World drop [world_drop] | 22.3 ranged_attack_power points (1.51 DPS) | yes | Blue Glittering Axe (7942, +0.00 DPS) [crafted]; Dazzling Mithril Rapier (7944, +0.00 DPS) [crafted]; Nordic Longshank (9401, +0.00 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (219.0 DPS) | yes | Gryphonwing Long Bow (13022, -0.51 DPS) [dungeon]; Houndmaster's Bow (11628, -2.26 DPS) [dungeon]; Dark Iron Rifle (16004, -3.83 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Sergeant Major's Chain Gauntlets; waist: Sagebrush Girdle; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Widowmaker; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 879, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 276.5. Weights run: 3.3s. Verify run: 10.7s. 1749 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.830 ± 0.016, crit=1.293 ± 0.033 per rating point (14 rating = 1%, 18.095 per %), hit=2.300 ± 0.124 per rating point (10 rating = 1%, 23.004 per %), melee_haste=13.152 ± 1.790

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (276.5 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -5.96 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.8 ranged_attack_power points (2.51 DPS) | yes | Sunprism Pendant (20649, -1.32 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -1.69 DPS) [world]; Kaleidoscope Chain (13084, -1.74 DPS) [world_drop] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 66.9 ranged_attack_power points (4.56 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.25 DPS) | yes | Dark Phantom Cape (13122, -0.88 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -1.94 DPS) [world_drop]; Deep River Cloak (15789, -2.13 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 145.8 ranged_attack_power points (9.93 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.15 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.15 DPS) [vendor]; Shadowcraft Tunic (16721, -4.42 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 42.5 ranged_attack_power points (2.89 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -6.74 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (276.5 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -5.36 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (276.5 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Heavy Timbermaw Belt (19043, +0.00 DPS) [crafted]; Sagebrush Girdle (17778, -5.15 DPS, sim-verified) [quest] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (276.5 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -8.71 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (276.5 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -6.85 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.6 ranged_attack_power points (2.70 DPS) | yes | Falcon's Hook (7552, -0.96 DPS) [dungeon]; Disengagement Ring (276202, -1.16 DPS) [vendor]; Blackstone Ring (17713, -1.34 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 28.3 ranged_attack_power points (1.93 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; Disengagement Ring (276202, -0.39 DPS) [vendor]; Blackstone Ring (17713, -0.57 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (276.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (276.5 DPS) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (276.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -8.75 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (276.5 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -12.00 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; trinket2: Mark of the Chosen; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (dwarf, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 712.0. Weights run: 3.3s. Verify run: 10.0s. 1749 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.903 ± 0.016, crit=1.460 ± 0.031 per rating point (14 rating = 1%, 20.434 per %), hit=2.925 ± 0.181 per rating point (10 rating = 1%, 29.249 per %), melee_haste=19.008 ± 2.444

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (712.0 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -9.67 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.7 ranged_attack_power points (5.32 DPS) | yes | Sunprism Pendant (20649, -0.41 DPS) [quest]; Archmagister's Faceted Pendant (273049, -3.63 DPS) [world]; Kaleidoscope Chain (13084, -3.68 DPS) [world_drop] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 69.7 ranged_attack_power points (9.83 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Spaulders of the Unseen (13116, -3.09 DPS, sim-verified) [world_drop] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.5 ranged_attack_power points (8.96 DPS) | yes | Blackflame Cape (13109, -4.05 DPS) [world_drop]; Dark Phantom Cape (13122, -4.07 DPS, sim-verified) [world_drop]; Deep River Cloak (15789, -4.46 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.1 ranged_attack_power points (22.16 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.35 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.35 DPS) [vendor]; Shadowcraft Tunic (16721, -10.95 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 43.5 ranged_attack_power points (6.14 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -4.74 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 61.0 ranged_attack_power points (8.60 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -7.19 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 43.5 ranged_attack_power points (6.14 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Wicked Leather Belt (15088, -0.41 DPS) [crafted]; Heavy Timbermaw Belt (19043, -5.98 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (712.0 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -8.73 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (712.0 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.39 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.6 ranged_attack_power points (5.73 DPS) | yes | Falcon's Hook (7552, -2.05 DPS) [dungeon]; Disengagement Ring (276202, -2.46 DPS) [vendor]; Field Researcher's Loop (281634, -2.87 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.0 ranged_attack_power points (4.09 DPS) | yes | Falcon's Hook (7552, -0.41 DPS) [dungeon]; Disengagement Ring (276202, -0.82 DPS) [vendor]; Field Researcher's Loop (281634, -1.23 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (712.0 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (712.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (712.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -21.75 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (712.0 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -26.43 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Beaststalker's Boots; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1749, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 110.3. Weights run: 2.9s. Verify run: 2.3s. 223 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.112 ± 0.015, crit=0.605 ± 0.016 per rating point (14 rating = 1%, 8.470 per %), hit=1.109 ± 0.048 per rating point (10 rating = 1%, 11.089 per %), melee_haste=7.909 ± 0.726

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.9 ranged_attack_power points (1.72 DPS) | yes | Red Winter Hat (21524, -1.86 DPS, sim-verified) [dungeon] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.2 ranged_attack_power points (0.43 DPS) | yes | Roadwatcher's Confidence (281265, -0.69 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Slime-encrusted Pads (6461, -0.63 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.7 ranged_attack_power points (1.29 DPS) | yes | Cape of the Brotherhood (5193, -0.42 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.43 DPS) [world_drop]; Bristlebark Cape (14571, -0.43 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 21.8 ranged_attack_power points (1.51 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Dark Leather Tunic (2317, -0.22 DPS) [crafted]; Prospector's Chestpiece (14562, -0.22 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.6 ranged_attack_power points (1.08 DPS) | yes | Wolf Bracers (4794, -0.22 DPS) [vendor]; Bristlebark Bindings (14569, -0.43 DPS) [world_drop]; Spare Part Bindings (279875, -0.43 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.43 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.43 DPS) [crafted]; Serpent Gloves (5970, -1.43 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.25 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Guardsman Belt (3429, -0.38 DPS) [world]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 28.0 ranged_attack_power points (1.94 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.22 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Gray Leather Boots (2315, -0.22 DPS) [crafted]; Feet of the Lynx (1121, -1.12 DPS, sim-verified) [world_drop] |
| finger1 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 9.3 ranged_attack_power points (0.65 DPS) | yes | The 1 Ring (8350, -0.43 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.2 ranged_attack_power points (0.43 DPS) | yes | The 1 Ring (8350, -0.97 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.9 ranged_attack_power points (2.14 DPS) | yes | Impaling Harpoon (5200, -0.20 DPS) [dungeon]; Crescent Staff (6505, -0.63 DPS) [quest]; Lupine Axe (1220, -0.85 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (12.42 DPS) | yes | Lil Timmy's Peashooter (13136, -1.52 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -3.00 DPS) [crafted]; Venomstrike (6469, -3.13 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Bounty Hunter's Ring; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 223, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 139.7. Weights run: 3.0s. Verify run: 2.5s. 431 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.894 ± 0.013, crit=0.674 ± 0.017 per rating point (14 rating = 1%, 9.436 per %), hit=1.207 ± 0.058 per rating point (10 rating = 1%, 12.067 per %), melee_haste=11.026 ± 0.959

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 28.9 ranged_attack_power points (1.97 DPS) | yes | Brawler's Leather Hood (252504, -0.44 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.59 DPS) [world]; Resilient Cap (14401, -1.58 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 ranged_attack_power points (0.82 DPS) | yes | Erudite's Amulet (277204, -0.42 DPS) [quest]; Kaleidoscope Chain (13084, -0.47 DPS, sim-verified) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.8 ranged_attack_power points (2.17 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.79 DPS) [crafted]; Insignia Mantle (4721, -0.79 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.39 DPS) [dungeon]; Swiftrunner Cape (6745, -0.39 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Panther Armor (6670, -0.99 DPS) [quest]; Garb of Florid Feathers (273026, -0.99 DPS) [dungeon]; Green Leather Armor (4255, -1.18 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Pilferer's Gloves (7358, -0.20 DPS) [crafted]; Braced Handguards (6784, -0.39 DPS) [quest]; Insignia Gloves (6408, -0.59 DPS) [world_drop] |
| waist | Skulker's Leather Belt (252520) (or Stalker's Leather Belt (252521)) | Leatherworking [crafted] | 26.1 ranged_attack_power points (1.78 DPS) | yes | Stalker's Leather Belt (252521, +0.00 DPS) [crafted]; Deftkin Belt (16659, -0.17 DPS) [quest]; Blackened Defias Belt (10403, -0.55 DPS) [dungeon] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 40.5 ranged_attack_power points (2.76 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -0.99 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 23.2 ranged_attack_power points (1.58 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.20 DPS) [quest] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Band of the Fist (17694, -0.39 DPS) [quest]; Thunderbrow Ring (13097, -0.59 DPS) [world_drop] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 14.5 ranged_attack_power points (0.99 DPS) | yes | Band of the Fist (17694, -0.20 DPS) [quest]; Thunderbrow Ring (13097, -0.39 DPS) [world_drop]; Defias Renegade Ring (1076, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) | Razorfen Kraul [dungeon] | 17.4 ranged_attack_power points (1.18 DPS) | yes | Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Outlaw Sabre (16886) | Baron Aquanis [quest] | 16.0 ranged_attack_power points (1.09 DPS) | yes | Jade Serpentblade (3850, +0.00 DPS) [crafted] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 249.8 ranged_attack_power points (17.04 DPS) | yes | Ironweaver (13137, -0.95 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.07 DPS) [world_drop]; Chesterfall Musket (7729, -1.26 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Skulker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Vendetta; off_hand: Outlaw Sabre; ranged: Golemsight Long Gun

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 184.5. Weights run: 3.3s. Verify run: 3.0s. 678 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.808 ± 0.013, crit=0.752 ± 0.020 per rating point (14 rating = 1%, 10.534 per %), hit=1.307 ± 0.095 per rating point (10 rating = 1%, 13.070 per %), melee_haste=12.121 ± 1.404

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Nightscape Headband (8176, -0.18 DPS) [crafted]; Guard's Chain Helm (250499, -0.18 DPS) [crafted]; Skullsplitter Helm (1624, -0.37 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 ranged_attack_power points (0.79 DPS) | yes | Kaleidoscope Chain (13084, -0.05 DPS) [world_drop]; Erudite's Amulet (277204, -0.42 DPS) [quest]; Ethereal Talisman (4430, -0.68 DPS, sim-verified) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 30.9 ranged_attack_power points (2.03 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.18 DPS) [dungeon]; Flintrock Shoulders (7755, -0.18 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.18 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 67.4 ranged_attack_power points (4.42 DPS) | yes | Wolffear Harness (13110, -1.51 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.66 DPS) [crafted]; Tough Scorpid Breastplate (8203, -1.66 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.5 ranged_attack_power points (1.47 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.18 DPS) [crafted] |
| hands | Gauntlets of Divinity (7724) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 32.0 ranged_attack_power points (2.10 DPS) | yes | Gloves of Holy Might (867, -0.10 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.26 DPS) [crafted]; Scarlet Gauntlets (10331, -0.26 DPS) [dungeon] |
| waist | Scorpashi Sash (14652) | World drop [world_drop] | 28.1 ranged_attack_power points (1.84 DPS) | yes | Blackforge Girdle (6425, -0.18 DPS) [dungeon]; Ogron's Sash (13117, -0.18 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.18 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 59.0 ranged_attack_power points (3.87 DPS) | yes | Triprunner Dungarees (9624, -1.03 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.29 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.29 DPS) [world_drop] |
| feet | Worn Running Boots (9398) (or Swampwalker Boots (2276)) | Uldaman: Eric \"The Swift\" [dungeon] | 36.5 ranged_attack_power points (2.40 DPS) | yes | Swampwalker Boots (2276, +0.00 DPS) [world_drop]; Imperial Leather Boots (6431, -0.37 DPS) [dungeon]; Dusky Boots (7390, -0.37 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.1 ranged_attack_power points (1.84 DPS) | yes | Disengagement Ring (276202, -0.37 DPS) [vendor]; Mark of Kern (2262, -0.53 DPS) [dungeon]; Assault Band (13095, -0.53 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.3 ranged_attack_power points (1.66 DPS) | yes | Disengagement Ring (276202, -0.18 DPS) [vendor]; Mark of Kern (2262, -0.35 DPS) [dungeon]; Assault Band (13095, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Steel Spear (250605) | Blacksmithing [crafted] | sim-verified (184.5 DPS) | yes | Hellslayer Battle Axe (13017, -0.82 DPS) [world_drop]; Armor Piercer (6679, -1.03 DPS) [dungeon]; Gut Ripper (2164, -3.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (22.98 DPS) | yes | Swiftwind (13038, -2.16 DPS) [world_drop]; Baelog's Shortbow (9400, -2.20 DPS) [dungeon]; The Silencer (13138, -3.36 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gauntlets of Divinity; waist: Scorpashi Sash; legs: Basilisk Hide Pants; feet: Worn Running Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Steel Spear; ranged: Bow of Searing Arrows

No-known-source sample (15 of 678, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5420001505001251-0053200000000000-000000000000000000)

Set DPS (verified): 225.8. Weights run: 3.3s. Verify run: 3.6s. 868 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.784 ± 0.013, crit=0.920 ± 0.024 per rating point (14 rating = 1%, 12.877 per %), hit=1.562 ± 0.105 per rating point (10 rating = 1%, 15.618 per %), melee_haste=10.622 ± 1.555

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 47.3 ranged_attack_power points (3.21 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.38 DPS) [quest]; Tough Scorpid Helm (8208, -0.57 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.2 ranged_attack_power points (2.46 DPS) | yes | Woven Ivy Necklace (19159, -1.45 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -1.64 DPS) [world]; Ethereal Talisman (4430, -1.70 DPS) [quest] |
| shoulder | Khan's Mantle (14787) | World drop [world_drop] | 33.4 ranged_attack_power points (2.27 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -0.19 DPS) [world_drop]; Nightscape Shoulders (8192, -0.19 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Serpentskin Cloak (8259, -0.76 DPS) [world_drop]; Nightscape Cloak (8195, -0.94 DPS) [crafted]; Blackflame Cape (13109, -1.35 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 66.8 ranged_attack_power points (4.54 DPS) | yes | Blazewind Breastplate (11193, -0.19 DPS) [quest]; Stone Guard's Chain Armor (220827, -0.83 DPS) [vendor]; Charred Leather Tunic (19127, -1.13 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Bloodlust Bracelets (14807, -0.76 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.76 DPS) [crafted]; Arena Bands (18711, -0.93 DPS) [world] |
| hands | First Sergeant's Chain Gauntlets (220830) | PvP rank 9 · First Sergeant · Horde [vendor] | 44.5 ranged_attack_power points (3.02 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.38 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -0.38 DPS) [crafted]; Heavy Scorpid Gauntlet (15078, -0.76 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 41.8 ranged_attack_power points (2.83 DPS) | yes | Stalker's Mail Belt (252588, -0.19 DPS) [crafted]; Seared Mail Girdle (19125, -0.57 DPS) [quest]; Skulker's Leather Waistguard (252474, -1.97 DPS, sim-verified) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.5 ranged_attack_power points (3.97 DPS) | yes | Keeper's Woolies (14668, -0.38 DPS) [world_drop]; Stone Guard's Chain Legplates (220833, -0.45 DPS) [vendor]; Oilskin Leggings (9414, -0.57 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 55.7 ranged_attack_power points (3.78 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.38 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -0.94 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.0 ranged_attack_power points (2.65 DPS) | yes | Falcon's Hook (7552, -0.94 DPS) [dungeon]; White Bone Band (11862, -1.02 DPS) [quest]; Disengagement Ring (276202, -1.13 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.8 ranged_attack_power points (1.89 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; White Bone Band (11862, -0.26 DPS) [quest]; Disengagement Ring (276202, -0.38 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (225.8 DPS) | yes | Molten Heart of the Mountain (249470, -7.24 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (225.8 DPS) | yes | Molten Heart of the Mountain (249470, -0.78 DPS, sim-verified) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (225.8 DPS) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Force of the Hippogryph (9684) | Weapons of Spirit [quest] | 25.1 ranged_attack_power points (1.70 DPS) | yes | Widowmaker (4091, +0.00 DPS, sim-verified) [world_drop]; White Bone Shredder (11863, -0.38 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (225.8 DPS) | yes | Gryphonwing Long Bow (13022, -0.51 DPS) [dungeon]; Houndmaster's Bow (11628, -2.26 DPS) [dungeon]; Dark Iron Rifle (16004, -2.80 DPS, sim-verified) [crafted] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; shoulder: Khan's Mantle; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: First Sergeant's Chain Gauntlets; waist: Sagebrush Girdle; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Force of the Hippogryph; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 288.0. Weights run: 3.3s. Verify run: 10.5s. 1746 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.830 ± 0.016, crit=1.293 ± 0.033 per rating point (14 rating = 1%, 18.095 per %), hit=2.300 ± 0.124 per rating point (10 rating = 1%, 23.004 per %), melee_haste=13.152 ± 1.790

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (288.0 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -6.07 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.8 ranged_attack_power points (2.51 DPS) | yes | Sunprism Pendant (20649, -0.19 DPS) [quest]; Woven Ivy Necklace (19159, -0.77 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.69 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 66.9 ranged_attack_power points (4.56 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS, sim-verified) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 62.5 ranged_attack_power points (4.25 DPS) | yes | Dark Phantom Cape (13122, -1.36 DPS) [world_drop]; Blackflame Cape (13109, -1.94 DPS) [world_drop]; Deep River Cloak (15789, -2.13 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 145.8 ranged_attack_power points (9.93 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.15 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.15 DPS) [vendor]; Shadowcraft Tunic (16721, -3.19 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 42.5 ranged_attack_power points (2.89 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -6.95 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (288.0 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Mourning Handguards (279257, -5.39 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (288.0 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Heavy Timbermaw Belt (19043, +0.00 DPS) [crafted]; Sagebrush Girdle (17778, -5.56 DPS, sim-verified) [quest] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (288.0 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -7.52 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (288.0 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -5.73 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.6 ranged_attack_power points (2.70 DPS) | yes | Falcon's Hook (7552, -0.96 DPS) [dungeon]; White Bone Band (11862, -1.06 DPS) [quest]; Disengagement Ring (276202, -1.16 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 28.3 ranged_attack_power points (1.93 DPS) | yes | Falcon's Hook (7552, -0.19 DPS) [dungeon]; White Bone Band (11862, -0.29 DPS) [quest]; Disengagement Ring (276202, -0.39 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (288.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (288.0 DPS) | yes | Molten Heart of the Mountain (249470, -2.18 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (288.0 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -9.50 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (288.0 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -8.56 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5420001505001251-0053502001000000-400000000000000000)

Set DPS (verified): 734.0. Weights run: 3.3s. Verify run: 10.4s. 1746 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.903 ± 0.016, crit=1.460 ± 0.031 per rating point (14 rating = 1%, 20.434 per %), hit=2.925 ± 0.181 per rating point (10 rating = 1%, 29.249 per %), melee_haste=19.008 ± 2.444

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | sim-verified (734.0 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -8.92 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.7 ranged_attack_power points (5.32 DPS) | yes | Sunprism Pendant (20649, -0.41 DPS) [quest]; Woven Ivy Necklace (19159, -1.64 DPS) [quest]; Archmagister's Faceted Pendant (273049, -3.63 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 69.7 ranged_attack_power points (9.83 DPS) | yes | Spaulders of the Unseen (13116, +0.00 DPS) [world_drop]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 63.5 ranged_attack_power points (8.96 DPS) | yes | Dark Phantom Cape (13122, -2.32 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.05 DPS) [world_drop]; Deep River Cloak (15789, -4.46 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 157.1 ranged_attack_power points (22.16 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.35 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.35 DPS) [vendor]; Shadowcraft Tunic (16721, -9.76 DPS, sim-verified) [dungeon] |
| wrist | Beaststalker's Bindings (16681) (or Deepfury Bracers (13120)) | Stratholme: Fleshflayer Ghoul [dungeon] | 43.5 ranged_attack_power points (6.14 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -11.17 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (734.0 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Mourning Handguards (279257, -5.99 DPS, sim-verified) [crafted] |
| waist | Beaststalker's Belt (16680) | Blackrock Spire: Smolderthorn Headhunter [dungeon] | sim-verified (734.0 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Heavy Timbermaw Belt (19043, +0.00 DPS) [crafted]; Sagebrush Girdle (17778, -7.52 DPS, sim-verified) [quest] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | sim-verified (734.0 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Windrunner Legguards (13130, -16.12 DPS, sim-verified) [world_drop] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | sim-verified (734.0 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -10.08 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.6 ranged_attack_power points (5.73 DPS) | yes | Falcon's Hook (7552, -2.05 DPS) [dungeon]; White Bone Band (11862, -2.35 DPS) [quest]; Disengagement Ring (276202, -2.46 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.0 ranged_attack_power points (4.09 DPS) | yes | Falcon's Hook (7552, -0.41 DPS) [dungeon]; White Bone Band (11862, -0.71 DPS) [quest]; Disengagement Ring (276202, -0.82 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (734.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (734.0 DPS) | yes | Molten Heart of the Mountain (249470, -5.55 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (734.0 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -21.25 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (734.0 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -20.17 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beaststalker's Bindings; hands: Beaststalker's Gloves; waist: Beaststalker's Belt; legs: Beaststalker's Pants; feet: Beaststalker's Boots; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1746, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

