# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-3152000000000000000-000000000000000000)

Set DPS (verified): 42.6. Weights run: 5.2s. Verify run: 10.6s. 234 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.791 ± 0.009, crit=0.460 ± 0.008 per rating point (14 rating = 1%, 6.441 per %), hit=0.629 ± 0.019 per rating point (10 rating = 1%, 6.292 per %), melee_haste=3.158 ± 0.294

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.3 attack_power points (1.03 DPS) | yes | Defender's Leather Hood (252447, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.7 attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.51 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.99 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.0 attack_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, -0.63 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (42.6 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.7 attack_power points (1.42 DPS) | yes | Brawler's Leather Armor (252490, -0.16 DPS) [crafted]; Prospector's Chestpiece (14562, -0.50 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.52 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.0 attack_power points (0.65 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (42.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.89 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.30 DPS) | yes | Brawler's Leather Belt (252428, -0.49 DPS) [crafted]; Murloc Scale Belt (5780, -0.77 DPS) [crafted]; Deviate Scale Belt (6468, -3.95 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (42.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.16 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (42.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.06 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.2 attack_power points (0.81 DPS) | yes | Demon Band (12054, -0.52 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.55 DPS) [dungeon]; The 1 Ring (8350, -0.60 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.7 attack_power points (0.70 DPS) | yes | Demon Band (12054, -0.25 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon]; The 1 Ring (8350, -0.50 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (16.51 DPS) | yes | Tail Spike (6448, -1.87 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.08 DPS) [crafted]; Goblin Screwdriver (1936, -3.72 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.5 attack_power points (16.43 DPS) | yes | Tail Spike (6448, -0.33 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.2 attack_power points (0.52 DPS) | yes | Fine Longbow (11304, -0.22 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.26 DPS) [crafted]; Light Bow (4576, -0.26 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 234, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 00000000000000000-3153030000012300000-000000000000000000)

Set DPS (verified): 65.4. Weights run: 5.3s. Verify run: 12.9s. 411 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.705 ± 0.011, crit=0.672 ± 0.012 per rating point (14 rating = 1%, 9.404 per %), hit=0.968 ± 0.038 per rating point (10 rating = 1%, 9.675 per %), melee_haste=5.457 ± 0.522

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.0 attack_power points (1.24 DPS) | yes | Brawler's Leather Hood (252504, -0.25 DPS) [crafted]; Defender's Leather Helm (252455, -0.37 DPS) [crafted]; Humbert's Helm (4724, -0.37 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.6 attack_power points (1.00 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; River Pride Choker (13087, -0.70 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 attack_power points (1.73 DPS) | yes | Mantle of Thieves (2264, -0.48 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.75 DPS) [crafted]; Bristlebark Amice (14573, -0.77 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (65.4 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.30 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.34 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (65.4 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Dusky Leather Armor (7374, +0.00 DPS) [crafted]; Brawler's Leather Tunic (252508, +0.00 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.2 attack_power points (1.04 DPS) | yes | Barbaric Bracers (18948, -0.25 DPS) [crafted]; Insignia Bracers (6410, -0.29 DPS) [world_drop]; Cultist's Armguards (270032, -0.31 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.4 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -0.41 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (65.4 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.81 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.4 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -0.54 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (65.4 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -0.79 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.2 attack_power points (1.18 DPS) | yes | Pyrewood Signet Ring (277210, -0.40 DPS) [quest]; Ring of Precision (1491, -0.44 DPS) [dungeon]; Theramore Signet (271911, -0.53 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.96 DPS) | yes | Ring of Precision (1491, -0.21 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.29 DPS, sim-verified) [quest]; Theramore Signet (271911, -0.30 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (65.4 DPS) | yes | Vendetta (776, -0.45 DPS) [dungeon]; Thornspike (6681, -0.65 DPS, sim-verified) [dungeon]; Torturing Poker (7682, -0.84 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (20.98 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 attack_power points (0.75 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.25 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Blackened Defias Armor; wrist: Hawkeye's Bracers; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-3153030000012321511-000000000000000000)

Set DPS (verified): 116.9. Weights run: 5.8s. Verify run: 16.3s. 544 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.740 ± 0.012, crit=0.974 ± 0.017 per rating point (14 rating = 1%, 13.642 per %), hit=1.488 ± 0.054 per rating point (10 rating = 1%, 14.875 per %), melee_haste=8.838 ± 0.753

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.1 attack_power points (2.65 DPS) | yes | Hawkeye's Helm (14591, -0.63 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.66 DPS) [world_drop]; Nightscape Headband (8176, -0.81 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 19.1 attack_power points (1.68 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.63 DPS) [world]; Kaleidoscope Chain (13084, -0.72 DPS) [world_drop]; River Pride Choker (13087, -1.33 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (2.12 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.59 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (116.9 DPS) | yes | Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Imperial Cloak (6432, -0.22 DPS) [dungeon]; Parachute Cloak (10518, -0.22 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 29.6 attack_power points (2.60 DPS) | yes | Nightscape Tunic (8175, -0.31 DPS) [crafted]; Dusky Leather Armor (7374, -0.46 DPS) [crafted]; Barbaric Harness (5739, -0.76 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.76 DPS) | yes | Imperial Leather Bracers (4061, -0.53 DPS) [dungeon]; Dusky Bracers (7378, -0.53 DPS) [crafted]; Hawkeye's Bracers (14590, -0.55 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (116.9 DPS) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.63 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.64 DPS) | yes | Ogron's Sash (13117, -0.50 DPS, sim-verified) [world_drop]; Highlander's Chain Girdle (20090, -0.53 DPS) [rep]; Skulker's Leather Belt (252520, -0.73 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 36.5 attack_power points (3.21 DPS) | yes | Triprunner Dungarees (9624, -0.20 DPS) [quest]; Ferine Leggings (6690, -0.93 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.07 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 26.1 attack_power points (2.30 DPS) | yes | Imperial Leather Boots (6431, -0.18 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.26 DPS) [crafted]; Excelsior Boots (4109, -0.31 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.9 attack_power points (1.93 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Ring of the Underwood (2951, -0.22 DPS) [world_drop]; Falcon's Hook (7552, -0.29 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.76 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (116.9 DPS) | yes | Coldrage Dagger (10761, -4.89 DPS, sim-verified) [dungeon]; The Ziggler (8006, -9.13 DPS) [world_drop]; Sacrificial Kris (3187, -10.06 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (36.49 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (116.9 DPS) | yes | Swiftwind (13038, -0.16 DPS) [world_drop]; Monolithic Bow (9426, -0.24 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.27 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 544, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 32200000000000000-3153030000012321541-000000000000000000)

Set DPS (verified): 193.3. Weights run: 6.3s. Verify run: 21.2s. 681 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.587 ± 0.012, crit=1.054 ± 0.019 per rating point (14 rating = 1%, 14.756 per %), hit=1.394 ± 0.055 per rating point (10 rating = 1%, 13.941 per %), melee_haste=7.156 ± 0.823

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.5 attack_power points (3.25 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.17 DPS) [crafted]; Lordrec Helmet (10741, -0.35 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.6 attack_power points (2.92 DPS) | yes | Sentinel's Medallion (19539, -0.75 DPS, sim-verified) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (3.02 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.13 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.58 DPS, sim-verified) [world_drop]; Serpentskin Cloak (8259, -0.72 DPS) [world_drop]; Pridelord Cape (14673, -0.74 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 46.7 attack_power points (5.33 DPS) | yes | Warbear Harness (15064, -0.81 DPS) [crafted]; Blazewind Breastplate (11193, -0.82 DPS) [quest]; Charred Leather Tunic (19127, -2.07 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.8 attack_power points (3.17 DPS) | yes | Pridelord Bands (14672, -0.79 DPS) [world_drop]; Skulker's Leather Bracers (252540, -0.81 DPS, sim-verified) [crafted]; Prowler's Leather Bracers (252539, -0.88 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -3.21 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 34.8 attack_power points (3.96 DPS) | yes | Skulker's Leather Waistguard (252474, -0.06 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.20 DPS) [crafted]; Highlander's Chain Girdle (20090, -1.23 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 46.7 attack_power points (5.33 DPS) | yes | Gryphon Rider's Leggings (9652, -0.63 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -0.95 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.53 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.7 attack_power points (3.62 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.15 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.83 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 33.9 attack_power points (3.87 DPS) | yes | Masons Fraternity Ring (9533, -1.34 DPS) [quest]; Mark of Kern (2262, -1.59 DPS) [dungeon]; Assault Band (13095, -1.59 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.3 attack_power points (2.77 DPS) | yes | Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Mark of Kern (2262, -0.49 DPS) [dungeon]; Assault Band (13095, -0.49 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.47 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.55 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Widowmaker (4091, -1.44 DPS) [world_drop]; Gahz'rilla Fang (9467, -2.66 DPS) [dungeon]; Julie's Dagger (6660, -2.73 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.1 attack_power points (57.38 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.67 DPS, sim-verified) [world_drop]; Guttbuster (13139, -0.74 DPS) [world_drop]; Houndmaster's Bow (11628, -0.82 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 681, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 32530100000000000-3153030000012321541-030000000000000000)

Set DPS (verified): 281.3. Weights run: 6.1s. Verify run: 43.5s. 1337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.532 ± 0.014, crit=1.189 ± 0.025 per rating point (14 rating = 1%, 16.645 per %), hit=1.848 ± 0.072 per rating point (10 rating = 1%, 18.485 per %), melee_haste=10.495 ± 1.136

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 57.1 attack_power points (7.31 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -0.45 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.9 attack_power points (3.19 DPS) | yes | Sentinel's Medallion (19538, -0.25 DPS) [rep]; Sunprism Pendant (20649, -0.84 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 58.0 attack_power points (7.42 DPS) | yes | Highlander's Leather Shoulders (20059, -0.06 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, -0.11 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, -0.94 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.17 DPS) [rep]; Stormpike Soldier's Cloak (19084, -2.43 DPS) [rep]; Dark Phantom Cape (13122, -2.56 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 91.8 attack_power points (11.75 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -2.16 DPS) [pvp]; Timbermaw Tunic (252484, -2.68 DPS, sim-verified) [crafted]; Knight-Captain's Leather Chestpiece (227058, -2.90 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 40.1 attack_power points (5.13 DPS) | yes | Marshal's Leather Armsplints (16460, -0.13 DPS) [pvp]; Deepfury Bracers (13120, -1.68 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -2.19 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, -3.28 DPS, sim-verified) [rep] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 50.6 attack_power points (6.48 DPS) | yes | Marshal's Leather Cinch (16458, -0.21 DPS) [pvp]; Might of the Timbermaw (19044, -1.20 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -2.20 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 74.7 attack_power points (9.55 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -0.71 DPS) [vendor]; Devilsaur Leggings (15062, -2.07 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 37.7 attack_power points (4.83 DPS) | yes | Highlander's Leather Boots (20052, -0.43 DPS) [rep]; Marshal's Leather Footguards (231546, -0.51 DPS) [pvp]; Skulker's Leather Boots (252469, -0.87 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 49.1 attack_power points (6.29 DPS) | yes | Protector's Band (19514, -2.60 DPS) [rep]; Masons Fraternity Ring (9533, -3.54 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 38.5 attack_power points (4.93 DPS) | yes | Protector's Band (19516, -1.88 DPS) [rep]; Masons Fraternity Ring (9533, -2.18 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.11 DPS, sim-verified) [quest] |
| main_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -9.68 DPS, sim-verified) [world_drop] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | The Lobotomizer (19324, -0.37 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.35 DPS) [quest]; Blackcrow (12651, -0.40 DPS, sim-verified) [dungeon]; Skull Splitting Crossbow (13039, -0.57 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Highlander's Leather Girdle; legs: Sentinel's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Frozen Heart of the Mountain; main_hand: Electrified Dagger; off_hand: Heartseeker

No-known-source sample (15 of 1337, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00520300200000000-3243030000012321451-005000000000000000)

Set DPS (verified): 774.1. Weights run: 6.6s. Verify run: 58.9s. 1337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.489 ± 0.016, crit=0.771 ± 0.025 per rating point (14 rating = 1%, 10.793 per %), hit=3.553 ± 0.075 per rating point (10 rating = 1%, 35.529 per %), melee_haste=14.685 ± 1.481

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.9 attack_power points (6.38 DPS) | yes | Sentinel's Medallion (19538, -0.65 DPS) [rep]; Stormpike Soldier's Pendant (19097, -1.76 DPS) [rep]; Sunprism Pendant (20649, -1.79 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Highlander's Leather Shoulders (20059, -2.44 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.23 DPS) [rep]; Cloak of Flames (3475, -1.75 DPS) [world_drop]; Stormpike Soldier's Cloak (19084, -4.71 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 101.4 attack_power points (26.03 DPS) | yes | Timbermaw Tunic (252484, -3.47 DPS, sim-verified) [crafted]; Field Marshal's Leather Chestpiece (231543, -4.20 DPS) [pvp]; Knight-Captain's Leather Chestpiece (227058, -5.42 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 40.4 attack_power points (10.37 DPS) | yes | Marshal's Leather Armsplints (16460, -0.28 DPS) [pvp]; Deepfury Bracers (13120, -3.51 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -4.63 DPS) [pvp] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (+9.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Theramore Gloves (271949, +0.00 DPS) [rep] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 44.8 attack_power points (11.50 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.22 DPS) [crafted]; Skulker's Leather Waistguard (252474, -2.76 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -11.86 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 41.9 attack_power points (10.76 DPS) | yes | Marshal's Leather Footguards (231546, -2.35 DPS) [pvp]; Skulker's Leather Boots (252469, -2.69 DPS) [crafted]; Highlander's Leather Boots (20052, -2.87 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 60.3 attack_power points (15.48 DPS) | yes | Protector's Band (19514, -7.89 DPS) [rep]; Masons Fraternity Ring (9533, -10.13 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 55.5 attack_power points (14.25 DPS) | yes | Protector's Band (19514, -6.66 DPS) [rep]; Masons Fraternity Ring (9533, -8.90 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+6.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.18 DPS, sim-verified) [quest] |
| main_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -24.73 DPS, sim-verified) [world_drop] |
| off_hand | Sentinel's Blade (19546) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Lobotomizer (19324, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (774.1 DPS) | yes | Bloodseeker (19107, -0.42 DPS) [quest]; Skull Splitting Crossbow (13039, -0.91 DPS) [world_drop]; Blackcrow (12651, -1.68 DPS, sim-verified) [dungeon] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Devilsaur Gauntlets; waist: Highlander's Leather Girdle; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Frozen Heart of the Mountain; main_hand: Electrified Dagger; off_hand: Sentinel's Blade

No-known-source sample (15 of 1337, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-3152000000000000000-000000000000000000)

Set DPS (verified): 41.8. Weights run: 5.2s. Verify run: 10.6s. 233 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.791 ± 0.009, crit=0.460 ± 0.008 per rating point (14 rating = 1%, 6.441 per %), hit=0.629 ± 0.019 per rating point (10 rating = 1%, 6.292 per %), melee_haste=3.158 ± 0.294

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.3 attack_power points (1.03 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.7 attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.54 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.0 attack_power points (0.65 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (41.8 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (41.8 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.65 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.0 attack_power points (0.65 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.7 attack_power points (1.06 DPS) | yes | Brawler's Leather Gloves (252494, -0.26 DPS) [crafted]; Bristlebark Gloves (14572, -0.27 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.29 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.30 DPS) | yes | Brawler's Leather Belt (252428, -0.49 DPS) [crafted]; Murloc Scale Belt (5780, -0.77 DPS) [crafted]; Deviate Scale Belt (6468, -3.65 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.8 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.85 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.8 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.77 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.2 attack_power points (0.81 DPS) | yes | Bounty Hunter's Ring (5351, -0.42 DPS) [quest]; Demon Band (12054, -0.52 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.55 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.7 attack_power points (0.70 DPS) | yes | Bounty Hunter's Ring (5351, -0.31 DPS) [quest]; Demon Band (12054, -0.41 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (16.51 DPS) | yes | Tail Spike (6448, -1.87 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.08 DPS) [crafted]; Goblin Screwdriver (1936, -3.72 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.5 attack_power points (16.43 DPS) | yes | Tail Spike (6448, +0.00 DPS) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.2 attack_power points (0.52 DPS) | yes | Deadly Blunderbuss (4369, -0.26 DPS) [crafted]; Light Bow (4576, -0.26 DPS) [world_drop]; Fine Longbow (11304, -0.26 DPS, sim-verified) [vendor] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 233, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-3153030000012300000-000000000000000000)

Set DPS (verified): 64.2. Weights run: 5.3s. Verify run: 11.1s. 412 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.705 ± 0.011, crit=0.672 ± 0.012 per rating point (14 rating = 1%, 9.404 per %), hit=0.968 ± 0.038 per rating point (10 rating = 1%, 9.675 per %), melee_haste=5.457 ± 0.522

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.0 attack_power points (1.24 DPS) | yes | Brawler's Leather Hood (252504, -0.23 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.37 DPS) [crafted]; Humbert's Helm (4724, -0.37 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.6 attack_power points (1.00 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.12 DPS) [world]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; River Pride Choker (13087, -0.70 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 attack_power points (1.73 DPS) | yes | Mantle of Thieves (2264, -0.48 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.75 DPS) [crafted]; Bristlebark Amice (14573, -0.77 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.34 DPS) [dungeon]; Swiftrunner Cape (6745, -0.34 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.9 attack_power points (1.74 DPS) | yes | Brawler's Leather Tunic (252508, -0.30 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.48 DPS) [quest]; Brawler's Leather Armor (252490, -0.51 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.2 attack_power points (1.04 DPS) | yes | Barbaric Bracers (18948, -0.23 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.29 DPS) [world_drop]; Cultist's Armguards (270032, -0.31 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.2 attack_power points (1.26 DPS) | yes | Toughened Leather Gloves (4253, -0.07 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.09 DPS) [crafted]; Cartilage Shapers (273054, -0.14 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.75 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.19 DPS) [crafted]; Prowler's Leather Belt (252459, -0.35 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.90 DPS) | yes | Petrolspill Leggings (9509, -0.16 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.16 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.27 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.6 attack_power points (1.21 DPS) | yes | Insignia Boots (4055, -0.22 DPS) [world_drop]; Warsong Boots (16977, -0.22 DPS) [quest]; Brawler's Leather Boots (252439, -0.23 DPS) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.2 attack_power points (1.18 DPS) | yes | Band of the Fist (17694, -0.39 DPS) [quest]; Pyrewood Signet Ring (277210, -0.40 DPS) [quest]; Ring of Precision (1491, -0.44 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (0.96 DPS) | yes | Band of the Fist (17694, -0.17 DPS) [quest]; Pyrewood Signet Ring (277210, -0.18 DPS) [quest]; Ring of Precision (1491, -0.21 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornspike (6681, -0.43 DPS, sim-verified) [dungeon]; Vendetta (776, -0.45 DPS) [dungeon]; Torturing Poker (7682, -0.84 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (20.98 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 attack_power points (0.75 DPS) | yes | Double-barreled Shotgun (2098, -0.15 DPS) [world_drop]; Gravespike Repeater (273033, -0.16 DPS) [world]; Silver Star (3463, -0.25 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 412, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 00000000000000000-3153030000012321511-000000000000000000)

Set DPS (verified): 115.5. Weights run: 5.8s. Verify run: 16.3s. 543 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.740 ± 0.012, crit=0.974 ± 0.017 per rating point (14 rating = 1%, 13.642 per %), hit=1.488 ± 0.054 per rating point (10 rating = 1%, 14.875 per %), melee_haste=8.838 ± 0.753

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.1 attack_power points (2.65 DPS) | yes | Hawkeye's Helm (14591, -0.62 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.66 DPS) [world_drop]; Nightscape Headband (8176, -0.81 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 19.1 attack_power points (1.68 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.63 DPS) [world]; Ethereal Talisman (4430, -0.63 DPS) [quest]; Kaleidoscope Chain (13084, -0.72 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (2.12 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.59 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Parachute Cloak (10518, -0.11 DPS) [crafted]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 29.6 attack_power points (2.60 DPS) | yes | Nightscape Tunic (8175, -0.31 DPS) [crafted]; Dusky Leather Armor (7374, -0.46 DPS) [crafted]; Barbaric Harness (5739, -0.76 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.76 DPS) | yes | Hawkeye's Bracers (14590, -0.49 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.53 DPS) [dungeon]; Dusky Bracers (7378, -0.53 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.60 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.64 DPS) | yes | Ogron's Sash (13117, -0.48 DPS, sim-verified) [world_drop]; Defiler's Chain Girdle (20152, -0.53 DPS) [rep]; Skulker's Leather Belt (252520, -0.73 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 36.5 attack_power points (3.21 DPS) | yes | Triprunner Dungarees (9624, -0.20 DPS) [quest]; Ferine Leggings (6690, -0.93 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.07 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 26.1 attack_power points (2.30 DPS) | yes | Imperial Leather Boots (6431, -0.18 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.26 DPS) [crafted]; Excelsior Boots (4109, -0.31 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.9 attack_power points (1.93 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Ring of the Underwood (2951, -0.22 DPS) [world_drop]; Falcon's Hook (7552, -0.29 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.76 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | The Ziggler (8006, -9.13 DPS) [world_drop]; Tok'kar's Murloc Shanker (9680, -10.02 DPS) [quest]; Claw of the Shadowmancer (2912, -11.84 DPS, sim-verified) [world_drop] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (35.67 DPS) | yes | The Ziggler (8006, -2.59 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (115.5 DPS) | yes | Swiftwind (13038, -0.16 DPS) [world_drop]; Monolithic Bow (9426, -0.24 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.27 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 543, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32200000000000000-3153030000012321541-000000000000000000)

Set DPS (verified): 192.5. Weights run: 6.3s. Verify run: 21.2s. 680 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.587 ± 0.012, crit=1.054 ± 0.019 per rating point (14 rating = 1%, 14.756 per %), hit=1.394 ± 0.055 per rating point (10 rating = 1%, 13.941 per %), melee_haste=7.156 ± 0.823

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.5 attack_power points (3.25 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.11 DPS) [quest]; Helm of Fire (8348, -0.17 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.6 attack_power points (2.92 DPS) | yes | Woven Ivy Necklace (19159, -0.63 DPS, sim-verified) [quest]; Scout's Medallion (19535, -0.75 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (3.02 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.13 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.54 DPS) [world_drop]; Serpentskin Cloak (8259, -0.72 DPS) [world_drop]; Pridelord Cape (14673, -0.74 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 46.7 attack_power points (5.33 DPS) | yes | Warbear Harness (15064, -0.81 DPS) [crafted]; Blazewind Breastplate (11193, -0.82 DPS) [quest]; Charred Leather Tunic (19127, -2.07 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.8 attack_power points (3.17 DPS) | yes | Skulker's Leather Bracers (252540, -0.74 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.79 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.88 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; First Sergeant's Leather Gauntlets (220857, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -2.81 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 34.8 attack_power points (3.96 DPS) | yes | Skulker's Leather Waistguard (252474, -0.06 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.20 DPS) [crafted]; Defiler's Chain Girdle (20152, -1.23 DPS) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 46.7 attack_power points (5.33 DPS) | yes | Serpentskin Leggings (8262, -0.95 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.53 DPS) [world_drop]; Triprunner Dungarees (9624, -1.73 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.7 attack_power points (3.62 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.15 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.83 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 33.9 attack_power points (3.87 DPS) | yes | White Bone Band (11862, -1.13 DPS) [quest]; Masons Fraternity Ring (9533, -1.34 DPS) [quest]; Mark of Kern (2262, -1.59 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.3 attack_power points (2.77 DPS) | yes | White Bone Band (11862, -0.03 DPS) [quest]; Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Mark of Kern (2262, -0.49 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.49 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (192.5 DPS) | yes | Widowmaker (4091, -1.44 DPS) [world_drop]; Gahz'rilla Fang (9467, -2.66 DPS) [dungeon]; Julie's Dagger (6660, -2.69 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.1 attack_power points (57.38 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.64 DPS, sim-verified) [world_drop]; Guttbuster (13139, -0.74 DPS) [world_drop]; Houndmaster's Bow (11628, -0.82 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 680, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 32530100000000000-3153030000012321541-030000000000000000)

Set DPS (verified): 277.6. Weights run: 6.1s. Verify run: 48.0s. 1342 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.745 ± 0.018, crit=1.596 ± 0.030 per rating point (14 rating = 1%, 22.350 per %), hit=1.869 ± 0.083 per rating point (10 rating = 1%, 18.687 per %), melee_haste=9.503 ± 1.292

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 70.2 attack_power points (7.61 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -1.21 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.7 attack_power points (3.00 DPS) | yes | Scout's Medallion (19534, -0.16 DPS) [rep]; Woven Ivy Necklace (19159, -0.65 DPS) [quest]; Sunprism Pendant (20649, -0.73 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 69.5 attack_power points (7.53 DPS) | yes | Champion's Leather Shoulders (227056, -0.70 DPS) [pvp]; Defiler's Leather Shoulders (20194, -0.87 DPS) [rep]; Warlord's Leather Spaulders (231551, -1.53 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.37 DPS) [rep]; Dark Phantom Cape (13122, -2.17 DPS) [world_drop]; Frostwolf Legionnaire's Cloak (19083, -2.40 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 105.6 attack_power points (11.45 DPS) | yes | Timbermaw Tunic (252484, -2.04 DPS) [crafted]; Warlord's Leather Breastplate (231549, -2.08 DPS) [pvp]; Legionnaire's Leather Chestpiece (227060, -3.31 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 44.1 attack_power points (4.79 DPS) | yes | General's Leather Armsplints (16559, -0.11 DPS) [pvp]; Deepfury Bracers (13120, -1.52 DPS) [world_drop]; Shadowcraft Bracers (16710, -1.95 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -1.77 DPS, sim-verified) [vendor] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 56.4 attack_power points (6.11 DPS) | yes | General's Leather Girdle (16557, -0.45 DPS) [pvp]; Might of the Timbermaw (19044, -1.17 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -2.16 DPS) [crafted] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 71.2 attack_power points (7.72 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -0.31 DPS) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 42.2 attack_power points (4.58 DPS) | yes | General's Leather Treads (231552, -0.42 DPS) [pvp]; Defiler's Leather Boots (20186, -0.58 DPS) [rep]; Albino Crocscale Boots (17728, -0.80 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 55.0 attack_power points (5.97 DPS) | yes | Legionnaire's Band (19510, -2.59 DPS) [rep]; Masons Fraternity Ring (9533, -3.32 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 38.7 attack_power points (4.20 DPS) | yes | Legionnaire's Band (19510, -0.53 DPS, sim-verified) [rep]; Masons Fraternity Ring (9533, -1.55 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.11 DPS) [crafted] |
| main_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -0.63 DPS) [rep] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | 607.7 attack_power points (65.92 DPS) | yes | The Lobotomizer (19324, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.40 DPS, sim-verified) [dungeon]; Bloodseeker (19107, -0.46 DPS) [quest]; Eaglehorn Long Bow (13023, -0.76 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Defiler's Leather Girdle; legs: Outrider's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: Glacial Blade; off_hand: Heartseeker

No-known-source sample (15 of 1342, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 00520300200000000-3243030000012321451-005000000000000000)

Set DPS (verified): 776.1. Weights run: 6.6s. Verify run: 58.2s. 1342 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.489 ± 0.016, crit=0.771 ± 0.025 per rating point (14 rating = 1%, 10.793 per %), hit=3.553 ± 0.075 per rating point (10 rating = 1%, 35.529 per %), melee_haste=14.685 ± 1.481

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-verified (776.1 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -5.65 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.9 attack_power points (6.38 DPS) | yes | Scout's Medallion (19534, -0.65 DPS) [rep]; Woven Ivy Necklace (19159, -1.25 DPS) [quest]; Frostwolf Legionnaire's Pendant (19095, -1.76 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (776.1 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Defiler's Leather Shoulders (20194, -3.00 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (776.1 DPS) | yes | Deathguard's Cloak (20068, -0.23 DPS) [rep]; Cloak of Flames (3475, -1.75 DPS) [world_drop]; Frostwolf Legionnaire's Cloak (19083, -4.71 DPS) [rep] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 101.4 attack_power points (26.03 DPS) | yes | Warlord's Leather Breastplate (231549, -4.20 DPS) [pvp]; Legionnaire's Leather Chestpiece (227060, -5.42 DPS) [pvp]; Timbermaw Tunic (252484, -7.17 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 40.4 attack_power points (10.37 DPS) | yes | General's Leather Armsplints (16559, -0.28 DPS) [pvp]; Deepfury Bracers (13120, -3.51 DPS) [world_drop]; Shadowcraft Bracers (16710, -4.63 DPS) [dungeon] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (776.1 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; Raider Gloves (272100, +0.00 DPS) [vendor] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 44.8 attack_power points (11.50 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -0.22 DPS) [crafted]; Skulker's Leather Waistguard (252474, -2.76 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (776.1 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Warbear Woolies (15065, -20.22 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 41.9 attack_power points (10.76 DPS) | yes | General's Leather Treads (231552, -2.35 DPS) [pvp]; Skulker's Leather Boots (252469, -2.69 DPS) [crafted]; Defiler's Leather Boots (20186, -5.17 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 60.3 attack_power points (15.48 DPS) | yes | Legionnaire's Band (19510, -7.89 DPS) [rep]; White Bone Band (11862, -9.32 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 55.5 attack_power points (14.25 DPS) | yes | Legionnaire's Band (19510, -6.70 DPS, sim-verified) [rep]; White Bone Band (11862, -8.09 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (776.1 DPS) | yes | Frozen Heart of the Mountain (249469, -1.77 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (776.1 DPS) | yes | Frozen Heart of the Mountain (249469, -4.23 DPS, sim-verified) [crafted] |
| main_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-verified (776.1 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -24.99 DPS, sim-verified) [world_drop] |
| off_hand | Scout's Blade (19542) | Warsong Outriders [rep] | sim-verified (776.1 DPS) | yes | The Lobotomizer (19324, -4.13 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (776.1 DPS) | yes | Bloodseeker (19107, -0.42 DPS) [quest]; Skull Splitting Crossbow (13039, -0.91 DPS) [world_drop]; Blackcrow (12651, -2.98 DPS, sim-verified) [dungeon] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Devilsaur Gauntlets; waist: Defiler's Leather Girdle; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Glacial Blade; off_hand: Scout's Blade

No-known-source sample (15 of 1342, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

