# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.0. Weights run: 2.2s. Verify run: 4.5s. 215 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.938 ± 0.015, crit=0.594 ± 0.014 per rating point (14 rating = 1%, 8.317 per %), hit=0.619 ± 0.030 per rating point (10 rating = 1%, 6.194 per %), melee_haste=4.054 ± 0.376

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.5 attack_power points (0.96 DPS) | yes | Defender's Leather Hood (252447, -0.47 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.9 attack_power points (0.24 DPS) | yes | Tarnished Locket (279870, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.7 attack_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.63 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.6 attack_power points (0.72 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 21.3 attack_power points (1.31 DPS) | yes | Brawler's Leather Armor (252490, -0.18 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.48 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.8 attack_power points (0.60 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.88 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Dusty Belt (279897, -0.51 DPS) [quest]; Deviate Scale Belt (6468, -3.86 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.10 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.02 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.2 attack_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; The 1 Ring (8350, -0.45 DPS) [world]; Ring of the Moon (12052, -0.57 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.0 attack_power points (0.25 DPS) | yes | Lavishly Jeweled Ring (1156, -0.01 DPS) [dungeon]; The 1 Ring (8350, -0.07 DPS) [world]; Ring of the Moon (12052, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.10 DPS) | yes | Evocator's Blade (2567, -0.88 DPS) [dungeon]; Buzzer Blade (2169, -2.01 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.59 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 228.2 attack_power points (14.07 DPS) | yes | Evocator's Blade (2567, -4.23 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.8 attack_power points (0.48 DPS) | yes | Deadly Blunderbuss (4369, -0.24 DPS) [crafted]; Light Bow (4576, -0.24 DPS) [world_drop]; Fine Longbow (11304, -0.25 DPS, sim-verified) [vendor] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Demon Band; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 215, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 67.2. Weights run: 2.2s. Verify run: 4.7s. 409 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.579 ± 0.013, crit=0.599 ± 0.014 per rating point (14 rating = 1%, 8.390 per %), hit=0.823 ± 0.033 per rating point (10 rating = 1%, 8.234 per %), melee_haste=6.231 ± 0.395

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.8 attack_power points (0.99 DPS) | yes | Defender's Leather Helm (252455, -0.24 DPS) [crafted]; Brawler's Leather Hood (252504, -0.26 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.30 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.75 DPS) | yes | Kaleidoscope Chain (13084, -0.11 DPS) [world_drop]; River Pride Choker (13087, -0.50 DPS) [world_drop]; Erudite's Amulet (277204, -0.55 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.4 attack_power points (1.40 DPS) | yes | Mantle of Thieves (2264, -0.49 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.59 DPS) [crafted]; Bristlebark Amice (14573, -0.62 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.0 attack_power points (0.88 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.1 attack_power points (1.38 DPS) | yes | Tunic of Westfall (2041, -0.30 DPS) [quest]; Brawler's Leather Tunic (252508, -0.36 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.38 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.5 attack_power points (0.84 DPS) | yes | Cultist's Armguards (270032, -0.22 DPS) [quest]; Insignia Bracers (6410, -0.25 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS, sim-verified) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.2 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.93 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (67.2 DPS) | yes | Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Warden's Leather Belt (252460, -0.16 DPS) [crafted]; Skulker's Leather Belt (252520, -3.50 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.2 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.04 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (67.2 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.26 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.7 attack_power points (0.80 DPS) | yes | Ring of Precision (1491, -0.20 DPS) [dungeon]; Defias Renegade Ring (1076, -0.30 DPS) [dungeon]; Monkey Ring (6748, -0.30 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.6 attack_power points (0.60 DPS) | yes | Ring of Precision (1491, -0.01 DPS) [dungeon]; Defias Renegade Ring (1076, -0.11 DPS) [dungeon]; Monkey Ring (6748, -0.11 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.12 DPS) | yes | Thornspike (6681, -2.11 DPS) [dungeon]; Claw of the Shadowmancer (2912, -2.16 DPS) [world_drop]; Toxic Revenger (9453, -2.16 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | 296.5 attack_power points (18.52 DPS) | yes | Thornspike (6681, -0.35 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.5 attack_power points (0.59 DPS) | yes | Gravespike Repeater (273033, -0.09 DPS) [world]; Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.20 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Thunderbrow Ring; finger2: Pyrewood Signet Ring; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Golemsight Long Gun

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14145 Cursed Felblade

### Band 40 (night-elf, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 110.0. Weights run: 2.3s. Verify run: 2.7s. 559 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.425 ± 0.014, crit=0.548 ± 0.018 per rating point (14 rating = 1%, 7.678 per %), hit=0.963 ± 0.034 per rating point (10 rating = 1%, 9.635 per %), melee_haste=7.145 ± 0.468

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.7 attack_power points (1.82 DPS) | yes | Warden's Wizard Hat (14604, -0.55 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.65 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.82 DPS) | yes | Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.54 DPS) [world_drop]; Gazlowe's Charm (13088, -0.54 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.7 attack_power points (1.41 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Nightscape Shoulders (8192, -0.34 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.6 attack_power points (0.99 DPS) | yes | Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.20 DPS) [quest]; Parachute Cloak (10518, -0.21 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 24.2 attack_power points (1.65 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.29 DPS) [crafted]; Barbaric Harness (5739, -0.40 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Hawkeye's Bracers (14590, -0.51 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.58 DPS) [dungeon]; Dusky Bracers (7378, -0.58 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 27.7 attack_power points (1.88 DPS) | yes | Skulker's Leather Gloves (252525, -0.30 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.33 DPS) [crafted]; Dog Training Gloves (7756, -0.77 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 21.8 attack_power points (1.49 DPS) | yes | Skulker's Leather Belt (252520, -0.20 DPS) [crafted]; Scorpashi Sash (14652, -0.24 DPS) [world_drop]; Blackened Defias Belt (10403, -0.26 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 29.9 attack_power points (2.04 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Ferine Leggings (6690, -0.27 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.68 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.7 attack_power points (1.54 DPS) | yes | Prowler's Leather Shoes (252465, -0.12 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Thunderbrow Ring (13097, -0.53 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.36 DPS) | yes | Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Thunderbrow Ring (13097, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (110.0 DPS) | yes | Coldrage Dagger (10761, -5.20 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -6.76 DPS) [dungeon]; The Ziggler (8006, -7.07 DPS) [world_drop] |
| off_hand | Black Menace (6831) (or Coldrage Dagger (10761)) | In the Name of the Light [quest] | 415.4 attack_power points (28.26 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (110.0 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Swiftwind (13038, -0.27 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.15 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 559, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 151.6. Weights run: 3.1s. Verify run: 3.2s. 710 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.587 ± 0.014, crit=0.982 ± 0.021 per rating point (14 rating = 1%, 13.743 per %), hit=1.326 ± 0.054 per rating point (10 rating = 1%, 13.262 per %), melee_haste=5.808 ± 0.788

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.5 attack_power points (2.43 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.13 DPS) [crafted]; Lordrec Helmet (10741, -0.26 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.6 attack_power points (2.19 DPS) | yes | Kaleidoscope Chain (13084, -1.30 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.39 DPS, sim-verified) [world]; River Pride Choker (13087, -1.85 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (2.26 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.10 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.8 attack_power points (2.03 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Serpentskin Cloak (8259, -0.54 DPS) [world_drop]; Pridelord Cape (14673, -0.56 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 45.0 attack_power points (3.84 DPS) | yes | Warbear Harness (15064, -0.47 DPS) [crafted]; Blazewind Breastplate (11193, -0.47 DPS) [quest]; Charred Leather Tunic (19127, -1.40 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.8 attack_power points (2.37 DPS) | yes | Pridelord Bands (14672, -0.59 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.66 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.69 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.7 attack_power points (2.88 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.22 DPS) [crafted]; Serpentskin Gloves (8260, -0.45 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.47 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 34.2 attack_power points (2.92 DPS) | yes | Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted]; Sagebrush Girdle (17778, -0.89 DPS) [quest]; Ogron's Sash (13117, -0.93 DPS) [world_drop] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 45.0 attack_power points (3.84 DPS) | yes | Gryphon Rider's Leggings (9652, -0.53 DPS) [quest]; Serpentskin Leggings (8262, -0.57 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.00 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.7 attack_power points (2.71 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.2 attack_power points (1.90 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; Ring of the Underwood (2951, -0.37 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.71 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.18 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (151.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (151.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (151.6 DPS) | yes | Gut Ripper (2164, +0.00 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -0.72 DPS) [dungeon]; Widowmaker (4091, -4.25 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.97 DPS) | yes | Charstone Dirk (17710, -28.42 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (151.6 DPS) | yes | Skull Splitting Crossbow (13039, -0.45 DPS) [world_drop]; Guttbuster (13139, -0.56 DPS) [world_drop]; Dark Iron Rifle (16004, -2.17 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 710, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 211.0. Weights run: 3.3s. Verify run: 5.3s. 1469 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.807 ± 0.020, crit=1.638 ± 0.034 per rating point (14 rating = 1%, 22.928 per %), hit=2.109 ± 0.081 per rating point (10 rating = 1%, 21.090 per %), melee_haste=14.164 ± 1.216

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.3 attack_power points (5.85 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -1.66 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 attack_power points (2.30 DPS) | yes | Sunprism Pendant (20649, -0.55 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.33 DPS) [world]; Kaleidoscope Chain (13084, -1.40 DPS) [world_drop] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 44.0 attack_power points (3.55 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS, sim-verified) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.1 attack_power points (3.81 DPS) | yes | Dark Phantom Cape (13122, -1.90 DPS, sim-verified) [world_drop]; Fel Cape (279269, -1.95 DPS) [crafted]; Sergeant Major's Cape (16337, -1.99 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 110.9 attack_power points (8.96 DPS) | yes | Timbermaw Tunic (252484, -1.20 DPS, sim-verified) [crafted]; Field Marshal's Leather Chestpiece (231543, -1.61 DPS) [pvp]; Knight-Captain's Leather Chestpiece (227058, -2.66 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (211.0 DPS) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -2.46 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 50.9 attack_power points (4.12 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Knight-Lieutenant's Leather Gauntlets (16396, -0.07 DPS) [pvp]; Timbermaw Brawlers (19049, -3.99 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 46.3 attack_power points (3.74 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Knight-Captain's Leather Belt (16398, -0.82 DPS) [pvp]; Skulker's Leather Waistguard (252474, -0.85 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 68.9 attack_power points (5.57 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight's Leather Pants (220858, -3.90 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 44.3 attack_power points (3.58 DPS) | yes | Marshal's Leather Footguards (231546, -0.37 DPS) [pvp]; Albino Crocscale Boots (17728, -0.66 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.79 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.3 attack_power points (2.04 DPS) | yes | Mark of Kern (2262, -0.43 DPS) [dungeon]; Assault Band (13095, -0.43 DPS) [world_drop]; Blackstone Ring (17713, -0.43 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 20.1 attack_power points (1.62 DPS) | yes | Mark of Kern (2262, -0.01 DPS) [dungeon]; Assault Band (13095, -0.01 DPS) [world_drop]; Blackstone Ring (17713, -0.01 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -7.31 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 596.7 attack_power points (48.23 DPS) | yes | The Shadowfoot Stabber (24222, -2.56 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -0.58 DPS) [world_drop]; Guttbuster (13139, -0.63 DPS) [world_drop]; Dark Iron Rifle (16004, -2.15 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1469, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 607.6. Weights run: 3.5s. Verify run: 8.3s. 1469 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.096 ± 0.024, crit=1.893 ± 0.040 per rating point (14 rating = 1%, 26.496 per %), hit=2.885 ± 0.134 per rating point (10 rating = 1%, 28.853 per %), melee_haste=22.231 ± 1.784

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 83.8 attack_power points (14.81 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.8 attack_power points (5.79 DPS) | yes | Sunprism Pendant (20649, -1.34 DPS) [quest]; Kaleidoscope Chain (13084, -3.53 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.67 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 51.9 attack_power points (9.17 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -22.03 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 51.4 attack_power points (9.09 DPS) | yes | Dark Phantom Cape (13122, -4.05 DPS, sim-verified) [world_drop]; Fel Cape (279269, -4.41 DPS) [crafted]; Sergeant Major's Cape (16337, -4.57 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 132.9 attack_power points (23.49 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -4.08 DPS) [pvp]; Knight-Captain's Leather Chestpiece (227058, -7.70 DPS) [pvp]; Timbermaw Tunic (252484, -8.24 DPS, sim-verified) [crafted] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (607.6 DPS) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -20.59 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (607.6 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -8.93 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (607.6 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -18.86 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (607.6 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Knight's Leather Pants (220858, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 52.7 attack_power points (9.31 DPS) | yes | Marshal's Leather Footguards (231546, -1.16 DPS) [pvp]; Hivethrasher's Boots (275615, -2.27 DPS) [crafted]; Albino Crocscale Boots (17728, -23.90 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 29.3 attack_power points (5.19 DPS) | yes | Falcon's Hook (7552, -1.27 DPS) [dungeon]; Mark of Kern (2262, -1.65 DPS) [dungeon]; Blackstone Ring (17713, -1.65 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 23.2 attack_power points (4.09 DPS) | yes | Falcon's Hook (7552, -0.18 DPS) [dungeon]; Mark of Kern (2262, -0.56 DPS) [dungeon]; Blackstone Ring (17713, -0.56 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (607.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (607.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (607.6 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -21.62 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 599.6 attack_power points (105.97 DPS) | yes | The Shadowfoot Stabber (24222, -4.93 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (607.6 DPS) | yes | Eaglehorn Long Bow (13023, -1.48 DPS) [world_drop]; Guttbuster (13139, -1.64 DPS) [world_drop]; Dark Iron Rifle (16004, -4.43 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Ring of the Underwood; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1469, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 40.5. Weights run: 2.2s. Verify run: 4.5s. 214 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.938 ± 0.015, crit=0.594 ± 0.014 per rating point (14 rating = 1%, 8.317 per %), hit=0.619 ± 0.030 per rating point (10 rating = 1%, 6.194 per %), melee_haste=4.054 ± 0.376

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.5 attack_power points (0.96 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.9 attack_power points (0.24 DPS) | yes | Roadwatcher's Confidence (281265, -0.24 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.7 attack_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.60 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.6 attack_power points (0.72 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (40.5 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.00 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.7 attack_power points (0.60 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.24 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.6 attack_power points (0.96 DPS) | yes | Brawler's Leather Gloves (252494, -0.24 DPS) [crafted]; Bristlebark Gloves (14572, -0.24 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.25 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Murloc Scale Belt (5780, -0.63 DPS) [crafted]; Deviate Scale Belt (6468, -3.94 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.20 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (40.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.13 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.2 attack_power points (0.63 DPS) | yes | Demon Band (12054, -0.38 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; Loop of Sacrifice (281673, -0.45 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 5.8 attack_power points (0.36 DPS) | yes | Demon Band (12054, -0.11 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Loop of Sacrifice (281673, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.10 DPS) | yes | Evocator's Blade (2567, -0.88 DPS) [dungeon]; Buzzer Blade (2169, -2.01 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.59 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 228.2 attack_power points (14.07 DPS) | yes | Evocator's Blade (2567, -4.09 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.8 attack_power points (0.48 DPS) | yes | Fine Longbow (11304, -0.21 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.24 DPS) [crafted]; Light Bow (4576, -0.24 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Bounty Hunter's Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 214, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 66.4. Weights run: 2.2s. Verify run: 4.7s. 408 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.579 ± 0.013, crit=0.599 ± 0.014 per rating point (14 rating = 1%, 8.390 per %), hit=0.823 ± 0.033 per rating point (10 rating = 1%, 8.234 per %), melee_haste=6.231 ± 0.395

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.8 attack_power points (0.99 DPS) | yes | Defender's Leather Helm (252455, -0.24 DPS) [crafted]; Brawler's Leather Hood (252504, -0.26 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.30 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.75 DPS) | yes | Kaleidoscope Chain (13084, -0.11 DPS) [world_drop]; River Pride Choker (13087, -0.50 DPS) [world_drop]; Erudite's Amulet (277204, -0.55 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.4 attack_power points (1.40 DPS) | yes | Mantle of Thieves (2264, -0.49 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.59 DPS) [crafted]; Bristlebark Amice (14573, -0.62 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.0 attack_power points (0.88 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon]; Wildhunter Cloak (16658, -0.25 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.1 attack_power points (1.38 DPS) | yes | Brawler's Leather Tunic (252508, -0.32 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.37 DPS) [quest]; Brawler's Leather Armor (252490, -0.38 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.5 attack_power points (0.84 DPS) | yes | Cultist's Armguards (270032, -0.22 DPS) [quest]; Insignia Bracers (6410, -0.25 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS, sim-verified) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.4 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.71 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (66.4 DPS) | yes | Deftkin Belt (16659, +0.00 DPS) [quest]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Skulker's Leather Belt (252520, -3.26 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.4 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.90 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (66.4 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.09 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.7 attack_power points (0.80 DPS) | yes | Pyrewood Signet Ring (277210, -0.20 DPS) [quest]; Ring of Precision (1491, -0.20 DPS) [dungeon]; Monkey Ring (6748, -0.30 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 10.3 attack_power points (0.64 DPS) | yes | Pyrewood Signet Ring (277210, -0.04 DPS) [quest]; Ring of Precision (1491, -0.05 DPS) [dungeon]; Monkey Ring (6748, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.12 DPS) | yes | Thornspike (6681, -2.11 DPS) [dungeon]; Claw of the Shadowmancer (2912, -2.16 DPS) [world_drop]; Toxic Revenger (9453, -2.16 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | 296.5 attack_power points (18.52 DPS) | yes | Thornspike (6681, -0.34 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.5 attack_power points (0.59 DPS) | yes | Gravespike Repeater (273033, -0.09 DPS) [world]; Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.20 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Golemsight Long Gun

No-known-source sample (15 of 408, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 109.2. Weights run: 2.3s. Verify run: 2.7s. 556 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.425 ± 0.014, crit=0.548 ± 0.018 per rating point (14 rating = 1%, 7.678 per %), hit=0.963 ± 0.034 per rating point (10 rating = 1%, 9.635 per %), melee_haste=7.145 ± 0.468

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.7 attack_power points (1.82 DPS) | yes | Warden's Wizard Hat (14604, -0.55 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.65 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.82 DPS) | yes | Ethereal Talisman (4430, -0.09 DPS) [quest]; Kaleidoscope Chain (13084, -0.16 DPS) [world_drop]; River Pride Choker (13087, -0.54 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 20.7 attack_power points (1.41 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Nightscape Shoulders (8192, -0.34 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.0 attack_power points (0.88 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Parachute Cloak (10518, -0.11 DPS) [crafted]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 24.2 attack_power points (1.65 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.29 DPS) [crafted]; Barbaric Harness (5739, -0.40 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Hawkeye's Bracers (14590, -0.42 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.58 DPS) [dungeon]; Dusky Bracers (7378, -0.58 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 27.7 attack_power points (1.88 DPS) | yes | Skulker's Leather Gloves (252525, -0.30 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.33 DPS) [crafted]; Dog Training Gloves (7756, -0.73 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 21.8 attack_power points (1.49 DPS) | yes | Skulker's Leather Belt (252520, -0.20 DPS) [crafted]; Scorpashi Sash (14652, -0.24 DPS) [world_drop]; Blackened Defias Belt (10403, -0.26 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 29.9 attack_power points (2.04 DPS) | yes | Triprunner Dungarees (9624, -0.09 DPS) [quest]; Ferine Leggings (6690, -0.27 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.68 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.7 attack_power points (1.54 DPS) | yes | Prowler's Leather Shoes (252465, -0.12 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Thunderbrow Ring (13097, -0.53 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.36 DPS) | yes | Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon]; Thunderbrow Ring (13097, -0.53 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (32.28 DPS) | yes | Hypnotic Blade (7714, -6.76 DPS) [dungeon]; The Ziggler (8006, -7.07 DPS) [world_drop]; Sacrificial Kris (3187, -7.79 DPS) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 415.4 attack_power points (28.26 DPS) | yes | Hypnotic Blade (7714, -27.13 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (109.2 DPS) | yes | Monolithic Bow (9426, -0.25 DPS) [dungeon]; Swiftwind (13038, -0.27 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.13 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 556, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 155.1. Weights run: 3.1s. Verify run: 3.2s. 707 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.587 ± 0.014, crit=0.982 ± 0.021 per rating point (14 rating = 1%, 13.743 per %), hit=1.326 ± 0.054 per rating point (10 rating = 1%, 13.262 per %), melee_haste=5.808 ± 0.788

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.5 attack_power points (2.43 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.13 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.6 attack_power points (2.19 DPS) | yes | Woven Ivy Necklace (19159, -0.46 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.16 DPS) [world]; Ethereal Talisman (4430, -1.22 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (2.26 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.10 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.8 attack_power points (2.03 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Serpentskin Cloak (8259, -0.54 DPS) [world_drop]; Pridelord Cape (14673, -0.56 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 10 · Stone Guard · Horde [vendor] | 45.0 attack_power points (3.84 DPS) | yes | Warbear Harness (15064, -0.47 DPS) [crafted]; Blazewind Breastplate (11193, -0.47 DPS) [quest]; Charred Leather Tunic (19127, -1.40 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.8 attack_power points (2.37 DPS) | yes | Pridelord Bands (14672, -0.59 DPS) [world_drop]; Skulker's Leather Bracers (252540, -0.64 DPS, sim-verified) [crafted]; Prowler's Leather Bracers (252539, -0.66 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.7 attack_power points (2.88 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.22 DPS) [crafted]; Serpentskin Gloves (8260, -0.45 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.47 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 34.2 attack_power points (2.92 DPS) | yes | Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted]; Sagebrush Girdle (17778, -0.89 DPS) [quest]; Ogron's Sash (13117, -0.93 DPS) [world_drop] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 45.0 attack_power points (3.84 DPS) | yes | Serpentskin Leggings (8262, -0.57 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.00 DPS) [world_drop]; Triprunner Dungarees (9624, -1.15 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.7 attack_power points (2.71 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.05 DPS) | yes | Mark of Kern (2262, -0.34 DPS) [dungeon]; Assault Band (13095, -0.34 DPS) [world_drop]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.2 attack_power points (1.90 DPS) | yes | Mark of Kern (2262, -0.19 DPS) [dungeon]; Assault Band (13095, -0.19 DPS) [world_drop]; Blackstone Ring (17713, -0.19 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (155.1 DPS) | yes | Frozen Heart of the Mountain (249469, -3.91 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (155.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (155.1 DPS) | yes | Gut Ripper (2164, -0.64 DPS, sim-verified) [world_drop]; Charstone Dirk (17710, -0.72 DPS) [dungeon]; Widowmaker (4091, -4.25 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.97 DPS) | yes | Charstone Dirk (17710, -29.37 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (155.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.45 DPS) [world_drop]; Guttbuster (13139, -0.56 DPS) [world_drop]; Dark Iron Rifle (16004, -2.17 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 707, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 215.4. Weights run: 3.3s. Verify run: 4.9s. 1468 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.807 ± 0.020, crit=1.638 ± 0.034 per rating point (14 rating = 1%, 22.928 per %), hit=2.109 ± 0.081 per rating point (10 rating = 1%, 21.090 per %), melee_haste=14.164 ± 1.216

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.3 attack_power points (5.85 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -1.91 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.5 attack_power points (2.30 DPS) | yes | Woven Ivy Necklace (19159, -0.50 DPS) [quest]; Sunprism Pendant (20649, -0.55 DPS) [quest]; Ethereal Talisman (4430, -1.31 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 44.0 attack_power points (3.55 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.1 attack_power points (3.81 DPS) | yes | Dark Phantom Cape (13122, -1.89 DPS, sim-verified) [world_drop]; Fel Cape (279269, -1.95 DPS) [crafted]; Blackflame Cape (13109, -2.05 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 110.9 attack_power points (8.96 DPS) | yes | Warlord's Leather Breastplate (231549, -1.61 DPS) [pvp]; Timbermaw Tunic (252484, -1.73 DPS, sim-verified) [crafted]; Legionnaire's Leather Chestpiece (227060, -2.66 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (215.4 DPS) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.31 DPS) [crafted]; Deepfury Bracers (13120, -2.45 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 50.9 attack_power points (4.12 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Blood Guard's Leather Vices (16499, -0.07 DPS) [pvp]; Timbermaw Brawlers (19049, -4.42 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 46.3 attack_power points (3.74 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Legionnaire's Leather Girdle (16500, -0.82 DPS) [pvp]; Skulker's Leather Waistguard (252474, -0.85 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 68.9 attack_power points (5.57 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Stone Guard's Leather Pants (220859, -4.44 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 44.3 attack_power points (3.58 DPS) | yes | General's Leather Treads (231552, -0.37 DPS) [pvp]; Albino Crocscale Boots (17728, -0.66 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.79 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.3 attack_power points (2.04 DPS) | yes | Ring of the Underwood (2951, -0.42 DPS) [world_drop]; Mark of Kern (2262, -0.43 DPS) [dungeon]; Blackstone Ring (17713, -0.43 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.94 DPS) | yes | Ring of the Underwood (2951, -0.32 DPS) [world_drop]; Mark of Kern (2262, -0.32 DPS) [dungeon]; Blackstone Ring (17713, -0.32 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -8.76 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 596.7 attack_power points (48.23 DPS) | yes | The Shadowfoot Stabber (24222, -2.42 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -0.58 DPS) [world_drop]; Guttbuster (13139, -0.63 DPS) [world_drop]; Dark Iron Rifle (16004, -2.18 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket2: Frozen Heart of the Mountain; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1468, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 620.4. Weights run: 3.5s. Verify run: 8.2s. 1468 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.096 ± 0.024, crit=1.893 ± 0.040 per rating point (14 rating = 1%, 26.496 per %), hit=2.885 ± 0.134 per rating point (10 rating = 1%, 28.853 per %), melee_haste=22.231 ± 1.784

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 83.8 attack_power points (14.81 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -2.69 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.8 attack_power points (5.79 DPS) | yes | Woven Ivy Necklace (19159, -1.29 DPS) [quest]; Sunprism Pendant (20649, -1.34 DPS) [quest]; Ethereal Talisman (4430, -3.33 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 51.9 attack_power points (9.17 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -20.57 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 51.4 attack_power points (9.09 DPS) | yes | Dark Phantom Cape (13122, -4.08 DPS, sim-verified) [world_drop]; Fel Cape (279269, -4.41 DPS) [crafted]; Blackflame Cape (13109, -4.65 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 132.9 attack_power points (23.49 DPS) | yes | Warlord's Leather Breastplate (231549, -4.08 DPS) [pvp]; Timbermaw Tunic (252484, -5.01 DPS, sim-verified) [crafted]; Legionnaire's Leather Chestpiece (227060, -7.70 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (620.4 DPS) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.86 DPS) [crafted]; Deepfury Bracers (13120, -18.16 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (620.4 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -7.08 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (620.4 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -16.41 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (620.4 DPS) | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 52.7 attack_power points (9.31 DPS) | yes | General's Leather Treads (231552, -1.16 DPS) [pvp]; Hivethrasher's Boots (275615, -2.27 DPS) [crafted]; Albino Crocscale Boots (17728, -23.08 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 29.3 attack_power points (5.19 DPS) | yes | Ring of the Underwood (2951, -1.09 DPS) [world_drop]; Falcon's Hook (7552, -1.27 DPS) [dungeon]; Blackstone Ring (17713, -1.65 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (4.24 DPS) | yes | Ring of the Underwood (2951, -0.15 DPS) [world_drop]; Falcon's Hook (7552, -0.32 DPS) [dungeon]; Blackstone Ring (17713, -0.71 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (620.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (620.4 DPS) | yes | Molten Heart of the Mountain (249470, -6.21 DPS, sim-verified) [crafted] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (620.4 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -18.10 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 599.6 attack_power points (105.97 DPS) | yes | The Shadowfoot Stabber (24222, -3.19 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (620.4 DPS) | yes | Eaglehorn Long Bow (13023, -1.48 DPS) [world_drop]; Guttbuster (13139, -1.64 DPS) [world_drop]; Dark Iron Rifle (16004, -4.45 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; trinket2: Frozen Heart of the Mountain; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1468, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

