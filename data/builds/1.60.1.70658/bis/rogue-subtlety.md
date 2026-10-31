# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 36.6. Weights run: 1.4s. Verify run: 2.8s. 209 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.790 ± 0.013, crit=0.510 ± 0.012 per rating point (14 rating = 1%, 7.141 per %), hit=0.659 ± 0.029 per rating point (10 rating = 1%, 6.591 per %), melee_haste=3.047 ± 0.329

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.3 attack_power points (0.92 DPS) | yes | Defender's Leather Hood (252447, -0.36 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.6 attack_power points (0.23 DPS) | yes | Tarnished Locket (279870, -0.22 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.55 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 10.7 attack_power points (0.69 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.22 DPS) [crafted]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.7 attack_power points (1.27 DPS) | yes | Brawler's Leather Armor (252490, -0.14 DPS) [crafted]; Prospector's Chestpiece (14562, -0.45 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.46 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.2 attack_power points (0.59 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.83 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.44 DPS) [crafted]; Dusty Belt (279897, -0.58 DPS) [quest]; Deviate Scale Belt (6468, -3.82 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.03 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (36.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.95 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.8 attack_power points (0.63 DPS) | yes | Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; The 1 Ring (8350, -0.45 DPS) [world]; Ring of the Moon (12052, -0.57 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.0 attack_power points (0.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; The 1 Ring (8350, -0.08 DPS) [world]; Ring of the Moon (12052, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.71 DPS) | yes | Evocator's Blade (2567, -0.92 DPS) [dungeon]; Buzzer Blade (2169, -2.10 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.74 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.6 attack_power points (14.64 DPS) | yes | Evocator's Blade (2567, -3.93 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.2 attack_power points (0.46 DPS) | yes | Fine Longbow (11304, -0.18 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.23 DPS) [crafted]; Light Bow (4576, -0.23 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Demon Band; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 58.2. Weights run: 1.4s. Verify run: 3.0s. 403 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.452 ± 0.011, crit=0.479 ± 0.012 per rating point (14 rating = 1%, 6.710 per %), hit=0.911 ± 0.035 per rating point (10 rating = 1%, 9.108 per %), melee_haste=7.054 ± 0.404

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 14.5 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world]; Defender's Leather Helm (252455, -0.29 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.78 DPS) | yes | Kaleidoscope Chain (13084, -0.14 DPS) [world_drop]; River Pride Choker (13087, -0.52 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.58 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.44 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.57 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.2 attack_power points (0.85 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 20.3 attack_power points (1.32 DPS) | yes | Brawler's Leather Tunic (252508, -0.25 DPS, sim-verified) [crafted]; Raptorbane Armor (3566, -0.28 DPS) [quest]; Tunic of Westfall (2041, -0.28 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.7 attack_power points (0.83 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.25 DPS, sim-verified) [quest]; Demonhide Bracers (270033, -0.25 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.2 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -2.91 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (58.2 DPS) | yes | Prowler's Leather Belt (252459, -0.02 DPS) [crafted]; Warden's Leather Belt (252460, -0.21 DPS) [crafted]; Skulker's Leather Belt (252520, -3.41 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.2 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.90 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (58.2 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Disjointed Shoes (277226, +0.00 DPS) [quest]; Feet of the Lynx (1121, -3.16 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.4 attack_power points (0.80 DPS) | yes | Ring of Precision (1491, -0.24 DPS) [dungeon]; Seal of Wrynn (2933, -0.32 DPS) [quest]; Monkey Ring (6748, -0.33 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.5 attack_power points (0.61 DPS) | yes | Ring of Precision (1491, -0.05 DPS) [dungeon]; Seal of Wrynn (2933, -0.14 DPS) [quest]; Monkey Ring (6748, -0.14 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.91 DPS) | yes | Thornspike (6681, -2.19 DPS) [dungeon]; Claw of the Shadowmancer (2912, -2.25 DPS) [world_drop]; Toxic Revenger (9453, -2.25 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | 296.5 attack_power points (19.26 DPS) | yes | Thornspike (6681, -0.27 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.7 attack_power points (0.57 DPS) | yes | Gravespike Repeater (273033, -0.05 DPS) [world]; Double-barreled Shotgun (2098, -0.09 DPS) [world_drop]; Silver Star (3463, -0.19 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Thunderbrow Ring; finger2: Pyrewood Signet Ring; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Golemsight Long Gun

No-known-source sample (15 of 403, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14145 Cursed Felblade

### Band 40 (night-elf, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 96.4. Weights run: 1.6s. Verify run: 1.8s. 553 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.287 ± 0.009, crit=0.369 ± 0.011 per rating point (14 rating = 1%, 5.165 per %), hit=0.968 ± 0.037 per rating point (10 rating = 1%, 9.679 per %), melee_haste=4.905 ± 0.444

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.2 attack_power points (1.86 DPS) | yes | Hawkeye's Helm (14591, -0.60 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.62 DPS) [world_drop]; Nightscape Headband (8176, -0.72 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.89 DPS) | yes | Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; River Pride Choker (13087, -0.59 DPS) [world_drop]; Gazlowe's Charm (13088, -0.59 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 19.2 attack_power points (1.41 DPS) | yes | Flintrock Shoulders (7755, -0.09 DPS) [dungeon]; Nightscape Shoulders (8192, -0.37 DPS) [crafted]; Mantle of Thieves (2264, -0.46 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.7 attack_power points (1.01 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.22 DPS) [quest]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 21.9 attack_power points (1.61 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.34 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Hawkeye's Bracers (14590, -0.50 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.72 DPS) [dungeon]; Dusky Bracers (7378, -0.72 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 25.2 attack_power points (1.86 DPS) | yes | Skulker's Leather Gloves (252525, -0.24 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.26 DPS) [crafted]; Dog Training Gloves (7756, -0.63 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.6 attack_power points (1.52 DPS) | yes | Skulker's Leather Belt (252520, -0.22 DPS) [crafted]; Scorpashi Sash (14652, -0.27 DPS) [world_drop]; Blackened Defias Belt (10403, -0.50 DPS, sim-verified) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.0 attack_power points (1.99 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.62 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.2 attack_power points (1.56 DPS) | yes | Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon]; Thunderbrow Ring (13097, -0.60 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon]; Thunderbrow Ring (13097, -0.60 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (96.4 DPS) | yes | Hypnotic Blade (7714, -7.34 DPS) [dungeon]; The Ziggler (8006, -7.67 DPS) [world_drop]; Coldrage Dagger (10761, -7.76 DPS, sim-verified) [dungeon] |
| off_hand | Black Menace (6831) (or Coldrage Dagger (10761)) | In the Name of the Light [quest] | 415.4 attack_power points (30.66 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (96.4 DPS) | yes | Monolithic Bow (9426, -0.31 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.36 DPS) [quest]; Bow of Searing Arrows (2825, -1.19 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 553, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 131.3. Weights run: 2.2s. Verify run: 2.1s. 704 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.417 ± 0.007, crit=0.701 ± 0.011 per rating point (14 rating = 1%, 9.811 per %), hit=1.584 ± 0.067 per rating point (10 rating = 1%, 15.837 per %), melee_haste=5.880 ± 0.944

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.6 attack_power points (2.25 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.21 DPS) [crafted]; Scorpashi Skullcap (14658, -0.32 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.4 attack_power points (1.98 DPS) | yes | Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.18 DPS, sim-verified) [world]; River Pride Choker (13087, -1.64 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 24.6 attack_power points (2.08 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.07 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 21.2 attack_power points (1.80 DPS) | yes | Blackflame Cape (13109, -0.36 DPS) [world_drop]; Pridelord Cape (14673, -0.47 DPS) [world_drop]; Serpentskin Cloak (8259, -0.48 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Warbear Harness (15064, -0.61 DPS) [crafted]; Blazewind Breastplate (11193, -0.68 DPS) [quest]; Charred Leather Tunic (19127, -1.54 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.2 attack_power points (2.14 DPS) | yes | Skulker's Leather Bracers (252540, -0.47 DPS) [crafted]; Pridelord Bands (14672, -0.52 DPS) [world_drop]; Branded Leather Bracers (19508, -0.61 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.8 attack_power points (2.52 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.08 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.26 DPS) [crafted]; Serpentskin Gloves (8260, -0.27 DPS) [world_drop] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 31.8 attack_power points (2.70 DPS) | yes | Prowler's Leather Waistguard (252473, -0.07 DPS) [crafted]; Ogron's Sash (13117, -0.85 DPS) [world_drop]; Sagebrush Girdle (17778, -0.90 DPS) [quest] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Gryphon Rider's Leggings (9652, -0.63 DPS) [quest]; Serpentskin Leggings (8262, -0.68 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.18 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 29.4 attack_power points (2.49 DPS) | yes | Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.09 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.58 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.01 DPS) [quest]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.01 DPS) [quest]; Ring of the Underwood (2951, -0.32 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (131.3 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (131.3 DPS) | yes | Mark of the Chosen (17774, -0.48 DPS, sim-verified) [quest] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (131.3 DPS) | yes | Charstone Dirk (17710, +0.00 DPS) [dungeon]; Widowmaker (4091, -1.19 DPS) [world_drop]; Julie's Dagger (6660, -5.92 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.60 DPS) | yes | Charstone Dirk (17710, -25.57 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (131.3 DPS) | yes | Skull Splitting Crossbow (13039, -0.24 DPS) [world_drop]; Houndmaster's Bow (11628, -0.41 DPS) [dungeon]; Dark Iron Rifle (16004, -2.01 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 215.4. Weights run: 2.1s. Verify run: 5.1s. 1463 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.575 ± 0.012, crit=1.203 ± 0.020 per rating point (14 rating = 1%, 16.841 per %), hit=1.913 ± 0.086 per rating point (10 rating = 1%, 19.127 per %), melee_haste=8.267 ± 1.157

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 58.3 attack_power points (5.89 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -1.26 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (2.57 DPS) | yes | Sunprism Pendant (20649, -0.76 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -1.36 DPS) [world]; Kaleidoscope Chain (13084, -1.53 DPS) [world_drop] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 38.5 attack_power points (3.89 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Dark Warder's Pauldrons (22241, -10.04 DPS, sim-verified) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 43.6 attack_power points (4.41 DPS) | yes | Sergeant Major's Cape (16337, -2.33 DPS) [pvp]; Dark Phantom Cape (13122, -2.33 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.50 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 94.3 attack_power points (9.52 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -1.75 DPS) [pvp]; Timbermaw Tunic (252484, -1.79 DPS, sim-verified) [crafted]; Knight-Captain's Leather Chestpiece (227058, -2.45 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (215.4 DPS) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -9.19 DPS, sim-verified) [world_drop] |
| hands | Shadowcraft Gloves (16712) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (215.4 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.72 DPS, sim-verified) [crafted] |
| waist | Shadowcraft Belt (16713) | Blackrock Spire: Blackhand Assassin [dungeon] | sim-verified (215.4 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted]; Might of the Timbermaw (19044, -8.30 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 62.8 attack_power points (6.35 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 38.8 attack_power points (3.92 DPS) | yes | Marshal's Leather Footguards (231546, -0.42 DPS) [pvp]; Skulker's Leather Boots (252469, -0.74 DPS) [crafted]; Albino Crocscale Boots (17728, -10.55 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.1 attack_power points (2.23 DPS) | yes | Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop]; Ring of the Underwood (2951, -0.43 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (2.02 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.23 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (215.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (215.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (215.4 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -11.76 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 594.4 attack_power points (60.02 DPS) | yes | The Shadowfoot Stabber (24222, -2.98 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (215.4 DPS) | yes | Skull Splitting Crossbow (13039, -0.51 DPS) [world_drop]; Eaglehorn Long Bow (13023, -0.64 DPS) [world_drop]; Dark Iron Rifle (16004, -2.42 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Shadowcraft Gloves; waist: Shadowcraft Belt; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: Blackstone Ring; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1463, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 524.0. Weights run: 2.2s. Verify run: 3.6s. 1463 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.009 ± 0.017, crit=1.706 ± 0.028 per rating point (14 rating = 1%, 23.888 per %), hit=2.826 ± 0.148 per rating point (10 rating = 1%, 28.260 per %), melee_haste=13.252 ± 1.654

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 78.0 attack_power points (14.30 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -4.04 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (5.80 DPS) | yes | Sunprism Pendant (20649, -1.38 DPS) [quest]; Kaleidoscope Chain (13084, -3.52 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.60 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 49.9 attack_power points (9.14 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.1 attack_power points (9.19 DPS) | yes | Dark Phantom Cape (13122, -4.14 DPS, sim-verified) [world_drop]; Sergeant Major's Cape (16337, -4.63 DPS) [pvp]; Blackflame Cape (13109, -4.77 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 126.5 attack_power points (23.19 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -4.05 DPS) [pvp]; Timbermaw Tunic (252484, -4.37 DPS, sim-verified) [crafted]; Knight-Captain's Leather Chestpiece (227058, -7.39 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -5.29 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 51.9 attack_power points (9.51 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Timbermaw Brawlers (19049, -0.12 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 51.2 attack_power points (9.39 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, -2.05 DPS, sim-verified) [crafted]; Knight-Captain's Leather Belt (16398, -2.05 DPS) [pvp] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (+11.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight's Leather Pants (220858, -11.75 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 50.7 attack_power points (9.29 DPS) | yes | Marshal's Leather Footguards (231546, -1.19 DPS) [pvp]; Skulker's Leather Boots (252469, -2.28 DPS) [crafted]; Albino Crocscale Boots (17728, -3.68 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.1 attack_power points (5.16 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Mark of Kern (2262, -1.49 DPS) [dungeon]; Blackstone Ring (17713, -1.49 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 22.3 attack_power points (4.09 DPS) | yes | Falcon's Hook (7552, -0.17 DPS) [dungeon]; Mark of Kern (2262, -0.42 DPS) [dungeon]; Blackstone Ring (17713, -0.42 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -24.94 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 598.7 attack_power points (109.75 DPS) | yes | The Shadowfoot Stabber (24222, -4.78 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.47 DPS) [world_drop]; Guttbuster (13139, -1.60 DPS) [world_drop]; Dark Iron Rifle (16004, -5.18 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: Ring of the Underwood; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1463, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 36.3. Weights run: 1.4s. Verify run: 2.8s. 208 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.790 ± 0.013, crit=0.510 ± 0.012 per rating point (14 rating = 1%, 7.141 per %), hit=0.659 ± 0.029 per rating point (10 rating = 1%, 6.591 per %), melee_haste=3.047 ± 0.329

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.3 attack_power points (0.92 DPS) | yes | Defender's Leather Hood (252447, -0.37 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.6 attack_power points (0.23 DPS) | yes | Roadwatcher's Confidence (281265, -0.23 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 10.7 attack_power points (0.69 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.22 DPS) [crafted]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (36.3 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.87 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.9 attack_power points (0.58 DPS) | yes | Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.23 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.7 attack_power points (0.95 DPS) | yes | Brawler's Leather Gloves (252494, -0.23 DPS) [crafted]; Bristlebark Gloves (14572, -0.23 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.26 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.44 DPS) [crafted]; Murloc Scale Belt (5780, -0.68 DPS) [crafted]; Deviate Scale Belt (6468, -3.84 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (36.3 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.04 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (36.3 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.97 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.8 attack_power points (0.63 DPS) | yes | Demon Band (12054, -0.37 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; Loop of Sacrifice (281673, -0.44 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 5.4 attack_power points (0.35 DPS) | yes | Demon Band (12054, -0.09 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Loop of Sacrifice (281673, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.71 DPS) | yes | Evocator's Blade (2567, -0.92 DPS) [dungeon]; Buzzer Blade (2169, -2.10 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.74 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.6 attack_power points (14.64 DPS) | yes | Evocator's Blade (2567, -3.84 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.2 attack_power points (0.46 DPS) | yes | Fine Longbow (11304, -0.19 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.23 DPS) [crafted]; Light Bow (4576, -0.23 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Bounty Hunter's Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 208, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 58.0. Weights run: 1.4s. Verify run: 2.9s. 402 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.452 ± 0.011, crit=0.479 ± 0.012 per rating point (14 rating = 1%, 6.710 per %), hit=0.911 ± 0.035 per rating point (10 rating = 1%, 9.108 per %), melee_haste=7.054 ± 0.404

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 14.5 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world]; Defender's Leather Helm (252455, -0.30 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.78 DPS) | yes | Kaleidoscope Chain (13084, -0.14 DPS) [world_drop]; River Pride Choker (13087, -0.52 DPS) [world_drop]; Erudite's Amulet (277204, -0.59 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.45 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.57 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.2 attack_power points (0.85 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon]; Wildhunter Cloak (16658, -0.21 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 20.3 attack_power points (1.32 DPS) | yes | Brawler's Leather Tunic (252508, -0.25 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.34 DPS) [crafted]; Panther Armor (6670, -0.34 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.7 attack_power points (0.83 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.25 DPS, sim-verified) [quest]; Insignia Bracers (6410, -0.26 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.0 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -2.91 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (58.0 DPS) | yes | Deftkin Belt (16659, -0.01 DPS) [quest]; Prowler's Leather Belt (252459, -0.02 DPS) [crafted]; Skulker's Leather Belt (252520, -3.38 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.0 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.86 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (58.0 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.15 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.4 attack_power points (0.80 DPS) | yes | Pyrewood Signet Ring (277210, -0.19 DPS) [quest]; Ring of Precision (1491, -0.24 DPS) [dungeon]; Monkey Ring (6748, -0.33 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 9.8 attack_power points (0.64 DPS) | yes | Pyrewood Signet Ring (277210, -0.02 DPS) [quest]; Ring of Precision (1491, -0.07 DPS) [dungeon]; Monkey Ring (6748, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.91 DPS) | yes | Thornspike (6681, -2.19 DPS) [dungeon]; Claw of the Shadowmancer (2912, -2.25 DPS) [world_drop]; Toxic Revenger (9453, -2.25 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | 296.5 attack_power points (19.26 DPS) | yes | Thornspike (6681, -0.27 DPS, sim-verified) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.7 attack_power points (0.57 DPS) | yes | Gravespike Repeater (273033, -0.05 DPS) [world]; Double-barreled Shotgun (2098, -0.09 DPS) [world_drop]; Silver Star (3463, -0.19 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Golemsight Long Gun

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 96.4. Weights run: 1.6s. Verify run: 1.8s. 550 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.287 ± 0.009, crit=0.369 ± 0.011 per rating point (14 rating = 1%, 5.165 per %), hit=0.968 ± 0.037 per rating point (10 rating = 1%, 9.679 per %), melee_haste=4.905 ± 0.444

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.2 attack_power points (1.86 DPS) | yes | Hawkeye's Helm (14591, -0.59 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.62 DPS) [world_drop]; Nightscape Headband (8176, -0.72 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.89 DPS) | yes | Ethereal Talisman (4430, -0.14 DPS) [quest]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; River Pride Choker (13087, -0.59 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 19.2 attack_power points (1.41 DPS) | yes | Flintrock Shoulders (7755, -0.09 DPS) [dungeon]; Nightscape Shoulders (8192, -0.37 DPS) [crafted]; Mantle of Thieves (2264, -0.46 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.0 attack_power points (0.89 DPS) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 21.9 attack_power points (1.61 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.34 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Hawkeye's Bracers (14590, -0.53 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.72 DPS) [dungeon]; Dusky Bracers (7378, -0.72 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 25.2 attack_power points (1.86 DPS) | yes | Skulker's Leather Gloves (252525, -0.24 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.26 DPS) [crafted]; Dog Training Gloves (7756, -0.70 DPS, sim-verified) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.6 attack_power points (1.52 DPS) | yes | Skulker's Leather Belt (252520, -0.22 DPS) [crafted]; Deftkin Belt (16659, -0.25 DPS) [quest]; Blackened Defias Belt (10403, -0.52 DPS, sim-verified) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.0 attack_power points (1.99 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.62 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.2 attack_power points (1.56 DPS) | yes | Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon]; Thunderbrow Ring (13097, -0.60 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon]; Thunderbrow Ring (13097, -0.60 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.5 attack_power points (35.02 DPS) | yes | Hypnotic Blade (7714, -7.34 DPS) [dungeon]; The Ziggler (8006, -7.67 DPS) [world_drop]; Sacrificial Kris (3187, -8.45 DPS) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 415.4 attack_power points (30.66 DPS) | yes | Hypnotic Blade (7714, -26.31 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (96.4 DPS) | yes | Monolithic Bow (9426, -0.31 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.36 DPS) [quest]; Bow of Searing Arrows (2825, -1.18 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 550, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 136.1. Weights run: 2.2s. Verify run: 2.1s. 701 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.417 ± 0.007, crit=0.701 ± 0.011 per rating point (14 rating = 1%, 9.811 per %), hit=1.584 ± 0.067 per rating point (10 rating = 1%, 15.837 per %), melee_haste=5.880 ± 0.944

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.6 attack_power points (2.25 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.08 DPS) [quest]; Helm of Fire (8348, -0.21 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.4 attack_power points (1.98 DPS) | yes | Woven Ivy Necklace (19159, -0.49 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -0.97 DPS) [world]; Ethereal Talisman (4430, -1.08 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 24.6 attack_power points (2.08 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.07 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 21.2 attack_power points (1.80 DPS) | yes | Blackflame Cape (13109, -0.36 DPS) [world_drop]; Pridelord Cape (14673, -0.47 DPS) [world_drop]; Serpentskin Cloak (8259, -0.48 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 10 · Stone Guard · Horde [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Blazewind Breastplate (11193, -0.68 DPS) [quest]; Charred Leather Tunic (19127, -1.54 DPS) [quest]; Warbear Harness (15064, -1.63 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.2 attack_power points (2.14 DPS) | yes | Skulker's Leather Bracers (252540, -0.47 DPS) [crafted]; Pridelord Bands (14672, -0.52 DPS) [world_drop]; Branded Leather Bracers (19508, -0.65 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.8 attack_power points (2.52 DPS) | yes | Skulker's Leather Gauntlets (252548, -0.08 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.26 DPS) [crafted]; Serpentskin Gloves (8260, -0.27 DPS) [world_drop] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 31.8 attack_power points (2.70 DPS) | yes | Prowler's Leather Waistguard (252473, -0.07 DPS) [crafted]; Ogron's Sash (13117, -0.85 DPS) [world_drop]; Sagebrush Girdle (17778, -0.90 DPS) [quest] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Basilisk Hide Pants (1718, -1.18 DPS) [world_drop]; Triprunner Dungarees (9624, -1.28 DPS) [quest]; Serpentskin Leggings (8262, -1.70 DPS, sim-verified) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 29.4 attack_power points (2.49 DPS) | yes | Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.09 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.58 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.03 DPS) | yes | Mark of Kern (2262, -0.34 DPS) [dungeon]; Assault Band (13095, -0.34 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.35 DPS) [quest] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.69 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.01 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (136.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.1 DPS) | yes | Molten Heart of the Mountain (249470, -1.78 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (136.1 DPS) | yes | Charstone Dirk (17710, +0.00 DPS) [dungeon]; Widowmaker (4091, -1.19 DPS) [world_drop]; Julie's Dagger (6660, -6.89 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.60 DPS) | yes | Charstone Dirk (17710, -26.59 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (136.1 DPS) | yes | Skull Splitting Crossbow (13039, -0.24 DPS) [world_drop]; Houndmaster's Bow (11628, -0.41 DPS) [dungeon]; Dark Iron Rifle (16004, -2.03 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 701, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 218.4. Weights run: 2.1s. Verify run: 3.5s. 1462 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.575 ± 0.012, crit=1.203 ± 0.020 per rating point (14 rating = 1%, 16.841 per %), hit=1.913 ± 0.086 per rating point (10 rating = 1%, 19.127 per %), melee_haste=8.267 ± 1.157

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 58.3 attack_power points (5.89 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -1.11 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (2.57 DPS) | yes | Woven Ivy Necklace (19159, -0.54 DPS) [quest]; Sunprism Pendant (20649, -0.66 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.36 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 38.5 attack_power points (3.89 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 43.6 attack_power points (4.41 DPS) | yes | Dark Phantom Cape (13122, -2.30 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -2.50 DPS) [world_drop]; Deep River Cloak (15789, -2.66 DPS) [quest] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 94.3 attack_power points (9.52 DPS) | yes | Warlord's Leather Breastplate (231549, -1.75 DPS) [pvp]; Timbermaw Tunic (252484, -1.89 DPS, sim-verified) [crafted]; Legionnaire's Leather Chestpiece (227060, -2.45 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (218.4 DPS) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.25 DPS) [crafted]; Deepfury Bracers (13120, -2.99 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 44.8 attack_power points (4.53 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Timbermaw Brawlers (19049, -4.34 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 43.1 attack_power points (4.35 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Prowler's Leather Waistguard (252473, -1.03 DPS) [crafted]; Skulker's Leather Waistguard (252474, -1.03 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 62.8 attack_power points (6.35 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.89 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 38.8 attack_power points (3.92 DPS) | yes | General's Leather Treads (231552, -0.42 DPS) [pvp]; Skulker's Leather Boots (252469, -0.74 DPS) [crafted]; Albino Crocscale Boots (17728, -0.86 DPS, sim-verified) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.42 DPS) | yes | Mark of Kern (2262, -0.40 DPS) [dungeon]; Assault Band (13095, -0.40 DPS) [world_drop]; Blackstone Ring (17713, -0.40 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.1 attack_power points (2.23 DPS) | yes | Mark of Kern (2262, -0.21 DPS) [dungeon]; Assault Band (13095, -0.21 DPS) [world_drop]; Blackstone Ring (17713, -0.21 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+6.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -1.50 DPS, sim-verified) [crafted] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -9.64 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 594.4 attack_power points (60.02 DPS) | yes | The Shadowfoot Stabber (24222, -2.36 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.51 DPS) [world_drop]; Eaglehorn Long Bow (13023, -0.64 DPS) [world_drop]; Dark Iron Rifle (16004, -2.34 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Masons Fraternity Ring; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1462, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 532.9. Weights run: 2.2s. Verify run: 3.7s. 1462 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.009 ± 0.017, crit=1.706 ± 0.028 per rating point (14 rating = 1%, 23.888 per %), hit=2.826 ± 0.148 per rating point (10 rating = 1%, 28.260 per %), melee_haste=13.252 ± 1.654

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 78.0 attack_power points (14.30 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Blue Suede Hat (252482, -4.64 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.6 attack_power points (5.80 DPS) | yes | Woven Ivy Necklace (19159, -1.27 DPS) [quest]; Sunprism Pendant (20649, -1.38 DPS) [quest]; Ethereal Talisman (4430, -3.31 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 49.9 attack_power points (9.14 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.1 attack_power points (9.19 DPS) | yes | Dark Phantom Cape (13122, -4.14 DPS, sim-verified) [world_drop]; Blackflame Cape (13109, -4.77 DPS) [world_drop]; Fel Cape (279269, -4.81 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 126.5 attack_power points (23.19 DPS) | yes | Timbermaw Tunic (252484, -3.84 DPS, sim-verified) [crafted]; Warlord's Leather Breastplate (231549, -4.05 DPS) [pvp]; Legionnaire's Leather Chestpiece (227060, -7.39 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.80 DPS) [crafted]; Deepfury Bracers (13120, -5.29 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 51.9 attack_power points (9.51 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Timbermaw Brawlers (19049, -0.12 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 51.2 attack_power points (9.39 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Skulker's Leather Waistguard (252474, -2.05 DPS, sim-verified) [crafted]; Legionnaire's Leather Girdle (16500, -2.05 DPS) [pvp] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Stone Guard's Leather Pants (220859, -9.65 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 50.7 attack_power points (9.29 DPS) | yes | General's Leather Treads (231552, -1.19 DPS) [pvp]; Skulker's Leather Boots (252469, -2.28 DPS) [crafted]; Albino Crocscale Boots (17728, -2.58 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.1 attack_power points (5.16 DPS) | yes | Ring of the Underwood (2951, -1.07 DPS) [world_drop]; Falcon's Hook (7552, -1.24 DPS) [dungeon]; Blackstone Ring (17713, -1.49 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (4.40 DPS) | yes | Ring of the Underwood (2951, -0.31 DPS) [world_drop]; Falcon's Hook (7552, -0.48 DPS) [dungeon]; Blackstone Ring (17713, -0.73 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+12.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -2.72 DPS, sim-verified) [crafted] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -24.27 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 598.7 attack_power points (109.75 DPS) | yes | The Shadowfoot Stabber (24222, -2.59 DPS, sim-verified) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eaglehorn Long Bow (13023, -1.47 DPS) [world_drop]; Guttbuster (13139, -1.60 DPS) [world_drop]; Dark Iron Rifle (16004, -5.75 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Masons Fraternity Ring; finger2: White Bone Band; main_hand: Heartseeker; off_hand: Scarlet Kris

No-known-source sample (15 of 1462, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

