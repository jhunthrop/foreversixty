# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 88.8. Weights run: 8.5s. Verify run: 8.6s. 268 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.205 ± 0.011, crit=0.707 ± 0.012 per rating point (14 rating = 1%, 9.900 per %), hit=1.110 ± 0.034 per rating point (10 rating = 1%, 11.101 per %), melee_haste=5.398 ± 0.540, intellect=1.572 ± 0.028, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.6 ranged_attack_power points (2.32 DPS) | yes | Trapper's Leather Hood (252505, -1.52 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.61 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 19.2 ranged_attack_power points (1.74 DPS) | yes | Erudite's Amulet (277204, -1.14 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 16.0 ranged_attack_power points (1.45 DPS) | yes | Magician's Mantle (12998, -0.79 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.88 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.28 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.58 DPS) [world_drop]; Bristlebark Cape (14571, -0.58 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 35.3 ranged_attack_power points (3.19 DPS) | yes | Trapper's Leather Armor (252491, -1.08 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -1.16 DPS) [crafted]; Mystic's Wrap (14369, -1.33 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.0 ranged_attack_power points (1.45 DPS) | yes | Wolf Bracers (4794, -0.28 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.58 DPS) [world_drop]; Bravo's Armbands (270015, -0.58 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.2 ranged_attack_power points (1.74 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.15 DPS) [crafted]; Riding Gloves (1304, -0.44 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.63 DPS) | yes | Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Trapper's Leather Belt (252429, -0.29 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.47 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.29 DPS) [world]; Trapper's Leather Pants (252501, -0.57 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.6 ranged_attack_power points (2.32 DPS) | yes | Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon]; Trapper's Leather Boots (252440, -0.69 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.8 ranged_attack_power points (1.44 DPS) | yes | The 1 Ring (8350, -1.00 DPS) [world]; Volcanic Rock Ring (12053, -1.01 DPS) [world_drop]; Minor Channeling Ring (1449, -1.15 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.8 ranged_attack_power points (1.16 DPS) | yes | The 1 Ring (8350, -0.67 DPS, sim-verified) [world]; Volcanic Rock Ring (12053, -0.73 DPS) [world_drop]; Minor Channeling Ring (1449, -0.88 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 36.7 ranged_attack_power points (3.33 DPS) | yes | Bronze Dory (250603, -0.80 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.58 DPS) [world]; Piercing Axe (6094, -1.58 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (16.26 DPS) | yes | Venomstrike (6469, -0.33 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.20 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.93 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 139.4. Weights run: 7.7s. Verify run: 18.8s. 463 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.044 ± 0.010, crit=0.883 ± 0.014 per rating point (14 rating = 1%, 12.367 per %), hit=1.442 ± 0.041 per rating point (10 rating = 1%, 14.421 per %), melee_haste=not significant (-0.877 ± 0.594), intellect=0.429 ± 0.016, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.4 ranged_attack_power points (3.37 DPS) | yes | Brawler's Leather Hood (252504, -0.63 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -1.01 DPS) [world]; Resilient Cap (14401, -2.18 DPS) [world_drop] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 24.4 ranged_attack_power points (2.70 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.08 DPS) [world]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Erudite's Amulet (277204, -2.02 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.5 ranged_attack_power points (3.71 DPS) | yes | Mantle of Thieves (2264, -0.34 DPS) [dungeon]; Dark Leather Shoulders (4252, -1.35 DPS) [crafted]; Insignia Mantle (4721, -1.35 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.67 DPS) [dungeon]; Cape of the Brotherhood (5193, -1.01 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Tunic of Westfall (2041, -0.04 DPS) [quest]; Green Leather Armor (4255, -0.67 DPS) [crafted]; Dusky Leather Armor (7374, -1.66 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.3 ranged_attack_power points (2.02 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.34 DPS) [world]; Forest Leather Bracers (3202, -0.34 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.4 ranged_attack_power points (3.04 DPS) | yes | Pilferer's Gloves (7358, -0.85 DPS, sim-verified) [crafted]; Insignia Gloves (6408, -1.01 DPS) [world_drop]; Gloves of the Fang (10413, -1.01 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 30.0 ranged_attack_power points (3.32 DPS) | yes | Highlander's Chain Girdle (20090, -0.66 DPS) [rep]; Highlander's Leather Girdle (20117, -0.66 DPS) [rep]; Skulker's Leather Belt (252520, -1.36 DPS, sim-verified) [crafted] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Petrolspill Leggings (9509, -0.23 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.34 DPS) [crafted]; Ferine Leggings (6690, -1.27 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 24.4 ranged_attack_power points (2.70 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, -0.67 DPS) [dungeon]; Footpads of the Fang (10411, -0.67 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 18.3 ranged_attack_power points (2.02 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.67 DPS) [dungeon]; Seal of Wrynn (2933, -0.82 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.3 ranged_attack_power points (2.02 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.67 DPS) [dungeon]; Seal of Wrynn (2933, -0.82 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Claw of the Shadowmancer (2912, +0.00 DPS) [world_drop]; Armor Piercer (6679, -1.26 DPS) [dungeon]; Kam's Walking Stick (2280, -1.46 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 250.7 ranged_attack_power points (27.77 DPS) | yes | Ironweaver (13137, -0.66 DPS, sim-verified) [world_drop]; Crystalpine Stinger (13037, -1.58 DPS) [world_drop]; Chesterfall Musket (7729, -2.15 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 463, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 227.1. Weights run: 8.3s. Verify run: 22.5s. 718 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.866 ± 0.029, crit=0.894 ± 0.045 per rating point (14 rating = 1%, 12.510 per %), hit=1.796 ± 0.081 per rating point (10 rating = 1%, 17.956 per %), melee_haste=6.394 ± 0.858, intellect=2.278 ± 0.034, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Corpseshroud (10574, +0.00 DPS) [dungeon]; Thinking Cap (2624, -0.54 DPS) [world] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 31.5 ranged_attack_power points (3.17 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; Prodigious Shadowshard Pendant (17773, -0.88 DPS) [quest]; Kaleidoscope Chain (13084, -1.10 DPS) [world_drop] |
| shoulder | Nightscape Shoulders (8192) | Leatherworking [crafted] | sim-verified (227.1 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271922, -0.58 DPS) [rep]; Imperial Cloak (6432, -0.80 DPS) [dungeon]; Parachute Cloak (10518, -0.80 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 86.7 ranged_attack_power points (8.73 DPS) | yes | Garb of Florid Feathers (273026, -2.69 DPS) [dungeon]; Tough Scorpid Breastplate (8203, -2.80 DPS) [crafted]; Quillward Harness (10583, -3.25 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.7 ranged_attack_power points (3.40 DPS) | yes | Green Leather Bracers (4259, -0.52 DPS) [crafted]; Warden's Wristbands (14600, -0.70 DPS) [world_drop]; Windchaser Cuffs (14429, -0.75 DPS) [world_drop] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.46 DPS) [crafted]; Aurora Gloves (4042, -0.81 DPS) [world_drop]; Gauntlets of Divinity (7724, -1.73 DPS) [dungeon] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.5 ranged_attack_power points (3.97 DPS) | yes | Deathmage Sash (10771, -0.53 DPS) [dungeon]; Highlander's Lizardhide Girdle (20104, -0.53 DPS) [rep]; Highlander's Mail Girdle (20119, -0.53 DPS) [vendor] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 60.2 ranged_attack_power points (6.06 DPS) | yes | Ferine Leggings (6690, -0.69 DPS) [dungeon]; Triprunner Dungarees (9624, -0.87 DPS) [quest]; Trapper's Leather Legguards (252517, -1.86 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | sim-verified (227.1 DPS) | yes | Swampwalker Boots (2276, -1.03 DPS) [world_drop]; Worn Running Boots (9398, -1.03 DPS) [dungeon]; Briar Tredders (10582, -1.27 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.7 ranged_attack_power points (2.88 DPS) | yes | Protector's Band (19515, -0.58 DPS) [rep]; Disengagement Ring (276202, -0.58 DPS) [vendor]; Field Researcher's Loop (281634, -0.87 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 25.8 ranged_attack_power points (2.60 DPS) | yes | Protector's Band (19515, -0.29 DPS) [rep]; Disengagement Ring (276202, -0.29 DPS) [vendor]; Field Researcher's Loop (281634, -0.58 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Steel Spear (250605, -2.54 DPS) [crafted]; Felbough Arm (273037, -2.61 DPS) [world]; Manslayer (10570, -3.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -2.37 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.16 DPS) [dungeon]; Swiftwind (13038, -3.27 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 324.6. Weights run: 8.4s. Verify run: 35.5s. 900 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.110 ± 0.032, crit=1.759 ± 0.057 per rating point (14 rating = 1%, 24.630 per %), hit=2.433 ± 0.108 per rating point (10 rating = 1%, 24.334 per %), melee_haste=7.554 ± 1.160, intellect=2.911 ± 0.044, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.44 DPS) [dungeon]; Helm of Exile (11124, -1.17 DPS, sim-verified) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Medallion (19539, -0.32 DPS) [rep]; Horizon Choker (13085, -5.13 DPS, sim-verified) [world_drop] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 60.2 ranged_attack_power points (6.10 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Imperial Red Mantle (8250, -0.42 DPS) [world_drop]; Ironfeather Shoulders (15067, -4.51 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.32 DPS) [dungeon]; Expeditionary's Cape (271921, -0.60 DPS) [rep] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -2.65 DPS) [world_drop]; Garb of Florid Feathers (273026, -2.77 DPS) [dungeon]; Knight's Chain Armor (220828, -3.40 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.7 ranged_attack_power points (4.73 DPS) | yes | Forgotten Wraps (9433, -0.24 DPS) [world_drop]; Aristocratic Cuffs (12546, -0.30 DPS) [dungeon]; Stalker's Mail Bracers (252570, -2.60 DPS, sim-verified) [crafted] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hibernal Gloves (8110, -0.97 DPS) [world_drop]; Gloves of Holy Might (867, -1.15 DPS, sim-verified) [world_drop]; Tough Scorpid Gloves (8204, -1.26 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 81.6 ranged_attack_power points (8.27 DPS) | yes | Stalker's Mail Belt (252588, -0.90 DPS) [crafted]; Ban'thok Sash (11662, -2.56 DPS) [dungeon]; Dawnspire Cord (12466, -2.66 DPS) [dungeon] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 97.3 ranged_attack_power points (9.86 DPS) | yes | Tough Scorpid Leggings (8206, -1.55 DPS) [crafted]; Searingscale Leggings (11749, -1.93 DPS) [dungeon]; Imperial Red Pants (8251, -3.50 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 76.6 ranged_attack_power points (7.76 DPS) | yes | Stalker's Mail Sabatons (252576, -1.01 DPS) [crafted]; Albino Crocscale Boots (17728, -1.45 DPS) [dungeon]; Rancor Boots (11865, -3.04 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.3 ranged_attack_power points (4.49 DPS) | yes | Mindseye Circle (10634, -0.95 DPS) [dungeon]; Ring of the Underwood (2951, -1.34 DPS) [world_drop]; Brainlash (6440, -6.33 DPS, sim-verified) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -0.87 DPS) [dungeon]; Ring of the Underwood (2951, -1.26 DPS) [world_drop]; Brainlash (6440, -5.48 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -5.06 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Claw of the Shadowmancer (2912, +0.00 DPS) [world_drop]; Glowing Brightwood Staff (812, -1.04 DPS) [world_drop]; Warmonger (13052, -2.20 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Megashot Rifle (17717, -3.34 DPS) [dungeon]; Dark Iron Rifle (16004, -3.94 DPS) [crafted]; Gryphonwing Long Bow (13022, -5.14 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Thunderbrew's Boot Flask; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 900, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 477.9. Weights run: 8.1s. Verify run: 90.3s. 2028 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.218 ± 0.030, crit=2.450 ± 0.059 per rating point (14 rating = 1%, 34.301 per %), hit=2.993 ± 0.119 per rating point (10 rating = 1%, 29.930 per %), melee_haste=not significant (3.594 ± 1.495), intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 147.9 ranged_attack_power points (19.26 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.09 DPS) [vendor]; Outlaw's Collar (279253, -21.00 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.84 DPS) [world_drop]; Talisman of Evasion (13177, -0.84 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 121.2 ranged_attack_power points (15.79 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -1.77 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.37 DPS) [rep]; Howler's Furs (272414, -5.97 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 183.3 ranged_attack_power points (23.88 DPS) | yes | Field Marshal's Chain Breastplate (16466, -3.27 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -3.27 DPS) [vendor]; Savage Gladiator Chain (11726, -11.99 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 61.1 ranged_attack_power points (7.97 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.84 DPS) [pvp]; Slashclaw Bracers (13211, -1.13 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -0.84 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Girdle (16464, -1.88 DPS) [pvp]; Marksman's Girdle (22232, -2.72 DPS, sim-verified) [dungeon]; Highlander's Chain Girdle (20043, -3.37 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 211.2 ranged_attack_power points (27.51 DPS) | yes | Marshal's Chain Legguards (231577, -6.90 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.25 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.33 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 94.3 ranged_attack_power points (12.28 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -2.66 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.64 DPS) [dungeon]; Blackstone Ring (17713, -3.69 DPS) [dungeon]; Naglering (11669, -11.08 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.26 DPS) [dungeon]; Innervating Band (18701, -3.48 DPS) [dungeon]; Painweaver Band (13098, -4.40 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (477.9 DPS) | yes | Hand of Justice (11815, -0.90 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Heart of Wyrmthalak (22321, -4.30 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -5.08 DPS) [pvp]; Electrified Dagger (19100, -8.50 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -3.21 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 940.3. Weights run: 8.5s. Verify run: 85.4s. 2028 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.921 ± 0.063, crit=3.596 ± 0.123 per rating point (14 rating = 1%, 50.344 per %), hit=5.148 ± 0.252 per rating point (10 rating = 1%, 51.475 per %), melee_haste=not significant (3.603 ± 3.143), intellect=1.767 ± 0.036, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 199.8 ranged_attack_power points (32.39 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -2.05 DPS) [pvp]; Outlaw's Collar (279253, -5.30 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 156.2 ranged_attack_power points (25.32 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.11 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -1.39 DPS) [dungeon]; Shadow Prowler's Cloak (22269, -2.08 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 246.9 ranged_attack_power points (40.02 DPS) | yes | Field Marshal's Chain Breastplate (16466, -6.81 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -6.81 DPS) [vendor]; Savage Gladiator Chain (11726, -21.71 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.27 DPS) [pvp] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Gloves of Holy Might (867, -7.61 DPS, sim-verified) [world_drop] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Marshal's Chain Girdle (16464, -2.67 DPS) [pvp]; Clutch of Andros (13956, -5.48 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 289.4 ranged_attack_power points (46.91 DPS) | yes | Sentinel's Leather Pants (237818, -13.43 DPS) [vendor]; Marshal's Chain Legguards (231577, -13.70 DPS) [pvp]; Marshal's Chain Legplates (231558, -18.67 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -6.42 DPS) [vendor]; Blackstone Ring (17713, -7.19 DPS) [dungeon]; Naglering (11669, -21.84 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -5.66 DPS) [dungeon]; Painweaver Band (13098, -6.49 DPS) [dungeon]; Channeler's Ring (272406, -7.59 DPS, sim-verified) [vendor] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Devilsaur Eye (19991, -6.59 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (940.3 DPS) | yes | Hand of Justice (11815, -4.27 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Warmonger (13052, -2.23 DPS) [world_drop]; Electrified Dagger (19100, -14.95 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-verified (+13.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -13.94 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Mourning Handguards; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (orc, 5420000000000000-0000000000000000-000000000000000000)

Set DPS (verified): 85.7. Weights run: 8.5s. Verify run: 8.6s. 264 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.205 ± 0.011, crit=0.707 ± 0.012 per rating point (14 rating = 1%, 9.900 per %), hit=1.110 ± 0.034 per rating point (10 rating = 1%, 11.101 per %), melee_haste=5.398 ± 0.540, intellect=1.572 ± 0.028, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.6 ranged_attack_power points (2.32 DPS) | yes | Trapper's Leather Hood (252505, -1.57 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.61 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 19.2 ranged_attack_power points (1.74 DPS) | yes | Erudite's Amulet (277204, -1.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 16.0 ranged_attack_power points (1.45 DPS) | yes | Magician's Mantle (12998, -0.38 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.88 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.29 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.58 DPS) [world_drop]; Bristlebark Cape (14571, -0.58 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 27.2 ranged_attack_power points (2.46 DPS) | yes | Mystic's Wrap (14369, -0.59 DPS) [world_drop]; Mystic's Robe (14371, -0.59 DPS) [world_drop]; Brawler's Leather Armor (252490, -0.75 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.0 ranged_attack_power points (1.45 DPS) | yes | Wolf Bracers (4794, -0.29 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.58 DPS) [world_drop]; Spare Part Bindings (279875, -0.58 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.2 ranged_attack_power points (1.74 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.15 DPS) [crafted]; Forest Leather Gloves (3058, -0.58 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.63 DPS) | yes | Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Trapper's Leather Belt (252429, -0.28 DPS, sim-verified) [crafted]; Dark Leather Belt (4249, -0.47 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.29 DPS) [world]; Trapper's Leather Pants (252501, -0.57 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.6 ranged_attack_power points (2.32 DPS) | yes | Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon]; Trapper's Leather Boots (252440, -0.68 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 15.8 ranged_attack_power points (1.44 DPS) | yes | Bounty Hunter's Ring (5351, -0.56 DPS) [quest]; Loop of Sacrifice (281673, -0.72 DPS) [quest]; The 1 Ring (8350, -1.00 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.8 ranged_attack_power points (1.16 DPS) | yes | Bounty Hunter's Ring (5351, -0.29 DPS, sim-verified) [quest]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; The 1 Ring (8350, -0.73 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 36.7 ranged_attack_power points (3.33 DPS) | yes | Bronze Dory (250603, -0.46 DPS) [crafted]; Crescent Staff (6505, -0.54 DPS, sim-verified) [quest]; Staff of Orgrimmar (15444, -0.73 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (16.26 DPS) | yes | Venomstrike (6469, -0.94 DPS) [dungeon]; Lil Timmy's Peashooter (13136, -1.20 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.93 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (orc, 5420001504000000-0000000000000000-000000000000000000)

Set DPS (verified): 137.7. Weights run: 7.7s. Verify run: 21.2s. 462 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.044 ± 0.010, crit=0.883 ± 0.014 per rating point (14 rating = 1%, 12.367 per %), hit=1.442 ± 0.041 per rating point (10 rating = 1%, 14.421 per %), melee_haste=not significant (-0.877 ± 0.594), intellect=0.429 ± 0.016, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 30.4 ranged_attack_power points (3.37 DPS) | yes | Brawler's Leather Hood (252504, -0.67 DPS) [crafted]; Humbert's Helm (4724, -1.01 DPS) [world]; Resilient Cap (14401, -2.18 DPS) [world_drop] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 24.4 ranged_attack_power points (2.70 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.08 DPS) [world]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Erudite's Amulet (277204, -2.02 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.5 ranged_attack_power points (3.71 DPS) | yes | Mantle of Thieves (2264, -0.34 DPS) [dungeon]; Dark Leather Shoulders (4252, -1.35 DPS) [crafted]; Insignia Mantle (4721, -1.35 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.67 DPS) [dungeon]; Swiftrunner Cape (6745, -0.67 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dusky Leather Armor (7374, +0.00 DPS) [crafted]; Green Leather Armor (4255, -0.67 DPS) [crafted]; Panther Armor (6670, -0.71 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.3 ranged_attack_power points (2.02 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.34 DPS) [world_drop]; Loamflake Bracers (15462, -0.34 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.4 ranged_attack_power points (3.04 DPS) | yes | Pilferer's Gloves (7358, -0.34 DPS) [crafted]; Braced Handguards (6784, -0.67 DPS) [quest]; Insignia Gloves (6408, -1.01 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 30.0 ranged_attack_power points (3.32 DPS) | yes | Skulker's Leather Belt (252520, -0.29 DPS) [crafted]; Deftkin Belt (16659, -0.64 DPS) [quest]; Defiler's Chain Girdle (20152, -0.66 DPS) [rep] |
| legs | Troll's Bane Leggings (13114) | World drop [world_drop] | sim-verified (137.7 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Dusky Leather Leggings (7373, -0.34 DPS) [crafted]; Ferine Leggings (6690, -1.27 DPS) [dungeon] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 24.4 ranged_attack_power points (2.70 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trailblazer Boots (10653, -0.34 DPS) [quest] |
| finger1 | Ring of Precision (1491) (or Legionnaire's Band (19513)) | Blackfathom Deeps [dungeon] | 18.3 ranged_attack_power points (2.02 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.67 DPS) [dungeon]; Band of the Fist (17694, -0.67 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.3 ranged_attack_power points (2.02 DPS) | yes | Monkey Ring (6748, -0.34 DPS) [quest]; Defias Renegade Ring (1076, -0.67 DPS) [dungeon]; Band of the Fist (17694, -0.67 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Armor Piercer (6679, -1.26 DPS) [dungeon]; Kam's Walking Stick (2280, -1.46 DPS) [dungeon]; Bronze Dory (250603, -1.61 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 250.7 ranged_attack_power points (27.77 DPS) | yes | Ironweaver (13137, -0.93 DPS) [world_drop]; Crystalpine Stinger (13037, -1.58 DPS) [world_drop]; Chesterfall Musket (7729, -2.15 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Troll's Bane Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Legionnaire's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (orc, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 226.0. Weights run: 8.3s. Verify run: 22.4s. 708 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=2.866 ± 0.029, crit=0.894 ± 0.045 per rating point (14 rating = 1%, 12.510 per %), hit=1.796 ± 0.081 per rating point (10 rating = 1%, 17.956 per %), melee_haste=6.394 ± 0.858, intellect=2.278 ± 0.034, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Miner's Hat of the Deep (9429, +0.00 DPS) [dungeon]; Corpseshroud (10574, +0.00 DPS) [dungeon]; Thinking Cap (2624, -0.54 DPS) [world] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 31.5 ranged_attack_power points (3.17 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; Prodigious Shadowshard Pendant (17773, -0.88 DPS) [quest]; Kaleidoscope Chain (13084, -1.10 DPS) [world_drop] |
| shoulder | Nightscape Shoulders (8192) | Leatherworking [crafted] | sim-verified (226.0 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Darkspear Raider's Cloak (272077, +0.00 DPS) [vendor] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 86.7 ranged_attack_power points (8.73 DPS) | yes | Garb of Florid Feathers (273026, -2.69 DPS) [dungeon]; Tough Scorpid Breastplate (8203, -2.80 DPS) [crafted]; Quillward Harness (10583, -3.25 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.7 ranged_attack_power points (3.40 DPS) | yes | Green Leather Bracers (4259, -0.52 DPS) [crafted]; Warden's Wristbands (14600, -0.70 DPS) [world_drop]; Windchaser Cuffs (14429, -0.75 DPS) [world_drop] |
| hands | Dragonscale Gauntlets (8347) | Leatherworking [crafted] | sim-verified (226.0 DPS) | yes | Aurora Gloves (4042, +0.00 DPS) [world_drop]; Tough Scorpid Gloves (8204, +0.00 DPS) [crafted]; Stalker's Leather Gloves (252526, +0.00 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.5 ranged_attack_power points (3.97 DPS) | yes | Deathmage Sash (10771, -0.53 DPS) [dungeon]; Defiler's Lizardhide Girdle (20173, -0.53 DPS) [rep]; Defiler's Mail Girdle (20198, -0.53 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 60.2 ranged_attack_power points (6.06 DPS) | yes | Ferine Leggings (6690, -0.69 DPS) [dungeon]; Triprunner Dungarees (9624, -0.87 DPS) [quest]; Trapper's Leather Legguards (252517, -1.86 DPS) [crafted] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | sim-verified (226.0 DPS) | yes | Swampwalker Boots (2276, -1.03 DPS) [world_drop]; Worn Running Boots (9398, -1.03 DPS) [dungeon]; Briar Tredders (10582, -1.27 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 28.7 ranged_attack_power points (2.88 DPS) | yes | Falcon's Hook (7552, -0.29 DPS) [dungeon]; Legionnaire's Band (19512, -0.58 DPS) [rep]; Field Researcher's Loop (281634, -0.87 DPS) [quest] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (226.0 DPS) | yes | Falcon's Hook (7552, +0.00 DPS) [dungeon]; Legionnaire's Band (19512, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.29 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Steel Spear (250605, -2.54 DPS) [crafted]; Felbough Arm (273037, -2.61 DPS) [world]; Manslayer (10570, -3.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Silencer (13138, -2.37 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.16 DPS) [dungeon]; Swiftwind (13038, -3.27 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Dragonscale Gauntlets; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 708, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (orc, 5420001505001251-0052300000000000-000000000000000000)

Set DPS (verified): 326.3. Weights run: 8.4s. Verify run: 35.2s. 891 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.110 ± 0.032, crit=1.759 ± 0.057 per rating point (14 rating = 1%, 24.630 per %), hit=2.433 ± 0.108 per rating point (10 rating = 1%, 24.334 per %), melee_haste=7.554 ± 1.160, intellect=2.911 ± 0.044, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Eye of Theradras (17715, -0.44 DPS) [dungeon]; Helm of Exile (11124, -1.23 DPS, sim-verified) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Scout's Medallion (19535, -0.32 DPS) [rep]; Horizon Choker (13085, -5.62 DPS, sim-verified) [world_drop] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | 60.2 ranged_attack_power points (6.10 DPS) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Imperial Red Mantle (8250, -0.42 DPS) [world_drop]; Ironfeather Shoulders (15067, -3.89 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.32 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.60 DPS) [vendor] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Robes of Insight (940, -2.65 DPS) [world_drop]; Garb of Florid Feathers (273026, -2.77 DPS) [dungeon]; Stone Guard's Chain Armor (220827, -4.18 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 46.7 ranged_attack_power points (4.73 DPS) | yes | Forgotten Wraps (9433, -0.24 DPS) [world_drop]; Aristocratic Cuffs (12546, -0.30 DPS) [dungeon]; Stalker's Mail Bracers (252570, -2.30 DPS, sim-verified) [crafted] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, -0.75 DPS, sim-verified) [world_drop]; Hibernal Gloves (8110, -0.97 DPS) [world_drop]; Tough Scorpid Gloves (8204, -1.26 DPS) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 81.6 ranged_attack_power points (8.27 DPS) | yes | Stalker's Mail Belt (252588, -0.90 DPS) [crafted]; Ban'thok Sash (11662, -2.56 DPS) [dungeon]; Dawnspire Cord (12466, -2.66 DPS) [dungeon] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 97.3 ranged_attack_power points (9.86 DPS) | yes | Tough Scorpid Leggings (8206, -1.55 DPS) [crafted]; Searingscale Leggings (11749, -1.93 DPS) [dungeon]; Imperial Red Pants (8251, -2.30 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 76.6 ranged_attack_power points (7.76 DPS) | yes | Stalker's Mail Sabatons (252576, -1.01 DPS) [crafted]; Albino Crocscale Boots (17728, -1.45 DPS) [dungeon]; Rancor Boots (11865, -2.72 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.3 ranged_attack_power points (4.49 DPS) | yes | Mindseye Circle (10634, -0.95 DPS) [dungeon]; Ring of the Underwood (2951, -1.34 DPS) [world_drop]; Brainlash (6440, -3.69 DPS, sim-verified) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindseye Circle (10634, -0.87 DPS) [dungeon]; Ring of the Underwood (2951, -1.26 DPS) [world_drop]; Brainlash (6440, -6.08 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.85 DPS, sim-verified) [crafted] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Brightwood Staff (812, -1.04 DPS) [world_drop]; Hanzo Sword (8190, -1.15 DPS, sim-verified) [world_drop]; Warmonger (13052, -2.20 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Megashot Rifle (17717, -3.34 DPS) [dungeon]; Dark Iron Rifle (16004, -3.94 DPS) [crafted]; Gryphonwing Long Bow (13022, -4.46 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (orc, 5420001505001251-0052451001200000-000000000000000000)

Set DPS (verified): 468.7. Weights run: 8.1s. Verify run: 89.3s. 2027 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.218 ± 0.030, crit=2.450 ± 0.059 per rating point (14 rating = 1%, 34.301 per %), hit=2.993 ± 0.119 per rating point (10 rating = 1%, 29.930 per %), melee_haste=not significant (3.594 ± 1.495), intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 147.9 ranged_attack_power points (19.26 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -20.34 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.84 DPS) [world_drop]; Talisman of Evasion (13177, -0.84 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 121.2 ranged_attack_power points (15.79 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.21 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -1.77 DPS) [dungeon]; Deathguard's Cloak (20068, -2.37 DPS) [rep]; Howler's Furs (272414, -5.46 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 183.3 ranged_attack_power points (23.88 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.27 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.27 DPS) [vendor]; Savage Gladiator Chain (11726, -11.21 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 61.1 ranged_attack_power points (7.97 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.84 DPS) [pvp]; Slashclaw Bracers (13211, -1.13 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -0.82 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Girdle (16572, -1.88 DPS) [pvp]; Marksman's Girdle (22232, -3.09 DPS, sim-verified) [dungeon]; Defiler's Chain Girdle (20150, -3.37 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 176.8 ranged_attack_power points (23.04 DPS) | yes | General's Chain Legguards (231574, -2.43 DPS) [pvp]; General's Chain Legplates (231567, -4.86 DPS) [vendor]; Plaguehound Leggings (18736, -7.22 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.64 DPS) [dungeon]; Blackstone Ring (17713, -3.69 DPS) [dungeon]; Naglering (11669, -10.82 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.88 DPS, sim-verified) [dungeon]; Blackstone Ring (17713, -3.26 DPS) [dungeon]; Innervating Band (18701, -3.48 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -1.95 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (468.7 DPS) | yes | Hand of Justice (11815, -0.90 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -5.08 DPS) [pvp]; Glacial Blade (19099, -7.12 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -2.10 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (orc, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 940.6. Weights run: 8.5s. Verify run: 86.5s. 2027 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.921 ± 0.063, crit=3.596 ± 0.123 per rating point (14 rating = 1%, 50.344 per %), hit=5.148 ± 0.252 per rating point (10 rating = 1%, 51.475 per %), melee_haste=not significant (3.603 ± 3.143), intellect=1.767 ± 0.036, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 199.8 ranged_attack_power points (32.39 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -5.65 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 156.2 ranged_attack_power points (25.32 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.31 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shroud of Arcane Mastery (22330, -1.28 DPS) [dungeon]; Shadow Prowler's Cloak (22269, -1.97 DPS) [dungeon]; Howler's Furs (272414, -9.10 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 246.9 ranged_attack_power points (40.02 DPS) | yes | Warlord's Chain Chestpiece (16565, -6.81 DPS) [vendor]; Warlord's Chain Hauberk (231573, -6.81 DPS) [vendor]; Savage Gladiator Chain (11726, -12.42 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+8.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.27 DPS) [pvp]; Slashclaw Bracers (13211, -8.79 DPS, sim-verified) [dungeon] |
| hands | Mourning Handguards (279257) | Leatherworking [crafted] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -5.48 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Girdle (16572, -2.67 DPS) [pvp]; Clutch of Andros (13956, -5.48 DPS) [dungeon]; Marksman's Girdle (22232, -6.02 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 239.1 ranged_attack_power points (38.75 DPS) | yes | General's Chain Legguards (231574, -5.54 DPS) [pvp]; General's Chain Legplates (231567, -10.51 DPS) [vendor]; Plaguehound Leggings (18736, -13.42 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -6.42 DPS) [vendor]; Blackstone Ring (17713, -7.19 DPS) [dungeon]; Naglering (11669, -12.91 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -5.66 DPS) [dungeon]; Channeler's Ring (272406, -6.13 DPS, sim-verified) [vendor]; Painweaver Band (13098, -6.49 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (940.6 DPS) | yes | Frozen Heart of the Mountain (249469, -0.60 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Warmonger (13052, -2.23 DPS) [world_drop]; Glacial Blade (19099, -3.70 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -15.60 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Mourning Handguards; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Hand of Justice; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

