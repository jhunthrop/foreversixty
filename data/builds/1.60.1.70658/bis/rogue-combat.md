# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.8. Weights run: 1.7s. Verify run: 3.5s. 198 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.938 ± 0.015, crit=0.594 ± 0.014 per rating point (14 rating = 1%, 8.317 per %), hit=0.619 ± 0.030 per rating point (10 rating = 1%, 6.194 per %), melee_haste=4.054 ± 0.376

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.5 attack_power points (0.96 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.6 attack_power points (0.72 DPS) | yes | Erudite's Amulet (277204, -0.49 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.7 attack_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.62 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.6 attack_power points (0.72 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 21.3 attack_power points (1.31 DPS) | yes | Brawler's Leather Armor (252490, -0.17 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.48 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.8 attack_power points (0.60 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.8 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.65 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Dusty Belt (279897, -0.51 DPS) [quest]; Deviate Scale Belt (6468, -3.66 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.8 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.91 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.8 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.83 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.8 attack_power points (0.72 DPS) | yes | Demon Band (12054, -0.48 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon]; The 1 Ring (8350, -0.54 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.2 attack_power points (0.63 DPS) | yes | Demon Band (12054, -0.23 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon]; The 1 Ring (8350, -0.45 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.10 DPS) | yes | Evocator's Blade (2567, -0.88 DPS) [dungeon]; Buzzer Blade (2169, -2.01 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.59 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 228.2 attack_power points (14.07 DPS) | yes | Evocator's Blade (2567, -4.31 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.8 attack_power points (0.48 DPS) | yes | Fine Longbow (11304, -0.23 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.24 DPS) [crafted]; Light Bow (4576, -0.24 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 70.3. Weights run: 1.6s. Verify run: 3.7s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.543 ± 0.013, crit=0.567 ± 0.014 per rating point (14 rating = 1%, 7.934 per %), hit=0.908 ± 0.036 per rating point (10 rating = 1%, 9.077 per %), melee_haste=7.207 ± 0.435

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.4 attack_power points (0.96 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.21 DPS) [crafted]; Humbert's Helm (4724, -0.29 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.77 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop]; River Pride Choker (13087, -0.52 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.0 attack_power points (1.37 DPS) | yes | Mantle of Thieves (2264, -0.49 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS) [crafted]; Bristlebark Amice (14573, -0.61 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.8 attack_power points (0.86 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.6 attack_power points (1.35 DPS) | yes | Tunic of Westfall (2041, -0.29 DPS) [quest]; Brawler's Leather Tunic (252508, -0.30 DPS, sim-verified) [crafted]; Raptorbane Armor (3566, -0.35 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.3 attack_power points (0.83 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.20 DPS) [quest]; Insignia Bracers (6410, -0.25 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.3 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.09 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (70.3 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.47 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.3 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.21 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (70.3 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.44 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.3 attack_power points (0.95 DPS) | yes | Pyrewood Signet Ring (277210, -0.34 DPS) [quest]; Ring of Precision (1491, -0.37 DPS) [dungeon]; Insurgent's Band (272067, -0.39 DPS) [vendor] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.6 attack_power points (0.79 DPS) | yes | Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Ring of Precision (1491, -0.21 DPS) [dungeon]; Insurgent's Band (272067, -0.23 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.12 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.15 DPS, sim-verified) [vendor]; Thornspike (6681, -2.11 DPS) [dungeon]; Toxic Revenger (9453, -2.16 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (70.3 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -7.19 DPS, sim-verified) [vendor] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.3 attack_power points (0.58 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.02 DPS) [vendor]; Gravespike Repeater (273033, -0.08 DPS) [world]; Double-barreled Shotgun (2098, -0.10 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Golemsight Long Gun

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 116.0. Weights run: 1.8s. Verify run: 2.0s. 493 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.388 ± 0.014, crit=0.503 ± 0.018 per rating point (14 rating = 1%, 7.037 per %), hit=1.098 ± 0.039 per rating point (10 rating = 1%, 10.975 per %), melee_haste=8.196 ± 0.525

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.3 attack_power points (1.78 DPS) | yes | Warden's Wizard Hat (14604, -0.56 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.65 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 15.3 attack_power points (1.04 DPS) | yes | Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.56 DPS, sim-verified) [world]; River Pride Choker (13087, -0.76 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.3 attack_power points (1.85 DPS) | yes | Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.3 attack_power points (0.97 DPS) | yes | Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.20 DPS) [quest]; Parachute Cloak (10518, -0.22 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 23.6 attack_power points (1.60 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.37 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Hawkeye's Bracers (14590, -0.49 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.60 DPS) [dungeon]; Dusky Bracers (7378, -0.60 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 27.0 attack_power points (1.83 DPS) | yes | Skulker's Leather Gloves (252525, -0.28 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.31 DPS) [crafted]; Dog Training Gloves (7756, -0.76 DPS, sim-verified) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.04 DPS) | yes | Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.58 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.78 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 29.2 attack_power points (1.98 DPS) | yes | Triprunner Dungarees (9624, -0.08 DPS) [quest]; Ferine Leggings (6690, -0.21 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.66 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.3 attack_power points (1.51 DPS) | yes | Prowler's Leather Shoes (252465, -0.11 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Protector's Band (19515, -0.06 DPS) [rep]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Falcon's Hook (7552, -0.31 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.36 DPS) | yes | Protector's Band (19515, -0.06 DPS) [rep]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Falcon's Hook (7552, -0.31 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (116.0 DPS) | yes | Black Menace (6831, -4.01 DPS) [quest]; Coldrage Dagger (10761, -4.68 DPS, sim-verified) [dungeon]; Darkspear Insurgent's Spellblade (272085, -4.71 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 442.1 attack_power points (29.99 DPS) | yes | Black Menace (6831, -0.88 DPS, sim-verified) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (116.0 DPS) | yes | Monolithic Bow (9426, -0.26 DPS) [dungeon]; Swiftwind (13038, -0.29 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.15 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 493, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 160.9. Weights run: 2.3s. Verify run: 2.5s. 616 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.554 ± 0.014, crit=0.927 ± 0.021 per rating point (14 rating = 1%, 12.978 per %), hit=1.364 ± 0.062 per rating point (10 rating = 1%, 13.641 per %), melee_haste=5.909 ± 0.880

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 42.6 attack_power points (3.62 DPS) | yes | White Bandit Mask (10008, -1.36 DPS, sim-verified) [crafted]; Helm of Fire (8348, -1.38 DPS) [crafted]; Lordrec Helmet (10741, -1.51 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.2 attack_power points (2.14 DPS) | yes | Sentinel's Medallion (19539, -0.63 DPS, sim-verified) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 34.6 attack_power points (2.94 DPS) | yes | Sunburn Spaulders (274751, -0.47 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.72 DPS) [crafted]; Failed Flying Experiment (9647, -0.77 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.3 attack_power points (1.98 DPS) | yes | Blackflame Cape (13109, -0.40 DPS) [world_drop]; Serpentskin Cloak (8259, -0.53 DPS) [world_drop]; Pridelord Cape (14673, -0.54 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 44.6 attack_power points (3.79 DPS) | yes | Warbear Harness (15064, -0.48 DPS) [crafted]; Blazewind Breastplate (11193, -0.50 DPS) [quest]; Charred Leather Tunic (19127, -1.41 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.3 attack_power points (2.32 DPS) | yes | Pridelord Bands (14672, -0.58 DPS) [world_drop]; Skulker's Leather Bracers (252540, -0.60 DPS, sim-verified) [crafted]; Branded Leather Bracers (19508, -0.62 DPS) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 44.2 attack_power points (3.76 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.95 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.14 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 33.8 attack_power points (2.87 DPS) | yes | Highlander's Leather Girdle (20115, -0.07 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 44.6 attack_power points (3.79 DPS) | yes | Gryphon Rider's Leggings (9652, -0.57 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -0.57 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.02 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 31.2 attack_power points (2.65 DPS) | yes | Albino Crocscale Boots (17728, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.60 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.0 attack_power points (2.04 DPS) | yes | Mark of Kern (2262, -0.34 DPS) [dungeon]; Assault Band (13095, -0.34 DPS) [world_drop]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.8 attack_power points (1.85 DPS) | yes | Mark of Kern (2262, -0.15 DPS) [dungeon]; Assault Band (13095, -0.15 DPS) [world_drop]; Blackstone Ring (17713, -0.15 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (160.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (160.9 DPS) | yes | Mark of the Chosen (17774, -1.05 DPS, sim-verified) [quest] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (160.9 DPS) | yes | Charstone Dirk (17710, -0.71 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, -1.83 DPS) [vendor]; Gut Ripper (2164, -3.15 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.77 DPS) | yes | Charstone Dirk (17710, -29.63 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (160.9 DPS) | yes | Skull Splitting Crossbow (13039, -0.40 DPS) [world_drop]; Guttbuster (13139, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.10 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Skulker's Leather Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 616, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 258.7. Weights run: 2.5s. Verify run: 6.6s. 1171 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.775 ± 0.020, crit=1.572 ± 0.035 per rating point (14 rating = 1%, 22.002 per %), hit=2.339 ± 0.093 per rating point (10 rating = 1%, 23.387 per %), melee_haste=15.872 ± 1.350

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 70.3 attack_power points (5.65 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.31 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 47.4 attack_power points (3.81 DPS) | yes | Amulet of the Darkmoon (19491, -0.29 DPS) [quest]; Zandalarian Shadow Talisman (19616, -1.19 DPS) [quest]; Medallion of the Dawn (22659, -1.30 DPS, sim-verified) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (258.7 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -12.57 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 51.4 attack_power points (4.13 DPS) | yes | Cloak of the Honor Guard (20073, -0.68 DPS) [rep]; Cape of the Black Baron (13340, -1.44 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -1.71 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 111.1 attack_power points (8.93 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -1.57 DPS) [pvp]; Timbermaw Tunic (252484, -1.94 DPS) [crafted]; Darkmantle Tunic (226825, -2.67 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 44.7 attack_power points (3.60 DPS) | yes | Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Bracers of Subterfuge (22668, -0.65 DPS) [quest]; Deepfury Bracers (13120, -1.13 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (258.7 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, -14.37 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 84.1 attack_power points (6.76 DPS) | yes | Ferocity of the Timbermaw (227805, -2.24 DPS) [vendor]; Highlander's Leather Girdle (20045, -2.26 DPS) [rep]; Belt of Preserved Heads (20216, -2.78 DPS, sim-verified) [quest] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (258.7 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -12.57 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 59.0 attack_power points (4.74 DPS) | yes | Drudge Boots (21532, -1.27 DPS) [quest]; Fine Dawn Treaders (227815, -1.29 DPS) [vendor]; Shadowcraft Boots (16711, -1.61 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 66.0 attack_power points (5.31 DPS) | yes | Cutthroat's Signet (272408, -1.43 DPS) [vendor]; Protector's Band (19514, -2.77 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 59.4 attack_power points (4.78 DPS) | yes | Cutthroat's Signet (272408, -1.04 DPS, sim-verified) [vendor]; Protector's Band (19514, -2.24 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (258.7 DPS) | yes | Frozen Heart of the Mountain (249469, -8.17 DPS, sim-verified) [crafted] |
| trinket2 | Darkmoon Card: Heroism (19287) | Darkmoon Warlords Deck [quest] | sim-verified (258.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (258.7 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -17.32 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (53.16 DPS) | yes | Electrified Dagger (19100, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (258.7 DPS) | yes | The Purifier (22656, -0.23 DPS) [quest]; Bloodseeker (19107, -0.36 DPS) [quest]; Dark Iron Rifle (16004, -2.39 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Darkmoon Card: Heroism; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1171, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 704.0. Weights run: 2.6s. Verify run: 7.0s. 1171 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.076 ± 0.024, crit=1.883 ± 0.040 per rating point (14 rating = 1%, 26.367 per %), hit=3.106 ± 0.136 per rating point (10 rating = 1%, 31.059 per %), melee_haste=22.436 ± 1.746

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 83.1 attack_power points (14.14 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.83 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 55.1 attack_power points (9.37 DPS) | yes | Amulet of the Darkmoon (19491, -0.78 DPS) [quest]; Medallion of the Dawn (22659, -0.80 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.95 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (704.0 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -37.17 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.1 attack_power points (10.05 DPS) | yes | Cape of the Black Baron (13340, -1.35 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.50 DPS) [rep]; Shifting Cloak (18511, -4.04 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.2 attack_power points (22.84 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -3.88 DPS) [pvp]; Darkmantle Tunic (226825, -5.38 DPS, sim-verified) [quest]; Timbermaw Tunic (252484, -5.82 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 51.5 attack_power points (8.77 DPS) | yes | Marshal's Leather Armsplints (16460, -0.19 DPS) [pvp]; Bracers of Subterfuge (22668, -1.60 DPS) [quest]; Deepfury Bracers (13120, -2.72 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (704.0 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, -42.52 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 105.7 attack_power points (17.98 DPS) | yes | Belt of Preserved Heads (20216, -4.29 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -7.11 DPS) [vendor]; Highlander's Leather Girdle (20045, -7.71 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (704.0 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -37.17 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 71.6 attack_power points (12.17 DPS) | yes | Shadowcraft Boots (16711, -3.17 DPS) [dungeon]; Fine Dawn Treaders (227815, -3.47 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -3.70 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 80.9 attack_power points (13.76 DPS) | yes | Cutthroat's Signet (272408, -3.53 DPS) [vendor]; Protector's Band (19514, -7.63 DPS) [rep]; Channeler's Ring (272406, -8.48 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 71.4 attack_power points (12.15 DPS) | yes | Cutthroat's Signet (272408, -1.92 DPS) [vendor]; Protector's Band (19514, -6.02 DPS) [rep]; Channeler's Ring (272406, -6.87 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (704.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (704.0 DPS) | yes | Molten Heart of the Mountain (249470, -3.75 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (704.0 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -37.64 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (112.46 DPS) | yes | Electrified Dagger (19100, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (704.0 DPS) | yes | The Purifier (22656, -0.46 DPS) [quest]; Bloodseeker (19107, -0.98 DPS) [quest]; Dark Iron Rifle (16004, -6.24 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1171, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.0. Weights run: 1.7s. Verify run: 3.5s. 191 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.938 ± 0.015, crit=0.594 ± 0.014 per rating point (14 rating = 1%, 8.317 per %), hit=0.619 ± 0.030 per rating point (10 rating = 1%, 6.194 per %), melee_haste=4.054 ± 0.376

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.5 attack_power points (0.96 DPS) | yes | Defender's Leather Hood (252447, -0.45 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.6 attack_power points (0.72 DPS) | yes | Erudite's Amulet (277204, -0.48 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.7 attack_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.61 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.6 attack_power points (0.72 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (41.0 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.63 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.7 attack_power points (0.60 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.6 attack_power points (0.96 DPS) | yes | Bristlebark Gloves (14572, -0.24 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.25 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Belt (252428, -0.39 DPS) [crafted]; Murloc Scale Belt (5780, -0.63 DPS) [crafted]; Deviate Scale Belt (6468, -3.57 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.82 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.0 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.74 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.8 attack_power points (0.72 DPS) | yes | Bounty Hunter's Ring (5351, -0.37 DPS) [quest]; Demon Band (12054, -0.48 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 10.2 attack_power points (0.63 DPS) | yes | Bounty Hunter's Ring (5351, -0.27 DPS) [quest]; Demon Band (12054, -0.38 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.10 DPS) | yes | Scout's Blade (20441, -0.42 DPS) [pvp]; Evocator's Blade (2567, -0.88 DPS) [dungeon]; Buzzer Blade (2169, -2.01 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 228.2 attack_power points (14.07 DPS) | yes | Evocator's Blade (2567, -3.88 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.8 attack_power points (0.48 DPS) | yes | Fine Longbow (11304, -0.21 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.24 DPS) [crafted]; Light Bow (4576, -0.24 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 254779 A'sharahm, the Roiling Tempest; 263006 Scout Ranger's Tunic

### Band 30 (troll, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 69.4. Weights run: 1.6s. Verify run: 3.7s. 354 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.543 ± 0.013, crit=0.567 ± 0.014 per rating point (14 rating = 1%, 7.934 per %), hit=0.908 ± 0.036 per rating point (10 rating = 1%, 9.077 per %), melee_haste=7.207 ± 0.435

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.4 attack_power points (0.96 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.21 DPS) [crafted]; Humbert's Helm (4724, -0.29 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 12.3 attack_power points (0.77 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop]; River Pride Choker (13087, -0.52 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.0 attack_power points (1.37 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS) [crafted]; Bristlebark Amice (14573, -0.61 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.8 attack_power points (0.86 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon]; Wildhunter Cloak (16658, -0.24 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.6 attack_power points (1.35 DPS) | yes | Brawler's Leather Tunic (252508, -0.33 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.36 DPS) [quest]; Brawler's Leather Armor (252490, -0.36 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.3 attack_power points (0.83 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.20 DPS) [quest]; Insignia Bracers (6410, -0.25 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.4 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.61 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (69.4 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -2.97 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (69.4 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.71 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (69.4 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.93 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.3 attack_power points (0.95 DPS) | yes | Band of the Fist (17694, -0.32 DPS) [quest]; Pyrewood Signet Ring (277210, -0.34 DPS) [quest]; Ring of Precision (1491, -0.37 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.6 attack_power points (0.79 DPS) | yes | Band of the Fist (17694, -0.15 DPS) [quest]; Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Ring of Precision (1491, -0.21 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.12 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -0.44 DPS, sim-verified) [vendor]; Scout's Blade (19545, -0.92 DPS) [pvp]; Thornspike (6681, -2.11 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (69.4 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -6.62 DPS, sim-verified) [vendor] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.3 attack_power points (0.58 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.02 DPS) [vendor]; Gravespike Repeater (273033, -0.08 DPS) [world]; Double-barreled Shotgun (2098, -0.10 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Golemsight Long Gun

No-known-source sample (15 of 354, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 115.1. Weights run: 1.8s. Verify run: 2.1s. 474 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.388 ± 0.014, crit=0.503 ± 0.018 per rating point (14 rating = 1%, 7.037 per %), hit=1.098 ± 0.039 per rating point (10 rating = 1%, 10.975 per %), melee_haste=8.196 ± 0.525

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.3 attack_power points (1.78 DPS) | yes | Warden's Wizard Hat (14604, -0.56 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.65 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 15.3 attack_power points (1.04 DPS) | yes | Ethereal Talisman (4430, -0.32 DPS) [quest]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.65 DPS, sim-verified) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.3 attack_power points (1.85 DPS) | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.7 attack_power points (0.86 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Parachute Cloak (10518, -0.11 DPS) [crafted]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 23.6 attack_power points (1.60 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.37 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.52 DPS) [world_drop]; Dusky Bracers (7378, -0.60 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 27.0 attack_power points (1.83 DPS) | yes | Skulker's Leather Gloves (252525, -0.28 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.31 DPS) [crafted]; Dog Training Gloves (7756, -0.84 DPS, sim-verified) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.04 DPS) | yes | Defiler's Chain Girdle (20152, -0.49 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.58 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.78 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 29.2 attack_power points (1.98 DPS) | yes | Triprunner Dungarees (9624, -0.08 DPS) [quest]; Ferine Leggings (6690, -0.21 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.66 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.3 attack_power points (1.51 DPS) | yes | Prowler's Leather Shoes (252465, -0.11 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Legionnaire's Band (19512, -0.06 DPS) [rep]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Falcon's Hook (7552, -0.31 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.36 DPS) | yes | Legionnaire's Band (19512, -0.06 DPS) [rep]; Ring of the Underwood (2951, -0.28 DPS) [world_drop]; Falcon's Hook (7552, -0.31 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (115.1 DPS) | yes | Scout's Blade (19544, -3.78 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -4.71 DPS) [vendor]; Coldrage Dagger (10761, -5.02 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 442.1 attack_power points (29.99 DPS) | yes | Coldrage Dagger (10761, -0.55 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (115.1 DPS) | yes | Monolithic Bow (9426, -0.26 DPS) [dungeon]; Swiftwind (13038, -0.29 DPS) [world_drop]; Bow of Searing Arrows (2825, -1.14 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 474, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 162.3. Weights run: 2.3s. Verify run: 2.5s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.554 ± 0.014, crit=0.927 ± 0.021 per rating point (14 rating = 1%, 12.978 per %), hit=1.364 ± 0.062 per rating point (10 rating = 1%, 13.641 per %), melee_haste=5.909 ± 0.880

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.1 attack_power points (2.39 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.08 DPS) [quest]; Helm of Fire (8348, -0.14 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.2 attack_power points (2.14 DPS) | yes | Woven Ivy Necklace (19159, -0.44 DPS) [quest]; Scout's Medallion (19535, -0.56 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.1 attack_power points (2.47 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.30 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.3 attack_power points (1.98 DPS) | yes | Blackflame Cape (13109, -0.40 DPS) [world_drop]; Serpentskin Cloak (8259, -0.53 DPS) [world_drop]; Pridelord Cape (14673, -0.54 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 44.6 attack_power points (3.79 DPS) | yes | Blazewind Breastplate (11193, -0.50 DPS) [quest]; Warbear Harness (15064, -1.02 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -1.41 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.3 attack_power points (2.32 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.64 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 44.2 attack_power points (3.76 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.95 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.14 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 33.8 attack_power points (2.87 DPS) | yes | Defiler's Leather Girdle (20193, -0.07 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 44.6 attack_power points (3.79 DPS) | yes | Basilisk Hide Pants (1718, -1.02 DPS) [world_drop]; Serpentskin Leggings (8262, -1.11 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -1.16 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 31.2 attack_power points (2.65 DPS) | yes | Albino Crocscale Boots (17728, -0.01 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.60 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.04 DPS) | yes | Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Mark of Kern (2262, -0.34 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.0 attack_power points (2.04 DPS) | yes | Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Mark of Kern (2262, -0.34 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (162.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (162.3 DPS) | yes | Molten Heart of the Mountain (249470, -0.90 DPS, sim-verified) [crafted] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (162.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Scout's Blade (19543, -0.38 DPS) [pvp]; Charstone Dirk (17710, -0.71 DPS) [dungeon] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.77 DPS) | yes | Charstone Dirk (17710, -29.05 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (162.3 DPS) | yes | Skull Splitting Crossbow (13039, -0.40 DPS) [world_drop]; Guttbuster (13139, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.08 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 265.0. Weights run: 2.5s. Verify run: 6.7s. 1165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.775 ± 0.020, crit=1.572 ± 0.035 per rating point (14 rating = 1%, 22.002 per %), hit=2.339 ± 0.093 per rating point (10 rating = 1%, 23.387 per %), melee_haste=15.872 ± 1.350

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 70.3 attack_power points (5.65 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.31 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 47.4 attack_power points (3.81 DPS) | yes | Medallion of the Dawn (22659, -0.11 DPS) [quest]; Amulet of the Darkmoon (19491, -0.29 DPS) [quest]; Zandalarian Shadow Talisman (19616, -1.19 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -10.76 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 51.4 attack_power points (4.13 DPS) | yes | Cape of the Black Baron (13340, -0.38 DPS) [dungeon]; Deathguard's Cloak (20068, -0.68 DPS) [rep]; Shifting Cloak (18511, -1.71 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 111.1 attack_power points (8.93 DPS) | yes | Warlord's Leather Breastplate (231549, -1.57 DPS) [pvp]; Timbermaw Tunic (252484, -1.94 DPS) [crafted]; Darkmantle Tunic (226825, -2.46 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 44.7 attack_power points (3.60 DPS) | yes | General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Bracers of Subterfuge (22668, -0.65 DPS) [quest]; Deepfury Bracers (13120, -1.13 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -13.01 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 84.1 attack_power points (6.76 DPS) | yes | Belt of Preserved Heads (20216, -1.80 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -2.24 DPS) [vendor]; Defiler's Leather Girdle (20190, -2.26 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -10.76 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 59.0 attack_power points (4.74 DPS) | yes | Drudge Boots (21532, -1.27 DPS) [quest]; Fine Dawn Treaders (227815, -1.29 DPS) [vendor]; Shadowcraft Boots (16711, -1.92 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 66.0 attack_power points (5.31 DPS) | yes | Cutthroat's Signet (272408, -1.43 DPS) [vendor]; Legionnaire's Band (19510, -2.77 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 59.4 attack_power points (4.78 DPS) | yes | Cutthroat's Signet (272408, -0.90 DPS) [vendor]; Legionnaire's Band (19510, -2.24 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+9.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.78 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+15.3 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-verified (265.0 DPS) | yes | The Lobotomizer (19324, -3.14 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Purifier (22656, -0.23 DPS) [quest]; Bloodseeker (19107, -0.36 DPS) [quest]; Dark Iron Rifle (16004, -2.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Glacial Blade

No-known-source sample (15 of 1165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 704.7. Weights run: 2.6s. Verify run: 6.9s. 1165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.076 ± 0.024, crit=1.883 ± 0.040 per rating point (14 rating = 1%, 26.367 per %), hit=3.106 ± 0.136 per rating point (10 rating = 1%, 31.059 per %), melee_haste=22.436 ± 1.746

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 83.1 attack_power points (14.14 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.83 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 55.1 attack_power points (9.37 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -0.80 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.95 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (704.7 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -36.58 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 59.1 attack_power points (10.05 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -2.50 DPS) [rep]; Shifting Cloak (18511, -4.04 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.2 attack_power points (22.84 DPS) | yes | Darkmantle Tunic (226825, -3.19 DPS, sim-verified) [quest]; Warlord's Leather Breastplate (231549, -3.88 DPS) [pvp]; Timbermaw Tunic (252484, -5.82 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 51.5 attack_power points (8.77 DPS) | yes | General's Leather Armsplints (16559, -0.19 DPS) [pvp]; Bracers of Subterfuge (22668, -1.60 DPS) [quest]; Deepfury Bracers (13120, -2.72 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (704.7 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -42.01 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 105.7 attack_power points (17.98 DPS) | yes | Belt of Preserved Heads (20216, -4.78 DPS) [quest]; Ferocity of the Timbermaw (227805, -7.11 DPS) [vendor]; Defiler's Leather Girdle (20190, -7.71 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (704.7 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -36.58 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 71.6 attack_power points (12.17 DPS) | yes | Fine Dawn Treaders (227815, -2.98 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -3.17 DPS) [dungeon]; Darkmantle Boots (22003, -3.70 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 80.9 attack_power points (13.76 DPS) | yes | Cutthroat's Signet (272408, -3.53 DPS) [vendor]; Legionnaire's Band (19510, -7.63 DPS) [rep]; Channeler's Ring (272406, -8.48 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 71.4 attack_power points (12.15 DPS) | yes | Cutthroat's Signet (272408, -1.92 DPS) [vendor]; Legionnaire's Band (19510, -6.02 DPS) [rep]; Channeler's Ring (272406, -6.87 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (704.7 DPS) | yes | Frozen Heart of the Mountain (249469, -6.55 DPS, sim-verified) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (704.7 DPS) | yes | Frozen Heart of the Mountain (249469, -5.77 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (704.7 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -35.26 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (112.46 DPS) | yes | Glacial Blade (19099, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (704.7 DPS) | yes | The Purifier (22656, -0.46 DPS) [quest]; Bloodseeker (19107, -0.98 DPS) [quest]; Dark Iron Rifle (16004, -6.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

