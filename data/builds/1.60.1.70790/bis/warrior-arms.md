# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 12321200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 31.3. Weights run: 3.7s. Verify run: 12.3s. 388 eligible items had no known source.

Race sweep: high-order-skyborne 31.5 ± 0.0, night-elf 31.3 ± 0.0, human 31.2 ± 0.0, dwarf 30.9 ± 0.0, gnome 30.9 ± 0.0 (dps); high-order-skyborne beats the guide race night-elf by 0.2 (0.59%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.944 ± 0.023, crit=0.534 ± 0.021 per rating point (14 rating = 1%, 7.475 per %), hit=0.610 ± 0.083 per rating point (10 rating = 1%, 6.099 per %), melee_haste=7.518 ± 0.704

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.55 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 5.7 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS) [quest]; Tarnished Locket (279870, -0.51 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.26 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Dark Leather Cloak (2316) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lambent Scale Cloak (4706, +0.00 DPS) [world_drop]; Grave Shroud (279865, -0.04 DPS) [quest]; Glowing Lizardscale Cloak (6449, -0.05 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.44 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.05 DPS) [quest]; Cryptwalker Bracers (280095, -0.09 DPS) [quest] |
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

Set DPS (verified): 65.7. Weights run: 4.4s. Verify run: 16.9s. 637 eligible items had no known source.

Race sweep: night-elf 65.7 ± 0.1, human 65.0 ± 0.1, gnome 64.0 ± 0.1, dwarf 63.8 ± 0.1, high-order-skyborne 63.3 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.744 ± 0.052, crit=1.246 ± 0.054 per rating point (14 rating = 1%, 17.447 per %), hit=1.195 ± 0.162 per rating point (10 rating = 1%, 11.951 per %), melee_haste=9.762 ± 1.430

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 33.7 attack_power points (1.74 DPS) | yes | Tusken Helm (6686, -0.40 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.50 DPS) [crafted]; Defender's Leather Helm (252455, -0.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.0 attack_power points (0.77 DPS) | yes | Sentinel's Medallion (19541, -0.05 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.15 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.2 attack_power points (1.51 DPS) | yes | Barbaric Iron Shoulders (7913, -0.35 DPS) [crafted]; Barbaric Shoulders (5964, -0.54 DPS) [crafted]; Mantle of Thieves (2264, -0.61 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, -0.17 DPS) [pvp]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.40 DPS) [dungeon] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Avenger's Armor (1488, -0.02 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.23 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.5 attack_power points (0.95 DPS) | yes | Yorgen Bracers (13012, -0.06 DPS) [world_drop]; Pugilist Bracers (4438, -0.13 DPS) [dungeon]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 24.5 attack_power points (1.26 DPS) | yes | Mail Combat Gauntlets (4075, -0.08 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Gauntlets of Ogre Strength (3341, -0.13 DPS) [world] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 30.5 attack_power points (1.57 DPS) | yes | Prowler's Leather Belt (252459, -0.10 DPS) [crafted]; Skulker's Leather Belt (252520, -0.14 DPS) [crafted]; Girdle of Golem Strength (9405, -0.33 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 34.2 attack_power points (1.77 DPS) | yes | Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.31 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.34 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Alacritous Treads (277234, -0.06 DPS) [quest]; Veteran's Boots (250503, -0.15 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.5 attack_power points (1.16 DPS) | yes | Seal of Wrynn (2933, -0.58 DPS) [quest]; Ring of Precision (1491, -0.62 DPS) [dungeon]; Tiger Band (6749, -0.64 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.2 attack_power points (1.10 DPS) | yes | Ring of Precision (1491, -0.56 DPS) [dungeon]; Tiger Band (6749, -0.58 DPS) [quest]; Seal of Wrynn (2933, -0.59 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.59 DPS) [dungeon]; Felbough Arm (273037, -0.70 DPS) [world]; Viscous Hammer (13045, -5.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 11.2 attack_power points (0.58 DPS) | yes | Golemsight Long Gun (273029, -0.04 DPS) [dungeon]; Gravespike Repeater (273033, -0.17 DPS) [world]; Silver Star (3463, -0.22 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 637, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 12323213132214001-000000000000000000-000000000000000000)

Set DPS (verified): 149.4. Weights run: 4.5s. Verify run: 12.9s. 946 eligible items had no known source.

Race sweep: night-elf 149.4 ± 0.2, human 148.6 ± 0.2, gnome 147.5 ± 0.2, dwarf 145.9 ± 0.2, high-order-skyborne 140.7 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.122 ± 0.080, crit=2.065 ± 0.094 per rating point (14 rating = 1%, 28.915 per %), hit=2.785 ± 0.233 per rating point (10 rating = 1%, 27.847 per %), melee_haste=18.604 ± 2.089

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 54.9 attack_power points (4.82 DPS) | yes | Chromite Barbute (8142, -0.48 DPS) [dungeon]; White Bandit Mask (10008, -0.84 DPS) [crafted]; Barbaric Iron Helm (7915, -1.56 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.3 attack_power points (2.05 DPS) | yes | Kaleidoscope Chain (13084, -0.60 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.82 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.00 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 attack_power points (2.93 DPS) | yes | Flintrock Shoulders (7755, -0.19 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.76 DPS) [crafted]; Nightscape Shoulders (8192, -0.88 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.34 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.53 DPS) [quest]; Parachute Cloak (10518, -0.68 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kolkar Marauder Chain (6773, -1.18 DPS) [quest]; Carapace of Tuten'kash (10775, -1.70 DPS) [dungeon]; Wolffear Harness (13110, -1.78 DPS) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 22.5 attack_power points (1.97 DPS) | yes | Hawkeye's Bracers (14590, -0.15 DPS) [world_drop]; Branded Leather Bracers (19508, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.36 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Scarlet Gauntlets (10331, -3.19 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.1 attack_power points (3.26 DPS) | yes | Highlander's Chain Girdle (20089, -0.02 DPS) [rep]; Officer's Belt (250556, -0.38 DPS) [crafted]; Skulker's Leather Belt (252520, -0.53 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.6 attack_power points (3.91 DPS) | yes | Triprunner Dungarees (9624, -0.03 DPS) [quest]; Symbolic Legplates (14829, -0.16 DPS) [world_drop]; Scarlet Leggings (10330, -0.23 DPS) [dungeon] |
| feet | Boots of Avoidance (14549) | World drop [world_drop] | 41.7 attack_power points (3.66 DPS) | yes | Officer's Boots (250546, -0.08 DPS) [crafted]; Blackforge Greaves (6423, -0.22 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.38 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 33.0 attack_power points (2.90 DPS) | yes | Falcon's Hook (7552, -0.69 DPS) [dungeon]; Thunderbrow Ring (13097, -0.93 DPS) [world_drop]; Mark of Kern (2262, -1.14 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.2 attack_power points (2.21 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.25 DPS) [world_drop]; Mark of Kern (2262, -0.46 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.61 DPS) [vendor]; Bonebiter (6830, -1.08 DPS) [quest]; Nightblade (1982, -4.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Swiftwind (13038, -0.31 DPS) [world_drop]; The Silencer (13138, -0.38 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Boots of Avoidance; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 946, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 12323213132215201-050200000000000000-000000000000000000)

Set DPS (verified): 220.8. Weights run: 4.7s. Verify run: 13.7s. 1225 eligible items had no known source.

Race sweep: night-elf 220.8 ± 0.3, high-order-skyborne 220.6 ± 0.3, human 219.9 ± 0.3, dwarf 219.5 ± 0.3, gnome 218.0 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.659 ± 0.094, crit=1.930 ± 0.123 per rating point (14 rating = 1%, 27.025 per %), hit=2.647 ± 0.277 per rating point (10 rating = 1%, 26.475 per %), melee_haste=24.749 ± 2.820

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 89.5 attack_power points (9.18 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.03 DPS) [vendor]; Raging Berserker's Helm (7719, -3.17 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -4.26 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (3.24 DPS) | yes | Sentinel's Medallion (19540, -1.37 DPS) [rep]; Kaleidoscope Chain (13084, -1.74 DPS) [world_drop] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 46.5 attack_power points (4.77 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.74 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.94 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.06 DPS) [dungeon]; Dark Phantom Cape (13122, -1.06 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.36 DPS) [pvp] |
| chest | Knight's Plate Hauberk (220794) | PvP rank 6 · Knight · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -0.46 DPS) [dungeon]; Warbear Harness (15064, -0.53 DPS) [crafted]; Warforged Chestplate (11195, -0.93 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.9 attack_power points (3.37 DPS) | yes | Officer's Wristguards (250581, -0.10 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.34 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.41 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.61 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 49.0 attack_power points (5.03 DPS) | yes | Prowler's Leather Waistguard (252473, -0.11 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.18 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.21 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | PvP rank 6 · Knight · Alliance [vendor] | 67.6 attack_power points (6.93 DPS) | yes | Gryphon Rider's Leggings (9652, -1.31 DPS) [quest]; Centurion Legplates (10740, -1.31 DPS) [quest]; Infernal Trickster Leggings (17754, -1.83 DPS, sim-verified) [dungeon] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 49.6 attack_power points (5.08 DPS) | yes | Prowler's Leather Boots (252468, -0.55 DPS) [crafted]; Skulker's Leather Boots (252469, -0.62 DPS) [crafted]; Officer's Sabatons (250561, -0.89 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.5 attack_power points (4.77 DPS) | yes | Masons Fraternity Ring (9533, -2.38 DPS) [quest]; Thunderbrow Ring (13097, -2.62 DPS) [world_drop]; Falcon's Hook (7552, -2.62 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 34.9 attack_power points (3.58 DPS) | yes | Masons Fraternity Ring (9533, -1.20 DPS) [quest]; Thunderbrow Ring (13097, -1.43 DPS) [world_drop]; Falcon's Hook (7552, -1.44 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Diamond Flask (20130) | Voodoo Feathers [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.53 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop]; Guttbuster (13139, -0.60 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Knight's Plate Hauberk; wrist: Deepfury Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Diamond Flask; main_hand: Blight; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1225, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 12323213132215201-050500000000000000-205000000000000000)

Set DPS (verified): 325.0. Weights run: 4.3s. Verify run: 78.8s. 2724 eligible items had no known source.

Race sweep: high-order-skyborne 325.8 ± 0.4, night-elf 325.0 ± 0.4, human 324.1 ± 0.4, gnome 323.1 ± 0.4, dwarf 320.4 ± 0.4 (dps); high-order-skyborne beats the guide race night-elf by 0.7 (0.22%) ± 0.6, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.670 ± 0.146, crit=3.815 ± 0.209 per rating point (14 rating = 1%, 53.405 per %), hit=3.256 ± 0.507 per rating point (10 rating = 1%, 32.560 per %), melee_haste=22.051 ± 5.103

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 207.9 attack_power points (17.95 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -6.90 DPS) [vendor]; Eye of Rend (12587, -10.86 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 79.4 attack_power points (6.86 DPS) | yes | Skibi's Pendant (13089, -3.00 DPS) [world_drop]; Emberfury Talisman (12929, -3.19 DPS, sim-verified) [dungeon]; Star of Mystaria (12103, -3.20 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 125.5 attack_power points (10.84 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Plate Shoulderguards (231537, -1.23 DPS) [pvp]; Lieutenant Commander's Plate Shoulders (227045, -3.29 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackveil Cape (11626, -0.96 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.14 DPS) [rep]; Cape of the Black Baron (13340, -3.37 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.49 DPS) [crafted]; Savage Gladiator Chain (11726, -3.05 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -4.49 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 86.0 attack_power points (7.42 DPS) | yes | Vambraces of the Sadist (13400, -1.60 DPS) [dungeon]; Forest Stalker's Bracers (19589, -1.95 DPS) [rep] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | 92.1 attack_power points (7.95 DPS) | yes | Highlander's Plate Girdle (20041, -0.40 DPS) [rep]; Highlander's Chain Girdle (20043, -0.40 DPS) [rep]; Marksman's Girdle (22232, -0.83 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -3.41 DPS) [vendor]; Marshal's Plate Legguards (231540, -4.62 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.66 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 86.0 attack_power points (7.42 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.21 DPS) [dungeon]; Swiftwalker Boots (12553, -1.72 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -2.59 DPS) [vendor]; Painweaver Band (13098, -2.98 DPS) [dungeon]; Naglering (11669, -10.94 DPS, sim-verified) [dungeon] |
| finger2 | Magni's Will (12548) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Painweaver Band (13098, -0.17 DPS) [dungeon]; Band of the Penitent (13217, -1.21 DPS) [quest] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -4.16 DPS, sim-verified) [dungeon] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -3.77 DPS, sim-verified) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -1.30 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.23 DPS) [quest]; Ancient Bone Bow (18680, -0.69 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.74 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Battleborn Armbraces; waist: Valiant's Waistguard; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Magni's Will; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2724, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 964.8. Weights run: 3.8s. Verify run: 80.5s. 2724 eligible items had no known source.

Race sweep: high-order-skyborne 966.9 ± 1.1, night-elf 964.8 ± 1.1, gnome 963.5 ± 1.1, human 962.5 ± 1.1, dwarf 952.1 ± 1.1 (dps); high-order-skyborne beats the guide race night-elf by 2.1 (0.21%) ± 1.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.881 ± 0.151, crit=3.677 ± 0.209 per rating point (14 rating = 1%, 51.485 per %), hit=5.647 ± 0.580 per rating point (10 rating = 1%, 56.470 per %), melee_haste=30.937 ± 5.294

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 255.5 attack_power points (47.08 DPS) | yes | Mask of the Unforgiven (13404, -16.78 DPS) [dungeon]; Lieutenant Commander's Plate Helm (23314, -18.67 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 77.5 attack_power points (14.28 DPS) | yes | Star of Mystaria (12103, -0.75 DPS) [dungeon]; Nacreous Shell Necklace (22403, -3.87 DPS) [dungeon]; Emberfury Talisman (12929, -4.79 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 129.3 attack_power points (23.82 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Truestrike Shoulders (12927, -3.79 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -3.92 DPS) [dungeon]; Cloak of Flames (3475, -5.16 DPS) [world_drop]; Stalwart Cloak (272415, -5.16 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -4.27 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -8.18 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -5.35 DPS) [rep]; Slashclaw Bracers (13211, -5.77 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Valiant's Waistguard (272403, -3.14 DPS) [vendor]; Sash of the Grand Hunt (22207, -3.18 DPS) [dungeon]; Marshal's Plate Girdle (16482, -5.05 DPS) [pvp] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -7.64 DPS) [vendor]; Marshal's Plate Legguards (231540, -10.47 DPS) [pvp]; Sentinel's Leather Pants (237818, -14.65 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 114.1 attack_power points (21.02 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.72 DPS) [dungeon]; Shadowcraft Boots (16711, -6.75 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.38 DPS) [dungeon]; Magni's Will (12548, -10.15 DPS) [quest]; Painweaver Band (13098, -10.77 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.75 DPS) [dungeon]; Magni's Will (12548, -5.51 DPS) [quest]; Painweaver Band (13098, -6.14 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -2.50 DPS) [world_drop]; Bloodseeker (19107, -4.00 DPS) [quest]; Ancient Bone Bow (18680, -5.12 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2724, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 12321200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 32.5. Weights run: 3.7s. Verify run: 15.0s. 371 eligible items had no known source.

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

Set DPS (verified): 65.6. Weights run: 4.4s. Verify run: 17.4s. 622 eligible items had no known source.

Race sweep: undead 65.6 ± 0.1, troll 64.6 ± 0.1, orc 64.5 ± 0.1, tauren 64.1 ± 0.1, windshaper-skyborne 63.3 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.744 ± 0.052, crit=1.246 ± 0.054 per rating point (14 rating = 1%, 17.447 per %), hit=1.195 ± 0.162 per rating point (10 rating = 1%, 11.951 per %), melee_haste=9.762 ± 1.430

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 33.7 attack_power points (1.74 DPS) | yes | Tusken Helm (6686, -0.40 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.50 DPS) [crafted]; Defender's Leather Helm (252455, -0.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.0 attack_power points (0.77 DPS) | yes | Scout's Medallion (19537, -0.05 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.15 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.2 attack_power points (1.51 DPS) | yes | Barbaric Iron Shoulders (7913, -0.35 DPS) [crafted]; Barbaric Shoulders (5964, -0.54 DPS) [crafted]; Mantle of Thieves (2264, -0.61 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.40 DPS) [dungeon]; Swiftrunner Cape (6745, -0.40 DPS) [quest] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Avenger's Armor (1488, -0.02 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.23 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.5 attack_power points (0.95 DPS) | yes | Yorgen Bracers (13012, -0.06 DPS) [world_drop]; Pugilist Bracers (4438, -0.13 DPS) [dungeon]; Barbaric Bracers (18948, -0.18 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 24.5 attack_power points (1.26 DPS) | yes | Mail Combat Gauntlets (4075, -0.08 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.10 DPS) [crafted]; Gauntlets of Ogre Strength (3341, -0.13 DPS) [world] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 30.5 attack_power points (1.57 DPS) | yes | Prowler's Leather Belt (252459, -0.10 DPS) [crafted]; Skulker's Leather Belt (252520, -0.14 DPS) [crafted]; Girdle of Golem Strength (9405, -0.33 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 34.2 attack_power points (1.77 DPS) | yes | Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.31 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.34 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Boots (252439, -0.06 DPS) [crafted]; Veteran's Boots (250503, -0.15 DPS) [crafted]; Stomping Boots (3741, -0.26 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.5 attack_power points (1.16 DPS) | yes | Band of the Fist (17694, -0.39 DPS) [quest]; Ring of Precision (1491, -0.62 DPS) [dungeon]; Tiger Band (6749, -0.64 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.2 attack_power points (1.10 DPS) | yes | Band of the Fist (17694, -0.37 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon]; Tiger Band (6749, -0.58 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.59 DPS) [dungeon]; Felbough Arm (273037, -0.70 DPS) [world]; Viscous Hammer (13045, -4.23 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 11.2 attack_power points (0.58 DPS) | yes | Golemsight Long Gun (273029, -0.04 DPS) [dungeon]; Gravespike Repeater (273033, -0.17 DPS) [world]; Silver Star (3463, -0.22 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 622, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 12323213132214001-000000000000000000-000000000000000000)

Set DPS (verified): 148.6. Weights run: 4.5s. Verify run: 12.9s. 928 eligible items had no known source.

Race sweep: undead 148.6 ± 0.2, orc 146.6 ± 0.2, troll 146.2 ± 0.2, tauren 145.7 ± 0.2, windshaper-skyborne 140.4 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.122 ± 0.080, crit=2.065 ± 0.094 per rating point (14 rating = 1%, 28.915 per %), hit=2.785 ± 0.233 per rating point (10 rating = 1%, 27.847 per %), melee_haste=18.604 ± 2.089

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 54.9 attack_power points (4.82 DPS) | yes | Chromite Barbute (8142, -0.48 DPS) [dungeon]; White Bandit Mask (10008, -0.84 DPS) [crafted]; Barbaric Iron Helm (7915, -1.56 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.3 attack_power points (2.05 DPS) | yes | Ethereal Talisman (4430, -0.40 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.60 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.82 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 attack_power points (2.93 DPS) | yes | Flintrock Shoulders (7755, -0.19 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.76 DPS) [crafted]; Nightscape Shoulders (8192, -0.88 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.34 DPS) [dungeon]; Parachute Cloak (10518, -0.34 DPS) [crafted]; Tigerstrike Mantle (13108, -0.34 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brutal Hauberk (7133, +0.00 DPS) [quest]; Kolkar Marauder Chain (6773, -1.18 DPS) [quest]; Carapace of Tuten'kash (10775, -1.70 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 22.5 attack_power points (1.97 DPS) | yes | Hawkeye's Bracers (14590, -0.15 DPS) [world_drop]; Branded Leather Bracers (19508, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.36 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Scarlet Gauntlets (10331, -1.72 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.1 attack_power points (3.26 DPS) | yes | Defiler's Chain Girdle (20153, -0.02 DPS) [rep]; Officer's Belt (250556, -0.38 DPS) [crafted]; Skulker's Leather Belt (252520, -0.53 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.6 attack_power points (3.91 DPS) | yes | Triprunner Dungarees (9624, -0.03 DPS) [quest]; Symbolic Legplates (14829, -0.16 DPS) [world_drop]; Scarlet Leggings (10330, -0.23 DPS) [dungeon] |
| feet | Boots of Avoidance (14549) | World drop [world_drop] | 41.7 attack_power points (3.66 DPS) | yes | Officer's Boots (250546, -0.08 DPS) [crafted]; Blackforge Greaves (6423, -0.22 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.38 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 33.0 attack_power points (2.90 DPS) | yes | Falcon's Hook (7552, -0.69 DPS) [dungeon]; Thunderbrow Ring (13097, -0.93 DPS) [world_drop]; Mark of Kern (2262, -1.14 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.2 attack_power points (2.21 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.25 DPS) [world_drop]; Mark of Kern (2262, -0.46 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.61 DPS) [vendor]; Nightblade (1982, -2.60 DPS, sim-verified) [world_drop]; X'caliboar (10758, -2.99 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Swiftwind (13038, -0.31 DPS) [world_drop]; The Silencer (13138, -0.38 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Boots of Avoidance; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 928, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 12323213132215201-050200000000000000-000000000000000000)

Set DPS (verified): 225.0. Weights run: 4.7s. Verify run: 13.8s. 1206 eligible items had no known source.

Race sweep: windshaper-skyborne 225.2 ± 0.3, undead 225.0 ± 0.3, troll 223.8 ± 0.3, orc 222.6 ± 0.3, tauren 221.3 ± 0.3 (dps); windshaper-skyborne beats the guide race undead by 0.2 (0.10%) ± 0.4, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.659 ± 0.094, crit=1.930 ± 0.123 per rating point (14 rating = 1%, 27.025 per %), hit=2.647 ± 0.277 per rating point (10 rating = 1%, 26.475 per %), melee_haste=24.749 ± 2.820

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 89.5 attack_power points (9.18 DPS) | yes | Blood Guard's Plate Helm (220803, -1.03 DPS) [vendor]; Embrace of the Lycan (9479, -4.26 DPS) [dungeon]; Raging Berserker's Helm (7719, -4.29 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (3.24 DPS) | yes | Woven Ivy Necklace (19159, -0.48 DPS) [quest]; Scout's Medallion (19535, -1.20 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 46.5 attack_power points (4.77 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -0.74 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.94 DPS) [world_drop] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.06 DPS) [dungeon]; Dark Phantom Cape (13122, -1.30 DPS, sim-verified) [world_drop]; Pridelord Cape (14673, -1.47 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -0.46 DPS) [dungeon]; Warbear Harness (15064, -0.53 DPS) [crafted]; Warforged Chestplate (11195, -0.93 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.9 attack_power points (3.37 DPS) | yes | Officer's Wristguards (250581, -0.10 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.34 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.41 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.35 DPS, sim-verified) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 49.0 attack_power points (5.03 DPS) | yes | Skulker's Leather Waistguard (252474, -0.18 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.21 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.51 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 6 · Stone Guard · Horde [vendor] | 67.6 attack_power points (6.93 DPS) | yes | Infernal Trickster Leggings (17754, -0.82 DPS) [dungeon]; Stormshroud Pants (15057, -1.39 DPS) [crafted]; Serpentskin Leggings (8262, -1.55 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 49.6 attack_power points (5.08 DPS) | yes | Prowler's Leather Boots (252468, -0.55 DPS) [crafted]; Skulker's Leather Boots (252469, -0.62 DPS) [crafted]; Officer's Sabatons (250561, -0.89 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.5 attack_power points (4.77 DPS) | yes | White Bone Band (11862, -2.31 DPS) [quest]; Masons Fraternity Ring (9533, -2.38 DPS) [quest]; Thunderbrow Ring (13097, -2.62 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 34.9 attack_power points (3.58 DPS) | yes | White Bone Band (11862, -1.12 DPS) [quest]; Masons Fraternity Ring (9533, -1.20 DPS) [quest]; Thunderbrow Ring (13097, -1.43 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.55 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blight (7959, +0.00 DPS) [crafted]; Warmonger (13052, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, -0.90 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.53 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop]; Guttbuster (13139, -0.60 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Stone Guard's Plate Armor; wrist: Deepfury Bracers; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1206, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 12323213132215201-050500000000000000-205000000000000000)

Set DPS (verified): 322.7. Weights run: 4.3s. Verify run: 76.9s. 2714 eligible items had no known source.

Race sweep: undead 322.7 ± 0.4, windshaper-skyborne 321.8 ± 0.4, troll 319.1 ± 0.4, orc 318.6 ± 0.4, tauren 317.2 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.670 ± 0.146, crit=3.815 ± 0.209 per rating point (14 rating = 1%, 53.405 per %), hit=3.256 ± 0.507 per rating point (10 rating = 1%, 32.560 per %), melee_haste=22.051 ± 5.103

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 207.9 attack_power points (17.95 DPS) | yes | Champion's Plate Helm (227043, -6.90 DPS) [pvp]; Fury Visor (20521, -7.42 DPS) [quest]; Eye of Rend (12587, -9.76 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 79.4 attack_power points (6.86 DPS) | yes | Emberfury Talisman (12929, -2.88 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -3.00 DPS) [world_drop]; Star of Mystaria (12103, -3.20 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 125.5 attack_power points (10.84 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -1.23 DPS) [pvp]; Champion's Plate Shoulders (227042, -3.29 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.96 DPS) [dungeon]; Deathguard's Cloak (20068, -1.14 DPS) [rep]; Cape of the Black Baron (13340, -4.17 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -1.55 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -2.99 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 86.0 attack_power points (7.42 DPS) | yes | Vambraces of the Sadist (13400, -1.60 DPS) [dungeon]; Forest Stalker's Bracers (19589, -1.95 DPS) [rep] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, -2.47 DPS, sim-verified) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | 92.1 attack_power points (7.95 DPS) | yes | Marksman's Girdle (22232, -0.30 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.40 DPS) [rep]; Defiler's Leather Girdle (20190, -0.40 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, -0.00 DPS) [pvp]; Outrider's Plate Legguards (22651, -2.11 DPS, sim-verified) [rep]; Legionnaire's Plate Leggings (227048, -4.20 DPS) [pvp] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 86.0 attack_power points (7.42 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.21 DPS) [dungeon]; Swiftwalker Boots (12553, -1.72 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.98 DPS) [dungeon]; Band of the Penitent (13217, -4.02 DPS) [quest]; Naglering (11669, -10.86 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of the Penitent (13217, -1.43 DPS) [quest]; Band of Rumination (18103, -1.43 DPS) [dungeon]; Painweaver Band (13098, -2.94 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+8.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -2.79 DPS, sim-verified) [dungeon] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Destiny (647, -3.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.70 DPS) [dungeon]; Bloodseeker (19107, -0.92 DPS) [quest]; Ancient Bone Bow (18680, -1.38 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2714, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 953.7. Weights run: 3.8s. Verify run: 80.3s. 2714 eligible items had no known source.

Race sweep: windshaper-skyborne 957.9 ± 1.1, undead 953.7 ± 1.1, orc 953.1 ± 1.1, troll 950.2 ± 1.1, tauren 943.9 ± 1.1 (dps); windshaper-skyborne beats the guide race undead by 4.2 (0.44%) ± 1.5, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.881 ± 0.151, crit=3.677 ± 0.209 per rating point (14 rating = 1%, 51.485 per %), hit=5.647 ± 0.580 per rating point (10 rating = 1%, 56.470 per %), melee_haste=30.937 ± 5.294

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 255.5 attack_power points (47.08 DPS) | yes | Mask of the Unforgiven (13404, -16.78 DPS) [dungeon]; Champion's Plate Helm (227043, -18.67 DPS) [pvp]; Fury Visor (20521, -19.89 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 77.5 attack_power points (14.28 DPS) | yes | Star of Mystaria (12103, -0.75 DPS) [dungeon]; Nacreous Shell Necklace (22403, -3.87 DPS) [dungeon]; Emberfury Talisman (12929, -4.79 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 129.3 attack_power points (23.82 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -3.79 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -3.92 DPS) [dungeon]; Cloak of Flames (3475, -5.16 DPS) [world_drop]; Stalwart Cloak (272415, -5.16 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -4.27 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -8.18 DPS) [quest] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -5.35 DPS) [rep]; Slashclaw Bracers (13211, -5.77 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Valiant's Waistguard (272403, -3.14 DPS) [vendor]; Sash of the Grand Hunt (22207, -3.18 DPS) [dungeon]; General's Plate Girdle (16547, -5.05 DPS) [pvp] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, -0.99 DPS) [pvp]; Outrider's Plate Legguards (22651, -7.23 DPS) [rep]; Plaguehound Leggings (18736, -12.14 DPS) [dungeon] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 114.1 attack_power points (21.02 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.72 DPS) [dungeon]; Shadowcraft Boots (16711, -6.75 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.38 DPS) [dungeon]; Painweaver Band (13098, -10.77 DPS) [dungeon]; Legionnaire's Band (19510, -11.77 DPS) [rep] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.75 DPS) [dungeon]; Painweaver Band (13098, -6.14 DPS) [dungeon]; Legionnaire's Band (19510, -7.13 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -2.50 DPS) [world_drop]; Bloodseeker (19107, -4.00 DPS) [quest]; Ancient Bone Bow (18680, -5.12 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Battleborn Armbraces; waist: Marksman's Girdle; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram; ranged: Blackcrow

No-known-source sample (15 of 2714, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

