# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0050510000000000-000000000000000000)

Set DPS (verified): 93.6. Weights run: 7.1s. Verify run: 18.7s. 268 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.084 ± 0.010, crit=0.595 ± 0.011 per rating point (14 rating = 1%, 8.332 per %), hit=0.919 ± 0.027 per rating point (10 rating = 1%, 9.185 per %), melee_haste=4.128 ± 0.432, intellect=2.172 ± 0.023, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.7 ranged_attack_power points (2.23 DPS) | yes | Trapper's Leather Hood (252505, -1.21 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.25 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 18.5 ranged_attack_power points (1.67 DPS) | yes | Erudite's Amulet (277204, -1.28 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Magician's Mantle (12998, -0.25 DPS, sim-verified) [world_drop]; Reinforced Woolen Shoulders (4315, -0.61 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.32 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.56 DPS) [world_drop]; Bristlebark Cape (14571, -0.56 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 33.9 ranged_attack_power points (3.06 DPS) | yes | Trapper's Leather Armor (252491, -0.58 DPS, sim-verified) [crafted]; Mystic's Wrap (14369, -0.85 DPS) [world_drop]; Mystic's Robe (14371, -0.85 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.4 ranged_attack_power points (1.39 DPS) | yes | Wolf Bracers (4794, -0.32 DPS, sim-verified) [vendor]; Mindthrust Bracers (1974, -0.41 DPS) [dungeon]; Owl Bracers (4796, -0.41 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Riding Gloves (1304, -0.25 DPS) [quest]; Trapper's Leather Gloves (252495, -1.93 DPS, sim-verified) [crafted] |
| waist | Trapper's Leather Belt (252429) | Leatherworking [crafted] | 18.9 ranged_attack_power points (1.70 DPS) | yes | Blackened Defias Belt (10403, -0.08 DPS) [dungeon]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Deviate Scale Belt (6468, -0.31 DPS) [crafted] |
| legs | Leggings of the Fang (10410) | Wailing Caverns: Lord Cobrahn [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Duty Bound Leggings (279868, -0.05 DPS) [quest]; Trapper's Leather Pants (252501, -0.11 DPS, sim-verified) [crafted] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Trapper's Leather Boots (252440, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.24 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.2 ranged_attack_power points (1.73 DPS) | yes | Volcanic Rock Ring (12053, -1.15 DPS) [world_drop]; The 1 Ring (8350, -1.26 DPS) [world]; Minor Channeling Ring (1449, -1.34 DPS) [quest] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.3 ranged_attack_power points (1.11 DPS) | yes | The 1 Ring (8350, -0.64 DPS) [world]; Volcanic Rock Ring (12053, -0.70 DPS, sim-verified) [world_drop]; Minor Channeling Ring (1449, -0.72 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 38.6 ranged_attack_power points (3.49 DPS) | yes | Bronze Dory (250603, -0.60 DPS, sim-verified) [crafted]; Twisted Chanter's Staff (890, -1.53 DPS) [world_drop]; Sword of the Fallen (286977, -1.72 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (16.19 DPS) | yes | Lil Timmy's Peashooter (13136, -1.23 DPS) [world_drop]; Venomstrike (6469, -1.31 DPS, sim-verified) [dungeon]; Lovingly Crafted Boomstick (4372, -3.91 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Trapper's Leather Belt; legs: Leggings of the Fang; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Protector's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 149.6. Weights run: 8.0s. Verify run: 21.6s. 463 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.132 ± 0.013, crit=1.008 ± 0.018 per rating point (14 rating = 1%, 14.115 per %), hit=1.452 ± 0.038 per rating point (10 rating = 1%, 14.519 per %), melee_haste=3.723 ± 0.634, intellect=3.795 ± 0.030, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 60.7 ranged_attack_power points (5.02 DPS) | yes | Resilient Cap (14401, -1.05 DPS) [world_drop]; Nightsky Cowl (4039, -1.26 DPS) [world_drop]; Insignia Cap (4052, -1.57 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 34.8 ranged_attack_power points (2.88 DPS) | yes | Kaleidoscope Chain (13084, -0.58 DPS) [world_drop]; Sentinel's Medallion (19541, -0.80 DPS) [rep]; Seal of the Expedition (271920, -1.31 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | sim-verified (149.6 DPS) | yes | Feline Mantle (3748, +0.00 DPS) [dungeon]; Death Speaker Mantle (6685, +0.00 DPS) [dungeon]; Resilient Mantle (14397, -0.02 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | sim-verified (149.6 DPS) | yes | Sergeant Major's Cape (16315, +0.00 DPS) [pvp]; Expeditionary's Cape (271923, +0.00 DPS) [rep]; Resilient Cape (14400, -0.19 DPS) [world_drop] |
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
| main_hand | Felbough Arm (273037) | Fel Ancient [world] | sim-verified (149.6 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Kam's Walking Stick (2280, -0.26 DPS) [dungeon]; Glimmering Staff (249392, -0.34 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.65 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.78 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Seal of Wrynn; main_hand: Felbough Arm; ranged: Crystalpine Stinger

No-known-source sample (15 of 463, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 227.5. Weights run: 7.2s. Verify run: 19.4s. 718 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.245 ± 0.017, crit=1.602 ± 0.027 per rating point (14 rating = 1%, 22.432 per %), hit=1.906 ± 0.048 per rating point (10 rating = 1%, 19.055 per %), melee_haste=5.744 ± 0.968, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 45.8 ranged_attack_power points (5.06 DPS) | yes | Nightscape Headband (8176, -0.76 DPS) [crafted]; Guard's Chain Helm (250499, -0.76 DPS) [crafted]; Corpseshroud (10574, -0.77 DPS, sim-verified) [dungeon] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 35.7 ranged_attack_power points (3.94 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.73 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.81 DPS) [world]; Kaleidoscope Chain (13084, -1.97 DPS) [world_drop] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 35.7 ranged_attack_power points (3.94 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.36 DPS) [dungeon]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (227.5 DPS) | yes | Imperial Cloak (6432, -0.08 DPS) [dungeon]; Parachute Cloak (10518, -0.08 DPS) [crafted]; Tigerstrike Mantle (13108, -0.08 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 96.9 ranged_attack_power points (10.70 DPS) | yes | Quillward Harness (10583, -4.19 DPS, sim-verified) [dungeon]; Tough Scorpid Breastplate (8203, -4.39 DPS) [crafted]; Wolffear Harness (13110, -4.61 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.0 ranged_attack_power points (3.31 DPS) | yes | Imperial Leather Bracers (4061, -0.44 DPS) [dungeon]; Dusky Bracers (7378, -0.44 DPS) [crafted]; Green Leather Bracers (4259, -0.49 DPS) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-verified (227.5 DPS) | yes | Dragonscale Gauntlets (8347, -0.16 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.27 DPS) [crafted]; Aurora Gloves (4042, -0.85 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 36.5 ranged_attack_power points (4.02 DPS) | yes | Scorpashi Sash (14652, -0.44 DPS) [world_drop]; Highlander's Chain Girdle (20089, -0.67 DPS) [rep]; Highlander's Leather Girdle (20116, -0.71 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 68.1 ranged_attack_power points (7.52 DPS) | yes | Triprunner Dungarees (9624, -1.16 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.51 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.51 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 44.2 ranged_attack_power points (4.87 DPS) | yes | Swampwalker Boots (2276, -0.22 DPS) [world_drop]; Worn Running Boots (9398, -0.22 DPS) [dungeon]; Imperial Leather Boots (6431, -0.93 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.5 ranged_attack_power points (3.58 DPS) | yes | Protector's Band (19515, -0.72 DPS) [rep]; Disengagement Ring (276202, -0.72 DPS) [vendor]; Field Researcher's Loop (281634, -1.07 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 29.2 ranged_attack_power points (3.22 DPS) | yes | Protector's Band (19515, -0.36 DPS) [rep]; Disengagement Ring (276202, -0.36 DPS) [vendor]; Field Researcher's Loop (281634, -0.72 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (227.5 DPS) | yes | Steel Spear (250605, -0.07 DPS) [crafted]; Felbough Arm (273037, -0.14 DPS) [world]; Claw of the Shadowmancer (2912, -3.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (227.5 DPS) | yes | The Silencer (13138, -3.09 DPS, sim-verified) [world_drop]; Swiftwind (13038, -3.29 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.46 DPS) [dungeon] |

**New at 40:** head: Warden's Wizard Hat; neck: Sentinel's Medallion; shoulder: Nightscape Shoulders; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 330.2. Weights run: 7.7s. Verify run: 25.4s. 900 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.197 ± 0.015, crit=1.967 ± 0.028 per rating point (14 rating = 1%, 27.542 per %), hit=2.224 ± 0.069 per rating point (10 rating = 1%, 22.242 per %), melee_haste=9.090 ± 0.876, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 61.7 ranged_attack_power points (7.64 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Exile (11124, -0.59 DPS) [quest]; Helm of Fire (8348, -0.91 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.6 ranged_attack_power points (5.15 DPS) | yes | Sentinel's Medallion (19539, -0.40 DPS) [rep] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (330.2 DPS) | yes | Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.40 DPS) [dungeon]; Blackmetal Cape (9512, -1.19 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 99.0 ranged_attack_power points (12.26 DPS) | yes | Knight's Chain Armor (220828, -1.51 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -3.15 DPS) [quest]; Quillward Harness (10583, -4.74 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 48.0 ranged_attack_power points (5.94 DPS) | yes | Stalker's Mail Bracers (252570, -1.23 DPS, sim-verified) [crafted]; Bloodlust Bracelets (14807, -1.58 DPS) [world_drop]; Wicked Leather Bracers (15084, -1.58 DPS) [crafted] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -1.11 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.35 DPS) [crafted]; Sergeant Major's Chain Gauntlets (220829, -2.25 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 62.5 ranged_attack_power points (7.74 DPS) | yes | Stalker's Mail Belt (252588, -0.88 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20088, -1.85 DPS) [rep]; Highlander's Leather Girdle (20115, -1.85 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 84.4 ranged_attack_power points (10.46 DPS) | yes | Tough Scorpid Leggings (8206, -2.22 DPS) [crafted]; Imperial Red Pants (8251, -2.27 DPS) [world_drop]; Basilisk Hide Pants (1718, -3.02 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.0 ranged_attack_power points (8.42 DPS) | yes | Albino Crocscale Boots (17728, -1.05 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -1.30 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.93 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 44.8 ranged_attack_power points (5.55 DPS) | yes | Ring of the Underwood (2951, -1.58 DPS) [world_drop]; Falcon's Hook (7552, -1.98 DPS) [dungeon]; Protector's Band (19516, -1.98 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.2 ranged_attack_power points (5.23 DPS) | yes | Falcon's Hook (7552, -1.67 DPS) [dungeon]; Protector's Band (19516, -1.67 DPS) [rep]; Ring of the Underwood (2951, -1.82 DPS, sim-verified) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.62 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.84 DPS) [world_drop]; Staff of Jordan (873, -1.98 DPS) [world_drop]; Claw of the Shadowmancer (2912, -5.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Megashot Rifle (17717, -4.23 DPS) [dungeon]; Dark Iron Rifle (16004, -4.97 DPS) [crafted]; Gryphonwing Long Bow (13022, -8.59 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Molten Heart of the Mountain; main_hand: Eyegouger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 900, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 471.9. Weights run: 7.4s. Verify run: 80.7s. 2028 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.151 ± 0.017, crit=2.322 ± 0.035 per rating point (14 rating = 1%, 32.507 per %), hit=2.789 ± 0.100 per rating point (10 rating = 1%, 27.892 per %), melee_haste=not significant (2.598 ± 1.296), intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 144.6 ranged_attack_power points (20.73 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Field Marshal's Chain Greathelm (231562, -1.16 DPS) [vendor]; Outlaw's Collar (279253, -15.35 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.90 DPS) [world_drop]; Talisman of Evasion (13177, -0.90 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 117.6 ranged_attack_power points (16.85 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Shoulders (231576, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -1.96 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.51 DPS) [rep]; Howler's Furs (272414, -2.88 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 177.0 ranged_attack_power points (25.36 DPS) | yes | Field Marshal's Chain Breastplate (16466, -3.27 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, -3.27 DPS) [vendor]; Savage Gladiator Chain (11726, -13.16 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.9 ranged_attack_power points (8.58 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -0.90 DPS) [pvp]; Slashclaw Bracers (13211, -1.42 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -0.13 DPS) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Marshal's Chain Girdle (16464, -2.02 DPS) [pvp]; Highlander's Chain Girdle (20043, -3.73 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 203.2 ranged_attack_power points (29.11 DPS) | yes | Marshal's Chain Legguards (231577, -7.03 DPS) [pvp]; Sentinel's Leather Pants (237818, -7.61 DPS) [vendor]; Marshal's Chain Legplates (231558, -9.54 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.71 DPS) [dungeon]; Blackstone Ring (17713, -3.80 DPS) [dungeon]; Naglering (11669, -11.61 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.45 DPS) [dungeon]; Innervating Band (18701, -3.55 DPS) [dungeon]; Painweaver Band (13098, -4.06 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (471.9 DPS) | yes | - |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -5.65 DPS) [pvp]; Electrified Dagger (19100, -7.99 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -1.59 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Hand of Justice; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1063.2. Weights run: 5.4s. Verify run: 61.8s. 2028 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.305 ± 0.026, crit=2.763 ± 0.054 per rating point (14 rating = 1%, 38.681 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=21.933 ± 2.291, intellect=1.200 ± 0.008, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 153.7 ranged_attack_power points (40.64 DPS) | yes | Field Marshal's Chain Helm (231580, +0.00 DPS) [pvp]; Lieutenant Commander's Chain Helm (227066, -1.59 DPS) [pvp]; Outlaw's Collar (279253, -15.61 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -1.75 DPS) [world_drop]; Talisman of Evasion (13177, -1.75 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 127.9 ranged_attack_power points (33.83 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Lieutenant Commander's Chain Shoulders (227068, -7.87 DPS) [pvp]; Field Marshal's Chain Spaulders (16468, -7.93 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -3.78 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -5.04 DPS) [rep]; Dark Phantom Cape (13122, -5.29 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 161.0 ranged_attack_power points (42.57 DPS) | yes | Field Marshal's Chain Breastplate (16466, +0.00 DPS) [vendor]; Field Marshal's Chain Hauberk (231581, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -25.78 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 62.8 ranged_attack_power points (16.61 DPS) | yes | Marshal's Chain Bracers (16461, +0.00 DPS) [pvp]; Knight-Captain's Chain Armguards (16402, -1.75 DPS) [pvp]; Beaststalker's Bindings (16681, -3.50 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -0.80 DPS) [crafted] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 96.1 ranged_attack_power points (25.42 DPS) | yes | Marshal's Chain Girdle (16464, -3.97 DPS) [pvp]; Highlander's Leather Girdle (20045, -6.20 DPS) [rep]; Highlander's Chain Girdle (20043, -8.42 DPS, sim-verified) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 193.0 ranged_attack_power points (51.05 DPS) | yes | Marshal's Chain Legguards (231577, -7.30 DPS) [pvp]; Marshal's Chain Legplates (231558, -12.63 DPS) [vendor]; Sentinel's Leather Pants (237818, -17.30 DPS, sim-verified) [vendor] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 73.0 ranged_attack_power points (19.31 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (1063.2 DPS) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Innervating Band (18701, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -1.70 DPS) [vendor]; Innervating Band (18701, -10.89 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+11.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Heart of Wyrmthalak (22321, -5.27 DPS, sim-verified) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, -0.51 DPS) [pvp]; Electrified Dagger (19100, -17.03 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Bullseye (234585, +0.00 DPS) [pvp]; Grand Marshal's Repeater (234586, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -19.38 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Sentinel's Medallion; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Beaststalker's Boots; finger2: Don Julio's Band; trinket2: Frozen Heart of the Mountain; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (troll, 0000000000000000-0050510000000000-000000000000000000)

Set DPS (verified): 90.3. Weights run: 7.1s. Verify run: 17.0s. 264 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.084 ± 0.010, crit=0.595 ± 0.011 per rating point (14 rating = 1%, 8.332 per %), hit=0.919 ± 0.027 per rating point (10 rating = 1%, 9.185 per %), melee_haste=4.128 ± 0.432, intellect=2.172 ± 0.023, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 24.7 ranged_attack_power points (2.23 DPS) | yes | Trapper's Leather Hood (252505, -1.16 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.25 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 18.5 ranged_attack_power points (1.67 DPS) | yes | Erudite's Amulet (277204, -1.24 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 19.6 ranged_attack_power points (1.76 DPS) | yes | Serpent's Shoulders (5404, -0.37 DPS) [dungeon]; Reinforced Woolen Shoulders (4315, -0.98 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (90.3 DPS) | yes | Cape of the Brotherhood (5193, -0.32 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.56 DPS) [world_drop]; Bristlebark Cape (14571, -0.56 DPS) [world_drop] |
| chest | Trapper's Leather Armor (252491) | Leatherworking [crafted] | 28.1 ranged_attack_power points (2.54 DPS) | yes | Mystic's Robe (14371, -0.33 DPS) [world_drop]; Brawler's Leather Armor (252490, -0.59 DPS) [crafted]; Mystic's Wrap (14369, -0.62 DPS, sim-verified) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 15.4 ranged_attack_power points (1.39 DPS) | yes | Wolf Bracers (4794, -0.32 DPS, sim-verified) [vendor]; Mindthrust Bracers (1974, -0.41 DPS) [dungeon]; Owl Bracers (4796, -0.41 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | sim-verified (90.3 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Forest Leather Gloves (3058, -0.56 DPS) [world_drop]; Trapper's Leather Gloves (252495, -2.12 DPS, sim-verified) [crafted] |
| waist | Trapper's Leather Belt (252429) | Leatherworking [crafted] | 18.9 ranged_attack_power points (1.70 DPS) | yes | Blackened Defias Belt (10403, -0.08 DPS) [dungeon]; Keller's Girdle (2911, -0.13 DPS) [world_drop]; Deviate Scale Belt (6468, -0.31 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 36.4 ranged_attack_power points (3.29 DPS) | yes | Leggings of the Fang (10410, -0.78 DPS) [dungeon]; Brawler's Leather Pants (252500, -0.78 DPS) [crafted]; Bluegill Breeches (3022, -1.06 DPS) [world] |
| feet | Footpads of the Fang (10411) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (90.3 DPS) | yes | Blackened Defias Boots (10402, +0.00 DPS) [dungeon]; Trapper's Leather Boots (252440, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.46 DPS, sim-verified) [world_drop] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 19.2 ranged_attack_power points (1.73 DPS) | yes | Loop of Sacrifice (281673, -0.75 DPS) [quest]; Bounty Hunter's Ring (5351, -0.90 DPS) [quest]; Volcanic Rock Ring (12053, -1.15 DPS) [world_drop] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.3 ranged_attack_power points (1.11 DPS) | yes | Bounty Hunter's Ring (5351, -0.28 DPS) [quest]; Loop of Sacrifice (281673, -0.45 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Impaling Harpoon (5200) | The Deadmines: Captain Greenskin [dungeon] | 38.6 ranged_attack_power points (3.49 DPS) | yes | Crescent Staff (6505, -0.32 DPS, sim-verified) [quest]; Staff of Orgrimmar (15444, -0.64 DPS) [quest]; Bronze Dory (250603, -0.72 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 179.4 ranged_attack_power points (16.19 DPS) | yes | Venomstrike (6469, -0.98 DPS, sim-verified) [dungeon]; Lil Timmy's Peashooter (13136, -1.23 DPS) [world_drop]; Lovingly Crafted Boomstick (4372, -3.91 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Glowing Lizardscale Cloak; chest: Trapper's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Trapper's Leather Belt; legs: Trapper's Leather Pants; feet: Footpads of the Fang; finger1: Lavishly Jeweled Ring; finger2: Legionnaire's Band; main_hand: Impaling Harpoon; ranged: Ranger Bow

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (troll, 0000000000000000-0050552001300000-000000000000000000)

Set DPS (verified): 146.3. Weights run: 8.0s. Verify run: 21.6s. 462 eligible items had no known source.

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
| finger2 | Ring of Precision (1491) | Blackfathom Deeps [dungeon] | sim-verified (146.3 DPS) | yes | Azora's Will (4999, +0.00 DPS) [world_drop]; Snake Hoop (6750, +0.00 DPS) [quest]; Loop of Sacrifice (281673, +0.00 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Impaling Harpoon (5200, -0.11 DPS) [dungeon]; Felbough Arm (273037, -0.22 DPS) [world]; Kam's Walking Stick (2280, -0.48 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Crystalpine Stinger (13037) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Golemsight Long Gun (273029, +0.00 DPS) [dungeon]; Ironweaver (13137, -0.65 DPS) [world_drop]; Harpyclaw Short Bow (13019, -0.78 DPS) [world_drop] |

**New at 30:** head: Enduring Cap; neck: Archmagister's Faceted Pendant; shoulder: Death Speaker Mantle; back: Tigerstrike Mantle; chest: Garb of Florid Feathers; wrist: Glowing Magical Bracelets; hands: Hotshot Pilot's Gloves; waist: Stalker's Leather Belt; legs: Ferine Leggings; feet: Trapper's Leather Boots; finger2: Ring of Precision; main_hand: Crescent Staff; ranged: Crystalpine Stinger

No-known-source sample (15 of 462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (troll, 0000000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 224.8. Weights run: 7.2s. Verify run: 19.8s. 708 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.245 ± 0.017, crit=1.602 ± 0.027 per rating point (14 rating = 1%, 22.432 per %), hit=1.906 ± 0.048 per rating point (10 rating = 1%, 19.055 per %), melee_haste=5.744 ± 0.968, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 45.8 ranged_attack_power points (5.06 DPS) | yes | Nightscape Headband (8176, -0.76 DPS) [crafted]; Guard's Chain Helm (250499, -0.76 DPS) [crafted]; Corpseshroud (10574, -0.89 DPS, sim-verified) [dungeon] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 35.7 ranged_attack_power points (3.94 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.73 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.81 DPS) [world]; Kaleidoscope Chain (13084, -1.97 DPS) [world_drop] |
| shoulder | Nightscape Shoulders (8192) (or Forest Tracker Epaulets (2278)) | Leatherworking [crafted] | 35.7 ranged_attack_power points (3.94 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS) [world_drop]; Mantle of Thieves (2264, -0.36 DPS) [dungeon]; Flintrock Shoulders (7755, -0.36 DPS) [dungeon] |
| back | Parachute Cloak (10518) | Engineering [crafted] | sim-verified (224.8 DPS) | yes | Imperial Cloak (6432, +0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, +0.00 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.36 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 96.9 ranged_attack_power points (10.70 DPS) | yes | Quillward Harness (10583, -4.35 DPS, sim-verified) [dungeon]; Tough Scorpid Breastplate (8203, -4.39 DPS) [crafted]; Wolffear Harness (13110, -4.61 DPS) [world_drop] |
| wrist | Tough Scorpid Bracers (8205) | Leatherworking [crafted] | 30.0 ranged_attack_power points (3.31 DPS) | yes | Imperial Leather Bracers (4061, -0.44 DPS) [dungeon]; Dusky Bracers (7378, -0.44 DPS) [crafted]; Green Leather Bracers (4259, -0.49 DPS) [crafted] |
| hands | Tough Scorpid Gloves (8204) | Leatherworking [crafted] | sim-verified (224.8 DPS) | yes | Dragonscale Gauntlets (8347, -0.16 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.27 DPS) [crafted]; Aurora Gloves (4042, -0.85 DPS) [world_drop] |
| waist | Stalker's Leather Belt (252521) | Leatherworking [crafted] | 36.5 ranged_attack_power points (4.02 DPS) | yes | Scorpashi Sash (14652, -0.44 DPS) [world_drop]; Defiler's Chain Girdle (20153, -0.67 DPS) [rep]; Defiler's Leather Girdle (20192, -0.71 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 68.1 ranged_attack_power points (7.52 DPS) | yes | Triprunner Dungarees (9624, -1.16 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -2.51 DPS) [dungeon]; Hawkeye's Breeches (14595, -2.51 DPS) [world_drop] |
| feet | Stalker's Mail Boots (252562) | Leatherworking [crafted] | 44.2 ranged_attack_power points (4.87 DPS) | yes | Swampwalker Boots (2276, -0.22 DPS) [world_drop]; Worn Running Boots (9398, -0.22 DPS) [dungeon]; Imperial Leather Boots (6431, -0.93 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 32.5 ranged_attack_power points (3.58 DPS) | yes | Legionnaire's Band (19512, -0.72 DPS) [rep]; Disengagement Ring (276202, -0.72 DPS) [vendor]; Field Researcher's Loop (281634, -1.07 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 29.2 ranged_attack_power points (3.22 DPS) | yes | Legionnaire's Band (19512, -0.36 DPS) [rep]; Disengagement Ring (276202, -0.36 DPS) [vendor]; Field Researcher's Loop (281634, -0.72 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (224.8 DPS) | yes | Steel Spear (250605, -0.07 DPS) [crafted]; Felbough Arm (273037, -0.14 DPS) [world]; Claw of the Shadowmancer (2912, -3.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bow of Searing Arrows (2825) | World drop [world_drop] | sim-verified (224.8 DPS) | yes | Swiftwind (13038, -3.29 DPS) [world_drop]; Shadowforge Bushmaster (9422, -3.46 DPS) [dungeon]; The Silencer (13138, -4.11 DPS, sim-verified) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; neck: Scout's Medallion; shoulder: Nightscape Shoulders; back: Parachute Cloak; chest: Icemail Jerkin; wrist: Tough Scorpid Bracers; hands: Tough Scorpid Gloves; legs: Basilisk Hide Pants; feet: Stalker's Mail Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Bow of Searing Arrows

No-known-source sample (15 of 708, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (troll, 5320000000000000-0050552001521050-000000000000000000)

Set DPS (verified): 325.5. Weights run: 7.7s. Verify run: 26.9s. 891 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.197 ± 0.015, crit=1.967 ± 0.028 per rating point (14 rating = 1%, 27.542 per %), hit=2.224 ± 0.069 per rating point (10 rating = 1%, 22.242 per %), melee_haste=9.090 ± 0.876, intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tough Scorpid Helm (8208) | Leatherworking [crafted] | 61.7 ranged_attack_power points (7.64 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Exile (11124, -0.59 DPS) [quest]; Helm of Fire (8348, -0.91 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.6 ranged_attack_power points (5.15 DPS) | yes | Scout's Medallion (19535, -0.40 DPS) [rep]; Woven Ivy Necklace (19159, -1.58 DPS) [quest] |
| shoulder | Tough Scorpid Shoulders (8207) | Leatherworking [crafted] | sim-verified (325.5 DPS) | yes | Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Stalker's Mail Shoulders (252566, -0.05 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackveil Cape (11626, -0.40 DPS) [dungeon]; Blackmetal Cape (9512, -1.19 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 99.0 ranged_attack_power points (12.26 DPS) | yes | Stone Guard's Chain Armor (220827, -1.62 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -3.15 DPS) [quest]; Quillward Harness (10583, -4.74 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 48.0 ranged_attack_power points (5.94 DPS) | yes | Stalker's Mail Bracers (252570, -1.25 DPS, sim-verified) [crafted]; Bloodlust Bracelets (14807, -1.58 DPS) [world_drop]; Wicked Leather Bracers (15084, -1.58 DPS) [crafted] |
| hands | Stalker's Mail Gauntlets (252584) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, -1.11 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.35 DPS) [crafted]; First Sergeant's Chain Gauntlets (220830, -2.37 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 62.5 ranged_attack_power points (7.74 DPS) | yes | Stalker's Mail Belt (252588, -0.91 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20151, -1.85 DPS) [rep]; Defiler's Leather Girdle (20193, -1.85 DPS) [rep] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 84.4 ranged_attack_power points (10.46 DPS) | yes | Tough Scorpid Leggings (8206, -2.22 DPS) [crafted]; Imperial Red Pants (8251, -2.27 DPS) [world_drop]; Basilisk Hide Pants (1718, -3.19 DPS, sim-verified) [world_drop] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.0 ranged_attack_power points (8.42 DPS) | yes | Albino Crocscale Boots (17728, -1.17 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -1.30 DPS) [world_drop]; Stalker's Mail Sabatons (252576, -1.93 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 44.8 ranged_attack_power points (5.55 DPS) | yes | Ring of the Underwood (2951, -1.58 DPS) [world_drop]; Falcon's Hook (7552, -1.98 DPS) [dungeon]; Legionnaire's Band (19511, -1.98 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.2 ranged_attack_power points (5.23 DPS) | yes | Falcon's Hook (7552, -1.67 DPS) [dungeon]; Legionnaire's Band (19511, -1.67 DPS) [rep]; Ring of the Underwood (2951, -2.13 DPS, sim-verified) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.12 DPS, sim-verified) [crafted] |
| main_hand | Eyegouger (9480) | Zul'Farrak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.84 DPS) [world_drop]; Staff of Jordan (873, -1.98 DPS) [world_drop]; Claw of the Shadowmancer (2912, -5.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Megashot Rifle (17717, -4.23 DPS) [dungeon]; Dark Iron Rifle (16004, -4.97 DPS) [crafted]; Gryphonwing Long Bow (13022, -5.67 DPS, sim-verified) [dungeon] |

**New at 50:** head: Tough Scorpid Helm; neck: Skibi's Pendant; shoulder: Tough Scorpid Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; hands: Stalker's Mail Gauntlets; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Eyegouger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (troll, 5320000500000000-0050552001521050-500000000000000000)

Set DPS (verified): 469.0. Weights run: 7.4s. Verify run: 82.2s. 2027 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.151 ± 0.017, crit=2.322 ± 0.035 per rating point (14 rating = 1%, 32.507 per %), hit=2.789 ± 0.100 per rating point (10 rating = 1%, 27.892 per %), melee_haste=not significant (2.598 ± 1.296), intellect=1.210 ± 0.001, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 144.6 ranged_attack_power points (20.73 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -15.68 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -0.90 DPS) [world_drop]; Talisman of Evasion (13177, -0.90 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 117.6 ranged_attack_power points (16.85 DPS) | yes | Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Chain Pauldrons (231565, -1.24 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -1.96 DPS) [dungeon]; Deathguard's Cloak (20068, -2.51 DPS) [rep]; Howler's Furs (272414, -2.69 DPS, sim-verified) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 177.0 ranged_attack_power points (25.36 DPS) | yes | Warlord's Chain Chestpiece (16565, -3.27 DPS) [vendor]; Warlord's Chain Hauberk (231573, -3.27 DPS) [vendor]; Savage Gladiator Chain (11726, -11.26 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.9 ranged_attack_power points (8.58 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -0.90 DPS) [pvp]; Slashclaw Bracers (13211, -1.42 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -4.16 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marksman's Girdle (22232, -0.65 DPS, sim-verified) [dungeon]; General's Chain Girdle (16572, -2.02 DPS) [pvp]; Defiler's Chain Girdle (20150, -3.73 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 170.7 ranged_attack_power points (24.46 DPS) | yes | General's Chain Legguards (231574, -2.37 DPS) [pvp]; General's Chain Legplates (231567, -4.88 DPS) [vendor]; Plaguehound Leggings (18736, -7.19 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (469.0 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.71 DPS) [dungeon]; Blackstone Ring (17713, -3.80 DPS) [dungeon]; Naglering (11669, -10.17 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.31 DPS, sim-verified) [dungeon]; Blackstone Ring (17713, -3.45 DPS) [dungeon]; Innervating Band (18701, -3.55 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.21 DPS) [crafted]; Hand of Justice (11815, -1.94 DPS) [dungeon] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -5.65 DPS) [pvp]; Glacial Blade (19099, -7.73 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -2.97 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (troll, 5000000000000000-0052550011503051-500050030000000000)

Set DPS (verified): 1045.7. Weights run: 5.4s. Verify run: 63.9s. 2027 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): ranged_attack_power=1.000 ± 0.001, agility=3.305 ± 0.026, crit=2.763 ± 0.054 per rating point (14 rating = 1%, 38.681 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=21.933 ± 2.291, intellect=1.200 ± 0.008, spell_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 153.7 ranged_attack_power points (40.64 DPS) | yes | Warlord's Chain Helmet (16566, +0.00 DPS) [vendor]; Warlord's Chain Helm (231571, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -14.58 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Skibi's Pendant (13089, -1.75 DPS) [world_drop]; Talisman of Evasion (13177, -1.75 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 127.9 ranged_attack_power points (33.83 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Chain Shoulders (227069, -7.87 DPS) [pvp]; Warlord's Chain Shoulders (231572, -7.93 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -3.50 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -5.04 DPS) [rep]; Dark Phantom Cape (13122, -5.29 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 161.0 ranged_attack_power points (42.57 DPS) | yes | Warlord's Chain Chestpiece (16565, +0.00 DPS) [vendor]; Warlord's Chain Hauberk (231573, +0.00 DPS) [vendor]; Savage Gladiator Chain (11726, -26.15 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 62.8 ranged_attack_power points (16.61 DPS) | yes | General's Chain Wristguards (16570, +0.00 DPS) [pvp]; Legionnaire's Chain Bracers (16517, -1.75 DPS) [pvp]; Beaststalker's Bindings (16681, -3.50 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Vices (231575, +0.00 DPS) [vendor]; Mourning Handguards (279257, +0.00 DPS) [crafted] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 96.1 ranged_attack_power points (25.42 DPS) | yes | General's Chain Girdle (16572, -3.97 DPS) [pvp]; Defiler's Leather Girdle (20190, -6.20 DPS) [rep]; Defiler's Chain Girdle (20150, -8.45 DPS, sim-verified) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 154.3 ranged_attack_power points (40.82 DPS) | yes | General's Chain Legguards (231574, +0.00 DPS) [pvp]; General's Chain Legplates (231567, -2.40 DPS) [vendor]; Outrider's Leather Pants (22740, -17.12 DPS, sim-verified) [rep] |
| feet | Beaststalker's Boots (16675) | Stratholme: Nerub'enkan [dungeon] | 73.0 ranged_attack_power points (19.31 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Fleetfoot Greaves (11627, -0.80 DPS) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (1045.7 DPS) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Innervating Band (18701, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, -1.70 DPS) [vendor]; Innervating Band (18701, -11.30 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+15.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Legionite Glaive (250619) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, -0.51 DPS) [pvp]; Glacial Blade (19099, -15.29 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Eaglehorn Long Bow (13023) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Recurve (234559, +0.00 DPS) [pvp]; High Warlord's Crossbow (234560, +0.00 DPS) [pvp]; Hyper Deluxe Sniper Rifle Mk XVII (279273, -19.71 DPS, sim-verified) [crafted] |

**New at 60:** head: Black Dragonscale Helm; neck: Scout's Medallion; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Beaststalker's Boots; finger2: Don Julio's Band; main_hand: Legionite Glaive; ranged: Eaglehorn Long Bow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

