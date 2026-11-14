# Leveling BiS: Beast Mastery

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5320000100000000-0000000000000000-000000000000000000)

Set DPS (verified): 89.6. Weights run: 4.9s. Verify run: 10.2s. 298 eligible items had no known source.

Race sweep: night-elf 89.6 ± 0.1, high-order-skyborne 88.4 ± 0.1, human 87.9 ± 0.1, dwarf 85.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.209 ± 0.011, crit=0.713 ± 0.012 per rating point (14 rating = 1%, 9.977 per %), hit=1.126 ± 0.033 per rating point (10 rating = 1%, 11.261 per %), melee_haste=5.608 ± 0.531, intellect=1.730 ± 0.027, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.7 ranged_attack_power points (2.33 DPS) | yes | Trapper's Leather Hood (252505, -0.93 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.54 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 19.3 ranged_attack_power points (1.75 DPS) | yes | Erudite's Amulet (277204, -1.11 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 16.0 ranged_attack_power points (1.46 DPS) | yes | Magician's Mantle (12998, -0.04 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -0.83 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.28 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.58 DPS) [world_drop]; Bristlebark Cape (14571, -0.58 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 35.3 ranged_attack_power points (3.20 DPS) | yes | Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -1.17 DPS) [crafted]; Mystic's Wrap (14369, -1.23 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.0 ranged_attack_power points (1.46 DPS) | yes | Wolf Bracers (4794, -0.28 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.58 DPS) [world_drop]; Bravo's Armbands (270015, -0.58 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.3 ranged_attack_power points (1.75 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.11 DPS) [crafted]; Riding Gloves (1304, -0.40 DPS) [quest] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Trapper's Leather Belt (252429, -0.16 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Duty Bound Leggings (279868, -0.25 DPS) [quest]; Trapper's Leather Pants (252501, -1.03 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.7 ranged_attack_power points (2.33 DPS) | yes | Trapper's Leather Boots (252440, -0.15 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.8 ranged_attack_power points (1.52 DPS) | yes | Volcanic Rock Ring (12053, -1.05 DPS) [world_drop]; Minor Channeling Ring (1449, -1.21 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.8 ranged_attack_power points (1.17 DPS) | yes | Volcanic Rock Ring (12053, -0.82 DPS, sim-verified) [world_drop]; Minor Channeling Ring (1449, -0.85 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 37.5 ranged_attack_power points (3.41 DPS) | yes | Bronze Dory (250603, -0.37 DPS, sim-verified) [crafted]; Lupine Axe (1220, -1.66 DPS) [world]; Piercing Axe (6094, -1.66 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (16.29 DPS) | yes | Venomstrike (6469, -0.29 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.20 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.94 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 5320000505000100-0000000000000000-000000000000000000)

Set DPS (verified): 138.5. Weights run: 4.8s. Verify run: 14.2s. 494 eligible items had no known source.

Race sweep: night-elf 138.5 ± 0.1, human 137.4 ± 0.1, high-order-skyborne 137.2 ± 0.1, dwarf 135.3 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.086 ± 0.011, crit=0.943 ± 0.015 per rating point (14 rating = 1%, 13.209 per %), hit=1.518 ± 0.047 per rating point (10 rating = 1%, 15.184 per %), melee_haste=5.662 ± 0.628, intellect=1.890 ± 0.020, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Resilient Cap (14401, -0.34 DPS) [world_drop]; Brawler's Leather Helm (252512, -0.57 DPS, sim-verified) [crafted]; Brawler's Leather Hood (252504, -0.58 DPS) [crafted] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 24.7 ranged_attack_power points (2.57 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop]; Seal of the Expedition (271920, -1.59 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.9 ranged_attack_power points (3.53 DPS) | yes | Mantle of Thieves (2264, -0.32 DPS) [dungeon]; Dark Leather Shoulders (4252, -1.28 DPS) [crafted]; Insignia Mantle (4721, -1.28 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.32 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.50 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.64 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 56.1 ranged_attack_power points (5.84 DPS) | yes | Dusky Leather Armor (7374, -1.37 DPS, sim-verified) [crafted]; Green Leather Armor (4255, -1.70 DPS) [crafted]; Trapper's Leather Tunic (252509, -2.09 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.5 ranged_attack_power points (1.93 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.32 DPS) [world]; Forest Leather Bracers (3202, -0.32 DPS) [world_drop] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.8 ranged_attack_power points (2.89 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.03 DPS) [dungeon]; Pilferer's Gloves (7358, -0.32 DPS) [crafted]; Resilient Handgrips (14403, -0.75 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.1 ranged_attack_power points (4.07 DPS) | yes | Green Leather Belt (4257, -1.07 DPS, sim-verified) [crafted]; Skulker's Leather Belt (252520, -1.18 DPS) [crafted]; Prowler's Leather Belt (252459, -1.36 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 48.7 ranged_attack_power points (5.07 DPS) | yes | Petrolspill Leggings (9509, -0.57 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Trapper's Leather Legguards (252517, -0.80 DPS) [crafted] |
| feet | Insignia Boots (4055) (or Feet of the Lynx (1121)) | World drop [world_drop] | 24.7 ranged_attack_power points (2.57 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Trapper's Leather Boots (252440, -0.18 DPS) [crafted]; Footpads of the Fang (10411, -0.64 DPS) [dungeon] |
| finger1 | Ring of Precision (1491) (or Protector's Band (19517)) | Blackfathom Deeps [dungeon] | 18.5 ranged_attack_power points (1.93 DPS) | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Seal of Wrynn (2933, -0.18 DPS) [quest]; Monkey Ring (6748, -0.32 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.5 ranged_attack_power points (1.93 DPS) | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Seal of Wrynn (2933, -0.18 DPS) [quest]; Monkey Ring (6748, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornspike (6681, -0.84 DPS, sim-verified) [dungeon]; Impaling Harpoon (5200, -0.92 DPS) [dungeon]; Kam's Walking Stick (2280, -0.99 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 250.9 ranged_attack_power points (26.11 DPS) | yes | Crystalpine Stinger (13037, -0.60 DPS, sim-verified) [world_drop]; Ironweaver (13137, -0.90 DPS) [world_drop]; Harpyclaw Short Bow (13019, -1.85 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Insignia Boots; finger1: Ring of Precision; finger2: Protector's Band; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 494, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 222.5. Weights run: 5.1s. Verify run: 10.6s. 827 eligible items had no known source.

Race sweep: night-elf 222.5 ± 0.1, human 221.6 ± 0.1, high-order-skyborne 220.7 ± 0.1, dwarf 218.3 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.252 ± 0.033, crit=1.600 ± 0.054 per rating point (14 rating = 1%, 22.397 per %), hit=2.231 ± 0.098 per rating point (10 rating = 1%, 22.306 per %), melee_haste=7.613 ± 1.067, intellect=1.767 ± 0.029, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 49.8 ranged_attack_power points (5.56 DPS) | yes | Warden's Wizard Hat (14604, -0.25 DPS) [world_drop]; Nightscape Headband (8176, -1.21 DPS) [crafted]; Guard's Chain Helm (250499, -1.21 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 35.8 ranged_attack_power points (3.99 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.47 DPS) [world]; Kaleidoscope Chain (13084, -1.75 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -2.41 DPS) [dungeon] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 35.8 ranged_attack_power points (3.99 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Windchaser Amice (14432, -0.34 DPS) [world_drop]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.46 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.46 DPS) [world_drop]; Parachute Cloak (10518, -1.42 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 100.3 ranged_attack_power points (11.19 DPS) | yes | Tough Scorpid Breastplate (8203, -4.37 DPS) [crafted]; Quillward Harness (10583, -4.83 DPS, sim-verified) [dungeon]; Garb of Florid Feathers (273026, -4.97 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.4 ranged_attack_power points (3.72 DPS) | yes | Green Leather Bracers (4259, -0.56 DPS) [crafted]; Imperial Leather Bracers (4061, -0.82 DPS) [dungeon]; Dusky Bracers (7378, -0.82 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Stalker's Leather Gloves (252526, +0.00 DPS) [crafted]; Tough Scorpid Gloves (8204, -2.88 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.9 ranged_attack_power points (4.44 DPS) | yes | Highlander's Chain Girdle (20089, -1.06 DPS) [rep]; Green Leather Belt (4257, -1.09 DPS) [crafted]; Scorpashi Sash (14652, -1.61 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 68.3 ranged_attack_power points (7.61 DPS) | yes | Triprunner Dungarees (9624, -1.34 DPS, sim-verified) [quest]; Ferine Leggings (6690, -2.35 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.54 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 48.1 ranged_attack_power points (5.37 DPS) | yes | Swampwalker Boots (2276, -0.65 DPS) [world_drop]; Imperial Leather Boots (6431, -1.38 DPS) [dungeon]; Worn Running Boots (9398, -1.41 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.5 ranged_attack_power points (3.62 DPS) | yes | Protector's Band (19515, -0.72 DPS) [rep]; Field Researcher's Loop (281634, -1.09 DPS) [quest]; Falcon's Hook (7552, -4.05 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19515, +0.00 DPS) [rep]; Falcon's Hook (7552, -0.28 DPS, sim-verified) [dungeon]; Field Researcher's Loop (281634, -0.36 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -1.56 DPS) [world]; Steel Spear (250605, -1.84 DPS) [crafted]; Claw of the Shadowmancer (2912, -7.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.32 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.50 DPS) [dungeon]; The Silencer (13138, -4.07 DPS, sim-verified) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 827, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5420001505001251-0050500000000000-000000000000000000)

Set DPS (verified): 313.4. Weights run: 5.1s. Verify run: 22.2s. 1065 eligible items had no known source.

Race sweep: night-elf 313.4 ± 0.2, human 312.1 ± 0.2, high-order-skyborne 311.2 ± 0.2, dwarf 309.0 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.031 ± 0.029, crit=1.600 ± 0.050 per rating point (14 rating = 1%, 22.402 per %), hit=2.221 ± 0.096 per rating point (10 rating = 1%, 22.211 per %), melee_haste=10.348 ± 1.049, intellect=1.865 ± 0.040, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 68.5 ranged_attack_power points (7.38 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Exile (11124, -0.18 DPS) [quest]; Chief Architect's Monocle (11839, -1.90 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.4 ranged_attack_power points (4.24 DPS) | yes | Sentinel's Medallion (19539, -0.33 DPS) [rep]; Horizon Choker (13085, -1.43 DPS) [world_drop] |
| shoulder | Stalker's Mail Shoulders (252566) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.10 DPS) [dungeon]; Tough Scorpid Shoulders (8207, -0.49 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.33 DPS) [dungeon]; Blackmetal Cape (9512, -0.98 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.9 ranged_attack_power points (10.22 DPS) | yes | Blazewind Breastplate (11193, -2.72 DPS) [quest]; Knight's Chain Armor (220828, -3.17 DPS, sim-verified) [vendor]; Tough Scorpid Breastplate (8203, -3.92 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 45.5 ranged_attack_power points (4.89 DPS) | yes | Bracers of the Stone Princess (17714, -0.88 DPS) [dungeon]; Bloodlust Bracelets (14807, -1.31 DPS) [world_drop]; Stalker's Mail Bracers (252570, -1.46 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hibernal Gloves (8110, +0.00 DPS) [world_drop]; Rockgrip Gauntlets (17736, +0.00 DPS) [dungeon]; Stalker's Mail Gauntlets (252584, -1.68 DPS, sim-verified) [crafted] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 61.1 ranged_attack_power points (6.58 DPS) | yes | Sagebrush Girdle (17778, -1.22 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -2.01 DPS) [crafted]; Highlander's Chain Girdle (20088, -2.01 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 97.7 ranged_attack_power points (10.52 DPS) | yes | Imperial Red Pants (8251, -2.62 DPS) [world_drop]; Tough Scorpid Leggings (8206, -2.97 DPS) [crafted]; Knight's Chain Legplates (220832, -5.35 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.8 ranged_attack_power points (7.40 DPS) | yes | Albino Crocscale Boots (17728, -0.88 DPS) [dungeon]; Stalker's Mail Sabatons (252576, -1.36 DPS) [crafted]; Elven Chain Boots (13125, -1.53 DPS) [world_drop] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.4 ranged_attack_power points (4.57 DPS) | yes | Mindseye Circle (10634, -0.96 DPS) [dungeon]; Ring of the Underwood (2951, -1.31 DPS) [world_drop]; Falcon's Hook (7552, -1.63 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.2 ranged_attack_power points (4.54 DPS) | yes | Ring of the Underwood (2951, -1.28 DPS) [world_drop]; Falcon's Hook (7552, -1.61 DPS) [dungeon]; Mindseye Circle (10634, -4.43 DPS, sim-verified) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Eye (19991, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Devilsaur Tooth (19992) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.13 DPS) [crafted]; Devilsaur Eye (19991, -0.56 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eyegouger (9480, -0.10 DPS) [dungeon]; Warmonger (13052, -0.43 DPS) [world_drop]; Shadowblade (2163, -1.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Gryphonwing Long Bow (13022, -0.04 DPS) [dungeon]; Hurricane (2824, -1.61 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -1.80 DPS) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Stalker's Mail Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Devilsaur Tooth; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1065, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5420001505001251-0050550001000000-400000000000000000)

Set DPS (verified): 483.0. Weights run: 5.1s. Verify run: 55.0s. 2313 eligible items had no known source.

Race sweep: night-elf 483.0 ± 0.3, human 481.6 ± 0.3, high-order-skyborne 481.5 ± 0.3, dwarf 478.1 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.115 ± 0.029, crit=2.276 ± 0.058 per rating point (14 rating = 1%, 31.870 per %), hit=3.233 ± 0.128 per rating point (10 rating = 1%, 32.333 per %), melee_haste=15.149 ± 1.389, intellect=1.237 ± 0.005, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 144.3 ranged_attack_power points (19.90 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.19 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.84 DPS) [pvp] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.9 ranged_attack_power points (7.98 DPS) | yes | Star of Mystaria (12103, -0.99 DPS) [dungeon]; Sentinel's Medallion (19538, -1.54 DPS) [rep]; Nacreous Shell Necklace (22403, -2.16 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 116.0 ranged_attack_power points (15.99 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -0.88 DPS) [vendor]; Shadow Prowler's Cloak (22269, -2.33 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.36 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 179.4 ranged_attack_power points (24.75 DPS) | yes | Field Marshal's Chain Breastplate (16466, -3.70 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -3.70 DPS) [vendor]; Field Marshal's Chain Armor (231563, -6.04 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.2 ranged_attack_power points (8.16 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Slashclaw Bracers (13211, -0.70 DPS) [dungeon]; Knight-Captain's Chain Armguards (16402, -0.86 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Sash of the Grand Hunt (22207, -0.77 DPS) [dungeon]; Marshal's Chain Girdle (16464, -1.90 DPS) [pvp] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 205.1 ranged_attack_power points (28.29 DPS) | yes | Marshal's Chain Legguards (231577, -7.24 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.90 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.57 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 99.6 ranged_attack_power points (13.73 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.57 DPS) [dungeon]; Channeler's Ring (272406, -3.94 DPS) [vendor]; Innervating Band (18701, -4.34 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.26 DPS) [dungeon]; Channeler's Ring (272406, -3.63 DPS) [vendor]; Innervating Band (18701, -4.03 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Warmonger (13052, -5.40 DPS) [world_drop]; Grand Marshal's Stave (234571, -6.07 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Grand Marshal's Hand Cannon (234587, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2313, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 977.2. Weights run: 5.3s. Verify run: 55.7s. 2313 eligible items had no known source.

Race sweep: night-elf 977.2 ± 0.5, high-order-skyborne 974.5 ± 0.5, human 972.0 ± 0.5, dwarf 969.2 ± 0.5 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.877 ± 0.055, crit=3.672 ± 0.118 per rating point (14 rating = 1%, 51.402 per %), hit=5.180 ± 0.242 per rating point (10 rating = 1%, 51.795 per %), melee_haste=not significant (4.507 ± 3.081), intellect=1.713 ± 0.034, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 197.9 ranged_attack_power points (32.77 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -1.64 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -3.98 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -1.97 DPS) [dungeon]; Sentinel's Medallion (19538, -3.19 DPS) [rep] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 156.1 ranged_attack_power points (25.84 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, +0.00 DPS) [vendor]; Arcanoweave Cloak (272411, -2.10 DPS) [vendor]; Shadow Prowler's Cloak (22269, -2.67 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 246.7 ranged_attack_power points (40.84 DPS) | yes | Field Marshal's Chain Breastplate (16466, -7.10 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -7.10 DPS) [vendor]; Knight-Captain's Chain Hauberk (23292, -11.84 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.28 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 133.2 ranged_attack_power points (22.06 DPS) | yes | Sash of the Grand Hunt (22207, -2.15 DPS) [dungeon]; Ranger's Belt (272397, -2.68 DPS) [vendor]; Marshal's Chain Girdle (16464, -5.40 DPS) [pvp] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 290.3 ranged_attack_power points (48.06 DPS) | yes | Sentinel's Leather Pants (237818, -13.71 DPS) [vendor]; Marshal's Chain Legguards (231577, -14.32 DPS) [pvp]; Knight-Captain's Chain Legguards (227072, -19.07 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 136.2 ranged_attack_power points (22.55 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -6.86 DPS) [vendor]; Blackstone Ring (17713, -7.52 DPS) [dungeon]; Painweaver Band (13098, -8.91 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -5.02 DPS) [vendor]; Blackstone Ring (17713, -5.68 DPS) [dungeon]; Painweaver Band (13098, -7.07 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Warmonger (13052, -2.06 DPS) [world_drop]; Grand Marshal's Stave (234571, -4.24 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Grand Marshal's Hand Cannon (234587, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2313, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (orc, 5320000100000000-0000000000000000-000000000000000000)

Set DPS (verified): 87.0. Weights run: 4.9s. Verify run: 9.8s. 294 eligible items had no known source.

Race sweep: windshaper-skyborne 88.0 ± 0.1, orc 87.0 ± 0.1, troll 86.8 ± 0.1, tauren 84.5 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 1.0 (1.14%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.209 ± 0.011, crit=0.713 ± 0.012 per rating point (14 rating = 1%, 9.977 per %), hit=1.126 ± 0.033 per rating point (10 rating = 1%, 11.261 per %), melee_haste=5.608 ± 0.531, intellect=1.730 ± 0.027, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.7 ranged_attack_power points (2.33 DPS) | yes | Trapper's Leather Hood (252505, -1.11 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.54 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 19.3 ranged_attack_power points (1.75 DPS) | yes | Erudite's Amulet (277204, -1.16 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 16.0 ranged_attack_power points (1.46 DPS) | yes | Magician's Mantle (12998, -0.04 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -0.83 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.29 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.58 DPS) [world_drop]; Bristlebark Cape (14571, -0.58 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 27.7 ranged_attack_power points (2.51 DPS) | yes | Brawler's Leather Armor (252490, -0.29 DPS, sim-verified) [crafted]; Mystic's Wrap (14369, -0.54 DPS) [world_drop]; Mystic's Robe (14371, -0.54 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.0 ranged_attack_power points (1.46 DPS) | yes | Wolf Bracers (4794, -0.29 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.58 DPS) [world_drop]; Spare Part Bindings (279875, -0.58 DPS) [quest] |
| hands | Gloves of the Fang (10413) (or Serpent Gloves (5970)) | Wailing Caverns: Druid of the Fang [dungeon] | 19.3 ranged_attack_power points (1.75 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Trapper's Leather Gloves (252495, -0.11 DPS) [crafted]; Forest Leather Gloves (3058, -0.58 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Trapper's Leather Belt (252429, -0.17 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.18 DPS) [crafted]; Keller's Girdle (2911, -0.38 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.29 DPS) [world]; Trapper's Leather Pants (252501, -1.06 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.7 ranged_attack_power points (2.33 DPS) | yes | Trapper's Leather Boots (252440, -0.25 DPS) [crafted]; Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 16.8 ranged_attack_power points (1.52 DPS) | yes | Bounty Hunter's Ring (5351, -0.65 DPS) [quest]; Loop of Sacrifice (281673, -0.74 DPS) [quest]; Volcanic Rock Ring (12053, -1.05 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.8 ranged_attack_power points (1.17 DPS) | yes | Bounty Hunter's Ring (5351, -0.29 DPS, sim-verified) [quest]; Loop of Sacrifice (281673, -0.38 DPS) [quest]; Volcanic Rock Ring (12053, -0.69 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 37.5 ranged_attack_power points (3.41 DPS) | yes | Crescent Staff (6505, -0.37 DPS, sim-verified) [quest]; Bronze Dory (250603, -0.53 DPS) [crafted]; Staff of Orgrimmar (15444, -0.72 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (16.29 DPS) | yes | Venomstrike (6469, -0.33 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.20 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.94 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Lavishly Jeweled Ring; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (orc, 5320000505000100-0000000000000000-000000000000000000)

Set DPS (verified): 136.9. Weights run: 4.8s. Verify run: 15.1s. 493 eligible items had no known source.

Race sweep: windshaper-skyborne 137.2 ± 0.1, orc 136.9 ± 0.1, troll 136.2 ± 0.1, tauren 133.2 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 0.3 (0.25%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.086 ± 0.011, crit=0.943 ± 0.015 per rating point (14 rating = 1%, 13.209 per %), hit=1.518 ± 0.047 per rating point (10 rating = 1%, 15.184 per %), melee_haste=5.662 ± 0.628, intellect=1.890 ± 0.020, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Helm (252512, +0.00 DPS) [crafted]; Resilient Cap (14401, -0.34 DPS) [world_drop]; Brawler's Leather Hood (252504, -0.58 DPS) [crafted] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 24.7 ranged_attack_power points (2.57 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; Kaleidoscope Chain (13084, -0.50 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.59 DPS) [vendor] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.9 ranged_attack_power points (3.53 DPS) | yes | Mantle of Thieves (2264, -0.32 DPS) [dungeon]; Dark Leather Shoulders (4252, -1.28 DPS) [crafted]; Insignia Mantle (4721, -1.28 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.32 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.64 DPS) [dungeon]; Swiftrunner Cape (6745, -0.64 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 56.1 ranged_attack_power points (5.84 DPS) | yes | Dusky Leather Armor (7374, -1.34 DPS) [crafted]; Green Leather Armor (4255, -1.70 DPS) [crafted]; Trapper's Leather Tunic (252509, -2.09 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) (or Insignia Bracers (6410)) | World drop [world_drop] | 18.5 ranged_attack_power points (1.93 DPS) | yes | Insignia Bracers (6410, +0.00 DPS) [world_drop]; Forest Leather Bracers (3202, -0.32 DPS) [world_drop]; Loamflake Bracers (15462, -0.32 DPS) [quest] |
| hands | Cartilage Shapers (273054) | Lyn the Ignored [world] | 27.8 ranged_attack_power points (2.89 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.03 DPS) [dungeon]; Pilferer's Gloves (7358, -0.32 DPS) [crafted]; Braced Handguards (6784, -0.64 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.1 ranged_attack_power points (4.07 DPS) | yes | Green Leather Belt (4257, -0.96 DPS) [crafted]; Skulker's Leather Belt (252520, -1.18 DPS) [crafted]; Prowler's Leather Belt (252459, -1.36 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 48.7 ranged_attack_power points (5.07 DPS) | yes | Petrolspill Leggings (9509, -0.57 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.57 DPS) [world_drop]; Trapper's Leather Legguards (252517, -0.80 DPS) [crafted] |
| feet | Insignia Boots (4055) (or Warsong Boots (16977), Feet of the Lynx (1121)) | World drop [world_drop] | 24.7 ranged_attack_power points (2.57 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Trapper's Leather Boots (252440, -0.18 DPS) [crafted] |
| finger1 | Monkey Ring (6748) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Legionnaire's Band (19513, +0.00 DPS) [rep]; Band of the Fist (17694, -0.32 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19513, +0.00 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Band of the Fist (17694, -0.64 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, -0.92 DPS) [dungeon]; Kam's Walking Stick (2280, -0.99 DPS) [dungeon]; Crescent Staff (6505, -1.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 250.9 ranged_attack_power points (26.11 DPS) | yes | Crystalpine Stinger (13037, -0.89 DPS) [world_drop]; Ironweaver (13137, -0.90 DPS) [world_drop]; Harpyclaw Short Bow (13019, -1.85 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Cartilage Shapers; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Insignia Boots; finger1: Monkey Ring; finger2: Ring of Precision; main_hand: Felbough Arm; ranged: Golemsight Long Gun

No-known-source sample (15 of 493, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (orc, 5420001505001251-0000000000000000-000000000000000000)

Set DPS (verified): 219.6. Weights run: 5.1s. Verify run: 10.6s. 817 eligible items had no known source.

Race sweep: windshaper-skyborne 219.6 ± 0.1, orc 219.6 ± 0.1, troll 217.6 ± 0.1, tauren 213.9 ± 0.1 (dps); windshaper-skyborne beats the guide race orc by 0.0 (0.00%) ± 0.2, inside error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.252 ± 0.033, crit=1.600 ± 0.054 per rating point (14 rating = 1%, 22.397 per %), hit=2.231 ± 0.098 per rating point (10 rating = 1%, 22.306 per %), melee_haste=7.613 ± 1.067, intellect=1.767 ± 0.029, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 49.8 ranged_attack_power points (5.56 DPS) | yes | Nightscape Headband (8176, -1.21 DPS) [crafted]; Guard's Chain Helm (250499, -1.21 DPS) [crafted]; Warden's Wizard Hat (14604, -1.22 DPS, sim-verified) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 35.8 ranged_attack_power points (3.99 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.47 DPS) [world]; Kaleidoscope Chain (13084, -1.75 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -2.41 DPS) [dungeon] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 35.8 ranged_attack_power points (3.99 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Windchaser Amice (14432, -0.34 DPS) [world_drop]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.36 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 100.3 ranged_attack_power points (11.19 DPS) | yes | Tough Scorpid Breastplate (8203, -4.37 DPS) [crafted]; Quillward Harness (10583, -4.78 DPS, sim-verified) [dungeon]; Garb of Florid Feathers (273026, -4.97 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 33.4 ranged_attack_power points (3.72 DPS) | yes | Green Leather Bracers (4259, -0.56 DPS) [crafted]; Imperial Leather Bracers (4061, -0.82 DPS) [dungeon]; Dusky Bracers (7378, -0.82 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Stalker's Leather Gloves (252526, +0.00 DPS) [crafted]; Tough Scorpid Gloves (8204, -2.28 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 39.9 ranged_attack_power points (4.44 DPS) | yes | Defiler's Chain Girdle (20153, -1.06 DPS) [rep]; Green Leather Belt (4257, -1.09 DPS) [crafted]; Scorpashi Sash (14652, -1.59 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 68.3 ranged_attack_power points (7.61 DPS) | yes | Triprunner Dungarees (9624, -1.09 DPS, sim-verified) [quest]; Ferine Leggings (6690, -2.35 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.54 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 48.1 ranged_attack_power points (5.37 DPS) | yes | Swampwalker Boots (2276, -0.65 DPS) [world_drop]; Imperial Leather Boots (6431, -1.38 DPS) [dungeon]; Worn Running Boots (9398, -2.02 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.5 ranged_attack_power points (3.62 DPS) | yes | Legionnaire's Band (19512, -0.72 DPS) [rep]; Field Researcher's Loop (281634, -1.09 DPS) [quest]; Falcon's Hook (7552, -3.24 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Band (19512, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.36 DPS) [quest]; Falcon's Hook (7552, -0.46 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -1.56 DPS) [world]; Steel Spear (250605, -1.84 DPS) [crafted]; Claw of the Shadowmancer (2912, -7.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.32 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.50 DPS) [dungeon]; The Silencer (13138, -4.60 DPS, sim-verified) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Gloves of Holy Might; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 817, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (orc, 5420001505001251-0050500000000000-000000000000000000)

Set DPS (verified): 313.8. Weights run: 5.1s. Verify run: 22.6s. 1056 eligible items had no known source.

Race sweep: orc 313.8 ± 0.2, windshaper-skyborne 312.4 ± 0.2, troll 310.6 ± 0.2, tauren 306.2 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.031 ± 0.029, crit=1.600 ± 0.050 per rating point (14 rating = 1%, 22.402 per %), hit=2.221 ± 0.096 per rating point (10 rating = 1%, 22.211 per %), melee_haste=10.348 ± 1.049, intellect=1.865 ± 0.040, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 68.5 ranged_attack_power points (7.38 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Exile (11124, -0.18 DPS) [quest]; Chief Architect's Monocle (11839, -1.10 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.4 ranged_attack_power points (4.24 DPS) | yes | Scout's Medallion (19535, -0.33 DPS) [rep]; Woven Ivy Necklace (19159, -1.31 DPS) [quest] |
| shoulder | Stalker's Mail Shoulders (252566) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.10 DPS) [dungeon]; Tough Scorpid Shoulders (8207, -0.51 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.33 DPS) [dungeon]; Blackmetal Cape (9512, -0.98 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.9 ranged_attack_power points (10.22 DPS) | yes | Stone Guard's Chain Armor (220827, -1.83 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.72 DPS) [quest]; Tough Scorpid Breastplate (8203, -3.92 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 45.5 ranged_attack_power points (4.89 DPS) | yes | Bracers of the Stone Princess (17714, -0.88 DPS) [dungeon]; Bloodlust Bracelets (14807, -1.31 DPS) [world_drop]; Stalker's Mail Bracers (252570, -1.37 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Hibernal Gloves (8110, +0.00 DPS) [world_drop]; Rockgrip Gauntlets (17736, +0.00 DPS) [dungeon]; Stalker's Mail Gauntlets (252584, -1.00 DPS, sim-verified) [crafted] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 61.1 ranged_attack_power points (6.58 DPS) | yes | Sagebrush Girdle (17778, -1.10 DPS, sim-verified) [quest]; Skulker's Leather Waistguard (252474, -2.01 DPS) [crafted]; Defiler's Chain Girdle (20151, -2.01 DPS) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 97.7 ranged_attack_power points (10.52 DPS) | yes | Imperial Red Pants (8251, -2.62 DPS) [world_drop]; Tough Scorpid Leggings (8206, -2.97 DPS) [crafted]; Stone Guard's Chain Legplates (220833, -4.18 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.8 ranged_attack_power points (7.40 DPS) | yes | Albino Crocscale Boots (17728, -0.88 DPS) [dungeon]; Stalker's Mail Sabatons (252576, -1.36 DPS) [crafted]; Elven Chain Boots (13125, -1.53 DPS) [world_drop] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 42.4 ranged_attack_power points (4.57 DPS) | yes | Mindseye Circle (10634, -0.96 DPS) [dungeon]; Ring of the Underwood (2951, -1.31 DPS) [world_drop]; Falcon's Hook (7552, -1.63 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.2 ranged_attack_power points (4.54 DPS) | yes | Ring of the Underwood (2951, -1.28 DPS) [world_drop]; Falcon's Hook (7552, -1.61 DPS) [dungeon]; Mindseye Circle (10634, -2.15 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, -0.90 DPS) [quest]; Devilsaur Eye (19991, -1.46 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Devilsaur Eye (19991, -0.43 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eyegouger (9480, -0.10 DPS) [dungeon]; Warmonger (13052, -0.43 DPS) [world_drop]; Shadowblade (2163, -1.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gryphonwing Long Bow (13022, -0.04 DPS) [dungeon]; Hurricane (2824, -1.54 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -1.80 DPS) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Stalker's Mail Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1056, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (orc, 5420001505001251-0050550001000000-400000000000000000)

Set DPS (verified): 479.4. Weights run: 5.1s. Verify run: 54.4s. 2312 eligible items had no known source.

Race sweep: orc 479.4 ± 0.3, windshaper-skyborne 476.8 ± 0.3, troll 474.9 ± 0.3, tauren 467.8 ± 0.3 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.115 ± 0.029, crit=2.276 ± 0.058 per rating point (14 rating = 1%, 31.870 per %), hit=3.233 ± 0.128 per rating point (10 rating = 1%, 32.333 per %), melee_haste=15.149 ± 1.389, intellect=1.237 ± 0.005, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 144.3 ranged_attack_power points (19.90 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -30.15 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.9 ranged_attack_power points (7.98 DPS) | yes | Scout's Medallion (19534, -1.54 DPS) [rep]; Nacreous Shell Necklace (22403, -2.16 DPS) [dungeon]; Star of Mystaria (12103, -2.96 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 116.0 ranged_attack_power points (15.99 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -0.38 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -2.33 DPS) [dungeon]; Deathguard's Cloak (20068, -2.36 DPS) [rep]; Howler's Furs (272414, -5.29 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 179.4 ranged_attack_power points (24.75 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.70 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.70 DPS) [vendor]; Icemail Jerkin (1981, -9.97 DPS, sim-verified) [world_drop] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.2 ranged_attack_power points (8.16 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.86 DPS) [pvp]; Slashclaw Bracers (13211, -3.17 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -2.14 DPS, sim-verified) [dungeon] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Sash of the Grand Hunt (22207, -0.77 DPS) [dungeon]; General's Chain Girdle (16572, -1.90 DPS) [pvp]; Marksman's Girdle (22232, -7.15 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 173.2 ranged_attack_power points (23.89 DPS) | yes | General's Chain Legguards (231574, -2.84 DPS) [pvp]; General's Chain Legplates (231567, -5.18 DPS) [vendor]; Plaguehound Leggings (18736, -6.25 DPS, sim-verified) [dungeon] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 99.6 ranged_attack_power points (13.73 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -3.37 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.57 DPS) [dungeon]; Channeler's Ring (272406, -3.94 DPS) [vendor]; Naglering (11669, -5.94 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.25 DPS, sim-verified) [dungeon]; Channeler's Ring (272406, -3.63 DPS) [vendor]; Innervating Band (18701, -4.03 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Second Wind (11819, -4.37 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Ankh of Life (1713, -4.48 DPS, sim-verified) [world_drop] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Warmonger (13052, -5.40 DPS) [world_drop]; Shadowblade (2163, -6.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Bloodseeker (19107, -3.95 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2312, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (orc, 5420001505001251-0052350001400000-000000000000000000)

Set DPS (verified): 969.5. Weights run: 5.3s. Verify run: 55.3s. 2312 eligible items had no known source.

Race sweep: orc 969.5 ± 0.5, windshaper-skyborne 965.5 ± 0.5, troll 960.8 ± 0.5, tauren 949.6 ± 0.5 (dps); the guide race orc is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.877 ± 0.055, crit=3.672 ± 0.118 per rating point (14 rating = 1%, 51.402 per %), hit=5.180 ± 0.242 per rating point (10 rating = 1%, 51.795 per %), melee_haste=not significant (4.507 ± 3.081), intellect=1.713 ± 0.034, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 197.9 ranged_attack_power points (32.77 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Champion's Chain Helm (227067, -1.64 DPS) [pvp] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -1.97 DPS) [dungeon]; Scout's Medallion (19534, -3.19 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 156.1 ranged_attack_power points (25.84 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.45 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, +0.00 DPS) [vendor]; Arcanoweave Cloak (272411, -2.10 DPS) [vendor]; Shadow Prowler's Cloak (22269, -2.67 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 246.7 ranged_attack_power points (40.84 DPS) | yes | Warlord's Chain Chestpiece (16565, -7.10 DPS) [vendor]; Warlord's Chain Hauberk (231573, -7.10 DPS) [vendor]; Legionnaire's Chain Hauberk (227071, -11.84 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.28 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 133.2 ranged_attack_power points (22.06 DPS) | yes | Sash of the Grand Hunt (22207, -2.15 DPS) [dungeon]; Ranger's Belt (272397, -2.68 DPS) [vendor]; General's Chain Girdle (16572, -5.40 DPS) [pvp] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 238.9 ranged_attack_power points (39.55 DPS) | yes | General's Chain Legguards (231574, -5.81 DPS) [pvp]; Legionnaire's Chain Legguards (227073, -10.56 DPS) [pvp]; General's Chain Legplates (231567, -10.77 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 136.2 ranged_attack_power points (22.55 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -6.00 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -6.86 DPS) [vendor]; Blackstone Ring (17713, -7.52 DPS) [dungeon]; Painweaver Band (13098, -8.91 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Channeler's Ring (272406, -5.02 DPS) [vendor]; Blackstone Ring (17713, -5.68 DPS) [dungeon]; Painweaver Band (13098, -7.07 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Warmonger (13052, -2.06 DPS) [world_drop]; High Warlord's War Staff (234549, -4.24 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; High Warlord's Street Sweeper (234561, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Marksman's Girdle; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2312, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

