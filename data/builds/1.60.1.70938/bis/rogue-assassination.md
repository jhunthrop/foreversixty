# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00530200100000000-0000000000000000000-000000000000000000)

Set DPS (verified): 45.6. Weights run: 5.9s. Verify run: 24.4s. 275 eligible items had no known source.

Race sweep: night-elf 45.6 ± 0.0, high-order-skyborne 45.0 ± 0.0, human 44.9 ± 0.0, gnome 44.8 ± 0.0, dwarf 44.2 ± 0.0 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.076 ± 0.014, crit=0.631 ± 0.012 per rating point (14 rating = 1%, 8.834 per %), hit=0.792 ± 0.039 per rating point (10 rating = 1%, 7.919 per %), melee_haste=3.453 ± 0.661

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.6 attack_power points (1.21 DPS) | yes | Defender's Leather Hood (252447, -0.52 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.5 attack_power points (0.91 DPS) | yes | Erudite's Amulet (277204, -0.57 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.13 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.4 attack_power points (0.76 DPS) | yes | Slime-encrusted Pads (6461, -0.71 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.15 DPS) [dungeon]; Sentry Cloak (2059, -0.30 DPS) [world_drop]; Bristlebark Cape (14571, -0.30 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.8 attack_power points (1.67 DPS) | yes | Brawler's Leather Armor (252490, -0.24 DPS) [crafted]; Trapper's Leather Armor (252491, -0.61 DPS) [crafted]; Prospector's Chestpiece (14562, -0.61 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.4 attack_power points (0.76 DPS) | yes | Bravo's Armbands (270015, -0.08 DPS) [quest]; Wolf Bracers (4794, -0.15 DPS) [vendor]; Bristlebark Bindings (14569, -0.16 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -3.20 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.31 DPS) | yes | Brawler's Leather Belt (252428, -0.42 DPS) [crafted]; Dark Leather Belt (4249, -0.71 DPS) [crafted]; Deviate Scale Belt (6468, -4.34 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.45 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.34 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.90 DPS) | yes | Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Demon Band (12054, -0.61 DPS) [world_drop]; Ring of the Moon (12052, -0.82 DPS) [world_drop] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 11.5 attack_power points (0.84 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS, sim-verified) [dungeon]; Demon Band (12054, -0.54 DPS) [world_drop]; Ring of the Moon (12052, -0.76 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 228.7 attack_power points (16.67 DPS) | yes | Tail Spike (6448, -1.82 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.02 DPS) [crafted]; Goblin Screwdriver (1936, -3.67 DPS) [dungeon] |
| off_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (16.67 DPS) | yes | Tail Spike (6448, -0.59 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.3 attack_power points (0.61 DPS) | yes | Deadly Blunderbuss (4369, -0.28 DPS, sim-verified) [crafted]; Light Bow (4576, -0.30 DPS) [world_drop]; Owlsight Rifle (15205, -0.30 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Assassin's Blade; off_hand: Blackfang; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (night-elf, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 76.6. Weights run: 5.8s. Verify run: 13.2s. 460 eligible items had no known source.

Race sweep: night-elf 76.6 ± 0.1, high-order-skyborne 75.9 ± 0.1, human 75.9 ± 0.1, gnome 75.7 ± 0.1, dwarf 74.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.606 ± 0.009, crit=0.625 ± 0.010 per rating point (14 rating = 1%, 8.746 per %), hit=1.025 ± 0.041 per rating point (10 rating = 1%, 10.246 per %), melee_haste=3.783 ± 0.591

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.1 attack_power points (1.50 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.38 DPS) [crafted]; Humbert's Helm (4724, -0.45 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 12.8 attack_power points (1.20 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.83 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.7 attack_power points (2.11 DPS) | yes | Mantle of Thieves (2264, -0.62 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.90 DPS) [crafted]; Bristlebark Amice (14573, -0.94 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.13 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.36 DPS) [pvp]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.5 attack_power points (2.10 DPS) | yes | Brawler's Leather Tunic (252508, -0.37 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.45 DPS) [quest]; Brawler's Leather Armor (252490, -0.58 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.6 attack_power points (1.27 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS, sim-verified) [crafted]; Cultist's Armguards (270032, -0.34 DPS) [quest]; Insignia Bracers (6410, -0.37 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.6 attack_power points (1.55 DPS) | yes | Heavy Earthen Gloves (7359, -0.06 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.19 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (2.24 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.33 DPS) [crafted]; Prowler's Leather Belt (252459, -0.50 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.43 DPS) | yes | Troll's Bane Leggings (13114, -0.33 DPS) [world_drop]; Petrolspill Leggings (9509, -0.38 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.42 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.8 attack_power points (1.48 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.28 DPS) [world_drop]; Disjointed Shoes (277226, -0.36 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.6 attack_power points (1.46 DPS) | yes | Pyrewood Signet Ring (277210, -0.48 DPS) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon]; Theramore Signet (271911, -0.62 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.8 attack_power points (1.20 DPS) | yes | Ring of Precision (1491, -0.30 DPS) [dungeon]; Theramore Signet (271911, -0.36 DPS) [rep]; Pyrewood Signet Ring (277210, -0.56 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Thornspike (6681) | Razorfen Kraul: Aggem Thorncurse [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Toxic Revenger (9453, -0.31 DPS, sim-verified) [dungeon]; Vendetta (776, -0.72 DPS) [dungeon]; Torturing Poker (7682, -1.15 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (26.81 DPS) | yes | Toxic Revenger (9453, +0.00 DPS) [dungeon] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.6 attack_power points (0.90 DPS) | yes | Gravespike Repeater (273033, -0.15 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.30 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Thornspike; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 132.3. Weights run: 5.7s. Verify run: 11.5s. 662 eligible items had no known source.

Race sweep: night-elf 132.3 ± 0.1, high-order-skyborne 131.8 ± 0.1, gnome 131.6 ± 0.1, human 131.6 ± 0.1, dwarf 130.1 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.634 ± 0.010, crit=0.949 ± 0.014 per rating point (14 rating = 1%, 13.293 per %), hit=1.574 ± 0.066 per rating point (10 rating = 1%, 15.743 per %), melee_haste=7.952 ± 1.019

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.0 attack_power points (2.80 DPS) | yes | Hawkeye's Helm (14591, -0.68 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.75 DPS) [world_drop]; Nightscape Headband (8176, -0.90 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 18.0 attack_power points (1.74 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.38 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.58 DPS) [world]; Kaleidoscope Chain (13084, -0.72 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.0 attack_power points (2.22 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.48 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Imperial Cloak (6432, -0.26 DPS) [dungeon]; Parachute Cloak (10518, -0.26 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 39.1 attack_power points (3.77 DPS) | yes | Wolffear Harness (13110, -1.08 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.40 DPS) [crafted]; Dusky Leather Armor (7374, -1.56 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.93 DPS) | yes | Hawkeye's Bracers (14590, -0.61 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.67 DPS) [dungeon]; Dusky Bracers (7378, -0.67 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dog Training Gloves (7756, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -4.26 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.90 DPS) | yes | Highlander's Chain Girdle (20090, -0.58 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.61 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.90 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.3 attack_power points (3.32 DPS) | yes | Triprunner Dungarees (9624, -0.18 DPS) [quest]; Ferine Leggings (6690, -0.80 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.11 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.0 attack_power points (2.41 DPS) | yes | Imperial Leather Boots (6431, -0.19 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.32 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.1 attack_power points (2.04 DPS) | yes | Assault Band (13095, -0.10 DPS) [world_drop]; Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.33 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.93 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.16 DPS) [world_drop]; Falcon's Hook (7552, -0.22 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -3.62 DPS, sim-verified) [dungeon]; The Ziggler (8006, -10.03 DPS) [world_drop]; Sacrificial Kris (3187, -11.06 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (40.12 DPS) | yes | Coldrage Dagger (10761, -0.95 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.25 DPS) [world_drop]; Monolithic Bow (9426, -0.30 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 662, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 221.2. Weights run: 6.0s. Verify run: 26.0s. 851 eligible items had no known source.

Race sweep: night-elf 221.2 ± 0.2, human 220.2 ± 0.2, dwarf 219.9 ± 0.1, gnome 219.9 ± 0.2, high-order-skyborne 219.7 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.710 ± 0.029, crit=1.379 ± 0.044 per rating point (14 rating = 1%, 19.313 per %), hit=2.106 ± 0.092 per rating point (10 rating = 1%, 21.056 per %), melee_haste=8.918 ± 1.259

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 55.3 attack_power points (6.32 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -1.14 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.91 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.2 attack_power points (3.11 DPS) | yes | Sentinel's Medallion (19539, -0.84 DPS, sim-verified) [rep]; Zealous Shadowshard Pendant (17772, -1.51 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 27.8 attack_power points (3.18 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.16 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.49 DPS) [dungeon]; Dark Phantom Cape (13122, -0.49 DPS) [world_drop]; Duskbat Drape (19982, -0.69 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | PvP rank 6 · Knight · Alliance [vendor] | 58.4 attack_power points (6.67 DPS) | yes | Blazewind Breastplate (11193, -1.59 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.90 DPS) [crafted]; Quillward Harness (10583, -2.04 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.6 attack_power points (3.39 DPS) | yes | Skulker's Leather Bracers (252540, -0.73 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.86 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.99 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, -4.04 DPS, sim-verified) [vendor] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 39.3 attack_power points (4.49 DPS) | yes | Girdle of Beastial Fury (11686, -0.15 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.39 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.55 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | PvP rank 6 · Knight · Alliance [vendor] | 58.4 attack_power points (6.67 DPS) | yes | Gryphon Rider's Leggings (9652, -1.81 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -2.06 DPS) [world_drop]; Stormshroud Pants (15057, -2.26 DPS) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 34.2 attack_power points (3.91 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Prowler's Leather Boots (252468, -0.27 DPS) [crafted]; Whisperwalk Boots (20255, -0.59 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.1 attack_power points (4.69 DPS) | yes | Masons Fraternity Ring (9533, -1.96 DPS) [quest]; Mark of Kern (2262, -2.41 DPS) [dungeon]; Assault Band (13095, -2.41 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.4 attack_power points (2.90 DPS) | yes | Masons Fraternity Ring (9533, -0.17 DPS) [quest]; Mark of Kern (2262, -0.62 DPS) [dungeon]; Assault Band (13095, -0.62 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -2.29 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -0.83 DPS, sim-verified) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -2.41 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -8.10 DPS) [quest]; Hookfang Shanker (11635, -8.46 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (58.43 DPS) | yes | Lifeforce Dirk (10750, -1.88 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Houndmaster's Rifle (11629, -0.33 DPS, sim-verified) [dungeon]; Stinging Bow (10624, -0.53 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 851, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 347.0. Weights run: 6.0s. Verify run: 76.0s. 1922 eligible items had no known source.

Race sweep: gnome 347.4 ± 0.3, night-elf 347.0 ± 0.3, high-order-skyborne 346.2 ± 0.3, human 346.1 ± 0.3, dwarf 345.7 ± 0.3 (dps); gnome beats the guide race night-elf by 0.3 (0.10%) ± 0.4, inside error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.466 ± 0.027, crit=1.259 ± 0.047 per rating point (14 rating = 1%, 17.632 per %), hit=1.411 ± 0.076 per rating point (10 rating = 1%, 14.112 per %), melee_haste=10.942 ± 1.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Ebon Mask (19984, -0.57 DPS, sim-verified) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 43.6 attack_power points (6.03 DPS) | yes | Will of the Martyr (17044, -1.88 DPS) [quest]; Imperial Jewel (11933, -1.92 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -2.71 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -18.80 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.11 DPS) [rep]; Cape of the Black Baron (13340, -1.59 DPS, sim-verified) [dungeon]; Blackveil Cape (11626, -2.15 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -17.29 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 38.9 attack_power points (5.37 DPS) | yes | Marshal's Leather Armsplints (16460, -0.14 DPS) [pvp]; Deepfury Bracers (13120, -1.78 DPS) [world_drop]; Wristguards of Renown (22204, -2.10 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, -2.97 DPS, sim-verified) [rep] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 59.0 attack_power points (8.16 DPS) | yes | Marshal's Leather Cinch (16458, -1.52 DPS) [pvp]; Might of the Timbermaw (19044, -2.42 DPS) [crafted]; Highlander's Leather Girdle (20045, -4.37 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -18.80 DPS, sim-verified) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (5.53 DPS) | yes | Swiftwalker Boots (12553, -0.58 DPS) [dungeon]; Shadowcraft Boots (16711, -0.69 DPS) [dungeon]; Highlander's Leather Boots (20052, -0.89 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.61 DPS) [dungeon]; Painweaver Band (13098, -2.23 DPS) [dungeon]; Naglering (11669, -8.21 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.07 DPS) [dungeon]; Painweaver Band (13098, -0.69 DPS) [dungeon]; Protector's Band (19514, -0.90 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Heart of Wyrmthalak (22321, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, -1.81 DPS, sim-verified) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Frightalon (14024, -10.43 DPS, sim-verified) [dungeon] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Electrified Dagger (19100, -1.84 DPS, sim-verified) [rep] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.40 DPS) [world_drop]; Bloodseeker (19107, -0.52 DPS) [quest]; Malgen's Long Bow (22318, -0.55 DPS) [dungeon] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Riphook

No-known-source sample (15 of 1922, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00530311251051031-3023030202012200000-000000000000000000)

Set DPS (verified): 957.8. Weights run: 6.6s. Verify run: 74.1s. 1922 eligible items had no known source.

Race sweep: gnome 962.9 ± 0.9, night-elf 957.8 ± 0.9, high-order-skyborne 957.5 ± 0.9, human 956.5 ± 0.9, dwarf 953.9 ± 0.9 (dps); gnome beats the guide race night-elf by 5.1 (0.54%) ± 1.3, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.100 ± 0.002, agility=2.248 ± 0.028, crit=2.435 ± 0.047 per rating point (14 rating = 1%, 34.090 per %), hit=5.346 ± 0.245 per rating point (10 rating = 1%, 53.455 per %), melee_haste=41.514 ± 1.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 141.0 attack_power points (16.06 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -1.99 DPS) [pvp]; Eye of Rend (12587, -3.43 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nacreous Shell Necklace (22403, -0.76 DPS) [dungeon]; Skibi's Pendant (13089, -2.89 DPS) [world_drop]; Star of Mystaria (12103, -5.15 DPS, sim-verified) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -72.37 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -2.41 DPS, sim-verified) [dungeon]; Cloak of Flames (3475, -3.19 DPS) [world_drop]; Stalwart Cloak (272415, -3.19 DPS) [vendor] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Chestpiece (227058, +0.00 DPS) [pvp]; Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Dawn Armor (252483, -65.94 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.13 DPS) [pvp]; Deepfury Bracers (13120, -1.90 DPS) [world_drop]; Blackmist Armguards (12966, -5.17 DPS, sim-verified) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.22 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 154.1 attack_power points (17.55 DPS) | yes | Marshal's Leather Cinch (16458, -10.45 DPS) [pvp]; Might of the Timbermaw (19044, -11.34 DPS) [crafted]; Highlander's Leather Girdle (20045, -15.92 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -72.37 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.2 attack_power points (7.20 DPS) | yes | Swiftwalker Boots (12553, -1.20 DPS) [dungeon]; Marshal's Leather Footguards (231546, -1.57 DPS) [pvp]; Kayser's Boots of Precision (22231, -6.06 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.20 DPS) [dungeon]; Channeler's Ring (272406, -5.48 DPS) [vendor]; Naglering (11669, -14.40 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.31 DPS) [dungeon]; Channeler's Ring (272406, -3.58 DPS) [vendor]; Painweaver Band (13098, -4.42 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -6.98 DPS, sim-verified) [quest] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; The Lobotomizer (19324, -37.53 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Heartseeker (12783, -3.43 DPS, sim-verified) [crafted] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, +0.00 DPS) [dungeon]; Ancient Bone Bow (18680, -0.92 DPS) [dungeon]; Bloodseeker (19107, -0.94 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Hand of Justice; main_hand: Felstriker; off_hand: Alcor's Sunrazor

No-known-source sample (15 of 1922, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00530200100000000-0000000000000000000-000000000000000000)

Set DPS (verified): 45.8. Weights run: 5.9s. Verify run: 24.9s. 273 eligible items had no known source.

Race sweep: undead 45.8 ± 0.0, windshaper-skyborne 45.1 ± 0.0, troll 45.0 ± 0.0, orc 44.5 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.738 ± 0.009, crit=0.442 ± 0.008 per rating point (14 rating = 1%, 6.186 per %), hit=0.628 ± 0.020 per rating point (10 rating = 1%, 6.282 per %), melee_haste=3.278 ± 0.345

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.9 attack_power points (1.09 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.4 attack_power points (0.82 DPS) | yes | Erudite's Amulet (277204, -0.52 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.7 attack_power points (0.68 DPS) | yes | Slime-encrusted Pads (6461, -0.68 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.40 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.7 attack_power points (0.68 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.27 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.4 attack_power points (1.13 DPS) | yes | Bristlebark Gloves (14572, -0.27 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.27 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Murloc Scale Belt (5780, -0.85 DPS) [crafted]; Deviate Scale Belt (6468, -4.17 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.36 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.44 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.0 attack_power points (0.86 DPS) | yes | Bounty Hunter's Ring (5351, -0.45 DPS) [quest]; Demon Band (12054, -0.55 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.5 attack_power points (0.74 DPS) | yes | Bounty Hunter's Ring (5351, -0.27 DPS, sim-verified) [quest]; Demon Band (12054, -0.43 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackfang (2236, -0.38 DPS, sim-verified) [world_drop]; Tail Spike (6448, -1.94 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.26 DPS) [crafted] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.0 attack_power points (0.55 DPS) | yes | Fine Longbow (11304, -0.24 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Assassin's Blade; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 273, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (undead, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 76.9. Weights run: 5.8s. Verify run: 13.5s. 460 eligible items had no known source.

Race sweep: undead 76.9 ± 0.1, troll 76.3 ± 0.1, windshaper-skyborne 76.2 ± 0.1, orc 75.6 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.606 ± 0.009, crit=0.625 ± 0.010 per rating point (14 rating = 1%, 8.746 per %), hit=1.025 ± 0.041 per rating point (10 rating = 1%, 10.246 per %), melee_haste=3.783 ± 0.591

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.1 attack_power points (1.50 DPS) | yes | Brawler's Leather Hood (252504, -0.38 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.38 DPS) [crafted]; Humbert's Helm (4724, -0.45 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 12.8 attack_power points (1.20 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.83 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.7 attack_power points (2.11 DPS) | yes | Mantle of Thieves (2264, -0.52 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.90 DPS) [crafted]; Bristlebark Amice (14573, -0.94 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.13 DPS) [world_drop]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Wildhunter Cloak (16658, -0.40 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.5 attack_power points (2.10 DPS) | yes | Brawler's Leather Tunic (252508, -0.34 DPS) [crafted]; Panther Armor (6670, -0.56 DPS) [quest]; Brawler's Leather Armor (252490, -0.58 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.6 attack_power points (1.27 DPS) | yes | Barbaric Bracers (18948, -0.30 DPS) [crafted]; Cultist's Armguards (270032, -0.34 DPS) [quest]; Insignia Bracers (6410, -0.37 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.6 attack_power points (1.55 DPS) | yes | Heavy Earthen Gloves (7359, -0.06 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.19 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (2.24 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.33 DPS) [crafted]; Prowler's Leather Belt (252459, -0.50 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.43 DPS) | yes | Petrolspill Leggings (9509, -0.33 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.33 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.42 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.8 attack_power points (1.48 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.28 DPS) [world_drop]; Warsong Boots (16977, -0.28 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.6 attack_power points (1.46 DPS) | yes | Pyrewood Signet Ring (277210, -0.48 DPS) [quest]; Band of the Fist (17694, -0.49 DPS) [quest]; Ring of Precision (1491, -0.56 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.8 attack_power points (1.20 DPS) | yes | Band of the Fist (17694, -0.22 DPS) [quest]; Ring of Precision (1491, -0.30 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.53 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornspike (6681, -0.30 DPS, sim-verified) [dungeon]; Vendetta (776, -0.64 DPS) [dungeon]; Torturing Poker (7682, -1.07 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (26.81 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.6 attack_power points (0.90 DPS) | yes | Gravespike Repeater (273033, -0.15 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.30 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (undead, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 133.3. Weights run: 5.7s. Verify run: 11.9s. 660 eligible items had no known source.

Race sweep: undead 133.3 ± 0.1, windshaper-skyborne 131.7 ± 0.1, troll 131.3 ± 0.1, orc 130.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.634 ± 0.010, crit=0.949 ± 0.014 per rating point (14 rating = 1%, 13.293 per %), hit=1.574 ± 0.066 per rating point (10 rating = 1%, 15.743 per %), melee_haste=7.952 ± 1.019

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.0 attack_power points (2.80 DPS) | yes | Hawkeye's Helm (14591, -0.48 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.75 DPS) [world_drop]; Nightscape Headband (8176, -0.90 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 18.0 attack_power points (1.74 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.56 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -0.58 DPS) [world]; Ethereal Talisman (4430, -0.62 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.0 attack_power points (2.22 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.48 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 39.1 attack_power points (3.77 DPS) | yes | Wolffear Harness (13110, -0.94 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.40 DPS) [crafted]; Dusky Leather Armor (7374, -1.56 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.93 DPS) | yes | Imperial Leather Bracers (4061, -0.67 DPS) [dungeon]; Dusky Bracers (7378, -0.67 DPS) [crafted]; Hawkeye's Bracers (14590, -0.71 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dog Training Gloves (7756, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -4.75 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.90 DPS) | yes | Defiler's Chain Girdle (20152, -0.55 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.61 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.90 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.3 attack_power points (3.32 DPS) | yes | Triprunner Dungarees (9624, -0.18 DPS) [quest]; Ferine Leggings (6690, -0.80 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.11 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.0 attack_power points (2.41 DPS) | yes | Imperial Leather Boots (6431, -0.19 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.32 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.1 attack_power points (2.04 DPS) | yes | Assault Band (13095, -0.10 DPS) [world_drop]; Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.33 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.93 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.16 DPS) [world_drop]; Falcon's Hook (7552, -0.22 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornspike (6681, -9.96 DPS, sim-verified) [dungeon]; The Ziggler (8006, -10.03 DPS) [world_drop]; Sacrificial Kris (3187, -11.06 DPS) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (39.22 DPS) | yes | The Ziggler (8006, -3.04 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.25 DPS) [world_drop]; Monolithic Bow (9426, -0.30 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 660, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 224.3. Weights run: 6.0s. Verify run: 22.8s. 849 eligible items had no known source.

Race sweep: undead 224.3 ± 0.2, troll 220.9 ± 0.1, windshaper-skyborne 220.6 ± 0.2, orc 220.1 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.710 ± 0.029, crit=1.379 ± 0.044 per rating point (14 rating = 1%, 19.313 per %), hit=2.106 ± 0.092 per rating point (10 rating = 1%, 21.056 per %), melee_haste=8.918 ± 1.259

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 55.3 attack_power points (6.32 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Embrace of the Lycan (9479, -2.04 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.91 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.2 attack_power points (3.11 DPS) | yes | Woven Ivy Necklace (19159, -0.67 DPS) [quest]; Scout's Medallion (19535, -0.77 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 27.8 attack_power points (3.18 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.08 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.16 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.49 DPS) [dungeon]; Dark Phantom Cape (13122, -0.49 DPS) [world_drop]; Duskbat Drape (19982, -0.69 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 6 · Stone Guard · Horde [vendor] | 58.4 attack_power points (6.67 DPS) | yes | Warbear Harness (15064, -1.90 DPS) [crafted]; Quillward Harness (10583, -2.04 DPS) [dungeon]; Blazewind Breastplate (11193, -2.89 DPS, sim-verified) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.6 attack_power points (3.39 DPS) | yes | Skulker's Leather Bracers (252540, -0.80 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.86 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.99 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; First Sergeant's Leather Gauntlets (220857, -4.66 DPS, sim-verified) [vendor] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 39.3 attack_power points (4.49 DPS) | yes | Girdle of Beastial Fury (11686, -0.15 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.39 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.55 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 6 · Stone Guard · Horde [vendor] | 58.4 attack_power points (6.67 DPS) | yes | Stormshroud Pants (15057, -2.26 DPS) [crafted]; Basilisk Hide Pants (1718, -2.57 DPS) [world_drop]; Serpentskin Leggings (8262, -3.15 DPS, sim-verified) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 34.2 attack_power points (3.91 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Prowler's Leather Boots (252468, -0.27 DPS) [crafted]; Whisperwalk Boots (20255, -0.59 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.1 attack_power points (4.69 DPS) | yes | White Bone Band (11862, -1.95 DPS) [quest]; Masons Fraternity Ring (9533, -1.96 DPS) [quest]; Mark of Kern (2262, -2.41 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 25.4 attack_power points (2.90 DPS) | yes | White Bone Band (11862, -0.16 DPS) [quest]; Masons Fraternity Ring (9533, -0.17 DPS) [quest]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.12 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -3.41 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -8.10 DPS) [quest]; Hookfang Shanker (11635, -8.46 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (58.43 DPS) | yes | Lifeforce Dirk (10750, -3.55 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Houndmaster's Rifle (11629, -0.51 DPS) [dungeon]; Stinging Bow (10624, -0.53 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.53 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 849, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 350.7. Weights run: 6.0s. Verify run: 72.5s. 1925 eligible items had no known source.

Race sweep: undead 350.7 ± 0.3, windshaper-skyborne 346.7 ± 0.3, troll 345.9 ± 0.3, orc 344.5 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.466 ± 0.027, crit=1.259 ± 0.047 per rating point (14 rating = 1%, 17.632 per %), hit=1.411 ± 0.076 per rating point (10 rating = 1%, 14.112 per %), melee_haste=10.942 ± 1.630

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 53.6 attack_power points (7.41 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Eye of Rend (12587, -0.71 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 43.6 attack_power points (6.03 DPS) | yes | Imperial Jewel (11933, -1.68 DPS, sim-verified) [dungeon]; Will of the Martyr (17044, -1.88 DPS) [quest]; Skibi's Pendant (13089, -2.71 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -20.12 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.11 DPS) [rep]; Blackveil Cape (11626, -2.15 DPS) [dungeon]; Cape of the Black Baron (13340, -2.54 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -16.07 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 38.9 attack_power points (5.37 DPS) | yes | General's Leather Armsplints (16559, -0.14 DPS) [pvp]; Deepfury Bracers (13120, -1.78 DPS) [world_drop]; Wristguards of Renown (22204, -2.10 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor]; Devilsaur Gauntlets (15063, -2.84 DPS, sim-verified) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 59.0 attack_power points (8.16 DPS) | yes | General's Leather Girdle (16557, -1.52 DPS) [pvp]; Might of the Timbermaw (19044, -2.42 DPS) [crafted]; Defiler's Leather Girdle (20190, -5.36 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -16.58 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (5.53 DPS) | yes | Swiftwalker Boots (12553, -0.58 DPS) [dungeon]; Shadowcraft Boots (16711, -0.69 DPS) [dungeon]; Defiler's Leather Boots (20186, -0.89 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.61 DPS) [dungeon]; Painweaver Band (13098, -2.23 DPS) [dungeon]; Naglering (11669, -9.19 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.07 DPS) [dungeon]; Painweaver Band (13098, -0.69 DPS) [dungeon]; Legionnaire's Band (19510, -0.90 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackhand's Breadth (13965, -1.16 DPS, sim-verified) [quest]; Smolderweb's Eye (13213, -1.35 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.55 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Frightalon (14024, -12.23 DPS, sim-verified) [dungeon] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Glacial Blade (19099, -2.34 DPS, sim-verified) [rep] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.40 DPS) [world_drop]; Bloodseeker (19107, -0.52 DPS) [quest]; Malgen's Long Bow (22318, -0.55 DPS) [dungeon] |

**New at 60:** neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Hand of Justice; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Riphook

No-known-source sample (15 of 1925, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 00530311251051031-3023030202012200000-000000000000000000)

Set DPS (verified): 959.1. Weights run: 6.6s. Verify run: 76.9s. 1925 eligible items had no known source.

Race sweep: undead 959.1 ± 0.9, windshaper-skyborne 957.5 ± 0.9, orc 956.2 ± 0.9, troll 954.6 ± 0.9 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.100 ± 0.002, agility=2.248 ± 0.028, crit=2.435 ± 0.047 per rating point (14 rating = 1%, 34.090 per %), hit=5.346 ± 0.245 per rating point (10 rating = 1%, 53.455 per %), melee_haste=41.514 ± 1.606

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.76 DPS) [dungeon]; Skibi's Pendant (13089, -2.89 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -3.16 DPS) [dungeon]; Cloak of Flames (3475, -3.19 DPS) [world_drop]; Stalwart Cloak (272415, -3.19 DPS) [vendor] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Chestpiece (227060, +0.00 DPS) [pvp]; Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Dawn Armor (252483, +0.00 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackmist Armguards (12966, +0.00 DPS) [dungeon]; General's Leather Armsplints (16559, -0.13 DPS) [pvp]; Deepfury Bracers (13120, -1.90 DPS) [world_drop] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.22 DPS) [crafted]; Timbermaw Brawlers (19049, -2.07 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 154.1 attack_power points (17.55 DPS) | yes | Defiler's Leather Girdle (20190, -9.80 DPS) [rep]; General's Leather Girdle (16557, -10.45 DPS) [pvp]; Might of the Timbermaw (19044, -11.34 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.2 attack_power points (7.20 DPS) | yes | Kayser's Boots of Precision (22231, -1.11 DPS) [dungeon]; Swiftwalker Boots (12553, -1.20 DPS) [dungeon]; General's Leather Treads (231552, -1.57 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.20 DPS) [dungeon]; Channeler's Ring (272406, -5.48 DPS) [vendor]; Painweaver Band (13098, -6.31 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.31 DPS) [dungeon]; Channeler's Ring (272406, -3.58 DPS) [vendor]; Painweaver Band (13098, -4.42 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Heartseeker (12783, -1.91 DPS) [crafted] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -2.51 DPS) [world_drop]; Ancient Bone Bow (18680, -3.43 DPS) [dungeon]; Bloodseeker (19107, -3.45 DPS) [quest] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1925, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

