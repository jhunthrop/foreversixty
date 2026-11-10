# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0050150000000000-000000000000000000)

Set DPS (verified): 98.3. Weights run: 4.8s. Verify run: 8.5s. 298 eligible items had no known source.

Race sweep: night-elf 98.3 ± 0.1, high-order-skyborne 96.9 ± 0.1, human 96.9 ± 0.1, dwarf 94.2 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.158 ± 0.011, crit=0.660 ± 0.011 per rating point (14 rating = 1%, 9.240 per %), hit=0.981 ± 0.033 per rating point (10 rating = 1%, 9.812 per %), melee_haste=4.134 ± 0.599, intellect=2.870 ± 0.021, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.3 ranged_attack_power points (2.55 DPS) | yes | Trapper's Leather Hood (252505, -0.23 DPS) [crafted]; Shadow Goggles (4373, -1.10 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.9 ranged_attack_power points (1.91 DPS) | yes | Erudite's Amulet (277204, -1.27 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 25.8 ranged_attack_power points (2.60 DPS) | yes | Serpent's Shoulders (5404, -0.79 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -1.45 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.32 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.64 DPS) [world_drop]; Bristlebark Cape (14571, -0.64 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 34.7 ranged_attack_power points (3.50 DPS) | yes | Trapper's Leather Armor (252491, -0.41 DPS) [crafted]; Mystic's Wrap (14369, -0.52 DPS) [world_drop]; Mystic's Robe (14371, -0.52 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.8 ranged_attack_power points (1.59 DPS) | yes | Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Owl Bracers (4796, -0.14 DPS) [vendor]; Wolf Bracers (4794, -0.32 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Riding Gloves (1304, -0.09 DPS) [quest]; Trapper's Leather Gloves (252495, -2.73 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 23.0 ranged_attack_power points (2.31 DPS) | yes | Trapper's Leather Belt (252429, -0.17 DPS) [crafted]; Blackened Defias Belt (10403, -0.50 DPS) [dungeon]; Deviate Scale Belt (6468, -0.72 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 39.9 ranged_attack_power points (4.02 DPS) | yes | Leggings of the Fang (10410, -1.16 DPS) [dungeon]; Brawler's Leather Pants (252500, -1.16 DPS) [crafted]; Duty Bound Leggings (279868, -1.23 DPS, sim-verified) [quest] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Trapper's Leather Boots (252440, -2.00 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 23.5 ranged_attack_power points (2.37 DPS) | yes | Volcanic Rock Ring (12053, -1.50 DPS) [world_drop]; Minor Channeling Ring (1449, -1.79 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.6 ranged_attack_power points (1.27 DPS) | yes | Volcanic Rock Ring (12053, -0.32 DPS, sim-verified) [world_drop]; Minor Channeling Ring (1449, -0.69 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 42.8 ranged_attack_power points (4.31 DPS) | yes | Bronze Dory (250603, -1.19 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.42 DPS) [world_drop]; Sword of the Fallen (286977, -1.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (18.09 DPS) | yes | Lil Timmy's Peashooter (13136, -1.35 DPS) [world_drop]; Venomstrike (6469, -1.53 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -4.37 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 149.5. Weights run: 5.1s. Verify run: 17.2s. 494 eligible items had no known source.

Race sweep: night-elf 149.5 ± 0.1, high-order-skyborne 147.9 ± 0.1, human 147.8 ± 0.1, dwarf 146.0 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.170 ± 0.013, crit=1.058 ± 0.017 per rating point (14 rating = 1%, 14.817 per %), hit=1.460 ± 0.034 per rating point (10 rating = 1%, 14.604 per %), melee_haste=6.721 ± 0.657, intellect=3.669 ± 0.029, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 58.7 ranged_attack_power points (4.90 DPS) | yes | Resilient Cap (14401, -1.00 DPS) [world_drop]; Nightsky Cowl (4039, -1.23 DPS) [world_drop]; Insignia Cap (4052, -1.53 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 34.0 ranged_attack_power points (2.84 DPS) | yes | Kaleidoscope Chain (13084, -0.56 DPS) [world_drop]; Sentinel's Medallion (19541, -0.72 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feline Mantle (3748, +0.00 DPS) [dungeon]; Death Speaker Mantle (6685, +0.00 DPS) [dungeon]; Resilient Mantle (14397, -0.15 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, +0.00 DPS) [pvp]; Expeditionary's Cape (271923, +0.00 DPS) [rep]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 83.6 ranged_attack_power points (6.98 DPS) | yes | Green Leather Armor (4255, -2.41 DPS) [crafted]; Resilient Tunic (14398, -2.55 DPS) [world_drop]; Resilient Robe (14405, -2.55 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 29.4 ranged_attack_power points (2.45 DPS) | yes | Nightsky Wristbands (6407, -0.61 DPS) [world_drop]; Insignia Bracers (6410, -0.86 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.86 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 42.0 ranged_attack_power points (3.51 DPS) | yes | Resilient Handgrips (14403, -0.88 DPS) [world_drop]; Cartilage Shapers (273054, -1.13 DPS) [world]; Town Clerk's Mittens (270029, -1.36 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 50.5 ranged_attack_power points (4.22 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.54 DPS) [rep]; Highlander's Mail Girdle (20120, -0.54 DPS) [vendor]; Green Leather Belt (4257, -0.79 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 70.0 ranged_attack_power points (5.85 DPS) | yes | Trapper's Leather Legguards (252517, -1.32 DPS) [crafted]; Petrolspill Leggings (9509, -2.14 DPS) [dungeon]; Troll's Bane Leggings (13114, -2.14 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 30.5 ranged_attack_power points (2.55 DPS) | yes | Acidic Walkers (9454, -0.10 DPS) [dungeon]; Frothing Slippers (254003, -0.40 DPS) [crafted]; Fiery Slippers (254005, -0.40 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 28.4 ranged_attack_power points (2.37 DPS) | yes | Snake Hoop (6750, -0.22 DPS) [quest]; Runemender's Seal (273031, -0.53 DPS) [world]; Ring of Precision (1491, -0.78 DPS) [dungeon] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Runemender's Seal (273031, -0.18 DPS) [world]; Ring of Precision (1491, -0.43 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Kam's Walking Stick (2280, -0.30 DPS) [dungeon]; Glimmering Staff (249392, -0.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.63 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.80 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Seal of Wrynn; main_hand: Felbough Arm; ranged: Crystalpine Stinger

No-known-source sample (15 of 494, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 235.1. Weights run: 4.9s. Verify run: 16.5s. 827 eligible items had no known source.

Race sweep: night-elf 235.1 ± 0.2, human 234.4 ± 0.1, high-order-skyborne 234.1 ± 0.2, dwarf 230.9 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.297 ± 0.018, crit=1.700 ± 0.029 per rating point (14 rating = 1%, 23.794 per %), hit=2.060 ± 0.054 per rating point (10 rating = 1%, 20.603 per %), melee_haste=5.621 ± 1.082, intellect=1.279 ± 0.009, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightscape Headband (8176, -0.14 DPS) [crafted]; Guard's Chain Helm (250499, -0.14 DPS) [crafted]; Warden's Wizard Hat (14604, -4.58 DPS, sim-verified) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 36.3 ranged_attack_power points (4.00 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.83 DPS) [world]; Kaleidoscope Chain (13084, -1.98 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -2.46 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.3 ranged_attack_power points (4.00 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.36 DPS) [dungeon]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.12 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.12 DPS) [world_drop]; Parachute Cloak (10518, -3.96 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 99.7 ranged_attack_power points (11.00 DPS) | yes | Quillward Harness (10583, -4.38 DPS, sim-verified) [dungeon]; Tough Scorpid Breastplate (8203, -4.56 DPS) [crafted]; Wolffear Harness (13110, -4.82 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.8 ranged_attack_power points (3.39 DPS) | yes | Imperial Leather Bracers (4061, -0.48 DPS) [dungeon]; Green Leather Bracers (4259, -0.50 DPS) [crafted]; Dusky Bracers (7378, -2.68 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.28 DPS) [crafted]; Aurora Gloves (4042, -0.87 DPS) [world_drop]; Dragonscale Gauntlets (8347, -3.40 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 37.3 ranged_attack_power points (4.12 DPS) | yes | Highlander's Chain Girdle (20089, -0.61 DPS) [rep]; Highlander's Leather Girdle (20116, -0.81 DPS) [rep]; Scorpashi Sash (14652, -3.60 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 69.2 ranged_attack_power points (7.63 DPS) | yes | Triprunner Dungarees (9624, -2.15 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.54 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.54 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 45.2 ranged_attack_power points (4.99 DPS) | yes | Swampwalker Boots (2276, -0.26 DPS) [world_drop]; Imperial Leather Boots (6431, -0.99 DPS) [dungeon]; Worn Running Boots (9398, -4.29 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 33.0 ranged_attack_power points (3.64 DPS) | yes | Protector's Band (19515, -0.73 DPS) [rep]; Field Researcher's Loop (281634, -1.09 DPS) [quest]; Falcon's Hook (7552, -3.54 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19515, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.36 DPS) [quest]; Falcon's Hook (7552, -0.80 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.32 DPS) [world]; Steel Spear (250605, -0.40 DPS) [crafted]; Claw of the Shadowmancer (2912, -13.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.24 DPS) [world_drop]; The Silencer (13138, -3.42 DPS, sim-verified) [world_drop]; Shadowforge Bushmaster (9422, -3.46 DPS) [dungeon] |

**New at 40:** head: Corpseshroud; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 827, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 331.1. Weights run: 5.1s. Verify run: 22.1s. 1062 eligible items had no known source.

Race sweep: night-elf 331.1 ± 0.2, human 330.5 ± 0.2, high-order-skyborne 329.2 ± 0.2, dwarf 326.1 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.247 ± 0.019, crit=2.078 ± 0.035 per rating point (14 rating = 1%, 29.098 per %), hit=2.515 ± 0.081 per rating point (10 rating = 1%, 25.146 per %), melee_haste=11.232 ± 1.095, intellect=1.209 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.29 DPS) [crafted]; Tough Scorpid Helm (8208, -6.36 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.2 ranged_attack_power points (5.26 DPS) | yes | Sentinel's Medallion (19540, -0.81 DPS) [rep] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Phytoskin Spaulders (17749, -4.58 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.40 DPS) [dungeon]; Blackmetal Cape (9512, -1.21 DPS) [dungeon]; Blisterbane Wrap (12552, -2.22 DPS, sim-verified) [dungeon] |
| chest | Knight's Chain Armor (220828) | PvP rank 6 · Knight · Alliance [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blazewind Breastplate (11193, -1.90 DPS) [quest]; Icemail Jerkin (1981, -3.17 DPS, sim-verified) [world_drop]; Quillward Harness (10583, -3.52 DPS) [dungeon] |
| wrist | Stalker's Mail Bracers (252570) | Leatherworking [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodlust Bracelets (14807, -0.10 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -3.32 DPS, sim-verified) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -0.97 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.36 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -8.06 DPS, sim-verified) [vendor] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 57.5 ranged_attack_power points (7.17 DPS) | yes | Highlander's Leather Girdle (20115, -1.05 DPS) [rep]; Sagebrush Girdle (17778, -1.37 DPS) [quest]; Highlander's Chain Girdle (20088, -7.53 DPS, sim-verified) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 99.8 ranged_attack_power points (12.44 DPS) | yes | Basilisk Hide Pants (1718, -3.94 DPS) [world_drop]; Tough Scorpid Leggings (8206, -4.05 DPS) [crafted]; Knight's Chain Legplates (220832, -8.81 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.9 ranged_attack_power points (8.60 DPS) | yes | Elven Chain Boots (13125, -1.31 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.98 DPS) [crafted]; Albino Crocscale Boots (17728, -3.60 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.5 ranged_attack_power points (5.67 DPS) | yes | Ring of the Underwood (2951, -1.62 DPS) [world_drop]; Falcon's Hook (7552, -2.02 DPS) [dungeon]; Protector's Band (19516, -2.02 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.1 ranged_attack_power points (5.63 DPS) | yes | Falcon's Hook (7552, -1.99 DPS) [dungeon]; Protector's Band (19516, -1.99 DPS) [rep]; Ring of the Underwood (2951, -2.80 DPS, sim-verified) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -5.03 DPS, sim-verified) [quest] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eyegouger (9480, -0.10 DPS) [dungeon]; Staff of Jordan (873, -1.69 DPS) [world_drop]; Shadowblade (2163, -2.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.44 DPS) [dungeon]; Hurricane (2824, -1.45 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -2.61 DPS) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Stalker's Mail Bracers; hands: Stalker's Mail Gauntlets; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Ankh of Life; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1062, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 484.5. Weights run: 5.0s. Verify run: 57.7s. 2287 eligible items had no known source.

Race sweep: night-elf 484.5 ± 0.3, high-order-skyborne 483.7 ± 0.3, human 482.8 ± 0.3, dwarf 474.4 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.164 ± 0.016, crit=2.325 ± 0.033 per rating point (14 rating = 1%, 32.547 per %), hit=2.561 ± 0.079 per rating point (10 rating = 1%, 25.613 per %), melee_haste=14.176 ± 0.948, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 145.0 ranged_attack_power points (20.89 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.17 DPS) [vendor]; Mask of the Unforgiven (13404, -24.75 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 58.5 ranged_attack_power points (8.44 DPS) | yes | Sentinel's Medallion (19538, -1.83 DPS, sim-verified) [rep]; Star of Mystaria (12103, -2.42 DPS) [dungeon]; Skibi's Pendant (13089, -2.51 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 118.0 ranged_attack_power points (17.00 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -0.96 DPS, sim-verified) [vendor]; Shadow Prowler's Cloak (22269, -2.43 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.54 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 175.2 ranged_attack_power points (25.25 DPS) | yes | Field Marshal's Chain Breastplate (16466, -2.97 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -2.97 DPS) [vendor]; Shadowcraft Tunic (16721, -9.20 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 60.1 ranged_attack_power points (8.66 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.91 DPS) [pvp]; Slashclaw Bracers (13211, -1.78 DPS) [dungeon] |
| hands | Trueaim Gauntlets (13255) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, -2.37 DPS, sim-verified) [rep] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 92.9 ranged_attack_power points (13.38 DPS) | yes | Marksman's Girdle (22232, -0.12 DPS) [dungeon]; Sash of the Grand Hunt (22207, -1.81 DPS) [dungeon]; Marshal's Chain Girdle (16464, -2.04 DPS) [pvp] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 201.4 ranged_attack_power points (29.03 DPS) | yes | Marshal's Chain Legguards (231577, -6.74 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.34 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.31 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 93.7 ranged_attack_power points (13.51 DPS) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -3.55 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.40 DPS) [dungeon]; Innervating Band (18701, -3.56 DPS) [dungeon]; Naglering (11669, -11.24 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, -3.23 DPS) [dungeon]; Painweaver Band (13098, -3.44 DPS, sim-verified) [dungeon]; Blackstone Ring (17713, -3.50 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Second Wind (11819, -5.75 DPS, sim-verified) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Blackhand's Breadth (13965, -2.12 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -5.38 DPS) [pvp]; Shadowblade (2163, -7.65 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bloodseeker (19107, -3.07 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Trueaim Gauntlets; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2287, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1102.8. Weights run: 3.6s. Verify run: 45.5s. 2287 eligible items had no known source.

Race sweep: high-order-skyborne 1104.8 ± 0.8, night-elf 1102.8 ± 0.8, human 1096.6 ± 0.8, dwarf 1090.3 ± 0.8 (dps); high-order-skyborne beats the guide race night-elf by 2.0 (0.18%) ± 1.2, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.250 ± 0.023, crit=2.599 ± 0.047 per rating point (14 rating = 1%, 36.388 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.482 ± 1.814, intellect=1.206 ± 0.005, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 150.5 ranged_attack_power points (41.98 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -2.33 DPS) [pvp]; Beaststalker's Cap (16677, -23.79 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 62.4 ranged_attack_power points (17.41 DPS) | yes | Sentinel's Medallion (19538, -3.94 DPS, sim-verified) [rep]; Skibi's Pendant (13089, -5.62 DPS) [world_drop]; Sunprism Pendant (20649, -6.52 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 124.1 ranged_attack_power points (34.63 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -7.69 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -7.69 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -5.01 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -5.16 DPS) [rep]; Dark Phantom Cape (13122, -5.58 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 156.6 ranged_attack_power points (43.70 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -29.63 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 61.8 ranged_attack_power points (17.23 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.81 DPS) [pvp]; Beaststalker's Bindings (16681, -3.63 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Gauntlets of Deftness (22410, -4.33 DPS) [dungeon] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 94.9 ranged_attack_power points (26.47 DPS) | yes | Marshal's Chain Girdle (16464, -4.09 DPS) [pvp]; Highlander's Leather Girdle (20045, -6.83 DPS) [rep]; Highlander's Chain Girdle (20043, -9.10 DPS, sim-verified) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 186.5 ranged_attack_power points (52.04 DPS) | yes | Marshal's Chain Legguards (231577, -7.02 DPS) [pvp]; Marshal's Chain Legplates (231558, -12.39 DPS) [vendor]; Sentinel's Leather Pants (237818, -20.61 DPS, sim-verified) [vendor] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 71.9 ranged_attack_power points (20.05 DPS) | yes | Timmy's Galoshes (13402, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Painweaver Band (13098) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, -1.01 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.92 DPS) [quest]; Naglering (11669, -15.14 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masons Fraternity Ring (9533, -1.36 DPS) [quest]; Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Innervating Band (18701, -13.58 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+20.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Hand of Justice (11815, -1.12 DPS) [dungeon]; Eye of the Beast (13968, -3.74 DPS) [quest] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Smolderweb's Eye (13213, -7.63 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -1.53 DPS) [pvp]; Electrified Dagger (19100, -18.49 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Bloodseeker (19107, -40.97 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Beaststalker's Boots; finger1: Painweaver Band; finger2: Don Julio's Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2287, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (troll, 0000000000000000-0050150000000000-000000000000000000)

Set DPS (verified): 95.2. Weights run: 4.8s. Verify run: 8.8s. 294 eligible items had no known source.

Race sweep: windshaper-skyborne 96.9 ± 0.1, orc 95.9 ± 0.1, troll 95.2 ± 0.1, tauren 92.7 ± 0.1 (dps); windshaper-skyborne beats the guide race troll by 1.6 (1.73%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.158 ± 0.011, crit=0.660 ± 0.011 per rating point (14 rating = 1%, 9.240 per %), hit=0.981 ± 0.033 per rating point (10 rating = 1%, 9.812 per %), melee_haste=4.134 ± 0.599, intellect=2.870 ± 0.021, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 25.3 ranged_attack_power points (2.55 DPS) | yes | Trapper's Leather Hood (252505, -0.31 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.10 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.9 ranged_attack_power points (1.91 DPS) | yes | Erudite's Amulet (277204, -1.28 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 25.8 ranged_attack_power points (2.60 DPS) | yes | Serpent's Shoulders (5404, -0.75 DPS, sim-verified) [dungeon]; Reinforced Woolen Shoulders (4315, -1.45 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.31 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.64 DPS) [world_drop]; Bristlebark Cape (14571, -0.64 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 30.7 ranged_attack_power points (3.10 DPS) | yes | Mystic's Wrap (14369, -0.12 DPS) [world_drop]; Mystic's Robe (14371, -0.12 DPS) [world_drop]; Brawler's Leather Armor (252490, -0.87 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.8 ranged_attack_power points (1.59 DPS) | yes | Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Owl Bracers (4796, -0.14 DPS) [vendor]; Wolf Bracers (4794, -0.32 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Magefist Gloves (12977, -0.46 DPS) [world_drop]; Trapper's Leather Gloves (252495, -2.69 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 23.0 ranged_attack_power points (2.31 DPS) | yes | Trapper's Leather Belt (252429, -0.17 DPS) [crafted]; Blackened Defias Belt (10403, -0.50 DPS) [dungeon]; Deviate Scale Belt (6468, -0.72 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 39.9 ranged_attack_power points (4.02 DPS) | yes | Leggings of the Fang (10410, -0.29 DPS, sim-verified) [dungeon]; Brawler's Leather Pants (252500, -1.16 DPS) [crafted]; Bluegill Breeches (3022, -1.48 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Trapper's Leather Boots (252440, -2.08 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 23.5 ranged_attack_power points (2.37 DPS) | yes | Legionnaire's Band (20429, -1.10 DPS) [rep]; Bounty Hunter's Ring (5351, -1.42 DPS) [quest]; Volcanic Rock Ring (12053, -1.50 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 14.4 ranged_attack_power points (1.45 DPS) | yes | Legionnaire's Band (20429, -0.17 DPS) [rep]; Bounty Hunter's Ring (5351, -0.49 DPS) [quest]; Volcanic Rock Ring (12053, -0.58 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 42.8 ranged_attack_power points (4.31 DPS) | yes | Crescent Staff (6505, -0.06 DPS) [quest]; Bronze Dory (250603, -1.16 DPS) [crafted]; Staff of Orgrimmar (15444, -1.24 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (18.09 DPS) | yes | Venomstrike (6469, -0.86 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.35 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -4.37 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (troll, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 146.4. Weights run: 5.1s. Verify run: 16.7s. 493 eligible items had no known source.

Race sweep: windshaper-skyborne 147.8 ± 0.1, orc 147.1 ± 0.1, troll 146.4 ± 0.1, tauren 143.2 ± 0.1 (dps); windshaper-skyborne beats the guide race troll by 1.5 (0.99%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.170 ± 0.013, crit=1.058 ± 0.017 per rating point (14 rating = 1%, 14.817 per %), hit=1.460 ± 0.034 per rating point (10 rating = 1%, 14.604 per %), melee_haste=6.721 ± 0.657, intellect=3.669 ± 0.029, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 58.7 ranged_attack_power points (4.90 DPS) | yes | Resilient Cap (14401, -1.00 DPS) [world_drop]; Nightsky Cowl (4039, -1.23 DPS) [world_drop]; Insignia Cap (4052, -1.53 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 34.0 ranged_attack_power points (2.84 DPS) | yes | Kaleidoscope Chain (13084, -0.56 DPS) [world_drop]; Scout's Medallion (19537, -0.72 DPS) [rep] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 40.4 ranged_attack_power points (3.37 DPS) | yes | Feline Mantle (3748, -0.39 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.46 DPS) [world_drop]; Resilient Mantle (14397, -0.61 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop]; Resilient Cape (14400, -0.28 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 83.6 ranged_attack_power points (6.98 DPS) | yes | Green Leather Armor (4255, -2.41 DPS) [crafted]; Resilient Tunic (14398, -2.55 DPS) [world_drop]; Resilient Robe (14405, -2.55 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 29.4 ranged_attack_power points (2.45 DPS) | yes | Nightsky Wristbands (6407, -0.61 DPS) [world_drop]; Insignia Bracers (6410, -0.86 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.86 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 42.0 ranged_attack_power points (3.51 DPS) | yes | Resilient Handgrips (14403, -0.88 DPS) [world_drop]; Cartilage Shapers (273054, -1.13 DPS) [world]; Shredder Operating Gloves (16740, -1.14 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 50.5 ranged_attack_power points (4.22 DPS) | yes | Defiler's Lizardhide Girdle (20172, -0.54 DPS) [rep]; Defiler's Mail Girdle (20197, -0.54 DPS) [rep]; Green Leather Belt (4257, -0.79 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 70.0 ranged_attack_power points (5.85 DPS) | yes | Trapper's Leather Legguards (252517, -1.32 DPS) [crafted]; Petrolspill Leggings (9509, -2.14 DPS) [dungeon]; Troll's Bane Leggings (13114, -2.14 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 30.5 ranged_attack_power points (2.55 DPS) | yes | Acidic Walkers (9454, -0.10 DPS) [dungeon]; Frothing Slippers (254003, -0.40 DPS) [crafted]; Fiery Slippers (254005, -0.40 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 28.4 ranged_attack_power points (2.37 DPS) | yes | Snake Hoop (6750, -0.22 DPS) [quest]; Runemender's Seal (273031, -0.53 DPS) [world]; Legionnaire's Band (19513, -0.78 DPS) [rep] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Legionnaire's Band (19513, +0.00 DPS) [rep]; Runemender's Seal (273031, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, -0.08 DPS) [dungeon]; Felbough Arm (273037, -0.13 DPS) [world]; Kam's Walking Stick (2280, -0.43 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.63 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.80 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Ring of Precision; main_hand: Crescent Staff; ranged: Crystalpine Stinger

No-known-source sample (15 of 493, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (troll, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 228.1. Weights run: 4.9s. Verify run: 16.6s. 817 eligible items had no known source.

Race sweep: windshaper-skyborne 231.3 ± 0.2, orc 229.7 ± 0.2, troll 228.1 ± 0.2, tauren 224.1 ± 0.2 (dps); windshaper-skyborne beats the guide race troll by 3.2 (1.40%) ± 0.2, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.297 ± 0.018, crit=1.700 ± 0.029 per rating point (14 rating = 1%, 23.794 per %), hit=2.060 ± 0.054 per rating point (10 rating = 1%, 20.603 per %), melee_haste=5.621 ± 1.082, intellect=1.279 ± 0.009, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightscape Headband (8176, -0.14 DPS) [crafted]; Guard's Chain Helm (250499, -0.14 DPS) [crafted]; Warden's Wizard Hat (14604, -5.77 DPS, sim-verified) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 36.3 ranged_attack_power points (4.00 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.83 DPS) [world]; Kaleidoscope Chain (13084, -1.98 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -2.46 DPS) [quest] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.3 ranged_attack_power points (4.00 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.36 DPS) [dungeon]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.36 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 99.7 ranged_attack_power points (11.00 DPS) | yes | Tough Scorpid Breastplate (8203, -4.56 DPS) [crafted]; Wolffear Harness (13110, -4.82 DPS) [world_drop]; Quillward Harness (10583, -4.86 DPS, sim-verified) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.8 ranged_attack_power points (3.39 DPS) | yes | Imperial Leather Bracers (4061, -0.48 DPS) [dungeon]; Green Leather Bracers (4259, -0.50 DPS) [crafted]; Dusky Bracers (7378, -2.93 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.28 DPS) [crafted]; Aurora Gloves (4042, -0.87 DPS) [world_drop]; Dragonscale Gauntlets (8347, -3.94 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 37.3 ranged_attack_power points (4.12 DPS) | yes | Defiler's Chain Girdle (20153, -0.61 DPS) [rep]; Defiler's Leather Girdle (20192, -0.81 DPS) [rep]; Scorpashi Sash (14652, -4.13 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 69.2 ranged_attack_power points (7.63 DPS) | yes | Triprunner Dungarees (9624, -1.59 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.54 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.54 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 45.2 ranged_attack_power points (4.99 DPS) | yes | Swampwalker Boots (2276, -0.26 DPS) [world_drop]; Imperial Leather Boots (6431, -0.99 DPS) [dungeon]; Worn Running Boots (9398, -4.23 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 33.0 ranged_attack_power points (3.64 DPS) | yes | Legionnaire's Band (19512, -0.73 DPS) [rep]; Field Researcher's Loop (281634, -1.09 DPS) [quest]; Falcon's Hook (7552, -3.62 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Band (19512, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.36 DPS) [quest]; Falcon's Hook (7552, -0.82 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.32 DPS) [world]; Steel Spear (250605, -0.40 DPS) [crafted]; Claw of the Shadowmancer (2912, -13.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.24 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.46 DPS) [dungeon]; The Silencer (13138, -4.32 DPS, sim-verified) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 817, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (troll, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 328.0. Weights run: 5.1s. Verify run: 21.6s. 1053 eligible items had no known source.

Race sweep: windshaper-skyborne 331.3 ± 0.2, orc 330.4 ± 0.2, troll 328.0 ± 0.2, tauren 322.7 ± 0.2 (dps); windshaper-skyborne beats the guide race troll by 3.3 (1.00%) ± 0.3, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.247 ± 0.019, crit=2.078 ± 0.035 per rating point (14 rating = 1%, 29.098 per %), hit=2.515 ± 0.081 per rating point (10 rating = 1%, 25.146 per %), melee_haste=11.232 ± 1.095, intellect=1.209 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.29 DPS) [crafted]; Tough Scorpid Helm (8208, -6.83 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.2 ranged_attack_power points (5.26 DPS) | yes | Scout's Medallion (19535, -0.40 DPS) [rep]; Woven Ivy Necklace (19159, -1.62 DPS) [quest] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Phytoskin Spaulders (17749, -4.88 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.40 DPS) [dungeon]; Blackmetal Cape (9512, -1.21 DPS) [dungeon]; Blisterbane Wrap (12552, -1.68 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blazewind Breastplate (11193, -1.90 DPS) [quest]; Quillward Harness (10583, -3.52 DPS) [dungeon]; Icemail Jerkin (1981, -3.74 DPS, sim-verified) [world_drop] |
| wrist | Stalker's Mail Bracers (252570) | Leatherworking [crafted] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodlust Bracelets (14807, -0.10 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -3.23 DPS, sim-verified) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -0.97 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.36 DPS) [crafted]; First Sergeant's Chain Gauntlets (220830, -7.17 DPS, sim-verified) [vendor] |
| waist | Stalker's Mail Belt (252588) | Leatherworking [crafted] | 57.5 ranged_attack_power points (7.17 DPS) | yes | Defiler's Leather Girdle (20193, -1.05 DPS) [rep]; Sagebrush Girdle (17778, -1.37 DPS) [quest]; Defiler's Chain Girdle (20151, -6.70 DPS, sim-verified) [rep] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 99.8 ranged_attack_power points (12.44 DPS) | yes | Basilisk Hide Pants (1718, -3.94 DPS) [world_drop]; Tough Scorpid Leggings (8206, -4.05 DPS) [crafted]; Stone Guard's Chain Legplates (220833, -7.22 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.9 ranged_attack_power points (8.60 DPS) | yes | Elven Chain Boots (13125, -1.31 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.98 DPS) [crafted]; Albino Crocscale Boots (17728, -3.78 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.5 ranged_attack_power points (5.67 DPS) | yes | Ring of the Underwood (2951, -1.62 DPS) [world_drop]; Falcon's Hook (7552, -2.02 DPS) [dungeon]; Legionnaire's Band (19511, -2.02 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.1 ranged_attack_power points (5.63 DPS) | yes | Ring of the Underwood (2951, -1.66 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -1.99 DPS) [dungeon]; Legionnaire's Band (19511, -1.99 DPS) [rep] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+7.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Devilsaur Tooth (19992, -1.11 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.12 DPS) [crafted] |
| main_hand | Warmonger (13052) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eyegouger (9480, -0.10 DPS) [dungeon]; Shadowblade (2163, -1.48 DPS, sim-verified) [world_drop]; Staff of Jordan (873, -1.69 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -0.44 DPS) [dungeon]; Hurricane (2824, -1.17 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -2.61 DPS) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Stalker's Mail Bracers; hands: Stalker's Mail Gauntlets; waist: Stalker's Mail Belt; legs: Infernal Trickster Leggings; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Warmonger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1053, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (troll, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 475.8. Weights run: 5.0s. Verify run: 56.1s. 2286 eligible items had no known source.

Race sweep: windshaper-skyborne 477.8 ± 0.3, orc 477.2 ± 0.3, troll 475.8 ± 0.3, tauren 467.0 ± 0.3 (dps); windshaper-skyborne beats the guide race troll by 2.1 (0.43%) ± 0.4, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.164 ± 0.016, crit=2.325 ± 0.033 per rating point (14 rating = 1%, 32.547 per %), hit=2.561 ± 0.079 per rating point (10 rating = 1%, 25.613 per %), melee_haste=14.176 ± 0.948, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 145.0 ranged_attack_power points (20.89 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -25.40 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 58.5 ranged_attack_power points (8.44 DPS) | yes | Scout's Medallion (19534, -2.03 DPS, sim-verified) [rep]; Star of Mystaria (12103, -2.42 DPS) [dungeon]; Skibi's Pendant (13089, -2.51 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 118.0 ranged_attack_power points (17.00 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.60 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -2.00 DPS) [vendor]; Shadow Prowler's Cloak (22269, -2.43 DPS) [dungeon]; Deathguard's Cloak (20068, -2.54 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 175.2 ranged_attack_power points (25.25 DPS) | yes | Warlord's Chain Chestpiece (16565, -2.97 DPS) [vendor]; Warlord's Chain Hauberk (231573, -2.97 DPS) [vendor]; Shadowcraft Tunic (16721, -9.65 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 60.1 ranged_attack_power points (8.66 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.91 DPS) [pvp]; Slashclaw Bracers (13211, -1.78 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -1.63 DPS, sim-verified) [dungeon] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 92.9 ranged_attack_power points (13.38 DPS) | yes | Sash of the Grand Hunt (22207, -1.81 DPS) [dungeon]; General's Chain Girdle (16572, -2.04 DPS) [pvp]; Marksman's Girdle (22232, -2.22 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 168.9 ranged_attack_power points (24.34 DPS) | yes | General's Chain Legguards (231574, -2.05 DPS) [pvp]; General's Chain Legplates (231567, -4.62 DPS) [vendor]; Outrider's Leather Pants (22740, -8.90 DPS, sim-verified) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 93.7 ranged_attack_power points (13.51 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -5.23 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.40 DPS) [dungeon]; Innervating Band (18701, -3.56 DPS) [dungeon]; Naglering (11669, -12.93 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, -3.23 DPS) [dungeon]; Blackstone Ring (17713, -3.50 DPS) [dungeon]; Painweaver Band (13098, -4.75 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Second Wind (11819, -4.95 DPS, sim-verified) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -2.38 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -5.38 DPS) [pvp]; Glacial Blade (19099, -9.44 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Bloodseeker (19107, -6.75 DPS, sim-verified) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Devilsaur Eye; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2286, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (troll, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1083.6. Weights run: 3.6s. Verify run: 44.9s. 2286 eligible items had no known source.

Race sweep: windshaper-skyborne 1093.9 ± 0.8, troll 1083.6 ± 0.8, orc 1083.1 ± 0.8, tauren 1069.0 ± 0.8 (dps); windshaper-skyborne beats the guide race troll by 10.3 (0.95%) ± 1.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.250 ± 0.023, crit=2.599 ± 0.047 per rating point (14 rating = 1%, 36.388 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=24.482 ± 1.814, intellect=1.206 ± 0.005, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 150.5 ranged_attack_power points (41.98 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Beaststalker's Cap (16677, -23.27 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 62.4 ranged_attack_power points (17.41 DPS) | yes | Scout's Medallion (19534, -4.08 DPS, sim-verified) [rep]; Skibi's Pendant (13089, -5.62 DPS) [world_drop]; Sunprism Pendant (20649, -6.52 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 124.1 ranged_attack_power points (34.63 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -7.69 DPS) [pvp]; Champion's Chain Shoulders (227069, -8.16 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -4.82 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -5.16 DPS) [rep]; Dark Phantom Cape (13122, -5.58 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 156.6 ranged_attack_power points (43.70 DPS) | yes | Warlord's Chain Chestpiece (16565, +0.00 DPS) [vendor]; Warlord's Chain Hauberk (231573, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -29.77 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 61.8 ranged_attack_power points (17.23 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.81 DPS) [pvp]; Beaststalker's Bindings (16681, -3.63 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 94.9 ranged_attack_power points (26.47 DPS) | yes | General's Chain Girdle (16572, -4.09 DPS) [pvp]; Defiler's Leather Girdle (20190, -6.83 DPS) [rep]; Defiler's Chain Girdle (20150, -8.24 DPS, sim-verified) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 150.1 ranged_attack_power points (41.89 DPS) | yes | General's Chain Legguards (231574, +0.00 DPS) [pvp]; General's Chain Legplates (231567, -2.24 DPS) [vendor]; Outrider's Leather Pants (22740, -20.06 DPS, sim-verified) [rep] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 71.9 ranged_attack_power points (20.05 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Timmy's Galoshes (13402, -0.37 DPS) [dungeon] |
| finger1 | Painweaver Band (13098) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Innervating Band (18701, -1.01 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.92 DPS) [quest]; Naglering (11669, -15.26 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masons Fraternity Ring (9533, -1.36 DPS) [quest]; Cutthroat's Signet (272408, -1.36 DPS) [vendor]; Innervating Band (18701, -14.04 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+21.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.56 DPS) [quest]; Hand of Justice (11815, -1.12 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -7.36 DPS, sim-verified) [quest] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -1.53 DPS) [pvp]; Glacial Blade (19099, -17.65 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Dark Iron Rifle (16004, -48.87 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Beaststalker's Boots; finger1: Painweaver Band; finger2: Don Julio's Band; trinket2: Blackhand's Breadth; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2286, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

