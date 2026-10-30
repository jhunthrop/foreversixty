# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.6. Weights run: 1.6s. Verify run: 2.6s. 287 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.716 ± 0.028, crit=0.311 ± 0.019 per rating point (14 rating = 1%, 4.355 per %), hit=1.230 ± 0.073 per rating point (10 rating = 1%, 12.301 per %), melee_haste=8.936 ± 0.660

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.51 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.3 attack_power points (0.15 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.09 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.07 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.11 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.6 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.02 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Veteran's Chain Belt (250513, -0.28 DPS) [crafted]; Cobrahn's Grasp (6460, -1.22 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.6 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Chausses of Westfall (6087, -1.04 DPS, sim-verified) [quest] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (26.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -0.86 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.9 attack_power points (0.39 DPS) | yes | The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.31 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.18 DPS, sim-verified) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 300.8 attack_power points (10.68 DPS) | yes | Duskbringer (2205, -0.23 DPS) [dungeon]; Night Reaver (1318, -0.44 DPS) [dungeon]; Living Root (6631, -0.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567), Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 287, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 76.7. Weights run: 1.7s. Verify run: 1.7s. 516 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.074 ± 0.044, crit=0.537 ± 0.034 per rating point (14 rating = 1%, 7.520 per %), hit=1.737 ± 0.168 per rating point (10 rating = 1%, 17.368 per %), melee_haste=11.460 ± 1.304

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.7 attack_power points (1.10 DPS) | yes | Veteran's Chain Helm (250498, -0.15 DPS) [crafted]; Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Tusken Helm (6686, -0.36 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.3 attack_power points (0.49 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.01 DPS) [world]; Sentinel's Medallion (19541, -0.15 DPS) [rep]; River Pride Choker (13087, -0.17 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.8 attack_power points (0.86 DPS) | yes | Barbaric Shoulders (5964, -0.26 DPS) [crafted]; Golden Scale Shoulders (3841, -0.31 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.36 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.5 attack_power points (0.54 DPS) | yes | Sergeant Major's Cape (16315, -0.05 DPS) [pvp]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.14 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.14 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.4 attack_power points (1.05 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Girdle of Golem Strength (9405, -0.10 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.10 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.5 attack_power points (1.17 DPS) | yes | Refractory Scaleguards (273048, -0.22 DPS) [world]; Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Ferine Leggings (6690, -0.43 DPS, sim-verified) [dungeon] |
| feet | Alacritous Treads (277234) (or Brawler's Leather Boots (252439)) | WANTED: Incinerator Gar'im [quest] | 15.4 attack_power points (0.61 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.2 attack_power points (0.76 DPS) | yes | Silverlaine's Family Seal (6321, -0.37 DPS) [dungeon]; Tiger Band (6749, -0.37 DPS) [quest]; Bonebinder's Signet (273032, -0.37 DPS) [world] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.4 attack_power points (0.73 DPS) | yes | Silverlaine's Family Seal (6321, -0.33 DPS) [dungeon]; Tiger Band (6749, -0.33 DPS) [quest]; Bonebinder's Signet (273032, -0.33 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (76.7 DPS) | yes | Corpsemaker (6687, -0.32 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.82 DPS) [vendor]; Viscous Hammer (13045, -20.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.2 attack_power points (0.37 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; Gravespike Repeater (273033, -0.05 DPS) [world]; Golemsight Long Gun (273029, -0.11 DPS) [dungeon] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Alacritous Treads; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 101.0. Weights run: 1.9s. Verify run: 1.9s. 716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.202 ± 0.041, crit=1.262 ± 0.048 per rating point (14 rating = 1%, 17.665 per %), hit=1.698 ± 0.160 per rating point (10 rating = 1%, 16.983 per %), melee_haste=12.717 ± 1.403

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 38.4 attack_power points (2.16 DPS) | yes | White Bandit Mask (10008, -0.18 DPS) [crafted]; Barbaric Iron Helm (7915, -0.54 DPS) [crafted]; Hard Gold Coif (250537, -0.58 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 13.2 attack_power points (0.74 DPS) | yes | Kaleidoscope Chain (13084, -0.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; River Pride Choker (13087, -0.29 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.2 attack_power points (1.41 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.18 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.2 attack_power points (1.08 DPS) | yes | Hawkeye's Cloak (14593, -0.27 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.34 DPS) [quest]; Scorpashi Cape (14656, -0.52 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.4 attack_power points (2.04 DPS) | yes | Avenger's Armor (1488, -0.36 DPS) [dungeon]; Jouster's Chestplate (8157, -0.36 DPS) [dungeon]; Carapace of Tuten'kash (10775, -0.38 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Ravager's Armguards (14770, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.25 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.7 attack_power points (2.11 DPS) | yes | Scarlet Gauntlets (10331, -0.09 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.32 DPS) [world_drop]; Truesilver Gauntlets (7938, -0.32 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) (or Highlander's Plate Girdle (20125)) | The League of Arathor [rep] | 30.0 attack_power points (1.68 DPS) | yes | Highlander's Plate Girdle (20125, +0.00 DPS) [rep]; Ogron's Sash (13117, -0.07 DPS) [world_drop]; Officer's Belt (250556, -0.16 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Symbolic Legplates (14829, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.4 attack_power points (1.93 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Blackforge Greaves (6423, -0.25 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 25.6 attack_power points (1.44 DPS) | yes | Assault Band (13095, -0.31 DPS) [world_drop]; Thunderbrow Ring (13097, -0.34 DPS) [world_drop]; Falcon's Hook (7552, -0.49 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (101.0 DPS) | yes | Bonebiter (6830, -1.35 DPS) [quest]; Primitive Fishing Pole (276203, -1.56 DPS) [vendor]; Fiery War Axe (870, -2.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (101.0 DPS) | yes | The Silencer (13138, -0.09 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Bow of Searing Arrows (2825, -1.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 175.0. Weights run: 2.1s. Verify run: 2.5s. 901 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.229 ± 0.046, crit=1.423 ± 0.060 per rating point (14 rating = 1%, 19.929 per %), hit=2.615 ± 0.166 per rating point (10 rating = 1%, 26.150 per %), melee_haste=12.311 ± 1.223

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 72.1 attack_power points (5.99 DPS) | yes | Sunscale Helmet (14849, -2.72 DPS) [world_drop]; Chromite Barbute (8142, -2.77 DPS) [dungeon]; Ornate Mithril Helm (7937, -3.35 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (2.16 DPS) | yes | Sentinel's Medallion (19540, -1.04 DPS) [rep]; Kaleidoscope Chain (13084, -1.09 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 41.9 attack_power points (3.49 DPS) | yes | Officer's Pauldrons (250576, -0.51 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.67 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.74 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.4 attack_power points (1.61 DPS) | yes | Dark Phantom Cape (13122, -0.08 DPS) [world_drop]; Bloodlust Cape (14801, -0.11 DPS) [world_drop]; Pridelord Cape (14673, -0.19 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 49.9 attack_power points (4.15 DPS) | yes | Warforged Chestplate (11195, -0.16 DPS) [quest]; Warbear Harness (15064, -0.48 DPS) [crafted]; Valorous Chestguard (8274, -0.83 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.4 attack_power points (2.44 DPS) | yes | Bracers of the Stone Princess (17714, -0.11 DPS) [dungeon]; Arena Bands (18711, -0.11 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 56.3 attack_power points (4.68 DPS) | yes | Raider Gloves (272100, -1.00 DPS, sim-verified) [vendor]; Officer's Gloves (250551, -1.10 DPS) [crafted]; Stonerender Gauntlets (17007, -1.35 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.7 attack_power points (3.55 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.07 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 56.2 attack_power points (4.67 DPS) | yes | Gryphon Rider's Leggings (9652, -0.65 DPS) [quest]; Centurion Legplates (10740, -0.65 DPS) [quest]; Serpentskin Leggings (8262, -0.88 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.5 attack_power points (3.29 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.41 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.1 attack_power points (2.58 DPS) | yes | Mark of Kern (2262, -0.92 DPS) [dungeon]; Assault Band (13095, -0.92 DPS) [world_drop]; Thunderbrow Ring (13097, -0.95 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.66 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (175.0 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (175.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (175.0 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Nightblade (1982, -5.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (175.0 DPS) | yes | Houndmaster's Bow (11628, -0.17 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.23 DPS) [world_drop]; Dark Iron Rifle (16004, -1.91 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Plate Pauldrons; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 901, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 260.4. Weights run: 2.1s. Verify run: 7.2s. 1630 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.29 DPS) [vendor]; Stalwart Helm (250599, -6.60 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.51 DPS) [quest]; Amulet of the Darkmoon (19491, -0.98 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 73.3 attack_power points (6.22 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -0.69 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.1 attack_power points (5.53 DPS) | yes | Cape of the Black Baron (13340, -1.85 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.98 DPS) [rep]; Stalwart Cloak (272415, -2.38 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+4.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Obsidian Mail Tunic (22191, -1.53 DPS) [crafted]; Breastplate of Heroism (226862, -2.19 DPS) [quest]; Dawn Armor (252483, -4.88 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 51.6 attack_power points (4.38 DPS) | yes | Berserker Bracers (19578, -0.10 DPS) [rep]; Marshal's Plate Bracers (16481, -0.57 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 71.2 attack_power points (6.04 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gauntlets of Heroism (226861, -0.34 DPS) [quest]; Stormshroud Gloves (21278, -4.18 DPS, sim-verified) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.57 DPS) [vendor]; Marshal's Plate Girdle (16482, -1.40 DPS) [pvp]; Belt of Preserved Heads (20216, -3.57 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 165.5 attack_power points (14.04 DPS) | yes | Sentinel's Chain Leggings (237819, -0.97 DPS) [vendor]; Sentinel's Plate Legguards (237825, -1.02 DPS) [vendor]; Marshal's Plate Legguards (231540, -2.20 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 77.1 attack_power points (6.55 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -1.94 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 85.1 attack_power points (7.23 DPS) | yes | Cutthroat's Signet (272408, -2.22 DPS) [vendor]; Protector's Band (19514, -3.73 DPS) [rep]; Channeler's Ring (272406, -4.07 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 82.3 attack_power points (6.99 DPS) | yes | Cutthroat's Signet (272408, -2.11 DPS, sim-verified) [vendor]; Protector's Band (19514, -3.49 DPS) [rep]; Channeler's Ring (272406, -3.83 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.36 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.79 DPS) [world_drop]; Dark Iron Rifle (16004, -2.51 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Heroism; main_hand: The Unstoppable Force; ranged: The Purifier

No-known-source sample (15 of 1630, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 60, raid preset (human, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 788.4. Weights run: 1.8s. Verify run: 6.8s. 1630 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.863 ± 0.094, crit=2.475 ± 0.126 per rating point (14 rating = 1%, 34.645 per %), hit=4.251 ± 0.385 per rating point (10 rating = 1%, 42.508 per %), melee_haste=15.560 ± 2.924

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 193.9 attack_power points (33.42 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -12.16 DPS) [vendor]; Stalwart Helm (250599, -15.16 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.36 DPS) [quest]; Amulet of the Darkmoon (19491, -1.57 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 84.9 attack_power points (14.64 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -2.22 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.5 attack_power points (12.15 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -4.69 DPS) [rep]; Stalwart Cloak (272415, -4.83 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Breastplate of Heroism (226862, -4.80 DPS) [quest]; Obsidian Mail Tunic (22191, -5.19 DPS) [crafted]; Dawn Armor (252483, -9.25 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.6 attack_power points (10.27 DPS) | yes | Berserker Bracers (19578, -0.50 DPS) [rep]; Marshal's Plate Bracers (16481, -1.58 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 79.9 attack_power points (13.77 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.22 DPS) [pvp]; Raider Gloves (272099, -0.93 DPS) [vendor]; Stormshroud Gloves (21278, -11.57 DPS, sim-verified) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.89 DPS) [vendor]; Marshal's Plate Girdle (16482, -2.88 DPS) [pvp]; Belt of Preserved Heads (20216, -7.97 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 185.7 attack_power points (32.00 DPS) | yes | Sentinel's Chain Leggings (237819, +0.00 DPS, sim-verified) [vendor]; Sentinel's Plate Legguards (237825, -2.49 DPS) [vendor]; Marshal's Plate Legguards (231540, -5.15 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 86.5 attack_power points (14.91 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -4.05 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 95.3 attack_power points (16.43 DPS) | yes | Cutthroat's Signet (272408, -4.60 DPS) [vendor]; Protector's Band (19514, -8.34 DPS) [rep]; Channeler's Ring (272406, -9.10 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 91.2 attack_power points (15.71 DPS) | yes | Cutthroat's Signet (272408, -4.30 DPS, sim-verified) [vendor]; Protector's Band (19514, -7.63 DPS) [rep]; Channeler's Ring (272406, -8.38 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+19.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -41.78 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.69 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.48 DPS) [world_drop]; Dark Iron Rifle (16004, -6.77 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Arcanite Champion; ranged: The Purifier

No-known-source sample (15 of 1630, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 27.1. Weights run: 1.6s. Verify run: 2.5s. 257 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.716 ± 0.028, crit=0.311 ± 0.019 per rating point (14 rating = 1%, 4.355 per %), hit=1.230 ± 0.073 per rating point (10 rating = 1%, 12.301 per %), melee_haste=8.936 ± 0.660

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.51 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.3 attack_power points (0.15 DPS) | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.09 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.07 DPS) [crafted]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.14 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.1 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.20 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Veteran's Chain Belt (250513, -0.28 DPS) [crafted]; Cobrahn's Grasp (6460, -1.38 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.1 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.16 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (27.1 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -1.01 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.9 attack_power points (0.39 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; The 1 Ring (8350, -0.29 DPS) [world]; Bounty Hunter's Ring (5351, -0.31 DPS) [quest] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; The 1 Ring (8350, -0.19 DPS) [world]; Bounty Hunter's Ring (5351, -0.21 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.87 DPS) | yes | Forsaken Greataxe (251533, -0.11 DPS) [quest]; Smite's Mighty Hammer (7230, -0.19 DPS) [dungeon]; Living Root (6631, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop]; Heavy Shortbow (3036, -0.07 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 257, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 209616 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 77.2. Weights run: 1.7s. Verify run: 1.7s. 478 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.074 ± 0.044, crit=0.537 ± 0.034 per rating point (14 rating = 1%, 7.520 per %), hit=1.737 ± 0.168 per rating point (10 rating = 1%, 17.368 per %), melee_haste=11.460 ± 1.304

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.7 attack_power points (1.10 DPS) | yes | Veteran's Chain Helm (250498, -0.15 DPS) [crafted]; Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Tusken Helm (6686, -0.47 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.3 attack_power points (0.49 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.01 DPS) [world]; Scout's Medallion (19537, -0.15 DPS) [rep]; River Pride Choker (13087, -0.17 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.8 attack_power points (0.86 DPS) | yes | Barbaric Shoulders (5964, -0.26 DPS) [crafted]; Golden Scale Shoulders (3841, -0.31 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.42 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.5 attack_power points (0.54 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest]; Construct Cloak (279848, -0.53 DPS, sim-verified) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.14 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Warsong Gauntlets (16978, -0.08 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.4 attack_power points (1.05 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Girdle of Golem Strength (9405, -0.10 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.10 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.5 attack_power points (1.17 DPS) | yes | Refractory Scaleguards (273048, -0.22 DPS) [world]; Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Ferine Leggings (6690, -0.53 DPS, sim-verified) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.4 attack_power points (0.61 DPS) | yes | Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Hard Gold Boots (250534, -0.05 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.2 attack_power points (0.76 DPS) | yes | Band of the Fist (17694, -0.27 DPS) [quest]; Tiger Band (6749, -0.37 DPS) [quest]; Bonebinder's Signet (273032, -0.37 DPS) [world] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.4 attack_power points (0.73 DPS) | yes | Tiger Band (6749, -0.33 DPS) [quest]; Bonebinder's Signet (273032, -0.33 DPS) [world]; Band of the Fist (17694, -0.49 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.2 DPS) | yes | Corpsemaker (6687, -0.32 DPS) [dungeon]; Darkspear Raider's Reaper (272082, -0.82 DPS) [vendor]; Viscous Hammer (13045, -21.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.2 attack_power points (0.37 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.01 DPS) [vendor]; Gravespike Repeater (273033, -0.05 DPS) [world]; Golemsight Long Gun (273029, -0.11 DPS) [dungeon] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 478, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 101.2. Weights run: 1.9s. Verify run: 1.9s. 667 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.202 ± 0.041, crit=1.262 ± 0.048 per rating point (14 rating = 1%, 17.665 per %), hit=1.698 ± 0.160 per rating point (10 rating = 1%, 16.983 per %), melee_haste=12.717 ± 1.403

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 38.4 attack_power points (2.16 DPS) | yes | White Bandit Mask (10008, -0.18 DPS) [crafted]; Barbaric Iron Helm (7915, -0.54 DPS) [crafted]; Hard Gold Coif (250537, -0.58 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 14.8 attack_power points (0.83 DPS) | yes | Scout's Medallion (19536, -0.09 DPS) [rep]; Kaleidoscope Chain (13084, -0.11 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 25.2 attack_power points (1.41 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.18 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.4 attack_power points (0.81 DPS) | yes | Scorpashi Cape (14656, -0.25 DPS) [world_drop]; Wildhunter Cloak (16658, -0.25 DPS) [quest]; Construct Cloak (279848, -0.25 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.4 attack_power points (2.04 DPS) | yes | Avenger's Armor (1488, -0.36 DPS) [dungeon]; Jouster's Chestplate (8157, -0.36 DPS) [dungeon]; Carapace of Tuten'kash (10775, -0.38 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.7 attack_power points (2.11 DPS) | yes | Scarlet Gauntlets (10331, -0.09 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.32 DPS) [world_drop]; Truesilver Gauntlets (7938, -0.32 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) (or Defiler's Plate Girdle (20206)) | The Defilers [rep] | 30.0 attack_power points (1.68 DPS) | yes | Defiler's Plate Girdle (20206, +0.00 DPS) [rep]; Ogron's Sash (13117, -0.07 DPS) [world_drop]; Officer's Belt (250556, -0.16 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Symbolic Legplates (14829, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.4 attack_power points (1.93 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Blackforge Greaves (6423, -0.25 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 25.6 attack_power points (1.44 DPS) | yes | Assault Band (13095, -0.31 DPS) [world_drop]; Thunderbrow Ring (13097, -0.34 DPS) [world_drop]; Falcon's Hook (7552, -0.49 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (101.2 DPS) | yes | Primitive Fishing Pole (276203, -1.56 DPS) [vendor]; Darkspear Raider's Reaper (272081, -1.72 DPS) [vendor]; Fiery War Axe (870, -2.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (101.2 DPS) | yes | The Silencer (13138, -0.09 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Bow of Searing Arrows (2825, -1.01 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; shoulder: Sunburn Spaulders; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 667, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 174.8. Weights run: 2.1s. Verify run: 2.4s. 843 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.229 ± 0.046, crit=1.423 ± 0.060 per rating point (14 rating = 1%, 19.929 per %), hit=2.615 ± 0.166 per rating point (10 rating = 1%, 26.150 per %), melee_haste=12.311 ± 1.223

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 39.9 attack_power points (3.32 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Sunscale Helmet (14849, -0.05 DPS) [world_drop]; Chromite Barbute (8142, -0.10 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (2.16 DPS) | yes | Woven Ivy Necklace (19159, -0.24 DPS) [quest]; Ethereal Talisman (4430, -0.92 DPS) [quest]; Scout's Medallion (19535, -0.93 DPS) [rep] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.8 attack_power points (2.98 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.17 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.4 attack_power points (1.53 DPS) | yes | Bloodlust Cape (14801, -0.04 DPS) [world_drop]; Pridelord Cape (14673, -0.11 DPS) [world_drop]; Blackflame Cape (13109, -0.31 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 49.9 attack_power points (4.15 DPS) | yes | Warforged Chestplate (11195, -0.16 DPS) [quest]; Warbear Harness (15064, -0.48 DPS) [crafted]; Valorous Chestguard (8274, -0.83 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.4 attack_power points (2.44 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 56.3 attack_power points (4.68 DPS) | yes | Raider Gloves (272100, -0.77 DPS) [vendor]; Officer's Gloves (250551, -1.10 DPS) [crafted]; Stonerender Gauntlets (17007, -1.35 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.7 attack_power points (3.55 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.23 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 56.2 attack_power points (4.67 DPS) | yes | Serpentskin Leggings (8262, -0.88 DPS) [world_drop]; Golem Shard Leggings (13074, -1.02 DPS) [world_drop]; Scarlet Leggings (10330, -1.18 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.5 attack_power points (3.29 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.41 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.1 attack_power points (2.58 DPS) | yes | Mark of Kern (2262, -0.92 DPS) [dungeon]; Assault Band (13095, -0.92 DPS) [world_drop]; Blackstone Ring (17713, -0.92 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.00 DPS) | yes | Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop]; Blackstone Ring (17713, -0.33 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (174.8 DPS) | yes | Frozen Heart of the Mountain (249469, -3.06 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (174.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (174.8 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Nightblade (1982, -3.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (174.8 DPS) | yes | Houndmaster's Bow (11628, -0.17 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.23 DPS) [world_drop]; Dark Iron Rifle (16004, -1.88 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 843, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 265.7. Weights run: 2.1s. Verify run: 7.0s. 1599 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Champion's Plate Helm (227043, -5.29 DPS) [pvp]; Stalwart Helm (250599, -6.35 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -6.65 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.51 DPS) [quest]; Amulet of the Darkmoon (19491, -0.98 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 73.3 attack_power points (6.22 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -0.69 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 65.1 attack_power points (5.53 DPS) | yes | Cape of the Black Baron (13340, -1.85 DPS) [dungeon]; Deathguard's Cloak (20068, -1.98 DPS) [rep]; Stalwart Cloak (272415, -2.38 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Obsidian Mail Tunic (22191, -1.53 DPS) [crafted]; Breastplate of Heroism (226862, -2.19 DPS) [quest]; Dawn Armor (252483, -4.21 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 51.6 attack_power points (4.38 DPS) | yes | Berserker Bracers (19578, -0.10 DPS) [rep]; General's Plate Armguards (16546, -0.57 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 71.2 attack_power points (6.04 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gauntlets of Heroism (226861, -0.34 DPS) [quest]; Stormshroud Gloves (21278, -4.59 DPS, sim-verified) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.57 DPS) [vendor]; General's Plate Girdle (16547, -1.40 DPS) [pvp]; Belt of Preserved Heads (20216, -3.16 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 165.5 attack_power points (14.04 DPS) | yes | Sentinel's Chain Leggings (237819, -0.97 DPS) [vendor]; Sentinel's Plate Legguards (237825, -1.02 DPS) [vendor]; General's Plate Leggings (231533, -2.20 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 77.1 attack_power points (6.55 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -1.94 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 85.1 attack_power points (7.23 DPS) | yes | Cutthroat's Signet (272408, -2.22 DPS) [vendor]; Legionnaire's Band (19510, -3.73 DPS) [rep]; Channeler's Ring (272406, -4.07 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 82.3 attack_power points (6.99 DPS) | yes | Cutthroat's Signet (272408, -2.56 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -3.49 DPS) [rep]; Channeler's Ring (272406, -3.83 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.93 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.28 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.36 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.79 DPS) [world_drop]; Dark Iron Rifle (16004, -2.81 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: The Purifier

No-known-source sample (15 of 1599, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (orc, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 794.6. Weights run: 1.8s. Verify run: 8.4s. 1599 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.863 ± 0.094, crit=2.475 ± 0.126 per rating point (14 rating = 1%, 34.645 per %), hit=4.251 ± 0.385 per rating point (10 rating = 1%, 42.508 per %), melee_haste=15.560 ± 2.924

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 193.9 attack_power points (33.42 DPS) | yes | Champion's Plate Helm (227043, -12.16 DPS) [pvp]; Blood Guard's Plate Helm (220803, -15.19 DPS) [vendor]; Stalwart Helm (250599, -17.35 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (794.6 DPS) | yes | Medallion of the Dawn (22659, -1.36 DPS) [quest]; Amulet of the Darkmoon (19491, -1.57 DPS) [quest]; Rage of Mugamba (19577, -6.56 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | sim-verified (794.6 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -5.28 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.5 attack_power points (12.15 DPS) | yes | Deathguard's Cloak (20068, -4.69 DPS) [rep]; Stalwart Cloak (272415, -4.83 DPS) [vendor]; Cape of the Black Baron (13340, -5.84 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (794.6 DPS) | yes | Breastplate of Heroism (226862, -4.80 DPS) [quest]; Obsidian Mail Tunic (22191, -5.19 DPS) [crafted]; Dawn Armor (252483, -9.66 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.6 attack_power points (10.27 DPS) | yes | Berserker Bracers (19578, -0.50 DPS) [rep]; General's Plate Armguards (16546, -1.58 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 79.9 attack_power points (13.77 DPS) | yes | General's Plate Gauntlets (231532, -0.22 DPS) [pvp]; Raider Gloves (272099, -0.93 DPS) [vendor]; Stormshroud Gloves (21278, -10.54 DPS, sim-verified) [crafted] |
| waist | Defiler's Plate Girdle (20204) | The Defilers [rep] | sim-verified (794.6 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -9.10 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 185.7 attack_power points (32.00 DPS) | yes | Sentinel's Plate Legguards (237825, -2.49 DPS) [vendor]; General's Plate Leggings (231533, -5.15 DPS) [pvp]; Sentinel's Chain Leggings (237819, -5.19 DPS, sim-verified) [vendor] |
| feet | Defiler's Plate Greaves (20208) | The Defilers [rep] | sim-verified (794.6 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Boots of Heroism (21995, -8.20 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 95.3 attack_power points (16.43 DPS) | yes | Cutthroat's Signet (272408, -4.60 DPS) [vendor]; Legionnaire's Band (19510, -8.34 DPS) [rep]; Channeler's Ring (272406, -9.10 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 91.2 attack_power points (15.71 DPS) | yes | Cutthroat's Signet (272408, -5.49 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -7.63 DPS) [rep]; Channeler's Ring (272406, -8.38 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (794.6 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (794.6 DPS) | yes | Frozen Heart of the Mountain (249469, -5.77 DPS) [crafted]; Molten Heart of the Mountain (249470, -17.17 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (794.6 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -37.14 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (794.6 DPS) | yes | Bloodseeker (19107, -0.69 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.48 DPS) [world_drop]; Dark Iron Rifle (16004, -7.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Defiler's Plate Spaulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Defiler's Plate Girdle; legs: Titanic Leggings; feet: Defiler's Plate Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Arcanite Champion; ranged: The Purifier

No-known-source sample (15 of 1599, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

