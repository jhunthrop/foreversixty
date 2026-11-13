# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 12321200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 31.3. Weights run: 4.9s. Verify run: 16.4s. 388 eligible items had no known source.

Race sweep: high-order-skyborne 31.5 ± 0.0, night-elf 31.3 ± 0.0, human 31.2 ± 0.0, dwarf 30.9 ± 0.0, gnome 30.9 ± 0.0 (dps); high-order-skyborne beats the guide race night-elf by 0.2 (0.59%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.944 ± 0.023, crit=0.534 ± 0.021 per rating point (14 rating = 1%, 7.475 per %), hit=0.610 ± 0.083 per rating point (10 rating = 1%, 6.099 per %), melee_haste=7.518 ± 0.704

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.55 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.7 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS) [quest]; Tarnished Locket (279870, -0.51 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Dark Leather Cloak (2316) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lambent Scale Cloak (4706, +0.00 DPS) [world_drop]; Grave Shroud (279865, -0.04 DPS) [quest]; Glowing Lizardscale Cloak (6449, -0.05 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.44 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.05 DPS) [quest]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Cobrahn's Grasp (6460, -0.05 DPS) [dungeon]; Brawler's Leather Belt (252428, -0.27 DPS) [crafted]; Veteran's Chain Belt (250513, -0.32 DPS) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chausses of Westfall (6087, +0.00 DPS) [quest]; Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.7 attack_power points (0.65 DPS) | yes | Veteran's Boots (250503, -0.04 DPS) [crafted]; Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.8 attack_power points (0.52 DPS) | yes | Ring of the Moon (12052, -0.43 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.35 DPS) | yes | Ring of the Moon (12052, -0.26 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 301.8 attack_power points (13.26 DPS) | yes | Duskbringer (2205, -0.32 DPS) [dungeon]; Rockslicer (872, -0.53 DPS) [dungeon]; Sword of the Fallen (286977, -0.61 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, +0.00 DPS) [vendor]; Cracked Blacksmith Hammer (285279, +0.00 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Dark Leather Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 30 (night-elf, 12323213130000000-000000000000000000-000000000000000000)

Set DPS (verified): 65.4. Weights run: 5.9s. Verify run: 24.8s. 637 eligible items had no known source.

Race sweep: night-elf 65.4 ± 0.1, human 64.6 ± 0.1, gnome 63.9 ± 0.1, dwarf 63.6 ± 0.1, high-order-skyborne 63.1 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.645 ± 0.044, crit=1.232 ± 0.044 per rating point (14 rating = 1%, 17.252 per %), hit=1.553 ± 0.160 per rating point (10 rating = 1%, 15.535 per %), melee_haste=9.321 ± 1.308

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 32.8 attack_power points (1.69 DPS) | yes | Tusken Helm (6686, -0.36 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.45 DPS) [crafted]; Defender's Leather Helm (252455, -0.45 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.6 attack_power points (0.75 DPS) | yes | Sentinel's Medallion (19541, -0.07 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.13 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.1 attack_power points (1.45 DPS) | yes | Barbaric Iron Shoulders (7913, -0.32 DPS) [crafted]; Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Mantle of Thieves (2264, -0.60 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, -0.15 DPS) [pvp]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Silvered Chain Shirt (250518, -0.01 DPS) [crafted]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.25 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.9 attack_power points (0.92 DPS) | yes | Yorgen Bracers (13012, -0.05 DPS) [world_drop]; Pugilist Bracers (4438, -0.10 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.9 attack_power points (1.23 DPS) | yes | Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.10 DPS) [world]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.9 attack_power points (1.54 DPS) | yes | Prowler's Leather Belt (252459, -0.10 DPS) [crafted]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted]; Girdle of Golem Strength (9405, -0.30 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 33.5 attack_power points (1.73 DPS) | yes | Glimmering Mail Legguards (6386, -0.31 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.34 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.35 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 19.2 attack_power points (0.99 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Alacritous Treads (277234, -0.05 DPS) [quest]; Veteran's Boots (250503, -0.13 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.9 attack_power points (1.13 DPS) | yes | Seal of Wrynn (2933, -0.56 DPS) [quest]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (1.08 DPS) | yes | Tiger Band (6749, -0.56 DPS) [quest]; Bonebinder's Signet (273032, -0.56 DPS) [world]; Seal of Wrynn (2933, -0.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.54 DPS) [world]; Corpsemaker (6687, -0.56 DPS) [dungeon]; Viscous Hammer (13045, -4.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.9 attack_power points (0.56 DPS) | yes | Golemsight Long Gun (273029, -0.05 DPS) [dungeon]; Gravespike Repeater (273033, -0.15 DPS) [world]; Silver Star (3463, -0.22 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 637, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 12323213132214001-000000000000000000-000000000000000000)

Set DPS (verified): 148.1. Weights run: 5.8s. Verify run: 16.7s. 946 eligible items had no known source.

Race sweep: night-elf 148.1 ± 0.2, human 147.0 ± 0.2, gnome 145.6 ± 0.2, dwarf 144.7 ± 0.2, high-order-skyborne 139.0 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.961 ± 0.076, crit=1.795 ± 0.088 per rating point (14 rating = 1%, 25.125 per %), hit=2.823 ± 0.234 per rating point (10 rating = 1%, 28.233 per %), melee_haste=16.897 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 51.1 attack_power points (4.39 DPS) | yes | Chromite Barbute (8142, -0.31 DPS) [dungeon]; White Bandit Mask (10008, -0.65 DPS) [crafted]; Barbaric Iron Helm (7915, -1.33 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 21.6 attack_power points (1.85 DPS) | yes | Kaleidoscope Chain (13084, -0.49 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.65 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.82 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.6 attack_power points (2.71 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.67 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.82 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.35 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.51 DPS) [quest]; Parachute Cloak (10518, -0.69 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Icemail Jerkin (1981, -0.51 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -0.99 DPS) [quest]; Carapace of Tuten'kash (10775, -1.51 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.8 attack_power points (1.87 DPS) | yes | Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop]; Yorgen Bracers (13012, -0.34 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.16 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.7 attack_power points (3.06 DPS) | yes | Highlander's Chain Girdle (20089, -0.22 DPS) [rep]; Officer's Belt (250556, -0.33 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.48 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.60 DPS) | yes | Symbolic Legplates (14829, -0.02 DPS) [world_drop]; Triprunner Dungarees (9624, -0.06 DPS) [quest]; Basilisk Hide Pants (1718, -0.07 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.7 attack_power points (3.41 DPS) | yes | Blackforge Greaves (6423, -0.18 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.36 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 31.7 attack_power points (2.72 DPS) | yes | Ring of the Underwood (2951, -0.69 DPS) [world_drop]; Thunderbrow Ring (13097, -0.84 DPS) [world_drop]; Mark of Kern (2262, -1.00 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.7 attack_power points (2.03 DPS) | yes | Ring of the Underwood (2951, -0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Mark of Kern (2262, -0.31 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Primitive Fishing Pole (276203, -0.74 DPS) [vendor]; Bonebiter (6830, -1.06 DPS) [quest]; Nightblade (1982, -2.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.17 DPS) [quest]; The Silencer (13138, -0.33 DPS) [world_drop]; Swiftwind (13038, -0.36 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 946, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 12323213132215201-050200000000000000-000000000000000000)

Set DPS (verified): 219.7. Weights run: 6.1s. Verify run: 32.7s. 1228 eligible items had no known source.

Race sweep: high-order-skyborne 219.7 ± 0.3, night-elf 219.7 ± 0.3, human 218.4 ± 0.3, dwarf 217.8 ± 0.3, gnome 217.0 ± 0.3 (dps); high-order-skyborne beats the guide race night-elf by 0.0 (0.02%) ± 0.4, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.673 ± 0.082, crit=1.915 ± 0.108 per rating point (14 rating = 1%, 26.812 per %), hit=2.416 ± 0.278 per rating point (10 rating = 1%, 24.160 per %), melee_haste=17.220 ± 2.448

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 87.0 attack_power points (8.96 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.03 DPS) [vendor]; Raging Berserker's Helm (7719, -3.17 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -4.01 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.8 attack_power points (3.27 DPS) | yes | Sentinel's Medallion (19539, -1.20 DPS) [rep]; Kaleidoscope Chain (13084, -1.76 DPS) [world_drop] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 44.2 attack_power points (4.55 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.49 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.70 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.06 DPS) [dungeon]; Dark Phantom Cape (13122, -1.06 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.38 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight's Plate Hauberk (220794, -0.77 DPS) [vendor]; Mixologist's Tunic (12793, -1.19 DPS) [dungeon]; Warbear Harness (15064, -1.26 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.1 attack_power points (3.41 DPS) | yes | Officer's Wristguards (250581, -0.11 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.35 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.42 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -2.99 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 48.8 attack_power points (5.03 DPS) | yes | Prowler's Leather Waistguard (252473, -0.08 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.14 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.21 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | PvP rank 6 · Knight · Alliance [vendor] | 67.5 attack_power points (6.96 DPS) | yes | Infernal Trickster Leggings (17754, -1.02 DPS, sim-verified) [dungeon]; Gryphon Rider's Leggings (9652, -1.28 DPS) [quest]; Centurion Legplates (10740, -1.28 DPS) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 49.8 attack_power points (5.13 DPS) | yes | Prowler's Leather Boots (252468, -0.55 DPS) [crafted]; Skulker's Leather Boots (252469, -0.62 DPS) [crafted]; Officer's Sabatons (250561, -0.90 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.2 attack_power points (4.55 DPS) | yes | Masons Fraternity Ring (9533, -2.14 DPS) [quest]; Falcon's Hook (7552, -2.38 DPS) [dungeon]; Thunderbrow Ring (13097, -2.38 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 35.1 attack_power points (3.61 DPS) | yes | Masons Fraternity Ring (9533, -1.20 DPS) [quest]; Falcon's Hook (7552, -1.44 DPS) [dungeon]; Thunderbrow Ring (13097, -1.45 DPS) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Molten Heart of the Mountain (249470, -0.77 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, -0.90 DPS) [world_drop]; Blight (7959, -1.56 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.53 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop]; Guttbuster (13139, -0.59 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Icemail Jerkin; wrist: Deepfury Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1228, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 12323213132215201-050500000000000000-205000000000000000)

Set DPS (verified): 312.4. Weights run: 5.7s. Verify run: 104.9s. 2727 eligible items had no known source.

Race sweep: high-order-skyborne 312.5 ± 0.4, night-elf 312.4 ± 0.4, human 311.7 ± 0.4, gnome 311.3 ± 0.4, dwarf 308.5 ± 0.4 (dps); high-order-skyborne beats the guide race night-elf by 0.2 (0.05%) ± 0.5, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.451 ± 0.133, crit=3.502 ± 0.190 per rating point (14 rating = 1%, 49.026 per %), hit=3.871 ± 0.474 per rating point (10 rating = 1%, 38.707 per %), melee_haste=22.113 ± 4.634

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 211.5 attack_power points (18.27 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -7.06 DPS) [vendor]; Mask of the Unforgiven (13404, -7.35 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 75.0 attack_power points (6.48 DPS) | yes | Star of Mystaria (12103, -2.13 DPS) [dungeon]; Emberfury Talisman (12929, -2.25 DPS) [dungeon]; Skibi's Pendant (13089, -2.86 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 115.2 attack_power points (9.95 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Plate Shoulderguards (231537, -0.11 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -2.78 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.86 DPS) [dungeon]; Blackveil Cape (11626, -1.76 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.77 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.67 DPS) [crafted]; Savage Gladiator Chain (11726, -3.33 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -4.70 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.66 DPS) [rep]; Vambraces of the Sadist (13400, -2.13 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Highlander's Plate Girdle (20041, -0.33 DPS) [rep]; Highlander's Chain Girdle (20043, -0.33 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -2.75 DPS) [vendor]; Marshal's Plate Legguards (231540, -3.96 DPS) [pvp]; Sentinel's Leather Pants (237818, -5.04 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 87.7 attack_power points (7.58 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.99 DPS) [dungeon]; Shadowcraft Boots (16711, -2.13 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -2.48 DPS) [vendor]; Painweaver Band (13098, -3.52 DPS) [dungeon]; Blackstone Ring (17713, -3.72 DPS) [dungeon] |
| finger2 | Magni's Will (12548) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Painweaver Band (13098, -0.17 DPS) [dungeon]; Blackstone Ring (17713, -0.37 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.04 DPS) [world_drop]; Bloodseeker (19107, -0.78 DPS) [quest]; Ancient Bone Bow (18680, -1.32 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; waist: Valiant's Waistguard; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Magni's Will; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2727, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 935.3. Weights run: 5.0s. Verify run: 104.3s. 2727 eligible items had no known source.

Race sweep: high-order-skyborne 938.8 ± 1.0, night-elf 935.3 ± 1.1, human 932.7 ± 1.1, gnome 930.6 ± 1.0, dwarf 924.0 ± 1.0 (dps); high-order-skyborne beats the guide race night-elf by 3.6 (0.38%) ± 1.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.594 ± 0.133, crit=3.270 ± 0.184 per rating point (14 rating = 1%, 45.774 per %), hit=5.803 ± 0.555 per rating point (10 rating = 1%, 58.034 per %), melee_haste=31.753 ± 5.081

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 247.2 attack_power points (45.57 DPS) | yes | Mask of the Unforgiven (13404, -15.74 DPS) [dungeon]; Lieutenant Commander's Plate Helm (23314, -17.92 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -2.53 DPS) [dungeon]; Emberfury Talisman (12929, -4.79 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 115.8 attack_power points (21.35 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Truestrike Shoulders (12927, -0.88 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -5.00 DPS) [dungeon]; Cloak of Flames (3475, -5.16 DPS) [world_drop]; Stalwart Cloak (272415, -5.16 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -4.59 DPS) [dungeon]; Icemail Jerkin (1981, -8.25 DPS) [world_drop] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 103.8 attack_power points (19.13 DPS) | yes | Slashclaw Bracers (13211, -5.09 DPS) [dungeon]; Forest Stalker's Bracers (19587, -5.59 DPS) [rep]; Blackmist Armguards (12966, -6.41 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 112.5 attack_power points (20.74 DPS) | yes | Sash of the Grand Hunt (22207, -2.87 DPS) [dungeon]; Valiant's Waistguard (272403, -3.59 DPS) [vendor]; Brigam Girdle (13142, -4.37 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -5.79 DPS) [vendor]; Marshal's Plate Legguards (231540, -8.63 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 109.9 attack_power points (20.26 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.09 DPS) [dungeon]; Shadowcraft Boots (16711, -7.01 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -7.33 DPS) [dungeon]; Magni's Will (12548, -10.44 DPS) [quest]; Channeler's Ring (272406, -11.02 DPS) [vendor] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.01 DPS) [dungeon]; Magni's Will (12548, -6.12 DPS) [quest]; Channeler's Ring (272406, -6.69 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -3.50 DPS) [world_drop]; Bloodseeker (19107, -4.47 DPS) [quest]; Ancient Bone Bow (18680, -5.80 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2727, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 12321200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 32.5. Weights run: 4.9s. Verify run: 20.3s. 371 eligible items had no known source.

Race sweep: windshaper-skyborne 32.8 ± 0.0, undead 32.5 ± 0.0, troll 31.8 ± 0.0, orc 31.6 ± 0.0, tauren 31.5 ± 0.0 (dps); windshaper-skyborne beats the guide race undead by 0.3 (0.92%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.944 ± 0.023, crit=0.534 ± 0.021 per rating point (14 rating = 1%, 7.475 per %), hit=0.610 ± 0.083 per rating point (10 rating = 1%, 6.099 per %), melee_haste=7.518 ± 0.704

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.55 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.7 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Subterranean Cape (14149, -0.09 DPS) [dungeon]; Grave Shroud (279865, -0.09 DPS) [quest] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.94 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.44 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.18 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.02 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.79 DPS) | yes | Brawler's Leather Belt (252428, -0.27 DPS) [crafted]; Veteran's Chain Belt (250513, -0.32 DPS) [crafted]; Cobrahn's Grasp (6460, -1.26 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.90 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.7 attack_power points (0.65 DPS) | yes | Veteran's Boots (250503, -0.04 DPS) [crafted]; Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.8 attack_power points (0.52 DPS) | yes | Loop of Sacrifice (281673, -0.25 DPS) [quest]; Bounty Hunter's Ring (5351, -0.39 DPS) [quest]; Ring of the Moon (12052, -0.43 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; Bounty Hunter's Ring (5351, -0.23 DPS) [quest]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (13.32 DPS) | yes | Hammerbone (270018, -0.12 DPS) [quest]; Duskbringer (2205, -0.38 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.01 DPS) [world_drop]; Heavy Shortbow (3036, -0.09 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Forsaken Greataxe; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 371, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (undead, 12323213130000000-000000000000000000-000000000000000000)

Set DPS (verified): 65.5. Weights run: 5.9s. Verify run: 24.4s. 622 eligible items had no known source.

Race sweep: undead 65.5 ± 0.1, troll 64.4 ± 0.1, orc 64.3 ± 0.1, tauren 63.8 ± 0.1, windshaper-skyborne 63.1 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.645 ± 0.044, crit=1.232 ± 0.044 per rating point (14 rating = 1%, 17.252 per %), hit=1.553 ± 0.160 per rating point (10 rating = 1%, 15.535 per %), melee_haste=9.321 ± 1.308

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 32.8 attack_power points (1.69 DPS) | yes | Tusken Helm (6686, -0.35 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.45 DPS) [crafted]; Defender's Leather Helm (252455, -0.45 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.6 attack_power points (0.75 DPS) | yes | Scout's Medallion (19537, -0.07 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.13 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.1 attack_power points (1.45 DPS) | yes | Barbaric Iron Shoulders (7913, -0.32 DPS) [crafted]; Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Mantle of Thieves (2264, -0.60 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wildhunter Cloak (16658, -0.39 DPS) [quest]; Construct Cloak (279848, -0.39 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Silvered Chain Shirt (250518, -0.01 DPS) [crafted]; Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.25 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.9 attack_power points (0.92 DPS) | yes | Yorgen Bracers (13012, -0.05 DPS) [world_drop]; Pugilist Bracers (4438, -0.10 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.9 attack_power points (1.23 DPS) | yes | Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.10 DPS) [world]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.9 attack_power points (1.54 DPS) | yes | Prowler's Leather Belt (252459, -0.10 DPS) [crafted]; Skulker's Leather Belt (252520, -0.16 DPS) [crafted]; Girdle of Golem Strength (9405, -0.30 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 33.5 attack_power points (1.73 DPS) | yes | Glimmering Mail Legguards (6386, -0.31 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.33 DPS, sim-verified) [crafted]; Brawler's Leather Pants (252500, -0.35 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 19.2 attack_power points (0.99 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Veteran's Boots (250503, -0.13 DPS) [crafted]; Stomping Boots (3741, -0.24 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.9 attack_power points (1.13 DPS) | yes | Band of the Fist (17694, -0.38 DPS) [quest]; Tiger Band (6749, -0.61 DPS) [quest]; Bonebinder's Signet (273032, -0.61 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.9 attack_power points (1.08 DPS) | yes | Band of the Fist (17694, -0.39 DPS, sim-verified) [quest]; Tiger Band (6749, -0.56 DPS) [quest]; Bonebinder's Signet (273032, -0.56 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Felbough Arm (273037, -0.54 DPS) [world]; Corpsemaker (6687, -0.56 DPS) [dungeon]; Viscous Hammer (13045, -4.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.9 attack_power points (0.56 DPS) | yes | Golemsight Long Gun (273029, -0.05 DPS) [dungeon]; Gravespike Repeater (273033, -0.15 DPS) [world]; Silver Star (3463, -0.22 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 622, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 12323213132214001-000000000000000000-000000000000000000)

Set DPS (verified): 147.1. Weights run: 5.8s. Verify run: 18.0s. 928 eligible items had no known source.

Race sweep: undead 147.1 ± 0.2, orc 144.9 ± 0.2, troll 144.3 ± 0.2, tauren 144.1 ± 0.2, windshaper-skyborne 138.3 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.961 ± 0.076, crit=1.795 ± 0.088 per rating point (14 rating = 1%, 25.125 per %), hit=2.823 ± 0.234 per rating point (10 rating = 1%, 28.233 per %), melee_haste=16.897 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 51.1 attack_power points (4.39 DPS) | yes | Chromite Barbute (8142, -0.31 DPS) [dungeon]; White Bandit Mask (10008, -0.65 DPS) [crafted]; Barbaric Iron Helm (7915, -1.33 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 21.6 attack_power points (1.85 DPS) | yes | Ethereal Talisman (4430, -0.32 DPS) [quest]; Kaleidoscope Chain (13084, -0.49 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.65 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.6 attack_power points (2.71 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.67 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.82 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.35 DPS) [dungeon]; Parachute Cloak (10518, -0.35 DPS) [crafted]; Tigerstrike Mantle (13108, -0.35 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, -1.89 DPS) [dungeon]; Kolkar Marauder Chain (6773, -2.88 DPS) [quest]; Carapace of Tuten'kash (10775, -3.40 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.8 attack_power points (1.87 DPS) | yes | Branded Leather Bracers (19508, -0.16 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop]; Yorgen Bracers (13012, -0.34 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -2.08 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.7 attack_power points (3.06 DPS) | yes | Defiler's Chain Girdle (20153, -0.22 DPS) [rep]; Officer's Belt (250556, -0.33 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.48 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.60 DPS) | yes | Symbolic Legplates (14829, -0.02 DPS) [world_drop]; Triprunner Dungarees (9624, -0.06 DPS) [quest]; Basilisk Hide Pants (1718, -0.07 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.7 attack_power points (3.41 DPS) | yes | Blackforge Greaves (6423, -0.18 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.36 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 31.7 attack_power points (2.72 DPS) | yes | Ring of the Underwood (2951, -0.69 DPS) [world_drop]; Thunderbrow Ring (13097, -0.84 DPS) [world_drop]; Mark of Kern (2262, -1.00 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.7 attack_power points (2.03 DPS) | yes | Ring of the Underwood (2951, -0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Mark of Kern (2262, -0.31 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.74 DPS) [vendor]; X'caliboar (10758, -2.92 DPS) [dungeon]; Nightblade (1982, -3.87 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.17 DPS) [quest]; The Silencer (13138, -0.33 DPS) [world_drop]; Swiftwind (13038, -0.36 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 928, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 12323213132215201-050200000000000000-000000000000000000)

Set DPS (verified): 221.3. Weights run: 6.1s. Verify run: 34.4s. 1209 eligible items had no known source.

Race sweep: undead 221.3 ± 0.3, windshaper-skyborne 220.9 ± 0.3, troll 219.0 ± 0.3, orc 218.8 ± 0.3, tauren 217.5 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.548 ± 0.080, crit=1.789 ± 0.105 per rating point (14 rating = 1%, 25.051 per %), hit=2.183 ± 0.187 per rating point (10 rating = 1%, 21.832 per %), melee_haste=14.199 ± 2.360

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 82.9 attack_power points (7.45 DPS) | yes | Blood Guard's Plate Helm (220803, -0.90 DPS) [vendor]; Embrace of the Lycan (9479, -3.13 DPS) [dungeon]; Raging Berserker's Helm (7719, -3.38 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.1 attack_power points (2.71 DPS) | yes | Woven Ivy Necklace (19159, -0.38 DPS) [quest]; Scout's Medallion (19535, -1.04 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 41.8 attack_power points (3.76 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.31 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.49 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.94 DPS) [dungeon]; Dark Phantom Cape (13122, -0.94 DPS) [world_drop]; Pridelord Cape (14673, -1.23 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -0.67 DPS) [dungeon]; Warbear Harness (15064, -0.82 DPS) [crafted]; Stone Guard's Plate Armor (220801, -0.86 DPS, sim-verified) [vendor] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 31.3 attack_power points (2.81 DPS) | yes | Deepfury Bracers (13120, -0.01 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.22 DPS) [crafted]; Runed Golem Shackles (12550, -0.30 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -2.79 DPS, sim-verified) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 47.1 attack_power points (4.23 DPS) | yes | Prowler's Leather Waistguard (252473, -0.04 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.09 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 6 · Stone Guard · Horde [vendor] | 64.5 attack_power points (5.80 DPS) | yes | Infernal Trickster Leggings (17754, -1.15 DPS, sim-verified) [dungeon]; Serpentskin Leggings (8262, -1.24 DPS) [world_drop]; Stormshroud Pants (15057, -1.30 DPS) [crafted] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 48.1 attack_power points (4.32 DPS) | yes | Prowler's Leather Boots (252468, -0.46 DPS) [crafted]; Skulker's Leather Boots (252469, -0.54 DPS) [crafted]; Officer's Sabatons (250561, -0.74 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 attack_power points (3.76 DPS) | yes | White Bone Band (11862, -1.60 DPS) [quest]; Masons Fraternity Ring (9533, -1.81 DPS) [quest]; Thunderbrow Ring (13097, -1.90 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 33.9 attack_power points (3.05 DPS) | yes | White Bone Band (11862, -0.89 DPS) [quest]; Masons Fraternity Ring (9533, -1.10 DPS) [quest]; Thunderbrow Ring (13097, -1.19 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.50 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.41 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.32 DPS) [world_drop]; Taran Icebreaker (2915, -0.79 DPS) [world_drop]; Ragehammer (10626, -1.47 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.30 DPS) [dungeon]; Guttbuster (13139, -0.45 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.47 DPS, sim-verified) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; wrist: Officer's Wristguards; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1209, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 12323213132215201-050500000000000000-205000000000000000)

Set DPS (verified): 312.2. Weights run: 5.7s. Verify run: 105.5s. 2717 eligible items had no known source.

Race sweep: undead 312.2 ± 0.4, windshaper-skyborne 310.9 ± 0.4, troll 307.9 ± 0.4, orc 307.2 ± 0.4, tauren 305.6 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.451 ± 0.133, crit=3.502 ± 0.190 per rating point (14 rating = 1%, 49.026 per %), hit=3.871 ± 0.474 per rating point (10 rating = 1%, 38.707 per %), melee_haste=22.113 ± 4.634

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 211.5 attack_power points (18.27 DPS) | yes | Champion's Plate Helm (227043, -7.06 DPS) [pvp]; Mask of the Unforgiven (13404, -7.35 DPS) [dungeon]; Fury Visor (20521, -7.58 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 75.0 attack_power points (6.48 DPS) | yes | Star of Mystaria (12103, -2.13 DPS) [dungeon]; Emberfury Talisman (12929, -2.25 DPS) [dungeon]; Skibi's Pendant (13089, -2.86 DPS) [world_drop] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 115.2 attack_power points (9.95 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -0.11 DPS) [pvp]; Champion's Plate Shoulders (227042, -2.78 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.86 DPS) [dungeon]; Blackveil Cape (11626, -1.76 DPS) [dungeon]; Deathguard's Cloak (20068, -1.77 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -1.67 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -3.03 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -1.66 DPS) [rep]; Vambraces of the Sadist (13400, -2.13 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.33 DPS) [rep]; Defiler's Leather Girdle (20190, -0.33 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Plate Legguards (22651, -2.57 DPS) [rep]; Legionnaire's Plate Leggings (227048, -4.45 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 87.7 attack_power points (7.58 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.99 DPS) [dungeon]; Shadowcraft Boots (16711, -2.13 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.52 DPS) [dungeon]; Blackstone Ring (17713, -3.72 DPS) [dungeon]; Legionnaire's Band (19510, -4.39 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.04 DPS) [dungeon]; Blackstone Ring (17713, -1.24 DPS) [dungeon]; Legionnaire's Band (19510, -1.91 DPS) [rep] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Bloodseeker (19107, -0.74 DPS) [quest]; Ancient Bone Bow (18680, -1.27 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2717, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 930.3. Weights run: 5.0s. Verify run: 106.0s. 2717 eligible items had no known source.

Race sweep: windshaper-skyborne 935.1 ± 1.0, undead 930.3 ± 1.1, troll 928.3 ± 1.1, orc 928.2 ± 1.1, tauren 920.2 ± 1.1 (dps); windshaper-skyborne beats the guide race undead by 4.8 (0.52%) ± 1.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.594 ± 0.133, crit=3.270 ± 0.184 per rating point (14 rating = 1%, 45.774 per %), hit=5.803 ± 0.555 per rating point (10 rating = 1%, 58.034 per %), melee_haste=31.753 ± 5.081

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 247.2 attack_power points (45.57 DPS) | yes | Champion's Plate Helm (227043, -17.92 DPS) [pvp]; Fury Visor (20521, -19.13 DPS) [quest]; Mask of the Unforgiven (13404, -19.47 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nacreous Shell Necklace (22403, -2.53 DPS) [dungeon]; Emberfury Talisman (12929, -4.79 DPS) [dungeon]; Star of Mystaria (12103, -16.10 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 115.8 attack_power points (21.35 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.88 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, -0.16 DPS) [world_drop]; Stalwart Cloak (272415, -0.16 DPS) [vendor]; Howler's Furs (272414, -5.73 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -4.59 DPS) [dungeon]; Icemail Jerkin (1981, -8.25 DPS) [world_drop]; Dawn Armor (252483, -11.42 DPS, sim-verified) [crafted] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 103.8 attack_power points (19.13 DPS) | yes | Forest Stalker's Bracers (19587, -5.59 DPS) [rep]; Slashclaw Bracers (13211, -6.26 DPS, sim-verified) [dungeon]; Blackmist Armguards (12966, -6.41 DPS) [dungeon] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Gloves of Holy Might (867, -18.80 DPS, sim-verified) [world_drop] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 112.5 attack_power points (20.74 DPS) | yes | Sash of the Grand Hunt (22207, -2.87 DPS) [dungeon]; Valiant's Waistguard (272403, -3.59 DPS) [vendor]; Brigam Girdle (13142, -4.37 DPS) [dungeon] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Plate Legguards (22651, +0.00 DPS) [rep]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Chain Leggings (22673, -6.95 DPS, sim-verified) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 109.9 attack_power points (20.26 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.09 DPS) [dungeon]; Shadowcraft Boots (16711, -7.01 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -7.33 DPS) [dungeon]; Channeler's Ring (272406, -11.02 DPS) [vendor]; Naglering (11669, -27.32 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.01 DPS) [dungeon]; Channeler's Ring (272406, -6.69 DPS) [vendor]; Blessed Band of Light (272407, -6.69 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+11.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Molten Heart of the Mountain (249470, -8.47 DPS, sim-verified) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -12.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.97 DPS) [quest]; Ancient Bone Bow (18680, -2.31 DPS) [dungeon]; Blackcrow (12651, -7.16 DPS, sim-verified) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; hands: Devilsaur Gauntlets; waist: Marksman's Girdle; legs: Devilsaur Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2717, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

