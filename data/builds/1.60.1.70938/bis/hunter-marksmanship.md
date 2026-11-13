# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0050150000000000-000000000000000000)

Set DPS (verified): 95.6. Weights run: 5.9s. Verify run: 13.6s. 298 eligible items had no known source.

Race sweep: night-elf 95.6 ± 0.1, high-order-skyborne 94.2 ± 0.1, human 94.2 ± 0.1, dwarf 91.6 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.215 ± 0.011, crit=0.708 ± 0.012 per rating point (14 rating = 1%, 9.916 per %), hit=1.102 ± 0.036 per rating point (10 rating = 1%, 11.023 per %), melee_haste=3.903 ± 0.605, intellect=3.110 ± 0.026, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Trapper's Leather Hood (252505) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Hood (252504, -0.11 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.85 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 19.3 ranged_attack_power points (1.76 DPS) | yes | Erudite's Amulet (277204, -1.16 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 28.0 ranged_attack_power points (2.56 DPS) | yes | Serpent's Shoulders (5404, -0.59 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -1.42 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.29 DPS) [dungeon]; Sentry Cloak (2059, -0.59 DPS) [world_drop]; Bristlebark Cape (14571, -0.59 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mystic's Wrap (14369, -0.04 DPS) [world_drop]; Mystic's Robe (14371, -0.04 DPS) [world_drop]; Tunic of Westfall (2041, -0.11 DPS, sim-verified) [quest] |
| wrist | Owl Bracers (4796) | Bernard Brubaker [vendor] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Forest Leather Bracers (3202, -0.23 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -0.25 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Riding Gloves (1304, -0.03 DPS) [quest]; Trapper's Leather Gloves (252495, -2.33 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 24.9 ranged_attack_power points (2.27 DPS) | yes | Trapper's Leather Belt (252429, -0.25 DPS) [crafted]; Blackened Defias Belt (10403, -0.63 DPS) [dungeon]; Deviate Scale Belt (6468, -0.80 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 41.4 ranged_attack_power points (3.78 DPS) | yes | Leggings of the Fang (10410, -1.14 DPS) [dungeon]; Brawler's Leather Pants (252500, -1.14 DPS) [crafted]; Duty Bound Leggings (279868, -1.23 DPS, sim-verified) [quest] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Trapper's Leather Boots (252440, -1.82 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 25.1 ranged_attack_power points (2.29 DPS) | yes | Volcanic Rock Ring (12053, -1.44 DPS) [world_drop]; Minor Channeling Ring (1449, -1.72 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.9 ranged_attack_power points (1.17 DPS) | yes | Volcanic Rock Ring (12053, -0.13 DPS, sim-verified) [world_drop]; Minor Channeling Ring (1449, -0.61 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 44.5 ranged_attack_power points (4.06 DPS) | yes | Bronze Dory (250603, -1.04 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.22 DPS) [world_drop]; Sword of the Fallen (286977, -1.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (16.40 DPS) | yes | Venomstrike (6469, -0.81 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.20 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.96 DPS) [crafted] |

**New at 20:** head: Trapper's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Owl Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 147.5. Weights run: 5.8s. Verify run: 19.2s. 494 eligible items had no known source.

Race sweep: night-elf 147.5 ± 0.1, high-order-skyborne 146.0 ± 0.1, human 145.9 ± 0.1, dwarf 142.5 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.148 ± 0.012, crit=1.028 ± 0.016 per rating point (14 rating = 1%, 14.396 per %), hit=1.356 ± 0.037 per rating point (10 rating = 1%, 13.562 per %), melee_haste=6.807 ± 0.662, intellect=3.386 ± 0.029, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 54.2 ranged_attack_power points (4.62 DPS) | yes | Resilient Cap (14401, -0.41 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -1.16 DPS) [world_drop]; Insignia Cap (4052, -1.44 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 32.3 ranged_attack_power points (2.76 DPS) | yes | Kaleidoscope Chain (13084, -0.53 DPS) [world_drop]; Sentinel's Medallion (19541, -0.61 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Feline Mantle (3748, -0.11 DPS) [dungeon]; Mantle of Thieves (2264, -0.27 DPS) [dungeon]; Death Speaker Mantle (6685, -1.49 DPS, sim-verified) [dungeon] |
| back | Sergeant Major's Cape (16315) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Tigerstrike Mantle (13108, -0.08 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.35 DPS) [world_drop]; Expeditionary's Cape (271923, -0.62 DPS, sim-verified) [rep] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 79.1 ranged_attack_power points (6.75 DPS) | yes | Green Leather Armor (4255, -1.69 DPS, sim-verified) [crafted]; Resilient Tunic (14398, -2.50 DPS) [world_drop]; Resilient Robe (14405, -2.50 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 27.1 ranged_attack_power points (2.31 DPS) | yes | Nightsky Wristbands (6407, -0.34 DPS, sim-verified) [world_drop]; Insignia Bracers (6410, -0.70 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.70 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 39.7 ranged_attack_power points (3.39 DPS) | yes | Resilient Handgrips (14403, -0.75 DPS, sim-verified) [world_drop]; Cartilage Shapers (273054, -0.97 DPS) [world]; Pilferer's Gloves (7358, -1.24 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 48.6 ranged_attack_power points (4.15 DPS) | yes | Highlander's Mail Girdle (20120, -0.68 DPS) [vendor]; Green Leather Belt (4257, -0.81 DPS) [crafted]; Highlander's Lizardhide Girdle (20105, -2.07 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 66.6 ranged_attack_power points (5.69 DPS) | yes | Trapper's Leather Legguards (252517, -0.56 DPS, sim-verified) [crafted]; Petrolspill Leggings (9509, -1.93 DPS) [dungeon]; Troll's Bane Leggings (13114, -1.93 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 29.3 ranged_attack_power points (2.50 DPS) | yes | Feet of the Lynx (1121, -0.35 DPS) [world_drop]; Insignia Boots (4055, -0.35 DPS) [world_drop]; Acidic Walkers (9454, -1.00 DPS, sim-verified) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 26.6 ranged_attack_power points (2.27 DPS) | yes | Runemender's Seal (273031, -0.54 DPS) [world]; Ring of Precision (1491, -0.66 DPS) [dungeon]; Snake Hoop (6750, -1.82 DPS, sim-verified) [quest] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Runemender's Seal (273031, -0.23 DPS) [world]; Ring of Precision (1491, -0.35 DPS) [dungeon]; Snake Hoop (6750, -0.39 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.02 DPS) [world]; Kam's Walking Stick (2280, -0.31 DPS) [dungeon]; Claw of the Shadowmancer (2912, -1.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironweaver (13137, -0.54 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.81 DPS) [world_drop]; Golemsight Long Gun (273029, -1.03 DPS, sim-verified) [dungeon] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Sergeant Major's Cape; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Seal of Wrynn; ranged: Crystalpine Stinger

No-known-source sample (15 of 494, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 227.5. Weights run: 5.9s. Verify run: 21.5s. 827 eligible items had no known source.

Race sweep: night-elf 227.5 ± 0.2, human 227.2 ± 0.2, high-order-skyborne 226.6 ± 0.2, dwarf 223.5 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.301 ± 0.019, crit=1.708 ± 0.030 per rating point (14 rating = 1%, 23.916 per %), hit=2.097 ± 0.055 per rating point (10 rating = 1%, 20.967 per %), melee_haste=5.738 ± 1.136, intellect=1.448 ± 0.017, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightscape Headband (8176, -0.49 DPS) [crafted]; Guard's Chain Helm (250499, -0.49 DPS) [crafted]; Warden's Wizard Hat (14604, -4.88 DPS, sim-verified) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 36.3 ranged_attack_power points (4.02 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.73 DPS) [world]; Kaleidoscope Chain (13084, -1.92 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -2.47 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.3 ranged_attack_power points (4.02 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.37 DPS) [dungeon]; Flintrock Shoulders (7755, -0.37 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.23 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Parachute Cloak (10518, -3.71 DPS, sim-verified) [crafted] |
| chest | Tough Scorpid Breastplate (8203) | Leatherworking [crafted] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Wolffear Harness (13110, -0.39 DPS) [world_drop]; Garb of Florid Feathers (273026, -0.91 DPS) [dungeon]; Quillward Harness (10583, -2.03 DPS, sim-verified) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 31.8 ranged_attack_power points (3.52 DPS) | yes | Imperial Leather Bracers (4061, -0.60 DPS) [dungeon]; Dusky Bracers (7378, -0.60 DPS) [crafted]; Green Leather Bracers (4259, -0.90 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.32 DPS) [crafted]; Aurora Gloves (4042, -0.89 DPS) [world_drop]; Dragonscale Gauntlets (8347, -3.42 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 38.4 ranged_attack_power points (4.25 DPS) | yes | Highlander's Chain Girdle (20089, -0.72 DPS) [rep]; Highlander's Leather Girdle (20116, -0.93 DPS) [rep]; Scorpashi Sash (14652, -3.37 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 69.3 ranged_attack_power points (7.67 DPS) | yes | Triprunner Dungarees (9624, -1.92 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.56 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.56 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 46.4 ranged_attack_power points (5.14 DPS) | yes | Swampwalker Boots (2276, -0.39 DPS) [world_drop]; Imperial Leather Boots (6431, -1.12 DPS) [dungeon]; Worn Running Boots (9398, -4.03 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 33.0 ranged_attack_power points (3.65 DPS) | yes | Protector's Band (19515, -0.73 DPS) [rep]; Field Researcher's Loop (281634, -1.10 DPS) [quest]; Falcon's Hook (7552, -3.42 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19515, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.37 DPS) [quest]; Falcon's Hook (7552, -0.65 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.71 DPS) [world]; Steel Spear (250605, -0.83 DPS) [crafted]; Claw of the Shadowmancer (2912, -10.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.25 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.47 DPS) [dungeon]; The Silencer (13138, -4.17 DPS, sim-verified) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Tough Scorpid Breastplate; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 827, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 326.7. Weights run: 6.1s. Verify run: 24.8s. 1063 eligible items had no known source.

Race sweep: night-elf 326.7 ± 0.2, human 325.7 ± 0.2, high-order-skyborne 325.2 ± 0.2, dwarf 321.7 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.217 ± 0.019, crit=2.025 ± 0.035 per rating point (14 rating = 1%, 28.347 per %), hit=2.524 ± 0.092 per rating point (10 rating = 1%, 25.240 per %), melee_haste=10.648 ± 1.118, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.31 DPS) [crafted]; Tough Scorpid Helm (8208, -5.76 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.8 ranged_attack_power points (5.27 DPS) | yes | Sentinel's Medallion (19539, -0.41 DPS) [rep] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Phytoskin Spaulders (17749, -3.69 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.41 DPS) [dungeon]; Blisterbane Wrap (12552, -1.20 DPS, sim-verified) [dungeon]; Blackmetal Cape (9512, -1.22 DPS) [dungeon] |
| chest | Knight's Chain Armor (220828) | PvP rank 6 · Knight · Alliance [vendor] | 88.7 ranged_attack_power points (11.17 DPS) | yes | Quillward Harness (10583, -3.48 DPS) [dungeon]; Charred Leather Tunic (19127, -3.88 DPS) [quest]; Blazewind Breastplate (11193, -6.88 DPS, sim-verified) [quest] |
| wrist | Stalker's Mail Bracers (252570) | Leatherworking [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodlust Bracelets (14807, -0.10 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -2.23 DPS, sim-verified) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -1.04 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.37 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -6.13 DPS, sim-verified) [vendor] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 57.1 ranged_attack_power points (7.20 DPS) | yes | Highlander's Leather Girdle (20115, -1.11 DPS) [rep]; Sagebrush Girdle (17778, -1.37 DPS) [quest]; Highlander's Chain Girdle (20088, -6.29 DPS, sim-verified) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 99.3 ranged_attack_power points (12.50 DPS) | yes | Basilisk Hide Pants (1718, -3.99 DPS) [world_drop]; Tough Scorpid Leggings (8206, -4.09 DPS) [crafted]; Knight's Chain Legplates (220832, -6.97 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.4 ranged_attack_power points (8.61 DPS) | yes | Elven Chain Boots (13125, -1.32 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.97 DPS) [crafted]; Albino Crocscale Boots (17728, -2.98 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.2 ranged_attack_power points (5.70 DPS) | yes | Ring of the Underwood (2951, -1.65 DPS) [world_drop]; Falcon's Hook (7552, -2.05 DPS) [dungeon]; Protector's Band (19516, -2.05 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.0 ranged_attack_power points (5.67 DPS) | yes | Ring of the Underwood (2951, -1.53 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -2.03 DPS) [dungeon]; Protector's Band (19516, -2.03 DPS) [rep] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.91 DPS, sim-verified) [crafted] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Eyegouger (9480, -0.22 DPS) [dungeon]; Staff of Jordan (873, -1.72 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.40 DPS) [dungeon]; Hurricane (2824, -1.35 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -2.58 DPS) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Stalker's Mail Bracers; hands: Stalker's Mail Gauntlets; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Thunderbrew's Boot Flask; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1063, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 483.1. Weights run: 6.0s. Verify run: 63.9s. 2311 eligible items had no known source.

Race sweep: night-elf 483.1 ± 0.3, high-order-skyborne 482.5 ± 0.3, human 481.2 ± 0.3, dwarf 472.8 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.131 ± 0.017, crit=2.245 ± 0.035 per rating point (14 rating = 1%, 31.436 per %), hit=not significant (0.008 ± 0.014) per rating point (10 rating = 1%, 0.080 per %), melee_haste=13.182 ± 1.298, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 143.2 ranged_attack_power points (21.38 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.19 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -1.95 DPS) [pvp] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.4 ranged_attack_power points (8.58 DPS) | yes | Sentinel's Medallion (19538, -1.56 DPS) [rep]; Skibi's Pendant (13089, -2.50 DPS) [world_drop]; Sunprism Pendant (20649, -2.97 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 116.0 ranged_attack_power points (17.32 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -3.34 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -3.34 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -2.52 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.59 DPS) [rep]; Dark Phantom Cape (13122, -2.99 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 147.4 ranged_attack_power points (22.01 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, +0.00 DPS) [vendor]; Field Marshal's Chain Armor (231563, -1.82 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.5 ranged_attack_power points (8.89 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.94 DPS) [pvp]; Windtalker's Wristguards (19582, -1.76 DPS) [rep] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 92.1 ranged_attack_power points (13.75 DPS) | yes | Marshal's Chain Girdle (16464, -2.08 DPS) [pvp]; Marksman's Girdle (22232, -3.92 DPS) [dungeon]; Highlander's Chain Girdle (20043, -3.98 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 172.5 ranged_attack_power points (25.77 DPS) | yes | Marshal's Chain Legguards (231577, -3.01 DPS) [pvp]; Sentinel's Leather Pants (237818, -3.75 DPS) [vendor]; Marshal's Chain Legplates (231558, -5.58 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, +0.00 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.25 DPS) [quest]; Painweaver Band (13098, -0.31 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, +0.00 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.01 DPS) [quest]; Painweaver Band (13098, -0.07 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Devilsaur Tooth (19992, +0.00 DPS) [quest] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Hand of Justice (11815, -0.60 DPS) [dungeon]; Eye of the Beast (13968, -2.05 DPS) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -2.01 DPS) [pvp]; Peacemaker (18725, -2.50 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Grand Marshal's Hand Cannon (234587, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Trueaim Gauntlets; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2311, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1102.0. Weights run: 4.1s. Verify run: 50.2s. 2311 eligible items had no known source.

Race sweep: high-order-skyborne 1103.5 ± 0.8, night-elf 1102.0 ± 0.8, human 1095.7 ± 0.8, dwarf 1089.3 ± 0.8 (dps); high-order-skyborne beats the guide race night-elf by 1.6 (0.14%) ± 1.2, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.243 ± 0.022, crit=2.576 ± 0.047 per rating point (14 rating = 1%, 36.067 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.086 ± 1.800, intellect=1.209 ± 0.006, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 150.1 ranged_attack_power points (41.95 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -2.35 DPS) [vendor]; Beaststalker's Cap (16677, -24.17 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 62.1 ranged_attack_power points (17.35 DPS) | yes | Sentinel's Medallion (19538, -3.58 DPS, sim-verified) [rep]; Skibi's Pendant (13089, -5.56 DPS) [world_drop]; Sunprism Pendant (20649, -6.47 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 123.6 ranged_attack_power points (34.55 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -7.61 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -7.61 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -4.87 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -5.15 DPS) [rep]; Dark Phantom Cape (13122, -5.59 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 156.1 ranged_attack_power points (43.61 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -30.50 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 61.6 ranged_attack_power points (17.22 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.81 DPS) [pvp]; Beaststalker's Bindings (16681, -3.63 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Gauntlets of Deftness (22410, -4.32 DPS) [dungeon] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 94.8 ranged_attack_power points (26.48 DPS) | yes | Marshal's Chain Girdle (16464, -4.09 DPS) [pvp]; Highlander's Leather Girdle (20045, -6.90 DPS) [rep]; Highlander's Chain Girdle (20043, -8.98 DPS, sim-verified) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 185.6 ranged_attack_power points (51.88 DPS) | yes | Marshal's Chain Legguards (231577, -6.93 DPS) [pvp]; Marshal's Chain Legplates (231558, -12.28 DPS) [vendor]; Sentinel's Leather Pants (237818, -21.85 DPS, sim-verified) [vendor] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 71.7 ranged_attack_power points (20.05 DPS) | yes | Timmy's Galoshes (13402, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.56 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.30 DPS) [quest]; Naglering (11669, -28.86 DPS, sim-verified) [dungeon] |
| finger2 | Innervating Band (18701) | Scholomance: Scholomance Handler [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.16 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.91 DPS) [quest]; Cutthroat's Signet (272408, -0.91 DPS) [vendor] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+19.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Hand of Justice (11815, -1.12 DPS) [dungeon]; Eye of the Beast (13968, -3.75 DPS) [quest] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Smolderweb's Eye (13213, -8.28 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -1.66 DPS) [pvp]; Electrified Dagger (19100, -19.32 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bloodseeker (19107, -41.56 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Beaststalker's Boots; finger1: Don Julio's Band; finger2: Innervating Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2311, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (troll, 0000000000000000-0050150000000000-000000000000000000)

Set DPS (verified): 93.0. Weights run: 5.9s. Verify run: 11.2s. 294 eligible items had no known source.

Race sweep: windshaper-skyborne 94.5 ± 0.1, orc 93.6 ± 0.1, troll 93.0 ± 0.1, tauren 90.6 ± 0.1 (dps); windshaper-skyborne beats the guide race troll by 1.5 (1.62%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.215 ± 0.011, crit=0.708 ± 0.012 per rating point (14 rating = 1%, 9.916 per %), hit=1.102 ± 0.036 per rating point (10 rating = 1%, 11.023 per %), melee_haste=3.903 ± 0.605, intellect=3.110 ± 0.026, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.7 ranged_attack_power points (2.35 DPS) | yes | Trapper's Leather Hood (252505, -0.08 DPS) [crafted]; Shadow Goggles (4373, -0.93 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 19.3 ranged_attack_power points (1.76 DPS) | yes | Erudite's Amulet (277204, -1.17 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 28.0 ranged_attack_power points (2.56 DPS) | yes | Serpent's Shoulders (5404, -0.89 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -1.42 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.28 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.59 DPS) [world_drop]; Bristlebark Cape (14571, -0.59 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 31.8 ranged_attack_power points (2.91 DPS) | yes | Mystic's Robe (14371, -0.04 DPS) [world_drop]; Mystic's Wrap (14369, -0.33 DPS, sim-verified) [world_drop]; Brawler's Leather Armor (252490, -0.85 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 16.1 ranged_attack_power points (1.47 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Owl Bracers (4796, -0.27 DPS, sim-verified) [vendor]; Wolf Bracers (4794, -0.29 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Magefist Gloves (12977, -0.34 DPS) [world_drop]; Trapper's Leather Gloves (252495, -2.56 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 24.9 ranged_attack_power points (2.27 DPS) | yes | Trapper's Leather Belt (252429, -0.41 DPS, sim-verified) [crafted]; Blackened Defias Belt (10403, -0.63 DPS) [dungeon]; Deviate Scale Belt (6468, -0.80 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 41.4 ranged_attack_power points (3.78 DPS) | yes | Leggings of the Fang (10410, -0.48 DPS, sim-verified) [dungeon]; Brawler's Leather Pants (252500, -1.14 DPS) [crafted]; Bluegill Breeches (3022, -1.43 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Trapper's Leather Boots (252440, -2.06 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 25.1 ranged_attack_power points (2.29 DPS) | yes | Legionnaire's Band (20429, -1.12 DPS) [rep]; Bounty Hunter's Ring (5351, -1.41 DPS) [quest]; Volcanic Rock Ring (12053, -1.44 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 15.5 ranged_attack_power points (1.42 DPS) | yes | Legionnaire's Band (20429, -0.41 DPS, sim-verified) [rep]; Bounty Hunter's Ring (5351, -0.54 DPS) [quest]; Volcanic Rock Ring (12053, -0.57 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 44.5 ranged_attack_power points (4.06 DPS) | yes | Crescent Staff (6505, -0.02 DPS) [quest]; Staff of Orgrimmar (15444, -1.17 DPS) [quest]; Bronze Dory (250603, -1.17 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (16.40 DPS) | yes | Venomstrike (6469, -1.01 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.20 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.96 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (troll, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 144.4. Weights run: 5.8s. Verify run: 18.5s. 493 eligible items had no known source.

Race sweep: windshaper-skyborne 145.8 ± 0.1, orc 145.1 ± 0.1, troll 144.4 ± 0.1, tauren 141.2 ± 0.1 (dps); windshaper-skyborne beats the guide race troll by 1.4 (0.95%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.148 ± 0.012, crit=1.028 ± 0.016 per rating point (14 rating = 1%, 14.396 per %), hit=1.356 ± 0.037 per rating point (10 rating = 1%, 13.562 per %), melee_haste=6.807 ± 0.662, intellect=3.386 ± 0.029, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 54.2 ranged_attack_power points (4.62 DPS) | yes | Resilient Cap (14401, -0.39 DPS, sim-verified) [world_drop]; Nightsky Cowl (4039, -1.16 DPS) [world_drop]; Insignia Cap (4052, -1.44 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 32.3 ranged_attack_power points (2.76 DPS) | yes | Kaleidoscope Chain (13084, -0.53 DPS) [world_drop]; Scout's Medallion (19537, -0.61 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Feline Mantle (3748, -0.11 DPS) [dungeon]; Mantle of Thieves (2264, -0.27 DPS) [dungeon]; Death Speaker Mantle (6685, -1.49 DPS, sim-verified) [dungeon] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hawkeye's Cloak (14593, -0.27 DPS) [world_drop]; Resilient Cape (14400, -0.42 DPS) [world_drop]; Darkspear Raider's Cloak (272078, -1.08 DPS, sim-verified) [vendor] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 79.1 ranged_attack_power points (6.75 DPS) | yes | Green Leather Armor (4255, -1.72 DPS, sim-verified) [crafted]; Resilient Tunic (14398, -2.50 DPS) [world_drop]; Resilient Robe (14405, -2.50 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 27.1 ranged_attack_power points (2.31 DPS) | yes | Nightsky Wristbands (6407, -0.34 DPS, sim-verified) [world_drop]; Insignia Bracers (6410, -0.70 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.70 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 39.7 ranged_attack_power points (3.39 DPS) | yes | Resilient Handgrips (14403, -0.78 DPS, sim-verified) [world_drop]; Cartilage Shapers (273054, -0.97 DPS) [world]; Shredder Operating Gloves (16740, -1.12 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 48.6 ranged_attack_power points (4.15 DPS) | yes | Defiler's Mail Girdle (20197, -0.68 DPS) [rep]; Green Leather Belt (4257, -0.81 DPS) [crafted]; Defiler's Lizardhide Girdle (20172, -2.05 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 66.6 ranged_attack_power points (5.69 DPS) | yes | Trapper's Leather Legguards (252517, -0.57 DPS, sim-verified) [crafted]; Petrolspill Leggings (9509, -1.93 DPS) [dungeon]; Troll's Bane Leggings (13114, -1.93 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 29.3 ranged_attack_power points (2.50 DPS) | yes | Insignia Boots (4055, -0.35 DPS) [world_drop]; Warsong Boots (16977, -0.35 DPS) [quest]; Acidic Walkers (9454, -0.97 DPS, sim-verified) [dungeon] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 26.6 ranged_attack_power points (2.27 DPS) | yes | Runemender's Seal (273031, -0.54 DPS) [world]; Ring of Precision (1491, -0.66 DPS) [dungeon]; Legionnaire's Band (19513, -0.66 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 23.7 ranged_attack_power points (2.02 DPS) | yes | Runemender's Seal (273031, -0.29 DPS) [world]; Ring of Precision (1491, -0.41 DPS) [dungeon]; Legionnaire's Band (19513, -0.41 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, -0.04 DPS) [dungeon]; Felbough Arm (273037, -0.06 DPS) [world]; Claw of the Shadowmancer (2912, -0.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ironweaver (13137, -0.54 DPS) [world_drop]; Golemsight Long Gun (273029, -0.70 DPS, sim-verified) [dungeon]; Harpyclaw Short Bow (13019, -0.81 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Snake Hoop; main_hand: Crescent Staff; ranged: Crystalpine Stinger

No-known-source sample (15 of 493, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (troll, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 221.1. Weights run: 5.9s. Verify run: 18.6s. 817 eligible items had no known source.

Race sweep: windshaper-skyborne 224.1 ± 0.2, orc 222.6 ± 0.2, troll 221.1 ± 0.2, tauren 217.1 ± 0.2 (dps); windshaper-skyborne beats the guide race troll by 2.9 (1.33%) ± 0.2, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.180 ± 0.016, crit=1.481 ± 0.026 per rating point (14 rating = 1%, 20.735 per %), hit=1.933 ± 0.048 per rating point (10 rating = 1%, 19.333 per %), melee_haste=7.253 ± 0.837, intellect=4.508 ± 0.035, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | 101.6 ranged_attack_power points (8.82 DPS) | yes | Thinking Cap (2624, -2.16 DPS) [world]; Miner's Hat of the Deep (9429, -2.16 DPS) [dungeon]; Papal Fez (9431, -2.16 DPS) [dungeon] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-verified (221.1 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, +0.00 DPS) [world]; Darkspear Warding Pendant (272074, -0.30 DPS) [vendor] |
| shoulder | Nightscape Shoulders (8192) | Leatherworking [crafted] | sim-verified (221.1 DPS) | yes | Sheepshear Mantle (13115, +0.00 DPS) [world_drop]; Windchaser Amice (14432, +0.00 DPS) [world_drop]; Inquisitor's Shawl (19507, +0.00 DPS) [dungeon] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-verified (221.1 DPS) | yes | Blackforge Cape (6424, +0.00 DPS) [dungeon]; Darkspear Raider's Cloak (272077, +0.00 DPS) [vendor] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (221.1 DPS) | yes | Tough Scorpid Breastplate (8203, +0.00 DPS) [crafted]; Red Mageweave Vest (10007, +0.00 DPS) [crafted]; Garb of Florid Feathers (273026, +0.00 DPS) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 49.3 ranged_attack_power points (4.28 DPS) | yes | Windchaser Cuffs (14429, -0.21 DPS) [world_drop]; Warden's Wristbands (14600, -0.32 DPS) [world_drop]; Green Leather Bracers (4259, -0.67 DPS) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.78 DPS) [crafted]; Aurora Gloves (4042, -0.94 DPS) [world_drop]; Stormcloth Gloves (10011, -1.59 DPS) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | sim-verified (221.1 DPS) | yes | Deathmage Sash (10771, +0.00 DPS) [dungeon]; Defiler's Lizardhide Girdle (20173, +0.00 DPS) [rep]; Defiler's Mail Girdle (20198, +0.00 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Trapper's Leather Legguards (252517, -0.57 DPS) [crafted]; Misplaced Pantaloons (276201, -0.71 DPS) [vendor] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 66.5 ranged_attack_power points (5.78 DPS) | yes | Briar Tredders (10582, -0.87 DPS) [dungeon]; Boots of the Maharishi (9658, -1.15 DPS) [quest]; Prowler's Leather Shoes (252465, -1.89 DPS) [crafted] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | sim-verified (221.1 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Ogremind Ring (1993, -0.02 DPS) [world_drop]; Voodoo Band (1996, -0.02 DPS) [world] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (221.1 DPS) | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Ogremind Ring (1993, +0.00 DPS) [world_drop]; Voodoo Band (1996, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crescent Staff (6505, -5.67 DPS) [quest]; Steel Spear (250605, -5.89 DPS) [crafted]; Impaling Harpoon (5200, -5.90 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skystriker Bow (13020, -1.87 DPS) [world_drop]; The Silencer (13138, -2.04 DPS) [world_drop]; Swiftwind (13038, -2.63 DPS) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 817, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (troll, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 323.3. Weights run: 6.1s. Verify run: 24.6s. 1054 eligible items had no known source.

Race sweep: windshaper-skyborne 326.4 ± 0.2, orc 325.6 ± 0.2, troll 323.3 ± 0.2, tauren 318.0 ± 0.2 (dps); windshaper-skyborne beats the guide race troll by 3.1 (0.94%) ± 0.3, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.217 ± 0.019, crit=2.025 ± 0.035 per rating point (14 rating = 1%, 28.347 per %), hit=2.524 ± 0.092 per rating point (10 rating = 1%, 25.240 per %), melee_haste=10.648 ± 1.118, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.31 DPS) [crafted]; Tough Scorpid Helm (8208, -5.65 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.8 ranged_attack_power points (5.27 DPS) | yes | Scout's Medallion (19536, -0.81 DPS) [rep]; Woven Ivy Necklace (19159, -1.62 DPS) [quest] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Phytoskin Spaulders (17749, -3.43 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.41 DPS) [dungeon]; Blackmetal Cape (9512, -1.22 DPS) [dungeon]; Blisterbane Wrap (12552, -2.14 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) | PvP rank 6 · Stone Guard · Horde [vendor] | 88.7 ranged_attack_power points (11.17 DPS) | yes | Quillward Harness (10583, -3.48 DPS) [dungeon]; Charred Leather Tunic (19127, -3.88 DPS) [quest]; Blazewind Breastplate (11193, -7.14 DPS, sim-verified) [quest] |
| wrist | Stalker's Mail Bracers (252570) | Leatherworking [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodlust Bracelets (14807, -0.10 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -2.18 DPS, sim-verified) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -1.04 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.37 DPS) [crafted]; First Sergeant's Chain Gauntlets (220830, -6.76 DPS, sim-verified) [vendor] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 57.1 ranged_attack_power points (7.20 DPS) | yes | Defiler's Leather Girdle (20193, -1.11 DPS) [rep]; Sagebrush Girdle (17778, -1.37 DPS) [quest]; Defiler's Chain Girdle (20151, -6.17 DPS, sim-verified) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 99.3 ranged_attack_power points (12.50 DPS) | yes | Basilisk Hide Pants (1718, -3.99 DPS) [world_drop]; Tough Scorpid Leggings (8206, -4.09 DPS) [crafted]; Stone Guard's Chain Legplates (220833, -6.05 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.4 ranged_attack_power points (8.61 DPS) | yes | Elven Chain Boots (13125, -1.32 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.97 DPS) [crafted]; Albino Crocscale Boots (17728, -3.89 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.2 ranged_attack_power points (5.70 DPS) | yes | Ring of the Underwood (2951, -1.65 DPS) [world_drop]; Falcon's Hook (7552, -2.05 DPS) [dungeon]; Legionnaire's Band (19511, -2.05 DPS) [rep] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.0 ranged_attack_power points (5.67 DPS) | yes | Ring of the Underwood (2951, -1.61 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -2.03 DPS) [dungeon]; Legionnaire's Band (19511, -2.03 DPS) [rep] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+7.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.13 DPS) [crafted]; Devilsaur Tooth (19992, -1.15 DPS) [quest] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Eyegouger (9480, -0.22 DPS) [dungeon]; Staff of Jordan (873, -1.72 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.40 DPS) [dungeon]; Hurricane (2824, -1.60 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -2.58 DPS) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Stalker's Mail Bracers; hands: Stalker's Mail Gauntlets; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1054, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (troll, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 474.1. Weights run: 6.0s. Verify run: 64.9s. 2310 eligible items had no known source.

Race sweep: windshaper-skyborne 476.6 ± 0.3, orc 475.8 ± 0.3, troll 474.1 ± 0.3, tauren 465.6 ± 0.3 (dps); windshaper-skyborne beats the guide race troll by 2.4 (0.51%) ± 0.4, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.131 ± 0.017, crit=2.245 ± 0.035 per rating point (14 rating = 1%, 31.436 per %), hit=not significant (0.008 ± 0.014) per rating point (10 rating = 1%, 0.080 per %), melee_haste=13.182 ± 1.298, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 143.2 ranged_attack_power points (21.38 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, -1.19 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.4 ranged_attack_power points (8.58 DPS) | yes | Scout's Medallion (19534, -1.56 DPS) [rep]; Skibi's Pendant (13089, -2.50 DPS) [world_drop]; Sunprism Pendant (20649, -2.97 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 116.0 ranged_attack_power points (17.32 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -3.34 DPS) [pvp]; Champion's Chain Shoulders (227069, -4.21 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -2.52 DPS) [dungeon]; Deathguard's Cloak (20068, -2.59 DPS) [rep]; Dark Phantom Cape (13122, -2.99 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 147.4 ranged_attack_power points (22.01 DPS) | yes | Warlord's Chain Chestpiece (16565, +0.00 DPS) [vendor]; Warlord's Chain Hauberk (231573, +0.00 DPS) [vendor]; Warlord's Chain Armor (231566, -1.82 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.5 ranged_attack_power points (8.89 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.94 DPS) [pvp]; Windtalker's Wristguards (19582, -1.76 DPS) [rep] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 92.1 ranged_attack_power points (13.75 DPS) | yes | General's Chain Girdle (16572, -2.08 DPS) [pvp]; Marksman's Girdle (22232, -3.92 DPS) [dungeon]; Defiler's Chain Girdle (20150, -3.98 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 141.1 ranged_attack_power points (21.08 DPS) | yes | General's Chain Legguards (231574, +0.00 DPS) [pvp]; General's Chain Legplates (231567, -0.88 DPS) [vendor]; Legionnaire's Chain Legguards (227073, -3.12 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beaststalker's Boots (16675, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, +0.00 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.25 DPS) [quest]; Painweaver Band (13098, -0.31 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, +0.00 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.01 DPS) [quest]; Painweaver Band (13098, -0.07 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Devilsaur Tooth (19992, +0.00 DPS) [quest] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.29 DPS) [quest]; Hand of Justice (11815, -0.60 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -2.01 DPS) [pvp]; Peacemaker (18725, -2.50 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; High Warlord's Street Sweeper (234561, +0.00 DPS) [pvp] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2310, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (troll, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1082.3. Weights run: 4.1s. Verify run: 51.4s. 2310 eligible items had no known source.

Race sweep: windshaper-skyborne 1092.9 ± 0.8, troll 1082.3 ± 0.8, orc 1081.8 ± 0.8, tauren 1067.8 ± 0.8 (dps); windshaper-skyborne beats the guide race troll by 10.6 (0.98%) ± 1.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.243 ± 0.022, crit=2.576 ± 0.047 per rating point (14 rating = 1%, 36.067 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.086 ± 1.800, intellect=1.209 ± 0.006, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 150.1 ranged_attack_power points (41.95 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -22.90 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 62.1 ranged_attack_power points (17.35 DPS) | yes | Scout's Medallion (19534, -4.28 DPS, sim-verified) [rep]; Skibi's Pendant (13089, -5.56 DPS) [world_drop]; Sunprism Pendant (20649, -6.47 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 123.6 ranged_attack_power points (34.55 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -7.61 DPS) [pvp]; Champion's Chain Shoulders (227069, -8.16 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -4.84 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -5.15 DPS) [rep]; Dark Phantom Cape (13122, -5.59 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 156.1 ranged_attack_power points (43.61 DPS) | yes | Warlord's Chain Chestpiece (16565, +0.00 DPS) [vendor]; Warlord's Chain Hauberk (231573, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -28.60 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 61.6 ranged_attack_power points (17.22 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.81 DPS) [pvp]; Beaststalker's Bindings (16681, -3.63 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 94.8 ranged_attack_power points (26.48 DPS) | yes | General's Chain Girdle (16572, -4.09 DPS) [pvp]; Defiler's Leather Girdle (20190, -6.90 DPS) [rep]; Defiler's Chain Girdle (20150, -8.25 DPS, sim-verified) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 149.6 ranged_attack_power points (41.80 DPS) | yes | General's Chain Legguards (231574, +0.00 DPS) [pvp]; General's Chain Legplates (231567, -2.20 DPS) [vendor]; Outrider's Leather Pants (22740, -19.18 DPS, sim-verified) [rep] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 71.7 ranged_attack_power points (20.05 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Timmy's Galoshes (13402, -0.35 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.56 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.30 DPS) [quest]; Naglering (11669, -27.09 DPS, sim-verified) [dungeon] |
| finger2 | Innervating Band (18701) | Scholomance: Scholomance Handler [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.16 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.91 DPS) [quest]; Cutthroat's Signet (272408, -0.91 DPS) [vendor] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+21.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.56 DPS) [quest]; Hand of Justice (11815, -1.12 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -7.79 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -1.66 DPS) [pvp]; Glacial Blade (19099, -17.71 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -46.06 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Beaststalker's Boots; finger1: Don Julio's Band; finger2: Innervating Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2310, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

