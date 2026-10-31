# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.1. Weights run: 2.0s. Verify run: 3.1s. 293 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.716 ± 0.028, crit=0.311 ± 0.019 per rating point (14 rating = 1%, 4.355 per %), hit=1.230 ± 0.073 per rating point (10 rating = 1%, 12.301 per %), melee_haste=8.936 ± 0.660

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.51 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 1.4 attack_power points (0.05 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.09 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.07 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.11 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.1 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.07 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Veteran's Chain Belt (250513, -0.28 DPS) [crafted]; Cobrahn's Grasp (6460, -1.23 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.1 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Chausses of Westfall (6087, -1.05 DPS, sim-verified) [quest] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (26.1 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -0.87 DPS, sim-verified) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Ring of the Moon (12052, -0.21 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.23 DPS) [dungeon] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 2.7 attack_power points (0.10 DPS) | yes | Ring of the Moon (12052, -0.03 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.05 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 300.8 attack_power points (10.68 DPS) | yes | Duskbringer (2205, -0.23 DPS) [dungeon]; Night Reaver (1318, -0.44 DPS) [dungeon]; Living Root (6631, -0.67 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Dwarven Fishing Pole (3567), Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, +0.00 DPS) [quest]; Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: The 1 Ring; main_hand: Smite's Mighty Hammer; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 76.0. Weights run: 2.1s. Verify run: 2.0s. 551 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.074 ± 0.044, crit=0.537 ± 0.034 per rating point (14 rating = 1%, 7.520 per %), hit=1.737 ± 0.168 per rating point (10 rating = 1%, 17.368 per %), melee_haste=11.460 ± 1.304

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.7 attack_power points (1.10 DPS) | yes | Veteran's Chain Helm (250498, -0.15 DPS) [crafted]; Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Tusken Helm (6686, -0.44 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.3 attack_power points (0.49 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.01 DPS) [world]; River Pride Choker (13087, -0.17 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.25 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.8 attack_power points (0.86 DPS) | yes | Barbaric Shoulders (5964, -0.26 DPS) [crafted]; Golden Scale Shoulders (3841, -0.31 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.42 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.5 attack_power points (0.54 DPS) | yes | Sergeant Major's Cape (16315, -0.05 DPS) [pvp]; Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.14 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.14 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.4 attack_power points (1.05 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Girdle of Golem Strength (9405, -0.10 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.19 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.5 attack_power points (1.17 DPS) | yes | Refractory Scaleguards (273048, -0.22 DPS) [world]; Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Ferine Leggings (6690, -0.52 DPS, sim-verified) [dungeon] |
| feet | Alacritous Treads (277234) (or Brawler's Leather Boots (252439)) | WANTED: Incinerator Gar'im [quest] | 15.4 attack_power points (0.61 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.2 attack_power points (0.76 DPS) | yes | Silverlaine's Family Seal (6321, -0.37 DPS) [dungeon]; Bonebinder's Signet (273032, -0.37 DPS) [world]; Seal of Wrynn (2933, -0.40 DPS) [quest] |
| finger2 | Tiger Band (6749) (or Bonebinder's Signet (273032), Silverlaine's Family Seal (6321)) | Willix the Importer [quest] | 10.0 attack_power points (0.40 DPS) | yes | Silverlaine's Family Seal (6321, +0.00 DPS) [dungeon]; Bonebinder's Signet (273032, +0.00 DPS) [world]; Seal of Wrynn (2933, -0.03 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (76.0 DPS) | yes | Corpsemaker (6687, -0.32 DPS) [dungeon]; Felbough Arm (273037, -1.15 DPS) [world]; Viscous Hammer (13045, -20.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.2 attack_power points (0.37 DPS) | yes | Gravespike Repeater (273033, -0.05 DPS) [world]; Golemsight Long Gun (273029, -0.11 DPS) [dungeon]; Long Battle Bow (15284, -0.13 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Alacritous Treads; finger1: Thunderbrow Ring; finger2: Tiger Band; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 551, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 100.2. Weights run: 2.5s. Verify run: 2.2s. 781 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.202 ± 0.041, crit=1.262 ± 0.048 per rating point (14 rating = 1%, 17.665 per %), hit=1.698 ± 0.160 per rating point (10 rating = 1%, 16.983 per %), melee_haste=12.717 ± 1.403

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 38.4 attack_power points (2.16 DPS) | yes | White Bandit Mask (10008, -0.18 DPS) [crafted]; Barbaric Iron Helm (7915, -0.54 DPS) [crafted]; Hard Gold Coif (250537, -0.58 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.8 attack_power points (0.72 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.27 DPS) [world_drop]; Gazlowe's Charm (13088, -0.27 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.2 attack_power points (1.30 DPS) | yes | Flintrock Shoulders (7755, -0.07 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.07 DPS) [crafted]; Chromite Pauldrons (8144, -0.18 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.2 attack_power points (1.08 DPS) | yes | Hawkeye's Cloak (14593, -0.27 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.34 DPS) [quest]; Scorpashi Cape (14656, -0.52 DPS) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.4 attack_power points (2.04 DPS) | yes | Avenger's Armor (1488, -0.36 DPS) [dungeon]; Jouster's Chestplate (8157, -0.36 DPS) [dungeon]; Carapace of Tuten'kash (10775, -0.38 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Ravager's Armguards (14770, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.25 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.7 attack_power points (2.11 DPS) | yes | Scarlet Gauntlets (10331, -0.09 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.32 DPS) [world_drop]; Truesilver Gauntlets (7938, -0.32 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 28.8 attack_power points (1.62 DPS) | yes | Officer's Belt (250556, -0.09 DPS) [crafted]; Prowler's Leather Belt (252459, -0.20 DPS) [crafted]; Scarlet Belt (10329, -0.27 DPS) [dungeon] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Symbolic Legplates (14829, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.4 attack_power points (1.93 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Blackforge Greaves (6423, -0.25 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.40 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon]; Ring of the Underwood (2951, -0.22 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.12 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon]; Ring of the Underwood (2951, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Bonebiter (6830, -1.35 DPS) [quest]; Primitive Fishing Pole (276203, -1.56 DPS) [vendor]; Fiery War Axe (870, -2.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (100.2 DPS) | yes | The Silencer (13138, -0.09 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Bow of Searing Arrows (2825, -1.06 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 781, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 169.2. Weights run: 2.6s. Verify run: 3.0s. 1004 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.229 ± 0.046, crit=1.423 ± 0.060 per rating point (14 rating = 1%, 19.929 per %), hit=2.615 ± 0.166 per rating point (10 rating = 1%, 26.150 per %), melee_haste=12.311 ± 1.223

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 39.9 attack_power points (3.32 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Sunscale Helmet (14849, -0.05 DPS) [world_drop]; Chromite Barbute (8142, -0.10 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (2.16 DPS) | yes | Kaleidoscope Chain (13084, -1.10 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -1.16 DPS) [world]; River Pride Choker (13087, -1.49 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.8 attack_power points (2.98 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.17 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 19.4 attack_power points (1.61 DPS) | yes | Dark Phantom Cape (13122, -0.08 DPS) [world_drop]; Bloodlust Cape (14801, -0.11 DPS) [world_drop]; Pridelord Cape (14673, -0.19 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 49.9 attack_power points (4.15 DPS) | yes | Warforged Chestplate (11195, -0.16 DPS) [quest]; Warbear Harness (15064, -0.48 DPS) [crafted]; Valorous Chestguard (8274, -0.83 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.4 attack_power points (2.44 DPS) | yes | Bracers of the Stone Princess (17714, -0.11 DPS) [dungeon]; Arena Bands (18711, -0.11 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 43.1 attack_power points (3.58 DPS) | yes | Stonerender Gauntlets (17007, -0.25 DPS) [world_drop]; Gloves of Holy Might (867, -0.26 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.33 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.7 attack_power points (3.55 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Belt of the Gladiator (13134, -0.56 DPS) [world_drop] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 56.2 attack_power points (4.67 DPS) | yes | Gryphon Rider's Leggings (9652, -0.65 DPS) [quest]; Centurion Legplates (10740, -0.65 DPS) [quest]; Serpentskin Leggings (8262, -0.88 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.5 attack_power points (3.29 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.41 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.66 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.23 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.66 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.23 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (169.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (169.2 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (169.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Nightblade (1982, -4.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (169.2 DPS) | yes | Houndmaster's Bow (11628, -0.17 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.23 DPS) [world_drop]; Dark Iron Rifle (16004, -1.86 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 1004, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 230.0. Weights run: 2.6s. Verify run: 9.1s. 2015 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.29 DPS) [vendor]; Stalwart Helm (250599, -5.37 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.3 attack_power points (2.57 DPS) | yes | Sunprism Pendant (20649, -1.10 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -1.36 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.55 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 59.2 attack_power points (5.03 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 43.4 attack_power points (3.68 DPS) | yes | Sergeant Major's Cape (16337, -1.27 DPS) [pvp]; Fel Cape (279269, -1.60 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.70 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (230.0 DPS) | yes | Dawn Armor (252483, -2.88 DPS, sim-verified) [crafted]; Field Marshal's Plate Armor (231536, -3.41 DPS) [pvp]; Knight-Captain's Plate Hauberk (227046, -4.42 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 32.0 attack_power points (2.72 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.05 DPS) [world_drop]; Officer's Wristguards (250581, -0.05 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.68 DPS) [world_drop]; Timbermaw Brawlers (19049, -5.17 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 63.8 attack_power points (5.42 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -1.40 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.75 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -5.22 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.5 attack_power points (4.28 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Shadowcraft Boots (16711, -1.73 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.8 attack_power points (1.85 DPS) | yes | Falcon's Hook (7552, -0.15 DPS) [dungeon]; Blackstone Ring (17713, -0.16 DPS) [dungeon]; Ring of Living Stone (18400, -0.16 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 20.7 attack_power points (1.75 DPS) | yes | Falcon's Hook (7552, -0.05 DPS) [dungeon]; Blackstone Ring (17713, -0.06 DPS) [dungeon]; Ring of Living Stone (18400, -0.06 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.16 DPS) [world_drop]; Guttbuster (13139, -0.28 DPS) [world_drop]; Dark Iron Rifle (16004, -1.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Forest Defender's Axe; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2015, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (human, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 703.5. Weights run: 2.3s. Verify run: 8.3s. 2015 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.863 ± 0.094, crit=2.475 ± 0.126 per rating point (14 rating = 1%, 34.645 per %), hit=4.251 ± 0.385 per rating point (10 rating = 1%, 42.508 per %), melee_haste=15.560 ± 2.924

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 193.9 attack_power points (33.42 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -12.16 DPS) [vendor]; Stalwart Helm (250599, -13.88 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (6.07 DPS) | yes | Sunprism Pendant (20649, -2.22 DPS) [quest]; Kaleidoscope Chain (13084, -3.27 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -4.00 DPS) [world] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 67.4 attack_power points (11.61 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Plate Shoulders (227045, +0.00 DPS) [pvp]; Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.9 attack_power points (8.26 DPS) | yes | Sergeant Major's Cape (16337, -2.66 DPS) [pvp]; Dark Phantom Cape (13122, -3.45 DPS) [world_drop]; Fel Cape (279269, -3.63 DPS, sim-verified) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 146.1 attack_power points (25.18 DPS) | yes | Timbermaw Tunic (252484, -0.91 DPS) [crafted]; Field Marshal's Plate Armor (231536, -8.64 DPS) [pvp]; Shadowcraft Tunic (16721, -10.97 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.7 attack_power points (6.33 DPS) | yes | Marshal's Plate Bracers (16481, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.24 DPS) [crafted]; Tranquil Wristguards (279256, -0.27 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (703.5 DPS) | yes | Marshal's Plate Gauntlets (231541, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -1.38 DPS) [world_drop]; Timbermaw Brawlers (19049, -17.64 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 72.3 attack_power points (12.46 DPS) | yes | Marshal's Plate Girdle (16482, +0.00 DPS) [pvp]; Knight-Captain's Plate Girdle (16407, -3.33 DPS) [pvp]; Prowler's Leather Waistguard (252473, -4.17 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (703.5 DPS) | yes | Knight-Captain's Plate Leggings (23301, +0.00 DPS) [vendor]; Marshal's Plate Legguards (231540, +0.00 DPS) [pvp]; Warbear Woolies (15065, -17.74 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 52.9 attack_power points (9.12 DPS) | yes | Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -0.18 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.66 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.1 attack_power points (4.50 DPS) | yes | Thunderbrow Ring (13097, -0.50 DPS) [world_drop]; Ring of the Underwood (2951, -0.53 DPS) [world_drop]; Ring of Living Stone (18400, -0.70 DPS) [quest] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.4 attack_power points (4.03 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Ring of Living Stone (18400, -0.24 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (703.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (703.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (703.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -21.82 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (703.5 DPS) | yes | Guttbuster (13139, -0.79 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.95 DPS) [world_drop]; Dark Iron Rifle (16004, -4.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Deepfury Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: Falcon's Hook; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 2015, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.7. Weights run: 2.0s. Verify run: 3.0s. 276 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=2.000 ± 0.003, agility=0.716 ± 0.028, crit=0.311 ± 0.019 per rating point (14 rating = 1%, 4.355 per %), hit=1.230 ± 0.073 per rating point (10 rating = 1%, 12.301 per %), melee_haste=8.936 ± 0.660

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.51 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 1.4 attack_power points (0.05 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.09 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.07 DPS) [crafted]; Subterranean Cape (14149, -0.07 DPS) [dungeon] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.14 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.7 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.17 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Brawler's Leather Belt (252428, -0.25 DPS) [crafted]; Veteran's Chain Belt (250513, -0.28 DPS) [crafted]; Cobrahn's Grasp (6460, -1.35 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.7 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.14 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (26.7 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -0.99 DPS, sim-verified) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.19 DPS) [world]; Bounty Hunter's Ring (5351, -0.21 DPS) [quest]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.21 DPS) | yes | The 1 Ring (8350, -0.12 DPS) [world]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Ring of the Moon (12052, -0.14 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.87 DPS) | yes | Forsaken Greataxe (251533, -0.11 DPS) [quest]; Smite's Mighty Hammer (7230, -0.19 DPS) [dungeon]; Living Root (6631, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cracked Blacksmith Hammer (285279) (or Fine Longbow (11304)) | Blacksmithing [crafted] | 4.0 attack_power points (0.14 DPS) | yes | Fine Longbow (11304, +0.00 DPS) [vendor]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop]; Heavy Shortbow (3036, -0.07 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone; ranged: Cracked Blacksmith Hammer

No-known-source sample (15 of 276, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 76.7. Weights run: 2.1s. Verify run: 2.1s. 534 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.074 ± 0.044, crit=0.537 ± 0.034 per rating point (14 rating = 1%, 7.520 per %), hit=1.737 ± 0.168 per rating point (10 rating = 1%, 17.368 per %), melee_haste=11.460 ± 1.304

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.7 attack_power points (1.10 DPS) | yes | Veteran's Chain Helm (250498, -0.15 DPS) [crafted]; Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Tusken Helm (6686, -0.46 DPS, sim-verified) [dungeon] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 12.3 attack_power points (0.49 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.01 DPS) [world]; River Pride Choker (13087, -0.17 DPS) [world_drop]; Erudite's Amulet (277204, -0.40 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.8 attack_power points (0.86 DPS) | yes | Barbaric Shoulders (5964, -0.26 DPS) [crafted]; Golden Scale Shoulders (3841, -0.31 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.44 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.5 attack_power points (0.54 DPS) | yes | Wolfmaster Cape (6314, -0.14 DPS) [dungeon]; Wildhunter Cloak (16658, -0.14 DPS) [quest]; Construct Cloak (279848, -0.51 DPS, sim-verified) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.14 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.63 DPS) | yes | Yorgen Bracers (13012, -0.03 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.06 DPS) [world_drop]; Barbaric Bracers (18948, -0.15 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.06 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Warsong Gauntlets (16978, -0.08 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.4 attack_power points (1.05 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Girdle of Golem Strength (9405, -0.10 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.19 DPS) [crafted] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.5 attack_power points (1.17 DPS) | yes | Refractory Scaleguards (273048, -0.22 DPS) [world]; Brawler's Leather Legguards (252516, -0.23 DPS) [crafted]; Ferine Leggings (6690, -0.51 DPS, sim-verified) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 15.4 attack_power points (0.61 DPS) | yes | Feet of the Lynx (1121, -0.03 DPS) [world_drop]; Veteran's Boots (250503, -0.04 DPS) [crafted]; Hard Gold Boots (250534, -0.05 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.2 attack_power points (0.76 DPS) | yes | Silverlaine's Family Seal (6321, -0.37 DPS) [dungeon]; Tiger Band (6749, -0.37 DPS) [quest]; Bonebinder's Signet (273032, -0.37 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 12.3 attack_power points (0.49 DPS) | yes | Silverlaine's Family Seal (6321, -0.09 DPS) [dungeon]; Tiger Band (6749, -0.09 DPS) [quest]; Bonebinder's Signet (273032, -0.09 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (76.7 DPS) | yes | Corpsemaker (6687, -0.32 DPS) [dungeon]; Felbough Arm (273037, -1.15 DPS) [world]; Viscous Hammer (13045, -21.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Double-barreled Shotgun (2098) | World drop [world_drop] | 9.2 attack_power points (0.37 DPS) | yes | Gravespike Repeater (273033, -0.05 DPS) [world]; Golemsight Long Gun (273029, -0.11 DPS) [dungeon]; Long Battle Bow (15284, -0.13 DPS) [world_drop] |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler; ranged: Double-barreled Shotgun

No-known-source sample (15 of 534, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 100.7. Weights run: 2.5s. Verify run: 2.3s. 761 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.005, agility=1.202 ± 0.041, crit=1.262 ± 0.048 per rating point (14 rating = 1%, 17.665 per %), hit=1.698 ± 0.160 per rating point (10 rating = 1%, 16.983 per %), melee_haste=12.717 ± 1.403

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 38.4 attack_power points (2.16 DPS) | yes | White Bandit Mask (10008, -0.18 DPS) [crafted]; Barbaric Iron Helm (7915, -0.54 DPS) [crafted]; Hard Gold Coif (250537, -0.58 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 14.8 attack_power points (0.83 DPS) | yes | Kaleidoscope Chain (13084, -0.11 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.16 DPS) [world]; River Pride Choker (13087, -0.38 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.2 attack_power points (1.30 DPS) | yes | Flintrock Shoulders (7755, -0.07 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.07 DPS) [crafted]; Chromite Pauldrons (8144, -0.18 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.4 attack_power points (0.81 DPS) | yes | Scorpashi Cape (14656, -0.25 DPS) [world_drop]; Wildhunter Cloak (16658, -0.25 DPS) [quest]; Construct Cloak (279848, -0.25 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.4 attack_power points (2.04 DPS) | yes | Avenger's Armor (1488, -0.36 DPS) [dungeon]; Jouster's Chestplate (8157, -0.36 DPS) [dungeon]; Carapace of Tuten'kash (10775, -0.38 DPS) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Ravager's Armguards (14770, -0.07 DPS) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.25 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 37.7 attack_power points (2.11 DPS) | yes | Scarlet Gauntlets (10331, -0.09 DPS) [dungeon]; Plated Fist of Hakoo (13071, -0.32 DPS) [world_drop]; Truesilver Gauntlets (7938, -0.32 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 28.8 attack_power points (1.62 DPS) | yes | Officer's Belt (250556, -0.09 DPS) [crafted]; Tharg's Shoelace (9705, -0.16 DPS) [quest]; Prowler's Leather Belt (252459, -0.20 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Symbolic Legplates (14829, -0.27 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.4 attack_power points (1.93 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Blackforge Greaves (6423, -0.25 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.40 DPS) [crafted] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon]; Ring of the Underwood (2951, -0.22 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.12 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon]; Ring of the Underwood (2951, -0.22 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (100.7 DPS) | yes | Primitive Fishing Pole (276203, -1.56 DPS) [vendor]; Staff of Jordan (873, -2.34 DPS) [world_drop]; Fiery War Axe (870, -2.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman [dungeon] | sim-verified (100.7 DPS) | yes | The Silencer (13138, -0.09 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Bow of Searing Arrows (2825, -0.98 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Mark of Kern; finger2: Assault Band; main_hand: Nightblade; ranged: Monolithic Bow

No-known-source sample (15 of 761, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 172.9. Weights run: 2.6s. Verify run: 2.9s. 983 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.229 ± 0.046, crit=1.423 ± 0.060 per rating point (14 rating = 1%, 19.929 per %), hit=2.615 ± 0.166 per rating point (10 rating = 1%, 26.150 per %), melee_haste=12.311 ± 1.223

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 39.9 attack_power points (3.32 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Sunscale Helmet (14849, -0.05 DPS) [world_drop]; Chromite Barbute (8142, -0.10 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.0 attack_power points (2.16 DPS) | yes | Woven Ivy Necklace (19159, -0.24 DPS) [quest]; Ethereal Talisman (4430, -0.92 DPS) [quest]; Kaleidoscope Chain (13084, -1.09 DPS) [world_drop] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 35.8 attack_power points (2.98 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.17 DPS) [world_drop]; Prowler's Leather Shoulders (252534, -0.23 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.4 attack_power points (1.53 DPS) | yes | Bloodlust Cape (14801, -0.04 DPS) [world_drop]; Pridelord Cape (14673, -0.11 DPS) [world_drop]; Blackflame Cape (13109, -0.31 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | PvP rank 10 · Stone Guard · Horde [vendor] | 49.9 attack_power points (4.15 DPS) | yes | Warforged Chestplate (11195, -0.16 DPS) [quest]; Warbear Harness (15064, -0.48 DPS) [crafted]; Valorous Chestguard (8274, -0.83 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.4 attack_power points (2.44 DPS) | yes | Bracers of the Stone Princess (17714, -0.11 DPS) [dungeon]; Arena Bands (18711, -0.11 DPS) [world]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted] |
| hands | Officer's Gloves (250551) | Blacksmithing [crafted] | 43.1 attack_power points (3.58 DPS) | yes | Stonerender Gauntlets (17007, -0.25 DPS) [world_drop]; Gloves of Holy Might (867, -0.26 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.33 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 42.7 attack_power points (3.55 DPS) | yes | Skulker's Leather Waistguard (252474, -0.13 DPS) [crafted]; Officer's Waistguard (250571, -0.20 DPS) [crafted]; Belt of the Gladiator (13134, -0.56 DPS) [world_drop] |
| legs | Stone Guard's Plate Leggings (220798) | PvP rank 10 · Stone Guard · Horde [vendor] | 56.2 attack_power points (4.67 DPS) | yes | Serpentskin Leggings (8262, -0.88 DPS) [world_drop]; Golem Shard Leggings (13074, -1.02 DPS) [world_drop]; Scarlet Leggings (10330, -1.18 DPS) [dungeon] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 39.5 attack_power points (3.29 DPS) | yes | Skulker's Leather Boots (252469, -0.13 DPS) [crafted]; Officer's Sabatons (250561, -0.20 DPS) [crafted]; Officer's Boots (250546, -0.41 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.00 DPS) | yes | Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.66 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (172.9 DPS) | yes | Frozen Heart of the Mountain (249469, -3.06 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (172.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-verified (172.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Nightblade (1982, -3.66 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (172.9 DPS) | yes | Houndmaster's Bow (11628, -0.17 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.23 DPS) [world_drop]; Dark Iron Rifle (16004, -1.87 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; back: Dark Phantom Cape; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Officer's Gloves; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Blight; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 983, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 234.1. Weights run: 2.6s. Verify run: 8.9s. 1998 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.559 ± 0.063, crit=2.228 ± 0.091 per rating point (14 rating = 1%, 31.186 per %), hit=3.715 ± 0.250 per rating point (10 rating = 1%, 37.149 per %), melee_haste=17.614 ± 2.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 172.7 attack_power points (14.65 DPS) | yes | Champion's Plate Helm (227043, -5.29 DPS) [pvp]; Stalwart Helm (250599, -5.87 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -6.65 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.3 attack_power points (2.57 DPS) | yes | Woven Ivy Necklace (19159, -0.36 DPS) [quest]; Sunprism Pendant (20649, -0.98 DPS) [quest]; Ethereal Talisman (4430, -1.19 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 59.2 attack_power points (5.03 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 43.4 attack_power points (3.68 DPS) | yes | Fel Cape (279269, -1.20 DPS, sim-verified) [crafted]; Dark Phantom Cape (13122, -1.70 DPS) [world_drop]; Pridelord Cape (14673, -1.98 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 126.0 attack_power points (10.69 DPS) | yes | Timbermaw Tunic (252484, +0.00 DPS, sim-verified) [crafted]; Warlord's Plate Armor (231530, -3.48 DPS) [pvp]; Legionnaire's Plate Hauberk (227049, -4.48 DPS) [pvp] |
| wrist | Tranquil Wristguards (279256) | Leatherworking [crafted] | 32.0 attack_power points (2.72 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -0.05 DPS) [world_drop]; Officer's Wristguards (250581, -0.05 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (234.1 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -0.68 DPS) [world_drop]; Timbermaw Brawlers (19049, -4.06 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 63.8 attack_power points (5.42 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -1.40 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.62 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (234.1 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.02 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 50.5 attack_power points (4.28 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.62 DPS) [crafted]; Shadowcraft Boots (16711, -1.32 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.04 DPS) | yes | Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon]; Ring of Living Stone (18400, -0.34 DPS) [quest] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.8 attack_power points (1.85 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.15 DPS) [dungeon]; Ring of Living Stone (18400, -0.16 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (234.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (234.1 DPS) | yes | Molten Heart of the Mountain (249470, -2.15 DPS, sim-verified) [crafted] |
| main_hand | Forest Defender's Axe (279260) | Blacksmithing [crafted] | sim-verified (234.1 DPS) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (234.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.16 DPS) [world_drop]; Guttbuster (13139, -0.28 DPS) [world_drop]; Dark Iron Rifle (16004, -1.48 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Tranquil Wristguards; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger2: Masons Fraternity Ring; trinket2: Frozen Heart of the Mountain; main_hand: Forest Defender's Axe; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1998, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (orc, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 713.7. Weights run: 2.3s. Verify run: 8.3s. 1998 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.863 ± 0.094, crit=2.475 ± 0.126 per rating point (14 rating = 1%, 34.645 per %), hit=4.251 ± 0.385 per rating point (10 rating = 1%, 42.508 per %), melee_haste=15.560 ± 2.924

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 193.9 attack_power points (33.42 DPS) | yes | Champion's Plate Helm (227043, -12.16 DPS) [pvp]; Stalwart Helm (250599, -13.42 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -15.19 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (6.07 DPS) | yes | Woven Ivy Necklace (19159, -0.91 DPS) [quest]; Sunprism Pendant (20649, -2.22 DPS) [quest]; Ethereal Talisman (4430, -2.89 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 67.4 attack_power points (11.61 DPS) | yes | Champion's Plate Shoulders (227042, +0.00 DPS) [pvp]; Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -6.57 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.9 attack_power points (8.26 DPS) | yes | Dark Phantom Cape (13122, -3.45 DPS) [world_drop]; Fel Cape (279269, -3.77 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -4.24 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 146.1 attack_power points (25.18 DPS) | yes | Timbermaw Tunic (252484, -0.91 DPS) [crafted]; Warlord's Plate Armor (231530, -8.64 DPS) [pvp]; Shadowcraft Tunic (16721, -10.97 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.7 attack_power points (6.33 DPS) | yes | General's Plate Armguards (16546, +0.00 DPS) [pvp]; Officer's Wristguards (250581, -0.24 DPS) [crafted]; Tranquil Wristguards (279256, -0.27 DPS) [crafted] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (713.7 DPS) | yes | General's Plate Gauntlets (231532, +0.00 DPS) [pvp]; Gloves of Holy Might (867, -1.38 DPS) [world_drop]; Timbermaw Brawlers (19049, -16.83 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 72.3 attack_power points (12.46 DPS) | yes | General's Plate Girdle (16547, +0.00 DPS) [pvp]; Legionnaire's Plate Cinch (16511, -3.33 DPS) [pvp]; Prowler's Leather Waistguard (252473, -4.43 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (713.7 DPS) | yes | Legionnaire's Plate Leggings (227048, +0.00 DPS) [pvp]; General's Plate Leggings (231533, +0.00 DPS) [pvp]; Warbear Woolies (15065, -16.23 DPS, sim-verified) [crafted] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | 52.9 attack_power points (9.12 DPS) | yes | General's Plate Boots (231531, +0.00 DPS) [pvp]; Prowler's Leather Boots (252468, -0.66 DPS) [crafted]; Shadowcraft Boots (16711, -4.68 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.1 attack_power points (4.50 DPS) | yes | Falcon's Hook (7552, -0.47 DPS) [dungeon]; Thunderbrow Ring (13097, -0.50 DPS) [world_drop]; Ring of the Underwood (2951, -0.53 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (4.14 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Ring of the Underwood (2951, -0.17 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (713.7 DPS) | yes | Frozen Heart of the Mountain (249469, -9.29 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (713.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (713.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Forest Defender's Axe (279260, -26.75 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (713.7 DPS) | yes | Guttbuster (13139, -0.79 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.95 DPS) [world_drop]; Dark Iron Rifle (16004, -4.27 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Deepfury Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Scalegut Treaders; finger1: Masons Fraternity Ring; finger2: White Bone Band; main_hand: Arcanite Champion; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1998, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

