# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 104.7. Weights run: 2.8s. Verify run: 2.8s. 266 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.429 ± 0.015, crit=0.929 ± 0.016 per rating point (14 rating = 1%, 13.011 per %), hit=1.482 ± 0.041 per rating point (10 rating = 1%, 14.816 per %), melee_haste=8.515 ± 0.609, intellect=1.591 ± 0.035, spell_power=0.275 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.10 DPS) | yes | Trapper's Leather Hood (252505, -1.40 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.49 DPS) [crafted]; Totemic Leather Hood (252448, -1.91 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.9 ranged_attack_power points (0.52 DPS) | yes | Tarnished Locket (279870, -0.57 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.31 DPS) | yes | Magician's Mantle (12998, -0.11 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -0.72 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.23 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.6 ranged_attack_power points (1.57 DPS) | yes | Cape of the Brotherhood (5193, -0.26 DPS) [dungeon]; Sentry Cloak (2059, -0.52 DPS) [world_drop]; Bristlebark Cape (14571, -0.52 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 37.7 ranged_attack_power points (2.88 DPS) | yes | Trapper's Leather Armor (252491, -0.55 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -1.05 DPS) [crafted]; Mystic's Wrap (14369, -1.25 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.31 DPS) | yes | Wolf Bracers (4794, -0.26 DPS) [vendor]; Bristlebark Bindings (14569, -0.52 DPS) [world_drop]; Bravo's Armbands (270015, -0.52 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 22.5 ranged_attack_power points (1.72 DPS) | yes | Gloves of the Fang (10413, -0.15 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.31 DPS) [crafted]; Riding Gloves (1304, -0.57 DPS) [quest] |
| waist | Trapper's Leather Belt (252429) | Leatherworking [crafted] | 18.5 ranged_attack_power points (1.41 DPS) | yes | Blackened Defias Belt (10403, -0.04 DPS) [dungeon]; Deviate Scale Belt (6468, -0.10 DPS) [crafted]; Dark Leather Belt (4249, -0.36 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 37.2 ranged_attack_power points (2.85 DPS) | yes | Brawler's Leather Pants (252500, -0.49 DPS) [crafted]; Bluegill Breeches (3022, -0.75 DPS) [world]; Leggings of the Fang (10410, -0.84 DPS, sim-verified) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.4 ranged_attack_power points (2.10 DPS) | yes | Trapper's Leather Boots (252440, -0.45 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.52 DPS) [dungeon]; Footpads of the Fang (10411, -0.52 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.4 ranged_attack_power points (1.25 DPS) | yes | Volcanic Rock Ring (12053, -0.89 DPS) [world_drop]; Minor Channeling Ring (1449, -0.91 DPS) [quest]; Furlfeather Ring (286745, -1.23 DPS) [world] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 5.0 ranged_attack_power points (0.38 DPS) | yes | Volcanic Rock Ring (12053, -0.02 DPS) [world_drop]; Minor Channeling Ring (1449, -0.04 DPS) [quest]; Furlfeather Ring (286745, -0.36 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 38.8 ranged_attack_power points (2.97 DPS) | yes | Bronze Dory (250603, -0.54 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.39 DPS) [world]; Piercing Axe (6094, -1.39 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.74 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS) [world_drop]; Venomstrike (6469, -1.06 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -3.33 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Trapper's Leather Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: The 1 Ring; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 266, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 134.1. Weights run: 2.9s. Verify run: 1.9s. 473 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.259 ± 0.013, crit=1.109 ± 0.018 per rating point (14 rating = 1%, 15.530 per %), hit=1.656 ± 0.046 per rating point (10 rating = 1%, 16.563 per %), melee_haste=9.752 ± 0.785, intellect=1.547 ± 0.041, spell_power=0.258 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 32.6 ranged_attack_power points (2.43 DPS) | yes | Brawler's Leather Hood (252504, -0.53 DPS, sim-verified) [crafted]; Enduring Cap (3020, -0.58 DPS) [world_drop]; Resilient Cap (14401, -0.68 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 21.3 ranged_attack_power points (1.59 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Erudite's Amulet (277204, -1.10 DPS) [quest]; Crystal Starfire Medallion (5003, -1.13 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 35.8 ranged_attack_power points (2.67 DPS) | yes | Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.97 DPS) [crafted]; Insignia Mantle (4721, -0.97 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.1 ranged_attack_power points (1.94 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.49 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.51 DPS) [pvp] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 52.5 ranged_attack_power points (3.92 DPS) | yes | Dusky Leather Armor (7374, -0.52 DPS) [crafted]; Green Leather Armor (4255, -1.05 DPS) [crafted]; Tunic of Westfall (2041, -1.24 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.6 ranged_attack_power points (1.46 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.24 DPS) [world]; Forest Leather Bracers (3202, -0.24 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.3 ranged_attack_power points (2.19 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.29 DPS) [dungeon]; Serpent Gloves (5970, -0.59 DPS) [dungeon]; Pilferer's Gloves (7358, -0.74 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 38.6 ranged_attack_power points (2.88 DPS) | yes | Green Leather Belt (4257, -0.73 DPS) [crafted]; Skulker's Leather Belt (252520, -0.85 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -0.96 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 45.6 ranged_attack_power points (3.40 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Dusky Leather Leggings (7373, -0.24 DPS) [crafted] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 26.1 ranged_attack_power points (1.94 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Trapper's Leather Boots (252440, -0.27 DPS) [crafted]; Footpads of the Fang (10411, -0.49 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.6 ranged_attack_power points (1.46 DPS) | yes | Seal of Wrynn (2933, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Defias Renegade Ring (1076, -0.49 DPS) [dungeon] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 16.3 ranged_attack_power points (1.21 DPS) | yes | Seal of Wrynn (2933, -0.02 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Defias Renegade Ring (1076, -0.24 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | 49.2 ranged_attack_power points (3.67 DPS) | yes | Kam's Walking Stick (2280, -0.89 DPS) [dungeon]; Impaling Harpoon (5200, -0.90 DPS) [dungeon]; Armor Piercer (6679, -0.99 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 252.0 ranged_attack_power points (18.79 DPS) | yes | Ironweaver (13137, -0.70 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -0.76 DPS) [world_drop]; Harpyclaw Short Bow (13019, -1.51 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 185.9. Weights run: 3.1s. Verify run: 3.2s. 728 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.270 ± 0.037, crit=1.671 ± 0.057 per rating point (14 rating = 1%, 23.393 per %), hit=2.044 ± 0.102 per rating point (10 rating = 1%, 20.437 per %), melee_haste=11.247 ± 1.120, intellect=1.641 ± 0.056, spell_power=0.206 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 47.4 ranged_attack_power points (3.52 DPS) | yes | Nightscape Headband (8176, -0.59 DPS, sim-verified) [crafted]; Guard's Chain Helm (250499, -0.61 DPS) [crafted]; Skullsplitter Helm (1624, -0.85 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 21.8 ranged_attack_power points (1.62 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -0.54 DPS) [dungeon]; Necklace of Calisea (1714, -0.77 DPS) [dungeon] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.0 ranged_attack_power points (2.67 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 29.5 ranged_attack_power points (2.19 DPS) | yes | Imperial Cloak (6432, -0.25 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.25 DPS) [world_drop]; Parachute Cloak (10518, -0.88 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 98.9 ranged_attack_power points (7.35 DPS) | yes | Tough Scorpid Breastplate (8203, -2.81 DPS, sim-verified) [crafted]; Wolffear Harness (13110, -3.22 DPS) [world_drop]; Garb of Florid Feathers (273026, -3.33 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 32.7 ranged_attack_power points (2.43 DPS) | yes | Imperial Leather Bracers (4061, -0.49 DPS) [dungeon]; Dusky Bracers (7378, -0.49 DPS) [crafted]; Green Leather Bracers (4259, -0.72 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | 47.5 ranged_attack_power points (3.53 DPS) | yes | Stalker's Leather Gloves (252526, -0.24 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.33 DPS) [crafted]; Aurora Gloves (4042, -0.61 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.3 ranged_attack_power points (2.92 DPS) | yes | Scorpashi Sash (14652, -0.49 DPS) [world_drop]; Green Leather Belt (4257, -0.73 DPS) [crafted]; Blackforge Girdle (6425, -0.73 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 68.7 ranged_attack_power points (5.10 DPS) | yes | Triprunner Dungarees (9624, -1.04 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.70 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.70 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 47.5 ranged_attack_power points (3.53 DPS) | yes | Swampwalker Boots (2276, -0.37 DPS) [world_drop]; Worn Running Boots (9398, -0.62 DPS, sim-verified) [dungeon]; Imperial Leather Boots (6431, -0.85 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.7 ranged_attack_power points (2.43 DPS) | yes | Disengagement Ring (276202, -0.49 DPS) [vendor]; Field Researcher's Loop (281634, -0.73 DPS) [quest]; Mark of Kern (2262, -0.94 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 29.4 ranged_attack_power points (2.19 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Field Researcher's Loop (281634, -0.49 DPS) [quest]; Mark of Kern (2262, -0.70 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (185.9 DPS) | yes | Felbough Arm (273037, -0.77 DPS) [world]; Steel Spear (250605, -0.81 DPS) [crafted]; Gut Ripper (2164, -8.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (185.9 DPS) | yes | Swiftwind (13038, -2.20 DPS) [world_drop]; Baelog's Shortbow (9400, -2.49 DPS) [dungeon]; The Silencer (13138, -3.58 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 728, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band

### Band 50 (dwarf, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 236.0. Weights run: 3.2s. Verify run: 2.8s. 917 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.078 ± 0.031, crit=1.734 ± 0.057 per rating point (14 rating = 1%, 24.271 per %), hit=2.628 ± 0.111 per rating point (10 rating = 1%, 26.279 per %), melee_haste=9.796 ± 1.174, intellect=1.920 ± 0.063, spell_power=0.293 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 70.0 ranged_attack_power points (5.41 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.85 DPS, sim-verified) [dungeon]; Helm of Fire (8348, -1.37 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.0 ranged_attack_power points (3.10 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.28 DPS) [world]; Kaleidoscope Chain (13084, -1.55 DPS) [world_drop]; Horizon Choker (13085, -1.65 DPS, sim-verified) [world_drop] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 50.0 ranged_attack_power points (3.87 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.21 DPS) [crafted]; Ironfeather Shoulders (15067, -0.56 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 46.2 ranged_attack_power points (3.57 DPS) | yes | Serpentskin Cloak (8259, -0.95 DPS) [world_drop]; Nightscape Cloak (8195, -1.19 DPS) [crafted]; Blackflame Cape (13109, -1.43 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 100.2 ranged_attack_power points (7.75 DPS) | yes | Knight's Chain Armor (220828, -1.24 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.27 DPS) [quest]; Tough Scorpid Breastplate (8203, -3.14 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.2 ranged_attack_power points (3.57 DPS) | yes | Stalker's Mail Bracers (252570, -0.54 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.66 DPS) [dungeon]; Bloodlust Bracelets (14807, -0.95 DPS) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | 60.4 ranged_attack_power points (4.67 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -0.86 DPS) [vendor]; Tough Scorpid Gloves (8204, -0.95 DPS) [crafted]; Hibernal Gloves (8110, -1.51 DPS, sim-verified) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 69.2 ranged_attack_power points (5.35 DPS) | yes | Stalker's Mail Belt (252588, -0.54 DPS) [crafted]; Skulker's Leather Waistguard (252474, -2.02 DPS) [crafted]; Stalker's Leather Belt (252521, -2.32 DPS) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 86.6 ranged_attack_power points (6.70 DPS) | yes | Tough Scorpid Leggings (8206, -1.16 DPS) [crafted]; Imperial Red Pants (8251, -1.58 DPS, sim-verified) [world_drop]; Searingscale Leggings (11749, -1.67 DPS) [dungeon] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 70.0 ranged_attack_power points (5.42 DPS) | yes | Albino Crocscale Boots (17728, -0.65 DPS) [dungeon]; Rancor Boots (11865, -0.92 DPS) [quest]; Stalker's Mail Sabatons (252576, -0.98 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.3 ranged_attack_power points (3.58 DPS) | yes | Ring of the Underwood (2951, -1.20 DPS) [world_drop]; Falcon's Hook (7552, -1.44 DPS) [dungeon]; Disengagement Ring (276202, -1.68 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.1 ranged_attack_power points (3.33 DPS) | yes | Falcon's Hook (7552, -1.19 DPS) [dungeon]; Ring of the Underwood (2951, -1.29 DPS, sim-verified) [world_drop]; Disengagement Ring (276202, -1.43 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (236.0 DPS) | yes | Uther's Strength (11302, -1.69 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (236.0 DPS) | yes | Uther's Strength (11302, -1.81 DPS, sim-verified) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (236.0 DPS) | yes | Staff of Jordan (873, -0.08 DPS) [world_drop]; Kindling Stave (11750, -1.08 DPS) [dungeon]; Hanzo Sword (8190, -7.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (236.0 DPS) | yes | Megashot Rifle (17717, -2.51 DPS) [dungeon]; Dark Iron Rifle (16004, -2.97 DPS) [crafted]; Gryphonwing Long Bow (13022, -3.58 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 917, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60 (dwarf, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 301.1. Weights run: 3.2s. Verify run: 6.2s. 1788 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.164 ± 0.039, crit=2.328 ± 0.077 per rating point (14 rating = 1%, 32.589 per %), hit=3.327 ± 0.144 per rating point (10 rating = 1%, 33.269 per %), melee_haste=11.525 ± 1.615, intellect=3.358 ± 0.084, spell_power=0.293 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 226.6 ranged_attack_power points (17.39 DPS) | yes | Field Marshal's Chain Helm (231580, -3.54 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -4.91 DPS) [vendor]; Outlaw's Collar (279253, -10.30 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 47.0 ranged_attack_power points (3.61 DPS) | yes | Opaline Medallion (16623, -0.26 DPS) [quest]; Skibi's Pendant (13089, -0.45 DPS) [world_drop]; Sunprism Pendant (20649, -0.69 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 88.4 ranged_attack_power points (6.78 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Spaulders of the Unseen (13116, -1.82 DPS, sim-verified) [world_drop] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 67.5 ranged_attack_power points (5.18 DPS) | yes | Sergeant Major's Cape (16337, -1.17 DPS) [pvp]; Dark Phantom Cape (13122, -1.53 DPS) [world_drop]; Cerise Drape (15804, -2.17 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 182.9 ranged_attack_power points (14.04 DPS) | yes | Field Marshal's Chain Breastplate (16466, -0.19 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -0.19 DPS) [vendor]; Knight's Chain Armor (220828, -6.24 DPS, sim-verified) [vendor] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 61.0 ranged_attack_power points (4.68 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.56 DPS) [pvp]; Magiskull Cuffs (13107, -1.34 DPS, sim-verified) [world_drop] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 113.5 ranged_attack_power points (8.71 DPS) | yes | Marshal's Chain Vices (231578, -0.08 DPS) [vendor]; Marshal's Chain Grips (231560, -1.06 DPS) [pvp]; Hero's Gauntlets (8305, -2.27 DPS, sim-verified) [dungeon] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 87.8 ranged_attack_power points (6.73 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Stalker's Mail Belt (252588, -0.76 DPS) [crafted]; Knight-Captain's Chain Girdle (16400, -1.02 DPS) [pvp] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | 115.8 ranged_attack_power points (8.89 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, +0.00 DPS) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 87.4 ranged_attack_power points (6.71 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 53.3 ranged_attack_power points (4.09 DPS) | yes | Glowing Crystal Ring (18402, -1.00 DPS) [quest]; Band of the Hierophant (13096, -1.51 DPS) [world_drop]; Frigid Ring (18679, -1.51 DPS) [world] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 44.3 ranged_attack_power points (3.40 DPS) | yes | Glowing Crystal Ring (18402, -0.31 DPS) [quest]; Band of the Hierophant (13096, -0.82 DPS) [world_drop]; Frigid Ring (18679, -0.82 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (301.1 DPS) | yes | Uther's Strength (11302, -2.16 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (301.1 DPS) | yes | Uther's Strength (11302, -1.13 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (301.1 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -10.47 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (301.1 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -3.41 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Horizon Choker; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Timmy's Galoshes; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1788, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60, raid preset (dwarf, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 658.7. Weights run: 3.1s. Verify run: 6.0s. 1788 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.407 ± 0.037, crit=2.621 ± 0.079 per rating point (14 rating = 1%, 36.690 per %), hit=3.947 ± 0.192 per rating point (10 rating = 1%, 39.473 per %), melee_haste=17.711 ± 2.318, intellect=3.437 ± 0.096, spell_power=0.873 ± 0.002

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 238.9 ranged_attack_power points (32.56 DPS) | yes | Field Marshal's Chain Helm (231580, -6.15 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -9.13 DPS) [vendor]; Outlaw's Collar (279253, -13.23 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 48.1 ranged_attack_power points (6.56 DPS) | yes | Archlight Talisman (15856, -0.45 DPS) [quest]; Opaline Medallion (16623, -0.47 DPS) [quest]; Skibi's Pendant (13089, -0.52 DPS) [world_drop] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 97.8 ranged_attack_power points (13.33 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 71.1 ranged_attack_power points (9.69 DPS) | yes | Dark Phantom Cape (13122, -2.73 DPS) [world_drop]; Sergeant Major's Cape (16337, -3.84 DPS, sim-verified) [pvp]; Cerise Drape (15804, -4.09 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 202.2 ranged_attack_power points (27.57 DPS) | yes | Field Marshal's Chain Breastplate (16466, -1.15 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -1.15 DPS) [vendor]; Spitfire Breastplate (20479, -15.07 DPS, sim-verified) [crafted] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 68.6 ranged_attack_power points (9.35 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.45 DPS) [pvp]; Magiskull Cuffs (13107, -2.59 DPS, sim-verified) [world_drop] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 119.7 ranged_attack_power points (16.31 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -1.79 DPS) [pvp]; Spitfire Gauntlets (20480, -3.75 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 92.4 ranged_attack_power points (12.59 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Stalker's Mail Belt (252588, -1.40 DPS) [crafted]; Knight-Captain's Chain Girdle (16400, -1.87 DPS) [pvp] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 131.7 ranged_attack_power points (17.96 DPS) | yes | Marshal's Chain Legplates (231558, +0.00 DPS) [vendor]; Marshal's Chain Legguards (231577, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -2.39 DPS, sim-verified) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 92.3 ranged_attack_power points (12.58 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 59.5 ranged_attack_power points (8.11 DPS) | yes | Glowing Crystal Ring (18402, -2.48 DPS) [quest]; Maiden's Circle (13001, -2.69 DPS) [world_drop]; Band of the Penitent (13217, -3.11 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 47.7 ranged_attack_power points (6.50 DPS) | yes | Maiden's Circle (13001, -1.08 DPS) [world_drop]; Band of the Penitent (13217, -1.50 DPS) [quest]; Glowing Crystal Ring (18402, -2.45 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (658.7 DPS) | yes | Uther's Strength (11302, -4.13 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (658.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (658.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -20.57 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (658.7 DPS) | yes | Eaglehorn Long Bow (13023, +0.00 DPS) [world_drop]; Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Horizon Choker; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Ironfeather Leggings; feet: Timmy's Galoshes; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1788, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 106.5. Weights run: 2.8s. Verify run: 2.5s. 262 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.429 ± 0.015, crit=0.929 ± 0.016 per rating point (14 rating = 1%, 13.011 per %), hit=1.482 ± 0.041 per rating point (10 rating = 1%, 14.816 per %), melee_haste=8.515 ± 0.609, intellect=1.591 ± 0.035, spell_power=0.275 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.10 DPS) | yes | Shadow Goggles (4373, -1.49 DPS) [crafted]; Trapper's Leather Hood (252505, -1.60 DPS, sim-verified) [crafted]; Totemic Leather Hood (252448, -1.91 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.9 ranged_attack_power points (0.52 DPS) | yes | Roadwatcher's Confidence (281265, -0.31 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 17.1 ranged_attack_power points (1.31 DPS) | yes | Magician's Mantle (12998, -0.63 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.72 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.23 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 20.6 ranged_attack_power points (1.57 DPS) | yes | Cape of the Brotherhood (5193, -0.26 DPS) [dungeon]; Sentry Cloak (2059, -0.52 DPS) [world_drop]; Bristlebark Cape (14571, -0.52 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 28.8 ranged_attack_power points (2.20 DPS) | yes | Brawler's Leather Armor (252490, -0.36 DPS, sim-verified) [crafted]; Mystic's Wrap (14369, -0.56 DPS) [world_drop]; Mystic's Robe (14371, -0.56 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 17.1 ranged_attack_power points (1.31 DPS) | yes | Wolf Bracers (4794, -0.26 DPS) [vendor]; Bristlebark Bindings (14569, -0.52 DPS) [world_drop]; Spare Part Bindings (279875, -0.52 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 22.5 ranged_attack_power points (1.72 DPS) | yes | Gloves of the Fang (10413, -0.15 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.31 DPS) [crafted]; Forest Leather Gloves (3058, -0.67 DPS) [world_drop] |
| waist | Trapper's Leather Belt (252429) | Leatherworking [crafted] | 18.5 ranged_attack_power points (1.41 DPS) | yes | Blackened Defias Belt (10403, -0.04 DPS) [dungeon]; Deviate Scale Belt (6468, -0.10 DPS) [crafted]; Dark Leather Belt (4249, -0.36 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 37.2 ranged_attack_power points (2.85 DPS) | yes | Brawler's Leather Pants (252500, -0.49 DPS) [crafted]; Leggings of the Fang (10410, -0.55 DPS, sim-verified) [dungeon]; Bluegill Breeches (3022, -0.75 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 27.4 ranged_attack_power points (2.10 DPS) | yes | Blackened Defias Boots (10402, -0.52 DPS) [dungeon]; Footpads of the Fang (10411, -0.52 DPS) [dungeon]; Trapper's Leather Boots (252440, -0.57 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.4 ranged_attack_power points (1.25 DPS) | yes | Loop of Sacrifice (281673, -0.65 DPS) [quest]; The 1 Ring (8350, -0.87 DPS) [world]; Volcanic Rock Ring (12053, -0.89 DPS) [world_drop] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 10.3 ranged_attack_power points (0.79 DPS) | yes | The 1 Ring (8350, -0.40 DPS) [world]; Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; Loop of Sacrifice (281673, -0.51 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 38.8 ranged_attack_power points (2.97 DPS) | yes | Crescent Staff (6505, -0.28 DPS) [quest]; Bronze Dory (250603, -0.41 DPS) [crafted]; Staff of Orgrimmar (15444, -0.66 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.7 ranged_attack_power points (13.74 DPS) | yes | Lil Timmy's Peashooter (13136, -0.96 DPS) [world_drop]; Venomstrike (6469, -1.02 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -3.33 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Trapper's Leather Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Bounty Hunter's Ring; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 262, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 136.1. Weights run: 2.9s. Verify run: 2.9s. 470 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.259 ± 0.013, crit=1.109 ± 0.018 per rating point (14 rating = 1%, 15.530 per %), hit=1.656 ± 0.046 per rating point (10 rating = 1%, 16.563 per %), melee_haste=9.752 ± 0.785, intellect=1.547 ± 0.041, spell_power=0.258 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 32.6 ranged_attack_power points (2.43 DPS) | yes | Brawler's Leather Hood (252504, -0.56 DPS, sim-verified) [crafted]; Enduring Cap (3020, -0.58 DPS) [world_drop]; Resilient Cap (14401, -0.68 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 21.3 ranged_attack_power points (1.59 DPS) | yes | Kaleidoscope Chain (13084, -0.15 DPS) [world_drop]; Erudite's Amulet (277204, -1.10 DPS) [quest]; Crystal Starfire Medallion (5003, -1.13 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 35.8 ranged_attack_power points (2.67 DPS) | yes | Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.97 DPS) [crafted]; Insignia Mantle (4721, -0.97 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 26.1 ranged_attack_power points (1.94 DPS) | yes | Hawkeye's Cloak (14593, -0.24 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.49 DPS) [dungeon]; Swiftrunner Cape (6745, -0.49 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 52.5 ranged_attack_power points (3.92 DPS) | yes | Dusky Leather Armor (7374, -0.42 DPS, sim-verified) [crafted]; Green Leather Armor (4255, -1.05 DPS) [crafted]; Trapper's Leather Tunic (252509, -1.28 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 19.6 ranged_attack_power points (1.46 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.24 DPS) [world_drop]; Loamflake Bracers (15462, -0.24 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 29.3 ranged_attack_power points (2.19 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.29 DPS) [dungeon]; Braced Handguards (6784, -0.49 DPS) [quest]; Pilferer's Gloves (7358, -0.54 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 38.6 ranged_attack_power points (2.88 DPS) | yes | Skulker's Leather Belt (252520, -0.47 DPS, sim-verified) [crafted]; Green Leather Belt (4257, -0.73 DPS) [crafted]; Prowler's Leather Belt (252459, -0.96 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 45.6 ranged_attack_power points (3.40 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Dusky Leather Leggings (7373, -0.24 DPS) [crafted] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 26.1 ranged_attack_power points (1.94 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.24 DPS) [quest] |
| finger1 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | 19.6 ranged_attack_power points (1.46 DPS) | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Defias Renegade Ring (1076, -0.49 DPS) [dungeon]; Band of the Fist (17694, -0.49 DPS) [quest] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 16.3 ranged_attack_power points (1.21 DPS) | yes | Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Defias Renegade Ring (1076, -0.24 DPS) [dungeon]; Band of the Fist (17694, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Kam's Walking Stick (2280) | The Stockade: Kam Deepfury [dungeon] | sim-verified (136.1 DPS) | yes | Impaling Harpoon (5200, -0.01 DPS) [dungeon]; Armor Piercer (6679, -0.10 DPS) [dungeon]; Felbough Arm (273037, -1.65 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 252.0 ranged_attack_power points (18.79 DPS) | yes | Ironweaver (13137, -0.42 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -0.76 DPS) [world_drop]; Harpyclaw Short Bow (13019, -1.51 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Monkey Ring; main_hand: Kam's Walking Stick; ranged: Golemsight Long Gun

No-known-source sample (15 of 470, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 185.1. Weights run: 3.1s. Verify run: 3.2s. 716 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.270 ± 0.037, crit=1.671 ± 0.057 per rating point (14 rating = 1%, 23.393 per %), hit=2.044 ± 0.102 per rating point (10 rating = 1%, 20.437 per %), melee_haste=11.247 ± 1.120, intellect=1.641 ± 0.056, spell_power=0.206 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 47.4 ranged_attack_power points (3.52 DPS) | yes | Nightscape Headband (8176, -0.61 DPS) [crafted]; Guard's Chain Helm (250499, -0.61 DPS) [crafted]; Skullsplitter Helm (1624, -0.85 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 21.8 ranged_attack_power points (1.62 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -0.54 DPS) [dungeon]; Ethereal Talisman (4430, -0.65 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.0 ranged_attack_power points (2.67 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.24 DPS) [dungeon]; Flintrock Shoulders (7755, -0.24 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 26.2 ranged_attack_power points (1.94 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 98.9 ranged_attack_power points (7.35 DPS) | yes | Wolffear Harness (13110, -3.22 DPS) [world_drop]; Garb of Florid Feathers (273026, -3.33 DPS) [dungeon]; Tough Scorpid Breastplate (8203, -3.56 DPS, sim-verified) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 32.7 ranged_attack_power points (2.43 DPS) | yes | Green Leather Bracers (4259, -0.36 DPS) [crafted]; Imperial Leather Bracers (4061, -0.49 DPS) [dungeon]; Dusky Bracers (7378, -0.49 DPS) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | 47.5 ranged_attack_power points (3.53 DPS) | yes | Stalker's Leather Gloves (252526, -0.24 DPS) [crafted]; Dragonscale Gauntlets (8347, -0.33 DPS) [crafted]; Aurora Gloves (4042, -0.61 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.3 ranged_attack_power points (2.92 DPS) | yes | Scorpashi Sash (14652, -0.49 DPS) [world_drop]; Green Leather Belt (4257, -0.73 DPS) [crafted]; Blackforge Girdle (6425, -0.73 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 68.7 ranged_attack_power points (5.10 DPS) | yes | Triprunner Dungarees (9624, -0.59 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.70 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.70 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 47.5 ranged_attack_power points (3.53 DPS) | yes | Swampwalker Boots (2276, -0.37 DPS) [world_drop]; Worn Running Boots (9398, -0.37 DPS) [dungeon]; Imperial Leather Boots (6431, -0.85 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.7 ranged_attack_power points (2.43 DPS) | yes | Disengagement Ring (276202, -0.49 DPS) [vendor]; Field Researcher's Loop (281634, -0.73 DPS) [quest]; Mark of Kern (2262, -0.94 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 29.4 ranged_attack_power points (2.19 DPS) | yes | Disengagement Ring (276202, -0.24 DPS) [vendor]; Field Researcher's Loop (281634, -0.49 DPS) [quest]; Mark of Kern (2262, -0.70 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (185.1 DPS) | yes | Felbough Arm (273037, -0.77 DPS) [world]; Steel Spear (250605, -0.81 DPS) [crafted]; Gut Ripper (2164, -7.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (185.1 DPS) | yes | Swiftwind (13038, -2.20 DPS) [world_drop]; Baelog's Shortbow (9400, -2.49 DPS) [dungeon]; The Silencer (13138, -3.25 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 244.4. Weights run: 3.2s. Verify run: 2.7s. 906 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.078 ± 0.031, crit=1.734 ± 0.057 per rating point (14 rating = 1%, 24.271 per %), hit=2.628 ± 0.111 per rating point (10 rating = 1%, 26.279 per %), melee_haste=9.796 ± 1.174, intellect=1.920 ± 0.063, spell_power=0.293 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 70.0 ranged_attack_power points (5.41 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.71 DPS) [dungeon]; Helm of Fire (8348, -1.37 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.0 ranged_attack_power points (3.10 DPS) | yes | Horizon Choker (13085, -1.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.28 DPS) [world]; Woven Ivy Necklace (19159, -1.71 DPS, sim-verified) [quest] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 50.0 ranged_attack_power points (3.87 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.21 DPS) [crafted]; Ironfeather Shoulders (15067, -0.56 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 46.2 ranged_attack_power points (3.57 DPS) | yes | Serpentskin Cloak (8259, -0.95 DPS) [world_drop]; Nightscape Cloak (8195, -1.19 DPS) [crafted]; Blackflame Cape (13109, -1.32 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 100.2 ranged_attack_power points (7.75 DPS) | yes | Stone Guard's Chain Armor (220827, -0.73 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.27 DPS) [quest]; Tough Scorpid Breastplate (8203, -3.14 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.2 ranged_attack_power points (3.57 DPS) | yes | Stalker's Mail Bracers (252570, -0.54 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.66 DPS) [dungeon]; Bloodlust Bracelets (14807, -0.95 DPS) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | 60.4 ranged_attack_power points (4.67 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -0.86 DPS) [vendor]; Tough Scorpid Gloves (8204, -0.95 DPS) [crafted]; Hibernal Gloves (8110, -1.62 DPS, sim-verified) [world_drop] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 69.2 ranged_attack_power points (5.35 DPS) | yes | Stalker's Mail Belt (252588, -0.54 DPS) [crafted]; Skulker's Leather Waistguard (252474, -2.02 DPS) [crafted]; Stalker's Leather Belt (252521, -2.32 DPS) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | PvP rank 10 · Stone Guard · Horde [vendor] | 86.6 ranged_attack_power points (6.70 DPS) | yes | Tough Scorpid Leggings (8206, -1.16 DPS) [crafted]; Imperial Red Pants (8251, -1.24 DPS, sim-verified) [world_drop]; Searingscale Leggings (11749, -1.67 DPS) [dungeon] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 70.0 ranged_attack_power points (5.42 DPS) | yes | Albino Crocscale Boots (17728, -0.65 DPS) [dungeon]; Rancor Boots (11865, -0.92 DPS) [quest]; Stalker's Mail Sabatons (252576, -0.98 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.3 ranged_attack_power points (3.58 DPS) | yes | Ring of the Underwood (2951, -1.20 DPS) [world_drop]; Falcon's Hook (7552, -1.44 DPS) [dungeon]; Disengagement Ring (276202, -1.68 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 43.1 ranged_attack_power points (3.33 DPS) | yes | Falcon's Hook (7552, -1.19 DPS) [dungeon]; Ring of the Underwood (2951, -1.22 DPS, sim-verified) [world_drop]; Disengagement Ring (276202, -1.43 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (244.4 DPS) | yes | Uther's Strength (11302, -7.79 DPS) [world_drop]; Molten Heart of the Mountain (249470, -8.06 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (244.4 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.69 DPS) [world_drop] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-verified (244.4 DPS) | yes | Staff of Jordan (873, -0.08 DPS) [world_drop]; Kindling Stave (11750, -1.08 DPS) [dungeon]; Hanzo Sword (8190, -7.80 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (244.4 DPS) | yes | Megashot Rifle (17717, -2.51 DPS) [dungeon]; Dark Iron Rifle (16004, -2.97 DPS) [crafted]; Gryphonwing Long Bow (13022, -4.58 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 906, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 308.6. Weights run: 3.2s. Verify run: 6.1s. 1785 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.164 ± 0.039, crit=2.328 ± 0.077 per rating point (14 rating = 1%, 32.589 per %), hit=3.327 ± 0.144 per rating point (10 rating = 1%, 33.269 per %), melee_haste=11.525 ± 1.615, intellect=3.358 ± 0.084, spell_power=0.293 ± 0.001

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 226.6 ranged_attack_power points (17.39 DPS) | yes | Warlord's Chain Helmet (16566, -3.54 DPS) [vendor]; Warlord's Chain Helm (231571, -3.54 DPS) [vendor]; Outlaw's Collar (279253, -10.27 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 47.0 ranged_attack_power points (3.61 DPS) | yes | Sunprism Pendant (20649, -0.69 DPS) [quest]; Archlight Talisman (15856, -0.76 DPS) [quest]; Skibi's Pendant (13089, -1.81 DPS, sim-verified) [world_drop] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 88.4 ranged_attack_power points (6.78 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Spaulders of the Unseen (13116, -1.95 DPS, sim-verified) [world_drop] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 67.5 ranged_attack_power points (5.18 DPS) | yes | Dark Phantom Cape (13122, -1.39 DPS, sim-verified) [world_drop]; Cerise Drape (15804, -2.17 DPS) [quest]; Blackflame Cape (13109, -2.26 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 182.9 ranged_attack_power points (14.04 DPS) | yes | Warlord's Chain Chestpiece (16565, -0.19 DPS) [vendor]; Warlord's Chain Hauberk (231573, -0.19 DPS) [vendor]; Stone Guard's Chain Armor (220827, -7.84 DPS, sim-verified) [vendor] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 61.0 ranged_attack_power points (4.68 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.56 DPS) [pvp]; Magiskull Cuffs (13107, -2.25 DPS, sim-verified) [world_drop] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 113.5 ranged_attack_power points (8.71 DPS) | yes | General's Chain Gloves (16571, -0.08 DPS) [vendor]; General's Chain Vices (231575, -0.08 DPS) [vendor]; Hero's Gauntlets (8305, -3.45 DPS, sim-verified) [dungeon] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 87.8 ranged_attack_power points (6.73 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Legionnaire's Chain Girdle (16529, -1.02 DPS) [pvp]; Stalker's Mail Belt (252588, -1.56 DPS, sim-verified) [crafted] |
| legs | Beaststalker's Pants (16678) | Stratholme: Baron Rivendare [dungeon] | 115.8 ranged_attack_power points (8.89 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -1.24 DPS, sim-verified) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 87.4 ranged_attack_power points (6.71 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -0.55 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 53.3 ranged_attack_power points (4.09 DPS) | yes | Glowing Crystal Ring (18402, -1.00 DPS) [quest]; Band of the Hierophant (13096, -1.51 DPS) [world_drop]; Frigid Ring (18679, -1.51 DPS) [world] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 44.3 ranged_attack_power points (3.40 DPS) | yes | Glowing Crystal Ring (18402, -0.31 DPS) [quest]; Band of the Hierophant (13096, -0.82 DPS) [world_drop]; Frigid Ring (18679, -0.82 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (308.6 DPS) | yes | Frozen Heart of the Mountain (249469, -7.46 DPS, sim-verified) [crafted]; Uther's Strength (11302, -8.10 DPS) [world_drop] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (308.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (308.6 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -11.13 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (308.6 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -3.74 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Horizon Choker; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Beaststalker's Pants; feet: Timmy's Galoshes; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1785, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 674.4. Weights run: 3.1s. Verify run: 6.1s. 1785 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.407 ± 0.037, crit=2.621 ± 0.079 per rating point (14 rating = 1%, 36.690 per %), hit=3.947 ± 0.192 per rating point (10 rating = 1%, 39.473 per %), melee_haste=17.711 ± 2.318, intellect=3.437 ± 0.096, spell_power=0.873 ± 0.002

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 238.9 ranged_attack_power points (32.56 DPS) | yes | Warlord's Chain Helmet (16566, -6.15 DPS) [vendor]; Warlord's Chain Helm (231571, -6.15 DPS) [vendor]; Outlaw's Collar (279253, -13.52 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 48.1 ranged_attack_power points (6.56 DPS) | yes | Archlight Talisman (15856, -0.45 DPS) [quest]; Skibi's Pendant (13089, -0.52 DPS) [world_drop]; Sunprism Pendant (20649, -0.99 DPS) [quest] |
| shoulder | Shroud of the Nathrezim (18720) | Stratholme: Balnazzar [dungeon] | 97.8 ranged_attack_power points (13.33 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 71.1 ranged_attack_power points (9.69 DPS) | yes | Dark Phantom Cape (13122, -3.52 DPS, sim-verified) [world_drop]; Deep Woodlands Cloak (19121, -4.05 DPS) [quest]; Cerise Drape (15804, -4.09 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 202.2 ranged_attack_power points (27.57 DPS) | yes | Warlord's Chain Chestpiece (16565, -1.15 DPS) [vendor]; Warlord's Chain Hauberk (231573, -1.15 DPS) [vendor]; Spitfire Breastplate (20479, -14.58 DPS, sim-verified) [crafted] |
| wrist | Spitfire Bracers (20481) | Leatherworking [crafted] | 68.6 ranged_attack_power points (9.35 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.45 DPS) [pvp]; Magiskull Cuffs (13107, -2.48 DPS, sim-verified) [world_drop] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | 119.7 ranged_attack_power points (16.31 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Spitfire Gauntlets (20480, -4.67 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 92.4 ranged_attack_power points (12.59 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Stalker's Mail Belt (252588, -1.40 DPS) [crafted]; Legionnaire's Chain Girdle (16529, -1.87 DPS) [pvp] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 131.7 ranged_attack_power points (17.96 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, +0.00 DPS) [pvp]; Red Dragonscale Leggings (252603, -2.62 DPS, sim-verified) [crafted] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 92.3 ranged_attack_power points (12.58 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -0.95 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 59.5 ranged_attack_power points (8.11 DPS) | yes | Glowing Crystal Ring (18402, -2.48 DPS) [quest]; Maiden's Circle (13001, -2.69 DPS) [world_drop]; Band of the Penitent (13217, -3.11 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 47.7 ranged_attack_power points (6.50 DPS) | yes | Maiden's Circle (13001, -1.08 DPS) [world_drop]; Band of the Penitent (13217, -1.50 DPS) [quest]; Glowing Crystal Ring (18402, -2.46 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (674.4 DPS) | yes | Uther's Strength (11302, -14.50 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (674.4 DPS) | yes | Molten Heart of the Mountain (249470, -2.99 DPS, sim-verified) [crafted]; Uther's Strength (11302, -4.13 DPS) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (674.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Enchanted Azsharite Felbane Sword (10696, -21.06 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (674.4 DPS) | yes | Eaglehorn Long Bow (13023, +0.00 DPS) [world_drop]; High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Horizon Choker; shoulder: Shroud of the Nathrezim; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Spitfire Bracers; hands: Mourning Handguards; legs: Ironfeather Leggings; feet: Timmy's Galoshes; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1785, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

