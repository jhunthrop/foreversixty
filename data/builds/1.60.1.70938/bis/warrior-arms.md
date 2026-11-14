# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 12321200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 33.0. Weights run: 5.1s. Verify run: 13.8s. 388 eligible items had no known source.

Race sweep: high-order-skyborne 33.4 ± 0.0, night-elf 33.0 ± 0.0, human 32.9 ± 0.0, dwarf 32.6 ± 0.0, gnome 32.5 ± 0.0 (dps); high-order-skyborne beats the guide race night-elf by 0.4 (1.28%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.947 ± 0.021, crit=0.528 ± 0.019 per rating point (14 rating = 1%, 7.392 per %), hit=0.614 ± 0.070 per rating point (10 rating = 1%, 6.144 per %), melee_haste=6.863 ± 0.597

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.96 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.51 DPS) [crafted]; Brawler's Leather Hood (252504, -0.60 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.7 attack_power points (0.27 DPS) | yes | Erudite's Amulet (277204, -0.21 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.56 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.29 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.06 DPS) [crafted]; Grave Shroud (279865, -0.10 DPS) [quest]; Glowing Lizardscale Cloak (6449, -0.11 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.96 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.15 DPS) [crafted]; Brawler's Leather Armor (252490, -0.16 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.48 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.06 DPS) [quest]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.77 DPS) | yes | Gold-flecked Gloves (5195, -0.10 DPS) [dungeon]; Gloves of the Fang (10413, -0.11 DPS) [dungeon]; Veteran's Gloves (250508, -0.15 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.87 DPS) | yes | Cobrahn's Grasp (6460, -0.06 DPS) [dungeon]; Brawler's Leather Belt (252428, -0.30 DPS) [crafted]; Veteran's Chain Belt (250513, -0.34 DPS) [crafted] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 22.7 attack_power points (1.09 DPS) | yes | Chausses of Westfall (6087, -0.04 DPS) [quest]; Defender's Leather Pants (252445, -0.05 DPS) [crafted]; Brawler's Leather Pants (252500, -0.11 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.7 attack_power points (0.71 DPS) | yes | Veteran's Boots (250503, -0.05 DPS) [crafted]; Feet of the Lynx (1121, -0.06 DPS) [world_drop]; Defender's Leather Boots (252441, -0.23 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.8 attack_power points (0.57 DPS) | yes | Ring of the Moon (12052, -0.47 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.39 DPS) | yes | Ring of the Moon (12052, -0.29 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 301.8 attack_power points (14.53 DPS) | yes | Duskbringer (2205, -0.52 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.58 DPS) [dungeon]; Sword of the Fallen (286977, -0.66 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567), Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.19 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.01 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 30 (night-elf, 12323213130000000-000000000000000000-000000000000000000)

Set DPS (verified): 68.2. Weights run: 5.9s. Verify run: 26.0s. 637 eligible items had no known source.

Race sweep: night-elf 68.2 ± 0.1, human 67.4 ± 0.1, gnome 66.6 ± 0.1, dwarf 66.3 ± 0.1, high-order-skyborne 65.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.503 ± 0.039, crit=1.126 ± 0.039 per rating point (14 rating = 1%, 15.761 per %), hit=1.383 ± 0.141 per rating point (10 rating = 1%, 13.835 per %), melee_haste=8.326 ± 1.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 31.5 attack_power points (1.83 DPS) | yes | Tusken Helm (6686, -0.33 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.44 DPS) [crafted]; Defender's Leather Helm (252455, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.0 attack_power points (0.81 DPS) | yes | Sentinel's Medallion (19541, -0.12 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.12 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.5 attack_power points (1.54 DPS) | yes | Barbaric Iron Shoulders (7913, -0.32 DPS) [crafted]; Barbaric Shoulders (5964, -0.52 DPS) [crafted]; Mantle of Thieves (2264, -0.67 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, -0.15 DPS) [pvp]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop]; Wolfmaster Cape (6314, -0.38 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Silvered Chain Shirt (250518, -0.06 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.32 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.0 attack_power points (0.99 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Pugilist Bracers (4438, -0.06 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.0 attack_power points (1.33 DPS) | yes | Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.0 attack_power points (1.68 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.20 DPS) [crafted]; Girdle of Golem Strength (9405, -0.29 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.5 attack_power points (1.89 DPS) | yes | Glimmering Mail Legguards (6386, -0.35 DPS) [world_drop]; Ferine Leggings (6690, -0.38 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.39 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.0 attack_power points (1.04 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Alacritous Treads (277234, -0.03 DPS) [quest]; Veteran's Boots (250503, -0.12 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.0 attack_power points (1.22 DPS) | yes | Seal of Wrynn (2933, -0.61 DPS) [quest]; Tiger Band (6749, -0.64 DPS) [quest]; Bonebinder's Signet (273032, -0.64 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.5 attack_power points (1.19 DPS) | yes | Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world]; Seal of Wrynn (2933, -0.65 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.59 DPS) [dungeon]; Felbough Arm (273037, -0.75 DPS) [world]; Viscous Hammer (13045, -4.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.5 attack_power points (0.61 DPS) | yes | Golemsight Long Gun (273029, -0.09 DPS) [dungeon]; Gravespike Repeater (273033, -0.15 DPS) [world]; Silver Star (3463, -0.26 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 637, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 12323213132214001-000000000000000000-000000000000000000)

Set DPS (verified): 151.8. Weights run: 5.8s. Verify run: 16.8s. 946 eligible items had no known source.

Race sweep: night-elf 151.8 ± 0.2, human 150.7 ± 0.2, gnome 149.2 ± 0.2, dwarf 148.4 ± 0.2, high-order-skyborne 142.6 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.851 ± 0.070, crit=1.695 ± 0.082 per rating point (14 rating = 1%, 23.736 per %), hit=2.631 ± 0.217 per rating point (10 rating = 1%, 26.307 per %), melee_haste=15.626 ± 1.862

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.7 attack_power points (4.59 DPS) | yes | Chromite Barbute (8142, -0.32 DPS) [dungeon]; White Bandit Mask (10008, -0.68 DPS) [crafted]; Barbaric Iron Helm (7915, -1.39 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 20.4 attack_power points (1.88 DPS) | yes | Kaleidoscope Chain (13084, -0.46 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.59 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.77 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.4 attack_power points (2.80 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.67 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.77 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.38 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.55 DPS) [quest]; Parachute Cloak (10518, -0.77 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Icemail Jerkin (1981, -0.61 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -0.94 DPS) [quest]; Carapace of Tuten'kash (10775, -1.51 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.4 attack_power points (1.98 DPS) | yes | Branded Leather Bracers (19508, -0.13 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.21 DPS) [world_drop]; Yorgen Bracers (13012, -0.36 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.07 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 attack_power points (3.20 DPS) | yes | Highlander's Chain Girdle (20089, -0.27 DPS) [rep]; Officer's Belt (250556, -0.33 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.43 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.88 DPS) | yes | Symbolic Legplates (14829, -0.08 DPS) [world_drop]; Triprunner Dungarees (9624, -0.25 DPS) [quest]; Basilisk Hide Pants (1718, -0.29 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.0 attack_power points (3.60 DPS) | yes | Blackforge Greaves (6423, -0.23 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.37 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.42 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.8 attack_power points (2.84 DPS) | yes | Ring of the Underwood (2951, -0.77 DPS) [world_drop]; Thunderbrow Ring (13097, -0.85 DPS) [world_drop]; Mark of Kern (2262, -1.00 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.7 attack_power points (2.09 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Mark of Kern (2262, -0.25 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Primitive Fishing Pole (276203, -0.89 DPS) [vendor]; Bonebiter (6830, -1.14 DPS) [quest]; Nightblade (1982, -3.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.20 DPS) [quest]; The Silencer (13138, -0.33 DPS) [world_drop]; Swiftwind (13038, -0.42 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 946, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 12323213132215201-050200000000000000-000000000000000000)

Set DPS (verified): 223.7. Weights run: 6.2s. Verify run: 29.8s. 1228 eligible items had no known source.

Race sweep: high-order-skyborne 223.7 ± 0.3, night-elf 223.7 ± 0.3, human 222.4 ± 0.3, dwarf 221.8 ± 0.3, gnome 220.9 ± 0.3 (dps); high-order-skyborne beats the guide race night-elf by 0.0 (0.01%) ± 0.4, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.680 ± 0.080, crit=1.973 ± 0.105 per rating point (14 rating = 1%, 27.621 per %), hit=2.191 ± 0.269 per rating point (10 rating = 1%, 21.909 per %), melee_haste=17.028 ± 2.345

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 85.5 attack_power points (9.21 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.08 DPS) [vendor]; Raging Berserker's Helm (7719, -3.44 DPS) [dungeon]; Embrace of the Lycan (9479, -4.04 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.8 attack_power points (3.43 DPS) | yes | Sentinel's Medallion (19539, -1.26 DPS) [rep]; Kaleidoscope Chain (13084, -1.84 DPS) [world_drop] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 41.9 attack_power points (4.51 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.27 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.48 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.11 DPS) [dungeon]; Dark Phantom Cape (13122, -1.11 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.45 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight's Plate Hauberk (220794, -0.50 DPS) [vendor]; Mixologist's Tunic (12793, -1.02 DPS) [dungeon]; Warbear Harness (15064, -1.08 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 attack_power points (3.58 DPS) | yes | Officer's Wristguards (250581, -0.12 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.37 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.44 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Officer's Gloves (250551, +0.00 DPS) [crafted]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 49.6 attack_power points (5.34 DPS) | yes | Prowler's Leather Waistguard (252473, -0.16 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.22 DPS) [rep]; Highlander's Leather Girdle (20115, -0.22 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | PvP rank 6 · Knight · Alliance [vendor] | 68.4 attack_power points (7.37 DPS) | yes | Infernal Trickster Leggings (17754, -1.39 DPS) [dungeon]; Stormshroud Pants (15057, -1.42 DPS) [crafted]; Gryphon Rider's Leggings (9652, -1.42 DPS) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 49.8 attack_power points (5.37 DPS) | yes | Prowler's Leather Boots (252468, -0.58 DPS) [crafted]; Skulker's Leather Boots (252469, -0.65 DPS) [crafted]; Officer's Sabatons (250561, -0.94 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.9 attack_power points (4.51 DPS) | yes | Masons Fraternity Ring (9533, -1.98 DPS) [quest]; Falcon's Hook (7552, -2.24 DPS) [dungeon]; Thunderbrow Ring (13097, -2.25 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 35.1 attack_power points (3.78 DPS) | yes | Masons Fraternity Ring (9533, -1.25 DPS) [quest]; Falcon's Hook (7552, -1.51 DPS) [dungeon]; Thunderbrow Ring (13097, -1.52 DPS) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.36 DPS) [world_drop]; Taran Icebreaker (2915, -0.94 DPS) [world_drop]; Stoneraven (13059, -3.16 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.60 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.60 DPS) [world_drop]; Guttbuster (13139, -0.66 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Icemail Jerkin; wrist: Deepfury Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1228, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 12323213132215201-050500000000000000-205000000000000000)

Set DPS (verified): 317.0. Weights run: 5.8s. Verify run: 110.8s. 2727 eligible items had no known source.

Race sweep: high-order-skyborne 317.2 ± 0.4, night-elf 317.0 ± 0.4, human 316.4 ± 0.4, gnome 315.9 ± 0.4, dwarf 313.1 ± 0.4 (dps); high-order-skyborne beats the guide race night-elf by 0.1 (0.04%) ± 0.5, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.318 ± 0.125, crit=3.311 ± 0.178 per rating point (14 rating = 1%, 46.352 per %), hit=3.633 ± 0.444 per rating point (10 rating = 1%, 36.327 per %), melee_haste=20.650 ± 4.338

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 201.4 attack_power points (18.56 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -7.07 DPS) [vendor]; Mask of the Unforgiven (13404, -7.59 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 72.4 attack_power points (6.67 DPS) | yes | Star of Mystaria (12103, -2.32 DPS) [dungeon]; Emberfury Talisman (12929, -2.40 DPS) [dungeon]; Skibi's Pendant (13089, -2.97 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 108.9 attack_power points (10.04 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, -2.63 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.88 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.73 DPS) [rep]; Blackveil Cape (11626, -1.83 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.32 DPS) [crafted]; Savage Gladiator Chain (11726, -3.18 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -4.58 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.53 DPS) [rep]; Vambraces of the Sadist (13400, -2.06 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Highlander's Plate Girdle (20041, -0.30 DPS) [rep]; Highlander's Chain Girdle (20043, -0.30 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -2.50 DPS) [vendor]; Marshal's Plate Legguards (231540, -3.79 DPS) [pvp]; Sentinel's Leather Pants (237818, -5.06 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 82.7 attack_power points (7.62 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.90 DPS) [dungeon]; Shadowcraft Boots (16711, -2.13 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -2.57 DPS) [vendor]; Painweaver Band (13098, -3.53 DPS) [dungeon]; Blackstone Ring (17713, -3.72 DPS) [dungeon] |
| finger2 | Magni's Will (12548) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Painweaver Band (13098, -0.18 DPS) [dungeon]; Blackstone Ring (17713, -0.37 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.02 DPS) [world_drop]; Bloodseeker (19107, -0.68 DPS) [quest]; Ancient Bone Bow (18680, -1.30 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; waist: Valiant's Waistguard; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Magni's Will; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2727, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 954.7. Weights run: 5.0s. Verify run: 94.7s. 2727 eligible items had no known source.

Race sweep: high-order-skyborne 958.6 ± 1.1, night-elf 954.7 ± 1.1, human 951.1 ± 1.1, gnome 948.8 ± 1.1, dwarf 943.3 ± 1.1 (dps); high-order-skyborne beats the guide race night-elf by 3.9 (0.40%) ± 1.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.181 ± 0.135, crit=2.886 ± 0.188 per rating point (14 rating = 1%, 40.401 per %), hit=6.261 ± 0.531 per rating point (10 rating = 1%, 62.605 per %), melee_haste=36.486 ± 5.287

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 245.6 attack_power points (48.04 DPS) | yes | Mask of the Unforgiven (13404, -15.65 DPS) [dungeon]; Stalwart Helm (250599, -18.39 DPS) [crafted]; Lieutenant Commander's Plate Helm (23314, -18.86 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.74 DPS) [dungeon]; Emberfury Talisman (12929, -5.09 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, -5.48 DPS) [world_drop]; Arcanoweave Cloak (272411, -5.48 DPS) [vendor]; Stalwart Cloak (272415, -5.48 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -5.07 DPS) [dungeon]; Icemail Jerkin (1981, -7.30 DPS) [world_drop] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 103.0 attack_power points (20.15 DPS) | yes | Slashclaw Bracers (13211, -4.92 DPS) [dungeon]; Blackmist Armguards (12966, -5.75 DPS) [dungeon]; Forest Stalker's Bracers (19587, -7.31 DPS) [rep] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 108.4 attack_power points (21.20 DPS) | yes | Sash of the Grand Hunt (22207, -2.56 DPS) [dungeon]; Brigam Girdle (13142, -2.93 DPS) [dungeon]; Valiant's Waistguard (272403, -4.94 DPS) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -3.31 DPS) [vendor]; Marshal's Plate Legguards (231540, -6.33 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 106.2 attack_power points (20.78 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.25 DPS) [dungeon]; Shadowcraft Boots (16711, -8.15 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.73 DPS) [dungeon]; Channeler's Ring (272406, -10.64 DPS) [vendor]; Magni's Will (12548, -11.97 DPS) [quest] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -2.06 DPS) [dungeon]; Channeler's Ring (272406, -5.97 DPS) [vendor]; Magni's Will (12548, -7.30 DPS) [quest] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -5.36 DPS) [world_drop]; Bloodseeker (19107, -5.87 DPS) [quest]; Guttbuster (13139, -7.60 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2727, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 12321200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 34.2. Weights run: 5.1s. Verify run: 16.2s. 371 eligible items had no known source.

Race sweep: windshaper-skyborne 34.5 ± 0.0, undead 34.2 ± 0.0, troll 33.5 ± 0.0, orc 33.4 ± 0.0, tauren 33.2 ± 0.0 (dps); windshaper-skyborne beats the guide race undead by 0.3 (0.87%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.327 ± 0.030, crit=0.718 ± 0.029 per rating point (14 rating = 1%, 10.052 per %), hit=1.042 ± 0.076 per rating point (10 rating = 1%, 10.422 per %), melee_haste=6.761 ± 0.877

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.96 DPS) | yes | Defender's Leather Hood (252447, -0.19 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.32 DPS) [crafted]; Brawler's Leather Hood (252504, -0.45 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.0 attack_power points (0.38 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, +0.00 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.00 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.00 DPS) [dungeon]; Cape of the Brotherhood (5193, -0.07 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (34.2 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.48 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Forest Leather Bracers (3202, -0.16 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (34.2 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.87 DPS) | yes | Cobrahn's Grasp (6460, -0.00 DPS) [dungeon]; Brawler's Leather Belt (252428, -0.23 DPS) [crafted]; Deviate Scale Belt (6468, -0.26 DPS) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (34.2 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.6 attack_power points (0.80 DPS) | yes | Feet of the Lynx (1121, -0.00 DPS) [world_drop]; Veteran's Boots (250503, -0.06 DPS) [crafted]; Defender's Leather Boots (252441, -0.32 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 13.3 attack_power points (0.64 DPS) | yes | Loop of Sacrifice (281673, -0.35 DPS) [quest]; Bounty Hunter's Ring (5351, -0.45 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; Bounty Hunter's Ring (5351, -0.19 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smite's Mighty Hammer (7230, +0.00 DPS) [dungeon]; Hammerbone (270018, -0.14 DPS) [quest]; Duskbringer (2205, -0.41 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) | Blacksmithing [crafted] | sim-verified (34.2 DPS) | yes | Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, +0.00 DPS) [world_drop]; Deadly Blunderbuss (4369, -0.06 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Forsaken Greataxe; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 371, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (undead, 12323213130000000-000000000000000000-000000000000000000)

Set DPS (verified): 68.2. Weights run: 5.9s. Verify run: 25.3s. 622 eligible items had no known source.

Race sweep: undead 68.2 ± 0.1, troll 67.2 ± 0.1, orc 67.1 ± 0.1, tauren 66.6 ± 0.1, windshaper-skyborne 65.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.503 ± 0.039, crit=1.126 ± 0.039 per rating point (14 rating = 1%, 15.761 per %), hit=1.383 ± 0.141 per rating point (10 rating = 1%, 13.835 per %), melee_haste=8.326 ± 1.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 31.5 attack_power points (1.83 DPS) | yes | Tusken Helm (6686, -0.32 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.44 DPS) [crafted]; Defender's Leather Helm (252455, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.0 attack_power points (0.81 DPS) | yes | Scout's Medallion (19537, -0.12 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.12 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 26.5 attack_power points (1.54 DPS) | yes | Barbaric Iron Shoulders (7913, -0.32 DPS) [crafted]; Barbaric Shoulders (5964, -0.52 DPS) [crafted]; Mantle of Thieves (2264, -0.67 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.26 DPS) [world_drop]; Wildhunter Cloak (16658, -0.38 DPS) [quest]; Construct Cloak (279848, -0.38 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Silvered Chain Shirt (250518, -0.06 DPS) [crafted]; Shining Silver Breastplate (2870, -0.12 DPS) [crafted]; Slayer's Surcoat (14751, -0.32 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.0 attack_power points (0.99 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Pugilist Bracers (4438, -0.06 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.0 attack_power points (1.33 DPS) | yes | Mail Combat Gauntlets (4075, -0.06 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.12 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.0 attack_power points (1.68 DPS) | yes | Prowler's Leather Belt (252459, -0.12 DPS) [crafted]; Skulker's Leather Belt (252520, -0.20 DPS) [crafted]; Girdle of Golem Strength (9405, -0.29 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.5 attack_power points (1.89 DPS) | yes | Glimmering Mail Legguards (6386, -0.35 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.37 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -0.38 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.0 attack_power points (1.04 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Veteran's Boots (250503, -0.12 DPS) [crafted]; Stomping Boots (3741, -0.23 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.0 attack_power points (1.22 DPS) | yes | Band of the Fist (17694, -0.41 DPS) [quest]; Tiger Band (6749, -0.64 DPS) [quest]; Bonebinder's Signet (273032, -0.64 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.5 attack_power points (1.19 DPS) | yes | Band of the Fist (17694, -0.43 DPS, sim-verified) [quest]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.59 DPS) [dungeon]; Felbough Arm (273037, -0.75 DPS) [world]; Viscous Hammer (13045, -4.01 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.5 attack_power points (0.61 DPS) | yes | Golemsight Long Gun (273029, -0.09 DPS) [dungeon]; Gravespike Repeater (273033, -0.15 DPS) [world]; Silver Star (3463, -0.26 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 622, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 12323213132214001-000000000000000000-000000000000000000)

Set DPS (verified): 150.7. Weights run: 5.8s. Verify run: 17.9s. 928 eligible items had no known source.

Race sweep: undead 150.7 ± 0.2, orc 148.6 ± 0.2, troll 147.9 ± 0.2, tauren 147.7 ± 0.2, windshaper-skyborne 141.8 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.851 ± 0.070, crit=1.695 ± 0.082 per rating point (14 rating = 1%, 23.736 per %), hit=2.631 ± 0.217 per rating point (10 rating = 1%, 26.307 per %), melee_haste=15.626 ± 1.862

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 49.7 attack_power points (4.59 DPS) | yes | Chromite Barbute (8142, -0.32 DPS) [dungeon]; White Bandit Mask (10008, -0.68 DPS) [crafted]; Barbaric Iron Helm (7915, -1.39 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 20.4 attack_power points (1.88 DPS) | yes | Ethereal Talisman (4430, -0.27 DPS) [quest]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.59 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.4 attack_power points (2.80 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.67 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.77 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.38 DPS) [dungeon]; Parachute Cloak (10518, -0.38 DPS) [crafted]; Tigerstrike Mantle (13108, -0.38 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, -1.81 DPS) [dungeon]; Kolkar Marauder Chain (6773, -2.75 DPS) [quest]; Carapace of Tuten'kash (10775, -3.32 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.4 attack_power points (1.98 DPS) | yes | Branded Leather Bracers (19508, -0.13 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.21 DPS) [world_drop]; Yorgen Bracers (13012, -0.36 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -2.00 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 attack_power points (3.20 DPS) | yes | Defiler's Chain Girdle (20153, -0.27 DPS) [rep]; Officer's Belt (250556, -0.33 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.43 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.88 DPS) | yes | Symbolic Legplates (14829, -0.08 DPS) [world_drop]; Triprunner Dungarees (9624, -0.25 DPS) [quest]; Basilisk Hide Pants (1718, -0.29 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.0 attack_power points (3.60 DPS) | yes | Blackforge Greaves (6423, -0.23 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.37 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.42 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.8 attack_power points (2.84 DPS) | yes | Ring of the Underwood (2951, -0.77 DPS) [world_drop]; Thunderbrow Ring (13097, -0.85 DPS) [world_drop]; Mark of Kern (2262, -1.00 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 22.7 attack_power points (2.09 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Mark of Kern (2262, -0.25 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.89 DPS) [vendor]; X'caliboar (10758, -3.14 DPS) [dungeon]; Nightblade (1982, -4.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.20 DPS) [quest]; The Silencer (13138, -0.33 DPS) [world_drop]; Swiftwind (13038, -0.42 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 928, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 12323213132215201-050200000000000000-000000000000000000)

Set DPS (verified): 225.2. Weights run: 6.2s. Verify run: 21.7s. 1209 eligible items had no known source.

Race sweep: undead 225.2 ± 0.3, windshaper-skyborne 224.7 ± 0.3, troll 223.2 ± 0.3, orc 222.8 ± 0.3, tauren 221.3 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.468 ± 0.075, crit=1.699 ± 0.098 per rating point (14 rating = 1%, 23.783 per %), hit=2.042 ± 0.175 per rating point (10 rating = 1%, 20.418 per %), melee_haste=13.405 ± 2.213

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 80.2 attack_power points (7.67 DPS) | yes | Blood Guard's Plate Helm (220803, -0.96 DPS) [vendor]; Embrace of the Lycan (9479, -3.08 DPS) [dungeon]; Raging Berserker's Helm (7719, -3.09 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.1 attack_power points (2.78 DPS) | yes | Woven Ivy Necklace (19159, -0.37 DPS) [quest]; Scout's Medallion (19535, -1.10 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 40.4 attack_power points (3.87 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.26 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.45 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.01 DPS) [dungeon]; Dark Phantom Cape (13122, -1.01 DPS) [world_drop]; Pridelord Cape (14673, -1.28 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stone Guard's Plate Armor (220801, -0.18 DPS) [vendor]; Mixologist's Tunic (12793, -0.48 DPS) [dungeon]; Warbear Harness (15064, -0.69 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 30.8 attack_power points (2.95 DPS) | yes | Deepfury Bracers (13120, -0.08 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.24 DPS) [crafted]; Runed Golem Shackles (12550, -0.27 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -2.76 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 46.0 attack_power points (4.40 DPS) | yes | Defiler's Lamellar Girdle (20179, -0.02 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.04 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.14 DPS) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 6 · Stone Guard · Horde [vendor] | 62.5 attack_power points (5.98 DPS) | yes | Serpentskin Leggings (8262, -1.24 DPS) [world_drop]; Stormshroud Pants (15057, -1.43 DPS) [crafted]; Infernal Trickster Leggings (17754, -1.79 DPS, sim-verified) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 47.1 attack_power points (4.51 DPS) | yes | Prowler's Leather Boots (252468, -0.47 DPS) [crafted]; Skulker's Leather Boots (252469, -0.57 DPS) [crafted]; Officer's Sabatons (250561, -0.75 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.4 attack_power points (3.87 DPS) | yes | White Bone Band (11862, -1.57 DPS) [quest]; Masons Fraternity Ring (9533, -1.90 DPS) [quest]; Thunderbrow Ring (13097, -1.92 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 33.2 attack_power points (3.18 DPS) | yes | White Bone Band (11862, -0.88 DPS) [quest]; Masons Fraternity Ring (9533, -1.21 DPS) [quest]; Thunderbrow Ring (13097, -1.23 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.90 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.50 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.75 DPS) [world_drop]; Taran Icebreaker (2915, -0.84 DPS) [world_drop]; Ragehammer (10626, -2.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.21 DPS) [dungeon]; Houndmaster's Bow (11628, -0.40 DPS) [dungeon]; Skull Splitting Crossbow (13039, -1.31 DPS, sim-verified) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; wrist: Officer's Wristguards; waist: Girdle of Beastial Fury; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1209, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 12323213132215201-050500000000000000-205000000000000000)

Set DPS (verified): 317.1. Weights run: 5.8s. Verify run: 111.2s. 2717 eligible items had no known source.

Race sweep: undead 317.1 ± 0.4, windshaper-skyborne 315.7 ± 0.4, troll 312.9 ± 0.4, orc 312.2 ± 0.4, tauren 310.5 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.318 ± 0.125, crit=3.311 ± 0.178 per rating point (14 rating = 1%, 46.352 per %), hit=3.633 ± 0.444 per rating point (10 rating = 1%, 36.327 per %), melee_haste=20.650 ± 4.338

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 201.4 attack_power points (18.56 DPS) | yes | Champion's Plate Helm (227043, -7.07 DPS) [pvp]; Mask of the Unforgiven (13404, -7.59 DPS) [dungeon]; Fury Visor (20521, -7.62 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 72.4 attack_power points (6.67 DPS) | yes | Star of Mystaria (12103, -2.32 DPS) [dungeon]; Emberfury Talisman (12929, -2.40 DPS) [dungeon]; Skibi's Pendant (13089, -2.97 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 108.9 attack_power points (10.04 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -2.63 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.88 DPS) [dungeon]; Deathguard's Cloak (20068, -1.73 DPS) [rep]; Blackveil Cape (11626, -1.83 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -1.85 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -3.26 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.53 DPS) [rep]; Vambraces of the Sadist (13400, -2.06 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.30 DPS) [rep]; Defiler's Leather Girdle (20190, -0.30 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Plate Legguards (22651, -2.32 DPS) [rep]; Legionnaire's Plate Leggings (227048, -4.34 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 82.7 attack_power points (7.62 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.90 DPS) [dungeon]; Shadowcraft Boots (16711, -2.13 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.53 DPS) [dungeon]; Blackstone Ring (17713, -3.72 DPS) [dungeon]; Legionnaire's Band (19510, -4.35 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.96 DPS) [dungeon]; Blackstone Ring (17713, -1.15 DPS) [dungeon]; Legionnaire's Band (19510, -1.78 DPS) [rep] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Bloodseeker (19107, -0.66 DPS) [quest]; Ancient Bone Bow (18680, -1.28 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2717, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 951.2. Weights run: 5.0s. Verify run: 95.8s. 2717 eligible items had no known source.

Race sweep: windshaper-skyborne 957.9 ± 1.1, undead 951.2 ± 1.1, orc 951.0 ± 1.1, troll 948.4 ± 1.1, tauren 941.8 ± 1.1 (dps); windshaper-skyborne beats the guide race undead by 6.6 (0.70%) ± 1.6, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.181 ± 0.135, crit=2.886 ± 0.188 per rating point (14 rating = 1%, 40.401 per %), hit=6.261 ± 0.531 per rating point (10 rating = 1%, 62.605 per %), melee_haste=36.486 ± 5.287

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 245.6 attack_power points (48.04 DPS) | yes | Mask of the Unforgiven (13404, -15.65 DPS) [dungeon]; Stalwart Helm (250599, -18.39 DPS) [crafted]; Champion's Plate Helm (227043, -18.86 DPS) [pvp] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.74 DPS) [dungeon]; Emberfury Talisman (12929, -5.09 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Howler's Furs (272414, +0.00 DPS) [vendor]; Stalwart Cloak (272415, +0.00 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -5.07 DPS) [dungeon]; Icemail Jerkin (1981, -7.30 DPS) [world_drop] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 103.0 attack_power points (20.15 DPS) | yes | Slashclaw Bracers (13211, -4.92 DPS) [dungeon]; Blackmist Armguards (12966, -5.75 DPS) [dungeon]; Forest Stalker's Bracers (19587, -7.31 DPS) [rep] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 108.4 attack_power points (21.20 DPS) | yes | Sash of the Grand Hunt (22207, -2.56 DPS) [dungeon]; Brigam Girdle (13142, -2.93 DPS) [dungeon]; Valiant's Waistguard (272403, -4.94 DPS) [vendor] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Plate Legguards (22651, +0.00 DPS) [rep]; Outrider's Chain Leggings (22673, +0.00 DPS) [rep]; General's Plate Leggings (231533, +0.00 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 106.2 attack_power points (20.78 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.25 DPS) [dungeon]; Shadowcraft Boots (16711, -8.15 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.73 DPS) [dungeon]; Channeler's Ring (272406, -10.64 DPS) [vendor]; Painweaver Band (13098, -12.64 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -2.06 DPS) [dungeon]; Channeler's Ring (272406, -5.97 DPS) [vendor]; Painweaver Band (13098, -7.97 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Bloodseeker (19107, -0.51 DPS) [quest]; Guttbuster (13139, -2.23 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Devilsaur Gauntlets; waist: Marksman's Girdle; legs: Devilsaur Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2717, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

