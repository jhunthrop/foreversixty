# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-0000000000000000000-035200001000000000)

Set DPS (verified): 45.5. Weights run: 7.4s. Verify run: 12.5s. 241 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.774 ± 0.009, crit=0.451 ± 0.008 per rating point (14 rating = 1%, 6.308 per %), hit=0.568 ± 0.018 per rating point (10 rating = 1%, 5.680 per %), melee_haste=2.468 ± 0.295

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.2 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.6 attack_power points (0.84 DPS) | yes | Erudite's Amulet (277204, -0.57 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.07 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.70 DPS) | yes | Slime-encrusted Pads (6461, -0.72 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (45.5 DPS) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.26 DPS) [crafted]; Bristlebark Cape (14571, -0.28 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.5 attack_power points (1.53 DPS) | yes | Brawler's Leather Armor (252490, -0.18 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.54 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.56 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.9 attack_power points (0.70 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.5 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.27 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.54 DPS) [crafted]; Murloc Scale Belt (5780, -0.84 DPS) [crafted]; Deviate Scale Belt (6468, -3.42 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.53 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (45.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.43 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.1 attack_power points (0.87 DPS) | yes | Demon Band (12054, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.74 DPS) | yes | Demon Band (12054, -0.28 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; The 1 Ring (8350, -0.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.96 DPS) | yes | Tail Spike (6448, -2.03 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.35 DPS) [crafted]; Goblin Screwdriver (1936, -4.05 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.5 attack_power points (17.87 DPS) | yes | Tail Spike (6448, -0.43 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.56 DPS) | yes | Fine Longbow (11304, -0.26 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.28 DPS) [crafted]; Light Bow (4576, -0.28 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 241, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-0000000000000000000-035230011311000100)

Set DPS (verified): 89.9. Weights run: 8.3s. Verify run: 18.4s. 422 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.668 ± 0.008, crit=0.643 ± 0.009 per rating point (14 rating = 1%, 9.000 per %), hit=0.920 ± 0.054 per rating point (10 rating = 1%, 9.195 per %), melee_haste=not significant (3.472 ± 0.902)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.7 attack_power points (1.82 DPS) | yes | Brawler's Leather Hood (252504, -0.36 DPS) [crafted]; Defender's Leather Helm (252455, -0.51 DPS) [crafted]; Humbert's Helm (4724, -0.55 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.3 attack_power points (1.46 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.15 DPS) [world]; Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; River Pride Choker (13087, -1.02 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.3 attack_power points (2.55 DPS) | yes | Mantle of Thieves (2264, -0.72 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -1.09 DPS) [crafted]; Bristlebark Amice (14573, -1.13 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (89.9 DPS) | yes | Tigerstrike Mantle (13108, -0.15 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.44 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.51 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.4 attack_power points (2.55 DPS) | yes | Brawler's Leather Tunic (252508, -0.39 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.55 DPS) [quest]; Brawler's Leather Armor (252490, -0.73 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.0 attack_power points (1.53 DPS) | yes | Barbaric Bracers (18948, -0.36 DPS) [crafted]; Insignia Bracers (6410, -0.44 DPS) [world_drop]; Cultist's Armguards (270032, -0.44 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.0 attack_power points (1.86 DPS) | yes | Toughened Leather Gloves (4253, -0.11 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.11 DPS) [crafted]; Cartilage Shapers (273054, -0.22 DPS) [world] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (2.62 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.33 DPS) [crafted]; Prowler's Leather Belt (252459, -0.54 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.84 DPS) | yes | Petrolspill Leggings (9509, -0.29 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.29 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.44 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.3 attack_power points (1.78 DPS) | yes | Insignia Boots (4055, -0.33 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.33 DPS) [crafted]; Disjointed Shoes (277226, -0.47 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.0 attack_power points (1.75 DPS) | yes | Pyrewood Signet Ring (277210, -0.62 DPS) [quest]; Ring of Precision (1491, -0.66 DPS) [dungeon]; Theramore Signet (271911, -0.77 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.0 attack_power points (1.42 DPS) | yes | Ring of Precision (1491, -0.33 DPS) [dungeon]; Theramore Signet (271911, -0.44 DPS) [rep]; Pyrewood Signet Ring (277210, -0.81 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (89.9 DPS) | yes | Vendetta (776, -0.71 DPS) [dungeon]; Thornspike (6681, -0.77 DPS, sim-verified) [dungeon]; Torturing Poker (7682, -1.25 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (31.37 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.0 attack_power points (1.09 DPS) | yes | Double-barreled Shotgun (2098, -0.22 DPS) [world_drop]; Gravespike Repeater (273033, -0.22 DPS) [world]; Silver Star (3463, -0.36 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 152.2. Weights run: 8.9s. Verify run: 25.7s. 562 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.716 ± 0.010, crit=0.956 ± 0.013 per rating point (14 rating = 1%, 13.384 per %), hit=1.580 ± 0.092 per rating point (10 rating = 1%, 15.795 per %), melee_haste=not significant (5.832 ± 1.542)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.9 attack_power points (3.63 DPS) | yes | Hawkeye's Helm (14591, -0.86 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.92 DPS) [world_drop]; Nightscape Headband (8176, -1.13 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.84 DPS) [world]; Kaleidoscope Chain (13084, -0.97 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (2.90 DPS) | yes | Flintrock Shoulders (7755, -0.21 DPS) [dungeon]; Nightscape Shoulders (8192, -0.61 DPS) [crafted]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.16 DPS) [world_drop]; Imperial Cloak (6432, -0.31 DPS) [dungeon]; Parachute Cloak (10518, -0.31 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (152.2 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.42 DPS) [crafted]; Dusky Leather Armor (7374, -0.63 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.43 DPS) | yes | Hawkeye's Bracers (14590, -0.73 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.76 DPS) [dungeon]; Dusky Bracers (7378, -0.76 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -2.57 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (3.64 DPS) | yes | Highlander's Chain Girdle (20090, -0.73 DPS) [rep]; Ogron's Sash (13117, -0.75 DPS, sim-verified) [world_drop]; Skulker's Leather Belt (252520, -1.04 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 36.0 attack_power points (4.38 DPS) | yes | Triprunner Dungarees (9624, -0.26 DPS) [quest]; Ferine Leggings (6690, -1.22 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.46 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.9 attack_power points (3.14 DPS) | yes | Imperial Leather Boots (6431, -0.24 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.35 DPS) [crafted]; Excelsior Boots (4109, -0.42 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.7 attack_power points (2.64 DPS) | yes | Assault Band (13095, -0.21 DPS) [world_drop]; Ring of the Underwood (2951, -0.31 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.43 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -6.92 DPS, sim-verified) [dungeon]; The Ziggler (8006, -12.61 DPS) [world_drop]; Sacrificial Kris (3187, -13.90 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (50.43 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.24 DPS) [world_drop]; Monolithic Bow (9426, -0.35 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.38 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00530200000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 228.9. Weights run: 9.3s. Verify run: 22.7s. 707 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.578 ± 0.008, crit=1.052 ± 0.013 per rating point (14 rating = 1%, 14.732 per %), hit=1.745 ± 0.098 per rating point (10 rating = 1%, 17.455 per %), melee_haste=not significant (4.880 ± 1.564)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 50.7 attack_power points (7.64 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.38 DPS) [vendor]; Embrace of the Lycan (9479, -1.73 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -3.37 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (3.84 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.91 DPS, sim-verified) [quest]; Sentinel's Medallion (19539, -0.99 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.4 attack_power points (3.97 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.09 DPS) [quest]; Phytoskin Spaulders (17749, -0.17 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.71 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 50.2 attack_power points (7.55 DPS) | yes | Warbear Harness (15064, -1.62 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.64 DPS) [quest]; Quillward Harness (10583, -1.84 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.7 attack_power points (4.17 DPS) | yes | Skulker's Leather Bracers (252540, -1.00 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -1.04 DPS) [world_drop]; Prowler's Leather Bracers (252539, -1.15 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -2.86 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | sim-verified (228.9 DPS) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.27 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 50.2 attack_power points (7.55 DPS) | yes | Gryphon Rider's Leggings (9652, -1.73 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -1.80 DPS) [world_drop]; Basilisk Hide Pants (1718, -2.56 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.6 attack_power points (4.75 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.18 DPS) [crafted]; Whisperwalk Boots (20255, -0.48 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 37.5 attack_power points (5.64 DPS) | yes | Masons Fraternity Ring (9533, -2.31 DPS) [quest]; Mark of Kern (2262, -2.63 DPS) [dungeon]; Assault Band (13095, -2.63 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.2 attack_power points (3.64 DPS) | yes | Masons Fraternity Ring (9533, -0.32 DPS) [quest]; Mark of Kern (2262, -0.63 DPS) [dungeon]; Assault Band (13095, -0.63 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, -2.64 DPS, sim-verified) [world_drop]; Lifeforce Dirk (10750, -6.30 DPS) [quest]; Hookfang Shanker (11635, -6.77 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (76.99 DPS) | yes | Lifeforce Dirk (10750, -1.85 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.77 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.85 DPS, sim-verified) [world_drop]; Guttbuster (13139, -0.97 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 707, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 401.2. Weights run: 8.9s. Verify run: 88.3s. 1667 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.500 ± 0.009, crit=1.116 ± 0.015 per rating point (14 rating = 1%, 15.619 per %), hit=0.987 ± 0.040 per rating point (10 rating = 1%, 9.871 per %), melee_haste=5.870 ± 1.346

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 54.9 attack_power points (11.09 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Ebon Mask (19984, -0.67 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Jewel (11933, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Will of the Martyr (17044, +0.00 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Lizardhide Shoulders (20060, +0.00 DPS) [rep]; Theramore Pauldrons (271955, +0.00 DPS) [rep]; Highlander's Leather Shoulders (20059, -18.61 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.20 DPS) [rep]; Howler's Furs (272414, -0.94 DPS) [vendor]; Blackveil Cape (11626, -3.13 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cadaverous Armor (14637, +0.00 DPS) [dungeon]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -16.16 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 39.5 attack_power points (7.98 DPS) | yes | Marshal's Leather Armsplints (16460, -0.20 DPS) [pvp]; Deepfury Bracers (13120, -2.62 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -3.43 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Knight-Lieutenant's Leather Gauntlets (16396, -0.91 DPS) [pvp]; Marshal's Leather Handgrips (231544, -1.19 DPS) [pvp] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 51.2 attack_power points (10.35 DPS) | yes | Marshal's Leather Cinch (16458, -0.55 DPS) [pvp]; Might of the Timbermaw (19044, -1.87 DPS) [crafted]; Highlander's Leather Girdle (20045, -3.00 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Leggings (15062, +0.00 DPS) [crafted]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -18.65 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Pads of the Dread Wolf (13210, +0.00 DPS) [dungeon]; Highlander's Leather Boots (20052, -0.09 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.59 DPS) [dungeon]; Blackstone Ring (17713, -1.94 DPS) [dungeon]; Naglering (11669, -9.36 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, +0.00 DPS) [dungeon]; Blackstone Ring (17713, -0.20 DPS) [dungeon]; Protector's Band (19514, -0.48 DPS) [rep] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -4.25 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -17.42 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (128.28 DPS) | yes | Electrified Dagger (19100, -6.18 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (401.2 DPS) | yes | Riphook (12653, +0.00 DPS) [dungeon]; Malgen's Long Bow (22318, -0.20 DPS) [dungeon]; Bloodseeker (19107, -0.50 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; main_hand: Felstriker; off_hand: Alcor's Sunrazor

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 867.0. Weights run: 9.8s. Verify run: 107.9s. 1667 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.459 ± 0.020, crit=2.670 ± 0.033 per rating point (14 rating = 1%, 37.376 per %), hit=3.606 ± 0.223 per rating point (10 rating = 1%, 36.058 per %), melee_haste=17.838 ± 2.901

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Emberfury Talisman (12929, -0.02 DPS) [dungeon]; Sentinel's Medallion (19538, -0.13 DPS) [rep] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -46.52 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -3.92 DPS) [rep]; Shadow Prowler's Cloak (22269, -4.91 DPS) [dungeon]; Cape of the Black Baron (13340, -6.91 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -39.11 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 58.8 attack_power points (12.99 DPS) | yes | Marshal's Leather Armsplints (16460, -0.24 DPS) [pvp]; Blackmist Armguards (12966, -3.81 DPS) [dungeon]; Deepfury Bracers (13120, -3.87 DPS) [world_drop] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.90 DPS) [crafted]; Gloves of Holy Might (867, -5.45 DPS, sim-verified) [world_drop] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 123.8 attack_power points (27.34 DPS) | yes | Marshal's Leather Cinch (16458, -12.87 DPS) [pvp]; Highlander's Leather Girdle (20045, -14.34 DPS, sim-verified) [rep]; Might of the Timbermaw (19044, -14.63 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -46.52 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.5 attack_power points (13.80 DPS) | yes | Marshal's Leather Footguards (231546, -1.85 DPS) [pvp]; Swiftfoot Treads (15861, -2.80 DPS) [quest]; Swiftwalker Boots (12553, -5.48 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.93 DPS) [dungeon]; Painweaver Band (13098, -7.52 DPS) [dungeon]; Naglering (11669, -23.43 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.19 DPS) [dungeon]; Painweaver Band (13098, -3.78 DPS) [dungeon]; Magni's Will (12548, -5.86 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -6.37 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -2.75 DPS) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -35.22 DPS, sim-verified) [rep] |
| off_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (867.0 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Ancient Bone Bow (18680, -1.63 DPS) [dungeon]; Bloodseeker (19107, -1.86 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; main_hand: Felstriker; off_hand: Heartseeker

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000000-035200001000000000)

Set DPS (verified): 45.5. Weights run: 7.4s. Verify run: 12.1s. 240 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.774 ± 0.009, crit=0.451 ± 0.008 per rating point (14 rating = 1%, 6.308 per %), hit=0.568 ± 0.018 per rating point (10 rating = 1%, 5.680 per %), melee_haste=2.468 ± 0.295

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.2 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.49 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.6 attack_power points (0.84 DPS) | yes | Erudite's Amulet (277204, -0.54 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.70 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (45.5 DPS) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.26 DPS) [crafted]; Bristlebark Cape (14571, -0.28 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (45.5 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.52 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.9 attack_power points (0.70 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.28 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.6 attack_power points (1.15 DPS) | yes | Brawler's Leather Gloves (252494, -0.28 DPS) [crafted]; Bristlebark Gloves (14572, -0.28 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.31 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.54 DPS) [crafted]; Murloc Scale Belt (5780, -0.84 DPS) [crafted]; Deviate Scale Belt (6468, -3.30 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (45.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.48 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (45.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.57 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.1 attack_power points (0.87 DPS) | yes | Bounty Hunter's Ring (5351, -0.45 DPS) [quest]; Demon Band (12054, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.74 DPS) | yes | Bounty Hunter's Ring (5351, -0.24 DPS, sim-verified) [quest]; Demon Band (12054, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.96 DPS) | yes | Tail Spike (6448, -2.03 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.35 DPS) [crafted]; Goblin Screwdriver (1936, -4.05 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.5 attack_power points (17.87 DPS) | yes | Tail Spike (6448, -0.41 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.56 DPS) | yes | Fine Longbow (11304, -0.26 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.28 DPS) [crafted]; Light Bow (4576, -0.28 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 240, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (undead, 00000000000000000-0000000000000000000-035230011311000100)

Set DPS (verified): 90.0. Weights run: 8.3s. Verify run: 21.3s. 423 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.668 ± 0.008, crit=0.643 ± 0.009 per rating point (14 rating = 1%, 9.000 per %), hit=0.920 ± 0.054 per rating point (10 rating = 1%, 9.195 per %), melee_haste=not significant (3.472 ± 0.902)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.7 attack_power points (1.82 DPS) | yes | Brawler's Leather Hood (252504, -0.42 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.51 DPS) [crafted]; Humbert's Helm (4724, -0.55 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.3 attack_power points (1.46 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.15 DPS) [world]; Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; River Pride Choker (13087, -1.02 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.3 attack_power points (2.55 DPS) | yes | Mantle of Thieves (2264, -0.63 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -1.09 DPS) [crafted]; Bristlebark Amice (14573, -1.13 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.15 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.51 DPS) [dungeon]; Swiftrunner Cape (6745, -0.51 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.4 attack_power points (2.55 DPS) | yes | Brawler's Leather Tunic (252508, -0.44 DPS) [crafted]; Panther Armor (6670, -0.69 DPS) [quest]; Brawler's Leather Armor (252490, -0.73 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.0 attack_power points (1.53 DPS) | yes | Barbaric Bracers (18948, -0.36 DPS) [crafted]; Insignia Bracers (6410, -0.44 DPS) [world_drop]; Cultist's Armguards (270032, -0.44 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.0 attack_power points (1.86 DPS) | yes | Toughened Leather Gloves (4253, -0.11 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.11 DPS) [crafted]; Cartilage Shapers (273054, -0.22 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (2.62 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.33 DPS) [crafted]; Prowler's Leather Belt (252459, -0.54 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.84 DPS) | yes | Petrolspill Leggings (9509, -0.29 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.29 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.44 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.3 attack_power points (1.78 DPS) | yes | Insignia Boots (4055, -0.33 DPS) [world_drop]; Warsong Boots (16977, -0.33 DPS) [quest]; Brawler's Leather Boots (252439, -0.33 DPS) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.0 attack_power points (1.75 DPS) | yes | Band of the Fist (17694, -0.58 DPS) [quest]; Pyrewood Signet Ring (277210, -0.62 DPS) [quest]; Ring of Precision (1491, -0.66 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.0 attack_power points (1.42 DPS) | yes | Band of the Fist (17694, -0.25 DPS) [quest]; Pyrewood Signet Ring (277210, -0.29 DPS) [quest]; Ring of Precision (1491, -0.33 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (90.0 DPS) | yes | Vendetta (776, -0.71 DPS) [dungeon]; Thornspike (6681, -0.93 DPS, sim-verified) [dungeon]; Torturing Poker (7682, -1.25 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (31.37 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.0 attack_power points (1.09 DPS) | yes | Double-barreled Shotgun (2098, -0.22 DPS) [world_drop]; Gravespike Repeater (273033, -0.22 DPS) [world]; Silver Star (3463, -0.36 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow

### Band 40 (undead, 00000000000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 151.9. Weights run: 8.9s. Verify run: 27.4s. 561 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.716 ± 0.010, crit=0.956 ± 0.013 per rating point (14 rating = 1%, 13.384 per %), hit=1.580 ± 0.092 per rating point (10 rating = 1%, 15.795 per %), melee_haste=not significant (5.832 ± 1.542)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.9 attack_power points (3.63 DPS) | yes | Hawkeye's Helm (14591, -0.85 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.92 DPS) [world_drop]; Nightscape Headband (8176, -1.13 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.84 DPS) [world]; Ethereal Talisman (4430, -0.85 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.9 attack_power points (2.90 DPS) | yes | Flintrock Shoulders (7755, -0.21 DPS) [dungeon]; Nightscape Shoulders (8192, -0.61 DPS) [crafted]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.16 DPS) [dungeon]; Parachute Cloak (10518, -0.16 DPS) [crafted]; Tigerstrike Mantle (13108, -0.16 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-verified (151.9 DPS) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.42 DPS) [crafted]; Dusky Leather Armor (7374, -0.63 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.43 DPS) | yes | Imperial Leather Bracers (4061, -0.76 DPS) [dungeon]; Dusky Bracers (7378, -0.76 DPS) [crafted]; Hawkeye's Bracers (14590, -0.83 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -2.19 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (3.64 DPS) | yes | Ogron's Sash (13117, -0.67 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.73 DPS) [rep]; Skulker's Leather Belt (252520, -1.04 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 36.0 attack_power points (4.38 DPS) | yes | Triprunner Dungarees (9624, -0.26 DPS) [quest]; Ferine Leggings (6690, -1.22 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.46 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.9 attack_power points (3.14 DPS) | yes | Imperial Leather Boots (6431, -0.24 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.35 DPS) [crafted]; Excelsior Boots (4109, -0.42 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.7 attack_power points (2.64 DPS) | yes | Assault Band (13095, -0.21 DPS) [world_drop]; Ring of the Underwood (2951, -0.31 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.43 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.19 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Toxic Revenger (9453, -12.36 DPS, sim-verified) [dungeon]; The Ziggler (8006, -12.61 DPS) [world_drop]; Tok'kar's Murloc Shanker (9680, -13.86 DPS) [quest] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (49.30 DPS) | yes | The Ziggler (8006, -1.30 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.24 DPS) [world_drop]; Monolithic Bow (9426, -0.35 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.38 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 00530200000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 231.0. Weights run: 9.3s. Verify run: 24.3s. 706 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.578 ± 0.008, crit=1.052 ± 0.013 per rating point (14 rating = 1%, 14.732 per %), hit=1.745 ± 0.098 per rating point (10 rating = 1%, 17.455 per %), melee_haste=not significant (4.880 ± 1.564)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 50.7 attack_power points (7.64 DPS) | yes | Blood Guard's Leather Headband (220851, -0.38 DPS) [vendor]; Embrace of the Lycan (9479, -1.70 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -3.37 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (3.84 DPS) | yes | Woven Ivy Necklace (19159, -0.80 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.83 DPS) [quest]; Scout's Medallion (19535, -0.99 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.4 attack_power points (3.97 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.09 DPS) [quest]; Phytoskin Spaulders (17749, -0.17 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.71 DPS) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 50.2 attack_power points (7.55 DPS) | yes | Blazewind Breastplate (11193, -1.64 DPS) [quest]; Quillward Harness (10583, -1.84 DPS) [dungeon]; Warbear Harness (15064, -2.55 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.7 attack_power points (4.17 DPS) | yes | Skulker's Leather Bracers (252540, -0.99 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -1.04 DPS) [world_drop]; Prowler's Leather Bracers (252539, -1.15 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; First Sergeant's Leather Gauntlets (220857, -3.97 DPS, sim-verified) [vendor] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.27 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 50.2 attack_power points (7.55 DPS) | yes | Basilisk Hide Pants (1718, -2.56 DPS) [world_drop]; Triprunner Dungarees (9624, -2.83 DPS) [quest]; Serpentskin Leggings (8262, -3.28 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.6 attack_power points (4.75 DPS) | yes | Skulker's Leather Boots (252469, -0.01 DPS) [crafted]; Prowler's Leather Boots (252468, -0.18 DPS) [crafted]; Whisperwalk Boots (20255, -0.48 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 37.5 attack_power points (5.64 DPS) | yes | White Bone Band (11862, -2.03 DPS) [quest]; Masons Fraternity Ring (9533, -2.31 DPS) [quest]; Mark of Kern (2262, -2.63 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.2 attack_power points (3.64 DPS) | yes | White Bone Band (11862, -0.03 DPS) [quest]; Masons Fraternity Ring (9533, -0.32 DPS) [quest]; Mark of Kern (2262, -0.63 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.58 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.87 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (231.0 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Lifeforce Dirk (10750, +0.00 DPS) [quest]; Hookfang Shanker (11635, +0.00 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (76.99 DPS) | yes | Lifeforce Dirk (10750, -2.47 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.77 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.86 DPS, sim-verified) [world_drop]; Guttbuster (13139, -0.97 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 410.0. Weights run: 8.9s. Verify run: 93.6s. 1671 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.957 ± 0.017, crit=2.082 ± 0.030 per rating point (14 rating = 1%, 29.144 per %), hit=2.625 ± 0.140 per rating point (10 rating = 1%, 26.251 per %), melee_haste=not significant (7.464 ± 2.272)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 84.8 attack_power points (13.13 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -4.51 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 55.1 attack_power points (8.54 DPS) | yes | Imperial Jewel (11933, -1.80 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -3.83 DPS) [world_drop]; Will of the Martyr (17044, -3.89 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -19.49 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.76 DPS) [dungeon]; Deathguard's Cloak (20068, -1.62 DPS) [rep]; Blackveil Cape (11626, -3.23 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -17.13 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 48.2 attack_power points (7.46 DPS) | yes | General's Leather Armsplints (16559, -0.15 DPS) [pvp]; Deepfury Bracers (13120, -2.30 DPS) [world_drop]; Blackmist Armguards (12966, -2.62 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.85 DPS, sim-verified) [crafted]; Blood Guard's Leather Vices (16499, -1.90 DPS) [pvp] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 93.6 attack_power points (14.49 DPS) | yes | Defiler's Leather Girdle (20190, -4.42 DPS, sim-verified) [rep]; General's Leather Girdle (16557, -5.92 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -22.22 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 49.0 attack_power points (7.58 DPS) | yes | Swiftwalker Boots (12553, -0.60 DPS) [dungeon]; General's Leather Treads (231552, -0.92 DPS) [pvp]; Swiftfoot Treads (15861, -1.35 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.58 DPS) [dungeon]; Painweaver Band (13098, -3.76 DPS) [dungeon]; Naglering (11669, -8.24 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.15 DPS) [dungeon]; Painweaver Band (13098, -1.32 DPS) [dungeon]; Legionnaire's Band (19510, -3.12 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Molten Heart of the Mountain (249470, -1.00 DPS, sim-verified) [crafted]; Frozen Heart of the Mountain (249469, -1.36 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -16.49 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (98.40 DPS) | yes | Heartseeker (12783, -6.28 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.73 DPS) [world_drop]; Riphook (12653, -1.57 DPS) [dungeon]; Bloodseeker (19107, -1.61 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Rune of the Guard Captain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 877.2. Weights run: 9.8s. Verify run: 108.4s. 1671 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.459 ± 0.020, crit=2.670 ± 0.033 per rating point (14 rating = 1%, 37.376 per %), hit=3.606 ± 0.223 per rating point (10 rating = 1%, 36.058 per %), melee_haste=17.838 ± 2.901

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Emberfury Talisman (12929, -0.02 DPS) [dungeon]; Scout's Medallion (19534, -0.13 DPS) [rep] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -41.33 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.58 DPS) [dungeon]; Deathguard's Cloak (20068, -3.92 DPS) [rep]; Shadow Prowler's Cloak (22269, -4.91 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -38.83 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 58.8 attack_power points (12.99 DPS) | yes | General's Leather Armsplints (16559, -0.24 DPS) [pvp]; Blackmist Armguards (12966, -3.81 DPS) [dungeon]; Deepfury Bracers (13120, -3.87 DPS) [world_drop] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.90 DPS) [crafted]; Gargoyle Slashers (13957, -3.94 DPS) [dungeon] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 123.8 attack_power points (27.34 DPS) | yes | Defiler's Leather Girdle (20190, -9.42 DPS, sim-verified) [rep]; General's Leather Girdle (16557, -12.87 DPS) [pvp]; Might of the Timbermaw (19044, -14.63 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -48.76 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.5 attack_power points (13.80 DPS) | yes | Swiftwalker Boots (12553, -1.42 DPS) [dungeon]; General's Leather Treads (231552, -1.85 DPS) [pvp]; Swiftfoot Treads (15861, -2.80 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.93 DPS) [dungeon]; Painweaver Band (13098, -7.52 DPS) [dungeon]; Naglering (11669, -14.71 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.19 DPS) [dungeon]; Painweaver Band (13098, -3.78 DPS) [dungeon]; Legionnaire's Band (19510, -6.68 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (877.2 DPS) | yes | Frozen Heart of the Mountain (249469, -1.50 DPS) [crafted]; Hand of Justice (11815, -4.25 DPS) [dungeon] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -2.77 DPS, sim-verified) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -33.26 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (140.34 DPS) | yes | Heartseeker (12783, -5.10 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.99 DPS) [world_drop]; Ancient Bone Bow (18680, -3.62 DPS) [dungeon]; Bloodseeker (19107, -3.85 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

