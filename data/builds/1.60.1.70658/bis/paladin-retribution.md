# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 39.1. Weights run: 1.6s. Verify run: 2.5s. 225 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.511 ± 0.026, crit=0.862 ± 0.024 per rating point (14 rating = 1%, 12.065 per %), hit=1.411 ± 0.066 per rating point (10 rating = 1%, 14.112 per %), melee_haste=8.478 ± 0.511

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.85 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.21 DPS) [crafted]; Brawler's Leather Hood (252504, -0.34 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.0 attack_power points (0.13 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.6 attack_power points (0.32 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.1 attack_power points (0.39 DPS) | yes | Grave Shroud (279865, -0.00 DPS) [quest]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (39.1 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.20 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.0 attack_power points (0.43 DPS) | yes | Patterned Bronze Bracers (2868, -0.00 DPS) [crafted]; Death Bindings (286981, -0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.1 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.36 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (39.1 DPS) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -2.50 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.1 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.34 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.1 attack_power points (0.77 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.34 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| finger2 | The 1 Ring (8350) | Fishing [world] | 3.5 attack_power points (0.15 DPS) | yes | Lavishly Jeweled Ring (1156, -0.02 DPS) [dungeon]; Ring of the Moon (12052, -0.06 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (15.25 DPS) | yes | Duskbringer (2205, -2.73 DPS) [dungeon]; Smite's Mighty Hammer (7230, -2.74 DPS, sim-verified) [dungeon]; Sword of the Fallen (286977, -3.01 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: The 1 Ring; main_hand: Verigan's Fist

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 9602 Brushwood Blade; 14145 Cursed Felblade; 14147 Cavedweller Bracers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 79.6. Weights run: 1.6s. Verify run: 2.9s. 447 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.267 ± 0.024, crit=0.951 ± 0.025 per rating point (14 rating = 1%, 13.311 per %), hit=1.536 ± 0.068 per rating point (10 rating = 1%, 15.359 per %), melee_haste=8.395 ± 0.669

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.4 attack_power points (1.39 DPS) | yes | Tusken Helm (6686, -0.16 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.25 DPS) [crafted]; Defender's Leather Helm (252455, -0.25 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.62 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.24 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.33 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.13 DPS) | yes | Barbaric Iron Shoulders (7913, -0.20 DPS) [crafted]; Barbaric Shoulders (5964, -0.36 DPS) [crafted]; Golden Scale Shoulders (3841, -0.47 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.70 DPS) | yes | Sergeant Major's Cape (16315, -0.08 DPS) [pvp]; Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wolfmaster Cape (6314, -0.23 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.41 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.28 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.75 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (79.6 DPS) | yes | Mail Combat Gauntlets (4075, +0.00 DPS) [world_drop]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -3.61 DPS, sim-verified) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (79.6 DPS) | yes | Girdle of Golem Strength (9405, +0.00 DPS) [world_drop]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Officer's Belt (250556, -3.71 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (79.6 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Brawler's Leather Legguards (252516, +0.00 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -3.53 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (79.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Alacritous Treads (277234, -3.67 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.93 DPS) | yes | Silverlaine's Family Seal (6321, -0.46 DPS) [dungeon]; Bonebinder's Signet (273032, -0.46 DPS) [world]; Seal of Wrynn (2933, -0.47 DPS) [quest] |
| finger2 | Tiger Band (6749) (or Bonebinder's Signet (273032), Silverlaine's Family Seal (6321)) | Willix the Importer [quest] | 10.0 attack_power points (0.47 DPS) | yes | Silverlaine's Family Seal (6321, +0.00 DPS) [dungeon]; Bonebinder's Signet (273032, +0.00 DPS) [world]; Seal of Wrynn (2933, -0.01 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (79.6 DPS) | yes | Felbough Arm (273037, -1.32 DPS) [world]; Cobalt Crusher (7730, -1.87 DPS) [dungeon]; Viscous Hammer (13045, -19.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Tiger Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 447, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 124.7. Weights run: 1.6s. Verify run: 3.0s. 636 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.653 ± 0.034, crit=1.603 ± 0.039 per rating point (14 rating = 1%, 22.442 per %), hit=1.910 ± 0.090 per rating point (10 rating = 1%, 19.095 per %), melee_haste=9.988 ± 0.814

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.8 attack_power points (2.30 DPS) | yes | White Bandit Mask (10008, -0.19 DPS) [crafted]; Barbaric Iron Helm (7915, -0.58 DPS) [crafted]; Hard Gold Coif (250537, -0.83 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 14.6 attack_power points (0.77 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Gazlowe's Charm (13088, -0.35 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.2 attack_power points (1.48 DPS) | yes | Flintrock Shoulders (7755, -0.09 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.32 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.33 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 21.9 attack_power points (1.15 DPS) | yes | Yeti Fur Cloak (2805, -0.32 DPS) [quest]; Parachute Cloak (10518, -0.46 DPS) [crafted]; Hawkeye's Cloak (14593, -1.85 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (124.7 DPS) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Icemail Jerkin (1981, -5.56 DPS, sim-verified) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (1.08 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.14 DPS) [world_drop]; Yorgen Bracers (13012, -0.19 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (124.7 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -1.64 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.9 attack_power points (1.73 DPS) | yes | Officer's Belt (250556, -0.16 DPS) [crafted]; Prowler's Leather Belt (252459, -0.26 DPS) [crafted]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (124.7 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.6 attack_power points (1.97 DPS) | yes | Blackforge Greaves (6423, -0.16 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.28 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.0 attack_power points (1.10 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 20.9 attack_power points (1.10 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (124.7 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -18.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Chromite Barbute; back: Sergeant Major's Cape; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger2: Falcon's Hook; ranged: Libram of Invocation

No-known-source sample (15 of 636, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 50 (human, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 160.8. Weights run: 1.7s. Verify run: 3.8s. 833 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.735 ± 0.036, crit=2.035 ± 0.047 per rating point (14 rating = 1%, 28.484 per %), hit=2.315 ± 0.102 per rating point (10 rating = 1%, 23.152 per %), melee_haste=-4.927 ± 1.209

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ornate Mithril Helm (7937) | Blacksmithing [crafted] | 50.5 attack_power points (2.70 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Chromite Barbute (8142, -0.17 DPS) [dungeon]; Sunscale Helmet (14849, -0.26 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.6 attack_power points (1.80 DPS) | yes | Kaleidoscope Chain (13084, -0.95 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.15 DPS) [world]; River Pride Choker (13087, -1.33 DPS) [world_drop] |
| shoulder | Blessed Plate Pauldrons (250586) | Blacksmithing [crafted] | sim-verified (160.8 DPS) | yes | Wyrmslayer Spaulders (13066, +0.00 DPS) [world_drop]; Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -3.35 DPS, sim-verified) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (160.8 DPS) | yes | Pridelord Cape (14673, -0.07 DPS) [world_drop]; Blackflame Cape (13109, -0.15 DPS) [world_drop]; Dark Phantom Cape (13122, -1.46 DPS, sim-verified) [world_drop] |
| chest | Knight's Lamellar Chestplate (220815) | Captain Dirgehammer [vendor] | sim-verified (160.8 DPS) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Knight's Plate Hauberk (220794, +0.00 DPS) [vendor]; Knight's Imbued Armor (220813, -1.46 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.8 attack_power points (1.87 DPS) | yes | Officer's Wristguards (250581, -0.01 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.15 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.20 DPS) [crafted] |
| hands | Blessed Plate Gauntlet (250588) | Blacksmithing [crafted] | sim-verified (160.8 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Officer's Gloves (250551, +0.00 DPS) [crafted]; Sergeant Major's Lamellar Gauntlets (220817, -0.85 DPS, sim-verified) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 51.6 attack_power points (2.77 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Officer's Waistguard (250571, -0.19 DPS) [crafted]; Belt of the Gladiator (13134, -0.64 DPS) [world_drop] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | sim-verified (160.8 DPS) | yes | Gryphon Rider's Leggings (9652, +0.00 DPS) [quest]; Knight's Plate Leggings (220797, +0.00 DPS) [vendor]; Knight's Imbued Leggings (220809, -1.46 DPS, sim-verified) [vendor] |
| feet | Blessed Plate Boots (250587) | Blacksmithing [crafted] | sim-verified (160.8 DPS) | yes | Officer's Sabatons (250561, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -2.54 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.3 attack_power points (1.30 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Ring of the Underwood (2951, -0.14 DPS) [world_drop]; Blackstone Ring (17713, -0.23 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.8 attack_power points (1.22 DPS) | yes | Falcon's Hook (7552, -0.03 DPS) [dungeon]; Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Blackstone Ring (17713, -0.15 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (160.8 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (160.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (160.8 DPS) | yes | Warmonger (13052, -0.17 DPS) [world_drop]; Stoneraven (13059, -1.62 DPS) [world_drop]; Blight (7959, -2.89 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Ornate Mithril Helm; neck: Skibi's Pendant; shoulder: Blessed Plate Pauldrons; chest: Knight's Lamellar Chestplate; wrist: Deepfury Bracers; hands: Blessed Plate Gauntlet; waist: Prowler's Leather Waistguard; legs: Knight's Lamellar Legplates; feet: Blessed Plate Boots; finger1: Masons Fraternity Ring; finger2: Thunderbrow Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 833, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 235.5. Weights run: 1.7s. Verify run: 8.1s. 1755 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.270 ± 0.054, crit=3.202 ± 0.076 per rating point (14 rating = 1%, 44.823 per %), hit=3.725 ± 0.171 per rating point (10 rating = 1%, 37.252 per %), melee_haste=19.079 ± 2.031

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 203.8 attack_power points (11.15 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -5.10 DPS) [vendor]; Stalwart Helm (250599, -5.63 DPS) [crafted]; Outlaw's Collar (279253, -6.30 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.5 attack_power points (2.22 DPS) | yes | Kaleidoscope Chain (13084, -1.24 DPS) [world_drop]; Sunprism Pendant (20649, -1.48 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -1.56 DPS) [world] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (235.5 DPS) | yes | Lightforge Spaulders (16729, +0.00 DPS) [dungeon]; Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Magister's Mantle (16689, -5.03 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.0 attack_power points (2.96 DPS) | yes | Fel Cape (279269, -0.83 DPS, sim-verified) [crafted]; Sergeant Major's Cape (16337, -1.00 DPS) [pvp]; Dark Phantom Cape (13122, -1.09 DPS) [world_drop] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (235.5 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Magister's Robes (16688, -5.35 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (235.5 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Arena Wristguards (18709, +0.00 DPS) [world]; Magister's Bindings (16683, -5.67 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (235.5 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Magister's Gloves (16684, -5.58 DPS, sim-verified) [dungeon] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-verified (235.5 DPS) | yes | Might of the Timbermaw (19044, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Magister's Belt (16685, -4.39 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (235.5 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Knight-Captain's Lamellar Leggings (16435, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -3.15 DPS, sim-verified) [crafted] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (235.5 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Knight-Lieutenant's Lamellar Sabatons (227146, +0.00 DPS) [vendor]; Deathbone Sabatons (14621, -7.34 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.8 attack_power points (1.74 DPS) | yes | Falcon's Hook (7552, -0.26 DPS) [dungeon]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop]; Ring of Living Stone (18400, -0.54 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.1 attack_power points (1.48 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; Ring of Living Stone (18400, -0.28 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (235.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (235.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (235.5 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Arcanite Champion (12790, -7.98 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-verified (235.5 DPS) | yes | Libram of Infusion (279248, +0.00 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Devout Mantle; back: Cape of the Black Baron; chest: Devout Robe; wrist: Devout Bracers; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Magister's Leggings; feet: Boots of Elements; finger2: Ring of the Underwood; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (human, 52003000000000000-5030000000000000-05225331001330320)

Set DPS (verified): 543.3. Weights run: 1.9s. Verify run: 9.0s. 1755 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.893 ± 0.042, crit=2.409 ± 0.057 per rating point (14 rating = 1%, 33.722 per %), hit=4.753 ± 0.260 per rating point (10 rating = 1%, 47.528 per %), melee_haste=28.679 ± 2.808

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stalwart Helm (250599) | Blacksmithing [crafted] | sim-verified (543.3 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -1.27 DPS) [vendor]; Lieutenant Commander's Lamellar Headguard (227144, -3.13 DPS) [vendor]; Lionheart Helm (12640, -5.15 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.7 attack_power points (4.08 DPS) | yes | Sunprism Pendant (20649, -1.56 DPS) [quest]; Kaleidoscope Chain (13084, -2.16 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -2.75 DPS) [world] |
| shoulder | Blue Dragonscale Shoulders (15049) | Leatherworking [crafted] | sim-verified (543.3 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Lightforge Spaulders (16729, -8.33 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 48.4 attack_power points (5.38 DPS) | yes | Sergeant Major's Cape (16337, -1.55 DPS) [pvp]; Fel Cape (279269, -1.63 DPS) [crafted]; Dark Phantom Cape (13122, -2.22 DPS) [world_drop] |
| chest | Blue Dragonscale Breastplate (15048) | Leatherworking [crafted] | sim-verified (543.3 DPS) | yes | Shadowcraft Tunic (16721, +0.00 DPS) [dungeon]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -6.71 DPS, sim-verified) [crafted] |
| wrist | Bindings of Elements (16671) | Stratholme: Crypt Beast [dungeon] | sim-verified (543.3 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Officer's Wristguards (250581, +0.00 DPS) [crafted]; Tranquil Wristguards (279256, -4.97 DPS, sim-verified) [crafted] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (543.3 DPS) | yes | Knight-Lieutenant's Lamellar Gauntlets (23274, +0.00 DPS) [vendor]; Marshal's Lamellar Gauntlets (231650, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -10.06 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 77.3 attack_power points (8.60 DPS) | yes | Skulker's Leather Waistguard (252474, -2.42 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.53 DPS, sim-verified) [crafted]; Officer's Waistguard (250571, -2.73 DPS) [crafted] |
| legs | Blue Dragonscale Leggings (20295) | Leatherworking [crafted] | sim-verified (543.3 DPS) | yes | Cloudkeeper Legplates (14554, +0.00 DPS) [world_drop]; Azerothium Legplates (279265, +0.00 DPS) [crafted]; Warbear Woolies (15065, -12.95 DPS, sim-verified) [crafted] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (543.3 DPS) | yes | Knight-Lieutenant's Lamellar Sabatons (227146, +0.00 DPS) [vendor]; Scalegut Treaders (275618, +0.00 DPS) [crafted]; Shadowcraft Boots (16711, -7.78 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.5 attack_power points (2.95 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Ring of Living Stone (18400, -0.26 DPS) [quest]; Ring of the Underwood (2951, -0.30 DPS) [world_drop] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 25.0 attack_power points (2.79 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Ring of Living Stone (18400, -0.09 DPS) [quest]; Ring of the Underwood (2951, -0.14 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (543.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (543.3 DPS) | yes | Mark of the Chosen (17774, -6.83 DPS, sim-verified) [quest] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (543.3 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Arcanite Champion (12790, -17.31 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (543.3 DPS) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Stalwart Helm; shoulder: Blue Dragonscale Shoulders; back: Cape of the Black Baron; chest: Blue Dragonscale Breastplate; wrist: Bindings of Elements; hands: Gauntlets of Elements; waist: Might of the Timbermaw; legs: Blue Dragonscale Leggings; feet: Boots of Elements; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Blade of Hanna

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 38.9. Weights run: 1.6s. Verify run: 2.6s. 215 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.511 ± 0.026, crit=0.862 ± 0.024 per rating point (14 rating = 1%, 12.065 per %), hit=1.411 ± 0.066 per rating point (10 rating = 1%, 14.112 per %), melee_haste=8.478 ± 0.511

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.85 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.21 DPS) [crafted]; Brawler's Leather Hood (252504, -0.34 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.0 attack_power points (0.13 DPS) | yes | Roadwatcher's Confidence (281265, +0.00 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.6 attack_power points (0.32 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.1 attack_power points (0.39 DPS) | yes | Grave Shroud (279865, -0.00 DPS) [quest]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (38.9 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.34 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Forest Leather Bracers (3202, -0.10 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.9 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.36 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (38.9 DPS) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -2.43 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.9 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.12 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.1 attack_power points (0.77 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.34 DPS) [crafted] |
| finger1 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.34 DPS) | yes | Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; The 1 Ring (8350, -0.19 DPS) [world]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 6.0 attack_power points (0.26 DPS) | yes | The 1 Ring (8350, -0.11 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Bounty Hunter's Ring (5351, -0.81 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 306.2 attack_power points (13.02 DPS) | yes | Smite's Mighty Hammer (7230, -0.09 DPS) [dungeon]; Duskbringer (2205, -0.50 DPS) [dungeon]; Sword of the Fallen (286977, -0.77 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Demon Band; finger2: Loop of Sacrifice; main_hand: Hammerbone

No-known-source sample (15 of 215, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 86.1. Weights run: 1.6s. Verify run: 2.9s. 436 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.267 ± 0.024, crit=0.951 ± 0.025 per rating point (14 rating = 1%, 13.311 per %), hit=1.536 ± 0.068 per rating point (10 rating = 1%, 15.359 per %), melee_haste=8.395 ± 0.669

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.4 attack_power points (1.39 DPS) | yes | Tusken Helm (6686, -0.16 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.25 DPS) [crafted]; Defender's Leather Helm (252455, -0.25 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.62 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; River Pride Choker (13087, -0.24 DPS) [world_drop]; Erudite's Amulet (277204, -0.50 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.13 DPS) | yes | Barbaric Iron Shoulders (7913, -0.20 DPS) [crafted]; Barbaric Shoulders (5964, -0.36 DPS) [crafted]; Golden Scale Shoulders (3841, -0.47 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.70 DPS) | yes | Tigerstrike Mantle (13108, -0.22 DPS) [world_drop]; Wildhunter Cloak (16658, -0.23 DPS) [quest]; Construct Cloak (279848, -0.23 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.41 DPS) | yes | Shining Silver Breastplate (2870, -0.09 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.28 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.75 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (86.1 DPS) | yes | Mail Combat Gauntlets (4075, +0.00 DPS) [world_drop]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -4.20 DPS, sim-verified) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (86.1 DPS) | yes | Girdle of Golem Strength (9405, +0.00 DPS) [world_drop]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Officer's Belt (250556, -3.72 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (86.1 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Brawler's Leather Legguards (252516, +0.00 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -3.46 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (86.1 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Veteran's Boots (250503, +0.00 DPS) [crafted]; Brawler's Leather Boots (252439, -4.12 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.93 DPS) | yes | Silverlaine's Family Seal (6321, -0.46 DPS) [dungeon]; Tiger Band (6749, -0.46 DPS) [quest]; Bonebinder's Signet (273032, -0.46 DPS) [world] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 13.1 attack_power points (0.62 DPS) | yes | Silverlaine's Family Seal (6321, -0.14 DPS) [dungeon]; Tiger Band (6749, -0.14 DPS) [quest]; Bonebinder's Signet (273032, -0.14 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (86.1 DPS) | yes | Felbough Arm (273037, -1.32 DPS) [world]; Cobalt Crusher (7730, -1.87 DPS) [dungeon]; Viscous Hammer (13045, -22.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 132.1. Weights run: 1.6s. Verify run: 3.0s. 622 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.653 ± 0.034, crit=1.603 ± 0.039 per rating point (14 rating = 1%, 22.442 per %), hit=1.910 ± 0.090 per rating point (10 rating = 1%, 19.095 per %), melee_haste=9.988 ± 0.814

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 43.8 attack_power points (2.30 DPS) | yes | White Bandit Mask (10008, -0.19 DPS) [crafted]; Barbaric Iron Helm (7915, -0.58 DPS) [crafted]; Hard Gold Coif (250537, -0.83 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (132.1 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Ethereal Talisman (4430, -0.87 DPS, sim-verified) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.2 attack_power points (1.48 DPS) | yes | Flintrock Shoulders (7755, +0.00 DPS, sim-verified) [dungeon]; Hard Gold Pauldrons (250539, -0.32 DPS) [crafted]; Barbaric Iron Shoulders (7913, -0.33 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.6 attack_power points (0.92 DPS) | yes | Imperial Cloak (6432, -0.23 DPS) [dungeon]; Parachute Cloak (10518, -0.23 DPS) [crafted]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (132.1 DPS) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Icemail Jerkin (1981, -4.08 DPS, sim-verified) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.6 attack_power points (1.08 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS, sim-verified) [dungeon]; Hawkeye's Bracers (14590, -0.14 DPS) [world_drop]; Yorgen Bracers (13012, -0.19 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (132.1 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -2.25 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 32.9 attack_power points (1.73 DPS) | yes | Officer's Belt (250556, -0.16 DPS) [crafted]; Prowler's Leather Belt (252459, -0.26 DPS) [crafted]; Skulker's Leather Belt (252520, -0.32 DPS) [crafted] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (132.1 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.6 attack_power points (1.97 DPS) | yes | Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.28 DPS) [crafted]; Blackforge Greaves (6423, -0.96 DPS, sim-verified) [dungeon] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.0 attack_power points (1.10 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 20.9 attack_power points (1.10 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (132.1 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -23.17 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger2: Falcon's Hook

No-known-source sample (15 of 622, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 166.8. Weights run: 1.7s. Verify run: 3.4s. 832 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.735 ± 0.036, crit=2.035 ± 0.047 per rating point (14 rating = 1%, 28.484 per %), hit=2.315 ± 0.102 per rating point (10 rating = 1%, 23.152 per %), melee_haste=-4.927 ± 1.209

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormcloth Headband (10032) | Tailoring [crafted] | sim-verified (166.8 DPS) | yes | Chromite Barbute (8142, +0.00 DPS) [dungeon]; Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -5.66 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.6 attack_power points (1.80 DPS) | yes | Woven Ivy Necklace (19159, -0.25 DPS) [quest]; Ethereal Talisman (4430, -0.84 DPS) [quest]; Kaleidoscope Chain (13084, -0.95 DPS) [world_drop] |
| shoulder | Stormcloth Shoulders (10038) | Tailoring [crafted] | sim-verified (166.8 DPS) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, +0.00 DPS) [crafted]; Blessed Plate Pauldrons (250586, -2.02 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.0 attack_power points (1.39 DPS) | yes | Pridelord Cape (14673, -0.20 DPS) [world_drop]; Blackflame Cape (13109, -0.28 DPS) [world_drop]; Bloodlust Cape (14801, -0.33 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (166.8 DPS) | yes | Warforged Chestplate (11195, +0.00 DPS) [quest]; Warbear Harness (15064, +0.00 DPS) [crafted]; Stone Guard's Plate Armor (220801, -4.37 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.8 attack_power points (1.87 DPS) | yes | Officer's Wristguards (250581, -0.01 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.15 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.20 DPS) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (166.8 DPS) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Officer's Gloves (250551, +0.00 DPS) [crafted]; Blessed Plate Gauntlet (250588, -6.96 DPS, sim-verified) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 51.6 attack_power points (2.77 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Officer's Waistguard (250571, -0.19 DPS) [crafted]; Belt of the Gladiator (13134, -0.64 DPS) [world_drop] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (166.8 DPS) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, +0.00 DPS) [crafted]; Stone Guard's Plate Leggings (220798, -1.14 DPS, sim-verified) [vendor] |
| feet | Stormcloth Boots (10039) | Tailoring [crafted] | sim-verified (166.8 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Blessed Plate Boots (250587, -2.86 DPS, sim-verified) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.3 attack_power points (1.30 DPS) | yes | Thunderbrow Ring (13097, -0.08 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon]; Ring of the Underwood (2951, -0.14 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.29 DPS) | yes | Thunderbrow Ring (13097, -0.06 DPS) [world_drop]; Falcon's Hook (7552, -0.10 DPS) [dungeon]; Ring of the Underwood (2951, -0.12 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (166.8 DPS) | yes | Molten Heart of the Mountain (249470, -2.91 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (166.8 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (166.8 DPS) | yes | Warmonger (13052, -0.17 DPS) [world_drop]; Stoneraven (13059, -1.62 DPS) [world_drop]; Blight (7959, -2.39 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Stormcloth Headband; neck: Skibi's Pendant; shoulder: Stormcloth Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Prowler's Leather Waistguard; feet: Stormcloth Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 832, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 253.1. Weights run: 1.7s. Verify run: 8.1s. 1774 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.270 ± 0.054, crit=3.202 ± 0.076 per rating point (14 rating = 1%, 44.823 per %), hit=3.725 ± 0.171 per rating point (10 rating = 1%, 37.252 per %), melee_haste=19.079 ± 2.031

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 203.8 attack_power points (11.15 DPS) | yes | Outlaw's Collar (279253, -4.65 DPS, sim-verified) [crafted]; Blood Guard's Plate Helm (220803, -5.10 DPS) [vendor]; Stalwart Helm (250599, -5.63 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.5 attack_power points (2.22 DPS) | yes | Woven Ivy Necklace (19159, -0.38 DPS) [quest]; Sunprism Pendant (20649, -0.73 DPS) [quest]; Ethereal Talisman (4430, -1.12 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (253.1 DPS) | yes | Ironfeather Shoulders (15067, +0.00 DPS, sim-verified) [crafted]; Lightforge Spaulders (16729, +0.00 DPS) [dungeon]; Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.0 attack_power points (2.96 DPS) | yes | Fel Cape (279269, -0.51 DPS) [crafted]; Dark Phantom Cape (13122, -1.09 DPS) [world_drop]; Blackflame Cape (13109, -1.47 DPS) [world_drop] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (253.1 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Ironfeather Breastplate (15066, -0.82 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (253.1 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Officer's Wristguards (250581, +0.00 DPS) [crafted]; Arena Wristguards (18709, -0.96 DPS, sim-verified) [world] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (253.1 DPS) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Magister's Gloves (16684, -3.97 DPS, sim-verified) [dungeon] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-verified (253.1 DPS) | yes | Might of the Timbermaw (19044, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Magister's Belt (16685, -2.45 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (253.1 DPS) | yes | Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Azerothium Legplates (279265, +0.00 DPS) [crafted]; Warbear Woolies (15065, -3.24 DPS, sim-verified) [crafted] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (253.1 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Scalegut Treaders (275618, +0.00 DPS) [crafted]; Deathbone Sabatons (14621, -5.86 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 31.8 attack_power points (1.74 DPS) | yes | Falcon's Hook (7552, -0.26 DPS) [dungeon]; Thunderbrow Ring (13097, -0.40 DPS) [world_drop]; White Bone Band (11862, -0.43 DPS) [quest] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 27.1 attack_power points (1.48 DPS) | yes | Falcon's Hook (7552, -0.00 DPS) [dungeon]; Thunderbrow Ring (13097, -0.15 DPS) [world_drop]; White Bone Band (11862, -0.17 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (253.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (253.1 DPS) | yes | Molten Heart of the Mountain (249470, -1.37 DPS, sim-verified) [crafted] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (253.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Arcanite Champion (12790, -13.42 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Fervor (23203) | World drop [world_drop] | sim-verified (253.1 DPS) | yes | Libram of Infusion (279248, +0.00 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; shoulder: Magister's Mantle; back: Cape of the Black Baron; chest: Magister's Robes; wrist: Magister's Bindings; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Magister's Leggings; feet: Boots of Elements; finger2: Ring of the Underwood; main_hand: Blade of Hanna; ranged: Libram of Fervor

No-known-source sample (15 of 1774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 52003000000000000-5030000000000000-05225331001330320)

Set DPS (verified): 564.1. Weights run: 1.9s. Verify run: 8.5s. 1774 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=1.893 ± 0.042, crit=2.409 ± 0.057 per rating point (14 rating = 1%, 33.722 per %), hit=4.753 ± 0.260 per rating point (10 rating = 1%, 47.528 per %), melee_haste=28.679 ± 2.808

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stalwart Helm (250599) | Blacksmithing [crafted] | sim-verified (564.1 DPS) | yes | Blood Guard's Plate Helm (220803, -1.27 DPS) [vendor]; Outlaw's Collar (279253, -3.76 DPS) [crafted]; Lionheart Helm (12640, -5.71 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 36.7 attack_power points (4.08 DPS) | yes | Woven Ivy Necklace (19159, -0.57 DPS) [quest]; Sunprism Pendant (20649, -1.56 DPS) [quest]; Ethereal Talisman (4430, -1.90 DPS) [quest] |
| shoulder | Lightforge Spaulders (16729) | Blackrock Spire: The Beast [dungeon] | 74.0 attack_power points (8.23 DPS) | yes | Blood Guard's Plate Pauldrons (220796, -1.52 DPS) [vendor]; Golden Mantle of the Dawn (19058, -2.42 DPS) [crafted]; Dark Warder's Pauldrons (22241, -6.35 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 48.4 attack_power points (5.38 DPS) | yes | Fel Cape (279269, -1.63 DPS) [crafted]; Dark Phantom Cape (13122, -2.22 DPS) [world_drop]; Pridelord Cape (14673, -2.68 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 151.3 attack_power points (16.83 DPS) | yes | Timbermaw Tunic (252484, -0.40 DPS) [crafted]; Shadowcraft Tunic (16721, -7.13 DPS) [dungeon]; Stone Guard's Plate Armor (220801, -9.04 DPS) [vendor] |
| wrist | Lightforge Bracers (16722) | Scholomance: Lord Alexei Barov [dungeon] | sim-verified (564.1 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Officer's Wristguards (250581, +0.00 DPS) [crafted]; Tranquil Wristguards (279256, +0.00 DPS, sim-verified) [crafted] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 77.3 attack_power points (8.60 DPS) | yes | Devilsaur Gauntlets (15063, -1.74 DPS) [crafted]; Officer's Gloves (250551, -2.40 DPS) [crafted]; Gloves of the Pathfinder (21319, -2.57 DPS) [quest] |
| waist | Lightforge Belt (16723) | Stratholme: Rockwing Gargoyle [dungeon] | sim-verified (564.1 DPS) | yes | Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -4.37 DPS, sim-verified) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 101.8 attack_power points (11.33 DPS) | yes | Azerothium Legplates (279265, -1.08 DPS) [crafted]; Cloudkeeper Legplates (14554, -1.73 DPS) [world_drop]; Stone Guard's Plate Leggings (220798, -2.24 DPS) [vendor] |
| feet | Lightforge Boots (16725) | Stratholme: Balnazzar [dungeon] | sim-verified (564.1 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Scalegut Treaders (275618, +0.00 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.5 attack_power points (2.95 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Ring of Living Stone (18400, -0.26 DPS) [quest]; White Bone Band (11862, -0.28 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 25.0 attack_power points (2.79 DPS) | yes | Falcon's Hook (7552, -0.08 DPS) [dungeon]; Ring of Living Stone (18400, -0.09 DPS) [quest]; White Bone Band (11862, -0.12 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (564.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (564.1 DPS) | yes | Molten Heart of the Mountain (249470, -4.42 DPS, sim-verified) [crafted] |
| main_hand | Blade of Hanna (2801) | World drop [world_drop] | sim-verified (564.1 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Arcanite Champion (12790, -11.48 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Stalwart Helm; shoulder: Lightforge Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Lightforge Bracers; hands: Timbermaw Brawlers; waist: Lightforge Belt; legs: Warbear Woolies; feet: Lightforge Boots; finger2: Thunderbrow Ring; main_hand: Blade of Hanna

No-known-source sample (15 of 1774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

