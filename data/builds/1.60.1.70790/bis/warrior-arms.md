# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 13320200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 31.1. Weights run: 3.6s. Verify run: 6.4s. 349 eligible items had no known source.

Race sweep: high-order-skyborne 31.6 ± 0.0, night-elf 31.1 ± 0.0, human 31.1 ± 0.0, dwarf 30.7 ± 0.0, gnome 30.7 ± 0.0 (dps); high-order-skyborne beats the guide race night-elf by 0.4 (1.41%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.040 ± 0.024, crit=0.578 ± 0.023 per rating point (14 rating = 1%, 8.085 per %), hit=0.697 ± 0.081 per rating point (10 rating = 1%, 6.969 per %), melee_haste=8.241 ± 0.688

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.40 DPS) [crafted]; Brawler's Leather Hood (252504, -0.49 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.2 attack_power points (0.26 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.51 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.03 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon]; Grave Shroud (279865, -0.08 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.84 DPS) | yes | Veteran's Chain Shirt (250488, -0.08 DPS) [crafted]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.12 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.42 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.67 DPS) | yes | Gloves of the Fang (10413, -0.07 DPS) [dungeon]; Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Veteran's Gloves (250508, -0.12 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.75 DPS) | yes | Cobrahn's Grasp (6460, -0.04 DPS) [dungeon]; Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Deviate Scale Belt (6468, -0.29 DPS) [crafted] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 23.2 attack_power points (0.97 DPS) | yes | Defender's Leather Pants (252445, -0.04 DPS) [crafted]; Chausses of Westfall (6087, -0.05 DPS) [quest]; Brawler's Leather Pants (252500, -0.08 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.2 attack_power points (0.64 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.22 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.2 attack_power points (0.51 DPS) | yes | The 1 Ring (8350, -0.38 DPS) [world]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon]; Ring of the Moon (12052, -0.43 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | The 1 Ring (8350, -0.21 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Ring of the Moon (12052, -0.25 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 302.2 attack_power points (12.67 DPS) | yes | Duskbringer (2205, -0.46 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.52 DPS) [dungeon]; Sword of the Fallen (286977, -0.59 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 attack_power points (0.17 DPS) | yes | Dwarven Fishing Pole (3567, -0.01 DPS) [quest]; Fine Longbow (11304, -0.01 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.01 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 30 (night-elf, 13323213030000000-000000000000000000-000000000000000000)

Set DPS (verified): 63.1. Weights run: 4.2s. Verify run: 14.4s. 596 eligible items had no known source.

Race sweep: night-elf 63.1 ± 0.1, human 62.4 ± 0.1, gnome 61.5 ± 0.1, dwarf 61.3 ± 0.1, high-order-skyborne 60.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.838 ± 0.055, crit=1.420 ± 0.062 per rating point (14 rating = 1%, 19.882 per %), hit=1.105 ± 0.169 per rating point (10 rating = 1%, 11.050 per %), melee_haste=11.706 ± 1.448

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 34.5 attack_power points (1.61 DPS) | yes | Tusken Helm (6686, -0.48 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.49 DPS) [crafted]; Defender's Leather Helm (252455, -0.49 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 attack_power points (0.71 DPS) | yes | Sentinel's Medallion (19541, -0.03 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.2 attack_power points (1.41 DPS) | yes | Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.53 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -0.55 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.19 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.32 DPS, sim-verified) [pvp]; Glowing Lizardscale Cloak (6449, -0.36 DPS) [dungeon] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Avenger's Armor (1488, -0.05 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.14 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.20 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.0 attack_power points (0.89 DPS) | yes | Yorgen Bracers (13012, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.14 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.0 attack_power points (1.16 DPS) | yes | Mail Combat Gauntlets (4075, -0.08 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Gauntlets of Ogre Strength (3341, -0.14 DPS) [world] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 31.0 attack_power points (1.44 DPS) | yes | Prowler's Leather Belt (252459, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.12 DPS) [crafted]; Girdle of Golem Strength (9405, -0.33 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 34.9 attack_power points (1.62 DPS) | yes | Brawler's Leather Legguards (252516, -0.20 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.28 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.29 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.7 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Alacritous Treads (277234, -0.07 DPS) [quest]; Veteran's Boots (250503, -0.16 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 23.0 attack_power points (1.07 DPS) | yes | Seal of Wrynn (2933, -0.54 DPS) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon]; Tiger Band (6749, -0.61 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.5 attack_power points (1.00 DPS) | yes | Ring of Precision (1491, -0.49 DPS) [dungeon]; Tiger Band (6749, -0.54 DPS) [quest]; Seal of Wrynn (2933, -0.54 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Felbough Arm (273037, -0.66 DPS) [world]; Viscous Hammer (13045, -5.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 11.5 attack_power points (0.54 DPS) | yes | Golemsight Long Gun (273029, -0.02 DPS) [dungeon]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.19 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 596, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 40 (night-elf, 13323213032214001-000000000000000000-000000000000000000)

Set DPS (verified): 145.8. Weights run: 4.3s. Verify run: 21.2s. 815 eligible items had no known source.

Race sweep: night-elf 145.8 ± 0.2, human 145.0 ± 0.2, dwarf 144.4 ± 0.2, gnome 143.3 ± 0.2, high-order-skyborne 138.4 ± 0.2 (dps); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.096 ± 0.080, crit=2.055 ± 0.096 per rating point (14 rating = 1%, 28.771 per %), hit=2.804 ± 0.239 per rating point (10 rating = 1%, 28.043 per %), melee_haste=17.608 ± 2.144

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 54.8 attack_power points (4.37 DPS) | yes | Chromite Barbute (8142, -0.45 DPS) [dungeon]; White Bandit Mask (10008, -0.77 DPS) [crafted]; Barbaric Iron Helm (7915, -1.43 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.1 attack_power points (1.84 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.24 DPS) [quest]; Kaleidoscope Chain (13084, -0.53 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.88 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.1 attack_power points (2.64 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.68 DPS) [crafted]; Nightscape Shoulders (8192, -0.80 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.31 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.48 DPS) [quest]; Parachute Cloak (10518, -0.62 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, -1.80 DPS) [dungeon]; Kolkar Marauder Chain (6773, -2.85 DPS) [quest]; Carapace of Tuten'kash (10775, -3.32 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 22.4 attack_power points (1.79 DPS) | yes | Hawkeye's Bracers (14590, -0.14 DPS) [world_drop]; Branded Leather Bracers (19508, -0.19 DPS) [dungeon]; Yorgen Bracers (13012, -0.33 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Scarlet Gauntlets (10331, -2.08 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.9 attack_power points (2.94 DPS) | yes | Highlander's Chain Girdle (20089, -0.01 DPS) [rep]; Officer's Belt (250556, -0.34 DPS) [crafted]; Skulker's Leather Belt (252520, -0.48 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.0 attack_power points (3.51 DPS) | yes | Triprunner Dungarees (9624, -0.02 DPS) [quest]; Symbolic Legplates (14829, -0.12 DPS) [world_drop]; Scarlet Leggings (10330, -0.16 DPS) [dungeon] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 40.7 attack_power points (3.24 DPS) | yes | Blackforge Greaves (6423, -0.14 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.29 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.32 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 32.8 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.63 DPS) [dungeon]; Thunderbrow Ring (13097, -0.84 DPS) [world_drop]; Mark of Kern (2262, -1.02 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.0 attack_power points (1.99 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -0.40 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Primitive Fishing Pole (276203, -0.58 DPS) [vendor]; Bonebiter (6830, -0.98 DPS) [quest]; Nightblade (1982, -3.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Swiftwind (13038, -0.29 DPS) [world_drop]; The Silencer (13138, -0.34 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Officer's Boots; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 815, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (night-elf, 13323213032215201-050200000000000000-000000000000000000)

Set DPS (verified): 218.6. Weights run: 4.6s. Verify run: 15.5s. 1030 eligible items had no known source.

Race sweep: high-order-skyborne 219.2 ± 0.3, night-elf 218.6 ± 0.3, human 218.3 ± 0.3, dwarf 218.3 ± 0.3, gnome 216.9 ± 0.3 (dps); high-order-skyborne beats the guide race night-elf by 0.7 (0.30%) ± 0.4, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.660 ± 0.084, crit=1.872 ± 0.109 per rating point (14 rating = 1%, 26.213 per %), hit=2.509 ± 0.276 per rating point (10 rating = 1%, 25.093 per %), melee_haste=23.246 ± 2.811

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 87.3 attack_power points (8.85 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.01 DPS) [vendor]; Raging Berserker's Helm (7719, -3.58 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -3.98 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (3.20 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.17 DPS) [quest]; Sentinel's Medallion (19539, -1.18 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 39.3 attack_power points (3.98 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.20 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.30 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight's Plate Hauberk (220794, -0.88 DPS) [vendor]; Mixologist's Tunic (12793, -1.08 DPS) [dungeon]; Warbear Harness (15064, -1.32 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.9 attack_power points (3.33 DPS) | yes | Officer's Wristguards (250581, -0.10 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.33 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.40 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.38 DPS, sim-verified) [crafted] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 48.2 attack_power points (4.89 DPS) | yes | Prowler's Leather Waistguard (252473, -0.03 DPS) [crafted]; Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Highlander's Chain Girdle (20088, -0.20 DPS) [rep] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 66.8 attack_power points (6.77 DPS) | yes | Gryphon Rider's Leggings (9652, -1.21 DPS) [quest]; Centurion Legplates (10740, -1.21 DPS) [quest]; Serpentskin Leggings (8262, -1.44 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battlechaser's Greaves (12555, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.07 DPS) [crafted]; Officer's Sabatons (250561, -0.34 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.1 attack_power points (4.57 DPS) | yes | Masons Fraternity Ring (9533, -2.21 DPS) [quest]; Thunderbrow Ring (13097, -2.44 DPS) [world_drop]; Falcon's Hook (7552, -2.45 DPS) [dungeon] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 34.9 attack_power points (3.54 DPS) | yes | Masons Fraternity Ring (9533, -1.19 DPS) [quest]; Thunderbrow Ring (13097, -1.42 DPS) [world_drop]; Falcon's Hook (7552, -1.42 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.11 DPS, sim-verified) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.33 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.33 DPS) [world_drop]; Guttbuster (13139, -0.40 DPS) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Highlander's Lamellar Girdle; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1030, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (night-elf, 13323213032215201-050500000000000000-205000000000000000)

Set DPS (verified): 312.0. Weights run: 4.2s. Verify run: 79.8s. 2366 eligible items had no known source.

Race sweep: night-elf 312.0 ± 0.4, high-order-skyborne 311.5 ± 0.4, human 311.1 ± 0.4, dwarf 310.5 ± 0.4, gnome 309.0 ± 0.4 (dps); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.610 ± 0.087, crit=2.300 ± 0.125 per rating point (14 rating = 1%, 32.194 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=15.069 ± 3.085

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 100.4 attack_power points (11.68 DPS) | yes | Field Marshal's Plate Helm (231538, -1.42 DPS) [pvp]; Outlaw's Collar (279253, -2.13 DPS) [crafted]; Eye of Rend (12587, -9.86 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 58.2 attack_power points (6.77 DPS) | yes | Emberfury Talisman (12929, -2.82 DPS, sim-verified) [dungeon]; Imperial Jewel (11933, -3.05 DPS) [dungeon]; Skibi's Pendant (13089, -3.17 DPS) [world_drop] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 75.7 attack_power points (8.80 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, -1.10 DPS) [pvp]; Highlander's Plate Spaulders (20057, -1.43 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.24 DPS) [rep]; Blackveil Cape (11626, -1.12 DPS) [dungeon]; Fel Cape (279269, -1.39 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -2.71 DPS, sim-verified) [dungeon]; Dawn Armor (252483, -4.12 DPS) [crafted]; Breastplate of Bloodthirst (12757, -4.28 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 52.6 attack_power points (6.12 DPS) | yes | Berserker Bracers (19578, -0.20 DPS) [rep]; Vambraces of the Sadist (13400, -0.74 DPS) [dungeon]; Marshal's Plate Bracers (16481, -0.85 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gloves (271948, +0.00 DPS) [rep]; Theramore Gauntlets (271940, -1.90 DPS, sim-verified) [rep] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | 66.6 attack_power points (7.75 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Plate Girdle (20041, -0.05 DPS) [rep]; Highlander's Chain Girdle (20043, -0.05 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -0.27 DPS) [vendor]; Sentinel's Leather Pants (237818, -1.50 DPS) [vendor]; Marshal's Plate Legguards (231540, -1.90 DPS) [pvp] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.9 attack_power points (5.92 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Battlechaser's Greaves (12555, -0.23 DPS) [dungeon]; Highlander's Plate Greaves (20048, -0.41 DPS) [rep] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magni's Will (12548, +0.00 DPS) [quest]; Painweaver Band (13098, +0.00 DPS) [dungeon]; Band of Rumination (18103, -1.11 DPS) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Band of Rumination (18103, -1.63 DPS) [dungeon]; Magni's Will (12548, -3.34 DPS, sim-verified) [quest] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heart of Wyrmthalak (22321, -2.86 DPS, sim-verified) [dungeon] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -2.61 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.55 DPS) [world_drop]; Riphook (12653, -0.61 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.85 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; waist: Valiant's Waistguard; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Hand of Justice; main_hand: Forest Defender's Axe; ranged: Bloodseeker

No-known-source sample (15 of 2366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60, raid preset (night-elf, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 954.2. Weights run: 3.6s. Verify run: 84.5s. 2366 eligible items had no known source.

Race sweep: night-elf 954.2 ± 1.1, high-order-skyborne 954.2 ± 1.0, human 950.6 ± 1.1, gnome 947.9 ± 1.1, dwarf 940.3 ± 1.1 (dps); the guide race night-elf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.870 ± 0.149, crit=3.690 ± 0.207 per rating point (14 rating = 1%, 51.659 per %), hit=6.265 ± 0.581 per rating point (10 rating = 1%, 62.652 per %), melee_haste=32.260 ± 5.266

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 268.2 attack_power points (49.45 DPS) | yes | Mask of the Unforgiven (13404, -16.83 DPS) [dungeon]; Lieutenant Commander's Plate Helm (23314, -19.86 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -5.09 DPS) [dungeon]; Cloak of Flames (3475, -5.16 DPS) [world_drop]; Stalwart Cloak (272415, -5.16 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -7.71 DPS) [dungeon]; Timbermaw Tunic (252484, -8.22 DPS) [crafted]; Breastplate of Bloodthirst (12757, -16.33 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.94 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Brigam Girdle (13142) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Valiant's Waistguard (272403, +0.00 DPS) [vendor]; Marshal's Plate Girdle (16482, -0.74 DPS) [pvp] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -7.57 DPS) [vendor]; Marshal's Plate Legguards (231540, -10.41 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.78 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 120.0 attack_power points (22.13 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -2.17 DPS) [dungeon]; Shadowcraft Boots (16711, -7.56 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.42 DPS) [dungeon]; Painweaver Band (13098, -11.18 DPS) [dungeon]; Magni's Will (12548, -11.70 DPS) [quest] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.72 DPS) [dungeon]; Painweaver Band (13098, -6.48 DPS) [dungeon]; Magni's Will (12548, -7.00 DPS) [quest] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Bloodseeker (19107, -0.46 DPS) [quest]; Ancient Bone Bow (18680, -1.59 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Brigam Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Molten Heart of the Mountain; main_hand: Blackblade of Shahram

No-known-source sample (15 of 2366, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (undead, 13320200000000000-000000000000000000-000000000000000000)

Set DPS (verified): 32.6. Weights run: 3.6s. Verify run: 8.6s. 332 eligible items had no known source.

Race sweep: windshaper-skyborne 32.8 ± 0.0, undead 32.6 ± 0.0, troll 31.8 ± 0.0, orc 31.6 ± 0.0, tauren 31.5 ± 0.0 (dps); windshaper-skyborne beats the guide race undead by 0.3 (0.80%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.040 ± 0.024, crit=0.578 ± 0.023 per rating point (14 rating = 1%, 8.085 per %), hit=0.697 ± 0.081 per rating point (10 rating = 1%, 6.969 per %), melee_haste=8.241 ± 0.688

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.40 DPS) [crafted]; Brawler's Leather Hood (252504, -0.49 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.2 attack_power points (0.26 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.25 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.03 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon]; Subterranean Cape (14149, -0.08 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, +0.00 DPS) [crafted]; Mutant Scale Breastplate (6627, -1.03 DPS, sim-verified) [dungeon] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.42 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.17 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.12 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.75 DPS) | yes | Brawler's Leather Belt (252428, -0.24 DPS) [crafted]; Deviate Scale Belt (6468, -0.29 DPS) [crafted]; Cobrahn's Grasp (6460, -1.30 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.98 DPS, sim-verified) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.2 attack_power points (0.64 DPS) | yes | Feet of the Lynx (1121, -0.04 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.22 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.2 attack_power points (0.51 DPS) | yes | Loop of Sacrifice (281673, -0.26 DPS) [quest]; Bounty Hunter's Ring (5351, -0.38 DPS) [quest]; The 1 Ring (8350, -0.38 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest]; The 1 Ring (8350, -0.21 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (12.71 DPS) | yes | Hammerbone (270018, -0.12 DPS) [quest]; Duskbringer (2205, -0.36 DPS) [dungeon]; Smite's Mighty Hammer (7230, -0.69 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.2 attack_power points (0.17 DPS) | yes | Fine Longbow (11304, -0.01 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.01 DPS) [crafted]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Forsaken Greataxe; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 332, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (undead, 13323213030000000-000000000000000000-000000000000000000)

Set DPS (verified): 63.2. Weights run: 4.2s. Verify run: 13.1s. 581 eligible items had no known source.

Race sweep: undead 63.2 ± 0.1, troll 62.1 ± 0.1, orc 61.9 ± 0.1, tauren 61.5 ± 0.1, windshaper-skyborne 60.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.838 ± 0.055, crit=1.420 ± 0.062 per rating point (14 rating = 1%, 19.882 per %), hit=1.105 ± 0.169 per rating point (10 rating = 1%, 11.050 per %), melee_haste=11.706 ± 1.448

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 34.5 attack_power points (1.61 DPS) | yes | Tusken Helm (6686, -0.49 DPS, sim-verified) [dungeon]; Veteran's Chain Helm (250498, -0.49 DPS) [crafted]; Defender's Leather Helm (252455, -0.49 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.4 attack_power points (0.71 DPS) | yes | Scout's Medallion (19537, -0.03 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.2 attack_power points (1.41 DPS) | yes | Barbaric Iron Shoulders (7913, -0.37 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Mantle of Thieves (2264, -0.55 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.19 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.36 DPS) [dungeon]; Swiftrunner Cape (6745, -0.36 DPS) [quest] |
| chest | Veteran's Silvered Chain Shirt (250518) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Avenger's Armor (1488, -0.05 DPS) [dungeon]; Shining Silver Breastplate (2870, -0.14 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.20 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.0 attack_power points (0.89 DPS) | yes | Yorgen Bracers (13012, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.14 DPS) [dungeon]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.0 attack_power points (1.16 DPS) | yes | Mail Combat Gauntlets (4075, -0.08 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Gauntlets of Ogre Strength (3341, -0.14 DPS) [world] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 31.0 attack_power points (1.44 DPS) | yes | Prowler's Leather Belt (252459, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.12 DPS) [crafted]; Girdle of Golem Strength (9405, -0.33 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 34.9 attack_power points (1.62 DPS) | yes | Brawler's Leather Legguards (252516, -0.20 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.28 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.29 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.7 attack_power points (0.96 DPS) | yes | Brawler's Leather Boots (252439, -0.07 DPS) [crafted]; Veteran's Boots (250503, -0.16 DPS) [crafted]; Stomping Boots (3741, -0.25 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 23.0 attack_power points (1.07 DPS) | yes | Band of the Fist (17694, -0.36 DPS) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon]; Tiger Band (6749, -0.61 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.5 attack_power points (1.00 DPS) | yes | Band of the Fist (17694, -0.29 DPS) [quest]; Ring of Precision (1491, -0.49 DPS) [dungeon]; Tiger Band (6749, -0.54 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, -0.55 DPS) [dungeon]; Felbough Arm (273037, -0.66 DPS) [world]; Viscous Hammer (13045, -5.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 11.5 attack_power points (0.54 DPS) | yes | Golemsight Long Gun (273029, -0.02 DPS) [dungeon]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.19 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Veteran's Silvered Chain Shirt; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 581, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4196 Feathered Mantle

### Band 40 (undead, 13323213032214001-000000000000000000-000000000000000000)

Set DPS (verified): 145.1. Weights run: 4.3s. Verify run: 21.3s. 797 eligible items had no known source.

Race sweep: undead 145.1 ± 0.2, orc 143.2 ± 0.2, troll 142.8 ± 0.2, tauren 142.4 ± 0.2, windshaper-skyborne 138.0 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.096 ± 0.080, crit=2.055 ± 0.096 per rating point (14 rating = 1%, 28.771 per %), hit=2.804 ± 0.239 per rating point (10 rating = 1%, 28.043 per %), melee_haste=17.608 ± 2.144

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 54.8 attack_power points (4.37 DPS) | yes | Chromite Barbute (8142, -0.45 DPS) [dungeon]; White Bandit Mask (10008, -0.77 DPS) [crafted]; Barbaric Iron Helm (7915, -1.43 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.1 attack_power points (1.84 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.24 DPS) [quest]; Ethereal Talisman (4430, -0.37 DPS) [quest]; Kaleidoscope Chain (13084, -0.53 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.1 attack_power points (2.64 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.68 DPS) [crafted]; Nightscape Shoulders (8192, -0.80 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (145.2 DPS) | yes | Imperial Cloak (6432, -0.31 DPS) [dungeon]; Parachute Cloak (10518, -0.31 DPS) [crafted]; Tigerstrike Mantle (13108, -0.31 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, -1.80 DPS) [dungeon]; Kolkar Marauder Chain (6773, -2.85 DPS) [quest]; Carapace of Tuten'kash (10775, -3.32 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 22.4 attack_power points (1.79 DPS) | yes | Hawkeye's Bracers (14590, -0.14 DPS) [world_drop]; Branded Leather Bracers (19508, -0.19 DPS) [dungeon]; Yorgen Bracers (13012, -0.33 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (145.2 DPS) | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Scarlet Gauntlets (10331, -2.42 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.9 attack_power points (2.94 DPS) | yes | Defiler's Chain Girdle (20153, -0.01 DPS) [rep]; Officer's Belt (250556, -0.34 DPS) [crafted]; Skulker's Leather Belt (252520, -0.48 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.0 attack_power points (3.51 DPS) | yes | Triprunner Dungarees (9624, -0.02 DPS) [quest]; Symbolic Legplates (14829, -0.12 DPS) [world_drop]; Scarlet Leggings (10330, -0.16 DPS) [dungeon] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 40.7 attack_power points (3.24 DPS) | yes | Blackforge Greaves (6423, -0.14 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.29 DPS) [crafted]; Prowler's Leather Shoes (252465, -0.32 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 32.8 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.63 DPS) [dungeon]; Thunderbrow Ring (13097, -0.84 DPS) [world_drop]; Mark of Kern (2262, -1.02 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.0 attack_power points (1.99 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop]; Mark of Kern (2262, -0.40 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (145.2 DPS) | yes | Primitive Fishing Pole (276203, -0.58 DPS) [vendor]; Thornstone Sledgehammer (1722, -2.81 DPS) [world_drop]; Nightblade (1982, -2.97 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (145.2 DPS) | yes | Master Hunter's Rifle (17687, -0.15 DPS) [quest]; Swiftwind (13038, -0.29 DPS) [world_drop]; The Silencer (13138, -0.34 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 797, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (undead, 13323213032215201-050200000000000000-000000000000000000)

Set DPS (verified): 221.7. Weights run: 4.6s. Verify run: 16.5s. 1011 eligible items had no known source.

Race sweep: undead 221.7 ± 0.3, windshaper-skyborne 221.1 ± 0.3, troll 220.1 ± 0.3, orc 218.9 ± 0.3, tauren 217.7 ± 0.3 (dps); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.660 ± 0.084, crit=1.872 ± 0.109 per rating point (14 rating = 1%, 26.213 per %), hit=2.509 ± 0.276 per rating point (10 rating = 1%, 25.093 per %), melee_haste=23.246 ± 2.811

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fury Visor (20521) | Voodoo Feathers [quest] | 87.3 attack_power points (8.85 DPS) | yes | Blood Guard's Plate Helm (220803, -1.01 DPS) [vendor]; Raging Berserker's Helm (7719, -2.19 DPS, sim-verified) [dungeon]; Embrace of the Lycan (9479, -3.98 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (3.20 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.17 DPS) [quest]; Scout's Medallion (19535, -1.18 DPS) [rep]; Woven Ivy Necklace (19159, -1.46 DPS, sim-verified) [quest] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 39.3 attack_power points (3.98 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.20 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.24 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Pridelord Cape (14673, -0.40 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -1.08 DPS) [dungeon]; Warbear Harness (15064, -1.32 DPS) [crafted]; Stone Guard's Plate Armor (220801, -1.42 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 32.9 attack_power points (3.33 DPS) | yes | Officer's Wristguards (250581, -0.10 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.33 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.40 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Officer's Gloves (250551, -3.30 DPS, sim-verified) [crafted] |
| waist | Defiler's Lamellar Girdle (20179) | The Defilers [rep] | 48.2 attack_power points (4.89 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.20 DPS) [rep]; Prowler's Leather Waistguard (252473, -1.40 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 66.8 attack_power points (6.77 DPS) | yes | Stormshroud Pants (15057, -1.46 DPS) [crafted]; Golem Shard Leggings (13074, -2.31 DPS) [world_drop]; Serpentskin Leggings (8262, -3.07 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battlechaser's Greaves (12555, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.07 DPS) [crafted]; Officer's Sabatons (250561, -0.34 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 45.1 attack_power points (4.57 DPS) | yes | White Bone Band (11862, -2.14 DPS) [quest]; Masons Fraternity Ring (9533, -2.21 DPS) [quest]; Thunderbrow Ring (13097, -2.44 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 34.9 attack_power points (3.54 DPS) | yes | White Bone Band (11862, -1.11 DPS) [quest]; Masons Fraternity Ring (9533, -1.19 DPS) [quest]; Thunderbrow Ring (13097, -1.42 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.35 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Taran Icebreaker (2915, -1.13 DPS) [world_drop]; Blight (7959, -1.54 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.33 DPS) [dungeon]; Guttbuster (13139, -0.40 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.54 DPS, sim-verified) [world_drop] |

**New at 50:** head: Fury Visor; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Defiler's Lamellar Girdle; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1011, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (undead, 13323213032215201-050500000000000000-205000000000000000)

Set DPS (verified): 308.9. Weights run: 4.2s. Verify run: 85.8s. 2356 eligible items had no known source.

Race sweep: undead 308.9 ± 0.4, windshaper-skyborne 306.2 ± 0.3, troll 305.1 ± 0.3, orc 303.7 ± 0.3, tauren 302.3 ± 0.3 (dps); the guide race undead is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=2.560 ± 0.142, crit=3.657 ± 0.203 per rating point (14 rating = 1%, 51.196 per %), hit=4.312 ± 0.521 per rating point (10 rating = 1%, 43.125 per %), melee_haste=26.140 ± 4.887

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 224.6 attack_power points (18.97 DPS) | yes | Champion's Plate Helm (227043, -7.46 DPS) [pvp]; Mask of the Unforgiven (13404, -7.80 DPS, sim-verified) [dungeon]; Outlaw's Collar (279253, -7.95 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 77.2 attack_power points (6.52 DPS) | yes | Emberfury Talisman (12929, -2.85 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -2.87 DPS) [world_drop]; Star of Mystaria (12103, -2.88 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 120.3 attack_power points (10.16 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Warlord's Plate Shoulders (231534, -0.02 DPS) [pvp]; Truestrike Shoulders (12927, -0.85 DPS) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.08 DPS) [dungeon]; Fel Cape (279269, -1.68 DPS) [crafted]; Blackveil Cape (11626, -1.97 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -1.59 DPS) [dungeon]; Timbermaw Tunic (252484, -2.18 DPS) [crafted]; Breastplate of Bloodthirst (12757, -5.12 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Vambraces of the Sadist (13400, -0.46 DPS) [dungeon]; Slashclaw Bracers (13211, -0.81 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, -1.91 DPS, sim-verified) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Chain Girdle (20150, -0.36 DPS) [rep]; Defiler's Leather Girdle (20190, -0.36 DPS) [rep]; Marksman's Girdle (22232, -2.45 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Plate Legguards (22651, -2.84 DPS) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 94.3 attack_power points (7.97 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Bloodmail Boots (14616, -0.86 DPS) [dungeon]; Shadowcraft Boots (16711, -2.33 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -3.47 DPS) [dungeon]; Blackstone Ring (17713, -3.82 DPS) [dungeon]; Naglering (11669, -7.47 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.99 DPS) [dungeon]; Blackstone Ring (17713, -1.34 DPS) [dungeon]; Legionnaire's Band (19510, -2.26 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -1.59 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted]; Hand of Justice (11815, -2.04 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Battle Axe (234543, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Bloodseeker (19107, -0.16 DPS) [quest]; Dark Iron Rifle (16004, -3.52 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force

No-known-source sample (15 of 2356, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (undead, 30303213132215201-050500050120000000-100000000000000000)

Set DPS (verified): 926.0. Weights run: 3.6s. Verify run: 85.4s. 2356 eligible items had no known source.

Race sweep: windshaper-skyborne 930.5 ± 1.1, undead 926.0 ± 1.1, orc 923.9 ± 1.1, troll 923.3 ± 1.1, tauren 915.5 ± 1.1 (dps); windshaper-skyborne beats the guide race undead by 4.5 (0.48%) ± 1.5, beyond error.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.003, agility=1.844 ± 0.091, crit=2.391 ± 0.125 per rating point (14 rating = 1%, 33.469 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=13.623 ± 2.853

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 106.5 attack_power points (25.13 DPS) | yes | Warlord's Plate Headpiece (231535, -2.70 DPS) [pvp]; Outlaw's Collar (279253, -4.15 DPS) [crafted]; Eye of Rend (12587, -34.49 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Emberfury Talisman (12929, -0.36 DPS) [dungeon]; Imperial Jewel (11933, -0.70 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 83.3 attack_power points (19.64 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Defiler's Lamellar Spaulders (20184, -2.90 DPS) [rep]; Defiler's Plate Spaulders (20212, -2.90 DPS) [rep] |
| back | Cloak of Flames (3475) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Deathguard's Cloak (20068, +0.00 DPS) [rep]; Cape of the Black Baron (13340, -6.70 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -6.38 DPS, sim-verified) [dungeon]; Dawn Armor (252483, -8.86 DPS) [crafted]; Breastplate of Bloodthirst (12757, -10.31 DPS) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.2 attack_power points (13.97 DPS) | yes | Berserker Bracers (19578, -0.63 DPS) [rep]; General's Plate Armguards (16546, -2.11 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Storm Gauntlets (12632, +0.00 DPS) [crafted]; General's Plate Gauntlets (231532, -0.52 DPS) [pvp]; Raider Gloves (272099, -1.34 DPS) [vendor] |
| waist | Valiant's Waistguard (272403) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.72 DPS) [crafted]; Bloodmail Belt (14614, -3.63 DPS, sim-verified) [dungeon] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Plate Legguards (22651, -0.70 DPS) [rep]; Warbear Woolies (15065, -0.76 DPS) [crafted] |
| feet | Bloodmail Boots (14616) | Scholomance [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Plate Greaves (20208, +0.00 DPS) [rep]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Battlechaser's Greaves (12555, -15.61 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Myrmidon's Signet (2246, -2.78 DPS) [world_drop]; Band of Rumination (18103, -3.12 DPS) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (+18.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Myrmidon's Signet (2246, -2.96 DPS) [world_drop]; Band of Rumination (18103, -3.30 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Molten Heart of the Mountain (249470, -12.16 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| main_hand | Blackblade of Shahram (12592) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -17.13 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.11 DPS) [world_drop]; Riphook (12653, -2.01 DPS) [dungeon]; Guttbuster (13139, -2.16 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Darkspear Pauldrons; back: Cloak of Flames; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Valiant's Waistguard; legs: Outrider's Chain Leggings; feet: Bloodmail Boots; finger1: Legionnaire's Band; finger2: Don Julio's Band; main_hand: Blackblade of Shahram; ranged: Bloodseeker

No-known-source sample (15 of 2356, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

