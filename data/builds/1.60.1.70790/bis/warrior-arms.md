# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 13322000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.1. Weights run: 4.7s. Verify run: 6.0s. 330 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.096 ± 0.023, crit=0.603 ± 0.022 per rating point (14 rating = 1%, 8.436 per %), hit=0.834 ± 0.073 per rating point (10 rating = 1%, 8.341 per %), melee_haste=7.450 ± 0.673

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.80 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.36 DPS) [crafted]; Brawler's Leather Hood (252504, -0.45 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.6 attack_power points (0.26 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.50 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.24 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.02 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-verified (30.1 DPS) | yes | Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.06 DPS) [dungeon]; Grave Shroud (279865, -0.08 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.80 DPS) | yes | Veteran's Chain Shirt (250488, -0.06 DPS) [crafted]; Brawler's Leather Armor (252490, -0.09 DPS) [crafted]; Defender's Leather Armor (252434, -0.11 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.40 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bravo's Armbands (270015, -0.03 DPS) [quest]; Bristlebark Bindings (14569, -0.11 DPS) [world_drop] |
| hands | Thorbia's Gauntlets (12994) | World drop [world_drop] | 16.0 attack_power points (0.64 DPS) | yes | Gloves of the Fang (10413, -0.06 DPS) [dungeon]; Gold-flecked Gloves (5195, -0.08 DPS) [dungeon]; Veteran's Gloves (250508, -0.11 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.72 DPS) | yes | Cobrahn's Grasp (6460, -0.03 DPS) [dungeon]; Brawler's Leather Belt (252428, -0.22 DPS) [crafted]; Deviate Scale Belt (6468, -0.26 DPS) [crafted] |
| legs | Veteran's Chain Leggings (250493) | Blacksmithing [crafted] | 23.5 attack_power points (0.94 DPS) | yes | Defender's Leather Pants (252445, -0.04 DPS) [crafted]; Chausses of Westfall (6087, -0.06 DPS) [quest]; Brawler's Leather Pants (252500, -0.06 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.5 attack_power points (0.62 DPS) | yes | Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Defender's Leather Boots (252441, -0.22 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.4 attack_power points (0.49 DPS) | yes | The 1 Ring (8350, -0.37 DPS) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon]; Ring of the Moon (12052, -0.41 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.32 DPS) | yes | The 1 Ring (8350, -0.19 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 302.4 attack_power points (12.07 DPS) | yes | Duskbringer (2205, -0.40 DPS, sim-verified) [dungeon]; Rockslicer (872, -0.50 DPS) [dungeon]; Sword of the Fallen (286977, -0.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.4 attack_power points (0.18 DPS) | yes | Dwarven Fishing Pole (3567, -0.02 DPS) [quest]; Fine Longbow (11304, -0.02 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.02 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Thorbia's Gauntlets; waist: Blackened Defias Belt; legs: Veteran's Chain Leggings; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 330, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade

### Band 30 (human, 13323213030000000-000000000000000000-000000000000000000)

Set DPS (verified): 63.1. Weights run: 5.3s. Verify run: 16.4s. 568 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.562 ± 0.058, crit=1.381 ± 0.064 per rating point (14 rating = 1%, 19.336 per %), hit=1.331 ± 0.172 per rating point (10 rating = 1%, 13.308 per %), melee_haste=10.129 ± 1.340

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 32.1 attack_power points (1.41 DPS) | yes | Tusken Helm (6686, -0.27 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.35 DPS) [crafted]; Defender's Leather Helm (252455, -0.35 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.2 attack_power points (0.63 DPS) | yes | Sentinel's Medallion (19541, -0.08 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.10 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.2 attack_power points (1.19 DPS) | yes | Barbaric Iron Shoulders (7913, -0.25 DPS) [crafted]; Barbaric Shoulders (5964, -0.41 DPS) [crafted]; Mantle of Thieves (2264, -0.51 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (63.1 DPS) | yes | Sergeant Major's Cape (16315, -0.12 DPS) [pvp]; Tigerstrike Mantle (13108, -0.19 DPS) [world_drop]; Wolfmaster Cape (6314, -0.30 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-verified (63.1 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.03 DPS) [crafted]; Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Slayer's Surcoat (14751, -0.23 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.4 attack_power points (0.76 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Pugilist Bracers (4438, -0.06 DPS) [dungeon]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.4 attack_power points (1.03 DPS) | yes | Mail Combat Gauntlets (4075, -0.05 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.4 attack_power points (1.29 DPS) | yes | Prowler's Leather Belt (252459, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.15 DPS) [crafted]; Girdle of Golem Strength (9405, -0.24 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.9 attack_power points (1.45 DPS) | yes | Brawler's Leather Legguards (252516, -0.21 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.26 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.30 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.5 attack_power points (0.81 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Alacritous Treads (277234, -0.03 DPS) [quest]; Veteran's Boots (250503, -0.10 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.4 attack_power points (0.94 DPS) | yes | Seal of Wrynn (2933, -0.47 DPS) [quest]; Tiger Band (6749, -0.50 DPS) [quest]; Bonebinder's Signet (273032, -0.50 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.7 attack_power points (0.91 DPS) | yes | Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world]; Seal of Wrynn (2933, -0.57 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (63.1 DPS) | yes | Corpsemaker (6687, -0.46 DPS) [dungeon]; Felbough Arm (273037, -0.58 DPS) [world]; Viscous Hammer (13045, -5.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.7 attack_power points (0.47 DPS) | yes | Golemsight Long Gun (273029, -0.06 DPS) [dungeon]; Gravespike Repeater (273033, -0.12 DPS) [world]; Silver Star (3463, -0.19 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 40 (human, 13323213032214001-000000000000000000-000000000000000000)

Set DPS (verified): 141.8. Weights run: 5.5s. Verify run: 9.5s. 777 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.593 ± 0.062, crit=1.595 ± 0.075 per rating point (14 rating = 1%, 22.329 per %), hit=1.855 ± 0.186 per rating point (10 rating = 1%, 18.549 per %), melee_haste=13.427 ± 1.635

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 48.3 attack_power points (4.38 DPS) | yes | Chromite Barbute (8142, -0.47 DPS) [dungeon]; White Bandit Mask (10008, -0.80 DPS) [crafted]; Barbaric Iron Helm (7915, -1.45 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 17.5 attack_power points (1.59 DPS) | yes | Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.50 DPS) [world]; River Pride Choker (13087, -0.86 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.5 attack_power points (2.49 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.50 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.54 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.40 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.54 DPS) [quest]; Parachute Cloak (10518, -0.80 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kolkar Marauder Chain (6773, -1.60 DPS) [quest]; Carapace of Tuten'kash (10775, -2.18 DPS) [dungeon]; Jouster's Chestplate (8157, -2.43 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.4 attack_power points (1.85 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.26 DPS) [world_drop]; Yorgen Bracers (13012, -0.33 DPS) [world_drop] |
| hands | Scarlet Gauntlets (10331) | Scarlet Monastery: Scarlet Centurion [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plated Fist of Hakoo (13071, -0.36 DPS) [world_drop]; Prowler's Leather Gloves (252524, -0.51 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.54 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.3 attack_power points (2.93 DPS) | yes | Highlander's Chain Girdle (20089, -0.18 DPS) [rep]; Highlander's Leather Girdle (20116, -0.21 DPS) [rep]; Highlander's Plate Girdle (20125, -0.21 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.81 DPS) | yes | Symbolic Legplates (14829, -0.22 DPS) [world_drop]; Firemane Leggings (13129, -0.36 DPS) [world_drop]; Triprunner Dungarees (9624, -0.66 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.1 attack_power points (3.37 DPS) | yes | Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.36 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.51 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 28.7 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.76 DPS) [dungeon]; Mark of Kern (2262, -0.79 DPS) [dungeon]; Assault Band (13095, -0.79 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (1.88 DPS) | yes | Falcon's Hook (7552, -0.04 DPS) [dungeon]; Mark of Kern (2262, -0.07 DPS) [dungeon]; Assault Band (13095, -0.07 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (141.8 DPS) | yes | Primitive Fishing Pole (276203, -1.11 DPS) [vendor]; Bonebiter (6830, -1.12 DPS) [quest]; Nightblade (1982, -2.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.22 DPS) [quest]; The Silencer (13138, -0.25 DPS) [world_drop]; Swiftwind (13038, -0.51 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Scarlet Gauntlets; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 50 (human, 13323213032215201-050200000000000000-000000000000000000)

Set DPS (verified): 207.8. Weights run: 6.1s. Verify run: 19.2s. 982 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.242 ± 0.066, crit=1.426 ± 0.087 per rating point (14 rating = 1%, 19.959 per %), hit=2.205 ± 0.205 per rating point (10 rating = 1%, 22.046 per %), melee_haste=10.008 ± 2.104

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.0 attack_power points (4.89 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.64 DPS) [crafted]; Sunscale Helmet (14849, -0.69 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.1 attack_power points (2.78 DPS) | yes | Sentinel's Medallion (19539, -1.33 DPS, sim-verified) [rep]; Kaleidoscope Chain (13084, -1.40 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.9 attack_power points (3.82 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.21 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.29 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Phantom Cape (13122, -0.09 DPS) [world_drop]; Bloodlust Cape (14801, -0.15 DPS) [world_drop]; Pridelord Cape (14673, -0.24 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warforged Chestplate (11195, -0.41 DPS) [quest]; Knight's Plate Hauberk (220794, -0.63 DPS, sim-verified) [vendor]; Warbear Harness (15064, -0.80 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.4 attack_power points (3.13 DPS) | yes | Bracers of the Stone Princess (17714, -0.15 DPS) [dungeon]; Arena Bands (18711, -0.15 DPS) [world]; Prowler's Leather Bracers (252539, -0.29 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -2.70 DPS, sim-verified) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.9 attack_power points (4.56 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.10 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.16 DPS) [crafted]; Officer's Waistguard (250571, -0.26 DPS) [crafted] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 56.4 attack_power points (5.99 DPS) | yes | Centurion Legplates (10740, -0.82 DPS) [quest]; Serpentskin Leggings (8262, -1.12 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.24 DPS, sim-verified) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.7 attack_power points (4.22 DPS) | yes | Skulker's Leather Boots (252469, -0.16 DPS) [crafted]; Officer's Sabatons (250561, -0.26 DPS) [crafted]; Officer's Boots (250546, -0.53 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.0 attack_power points (4.47 DPS) | yes | Mark of Kern (2262, -2.34 DPS) [dungeon]; Assault Band (13095, -2.34 DPS) [world_drop]; Thunderbrow Ring (13097, -2.37 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.2 attack_power points (3.31 DPS) | yes | Mark of Kern (2262, -1.19 DPS) [dungeon]; Assault Band (13095, -1.19 DPS) [world_drop]; Thunderbrow Ring (13097, -1.22 DPS) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.34 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.31 DPS) [world_drop]; Blight (7959, -2.60 DPS, sim-verified) [crafted]; Bleakwood Hew (12769, -3.44 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Houndmaster's Bow (11628, -0.21 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.28 DPS) [world_drop]; Guttbuster (13139, -0.43 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; wrist: Officer's Wristguards; hands: Gloves of Holy Might; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 982, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 60 (human, 13323213032215201-050500000000000000-205000000000000000)

Set DPS (verified): 289.9. Weights run: 5.9s. Verify run: 91.4s. 1833 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=1.078 ± 0.064, crit=1.540 ± 0.091 per rating point (14 rating = 1%, 21.555 per %), hit=3.123 ± 0.244 per rating point (10 rating = 1%, 31.226 per %), melee_haste=17.142 ± 2.306

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 141.6 attack_power points (16.64 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.50 DPS) [vendor]; Stalwart Helm (250599, -6.31 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.0 attack_power points (2.82 DPS) | yes | Stormpike Soldier's Pendant (19097, -0.71 DPS) [rep]; Sentinel's Medallion (19538, -0.92 DPS) [rep]; Sunprism Pendant (20649, -1.30 DPS) [quest] |
| shoulder | Highlander's Plate Spaulders (20057) (or Highlander's Lamellar Spaulders (20058)) | The League of Arathor [rep] | 54.3 attack_power points (6.38 DPS) | yes | Highlander's Lamellar Spaulders (20058, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.38 DPS) [dungeon]; Cloak of Flames (3475, -0.96 DPS) [world_drop]; Sergeant Major's Cape (16337, -1.74 DPS) [pvp] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, -1.78 DPS, sim-verified) [crafted]; Savage Gladiator Chain (11726, -1.97 DPS) [dungeon]; Field Marshal's Plate Armor (231536, -4.50 DPS) [pvp] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 46.6 attack_power points (5.48 DPS) | yes | Forest Stalker's Bracers (19587, -0.49 DPS) [rep]; Marshal's Plate Bracers (16481, -0.60 DPS) [pvp]; Windtalker's Wristguards (19582, -1.01 DPS) [rep] |
| hands | Theramore Gauntlets (271940) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Plate Gauntlets (231541, -0.51 DPS) [pvp]; Timbermaw Brawlers (19049, -1.03 DPS) [crafted]; Theramore Gloves (271948, -1.73 DPS) [rep] |
| waist | Brigam Girdle (13142) | Blackrock Spire: General Drakkisath [dungeon] | 59.2 attack_power points (6.96 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.25 DPS) [crafted]; Highlander's Plate Girdle (20041, -0.43 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -1.57 DPS, sim-verified) [vendor]; Marshal's Plate Legguards (231540, -1.65 DPS) [pvp]; Sentinel's Lamellar Legguards (22753, -3.94 DPS) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 46.6 attack_power points (5.48 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Highlander's Plate Greaves (20048, -0.67 DPS) [rep]; Highlander's Lamellar Greaves (20049, -0.67 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 66.8 attack_power points (7.85 DPS) | yes | Protector's Band (19514, -3.63 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 51.2 attack_power points (6.02 DPS) | yes | Protector's Band (19514, -1.81 DPS) [rep] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Sunderer (234566, +0.00 DPS) [pvp]; Forest Defender's Axe (279260, -1.98 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.28 DPS) [quest]; Skull Splitting Crossbow (13039, -1.70 DPS) [world_drop]; Houndmaster's Bow (11628, -1.93 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; shoulder: Highlander's Plate Spaulders; back: Cloak of the Honor Guard; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Theramore Gauntlets; waist: Brigam Girdle; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: The Unstoppable Force; ranged: Blackcrow

No-known-source sample (15 of 1833, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

### Band 60, raid preset (human, 34323213032215201-050500200000000000-200000000000000000)

Set DPS (verified): 780.0. Weights run: 4.7s. Verify run: 79.3s. 1833 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.494 ± 0.096, crit=1.986 ± 0.131 per rating point (14 rating = 1%, 27.808 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=15.458 ± 2.584

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 95.2 attack_power points (21.42 DPS) | yes | Field Marshal's Plate Helm (231538, -1.31 DPS) [pvp]; Lieutenant Commander's Plate Helm (23314, -4.77 DPS) [vendor]; Outlaw's Collar (279253, -29.33 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.4 attack_power points (6.85 DPS) | yes | Sentinel's Medallion (19538, -1.80 DPS) [rep]; Stormpike Soldier's Pendant (19097, -2.80 DPS) [rep]; Sunprism Pendant (20649, -2.81 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 68.2 attack_power points (15.34 DPS) | yes | Theramore Epaulets (271956, +0.00 DPS) [rep]; Lieutenant Commander's Plate Shoulders (227045, -0.66 DPS) [pvp]; Highlander's Plate Spaulders (20057, -0.71 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (780.0 DPS) | yes | Cloak of the Honor Guard (20073, -0.21 DPS) [rep]; Sergeant Major's Cape (16337, -2.89 DPS) [pvp]; Fel Cape (279269, -3.29 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (780.0 DPS) | yes | Savage Gladiator Chain (11726, -8.17 DPS, sim-verified) [dungeon]; Field Marshal's Plate Armor (231536, -9.76 DPS) [pvp]; Dawn Armor (252483, -9.95 DPS) [crafted] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 53.8 attack_power points (12.10 DPS) | yes | Forest Stalker's Bracers (19587, -0.26 DPS) [rep]; Marshal's Plate Bracers (16481, -1.33 DPS) [pvp] |
| hands | Theramore Gauntlets (271940) | Theramore Expeditionary Force [rep] | sim-verified (780.0 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.91 DPS) [pvp]; Timbermaw Brawlers (19049, -1.96 DPS) [crafted]; Gloves of Holy Might (867, -5.85 DPS, sim-verified) [world_drop] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 67.1 attack_power points (15.10 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Highlander's Plate Girdle (20041, -0.43 DPS) [rep]; Highlander's Chain Girdle (20043, -1.20 DPS) [rep] |
| legs | Sentinel's Plate Legguards (237825) | Illiyana Moonblaze [vendor] | sim-verified (780.0 DPS) | yes | Sentinel's Chain Leggings (237819, -1.60 DPS) [vendor]; Marshal's Plate Legguards (231540, -3.47 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.29 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.0 attack_power points (11.24 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Highlander's Plate Greaves (20048, -0.27 DPS) [rep]; Highlander's Lamellar Greaves (20049, -0.27 DPS) [rep] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | 42.8 attack_power points (9.64 DPS) | yes | Band of the Penitent (13217, -3.38 DPS) [quest]; Don Rodrigo's Band (21563, -3.38 DPS) [rep]; Thunderbrow Ring (13097, -4.67 DPS) [world_drop] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 41.8 attack_power points (9.41 DPS) | yes | Don Rodrigo's Band (21563, -3.15 DPS) [rep]; Thunderbrow Ring (13097, -4.44 DPS) [world_drop]; Band of the Penitent (13217, -17.54 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (780.0 DPS) | yes | - |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (780.0 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (780.0 DPS) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -28.55 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | Korrak the Bloodrager [quest] | sim-verified (780.0 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.61 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.68 DPS) [world_drop]; Houndmaster's Bow (11628, -2.13 DPS) [dungeon] |

**New at 60:** head: Lionheart Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Theramore Gauntlets; waist: Might of the Timbermaw; legs: Sentinel's Plate Legguards; feet: Scalegut Treaders; finger1: Protector's Band; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Arcanite Champion; ranged: Bloodseeker

No-known-source sample (15 of 1833, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher

## Horde

### Band 20 (orc, 13322000000000000-000000000000000000-000000000000000000)

Set DPS (verified): 30.6. Weights run: 4.7s. Verify run: 8.1s. 313 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.096 ± 0.023, crit=0.603 ± 0.022 per rating point (14 rating = 1%, 8.436 per %), hit=0.834 ± 0.073 per rating point (10 rating = 1%, 8.341 per %), melee_haste=7.450 ± 0.673

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.80 DPS) | yes | Defender's Leather Hood (252447, -0.16 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.36 DPS) [crafted]; Brawler's Leather Hood (252504, -0.45 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.6 attack_power points (0.26 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.24 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.02 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | sim-verified (30.6 DPS) | yes | Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.06 DPS) [dungeon]; Subterranean Cape (14149, -0.08 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.80 DPS) | yes | Veteran's Chain Shirt (250488, -0.06 DPS) [crafted]; Brawler's Leather Armor (252490, -0.09 DPS) [crafted]; Defender's Leather Armor (252434, -0.11 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.40 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.16 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (30.6 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -0.89 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.72 DPS) | yes | Brawler's Leather Belt (252428, -0.22 DPS) [crafted]; Deviate Scale Belt (6468, -0.26 DPS) [crafted]; Cobrahn's Grasp (6460, -1.08 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (30.6 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -0.87 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (30.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -0.71 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.4 attack_power points (0.49 DPS) | yes | Loop of Sacrifice (281673, -0.25 DPS) [quest]; Bounty Hunter's Ring (5351, -0.36 DPS) [quest]; The 1 Ring (8350, -0.37 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.32 DPS) | yes | Loop of Sacrifice (281673, -0.08 DPS) [quest]; Bounty Hunter's Ring (5351, -0.19 DPS) [quest]; The 1 Ring (8350, -0.20 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Forsaken Greataxe (251533) | The Wrath of Rath'mael [quest] | 303.0 attack_power points (12.09 DPS) | yes | Hammerbone (270018, -0.11 DPS) [quest]; Smite's Mighty Hammer (7230, -0.26 DPS, sim-verified) [dungeon]; Duskbringer (2205, -0.34 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.4 attack_power points (0.18 DPS) | yes | Fine Longbow (11304, -0.02 DPS) [vendor]; Cracked Blacksmith Hammer (285279, -0.02 DPS) [crafted]; Deadly Blunderbuss (4369, -0.09 DPS) [crafted] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Forsaken Greataxe; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (orc, 13323213030000000-000000000000000000-000000000000000000)

Set DPS (verified): 62.9. Weights run: 5.3s. Verify run: 13.9s. 553 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.562 ± 0.058, crit=1.381 ± 0.064 per rating point (14 rating = 1%, 19.336 per %), hit=1.331 ± 0.172 per rating point (10 rating = 1%, 13.308 per %), melee_haste=10.129 ± 1.340

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 32.1 attack_power points (1.41 DPS) | yes | Tusken Helm (6686, -0.27 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.35 DPS) [crafted]; Defender's Leather Helm (252455, -0.35 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.2 attack_power points (0.63 DPS) | yes | Scout's Medallion (19537, -0.08 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.10 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.2 attack_power points (1.19 DPS) | yes | Barbaric Iron Shoulders (7913, -0.37 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -0.41 DPS) [crafted]; Mantle of Thieves (2264, -0.51 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (62.9 DPS) | yes | Tigerstrike Mantle (13108, -0.19 DPS) [world_drop]; Wildhunter Cloak (16658, -0.30 DPS) [quest]; Construct Cloak (279848, -0.30 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | sim-verified (62.9 DPS) | yes | Veteran's Silvered Chain Shirt (250518, -0.03 DPS) [crafted]; Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Slayer's Surcoat (14751, -0.23 DPS) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 17.4 attack_power points (0.76 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Pugilist Bracers (4438, -0.06 DPS) [dungeon]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 23.4 attack_power points (1.03 DPS) | yes | Mail Combat Gauntlets (4075, -0.05 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -0.06 DPS) [world]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 29.4 attack_power points (1.29 DPS) | yes | Prowler's Leather Belt (252459, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.15 DPS) [crafted]; Girdle of Golem Strength (9405, -0.24 DPS) [world_drop] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 32.9 attack_power points (1.45 DPS) | yes | Brawler's Leather Legguards (252516, -0.21 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.26 DPS) [world_drop]; Brawler's Leather Pants (252500, -0.30 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.5 attack_power points (0.81 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Veteran's Boots (250503, -0.10 DPS) [crafted]; Stomping Boots (3741, -0.19 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.4 attack_power points (0.94 DPS) | yes | Band of the Fist (17694, -0.31 DPS) [quest]; Tiger Band (6749, -0.50 DPS) [quest]; Bonebinder's Signet (273032, -0.50 DPS) [world] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.7 attack_power points (0.91 DPS) | yes | Band of the Fist (17694, -0.32 DPS, sim-verified) [quest]; Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (62.9 DPS) | yes | Corpsemaker (6687, -0.46 DPS) [dungeon]; Felbough Arm (273037, -0.58 DPS) [world]; Viscous Hammer (13045, -5.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 10.7 attack_power points (0.47 DPS) | yes | Golemsight Long Gun (273029, -0.06 DPS) [dungeon]; Gravespike Repeater (273033, -0.12 DPS) [world]; Silver Star (3463, -0.19 DPS) [quest] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 553, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade

### Band 40 (orc, 13323213032214001-000000000000000000-000000000000000000)

Set DPS (verified): 144.0. Weights run: 5.5s. Verify run: 10.8s. 759 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.593 ± 0.062, crit=1.595 ± 0.075 per rating point (14 rating = 1%, 22.329 per %), hit=1.855 ± 0.186 per rating point (10 rating = 1%, 18.549 per %), melee_haste=13.427 ± 1.635

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 48.3 attack_power points (4.38 DPS) | yes | Chromite Barbute (8142, -0.47 DPS) [dungeon]; White Bandit Mask (10008, -0.80 DPS) [crafted]; Barbaric Iron Helm (7915, -1.45 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 17.5 attack_power points (1.59 DPS) | yes | Ethereal Talisman (4430, -0.10 DPS) [quest]; Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.50 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 27.5 attack_power points (2.49 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.50 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.54 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.40 DPS) [dungeon]; Parachute Cloak (10518, -0.40 DPS) [crafted]; Tigerstrike Mantle (13108, -0.40 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kolkar Marauder Chain (6773, -1.60 DPS) [quest]; Carapace of Tuten'kash (10775, -2.18 DPS) [dungeon]; Jouster's Chestplate (8157, -2.43 DPS) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.4 attack_power points (1.85 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.26 DPS) [world_drop]; Yorgen Bracers (13012, -0.33 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Scarlet Gauntlets (10331, -2.25 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.3 attack_power points (2.93 DPS) | yes | Defiler's Chain Girdle (20153, -0.18 DPS) [rep]; Defiler's Leather Girdle (20192, -0.21 DPS) [rep]; Defiler's Plate Girdle (20206, -0.21 DPS) [rep] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (3.81 DPS) | yes | Symbolic Legplates (14829, -0.22 DPS) [world_drop]; Firemane Leggings (13129, -0.36 DPS) [world_drop]; Triprunner Dungarees (9624, -0.66 DPS) [quest] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.1 attack_power points (3.37 DPS) | yes | Blackforge Greaves (6423, -0.29 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.36 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.51 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 28.7 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.76 DPS) [dungeon]; Mark of Kern (2262, -0.79 DPS) [dungeon]; Assault Band (13095, -0.79 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.8 attack_power points (1.88 DPS) | yes | Falcon's Hook (7552, -0.04 DPS) [dungeon]; Mark of Kern (2262, -0.07 DPS) [dungeon]; Assault Band (13095, -0.07 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Fiery War Axe (870) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Primitive Fishing Pole (276203, -1.11 DPS) [vendor]; Nightblade (1982, -1.66 DPS, sim-verified) [world_drop]; Thornstone Sledgehammer (1722, -3.19 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Rifle (17687, -0.22 DPS) [quest]; The Silencer (13138, -0.25 DPS) [world_drop]; Swiftwind (13038, -0.51 DPS) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Ravager's Armguards; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; trinket1: Ankh of Life; main_hand: Fiery War Axe; ranged: Monolithic Bow

No-known-source sample (15 of 759, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (orc, 13323213032215201-050200000000000000-000000000000000000)

Set DPS (verified): 208.9. Weights run: 6.1s. Verify run: 18.6s. 963 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.242 ± 0.066, crit=1.426 ± 0.087 per rating point (14 rating = 1%, 19.959 per %), hit=2.205 ± 0.205 per rating point (10 rating = 1%, 22.046 per %), melee_haste=10.008 ± 2.104

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.0 attack_power points (4.89 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -0.64 DPS) [crafted]; Sunscale Helmet (14849, -0.69 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.1 attack_power points (2.78 DPS) | yes | Woven Ivy Necklace (19159, -0.32 DPS) [quest]; Ethereal Talisman (4430, -1.19 DPS) [quest]; Scout's Medallion (19535, -1.19 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.9 attack_power points (3.82 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.21 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.29 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodlust Cape (14801, -0.07 DPS) [world_drop]; Pridelord Cape (14673, -0.15 DPS) [world_drop]; Blackflame Cape (13109, -0.40 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stone Guard's Plate Armor (220801, -0.20 DPS) [vendor]; Warforged Chestplate (11195, -0.41 DPS) [quest]; Warbear Harness (15064, -0.80 DPS) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.4 attack_power points (3.13 DPS) | yes | Bracers of the Stone Princess (17714, -0.15 DPS) [dungeon]; Arena Bands (18711, -0.15 DPS) [world]; Prowler's Leather Bracers (252539, -0.29 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (208.9 DPS) | yes | Stonerender Gauntlets (17007, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Officer's Gloves (250551, -2.07 DPS, sim-verified) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.9 attack_power points (4.56 DPS) | yes | Defiler's Lamellar Girdle (20179, -0.10 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.16 DPS) [crafted]; Officer's Waistguard (250571, -0.26 DPS) [crafted] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 56.4 attack_power points (5.99 DPS) | yes | Golem Shard Leggings (13074, -1.32 DPS) [world_drop]; Serpentskin Leggings (8262, -1.52 DPS, sim-verified) [world_drop]; Scarlet Leggings (10330, -1.53 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.7 attack_power points (4.22 DPS) | yes | Skulker's Leather Boots (252469, -0.16 DPS) [crafted]; Officer's Sabatons (250561, -0.26 DPS) [crafted]; Officer's Boots (250546, -0.53 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.0 attack_power points (4.47 DPS) | yes | White Bone Band (11862, -1.92 DPS) [quest]; Mark of Kern (2262, -2.34 DPS) [dungeon]; Assault Band (13095, -2.34 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.2 attack_power points (3.31 DPS) | yes | White Bone Band (11862, -0.76 DPS) [quest]; Mark of Kern (2262, -1.19 DPS) [dungeon]; Assault Band (13095, -1.19 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.58 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, -0.31 DPS) [world_drop]; Blight (7959, -1.93 DPS, sim-verified) [crafted]; Bleakwood Hew (12769, -3.44 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Houndmaster's Bow (11628, -0.21 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.28 DPS) [world_drop]; Guttbuster (13139, -0.43 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; wrist: Officer's Wristguards; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 963, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (orc, 13323213032215201-050500000000000000-205000000000000000)

Set DPS (verified): 290.2. Weights run: 5.9s. Verify run: 91.0s. 1824 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.695 ± 0.102, crit=2.422 ± 0.145 per rating point (14 rating = 1%, 33.909 per %), hit=3.229 ± 0.340 per rating point (10 rating = 1%, 32.292 per %), melee_haste=16.596 ± 3.178

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 168.4 attack_power points (15.93 DPS) | yes | Champion's Plate Helm (227043, -5.70 DPS) [pvp]; Mask of the Unforgiven (13404, -6.58 DPS, sim-verified) [dungeon]; Blood Guard's Plate Helm (220803, -7.21 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.0 attack_power points (3.03 DPS) | yes | Woven Ivy Necklace (19159, -0.45 DPS) [quest]; Scout's Medallion (19534, -0.63 DPS) [rep]; Sunprism Pendant (20649, -1.11 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 79.7 attack_power points (7.54 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -1.11 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.28 DPS) [rep]; Fel Cape (279269, -1.09 DPS) [crafted]; Cloak of Flames (3475, -1.24 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Savage Gladiator Chain (11726, -0.65 DPS) [dungeon]; Dawn Armor (252483, -2.72 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -3.86 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 54.2 attack_power points (5.13 DPS) | yes | Berserker Bracers (19578, -0.25 DPS) [rep]; Vambraces of the Sadist (13400, -0.60 DPS) [dungeon]; General's Plate Armguards (16546, -0.79 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor]; Raider Gauntlets (272095, -1.06 DPS, sim-verified) [vendor] |
| waist | Defiler's Chain Girdle (20150) (or Defiler's Leather Girdle (20190), Defiler's Plate Girdle (20204)) | The Defilers [rep] | 67.9 attack_power points (6.42 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; Defiler's Plate Girdle (20204, +0.00 DPS) [rep] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Outrider's Chain Leggings (22673, -1.75 DPS, sim-verified) [rep]; Legionnaire's Plate Leggings (227048, -2.87 DPS) [pvp] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 51.6 attack_power points (4.88 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Defiler's Lamellar Greaves (20181, -0.30 DPS) [rep]; Defiler's Plate Greaves (20208, -0.30 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 80.2 attack_power points (7.59 DPS) | yes | Legionnaire's Band (19510, -3.55 DPS) [rep]; Band of the Penitent (13217, -4.38 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 52.3 attack_power points (4.95 DPS) | yes | Legionnaire's Band (19511, -1.61 DPS) [rep]; Band of the Penitent (13217, -1.74 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Arcanite Champion (12790, -1.68 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.90 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.29 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.64 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; waist: Defiler's Chain Girdle; legs: Outrider's Plate Legguards; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: The Unstoppable Force; ranged: Blackcrow

No-known-source sample (15 of 1824, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (orc, 34323213032215201-050500200000000000-200000000000000000)

Set DPS (verified): 784.8. Weights run: 4.7s. Verify run: 81.6s. 1824 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.494 ± 0.096, crit=1.986 ± 0.131 per rating point (14 rating = 1%, 27.808 per %), hit=not significant (0.000 ± 0.000) per rating point (10 rating = 1%, 0.000 per %), melee_haste=15.458 ± 2.584

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 95.2 attack_power points (21.42 DPS) | yes | Warlord's Plate Headpiece (231535, -1.31 DPS) [pvp]; Champion's Plate Helm (227043, -4.77 DPS) [pvp]; Outlaw's Collar (279253, -28.21 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.4 attack_power points (6.85 DPS) | yes | Woven Ivy Necklace (19159, -0.85 DPS) [quest]; Scout's Medallion (19534, -1.80 DPS) [rep]; Frostwolf Legionnaire's Pendant (19095, -2.80 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 68.2 attack_power points (15.34 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -0.66 DPS) [pvp]; Defiler's Plate Spaulders (20212, -0.71 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.21 DPS) [rep]; Fel Cape (279269, -3.29 DPS) [crafted]; Frostwolf Legionnaire's Cloak (19083, -4.14 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (784.8 DPS) | yes | Savage Gladiator Chain (11726, -5.45 DPS, sim-verified) [dungeon]; Warlord's Plate Armor (231530, -9.76 DPS) [pvp]; Dawn Armor (252483, -9.95 DPS) [crafted] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 53.8 attack_power points (12.10 DPS) | yes | Forest Stalker's Bracers (19587, -0.26 DPS) [rep]; General's Plate Armguards (16546, -1.33 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | sim-verified (784.8 DPS) | yes | General's Plate Gauntlets (231532, -0.91 DPS) [pvp]; Timbermaw Brawlers (19049, -1.96 DPS) [crafted]; Raider Gloves (272099, -2.54 DPS) [vendor] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 67.1 attack_power points (15.10 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Defiler's Plate Girdle (20204, -0.43 DPS) [rep]; Defiler's Chain Girdle (20150, -1.20 DPS) [rep] |
| legs | Outrider's Plate Legguards (22651) | Warsong Outriders [rep] | sim-verified (784.8 DPS) | yes | General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -0.20 DPS) [crafted]; Cloudkeeper Legplates (14554, -15.36 DPS, sim-verified) [world_drop] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.0 attack_power points (11.24 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Defiler's Lamellar Greaves (20181, -0.27 DPS) [rep]; Defiler's Plate Greaves (20208, -0.27 DPS) [rep] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 42.8 attack_power points (9.64 DPS) | yes | Band of the Penitent (13217, -3.38 DPS) [quest]; Don Rodrigo's Band (21563, -3.38 DPS) [rep]; White Bone Band (11862, -4.24 DPS) [quest] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 41.8 attack_power points (9.41 DPS) | yes | Don Rodrigo's Band (21563, -3.15 DPS) [rep]; White Bone Band (11862, -4.01 DPS) [quest]; Band of the Penitent (13217, -19.91 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (784.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (784.8 DPS) | yes | Molten Heart of the Mountain (249470, -12.44 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (784.8 DPS) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -31.19 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Bloodseeker (19107) | The Legend of Korrak [quest] | sim-verified (784.8 DPS) | yes | Precisely Calibrated Boomstick (2100, -1.61 DPS) [world_drop]; Skull Splitting Crossbow (13039, -1.68 DPS) [world_drop]; Dark Iron Rifle (16004, -5.58 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Might of the Timbermaw; legs: Outrider's Plate Legguards; feet: Scalegut Treaders; finger1: Legionnaire's Band; finger2: Don Julio's Band; main_hand: Arcanite Champion; ranged: Bloodseeker

No-known-source sample (15 of 1824, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

