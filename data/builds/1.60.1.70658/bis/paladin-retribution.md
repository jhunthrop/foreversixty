# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 40.7. Weights run: 1.6s. Verify run: 2.6s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.522 ± 0.026, crit=0.868 ± 0.024 per rating point (14 rating = 1%, 12.154 per %), hit=1.432 ± 0.065 per rating point (10 rating = 1%, 14.319 per %), melee_haste=8.491 ± 0.501

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.86 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.21 DPS) [crafted]; Brawler's Leather Hood (252504, -0.34 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.1 attack_power points (0.39 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.6 attack_power points (0.33 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.1 attack_power points (0.39 DPS) | yes | Grave Shroud (279865, +0.00 DPS, sim-verified) [quest]; Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (40.7 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.78 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.1 attack_power points (0.44 DPS) | yes | Patterned Bronze Bracers (2868, +0.00 DPS, sim-verified) [crafted]; Death Bindings (286981, -0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.96 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (40.7 DPS) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -2.11 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.95 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.2 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.35 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.1 attack_power points (0.61 DPS) | yes | The 1 Ring (8350, -0.46 DPS) [world]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; Ring of the Moon (12052, -0.52 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.35 DPS) | yes | The 1 Ring (8350, -0.19 DPS) [world]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon]; Ring of the Moon (12052, -0.26 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (15.49 DPS) | yes | Smite's Mighty Hammer (7230, -2.62 DPS, sim-verified) [dungeon]; Duskbringer (2205, -2.78 DPS) [dungeon]; Sword of the Fallen (286977, -3.06 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Verigan's Fist

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 82.8. Weights run: 1.6s. Verify run: 3.0s. 418 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.279 ± 0.024, crit=0.960 ± 0.025 per rating point (14 rating = 1%, 13.443 per %), hit=1.559 ± 0.068 per rating point (10 rating = 1%, 15.585 per %), melee_haste=8.395 ± 0.659

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.5 attack_power points (1.42 DPS) | yes | Tusken Helm (6686, -0.17 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.26 DPS) [crafted]; Defender's Leather Helm (252455, -0.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.63 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Sentinel's Medallion (19541, -0.14 DPS) [rep]; River Pride Choker (13087, -0.25 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.16 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.37 DPS) [crafted]; Golden Scale Shoulders (3841, -0.48 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.72 DPS) | yes | Sergeant Major's Cape (16315, -0.09 DPS) [pvp]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.44 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.77 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (82.8 DPS) | yes | Mail Combat Gauntlets (4075, +0.00 DPS) [world_drop]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Gauntlets of Ogre Strength (3341, -3.54 DPS, sim-verified) [world] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (82.8 DPS) | yes | Girdle of Golem Strength (9405, +0.00 DPS) [world_drop]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Officer's Belt (250556, -3.68 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (82.8 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Brawler's Leather Legguards (252516, +0.00 DPS) [crafted]; Veteran's Silvered Chain Leggings (250523, -3.46 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (82.8 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Alacritous Treads (277234, -3.63 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.95 DPS) | yes | Silverlaine's Family Seal (6321, -0.47 DPS) [dungeon]; Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.7 attack_power points (0.95 DPS) | yes | Silverlaine's Family Seal (6321, -0.47 DPS) [dungeon]; Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (82.8 DPS) | yes | Felbough Arm (273037, -1.35 DPS) [world]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -19.79 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; feet: Blackened Defias Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 418, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 130.8. Weights run: 1.6s. Verify run: 3.1s. 580 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.683 ± 0.034, crit=1.634 ± 0.040 per rating point (14 rating = 1%, 22.876 per %), hit=1.951 ± 0.090 per rating point (10 rating = 1%, 19.515 per %), melee_haste=10.065 ± 0.805

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 44.2 attack_power points (2.37 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.59 DPS) [crafted]; Hard Gold Coif (250537, -0.87 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 18.5 attack_power points (0.99 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Archmagister's Faceted Pendant (273049, -0.35 DPS) [world]; River Pride Choker (13087, -0.56 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.5 attack_power points (1.64 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.45 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 22.1 attack_power points (1.19 DPS) | yes | Yeti Fur Cloak (2805, -0.32 DPS) [quest]; Parachute Cloak (10518, -0.46 DPS) [crafted]; Hawkeye's Cloak (14593, -0.98 DPS, sim-verified) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (130.8 DPS) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Icemail Jerkin (1981, -6.01 DPS, sim-verified) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.7 attack_power points (1.11 DPS) | yes | Branded Leather Bracers (19508, -0.04 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.14 DPS) [world_drop]; Yorgen Bracers (13012, -0.20 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (130.8 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -1.14 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 33.1 attack_power points (1.78 DPS) | yes | Highlander's Chain Girdle (20089, -0.12 DPS) [rep]; Officer's Belt (250556, -0.16 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.17 DPS) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (130.8 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.8 attack_power points (2.03 DPS) | yes | Blackforge Greaves (6423, -0.16 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.28 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 29.5 attack_power points (1.58 DPS) | yes | Thunderbrow Ring (13097, -0.45 DPS) [world_drop]; Ring of the Underwood (2951, -0.46 DPS) [world_drop]; Mark of Kern (2262, -0.51 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 21.1 attack_power points (1.13 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.06 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (130.8 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -19.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger1: Protector's Band; finger2: Falcon's Hook

No-known-source sample (15 of 580, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 50 (human, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 186.7. Weights run: 1.7s. Verify run: 4.0s. 740 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.782 ± 0.037, crit=2.090 ± 0.048 per rating point (14 rating = 1%, 29.262 per %), hit=2.379 ± 0.104 per rating point (10 rating = 1%, 23.788 per %), melee_haste=-6.159 ± 1.250

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Lamellar Helm (220819) | Captain Dirgehammer [vendor] | sim-verified (186.7 DPS) | yes | Ornate Mithril Helm (7937, +0.00 DPS) [crafted]; Knight-Lieutenant's Plate Helm (220804, +0.00 DPS) [vendor]; Knight-Lieutenant's Imbued Helmet (220810, -7.53 DPS, sim-verified) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.2 attack_power points (1.87 DPS) | yes | Sentinel's Medallion (19539, -0.70 DPS) [rep]; Kaleidoscope Chain (13084, -1.00 DPS) [world_drop] |
| shoulder | Knight-Lieutenant's Lamellar Pauldrons (220818) | Captain Dirgehammer [vendor] | sim-verified (186.7 DPS) | yes | Knight-Lieutenant's Plate Pauldrons (220795, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, +0.00 DPS) [crafted]; Knight-Lieutenant's Imbued Pauldrons (220808, -7.54 DPS, sim-verified) [vendor] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (186.7 DPS) | yes | Pridelord Cape (14673, -0.07 DPS) [world_drop]; Blackflame Cape (13109, -0.14 DPS) [world_drop]; Dark Phantom Cape (13122, -1.60 DPS, sim-verified) [world_drop] |
| chest | Knight's Lamellar Chestplate (220815) | Captain Dirgehammer [vendor] | sim-verified (186.7 DPS) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Knight's Plate Hauberk (220794, +0.00 DPS) [vendor]; Knight's Imbued Armor (220813, -7.23 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 35.5 attack_power points (1.95 DPS) | yes | Officer's Wristguards (250581, -0.03 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.18 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.22 DPS) [crafted] |
| hands | Sergeant Major's Lamellar Gauntlets (220817) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (186.7 DPS) | yes | Raider Gauntlets (272096, +0.00 DPS) [vendor]; Raider Gloves (272100, +0.00 DPS) [vendor]; Sergeant Major's Imbued Gauntlets (220812, -6.40 DPS, sim-verified) [vendor] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 53.5 attack_power points (2.93 DPS) | yes | Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted]; Highlander's Plate Girdle (20124, -0.12 DPS) [rep]; Prowler's Leather Waistguard (252473, -3.49 DPS, sim-verified) [crafted] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | sim-verified (186.7 DPS) | yes | Gryphon Rider's Leggings (9652, +0.00 DPS) [quest]; Knight's Plate Leggings (220797, +0.00 DPS) [vendor]; Knight's Imbued Leggings (220809, -7.23 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Lamellar Boots (220814) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (186.7 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Sergeant Major's Imbued Greaves (220811, -6.91 DPS, sim-verified) [vendor] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 38.0 attack_power points (2.08 DPS) | yes | Thunderbrow Ring (13097, -0.83 DPS) [world_drop]; Falcon's Hook (7552, -0.84 DPS) [dungeon]; Ring of the Underwood (2951, -0.87 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.9 attack_power points (1.37 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.13 DPS) [dungeon]; Ring of the Underwood (2951, -0.15 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (186.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (186.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (186.7 DPS) | yes | Warmonger (13052, -0.07 DPS) [world_drop]; Stoneraven (13059, -1.62 DPS) [world_drop]; Blight (7959, -4.96 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Knight-Lieutenant's Lamellar Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Lamellar Pauldrons; chest: Knight's Lamellar Chestplate; wrist: Deepfury Bracers; hands: Sergeant Major's Lamellar Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Lamellar Legplates; feet: Sergeant Major's Lamellar Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 740, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 60 (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 273.4. Weights run: 1.7s. Verify run: 9.2s. 1396 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.319 ± 0.055, crit=3.271 ± 0.078 per rating point (14 rating = 1%, 45.793 per %), hit=3.813 ± 0.175 per rating point (10 rating = 1%, 38.133 per %), melee_haste=19.062 ± 2.054

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 207.5 attack_power points (11.66 DPS) | yes | Knight-Lieutenant's Plate Helm (220804, -5.34 DPS) [vendor]; Stalwart Helm (250599, -5.89 DPS) [crafted]; Outlaw's Collar (279253, -6.32 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 69.8 attack_power points (3.92 DPS) | yes | Amulet of the Darkmoon (19491, -0.21 DPS) [quest]; Beads of Ogre Might (22150, -0.43 DPS) [quest]; Skibi's Pendant (13089, -1.61 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (273.4 DPS) | yes | Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -2.87 DPS, sim-verified) [crafted] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.1 attack_power points (3.72 DPS) | yes | Chromatic Cloak (18509, -1.14 DPS) [crafted]; Fel Cape (279269, -1.14 DPS) [crafted]; Cape of the Black Baron (13340, -1.36 DPS, sim-verified) [dungeon] |
| chest | Soulforge Breastplate (226973) | Saving the Best for Last [quest] | sim-verified (273.4 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Devout Robe (16690, -4.07 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (273.4 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19589, +0.00 DPS) [rep] |
| hands | Soulforge Gauntlets (226975) | Just Compensation [quest] | sim-verified (273.4 DPS) | yes | Raider Gauntlets (272095, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; Devout Gloves (16692, -3.32 DPS, sim-verified) [dungeon] |
| waist | Soulforge Belt (226971) | Just Compensation [quest] | sim-verified (273.4 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Devout Belt (16696, -2.49 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 210.9 attack_power points (11.85 DPS) | yes | Sentinel's Plate Legguards (237825, -1.22 DPS) [vendor]; Titanic Leggings (22385, -1.28 DPS) [crafted]; Sentinel's Leather Pants (237818, -3.19 DPS) [vendor] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | sim-verified (273.4 DPS) | yes | Drudge Boots (21532, +0.00 DPS) [quest]; Fine Dawn Treaders (227815, +0.00 DPS) [vendor]; Knight-Lieutenant's Lamellar Sabatons (227146, -2.71 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 97.9 attack_power points (5.50 DPS) | yes | Cutthroat's Signet (272408, -1.54 DPS) [vendor]; Protector's Band (19514, -2.59 DPS) [rep]; Don Rodrigo's Band (21563, -2.93 DPS) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 93.8 attack_power points (5.27 DPS) | yes | Cutthroat's Signet (272408, -1.30 DPS) [vendor]; Protector's Band (19514, -2.35 DPS) [rep]; Don Rodrigo's Band (21563, -2.70 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (273.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (273.4 DPS) | yes | Frozen Heart of the Mountain (249469, -3.42 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (273.4 DPS) | yes | Arcanite Champion (12790, +0.00 DPS) [crafted]; Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-verified (273.4 DPS) | yes | Libram of Law (272435, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Devout Mantle; back: Howler's Furs; chest: Soulforge Breastplate; wrist: Devout Bracers; hands: Soulforge Gauntlets; waist: Soulforge Belt; legs: Sentinel's Chain Leggings; feet: Devout Sandals; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

### Band 60, raid preset (human, 55003003100100000-0000000000000000-05225331001330320)

Set DPS (verified): 624.7. Weights run: 2.0s. Verify run: 10.0s. 1396 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.024 ± 0.046, crit=2.599 ± 0.061 per rating point (14 rating = 1%, 36.384 per %), hit=5.110 ± 0.265 per rating point (10 rating = 1%, 51.103 per %), melee_haste=34.011 ± 2.990

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 218.5 attack_power points (24.07 DPS) | yes | Stalwart Helm (250599, -9.62 DPS) [crafted]; Knight-Lieutenant's Plate Helm (220804, -10.97 DPS) [vendor]; Lieutenant Commander's Lamellar Headguard (227144, -13.11 DPS) [vendor] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 75.1 attack_power points (8.27 DPS) | yes | Amulet of the Darkmoon (19491, -1.37 DPS) [quest]; Medallion of the Dawn (22659, -1.62 DPS) [quest]; Skibi's Pendant (13089, -4.04 DPS) [world_drop] |
| shoulder | Highlander's Lamellar Spaulders (20058) | The League of Arathor [rep] | sim-verified (624.7 DPS) | yes | Glowing Mantle of the Dawn (227818, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 79.1 attack_power points (8.71 DPS) | yes | Arcanoweave Cloak (272411, -3.08 DPS) [vendor]; Cape of the Black Baron (13340, -3.17 DPS) [dungeon]; Stalwart Cloak (272415, -3.70 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (624.7 DPS) | yes | Obsidian Mail Tunic (22191, -4.44 DPS) [crafted]; Dawn Armor (252483, -5.18 DPS, sim-verified) [crafted]; Shadowcraft Tunic (16721, -6.52 DPS) [dungeon] |
| wrist | Soulforge Bracers (226970) | An Earnest Proposition [quest] | sim-verified (624.7 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19589, +0.00 DPS) [rep] |
| hands | Soulforge Gauntlets (226975) | Just Compensation [quest] | sim-verified (624.7 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Raider Gauntlets (272095, -9.81 DPS, sim-verified) [vendor] |
| waist | Soulforge Belt (226971) | Just Compensation [quest] | sim-verified (624.7 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -8.50 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 211.2 attack_power points (23.26 DPS) | yes | Sentinel's Plate Legguards (237825, -2.42 DPS) [vendor]; Sentinel's Chain Leggings (237819, -3.60 DPS, sim-verified) [vendor] |
| feet | Soulforge Sabatons (226991) | Mokvar [vendor] | sim-verified (624.7 DPS) | yes | Highlander's Lamellar Greaves (20049, +0.00 DPS) [rep]; Drudge Boots (21532, +0.00 DPS) [quest]; Fine Dawn Treaders (227815, +0.00 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 109.2 attack_power points (12.03 DPS) | yes | Cutthroat's Signet (272408, -3.28 DPS) [vendor]; Protector's Band (19514, -6.38 DPS) [rep]; Channeler's Ring (272406, -6.40 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 101.5 attack_power points (11.18 DPS) | yes | Cutthroat's Signet (272408, -2.43 DPS) [vendor]; Protector's Band (19514, -5.53 DPS) [rep]; Channeler's Ring (272406, -5.55 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (624.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (624.7 DPS) | yes | Frozen Heart of the Mountain (249469, -10.05 DPS, sim-verified) [crafted] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (624.7 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; The Needler (13060, -6.94 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Highlander's Lamellar Spaulders; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Soulforge Bracers; hands: Soulforge Gauntlets; waist: Soulforge Belt; legs: Titanic Leggings; feet: Soulforge Sabatons; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: Arcanite Champion

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 40.2. Weights run: 1.6s. Verify run: 2.7s. 199 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.522 ± 0.026, crit=0.868 ± 0.024 per rating point (14 rating = 1%, 12.154 per %), hit=1.432 ± 0.065 per rating point (10 rating = 1%, 14.319 per %), melee_haste=8.491 ± 0.501

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.86 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.21 DPS) [crafted]; Brawler's Leather Hood (252504, -0.34 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 9.1 attack_power points (0.39 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.6 attack_power points (0.33 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.52 DPS, sim-verified) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.59 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) (or Death Bindings (286981)) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Death Bindings (286981, +0.00 DPS) [dungeon]; Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Forest Leather Bracers (3202, -0.10 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.77 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -1.83 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.44 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.2 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.35 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.1 attack_power points (0.61 DPS) | yes | Bounty Hunter's Ring (5351, -0.41 DPS) [quest]; The 1 Ring (8350, -0.46 DPS) [world]; Demon Band (12054, -0.57 DPS, sim-verified) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Bounty Hunter's Ring (5351, -0.06 DPS) [quest]; The 1 Ring (8350, -0.11 DPS) [world]; Demon Band (12054, -0.54 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Duskbringer (2205, -0.42 DPS) [dungeon]; Hammerbone (270018, -0.60 DPS, sim-verified) [quest]; Sword of the Fallen (286977, -0.69 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Loop of Sacrifice; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets; 18864 Insignia of the Alliance; 205218 Libram of Discovery; 209614 Insignia of the Alliance; 213513 Libram of Deliverance; 248005 Northshire Hammer

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 87.9. Weights run: 1.6s. Verify run: 1.6s. 394 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.279 ± 0.024, crit=0.960 ± 0.025 per rating point (14 rating = 1%, 13.443 per %), hit=1.559 ± 0.068 per rating point (10 rating = 1%, 15.585 per %), melee_haste=8.395 ± 0.659

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 29.5 attack_power points (1.42 DPS) | yes | Tusken Helm (6686, -0.17 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.26 DPS) [crafted]; Defender's Leather Helm (252455, -0.26 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 13.1 attack_power points (0.63 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.05 DPS) [world]; Scout's Medallion (19537, -0.14 DPS) [rep]; River Pride Choker (13087, -0.25 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (1.16 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.37 DPS) [crafted]; Golden Scale Shoulders (3841, -0.48 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.9 attack_power points (0.72 DPS) | yes | Tigerstrike Mantle (13108, -0.23 DPS) [world_drop]; Wildhunter Cloak (16658, -0.24 DPS) [quest]; Construct Cloak (279848, -0.24 DPS) [quest] |
| chest | Avenger's Armor (1488) | Razorfen Kraul [dungeon] | 30.0 attack_power points (1.44 DPS) | yes | Shining Silver Breastplate (2870, -0.10 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.11 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.29 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul [dungeon] | 16.0 attack_power points (0.77 DPS) | yes | Yorgen Bracers (13012, -0.01 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.02 DPS) [world_drop]; Barbaric Bracers (18948, -0.14 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.06 DPS) | yes | Insignia Gloves (6408, -0.02 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.04 DPS) [world_drop]; Warsong Gauntlets (16978, -0.10 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.7 attack_power points (1.33 DPS) | yes | Prowler's Leather Belt (252459, +0.00 DPS, sim-verified) [crafted]; Girdle of Golem Strength (9405, -0.18 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.18 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.9 attack_power points (1.49 DPS) | yes | Ferine Leggings (6690, -0.24 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.26 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.29 DPS) [world_drop] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 16.4 attack_power points (0.79 DPS) | yes | Veteran's Boots (250503, -0.06 DPS) [crafted]; Hard Gold Boots (250534, -0.12 DPS) [crafted]; Feet of the Lynx (1121, -0.33 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.8 attack_power points (0.95 DPS) | yes | Band of the Fist (17694, -0.32 DPS) [quest]; Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.7 attack_power points (0.95 DPS) | yes | Band of the Fist (17694, -0.32 DPS) [quest]; Tiger Band (6749, -0.47 DPS) [quest]; Bonebinder's Signet (273032, -0.47 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (87.9 DPS) | yes | Felbough Arm (273037, -1.35 DPS) [world]; Cobalt Crusher (7730, -1.91 DPS) [dungeon]; Viscous Hammer (13045, -22.10 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Brawler's Leather Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 394, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9602 Brushwood Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 15402 Noosegrip Gauntlets

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 138.4. Weights run: 1.6s. Verify run: 3.0s. 548 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.683 ± 0.034, crit=1.634 ± 0.040 per rating point (14 rating = 1%, 22.876 per %), hit=1.951 ± 0.090 per rating point (10 rating = 1%, 19.515 per %), melee_haste=10.065 ± 0.805

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Chromite Barbute (8142) | Uldaman: Ancient Treasure [dungeon] | 44.2 attack_power points (2.37 DPS) | yes | White Bandit Mask (10008, -0.20 DPS) [crafted]; Barbaric Iron Helm (7915, -0.59 DPS) [crafted]; Hard Gold Coif (250537, -0.87 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 18.5 attack_power points (0.99 DPS) | yes | Ethereal Talisman (4430, -0.10 DPS) [quest]; Kaleidoscope Chain (13084, -0.20 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.35 DPS) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 30.5 attack_power points (1.64 DPS) | yes | Forest Tracker Epaulets (2278, -0.11 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Barbaric Iron Shoulders (7913, -0.45 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.8 attack_power points (0.95 DPS) | yes | Imperial Cloak (6432, -0.23 DPS) [dungeon]; Parachute Cloak (10518, -0.23 DPS) [crafted]; Tigerstrike Mantle (13108, -0.23 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (138.4 DPS) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Icemail Jerkin (1981, -6.28 DPS, sim-verified) [world_drop] |
| wrist | Ravager's Armguards (14770) | World drop [world_drop] | 20.7 attack_power points (1.11 DPS) | yes | Branded Leather Bracers (19508, +0.00 DPS, sim-verified) [dungeon]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (138.4 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -1.37 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 33.1 attack_power points (1.78 DPS) | yes | Officer's Belt (250556, -0.16 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.17 DPS) [rep]; Defiler's Chain Girdle (20153, -0.62 DPS, sim-verified) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (138.4 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 37.8 attack_power points (2.03 DPS) | yes | Prowler's Leather Shoes (252465, -0.21 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.28 DPS) [crafted]; Blackforge Greaves (6423, -0.99 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 29.5 attack_power points (1.58 DPS) | yes | Thunderbrow Ring (13097, -0.45 DPS) [world_drop]; Ring of the Underwood (2951, -0.46 DPS) [world_drop]; Mark of Kern (2262, -0.51 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 21.1 attack_power points (1.13 DPS) | yes | Thunderbrow Ring (13097, -0.01 DPS) [world_drop]; Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Mark of Kern (2262, -0.06 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (138.4 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -24.28 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Chromite Barbute; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Stormcloth Vest; wrist: Ravager's Armguards; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Falcon's Hook

No-known-source sample (15 of 548, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (undead, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 178.5. Weights run: 1.7s. Verify run: 3.5s. 716 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.782 ± 0.037, crit=2.090 ± 0.048 per rating point (14 rating = 1%, 29.262 per %), hit=2.379 ± 0.104 per rating point (10 rating = 1%, 23.788 per %), melee_haste=-6.159 ± 1.250

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormcloth Headband (10032) | Tailoring [crafted] | sim-verified (178.5 DPS) | yes | Chromite Barbute (8142, +0.00 DPS) [dungeon]; Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Ornate Mithril Helm (7937, -6.47 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.2 attack_power points (1.87 DPS) | yes | Woven Ivy Necklace (19159, -0.27 DPS) [quest]; Scout's Medallion (19535, -0.70 DPS) [rep] |
| shoulder | Stormcloth Shoulders (10038) | Tailoring [crafted] | sim-verified (178.5 DPS) | yes | Wyrmslayer Spaulders (13066, +0.00 DPS) [world_drop]; Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -3.24 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 26.7 attack_power points (1.47 DPS) | yes | Pridelord Cape (14673, -0.22 DPS) [world_drop]; Blackflame Cape (13109, -0.29 DPS) [world_drop]; Bloodlust Cape (14801, -0.38 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (178.5 DPS) | yes | Warbear Harness (15064, +0.00 DPS) [crafted]; Stone Guard's Plate Armor (220801, +0.00 DPS) [vendor]; Green Dragonscale Breastplate (15045, -4.28 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 35.5 attack_power points (1.95 DPS) | yes | Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Officer's Wristguards (250581, +0.00 DPS) [crafted] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (178.5 DPS) | yes | Officer's Gloves (250551, +0.00 DPS) [crafted]; Raider Gloves (272100, +0.00 DPS) [vendor]; Raider Gauntlets (272096, -6.04 DPS, sim-verified) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 52.2 attack_power points (2.86 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Defiler's Plate Girdle (20205, -0.05 DPS) [rep]; Defiler's Chain Girdle (20151, -0.16 DPS) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (178.5 DPS) | yes | Stormshroud Pants (15057, +0.00 DPS) [crafted]; Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor]; Green Dragonscale Leggings (15046, -1.40 DPS, sim-verified) [crafted] |
| feet | Stormcloth Boots (10039) | Tailoring [crafted] | sim-verified (178.5 DPS) | yes | Officer's Sabatons (250561, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -3.30 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 38.0 attack_power points (2.08 DPS) | yes | White Bone Band (11862, -0.77 DPS) [quest]; Thunderbrow Ring (13097, -0.83 DPS) [world_drop]; Falcon's Hook (7552, -0.84 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.9 attack_power points (1.37 DPS) | yes | White Bone Band (11862, -0.05 DPS) [quest]; Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.13 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (178.5 DPS) | yes | Molten Heart of the Mountain (249470, -3.20 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (178.5 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (178.5 DPS) | yes | Warmonger (13052, -0.07 DPS) [world_drop]; Stoneraven (13059, -1.62 DPS) [world_drop]; Blight (7959, -4.37 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Stormcloth Headband; neck: Skibi's Pendant; shoulder: Stormcloth Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Prowler's Leather Waistguard; feet: Stormcloth Boots; finger1: Legionnaire's Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 716, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 296.8. Weights run: 1.7s. Verify run: 9.5s. 1418 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=2.319 ± 0.055, crit=3.271 ± 0.078 per rating point (14 rating = 1%, 45.793 per %), hit=3.813 ± 0.175 per rating point (10 rating = 1%, 38.133 per %), melee_haste=19.062 ± 2.054

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 207.5 attack_power points (11.66 DPS) | yes | Blood Guard's Plate Helm (220803, -5.34 DPS) [vendor]; Stalwart Helm (250599, -5.89 DPS) [crafted]; Outlaw's Collar (279253, -6.74 DPS, sim-verified) [crafted] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 69.8 attack_power points (3.92 DPS) | yes | Amulet of the Darkmoon (19491, -0.21 DPS) [quest]; Beads of Ogre Might (22150, -0.43 DPS) [quest]; Skibi's Pendant (13089, -1.61 DPS) [world_drop] |
| shoulder | Defiler's Epaulets (20176) | The Defilers [rep] | sim-verified (296.8 DPS) | yes | Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Devout Mantle (16695, -4.64 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.1 attack_power points (3.72 DPS) | yes | Chromatic Cloak (18509, -1.14 DPS) [crafted]; Fel Cape (279269, -1.14 DPS) [crafted]; Cape of the Black Baron (13340, -2.43 DPS, sim-verified) [dungeon] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (296.8 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Magister's Robes (16688, -7.10 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (296.8 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Magister's Bindings (16683, -8.12 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (296.8 DPS) | yes | Raider Gauntlets (272095, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; Magister's Gloves (16684, -6.80 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-verified (296.8 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Devout Belt (16696, -6.17 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 210.9 attack_power points (11.85 DPS) | yes | Sentinel's Plate Legguards (237825, -1.22 DPS) [vendor]; Titanic Leggings (22385, -1.28 DPS) [crafted]; Outrider's Chain Leggings (22673, -2.57 DPS) [rep] |
| feet | Defiler's Cloth Boots (20159) | The Defilers [rep] | sim-verified (296.8 DPS) | yes | Drudge Boots (21532, +0.00 DPS) [quest]; Fine Dawn Treaders (227815, +0.00 DPS) [vendor]; Devout Sandals (16691, -5.46 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 97.9 attack_power points (5.50 DPS) | yes | Cutthroat's Signet (272408, -1.54 DPS) [vendor]; Legionnaire's Band (19510, -2.59 DPS) [rep]; Don Rodrigo's Band (21563, -2.93 DPS) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 93.8 attack_power points (5.27 DPS) | yes | Cutthroat's Signet (272408, -1.30 DPS) [vendor]; Legionnaire's Band (19510, -2.35 DPS) [rep]; Don Rodrigo's Band (21563, -2.70 DPS) [rep] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (296.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (296.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (296.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; The Needler (13060, -8.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-verified (296.8 DPS) | yes | Libram of Fervor (23203, -1.84 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Medallion of the Dawn; shoulder: Defiler's Epaulets; back: Howler's Furs; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Defiler's Cloth Girdle; legs: Sentinel's Chain Leggings; feet: Defiler's Cloth Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: Hand of Edward the Odd; ranged: Libram of Law

No-known-source sample (15 of 1418, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (undead, 55003003100100000-0000000000000000-05225331001330320)

Set DPS (verified): 626.0. Weights run: 2.0s. Verify run: 9.3s. 1418 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.005, agility=2.024 ± 0.046, crit=2.599 ± 0.061 per rating point (14 rating = 1%, 36.384 per %), hit=5.110 ± 0.265 per rating point (10 rating = 1%, 51.103 per %), melee_haste=34.011 ± 2.990

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 218.5 attack_power points (24.07 DPS) | yes | Stalwart Helm (250599, -9.62 DPS) [crafted]; Blood Guard's Plate Helm (220803, -10.97 DPS) [vendor]; Outlaw's Collar (279253, -13.38 DPS) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 75.1 attack_power points (8.27 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -1.62 DPS) [quest]; Skibi's Pendant (13089, -4.04 DPS) [world_drop] |
| shoulder | Pauldrons of Elements (16669) | Blackrock Spire: Gyth [dungeon] | sim-verified (626.0 DPS) | yes | Glowing Mantle of the Dawn (227818, +0.00 DPS) [vendor]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -3.73 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 79.1 attack_power points (8.71 DPS) | yes | Arcanoweave Cloak (272411, -3.08 DPS) [vendor]; Cape of the Black Baron (13340, -3.17 DPS) [dungeon]; Stalwart Cloak (272415, -3.37 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (626.0 DPS) | yes | Obsidian Mail Tunic (22191, -4.44 DPS) [crafted]; Dawn Armor (252483, -5.55 DPS, sim-verified) [crafted]; Shadowcraft Tunic (16721, -6.52 DPS) [dungeon] |
| wrist | Bindings of Elements (16671) | Stratholme: Crypt Beast [dungeon] | sim-verified (626.0 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19589, +0.00 DPS) [rep] |
| hands | Gauntlets of Elements (16672) | Blackrock Spire: Pyroguard Emberseer [dungeon] | sim-verified (626.0 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Raider Gauntlets (272095, -3.20 DPS, sim-verified) [vendor] |
| waist | Cord of Elements (16673) | Blackrock Spire: Smolderthorn Witch Doctor [dungeon] | sim-verified (626.0 DPS) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -7.26 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | 211.2 attack_power points (23.26 DPS) | yes | Sentinel's Chain Leggings (237819, +0.00 DPS, sim-verified) [vendor]; Sentinel's Plate Legguards (237825, -2.42 DPS) [vendor]; Outrider's Chain Leggings (22673, -5.83 DPS) [rep] |
| feet | Boots of Elements (16670) | Blackrock Spire: Highlord Omokk [dungeon] | sim-verified (626.0 DPS) | yes | Defiler's Plate Greaves (20208, +0.00 DPS) [rep]; Drudge Boots (21532, +0.00 DPS) [quest]; Fine Dawn Treaders (227815, -9.39 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 109.2 attack_power points (12.03 DPS) | yes | Cutthroat's Signet (272408, -3.28 DPS) [vendor]; Legionnaire's Band (19510, -6.38 DPS) [rep]; Channeler's Ring (272406, -6.40 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 101.5 attack_power points (11.18 DPS) | yes | Cutthroat's Signet (272408, -3.12 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -5.53 DPS) [rep]; Channeler's Ring (272406, -5.55 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (626.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Maelstrom (19289, -16.82 DPS, sim-verified) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (626.0 DPS) | yes | Frozen Heart of the Mountain (249469, -3.50 DPS) [crafted]; Darkmoon Card: Maelstrom (19289, -4.36 DPS, sim-verified) [quest] |
| main_hand | Arcanite Champion (12790) | Blacksmithing [crafted] | sim-verified (626.0 DPS) | yes | The Needler (13060, +0.00 DPS) [world_drop]; High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Beads of Ogre Might; shoulder: Pauldrons of Elements; back: Howler's Furs; chest: Timbermaw Tunic; wrist: Bindings of Elements; hands: Gauntlets of Elements; waist: Cord of Elements; legs: Titanic Leggings; feet: Boots of Elements; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Rune of the Guard Captain; main_hand: Arcanite Champion

No-known-source sample (15 of 1418, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4116 Olmann Sewar; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

