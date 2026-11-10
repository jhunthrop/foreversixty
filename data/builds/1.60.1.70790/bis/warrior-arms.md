# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 13320200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 31.3. Weights run: 3.2s. Verify run: 17.9s. 388 eligible items had no known source.

Race sweep: high-order-skyborne 31.5 ± 0.0, night-elf 31.3 ± 0.0, human 31.3 ± 0.0, dwarf 30.9 ± 0.0, gnome 30.8 ± 0.0 (dps); high-order-skyborne beats the guide race night-elf by 0.2 (0.60%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.501 ± 0.034, crit=0.818 ± 0.033 per rating point (14 rating = 1%, 11.459 per %), hit=1.157 ± 0.087 per rating point (10 rating = 1%, 11.566 per %), melee_haste=7.865 ± 1.008

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.85 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.21 DPS) [crafted]; Brawler's Leather Hood (252504, -0.34 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.0 attack_power points (0.38 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.64 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) | Blacksmithing [crafted] | sim-verified (31.3 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, +0.00 DPS) [dungeon] |
| back | Dark Leather Cloak (2316) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Lizardscale Cloak (6449, +0.00 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.02 DPS) [world_drop]; Cape of the Brotherhood (5193, -0.04 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Veteran's Chain Shirt (250488, -0.58 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bravo's Armbands (270015, +0.00 DPS) [quest]; Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Veteran's Gloves (250508, +0.00 DPS) [crafted]; Thorbia's Gauntlets (12994, -0.69 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -0.95 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -0.78 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, -0.06 DPS) [crafted]; Defender's Leather Boots (252441, -0.32 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.0 attack_power points (0.59 DPS) | yes | Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Ring of the Moon (12052, -0.51 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Ring of the Moon (12052, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.26 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 304.0 attack_power points (12.87 DPS) | yes | Duskbringer (2205, -0.30 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.60 DPS) [dungeon]; Sword of the Fallen (286977, -0.68 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.0 attack_power points (0.25 DPS) | yes | Dwarven Fishing Pole (3567, -0.08 DPS) [quest]; Fine Longbow (11304, -0.08 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.08 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Dark Leather Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 30 (night-elf, 13323213030000000-000000000000000000-000000000000000000)

Set DPS (verified): 63.1. Weights run: 3.6s. Verify run: 12.6s. 637 eligible items had no known source.

Race sweep: night-elf 63.1 ± 0.1, human 62.4 ± 0.1, gnome 61.5 ± 0.1, dwarf 61.3 ± 0.1, high-order-skyborne 60.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.795 ± 0.055, crit=1.258 ± 0.057 per rating point (14 rating = 1%, 17.616 per %), hit=1.303 ± 0.166 per rating point (10 rating = 1%, 13.033 per %), melee_haste=11.053 ± 1.471

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 34.2 attack_power points (1.61 DPS) | yes | Veteran's Chain Helm (250498, -0.48 DPS) [crafted]; Defender's Leather Helm (252455, -0.48 DPS) [crafted]; Tusken Helm (6686, -0.48 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.2 attack_power points (0.72 DPS) | yes | Sentinel's Medallion (19541, -0.04 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.15 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.7 attack_power points (1.40 DPS) | yes | Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.53 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -0.56 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.20 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.32 DPS, sim-verified) [pvp]; Glowing Lizardscale Cloak (6449, -0.37 DPS) [dungeon] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.21 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.8 attack_power points (0.88 DPS) | yes | Yorgen Bracers (13012, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.13 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 24.8 attack_power points (1.17 DPS) | yes | Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Gauntlets of Ogre Strength (3341, -0.13 DPS) [world] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 30.8 attack_power points (1.45 DPS) | yes | Prowler's Leather Belt (252459, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.12 DPS) [crafted]; Girdle of Golem Strength (9405, -0.32 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 34.6 attack_power points (1.63 DPS) | yes | Brawler's Leather Legguards (252516, -0.21 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.28 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.30 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.4 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Alacritous Treads (277234, -0.07 DPS) [quest]; Veteran's Boots (250503, -0.15 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.8 attack_power points (1.07 DPS) | yes | Seal of Wrynn (2933, -0.54 DPS) [quest]; Ring of Precision (1491, -0.57 DPS) [dungeon]; Tiger Band (6749, -0.60 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.4 attack_power points (1.01 DPS) | yes | Ring of Precision (1491, -0.50 DPS) [dungeon]; Tiger Band (6749, -0.54 DPS) [quest]; Seal of Wrynn (2933, -0.54 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Felbough Arm (273037, -0.58 DPS) [world]; Viscous Hammer (13045, -5.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 11.4 attack_power points (0.54 DPS) | yes | Golemsight Long Gun (273029, -0.03 DPS) [dungeon]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.20 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 637, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 13323213032214001-000000000000000000-000000000000000000)

Set DPS (verified): 146.1. Weights run: 3.7s. Verify run: 18.8s. 946 eligible items had no known source.

Race sweep: night-elf 146.1 ± 0.2, human 145.1 ± 0.2, gnome 143.4 ± 0.2, dwarf 143.0 ± 0.2, high-order-skyborne 138.6 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.139 ± 0.076, crit=2.137 ± 0.093 per rating point (14 rating = 1%, 29.925 per %), hit=2.140 ± 0.149 per rating point (10 rating = 1%, 21.395 per %), melee_haste=9.567 ± 1.890

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 55.9 attack_power points (4.34 DPS) | yes | Chromite Barbute (8142, -0.49 DPS) [dungeon]; White Bandit Mask (10008, -0.81 DPS) [crafted]; Barbaric Iron Helm (7915, -1.45 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.5 attack_power points (1.83 DPS) | yes | Kaleidoscope Chain (13084, -0.54 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.74 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.90 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.5 attack_power points (2.60 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.68 DPS) [crafted]; Nightscape Shoulders (8192, -0.78 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.30 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.47 DPS) [quest]; Parachute Cloak (10518, -0.60 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Icemail Jerkin (1981, -0.63 DPS, sim-verified) [world_drop]; Kolkar Marauder Chain (6773, -1.06 DPS) [quest]; Carapace of Tuten'kash (10775, -1.52 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 22.6 attack_power points (1.75 DPS) | yes | Hawkeye's Bracers (14590, -0.13 DPS) [world_drop]; Branded Leather Bracers (19508, -0.20 DPS) [dungeon]; Yorgen Bracers (13012, -0.32 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Scarlet Gauntlets (10331, -2.41 DPS, sim-verified) [dungeon] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 37.9 attack_power points (2.95 DPS) | yes | Ogron's Sash (13117, -0.05 DPS) [world_drop]; Officer's Belt (250556, -0.40 DPS) [crafted]; Skulker's Leather Belt (252520, -0.52 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.9 attack_power points (3.49 DPS) | yes | Triprunner Dungarees (9624, -0.03 DPS) [quest]; Symbolic Legplates (14829, -0.16 DPS) [world_drop]; Scarlet Leggings (10330, -0.23 DPS) [dungeon] |
| feet | Boots of Avoidance (14549) | World drop [world_drop] | 41.9 attack_power points (3.26 DPS) | yes | Officer's Boots (250546, -0.08 DPS) [crafted]; Blackforge Greaves (6423, -0.20 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.34 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 33.1 attack_power points (2.57 DPS) | yes | Falcon's Hook (7552, -0.61 DPS) [dungeon]; Thunderbrow Ring (13097, -0.83 DPS) [world_drop]; Mark of Kern (2262, -1.02 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.4 attack_power points (1.97 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Mark of Kern (2262, -0.42 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.53 DPS) [vendor]; Bonebiter (6830, -0.96 DPS) [quest]; Nightblade (1982, -4.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.14 DPS) [quest]; Swiftwind (13038, -0.27 DPS) [world_drop]; The Silencer (13138, -0.34 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Boots of Avoidance; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 946, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 13323213032215201-050200000000000000-000000000000000000)

Set DPS (verified): 221.9. Weights run: 3.9s. Verify run: 14.1s. 1224 eligible items had no known source.

Race sweep: night-elf 221.9 ± 0.3, high-order-skyborne 221.1 ± 0.3, human 220.8 ± 0.3, dwarf 220.1 ± 0.3, gnome 219.8 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.783 ± 0.092, crit=2.046 ± 0.119 per rating point (14 rating = 1%, 28.647 per %), hit=2.478 ± 0.279 per rating point (10 rating = 1%, 24.782 per %), melee_haste=22.627 ± 2.763

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 89.4 attack_power points (9.12 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.02 DPS) [vendor]; Ornate Mithril Helm (7937, -4.16 DPS) [crafted]; Raging Berserker's Helm (7719, -5.48 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.2 attack_power points (3.38 DPS) | yes | Sentinel's Medallion (19539, -1.20 DPS) [rep]; Kaleidoscope Chain (13084, -1.84 DPS) [world_drop] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 44.8 attack_power points (4.57 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.67 DPS) [world_drop]; Officer's Pauldrons (250576, -1.81 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.04 DPS) [dungeon]; Dark Phantom Cape (13122, -1.04 DPS) [world_drop]; Sergeant Major's Cape (16336, -1.45 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warbear Harness (15064, -1.37 DPS) [crafted]; Mixologist's Tunic (12793, -1.40 DPS) [dungeon]; Knight's Plate Hauberk (220794, -2.23 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.7 attack_power points (3.54 DPS) | yes | Officer's Wristguards (250581, -0.21 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.43 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.48 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -4.23 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 50.6 attack_power points (5.16 DPS) | yes | Prowler's Leather Waistguard (252473, -0.13 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.17 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.20 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | PvP rank 6 · Knight · Alliance [vendor] | 70.5 attack_power points (7.19 DPS) | yes | Infernal Trickster Leggings (17754, -1.02 DPS) [dungeon]; Stormshroud Pants (15057, -1.34 DPS) [crafted]; Gryphon Rider's Leggings (9652, -1.40 DPS) [quest] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 51.2 attack_power points (5.22 DPS) | yes | Prowler's Leather Boots (252468, -0.57 DPS) [crafted]; Skulker's Leather Boots (252469, -0.61 DPS) [crafted]; Officer's Sabatons (250561, -0.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.8 attack_power points (4.57 DPS) | yes | Masons Fraternity Ring (9533, -2.02 DPS) [quest]; Falcon's Hook (7552, -2.32 DPS) [dungeon]; Ring of the Underwood (2951, -2.34 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 36.0 attack_power points (3.68 DPS) | yes | Masons Fraternity Ring (9533, -1.13 DPS) [quest]; Falcon's Hook (7552, -1.43 DPS) [dungeon]; Ring of the Underwood (2951, -1.45 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.63 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, -0.89 DPS) [world_drop]; Blight (7959, -1.45 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.56 DPS, sim-verified) [world_drop]; Stinging Bow (10624, -0.69 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.69 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; chest: Icemail Jerkin; wrist: Deepfury Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1224, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 13323213032215201-050500000000000000-205000000000000000)

Set DPS (verified): 314.7. Weights run: 3.7s. Verify run: 81.7s. 2695 eligible items had no known source.

Race sweep: night-elf 314.7 ± 0.4, high-order-skyborne 314.4 ± 0.4, dwarf 313.9 ± 0.4, human 313.4 ± 0.4, gnome 312.4 ± 0.4 (dps); the guide race night-elf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.560 ± 0.142, crit=3.657 ± 0.203 per rating point (14 rating = 1%, 51.196 per %), hit=4.312 ± 0.521 per rating point (10 rating = 1%, 43.125 per %), melee_haste=26.140 ± 4.887

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 224.6 attack_power points (18.97 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -7.46 DPS) [vendor]; Mask of the Unforgiven (13404, -7.81 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 77.2 attack_power points (6.52 DPS) | yes | Emberfury Talisman (12929, -2.20 DPS) [dungeon]; Skibi's Pendant (13089, -2.87 DPS) [world_drop]; Star of Mystaria (12103, -6.63 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 120.3 attack_power points (10.16 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Plate Shoulderguards (231537, -0.02 DPS) [pvp]; Truestrike Shoulders (12927, -2.67 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackveil Cape (11626, -0.89 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.98 DPS) [rep]; Howler's Furs (272414, -1.14 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -1.57 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -2.95 DPS) [quest]; Dawn Armor (252483, -5.36 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Vambraces of the Sadist (13400, -0.46 DPS) [dungeon]; Slashclaw Bracers (13211, -0.81 DPS) [dungeon]; Battleborn Armbraces (12936, -2.06 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gloves (271948, -2.57 DPS, sim-verified) [rep] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Plate Girdle (20041, -0.36 DPS) [rep]; Highlander's Chain Girdle (20043, -0.36 DPS) [rep]; Marksman's Girdle (22232, -3.66 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -3.01 DPS) [vendor]; Marshal's Plate Legguards (231540, -4.19 DPS) [pvp]; Sentinel's Leather Pants (237818, -5.37 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 94.3 attack_power points (7.97 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.07 DPS) [dungeon]; Shadowcraft Boots (16711, -2.33 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.47 DPS) [dungeon]; Magni's Will (12548, -3.64 DPS) [quest]; Naglering (11669, -6.23 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.99 DPS) [dungeon]; Magni's Will (12548, -1.16 DPS) [quest]; Blackstone Ring (17713, -1.34 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (+6.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -2.99 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.81 DPS) [quest]; Ancient Bone Bow (18680, -1.30 DPS) [dungeon]; Blackcrow (12651, -2.48 DPS, sim-verified) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; waist: Valiant's Waistguard; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 2695, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 959.4. Weights run: 3.1s. Verify run: 76.7s. 2695 eligible items had no known source.

Race sweep: high-order-skyborne 961.9 ± 1.1, dwarf 961.7 ± 1.1, night-elf 959.4 ± 1.1, human 957.3 ± 1.1, gnome 954.1 ± 1.1 (dps); high-order-skyborne beats the guide race night-elf by 2.5 (0.26%) ± 1.5, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.870 ± 0.149, crit=3.690 ± 0.207 per rating point (14 rating = 1%, 51.659 per %), hit=6.265 ± 0.581 per rating point (10 rating = 1%, 62.652 per %), melee_haste=32.260 ± 5.266

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 268.2 attack_power points (49.45 DPS) | yes | Mask of the Unforgiven (13404, -18.01 DPS, sim-verified) [dungeon]; Lieutenant Commander's Plate Helm (23314, -19.86 DPS) [vendor] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (+15.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Nacreous Shell Necklace (22403, -2.77 DPS) [dungeon]; Emberfury Talisman (12929, -4.79 DPS) [dungeon]; Star of Mystaria (12103, -15.88 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 129.1 attack_power points (23.81 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Truestrike Shoulders (12927, -2.06 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of Flames (3475, -0.07 DPS) [world_drop]; Stalwart Cloak (272415, -0.07 DPS) [vendor]; Howler's Furs (272414, -4.98 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Gladiator Chain (11726, -4.25 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -8.11 DPS) [quest]; Dawn Armor (252483, -10.75 DPS, sim-verified) [crafted] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, -0.74 DPS) [rep]; Blackmist Armguards (12966, -1.68 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; Theramore Gloves (271948, +0.00 DPS) [rep]; Voone's Vice Grips (13963, -23.22 DPS, sim-verified) [quest] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 122.9 attack_power points (22.66 DPS) | yes | Sash of the Grand Hunt (22207, -3.17 DPS) [dungeon]; Valiant's Waistguard (272403, -4.29 DPS) [vendor]; Brigam Girdle (13142, -5.43 DPS) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -7.57 DPS) [vendor]; Marshal's Plate Legguards (231540, -10.41 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.78 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 120.0 attack_power points (22.13 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.70 DPS) [dungeon]; Shadowcraft Boots (16711, -7.56 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.42 DPS) [dungeon]; Painweaver Band (13098, -11.18 DPS) [dungeon]; Naglering (11669, -11.50 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.72 DPS) [dungeon]; Painweaver Band (13098, -6.48 DPS) [dungeon]; Magni's Will (12548, -6.60 DPS) [quest] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+20.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -9.57 DPS, sim-verified) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -15.41 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -1.52 DPS) [quest]; Ancient Bone Bow (18680, -2.65 DPS) [dungeon]; Blackcrow (12651, -6.15 DPS, sim-verified) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Slashclaw Bracers; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2695, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 13320200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 32.6. Weights run: 3.2s. Verify run: 13.5s. 371 eligible items had no known source.

Race sweep: windshaper-skyborne 32.8 ± 0.0, undead 32.6 ± 0.0, troll 31.8 ± 0.0, orc 31.6 ± 0.0, tauren 31.5 ± 0.0 (dps); windshaper-skyborne beats the guide race undead by 0.3 (0.85%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.944 ± 0.022, crit=0.526 ± 0.021 per rating point (14 rating = 1%, 7.362 per %), hit=0.641 ± 0.084 per rating point (10 rating = 1%, 6.412 per %), melee_haste=7.659 ± 0.705

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.88 DPS) | yes | Defender's Leather Hood (252447, -0.18 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.47 DPS) [crafted]; Brawler's Leather Hood (252504, -0.55 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 5.7 attack_power points (0.25 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.27 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.06 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Subterranean Cape (14149, -0.09 DPS) [dungeon]; Grave Shroud (279865, -0.09 DPS) [quest] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -0.98 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.44 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.18 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.10 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.80 DPS) | yes | Brawler's Leather Belt (252428, -0.28 DPS) [crafted]; Veteran's Chain Belt (250513, -0.32 DPS) [crafted]; Cobrahn's Grasp (6460, -1.35 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.96 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 14.7 attack_power points (0.65 DPS) | yes | Veteran's Boots (250503, -0.04 DPS) [crafted]; Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Defender's Leather Boots (252441, -0.21 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.8 attack_power points (0.52 DPS) | yes | Loop of Sacrifice (281673, -0.26 DPS) [quest]; Bounty Hunter's Ring (5351, -0.40 DPS) [quest]; Ring of the Moon (12052, -0.43 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.35 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS) [quest]; Bounty Hunter's Ring (5351, -0.23 DPS) [quest]; Ring of the Moon (12052, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (13.39 DPS) | yes | Hammerbone (270018, -0.12 DPS) [quest]; Duskbringer (2205, -0.38 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.01 DPS) [world_drop]; Heavy Shortbow (3036, -0.09 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Forsaken Greataxe; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 371, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (undead, 13323213030000000-000000000000000000-000000000000000000)

Set DPS (verified): 63.2. Weights run: 3.6s. Verify run: 11.1s. 622 eligible items had no known source.

Race sweep: undead 63.2 ± 0.1, troll 62.1 ± 0.1, orc 61.9 ± 0.1, tauren 61.5 ± 0.1, windshaper-skyborne 60.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.795 ± 0.055, crit=1.258 ± 0.057 per rating point (14 rating = 1%, 17.616 per %), hit=1.303 ± 0.166 per rating point (10 rating = 1%, 13.033 per %), melee_haste=11.053 ± 1.471

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 34.2 attack_power points (1.61 DPS) | yes | Veteran's Chain Helm (250498, -0.48 DPS) [crafted]; Defender's Leather Helm (252455, -0.48 DPS) [crafted]; Tusken Helm (6686, -0.49 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.2 attack_power points (0.72 DPS) | yes | Scout's Medallion (19537, -0.04 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.15 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 29.7 attack_power points (1.40 DPS) | yes | Barbaric Iron Shoulders (7913, -0.37 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Mantle of Thieves (2264, -0.56 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.20 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.37 DPS) [dungeon]; Swiftrunner Cape (6745, -0.37 DPS) [quest] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Avenger's Armor (1488, -0.04 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.13 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.21 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.8 attack_power points (0.88 DPS) | yes | Yorgen Bracers (13012, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.13 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 24.8 attack_power points (1.17 DPS) | yes | Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Gauntlets of Ogre Strength (3341, -0.13 DPS) [world] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 30.8 attack_power points (1.45 DPS) | yes | Prowler's Leather Belt (252459, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.12 DPS) [crafted]; Girdle of Golem Strength (9405, -0.32 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 34.6 attack_power points (1.63 DPS) | yes | Brawler's Leather Legguards (252516, -0.21 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.28 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.30 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.4 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Veteran's Boots (250503, -0.15 DPS) [crafted]; Stomping Boots (3741, -0.24 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.8 attack_power points (1.07 DPS) | yes | Band of the Fist (17694, -0.36 DPS) [quest]; Ring of Precision (1491, -0.57 DPS) [dungeon]; Tiger Band (6749, -0.60 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.4 attack_power points (1.01 DPS) | yes | Band of the Fist (17694, -0.29 DPS) [quest]; Ring of Precision (1491, -0.50 DPS) [dungeon]; Tiger Band (6749, -0.54 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Felbough Arm (273037, -0.58 DPS) [world]; Viscous Hammer (13045, -5.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 11.4 attack_power points (0.54 DPS) | yes | Golemsight Long Gun (273029, -0.03 DPS) [dungeon]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.20 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 622, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 13323213032214001-000000000000000000-000000000000000000)

Set DPS (verified): 145.0. Weights run: 3.7s. Verify run: 16.6s. 928 eligible items had no known source.

Race sweep: undead 145.0 ± 0.2, orc 143.1 ± 0.2, troll 142.4 ± 0.2, tauren 142.2 ± 0.2, windshaper-skyborne 137.7 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.889 ± 0.076, crit=1.760 ± 0.089 per rating point (14 rating = 1%, 24.633 per %), hit=2.516 ± 0.244 per rating point (10 rating = 1%, 25.161 per %), melee_haste=10.409 ± 2.079

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 50.6 attack_power points (4.30 DPS) | yes | Chromite Barbute (8142, -0.34 DPS) [dungeon]; White Bandit Mask (10008, -0.67 DPS) [crafted]; Barbaric Iron Helm (7915, -1.33 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 20.8 attack_power points (1.76 DPS) | yes | Ethereal Talisman (4430, -0.27 DPS) [quest]; Kaleidoscope Chain (13084, -0.44 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.58 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.8 attack_power points (2.61 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.63 DPS) [crafted]; Hard Gold Pauldrons (250539, -0.75 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.35 DPS) [dungeon]; Parachute Cloak (10518, -0.35 DPS) [crafted]; Tigerstrike Mantle (13108, -0.35 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, -1.58 DPS) [dungeon]; Kolkar Marauder Chain (6773, -2.48 DPS) [quest]; Carapace of Tuten'kash (10775, -3.00 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 21.6 attack_power points (1.83 DPS) | yes | Branded Leather Bracers (19508, -0.13 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.19 DPS) [world_drop]; Yorgen Bracers (13012, -0.33 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -3.01 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 35.0 attack_power points (2.97 DPS) | yes | Defiler's Chain Girdle (20153, -0.20 DPS) [rep]; Officer's Belt (250556, -0.31 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.42 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.56 DPS) | yes | Symbolic Legplates (14829, -0.06 DPS) [world_drop]; Triprunner Dungarees (9624, -0.17 DPS) [quest]; Basilisk Hide Pants (1718, -0.20 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 39.2 attack_power points (3.33 DPS) | yes | Boots of Avoidance (14549, -0.07 DPS) [world_drop]; Blackforge Greaves (6423, -0.20 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.34 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 31.1 attack_power points (2.64 DPS) | yes | Ring of the Underwood (2951, -0.70 DPS) [world_drop]; Thunderbrow Ring (13097, -0.80 DPS) [world_drop]; Mark of Kern (2262, -0.94 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.0 attack_power points (1.95 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Mark of Kern (2262, -0.25 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.79 DPS) [vendor]; X'caliboar (10758, -2.89 DPS) [dungeon]; Nightblade (1982, -4.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.18 DPS) [quest]; The Silencer (13138, -0.31 DPS) [world_drop]; Swiftwind (13038, -0.38 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Falcon's Hook; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 928, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 13323213032215201-050200000000000000-000000000000000000)

Set DPS (verified): 224.3. Weights run: 3.9s. Verify run: 14.6s. 1205 eligible items had no known source.

Race sweep: undead 224.3 ± 0.3, windshaper-skyborne 223.6 ± 0.3, troll 222.9 ± 0.3, orc 221.8 ± 0.3, tauren 220.7 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.783 ± 0.092, crit=2.046 ± 0.119 per rating point (14 rating = 1%, 28.647 per %), hit=2.478 ± 0.279 per rating point (10 rating = 1%, 24.782 per %), melee_haste=22.627 ± 2.763

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 89.4 attack_power points (9.12 DPS) | yes | Blood Guard's Plate Helm (220803, -1.02 DPS) [vendor]; Raging Berserker's Helm (7719, -1.94 DPS, sim-verified) [dungeon]; Ornate Mithril Helm (7937, -4.16 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.2 attack_power points (3.38 DPS) | yes | Woven Ivy Necklace (19159, -0.52 DPS) [quest]; Scout's Medallion (19535, -1.20 DPS) [rep] |
| shoulder | Razorsteel Shoulders (20517) | Voodoo Feathers [quest] | 44.8 attack_power points (4.57 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.67 DPS) [world_drop]; Officer's Pauldrons (250576, -0.79 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.04 DPS) [dungeon]; Dark Phantom Cape (13122, -1.04 DPS) [world_drop]; Pridelord Cape (14673, -1.52 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stone Guard's Plate Armor (220801, -0.83 DPS, sim-verified) [vendor]; Warbear Harness (15064, -1.37 DPS) [crafted]; Mixologist's Tunic (12793, -1.40 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.7 attack_power points (3.54 DPS) | yes | Officer's Wristguards (250581, -0.21 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.43 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.48 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.82 DPS, sim-verified) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 50.6 attack_power points (5.16 DPS) | yes | Prowler's Leather Waistguard (252473, -0.13 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.17 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.20 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 6 · Stone Guard · Horde [vendor] | 70.5 attack_power points (7.19 DPS) | yes | Stormshroud Pants (15057, -1.34 DPS) [crafted]; Infernal Trickster Leggings (17754, -1.46 DPS, sim-verified) [dungeon]; Serpentskin Leggings (8262, -1.63 DPS) [world_drop] |
| feet | Battlechaser's Greaves (12555) | Blackrock Depths [dungeon] | 51.2 attack_power points (5.22 DPS) | yes | Prowler's Leather Boots (252468, -0.57 DPS) [crafted]; Skulker's Leather Boots (252469, -0.61 DPS) [crafted]; Officer's Sabatons (250561, -0.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.8 attack_power points (4.57 DPS) | yes | Masons Fraternity Ring (9533, -2.02 DPS) [quest]; White Bone Band (11862, -2.12 DPS) [quest]; Falcon's Hook (7552, -2.32 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 36.0 attack_power points (3.68 DPS) | yes | Masons Fraternity Ring (9533, -1.13 DPS) [quest]; White Bone Band (11862, -1.23 DPS) [quest]; Falcon's Hook (7552, -1.43 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.22 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Molten Heart of the Mountain (249470, -1.40 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, -0.89 DPS) [world_drop]; Blight (7959, -1.91 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.67 DPS) [world_drop]; Stinging Bow (10624, -0.69 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.69 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Razorsteel Shoulders; back: Blackveil Cape; wrist: Deepfury Bracers; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Battlechaser's Greaves; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1205, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 13323213032215201-050500000000000000-205000000000000000)

Set DPS (verified): 313.1. Weights run: 3.7s. Verify run: 81.6s. 2685 eligible items had no known source.

Race sweep: undead 313.1 ± 0.4, windshaper-skyborne 310.9 ± 0.4, troll 309.6 ± 0.4, orc 307.8 ± 0.4, tauren 306.3 ± 0.4 (dps); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.560 ± 0.142, crit=3.657 ± 0.203 per rating point (14 rating = 1%, 51.196 per %), hit=4.312 ± 0.521 per rating point (10 rating = 1%, 43.125 per %), melee_haste=26.140 ± 4.887

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 224.6 attack_power points (18.97 DPS) | yes | Champion's Plate Helm (227043, -7.46 DPS) [pvp]; Fury Visor (20521, -7.97 DPS) [quest]; Mask of the Unforgiven (13404, -9.42 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 77.2 attack_power points (6.52 DPS) | yes | Emberfury Talisman (12929, -2.20 DPS) [dungeon]; Skibi's Pendant (13089, -2.87 DPS) [world_drop]; Star of Mystaria (12103, -8.32 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 120.3 attack_power points (10.16 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -0.02 DPS) [pvp]; Truestrike Shoulders (12927, -2.67 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackveil Cape (11626, -0.89 DPS) [dungeon]; Deathguard's Cloak (20068, -0.98 DPS) [rep]; Howler's Furs (272414, -1.74 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Savage Gladiator Chain (11726, -1.57 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -2.95 DPS) [quest]; Dawn Armor (252483, -3.48 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Vambraces of the Sadist (13400, -0.46 DPS) [dungeon]; Slashclaw Bracers (13211, -0.81 DPS) [dungeon]; Battleborn Armbraces (12936, -2.37 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Trueaim Gauntlets (13255, +0.00 DPS) [dungeon]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, -4.67 DPS, sim-verified) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Chain Girdle (20150, -0.36 DPS) [rep]; Defiler's Leather Girdle (20190, -0.36 DPS) [rep]; Marksman's Girdle (22232, -3.78 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Plate Legguards (22651, -2.84 DPS) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 94.3 attack_power points (7.97 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -1.07 DPS) [dungeon]; Shadowcraft Boots (16711, -2.33 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.47 DPS) [dungeon]; Blackstone Ring (17713, -3.82 DPS) [dungeon]; Naglering (11669, -7.26 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.99 DPS) [dungeon]; Blackstone Ring (17713, -1.34 DPS) [dungeon]; Legionnaire's Band (19510, -2.26 DPS) [rep] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+7.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Heart of Wyrmthalak (22321, -3.87 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.81 DPS) [quest]; Ancient Bone Bow (18680, -1.30 DPS) [dungeon]; Blackcrow (12651, -2.69 DPS, sim-verified) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: The Unstoppable Force

No-known-source sample (15 of 2685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 948.9. Weights run: 3.1s. Verify run: 77.3s. 2685 eligible items had no known source.

Race sweep: windshaper-skyborne 953.5 ± 1.1, undead 948.9 ± 1.1, orc 947.9 ± 1.1, troll 945.4 ± 1.1, tauren 939.6 ± 1.1 (dps); windshaper-skyborne beats the guide race undead by 4.5 (0.48%) ± 1.5, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.870 ± 0.149, crit=3.690 ± 0.207 per rating point (14 rating = 1%, 51.659 per %), hit=6.265 ± 0.581 per rating point (10 rating = 1%, 62.652 per %), melee_haste=32.260 ± 5.266

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 268.2 attack_power points (49.45 DPS) | yes | Mask of the Unforgiven (13404, -19.20 DPS, sim-verified) [dungeon]; Champion's Plate Helm (227043, -19.86 DPS) [pvp]; Fury Visor (20521, -21.08 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-verified (+14.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Nacreous Shell Necklace (22403, -2.77 DPS) [dungeon]; Emberfury Talisman (12929, -4.79 DPS) [dungeon]; Star of Mystaria (12103, -14.35 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 129.1 attack_power points (23.81 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -2.06 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of Flames (3475, -0.07 DPS) [world_drop]; Stalwart Cloak (272415, -0.07 DPS) [vendor]; Howler's Furs (272414, -4.84 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -4.25 DPS) [dungeon]; Breastplate of Bloodthirst (12757, -8.11 DPS) [quest]; Dawn Armor (252483, -10.08 DPS, sim-verified) [crafted] |
| wrist | Slashclaw Bracers (13211) | Blackrock Spire: Halycon [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, -0.74 DPS) [rep]; Blackmist Armguards (12966, -1.68 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+15.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gloves (272099, +0.00 DPS) [vendor]; Trueaim Gauntlets (13255, -15.43 DPS, sim-verified) [dungeon] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | 122.9 attack_power points (22.66 DPS) | yes | Sash of the Grand Hunt (22207, -3.17 DPS) [dungeon]; Valiant's Waistguard (272403, -4.29 DPS) [vendor]; Brigam Girdle (13142, -5.43 DPS) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; General's Plate Leggings (231533, -0.88 DPS) [pvp]; Outrider's Plate Legguards (22651, -7.16 DPS) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 120.0 attack_power points (22.13 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.70 DPS) [dungeon]; Shadowcraft Boots (16711, -7.56 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.42 DPS) [dungeon]; Painweaver Band (13098, -11.18 DPS) [dungeon]; Naglering (11669, -13.40 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.72 DPS) [dungeon]; Painweaver Band (13098, -6.48 DPS) [dungeon]; Channeler's Ring (272406, -7.41 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+17.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Molten Heart of the Mountain (249470, -7.05 DPS, sim-verified) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Runeblade of Baron Rivendare (13505, -17.97 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.52 DPS) [quest]; Ancient Bone Bow (18680, -2.65 DPS) [dungeon]; Blackcrow (12651, -7.55 DPS, sim-verified) [dungeon] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Slashclaw Bracers; waist: Marksman's Girdle; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

