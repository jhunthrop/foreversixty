# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 40.7. Weights run: 2.0s. Verify run: 3.1s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.542 ± 0.026, crit=0.879 ± 0.024 per rating point (14 rating = 1%, 12.309 per %), hit=1.449 ± 0.066 per rating point (10 rating = 1%, 14.490 per %), melee_haste=8.598 ± 0.508

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.86 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.20 DPS) [crafted]; Brawler's Leather Hood (252504, -0.33 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.2 attack_power points (0.40 DPS) | yes | Erudite's Amulet (277204, -0.22 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.7 attack_power points (0.33 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.2 attack_power points (0.40 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (40.7 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.88 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.2 attack_power points (0.44 DPS) | yes | Patterned Bronze Bracers (2868, +0.00 DPS, sim-verified) [crafted]; Death Bindings (286981, -0.01 DPS) [dungeon]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.05 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (40.7 DPS) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -2.19 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.01 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.36 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.2 attack_power points (0.61 DPS) | yes | Demon Band (12054, -0.27 DPS) [world_drop]; The 1 Ring (8350, -0.46 DPS) [world]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.2 attack_power points (0.40 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.25 DPS) [world]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (15.49 DPS) | yes | Smite's Mighty Hammer (7230, -2.68 DPS, sim-verified) [dungeon]; Duskbringer (2205, -2.78 DPS) [dungeon]; Sword of the Fallen (286977, -3.06 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 82.1. Weights run: 2.0s. Verify run: 3.6s. 431 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.292 ± 0.024, crit=0.970 ± 0.026 per rating point (14 rating = 1%, 13.585 per %), hit=1.574 ± 0.069 per rating point (10 rating = 1%, 15.736 per %), melee_haste=8.481 ± 0.667

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.6 attack_power points (1.43 DPS) | yes | Tusken Helm (6686, -0.17 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.27 DPS) [crafted]; Defender's Leather Helm (252455, -0.27 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Fallen Guard's Pendant (279837, -0.10 DPS) [quest]; Sentinel's Medallion (19541, -0.18 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.2 attack_power points (1.16 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.37 DPS) [crafted]; Golden Scale Shoulders (3841, -0.49 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.0 attack_power points (0.72 DPS) | yes | Sergeant Major's Cape (16315, +0.00 DPS, sim-verified) [pvp]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.44 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.77 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.01 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (82.1 DPS) | yes | Mail Combat Gauntlets (4075, +0.00 DPS) [world_drop]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -3.45 DPS, sim-verified) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (82.1 DPS) | yes | Girdle of Golem Strength (9405, +0.00 DPS) [world_drop]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Officer's Belt (250556, -3.53 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (82.1 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Brawler's Leather Legguards (252516, +0.00 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -3.37 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (82.1 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Alacritous Treads (277234, +0.00 DPS) [quest]; Trouncing Boots (4464, -3.09 DPS, sim-verified) [world] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.9 attack_power points (0.96 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Tiger Band (6749, -0.28 DPS) [quest]; Silverlaine's Family Seal (6321, -0.48 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.8 attack_power points (0.95 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Tiger Band (6749, -0.28 DPS) [quest]; Silverlaine's Family Seal (6321, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (82.1 DPS) | yes | Morbid Dawn (7689, -0.10 DPS) [dungeon]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -19.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 431, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 133.1. Weights run: 2.0s. Verify run: 4.0s. 589 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.701 ± 0.034, crit=1.652 ± 0.040 per rating point (14 rating = 1%, 23.124 per %), hit=1.971 ± 0.092 per rating point (10 rating = 1%, 19.715 per %), melee_haste=10.148 ± 0.814

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 44.4 attack_power points (2.38 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.60 DPS) [crafted]; Hard Gold Coif (250537, -0.88 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-verified (133.1 DPS) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; Ghostshard Talisman (7731, -0.25 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.7 attack_power points (1.65 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.46 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.32 DPS) [quest]; Dark Hooded Cape (5257, -0.51 DPS, sim-verified) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Quillward Harness (10583, -5.74 DPS, sim-verified) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.8 attack_power points (1.12 DPS) | yes | Branded Leather Bracers (19508, -0.04 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.14 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -1.39 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 33.3 attack_power points (1.79 DPS) | yes | Highlander's Chain Girdle (20089, -0.12 DPS) [rep]; Officer's Belt (250556, -0.17 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.18 DPS) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.9 attack_power points (2.03 DPS) | yes | Blackforge Greaves (6423, -0.16 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.28 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 29.6 attack_power points (1.59 DPS) | yes | Falcon's Hook (7552, -0.45 DPS) [dungeon]; Thunderbrow Ring (13097, -0.46 DPS) [world_drop]; Ring of the Underwood (2951, -0.46 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.3 attack_power points (1.25 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Ring of the Underwood (2951, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -19.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger1: Protector's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 589, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 184.0. Weights run: 2.1s. Verify run: 5.3s. 753 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.900 ± 0.042, crit=2.248 ± 0.055 per rating point (14 rating = 1%, 31.475 per %), hit=2.434 ± 0.121 per rating point (10 rating = 1%, 24.337 per %), melee_haste=12.658 ± 1.243

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Lamellar Helm (220819) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ornate Mithril Helm (7937, +0.00 DPS) [crafted]; Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Knight-Lieutenant's Imbued Helmet (220810, -5.03 DPS, sim-verified) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.7 attack_power points (1.95 DPS) | yes | Sentinel's Medallion (19540, -0.81 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.86 DPS) [quest] |
| shoulder | Knight-Lieutenant's Lamellar Pauldrons (220818) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, +0.00 DPS) [crafted]; Knight-Lieutenant's Imbued Pauldrons (220808, -4.91 DPS, sim-verified) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Dark Hooded Cape (5257, -0.04 DPS) [world] |
| chest | Knight's Lamellar Chestplate (220815) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Knight's Plate Hauberk (220794, +0.00 DPS) [vendor]; Knight's Imbued Armor (220813, -5.15 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.3 attack_power points (2.04 DPS) | yes | Officer's Wristguards (250581, -0.09 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.26 DPS) [crafted] |
| hands | Sergeant Major's Lamellar Gauntlets (220817) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Gauntlets (272096, +0.00 DPS) [vendor]; Raider Gloves (272100, +0.00 DPS) [vendor]; Sergeant Major's Imbued Gauntlets (220812, -4.49 DPS, sim-verified) [vendor] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 55.7 attack_power points (3.04 DPS) | yes | Highlander's Plate Girdle (20124, -0.12 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.15 DPS) [crafted]; Prowler's Leather Waistguard (252473, -2.53 DPS, sim-verified) [crafted] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormshroud Pants (15057, +0.00 DPS) [crafted]; Knight's Plate Leggings (220797, +0.00 DPS) [vendor]; Knight's Imbued Leggings (220809, -5.15 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Lamellar Boots (220814) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battlechaser's Greaves (12555, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Sergeant Major's Imbued Greaves (220811, -4.42 DPS, sim-verified) [vendor] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 39.1 attack_power points (2.13 DPS) | yes | Ironspine's Eye (7686, -0.72 DPS) [dungeon]; Falcon's Hook (7552, -0.84 DPS) [dungeon]; Ring of the Underwood (2951, -0.86 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.6 attack_power points (1.45 DPS) | yes | Ironspine's Eye (7686, -0.04 DPS) [dungeon]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.17 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sanctified Orb (20512, +0.00 DPS) [quest] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (184.0 DPS) | yes | - |
| main_hand | Blight (7959) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS, sim-verified) [world_drop]; Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Knight-Lieutenant's Lamellar Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Lamellar Pauldrons; back: Dark Phantom Cape; chest: Knight's Lamellar Chestplate; wrist: Deepfury Bracers; hands: Sergeant Major's Lamellar Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Lamellar Legplates; feet: Sergeant Major's Lamellar Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Blight

No-known-source sample (15 of 753, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 285.8. Weights run: 2.2s. Verify run: 12.4s. 1685 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.161 ± 0.053, crit=3.055 ± 0.076 per rating point (14 rating = 1%, 42.765 per %), hit=3.528 ± 0.177 per rating point (10 rating = 1%, 35.283 per %), melee_haste=21.367 ± 2.012

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 195.7 attack_power points (11.31 DPS) | yes | Outlaw's Collar (279253, -4.98 DPS) [crafted]; Knight-Lieutenant's Plate Helm (220804, -5.15 DPS) [vendor]; Eye of Rend (12587, -6.11 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.21 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -9.33 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 63.3 attack_power points (3.66 DPS) | yes | Windshear Cape (20691, -0.77 DPS) [world]; Cloak of the Honor Guard (20073, -1.07 DPS) [rep]; Cape of the Black Baron (13340, -1.53 DPS, sim-verified) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; The Postmaster's Tunic (13388, -5.29 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -6.64 DPS, sim-verified) [dungeon] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Devout Gloves (16692, -6.01 DPS, sim-verified) [dungeon] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Devout Belt (16696, -3.66 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -0.94 DPS) [vendor]; Titanic Leggings (22385, -0.99 DPS) [crafted]; Cloudkeeper Legplates (14554, -8.07 DPS, sim-verified) [world_drop] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Devout Sandals (16691, -8.41 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (285.8 DPS) | yes | Tarnished Elven Ring (18500, -1.52 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; Painweaver Band (13098, -2.04 DPS) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.18 DPS) [dungeon]; Cutthroat's Signet (272408, -1.30 DPS) [vendor]; Naglering (11669, -6.34 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Draconic Infused Emblem (22268, -2.40 DPS, sim-verified) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Draconic Infused Emblem (22268, -3.78 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, -1.63 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Devout Mantle; back: Howler's Furs; chest: Devout Robe; wrist: Devout Bracers; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Sentinel's Chain Leggings; feet: Boots of Elements; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Burst of Knowledge; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 55003003100100000-0000000000000000-05225331001330320)

Set DPS (verified): 644.8. Weights run: 2.5s. Verify run: 13.1s. 1685 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.096 ± 0.048, crit=2.700 ± 0.065 per rating point (14 rating = 1%, 37.795 per %), hit=5.321 ± 0.303 per rating point (10 rating = 1%, 53.210 per %), melee_haste=18.533 ± 2.503

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 225.6 attack_power points (25.39 DPS) | yes | Helm of the Executioner (22411, -10.04 DPS, sim-verified) [dungeon]; Stalwart Helm (250599, -10.14 DPS) [crafted]; Knight-Lieutenant's Plate Helm (220804, -11.61 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (644.8 DPS) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -1.48 DPS) [quest]; Mark of Fordring (15411, -1.51 DPS) [quest] |
| shoulder | Soulforge Pauldrons (226987) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -7.03 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 81.2 attack_power points (9.14 DPS) | yes | Shroud of Arcane Mastery (22330, -3.15 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.15 DPS) [vendor]; Stalwart Cloak (272415, -3.67 DPS, sim-verified) [vendor] |
| chest | Soulforge Breastplate (226973) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Magister's Robes (16688, -16.64 DPS, sim-verified) [dungeon] |
| wrist | Soulforge Bracers (226970) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Magister's Bindings (16683, -14.05 DPS, sim-verified) [dungeon] |
| hands | Soulforge Gauntlets (226975) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Magister's Gloves (16684, -15.56 DPS, sim-verified) [dungeon] |
| waist | Soulforge Belt (226971) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Magister's Belt (16685, -12.24 DPS, sim-verified) [dungeon] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -1.65 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.55 DPS) [vendor]; Cloudkeeper Legplates (14554, -14.06 DPS, sim-verified) [world_drop] |
| feet | Soulforge Sabatons (226991) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Magister's Boots (16682, -6.24 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Tarnished Elven Ring (18500, -3.00 DPS) [dungeon]; Cutthroat's Signet (272408, -3.23 DPS) [vendor]; The Postmaster's Seal (13392, -10.68 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.52 DPS) [dungeon]; Cutthroat's Signet (272408, -2.75 DPS) [vendor]; Naglering (11669, -10.47 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Draconic Infused Emblem (22268, -7.83 DPS, sim-verified) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.51 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Invocation (249442, +0.00 DPS) [crafted] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Soulforge Pauldrons; back: Howler's Furs; chest: Soulforge Breastplate; wrist: Soulforge Bracers; hands: Soulforge Gauntlets; waist: Soulforge Belt; legs: Titanic Leggings; feet: Soulforge Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1685, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 40.2. Weights run: 2.0s. Verify run: 3.3s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.542 ± 0.026, crit=0.879 ± 0.024 per rating point (14 rating = 1%, 12.309 per %), hit=1.449 ± 0.066 per rating point (10 rating = 1%, 14.490 per %), melee_haste=8.598 ± 0.508

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.86 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.20 DPS) [crafted]; Brawler's Leather Hood (252504, -0.33 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 9.2 attack_power points (0.40 DPS) | yes | Erudite's Amulet (277204, -0.27 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.7 attack_power points (0.33 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.46 DPS, sim-verified) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.65 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Forest Leather Bracers (3202, -0.10 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.81 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -1.89 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.49 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.36 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.2 attack_power points (0.61 DPS) | yes | Demon Band (12054, -0.27 DPS) [world_drop]; Loop of Sacrifice (281673, -0.35 DPS) [quest]; Bounty Hunter's Ring (5351, -0.41 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.2 attack_power points (0.40 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; Loop of Sacrifice (281673, -0.14 DPS) [quest]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Duskbringer (2205, -0.42 DPS) [dungeon]; Hammerbone (270018, -0.56 DPS, sim-verified) [quest]; Sword of the Fallen (286977, -0.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 88.0. Weights run: 2.0s. Verify run: 2.2s. 409 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.292 ± 0.024, crit=0.970 ± 0.026 per rating point (14 rating = 1%, 13.585 per %), hit=1.574 ± 0.069 per rating point (10 rating = 1%, 15.736 per %), melee_haste=8.481 ± 0.667

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.6 attack_power points (1.43 DPS) | yes | Tusken Helm (6686, -0.17 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.27 DPS) [crafted]; Defender's Leather Helm (252455, -0.27 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.67 DPS) | yes | Kaleidoscope Chain (13084, -0.04 DPS) [world_drop]; Scout's Medallion (19537, -0.18 DPS) [rep]; River Pride Choker (13087, -0.29 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.2 attack_power points (1.16 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.37 DPS) [crafted]; Golden Scale Shoulders (3841, -0.49 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.0 attack_power points (0.72 DPS) | yes | Construct Cloak (279848, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wildhunter Cloak (16658, -0.24 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.44 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.77 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.01 DPS) [world_drop]; Undead Knight's Bracers (251965, -0.14 DPS) [quest] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.06 DPS) | yes | Insignia Gloves (6408, -0.01 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.04 DPS) [world_drop]; Warsong Gauntlets (16978, -0.10 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.8 attack_power points (1.33 DPS) | yes | Prowler's Leather Belt (252459, +0.00 DPS, sim-verified) [crafted]; Girdle of Golem Strength (9405, -0.18 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.18 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 31.0 attack_power points (1.49 DPS) | yes | Ferine Leggings (6690, -0.24 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.26 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.29 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 23.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Boots (252439, -0.32 DPS) [crafted]; Feet of the Lynx (1121, -0.32 DPS) [world_drop]; Veteran's Boots (250503, -0.38 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.9 attack_power points (0.96 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Tiger Band (6749, -0.28 DPS) [quest]; Band of the Fist (17694, -0.32 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.8 attack_power points (0.95 DPS) | yes | Tiger Band (6749, -0.28 DPS) [quest]; Ironspine's Eye (7686, -0.31 DPS, sim-verified) [dungeon]; Band of the Fist (17694, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (88.0 DPS) | yes | Morbid Dawn (7689, -0.10 DPS) [dungeon]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -21.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 409, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 143.0. Weights run: 2.0s. Verify run: 4.1s. 558 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.701 ± 0.034, crit=1.652 ± 0.040 per rating point (14 rating = 1%, 23.124 per %), hit=1.971 ± 0.092 per rating point (10 rating = 1%, 19.715 per %), melee_haste=10.148 ± 0.814

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 44.4 attack_power points (2.38 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.60 DPS) [crafted]; Hard Gold Coif (250537, -0.88 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-verified (143.0 DPS) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Ethereal Talisman (4430, -0.10 DPS) [quest]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.7 attack_power points (1.65 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.46 DPS) [crafted] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Construct Cloak (279848, -0.44 DPS) [quest]; Dark Hooded Cape (5257, -2.04 DPS, sim-verified) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Quillward Harness (10583, -6.20 DPS, sim-verified) [dungeon] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.8 attack_power points (1.12 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS, sim-verified) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -2.91 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 33.3 attack_power points (1.79 DPS) | yes | Officer's Belt (250556, -0.17 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.18 DPS) [rep]; Defiler's Chain Girdle (20153, -0.60 DPS, sim-verified) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.9 attack_power points (2.03 DPS) | yes | Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.28 DPS) [crafted]; Blackforge Greaves (6423, -0.97 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 29.6 attack_power points (1.59 DPS) | yes | Falcon's Hook (7552, -0.45 DPS) [dungeon]; Thunderbrow Ring (13097, -0.46 DPS) [world_drop]; Ring of the Underwood (2951, -0.46 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.3 attack_power points (1.25 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Thunderbrow Ring (13097, -0.12 DPS) [world_drop]; Ring of the Underwood (2951, -0.12 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -24.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: First Sergeant's Cloak; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 558, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 180.7. Weights run: 2.1s. Verify run: 4.8s. 733 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.900 ± 0.042, crit=2.248 ± 0.055 per rating point (14 rating = 1%, 31.475 per %), hit=2.434 ± 0.121 per rating point (10 rating = 1%, 24.337 per %), melee_haste=12.658 ± 1.243

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormcloth Headband (10032) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ornate Mithril Helm (7937, +0.00 DPS) [crafted]; Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -6.53 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.7 attack_power points (1.95 DPS) | yes | Woven Ivy Necklace (19159, -0.29 DPS) [quest]; Scout's Medallion (19535, -0.70 DPS) [rep] |
| shoulder | Stormcloth Shoulders (10038) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, +0.00 DPS) [crafted]; Blessed Plate Pauldrons (250586, -2.56 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (180.7 DPS) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Dark Hooded Cape (5257, -0.04 DPS) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, +0.00 DPS) [crafted]; Stone Guard's Plate Armor (220801, -5.66 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.3 attack_power points (2.04 DPS) | yes | Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Officer's Wristguards (250581, +0.00 DPS) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Raider Gauntlets (272096, +0.00 DPS) [vendor]; Raider Gloves (272100, +0.00 DPS) [vendor]; Blessed Plate Gauntlet (250588, -8.41 DPS, sim-verified) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.6 attack_power points (2.92 DPS) | yes | Defiler's Plate Girdle (20205, -0.01 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.03 DPS) [crafted]; Defiler's Chain Girdle (20151, -0.12 DPS) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Stormshroud Pants (15057, +0.00 DPS) [crafted]; Stone Guard's Plate Leggings (220798, -1.42 DPS, sim-verified) [vendor] |
| feet | Stormcloth Boots (10039) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battlechaser's Greaves (12555, +0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Blessed Plate Boots (250587, -3.79 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 39.1 attack_power points (2.13 DPS) | yes | Ironspine's Eye (7686, -0.72 DPS) [dungeon]; White Bone Band (11862, -0.82 DPS) [quest]; Falcon's Hook (7552, -0.84 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.6 attack_power points (1.45 DPS) | yes | Ironspine's Eye (7686, -0.04 DPS) [dungeon]; White Bone Band (11862, -0.14 DPS) [quest]; Falcon's Hook (7552, -0.16 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.52 DPS, sim-verified) [crafted] |
| main_hand | Nightblade (1982) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Blight (7959, -1.93 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Stormcloth Headband; neck: Skibi's Pendant; shoulder: Stormcloth Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Prowler's Leather Waistguard; feet: Stormcloth Boots; finger1: Legionnaire's Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Nightblade

No-known-source sample (15 of 733, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 293.3. Weights run: 2.2s. Verify run: 11.5s. 1710 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.161 ± 0.053, crit=3.055 ± 0.076 per rating point (14 rating = 1%, 42.765 per %), hit=3.528 ± 0.177 per rating point (10 rating = 1%, 35.283 per %), melee_haste=21.367 ± 2.012

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 195.7 attack_power points (11.31 DPS) | yes | Eye of Rend (12587, -4.81 DPS, sim-verified) [dungeon]; Outlaw's Collar (279253, -4.98 DPS) [crafted]; Blood Guard's Plate Helm (220803, -5.15 DPS) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.21 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -3.41 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 63.3 attack_power points (3.66 DPS) | yes | Cape of the Black Baron (13340, -0.63 DPS) [dungeon]; Windshear Cape (20691, -0.77 DPS) [world]; Deathguard's Cloak (20068, -1.07 DPS) [rep] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Tunic (13388, +0.00 DPS) [dungeon]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -1.08 DPS, sim-verified) [dungeon] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Savage Gladiator Grips (11730, -1.47 DPS, sim-verified) [dungeon] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, -1.95 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Plate Legguards (237825, -0.94 DPS) [vendor]; Titanic Leggings (22385, -0.99 DPS) [crafted]; Cloudkeeper Legplates (14554, -7.14 DPS, sim-verified) [world_drop] |
| feet | Magister's Boots (16682) | Stratholme: Hearthsinger Forresten [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; The Postmaster's Treads (13391, -2.49 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Tarnished Elven Ring (18500, -1.52 DPS) [dungeon]; Cutthroat's Signet (272408, -1.65 DPS) [vendor]; The Postmaster's Seal (13392, -3.08 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.18 DPS) [dungeon]; Cutthroat's Signet (272408, -1.30 DPS) [vendor]; Naglering (11669, -5.43 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, -1.08 DPS, sim-verified) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (293.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -1.92 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, -1.59 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Magister's Mantle; back: Howler's Furs; chest: Magister's Robes; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Sentinel's Chain Leggings; feet: Magister's Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1710, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 55003003100100000-0000000000000000-05225331001330320)

Set DPS (verified): 641.0. Weights run: 2.5s. Verify run: 12.3s. 1710 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.096 ± 0.048, crit=2.700 ± 0.065 per rating point (14 rating = 1%, 37.795 per %), hit=5.321 ± 0.303 per rating point (10 rating = 1%, 53.210 per %), melee_haste=18.533 ± 2.503

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 225.6 attack_power points (25.39 DPS) | yes | Helm of the Executioner (22411, -9.45 DPS, sim-verified) [dungeon]; Stalwart Helm (250599, -10.14 DPS) [crafted]; Blood Guard's Plate Helm (220803, -11.61 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pendant of Celerity (22340, +0.00 DPS) [dungeon]; Amulet of the Darkmoon (19491, -1.48 DPS) [quest]; Mark of Fordring (15411, -1.51 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -21.58 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 81.2 attack_power points (9.14 DPS) | yes | Shroud of Arcane Mastery (22330, -3.15 DPS) [dungeon]; Arcanoweave Cloak (272411, -3.15 DPS) [vendor]; Stalwart Cloak (272415, -3.47 DPS, sim-verified) [vendor] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Tunic (13388, +0.00 DPS, sim-verified) [dungeon]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Blackmist Armguards (12966, -0.13 DPS) [dungeon] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -8.48 DPS, sim-verified) [quest] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 118.5 attack_power points (13.34 DPS) | yes | Radiant Girdle of the Dawn (227814, -1.73 DPS) [vendor]; Ferocity of the Timbermaw (227805, -2.25 DPS) [vendor]; Marksman's Girdle (22232, -2.40 DPS) [dungeon] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -1.65 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.55 DPS) [vendor]; Cloudkeeper Legplates (14554, -17.23 DPS, sim-verified) [world_drop] |
| feet | Savage Gladiator Greaves (11731) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Fine Dawn Treaders (227815, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (641.0 DPS) | yes | Tarnished Elven Ring (18500, -3.00 DPS) [dungeon]; Cutthroat's Signet (272408, -3.23 DPS) [vendor]; Painweaver Band (13098, -6.47 DPS) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.52 DPS) [dungeon]; Cutthroat's Signet (272408, -2.75 DPS) [vendor]; Naglering (11669, -14.11 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, -7.81 DPS, sim-verified) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Libram of Fervor (23203, +0.00 DPS) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Ironfeather Shoulders; back: Howler's Furs; chest: Ironfeather Breastplate; wrist: Forest Stalker's Bracers; hands: Savage Gladiator Grips; waist: Belt of Preserved Heads; legs: Titanic Leggings; feet: Savage Gladiator Greaves; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1710, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

