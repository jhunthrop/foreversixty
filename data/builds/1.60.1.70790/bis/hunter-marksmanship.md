# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0050150000000000-000000000000000000)

Set DPS (verified): 98.3. Weights run: 5.7s. Verify run: 9.9s. 268 eligible items had no known source.

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
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 23.5 ranged_attack_power points (2.37 DPS) | yes | Volcanic Rock Ring (12053, -1.50 DPS) [world_drop]; The 1 Ring (8350, -1.76 DPS) [world]; Minor Channeling Ring (1449, -1.79 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.6 ranged_attack_power points (1.27 DPS) | yes | Volcanic Rock Ring (12053, -0.32 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.67 DPS) [world]; Minor Channeling Ring (1449, -0.69 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 42.8 ranged_attack_power points (4.31 DPS) | yes | Bronze Dory (250603, -1.19 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.42 DPS) [world_drop]; Sword of the Fallen (286977, -1.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (18.09 DPS) | yes | Lil Timmy's Peashooter (13136, -1.35 DPS) [world_drop]; Venomstrike (6469, -1.53 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -4.37 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 149.5. Weights run: 5.9s. Verify run: 18.9s. 463 eligible items had no known source.

Race sweep: night-elf 149.5 ± 0.1, high-order-skyborne 147.9 ± 0.1, human 147.8 ± 0.1, dwarf 146.0 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.132 ± 0.013, crit=1.008 ± 0.018 per rating point (14 rating = 1%, 14.115 per %), hit=1.452 ± 0.038 per rating point (10 rating = 1%, 14.519 per %), melee_haste=3.723 ± 0.634, intellect=3.795 ± 0.030, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 60.7 ranged_attack_power points (5.02 DPS) | yes | Resilient Cap (14401, -1.05 DPS) [world_drop]; Nightsky Cowl (4039, -1.26 DPS) [world_drop]; Insignia Cap (4052, -1.57 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 34.8 ranged_attack_power points (2.88 DPS) | yes | Kaleidoscope Chain (13084, -0.58 DPS) [world_drop]; Sentinel's Medallion (19541, -0.80 DPS) [rep]; Seal of the Expedition (271920, -1.31 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feline Mantle (3748, +0.00 DPS) [dungeon]; Death Speaker Mantle (6685, +0.00 DPS) [dungeon]; Resilient Mantle (14397, -0.02 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, +0.00 DPS) [pvp]; Expeditionary's Cape (271923, +0.00 DPS) [rep]; Resilient Cape (14400, -0.19 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 85.1 ranged_attack_power points (7.04 DPS) | yes | Green Leather Armor (4255, -2.46 DPS) [crafted]; Resilient Tunic (14398, -2.55 DPS) [world_drop]; Resilient Robe (14405, -2.55 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 30.4 ranged_attack_power points (2.51 DPS) | yes | Nightsky Wristbands (6407, -0.63 DPS) [world_drop]; Owl Bracers (4796, -0.94 DPS) [vendor]; Stonecloth Bindings (14416, -0.94 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 42.9 ranged_attack_power points (3.55 DPS) | yes | Resilient Handgrips (14403, -0.89 DPS) [world_drop]; Cartilage Shapers (273054, -1.22 DPS) [world]; Town Clerk's Mittens (270029, -1.35 DPS) [quest] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 51.0 ranged_attack_power points (4.22 DPS) | yes | Highlander's Lizardhide Girdle (20105, -0.45 DPS) [rep]; Highlander's Mail Girdle (20120, -0.45 DPS) [vendor]; Green Leather Belt (4257, -0.78 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 71.5 ranged_attack_power points (5.92 DPS) | yes | Trapper's Leather Legguards (252517, -1.39 DPS) [crafted]; Petrolspill Leggings (9509, -2.29 DPS) [dungeon]; Troll's Bane Leggings (13114, -2.29 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 30.8 ranged_attack_power points (2.55 DPS) | yes | Acidic Walkers (9454, -0.04 DPS) [dungeon]; Frothing Slippers (254003, -0.35 DPS) [crafted]; Fiery Slippers (254005, -0.35 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 29.0 ranged_attack_power points (2.40 DPS) | yes | Azora's Will (4999, -0.83 DPS) [world_drop]; Snake Hoop (6750, -0.83 DPS) [quest]; Ring of Precision (1491, -0.85 DPS) [dungeon] |
| finger2 | Seal of Wrynn (2933) | An Audience with the King [quest] | 24.6 ranged_attack_power points (2.03 DPS) | yes | Azora's Will (4999, -0.46 DPS) [world_drop]; Snake Hoop (6750, -0.46 DPS) [quest]; Ring of Precision (1491, -0.48 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Kam's Walking Stick (2280, -0.26 DPS) [dungeon]; Glimmering Staff (249392, -0.34 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.65 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.78 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Seal of Wrynn; main_hand: Felbough Arm; ranged: Crystalpine Stinger

No-known-source sample (15 of 463, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 236.0. Weights run: 5.8s. Verify run: 21.1s. 718 eligible items had no known source.

Race sweep: night-elf 236.0 ± 0.1, human 235.1 ± 0.1, high-order-skyborne 235.1 ± 0.1, dwarf 231.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.294 ± 0.018, crit=1.695 ± 0.029 per rating point (14 rating = 1%, 23.732 per %), hit=2.035 ± 0.054 per rating point (10 rating = 1%, 20.348 per %), melee_haste=5.580 ± 1.072, intellect=1.254 ± 0.007, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightscape Headband (8176, -0.08 DPS) [crafted]; Guard's Chain Helm (250499, -0.08 DPS) [crafted]; Warden's Wizard Hat (14604, -5.88 DPS, sim-verified) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 36.2 ranged_attack_power points (4.00 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.79 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.84 DPS) [world]; Kaleidoscope Chain (13084, -1.99 DPS) [world_drop] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.2 ranged_attack_power points (4.00 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.36 DPS) [dungeon]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.10 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Parachute Cloak (10518, -3.96 DPS, sim-verified) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 99.4 ranged_attack_power points (10.97 DPS) | yes | Quillward Harness (10583, -4.38 DPS, sim-verified) [dungeon]; Tough Scorpid Breastplate (8203, -4.55 DPS) [crafted]; Wolffear Harness (13110, -4.79 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.6 ranged_attack_power points (3.37 DPS) | yes | Imperial Leather Bracers (4061, -0.47 DPS) [dungeon]; Green Leather Bracers (4259, -0.50 DPS) [crafted]; Dusky Bracers (7378, -2.68 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.28 DPS) [crafted]; Aurora Gloves (4042, -0.87 DPS) [world_drop]; Dragonscale Gauntlets (8347, -3.40 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 37.2 ranged_attack_power points (4.10 DPS) | yes | Highlander's Chain Girdle (20089, -0.60 DPS) [rep]; Highlander's Leather Girdle (20116, -0.79 DPS) [rep]; Scorpashi Sash (14652, -3.60 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 69.2 ranged_attack_power points (7.63 DPS) | yes | Triprunner Dungarees (9624, -2.15 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.54 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.54 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 45.0 ranged_attack_power points (4.97 DPS) | yes | Swampwalker Boots (2276, -0.24 DPS) [world_drop]; Imperial Leather Boots (6431, -0.97 DPS) [dungeon]; Worn Running Boots (9398, -4.29 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.9 ranged_attack_power points (3.63 DPS) | yes | Protector's Band (19515, -0.73 DPS) [rep]; Field Researcher's Loop (281634, -1.09 DPS) [quest]; Falcon's Hook (7552, -3.54 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19515, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.36 DPS) [quest]; Falcon's Hook (7552, -0.80 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.25 DPS) [world]; Steel Spear (250605, -0.31 DPS) [crafted]; Claw of the Shadowmancer (2912, -13.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.25 DPS) [world_drop]; The Silencer (13138, -3.42 DPS, sim-verified) [world_drop]; Shadowforge Bushmaster (9422, -3.46 DPS) [dungeon] |

**New at 40:** head: Corpseshroud; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 329.1. Weights run: 6.2s. Verify run: 27.4s. 900 eligible items had no known source.

Race sweep: night-elf 329.1 ± 0.3, human 328.3 ± 0.2, high-order-skyborne 326.5 ± 0.3, dwarf 324.4 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.242 ± 0.016, crit=2.051 ± 0.030 per rating point (14 rating = 1%, 28.714 per %), hit=2.356 ± 0.083 per rating point (10 rating = 1%, 23.564 per %), melee_haste=9.615 ± 1.037, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.29 DPS) [crafted]; Tough Scorpid Helm (8208, -6.33 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.1 ranged_attack_power points (5.31 DPS) | yes | Sentinel's Medallion (19540, -0.82 DPS) [rep] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Phytoskin Spaulders (17749, -3.85 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.41 DPS) [dungeon]; Blackmetal Cape (9512, -1.23 DPS) [dungeon]; Blisterbane Wrap (12552, -2.23 DPS, sim-verified) [dungeon] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Blazewind Breastplate (11193, -1.87 DPS) [quest]; Icemail Jerkin (1981, -2.61 DPS, sim-verified) [world_drop]; Quillward Harness (10583, -3.51 DPS) [dungeon] |
| wrist | Stalker's Mail Bracers (252570) | Leatherworking [crafted] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodlust Bracelets (14807, -0.10 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -2.98 DPS, sim-verified) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -1.02 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.37 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -6.54 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 63.1 ranged_attack_power points (7.96 DPS) | yes | Highlander's Chain Girdle (20088, -1.82 DPS) [rep]; Highlander's Leather Girdle (20115, -1.82 DPS) [rep]; Stalker's Mail Belt (252588, -1.91 DPS, sim-verified) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 86.2 ranged_attack_power points (10.86 DPS) | yes | Tough Scorpid Leggings (8206, -2.39 DPS) [crafted]; Imperial Red Pants (8251, -2.45 DPS) [world_drop]; Basilisk Hide Pants (1718, -6.50 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.9 ranged_attack_power points (8.68 DPS) | yes | Elven Chain Boots (13125, -1.32 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.99 DPS) [crafted]; Albino Crocscale Boots (17728, -3.38 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.4 ranged_attack_power points (5.72 DPS) | yes | Ring of the Underwood (2951, -1.63 DPS) [world_drop]; Falcon's Hook (7552, -2.04 DPS) [dungeon]; Protector's Band (19516, -2.04 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.6 ranged_attack_power points (5.49 DPS) | yes | Falcon's Hook (7552, -1.81 DPS) [dungeon]; Protector's Band (19516, -1.81 DPS) [rep]; Ring of the Underwood (2951, -2.48 DPS, sim-verified) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -4.62 DPS, sim-verified) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -2.00 DPS, sim-verified) [world_drop] |
| main_hand | Eyegouger (9480) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.49 DPS) [world_drop]; Staff of Jordan (873, -1.89 DPS) [world_drop]; Claw of the Shadowmancer (2912, -6.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Megashot Rifle (17717, -4.38 DPS) [dungeon]; Dark Iron Rifle (16004, -5.13 DPS) [crafted]; Gryphonwing Long Bow (13022, -7.97 DPS, sim-verified) [dungeon] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Stalker's Mail Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Thunderbrew's Boot Flask; main_hand: Eyegouger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 900, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 468.7. Weights run: 6.0s. Verify run: 67.3s. 2028 eligible items had no known source.

Race sweep: night-elf 468.7 ± 0.3, human 467.5 ± 0.3, high-order-skyborne 463.9 ± 0.3, dwarf 463.7 ± 0.3 (dps); the guide race night-elf is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.193 ± 0.018, crit=2.468 ± 0.038 per rating point (14 rating = 1%, 34.551 per %), hit=2.847 ± 0.101 per rating point (10 rating = 1%, 28.475 per %), melee_haste=not significant (3.616 ± 1.419), intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 147.6 ranged_attack_power points (20.60 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.15 DPS) [vendor]; Outlaw's Collar (279253, -27.11 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.89 DPS) [world_drop]; Talisman of Evasion (13177, -0.89 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 120.8 ranged_attack_power points (16.86 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.67 DPS, sim-verified) [vendor]; Shadow Prowler's Cloak (22269, -1.90 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.50 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 181.2 ranged_attack_power points (25.29 DPS) | yes | Field Marshal's Chain Breastplate (16466, -3.28 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -3.28 DPS) [vendor]; Savage Gladiator Chain (11726, -12.64 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 60.7 ranged_attack_power points (8.47 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.89 DPS) [pvp]; Slashclaw Bracers (13211, -1.37 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -0.07 DPS) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Girdle (16464, -2.00 DPS) [pvp]; Highlander's Chain Girdle (20043, -3.49 DPS) [rep]; Marksman's Girdle (22232, -4.91 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 209.3 ranged_attack_power points (29.22 DPS) | yes | Marshal's Chain Legguards (231577, -7.22 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.54 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.77 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.70 DPS) [dungeon]; Blackstone Ring (17713, -3.99 DPS) [dungeon]; Naglering (11669, -11.17 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.45 DPS) [dungeon]; Innervating Band (18701, -3.53 DPS) [dungeon]; Painweaver Band (13098, -3.73 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -0.79 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -5.11 DPS) [pvp]; Electrified Dagger (19100, -7.55 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Hyper Deluxe Sniper Rifle Mk XVII (279273) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Eaglehorn Long Bow (13023, -2.85 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Burst of Knowledge; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Hyper Deluxe Sniper Rifle Mk XVII

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1083.6. Weights run: 4.0s. Verify run: 51.2s. 2028 eligible items had no known source.

Race sweep: high-order-skyborne 1086.0 ± 0.8, night-elf 1083.6 ± 0.8, human 1077.7 ± 0.8, dwarf 1072.7 ± 0.8 (dps); high-order-skyborne beats the guide race night-elf by 2.4 (0.22%) ± 1.1, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.303 ± 0.026, crit=2.698 ± 0.053 per rating point (14 rating = 1%, 37.776 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=22.566 ± 2.289, intellect=1.196 ± 0.008, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 152.6 ranged_attack_power points (41.18 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -1.84 DPS) [pvp]; Outlaw's Collar (279253, -15.89 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -1.78 DPS) [world_drop]; Talisman of Evasion (13177, -1.78 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 127.0 ranged_attack_power points (34.26 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -7.86 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, -7.86 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -3.97 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -5.14 DPS) [rep]; Dark Phantom Cape (13122, -5.40 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 160.0 ranged_attack_power points (43.17 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -26.26 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 62.8 ranged_attack_power points (16.93 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.78 DPS) [pvp]; Beaststalker's Bindings (16681, -3.57 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -0.83 DPS) [crafted] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 96.0 ranged_attack_power points (25.91 DPS) | yes | Marshal's Chain Girdle (16464, -4.06 DPS) [pvp]; Highlander's Leather Girdle (20045, -6.54 DPS) [rep]; Highlander's Chain Girdle (20043, -8.68 DPS, sim-verified) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 191.2 ranged_attack_power points (51.58 DPS) | yes | Marshal's Chain Legguards (231577, -7.21 DPS) [pvp]; Marshal's Chain Legplates (231558, -12.65 DPS) [vendor]; Sentinel's Leather Pants (237818, -17.37 DPS, sim-verified) [vendor] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 72.9 ranged_attack_power points (19.69 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Innervating Band (18701, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -1.49 DPS) [vendor]; Innervating Band (18701, -10.68 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Heart of Wyrmthalak (22321, -6.28 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -1.02 DPS) [pvp]; Electrified Dagger (19100, -17.18 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -20.00 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Beaststalker's Boots; finger2: Don Julio's Band; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (troll, 0000000000000000-0050150000000000-000000000000000000)

Set DPS (verified): 95.2. Weights run: 5.7s. Verify run: 10.4s. 264 eligible items had no known source.

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
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 42.8 ranged_attack_power points (4.31 DPS) | yes | Crescent Staff (6505, -0.06 DPS) [quest]; Staff of Orgrimmar (15444, -0.67 DPS) [quest]; Bronze Dory (250603, -1.16 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.5 ranged_attack_power points (18.09 DPS) | yes | Venomstrike (6469, -0.86 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.35 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -4.37 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Keller's Girdle; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (troll, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 146.4. Weights run: 5.9s. Verify run: 18.9s. 462 eligible items had no known source.

Race sweep: windshaper-skyborne 147.8 ± 0.1, orc 147.1 ± 0.1, troll 146.4 ± 0.1, tauren 143.2 ± 0.1 (dps); windshaper-skyborne beats the guide race troll by 1.5 (0.99%) ± 0.1, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.132 ± 0.013, crit=1.008 ± 0.018 per rating point (14 rating = 1%, 14.115 per %), hit=1.452 ± 0.038 per rating point (10 rating = 1%, 14.519 per %), melee_haste=3.723 ± 0.634, intellect=3.795 ± 0.030, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 60.7 ranged_attack_power points (5.02 DPS) | yes | Resilient Cap (14401, -1.05 DPS) [world_drop]; Nightsky Cowl (4039, -1.26 DPS) [world_drop]; Insignia Cap (4052, -1.57 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 34.8 ranged_attack_power points (2.88 DPS) | yes | Kaleidoscope Chain (13084, -0.58 DPS) [world_drop]; Scout's Medallion (19537, -0.80 DPS) [rep]; Darkspear Warding Pendant (272075, -1.31 DPS) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 41.8 ranged_attack_power points (3.45 DPS) | yes | Feline Mantle (3748, -0.42 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.60 DPS) [world_drop]; Resilient Mantle (14397, -0.63 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Resilient Cape (14400, -0.19 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.26 DPS) [world_drop] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 85.1 ranged_attack_power points (7.04 DPS) | yes | Green Leather Armor (4255, -2.46 DPS) [crafted]; Resilient Tunic (14398, -2.55 DPS) [world_drop]; Resilient Robe (14405, -2.55 DPS) [world_drop] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 30.4 ranged_attack_power points (2.51 DPS) | yes | Nightsky Wristbands (6407, -0.63 DPS) [world_drop]; Owl Bracers (4796, -0.94 DPS) [vendor]; Stonecloth Bindings (14416, -0.94 DPS) [world_drop] |
| hands | Hotshot Pilot's Gloves (9491) | Gnomeregan [dungeon] | 42.9 ranged_attack_power points (3.55 DPS) | yes | Resilient Handgrips (14403, -0.89 DPS) [world_drop]; Shredder Operating Gloves (16740, -1.15 DPS) [quest]; Cartilage Shapers (273054, -1.22 DPS) [world] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 51.0 ranged_attack_power points (4.22 DPS) | yes | Defiler's Lizardhide Girdle (20172, -0.45 DPS) [rep]; Defiler's Mail Girdle (20197, -0.45 DPS) [rep]; Green Leather Belt (4257, -0.78 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 71.5 ranged_attack_power points (5.92 DPS) | yes | Trapper's Leather Legguards (252517, -1.39 DPS) [crafted]; Petrolspill Leggings (9509, -2.29 DPS) [dungeon]; Troll's Bane Leggings (13114, -2.29 DPS) [world_drop] |
| feet | Trapper's Leather Boots (252440) | Leatherworking [crafted] | 30.8 ranged_attack_power points (2.55 DPS) | yes | Acidic Walkers (9454, -0.04 DPS) [dungeon]; Frothing Slippers (254003, -0.35 DPS) [crafted]; Fiery Slippers (254005, -0.35 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 29.0 ranged_attack_power points (2.40 DPS) | yes | Azora's Will (4999, -0.83 DPS) [world_drop]; Snake Hoop (6750, -0.83 DPS) [quest]; Loop of Sacrifice (281673, -0.83 DPS) [quest] |
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Azora's Will (4999, +0.00 DPS) [world_drop]; Snake Hoop (6750, +0.00 DPS) [quest]; Loop of Sacrifice (281673, +0.00 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, -0.11 DPS) [dungeon]; Felbough Arm (273037, -0.22 DPS) [world]; Kam's Walking Stick (2280, -0.48 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.65 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.78 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Ring of Precision; main_hand: Crescent Staff; ranged: Crystalpine Stinger

No-known-source sample (15 of 462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (troll, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 229.5. Weights run: 5.8s. Verify run: 21.6s. 708 eligible items had no known source.

Race sweep: windshaper-skyborne 232.6 ± 0.2, orc 230.9 ± 0.2, troll 229.5 ± 0.2, tauren 225.1 ± 0.2 (dps); windshaper-skyborne beats the guide race troll by 3.0 (1.32%) ± 0.2, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.294 ± 0.018, crit=1.695 ± 0.029 per rating point (14 rating = 1%, 23.732 per %), hit=2.035 ± 0.054 per rating point (10 rating = 1%, 20.348 per %), melee_haste=5.580 ± 1.072, intellect=1.254 ± 0.007, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Corpseshroud (10574) | Razorfen Downs [dungeon] | sim-verified (+7.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Nightscape Headband (8176, -0.08 DPS) [crafted]; Guard's Chain Helm (250499, -0.08 DPS) [crafted]; Warden's Wizard Hat (14604, -7.00 DPS, sim-verified) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 36.2 ranged_attack_power points (4.00 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.79 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.84 DPS) [world]; Kaleidoscope Chain (13084, -1.99 DPS) [world_drop] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 36.2 ranged_attack_power points (4.00 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.36 DPS) [dungeon]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.36 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 99.4 ranged_attack_power points (10.97 DPS) | yes | Tough Scorpid Breastplate (8203, -4.55 DPS) [crafted]; Wolffear Harness (13110, -4.79 DPS) [world_drop]; Quillward Harness (10583, -4.86 DPS, sim-verified) [dungeon] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.6 ranged_attack_power points (3.37 DPS) | yes | Imperial Leather Bracers (4061, -0.47 DPS) [dungeon]; Green Leather Bracers (4259, -0.50 DPS) [crafted]; Dusky Bracers (7378, -2.93 DPS, sim-verified) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stalker's Leather Gloves (252526, -0.28 DPS) [crafted]; Aurora Gloves (4042, -0.87 DPS) [world_drop]; Dragonscale Gauntlets (8347, -3.94 DPS, sim-verified) [crafted] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 37.2 ranged_attack_power points (4.10 DPS) | yes | Defiler's Chain Girdle (20153, -0.60 DPS) [rep]; Defiler's Leather Girdle (20192, -0.79 DPS) [rep]; Scorpashi Sash (14652, -4.13 DPS, sim-verified) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 69.2 ranged_attack_power points (7.63 DPS) | yes | Triprunner Dungarees (9624, -1.59 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.54 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.54 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 45.0 ranged_attack_power points (4.97 DPS) | yes | Swampwalker Boots (2276, -0.24 DPS) [world_drop]; Imperial Leather Boots (6431, -0.97 DPS) [dungeon]; Worn Running Boots (9398, -4.23 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.9 ranged_attack_power points (3.63 DPS) | yes | Legionnaire's Band (19512, -0.73 DPS) [rep]; Field Researcher's Loop (281634, -1.09 DPS) [quest]; Falcon's Hook (7552, -3.62 DPS, sim-verified) [dungeon] |
| finger2 | Disengagement Ring (276202) | Friz Frazzlespark [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Band (19512, +0.00 DPS) [rep]; Field Researcher's Loop (281634, -0.36 DPS) [quest]; Falcon's Hook (7552, -0.82 DPS, sim-verified) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.25 DPS) [world]; Steel Spear (250605, -0.31 DPS) [crafted]; Claw of the Shadowmancer (2912, -13.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -3.25 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.46 DPS) [dungeon]; The Silencer (13138, -4.32 DPS, sim-verified) [world_drop] |

**New at 40:** head: Corpseshroud; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Disengagement Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 708, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (troll, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 323.5. Weights run: 6.2s. Verify run: 29.1s. 891 eligible items had no known source.

Race sweep: orc 328.1 ± 0.3, windshaper-skyborne 327.1 ± 0.3, troll 323.5 ± 0.3, tauren 319.8 ± 0.3 (dps); orc beats the guide race troll by 4.6 (1.41%) ± 0.4, beyond error.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.242 ± 0.016, crit=2.051 ± 0.030 per rating point (14 rating = 1%, 28.714 per %), hit=2.356 ± 0.083 per rating point (10 rating = 1%, 23.564 per %), melee_haste=9.615 ± 1.037, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Helm of Exile (11124) | Jammal'an the Prophet [quest] | sim-verified (+6.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.29 DPS) [crafted]; Tough Scorpid Helm (8208, -6.37 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.1 ranged_attack_power points (5.31 DPS) | yes | Scout's Medallion (19536, -0.82 DPS) [rep]; Woven Ivy Necklace (19159, -1.63 DPS) [quest] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted]; Phytoskin Spaulders (17749, -3.79 DPS, sim-verified) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.41 DPS) [dungeon]; Blackmetal Cape (9512, -1.23 DPS) [dungeon]; Blisterbane Wrap (12552, -2.10 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) | Lady Palanseer [vendor] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Blazewind Breastplate (11193, -1.87 DPS) [quest]; Icemail Jerkin (1981, -2.78 DPS, sim-verified) [world_drop]; Quillward Harness (10583, -3.51 DPS) [dungeon] |
| wrist | Stalker's Mail Bracers (252570) | Leatherworking [crafted] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodlust Bracelets (14807, -0.10 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.10 DPS) [crafted]; Deepfury Bracers (13120, -3.01 DPS, sim-verified) [world_drop] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -1.02 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.37 DPS) [crafted]; First Sergeant's Chain Gauntlets (220830, -6.55 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 63.1 ranged_attack_power points (7.96 DPS) | yes | Defiler's Chain Girdle (20151, -1.82 DPS) [rep]; Defiler's Leather Girdle (20193, -1.82 DPS) [rep]; Stalker's Mail Belt (252588, -1.85 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 86.2 ranged_attack_power points (10.86 DPS) | yes | Tough Scorpid Leggings (8206, -2.39 DPS) [crafted]; Imperial Red Pants (8251, -2.45 DPS) [world_drop]; Basilisk Hide Pants (1718, -6.46 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.9 ranged_attack_power points (8.68 DPS) | yes | Elven Chain Boots (13125, -1.32 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.99 DPS) [crafted]; Albino Crocscale Boots (17728, -3.46 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 45.4 ranged_attack_power points (5.72 DPS) | yes | Ring of the Underwood (2951, -1.63 DPS) [world_drop]; Falcon's Hook (7552, -2.04 DPS) [dungeon]; Legionnaire's Band (19511, -2.04 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.6 ranged_attack_power points (5.49 DPS) | yes | Falcon's Hook (7552, -1.81 DPS) [dungeon]; Legionnaire's Band (19511, -1.81 DPS) [rep]; Ring of the Underwood (2951, -2.07 DPS, sim-verified) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -2.36 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.92 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.49 DPS) [world_drop]; Staff of Jordan (873, -1.89 DPS) [world_drop]; Claw of the Shadowmancer (2912, -4.51 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -4.22 DPS, sim-verified) [dungeon]; Megashot Rifle (17717, -4.38 DPS) [dungeon]; Dark Iron Rifle (16004, -5.13 DPS) [crafted] |

**New at 50:** head: Helm of Exile; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Stalker's Mail Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Eyegouger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (troll, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 459.3. Weights run: 6.0s. Verify run: 69.3s. 2027 eligible items had no known source.

Race sweep: windshaper-skyborne 461.6 ± 0.3, orc 461.1 ± 0.3, troll 459.3 ± 0.3, tauren 451.1 ± 0.3 (dps); windshaper-skyborne beats the guide race troll by 2.4 (0.52%) ± 0.5, beyond error.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.193 ± 0.018, crit=2.468 ± 0.038 per rating point (14 rating = 1%, 34.551 per %), hit=2.847 ± 0.101 per rating point (10 rating = 1%, 28.475 per %), melee_haste=not significant (3.616 ± 1.419), intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 147.6 ranged_attack_power points (20.60 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -28.79 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.89 DPS) [world_drop]; Talisman of Evasion (13177, -0.89 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 120.8 ranged_attack_power points (16.86 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.49 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -1.90 DPS) [dungeon]; Deathguard's Cloak (20068, -2.50 DPS) [rep]; Howler's Furs (272414, -3.08 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 181.2 ranged_attack_power points (25.29 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.28 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.28 DPS) [vendor]; Savage Gladiator Chain (11726, -15.26 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 60.7 ranged_attack_power points (8.47 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.89 DPS) [pvp]; Slashclaw Bracers (13211, -1.37 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -1.52 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+4.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Girdle (16572, -2.00 DPS) [pvp]; Defiler's Chain Girdle (20150, -3.49 DPS) [rep]; Marksman's Girdle (22232, -4.57 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 174.8 ranged_attack_power points (24.40 DPS) | yes | General's Chain Legguards (231574, -2.39 DPS) [pvp]; General's Chain Legplates (231567, -4.95 DPS) [vendor]; Plaguehound Leggings (18736, -7.40 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.70 DPS) [dungeon]; Blackstone Ring (17713, -3.99 DPS) [dungeon]; Naglering (11669, -14.17 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.45 DPS) [dungeon]; Innervating Band (18701, -3.53 DPS) [dungeon]; Painweaver Band (13098, -6.67 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.16 DPS) [crafted]; Hand of Justice (11815, -1.94 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -5.11 DPS) [pvp]; Glacial Blade (19099, -9.33 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -1.04 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Second Wind; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (troll, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1065.9. Weights run: 4.0s. Verify run: 51.3s. 2027 eligible items had no known source.

Race sweep: windshaper-skyborne 1077.5 ± 0.8, orc 1066.6 ± 0.8, troll 1065.9 ± 0.8, tauren 1052.8 ± 0.8 (dps); windshaper-skyborne beats the guide race troll by 11.6 (1.09%) ± 1.1, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.303 ± 0.026, crit=2.698 ± 0.053 per rating point (14 rating = 1%, 37.776 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=22.566 ± 2.289, intellect=1.196 ± 0.008, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 152.6 ranged_attack_power points (41.18 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -15.33 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -1.78 DPS) [world_drop]; Talisman of Evasion (13177, -1.78 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 127.0 ranged_attack_power points (34.26 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, -7.86 DPS) [pvp]; Champion's Chain Shoulders (227069, -8.02 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -3.88 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -5.14 DPS) [rep]; Dark Phantom Cape (13122, -5.40 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 160.0 ranged_attack_power points (43.17 DPS) | yes | Warlord's Chain Chestpiece (16565, +0.00 DPS) [vendor]; Warlord's Chain Hauberk (231573, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -26.92 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 62.8 ranged_attack_power points (16.93 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.78 DPS) [pvp]; Beaststalker's Bindings (16681, -3.57 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Mourning Handguards (279257, +0.00 DPS) [crafted] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 96.0 ranged_attack_power points (25.91 DPS) | yes | General's Chain Girdle (16572, -4.06 DPS) [pvp]; Defiler's Leather Girdle (20190, -6.54 DPS) [rep]; Defiler's Chain Girdle (20150, -9.32 DPS, sim-verified) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 153.4 ranged_attack_power points (41.39 DPS) | yes | General's Chain Legguards (231574, +0.00 DPS) [pvp]; General's Chain Legplates (231567, -2.45 DPS) [vendor]; Outrider's Leather Pants (22740, -17.57 DPS, sim-verified) [rep] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 72.9 ranged_attack_power points (19.69 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -0.81 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Innervating Band (18701, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -1.49 DPS) [vendor]; Innervating Band (18701, -11.60 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+15.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -1.02 DPS) [pvp]; Glacial Blade (19099, -16.25 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -19.22 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Beaststalker's Boots; finger2: Don Julio's Band; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

