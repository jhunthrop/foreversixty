# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 40.7. Weights run: 1.6s. Verify run: 3.0s. 198 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.974 ± 0.014, crit=0.569 ± 0.013 per rating point (14 rating = 1%, 7.966 per %), hit=0.686 ± 0.032 per rating point (10 rating = 1%, 6.856 per %), melee_haste=4.645 ± 0.418

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.8 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.8 attack_power points (0.76 DPS) | yes | Erudite's Amulet (277204, -0.55 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.9 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.8 attack_power points (0.76 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 21.7 attack_power points (1.40 DPS) | yes | Brawler's Leather Armor (252490, -0.19 DPS) [crafted]; Prospector's Chestpiece (14562, -0.51 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.51 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.9 attack_power points (0.64 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -1.96 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Dusty Belt (279897, -0.52 DPS) [quest]; Deviate Scale Belt (6468, -3.13 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (40.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.22 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (40.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.12 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.9 attack_power points (0.77 DPS) | yes | Demon Band (12054, -0.51 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; The 1 Ring (8350, -0.58 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.6 attack_power points (0.69 DPS) | yes | Demon Band (12054, -0.22 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Shadowfang (1482) | Shadowfang Keep [dungeon] | 248.9 attack_power points (16.05 DPS) | yes | Diamond Hammer (2194, +0.00 DPS, sim-verified) [world_drop]; Blackfang (2236, -1.31 DPS) [world_drop]; Assassin's Blade (1935, -1.33 DPS) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (40.7 DPS) | yes | Diamond Hammer (2194, -0.38 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.9 attack_power points (0.51 DPS) | yes | Fine Longbow (11304, -0.24 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 63.7. Weights run: 1.6s. Verify run: 3.3s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.714 ± 0.013, crit=0.667 ± 0.015 per rating point (14 rating = 1%, 9.342 per %), hit=0.897 ± 0.040 per rating point (10 rating = 1%, 8.973 per %), melee_haste=7.864 ± 0.505

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.1 attack_power points (1.14 DPS) | yes | Brawler's Leather Hood (252504, -0.23 DPS) [crafted]; Defender's Leather Helm (252455, -0.34 DPS) [crafted]; Humbert's Helm (4724, -0.34 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.7 attack_power points (0.91 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.11 DPS) [world]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; River Pride Choker (13087, -0.65 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.48 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.68 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.28 DPS) [pvp]; Fenrus' Hide (6340, -0.31 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.0 attack_power points (1.59 DPS) | yes | Brawler's Leather Tunic (252508, -0.30 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.34 DPS) [quest]; Brawler's Leather Armor (252490, -0.46 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.3 attack_power points (0.95 DPS) | yes | Barbaric Bracers (18948, -0.23 DPS) [crafted]; Insignia Bracers (6410, -0.27 DPS) [world_drop]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (63.7 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.42 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (63.7 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.78 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (63.7 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.53 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (63.7 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Feet of the Lynx (1121, -3.75 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.3 attack_power points (1.08 DPS) | yes | Pyrewood Signet Ring (277210, -0.39 DPS) [quest]; Ring of Precision (1491, -0.40 DPS) [dungeon]; Insurgent's Band (272067, -0.48 DPS) [vendor] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.87 DPS) | yes | Ring of Precision (1491, -0.19 DPS) [dungeon]; Insurgent's Band (272067, -0.28 DPS) [vendor]; Pyrewood Signet Ring (277210, -0.42 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (21.39 DPS) | yes | Zealot Blade (13033, -0.80 DPS) [world_drop]; Electrocutioner Leg (9446, -0.87 DPS) [dungeon]; Dreamslayer (7752, -1.00 DPS) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 attack_power points (21.15 DPS) | yes | Zealot Blade (13033, -2.65 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.3 attack_power points (0.68 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.09 DPS) [vendor]; Double-barreled Shotgun (2098, -0.14 DPS) [world_drop]; Gravespike Repeater (273033, -0.15 DPS) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 108.7. Weights run: 1.6s. Verify run: 1.7s. 493 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.551 ± 0.015, crit=0.716 ± 0.021 per rating point (14 rating = 1%, 10.030 per %), hit=1.873 ± 0.077 per rating point (10 rating = 1%, 18.732 per %), melee_haste=30.832 ± 0.898

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.1 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 17.1 attack_power points (0.59 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Kaleidoscope Chain (13084, -0.24 DPS) [world_drop]; River Pride Choker (13087, -0.45 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.1 attack_power points (1.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.30 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 15.3 attack_power points (0.53 DPS) | yes | Hawkeye's Cloak (14593, -0.05 DPS) [world_drop]; Imperial Cloak (6432, -0.10 DPS) [dungeon]; Parachute Cloak (10518, -0.10 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 26.4 attack_power points (0.91 DPS) | yes | Nightscape Tunic (8175, -0.11 DPS) [crafted]; Dusky Leather Armor (7374, -0.16 DPS) [crafted]; Barbaric Harness (5739, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.26 DPS) [dungeon]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.0 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Dog Training Gloves (7756, -0.21 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.04 DPS) | yes | Highlander's Chain Girdle (20090, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.35 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.6 attack_power points (1.12 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.23 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.37 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.1 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.4 attack_power points (0.70 DPS) | yes | Assault Band (13095, -0.01 DPS) [world_drop]; Ring of the Underwood (2951, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Vanquisher's Sword (10823, -1.12 DPS) [quest] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (108.7 DPS) | yes | Ardent Custodian (868, -62.68 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.11 DPS) [world_drop]; Monolithic Bow (9426, -0.12 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 493, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 172.4. Weights run: 2.4s. Verify run: 2.2s. 616 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.579 ± 0.012, crit=0.963 ± 0.018 per rating point (14 rating = 1%, 13.485 per %), hit=2.203 ± 0.082 per rating point (10 rating = 1%, 22.026 per %), melee_haste=5.914 ± 1.264

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 51.5 attack_power points (4.45 DPS) | yes | Helm of Fire (8348, -2.13 DPS) [crafted]; Lordrec Helmet (10741, -2.27 DPS) [quest]; White Bandit Mask (10008, -2.37 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (2.21 DPS) | yes | Sentinel's Medallion (19539, -0.65 DPS, sim-verified) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 43.5 attack_power points (3.76 DPS) | yes | Sunburn Spaulders (274751, -1.48 DPS, sim-verified) [vendor]; Skulker's Leather Shoulder (252535, -1.48 DPS) [crafted]; Failed Flying Experiment (9647, -1.53 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.7 attack_power points (2.05 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Pridelord Cape (14673, -0.56 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Blazewind Breastplate (11193, -1.23 DPS) [quest]; Warbear Harness (15064, -1.46 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -2.17 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.7 attack_power points (2.39 DPS) | yes | Pridelord Bands (14672, -0.60 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.66 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.67 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 44.7 attack_power points (3.87 DPS) | yes | Darkmantle Grips (226828, -0.09 DPS) [vendor]; Gloves of Holy Might (867, -0.97 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.18 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 34.1 attack_power points (2.95 DPS) | yes | Highlander's Leather Girdle (20115, -0.05 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Serpentskin Leggings (8262, -1.32 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.55 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.76 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.6 attack_power points (2.73 DPS) | yes | Skulker's Leather Boots (252469, -0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.2 attack_power points (2.09 DPS) | yes | Mark of Kern (2262, -0.36 DPS) [dungeon]; Assault Band (13095, -0.36 DPS) [world_drop]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 22.1 attack_power points (1.91 DPS) | yes | Mark of Kern (2262, -0.18 DPS) [dungeon]; Assault Band (13095, -0.18 DPS) [world_drop]; Blackstone Ring (17713, -0.18 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.4 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (172.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (172.4 DPS) | yes | Hanzo Sword (8190, -1.09 DPS, sim-verified) [world_drop]; Julie's Dagger (6660, -3.61 DPS) [world_drop]; Viking Warhammer (1721, -4.03 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (45.50 DPS) | yes | Julie's Dagger (6660, -3.07 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.83 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (172.4 DPS) | yes | Skull Splitting Crossbow (13039, -0.44 DPS) [world_drop]; Guttbuster (13139, -0.56 DPS) [world_drop]; Dark Iron Rifle (16004, -2.23 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 616, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 286.6. Weights run: 2.3s. Verify run: 6.2s. 1171 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.815 ± 0.019, crit=1.646 ± 0.030 per rating point (14 rating = 1%, 23.048 per %), hit=3.401 ± 0.116 per rating point (10 rating = 1%, 34.008 per %), melee_haste=9.014 ± 1.758

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 73.1 attack_power points (6.11 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 58.0 attack_power points (4.85 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -1.13 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.07 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (286.6 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -15.72 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 62.0 attack_power points (5.18 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Cloak of the Honor Guard (20073, -1.58 DPS) [rep]; Stalwart Cloak (272415, -2.34 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.2 attack_power points (10.39 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -1.67 DPS) [pvp]; Darkmantle Tunic (226825, -1.79 DPS, sim-verified) [quest]; Knight-Captain's Leather Chestpiece (227058, -2.77 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 45.5 attack_power points (3.80 DPS) | yes | Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Bracers of Subterfuge (22668, -0.69 DPS) [quest]; Deepfury Bracers (13120, -1.19 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (286.6 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -0.43 DPS) [quest]; Raider Gloves (272099, -18.76 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 106.1 attack_power points (8.87 DPS) | yes | Belt of Preserved Heads (20216, -2.58 DPS) [quest]; Highlander's Leather Girdle (20045, -4.10 DPS) [rep]; Ferocity of the Timbermaw (227805, -4.12 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (286.6 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.72 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 67.4 attack_power points (5.63 DPS) | yes | Fine Dawn Treaders (227815, -1.23 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.59 DPS) [dungeon]; Drudge Boots (21532, -1.97 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 77.6 attack_power points (6.49 DPS) | yes | Cutthroat's Signet (272408, -1.52 DPS) [vendor]; Channeler's Ring (272406, -3.64 DPS) [vendor]; Blessed Band of Light (272407, -3.64 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 71.1 attack_power points (5.94 DPS) | yes | Cutthroat's Signet (272408, -0.97 DPS) [vendor]; Channeler's Ring (272406, -3.10 DPS) [vendor]; Blessed Band of Light (272407, -3.10 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (286.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (286.6 DPS) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (286.6 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -4.29 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 832.7 attack_power points (69.61 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -4.37 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -29.30 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (286.6 DPS) | yes | The Purifier (22656, -0.20 DPS) [quest]; Bloodseeker (19107, -0.39 DPS) [quest]; Dark Iron Rifle (16004, -2.55 DPS, sim-verified) [crafted] |

**New at 60:** neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1171, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 738.6. Weights run: 2.4s. Verify run: 6.3s. 1171 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.200 ± 0.021, crit=2.108 ± 0.035 per rating point (14 rating = 1%, 29.512 per %), hit=3.859 ± 0.162 per rating point (10 rating = 1%, 38.590 per %), melee_haste=16.559 ± 2.071

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 90.5 attack_power points (15.85 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -1.04 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 62.6 attack_power points (10.96 DPS) | yes | Medallion of the Dawn (22659, -1.59 DPS) [quest]; Amulet of the Darkmoon (19491, -1.71 DPS) [quest]; Zandalarian Shadow Talisman (19616, -4.03 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (738.6 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -43.28 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.6 attack_power points (11.66 DPS) | yes | Cape of the Black Baron (13340, -2.38 DPS) [dungeon]; Cloak of the Honor Guard (20073, -3.78 DPS) [rep]; Stalwart Cloak (272415, -4.90 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 149.5 attack_power points (26.19 DPS) | yes | Darkmantle Tunic (226825, -3.55 DPS, sim-verified) [quest]; Field Marshal's Leather Chestpiece (231543, -4.24 DPS) [pvp]; Timbermaw Tunic (252484, -7.72 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 53.9 attack_power points (9.44 DPS) | yes | Marshal's Leather Armsplints (16460, -0.19 DPS) [pvp]; Bracers of Subterfuge (22668, -1.73 DPS) [quest]; Deepfury Bracers (13120, -2.89 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (738.6 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.14 DPS) [quest]; Raider Gloves (272099, -54.31 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 123.4 attack_power points (21.61 DPS) | yes | Belt of Preserved Heads (20216, -5.03 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -10.05 DPS) [vendor]; Highlander's Leather Girdle (20045, -10.49 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (738.6 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -43.28 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 79.8 attack_power points (13.98 DPS) | yes | Fine Dawn Treaders (227815, -2.98 DPS) [vendor]; Shadowcraft Boots (16711, -3.86 DPS) [dungeon]; Darkmantle Boots (22003, -4.73 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 91.4 attack_power points (16.01 DPS) | yes | Cutthroat's Signet (272408, -3.85 DPS) [vendor]; Channeler's Ring (272406, -9.25 DPS) [vendor]; Blessed Band of Light (272407, -9.25 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 82.1 attack_power points (14.38 DPS) | yes | Cutthroat's Signet (272408, -2.23 DPS) [vendor]; Channeler's Ring (272406, -7.62 DPS) [vendor]; Blessed Band of Light (272407, -7.62 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (738.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (738.6 DPS) | yes | Darkmoon Card: Heroism (19287, -6.33 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (738.6 DPS) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Shadowsong's Sorrow (21522, -4.29 DPS, sim-verified) [quest] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 837.4 attack_power points (146.68 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Greenhammer (279261, -14.73 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -62.23 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (738.6 DPS) | yes | The Purifier (22656, -0.23 DPS) [quest]; Bloodseeker (19107, -1.16 DPS) [quest]; Dark Iron Rifle (16004, -6.67 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1171, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 32501000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 40.5. Weights run: 1.6s. Verify run: 3.1s. 191 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.974 ± 0.014, crit=0.569 ± 0.013 per rating point (14 rating = 1%, 7.966 per %), hit=0.686 ± 0.032 per rating point (10 rating = 1%, 6.856 per %), melee_haste=4.645 ± 0.418

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.8 attack_power points (1.02 DPS) | yes | Defender's Leather Hood (252447, -0.47 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.8 attack_power points (0.76 DPS) | yes | Erudite's Amulet (277204, -0.54 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.9 attack_power points (0.64 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.8 attack_power points (0.76 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.98 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.9 attack_power points (0.64 DPS) | yes | Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Ratchet Wristwraps (274742, -0.25 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.8 attack_power points (1.02 DPS) | yes | Brawler's Leather Gloves (252494, -0.25 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Bristlebark Gloves (14572, -0.28 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Murloc Scale Belt (5780, -0.65 DPS) [crafted]; Deviate Scale Belt (6468, -3.07 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.19 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.10 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.9 attack_power points (0.77 DPS) | yes | Bounty Hunter's Ring (5351, -0.39 DPS) [quest]; Demon Band (12054, -0.51 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.6 attack_power points (0.69 DPS) | yes | Bounty Hunter's Ring (5351, -0.30 DPS) [quest]; Demon Band (12054, -0.43 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Diamond Hammer (2194) | World drop [world_drop] | sim-verified (40.5 DPS) | yes | Wingblade (6504, -0.03 DPS) [quest]; Blackfang (2236, -0.08 DPS) [world_drop]; Shadowfang (1482, -0.57 DPS, sim-verified) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.9 attack_power points (0.51 DPS) | yes | Fine Longbow (11304, -0.24 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Diamond Hammer; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 254779 A'sharahm, the Roiling Tempest; 263006 Scout Ranger's Tunic

### Band 30 (troll, 32502110520000000-00000000000000000-0000000000000000000)

Set DPS (verified): 62.8. Weights run: 1.6s. Verify run: 3.3s. 354 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.714 ± 0.013, crit=0.667 ± 0.015 per rating point (14 rating = 1%, 9.342 per %), hit=0.897 ± 0.040 per rating point (10 rating = 1%, 8.973 per %), melee_haste=7.864 ± 0.505

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.1 attack_power points (1.14 DPS) | yes | Brawler's Leather Hood (252504, -0.23 DPS) [crafted]; Defender's Leather Helm (252455, -0.34 DPS) [crafted]; Humbert's Helm (4724, -0.34 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.7 attack_power points (0.91 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.11 DPS) [world]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; River Pride Choker (13087, -0.65 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.68 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Fenrus' Hide (6340, -0.31 DPS) [dungeon]; Swiftrunner Cape (6745, -0.31 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.0 attack_power points (1.59 DPS) | yes | Brawler's Leather Tunic (252508, -0.30 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.44 DPS) [quest]; Brawler's Leather Armor (252490, -0.46 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.3 attack_power points (0.95 DPS) | yes | Barbaric Bracers (18948, -0.23 DPS) [crafted]; Insignia Bracers (6410, -0.27 DPS) [world_drop]; Cultist's Armguards (270032, -0.28 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (62.8 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.97 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (62.8 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -3.37 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (62.8 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.11 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (62.8 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Feet of the Lynx (1121, -3.32 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.3 attack_power points (1.08 DPS) | yes | Band of the Fist (17694, -0.36 DPS) [quest]; Pyrewood Signet Ring (277210, -0.39 DPS) [quest]; Ring of Precision (1491, -0.40 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.87 DPS) | yes | Band of the Fist (17694, -0.15 DPS) [quest]; Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Ring of Precision (1491, -0.19 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (21.39 DPS) | yes | Zealot Blade (13033, -0.80 DPS) [world_drop]; Electrocutioner Leg (9446, -0.87 DPS) [dungeon]; Scout's Blade (19545, -0.90 DPS) [pvp] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.3 attack_power points (21.15 DPS) | yes | Zealot Blade (13033, -2.22 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.3 attack_power points (0.68 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.09 DPS) [vendor]; Double-barreled Shotgun (2098, -0.14 DPS) [world_drop]; Gravespike Repeater (273033, -0.15 DPS) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 354, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 32502110551401001-00000000000000000-0000000000000000000)

Set DPS (verified): 107.8. Weights run: 1.6s. Verify run: 1.7s. 474 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.551 ± 0.015, crit=0.716 ± 0.021 per rating point (14 rating = 1%, 10.030 per %), hit=1.873 ± 0.077 per rating point (10 rating = 1%, 18.732 per %), melee_haste=30.832 ± 0.898

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.1 attack_power points (0.97 DPS) | yes | Hawkeye's Helm (14591, -0.24 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.27 DPS) [world_drop]; Nightscape Headband (8176, -0.33 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 17.1 attack_power points (0.59 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Ethereal Talisman (4430, -0.20 DPS) [quest]; Kaleidoscope Chain (13084, -0.24 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.1 attack_power points (1.00 DPS) | yes | Forest Tracker Epaulets (2278, -0.24 DPS) [world_drop]; Flintrock Shoulders (7755, -0.30 DPS) [dungeon]; Nightscape Shoulders (8192, -0.41 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.9 attack_power points (0.48 DPS) | yes | Imperial Cloak (6432, -0.05 DPS) [dungeon]; Parachute Cloak (10518, -0.05 DPS) [crafted]; Tigerstrike Mantle (13108, -0.05 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 26.4 attack_power points (0.91 DPS) | yes | Nightscape Tunic (8175, -0.11 DPS) [crafted]; Dusky Leather Armor (7374, -0.16 DPS) [crafted]; Barbaric Harness (5739, -0.24 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.23 DPS) [world_drop]; Dusky Bracers (7378, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.0 attack_power points (1.04 DPS) | yes | Skulker's Leather Gloves (252525, -0.19 DPS) [crafted]; Dog Training Gloves (7756, -0.21 DPS) [dungeon]; Prowler's Leather Gloves (252524, -0.21 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.04 DPS) | yes | Defiler's Chain Girdle (20152, -0.21 DPS) [rep]; Ogron's Sash (13117, -0.24 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.35 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 32.6 attack_power points (1.12 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.23 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.37 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.1 attack_power points (0.83 DPS) | yes | Imperial Leather Boots (6431, -0.07 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.11 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.4 attack_power points (0.70 DPS) | yes | Assault Band (13095, -0.01 DPS) [world_drop]; Ring of the Underwood (2951, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (0.69 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.09 DPS) [world_drop]; Falcon's Hook (7552, -0.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Ardent Custodian (868, -0.50 DPS) [world_drop]; Vanquisher's Sword (10823, -1.12 DPS) [quest] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | sim-verified (107.8 DPS) | yes | Ardent Custodian (868, -61.52 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.11 DPS) [world_drop]; Monolithic Bow (9426, -0.12 DPS) [dungeon]; Bow of Searing Arrows (2825, -0.52 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 474, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32502110551501001-30230100000000000-0000000000000000000)

Set DPS (verified): 172.6. Weights run: 2.4s. Verify run: 2.2s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.579 ± 0.012, crit=0.963 ± 0.018 per rating point (14 rating = 1%, 13.485 per %), hit=2.203 ± 0.082 per rating point (10 rating = 1%, 22.026 per %), melee_haste=5.914 ± 1.264

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.4 attack_power points (2.45 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.13 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (2.21 DPS) | yes | Woven Ivy Necklace (19159, -0.46 DPS) [quest]; Scout's Medallion (19535, -0.57 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.4 attack_power points (2.54 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.26 DPS) [crafted]; Failed Flying Experiment (9647, -0.31 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.7 attack_power points (2.05 DPS) | yes | Blackflame Cape (13109, -0.41 DPS) [world_drop]; Serpentskin Cloak (8259, -0.55 DPS) [world_drop]; Pridelord Cape (14673, -0.56 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Warbear Harness (15064, -1.06 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.23 DPS) [quest]; Charred Leather Tunic (19127, -2.17 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.7 attack_power points (2.39 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.56 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 44.7 attack_power points (3.87 DPS) | yes | Darkmantle Grips (226828, -0.09 DPS) [vendor]; Gloves of Holy Might (867, -0.97 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.18 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 34.1 attack_power points (2.95 DPS) | yes | Defiler's Leather Girdle (20193, -0.05 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 53.5 attack_power points (4.62 DPS) | yes | Serpentskin Leggings (8262, -1.15 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -1.76 DPS) [world_drop]; Triprunner Dungarees (9624, -1.91 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.6 attack_power points (2.73 DPS) | yes | Skulker's Leather Boots (252469, -0.00 DPS) [crafted]; Prowler's Leather Boots (252468, -0.11 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.62 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.2 attack_power points (2.09 DPS) | yes | Masons Fraternity Ring (9533, -0.18 DPS) [quest]; Mark of Kern (2262, -0.36 DPS) [dungeon]; Blackstone Ring (17713, -0.36 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.07 DPS) | yes | Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Mark of Kern (2262, -0.35 DPS) [dungeon]; Blackstone Ring (17713, -0.35 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (172.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (172.6 DPS) | yes | Molten Heart of the Mountain (249470, -0.73 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (172.6 DPS) | yes | Julie's Dagger (6660, -3.61 DPS) [world_drop]; Hanzo Sword (8190, -3.86 DPS) [world_drop]; Scout's Blade (19543, -3.97 DPS) [pvp] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (45.50 DPS) | yes | Julie's Dagger (6660, -3.56 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -3.83 DPS) [dungeon]; White Bone Shredder (11863, -6.60 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (172.6 DPS) | yes | Skull Splitting Crossbow (13039, -0.44 DPS) [world_drop]; Guttbuster (13139, -0.56 DPS) [world_drop]; Dark Iron Rifle (16004, -2.16 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32502110551501001-30230300000000000-5120000000000000000)

Set DPS (verified): 290.4. Weights run: 2.3s. Verify run: 6.3s. 1165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.815 ± 0.019, crit=1.646 ± 0.030 per rating point (14 rating = 1%, 23.048 per %), hit=3.401 ± 0.116 per rating point (10 rating = 1%, 34.008 per %), melee_haste=9.014 ± 1.758

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.7 attack_power points (6.08 DPS) | yes | Darkmantle Cap (226829, +0.00 DPS) [quest]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 58.0 attack_power points (4.85 DPS) | yes | Medallion of the Dawn (22659, -0.92 DPS) [quest]; Amulet of the Darkmoon (19491, -1.13 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.07 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (290.4 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -15.15 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 62.0 attack_power points (5.18 DPS) | yes | Cape of the Black Baron (13340, -1.24 DPS) [dungeon]; Deathguard's Cloak (20068, -1.58 DPS) [rep]; Stalwart Cloak (272415, -2.34 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.2 attack_power points (10.39 DPS) | yes | Warlord's Leather Breastplate (231549, -1.67 DPS) [pvp]; Darkmantle Tunic (226825, -1.83 DPS, sim-verified) [quest]; Legionnaire's Leather Chestpiece (227060, -2.77 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 45.5 attack_power points (3.80 DPS) | yes | General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Bracers of Subterfuge (22668, -0.69 DPS) [quest]; Deepfury Bracers (13120, -1.19 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (290.4 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -0.43 DPS) [quest]; Raider Gloves (272099, -19.00 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 106.1 attack_power points (8.87 DPS) | yes | Belt of Preserved Heads (20216, -1.25 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -4.10 DPS) [rep]; Ferocity of the Timbermaw (227805, -4.12 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (290.4 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.15 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 67.4 attack_power points (5.63 DPS) | yes | Fine Dawn Treaders (227815, -1.18 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.59 DPS) [dungeon]; Drudge Boots (21532, -1.97 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 77.6 attack_power points (6.49 DPS) | yes | Cutthroat's Signet (272408, -1.52 DPS) [vendor]; Channeler's Ring (272406, -3.64 DPS) [vendor]; Blessed Band of Light (272407, -3.64 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 71.1 attack_power points (5.94 DPS) | yes | Cutthroat's Signet (272408, -1.06 DPS, sim-verified) [vendor]; Channeler's Ring (272406, -3.10 DPS) [vendor]; Blessed Band of Light (272407, -3.10 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (290.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (290.4 DPS) | yes | Frozen Heart of the Mountain (249469, -2.94 DPS) [crafted]; Molten Heart of the Mountain (249470, -4.74 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (290.4 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Teebu's Blazing Longsword (1728, -4.57 DPS, sim-verified) [world_drop] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 832.7 attack_power points (69.61 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -4.77 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -29.30 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (290.4 DPS) | yes | The Purifier (22656, -0.20 DPS) [quest]; Bloodseeker (19107, -0.39 DPS) [quest]; Dark Iron Rifle (16004, -2.56 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 01532310421501000-31530300001500000-0020000000000000000)

Set DPS (verified): 737.5. Weights run: 2.4s. Verify run: 6.6s. 1165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.200 ± 0.021, crit=2.108 ± 0.035 per rating point (14 rating = 1%, 29.512 per %), hit=3.859 ± 0.162 per rating point (10 rating = 1%, 38.590 per %), melee_haste=16.559 ± 2.071

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 90.5 attack_power points (15.85 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -1.04 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 62.6 attack_power points (10.96 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -1.71 DPS) [quest]; Zandalarian Shadow Talisman (19616, -4.03 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (737.5 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -45.44 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.6 attack_power points (11.66 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -3.78 DPS) [rep]; Stalwart Cloak (272415, -4.90 DPS) [vendor] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 149.5 attack_power points (26.19 DPS) | yes | Darkmantle Tunic (226825, -4.00 DPS, sim-verified) [quest]; Warlord's Leather Breastplate (231549, -4.24 DPS) [pvp]; Timbermaw Tunic (252484, -7.72 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 53.9 attack_power points (9.44 DPS) | yes | General's Leather Armsplints (16559, -0.19 DPS) [pvp]; Bracers of Subterfuge (22668, -1.73 DPS) [quest]; Deepfury Bracers (13120, -2.89 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (737.5 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.14 DPS) [quest]; Raider Gloves (272099, -54.08 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 123.4 attack_power points (21.61 DPS) | yes | Belt of Preserved Heads (20216, -6.37 DPS) [quest]; Ferocity of the Timbermaw (227805, -10.05 DPS) [vendor]; Defiler's Leather Girdle (20190, -10.49 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (737.5 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -45.44 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 79.8 attack_power points (13.98 DPS) | yes | Fine Dawn Treaders (227815, -3.44 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -3.86 DPS) [dungeon]; Darkmantle Boots (22003, -4.73 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 91.4 attack_power points (16.01 DPS) | yes | Cutthroat's Signet (272408, -3.85 DPS) [vendor]; Channeler's Ring (272406, -9.25 DPS) [vendor]; Blessed Band of Light (272407, -9.25 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 82.1 attack_power points (14.38 DPS) | yes | Cutthroat's Signet (272408, -2.23 DPS) [vendor]; Channeler's Ring (272406, -7.62 DPS) [vendor]; Blessed Band of Light (272407, -7.62 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (737.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (737.5 DPS) | yes | Frozen Heart of the Mountain (249469, -6.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -8.88 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (737.5 DPS) | yes | Teebu's Blazing Longsword (1728, +0.00 DPS) [world_drop]; High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 837.4 attack_power points (146.68 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Greenhammer (279261, -5.14 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -62.23 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (737.5 DPS) | yes | The Purifier (22656, -0.23 DPS) [quest]; Bloodseeker (19107, -1.16 DPS) [quest]; Dark Iron Rifle (16004, -6.96 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy

No-known-source sample (15 of 1165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

