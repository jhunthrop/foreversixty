# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 106.5. Weights run: 1.1s. Verify run: 1.9s. 212 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.067 ± 0.021, crit=0.582 ± 0.023 per rating point (14 rating = 1%, 8.149 per %), hit=0.973 ± 0.053 per rating point (10 rating = 1%, 9.727 per %), melee_haste=5.277 ± 1.040

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.5 ranged_attack_power points (1.67 DPS) | yes | Red Winter Hat (21524, -1.75 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.4 ranged_attack_power points (1.25 DPS) | yes | Erudite's Amulet (277204, -0.90 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.3 ranged_attack_power points (1.04 DPS) | yes | Slime-encrusted Pads (6461, -0.78 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.4 ranged_attack_power points (1.25 DPS) | yes | Cape of the Brotherhood (5193, -0.21 DPS) [dungeon]; Sentry Cloak (2059, -0.42 DPS) [world_drop]; Bristlebark Cape (14571, -0.42 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 33.7 ranged_attack_power points (2.30 DPS) | yes | Brawler's Leather Armor (252490, -0.62 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.84 DPS) [crafted]; Dark Leather Tunic (2317, -1.04 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.3 ranged_attack_power points (1.04 DPS) | yes | Wolf Bracers (4794, -0.21 DPS) [vendor]; Bravo's Armbands (270015, -0.21 DPS) [quest]; Ratchet Wristwraps (274742, -0.42 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (106.5 DPS) | yes | Forest Leather Gloves (3058, -0.42 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.42 DPS) [crafted]; Serpent Gloves (5970, -1.31 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.23 DPS) | yes | Dusty Belt (279897, +0.00 DPS, sim-verified) [quest]; Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.6 ranged_attack_power points (1.88 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.21 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 24.5 ranged_attack_power points (1.67 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS, sim-verified) [dungeon]; Blackened Defias Boots (10402, -0.42 DPS) [dungeon]; Agile Boots (4788, -0.63 DPS) [vendor] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.3 ranged_attack_power points (0.84 DPS) | yes | The 1 Ring (8350, -0.63 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.1 ranged_attack_power points (0.42 DPS) | yes | The 1 Ring (8350, -0.43 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.5 ranged_attack_power points (2.08 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Lupine Axe (1220, -0.83 DPS) [world]; Piercing Axe (6094, -0.83 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.3 ranged_attack_power points (12.22 DPS) | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Lovingly Crafted Boomstick (4372, -2.95 DPS) [crafted]; Venomstrike (6469, -3.08 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 212, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 131.8. Weights run: 1.0s. Verify run: 2.2s. 382 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=3.031 ± 0.024, crit=0.870 ± 0.033 per rating point (14 rating = 1%, 12.175 per %), hit=1.207 ± 0.080 per rating point (10 rating = 1%, 12.067 per %), melee_haste=8.488 ± 1.465

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.3 ranged_attack_power points (2.03 DPS) | yes | Brawler's Leather Hood (252504, -0.46 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.61 DPS) [world]; Resilient Cap (14401, -1.63 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Kaleidoscope Chain (13084, -0.81 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.82 DPS) [world]; Erudite's Amulet (277204, -1.22 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 ranged_attack_power points (2.24 DPS) | yes | Mantle of Thieves (2264, +0.00 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.81 DPS) [crafted]; Insignia Mantle (4721, -0.81 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.41 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.41 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Tunic of Westfall (2041, -0.68 DPS, sim-verified) [quest]; Garb of Florid Feathers (273026, -1.02 DPS) [dungeon]; Green Leather Armor (4255, -1.22 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.20 DPS) [world]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.3 ranged_attack_power points (1.83 DPS) | yes | Pilferer's Gloves (7358, -0.48 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.61 DPS) [dungeon]; Insignia Gloves (6408, -0.61 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (131.8 DPS) | yes | Highlander's Chain Girdle (20090, -0.22 DPS) [rep]; Highlander's Leather Girdle (20117, -0.22 DPS) [rep]; Skulker's Leather Belt (252520, -1.75 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -1.02 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Feet of the Lynx (1121)) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Footpads of the Fang (10411, -0.41 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Thunderbrow Ring (13097, -0.61 DPS) [world_drop] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Monkey Ring (6748, +0.00 DPS, sim-verified) [quest]; Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Thunderbrow Ring (13097, -0.61 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Jade Serpentblade (3850, +0.00 DPS) [crafted] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.55 DPS) | yes | Ironweaver (13137, -1.29 DPS) [world_drop]; Golemsight Long Gun (273029, -1.69 DPS, sim-verified) [dungeon]; Crystalpine Stinger (13037, -1.81 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 169.1. Weights run: 1.0s. Verify run: 2.5s. 602 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.950 ± 0.024, crit=1.046 ± 0.040 per rating point (14 rating = 1%, 14.650 per %), hit=1.307 ± 0.097 per rating point (10 rating = 1%, 13.071 per %), melee_haste=8.912 ± 1.463

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 38.3 ranged_attack_power points (2.73 DPS) | yes | Guard's Chain Helm (250499, -0.21 DPS) [crafted]; Skullsplitter Helm (1624, -0.42 DPS) [world]; Nightscape Headband (8176, -1.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 32.4 ranged_attack_power points (2.31 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.46 DPS) [world]; Kaleidoscope Chain (13084, -1.47 DPS) [world_drop]; Erudite's Amulet (277204, -1.89 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 44.4 ranged_attack_power points (3.16 DPS) | yes | Forest Tracker Epaulets (2278, -0.85 DPS) [world_drop]; Nightscape Shoulders (8192, -0.92 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -1.06 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 23.6 ranged_attack_power points (1.68 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.21 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 50.1 ranged_attack_power points (3.57 DPS) | yes | Nightscape Tunic (8175, -0.42 DPS) [crafted]; Dusky Leather Armor (7374, -0.63 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.63 DPS, sim-verified) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 23.6 ranged_attack_power points (1.68 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Tough Scorpid Bracers (8205, -0.21 DPS) [crafted]; Branded Leather Bracers (19508, -0.26 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 34.7 ranged_attack_power points (2.47 DPS) | yes | Dragonscale Gauntlets (8347, -0.16 DPS) [crafted]; Gauntlets of Divinity (7724, -0.19 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.37 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 ranged_attack_power points (2.14 DPS) | yes | Scorpashi Sash (14652, -0.04 DPS) [world_drop]; Blackforge Girdle (6425, -0.25 DPS) [dungeon]; Ogron's Sash (13117, -0.25 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 61.9 ranged_attack_power points (4.41 DPS) | yes | Petrolspill Leggings (9509, -1.47 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.47 DPS) [world_drop]; Triprunner Dungarees (9624, -1.71 DPS, sim-verified) [quest] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 38.3 ranged_attack_power points (2.73 DPS) | yes | Imperial Leather Boots (6431, -0.42 DPS) [dungeon]; Dusky Boots (7390, -0.42 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.42 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 29.5 ranged_attack_power points (2.10 DPS) | yes | Protector's Band (19515, -0.42 DPS) [rep]; Disengagement Ring (276202, -0.42 DPS) [vendor]; Field Researcher's Loop (281634, -0.63 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 26.5 ranged_attack_power points (1.89 DPS) | yes | Protector's Band (19515, -0.21 DPS) [rep]; Disengagement Ring (276202, -0.21 DPS) [vendor]; Field Researcher's Loop (281634, -0.42 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (169.1 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -3.07 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 23.6 ranged_attack_power points (1.68 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (24.91 DPS) | yes | Swiftwind (13038, -2.27 DPS) [world_drop]; Baelog's Shortbow (9400, -2.38 DPS) [dungeon]; The Silencer (13138, -3.96 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 602, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 50 (dwarf, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 220.1. Weights run: 1.0s. Verify run: 3.5s. 758 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.918 ± 0.025, crit=1.219 ± 0.044 per rating point (14 rating = 1%, 17.067 per %), hit=1.532 ± 0.117 per rating point (10 rating = 1%, 15.321 per %), melee_haste=12.475 ± 1.896

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 63.7 ranged_attack_power points (4.91 DPS) | yes | Lordrec Helmet (10741, -1.31 DPS) [quest]; Sprightring Helm (17776, -1.54 DPS) [quest]; Helm of Fire (8348, -2.42 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.9 ranged_attack_power points (2.92 DPS) | yes | Sentinel's Medallion (19539, -0.22 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 55.0 ranged_attack_power points (4.24 DPS) | yes | Sunburn Spaulders (274751, -1.13 DPS, sim-verified) [vendor]; Khan's Mantle (14787, -1.54 DPS) [world_drop]; Nightscape Shoulders (8192, -1.76 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 43.8 ranged_attack_power points (3.37 DPS) | yes | Serpentskin Cloak (8259, -0.90 DPS) [world_drop]; Nightscape Cloak (8195, -1.12 DPS) [crafted]; Blackflame Cape (13109, -1.52 DPS, sim-verified) [world_drop] |
| chest | Blazewind Breastplate (11193) | Tremors of the Earth [quest] | 67.1 ranged_attack_power points (5.17 DPS) | yes | Knight's Chain Armor (220828, +0.00 DPS, sim-verified) [vendor]; Warbear Harness (15064, -1.12 DPS) [crafted]; Charred Leather Tunic (19127, -1.12 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.8 ranged_attack_power points (3.37 DPS) | yes | Bloodlust Bracelets (14807, -0.80 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.90 DPS) [crafted]; Pridelord Bands (14672, -1.12 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 64.2 ranged_attack_power points (4.94 DPS) | yes | Sergeant Major's Chain Gauntlets (220829, -1.58 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.80 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.80 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 56.8 ranged_attack_power points (4.37 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.23 DPS) [crafted]; Stalker's Mail Belt (252588, -1.23 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 61.3 ranged_attack_power points (4.72 DPS) | yes | Knight's Chain Legplates (220832, +0.00 DPS, sim-verified) [vendor]; Keeper's Woolies (14668, -0.45 DPS) [world_drop]; Oilskin Leggings (9414, -0.67 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 58.4 ranged_attack_power points (4.49 DPS) | yes | Fleetfoot Greaves (11627, -0.22 DPS) [dungeon]; Elven Chain Boots (13125, -0.45 DPS) [world_drop]; Sergeant Major's Chain Sabatons (220837, -1.12 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.8 ranged_attack_power points (3.15 DPS) | yes | Falcon's Hook (7552, -1.12 DPS) [dungeon]; Protector's Band (19516, -1.12 DPS) [rep]; Disengagement Ring (276202, -1.35 DPS) [vendor] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.2 ranged_attack_power points (2.25 DPS) | yes | Falcon's Hook (7552, -0.22 DPS) [dungeon]; Protector's Band (19516, -0.22 DPS) [rep]; Disengagement Ring (276202, -0.45 DPS) [vendor] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (220.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (220.1 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-verified (220.1 DPS) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Hanzo Sword (8190, -2.37 DPS, sim-verified) [world_drop] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 26.3 ranged_attack_power points (2.02 DPS) | yes | Widowmaker (4091, +0.00 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (220.1 DPS) | yes | Gryphonwing Long Bow (13022, -0.64 DPS) [dungeon]; Dark Iron Rifle (16004, -1.84 DPS, sim-verified) [crafted]; Houndmaster's Bow (11628, -2.71 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 758, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60 (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 297.7. Weights run: 1.1s. Verify run: 10.6s. 1388 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.964 ± 0.029, crit=1.678 ± 0.060 per rating point (14 rating = 1%, 23.495 per %), hit=2.228 ± 0.156 per rating point (10 rating = 1%, 22.278 per %), melee_haste=15.420 ± 2.208

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (297.7 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Beaststalker's Cap (16677, -11.46 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.3 ranged_attack_power points (4.30 DPS) | yes | Medallion of the Dawn (22659, -0.67 DPS) [quest]; Beads of Ogre Might (22150, -0.77 DPS) [quest]; Sentinel's Medallion (19538, -0.90 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 103.5 ranged_attack_power points (7.90 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Field Marshal's Chain Spaulders (16468, -0.32 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -0.32 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 64.5 ranged_attack_power points (4.92 DPS) | yes | Howler's Furs (272414, -1.08 DPS) [vendor]; Cloak of the Honor Guard (20073, -1.19 DPS) [rep]; Shifting Cloak (18511, -1.40 DPS, sim-verified) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.4 ranged_attack_power points (11.86 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.38 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.38 DPS) [vendor]; Obsidian Mail Tunic (22191, -6.82 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (297.7 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -11.19 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (297.7 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (297.7 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -9.43 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 173.0 ranged_attack_power points (13.20 DPS) | yes | Sentinel's Leather Pants (237818, -3.51 DPS) [vendor]; Marshal's Chain Legguards (231577, -3.72 DPS) [pvp]; Marshal's Chain Legplates (231558, -4.82 DPS) [vendor] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (297.7 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -10.94 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 93.4 ranged_attack_power points (7.13 DPS) | yes | Don Julio's Band (19325, -2.41 DPS) [rep]; Masons Fraternity Ring (9533, -3.96 DPS) [quest]; Protector's Band (19514, -4.64 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 63.8 ranged_attack_power points (4.87 DPS) | yes | Don Julio's Band (19325, -0.15 DPS) [rep]; Masons Fraternity Ring (9533, -1.70 DPS) [quest]; Protector's Band (19514, -2.38 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (297.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (297.7 DPS) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (297.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -10.59 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (297.7 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -13.98 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 755.3. Weights run: 1.0s. Verify run: 8.5s. 1388 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=3.024 ± 0.028, crit=1.790 ± 0.056 per rating point (14 rating = 1%, 25.056 per %), hit=3.346 ± 0.238 per rating point (10 rating = 1%, 33.456 per %), melee_haste=14.961 ± 3.085

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 103.0 ranged_attack_power points (17.31 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -4.49 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 57.5 ranged_attack_power points (9.65 DPS) | yes | Beads of Ogre Might (22150, -0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.41 DPS) [quest]; Sentinel's Medallion (19538, -2.03 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 106.7 ranged_attack_power points (17.93 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 65.4 ranged_attack_power points (10.98 DPS) | yes | Howler's Furs (272414, -0.66 DPS) [vendor]; Shifting Cloak (18511, -2.34 DPS) [crafted]; Cloak of the Honor Guard (20073, -2.73 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 170.4 ranged_attack_power points (28.63 DPS) | yes | Field Marshal's Chain Breastplate (16466, -7.14 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -7.14 DPS) [vendor]; Shadowcraft Tunic (16721, -11.97 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 57.5 ranged_attack_power points (9.65 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.02 DPS) [pvp]; Beastmaster's Bindings (226885, -1.52 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 81.7 ranged_attack_power points (13.72 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -0.73 DPS) [pvp]; Mourning Handguards (279257, -3.05 DPS) [crafted] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (755.3 DPS) | yes | Dense Timbermaw Belt (227807, -1.44 DPS) [vendor]; Highlander's Chain Girdle (20043, -2.27 DPS) [rep]; Belt of Preserved Heads (20216, -7.82 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 189.4 ranged_attack_power points (31.82 DPS) | yes | Sentinel's Leather Pants (237818, -9.68 DPS) [vendor]; Marshal's Chain Legguards (231577, -10.34 DPS) [pvp]; Marshal's Chain Legplates (231558, -12.93 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 73.6 ranged_attack_power points (12.36 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Beastmaster's Boots (22061, +0.00 DPS) [quest]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 106.0 ranged_attack_power points (17.81 DPS) | yes | Don Julio's Band (19325, -5.30 DPS) [rep]; Masons Fraternity Ring (9533, -10.70 DPS) [quest]; Channeler's Ring (272406, -12.19 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 75.8 ranged_attack_power points (12.73 DPS) | yes | Don Julio's Band (19325, -0.22 DPS) [rep]; Masons Fraternity Ring (9533, -5.62 DPS) [quest]; Channeler's Ring (272406, -7.11 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+4.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -19.78 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -32.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4833 Glorious Shoulders

## Horde

### Band 20 (troll, 0000000000000000-0051500000000000-000000000000000000)

Set DPS (verified): 107.9. Weights run: 1.1s. Verify run: 1.9s. 201 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.067 ± 0.021, crit=0.582 ± 0.023 per rating point (14 rating = 1%, 8.149 per %), hit=0.973 ± 0.053 per rating point (10 rating = 1%, 9.727 per %), melee_haste=5.277 ± 1.040

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.5 ranged_attack_power points (1.67 DPS) | yes | Red Winter Hat (21524, -1.75 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.4 ranged_attack_power points (1.25 DPS) | yes | Erudite's Amulet (277204, -0.90 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.3 ranged_attack_power points (1.04 DPS) | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 18.4 ranged_attack_power points (1.25 DPS) | yes | Cape of the Brotherhood (5193, -0.41 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.42 DPS) [world_drop]; Bristlebark Cape (14571, -0.42 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 21.5 ranged_attack_power points (1.46 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.21 DPS) [crafted]; Prospector's Chestpiece (14562, -0.21 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.3 ranged_attack_power points (1.04 DPS) | yes | Wolf Bracers (4794, -0.21 DPS) [vendor]; Bristlebark Bindings (14569, -0.42 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.42 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Forest Leather Gloves (3058, -0.42 DPS) [world_drop]; Nimble Leather Gloves (7285, -0.42 DPS) [crafted]; Serpent Gloves (5970, -1.57 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.23 DPS) | yes | Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Guardsman Belt (3429, -0.39 DPS) [world]; Dark Leather Belt (4249, -0.39 DPS) [crafted] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 27.6 ranged_attack_power points (1.88 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.21 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Agile Boots (4788, -0.21 DPS) [vendor]; Feet of the Lynx (1121, -1.19 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.3 ranged_attack_power points (0.84 DPS) | yes | The 1 Ring (8350, -0.63 DPS) [world]; Bounty Hunter's Ring (5351, -0.90 DPS, sim-verified) [quest] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Bounty Hunter's Ring (5351, -1.08 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 30.5 ranged_attack_power points (2.08 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.62 DPS) [quest]; Lupine Axe (1220, -0.83 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.3 ranged_attack_power points (12.22 DPS) | yes | Lil Timmy's Peashooter (13136, -0.56 DPS, sim-verified) [world_drop]; Outrider's Bow (20437, -0.67 DPS) [pvp]; Lovingly Crafted Boomstick (4372, -2.95 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Legionnaire's Band; finger2: Lavishly Jeweled Ring; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 201, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 209611 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff

### Band 30 (troll, 0000000000000000-0051550001400000-000000000000000000)

Set DPS (verified): 133.8. Weights run: 1.0s. Verify run: 2.1s. 365 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=3.031 ± 0.024, crit=0.870 ± 0.033 per rating point (14 rating = 1%, 12.175 per %), hit=1.207 ± 0.080 per rating point (10 rating = 1%, 12.067 per %), melee_haste=8.488 ± 1.465

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.3 ranged_attack_power points (2.03 DPS) | yes | Brawler's Leather Hood (252504, -0.48 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.61 DPS) [world]; Resilient Cap (14401, -1.63 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Kaleidoscope Chain (13084, -0.81 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.82 DPS) [world]; Erudite's Amulet (277204, -1.22 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 ranged_attack_power points (2.24 DPS) | yes | Mantle of Thieves (2264, -0.20 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.81 DPS) [crafted]; Insignia Mantle (4721, -0.81 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Hawkeye's Cloak (14593, -0.20 DPS) [world_drop]; Fenrus' Hide (6340, -0.41 DPS) [dungeon]; Swiftrunner Cape (6745, -0.41 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Garb of Florid Feathers (273026, +0.00 DPS, sim-verified) [dungeon]; Panther Armor (6670, -1.02 DPS) [quest]; Green Leather Armor (4255, -1.22 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.20 DPS) [world_drop]; Loamflake Bracers (15462, -0.20 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.3 ranged_attack_power points (1.83 DPS) | yes | Braced Handguards (6784, -0.41 DPS) [quest]; Insignia Gloves (6408, -0.61 DPS) [world_drop]; Pilferer's Gloves (7358, -0.80 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (133.8 DPS) | yes | Deftkin Belt (16659, -0.21 DPS) [quest]; Defiler's Chain Girdle (20152, -0.22 DPS) [rep]; Skulker's Leather Belt (252520, -1.40 DPS, sim-verified) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan [dungeon] | 42.4 ranged_attack_power points (2.85 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.20 DPS) [crafted]; Insignia Leggings (4054, -1.02 DPS) [world_drop] |
| feet | Insignia Boots (4055) (or Highlander's Mail Greaves (20123), Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 24.2 ranged_attack_power points (1.63 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor] |
| finger1 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Band of the Fist (17694, -0.41 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Defias Renegade Ring (1076, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.20 DPS) [quest]; Band of the Fist (17694, -0.41 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Vendetta (776) (or Sentinel's Blade (212583)) | Razorfen Kraul [dungeon] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Sentinel's Blade (212583) | Illiyana Moonblaze [vendor] | 18.2 ranged_attack_power points (1.22 DPS) | yes | Jade Serpentblade (3850, +0.00 DPS) [crafted] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (17.55 DPS) | yes | Ironweaver (13137, -1.29 DPS) [world_drop]; Outrider's Bow (19561, -1.48 DPS) [pvp]; Golemsight Long Gun (273029, -1.96 DPS, sim-verified) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Petrolspill Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Legionnaire's Band; main_hand: Vendetta; off_hand: Sentinel's Blade; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 365, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 1000000000000000-0051550001503050-000000000000000000)

Set DPS (verified): 171.0. Weights run: 1.0s. Verify run: 2.5s. 566 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.950 ± 0.024, crit=1.046 ± 0.040 per rating point (14 rating = 1%, 14.650 per %), hit=1.307 ± 0.097 per rating point (10 rating = 1%, 13.071 per %), melee_haste=8.912 ± 1.463

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 38.3 ranged_attack_power points (2.73 DPS) | yes | Nightscape Headband (8176, -0.21 DPS) [crafted]; Guard's Chain Helm (250499, -0.21 DPS) [crafted]; Skullsplitter Helm (1624, -0.42 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 32.4 ranged_attack_power points (2.31 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.46 DPS) [world]; Ethereal Talisman (4430, -1.47 DPS) [quest]; Kaleidoscope Chain (13084, -1.47 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 44.4 ranged_attack_power points (3.16 DPS) | yes | Forest Tracker Epaulets (2278, -0.85 DPS) [world_drop]; Nightscape Shoulders (8192, -0.94 DPS, sim-verified) [crafted]; Flintrock Shoulders (7755, -1.06 DPS) [dungeon] |
| back | Parachute Cloak (10518) (or Imperial Cloak (6432), Tigerstrike Mantle (13108)) | Engineering [crafted] | 23.6 ranged_attack_power points (1.68 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.21 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 50.1 ranged_attack_power points (3.57 DPS) | yes | Nightscape Tunic (8175, -0.42 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.42 DPS) [crafted]; Dusky Leather Armor (7374, -0.63 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 23.6 ranged_attack_power points (1.68 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Windtalker's Wristguards (19584, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 34.7 ranged_attack_power points (2.47 DPS) | yes | Dragonscale Gauntlets (8347, -0.16 DPS) [crafted]; Gauntlets of Divinity (7724, -0.19 DPS) [dungeon]; Tough Scorpid Gloves (8204, -0.37 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 ranged_attack_power points (2.14 DPS) | yes | Scorpashi Sash (14652, -0.04 DPS) [world_drop]; Blackforge Girdle (6425, -0.25 DPS) [dungeon]; Ogron's Sash (13117, -0.25 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 61.9 ranged_attack_power points (4.41 DPS) | yes | Triprunner Dungarees (9624, -0.80 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.47 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.47 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 38.3 ranged_attack_power points (2.73 DPS) | yes | Imperial Leather Boots (6431, -0.42 DPS) [dungeon]; Dusky Boots (7390, -0.42 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.42 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 29.5 ranged_attack_power points (2.10 DPS) | yes | Legionnaire's Band (19512, -0.42 DPS) [rep]; Disengagement Ring (276202, -0.42 DPS) [vendor]; Field Researcher's Loop (281634, -0.63 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 26.5 ranged_attack_power points (1.89 DPS) | yes | Legionnaire's Band (19512, -0.21 DPS) [rep]; Disengagement Ring (276202, -0.21 DPS) [vendor]; Field Researcher's Loop (281634, -0.42 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (171.0 DPS) | yes | Hellslayer Battle Axe (13017, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted]; Gut Ripper (2164, -2.38 DPS, sim-verified) [world_drop] |
| off_hand | Blue Glittering Axe (7942) (or Nordic Longshank (9401), Speedsteel Rapier (13034)) | Blacksmithing [crafted] | 23.6 ranged_attack_power points (1.68 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Speedsteel Rapier (13034, +0.00 DPS) [world_drop] |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | 350.0 ranged_attack_power points (24.91 DPS) | yes | Outrider's Bow (19560, -1.65 DPS) [pvp]; Swiftwind (13038, -2.27 DPS) [world_drop]; The Silencer (13138, -3.53 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Parachute Cloak; chest: Wolffear Harness; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Blue Glittering Axe; ranged: Bow of Searing Arrows

No-known-source sample (15 of 566, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (troll, 5320000100000000-0051550001503050-000000000000000000)

Set DPS (verified): 226.7. Weights run: 1.0s. Verify run: 3.1s. 715 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.918 ± 0.025, crit=1.219 ± 0.044 per rating point (14 rating = 1%, 17.067 per %), hit=1.532 ± 0.117 per rating point (10 rating = 1%, 15.321 per %), melee_haste=12.475 ± 1.896

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Fire (8348) | Leatherworking [crafted] | 49.6 ranged_attack_power points (3.82 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Sprightring Helm (17776, -0.45 DPS) [quest]; Tough Scorpid Helm (8208, -0.67 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.9 ranged_attack_power points (2.92 DPS) | yes | Scout's Medallion (19535, -0.22 DPS) [rep]; Woven Ivy Necklace (19159, -0.90 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 44.1 ranged_attack_power points (3.40 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Khan's Mantle (14787, -0.70 DPS) [world_drop]; Nightscape Shoulders (8192, -0.92 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 43.8 ranged_attack_power points (3.37 DPS) | yes | Serpentskin Cloak (8259, -0.90 DPS) [world_drop]; Nightscape Cloak (8195, -1.12 DPS) [crafted]; Blackflame Cape (13109, -1.13 DPS, sim-verified) [world_drop] |
| chest | Blazewind Breastplate (11193) | Broken Alliances [quest] | 67.1 ranged_attack_power points (5.17 DPS) | yes | Stone Guard's Chain Armor (220827, +0.00 DPS, sim-verified) [vendor]; Warbear Harness (15064, -1.12 DPS) [crafted]; Charred Leather Tunic (19127, -1.12 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.8 ranged_attack_power points (3.37 DPS) | yes | Bloodlust Bracelets (14807, +0.00 DPS) [world_drop]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 64.2 ranged_attack_power points (4.94 DPS) | yes | First Sergeant's Chain Gauntlets (220830, -1.62 DPS, sim-verified) [vendor]; Beastmaster's Gauntlets (226883, -1.80 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.80 DPS) [crafted] |
| waist | Substandard Belt Chain (274757) | Zippie Fizzbolt [vendor] | 56.8 ranged_attack_power points (4.37 DPS) | yes | Sagebrush Girdle (17778, +0.00 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -1.23 DPS) [crafted]; Stalker's Mail Belt (252588, -1.23 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 61.3 ranged_attack_power points (4.72 DPS) | yes | Stone Guard's Chain Legplates (220833, +0.00 DPS, sim-verified) [vendor]; Keeper's Woolies (14668, -0.45 DPS) [world_drop]; Oilskin Leggings (9414, -0.67 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 58.4 ranged_attack_power points (4.49 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.45 DPS) [world_drop]; First Sergeant's Chain Sabatons (220836, -1.12 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 40.8 ranged_attack_power points (3.15 DPS) | yes | Falcon's Hook (7552, -1.12 DPS) [dungeon]; Legionnaire's Band (19511, -1.12 DPS) [rep]; White Bone Band (11862, -1.30 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 29.2 ranged_attack_power points (2.25 DPS) | yes | Falcon's Hook (7552, -0.22 DPS) [dungeon]; Legionnaire's Band (19511, -0.22 DPS) [rep]; White Bone Band (11862, -0.40 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -7.80 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hanzo Sword (8190, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Steel Spear (250605, +0.00 DPS) [crafted] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (226.7 DPS) | yes | White Bone Shredder (11863, -0.45 DPS) [quest]; Force of the Hippogryph (9684, -2.52 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.64 DPS) [dungeon]; Dark Iron Rifle (16004, -2.14 DPS, sim-verified) [crafted]; Outrider's Bow (19559, -2.65 DPS) [pvp] |

**New at 50:** head: Helm of Fire; neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Blazewind Breastplate; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Substandard Belt Chain; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ebon Shiv; off_hand: Jhordy's Misplaced Screwdriver; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 715, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 306.9. Weights run: 1.1s. Verify run: 10.5s. 1368 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=2.964 ± 0.029, crit=1.678 ± 0.060 per rating point (14 rating = 1%, 23.495 per %), hit=2.228 ± 0.156 per rating point (10 rating = 1%, 22.278 per %), melee_haste=15.420 ± 2.208

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | sim-verified (306.9 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -9.93 DPS, sim-verified) [dungeon] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.3 ranged_attack_power points (4.30 DPS) | yes | Medallion of the Dawn (22659, -0.67 DPS) [quest]; Beads of Ogre Might (22150, -0.77 DPS) [quest]; Scout's Medallion (19534, -0.90 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 103.5 ranged_attack_power points (7.90 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -0.32 DPS) [pvp]; Warlord's Chain Pauldrons (231565, -1.12 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 64.5 ranged_attack_power points (4.92 DPS) | yes | Howler's Furs (272414, -1.08 DPS) [vendor]; Deathguard's Cloak (20068, -1.19 DPS) [rep]; Shifting Cloak (18511, -1.31 DPS, sim-verified) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 155.4 ranged_attack_power points (11.86 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.38 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.38 DPS) [vendor]; Obsidian Mail Tunic (22191, -6.30 DPS, sim-verified) [crafted] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (306.9 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Beaststalker's Bindings (16681, -9.79 DPS, sim-verified) [dungeon] |
| hands | Beaststalker's Gloves (16676) | Blackrock Spire: War Master Voone [dungeon] | sim-verified (306.9 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS, sim-verified) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (306.9 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Beaststalker's Belt (16680, -8.15 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 173.0 ranged_attack_power points (13.20 DPS) | yes | Outrider's Chain Leggings (22673, -2.72 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -3.51 DPS) [vendor]; General's Chain Legguards (231574, -3.72 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (306.9 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -10.38 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 93.4 ranged_attack_power points (7.13 DPS) | yes | Don Julio's Band (19325, -2.41 DPS) [rep]; Masons Fraternity Ring (9533, -3.96 DPS) [quest]; Legionnaire's Band (19510, -4.64 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 63.8 ranged_attack_power points (4.87 DPS) | yes | Don Julio's Band (19325, -0.15 DPS) [rep]; Masons Fraternity Ring (9533, -1.70 DPS) [quest]; Legionnaire's Band (19510, -2.38 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (306.9 DPS) | yes | Frozen Heart of the Mountain (249469, -7.35 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (306.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (306.9 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -9.77 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (306.9 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -10.57 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Beaststalker's Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; trinket2: Molten Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 5320000501000000-0051550001503050-500000000000000000)

Set DPS (verified): 785.6. Weights run: 1.0s. Verify run: 10.4s. 1368 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.002, agility=3.024 ± 0.028, crit=1.790 ± 0.056 per rating point (14 rating = 1%, 25.056 per %), hit=3.346 ± 0.238 per rating point (10 rating = 1%, 33.456 per %), melee_haste=14.961 ± 3.085

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Beastmaster's Cap (226887) | Saving the Best for Last [quest] | 103.0 ranged_attack_power points (17.31 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -10.13 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 57.5 ranged_attack_power points (9.65 DPS) | yes | Beads of Ogre Might (22150, -0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.41 DPS) [quest]; Scout's Medallion (19534, -2.03 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 106.7 ranged_attack_power points (17.93 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 65.4 ranged_attack_power points (10.98 DPS) | yes | Howler's Furs (272414, -0.66 DPS) [vendor]; Shifting Cloak (18511, -2.34 DPS) [crafted]; Deathguard's Cloak (20068, -2.73 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 170.4 ranged_attack_power points (28.63 DPS) | yes | Warlord's Chain Chestpiece (16565, -7.14 DPS) [vendor]; Warlord's Chain Hauberk (231573, -7.14 DPS) [vendor]; Shadowcraft Tunic (16721, -13.68 DPS, sim-verified) [dungeon] |
| wrist | Beastmaster's Bindings (226885) | An Earnest Proposition [quest] | sim-verified (785.6 DPS) | yes | Legionnaire's Chain Bracers (16517, +0.00 DPS) [pvp]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19587, -7.97 DPS, sim-verified) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 81.7 ranged_attack_power points (13.72 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; General's Chain Grips (231569, -0.73 DPS) [vendor] |
| waist | Beastmaster's Belt (226888) | Just Compensation [quest] | sim-verified (785.6 DPS) | yes | Dense Timbermaw Belt (227807, +0.00 DPS) [vendor]; Ranger's Belt (272397, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -5.14 DPS, sim-verified) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 189.4 ranged_attack_power points (31.82 DPS) | yes | Outrider's Chain Leggings (22673, -5.70 DPS, sim-verified) [rep]; Sentinel's Leather Pants (237818, -9.68 DPS) [vendor]; General's Chain Legguards (231574, -10.34 DPS) [pvp] |
| feet | Beastmaster's Treads (226881) | Mokvar [vendor] | sim-verified (785.6 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -9.37 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 106.0 ranged_attack_power points (17.81 DPS) | yes | Don Julio's Band (19325, -5.30 DPS) [rep]; Masons Fraternity Ring (9533, -10.70 DPS) [quest]; Channeler's Ring (272406, -12.19 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | 75.8 ranged_attack_power points (12.73 DPS) | yes | Don Julio's Band (19325, -0.22 DPS) [rep]; Masons Fraternity Ring (9533, -5.62 DPS) [quest]; Channeler's Ring (272406, -7.11 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (785.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (785.6 DPS) | yes | Molten Heart of the Mountain (249470, -5.38 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (785.6 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -21.83 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (785.6 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -26.96 DPS, sim-verified) [crafted] |

**New at 60:** head: Beastmaster's Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Beastmaster's Bindings; hands: Raider Gloves; waist: Beastmaster's Belt; legs: Sentinel's Chain Leggings; feet: Beastmaster's Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Cutthroat's Signet; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

