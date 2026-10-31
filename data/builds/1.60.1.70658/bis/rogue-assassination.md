# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 39.6. Weights run: 1.3s. Verify run: 2.6s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.974 ± 0.014, crit=0.569 ± 0.013 per rating point (14 rating = 1%, 7.966 per %), hit=0.686 ± 0.032 per rating point (10 rating = 1%, 6.856 per %), melee_haste=4.645 ± 0.418

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.8 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.43 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.9 attack_power points (0.25 DPS) | yes | Tarnished Locket (279870, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.9 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.8 attack_power points (0.76 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 21.7 attack_power points (1.40 DPS) | yes | Brawler's Leather Armor (252490, -0.19 DPS) [crafted]; Prospector's Chestpiece (14562, -0.51 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.51 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.9 attack_power points (0.64 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -1.95 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Dusty Belt (279897, -0.52 DPS) [quest]; Deviate Scale Belt (6468, -3.09 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.20 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (39.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.10 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.6 attack_power points (0.69 DPS) | yes | Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world]; Ring of the Moon (12052, -0.62 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.0 attack_power points (0.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.00 DPS) [dungeon]; The 1 Ring (8350, -0.07 DPS) [world]; Ring of the Moon (12052, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (16.05 DPS) | yes | Diamond Hammer (2194, +0.00 DPS, sim-verified) [world_drop]; Blackfang (2236, -1.31 DPS) [world_drop]; Assassin's Blade (1935, -1.33 DPS) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (39.6 DPS) | yes | Diamond Hammer (2194, -0.35 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.9 attack_power points (0.51 DPS) | yes | Fine Longbow (11304, -0.20 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Demon Band; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 62.9. Weights run: 1.2s. Verify run: 2.8s. 403 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.714 ± 0.013, crit=0.667 ± 0.015 per rating point (14 rating = 1%, 9.342 per %), hit=0.897 ± 0.040 per rating point (10 rating = 1%, 8.973 per %), melee_haste=7.864 ± 0.505

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.1 attack_power points (1.14 DPS) | yes | Brawler's Leather Hood (252504, -0.23 DPS) [crafted]; Defender's Leather Helm (252455, -0.34 DPS) [crafted]; Humbert's Helm (4724, -0.34 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.80 DPS) | yes | Kaleidoscope Chain (13084, -0.08 DPS) [world_drop]; River Pride Choker (13087, -0.53 DPS) [world_drop]; Erudite's Amulet (277204, -0.57 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.68 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.28 DPS) [pvp]; Fenrus' Hide (6340, -0.31 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.0 attack_power points (1.59 DPS) | yes | Brawler's Leather Tunic (252508, -0.33 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.34 DPS) [quest]; Brawler's Leather Armor (252490, -0.46 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.3 attack_power points (0.95 DPS) | yes | Barbaric Bracers (18948, -0.23 DPS) [crafted]; Insignia Bracers (6410, -0.27 DPS) [world_drop]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (62.9 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.44 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (62.9 DPS) | yes | Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Warden's Leather Belt (252460, -0.11 DPS) [crafted]; Skulker's Leather Belt (252520, -3.99 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (62.9 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.58 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (62.9 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.77 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.87 DPS) | yes | Ring of Precision (1491, -0.19 DPS) [dungeon]; Defias Renegade Ring (1076, -0.30 DPS) [dungeon]; Monkey Ring (6748, -0.30 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.4 attack_power points (0.69 DPS) | yes | Ring of Precision (1491, -0.01 DPS) [dungeon]; Defias Renegade Ring (1076, -0.12 DPS) [dungeon]; Monkey Ring (6748, -0.12 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (21.39 DPS) | yes | Zealot Blade (13033, -0.80 DPS) [world_drop]; Electrocutioner Leg (9446, -0.87 DPS) [dungeon]; Dreamslayer (7752, -1.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 attack_power points (21.15 DPS) | yes | Zealot Blade (13033, -2.69 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.3 attack_power points (0.68 DPS) | yes | Double-barreled Shotgun (2098, -0.14 DPS) [world_drop]; Gravespike Repeater (273033, -0.15 DPS) [world]; Silver Star (3463, -0.23 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Thunderbrow Ring; finger2: Pyrewood Signet Ring; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14145 Cursed Felblade

### Band 40 (night-elf, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 109.1. Weights run: 1.3s. Verify run: 1.5s. 553 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.551 ± 0.015, crit=0.716 ± 0.021 per rating point (14 rating = 1%, 10.030 per %), hit=1.873 ± 0.077 per rating point (10 rating = 1%, 18.732 per %), melee_haste=30.832 ± 0.898

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.1 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.41 DPS) | yes | Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; River Pride Choker (13087, -0.28 DPS) [world_drop]; Gazlowe's Charm (13088, -0.28 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.1 attack_power points (0.76 DPS) | yes | Flintrock Shoulders (7755, -0.05 DPS) [dungeon]; Nightscape Shoulders (8192, -0.17 DPS) [crafted]; Mantle of Thieves (2264, -0.23 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 15.3 attack_power points (0.53 DPS) | yes | Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Imperial Cloak (6432, -0.10 DPS) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 26.4 attack_power points (0.91 DPS) | yes | Nightscape Tunic (8175, -0.11 DPS) [crafted]; Dusky Leather Armor (7374, -0.16 DPS) [crafted]; Barbaric Harness (5739, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.26 DPS) [dungeon]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.0 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Dog Training Gloves (7756, -0.21 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.0 attack_power points (0.79 DPS) | yes | Skulker's Leather Belt (252520, -0.10 DPS) [crafted]; Scorpashi Sash (14652, -0.12 DPS) [world_drop]; Prowler's Leather Belt (252459, -0.16 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.6 attack_power points (1.12 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.23 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.37 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.1 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Ring of the Underwood (2951, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.22 DPS) [quest] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.69 DPS) | yes | Ring of the Underwood (2951, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.22 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Dazzling Longsword (869, -1.17 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (109.1 DPS) | yes | Ardent Custodian (868, -63.78 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.11 DPS) [world_drop]; Monolithic Bow (9426, -0.12 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 553, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 166.6. Weights run: 1.9s. Verify run: 1.8s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.579 ± 0.012, crit=0.963 ± 0.018 per rating point (14 rating = 1%, 13.485 per %), hit=2.203 ± 0.082 per rating point (10 rating = 1%, 22.026 per %), melee_haste=5.914 ± 1.264

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.4 attack_power points (2.45 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.13 DPS) [crafted]; Lordrec Helmet (10741, -0.27 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (2.21 DPS) | yes | Kaleidoscope Chain (13084, -1.32 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.41 DPS, sim-verified) [world]; River Pride Choker (13087, -1.86 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.4 attack_power points (2.28 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.10 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.7 attack_power points (2.05 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Pridelord Cape (14673, -0.56 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Blazewind Breastplate (11193, -1.23 DPS) [quest]; Warbear Harness (15064, -1.69 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -2.17 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.7 attack_power points (2.39 DPS) | yes | Skulker's Leather Bracers (252540, -0.56 DPS) [crafted]; Pridelord Bands (14672, -0.60 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.66 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.5 attack_power points (2.89 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.21 DPS) [crafted]; Serpentskin Gloves (8260, -0.44 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.46 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 34.1 attack_power points (2.95 DPS) | yes | Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted]; Sagebrush Girdle (17778, -0.90 DPS) [quest]; Ogron's Sash (13117, -0.94 DPS) [world_drop] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Serpentskin Leggings (8262, -1.32 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.76 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.78 DPS, sim-verified) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.6 attack_power points (2.73 DPS) | yes | Skulker's Leather Boots (252469, -0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.1 attack_power points (1.91 DPS) | yes | Mark of Kern (2262, -0.18 DPS) [dungeon]; Assault Band (13095, -0.18 DPS) [world_drop]; Ring of the Underwood (2951, -0.37 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.73 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.19 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (166.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (166.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (166.6 DPS) | yes | Hanzo Sword (8190, -0.66 DPS, sim-verified) [world_drop]; Julie's Dagger (6660, -3.61 DPS) [world_drop]; Viking Warhammer (1721, -4.03 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (45.50 DPS) | yes | Julie's Dagger (6660, -2.82 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.83 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (166.6 DPS) | yes | Skull Splitting Crossbow (13039, -0.44 DPS) [world_drop]; Guttbuster (13139, -0.56 DPS) [world_drop]; Dark Iron Rifle (16004, -2.24 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 246.6. Weights run: 1.8s. Verify run: 4.3s. 1463 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.815 ± 0.019, crit=1.646 ± 0.030 per rating point (14 rating = 1%, 23.048 per %), hit=3.401 ± 0.116 per rating point (10 rating = 1%, 34.008 per %), melee_haste=9.014 ± 1.758

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.7 attack_power points (6.08 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -2.11 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.6 attack_power points (2.39 DPS) | yes | Sunprism Pendant (20649, -0.57 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.39 DPS) [world]; Kaleidoscope Chain (13084, -1.45 DPS) [world_drop] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 46.7 attack_power points (3.91 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -12.17 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.2 attack_power points (3.95 DPS) | yes | Fel Cape (279269, -2.02 DPS) [crafted]; Sergeant Major's Cape (16337, -2.07 DPS) [pvp]; Dark Phantom Cape (13122, -2.20 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.2 attack_power points (10.39 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -1.67 DPS) [pvp]; Knight-Captain's Leather Chestpiece (227058, -2.77 DPS) [pvp]; Timbermaw Tunic (252484, -4.15 DPS, sim-verified) [crafted] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (246.6 DPS) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -11.75 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (246.6 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -4.79 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (246.6 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -10.90 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (246.6 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Knight's Leather Pants (220858, +0.00 DPS, sim-verified) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 48.3 attack_power points (4.04 DPS) | yes | Marshal's Leather Footguards (231546, -0.70 DPS) [pvp]; Skulker's Leather Boots (252469, -1.15 DPS) [crafted]; Albino Crocscale Boots (17728, -12.94 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.4 attack_power points (2.12 DPS) | yes | Ring of the Underwood (2951, -0.44 DPS) [world_drop]; Mark of Kern (2262, -0.45 DPS) [dungeon]; Blackstone Ring (17713, -0.45 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 20.9 attack_power points (1.74 DPS) | yes | Ring of the Underwood (2951, -0.06 DPS) [world_drop]; Mark of Kern (2262, -0.07 DPS) [dungeon]; Blackstone Ring (17713, -0.07 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (246.6 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (246.6 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Greenhammer (279261) | Blacksmithing [crafted] | sim-verified (246.6 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -2.54 DPS, sim-verified) [world_drop] |
| off_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | 704.1 attack_power points (58.86 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -18.55 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (246.6 DPS) | yes | Eaglehorn Long Bow (13023, -0.61 DPS) [world_drop]; Guttbuster (13139, -0.66 DPS) [world_drop]; Dark Iron Rifle (16004, -2.49 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Pyrewood Signet Ring; main_hand: Greenhammer; off_hand: Ironwood Blade

No-known-source sample (15 of 1463, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 644.5. Weights run: 1.9s. Verify run: 4.6s. 1463 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.200 ± 0.021, crit=2.108 ± 0.035 per rating point (14 rating = 1%, 29.512 per %), hit=3.859 ± 0.162 per rating point (10 rating = 1%, 38.590 per %), melee_haste=16.559 ± 2.071

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 90.5 attack_power points (15.85 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -5.15 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.1 attack_power points (5.97 DPS) | yes | Sunprism Pendant (20649, -1.35 DPS) [quest]; Kaleidoscope Chain (13084, -3.66 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.87 DPS) [world] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (644.5 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -44.26 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.0 attack_power points (9.28 DPS) | yes | Fel Cape (279269, -4.11 DPS) [crafted]; Dark Phantom Cape (13122, -4.60 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.66 DPS) [world_drop] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (644.5 DPS) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -30.69 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.4 attack_power points (6.55 DPS) | yes | Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Knight-Captain's Leather Bracers (16394, -0.77 DPS) [pvp]; Shadowcraft Bracers (16710, -0.77 DPS) [dungeon] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 57.5 attack_power points (10.07 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Knight-Lieutenant's Leather Gauntlets (16396, -0.62 DPS) [pvp]; Timbermaw Brawlers (19049, -0.63 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 53.9 attack_power points (9.44 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.73 DPS) [crafted]; Knight-Captain's Leather Belt (16398, -1.93 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (644.5 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -29.73 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 57.8 attack_power points (10.12 DPS) | yes | Marshal's Leather Footguards (231546, -1.64 DPS) [pvp]; Albino Crocscale Boots (17728, -2.41 DPS) [dungeon]; Hivethrasher's Boots (275615, -2.80 DPS) [crafted] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.8 attack_power points (5.40 DPS) | yes | Ring of the Underwood (2951, -1.16 DPS) [world_drop]; Falcon's Hook (7552, -1.35 DPS) [dungeon]; Blackstone Ring (17713, -1.89 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 24.2 attack_power points (4.25 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [dungeon]; Blackstone Ring (17713, -0.74 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (644.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (644.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Greenhammer (279261) | Blacksmithing [crafted] | sim-verified (644.5 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -5.15 DPS, sim-verified) [world_drop] |
| off_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | 709.5 attack_power points (124.27 DPS) | yes | Persuader (22384, +0.00 DPS, sim-verified) [crafted]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Claw of Celebras (17738, -39.81 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (644.5 DPS) | yes | Eaglehorn Long Bow (13023, -1.54 DPS) [world_drop]; Guttbuster (13139, -1.73 DPS) [world_drop]; Dark Iron Rifle (16004, -6.03 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger2: Pyrewood Signet Ring; main_hand: Greenhammer; off_hand: Ironwood Blade

No-known-source sample (15 of 1463, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 39.5. Weights run: 1.3s. Verify run: 2.7s. 208 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.974 ± 0.014, crit=0.569 ± 0.013 per rating point (14 rating = 1%, 7.966 per %), hit=0.686 ± 0.032 per rating point (10 rating = 1%, 6.856 per %), melee_haste=4.645 ± 0.418

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.8 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.45 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.9 attack_power points (0.25 DPS) | yes | Roadwatcher's Confidence (281265, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.9 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.65 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.8 attack_power points (0.76 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.97 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.9 attack_power points (0.64 DPS) | yes | Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.25 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.8 attack_power points (1.02 DPS) | yes | Brawler's Leather Gloves (252494, -0.25 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Bristlebark Gloves (14572, -0.26 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Murloc Scale Belt (5780, -0.65 DPS) [crafted]; Deviate Scale Belt (6468, -3.08 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.19 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.08 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.6 attack_power points (0.69 DPS) | yes | Demon Band (12054, -0.43 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Loop of Sacrifice (281673, -0.49 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 5.9 attack_power points (0.38 DPS) | yes | Demon Band (12054, -0.12 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Loop of Sacrifice (281673, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Diamond Hammer (2194) | World drop [world_drop] | sim-verified (39.5 DPS) | yes | Wingblade (6504, -0.03 DPS) [quest]; Blackfang (2236, -0.08 DPS) [world_drop]; Shadowfang (1482, -0.57 DPS, sim-verified) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.9 attack_power points (0.51 DPS) | yes | Fine Longbow (11304, -0.22 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Bounty Hunter's Ring; main_hand: Diamond Hammer; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 208, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 62.3. Weights run: 1.2s. Verify run: 2.7s. 402 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.714 ± 0.013, crit=0.667 ± 0.015 per rating point (14 rating = 1%, 9.342 per %), hit=0.897 ± 0.040 per rating point (10 rating = 1%, 8.973 per %), melee_haste=7.864 ± 0.505

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.1 attack_power points (1.14 DPS) | yes | Brawler's Leather Hood (252504, -0.23 DPS) [crafted]; Defender's Leather Helm (252455, -0.34 DPS) [crafted]; Humbert's Helm (4724, -0.34 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.80 DPS) | yes | Kaleidoscope Chain (13084, -0.08 DPS) [world_drop]; River Pride Choker (13087, -0.53 DPS) [world_drop]; Erudite's Amulet (277204, -0.57 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.46 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.68 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Fenrus' Hide (6340, -0.31 DPS) [dungeon]; Swiftrunner Cape (6745, -0.31 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.0 attack_power points (1.59 DPS) | yes | Brawler's Leather Tunic (252508, -0.32 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.44 DPS) [quest]; Brawler's Leather Armor (252490, -0.46 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.3 attack_power points (0.95 DPS) | yes | Barbaric Bracers (18948, -0.23 DPS) [crafted]; Insignia Bracers (6410, -0.27 DPS) [world_drop]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (62.3 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.99 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (62.3 DPS) | yes | Deftkin Belt (16659, +0.00 DPS) [quest]; Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Skulker's Leather Belt (252520, -3.55 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (62.3 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.14 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (62.3 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest]; Feet of the Lynx (1121, -3.34 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.87 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Ring of Precision (1491, -0.19 DPS) [dungeon]; Monkey Ring (6748, -0.30 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 10.9 attack_power points (0.72 DPS) | yes | Pyrewood Signet Ring (277210, -0.03 DPS) [quest]; Ring of Precision (1491, -0.04 DPS) [dungeon]; Monkey Ring (6748, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (21.39 DPS) | yes | Zealot Blade (13033, -0.80 DPS) [world_drop]; Electrocutioner Leg (9446, -0.87 DPS) [dungeon]; Dreamslayer (7752, -1.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 attack_power points (21.15 DPS) | yes | Zealot Blade (13033, -2.23 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.3 attack_power points (0.68 DPS) | yes | Double-barreled Shotgun (2098, -0.14 DPS) [world_drop]; Gravespike Repeater (273033, -0.15 DPS) [world]; Silver Star (3463, -0.23 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 108.2. Weights run: 1.3s. Verify run: 1.4s. 550 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.551 ± 0.015, crit=0.716 ± 0.021 per rating point (14 rating = 1%, 10.030 per %), hit=1.873 ± 0.077 per rating point (10 rating = 1%, 18.732 per %), melee_haste=30.832 ± 0.898

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.1 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.41 DPS) | yes | Ethereal Talisman (4430, -0.03 DPS) [quest]; Kaleidoscope Chain (13084, -0.06 DPS) [world_drop]; River Pride Choker (13087, -0.28 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.1 attack_power points (0.76 DPS) | yes | Flintrock Shoulders (7755, -0.05 DPS) [dungeon]; Nightscape Shoulders (8192, -0.17 DPS) [crafted]; Mantle of Thieves (2264, -0.23 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.9 attack_power points (0.48 DPS) | yes | Imperial Cloak (6432, -0.05 DPS) [dungeon]; Parachute Cloak (10518, -0.05 DPS) [crafted]; Tigerstrike Mantle (13108, -0.05 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 26.4 attack_power points (0.91 DPS) | yes | Nightscape Tunic (8175, -0.11 DPS) [crafted]; Dusky Leather Armor (7374, -0.16 DPS) [crafted]; Barbaric Harness (5739, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.26 DPS) [dungeon]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.0 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Dog Training Gloves (7756, -0.21 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.0 attack_power points (0.79 DPS) | yes | Skulker's Leather Belt (252520, -0.10 DPS) [crafted]; Scorpashi Sash (14652, -0.12 DPS) [world_drop]; Prowler's Leather Belt (252459, -0.16 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.6 attack_power points (1.12 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.23 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.37 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.1 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Ring of the Underwood (2951, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.22 DPS) [quest] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (0.69 DPS) | yes | Ring of the Underwood (2951, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.22 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Dazzling Longsword (869, -1.17 DPS) [world_drop] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (108.2 DPS) | yes | Ardent Custodian (868, -62.63 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.11 DPS) [world_drop]; Monolithic Bow (9426, -0.12 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.51 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 550, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 171.1. Weights run: 1.9s. Verify run: 1.8s. 701 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.579 ± 0.012, crit=0.963 ± 0.018 per rating point (14 rating = 1%, 13.485 per %), hit=2.203 ± 0.082 per rating point (10 rating = 1%, 22.026 per %), melee_haste=5.914 ± 1.264

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.4 attack_power points (2.45 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.13 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (2.21 DPS) | yes | Woven Ivy Necklace (19159, -0.46 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.17 DPS) [world]; Ethereal Talisman (4430, -1.23 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.4 attack_power points (2.28 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.05 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.10 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.7 attack_power points (2.05 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Pridelord Cape (14673, -0.56 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 10 · Stone Guard · Horde [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Warbear Harness (15064, -1.08 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.23 DPS) [quest]; Charred Leather Tunic (19127, -2.17 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.7 attack_power points (2.39 DPS) | yes | Pridelord Bands (14672, -0.60 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.66 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.73 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.5 attack_power points (2.89 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.21 DPS) [crafted]; Serpentskin Gloves (8260, -0.44 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.46 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 34.1 attack_power points (2.95 DPS) | yes | Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted]; Sagebrush Girdle (17778, -0.90 DPS) [quest]; Ogron's Sash (13117, -0.94 DPS) [world_drop] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Serpentskin Leggings (8262, -1.22 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -1.76 DPS) [world_drop]; Triprunner Dungarees (9624, -1.91 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.6 attack_power points (2.73 DPS) | yes | Skulker's Leather Boots (252469, -0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.07 DPS) | yes | Mark of Kern (2262, -0.35 DPS) [dungeon]; Assault Band (13095, -0.35 DPS) [world_drop]; Blackstone Ring (17713, -0.35 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.1 attack_power points (1.91 DPS) | yes | Mark of Kern (2262, -0.18 DPS) [dungeon]; Assault Band (13095, -0.18 DPS) [world_drop]; Blackstone Ring (17713, -0.18 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (171.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (171.1 DPS) | yes | Molten Heart of the Mountain (249470, -0.76 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (171.1 DPS) | yes | Julie's Dagger (6660, -3.61 DPS) [world_drop]; Hanzo Sword (8190, -3.86 DPS) [world_drop]; Viking Warhammer (1721, -4.03 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (45.50 DPS) | yes | Julie's Dagger (6660, -3.53 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.83 DPS) [dungeon]; White Bone Shredder (11863, -6.60 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (171.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.44 DPS) [world_drop]; Guttbuster (13139, -0.56 DPS) [world_drop]; Dark Iron Rifle (16004, -2.24 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 701, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 251.2. Weights run: 1.8s. Verify run: 4.3s. 1462 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.815 ± 0.019, crit=1.646 ± 0.030 per rating point (14 rating = 1%, 23.048 per %), hit=3.401 ± 0.116 per rating point (10 rating = 1%, 34.008 per %), melee_haste=9.014 ± 1.758

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.7 attack_power points (6.08 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -2.38 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.6 attack_power points (2.39 DPS) | yes | Woven Ivy Necklace (19159, -0.52 DPS) [quest]; Sunprism Pendant (20649, -0.57 DPS) [quest]; Ethereal Talisman (4430, -1.37 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 46.7 attack_power points (3.91 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -11.39 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.2 attack_power points (3.95 DPS) | yes | Fel Cape (279269, -2.02 DPS) [crafted]; Blackflame Cape (13109, -2.13 DPS) [world_drop]; Dark Phantom Cape (13122, -2.18 DPS, sim-verified) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.2 attack_power points (10.39 DPS) | yes | Warlord's Leather Breastplate (231549, -1.67 DPS) [pvp]; Timbermaw Tunic (252484, -2.24 DPS, sim-verified) [crafted]; Legionnaire's Leather Chestpiece (227060, -2.77 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (251.2 DPS) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.33 DPS) [crafted]; Deepfury Bracers (13120, -10.72 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (251.2 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -4.64 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (251.2 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -9.88 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (251.2 DPS) | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 48.3 attack_power points (4.04 DPS) | yes | General's Leather Treads (231552, -0.70 DPS) [pvp]; Skulker's Leather Boots (252469, -1.15 DPS) [crafted]; Albino Crocscale Boots (17728, -12.12 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 25.4 attack_power points (2.12 DPS) | yes | Pyrewood Signet Ring (277210, -0.38 DPS) [quest]; Ring of the Underwood (2951, -0.44 DPS) [world_drop]; Blackstone Ring (17713, -0.45 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.01 DPS) | yes | Pyrewood Signet Ring (277210, -0.26 DPS) [quest]; Ring of the Underwood (2951, -0.32 DPS) [world_drop]; Blackstone Ring (17713, -0.33 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (251.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (251.2 DPS) | yes | Molten Heart of the Mountain (249470, -1.71 DPS, sim-verified) [crafted] |
| main_hand | Greenhammer (279261) | Blacksmithing [crafted] | sim-verified (251.2 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | 704.1 attack_power points (58.86 DPS) | yes | Persuader (22384, +0.00 DPS, sim-verified) [crafted]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -18.55 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (251.2 DPS) | yes | Eaglehorn Long Bow (13023, -0.61 DPS) [world_drop]; Guttbuster (13139, -0.66 DPS) [world_drop]; Dark Iron Rifle (16004, -2.53 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; main_hand: Greenhammer; off_hand: Ironwood Blade

No-known-source sample (15 of 1462, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 660.7. Weights run: 1.9s. Verify run: 4.6s. 1462 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.200 ± 0.021, crit=2.108 ± 0.035 per rating point (14 rating = 1%, 29.512 per %), hit=3.859 ± 0.162 per rating point (10 rating = 1%, 38.590 per %), melee_haste=16.559 ± 2.071

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 90.5 attack_power points (15.85 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -6.86 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 34.1 attack_power points (5.97 DPS) | yes | Sunprism Pendant (20649, -1.35 DPS) [quest]; Woven Ivy Necklace (19159, -1.35 DPS) [quest]; Ethereal Talisman (4430, -3.47 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (660.7 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -48.34 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 53.0 attack_power points (9.28 DPS) | yes | Fel Cape (279269, -4.11 DPS) [crafted]; Dark Phantom Cape (13122, -4.65 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.66 DPS) [world_drop] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (660.7 DPS) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -35.16 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 37.4 attack_power points (6.55 DPS) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Shadowcraft Bracers (16710, -0.77 DPS) [dungeon]; Arena Wristguards (18709, -1.38 DPS) [world] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 57.5 attack_power points (10.07 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Blood Guard's Leather Vices (16499, -0.62 DPS) [pvp]; Timbermaw Brawlers (19049, -0.63 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 53.9 attack_power points (9.44 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.73 DPS) [crafted]; Legionnaire's Leather Girdle (16500, -1.93 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (660.7 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -35.02 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 57.8 attack_power points (10.12 DPS) | yes | General's Leather Treads (231552, -1.64 DPS) [pvp]; Hivethrasher's Boots (275615, -2.80 DPS) [crafted]; Albino Crocscale Boots (17728, -4.86 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 30.8 attack_power points (5.40 DPS) | yes | Ring of the Underwood (2951, -1.16 DPS) [world_drop]; White Bone Band (11862, -1.19 DPS) [quest]; Falcon's Hook (7552, -1.35 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 24.2 attack_power points (4.25 DPS) | yes | Ring of the Underwood (2951, -0.01 DPS) [world_drop]; White Bone Band (11862, -0.04 DPS) [quest]; Falcon's Hook (7552, -0.20 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (660.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (660.7 DPS) | yes | Molten Heart of the Mountain (249470, -5.13 DPS, sim-verified) [crafted] |
| main_hand | Greenhammer (279261) | Blacksmithing [crafted] | sim-verified (660.7 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -4.11 DPS, sim-verified) [world_drop] |
| off_hand | Ironwood Blade (279259) | Blacksmithing [crafted] | 709.5 attack_power points (124.27 DPS) | yes | Persuader (22384, +0.00 DPS) [crafted]; High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Claw of Celebras (17738, -39.81 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (660.7 DPS) | yes | Eaglehorn Long Bow (13023, -1.54 DPS) [world_drop]; Guttbuster (13139, -1.73 DPS) [world_drop]; Dark Iron Rifle (16004, -6.58 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: Pyrewood Signet Ring; main_hand: Greenhammer; off_hand: Ironwood Blade

No-known-source sample (15 of 1462, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

