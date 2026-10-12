# Leveling BiS: Arms

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 26.6. Weights run: 1.2s. Verify run: 3.6s. 284 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.698 ± 0.037, crit=0.296 ± 0.026 per rating point (14 rating = 1%, 4.149 per %), hit=1.154 ± 0.105 per rating point (10 rating = 1%, 11.538 per %), melee_haste=8.649 ± 0.944

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.51 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 4.2 attack_power points (0.15 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.09 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.07 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Cryptwalker Bracers (280095, -0.07 DPS) [quest]; Bravo's Armbands (270015, -0.11 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.6 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.02 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -0.26 DPS) [crafted]; Cobrahn's Grasp (6460, -1.22 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (26.6 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Chausses of Westfall (6087, -1.04 DPS, sim-verified) [quest] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (26.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -0.86 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.8 attack_power points (0.38 DPS) | yes | Signet of the Zhevra (285330, -0.23 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world]; Ring of the Moon (12052, -0.31 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Signet of the Zhevra (285330, -0.14 DPS) [world]; The 1 Ring (8350, -0.19 DPS) [world]; Ring of the Moon (12052, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 300.7 attack_power points (10.68 DPS) | yes | Duskbringer (2205, -0.22 DPS) [dungeon]; Night Reaver (1318, -0.43 DPS) [dungeon]; Living Root (6631, -0.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Dwarven Fishing Pole (3567, -0.00 DPS) [quest]; Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer; ranged: Fine Longbow

No-known-source sample (15 of 284, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (human, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 77.3. Weights run: 1.3s. Verify run: 2.2s. 499 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.001 ± 0.062, crit=0.518 ± 0.047 per rating point (14 rating = 1%, 7.259 per %), hit=1.498 ± 0.244 per rating point (10 rating = 1%, 14.978 per %), melee_haste=9.585 ± 1.907

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.0 attack_power points (1.07 DPS) | yes | Veteran's Chain Helm (250498, -0.12 DPS) [crafted]; Defender's Leather Helm (252455, -0.12 DPS) [crafted]; Tusken Helm (6686, -0.46 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.56 DPS) | yes | Kaleidoscope Chain (13084, -0.08 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.08 DPS) [quest]; Sentinel's Medallion (19541, -0.24 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (0.84 DPS) | yes | Barbaric Shoulders (5964, -0.24 DPS) [crafted]; Golden Scale Shoulders (3841, -0.28 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.43 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.0 attack_power points (0.52 DPS) | yes | Sergeant Major's Cape (16315, -0.04 DPS) [pvp]; Wolfmaster Cape (6314, -0.12 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.20 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.16 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.64 DPS) | yes | Yorgen Bracers (13012, -0.04 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.08 DPS) [world_drop]; Barbaric Bracers (18948, -0.16 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.08 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.08 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.16 DPS) [crafted] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.0 attack_power points (1.03 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Girdle of Golem Strength (9405, -0.08 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.08 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.0 attack_power points (1.15 DPS) | yes | Brawler's Leather Legguards (252516, -0.24 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.24 DPS) [world_drop]; Ferine Leggings (6690, -0.52 DPS, sim-verified) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 21.0 attack_power points (0.83 DPS) | yes | Brawler's Leather Boots (252439, -0.24 DPS) [crafted]; Feet of the Lynx (1121, -0.28 DPS) [world_drop]; Alacritous Treads (277234, -0.47 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.0 attack_power points (0.76 DPS) | yes | Ironspine's Eye (7686, -0.08 DPS) [dungeon]; Tiger Band (6749, -0.20 DPS) [quest]; Silverlaine's Family Seal (6321, -0.36 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 18.0 attack_power points (0.72 DPS) | yes | Ironspine's Eye (7686, -0.04 DPS) [dungeon]; Tiger Band (6749, -0.16 DPS) [quest]; Silverlaine's Family Seal (6321, -0.32 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.3 DPS) | yes | Morbid Dawn (7689, -0.02 DPS) [dungeon]; Corpsemaker (6687, -0.31 DPS) [dungeon]; Viscous Hammer (13045, -21.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.48 DPS) | yes | Double-barreled Shotgun (2098, -0.12 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.12 DPS) [vendor]; Golemsight Long Gun (273029, -0.24 DPS) [dungeon] |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler; ranged: Glass Shooter

No-known-source sample (15 of 499, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler

### Band 40 (human, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 120.3. Weights run: 1.3s. Verify run: 2.7s. 682 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.007, agility=1.152 ± 0.055, crit=1.194 ± 0.063 per rating point (14 rating = 1%, 16.714 per %), hit=1.723 ± 0.229 per rating point (10 rating = 1%, 17.229 per %), melee_haste=12.704 ± 1.950

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 37.8 attack_power points (2.13 DPS) | yes | White Bandit Mask (10008, -0.18 DPS) [crafted]; Barbaric Iron Helm (7915, -0.53 DPS) [crafted]; Hard Gold Coif (250537, -0.55 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.79 DPS) | yes | Sentinel's Medallion (19540, -0.07 DPS) [rep]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.11 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.7 attack_power points (1.39 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.15 DPS) [crafted]; Flintrock Shoulders (7755, -0.18 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.5 attack_power points (1.10 DPS) | yes | Sergeant Major's Cape (16336, -0.03 DPS) [pvp]; Hawkeye's Cloak (14593, -0.31 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.37 DPS) [quest] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.1 attack_power points (2.03 DPS) | yes | Avenger's Armor (1488, -0.34 DPS) [dungeon]; Carapace of Tuten'kash (10775, -0.39 DPS) [dungeon]; Jouster's Chestplate (8157, -0.62 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Ravager's Armguards (14770, -0.08 DPS) [world_drop]; Pugilist Bracers (4438, -0.22 DPS) [dungeon]; Yorgen Bracers (13012, -0.26 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 36.7 attack_power points (2.07 DPS) | yes | Scarlet Gauntlets (10331, -0.07 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.27 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.27 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.69 DPS) | yes | Highlander's Plate Girdle (20125, -0.00 DPS) [rep]; Ogron's Sash (13117, -0.09 DPS) [world_drop]; Officer's Belt (250556, -0.17 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Symbolic Legplates (14829, -0.29 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.1 attack_power points (1.92 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Blackforge Greaves (6423, -0.26 DPS) [dungeon]; Obsidian Greaves (13068, -0.40 DPS) [world_drop] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 25.2 attack_power points (1.42 DPS) | yes | Assault Band (13095, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Ironspine's Eye (7686, -0.39 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Ironspine's Eye (7686, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (120.3 DPS) | yes | Bonebiter (6830, +0.00 DPS) [quest]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -19.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (120.3 DPS) | yes | The Silencer (13138, -0.08 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Bow of Searing Arrows (2825, -1.54 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Protector's Band; finger2: Mark of Kern; ranged: Monolithic Bow

No-known-source sample (15 of 682, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 50 (human, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 174.5. Weights run: 1.3s. Verify run: 3.5s. 861 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.270 ± 0.067, crit=1.462 ± 0.088 per rating point (14 rating = 1%, 20.469 per %), hit=2.165 ± 0.236 per rating point (10 rating = 1%, 21.652 per %), melee_haste=12.576 ± 1.719

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Plate Helm (220804) | Captain Dirgehammer [vendor] | 68.1 attack_power points (5.68 DPS) | yes | Sunscale Helmet (14849, -2.38 DPS) [world_drop]; Chromite Barbute (8142, -2.41 DPS) [dungeon]; Ornate Mithril Helm (7937, -3.43 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.5 attack_power points (2.21 DPS) | yes | Ghostshard Talisman (7731, -1.04 DPS) [dungeon]; Sentinel's Medallion (19540, -1.05 DPS) [rep] |
| shoulder | Knight-Lieutenant's Plate Pauldrons (220795) | Captain Dirgehammer [vendor] | 42.5 attack_power points (3.54 DPS) | yes | Officer's Pauldrons (250576, -0.53 DPS) [crafted]; Wyrmslayer Spaulders (13066, -0.69 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.75 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 20.7 attack_power points (1.73 DPS) | yes | Sergeant Major's Cape (16336, -0.09 DPS) [pvp]; Dark Phantom Cape (13122, -0.14 DPS) [world_drop]; Bloodlust Cape (14801, -0.23 DPS) [world_drop] |
| chest | Knight's Plate Hauberk (220794) | Captain Dirgehammer [vendor] | 50.5 attack_power points (4.21 DPS) | yes | Warforged Chestplate (11195, -0.21 DPS) [quest]; Warbear Harness (15064, -0.47 DPS) [crafted]; Valorous Chestguard (8274, -0.87 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.6 attack_power points (2.47 DPS) | yes | Bracers of the Stone Princess (17714, -0.14 DPS) [dungeon]; Arena Bands (18711, -0.14 DPS) [world]; Deepfury Bracers (13120, -0.21 DPS) [world_drop] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 56.7 attack_power points (4.73 DPS) | yes | Raider Gloves (272100, -0.73 DPS) [vendor]; Officer's Gloves (250551, -1.11 DPS) [crafted]; Gloves of Holy Might (867, -1.35 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 43.2 attack_power points (3.61 DPS) | yes | Highlander's Lamellar Girdle (20106, -0.06 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted]; Officer's Waistguard (250571, -0.21 DPS) [crafted] |
| legs | Knight's Plate Leggings (220797) | Captain Dirgehammer [vendor] | 57.2 attack_power points (4.77 DPS) | yes | Gryphon Rider's Leggings (9652, -0.68 DPS) [quest]; Centurion Legplates (10740, -0.68 DPS) [quest]; Serpentskin Leggings (8262, -0.90 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 40.0 attack_power points (3.33 DPS) | yes | Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Officer's Sabatons (250561, -0.21 DPS) [crafted]; Officer's Boots (250546, -0.42 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 31.4 attack_power points (2.62 DPS) | yes | Mark of Kern (2262, -0.95 DPS) [dungeon]; Assault Band (13095, -0.95 DPS) [world_drop]; Thunderbrow Ring (13097, -0.97 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.67 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.02 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (174.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (174.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (174.5 DPS) | yes | Blight (7959, +0.00 DPS) [crafted]; Thorium Greatmace (250613, -1.09 DPS) [crafted]; Warmonger (13052, -1.43 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (174.5 DPS) | yes | Houndmaster's Bow (11628, -0.17 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.18 DPS) [world_drop]; Dark Iron Rifle (16004, -1.90 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Plate Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Plate Pauldrons; chest: Knight's Plate Hauberk; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Knight's Plate Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 861, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 60 (human, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 268.5. Weights run: 1.3s. Verify run: 14.1s. 1585 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.440 ± 0.091, crit=2.058 ± 0.130 per rating point (14 rating = 1%, 28.807 per %), hit=3.989 ± 0.345 per rating point (10 rating = 1%, 39.886 per %), melee_haste=16.659 ± 2.849

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 173.4 attack_power points (14.72 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -5.32 DPS) [vendor]; Stalwart Helm (250599, -6.75 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.94 DPS) [quest]; Amulet of the Darkmoon (19491, -1.40 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 67.7 attack_power points (5.75 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -0.42 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.9 attack_power points (5.76 DPS) | yes | Cape of the Black Baron (13340, -2.23 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.27 DPS) [rep]; Stalwart Cloak (272415, -2.38 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Obsidian Mail Tunic (22191, -1.35 DPS) [crafted]; Breastplate of Heroism (226862, -1.71 DPS) [quest]; Dawn Armor (252483, -4.02 DPS, sim-verified) [crafted] |
| wrist | Berserker Bracers (19578) | Silverwing Sentinels [rep] | 49.5 attack_power points (4.21 DPS) | yes | Forest Stalker's Bracers (19587, -0.01 DPS) [rep]; Marshal's Plate Bracers (16481, -0.46 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 69.8 attack_power points (5.93 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.09 DPS) [pvp]; Gauntlets of Heroism (226861, -0.43 DPS) [quest]; Stormshroud Gloves (21278, -4.40 DPS, sim-verified) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.54 DPS) [vendor]; Marshal's Plate Girdle (16482, -1.32 DPS) [pvp]; Belt of Preserved Heads (20216, -3.50 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 168.6 attack_power points (14.31 DPS) | yes | Sentinel's Plate Legguards (237825, +0.00 DPS, sim-verified) [vendor]; Sentinel's Chain Leggings (237819, -1.75 DPS) [vendor]; Marshal's Plate Legguards (231540, -2.64 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 79.9 attack_power points (6.78 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -2.05 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 87.9 attack_power points (7.46 DPS) | yes | Cutthroat's Signet (272408, -2.36 DPS) [vendor]; Channeler's Ring (272406, -4.08 DPS) [vendor]; Blessed Band of Light (272407, -4.08 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 84.7 attack_power points (7.19 DPS) | yes | Cutthroat's Signet (272408, -2.67 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.80 DPS) [vendor]; Blessed Band of Light (272407, -3.80 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+8.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Arcanite Champion (12790, -7.96 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.23 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.73 DPS) [world_drop]; Dark Iron Rifle (16004, -3.03 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Heroism; main_hand: The Unstoppable Force; ranged: The Purifier

No-known-source sample (15 of 1585, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

### Band 60, raid preset (human, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 788.8. Weights run: 1.5s. Verify run: 13.9s. 1585 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.200 ± 0.006, agility=1.769 ± 0.133, crit=2.373 ± 0.179 per rating point (14 rating = 1%, 33.216 per %), hit=4.041 ± 0.527 per rating point (10 rating = 1%, 40.414 per %), melee_haste=18.525 ± 4.164

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 186.9 attack_power points (32.17 DPS) | yes | Lieutenant Commander's Plate Helm (23314, -11.54 DPS) [vendor]; Stalwart Helm (250599, -15.17 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.24 DPS) [quest]; Amulet of the Darkmoon (19491, -1.51 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 81.0 attack_power points (13.94 DPS) | yes | Field Marshal's Plate Shoulderguards (231537, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Lieutenant Commander's Plate Shoulders (227045, -1.79 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 68.4 attack_power points (11.78 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -4.40 DPS) [rep]; Stalwart Cloak (272415, -4.82 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Breastplate of Heroism (226862, -4.83 DPS) [quest]; Obsidian Mail Tunic (22191, -4.90 DPS) [crafted]; Dawn Armor (252483, -9.25 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 57.8 attack_power points (9.95 DPS) | yes | Berserker Bracers (19578, -0.32 DPS) [rep]; Marshal's Plate Bracers (16481, -1.38 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 78.9 attack_power points (13.58 DPS) | yes | Marshal's Plate Gauntlets (231541, -0.28 DPS) [pvp]; Gauntlets of Heroism (226861, -1.04 DPS) [quest]; Stormshroud Gloves (21278, -11.57 DPS, sim-verified) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+8.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.92 DPS) [vendor]; Marshal's Plate Girdle (16482, -2.82 DPS) [pvp]; Belt of Preserved Heads (20216, -7.97 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 180.0 attack_power points (31.00 DPS) | yes | Sentinel's Chain Leggings (237819, +0.00 DPS, sim-verified) [vendor]; Sentinel's Plate Legguards (237825, -2.38 DPS) [vendor]; Marshal's Plate Legguards (231540, -5.03 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 84.4 attack_power points (14.53 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; Marshal's Plate Boots (231539, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -4.22 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 93.2 attack_power points (16.05 DPS) | yes | Cutthroat's Signet (272408, -4.83 DPS) [vendor]; Protector's Band (19514, -8.15 DPS) [rep]; Channeler's Ring (272406, -9.09 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 89.6 attack_power points (15.43 DPS) | yes | Cutthroat's Signet (272408, -4.75 DPS, sim-verified) [vendor]; Protector's Band (19514, -7.54 DPS) [rep]; Channeler's Ring (272406, -8.47 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+19.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmoon Card: Heroism (19287, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -25.36 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.56 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.45 DPS) [world_drop]; Dark Iron Rifle (16004, -6.77 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Arcanite Champion; ranged: The Purifier

No-known-source sample (15 of 1585, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants

## Horde

### Band 20 (orc, 03323000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 27.1. Weights run: 1.2s. Verify run: 3.7s. 254 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=0.698 ± 0.037, crit=0.296 ± 0.026 per rating point (14 rating = 1%, 4.149 per %), hit=1.154 ± 0.105 per rating point (10 rating = 1%, 11.538 per %), melee_haste=8.649 ± 0.944

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.71 DPS) | yes | Defender's Leather Hood (252447, -0.15 DPS, sim-verified) [crafted]; Guard's Silvered Chain Helm (250529, -0.46 DPS) [crafted]; Brawler's Leather Hood (252504, -0.51 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 4.2 attack_power points (0.15 DPS) | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest] |
| shoulder | Rough Bronze Shoulders (3480) (or Silvered Bronze Shoulders (3481)) | Blacksmithing [crafted] | 6.0 attack_power points (0.21 DPS) | yes | Silvered Bronze Shoulders (3481, +0.00 DPS) [crafted]; Serpent's Shoulders (5404, -0.09 DPS) [dungeon] |
| back | Lambent Scale Cloak (4706) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Grave Shroud (279865, -0.02 DPS) [quest]; Dark Leather Cloak (2316, -0.07 DPS) [crafted]; Catacomb Cloak (279899, -0.07 DPS) [quest] |
| chest | Mutant Scale Breastplate (6627) | Wailing Caverns: Mutanus the Devourer [dungeon] | 20.0 attack_power points (0.71 DPS) | yes | Veteran's Chain Shirt (250488, -0.11 DPS) [crafted]; Defender's Leather Armor (252434, -0.14 DPS) [crafted]; Brawler's Leather Armor (252490, -0.18 DPS) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.35 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Raptorcrest Bracers (270010, -0.14 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.1 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Gloves of the Fang (10413, +0.00 DPS) [dungeon]; Thorbia's Gauntlets (12994, -1.20 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.64 DPS) | yes | Ruffian Belt (5975, -0.21 DPS) [world]; Brawler's Leather Belt (252428, -0.26 DPS) [crafted]; Cobrahn's Grasp (6460, -1.38 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (27.1 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Veteran's Chain Leggings (250493, -1.16 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (27.1 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -1.01 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.8 attack_power points (0.38 DPS) | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Signet of the Zhevra (285330, -0.23 DPS) [world]; The 1 Ring (8350, -0.29 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.28 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Signet of the Zhevra (285330, -0.14 DPS) [world]; The 1 Ring (8350, -0.19 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (10.87 DPS) | yes | Forsaken Greataxe (251533, -0.11 DPS) [quest]; Smite's Mighty Hammer (7230, -0.19 DPS) [dungeon]; Living Root (6631, -0.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fine Longbow (11304) | Naela Trance [vendor] | 4.0 attack_power points (0.14 DPS) | yes | Cracked Blacksmith Hammer (285279, -0.00 DPS) [crafted]; Lil Timmy's Peashooter (13136, -0.04 DPS) [world_drop]; Heavy Shortbow (3036, -0.07 DPS) [world_drop] |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Rough Bronze Shoulders; back: Lambent Scale Cloak; chest: Mutant Scale Breastplate; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone; ranged: Fine Longbow

No-known-source sample (15 of 254, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18854 Insignia of the Alliance; 209616 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 03325213020000000-00000000000000000-000000000000000000)

Set DPS (verified): 77.5. Weights run: 1.3s. Verify run: 2.2s. 463 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.006, agility=1.001 ± 0.062, crit=0.518 ± 0.047 per rating point (14 rating = 1%, 7.259 per %), hit=1.498 ± 0.244 per rating point (10 rating = 1%, 14.978 per %), melee_haste=9.585 ± 1.907

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 27.0 attack_power points (1.07 DPS) | yes | Veteran's Chain Helm (250498, -0.12 DPS) [crafted]; Defender's Leather Helm (252455, -0.12 DPS) [crafted]; Tusken Helm (6686, -0.51 DPS, sim-verified) [dungeon] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.56 DPS) | yes | Kaleidoscope Chain (13084, -0.08 DPS) [world_drop]; Scout's Medallion (19537, -0.24 DPS) [rep]; River Pride Choker (13087, -0.24 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (0.84 DPS) | yes | Barbaric Shoulders (5964, -0.24 DPS) [crafted]; Golden Scale Shoulders (3841, -0.28 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.45 DPS, sim-verified) [crafted] |
| back | Construct Cloak (279848) | Source of Power [quest] | 14.0 attack_power points (0.56 DPS) | yes | Hawkeye's Cloak (14593, -0.04 DPS) [world_drop]; Wolfmaster Cape (6314, -0.16 DPS) [dungeon]; Wildhunter Cloak (16658, -0.16 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.19 DPS) | yes | Shining Silver Breastplate (2870, -0.08 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.16 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.24 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.64 DPS) | yes | Yorgen Bracers (13012, -0.04 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.08 DPS) [world_drop]; Undead Knight's Bracers (251965, -0.16 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (0.87 DPS) | yes | Insignia Gloves (6408, -0.08 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.08 DPS) [world_drop]; Warsong Gauntlets (16978, -0.08 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 26.0 attack_power points (1.03 DPS) | yes | Prowler's Leather Belt (252459, -0.08 DPS) [crafted]; Girdle of Golem Strength (9405, -0.08 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.08 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 29.0 attack_power points (1.15 DPS) | yes | Brawler's Leather Legguards (252516, -0.24 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.24 DPS) [world_drop]; Ferine Leggings (6690, -0.53 DPS, sim-verified) [dungeon] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 21.0 attack_power points (0.83 DPS) | yes | Feet of the Lynx (1121, -0.28 DPS) [world_drop]; Veteran's Boots (250503, -0.28 DPS) [crafted]; Brawler's Leather Boots (252439, -0.50 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.0 attack_power points (0.76 DPS) | yes | Ironspine's Eye (7686, -0.08 DPS) [dungeon]; Tiger Band (6749, -0.20 DPS) [quest]; Band of the Fist (17694, -0.28 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 18.0 attack_power points (0.72 DPS) | yes | Ironspine's Eye (7686, -0.04 DPS) [dungeon]; Tiger Band (6749, -0.16 DPS) [quest]; Band of the Fist (17694, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (77.5 DPS) | yes | Morbid Dawn (7689, -0.02 DPS) [dungeon]; Corpsemaker (6687, -0.31 DPS) [dungeon]; Viscous Hammer (13045, -21.47 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.48 DPS) | yes | Double-barreled Shotgun (2098, -0.12 DPS) [world_drop]; Booty Bay Bruiser's Buckshot (274748, -0.12 DPS) [vendor]; Golemsight Long Gun (273029, -0.24 DPS) [dungeon] |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Construct Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler; ranged: Glass Shooter

No-known-source sample (15 of 463, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads

### Band 40 (orc, 03325213032511000-00000000000000000-000000000000000000)

Set DPS (verified): 120.7. Weights run: 1.3s. Verify run: 2.7s. 635 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.007, agility=1.152 ± 0.055, crit=1.194 ± 0.063 per rating point (14 rating = 1%, 16.714 per %), hit=1.723 ± 0.229 per rating point (10 rating = 1%, 17.229 per %), melee_haste=12.704 ± 1.950

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 37.8 attack_power points (2.13 DPS) | yes | White Bandit Mask (10008, -0.18 DPS) [crafted]; Barbaric Iron Helm (7915, -0.53 DPS) [crafted]; Hard Gold Coif (250537, -0.55 DPS) [crafted] |
| neck | Ethereal Talisman (4430) | The Crown of Will [quest] | 14.6 attack_power points (0.82 DPS) | yes | Ghostshard Talisman (7731, -0.03 DPS) [dungeon]; Scout's Medallion (19536, -0.11 DPS) [rep]; Kaleidoscope Chain (13084, -0.11 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.7 attack_power points (1.39 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Hard Gold Pauldrons (250539, -0.15 DPS) [crafted]; Flintrock Shoulders (7755, -0.18 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.5 attack_power points (1.10 DPS) | yes | Construct Cloak (279848, -0.31 DPS) [quest]; Wildhunter Cloak (16658, -0.54 DPS) [quest]; Hawkeye's Cloak (14593, -0.57 DPS, sim-verified) [world_drop] |
| chest | Kolkar Marauder Chain (6773) | Khan Hratha [quest] | 36.1 attack_power points (2.03 DPS) | yes | Avenger's Armor (1488, -0.34 DPS) [dungeon]; Carapace of Tuten'kash (10775, -0.39 DPS) [dungeon]; Jouster's Chestplate (8157, -0.65 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 36.7 attack_power points (2.07 DPS) | yes | Scarlet Gauntlets (10331, -0.07 DPS) [dungeon]; Gauntlets of Divinity (7724, -0.27 DPS) [dungeon]; Truesilver Gauntlets (7938, -0.27 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.69 DPS) | yes | Defiler's Plate Girdle (20206, -0.00 DPS) [rep]; Ogron's Sash (13117, -0.09 DPS) [world_drop]; Officer's Belt (250556, -0.17 DPS) [crafted] |
| legs | Scarlet Leggings (10330) | Scarlet Monastery: Herod [dungeon] | 42.0 attack_power points (2.36 DPS) | yes | Firemane Leggings (13129, -0.22 DPS) [world_drop]; Symbolic Legplates (14829, -0.29 DPS) [world_drop]; Orcish War Leggings (7929, -0.45 DPS) [crafted] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 34.1 attack_power points (1.92 DPS) | yes | Prowler's Leather Shoes (252465, -0.22 DPS) [crafted]; Blackforge Greaves (6423, -0.26 DPS) [dungeon]; Obsidian Greaves (13068, -0.40 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 25.2 attack_power points (1.42 DPS) | yes | Assault Band (13095, -0.29 DPS) [world_drop]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Ironspine's Eye (7686, -0.39 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.12 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Ironspine's Eye (7686, -0.09 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (120.7 DPS) | yes | Darkspear Raider's Reaper (272081, +0.00 DPS) [vendor]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -20.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Monolithic Bow (9426) | Uldaman: Shadowforge Relic Hunter [dungeon] | sim-verified (120.7 DPS) | yes | The Silencer (13138, -0.08 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.16 DPS) [quest]; Bow of Searing Arrows (2825, -1.49 DPS, sim-verified) [world_drop] |

**New at 40:** head: Chromite Barbute; neck: Ethereal Talisman; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Kolkar Marauder Chain; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Scarlet Leggings; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Mark of Kern; ranged: Monolithic Bow

No-known-source sample (15 of 635, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 50 (orc, 03325213032515001-05000000000000000-000000000000000000)

Set DPS (verified): 175.0. Weights run: 1.3s. Verify run: 3.6s. 805 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.270 ± 0.067, crit=1.462 ± 0.088 per rating point (14 rating = 1%, 20.469 per %), hit=2.165 ± 0.236 per rating point (10 rating = 1%, 21.652 per %), melee_haste=12.576 ± 1.719

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 40.5 attack_power points (3.37 DPS) | yes | Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Sunscale Helmet (14849, -0.07 DPS) [world_drop]; Chromite Barbute (8142, -0.10 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.5 attack_power points (2.21 DPS) | yes | Woven Ivy Necklace (19159, -0.26 DPS) [quest]; Scout's Medallion (19535, -0.94 DPS) [rep]; Ethereal Talisman (4430, -0.95 DPS) [quest] |
| shoulder | Officer's Pauldrons (250576) | Blacksmithing [crafted] | 36.2 attack_power points (3.02 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Wyrmslayer Spaulders (13066, -0.17 DPS) [world_drop]; Prowler's Leather Shoulder (252534, -0.23 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 20.7 attack_power points (1.73 DPS) | yes | Dark Phantom Cape (13122, -0.14 DPS) [world_drop]; Bloodlust Cape (14801, -0.23 DPS) [world_drop]; Pridelord Cape (14673, -0.27 DPS) [world_drop] |
| chest | Stone Guard's Plate Armor (220801) | Lady Palanseer [vendor] | 50.5 attack_power points (4.21 DPS) | yes | Warforged Chestplate (11195, -0.21 DPS) [quest]; Warbear Harness (15064, -0.47 DPS) [crafted]; Valorous Chestguard (8274, -0.87 DPS) [world_drop] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 29.6 attack_power points (2.47 DPS) | yes | Arena Bands (18711, +0.00 DPS) [world]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Raider Gauntlets (272096) | Creeg Bothunk [vendor] | 56.7 attack_power points (4.73 DPS) | yes | Raider Gloves (272100, -0.73 DPS) [vendor]; Officer's Gloves (250551, -1.11 DPS) [crafted]; Gloves of Holy Might (867, -1.35 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 43.2 attack_power points (3.61 DPS) | yes | Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted]; Officer's Waistguard (250571, -0.21 DPS) [crafted]; Defiler's Plate Girdle (20205, -0.23 DPS) [rep] |
| legs | Stone Guard's Plate Leggings (220798) | Lady Palanseer [vendor] | 57.2 attack_power points (4.77 DPS) | yes | Golem Shard Leggings (13074, -1.10 DPS) [world_drop]; Scarlet Leggings (10330, -1.26 DPS) [dungeon]; Serpentskin Leggings (8262, -1.43 DPS, sim-verified) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 40.0 attack_power points (3.33 DPS) | yes | Skulker's Leather Boots (252469, -0.12 DPS) [crafted]; Officer's Sabatons (250561, -0.21 DPS) [crafted]; Officer's Boots (250546, -0.42 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 31.4 attack_power points (2.62 DPS) | yes | Mark of Kern (2262, -0.95 DPS) [dungeon]; Assault Band (13095, -0.95 DPS) [world_drop]; Blackstone Ring (17713, -0.95 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.00 DPS) | yes | Mark of Kern (2262, -0.33 DPS) [dungeon]; Assault Band (13095, -0.33 DPS) [world_drop]; Blackstone Ring (17713, -0.33 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (175.0 DPS) | yes | Frozen Heart of the Mountain (249469, -3.14 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (175.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (175.0 DPS) | yes | Blight (7959, +0.00 DPS) [crafted]; Thorium Greatmace (250613, -1.09 DPS) [crafted]; Warmonger (13052, -1.43 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (175.0 DPS) | yes | Houndmaster's Bow (11628, -0.17 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.18 DPS) [world_drop]; Dark Iron Rifle (16004, -1.87 DPS, sim-verified) [crafted] |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Officer's Pauldrons; chest: Stone Guard's Plate Armor; wrist: Officer's Wristguards; hands: Raider Gauntlets; waist: Prowler's Leather Waistguard; legs: Stone Guard's Plate Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 805, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (orc, 03325213032515001-05050000000000000-005000000000000000)

Set DPS (verified): 274.0. Weights run: 1.3s. Verify run: 13.1s. 1556 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.000 ± 0.005, agility=1.440 ± 0.091, crit=2.058 ± 0.130 per rating point (14 rating = 1%, 28.807 per %), hit=3.989 ± 0.345 per rating point (10 rating = 1%, 39.886 per %), melee_haste=16.659 ± 2.849

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 173.4 attack_power points (14.72 DPS) | yes | Champion's Plate Helm (227043, -5.32 DPS) [pvp]; Blood Guard's Plate Helm (220803, -6.68 DPS) [vendor]; Stalwart Helm (250599, -6.77 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rage of Mugamba (19577, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.94 DPS) [quest]; Amulet of the Darkmoon (19491, -1.40 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 67.7 attack_power points (5.75 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Champion's Plate Shoulders (227042, -0.42 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 67.9 attack_power points (5.76 DPS) | yes | Cape of the Black Baron (13340, -2.23 DPS) [dungeon]; Deathguard's Cloak (20068, -2.27 DPS) [rep]; Stalwart Cloak (272415, -2.38 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Obsidian Mail Tunic (22191, -1.35 DPS) [crafted]; Breastplate of Heroism (226862, -1.71 DPS) [quest]; Dawn Armor (252483, -4.39 DPS, sim-verified) [crafted] |
| wrist | Berserker Bracers (19578) | Warsong Outriders [rep] | 49.5 attack_power points (4.21 DPS) | yes | Forest Stalker's Bracers (19587, -0.01 DPS) [rep]; General's Plate Armguards (16546, -0.46 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 69.8 attack_power points (5.93 DPS) | yes | General's Plate Gauntlets (231532, -0.09 DPS) [pvp]; Gauntlets of Heroism (226861, -0.43 DPS) [quest]; Stormshroud Gloves (21278, -4.77 DPS, sim-verified) [crafted] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Ferocity of the Timbermaw (227805, -0.54 DPS) [vendor]; General's Plate Girdle (16547, -1.32 DPS) [pvp]; Belt of Preserved Heads (20216, -2.93 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 168.6 attack_power points (14.31 DPS) | yes | Sentinel's Plate Legguards (237825, +0.00 DPS, sim-verified) [vendor]; Sentinel's Chain Leggings (237819, -1.75 DPS) [vendor]; General's Plate Leggings (231533, -2.64 DPS) [pvp] |
| feet | Boots of Heroism (21995) (or Battleboots of Heroism (226857)) | Anthion's Parting Words [quest] | 79.9 attack_power points (6.78 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Fine Dawn Treaders (227815, -2.05 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 87.9 attack_power points (7.46 DPS) | yes | Cutthroat's Signet (272408, -2.36 DPS) [vendor]; Channeler's Ring (272406, -4.08 DPS) [vendor]; Blessed Band of Light (272407, -4.08 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 84.7 attack_power points (7.19 DPS) | yes | Cutthroat's Signet (272408, -2.70 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.80 DPS) [vendor]; Blessed Band of Light (272407, -3.80 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.89 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.23 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Arcanite Champion (12790, -8.83 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.23 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.73 DPS) [world_drop]; Dark Iron Rifle (16004, -2.68 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Berserker Bracers; hands: Raider Gauntlets; waist: Radiant Girdle of the Dawn; legs: Titanic Leggings; feet: Boots of Heroism; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: The Purifier

No-known-source sample (15 of 1556, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (orc, 02305213032515001-55050000001000000-200000000000000000)

Set DPS (verified): 795.1. Weights run: 1.5s. Verify run: 15.6s. 1556 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.003, strength=2.200 ± 0.006, agility=1.769 ± 0.133, crit=2.373 ± 0.179 per rating point (14 rating = 1%, 33.216 per %), hit=4.041 ± 0.527 per rating point (10 rating = 1%, 40.414 per %), melee_haste=18.525 ± 4.164

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 186.9 attack_power points (32.17 DPS) | yes | Champion's Plate Helm (227043, -11.54 DPS) [pvp]; Blood Guard's Plate Helm (220803, -14.57 DPS) [vendor]; Stalwart Helm (250599, -17.36 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (795.1 DPS) | yes | Medallion of the Dawn (22659, -1.24 DPS) [quest]; Amulet of the Darkmoon (19491, -1.51 DPS) [quest]; Rage of Mugamba (19577, -6.56 DPS, sim-verified) [quest] |
| shoulder | Defiler's Plate Spaulders (20212) | The Defilers [rep] | sim-verified (795.1 DPS) | yes | Warlord's Plate Shoulders (231534, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -5.28 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 68.4 attack_power points (11.78 DPS) | yes | Deathguard's Cloak (20068, -4.40 DPS) [rep]; Stalwart Cloak (272415, -4.82 DPS) [vendor]; Cape of the Black Baron (13340, -5.84 DPS, sim-verified) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (795.1 DPS) | yes | Breastplate of Heroism (226862, -4.83 DPS) [quest]; Obsidian Mail Tunic (22191, -4.90 DPS) [crafted]; Dawn Armor (252483, -9.66 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 57.8 attack_power points (9.95 DPS) | yes | Berserker Bracers (19578, -0.32 DPS) [rep]; General's Plate Armguards (16546, -1.38 DPS) [pvp] |
| hands | Raider Gauntlets (272095) | Creeg Bothunk [vendor] | 78.9 attack_power points (13.58 DPS) | yes | General's Plate Gauntlets (231532, -0.28 DPS) [pvp]; Gauntlets of Heroism (226861, -1.04 DPS) [quest]; Stormshroud Gloves (21278, -10.54 DPS, sim-verified) [crafted] |
| waist | Defiler's Plate Girdle (20204) | The Defilers [rep] | sim-verified (795.1 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -9.11 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 180.0 attack_power points (31.00 DPS) | yes | Sentinel's Plate Legguards (237825, -2.38 DPS) [vendor]; General's Plate Leggings (231533, -5.03 DPS) [pvp]; Sentinel's Chain Leggings (237819, -5.18 DPS, sim-verified) [vendor] |
| feet | Defiler's Plate Greaves (20208) | The Defilers [rep] | sim-verified (795.1 DPS) | yes | Battleboots of Heroism (226857, +0.00 DPS) [vendor]; General's Plate Boots (231531, +0.00 DPS) [pvp]; Boots of Heroism (21995, -8.20 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 93.2 attack_power points (16.05 DPS) | yes | Cutthroat's Signet (272408, -4.83 DPS) [vendor]; Legionnaire's Band (19510, -8.15 DPS) [rep]; Channeler's Ring (272406, -9.09 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 89.6 attack_power points (15.43 DPS) | yes | Cutthroat's Signet (272408, -5.94 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -7.54 DPS) [rep]; Channeler's Ring (272406, -8.47 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (795.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (795.1 DPS) | yes | Frozen Heart of the Mountain (249469, -5.84 DPS) [crafted]; Molten Heart of the Mountain (249470, -17.18 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (795.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; The Unstoppable Force (19323, -20.90 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (795.1 DPS) | yes | Bloodseeker (19107, -0.56 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.45 DPS) [world_drop]; Dark Iron Rifle (16004, -7.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Defiler's Plate Spaulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gauntlets; waist: Defiler's Plate Girdle; legs: Titanic Leggings; feet: Defiler's Plate Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Arcanite Champion; ranged: The Purifier

No-known-source sample (15 of 1556, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

