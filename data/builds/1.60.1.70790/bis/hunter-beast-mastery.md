# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5320000100000000-0000000000000000-000000000000000000)

Set DPS (verified): 91.0. Weights run: 5.0s. Verify run: 6.6s. 298 eligible items had no known source.

Race sweep: night-elf 91.0 ± 0.1, high-order-skyborne 89.9 ± 0.1, human 89.7 ± 0.1, dwarf 87.0 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.158 ± 0.010, crit=0.667 ± 0.011 per rating point (14 rating = 1%, 9.342 per %), hit=1.082 ± 0.032 per rating point (10 rating = 1%, 10.817 per %), melee_haste=4.895 ± 0.476, intellect=1.650 ± 0.022, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.3 ranged_attack_power points (2.51 DPS) | yes | Trapper's Leather Hood (252505, -1.02 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.69 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.9 ranged_attack_power points (1.88 DPS) | yes | Erudite's Amulet (277204, -1.24 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.8 ranged_attack_power points (1.57 DPS) | yes | Magician's Mantle (12998, -0.09 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -0.91 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.32 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.63 DPS) [world_drop]; Bristlebark Cape (14571, -0.63 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.7 ranged_attack_power points (3.45 DPS) | yes | Trapper's Leather Armor (252491, -0.29 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -1.25 DPS) [crafted]; Mystic's Wrap (14369, -1.36 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.8 ranged_attack_power points (1.57 DPS) | yes | Wolf Bracers (4794, -0.32 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.63 DPS) [world_drop]; Bravo's Armbands (270015, -0.63 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 18.9 ranged_attack_power points (1.88 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.14 DPS) [crafted]; Riding Gloves (1304, -0.45 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.79 DPS) | yes | Trapper's Leather Belt (252429, -0.04 DPS) [crafted]; Deviate Scale Belt (6468, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.48 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Duty Bound Leggings (279868, -0.29 DPS) [quest]; Trapper's Leather Pants (252501, -0.96 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.3 ranged_attack_power points (2.51 DPS) | yes | Trapper's Leather Boots (252440, -0.26 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.63 DPS) [dungeon]; Footpads of the Fang (10411, -0.63 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.2 ranged_attack_power points (1.61 DPS) | yes | Volcanic Rock Ring (12053, -1.12 DPS) [world_drop]; Minor Channeling Ring (1449, -1.28 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.6 ranged_attack_power points (1.25 DPS) | yes | Volcanic Rock Ring (12053, -0.75 DPS, sim-verified) [world_drop]; Minor Channeling Ring (1449, -0.93 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 36.7 ranged_attack_power points (3.64 DPS) | yes | Bronze Dory (250603, -0.43 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.76 DPS) [world]; Piercing Axe (6094, -1.76 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (17.82 DPS) | yes | Venomstrike (6469, -0.31 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.33 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -4.30 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 5320000505000100-0000000000000000-000000000000000000)

Set DPS (verified): 146.2. Weights run: 4.8s. Verify run: 14.3s. 494 eligible items had no known source.

Race sweep: night-elf 146.2 ± 0.1, human 144.8 ± 0.1, high-order-skyborne 143.6 ± 0.1, dwarf 142.9 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.055 ± 0.010, crit=0.899 ± 0.014 per rating point (14 rating = 1%, 12.583 per %), hit=1.423 ± 0.044 per rating point (10 rating = 1%, 14.225 per %), melee_haste=not significant (-0.569 ± 0.646), intellect=0.556 ± 0.019, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.5 ranged_attack_power points (3.42 DPS) | yes | Brawler's Leather Hood (252504, -0.65 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -1.03 DPS) [world]; Resilient Cap (14401, -2.05 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 24.4 ranged_attack_power points (2.74 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.02 DPS) [world]; Kaleidoscope Chain (13084, -1.12 DPS) [world_drop]; Erudite's Amulet (277204, -2.05 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.6 ranged_attack_power points (3.76 DPS) | yes | Mantle of Thieves (2264, -0.34 DPS) [dungeon]; Dark Leather Shoulders (4252, -1.37 DPS) [crafted]; Insignia Mantle (4721, -1.37 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.68 DPS) [dungeon]; Cape of the Brotherhood (5193, -1.03 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Tunic of Westfall (2041, -0.25 DPS) [quest]; Green Leather Armor (4255, -0.78 DPS) [crafted]; Dusky Leather Armor (7374, -1.55 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.3 ranged_attack_power points (2.05 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.34 DPS) [world]; Forest Leather Bracers (3202, -0.34 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.5 ranged_attack_power points (3.08 DPS) | yes | Pilferer's Gloves (7358, -0.79 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -1.03 DPS) [world_drop]; Gloves of the Fang (10413, -1.03 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 30.8 ranged_attack_power points (3.45 DPS) | yes | Highlander's Chain Girdle (20090, -0.77 DPS) [rep]; Highlander's Leather Girdle (20117, -0.77 DPS) [rep]; Skulker's Leather Belt (252520, -1.21 DPS, sim-verified) [crafted] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Petrolspill Leggings (9509, -0.31 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.34 DPS) [crafted]; Ferine Leggings (6690, -1.13 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 24.4 ranged_attack_power points (2.74 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, -0.68 DPS) [dungeon]; Footpads of the Fang (10411, -0.68 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 18.3 ranged_attack_power points (2.05 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.68 DPS) [dungeon]; Seal of Wrynn (2933, -0.78 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.3 ranged_attack_power points (2.05 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.68 DPS) [dungeon]; Seal of Wrynn (2933, -0.78 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Claw of the Shadowmancer (2912, +0.00 DPS) [world_drop]; Armor Piercer (6679, -1.25 DPS) [dungeon]; Kam's Walking Stick (2280, -1.41 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 250.7 ranged_attack_power points (28.09 DPS) | yes | Ironweaver (13137, -1.31 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.54 DPS) [world_drop]; Chesterfall Musket (7729, -2.18 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 494, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 227.2. Weights run: 5.3s. Verify run: 17.8s. 827 eligible items had no known source.

Race sweep: night-elf 227.2 ± 0.1, human 225.3 ± 0.1, high-order-skyborne 225.2 ± 0.1, dwarf 223.2 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.869 ± 0.027, crit=0.873 ± 0.042 per rating point (14 rating = 1%, 12.224 per %), hit=1.865 ± 0.077 per rating point (10 rating = 1%, 18.649 per %), melee_haste=8.174 ± 0.789, intellect=2.305 ± 0.031, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpseshroud (10574, +0.00 DPS) [dungeon]; Thinking Cap (2624, -0.52 DPS) [world]; Miner's Hat of the Deep (9429, -0.52 DPS) [dungeon] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 31.6 ranged_attack_power points (3.28 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Kaleidoscope Chain (13084, -1.13 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.36 DPS) [dungeon] |
| shoulder | Nightscape Shoulders (8192) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.59 DPS) [rep]; Imperial Cloak (6432, -0.84 DPS) [dungeon]; Parachute Cloak (10518, -0.84 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 87.5 ranged_attack_power points (9.09 DPS) | yes | Garb of Florid Feathers (273026, -2.82 DPS) [dungeon]; Tough Scorpid Breastplate (8203, -2.94 DPS) [crafted]; Quillward Harness (10583, -3.43 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.9 ranged_attack_power points (3.52 DPS) | yes | Green Leather Bracers (4259, -0.54 DPS) [crafted]; Warden's Wristbands (14600, -0.71 DPS) [world_drop]; Windchaser Cuffs (14429, -0.77 DPS) [world_drop] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.48 DPS) [crafted]; Aurora Gloves (4042, -0.84 DPS) [world_drop]; Mistscape Gloves (6428, -1.79 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.7 ranged_attack_power points (4.12 DPS) | yes | Deathmage Sash (10771, -0.53 DPS) [dungeon]; Highlander's Lizardhide Girdle (20104, -0.53 DPS) [rep]; Highlander's Mail Girdle (20119, -0.53 DPS) [vendor] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 60.3 ranged_attack_power points (6.26 DPS) | yes | Ferine Leggings (6690, -0.69 DPS) [dungeon]; Triprunner Dungarees (9624, -0.89 DPS) [quest]; Trapper's Leather Legguards (252517, -1.90 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swampwalker Boots (2276, -1.08 DPS) [world_drop]; Worn Running Boots (9398, -1.08 DPS) [dungeon]; Briar Tredders (10582, -1.31 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.7 ranged_attack_power points (2.98 DPS) | yes | Protector's Band (19515, -0.60 DPS) [rep]; Disengagement Ring (276202, -0.60 DPS) [vendor]; Field Researcher's Loop (281634, -0.89 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.8 ranged_attack_power points (2.68 DPS) | yes | Protector's Band (19515, -0.30 DPS) [rep]; Disengagement Ring (276202, -0.30 DPS) [vendor]; Field Researcher's Loop (281634, -0.60 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -2.77 DPS) [world]; Steel Spear (250605, -2.78 DPS) [crafted]; Ruthless Shiv (7758, -3.74 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -2.44 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.26 DPS) [dungeon]; Swiftwind (13038, -3.37 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 827, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 326.6. Weights run: 5.4s. Verify run: 24.3s. 1062 eligible items had no known source.

Race sweep: night-elf 326.6 ± 0.2, human 324.6 ± 0.2, high-order-skyborne 324.4 ± 0.2, dwarf 321.9 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.107 ± 0.030, crit=1.777 ± 0.052 per rating point (14 rating = 1%, 24.877 per %), hit=2.450 ± 0.096 per rating point (10 rating = 1%, 24.501 per %), melee_haste=11.764 ± 1.106, intellect=2.780 ± 0.041, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Tough Scorpid Helm (8208, -0.19 DPS) [crafted]; Chief Architect's Monocle (11839, -1.70 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.4 ranged_attack_power points (4.24 DPS) | yes | Sentinel's Medallion (19539, -0.33 DPS) [rep]; Horizon Choker (13085, -4.95 DPS, sim-verified) [world_drop] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 58.9 ranged_attack_power points (6.17 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Imperial Red Mantle (8250, -0.50 DPS) [world_drop]; Ironfeather Shoulders (15067, -3.83 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.33 DPS) [dungeon]; Expeditionary's Cape (271921, -0.81 DPS) [rep] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blazewind Breastplate (11193, -2.90 DPS) [quest]; Garb of Florid Feathers (273026, -3.08 DPS) [dungeon]; Knight's Chain Armor (220828, -3.20 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.6 ranged_attack_power points (4.89 DPS) | yes | Forgotten Wraps (9433, -0.41 DPS) [world_drop]; Bracers of the Stone Princess (17714, -0.49 DPS) [dungeon]; Stalker's Mail Bracers (252570, -2.28 DPS, sim-verified) [crafted] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rockgrip Gauntlets (17736, -1.04 DPS) [dungeon]; Tough Scorpid Gloves (8204, -1.30 DPS) [crafted]; Hibernal Gloves (8110, -1.55 DPS, sim-verified) [world_drop] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 71.3 ranged_attack_power points (7.48 DPS) | yes | Sagebrush Girdle (17778, -1.10 DPS, sim-verified) [quest]; Dawnspire Cord (12466, -1.94 DPS) [dungeon]; Highlander's Lizardhide Girdle (20103, -1.95 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 108.9 ranged_attack_power points (11.42 DPS) | yes | Imperial Red Pants (8251, -2.16 DPS) [world_drop]; Tough Scorpid Leggings (8206, -2.96 DPS) [crafted]; Knight's Chain Legplates (220832, -3.44 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 75.7 ranged_attack_power points (7.94 DPS) | yes | Albino Crocscale Boots (17728, -1.42 DPS) [dungeon]; Tough Scorpid Boots (8209, -1.99 DPS) [crafted]; Stalker's Mail Sabatons (252576, -2.31 DPS, sim-verified) [crafted] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 45.6 ranged_attack_power points (4.78 DPS) | yes | Masons Fraternity Ring (9533, -0.22 DPS) [quest]; Ring of the Underwood (2951, -1.53 DPS) [world_drop]; Falcon's Hook (7552, -1.85 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.5 ranged_attack_power points (4.67 DPS) | yes | Masons Fraternity Ring (9533, -0.69 DPS, sim-verified) [quest]; Ring of the Underwood (2951, -1.41 DPS) [world_drop]; Falcon's Hook (7552, -1.73 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, -0.03 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Staff of Jordan (873, -3.37 DPS, sim-verified) [world_drop] |
| off_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | 28.0 ranged_attack_power points (2.93 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gryphonwing Long Bow (13022, -0.16 DPS) [dungeon]; Hurricane (2824, -1.50 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -1.92 DPS) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Mindseye Circle; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Ebon Shiv; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1062, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 487.1. Weights run: 5.5s. Verify run: 60.6s. 2287 eligible items had no known source.

Race sweep: night-elf 487.1 ± 0.2, high-order-skyborne 485.0 ± 0.2, human 484.9 ± 0.2, dwarf 481.9 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.171 ± 0.025, crit=2.406 ± 0.053 per rating point (14 rating = 1%, 33.677 per %), hit=3.050 ± 0.109 per rating point (10 rating = 1%, 30.496 per %), melee_haste=13.551 ± 1.160, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 146.3 ranged_attack_power points (19.56 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.09 DPS) [vendor]; Mask of the Unforgiven (13404, -28.45 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 59.7 ranged_attack_power points (7.98 DPS) | yes | Sentinel's Medallion (19538, -1.62 DPS) [rep]; Skibi's Pendant (13089, -2.47 DPS) [world_drop]; Star of Mystaria (12103, -5.61 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 119.3 ranged_attack_power points (15.96 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -2.25 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.37 DPS) [rep]; Howler's Furs (272414, -5.89 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 181.5 ranged_attack_power points (24.28 DPS) | yes | Field Marshal's Chain Breastplate (16466, -3.41 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -3.41 DPS) [vendor]; Shadowcraft Tunic (16721, -8.00 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 60.3 ranged_attack_power points (8.06 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.85 DPS) [pvp]; Slashclaw Bracers (13211, -1.01 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -2.03 DPS, sim-verified) [dungeon] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Sash of the Grand Hunt (22207, -1.03 DPS) [dungeon]; Marshal's Chain Girdle (16464, -1.90 DPS) [pvp]; Marksman's Girdle (22232, -5.60 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 208.8 ranged_attack_power points (27.93 DPS) | yes | Marshal's Chain Legguards (231577, -7.07 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.47 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.46 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 98.8 ranged_attack_power points (13.21 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.70 DPS) [dungeon]; Painweaver Band (13098, -3.81 DPS) [dungeon]; Naglering (11669, -6.57 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.37 DPS) [dungeon]; Blackstone Ring (17713, -3.61 DPS, sim-verified) [dungeon]; Innervating Band (18701, -3.65 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Second Wind (11819, -3.35 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Eye (19991, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -2.92 DPS, sim-verified) [rep]; Grand Marshal's Stave (234571, -5.36 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bloodseeker (19107, -3.05 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2287, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 983.9. Weights run: 5.3s. Verify run: 59.8s. 2287 eligible items had no known source.

Race sweep: night-elf 983.9 ± 0.5, high-order-skyborne 981.3 ± 0.5, human 978.8 ± 0.5, dwarf 972.8 ± 0.5 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.887 ± 0.054, crit=3.659 ± 0.118 per rating point (14 rating = 1%, 51.222 per %), hit=5.226 ± 0.240 per rating point (10 rating = 1%, 52.258 per %), melee_haste=not significant (4.547 ± 3.069), intellect=1.722 ± 0.034, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 198.3 ranged_attack_power points (32.83 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -1.72 DPS) [pvp]; Mask of the Unforgiven (13404, -26.98 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (+13.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nacreous Shell Necklace (22403, -1.85 DPS) [dungeon]; Sentinel's Medallion (19538, -3.13 DPS) [rep]; Star of Mystaria (12103, -13.88 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 156.2 ranged_attack_power points (25.86 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -1.18 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.03 DPS) [vendor]; Howler's Furs (272414, -10.07 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 247.3 ranged_attack_power points (40.94 DPS) | yes | Field Marshal's Chain Breastplate (16466, -7.16 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -7.16 DPS) [vendor]; Icemail Jerkin (1981, -20.35 DPS, sim-verified) [world_drop] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.29 DPS) [pvp]; Slashclaw Bracers (13211, -8.52 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+11.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -10.99 DPS, sim-verified) [dungeon] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 133.9 ranged_attack_power points (22.17 DPS) | yes | Ranger's Belt (272397, -2.73 DPS) [vendor]; Sash of the Grand Hunt (22207, -2.84 DPS, sim-verified) [dungeon]; Marshal's Chain Girdle (16464, -5.45 DPS) [pvp] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 290.7 ranged_attack_power points (48.14 DPS) | yes | Sentinel's Leather Pants (237818, -13.80 DPS) [vendor]; Marshal's Chain Legguards (231577, -14.35 DPS) [pvp]; Knight-Captain's Chain Legguards (227072, -19.17 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 136.9 ranged_attack_power points (22.66 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -7.58 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -6.81 DPS) [vendor]; Blackstone Ring (17713, -7.49 DPS) [dungeon]; Naglering (11669, -13.40 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -5.70 DPS) [dungeon]; Painweaver Band (13098, -6.53 DPS) [dungeon]; Channeler's Ring (272406, -8.19 DPS, sim-verified) [vendor] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (+15.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -8.04 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Warmonger (13052, -1.93 DPS) [world_drop]; Shadowblade (2163, -7.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Blackcrow (12651, -21.69 DPS, sim-verified) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Gloves of Holy Might; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2287, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (orc, 5320000100000000-0000000000000000-000000000000000000)

Set DPS (verified): 88.4. Weights run: 5.0s. Verify run: 6.7s. 294 eligible items had no known source.

Race sweep: windshaper-skyborne 89.3 ± 0.1, troll 88.4 ± 0.1, orc 88.4 ± 0.1, tauren 85.8 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 0.9 (1.04%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.158 ± 0.010, crit=0.667 ± 0.011 per rating point (14 rating = 1%, 9.342 per %), hit=1.082 ± 0.032 per rating point (10 rating = 1%, 10.817 per %), melee_haste=4.895 ± 0.476, intellect=1.650 ± 0.022, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.3 ranged_attack_power points (2.51 DPS) | yes | Trapper's Leather Hood (252505, -1.42 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.69 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.9 ranged_attack_power points (1.88 DPS) | yes | Erudite's Amulet (277204, -1.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 15.8 ranged_attack_power points (1.57 DPS) | yes | Magician's Mantle (12998, -0.67 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.91 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.31 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.63 DPS) [world_drop]; Bristlebark Cape (14571, -0.63 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 27.1 ranged_attack_power points (2.69 DPS) | yes | Brawler's Leather Armor (252490, -0.58 DPS, sim-verified) [crafted]; Mystic's Wrap (14369, -0.60 DPS) [world_drop]; Mystic's Robe (14371, -0.60 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.8 ranged_attack_power points (1.57 DPS) | yes | Wolf Bracers (4794, -0.31 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.63 DPS) [world_drop]; Spare Part Bindings (279875, -0.63 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 18.9 ranged_attack_power points (1.88 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.14 DPS) [crafted]; Forest Leather Gloves (3058, -0.63 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.79 DPS) | yes | Trapper's Leather Belt (252429, -0.04 DPS) [crafted]; Deviate Scale Belt (6468, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.48 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.31 DPS) [world]; Trapper's Leather Pants (252501, -0.64 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.3 ranged_attack_power points (2.51 DPS) | yes | Trapper's Leather Boots (252440, -0.31 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.63 DPS) [dungeon]; Footpads of the Fang (10411, -0.63 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.2 ranged_attack_power points (1.61 DPS) | yes | Bounty Hunter's Ring (5351, -0.67 DPS) [quest]; Loop of Sacrifice (281673, -0.79 DPS) [quest]; Volcanic Rock Ring (12053, -1.12 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.6 ranged_attack_power points (1.25 DPS) | yes | Bounty Hunter's Ring (5351, -0.31 DPS, sim-verified) [quest]; Loop of Sacrifice (281673, -0.44 DPS) [quest]; Volcanic Rock Ring (12053, -0.76 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 36.7 ranged_attack_power points (3.64 DPS) | yes | Crescent Staff (6505, -0.36 DPS, sim-verified) [quest]; Bronze Dory (250603, -0.54 DPS) [crafted]; Staff of Orgrimmar (15444, -1.10 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (17.82 DPS) | yes | Venomstrike (6469, -0.58 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.33 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -4.30 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (orc, 5320000505000100-0000000000000000-000000000000000000)

Set DPS (verified): 144.4. Weights run: 4.8s. Verify run: 13.8s. 493 eligible items had no known source.

Race sweep: orc 144.4 ± 0.1, troll 143.7 ± 0.1, windshaper-skyborne 143.6 ± 0.1, tauren 140.5 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.055 ± 0.010, crit=0.899 ± 0.014 per rating point (14 rating = 1%, 12.583 per %), hit=1.423 ± 0.044 per rating point (10 rating = 1%, 14.225 per %), melee_haste=not significant (-0.569 ± 0.646), intellect=0.556 ± 0.019, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.5 ranged_attack_power points (3.42 DPS) | yes | Brawler's Leather Hood (252504, -0.70 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -1.03 DPS) [world]; Resilient Cap (14401, -2.05 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 24.4 ranged_attack_power points (2.74 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.02 DPS) [world]; Kaleidoscope Chain (13084, -1.12 DPS) [world_drop]; Erudite's Amulet (277204, -2.05 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.6 ranged_attack_power points (3.76 DPS) | yes | Mantle of Thieves (2264, -0.52 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -1.37 DPS) [crafted]; Insignia Mantle (4721, -1.37 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.68 DPS) [dungeon]; Swiftrunner Cape (6745, -0.68 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Leather Armor (4255, -0.78 DPS) [crafted]; Trapper's Leather Tunic (252509, -0.90 DPS) [crafted]; Dusky Leather Armor (7374, -1.43 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.3 ranged_attack_power points (2.05 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.34 DPS) [world_drop]; Loamflake Bracers (15462, -0.34 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.5 ranged_attack_power points (3.08 DPS) | yes | Braced Handguards (6784, -0.68 DPS) [quest]; Insignia Gloves (6408, -1.03 DPS) [world_drop]; Pilferer's Gloves (7358, -1.07 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 30.8 ranged_attack_power points (3.45 DPS) | yes | Deftkin Belt (16659, -0.74 DPS) [quest]; Defiler's Chain Girdle (20152, -0.77 DPS) [rep]; Skulker's Leather Belt (252520, -1.25 DPS, sim-verified) [crafted] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Petrolspill Leggings (9509, -0.34 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.34 DPS) [crafted]; Ferine Leggings (6690, -1.13 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 24.4 ranged_attack_power points (2.74 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.34 DPS) [quest] |
| finger1 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 18.3 ranged_attack_power points (2.05 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.68 DPS) [dungeon]; Band of the Fist (17694, -0.68 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.3 ranged_attack_power points (2.05 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.68 DPS) [dungeon]; Band of the Fist (17694, -0.68 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Claw of the Shadowmancer (2912, -0.40 DPS, sim-verified) [world_drop]; Armor Piercer (6679, -1.25 DPS) [dungeon]; Kam's Walking Stick (2280, -1.41 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 250.7 ranged_attack_power points (28.09 DPS) | yes | Ironweaver (13137, -1.37 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.54 DPS) [world_drop]; Chesterfall Musket (7729, -2.18 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Legionnaire's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 493, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (orc, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 226.1. Weights run: 5.3s. Verify run: 17.7s. 817 eligible items had no known source.

Race sweep: orc 226.1 ± 0.1, windshaper-skyborne 225.5 ± 0.1, troll 224.3 ± 0.1, tauren 220.6 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.869 ± 0.027, crit=0.873 ± 0.042 per rating point (14 rating = 1%, 12.224 per %), hit=1.865 ± 0.077 per rating point (10 rating = 1%, 18.649 per %), melee_haste=8.174 ± 0.789, intellect=2.305 ± 0.031, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpseshroud (10574, +0.00 DPS) [dungeon]; Thinking Cap (2624, -0.52 DPS) [world]; Miner's Hat of the Deep (9429, -0.52 DPS) [dungeon] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 31.6 ranged_attack_power points (3.28 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Kaleidoscope Chain (13084, -1.13 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.36 DPS) [dungeon] |
| shoulder | Nightscape Shoulders (8192) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Darkspear Raider's Cloak (272077, +0.00 DPS) [vendor] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 87.5 ranged_attack_power points (9.09 DPS) | yes | Garb of Florid Feathers (273026, -2.82 DPS) [dungeon]; Tough Scorpid Breastplate (8203, -2.94 DPS) [crafted]; Quillward Harness (10583, -3.43 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.9 ranged_attack_power points (3.52 DPS) | yes | Green Leather Bracers (4259, -0.54 DPS) [crafted]; Warden's Wristbands (14600, -0.71 DPS) [world_drop]; Windchaser Cuffs (14429, -0.77 DPS) [world_drop] |
| hands | Dragonscale Gauntlets (8347) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Aurora Gloves (4042, +0.00 DPS) [world_drop]; Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stalker's Leather Gloves (252526, +0.00 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.7 ranged_attack_power points (4.12 DPS) | yes | Deathmage Sash (10771, -0.53 DPS) [dungeon]; Defiler's Lizardhide Girdle (20173, -0.53 DPS) [rep]; Defiler's Mail Girdle (20198, -0.53 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 60.3 ranged_attack_power points (6.26 DPS) | yes | Ferine Leggings (6690, -0.69 DPS) [dungeon]; Triprunner Dungarees (9624, -0.89 DPS) [quest]; Trapper's Leather Legguards (252517, -1.90 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swampwalker Boots (2276, -1.08 DPS) [world_drop]; Worn Running Boots (9398, -1.08 DPS) [dungeon]; Briar Tredders (10582, -1.31 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.7 ranged_attack_power points (2.98 DPS) | yes | Falcon's Hook (7552, -0.30 DPS) [dungeon]; Legionnaire's Band (19512, -0.60 DPS) [rep]; Field Researcher's Loop (281634, -0.89 DPS) [quest] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Falcon's Hook (7552, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.30 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -2.77 DPS) [world]; Steel Spear (250605, -2.78 DPS) [crafted]; Ruthless Shiv (7758, -3.74 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -2.44 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.26 DPS) [dungeon]; Swiftwind (13038, -3.37 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Dragonscale Gauntlets; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 817, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (orc, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 327.5. Weights run: 5.4s. Verify run: 24.6s. 1053 eligible items had no known source.

Race sweep: orc 327.5 ± 0.2, windshaper-skyborne 325.6 ± 0.2, troll 324.1 ± 0.2, tauren 319.6 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.107 ± 0.030, crit=1.777 ± 0.052 per rating point (14 rating = 1%, 24.877 per %), hit=2.450 ± 0.096 per rating point (10 rating = 1%, 24.501 per %), melee_haste=11.764 ± 1.106, intellect=2.780 ± 0.041, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Tough Scorpid Helm (8208, -0.19 DPS) [crafted]; Chief Architect's Monocle (11839, -1.49 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.4 ranged_attack_power points (4.24 DPS) | yes | Scout's Medallion (19535, -0.33 DPS) [rep]; Horizon Choker (13085, -5.10 DPS, sim-verified) [world_drop] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 58.9 ranged_attack_power points (6.17 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Imperial Red Mantle (8250, -0.50 DPS) [world_drop]; Ironfeather Shoulders (15067, -3.98 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.33 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.81 DPS) [vendor] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blazewind Breastplate (11193, -2.90 DPS) [quest]; Garb of Florid Feathers (273026, -3.08 DPS) [dungeon]; Stone Guard's Chain Armor (220827, -3.49 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.6 ranged_attack_power points (4.89 DPS) | yes | Forgotten Wraps (9433, -0.41 DPS) [world_drop]; Bracers of the Stone Princess (17714, -0.49 DPS) [dungeon]; Stalker's Mail Bracers (252570, -2.45 DPS, sim-verified) [crafted] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hibernal Gloves (8110, -1.01 DPS) [world_drop]; Rockgrip Gauntlets (17736, -1.04 DPS) [dungeon]; Gloves of Holy Might (867, -1.28 DPS, sim-verified) [world_drop] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 71.3 ranged_attack_power points (7.48 DPS) | yes | Sagebrush Girdle (17778, -1.33 DPS, sim-verified) [quest]; Dawnspire Cord (12466, -1.94 DPS) [dungeon]; Defiler's Lizardhide Girdle (20174, -1.95 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 108.9 ranged_attack_power points (11.42 DPS) | yes | Imperial Red Pants (8251, -2.16 DPS) [world_drop]; Tough Scorpid Leggings (8206, -2.96 DPS) [crafted]; Stone Guard's Chain Legplates (220833, -3.74 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 75.7 ranged_attack_power points (7.94 DPS) | yes | Albino Crocscale Boots (17728, -1.42 DPS) [dungeon]; Tough Scorpid Boots (8209, -1.99 DPS) [crafted]; Stalker's Mail Sabatons (252576, -2.48 DPS, sim-verified) [crafted] |
| finger1 | Mindseye Circle (10634) | Sunken Temple: Atal'ai Warrior [dungeon] | 45.6 ranged_attack_power points (4.78 DPS) | yes | Masons Fraternity Ring (9533, -0.22 DPS) [quest]; Ring of the Underwood (2951, -1.53 DPS) [world_drop]; Falcon's Hook (7552, -1.85 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.5 ranged_attack_power points (4.67 DPS) | yes | Masons Fraternity Ring (9533, -0.11 DPS) [quest]; Ring of the Underwood (2951, -1.41 DPS) [world_drop]; Falcon's Hook (7552, -1.73 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Devilsaur Tooth (19992, -8.29 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, -0.62 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -0.95 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Staff of Jordan (873, -3.21 DPS, sim-verified) [world_drop] |
| off_hand | Ebon Shiv (7947) | Blacksmithing [crafted] | 28.0 ranged_attack_power points (2.93 DPS) | yes | White Bone Shredder (11863, -0.65 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Gryphonwing Long Bow (13022, -0.16 DPS) [dungeon]; Hurricane (2824, -1.73 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -1.92 DPS) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Mindseye Circle; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Shadowblade; off_hand: Ebon Shiv; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1053, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (orc, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 484.8. Weights run: 5.5s. Verify run: 59.5s. 2286 eligible items had no known source.

Race sweep: orc 484.8 ± 0.2, windshaper-skyborne 482.1 ± 0.2, troll 480.1 ± 0.2, tauren 473.4 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.171 ± 0.025, crit=2.406 ± 0.053 per rating point (14 rating = 1%, 33.677 per %), hit=3.050 ± 0.109 per rating point (10 rating = 1%, 30.496 per %), melee_haste=13.551 ± 1.160, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 146.3 ranged_attack_power points (19.56 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -22.58 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 59.7 ranged_attack_power points (7.98 DPS) | yes | Star of Mystaria (12103, -1.55 DPS, sim-verified) [dungeon]; Scout's Medallion (19534, -1.62 DPS) [rep]; Skibi's Pendant (13089, -2.47 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 119.3 ranged_attack_power points (15.96 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -2.25 DPS) [dungeon]; Deathguard's Cloak (20068, -2.37 DPS) [rep]; Howler's Furs (272414, -2.71 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 181.5 ranged_attack_power points (24.28 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.41 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.41 DPS) [vendor]; Shadowcraft Tunic (16721, -9.22 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 60.3 ranged_attack_power points (8.06 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.85 DPS) [pvp]; Slashclaw Bracers (13211, -1.01 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -4.84 DPS, sim-verified) [dungeon] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marksman's Girdle (22232, -0.72 DPS, sim-verified) [dungeon]; Sash of the Grand Hunt (22207, -1.03 DPS) [dungeon]; General's Chain Girdle (16572, -1.90 DPS) [pvp] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 175.2 ranged_attack_power points (23.43 DPS) | yes | General's Chain Legguards (231574, -2.56 DPS) [pvp]; General's Chain Legplates (231567, -4.96 DPS) [vendor]; Plaguehound Leggings (18736, -7.05 DPS, sim-verified) [dungeon] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 98.8 ranged_attack_power points (13.21 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -4.01 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.70 DPS) [dungeon]; Painweaver Band (13098, -3.81 DPS) [dungeon]; Naglering (11669, -12.46 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.37 DPS) [dungeon]; Blackstone Ring (17713, -3.40 DPS, sim-verified) [dungeon]; Innervating Band (18701, -3.65 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Second Wind (11819, -3.23 DPS, sim-verified) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -3.54 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -5.36 DPS) [pvp]; Shadowblade (2163, -5.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Bloodseeker (19107, -4.72 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2286, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (orc, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 975.9. Weights run: 5.3s. Verify run: 60.3s. 2286 eligible items had no known source.

Race sweep: orc 975.9 ± 0.5, windshaper-skyborne 972.2 ± 0.5, troll 968.4 ± 0.5, tauren 955.5 ± 0.5 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.887 ± 0.054, crit=3.659 ± 0.118 per rating point (14 rating = 1%, 51.222 per %), hit=5.226 ± 0.240 per rating point (10 rating = 1%, 52.258 per %), melee_haste=not significant (4.547 ± 3.069), intellect=1.722 ± 0.034, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 198.3 ranged_attack_power points (32.83 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -25.91 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (+13.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Nacreous Shell Necklace (22403, -1.85 DPS) [dungeon]; Scout's Medallion (19534, -3.13 DPS) [rep]; Star of Mystaria (12103, -13.56 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 156.2 ranged_attack_power points (25.86 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.35 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shroud of Arcane Mastery (22330, -1.18 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.03 DPS) [vendor]; Howler's Furs (272414, -9.30 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 247.3 ranged_attack_power points (40.94 DPS) | yes | Warlord's Chain Chestpiece (16565, -7.16 DPS) [vendor]; Warlord's Chain Hauberk (231573, -7.16 DPS) [vendor]; Icemail Jerkin (1981, -17.99 DPS, sim-verified) [world_drop] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.29 DPS) [pvp]; Slashclaw Bracers (13211, -8.66 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -10.30 DPS, sim-verified) [dungeon] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 133.9 ranged_attack_power points (22.17 DPS) | yes | Sash of the Grand Hunt (22207, -2.65 DPS, sim-verified) [dungeon]; Ranger's Belt (272397, -2.73 DPS) [vendor]; General's Chain Girdle (16572, -5.45 DPS) [pvp] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 239.5 ranged_attack_power points (39.66 DPS) | yes | General's Chain Legguards (231574, -5.87 DPS) [pvp]; Legionnaire's Chain Legguards (227073, -10.69 DPS) [pvp]; Plaguehound Leggings (18736, -11.62 DPS, sim-verified) [dungeon] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 136.9 ranged_attack_power points (22.66 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -7.05 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -6.81 DPS) [vendor]; Blackstone Ring (17713, -7.49 DPS) [dungeon]; Naglering (11669, -11.49 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -5.70 DPS) [dungeon]; Painweaver Band (13098, -6.53 DPS) [dungeon]; Channeler's Ring (272406, -6.56 DPS, sim-verified) [vendor] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -8.13 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Warmonger (13052, -1.93 DPS) [world_drop]; Shadowblade (2163, -6.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Blackcrow (12651, -16.76 DPS, sim-verified) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Gloves of Holy Might; waist: Marksman's Girdle; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2286, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

