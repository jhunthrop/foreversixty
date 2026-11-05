# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 67.3. Weights run: 7.5s. Verify run: 8.2s. 258 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.173 ± 0.012, crit=0.679 ± 0.013 per rating point (14 rating = 1%, 9.505 per %), hit=1.150 ± 0.041 per rating point (10 rating = 1%, 11.498 per %), melee_haste=5.252 ± 0.566, intellect=1.086 ± 0.030, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.4 ranged_attack_power points (1.82 DPS) | yes | Trapper's Leather Hood (252505, -1.15 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.43 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 19.0 ranged_attack_power points (1.36 DPS) | yes | Erudite's Amulet (277204, -0.90 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.9 ranged_attack_power points (1.14 DPS) | yes | Magician's Mantle (12998, -0.59 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.82 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.23 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.45 DPS) [world_drop]; Bristlebark Cape (14571, -0.45 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.9 ranged_attack_power points (2.50 DPS) | yes | Trapper's Leather Armor (252491, -0.66 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.91 DPS) [crafted]; Dark Leather Tunic (2317, -1.14 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.9 ranged_attack_power points (1.14 DPS) | yes | Wolf Bracers (4794, -0.23 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.45 DPS) [world_drop]; Bravo's Armbands (270015, -0.45 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.0 ranged_attack_power points (1.36 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.22 DPS) [crafted]; Riding Gloves (1304, -0.45 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.29 DPS) | yes | Trapper's Leather Belt (252429, -0.15 DPS) [crafted]; Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.23 DPS) [world]; Trapper's Leather Pants (252501, -0.36 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.4 ranged_attack_power points (1.82 DPS) | yes | Blackened Defias Boots (10402, -0.45 DPS) [dungeon]; Footpads of the Fang (10411, -0.45 DPS) [dungeon]; Trapper's Leather Boots (252440, -0.48 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 12.9 ranged_attack_power points (0.92 DPS) | yes | The 1 Ring (8350, -0.62 DPS) [world]; Volcanic Rock Ring (12053, -0.69 DPS) [world_drop]; Minor Channeling Ring (1449, -0.77 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.7 ranged_attack_power points (0.91 DPS) | yes | The 1 Ring (8350, -0.56 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.68 DPS) [world_drop]; Minor Channeling Ring (1449, -0.75 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 34.0 ranged_attack_power points (2.43 DPS) | yes | Bronze Dory (250603, -0.24 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.07 DPS) [world]; Piercing Axe (6094, -1.07 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (12.85 DPS) | yes | Venomstrike (6469, -0.51 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -0.95 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.10 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 258, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (dwarf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 104.4. Weights run: 7.9s. Verify run: 19.7s. 447 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.114 ± 0.013, crit=0.981 ± 0.018 per rating point (14 rating = 1%, 13.733 per %), hit=1.617 ± 0.052 per rating point (10 rating = 1%, 16.165 per %), melee_haste=8.587 ± 0.754, intellect=1.204 ± 0.038, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 31.1 ranged_attack_power points (2.32 DPS) | yes | Brawler's Leather Hood (252504, -0.47 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.70 DPS) [world]; Resilient Cap (14401, -0.87 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 24.9 ranged_attack_power points (1.85 DPS) | yes | Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.69 DPS, sim-verified) [world]; Erudite's Amulet (277204, -1.39 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.3 ranged_attack_power points (2.55 DPS) | yes | Mantle of Thieves (2264, -0.23 DPS) [dungeon]; Dark Leather Shoulders (4252, -0.93 DPS) [crafted]; Insignia Mantle (4721, -0.93 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.46 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.57 DPS) [pvp] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 46.1 ranged_attack_power points (3.43 DPS) | yes | Dusky Leather Armor (7374, -0.19 DPS) [crafted]; Green Leather Armor (4255, -0.86 DPS) [crafted]; Tunic of Westfall (2041, -0.88 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.7 ranged_attack_power points (1.39 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.23 DPS) [world]; Forest Leather Bracers (3202, -0.23 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 28.0 ranged_attack_power points (2.09 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.44 DPS) [dungeon]; Pilferer's Gloves (7358, -0.68 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -0.70 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 35.2 ranged_attack_power points (2.62 DPS) | yes | Skulker's Leather Belt (252520, -0.57 DPS, sim-verified) [crafted]; Green Leather Belt (4257, -0.70 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.84 DPS) [rep] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dusky Leather Leggings (7373, -0.23 DPS) [crafted]; Ferine Leggings (6690, -0.23 DPS) [dungeon]; Petrolspill Leggings (9509, -0.31 DPS, sim-verified) [dungeon] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 24.9 ranged_attack_power points (1.85 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Trapper's Leather Boots (252440, -0.34 DPS) [crafted]; Footpads of the Fang (10411, -0.46 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 18.7 ranged_attack_power points (1.39 DPS) | yes | Monkey Ring (6748, -0.23 DPS) [quest]; Seal of Wrynn (2933, -0.34 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.7 ranged_attack_power points (1.39 DPS) | yes | Monkey Ring (6748, -0.23 DPS) [quest]; Seal of Wrynn (2933, -0.34 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kam's Walking Stick (2280, -0.93 DPS) [dungeon]; Armor Piercer (6679, -0.97 DPS) [dungeon]; Thornspike (6681, -2.00 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 251.1 ranged_attack_power points (18.69 DPS) | yes | Ironweaver (13137, -0.31 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -0.84 DPS) [world_drop]; Chesterfall Musket (7729, -1.47 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 447, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (dwarf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 170.4. Weights run: 8.3s. Verify run: 20.5s. 681 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.873 ± 0.033, crit=0.833 ± 0.052 per rating point (14 rating = 1%, 11.661 per %), hit=1.802 ± 0.090 per rating point (10 rating = 1%, 18.021 per %), melee_haste=7.813 ± 0.893, intellect=1.209 ± 0.047, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 41.0 ranged_attack_power points (3.03 DPS) | yes | Nightscape Headband (8176, -0.48 DPS) [crafted]; Guard's Chain Helm (250499, -0.48 DPS) [crafted]; Skullsplitter Helm (1624, -0.69 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 31.6 ranged_attack_power points (2.34 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.91 DPS) [world]; Kaleidoscope Chain (13084, -1.13 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.62 DPS) [dungeon] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 31.6 ranged_attack_power points (2.34 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.21 DPS) [dungeon]; Flintrock Shoulders (7755, -0.21 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (170.4 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Parachute Cloak (10518, -0.70 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 87.0 ranged_attack_power points (6.44 DPS) | yes | Wolffear Harness (13110, -2.82 DPS) [world_drop]; Garb of Florid Feathers (273026, -3.18 DPS) [dungeon]; Tough Scorpid Breastplate (8203, -3.62 DPS, sim-verified) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 27.4 ranged_attack_power points (2.03 DPS) | yes | Green Leather Bracers (4259, -0.30 DPS) [crafted]; Imperial Leather Bracers (4061, -0.32 DPS) [dungeon]; Dusky Bracers (7378, -0.32 DPS) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-verified (170.4 DPS) | yes | Stalker's Leather Gloves (252526, -0.18 DPS) [crafted]; Aurora Gloves (4042, -0.52 DPS) [world_drop]; Gauntlets of Divinity (7724, -0.56 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 33.1 ranged_attack_power points (2.45 DPS) | yes | Scorpashi Sash (14652, -0.32 DPS) [world_drop]; Blackforge Girdle (6425, -0.54 DPS) [dungeon]; Highlander's Leather Girdle (20116, -0.87 DPS, sim-verified) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 60.3 ranged_attack_power points (4.47 DPS) | yes | Triprunner Dungarees (9624, -1.23 DPS, sim-verified) [quest]; Ferine Leggings (6690, -1.47 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.49 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 40.1 ranged_attack_power points (2.97 DPS) | yes | Swampwalker Boots (2276, -0.20 DPS) [world_drop]; Imperial Leather Boots (6431, -0.63 DPS) [dungeon]; Worn Running Boots (9398, -1.28 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.7 ranged_attack_power points (2.13 DPS) | yes | Protector's Band (19515, -0.43 DPS) [rep]; Disengagement Ring (276202, -0.43 DPS) [vendor]; Field Researcher's Loop (281634, -0.64 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.9 ranged_attack_power points (1.91 DPS) | yes | Protector's Band (19515, -0.21 DPS) [rep]; Disengagement Ring (276202, -0.21 DPS) [vendor]; Field Researcher's Loop (281634, -0.43 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Steel Spear (250605, -0.29 DPS) [crafted]; Felbough Arm (273037, -0.33 DPS) [world]; Claw of the Shadowmancer (2912, -3.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Swiftwind (13038, -2.40 DPS) [world_drop]; Baelog's Shortbow (9400, -2.48 DPS) [dungeon]; The Silencer (13138, -4.03 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 681, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 5000 Coral Band

### Band 50 (dwarf, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 229.4. Weights run: 8.8s. Verify run: 29.8s. 854 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.202 ± 0.036, crit=1.989 ± 0.066 per rating point (14 rating = 1%, 27.844 per %), hit=2.396 ± 0.110 per rating point (10 rating = 1%, 23.957 per %), melee_haste=9.365 ± 1.267, intellect=1.883 ± 0.062, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 71.2 ranged_attack_power points (5.19 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.55 DPS) [dungeon]; Helm of Fire (8348, -1.22 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.6 ranged_attack_power points (3.03 DPS) | yes | Sentinel's Medallion (19539, -0.23 DPS) [rep]; Horizon Choker (13085, -1.11 DPS) [world_drop] |
| shoulder | Stalker's Mail Shoulders (252566) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Tough Scorpid Shoulders (8207, -0.34 DPS, sim-verified) [crafted]; Imperial Red Mantle (8250, -0.54 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Cloak (8259, -0.93 DPS) [world_drop]; Nightscape Cloak (8195, -1.17 DPS) [crafted]; Blackflame Cape (13109, -1.36 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 100.8 ranged_attack_power points (7.35 DPS) | yes | Knight's Chain Armor (220828, -1.47 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -1.98 DPS) [quest]; Tough Scorpid Breastplate (8203, -2.89 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 48.0 ranged_attack_power points (3.50 DPS) | yes | Stalker's Mail Bracers (252570, -0.58 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.77 DPS) [dungeon]; Bloodlust Bracelets (14807, -0.93 DPS) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hibernal Gloves (8110, -0.80 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.93 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -1.22 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 70.6 ranged_attack_power points (5.15 DPS) | yes | Stalker's Mail Belt (252588, -0.76 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20088, -1.66 DPS) [rep]; Highlander's Leather Girdle (20115, -1.66 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 91.5 ranged_attack_power points (6.67 DPS) | yes | Imperial Red Pants (8251, -1.27 DPS, sim-verified) [world_drop]; Tough Scorpid Leggings (8206, -1.33 DPS) [crafted]; Basilisk Hide Pants (1718, -1.77 DPS) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 72.1 ranged_attack_power points (5.26 DPS) | yes | Albino Crocscale Boots (17728, -0.59 DPS) [dungeon]; Rancor Boots (11865, -0.95 DPS) [quest]; Stalker's Mail Sabatons (252576, -0.99 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 44.8 ranged_attack_power points (3.27 DPS) | yes | Ring of the Underwood (2951, -0.93 DPS) [world_drop]; Falcon's Hook (7552, -1.17 DPS) [dungeon]; Protector's Band (19516, -1.17 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.0 ranged_attack_power points (3.20 DPS) | yes | Falcon's Hook (7552, -1.10 DPS) [dungeon]; Protector's Band (19516, -1.10 DPS) [rep]; Ring of the Underwood (2951, -1.52 DPS, sim-verified) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.13 DPS) [world_drop]; Felbough Arm (273037, -1.28 DPS) [world]; Hanzo Sword (8190, -5.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -1.44 DPS, sim-verified) [dungeon]; Megashot Rifle (17717, -2.49 DPS) [dungeon]; Dark Iron Rifle (16004, -2.93 DPS) [crafted] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Stalker's Mail Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 854, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60 (dwarf, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 349.2. Weights run: 8.5s. Verify run: 57.0s. 1594 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.220 ± 0.033, crit=2.552 ± 0.069 per rating point (14 rating = 1%, 35.733 per %), hit=3.933 ± 0.158 per rating point (10 rating = 1%, 39.332 per %), melee_haste=13.794 ± 1.882, intellect=2.995 ± 0.082, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 217.2 ranged_attack_power points (20.32 DPS) | yes | Field Marshal's Chain Helm (231580, -3.37 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -5.12 DPS) [vendor]; Outlaw's Collar (279253, -10.10 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | 48.3 ranged_attack_power points (4.52 DPS) | yes | Horizon Choker (13085, -0.60 DPS) [world_drop]; Skibi's Pendant (13089, -0.60 DPS) [world_drop]; Opaline Medallion (16623, -0.88 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 122.7 ranged_attack_power points (11.48 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (349.2 DPS) | yes | Sergeant Major's Cape (16337, -1.74 DPS) [pvp]; Dark Phantom Cape (13122, -1.87 DPS) [world_drop]; Cloak of the Honor Guard (20073, -2.02 DPS, sim-verified) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 194.2 ranged_attack_power points (18.17 DPS) | yes | Field Marshal's Chain Breastplate (16466, -1.22 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -1.22 DPS) [vendor]; Icemail Jerkin (1981, -7.48 DPS, sim-verified) [world_drop] |
| wrist | Windtalker's Wristguards (19582) | Silverwing Sentinels [rep] | 62.0 ranged_attack_power points (5.80 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Spitfire Bracers (20481, -0.56 DPS) [crafted]; Forest Stalker's Bracers (19587, -1.64 DPS, sim-verified) [rep] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (349.2 DPS) | yes | Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Marshal's Chain Grips (231560, -0.71 DPS) [pvp]; Theramore Gloves (271948, -2.41 DPS, sim-verified) [rep] |
| waist | Highlander's Mail Girdle (20044) (or Highlander's Lizardhide Girdle (20046)) | Samuel Hawke [vendor] | 86.6 ranged_attack_power points (8.11 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Highlander's Lizardhide Girdle (20046, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.22 DPS) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 223.5 ranged_attack_power points (20.91 DPS) | yes | Marshal's Chain Legguards (231577, -3.96 DPS) [pvp]; Marshal's Chain Legplates (231558, -5.72 DPS) [vendor]; Sentinel's Leather Pants (237818, -6.09 DPS) [vendor] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 84.7 ranged_attack_power points (7.92 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 89.1 ranged_attack_power points (8.33 DPS) | yes | Masons Fraternity Ring (9533, -4.11 DPS) [quest]; Glowing Crystal Ring (18402, -4.97 DPS) [quest]; Band of the Penitent (13217, -4.99 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 59.3 ranged_attack_power points (5.55 DPS) | yes | Masons Fraternity Ring (9533, -1.76 DPS, sim-verified) [quest]; Glowing Crystal Ring (18402, -2.19 DPS) [quest]; Band of the Penitent (13217, -2.21 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (349.2 DPS) | yes | - |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (349.2 DPS) | yes | Ankh of Life (1713, -1.66 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (349.2 DPS) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -0.42 DPS) [pvp]; Electrified Dagger (19100, -6.37 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (349.2 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -9.54 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Windtalker's Wristguards; hands: Mourning Handguards; waist: Highlander's Mail Girdle; legs: Sentinel's Chain Leggings; feet: Timmy's Galoshes; finger1: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1594, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

### Band 60, raid preset (dwarf, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 731.4. Weights run: 8.3s. Verify run: 50.2s. 1594 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.305 ± 0.034, crit=2.804 ± 0.069 per rating point (14 rating = 1%, 39.250 per %), hit=4.278 ± 0.195 per rating point (10 rating = 1%, 42.781 per %), melee_haste=17.463 ± 2.369, intellect=1.505 ± 0.027, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 165.8 ranged_attack_power points (27.86 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -2.40 DPS) [pvp]; Outlaw's Collar (279253, -8.78 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | 49.6 ranged_attack_power points (8.33 DPS) | yes | Skibi's Pendant (13089, -1.11 DPS) [world_drop]; Sunprism Pendant (20649, -1.67 DPS) [quest]; Horizon Choker (13085, -4.79 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 128.5 ranged_attack_power points (21.58 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (731.4 DPS) | yes | Dark Phantom Cape (13122, -3.36 DPS) [world_drop]; Cloak of the Honor Guard (20073, -3.83 DPS, sim-verified) [rep]; Cloak of Flames (3475, -4.50 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 204.3 ranged_attack_power points (34.32 DPS) | yes | Field Marshal's Chain Breastplate (16466, -5.82 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -5.82 DPS) [vendor]; Savage Gladiator Chain (11726, -17.68 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 62.8 ranged_attack_power points (10.55 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.11 DPS) [pvp]; Windtalker's Wristguards (19582, -2.14 DPS) [rep] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (731.4 DPS) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, -0.21 DPS) [rep] |
| waist | Highlander's Chain Girdle (20043) (or Highlander's Leather Girdle (20045)) | The League of Arathor [rep] | 73.2 ranged_attack_power points (12.31 DPS) | yes | Marshal's Chain Girdle (16464, +0.00 DPS) [pvp]; Highlander's Leather Girdle (20045, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.95 DPS) [quest] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 236.9 ranged_attack_power points (39.80 DPS) | yes | Marshal's Chain Legguards (231577, -11.30 DPS) [pvp]; Sentinel's Leather Pants (237818, -11.63 DPS) [vendor]; Marshal's Chain Legplates (231558, -14.69 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 82.2 ranged_attack_power points (13.81 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -2.11 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 96.0 ranged_attack_power points (16.13 DPS) | yes | Masons Fraternity Ring (9533, -8.36 DPS) [quest]; Band of the Penitent (13217, -9.54 DPS) [quest]; Don Rodrigo's Band (21563, -9.54 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 62.8 ranged_attack_power points (10.55 DPS) | yes | Masons Fraternity Ring (9533, -3.31 DPS, sim-verified) [quest]; Band of the Penitent (13217, -3.95 DPS) [quest]; Don Rodrigo's Band (21563, -3.95 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (731.4 DPS) | yes | - |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (731.4 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (731.4 DPS) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Warmonger (13052, -3.58 DPS) [world_drop]; Electrified Dagger (19100, -12.51 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (731.4 DPS) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -5.01 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Mourning Handguards; waist: Highlander's Chain Girdle; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1594, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4988 Burning Obsidian Band

## Horde

### Band 20 (troll, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 68.8. Weights run: 7.5s. Verify run: 8.0s. 254 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.173 ± 0.012, crit=0.679 ± 0.013 per rating point (14 rating = 1%, 9.505 per %), hit=1.150 ± 0.041 per rating point (10 rating = 1%, 11.498 per %), melee_haste=5.252 ± 0.566, intellect=1.086 ± 0.030, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.4 ranged_attack_power points (1.82 DPS) | yes | Trapper's Leather Hood (252505, -1.38 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.43 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 19.0 ranged_attack_power points (1.36 DPS) | yes | Erudite's Amulet (277204, -0.95 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.9 ranged_attack_power points (1.14 DPS) | yes | Magician's Mantle (12998, -0.68 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.82 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.25 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.45 DPS) [world_drop]; Bristlebark Cape (14571, -0.45 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 25.5 ranged_attack_power points (1.82 DPS) | yes | Brawler's Leather Armor (252490, -0.25 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.46 DPS) [crafted]; Prospector's Chestpiece (14562, -0.46 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.9 ranged_attack_power points (1.14 DPS) | yes | Wolf Bracers (4794, -0.25 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.45 DPS) [world_drop]; Spare Part Bindings (279875, -0.45 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.0 ranged_attack_power points (1.36 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.22 DPS) [crafted]; Forest Leather Gloves (3058, -0.45 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.29 DPS) | yes | Trapper's Leather Belt (252429, -0.15 DPS) [crafted]; Deviate Scale Belt (6468, -0.15 DPS) [crafted]; Dark Leather Belt (4249, -0.38 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.23 DPS) [world]; Trapper's Leather Pants (252501, -0.27 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.4 ranged_attack_power points (1.82 DPS) | yes | Trapper's Leather Boots (252440, -0.26 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.45 DPS) [dungeon]; Footpads of the Fang (10411, -0.45 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 12.9 ranged_attack_power points (0.92 DPS) | yes | Bounty Hunter's Ring (5351, -0.24 DPS) [quest]; Loop of Sacrifice (281673, -0.53 DPS) [quest]; The 1 Ring (8350, -0.62 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.7 ranged_attack_power points (0.91 DPS) | yes | Bounty Hunter's Ring (5351, -0.25 DPS, sim-verified) [quest]; Loop of Sacrifice (281673, -0.52 DPS) [quest]; The 1 Ring (8350, -0.60 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 34.0 ranged_attack_power points (2.43 DPS) | yes | Bronze Dory (250603, -0.19 DPS) [crafted]; Crescent Staff (6505, -0.30 DPS) [quest]; Staff of Orgrimmar (15444, -0.60 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (12.85 DPS) | yes | Venomstrike (6469, -0.86 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -0.95 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.10 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 254, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18856 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade

### Band 30 (troll, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 104.9. Weights run: 7.9s. Verify run: 18.4s. 446 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.114 ± 0.013, crit=0.981 ± 0.018 per rating point (14 rating = 1%, 13.733 per %), hit=1.617 ± 0.052 per rating point (10 rating = 1%, 16.165 per %), melee_haste=8.587 ± 0.754, intellect=1.204 ± 0.038, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 31.1 ranged_attack_power points (2.32 DPS) | yes | Brawler's Leather Hood (252504, -0.47 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.70 DPS) [world]; Resilient Cap (14401, -0.87 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 24.9 ranged_attack_power points (1.85 DPS) | yes | Kaleidoscope Chain (13084, -0.57 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.59 DPS, sim-verified) [world]; Erudite's Amulet (277204, -1.39 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.3 ranged_attack_power points (2.55 DPS) | yes | Mantle of Thieves (2264, -0.13 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.93 DPS) [crafted]; Insignia Mantle (4721, -0.93 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.46 DPS) [dungeon]; Swiftrunner Cape (6745, -0.46 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 46.1 ranged_attack_power points (3.43 DPS) | yes | Dusky Leather Armor (7374, -0.19 DPS) [crafted]; Green Leather Armor (4255, -0.86 DPS) [crafted]; Trapper's Leather Tunic (252509, -1.04 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.7 ranged_attack_power points (1.39 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.23 DPS) [world_drop]; Loamflake Bracers (15462, -0.23 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 28.0 ranged_attack_power points (2.09 DPS) | yes | Pilferer's Gloves (7358, -0.39 DPS, sim-verified) [crafted]; Hotshot Pilot's Gloves (9491, -0.44 DPS) [dungeon]; Braced Handguards (6784, -0.46 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 35.2 ranged_attack_power points (2.62 DPS) | yes | Skulker's Leather Belt (252520, -0.40 DPS, sim-verified) [crafted]; Green Leather Belt (4257, -0.70 DPS) [crafted]; Deftkin Belt (16659, -0.80 DPS) [quest] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Petrolspill Leggings (9509, -0.17 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.23 DPS) [crafted]; Ferine Leggings (6690, -0.23 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 24.9 ranged_attack_power points (1.85 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.23 DPS) [quest] |
| finger1 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 18.7 ranged_attack_power points (1.39 DPS) | yes | Monkey Ring (6748, -0.23 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Band of the Fist (17694, -0.46 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.7 ranged_attack_power points (1.39 DPS) | yes | Monkey Ring (6748, -0.23 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Band of the Fist (17694, -0.46 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kam's Walking Stick (2280, -0.93 DPS) [dungeon]; Armor Piercer (6679, -0.97 DPS) [dungeon]; Thornspike (6681, -1.25 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 251.1 ranged_attack_power points (18.69 DPS) | yes | Ironweaver (13137, -0.54 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -0.84 DPS) [world_drop]; Chesterfall Musket (7729, -1.47 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Legionnaire's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 170.4. Weights run: 8.3s. Verify run: 21.9s. 671 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.321 ± 0.038, crit=1.797 ± 0.060 per rating point (14 rating = 1%, 25.156 per %), hit=2.437 ± 0.109 per rating point (10 rating = 1%, 24.367 per %), melee_haste=11.471 ± 1.194, intellect=1.235 ± 0.051, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 46.9 ranged_attack_power points (3.76 DPS) | yes | Nightscape Headband (8176, -0.56 DPS) [crafted]; Guard's Chain Helm (250499, -0.56 DPS) [crafted]; Skullsplitter Helm (1624, -0.83 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 36.5 ranged_attack_power points (2.93 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.38 DPS) [world]; Kaleidoscope Chain (13084, -1.47 DPS) [world_drop]; Ethereal Talisman (4430, -1.87 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.5 ranged_attack_power points (2.93 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.27 DPS) [dungeon]; Flintrock Shoulders (7755, -0.27 DPS) [dungeon] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.27 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 104.1 ranged_attack_power points (8.36 DPS) | yes | Tough Scorpid Breastplate (8203, -3.57 DPS, sim-verified) [crafted]; Wolffear Harness (13110, -3.82 DPS) [world_drop]; Nightscape Tunic (8175, -4.36 DPS) [crafted] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.7 ranged_attack_power points (2.46 DPS) | yes | Imperial Leather Bracers (4061, -0.33 DPS) [dungeon]; Dusky Bracers (7378, -0.34 DPS, sim-verified) [crafted]; Green Leather Bracers (4259, -0.37 DPS) [crafted] |
| hands | Dragonscale Gauntlets (8347) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tough Scorpid Gloves (8204, -0.06 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.26 DPS) [crafted]; Aurora Gloves (4042, -0.69 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 37.3 ranged_attack_power points (3.00 DPS) | yes | Scorpashi Sash (14652, -0.33 DPS) [world_drop]; Defiler's Chain Girdle (20153, -0.33 DPS) [rep]; Defiler's Leather Girdle (20192, -0.59 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 69.7 ranged_attack_power points (5.60 DPS) | yes | Triprunner Dungarees (9624, -1.07 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -1.87 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.87 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 45.2 ranged_attack_power points (3.63 DPS) | yes | Swampwalker Boots (2276, -0.16 DPS) [world_drop]; Worn Running Boots (9398, -0.16 DPS) [dungeon]; Imperial Leather Boots (6431, -0.69 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 33.2 ranged_attack_power points (2.67 DPS) | yes | Legionnaire's Band (19512, -0.53 DPS) [rep]; Field Researcher's Loop (281634, -0.80 DPS) [quest]; Falcon's Hook (7552, -2.39 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Band (19512, +0.00 DPS) [rep]; Falcon's Hook (7552, -0.26 DPS, sim-verified) [dungeon]; Field Researcher's Loop (281634, -0.27 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.29 DPS) [world]; Steel Spear (250605, -0.65 DPS) [crafted]; Claw of the Shadowmancer (2912, -3.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -2.35 DPS) [world_drop]; Baelog's Shortbow (9400, -2.69 DPS) [dungeon]; The Silencer (13138, -3.33 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Dragonscale Gauntlets; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 671, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (troll, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 230.2. Weights run: 8.8s. Verify run: 30.2s. 845 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.202 ± 0.036, crit=1.989 ± 0.066 per rating point (14 rating = 1%, 27.844 per %), hit=2.396 ± 0.110 per rating point (10 rating = 1%, 23.957 per %), melee_haste=9.365 ± 1.267, intellect=1.883 ± 0.062, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Tough Scorpid Helm (8208, -0.42 DPS, sim-verified) [crafted]; Helm of Fire (8348, -0.67 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.6 ranged_attack_power points (3.03 DPS) | yes | Scout's Medallion (19536, -0.47 DPS) [rep]; Woven Ivy Necklace (19159, -0.93 DPS) [quest] |
| shoulder | Stalker's Mail Shoulders (252566) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Tough Scorpid Shoulders (8207, -0.43 DPS, sim-verified) [crafted]; Imperial Red Mantle (8250, -0.54 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Cloak (8259, -0.93 DPS) [world_drop]; Nightscape Cloak (8195, -1.17 DPS) [crafted]; Blackflame Cape (13109, -1.32 DPS, sim-verified) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 100.8 ranged_attack_power points (7.35 DPS) | yes | Stone Guard's Chain Armor (220827, -1.72 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -1.98 DPS) [quest]; Tough Scorpid Breastplate (8203, -2.89 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 48.0 ranged_attack_power points (3.50 DPS) | yes | Stalker's Mail Bracers (252570, -0.58 DPS) [crafted]; Bracers of the Stone Princess (17714, -0.77 DPS) [dungeon]; Bloodlust Bracelets (14807, -0.93 DPS) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hibernal Gloves (8110, -0.80 DPS) [world_drop]; Tough Scorpid Gloves (8204, -0.93 DPS) [crafted]; First Sergeant's Chain Gauntlets (220830, -2.06 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 70.6 ranged_attack_power points (5.15 DPS) | yes | Stalker's Mail Belt (252588, -0.51 DPS) [crafted]; Defiler's Chain Girdle (20151, -1.66 DPS) [rep]; Defiler's Leather Girdle (20193, -1.66 DPS) [rep] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 91.5 ranged_attack_power points (6.67 DPS) | yes | Tough Scorpid Leggings (8206, -1.33 DPS) [crafted]; Imperial Red Pants (8251, -1.69 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -1.77 DPS) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 72.1 ranged_attack_power points (5.26 DPS) | yes | Albino Crocscale Boots (17728, -0.59 DPS) [dungeon]; Rancor Boots (11865, -0.95 DPS) [quest]; Stalker's Mail Sabatons (252576, -0.99 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 44.8 ranged_attack_power points (3.27 DPS) | yes | Ring of the Underwood (2951, -0.93 DPS) [world_drop]; Falcon's Hook (7552, -1.17 DPS) [dungeon]; Legionnaire's Band (19511, -1.17 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.0 ranged_attack_power points (3.20 DPS) | yes | Ring of the Underwood (2951, -1.07 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -1.10 DPS) [dungeon]; Legionnaire's Band (19511, -1.10 DPS) [rep] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.67 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.13 DPS) [world_drop]; Felbough Arm (273037, -1.28 DPS) [world]; Hanzo Sword (8190, -5.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Megashot Rifle (17717, -2.49 DPS) [dungeon]; Gryphonwing Long Bow (13022, -2.52 DPS, sim-verified) [dungeon]; Dark Iron Rifle (16004, -2.93 DPS) [crafted] |

**New at 50:** head: Eye of Theradras; neck: Skibi's Pendant; shoulder: Stalker's Mail Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 845, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 345.0. Weights run: 8.5s. Verify run: 56.7s. 1595 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.220 ± 0.033, crit=2.552 ± 0.069 per rating point (14 rating = 1%, 35.733 per %), hit=3.933 ± 0.158 per rating point (10 rating = 1%, 39.332 per %), melee_haste=13.794 ± 1.882, intellect=2.995 ± 0.082, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 217.2 ranged_attack_power points (20.32 DPS) | yes | Warlord's Chain Helmet (16566, -3.37 DPS) [vendor]; Warlord's Chain Helm (231571, -3.37 DPS) [vendor]; Outlaw's Collar (279253, -8.65 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | 48.3 ranged_attack_power points (4.52 DPS) | yes | Horizon Choker (13085, -0.60 DPS) [world_drop]; Skibi's Pendant (13089, -0.60 DPS) [world_drop]; Sunprism Pendant (20649, -0.90 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 122.7 ranged_attack_power points (11.48 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (345.0 DPS) | yes | Dark Phantom Cape (13122, -1.87 DPS) [world_drop]; Darkspear Raider's Cloak (272063, -1.91 DPS) [vendor]; Deathguard's Cloak (20068, -1.98 DPS, sim-verified) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 194.2 ranged_attack_power points (18.17 DPS) | yes | Warlord's Chain Chestpiece (16565, -1.22 DPS) [vendor]; Warlord's Chain Hauberk (231573, -1.22 DPS) [vendor]; Icemail Jerkin (1981, -7.70 DPS, sim-verified) [world_drop] |
| wrist | Windtalker's Wristguards (19582) | Warsong Outriders [rep] | 62.0 ranged_attack_power points (5.80 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Spitfire Bracers (20481, -0.56 DPS) [crafted]; Forest Stalker's Bracers (19587, -1.01 DPS, sim-verified) [rep] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (345.0 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, -1.85 DPS, sim-verified) [vendor] |
| waist | Defiler's Lizardhide Girdle (20171) (or Defiler's Mail Girdle (20195)) | The Defilers [rep] | 86.6 ranged_attack_power points (8.11 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Defiler's Mail Girdle (20195, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.22 DPS) [quest] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 187.8 ranged_attack_power points (17.57 DPS) | yes | General's Chain Legguards (231574, -0.62 DPS) [pvp]; General's Chain Legplates (231567, -2.37 DPS) [vendor]; Outrider's Leather Pants (22740, -6.48 DPS, sim-verified) [rep] |
| feet | Timmy's Galoshes (13402) | Stratholme: Timmy the Cruel [dungeon] | 84.7 ranged_attack_power points (7.92 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.49 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 89.1 ranged_attack_power points (8.33 DPS) | yes | Masons Fraternity Ring (9533, -4.11 DPS) [quest]; Glowing Crystal Ring (18402, -4.97 DPS) [quest]; Band of the Penitent (13217, -4.99 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 59.3 ranged_attack_power points (5.55 DPS) | yes | Masons Fraternity Ring (9533, -1.51 DPS, sim-verified) [quest]; Glowing Crystal Ring (18402, -2.19 DPS) [quest]; Band of the Penitent (13217, -2.21 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (345.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (345.0 DPS) | yes | Ankh of Life (1713, -1.70 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (345.0 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -0.42 DPS) [pvp]; Glacial Blade (19099, -5.96 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (345.0 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Bow of Searing Arrows (2825, -10.73 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Windtalker's Wristguards; hands: Mourning Handguards; waist: Defiler's Lizardhide Girdle; legs: Outrider's Chain Leggings; feet: Timmy's Galoshes; finger1: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1595, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (troll, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 727.8. Weights run: 8.3s. Verify run: 50.7s. 1595 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.305 ± 0.034, crit=2.804 ± 0.069 per rating point (14 rating = 1%, 39.250 per %), hit=4.278 ± 0.195 per rating point (10 rating = 1%, 42.781 per %), melee_haste=17.463 ± 2.369, intellect=1.505 ± 0.027, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 165.8 ranged_attack_power points (27.86 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -9.12 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | 49.6 ranged_attack_power points (8.33 DPS) | yes | Skibi's Pendant (13089, -1.11 DPS) [world_drop]; Sunprism Pendant (20649, -1.67 DPS) [quest]; Woven Ivy Necklace (19159, -3.33 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 128.5 ranged_attack_power points (21.58 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (727.8 DPS) | yes | Deathguard's Cloak (20068, -3.20 DPS) [rep]; Dark Phantom Cape (13122, -3.36 DPS) [world_drop]; Cloak of Flames (3475, -4.46 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 204.3 ranged_attack_power points (34.32 DPS) | yes | Warlord's Chain Chestpiece (16565, -5.82 DPS) [vendor]; Warlord's Chain Hauberk (231573, -5.82 DPS) [vendor]; Savage Gladiator Chain (11726, -17.94 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 62.8 ranged_attack_power points (10.55 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.11 DPS) [pvp]; Windtalker's Wristguards (19582, -2.14 DPS) [rep] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (727.8 DPS) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190)) | The Defilers [rep] | 73.2 ranged_attack_power points (12.31 DPS) | yes | General's Chain Girdle (16572, +0.00 DPS) [pvp]; Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; Sagebrush Girdle (17778, -0.95 DPS) [quest] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 197.7 ranged_attack_power points (33.21 DPS) | yes | General's Chain Legguards (231574, -4.71 DPS) [pvp]; General's Chain Legplates (231567, -8.10 DPS) [vendor]; Outrider's Leather Pants (22740, -12.29 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 82.2 ranged_attack_power points (13.81 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Beaststalker's Boots (16675, -3.35 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 96.0 ranged_attack_power points (16.13 DPS) | yes | Masons Fraternity Ring (9533, -8.36 DPS) [quest]; Band of the Penitent (13217, -9.54 DPS) [quest]; Don Rodrigo's Band (21563, -9.54 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 62.8 ranged_attack_power points (10.55 DPS) | yes | Masons Fraternity Ring (9533, -3.30 DPS, sim-verified) [quest]; Band of the Penitent (13217, -3.95 DPS) [quest]; Don Rodrigo's Band (21563, -3.95 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (727.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (727.8 DPS) | yes | Molten Heart of the Mountain (249470, -5.57 DPS, sim-verified) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-verified (727.8 DPS) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Warmonger (13052, -3.58 DPS) [world_drop]; Glacial Blade (19099, -9.09 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-verified (727.8 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -2.33 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Mourning Handguards; waist: Defiler's Chain Girdle; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 1595, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

