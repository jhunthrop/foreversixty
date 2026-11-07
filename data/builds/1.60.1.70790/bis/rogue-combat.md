# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-3113030000000000000-000000000000000000)

Set DPS (verified): 45.7. Weights run: 6.3s. Verify run: 15.0s. 241 eligible items had no known source.

Race sweep: night-elf 46.2 ± 0.0, gnome 45.7 ± 0.0, high-order-skyborne 45.7 ± 0.0, human 45.6 ± 0.0, dwarf 45.2 ± 0.0 (dps); night-elf beats the guide race gnome by 0.5 (1.09%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.767 ± 0.009, crit=0.458 ± 0.008 per rating point (14 rating = 1%, 6.410 per %), hit=0.574 ± 0.013 per rating point (10 rating = 1%, 5.738 per %), melee_haste=3.295 ± 0.296

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.1 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.45 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.6 attack_power points (0.80 DPS) | yes | Erudite's Amulet (277204, -0.52 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.03 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.8 attack_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, -0.65 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.4 attack_power points (1.47 DPS) | yes | Brawler's Leather Armor (252490, -0.16 DPS) [crafted]; Prospector's Chestpiece (14562, -0.52 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.53 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.8 attack_power points (0.67 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.59 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.36 DPS) | yes | Brawler's Leather Belt (252428, -0.52 DPS) [crafted]; Murloc Scale Belt (5780, -0.81 DPS) [crafted]; Deviate Scale Belt (6468, -3.69 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.85 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.75 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.1 attack_power points (0.83 DPS) | yes | Demon Band (12054, -0.53 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; The 1 Ring (8350, -0.63 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon]; The 1 Ring (8350, -0.50 DPS) [world]; Demon Band (12054, -0.55 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.25 DPS) | yes | Tail Spike (6448, -1.96 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.22 DPS) [crafted]; Goblin Screwdriver (1936, -3.89 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.4 attack_power points (17.16 DPS) | yes | Tail Spike (6448, -0.51 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.53 DPS) | yes | Fine Longbow (11304, -0.22 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 241, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (gnome, 00000000000000000-3153030000012300000-000000000000000000)

Set DPS (verified): 76.1. Weights run: 6.4s. Verify run: 18.5s. 422 eligible items had no known source.

Race sweep: night-elf 76.6 ± 0.1, gnome 76.1 ± 0.1, human 75.9 ± 0.1, high-order-skyborne 75.8 ± 0.1, dwarf 75.4 ± 0.1 (dps); night-elf beats the guide race gnome by 0.5 (0.63%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.644 ± 0.010, crit=0.617 ± 0.012 per rating point (14 rating = 1%, 8.645 per %), hit=0.918 ± 0.040 per rating point (10 rating = 1%, 9.176 per %), melee_haste=5.411 ± 0.528

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.4 attack_power points (1.43 DPS) | yes | Brawler's Leather Hood (252504, -0.30 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.38 DPS) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.1 attack_power points (1.14 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.79 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.1 attack_power points (2.00 DPS) | yes | Mantle of Thieves (2264, -0.58 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.86 DPS) [crafted]; Bristlebark Amice (14573, -0.89 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.12 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.34 DPS) [pvp]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.0 attack_power points (2.00 DPS) | yes | Brawler's Leather Tunic (252508, -0.36 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.43 DPS) [quest]; Brawler's Leather Armor (252490, -0.56 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (1.20 DPS) | yes | Barbaric Bracers (18948, -0.30 DPS, sim-verified) [crafted]; Cultist's Armguards (270032, -0.33 DPS) [quest]; Insignia Bracers (6410, -0.35 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.9 attack_power points (1.46 DPS) | yes | Heavy Earthen Gloves (7359, -0.07 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.17 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (2.08 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.26 DPS) | yes | Troll's Bane Leggings (13114, -0.26 DPS) [world_drop]; Petrolspill Leggings (9509, -0.34 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.36 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.1 attack_power points (1.40 DPS) | yes | Brawler's Leather Boots (252439, -0.25 DPS) [crafted]; Insignia Boots (4055, -0.26 DPS) [world_drop]; Disjointed Shoes (277226, -0.36 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.9 attack_power points (1.38 DPS) | yes | Pyrewood Signet Ring (277210, -0.49 DPS) [quest]; Ring of Precision (1491, -0.52 DPS) [dungeon]; Theramore Signet (271911, -0.60 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.9 attack_power points (1.12 DPS) | yes | Pyrewood Signet Ring (277210, -0.23 DPS) [quest]; Ring of Precision (1491, -0.27 DPS) [dungeon]; Theramore Signet (271911, -0.34 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Vendetta (776, -0.58 DPS) [dungeon]; Thornspike (6681, -0.98 DPS, sim-verified) [dungeon]; Torturing Poker (7682, -1.00 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (24.93 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.9 attack_power points (0.86 DPS) | yes | Gravespike Repeater (273033, -0.16 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.29 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (gnome, 00000000000000000-3153030000012321511-000000000000000000)

Set DPS (verified): 137.3. Weights run: 6.9s. Verify run: 16.5s. 562 eligible items had no known source.

Race sweep: night-elf 137.9 ± 0.1, gnome 137.3 ± 0.1, high-order-skyborne 136.7 ± 0.1, human 136.6 ± 0.1, dwarf 136.3 ± 0.1 (dps); night-elf beats the guide race gnome by 0.6 (0.41%) ± 0.2, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.695 ± 0.012, crit=0.920 ± 0.017 per rating point (14 rating = 1%, 12.877 per %), hit=1.195 ± 0.057 per rating point (10 rating = 1%, 11.953 per %), melee_haste=6.368 ± 0.818

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.6 attack_power points (3.06 DPS) | yes | Hawkeye's Helm (14591, -0.73 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.79 DPS) [world_drop]; Nightscape Headband (8176, -0.96 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.69 DPS) [world]; Kaleidoscope Chain (13084, -0.81 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.6 attack_power points (2.44 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Nightscape Shoulders (8192, -0.52 DPS) [crafted]; Mantle of Thieves (2264, -0.69 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Imperial Cloak (6432, -0.27 DPS) [dungeon]; Parachute Cloak (10518, -0.27 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.35 DPS) [crafted]; Dusky Leather Armor (7374, -0.52 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.06 DPS) | yes | Hawkeye's Bracers (14590, -0.61 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.67 DPS) [dungeon]; Dusky Bracers (7378, -0.67 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.53 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (3.10 DPS) | yes | Ogron's Sash (13117, -0.60 DPS, sim-verified) [world_drop]; Highlander's Chain Girdle (20090, -0.62 DPS) [rep]; Skulker's Leather Belt (252520, -0.90 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.6 attack_power points (3.67 DPS) | yes | Triprunner Dungarees (9624, -0.22 DPS) [quest]; Ferine Leggings (6690, -0.99 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.22 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.6 attack_power points (2.65 DPS) | yes | Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.29 DPS) [crafted]; Excelsior Boots (4109, -0.35 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.6 attack_power points (2.23 DPS) | yes | Assault Band (13095, -0.16 DPS) [world_drop]; Ring of the Underwood (2951, -0.27 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.06 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -6.12 DPS, sim-verified) [dungeon]; The Ziggler (8006, -10.72 DPS) [world_drop]; Sacrificial Kris (3187, -11.82 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (42.88 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.22 DPS) [world_drop]; Monolithic Bow (9426, -0.30 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (gnome, 32200000000000000-3153030000012321541-000000000000000000)

Set DPS (verified): 207.7. Weights run: 7.2s. Verify run: 25.3s. 707 eligible items had no known source.

Race sweep: night-elf 208.4 ± 0.2, gnome 207.7 ± 0.2, human 206.9 ± 0.2, high-order-skyborne 206.9 ± 0.2, dwarf 206.2 ± 0.1 (dps); night-elf beats the guide race gnome by 0.7 (0.33%) ± 0.2, beyond error.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.594 ± 0.013, crit=1.072 ± 0.020 per rating point (14 rating = 1%, 15.008 per %), hit=1.436 ± 0.058 per rating point (10 rating = 1%, 14.360 per %), melee_haste=6.841 ± 0.837

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.0 attack_power points (6.11 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.68 DPS) [vendor]; Embrace of the Lycan (9479, -1.41 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.69 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (3.08 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.72 DPS, sim-verified) [quest]; Sentinel's Medallion (19539, -0.79 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (3.18 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Phytoskin Spaulders (17749, -0.12 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 47.4 attack_power points (5.68 DPS) | yes | Warbear Harness (15064, -0.87 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -0.92 DPS) [quest]; Quillward Harness (10583, -1.09 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.9 attack_power points (3.34 DPS) | yes | Skulker's Leather Bracers (252540, -0.80 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.84 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.93 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -1.86 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.23 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 47.4 attack_power points (5.68 DPS) | yes | Gryphon Rider's Leggings (9652, -1.01 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -1.06 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.67 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.9 attack_power points (3.82 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Prowler's Leather Boots (252468, -0.16 DPS) [crafted]; Whisperwalk Boots (20255, -0.38 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.4 attack_power points (4.12 DPS) | yes | Masons Fraternity Ring (9533, -1.44 DPS) [quest]; Mark of Kern (2262, -1.72 DPS) [dungeon]; Assault Band (13095, -1.72 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.3 attack_power points (2.92 DPS) | yes | Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Mark of Kern (2262, -0.52 DPS) [dungeon]; Assault Band (13095, -0.52 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Lifeforce Dirk (10750, +0.00 DPS) [quest]; Hookfang Shanker (11635, +0.00 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (61.31 DPS) | yes | Lifeforce Dirk (10750, +0.00 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.64 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.64 DPS) [world_drop]; Guttbuster (13139, -0.79 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 707, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (gnome, 32530100000000000-3153030000012321541-030000000000000000)

Set DPS (verified): 310.2. Weights run: 6.6s. Verify run: 74.9s. 1667 eligible items had no known source.

Race sweep: gnome 310.2 ± 0.3, night-elf 310.1 ± 0.3, high-order-skyborne 308.7 ± 0.3, human 308.4 ± 0.3, dwarf 308.0 ± 0.3 (dps); the guide race gnome is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.470 ± 0.012, crit=1.091 ± 0.021 per rating point (14 rating = 1%, 15.271 per %), hit=1.282 ± 0.051 per rating point (10 rating = 1%, 12.818 per %), melee_haste=9.814 ± 1.077

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 53.8 attack_power points (7.86 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Ebon Mask (19984, -0.37 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Jewel (11933, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Will of the Martyr (17044, +0.00 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, +0.00 DPS) [rep]; Highlander's Leather Shoulders (20059, -15.26 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.10 DPS) [rep]; Howler's Furs (272414, -0.18 DPS) [vendor]; Blackveil Cape (11626, -2.26 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cadaverous Armor (14637, +0.00 DPS) [dungeon]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -14.21 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 38.9 attack_power points (5.69 DPS) | yes | Marshal's Leather Armsplints (16460, -0.15 DPS) [pvp]; Deepfury Bracers (13120, -1.88 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -2.47 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, -4.04 DPS, sim-verified) [rep] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 56.5 attack_power points (8.26 DPS) | yes | Marshal's Leather Cinch (16458, -1.24 DPS) [pvp]; Might of the Timbermaw (19044, -2.18 DPS) [crafted]; Highlander's Leather Girdle (20045, -3.65 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.03 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Pads of the Dread Wolf (13210, +0.00 DPS) [dungeon]; Highlander's Leather Boots (20052, -0.16 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.35 DPS) [dungeon]; Painweaver Band (13098, -1.58 DPS) [dungeon]; Naglering (11669, -7.89 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.08 DPS) [dungeon]; Painweaver Band (13098, -0.31 DPS) [dungeon]; Protector's Band (19514, -0.76 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.24 DPS) [crafted]; Heart of Wyrmthalak (22321, -3.07 DPS, sim-verified) [dungeon] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.80 DPS, sim-verified) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -12.49 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (92.82 DPS) | yes | Electrified Dagger (19100, -7.79 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Riphook (12653, +0.00 DPS) [dungeon]; Malgen's Long Bow (22318, -0.08 DPS) [dungeon]; Bloodseeker (19107, -0.33 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Molten Heart of the Mountain; main_hand: Felstriker; off_hand: Alcor's Sunrazor

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (gnome, 00520300200000000-3043030002012321451-005000000000000000)

Set DPS (verified): 889.5. Weights run: 7.4s. Verify run: 81.0s. 1667 eligible items had no known source.

Race sweep: gnome 889.5 ± 0.7, high-order-skyborne 885.3 ± 0.7, night-elf 884.0 ± 0.7, human 883.6 ± 0.7, dwarf 880.7 ± 0.7 (dps); the guide race gnome is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.474 ± 0.026, crit=2.639 ± 0.042 per rating point (14 rating = 1%, 36.949 per %), hit=3.519 ± 0.216 per rating point (10 rating = 1%, 35.192 per %), melee_haste=16.489 ± 2.578

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 107.4 attack_power points (24.02 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -0.01 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Sentinel's Medallion (19538, -0.12 DPS) [rep]; Emberfury Talisman (12929, -0.16 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -50.35 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -3.76 DPS) [rep]; Shadow Prowler's Cloak (22269, -4.73 DPS) [dungeon]; Cape of the Black Baron (13340, -4.92 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -43.05 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.1 attack_power points (13.22 DPS) | yes | Marshal's Leather Armsplints (16460, -0.25 DPS) [pvp]; Deepfury Bracers (13120, -3.94 DPS) [world_drop]; Blackmist Armguards (12966, -4.12 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.12 DPS) [crafted]; Gloves of Holy Might (867, -3.33 DPS, sim-verified) [world_drop] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 122.3 attack_power points (27.37 DPS) | yes | Marshal's Leather Cinch (16458, -12.67 DPS) [pvp]; Highlander's Leather Girdle (20045, -14.37 DPS, sim-verified) [rep]; Might of the Timbermaw (19044, -14.45 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -50.35 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.5 attack_power points (13.98 DPS) | yes | Swiftwalker Boots (12553, -1.38 DPS) [dungeon]; Marshal's Leather Footguards (231546, -1.81 DPS) [pvp]; Swiftfoot Treads (15861, -2.79 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.92 DPS) [dungeon]; Painweaver Band (13098, -7.43 DPS) [dungeon]; Naglering (11669, -18.66 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.27 DPS) [dungeon]; Painweaver Band (13098, -3.78 DPS) [dungeon]; Magni's Will (12548, -5.88 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+8.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -9.29 DPS, sim-verified) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -26.60 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (142.15 DPS) | yes | Heartseeker (12783, -14.31 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -3.20 DPS, sim-verified) [world_drop]; Ancient Bone Bow (18680, -3.45 DPS) [dungeon]; Bloodseeker (19107, -3.69 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Hand of Justice; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00000000000000000-3113030000000000000-000000000000000000)

Set DPS (verified): 46.3. Weights run: 6.3s. Verify run: 14.0s. 240 eligible items had no known source.

Race sweep: undead 46.3 ± 0.0, troll 45.6 ± 0.0, windshaper-skyborne 45.5 ± 0.0, orc 45.1 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.767 ± 0.009, crit=0.458 ± 0.008 per rating point (14 rating = 1%, 6.410 per %), hit=0.574 ± 0.013 per rating point (10 rating = 1%, 5.738 per %), melee_haste=3.295 ± 0.296

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.1 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.49 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.6 attack_power points (0.80 DPS) | yes | Erudite's Amulet (277204, -0.49 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.8 attack_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.54 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.8 attack_power points (0.67 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.27 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.6 attack_power points (1.10 DPS) | yes | Bristlebark Gloves (14572, -0.25 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.27 DPS) [crafted]; Serpent Gloves (5970, -0.30 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.36 DPS) | yes | Brawler's Leather Belt (252428, -0.52 DPS) [crafted]; Murloc Scale Belt (5780, -0.81 DPS) [crafted]; Deviate Scale Belt (6468, -3.27 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.48 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.58 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.1 attack_power points (0.83 DPS) | yes | Bounty Hunter's Ring (5351, -0.44 DPS) [quest]; Demon Band (12054, -0.53 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.71 DPS) | yes | Demon Band (12054, -0.40 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon]; Bounty Hunter's Ring (5351, -0.53 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.25 DPS) | yes | Tail Spike (6448, -1.96 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.22 DPS) [crafted]; Goblin Screwdriver (1936, -3.89 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.4 attack_power points (17.16 DPS) | yes | Tail Spike (6448, -0.43 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.53 DPS) | yes | Fine Longbow (11304, -0.23 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 240, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (undead, 00000000000000000-3153030000012300000-000000000000000000)

Set DPS (verified): 76.8. Weights run: 6.4s. Verify run: 19.3s. 423 eligible items had no known source.

Race sweep: undead 76.8 ± 0.1, troll 75.9 ± 0.1, windshaper-skyborne 75.8 ± 0.1, orc 75.5 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.644 ± 0.010, crit=0.617 ± 0.012 per rating point (14 rating = 1%, 8.645 per %), hit=0.918 ± 0.040 per rating point (10 rating = 1%, 9.176 per %), melee_haste=5.411 ± 0.528

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.4 attack_power points (1.43 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS) [crafted]; Defender's Leather Helm (252455, -0.38 DPS) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.1 attack_power points (1.14 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop]; River Pride Choker (13087, -0.79 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.1 attack_power points (2.00 DPS) | yes | Mantle of Thieves (2264, -0.58 DPS) [dungeon]; Barbaric Shoulders (5964, -0.86 DPS) [crafted]; Bristlebark Amice (14573, -0.89 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.12 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon]; Wildhunter Cloak (16658, -0.39 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.0 attack_power points (2.00 DPS) | yes | Brawler's Leather Tunic (252508, -0.33 DPS) [crafted]; Panther Armor (6670, -0.54 DPS) [quest]; Brawler's Leather Armor (252490, -0.56 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (1.20 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS) [crafted]; Cultist's Armguards (270032, -0.33 DPS) [quest]; Insignia Bracers (6410, -0.35 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.9 attack_power points (1.46 DPS) | yes | Heavy Earthen Gloves (7359, -0.07 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.17 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (2.08 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.26 DPS) | yes | Petrolspill Leggings (9509, -0.26 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.26 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.36 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.1 attack_power points (1.40 DPS) | yes | Brawler's Leather Boots (252439, -0.25 DPS) [crafted]; Insignia Boots (4055, -0.26 DPS) [world_drop]; Warsong Boots (16977, -0.26 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.9 attack_power points (1.38 DPS) | yes | Band of the Fist (17694, -0.46 DPS) [quest]; Pyrewood Signet Ring (277210, -0.49 DPS) [quest]; Ring of Precision (1491, -0.52 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.9 attack_power points (1.12 DPS) | yes | Band of the Fist (17694, -0.20 DPS) [quest]; Pyrewood Signet Ring (277210, -0.23 DPS) [quest]; Ring of Precision (1491, -0.27 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornspike (6681, +0.00 DPS) [dungeon]; Vendetta (776, -0.58 DPS) [dungeon]; Torturing Poker (7682, -1.00 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (24.93 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.9 attack_power points (0.86 DPS) | yes | Gravespike Repeater (273033, -0.16 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.29 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow

### Band 40 (undead, 00000000000000000-3153030000012321511-000000000000000000)

Set DPS (verified): 138.3. Weights run: 6.9s. Verify run: 16.0s. 561 eligible items had no known source.

Race sweep: undead 138.3 ± 0.1, windshaper-skyborne 136.4 ± 0.1, troll 136.4 ± 0.1, orc 135.7 ± 0.1 (dps); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.695 ± 0.012, crit=0.920 ± 0.017 per rating point (14 rating = 1%, 12.877 per %), hit=1.195 ± 0.057 per rating point (10 rating = 1%, 11.953 per %), melee_haste=6.368 ± 0.818

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.6 attack_power points (3.06 DPS) | yes | Hawkeye's Helm (14591, -0.53 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.79 DPS) [world_drop]; Nightscape Headband (8176, -0.96 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.69 DPS) [world]; Ethereal Talisman (4430, -0.71 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.6 attack_power points (2.44 DPS) | yes | Flintrock Shoulders (7755, -0.17 DPS) [dungeon]; Nightscape Shoulders (8192, -0.52 DPS) [crafted]; Mantle of Thieves (2264, -0.69 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.35 DPS) [crafted]; Dusky Leather Armor (7374, -0.52 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.06 DPS) | yes | Imperial Leather Bracers (4061, -0.67 DPS) [dungeon]; Dusky Bracers (7378, -0.67 DPS) [crafted]; Hawkeye's Bracers (14590, -0.77 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.70 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (3.10 DPS) | yes | Ogron's Sash (13117, -0.55 DPS, sim-verified) [world_drop]; Defiler's Chain Girdle (20152, -0.62 DPS) [rep]; Skulker's Leather Belt (252520, -0.90 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.6 attack_power points (3.67 DPS) | yes | Triprunner Dungarees (9624, -0.22 DPS) [quest]; Ferine Leggings (6690, -0.99 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.22 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.6 attack_power points (2.65 DPS) | yes | Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.29 DPS) [crafted]; Excelsior Boots (4109, -0.35 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.6 attack_power points (2.23 DPS) | yes | Assault Band (13095, -0.16 DPS) [world_drop]; Ring of the Underwood (2951, -0.27 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.06 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.18 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Ziggler (8006, -10.72 DPS) [world_drop]; Tok'kar's Murloc Shanker (9680, -11.80 DPS) [quest]; Thornspike (6681, -13.09 DPS, sim-verified) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (41.92 DPS) | yes | The Ziggler (8006, -2.37 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.22 DPS) [world_drop]; Monolithic Bow (9426, -0.30 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 32200000000000000-3153030000012321541-000000000000000000)

Set DPS (verified): 211.4. Weights run: 7.2s. Verify run: 26.3s. 706 eligible items had no known source.

Race sweep: undead 211.4 ± 0.2, troll 208.2 ± 0.2, windshaper-skyborne 208.0 ± 0.2, orc 207.4 ± 0.2 (dps); the guide race undead is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.594 ± 0.013, crit=1.072 ± 0.020 per rating point (14 rating = 1%, 15.008 per %), hit=1.436 ± 0.058 per rating point (10 rating = 1%, 14.360 per %), melee_haste=6.841 ± 0.837

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.0 attack_power points (6.11 DPS) | yes | Blood Guard's Leather Headband (220851, -0.68 DPS) [vendor]; Embrace of the Lycan (9479, -1.26 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.69 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (3.08 DPS) | yes | Woven Ivy Necklace (19159, -0.64 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.69 DPS) [quest]; Scout's Medallion (19535, -0.79 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (3.18 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Phytoskin Spaulders (17749, -0.12 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.57 DPS) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 47.4 attack_power points (5.68 DPS) | yes | Blazewind Breastplate (11193, -0.92 DPS) [quest]; Warbear Harness (15064, -1.03 DPS, sim-verified) [crafted]; Quillward Harness (10583, -1.09 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.9 attack_power points (3.34 DPS) | yes | Skulker's Leather Bracers (252540, -0.66 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.84 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.93 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; First Sergeant's Leather Gauntlets (220857, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -2.77 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.23 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 47.4 attack_power points (5.68 DPS) | yes | Basilisk Hide Pants (1718, -1.67 DPS) [world_drop]; Serpentskin Leggings (8262, -1.73 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -1.88 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.9 attack_power points (3.82 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Prowler's Leather Boots (252468, -0.16 DPS) [crafted]; Whisperwalk Boots (20255, -0.38 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.4 attack_power points (4.12 DPS) | yes | White Bone Band (11862, -1.24 DPS) [quest]; Masons Fraternity Ring (9533, -1.44 DPS) [quest]; Mark of Kern (2262, -1.72 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.3 attack_power points (2.92 DPS) | yes | White Bone Band (11862, -0.04 DPS) [quest]; Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Mark of Kern (2262, -0.52 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.33 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.65 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Lifeforce Dirk (10750, +0.00 DPS) [quest]; Hookfang Shanker (11635, +0.00 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (61.31 DPS) | yes | Lifeforce Dirk (10750, -1.61 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.64 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.64 DPS) [world_drop]; Guttbuster (13139, -0.79 DPS) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 32530100000000000-3153030000012321541-030000000000000000)

Set DPS (verified): 313.3. Weights run: 6.6s. Verify run: 78.4s. 1671 eligible items had no known source.

Race sweep: undead 313.3 ± 0.3, windshaper-skyborne 308.6 ± 0.3, troll 308.3 ± 0.3, orc 306.8 ± 0.3 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.979 ± 0.021, crit=2.152 ± 0.036 per rating point (14 rating = 1%, 30.128 per %), hit=2.739 ± 0.127 per rating point (10 rating = 1%, 27.387 per %), melee_haste=11.879 ± 1.632

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 86.8 attack_power points (10.39 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -3.01 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Jewel (11933, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Emberfury Talisman (12929, -0.07 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -14.92 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -0.68 DPS) [dungeon]; Deathguard's Cloak (20068, -1.38 DPS) [rep]; Blackveil Cape (11626, -2.60 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -12.22 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 48.6 attack_power points (5.82 DPS) | yes | General's Leather Armsplints (16559, -0.12 DPS) [pvp]; Deepfury Bracers (13120, -1.79 DPS) [world_drop]; Blackmist Armguards (12966, -1.94 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 96.3 attack_power points (11.54 DPS) | yes | Defiler's Leather Girdle (20190, -3.62 DPS, sim-verified) [rep]; General's Leather Girdle (16557, -4.87 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -15.66 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 49.8 attack_power points (5.96 DPS) | yes | Swiftwalker Boots (12553, -0.50 DPS) [dungeon]; General's Leather Treads (231552, -0.75 DPS) [pvp]; Swiftfoot Treads (15861, -1.10 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -2.89 DPS) [dungeon]; Painweaver Band (13098, -3.04 DPS) [dungeon]; Naglering (11669, -6.61 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.92 DPS) [dungeon]; Painweaver Band (13098, -1.07 DPS) [dungeon]; Legionnaire's Band (19510, -2.55 DPS) [rep] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.02 DPS) [crafted]; Hand of Justice (11815, -1.58 DPS) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -10.91 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (76.08 DPS) | yes | Heartseeker (12783, -8.36 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.67 DPS) [world_drop]; Riphook (12653, -1.36 DPS) [dungeon]; Bloodseeker (19107, -1.37 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 00520300200000000-3043030002012321451-005000000000000000)

Set DPS (verified): 889.8. Weights run: 7.4s. Verify run: 85.4s. 1671 eligible items had no known source.

Race sweep: undead 889.8 ± 0.7, windshaper-skyborne 886.3 ± 0.7, orc 885.8 ± 0.7, troll 885.6 ± 0.7 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.474 ± 0.026, crit=2.639 ± 0.042 per rating point (14 rating = 1%, 36.949 per %), hit=3.519 ± 0.216 per rating point (10 rating = 1%, 35.192 per %), melee_haste=16.489 ± 2.578

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 107.4 attack_power points (24.02 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -0.01 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Scout's Medallion (19534, -0.12 DPS) [rep]; Emberfury Talisman (12929, -0.16 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -41.02 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.36 DPS) [dungeon]; Deathguard's Cloak (20068, -3.76 DPS) [rep]; Shadow Prowler's Cloak (22269, -4.73 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -31.75 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.1 attack_power points (13.22 DPS) | yes | General's Leather Armsplints (16559, -0.25 DPS) [pvp]; Deepfury Bracers (13120, -3.94 DPS) [world_drop]; Blackmist Armguards (12966, -4.12 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.12 DPS) [crafted]; Gargoyle Slashers (13957, -4.16 DPS) [dungeon] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 122.3 attack_power points (27.37 DPS) | yes | General's Leather Girdle (16557, -12.67 DPS) [pvp]; Defiler's Leather Girdle (20190, -12.78 DPS, sim-verified) [rep]; Might of the Timbermaw (19044, -14.45 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Outrider's Leather Pants (22740, +0.00 DPS) [rep]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -41.64 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.5 attack_power points (13.98 DPS) | yes | Swiftwalker Boots (12553, -1.38 DPS) [dungeon]; General's Leather Treads (231552, -1.81 DPS) [pvp]; Swiftfoot Treads (15861, -2.79 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.92 DPS) [dungeon]; Painweaver Band (13098, -7.43 DPS) [dungeon]; Naglering (11669, -14.92 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.27 DPS) [dungeon]; Painweaver Band (13098, -3.78 DPS) [dungeon]; Legionnaire's Band (19510, -6.58 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+10.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted]; Frozen Heart of the Mountain (249469, -1.56 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -25.10 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (142.15 DPS) | yes | Heartseeker (12783, -11.50 DPS, sim-verified) [crafted] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ancient Bone Bow (18680, -1.66 DPS) [dungeon]; Blackcrow (12651, -1.69 DPS, sim-verified) [dungeon]; Bloodseeker (19107, -1.91 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Rune of the Guard Captain; main_hand: Felstriker; off_hand: Alcor's Sunrazor

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

