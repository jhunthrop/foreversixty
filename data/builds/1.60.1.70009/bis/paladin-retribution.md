# Leveling BiS: Retribution

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 40.7. Weights run: 1.3s. Verify run: 2.1s. 239 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.542 ± 0.026, crit=0.879 ± 0.024 per rating point (14 rating = 1%, 12.309 per %), hit=1.449 ± 0.066 per rating point (10 rating = 1%, 14.490 per %), melee_haste=8.598 ± 0.508

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.86 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.20 DPS) [crafted]; Brawler's Leather Hood (252504, -0.33 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 9.2 attack_power points (0.40 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.7 attack_power points (0.33 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 9.2 attack_power points (0.40 DPS) | yes | Grave Shroud (279865, -0.01 DPS) [quest]; Dark Leather Cloak (2316, -0.03 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (40.7 DPS) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.88 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.2 attack_power points (0.44 DPS) | yes | Patterned Bronze Bracers (2868, +0.00 DPS, sim-verified) [crafted]; Bristlebark Bindings (14569, -0.07 DPS) [world_drop]; Cryptwalker Bracers (280095, -0.09 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.05 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (40.7 DPS) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -2.19 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.01 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.36 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.2 attack_power points (0.61 DPS) | yes | Demon Band (12054, -0.27 DPS) [world_drop]; The 1 Ring (8350, -0.46 DPS) [world]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.2 attack_power points (0.40 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; The 1 Ring (8350, -0.25 DPS) [world]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Verigan's Fist (6953) | The Test of Righteousness [quest] | 358.7 attack_power points (15.49 DPS) | yes | Smite's Mighty Hammer (7230, -2.68 DPS, sim-verified) [dungeon]; Duskbringer (2205, -2.78 DPS) [dungeon]; Monstrous Cleaver (279864, -3.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Verigan's Fist

No-known-source sample (15 of 239, see the JSON for more): 1189 Overseer's Ring; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler

### Band 30 (human, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 84.8. Weights run: 1.3s. Verify run: 1.5s. 404 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.186 ± 0.022, crit=0.890 ± 0.024 per rating point (14 rating = 1%, 12.463 per %), hit=1.441 ± 0.063 per rating point (10 rating = 1%, 14.410 per %), melee_haste=7.783 ± 0.615

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 28.7 attack_power points (1.54 DPS) | yes | Tusken Helm (6686, -0.14 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.25 DPS) [crafted]; Defender's Leather Helm (252455, -0.25 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.75 DPS) | yes | Kaleidoscope Chain (13084, +0.00 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.24 DPS) [rep]; River Pride Choker (13087, -0.32 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.0 attack_power points (1.23 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.38 DPS) [crafted]; Golden Scale Shoulders (3841, -0.48 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.3 attack_power points (0.77 DPS) | yes | Sergeant Major's Cape (16315, +0.00 DPS, sim-verified) [pvp]; Wolfmaster Cape (6314, -0.23 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.61 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.15 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.86 DPS) | yes | Yorgen Bracers (13012, -0.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.18 DPS) | yes | Insignia Gloves (6408, -0.05 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; The Frozen Clutch (23170, -0.11 DPS) [dungeon] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.1 attack_power points (1.45 DPS) | yes | Prowler's Leather Belt (252459, -0.11 DPS) [crafted]; Girdle of Golem Strength (9405, -0.17 DPS) [world_drop]; Highlander's Plate Girdle (20126, -0.17 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.3 attack_power points (1.62 DPS) | yes | Ferine Leggings (6690, +0.00 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.30 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.32 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 22.3 attack_power points (1.19 DPS) | yes | Brawler's Leather Boots (252439, -0.34 DPS) [crafted]; Feet of the Lynx (1121, -0.36 DPS) [world_drop]; Alacritous Treads (277234, -0.49 DPS, sim-verified) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.6 attack_power points (1.05 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Tiger Band (6749, -0.40 DPS) [quest]; Silverlaine's Family Seal (6321, -0.51 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 19.1 attack_power points (1.02 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Tiger Band (6749, -0.38 DPS) [quest]; Silverlaine's Family Seal (6321, -0.49 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.8 DPS) | yes | Morbid Dawn (7689, -0.08 DPS) [dungeon]; Cobalt Crusher (7730, -2.10 DPS) [dungeon]; Viscous Hammer (13045, -19.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 404, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak

### Band 40 (human, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 137.3. Weights run: 1.3s. Verify run: 2.5s. 562 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.433 ± 0.028, crit=1.392 ± 0.033 per rating point (14 rating = 1%, 19.491 per %), hit=1.788 ± 0.080 per rating point (10 rating = 1%, 17.884 per %), melee_haste=9.278 ± 0.718

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 45.5 attack_power points (2.65 DPS) | yes | Chromite Barbute (8142, -0.25 DPS) [dungeon]; White Bandit Mask (10008, -0.45 DPS) [crafted]; Barbaric Iron Helm (7915, -0.85 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.16 DPS) | yes | Ghostshard Talisman (7731, -0.35 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop]; Sentinel's Medallion (19540, -0.54 DPS, sim-verified) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.8 attack_power points (1.62 DPS) | yes | Forest Tracker Epaulets (2278, -0.12 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.34 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | sim-verified (137.3 DPS) | yes | Dark Hooded Cape (5257, +0.00 DPS) [world]; Hawkeye's Cloak (14593, -0.27 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.35 DPS) [quest] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (137.3 DPS) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Quillward Harness (10583, -5.13 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.16 DPS) | yes | Ravager's Armguards (14770, -0.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.20 DPS) [world_drop]; Yorgen Bracers (13012, -0.22 DPS) [world_drop] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (137.3 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -0.96 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 30.9 attack_power points (1.80 DPS) | yes | Highlander's Leather Girdle (20116, -0.05 DPS) [rep]; Boar Champion's Belt (10768, -0.05 DPS) [dungeon]; Highlander's Plate Girdle (20125, -0.05 DPS) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (137.3 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 36.0 attack_power points (2.10 DPS) | yes | Blackforge Greaves (6423, -0.22 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.36 DPS) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 27.5 attack_power points (1.60 DPS) | yes | Thunderbrow Ring (13097, -0.42 DPS) [world_drop]; Mark of Kern (2262, -0.43 DPS) [dungeon]; Assault Band (13095, -0.43 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 20.9 attack_power points (1.22 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (137.3 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -22.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Stormcloth Vest; wrist: Branded Leather Bracers; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger1: Protector's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 562, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 50 (human, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 192.5. Weights run: 1.3s. Verify run: 3.1s. 722 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.617 ± 0.035, crit=1.914 ± 0.045 per rating point (14 rating = 1%, 26.798 per %), hit=2.201 ± 0.105 per rating point (10 rating = 1%, 22.014 per %), melee_haste=11.514 ± 1.107

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Lamellar Helm (220819) | Captain Dirgehammer [vendor] | sim-verified (192.5 DPS) | yes | Raging Berserker's Helm (7719, +0.00 DPS) [dungeon]; Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Knight-Lieutenant's Plate Helm (220804, -6.83 DPS, sim-verified) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.0 attack_power points (1.90 DPS) | yes | Sentinel's Medallion (19539, -0.75 DPS) [rep]; Zealous Shadowshard Pendant (17772, -3.11 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Lamellar Pauldrons (220818) | Captain Dirgehammer [vendor] | sim-verified (192.5 DPS) | yes | Wyrmslayer Spaulders (13066, +0.00 DPS) [world_drop]; Officer's Pauldrons (250576, +0.00 DPS) [crafted]; Knight-Lieutenant's Plate Pauldrons (220795, -8.95 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 35.8 attack_power points (2.12 DPS) | yes | Dark Hooded Cape (5257, -0.64 DPS) [world]; Blisterbane Wrap (12552, -0.69 DPS) [dungeon]; Dark Phantom Cape (13122, -0.69 DPS) [world_drop] |
| chest | Knight's Lamellar Chestplate (220815) | Captain Dirgehammer [vendor] | sim-verified (192.5 DPS) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Knight's Plate Hauberk (220794, +0.00 DPS) [vendor]; Green Dragonscale Breastplate (15045, -7.50 DPS, sim-verified) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 33.9 attack_power points (2.01 DPS) | yes | Deepfury Bracers (13120, -0.05 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.16 DPS) [crafted]; Runed Golem Shackles (12550, -0.18 DPS) [dungeon] |
| hands | Sergeant Major's Lamellar Gauntlets (220817) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (192.5 DPS) | yes | Officer's Gloves (250551, +0.00 DPS) [crafted]; Raider Gloves (272100, +0.00 DPS) [vendor]; Raider Gauntlets (272096, -4.12 DPS, sim-verified) [vendor] |
| waist | Highlander's Lamellar Girdle (20106) | The League of Arathor [rep] | 51.0 attack_power points (3.02 DPS) | yes | Skulker's Leather Waistguard (252474, -0.12 DPS) [crafted]; Highlander's Plate Girdle (20124, -0.13 DPS) [rep]; Prowler's Leather Waistguard (252473, -4.42 DPS, sim-verified) [crafted] |
| legs | Knight's Lamellar Legplates (220816) | Captain Dirgehammer [vendor] | sim-verified (192.5 DPS) | yes | Gryphon Rider's Leggings (9652, +0.00 DPS) [quest]; Knight's Plate Leggings (220797, +0.00 DPS) [vendor]; Green Dragonscale Leggings (15046, -7.50 DPS, sim-verified) [crafted] |
| feet | Sergeant Major's Lamellar Boots (220814) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (192.5 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Battlechaser's Greaves (12555, -7.72 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.0 attack_power points (2.49 DPS) | yes | Ironspine's Eye (7686, -1.11 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.15 DPS) [quest]; Thunderbrow Ring (13097, -1.16 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 36.5 attack_power points (2.16 DPS) | yes | Ironspine's Eye (7686, -0.78 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.82 DPS) [quest]; Thunderbrow Ring (13097, -0.83 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (192.5 DPS) | yes | - |
| trinket2 | Sanctified Orb (20512) | Forging the Mightstone [quest] | sim-verified (192.5 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (192.5 DPS) | yes | Warmonger (13052, -0.39 DPS) [world_drop]; Taran Icebreaker (2915, -0.90 DPS) [world_drop]; Blight (7959, -5.83 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Invocation (249442) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Knight-Lieutenant's Lamellar Helm; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Lamellar Pauldrons; back: Blackveil Cape; chest: Knight's Lamellar Chestplate; wrist: Officer's Wristguards; hands: Sergeant Major's Lamellar Gauntlets; waist: Highlander's Lamellar Girdle; legs: Knight's Lamellar Legplates; feet: Sergeant Major's Lamellar Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Sanctified Orb; main_hand: Thorium Greatmace

No-known-source sample (15 of 722, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60 (human, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 285.0. Weights run: 1.3s. Verify run: 8.1s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.904 ± 0.046, crit=2.694 ± 0.065 per rating point (14 rating = 1%, 37.717 per %), hit=3.161 ± 0.158 per rating point (10 rating = 1%, 31.610 per %), melee_haste=19.759 ± 1.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 178.3 attack_power points (11.20 DPS) | yes | Mask of the Unforgiven (13404, -4.86 DPS) [dungeon]; Knight-Lieutenant's Plate Helm (220804, -5.05 DPS) [vendor]; Eye of Rend (12587, -5.56 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 63.7 attack_power points (4.00 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Pendant of Celerity (22340, -0.22 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.35 DPS) [quest] |
| shoulder | Soulforge Pauldrons (226987) | Mokvar [vendor] | sim-verified (285.0 DPS) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -6.83 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.6 attack_power points (3.74 DPS) | yes | Windshear Cape (20691, -0.84 DPS) [world]; Cloak of the Honor Guard (20073, -1.01 DPS) [rep]; Cape of the Black Baron (13340, -1.49 DPS, sim-verified) [dungeon] |
| chest | Soulforge Breastplate (226973) | Saving the Best for Last [quest] | sim-verified (285.0 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; The Postmaster's Tunic (13388, -11.08 DPS, sim-verified) [dungeon] |
| wrist | Soulforge Bracers (226970) | An Earnest Proposition [quest] | sim-verified (285.0 DPS) | yes | Berserker Bracers (19578, +0.00 DPS) [rep]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -10.17 DPS, sim-verified) [dungeon] |
| hands | Soulforge Gauntlets (226975) | Just Compensation [quest] | sim-verified (285.0 DPS) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Savage Gladiator Grips (11730, -11.56 DPS, sim-verified) [dungeon] |
| waist | Soulforge Belt (226971) | Just Compensation [quest] | sim-verified (285.0 DPS) | yes | Belt of Preserved Heads (20216, +0.00 DPS) [quest]; Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, -11.04 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (285.0 DPS) | yes | Titanic Leggings (22385, -0.42 DPS) [crafted]; Sentinel's Plate Legguards (237825, -0.46 DPS) [vendor]; Cloudkeeper Legplates (14554, -9.04 DPS, sim-verified) [world_drop] |
| feet | Soulforge Sabatons (226991) | Mokvar [vendor] | sim-verified (285.0 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; The Postmaster's Treads (13391, -2.26 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-verified (285.0 DPS) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Don Julio's Band (19325, +0.00 DPS, sim-verified) [rep]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (285.0 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Tarnished Elven Ring (18500, -0.36 DPS) [dungeon]; Naglering (11669, -4.91 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (285.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (285.0 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Draconic Infused Emblem (22268, -2.12 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (285.0 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-verified (285.0 DPS) | yes | Libram of Fervor (23203, -1.53 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Soulforge Pauldrons; back: Howler's Furs; chest: Soulforge Breastplate; wrist: Soulforge Bracers; hands: Soulforge Gauntlets; waist: Soulforge Belt; legs: Sentinel's Chain Leggings; feet: Soulforge Sabatons; finger1: The Postmaster's Seal; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Burst of Knowledge; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

### Band 60, raid preset (human, 55003003100100000-0000000000000000-05225331001330320)

Set DPS (verified): 666.2. Weights run: 1.6s. Verify run: 9.0s. 1674 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=2.173 ± 0.049, crit=2.800 ± 0.067 per rating point (14 rating = 1%, 39.197 per %), hit=5.127 ± 0.294 per rating point (10 rating = 1%, 51.269 per %), melee_haste=17.566 ± 2.436

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 224.5 attack_power points (27.83 DPS) | yes | Helm of the Executioner (22411, -10.92 DPS) [dungeon]; Mask of the Unforgiven (13404, -11.01 DPS, sim-verified) [dungeon]; Stalwart Helm (250599, -11.52 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 83.9 attack_power points (10.40 DPS) | yes | Beads of Ogre Might (22150, -1.07 DPS) [quest]; Amulet of the Darkmoon (19491, -2.28 DPS) [quest]; Mark of Fordring (15411, -2.31 DPS) [quest] |
| shoulder | Soulforge Pauldrons (226987) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -6.61 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 79.3 attack_power points (9.83 DPS) | yes | Cape of the Black Baron (13340, -3.31 DPS) [dungeon]; Windshear Cape (20691, -3.39 DPS) [world]; Stalwart Cloak (272415, -3.47 DPS) [vendor] |
| chest | Soulforge Breastplate (226973) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; The Postmaster's Tunic (13388, -17.26 DPS, sim-verified) [dungeon] |
| wrist | Soulforge Bracers (226970) | An Earnest Proposition [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Battleborn Armbraces (12936, -11.08 DPS, sim-verified) [dungeon] |
| hands | Soulforge Gauntlets (226975) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Voone's Vice Grips (13963, +0.00 DPS) [quest]; Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Savage Gladiator Grips (11730, -17.02 DPS, sim-verified) [dungeon] |
| waist | Soulforge Belt (226971) | Just Compensation [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ferocity of the Timbermaw (227805, +0.00 DPS) [vendor]; Radiant Girdle of the Dawn (227814, +0.00 DPS) [vendor]; Belt of Preserved Heads (20216, -11.07 DPS, sim-verified) [quest] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentinel's Chain Leggings (237819, -1.07 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.40 DPS) [vendor]; Cloudkeeper Legplates (14554, -14.54 DPS, sim-verified) [world_drop] |
| feet | Soulforge Sabatons (226991) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; The Postmaster's Treads (13391, -6.23 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (666.2 DPS) | yes | Tarnished Elven Ring (18500, -2.80 DPS) [dungeon]; Cutthroat's Signet (272408, -3.07 DPS) [vendor]; The Postmaster's Seal (13392, -10.44 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.81 DPS) [dungeon]; Cutthroat's Signet (272408, -1.08 DPS) [vendor]; Naglering (11669, -9.03 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Draconic Infused Emblem (22268, -7.78 DPS, sim-verified) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -5.46 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Soulforge Pauldrons; back: Howler's Furs; chest: Soulforge Breastplate; wrist: Soulforge Bracers; hands: Soulforge Gauntlets; waist: Soulforge Belt; legs: Titanic Leggings; feet: Soulforge Sabatons; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1674, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4765 Enamelled Broadsword; 4777 Ironwood Maul

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000-05024000000000000)

Set DPS (verified): 40.2. Weights run: 1.3s. Verify run: 2.2s. 219 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.542 ± 0.026, crit=0.879 ± 0.024 per rating point (14 rating = 1%, 12.309 per %), hit=1.449 ± 0.066 per rating point (10 rating = 1%, 14.490 per %), melee_haste=8.598 ± 0.508

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Veteran's Silvered Chain Helm (250528) | Blacksmithing [crafted] | 20.0 attack_power points (0.86 DPS) | yes | Defender's Leather Hood (252447, -0.17 DPS) [crafted]; Guard's Silvered Chain Helm (250529, -0.20 DPS) [crafted]; Brawler's Leather Hood (252504, -0.33 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 9.2 attack_power points (0.40 DPS) | yes | Erudite's Amulet (277204, -0.13 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.7 attack_power points (0.33 DPS) | yes | Rough Bronze Shoulders (3480, -0.07 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.07 DPS) [crafted] |
| back | Grave Shroud (279865) | Unending Torment [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dark Leather Cloak (2316, -0.02 DPS) [crafted]; Lambent Scale Cloak (4706, -0.05 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.46 DPS, sim-verified) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mutant Scale Breastplate (6627, +0.00 DPS) [dungeon]; Veteran's Chain Shirt (250488, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.65 DPS, sim-verified) [crafted] |
| wrist | Patterned Bronze Bracers (2868) | Blacksmithing [crafted] | 10.0 attack_power points (0.43 DPS) | yes | Bristlebark Bindings (14569, -0.06 DPS) [world_drop]; Forest Leather Bracers (3202, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.17 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dagmire Gauntlets (6481, +0.00 DPS) [quest]; Thorbia's Gauntlets (12994, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -1.81 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Brawler's Leather Belt (252428, -0.17 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Cobrahn's Grasp (6460, -1.89 DPS, sim-verified) [dungeon] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Veteran's Chain Leggings (250493, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.49 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.3 attack_power points (0.79 DPS) | yes | Brawler's Leather Boots (252439, -0.03 DPS) [crafted]; Veteran's Boots (250503, -0.09 DPS) [crafted]; Defender's Leather Boots (252441, -0.36 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.2 attack_power points (0.61 DPS) | yes | Demon Band (12054, -0.27 DPS) [world_drop]; Loop of Sacrifice (281673, -0.35 DPS) [quest]; Bounty Hunter's Ring (5351, -0.41 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 9.2 attack_power points (0.40 DPS) | yes | Demon Band (12054, -0.05 DPS) [world_drop]; Loop of Sacrifice (281673, -0.14 DPS) [quest]; Bounty Hunter's Ring (5351, -0.20 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Duskbringer (2205, -0.42 DPS) [dungeon]; Hammerbone (270018, -0.56 DPS, sim-verified) [quest]; Monstrous Cleaver (279864, -0.79 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Veteran's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Patterned Bronze Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2273 Guerrilla Armor; 2543 Militia Pants; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5968 Rugged Boots; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 9602 Brushwood Blade

### Band 30 (undead, 00000000000000000-0000000000000000-05025331001100000)

Set DPS (verified): 91.7. Weights run: 1.3s. Verify run: 1.4s. 381 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.003, agility=1.186 ± 0.022, crit=0.890 ± 0.024 per rating point (14 rating = 1%, 12.463 per %), hit=1.441 ± 0.063 per rating point (10 rating = 1%, 14.410 per %), melee_haste=7.783 ± 0.615

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Barbaric Iron Helm (7915) | Blacksmithing [crafted] | 28.7 attack_power points (1.54 DPS) | yes | Tusken Helm (6686, -0.14 DPS) [dungeon]; Veteran's Chain Helm (250498, -0.25 DPS) [crafted]; Defender's Leather Helm (252455, -0.25 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.75 DPS) | yes | Kaleidoscope Chain (13084, -0.07 DPS) [world_drop]; Scout's Medallion (19537, -0.24 DPS) [rep]; River Pride Choker (13087, -0.32 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.0 attack_power points (1.23 DPS) | yes | Barbaric Iron Shoulders (7913, -0.21 DPS) [crafted]; Barbaric Shoulders (5964, -0.38 DPS) [crafted]; Golden Scale Shoulders (3841, -0.48 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.3 attack_power points (0.77 DPS) | yes | Wolfmaster Cape (6314, -0.23 DPS) [dungeon]; Wildhunter Cloak (16658, -0.23 DPS) [quest]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop] |
| chest | Avenger's Armor (1488) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 30.0 attack_power points (1.61 DPS) | yes | Shining Silver Breastplate (2870, -0.11 DPS) [crafted]; Veteran's Silvered Chain Shirt (250518, -0.15 DPS) [crafted]; Barbaric Iron Breastplate (7914, -0.32 DPS) [crafted] |
| wrist | Pugilist Bracers (4438) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 16.0 attack_power points (0.86 DPS) | yes | Yorgen Bracers (13012, -0.02 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Gauntlets of Ogre Strength (3341) | Boulderfist Enforcer [world] | 22.0 attack_power points (1.18 DPS) | yes | Insignia Gloves (6408, -0.05 DPS) [world_drop]; Mail Combat Gauntlets (4075, -0.07 DPS) [world_drop]; Warsong Gauntlets (16978, -0.11 DPS) [quest] |
| waist | Officer's Belt (250556) | Blacksmithing [crafted] | 27.1 attack_power points (1.45 DPS) | yes | Prowler's Leather Belt (252459, +0.00 DPS, sim-verified) [crafted]; Girdle of Golem Strength (9405, -0.17 DPS) [world_drop]; Defiler's Plate Girdle (20207, -0.17 DPS) [rep] |
| legs | Veteran's Silvered Chain Leggings (250523) | Blacksmithing [crafted] | 30.3 attack_power points (1.62 DPS) | yes | Ferine Leggings (6690, -0.23 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.30 DPS) [crafted]; Glimmering Mail Legguards (6386, -0.32 DPS) [world_drop] |
| feet | Trouncing Boots (4464) | Commander Felstrom [world] | 22.3 attack_power points (1.19 DPS) | yes | Brawler's Leather Boots (252439, -0.34 DPS) [crafted]; Feet of the Lynx (1121, -0.36 DPS) [world_drop]; Veteran's Boots (250503, -0.40 DPS) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 19.6 attack_power points (1.05 DPS) | yes | Ironspine's Eye (7686, -0.05 DPS) [dungeon]; Band of the Fist (17694, -0.36 DPS) [quest]; Tiger Band (6749, -0.40 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 19.1 attack_power points (1.02 DPS) | yes | Ironspine's Eye (7686, -0.33 DPS, sim-verified) [dungeon]; Band of the Fist (17694, -0.34 DPS) [quest]; Tiger Band (6749, -0.38 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (91.7 DPS) | yes | Morbid Dawn (7689, -0.08 DPS) [dungeon]; Cobalt Crusher (7730, -2.10 DPS) [dungeon]; Viscous Hammer (13045, -22.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Barbaric Iron Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Avenger's Armor; wrist: Pugilist Bracers; hands: Gauntlets of Ogre Strength; waist: Officer's Belt; legs: Veteran's Silvered Chain Leggings; feet: Trouncing Boots; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 381, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 40 (undead, 00000000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 147.5. Weights run: 1.3s. Verify run: 2.6s. 531 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.000 ± 0.004, agility=1.433 ± 0.028, crit=1.392 ± 0.033 per rating point (14 rating = 1%, 19.491 per %), hit=1.788 ± 0.080 per rating point (10 rating = 1%, 17.884 per %), melee_haste=9.278 ± 0.718

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 45.5 attack_power points (2.65 DPS) | yes | White Bandit Mask (10008, -0.45 DPS) [crafted]; Barbaric Iron Helm (7915, -0.85 DPS) [crafted]; Chromite Barbute (8142, -0.92 DPS, sim-verified) [dungeon] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.16 DPS) | yes | Scout's Medallion (19536, -0.25 DPS) [rep]; Ethereal Talisman (4430, -0.25 DPS) [quest]; Ghostshard Talisman (7731, -0.35 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.8 attack_power points (1.62 DPS) | yes | Forest Tracker Epaulets (2278, -0.12 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Hard Gold Pauldrons (250539, -0.34 DPS) [crafted] |
| back | First Sergeant's Cloak (16340) | PvP rank 9 · First Sergeant · Horde [pvp] | sim-verified (147.5 DPS) | yes | Hawkeye's Cloak (14593, -0.27 DPS) [world_drop]; Parachute Cloak (10518, -0.53 DPS) [crafted]; Dark Hooded Cape (5257, -1.98 DPS, sim-verified) [world] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (147.5 DPS) | yes | Kolkar Marauder Chain (6773, +0.00 DPS) [quest]; Carapace of Tuten'kash (10775, +0.00 DPS) [dungeon]; Quillward Harness (10583, -5.64 DPS, sim-verified) [dungeon] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.16 DPS) | yes | Ravager's Armguards (14770, +0.00 DPS) [world_drop]; Berserker Bracers (19581, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19590, +0.00 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (147.5 DPS) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Plated Fist of Hakoo (13071, +0.00 DPS) [world_drop]; Gloves of Holy Might (867, -2.40 DPS, sim-verified) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 30.9 attack_power points (1.80 DPS) | yes | Defiler's Leather Girdle (20192, -0.05 DPS) [rep]; Boar Champion's Belt (10768, -0.05 DPS) [dungeon]; Defiler's Plate Girdle (20206, -0.05 DPS) [rep] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (147.5 DPS) | yes | Scarlet Leggings (10330, +0.00 DPS, sim-verified) [dungeon]; Firemane Leggings (13129, +0.00 DPS) [world_drop]; Symbolic Legplates (14829, +0.00 DPS) [world_drop] |
| feet | Officer's Boots (250546) | Blacksmithing [crafted] | 36.0 attack_power points (2.10 DPS) | yes | Prowler's Leather Shoes (252465, -0.23 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.36 DPS) [crafted]; Blackforge Greaves (6423, -0.99 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 27.5 attack_power points (1.60 DPS) | yes | Thunderbrow Ring (13097, -0.42 DPS) [world_drop]; Mark of Kern (2262, -0.43 DPS) [dungeon]; Assault Band (13095, -0.43 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 20.9 attack_power points (1.22 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Mark of Kern (2262, -0.05 DPS) [dungeon]; Assault Band (13095, -0.05 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (147.5 DPS) | yes | X'caliboar (10758, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Nightblade (1982, -27.46 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Raging Berserker's Helm; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: First Sergeant's Cloak; chest: Stormcloth Vest; wrist: Branded Leather Bracers; hands: Stormcloth Gloves; waist: Ogron's Sash; legs: Stormcloth Pants; feet: Officer's Boots; finger1: Legionnaire's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 531, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band

### Band 50 (undead, 52003000000000000-0000000000000000-05025331001330311)

Set DPS (verified): 183.9. Weights run: 1.3s. Verify run: 2.8s. 702 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.617 ± 0.035, crit=1.914 ± 0.045 per rating point (14 rating = 1%, 26.798 per %), hit=2.201 ± 0.105 per rating point (10 rating = 1%, 22.014 per %), melee_haste=11.514 ± 1.107

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Stormcloth Headband (10032) | Tailoring [crafted] | sim-verified (183.9 DPS) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Blood Guard's Plate Helm (220803, +0.00 DPS) [vendor]; Raging Berserker's Helm (7719, -7.65 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.0 attack_power points (1.90 DPS) | yes | Woven Ivy Necklace (19159, -0.25 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.71 DPS) [quest]; Scout's Medallion (19535, -0.75 DPS) [rep] |
| shoulder | Stormcloth Shoulders (10038) | Tailoring [crafted] | sim-verified (183.9 DPS) | yes | Wyrmslayer Spaulders (13066, +0.00 DPS) [world_drop]; Blood Guard's Plate Pauldrons (220796, +0.00 DPS) [vendor]; Officer's Pauldrons (250576, -3.49 DPS, sim-verified) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 35.8 attack_power points (2.12 DPS) | yes | Dark Hooded Cape (5257, -0.64 DPS) [world]; Blisterbane Wrap (12552, -0.69 DPS) [dungeon]; Dark Phantom Cape (13122, -0.69 DPS) [world_drop] |
| chest | Stormcloth Vest (10020) | Tailoring [crafted] | sim-verified (183.9 DPS) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Stone Guard's Plate Armor (220801, +0.00 DPS) [vendor]; Green Dragonscale Breastplate (15045, -4.91 DPS, sim-verified) [crafted] |
| wrist | Officer's Wristguards (250581) | Blacksmithing [crafted] | 33.9 attack_power points (2.01 DPS) | yes | Deepfury Bracers (13120, +0.00 DPS) [world_drop]; Berserker Bracers (19580, +0.00 DPS) [pvp]; Forest Stalker's Bracers (19589, +0.00 DPS) [pvp] |
| hands | Stormcloth Gloves (10011) | Tailoring [crafted] | sim-verified (183.9 DPS) | yes | Officer's Gloves (250551, +0.00 DPS) [crafted]; Raider Gloves (272100, +0.00 DPS) [vendor]; Raider Gauntlets (272096, -5.72 DPS, sim-verified) [vendor] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 50.2 attack_power points (2.97 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Defiler's Plate Girdle (20205, -0.08 DPS) [rep]; Girdle of Beastial Fury (11686, -0.15 DPS) [dungeon] |
| legs | Stormcloth Pants (10010) | Tailoring [crafted] | sim-verified (183.9 DPS) | yes | Serpentskin Leggings (8262, +0.00 DPS) [world_drop]; Green Dragonscale Leggings (15046, +0.00 DPS) [crafted]; Stone Guard's Plate Leggings (220798, +0.00 DPS) [vendor] |
| feet | Stormcloth Boots (10039) | Tailoring [crafted] | sim-verified (183.9 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Battlechaser's Greaves (12555, -3.18 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.0 attack_power points (2.49 DPS) | yes | White Bone Band (11862, -1.07 DPS) [quest]; Ironspine's Eye (7686, -1.11 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.15 DPS) [quest] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 36.5 attack_power points (2.16 DPS) | yes | White Bone Band (11862, -0.74 DPS) [quest]; Ironspine's Eye (7686, -0.78 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.82 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (183.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (183.9 DPS) | yes | Molten Heart of the Mountain (249470, -1.39 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (183.9 DPS) | yes | Warmonger (13052, -0.39 DPS) [world_drop]; Taran Icebreaker (2915, -0.90 DPS) [world_drop]; Nightblade (1982, -1.68 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Stormcloth Headband; neck: Skibi's Pendant; shoulder: Stormcloth Shoulders; back: Blackveil Cape; wrist: Officer's Wristguards; waist: Prowler's Leather Waistguard; feet: Stormcloth Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 702, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60 (undead, 52003003000000000-0520000000000000-05025331001330311)

Set DPS (verified): 285.8. Weights run: 1.3s. Verify run: 7.7s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.200 ± 0.004, agility=1.904 ± 0.046, crit=2.694 ± 0.065 per rating point (14 rating = 1%, 37.717 per %), hit=3.161 ± 0.158 per rating point (10 rating = 1%, 31.610 per %), melee_haste=19.759 ± 1.845

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 178.3 attack_power points (11.20 DPS) | yes | Mask of the Unforgiven (13404, -4.86 DPS) [dungeon]; Blood Guard's Plate Helm (220803, -5.05 DPS) [vendor]; Eye of Rend (12587, -5.21 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 63.7 attack_power points (4.00 DPS) | yes | Medallion of the Dawn (22659, -0.13 DPS) [quest]; Pendant of Celerity (22340, -0.22 DPS) [dungeon]; Amulet of the Darkmoon (19491, -0.35 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 89.1 attack_power points (5.60 DPS) | yes | Darkspear Epaulets (272106, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -0.12 DPS) [dungeon]; Defiler's Plate Spaulders (20212, -1.08 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.6 attack_power points (3.74 DPS) | yes | Windshear Cape (20691, -0.84 DPS) [world]; Deathguard's Cloak (20068, -1.01 DPS) [rep]; Cape of the Black Baron (13340, -1.29 DPS, sim-verified) [dungeon] |
| chest | The Postmaster's Tunic (13388) | Stratholme: Postmaster Malown [dungeon] | sim-verified (285.8 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -13.62 DPS, sim-verified) [dungeon] |
| wrist | Battleborn Armbraces (12936) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (285.8 DPS) | yes | Forest Stalker's Bracers (19587, -0.56 DPS) [rep]; Berserker Bracers (19578, -0.77 DPS) [rep]; Bracers of Undead Slaying (23090, -4.22 DPS, sim-verified) [world] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-verified (285.8 DPS) | yes | Raider Gauntlets (272095, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -4.56 DPS, sim-verified) [quest] |
| waist | Radiant Girdle of the Dawn (227814) | Argent Quartermaster Hasana [vendor] | 97.1 attack_power points (6.10 DPS) | yes | Belt of Preserved Heads (20216, -0.39 DPS) [quest]; Ferocity of the Timbermaw (227805, -0.47 DPS) [vendor]; Defiler's Plate Girdle (20204, -1.38 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | sim-verified (285.8 DPS) | yes | Titanic Leggings (22385, -0.42 DPS) [crafted]; Sentinel's Plate Legguards (237825, -0.46 DPS) [vendor]; Cloudkeeper Legplates (14554, -8.99 DPS, sim-verified) [world_drop] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-verified (285.8 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Savage Gladiator Greaves (11731, -12.18 DPS, sim-verified) [dungeon] |
| finger1 | The Postmaster's Seal (13392) | Stratholme: Postmaster Malown [dungeon] | sim-verified (285.8 DPS) | yes | Tarnished Elven Ring (18500, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Don Julio's Band (19325, -5.01 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (285.8 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Tarnished Elven Ring (18500, -0.36 DPS) [dungeon]; Naglering (11669, -4.31 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (285.8 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, -1.96 DPS, sim-verified) [dungeon] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (285.8 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, -3.04 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (285.8 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.40 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Libram of Law (272435) | Pix Xizzix [vendor] | sim-verified (285.8 DPS) | yes | Libram of Fervor (23203, -1.47 DPS, sim-verified) [world_drop] |

**New at 60:** head: Lionheart Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: The Postmaster's Tunic; wrist: Battleborn Armbraces; hands: Savage Gladiator Grips; waist: Radiant Girdle of the Dawn; legs: Sentinel's Chain Leggings; feet: The Postmaster's Treads; finger1: The Postmaster's Seal; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Burst of Knowledge; main_hand: The Unstoppable Force; ranged: Libram of Law

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 55003003100100000-0000000000000000-05225331001330320)

Set DPS (verified): 661.4. Weights run: 1.6s. Verify run: 8.0s. 1699 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=2.420 ± 0.006, agility=2.173 ± 0.049, crit=2.800 ± 0.067 per rating point (14 rating = 1%, 39.197 per %), hit=5.127 ± 0.294 per rating point (10 rating = 1%, 51.269 per %), melee_haste=17.566 ± 2.436

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lionheart Helm (12640) | Blacksmithing [crafted] | 224.5 attack_power points (27.83 DPS) | yes | Mask of the Unforgiven (13404, -10.86 DPS, sim-verified) [dungeon]; Helm of the Executioner (22411, -10.92 DPS) [dungeon]; Stalwart Helm (250599, -11.52 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 83.9 attack_power points (10.40 DPS) | yes | Beads of Ogre Might (22150, -1.07 DPS) [quest]; Amulet of the Darkmoon (19491, -2.28 DPS) [quest]; Mark of Fordring (15411, -2.31 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (661.4 DPS) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Necropile Mantle (14633, -12.51 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 79.3 attack_power points (9.83 DPS) | yes | Cape of the Black Baron (13340, -3.31 DPS) [dungeon]; Windshear Cape (20691, -3.39 DPS) [world]; Stalwart Cloak (272415, -3.47 DPS) [vendor] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (661.4 DPS) | yes | Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Savage Gladiator Chain (11726, -6.19 DPS, sim-verified) [dungeon] |
| wrist | Necropile Cuffs (14629) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (661.4 DPS) | yes | Battleborn Armbraces (12936, +0.00 DPS) [dungeon]; Slashclaw Bracers (13211, +0.00 DPS) [dungeon]; Forest Stalker's Bracers (19587, +0.00 DPS) [rep] |
| hands | Savage Gladiator Grips (11730) | Blackrock Depths: Eviscerator [dungeon] | sim-verified (661.4 DPS) | yes | Stormshroud Gloves (21278, +0.00 DPS) [crafted]; Raider Gauntlets (272095, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -8.19 DPS, sim-verified) [quest] |
| waist | Belt of Preserved Heads (20216) | A Collection of Heads [quest] | 117.7 attack_power points (14.60 DPS) | yes | Radiant Girdle of the Dawn (227814, +0.00 DPS, sim-verified) [vendor]; Ferocity of the Timbermaw (227805, -2.22 DPS) [vendor]; Marksman's Girdle (22232, -2.58 DPS) [dungeon] |
| legs | Titanic Leggings (22385) | Blacksmithing [crafted] | sim-verified (661.4 DPS) | yes | Sentinel's Chain Leggings (237819, -1.07 DPS) [vendor]; Sentinel's Plate Legguards (237825, -2.40 DPS) [vendor]; Cloudkeeper Legplates (14554, -16.11 DPS, sim-verified) [world_drop] |
| feet | Necropile Boots (14631) | Scholomance: Lady Illucia Barov [dungeon] | sim-verified (661.4 DPS) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Savage Gladiator Greaves (11731, -4.54 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (661.4 DPS) | yes | Tarnished Elven Ring (18500, -2.80 DPS) [dungeon]; Cutthroat's Signet (272408, -3.07 DPS) [vendor]; Naglering (11669, -9.96 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (661.4 DPS) | yes | Tarnished Elven Ring (18500, -0.81 DPS) [dungeon]; Cutthroat's Signet (272408, -1.08 DPS) [vendor]; Naglering (11669, -7.49 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (661.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, -7.72 DPS, sim-verified) [dungeon] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (661.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (661.4 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Blackblade of Shahram (12592, -4.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Lionheart Helm; neck: Pendant of Celerity; shoulder: Ironfeather Shoulders; back: Howler's Furs; chest: Ironfeather Breastplate; wrist: Necropile Cuffs; hands: Savage Gladiator Grips; waist: Belt of Preserved Heads; legs: Titanic Leggings; feet: Necropile Boots; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force

No-known-source sample (15 of 1699, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3579 Ornate Copper Shoulders; 4081 Blackforge Leggings; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band

