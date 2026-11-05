# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32510000000000000-0000000000000000000-000000000000000000)

Set DPS (verified): 43.6. Weights run: 6.9s. Verify run: 12.6s. 241 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.717 ± 0.008, crit=0.435 ± 0.007 per rating point (14 rating = 1%, 6.092 per %), hit=0.598 ± 0.017 per rating point (10 rating = 1%, 5.976 per %), melee_haste=3.643 ± 0.265

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.7 attack_power points (1.05 DPS) | yes | Defender's Leather Hood (252447, -0.51 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.3 attack_power points (0.79 DPS) | yes | Erudite's Amulet (277204, -0.55 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.02 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.6 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (43.6 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 18.9 attack_power points (1.44 DPS) | yes | Brawler's Leather Armor (252490, -0.14 DPS) [crafted]; Prospector's Chestpiece (14562, -0.50 DPS) [world_drop]; Defender's Leather Armor (252434, -0.51 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.6 attack_power points (0.66 DPS) | yes | Bravo's Armbands (270015, -0.03 DPS) [quest]; Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (43.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.82 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.37 DPS) | yes | Brawler's Leather Belt (252428, -0.54 DPS) [crafted]; Murloc Scale Belt (5780, -0.83 DPS) [crafted]; Deviate Scale Belt (6468, -3.93 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (43.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.06 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (43.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.95 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.9 attack_power points (0.83 DPS) | yes | Demon Band (12054, -0.52 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; The 1 Ring (8350, -0.62 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.3 attack_power points (0.71 DPS) | yes | Demon Band (12054, -0.27 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon]; The 1 Ring (8350, -0.50 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.46 DPS) | yes | Tail Spike (6448, -1.99 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.28 DPS) [crafted]; Goblin Screwdriver (1936, -3.95 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.2 attack_power points (17.35 DPS) | yes | Tail Spike (6448, -0.23 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.9 attack_power points (0.52 DPS) | yes | Fine Longbow (11304, -0.25 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.26 DPS) [crafted]; Light Bow (4576, -0.26 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 241, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 76.6. Weights run: 7.7s. Verify run: 9.3s. 422 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.599 ± 0.009, crit=0.617 ± 0.010 per rating point (14 rating = 1%, 8.632 per %), hit=0.906 ± 0.042 per rating point (10 rating = 1%, 9.055 per %), melee_haste=5.114 ± 0.579

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.0 attack_power points (1.49 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.37 DPS) [crafted]; Humbert's Helm (4724, -0.45 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 12.8 attack_power points (1.19 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.82 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.6 attack_power points (2.11 DPS) | yes | Mantle of Thieves (2264, -0.61 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.90 DPS) [crafted]; Bristlebark Amice (14573, -0.93 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (76.6 DPS) | yes | Tigerstrike Mantle (13108, -0.13 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.35 DPS) [pvp]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.4 attack_power points (2.09 DPS) | yes | Brawler's Leather Tunic (252508, -0.35 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.45 DPS) [quest]; Brawler's Leather Armor (252490, -0.58 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.6 attack_power points (1.27 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS, sim-verified) [crafted]; Cultist's Armguards (270032, -0.34 DPS) [quest]; Insignia Bracers (6410, -0.37 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.6 attack_power points (1.55 DPS) | yes | Heavy Earthen Gloves (7359, -0.06 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.19 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (2.24 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.34 DPS) [crafted]; Prowler's Leather Belt (252459, -0.50 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.43 DPS) | yes | Petrolspill Leggings (9509, -0.31 DPS, sim-verified) [dungeon]; Troll's Bane Leggings (13114, -0.34 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.43 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.8 attack_power points (1.48 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.28 DPS) [world_drop]; Disjointed Shoes (277226, -0.35 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.6 attack_power points (1.46 DPS) | yes | Pyrewood Signet Ring (277210, -0.52 DPS) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon]; Theramore Signet (271911, -0.62 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.8 attack_power points (1.20 DPS) | yes | Ring of Precision (1491, -0.30 DPS) [dungeon]; Theramore Signet (271911, -0.35 DPS) [rep]; Pyrewood Signet Ring (277210, -0.60 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Thornspike (6681) | Razorfen Kraul: Aggem Thorncurse [dungeon] | sim-verified (76.6 DPS) | yes | Toxic Revenger (9453, -0.08 DPS) [dungeon]; Vendetta (776, -0.73 DPS) [dungeon]; Torturing Poker (7682, -1.15 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (26.84 DPS) | yes | Toxic Revenger (9453, +0.00 DPS) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.6 attack_power points (0.90 DPS) | yes | Gravespike Repeater (273033, -0.15 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.30 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Thornspike; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 130.2. Weights run: 7.0s. Verify run: 22.7s. 562 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.667 ± 0.011, crit=0.987 ± 0.015 per rating point (14 rating = 1%, 13.811 per %), hit=1.728 ± 0.067 per rating point (10 rating = 1%, 17.275 per %), melee_haste=7.941 ± 0.940

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.3 attack_power points (2.79 DPS) | yes | Hawkeye's Helm (14591, -0.68 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.73 DPS) [world_drop]; Nightscape Headband (8176, -0.89 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Kaleidoscope Chain (13084, -0.73 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.3 attack_power points (2.22 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.48 DPS) [crafted]; Mantle of Thieves (2264, -0.63 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Imperial Cloak (6432, -0.25 DPS) [dungeon]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (130.2 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.32 DPS) [crafted]; Dusky Leather Armor (7374, -0.48 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Hawkeye's Bracers (14590, -0.63 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.63 DPS) [dungeon]; Dusky Bracers (7378, -0.63 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dog Training Gloves (7756, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -4.15 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.85 DPS) | yes | Ogron's Sash (13117, -0.57 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.58 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.86 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.0 attack_power points (3.33 DPS) | yes | Triprunner Dungarees (9624, -0.19 DPS) [quest]; Ferine Leggings (6690, -0.86 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.11 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.3 attack_power points (2.41 DPS) | yes | Imperial Leather Boots (6431, -0.19 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.32 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.3 attack_power points (2.03 DPS) | yes | Assault Band (13095, -0.13 DPS) [world_drop]; Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Falcon's Hook (7552, -0.32 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.13 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -3.42 DPS, sim-verified) [dungeon]; The Ziggler (8006, -9.87 DPS) [world_drop]; Sacrificial Kris (3187, -10.88 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (39.47 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.22 DPS) [world_drop]; Monolithic Bow (9426, -0.29 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.32 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 214.3. Weights run: 7.7s. Verify run: 25.3s. 707 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.646 ± 0.027, crit=1.293 ± 0.041 per rating point (14 rating = 1%, 18.103 per %), hit=2.117 ± 0.082 per rating point (10 rating = 1%, 21.170 per %), melee_haste=9.073 ± 1.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 54.1 attack_power points (6.20 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.75 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.86 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.4 attack_power points (3.02 DPS) | yes | Sentinel's Medallion (19539, -0.76 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.98 DPS, sim-verified) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 27.1 attack_power points (3.10 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Phytoskin Spaulders (17749, -0.09 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (214.3 DPS) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 57.3 attack_power points (6.56 DPS) | yes | Blazewind Breastplate (11193, -1.42 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.91 DPS) [crafted]; Quillward Harness (10583, -2.06 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.7 attack_power points (3.29 DPS) | yes | Skulker's Leather Bracers (252540, -0.77 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.83 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.94 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, -3.12 DPS, sim-verified) [vendor] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 38.1 attack_power points (4.36 DPS) | yes | Girdle of Beastial Fury (11686, -0.01 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.35 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.50 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 57.3 attack_power points (6.56 DPS) | yes | Gryphon Rider's Leggings (9652, -1.34 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -2.06 DPS) [world_drop]; Stormshroud Pants (15057, -2.41 DPS) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 32.9 attack_power points (3.77 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Prowler's Leather Boots (252468, -0.21 DPS) [crafted]; Whisperwalk Boots (20255, -0.38 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.2 attack_power points (4.72 DPS) | yes | Masons Fraternity Ring (9533, -2.08 DPS) [quest]; Mark of Kern (2262, -2.42 DPS) [dungeon]; Assault Band (13095, -2.42 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.8 attack_power points (2.84 DPS) | yes | Masons Fraternity Ring (9533, -0.20 DPS) [quest]; Mark of Kern (2262, -0.55 DPS) [dungeon]; Assault Band (13095, -0.55 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, -2.80 DPS, sim-verified) [world_drop]; Lifeforce Dirk (10750, -4.80 DPS) [quest]; Hookfang Shanker (11635, -5.15 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (58.59 DPS) | yes | Lifeforce Dirk (10750, -2.11 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.69 DPS) [dungeon]; Guttbuster (13139, -0.79 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.92 DPS, sim-verified) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 707, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 345.0. Weights run: 7.5s. Verify run: 92.4s. 1667 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.782 ± 0.020, crit=1.993 ± 0.036 per rating point (14 rating = 1%, 27.898 per %), hit=4.016 ± 0.176 per rating point (10 rating = 1%, 40.160 per %), melee_haste=37.660 ± 1.663

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-verified (345.0 DPS) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -18.98 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.52 DPS) [rep]; Cloak of Flames (3475, -1.69 DPS) [world_drop]; Cape of the Black Baron (13340, -1.81 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Chestpiece (227058, +0.00 DPS) [pvp]; Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Dawn Armor (252483, -16.25 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.06 DPS) [pvp]; Deepfury Bracers (13120, -0.85 DPS) [world_drop]; Blackmist Armguards (12966, -2.78 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, +0.00 DPS) [rep] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 117.7 attack_power points (7.11 DPS) | yes | Highlander's Leather Girdle (20045, -1.78 DPS, sim-verified) [rep]; Marshal's Leather Cinch (16458, -3.92 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -18.24 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 49.5 attack_power points (2.99 DPS) | yes | Swiftwalker Boots (12553, -0.49 DPS) [dungeon]; Pads of the Dread Wolf (13210, -0.57 DPS) [dungeon]; Marshal's Leather Footguards (231546, -0.62 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.32 DPS) [dungeon]; Painweaver Band (13098, -2.30 DPS) [dungeon]; Naglering (11669, -7.75 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.30 DPS) [dungeon]; Painweaver Band (13098, -1.28 DPS) [dungeon]; Channeler's Ring (272406, -1.51 DPS) [vendor] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -1.79 DPS, sim-verified) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Finkle's Skinner (12709, -8.54 DPS, sim-verified) [dungeon] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Heartseeker (12783, -5.93 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Riphook (12653, -1.42 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -1.43 DPS, sim-verified) [world_drop]; Bloodseeker (19107, -1.51 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Hand of Justice; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 815.2. Weights run: 9.0s. Verify run: 98.3s. 1667 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.909 ± 0.037, crit=1.750 ± 0.057 per rating point (14 rating = 1%, 24.502 per %), hit=2.054 ± 0.090 per rating point (10 rating = 1%, 20.544 per %), melee_haste=16.537 ± 2.134

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 76.8 attack_power points (18.92 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -2.77 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Jewel (11933, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Will of the Martyr (17044, -0.08 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -48.49 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, +0.00 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.23 DPS) [rep]; Blackveil Cape (11626, -3.75 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -42.56 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 48.4 attack_power points (11.91 DPS) | yes | Marshal's Leather Armsplints (16460, -0.27 DPS) [pvp]; Deepfury Bracers (13120, -3.78 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -4.86 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Leather Handgrips (231544, -0.24 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.74 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 81.2 attack_power points (19.99 DPS) | yes | Marshal's Leather Cinch (16458, -5.89 DPS) [pvp]; Might of the Timbermaw (19044, -7.72 DPS) [crafted]; Highlander's Leather Girdle (20045, -7.79 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -48.49 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 46.3 attack_power points (11.39 DPS) | yes | Swiftwalker Boots (12553, -0.43 DPS) [dungeon]; Marshal's Leather Footguards (231546, -1.05 DPS) [pvp]; Pads of the Dread Wolf (13210, -1.54 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -4.55 DPS) [dungeon]; Painweaver Band (13098, -4.57 DPS) [dungeon]; Naglering (11669, -15.43 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.66 DPS) [dungeon]; Painweaver Band (13098, -1.67 DPS) [dungeon]; Protector's Band (19514, -3.22 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heart of Wyrmthalak (22321, -5.16 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+602.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Alcor's Sunrazor (14555, -0.70 DPS) [world_drop] |
| off_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | sim-verified (815.2 DPS) | yes | - |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, +0.00 DPS) [world_drop]; Bloodseeker (19107, -1.01 DPS) [quest]; Riphook (12653, -1.05 DPS) [dungeon] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; main_hand: Felstriker; off_hand: Electrified Dagger; ranged: Blackcrow

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 32510000000000000-0000000000000000000-000000000000000000)

Set DPS (verified): 43.6. Weights run: 6.9s. Verify run: 12.4s. 240 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.717 ± 0.008, crit=0.435 ± 0.007 per rating point (14 rating = 1%, 6.092 per %), hit=0.598 ± 0.017 per rating point (10 rating = 1%, 5.976 per %), melee_haste=3.643 ± 0.265

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.7 attack_power points (1.05 DPS) | yes | Defender's Leather Hood (252447, -0.49 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.3 attack_power points (0.79 DPS) | yes | Erudite's Amulet (277204, -0.50 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.6 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (43.6 DPS) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (43.6 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.20 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.6 attack_power points (0.66 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.3 attack_power points (1.09 DPS) | yes | Brawler's Leather Gloves (252494, -0.26 DPS) [crafted]; Bristlebark Gloves (14572, -0.26 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.31 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.37 DPS) | yes | Brawler's Leather Belt (252428, -0.54 DPS) [crafted]; Murloc Scale Belt (5780, -0.83 DPS) [crafted]; Deviate Scale Belt (6468, -3.94 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (43.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.14 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (43.6 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.23 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.9 attack_power points (0.83 DPS) | yes | Bounty Hunter's Ring (5351, -0.44 DPS) [quest]; Demon Band (12054, -0.52 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.3 attack_power points (0.71 DPS) | yes | Bounty Hunter's Ring (5351, -0.20 DPS, sim-verified) [quest]; Demon Band (12054, -0.40 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.46 DPS) | yes | Tail Spike (6448, -1.99 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.28 DPS) [crafted]; Goblin Screwdriver (1936, -3.95 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.2 attack_power points (17.35 DPS) | yes | Tail Spike (6448, -0.44 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.9 attack_power points (0.52 DPS) | yes | Fine Longbow (11304, -0.23 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.26 DPS) [crafted]; Light Bow (4576, -0.26 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 240, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (undead, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 76.9. Weights run: 7.7s. Verify run: 11.3s. 423 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.599 ± 0.009, crit=0.617 ± 0.010 per rating point (14 rating = 1%, 8.632 per %), hit=0.906 ± 0.042 per rating point (10 rating = 1%, 9.055 per %), melee_haste=5.114 ± 0.579

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.0 attack_power points (1.49 DPS) | yes | Brawler's Leather Hood (252504, -0.37 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.37 DPS) [crafted]; Humbert's Helm (4724, -0.45 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 12.8 attack_power points (1.19 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.07 DPS) [world]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.82 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.6 attack_power points (2.11 DPS) | yes | Mantle of Thieves (2264, -0.51 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.90 DPS) [crafted]; Bristlebark Amice (14573, -0.93 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.13 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon]; Wildhunter Cloak (16658, -0.39 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.4 attack_power points (2.09 DPS) | yes | Brawler's Leather Tunic (252508, -0.34 DPS) [crafted]; Panther Armor (6670, -0.56 DPS) [quest]; Brawler's Leather Armor (252490, -0.58 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.6 attack_power points (1.27 DPS) | yes | Barbaric Bracers (18948, -0.30 DPS) [crafted]; Cultist's Armguards (270032, -0.34 DPS) [quest]; Insignia Bracers (6410, -0.37 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.6 attack_power points (1.55 DPS) | yes | Heavy Earthen Gloves (7359, -0.06 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.19 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (2.24 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.34 DPS) [crafted]; Prowler's Leather Belt (252459, -0.50 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.43 DPS) | yes | Petrolspill Leggings (9509, -0.34 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.34 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.43 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.8 attack_power points (1.48 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.28 DPS) [world_drop]; Warsong Boots (16977, -0.28 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.6 attack_power points (1.46 DPS) | yes | Band of the Fist (17694, -0.49 DPS) [quest]; Pyrewood Signet Ring (277210, -0.52 DPS) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.8 attack_power points (1.20 DPS) | yes | Band of the Fist (17694, -0.22 DPS) [quest]; Pyrewood Signet Ring (277210, -0.26 DPS) [quest]; Ring of Precision (1491, -0.30 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornspike (6681, -0.18 DPS, sim-verified) [dungeon]; Vendetta (776, -0.65 DPS) [dungeon]; Torturing Poker (7682, -1.07 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (26.84 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.6 attack_power points (0.90 DPS) | yes | Gravespike Repeater (273033, -0.15 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.30 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow

### Band 40 (undead, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 131.2. Weights run: 7.0s. Verify run: 25.2s. 561 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.667 ± 0.011, crit=0.987 ± 0.015 per rating point (14 rating = 1%, 13.811 per %), hit=1.728 ± 0.067 per rating point (10 rating = 1%, 17.275 per %), melee_haste=7.941 ± 0.940

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.3 attack_power points (2.79 DPS) | yes | Hawkeye's Helm (14591, -0.48 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.73 DPS) [world_drop]; Nightscape Headband (8176, -0.89 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Ethereal Talisman (4430, -0.63 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.3 attack_power points (2.22 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.48 DPS) [crafted]; Mantle of Thieves (2264, -0.63 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (131.2 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.32 DPS) [crafted]; Dusky Leather Armor (7374, -0.48 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Imperial Leather Bracers (4061, -0.63 DPS) [dungeon]; Dusky Bracers (7378, -0.63 DPS) [crafted]; Hawkeye's Bracers (14590, -0.71 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dog Training Gloves (7756, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.69 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.85 DPS) | yes | Defiler's Chain Girdle (20152, -0.55 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.57 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.86 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.0 attack_power points (3.33 DPS) | yes | Triprunner Dungarees (9624, -0.19 DPS) [quest]; Ferine Leggings (6690, -0.86 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.11 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.3 attack_power points (2.41 DPS) | yes | Imperial Leather Boots (6431, -0.19 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.32 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.3 attack_power points (2.03 DPS) | yes | Assault Band (13095, -0.13 DPS) [world_drop]; Ring of the Underwood (2951, -0.25 DPS) [world_drop]; Falcon's Hook (7552, -0.32 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.13 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Toxic Revenger (9453, -9.41 DPS, sim-verified) [dungeon]; The Ziggler (8006, -9.87 DPS) [world_drop]; Sacrificial Kris (3187, -10.88 DPS) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (38.59 DPS) | yes | The Ziggler (8006, -2.87 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.22 DPS) [world_drop]; Monolithic Bow (9426, -0.29 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.32 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 215.7. Weights run: 7.7s. Verify run: 26.7s. 706 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.646 ± 0.027, crit=1.293 ± 0.041 per rating point (14 rating = 1%, 18.103 per %), hit=2.117 ± 0.082 per rating point (10 rating = 1%, 21.170 per %), melee_haste=9.073 ± 1.150

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 54.1 attack_power points (6.20 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.34 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.86 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 26.4 attack_power points (3.02 DPS) | yes | Woven Ivy Necklace (19159, -0.64 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.73 DPS) [quest]; Scout's Medallion (19535, -0.76 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 27.1 attack_power points (3.10 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Phytoskin Spaulders (17749, -0.09 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 57.3 attack_power points (6.56 DPS) | yes | Blazewind Breastplate (11193, -1.84 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.91 DPS) [crafted]; Quillward Harness (10583, -2.06 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.7 attack_power points (3.29 DPS) | yes | Skulker's Leather Bracers (252540, -0.81 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.83 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.94 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; First Sergeant's Leather Gauntlets (220857, -2.45 DPS, sim-verified) [vendor] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 38.1 attack_power points (4.36 DPS) | yes | Girdle of Beastial Fury (11686, -0.01 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.35 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.50 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 57.3 attack_power points (6.56 DPS) | yes | Serpentskin Leggings (8262, -2.15 DPS, sim-verified) [world_drop]; Stormshroud Pants (15057, -2.41 DPS) [crafted]; Basilisk Hide Pants (1718, -2.60 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 32.9 attack_power points (3.77 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Prowler's Leather Boots (252468, -0.21 DPS) [crafted]; Whisperwalk Boots (20255, -0.38 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.2 attack_power points (4.72 DPS) | yes | White Bone Band (11862, -1.97 DPS) [quest]; Masons Fraternity Ring (9533, -2.08 DPS) [quest]; Mark of Kern (2262, -2.42 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.8 attack_power points (2.84 DPS) | yes | White Bone Band (11862, -0.09 DPS) [quest]; Masons Fraternity Ring (9533, -0.20 DPS) [quest]; Mark of Kern (2262, -0.55 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.25 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.33 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (215.7 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Lifeforce Dirk (10750, +0.00 DPS) [quest]; Hookfang Shanker (11635, +0.00 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (58.59 DPS) | yes | Lifeforce Dirk (10750, -1.98 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.69 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.69 DPS) [world_drop]; Guttbuster (13139, -0.79 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 349.2. Weights run: 7.5s. Verify run: 90.3s. 1671 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.524 ± 0.025, crit=1.275 ± 0.044 per rating point (14 rating = 1%, 17.845 per %), hit=1.376 ± 0.048 per rating point (10 rating = 1%, 13.762 per %), melee_haste=11.376 ± 1.556

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 58.8 attack_power points (8.22 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Ebon Mask (19984, -1.42 DPS, sim-verified) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 43.8 attack_power points (6.13 DPS) | yes | Imperial Jewel (11933, -1.47 DPS, sim-verified) [dungeon]; Will of the Martyr (17044, -1.94 DPS) [quest]; Skibi's Pendant (13089, -2.66 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -21.33 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Deathguard's Cloak (20068, -0.02 DPS) [rep]; Blackveil Cape (11626, -2.02 DPS) [dungeon]; Cape of the Black Baron (13340, -2.51 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cadaverous Armor (14637, +0.00 DPS) [dungeon]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -18.58 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 40.0 attack_power points (5.59 DPS) | yes | General's Leather Armsplints (16559, -0.14 DPS) [pvp]; Deepfury Bracers (13120, -1.83 DPS) [world_drop]; Shadowcraft Bracers (16710, -2.39 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -3.03 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 59.5 attack_power points (8.32 DPS) | yes | General's Leather Girdle (16557, -1.49 DPS) [pvp]; Might of the Timbermaw (19044, -2.40 DPS) [crafted]; Defiler's Leather Girdle (20190, -2.82 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -14.57 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pads of the Dread Wolf (13210, +0.00 DPS) [dungeon]; Swiftwalker Boots (12553, -0.02 DPS) [dungeon]; Defiler's Leather Boots (20186, -0.26 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.64 DPS) [dungeon]; Blackstone Ring (17713, -1.66 DPS) [dungeon]; Naglering (11669, -9.08 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.19 DPS) [dungeon]; Legionnaire's Band (19510, -0.89 DPS) [rep]; Painweaver Band (13098, -2.34 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -0.51 DPS) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Finkle's Skinner (12709, -10.34 DPS, sim-verified) [dungeon] |
| off_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (349.2 DPS) | yes | Riphook (12653, +0.00 DPS) [dungeon]; Malgen's Long Bow (22318, -0.19 DPS) [dungeon]; Bloodseeker (19107, -0.37 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Frozen Heart of the Mountain; main_hand: Felstriker; off_hand: Glacial Blade

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 817.6. Weights run: 9.0s. Verify run: 94.5s. 1671 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.909 ± 0.037, crit=1.750 ± 0.057 per rating point (14 rating = 1%, 24.502 per %), hit=2.054 ± 0.090 per rating point (10 rating = 1%, 20.544 per %), melee_haste=16.537 ± 2.134

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -2.29 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Jewel (11933, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Will of the Martyr (17044, -0.08 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -45.31 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Deathguard's Cloak (20068, -1.23 DPS) [rep]; Cape of the Black Baron (13340, -3.31 DPS, sim-verified) [dungeon]; Blackveil Cape (11626, -3.75 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -37.93 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 48.4 attack_power points (11.91 DPS) | yes | General's Leather Armsplints (16559, -0.27 DPS) [pvp]; Deepfury Bracers (13120, -3.78 DPS) [world_drop]; Shadowcraft Bracers (16710, -4.86 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Leather Mitts (231555, -0.24 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.74 DPS) [crafted]; Blood Guard's Leather Vices (16499, -3.09 DPS) [pvp] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 81.2 attack_power points (19.99 DPS) | yes | General's Leather Girdle (16557, -5.89 DPS) [pvp]; Might of the Timbermaw (19044, -7.72 DPS) [crafted]; Defiler's Leather Girdle (20190, -11.98 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -49.65 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 46.3 attack_power points (11.39 DPS) | yes | Swiftwalker Boots (12553, -0.43 DPS) [dungeon]; General's Leather Treads (231552, -1.05 DPS) [pvp]; Pads of the Dread Wolf (13210, -1.54 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -4.55 DPS) [dungeon]; Painweaver Band (13098, -4.57 DPS) [dungeon]; Naglering (11669, -17.52 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.66 DPS) [dungeon]; Painweaver Band (13098, -1.67 DPS) [dungeon]; Legionnaire's Band (19510, -3.22 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.37 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.44 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+593.6 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Alcor's Sunrazor (14555, -0.70 DPS) [world_drop] |
| off_hand | Glacial Blade (19099) | Frostwolf Clan [rep] | sim-verified (817.6 DPS) | yes | - |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -1.01 DPS) [quest]; Riphook (12653, -1.05 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -2.58 DPS, sim-verified) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Rune of the Guard Captain; main_hand: Felstriker; off_hand: Glacial Blade; ranged: Blackcrow

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

