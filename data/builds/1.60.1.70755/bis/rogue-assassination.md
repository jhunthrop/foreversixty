# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32510000000000000-0000000000000000000-000000000000000000)

Set DPS (verified): 39.5. Weights run: 4.7s. Verify run: 9.4s. 234 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.690 ± 0.007, crit=0.418 ± 0.006 per rating point (14 rating = 1%, 5.848 per %), hit=0.633 ± 0.019 per rating point (10 rating = 1%, 6.326 per %), melee_haste=3.510 ± 0.301

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.5 attack_power points (0.99 DPS) | yes | Defender's Leather Hood (252447, -0.47 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.1 attack_power points (0.75 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.97 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.4 attack_power points (0.62 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (39.5 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.23 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 18.6 attack_power points (1.37 DPS) | yes | Brawler's Leather Armor (252490, -0.13 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop]; Defender's Leather Armor (252434, -0.48 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.4 attack_power points (0.62 DPS) | yes | Bravo's Armbands (270015, -0.03 DPS) [quest]; Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.5 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.18 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.32 DPS) | yes | Brawler's Leather Belt (252428, -0.53 DPS) [crafted]; Murloc Scale Belt (5780, -0.80 DPS) [crafted]; Deviate Scale Belt (6468, -4.24 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (39.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.41 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (39.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.32 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.8 attack_power points (0.79 DPS) | yes | Demon Band (12054, -0.50 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon]; The 1 Ring (8350, -0.59 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.3 attack_power points (0.68 DPS) | yes | Demon Band (12054, -0.30 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (16.83 DPS) | yes | Tail Spike (6448, -1.92 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.17 DPS) [crafted]; Goblin Screwdriver (1936, -3.82 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.1 attack_power points (16.71 DPS) | yes | Tail Spike (6448, -0.60 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.8 attack_power points (0.50 DPS) | yes | Fine Longbow (11304, -0.25 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 234, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 58.7. Weights run: 4.9s. Verify run: 10.1s. 411 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.716 ± 0.010, crit=0.701 ± 0.011 per rating point (14 rating = 1%, 9.815 per %), hit=1.052 ± 0.041 per rating point (10 rating = 1%, 10.522 per %), melee_haste=5.968 ± 0.572

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.2 attack_power points (1.19 DPS) | yes | Brawler's Leather Hood (252504, -0.24 DPS) [crafted]; Humbert's Helm (4724, -0.36 DPS) [world]; Defender's Leather Helm (252455, -0.36 DPS) [crafted] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.7 attack_power points (0.95 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; Kaleidoscope Chain (13084, -0.20 DPS) [world_drop]; River Pride Choker (13087, -0.67 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.66 DPS) | yes | Mantle of Thieves (2264, -0.46 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.71 DPS) [crafted]; Bristlebark Amice (14573, -0.73 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (58.7 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.29 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.33 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.0 attack_power points (1.67 DPS) | yes | Brawler's Leather Tunic (252508, -0.28 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.36 DPS) [quest]; Brawler's Leather Armor (252490, -0.49 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.3 attack_power points (0.99 DPS) | yes | Barbaric Bracers (18948, -0.24 DPS) [crafted]; Insignia Bracers (6410, -0.28 DPS) [world_drop]; Cultist's Armguards (270032, -0.30 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.7 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.57 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (58.7 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -2.94 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.7 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.68 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (58.7 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.90 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.3 attack_power points (1.13 DPS) | yes | Pyrewood Signet Ring (277210, -0.36 DPS) [quest]; Ring of Precision (1491, -0.42 DPS) [dungeon]; Theramore Signet (271911, -0.51 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.91 DPS) | yes | Ring of Precision (1491, -0.20 DPS) [dungeon]; Theramore Signet (271911, -0.29 DPS) [rep]; Pyrewood Signet Ring (277210, -0.36 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (58.7 DPS) | yes | Thornspike (6681, -0.39 DPS, sim-verified) [dungeon]; Vendetta (776, -0.42 DPS) [dungeon]; Torturing Poker (7682, -0.80 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (19.92 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.3 attack_power points (0.71 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.24 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 110.6. Weights run: 5.4s. Verify run: 7.9s. 544 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.794 ± 0.013, crit=1.101 ± 0.018 per rating point (14 rating = 1%, 15.411 per %), hit=2.088 ± 0.070 per rating point (10 rating = 1%, 20.878 per %), melee_haste=8.415 ± 1.051

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.7 attack_power points (2.37 DPS) | yes | Hawkeye's Helm (14591, -0.54 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.57 DPS) [world_drop]; Nightscape Headband (8176, -0.71 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 19.7 attack_power points (1.52 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Kaleidoscope Chain (13084, -0.66 DPS) [world_drop]; River Pride Choker (13087, -1.21 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.7 attack_power points (1.91 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Nightscape Shoulders (8192, -0.39 DPS) [crafted]; Mantle of Thieves (2264, -0.52 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (110.6 DPS) | yes | Hawkeye's Cloak (14593, -0.09 DPS) [world_drop]; Imperial Cloak (6432, -0.19 DPS) [dungeon]; Parachute Cloak (10518, -0.19 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 30.5 attack_power points (2.35 DPS) | yes | Nightscape Tunic (8175, -0.28 DPS) [crafted]; Dusky Leather Armor (7374, -0.42 DPS) [crafted]; Warden's Wraps (14601, -0.69 DPS) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.54 DPS) | yes | Hawkeye's Bracers (14590, -0.40 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.44 DPS) [dungeon]; Dusky Bracers (7378, -0.44 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (110.6 DPS) | yes | Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.02 DPS) [dungeon]; Skulker's Leather Gloves (252525, -3.61 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.31 DPS) | yes | Ogron's Sash (13117, -0.37 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.46 DPS) [rep]; Skulker's Leather Belt (252520, -0.61 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 37.7 attack_power points (2.91 DPS) | yes | Triprunner Dungarees (9624, -0.18 DPS) [quest]; Ferine Leggings (6690, -0.90 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.97 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 26.7 attack_power points (2.06 DPS) | yes | Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.24 DPS) [crafted]; Worn Running Boots (9398, -0.26 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 22.4 attack_power points (1.72 DPS) | yes | Assault Band (13095, -0.18 DPS) [world_drop]; Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Falcon's Hook (7552, -0.25 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.54 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.07 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (110.6 DPS) | yes | Coldrage Dagger (10761, -3.10 DPS, sim-verified) [dungeon]; The Ziggler (8006, -8.01 DPS) [world_drop]; Sacrificial Kris (3187, -8.83 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (32.03 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (110.6 DPS) | yes | Swiftwind (13038, -0.11 DPS) [world_drop]; Monolithic Bow (9426, -0.20 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.22 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 544, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 201.0. Weights run: 6.2s. Verify run: 23.5s. 681 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.770 ± 0.029, crit=1.475 ± 0.044 per rating point (14 rating = 1%, 20.654 per %), hit=2.134 ± 0.091 per rating point (10 rating = 1%, 21.340 per %), melee_haste=7.847 ± 1.338

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.5 attack_power points (3.03 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.04 DPS) [crafted]; Lordrec Helmet (10741, -0.21 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.0 attack_power points (2.78 DPS) | yes | Sentinel's Medallion (19539, -0.69 DPS, sim-verified) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 28.5 attack_power points (2.83 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.15 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.53 DPS) [world_drop]; Serpentskin Cloak (8259, -0.70 DPS) [world_drop]; Pridelord Cape (14673, -0.76 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 60.0 attack_power points (5.96 DPS) | yes | Warbear Harness (15064, -1.70 DPS) [crafted]; Blazewind Breastplate (11193, -1.77 DPS, sim-verified) [quest]; Charred Leather Tunic (19127, -2.79 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 30.5 attack_power points (3.03 DPS) | yes | Pridelord Bands (14672, -0.78 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.91 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.99 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Serpentskin Gloves (8260, -0.02 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, -3.94 DPS, sim-verified) [vendor] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 40.7 attack_power points (4.04 DPS) | yes | Skulker's Leather Waistguard (252474, -0.38 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.54 DPS) [crafted]; Sagebrush Girdle (17778, -1.40 DPS) [quest] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 60.0 attack_power points (5.96 DPS) | yes | Serpentskin Leggings (8262, -1.85 DPS) [world_drop]; Stormshroud Pants (15057, -1.86 DPS) [crafted]; Gryphon Rider's Leggings (9652, -2.11 DPS, sim-verified) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 35.4 attack_power points (3.51 DPS) | yes | Skulker's Leather Boots (252469, -0.14 DPS) [crafted]; Prowler's Leather Boots (252468, -0.29 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.89 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.3 attack_power points (4.11 DPS) | yes | Masons Fraternity Ring (9533, -1.64 DPS) [quest]; Mark of Kern (2262, -2.12 DPS) [dungeon]; Assault Band (13095, -2.12 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.9 attack_power points (2.57 DPS) | yes | Masons Fraternity Ring (9533, -0.11 DPS) [quest]; Mark of Kern (2262, -0.59 DPS) [dungeon]; Assault Band (13095, -0.59 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -1.28 DPS, sim-verified) [quest] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Widowmaker (4091, -1.11 DPS) [world_drop]; Julie's Dagger (6660, -1.90 DPS, sim-verified) [world_drop]; Gahz'rilla Fang (9467, -2.32 DPS) [dungeon] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.1 attack_power points (49.96 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.77 DPS) [world_drop]; Houndmaster's Bow (11628, -0.97 DPS) [dungeon]; Guttbuster (13139, -0.99 DPS, sim-verified) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 681, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 311.5. Weights run: 6.0s. Verify run: 65.3s. 1337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.652 ± 0.027, crit=1.518 ± 0.049 per rating point (14 rating = 1%, 21.255 per %), hit=1.990 ± 0.090 per rating point (10 rating = 1%, 19.901 per %), melee_haste=8.533 ± 1.637

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -3.90 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.5 attack_power points (3.17 DPS) | yes | Sentinel's Medallion (19538, -0.20 DPS) [rep]; Sunprism Pendant (20649, -0.80 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -16.61 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.30 DPS) [rep]; Dark Phantom Cape (13122, -2.39 DPS) [world_drop]; Stormpike Soldier's Cloak (19084, -2.49 DPS) [rep] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -13.52 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 42.4 attack_power points (5.07 DPS) | yes | Marshal's Leather Armsplints (16460, -0.12 DPS) [pvp]; Deepfury Bracers (13120, -1.63 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -2.11 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, -1.58 DPS, sim-verified) [rep] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 55.3 attack_power points (6.61 DPS) | yes | Marshal's Leather Cinch (16458, -0.54 DPS) [pvp]; Might of the Timbermaw (19044, -1.76 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -2.41 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -16.61 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 40.7 attack_power points (4.86 DPS) | yes | Marshal's Leather Footguards (231546, -0.52 DPS) [pvp]; Highlander's Leather Boots (20052, -0.55 DPS, sim-verified) [rep]; Albino Crocscale Boots (17728, -0.91 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 55.2 attack_power points (6.60 DPS) | yes | Protector's Band (19514, -2.99 DPS) [rep]; Masons Fraternity Ring (9533, -3.83 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 39.9 attack_power points (4.77 DPS) | yes | Protector's Band (19514, -1.46 DPS, sim-verified) [rep]; Masons Fraternity Ring (9533, -2.01 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.05 DPS, sim-verified) [quest] |
| main_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -5.78 DPS, sim-verified) [world_drop] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | The Lobotomizer (19324, -0.75 DPS, sim-verified) [rep] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.21 DPS) [world_drop]; Bloodseeker (19107, -0.63 DPS) [quest]; Skull Splitting Crossbow (13039, -0.94 DPS) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Electrified Dagger; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1337, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00530311251551031-3023020000000000000-005000000000000000)

Set DPS (verified): 761.7. Weights run: 6.6s. Verify run: 35.1s. 1337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.879 ± 0.038, crit=1.543 ± 0.059 per rating point (14 rating = 1%, 21.606 per %), hit=3.741 ± 0.161 per rating point (10 rating = 1%, 37.409 per %), melee_haste=13.659 ± 2.051

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 96.4 attack_power points (22.84 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -0.33 DPS) [pvp]; Outlaw's Collar (279253, -6.36 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.9 attack_power points (7.09 DPS) | yes | Sentinel's Medallion (19538, -0.41 DPS) [rep]; Sunprism Pendant (20649, -1.75 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (761.7 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -44.24 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (761.7 DPS) | yes | Cloak of the Honor Guard (20073, -1.13 DPS) [rep]; Cloak of Flames (3475, -2.55 DPS) [world_drop]; Dark Phantom Cape (13122, -4.74 DPS) [world_drop] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-verified (761.7 DPS) | yes | Knight-Captain's Leather Chestpiece (227058, +0.00 DPS) [pvp]; Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Dawn Armor (252483, -36.22 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 47.8 attack_power points (11.32 DPS) | yes | Marshal's Leather Armsplints (16460, -0.26 DPS) [pvp]; Deepfury Bracers (13120, -3.60 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -4.65 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Leather Handgrips (231544, -0.86 DPS) [pvp]; Knight-Lieutenant's Leather Gauntlets (16396, -2.85 DPS) [pvp] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 55.6 attack_power points (13.17 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -1.47 DPS) [crafted]; Skulker's Leather Waistguard (252474, -3.81 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (761.7 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -44.24 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 50.7 attack_power points (12.00 DPS) | yes | Marshal's Leather Footguards (231546, -2.21 DPS) [pvp]; Highlander's Leather Boots (20052, -2.87 DPS) [rep]; Albino Crocscale Boots (17728, -3.10 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 73.0 attack_power points (17.29 DPS) | yes | Protector's Band (19514, -9.27 DPS) [rep]; Masons Fraternity Ring (9533, -11.06 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 attack_power points (13.60 DPS) | yes | Protector's Band (19514, -4.09 DPS, sim-verified) [rep]; Masons Fraternity Ring (9533, -7.37 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (761.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -1.93 DPS, sim-verified) [quest] |
| main_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | sim-verified (761.7 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; The Lobotomizer (19324, -1.69 DPS) [rep] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 607.5 attack_power points (143.88 DPS) | yes | The Lobotomizer (19324, -3.15 DPS, sim-verified) [rep] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (761.7 DPS) | yes | Precisely Calibrated Boomstick (2100, -3.97 DPS) [world_drop]; Bloodseeker (19107, -5.00 DPS) [quest]; Eaglehorn Long Bow (13023, -5.74 DPS) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Leather Girdle; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Electrified Dagger; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1337, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 32510000000000000-0000000000000000000-000000000000000000)

Set DPS (verified): 38.9. Weights run: 4.7s. Verify run: 9.1s. 233 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.690 ± 0.007, crit=0.418 ± 0.006 per rating point (14 rating = 1%, 5.848 per %), hit=0.633 ± 0.019 per rating point (10 rating = 1%, 6.326 per %), melee_haste=3.510 ± 0.301

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.5 attack_power points (0.99 DPS) | yes | Defender's Leather Hood (252447, -0.45 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.1 attack_power points (0.75 DPS) | yes | Erudite's Amulet (277204, -0.50 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.4 attack_power points (0.62 DPS) | yes | Slime-encrusted Pads (6461, -0.63 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (38.9 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.23 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (38.9 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.11 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.4 attack_power points (0.62 DPS) | yes | Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.25 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.1 attack_power points (1.04 DPS) | yes | Brawler's Leather Gloves (252494, -0.25 DPS) [crafted]; Bristlebark Gloves (14572, -0.25 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.29 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.32 DPS) | yes | Brawler's Leather Belt (252428, -0.53 DPS) [crafted]; Murloc Scale Belt (5780, -0.80 DPS) [crafted]; Deviate Scale Belt (6468, -4.12 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.31 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (38.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.23 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.8 attack_power points (0.79 DPS) | yes | Bounty Hunter's Ring (5351, -0.42 DPS) [quest]; Demon Band (12054, -0.50 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.54 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.3 attack_power points (0.68 DPS) | yes | Bounty Hunter's Ring (5351, -0.31 DPS) [quest]; Demon Band (12054, -0.39 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (16.83 DPS) | yes | Tail Spike (6448, -1.92 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.17 DPS) [crafted]; Goblin Screwdriver (1936, -3.82 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.1 attack_power points (16.71 DPS) | yes | Tail Spike (6448, -0.38 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.8 attack_power points (0.50 DPS) | yes | Fine Longbow (11304, -0.22 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 233, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance

### Band 30 (troll, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 57.6. Weights run: 4.9s. Verify run: 9.2s. 412 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.716 ± 0.010, crit=0.701 ± 0.011 per rating point (14 rating = 1%, 9.815 per %), hit=1.052 ± 0.041 per rating point (10 rating = 1%, 10.522 per %), melee_haste=5.968 ± 0.572

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.2 attack_power points (1.19 DPS) | yes | Brawler's Leather Hood (252504, -0.24 DPS) [crafted]; Humbert's Helm (4724, -0.36 DPS) [world]; Defender's Leather Helm (252455, -0.36 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-verified (57.6 DPS) | yes | Scout's Medallion (19537, +0.00 DPS) [rep]; Kaleidoscope Chain (13084, -0.08 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (1.66 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS) [dungeon]; Barbaric Shoulders (5964, -0.71 DPS) [crafted]; Bristlebark Amice (14573, -0.73 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.33 DPS) [dungeon]; Swiftrunner Cape (6745, -0.33 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 24.0 attack_power points (1.67 DPS) | yes | Brawler's Leather Tunic (252508, -0.30 DPS) [crafted]; Panther Armor (6670, -0.46 DPS) [quest]; Brawler's Leather Armor (252490, -0.49 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.3 attack_power points (0.99 DPS) | yes | Barbaric Bracers (18948, -0.24 DPS) [crafted]; Insignia Bracers (6410, -0.28 DPS) [world_drop]; Cultist's Armguards (270032, -0.30 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (57.6 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (57.6 DPS) | yes | Defiler's Chain Girdle (20152, +0.00 DPS) [rep]; Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (57.6 DPS) | yes | Ferine Leggings (6690, +0.00 DPS) [dungeon]; Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (57.6 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Insignia Boots (4055, +0.00 DPS) [world_drop]; Warsong Boots (16977, +0.00 DPS) [quest] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | sim-verified (57.6 DPS) | yes | Legionnaire's Band (19513, +0.00 DPS) [rep]; Pyrewood Signet Ring (277210, -0.14 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | sim-verified (57.6 DPS) | yes | Legionnaire's Band (19513, +0.00 DPS) [rep]; Pyrewood Signet Ring (277210, +0.00 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornspike (6681, +0.00 DPS) [dungeon]; Vendetta (776, -0.42 DPS) [dungeon]; Torturing Poker (7682, -0.80 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (19.92 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.3 attack_power points (0.71 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.24 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 412, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 109.5. Weights run: 5.4s. Verify run: 8.0s. 543 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.794 ± 0.013, crit=1.101 ± 0.018 per rating point (14 rating = 1%, 15.411 per %), hit=2.088 ± 0.070 per rating point (10 rating = 1%, 20.878 per %), melee_haste=8.415 ± 1.051

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.7 attack_power points (2.37 DPS) | yes | Hawkeye's Helm (14591, -0.54 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.57 DPS) [world_drop]; Nightscape Headband (8176, -0.71 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 19.7 attack_power points (1.52 DPS) | yes | Ethereal Talisman (4430, -0.58 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Kaleidoscope Chain (13084, -0.66 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.7 attack_power points (1.91 DPS) | yes | Flintrock Shoulders (7755, -0.14 DPS) [dungeon]; Nightscape Shoulders (8192, -0.39 DPS) [crafted]; Mantle of Thieves (2264, -0.52 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (109.5 DPS) | yes | Imperial Cloak (6432, -0.09 DPS) [dungeon]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 30.5 attack_power points (2.35 DPS) | yes | Nightscape Tunic (8175, -0.28 DPS) [crafted]; Dusky Leather Armor (7374, -0.42 DPS) [crafted]; Warden's Wraps (14601, -0.69 DPS) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.54 DPS) | yes | Hawkeye's Bracers (14590, -0.40 DPS) [world_drop]; Imperial Leather Bracers (4061, -0.44 DPS) [dungeon]; Dusky Bracers (7378, -0.44 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (109.5 DPS) | yes | Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Imperial Leather Gloves (4063, -0.02 DPS) [dungeon]; Skulker's Leather Gloves (252525, -3.51 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.31 DPS) | yes | Ogron's Sash (13117, -0.37 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.46 DPS) [rep]; Skulker's Leather Belt (252520, -0.61 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 37.7 attack_power points (2.91 DPS) | yes | Triprunner Dungarees (9624, -0.18 DPS) [quest]; Ferine Leggings (6690, -0.90 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.97 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 26.7 attack_power points (2.06 DPS) | yes | Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.24 DPS) [crafted]; Worn Running Boots (9398, -0.26 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 22.4 attack_power points (1.72 DPS) | yes | Assault Band (13095, -0.18 DPS) [world_drop]; Ring of the Underwood (2951, -0.19 DPS) [world_drop]; Falcon's Hook (7552, -0.25 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.54 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.07 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (109.5 DPS) | yes | The Ziggler (8006, -8.01 DPS) [world_drop]; Tok'kar's Murloc Shanker (9680, -8.77 DPS) [quest]; Thornspike (6681, -9.58 DPS, sim-verified) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (31.31 DPS) | yes | The Ziggler (8006, -1.81 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (109.5 DPS) | yes | Swiftwind (13038, -0.11 DPS) [world_drop]; Monolithic Bow (9426, -0.20 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.22 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 543, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 199.8. Weights run: 6.2s. Verify run: 25.2s. 680 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.770 ± 0.029, crit=1.475 ± 0.044 per rating point (14 rating = 1%, 20.654 per %), hit=2.134 ± 0.091 per rating point (10 rating = 1%, 21.340 per %), melee_haste=7.847 ± 1.338

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.5 attack_power points (3.03 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.04 DPS) [crafted]; Undercity Reservist's Cap (20643, -0.10 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 28.0 attack_power points (2.78 DPS) | yes | Woven Ivy Necklace (19159, -0.60 DPS) [quest]; Scout's Medallion (19535, -0.67 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 28.5 attack_power points (2.83 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.15 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.53 DPS) [world_drop]; Serpentskin Cloak (8259, -0.70 DPS) [world_drop]; Pridelord Cape (14673, -0.76 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 60.0 attack_power points (5.96 DPS) | yes | Blazewind Breastplate (11193, -1.64 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.70 DPS) [crafted]; Charred Leather Tunic (19127, -2.79 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 30.5 attack_power points (3.03 DPS) | yes | Skulker's Leather Bracers (252540, -0.73 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.78 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.91 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Serpentskin Gloves (8260, -0.02 DPS) [world_drop]; First Sergeant's Leather Gauntlets (220857, -3.88 DPS, sim-verified) [vendor] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 40.7 attack_power points (4.04 DPS) | yes | Skulker's Leather Waistguard (252474, -0.38 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.54 DPS) [crafted]; Sagebrush Girdle (17778, -1.40 DPS) [quest] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 60.0 attack_power points (5.96 DPS) | yes | Serpentskin Leggings (8262, -1.74 DPS, sim-verified) [world_drop]; Stormshroud Pants (15057, -1.86 DPS) [crafted]; Basilisk Hide Pants (1718, -2.27 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 35.4 attack_power points (3.51 DPS) | yes | Skulker's Leather Boots (252469, -0.14 DPS) [crafted]; Prowler's Leather Boots (252468, -0.29 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.89 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.3 attack_power points (4.11 DPS) | yes | Masons Fraternity Ring (9533, -1.64 DPS) [quest]; White Bone Band (11862, -1.72 DPS) [quest]; Mark of Kern (2262, -2.12 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 25.9 attack_power points (2.57 DPS) | yes | Masons Fraternity Ring (9533, -0.11 DPS) [quest]; White Bone Band (11862, -0.19 DPS) [quest]; Mark of Kern (2262, -0.59 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.97 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Widowmaker (4091, -1.11 DPS) [world_drop]; Julie's Dagger (6660, -2.05 DPS, sim-verified) [world_drop]; Gahz'rilla Fang (9467, -2.32 DPS) [dungeon] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.1 attack_power points (49.96 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -0.73 DPS, sim-verified) [world_drop]; Skull Splitting Crossbow (13039, -0.77 DPS) [world_drop]; Houndmaster's Bow (11628, -0.97 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 680, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 312.2. Weights run: 6.0s. Verify run: 64.5s. 1342 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.743 ± 0.018, crit=1.820 ± 0.032 per rating point (14 rating = 1%, 25.485 per %), hit=4.118 ± 0.173 per rating point (10 rating = 1%, 41.185 per %), melee_haste=43.111 ± 1.940

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Champion's Leather Helm (227057, -0.23 DPS) [pvp]; Outlaw's Collar (279253, -2.05 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.7 attack_power points (1.23 DPS) | yes | Scout's Medallion (19534, -0.07 DPS) [rep]; Woven Ivy Necklace (19159, -0.26 DPS) [quest]; Sunprism Pendant (20649, -0.30 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -16.65 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.15 DPS) [rep]; Cloak of Flames (3475, -0.22 DPS) [world_drop]; Dark Phantom Cape (13122, -0.89 DPS) [world_drop] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Chestpiece (227060, +0.00 DPS) [pvp]; Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Dawn Armor (252483, -13.87 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 44.1 attack_power points (1.96 DPS) | yes | General's Leather Armsplints (16559, -0.04 DPS) [pvp]; Deepfury Bracers (13120, -0.62 DPS) [world_drop]; Shadowcraft Bracers (16710, -0.80 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -3.22 DPS, sim-verified) [vendor] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 59.5 attack_power points (2.64 DPS) | yes | General's Leather Girdle (16557, -0.33 DPS) [pvp]; Might of the Timbermaw (19044, -0.62 DPS) [crafted]; Skulker's Leather Waistguard (252474, -1.02 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -18.92 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 49.0 attack_power points (2.17 DPS) | yes | General's Leather Treads (231552, -0.47 DPS) [pvp]; Defiler's Leather Boots (20186, -0.53 DPS) [rep]; Albino Crocscale Boots (17728, -0.63 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 80.7 attack_power points (3.58 DPS) | yes | Legionnaire's Band (19510, -2.20 DPS) [rep]; Band of the Penitent (13217, -2.45 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 61.2 attack_power points (2.71 DPS) | yes | Legionnaire's Band (19511, -1.58 DPS) [rep]; Band of the Penitent (13217, -1.58 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -4.46 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Molten Heart of the Mountain (249470, -1.24 DPS, sim-verified) [crafted] |
| main_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -0.26 DPS) [rep] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 610.9 attack_power points (27.10 DPS) | yes | The Lobotomizer (19324, +0.00 DPS) [rep] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.16 DPS) [quest]; Eaglehorn Long Bow (13023, -1.29 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -1.34 DPS, sim-verified) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Frozen Heart of the Mountain; main_hand: Glacial Blade; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1342, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 00530311251551031-3023020000000000000-005000000000000000)

Set DPS (verified): 763.5. Weights run: 6.6s. Verify run: 37.1s. 1342 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.879 ± 0.038, crit=1.543 ± 0.059 per rating point (14 rating = 1%, 21.606 per %), hit=3.741 ± 0.161 per rating point (10 rating = 1%, 37.409 per %), melee_haste=13.659 ± 2.051

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 96.4 attack_power points (22.84 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Champion's Leather Helm (227057, -0.33 DPS) [pvp]; Outlaw's Collar (279253, -7.74 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 29.9 attack_power points (7.09 DPS) | yes | Scout's Medallion (19534, -0.41 DPS) [rep]; Woven Ivy Necklace (19159, -1.52 DPS) [quest]; Sunprism Pendant (20649, -1.75 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -47.20 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -1.13 DPS) [rep]; Cloak of Flames (3475, -2.55 DPS) [world_drop]; Dark Phantom Cape (13122, -4.74 DPS) [world_drop] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Chestpiece (227060, +0.00 DPS) [pvp]; Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Dawn Armor (252483, -40.25 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 47.8 attack_power points (11.32 DPS) | yes | General's Leather Armsplints (16559, -0.26 DPS) [pvp]; Deepfury Bracers (13120, -3.60 DPS) [world_drop]; Shadowcraft Bracers (16710, -4.65 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, -0.86 DPS) [pvp]; Blood Guard's Leather Vices (16499, -2.85 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.13 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 55.6 attack_power points (13.17 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -1.47 DPS) [crafted]; Skulker's Leather Waistguard (252474, -3.81 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -53.22 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 50.7 attack_power points (12.00 DPS) | yes | General's Leather Treads (231552, -2.21 DPS) [pvp]; Albino Crocscale Boots (17728, -3.10 DPS) [dungeon]; Defiler's Leather Boots (20186, -3.36 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 73.0 attack_power points (17.29 DPS) | yes | Legionnaire's Band (19510, -9.27 DPS) [rep]; Masons Fraternity Ring (9533, -11.06 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 attack_power points (13.60 DPS) | yes | Legionnaire's Band (19510, -4.04 DPS, sim-verified) [rep]; Masons Fraternity Ring (9533, -7.37 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Molten Heart of the Mountain (249470, -3.63 DPS, sim-verified) [crafted] |
| main_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -3.26 DPS, sim-verified) [rep] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 607.5 attack_power points (143.88 DPS) | yes | The Lobotomizer (19324, -3.24 DPS, sim-verified) [rep] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Precisely Calibrated Boomstick (2100, -2.45 DPS, sim-verified) [world_drop]; Bloodseeker (19107, -5.00 DPS) [quest]; Eaglehorn Long Bow (13023, -5.74 DPS) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Frozen Heart of the Mountain; main_hand: Glacial Blade; off_hand: Heartseeker; ranged: Blackcrow

No-known-source sample (15 of 1342, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

