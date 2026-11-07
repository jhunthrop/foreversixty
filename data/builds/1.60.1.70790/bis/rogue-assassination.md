# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00530200100000000-0000000000000000000-000000000000000000)

Set DPS (verified): 45.0. Weights run: 6.0s. Verify run: 15.3s. 241 eligible items had no known source.

Race sweep: night-elf 45.0 ± 0.0, high-order-skyborne 44.4 ± 0.0, human 44.4 ± 0.0, gnome 44.2 ± 0.0, dwarf 43.6 ± 0.0 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.733 ± 0.009, crit=0.443 ± 0.007 per rating point (14 rating = 1%, 6.197 per %), hit=0.611 ± 0.017 per rating point (10 rating = 1%, 6.111 per %), melee_haste=3.445 ± 0.272

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.9 attack_power points (1.09 DPS) | yes | Defender's Leather Hood (252447, -0.51 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.4 attack_power points (0.82 DPS) | yes | Erudite's Amulet (277204, -0.55 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.05 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.7 attack_power points (0.68 DPS) | yes | Slime-encrusted Pads (6461, -0.69 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.1 attack_power points (1.50 DPS) | yes | Brawler's Leather Armor (252490, -0.15 DPS) [crafted]; Prospector's Chestpiece (14562, -0.52 DPS) [world_drop]; Defender's Leather Armor (252434, -0.54 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.7 attack_power points (0.68 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.87 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Murloc Scale Belt (5780, -0.85 DPS) [crafted]; Deviate Scale Belt (6468, -4.03 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.14 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.03 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 10.9 attack_power points (0.86 DPS) | yes | Demon Band (12054, -0.54 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; The 1 Ring (8350, -0.64 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.74 DPS) | yes | Demon Band (12054, -0.27 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; The 1 Ring (8350, -0.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.95 DPS) | yes | Tail Spike (6448, -2.04 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.36 DPS) [crafted]; Goblin Screwdriver (1936, -4.06 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.3 attack_power points (17.84 DPS) | yes | Tail Spike (6448, -0.20 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.9 attack_power points (0.54 DPS) | yes | Fine Longbow (11304, -0.25 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 241, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 76.6. Weights run: 6.2s. Verify run: 13.2s. 422 eligible items had no known source.

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

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 130.5. Weights run: 5.8s. Verify run: 22.6s. 562 eligible items had no known source.

Race sweep: night-elf 130.5 ± 0.1, high-order-skyborne 130.1 ± 0.1, human 129.8 ± 0.1, gnome 129.8 ± 0.1, dwarf 128.3 ± 0.1 (dps); the guide race night-elf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.654 ± 0.010, crit=0.978 ± 0.015 per rating point (14 rating = 1%, 13.689 per %), hit=1.612 ± 0.066 per rating point (10 rating = 1%, 16.122 per %), melee_haste=8.260 ± 1.020

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.2 attack_power points (2.77 DPS) | yes | Hawkeye's Helm (14591, -0.67 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.73 DPS) [world_drop]; Nightscape Headband (8176, -0.89 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; Kaleidoscope Chain (13084, -0.72 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.2 attack_power points (2.20 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.47 DPS) [crafted]; Mantle of Thieves (2264, -0.63 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Imperial Cloak (6432, -0.26 DPS) [dungeon]; Parachute Cloak (10518, -0.26 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.31 DPS) [crafted]; Dusky Leather Armor (7374, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Hawkeye's Bracers (14590, -0.58 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.64 DPS) [dungeon]; Dusky Bracers (7378, -0.64 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dog Training Gloves (7756, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.23 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.85 DPS) | yes | Highlander's Chain Girdle (20090, -0.58 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.58 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.87 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.7 attack_power points (3.30 DPS) | yes | Triprunner Dungarees (9624, -0.19 DPS) [quest]; Ferine Leggings (6690, -0.83 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.10 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.2 attack_power points (2.39 DPS) | yes | Imperial Leather Boots (6431, -0.19 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.31 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.2 attack_power points (2.02 DPS) | yes | Assault Band (13095, -0.12 DPS) [world_drop]; Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.32 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.14 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -3.41 DPS, sim-verified) [dungeon]; The Ziggler (8006, -9.86 DPS) [world_drop]; Sacrificial Kris (3187, -10.87 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (39.44 DPS) | yes | Coldrage Dagger (10761, -0.31 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.23 DPS) [world_drop]; Monolithic Bow (9426, -0.29 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.32 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 215.9. Weights run: 6.3s. Verify run: 23.4s. 707 eligible items had no known source.

Race sweep: night-elf 215.9 ± 0.1, gnome 215.1 ± 0.2, human 215.0 ± 0.1, dwarf 214.8 ± 0.1, high-order-skyborne 214.7 ± 0.1 (dps); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.693 ± 0.026, crit=1.366 ± 0.039 per rating point (14 rating = 1%, 19.129 per %), hit=1.945 ± 0.083 per rating point (10 rating = 1%, 19.455 per %), melee_haste=9.252 ± 1.205

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 55.1 attack_power points (6.28 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.06 DPS) [vendor]; Embrace of the Lycan (9479, -1.86 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.91 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.0 attack_power points (3.08 DPS) | yes | Sentinel's Medallion (19539, -0.80 DPS, sim-verified) [rep]; Zealous Shadowshard Pendant (17772, -0.80 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 27.6 attack_power points (3.15 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.06 DPS) [dungeon]; Failed Flying Experiment (9647, -0.08 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.58 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 56.6 attack_power points (6.45 DPS) | yes | Blazewind Breastplate (11193, -0.81 DPS, sim-verified) [quest]; Warbear Harness (15064, -1.72 DPS) [crafted]; Quillward Harness (10583, -1.87 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 29.4 attack_power points (3.35 DPS) | yes | Pridelord Bands (14672, -0.85 DPS) [world_drop]; Skulker's Leather Bracers (252540, -0.96 DPS, sim-verified) [crafted]; Prowler's Leather Bracers (252539, -0.97 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, -3.54 DPS, sim-verified) [vendor] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 39.1 attack_power points (4.46 DPS) | yes | Girdle of Beastial Fury (11686, -0.13 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.39 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.55 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 56.6 attack_power points (6.45 DPS) | yes | Gryphon Rider's Leggings (9652, -1.13 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -1.88 DPS) [world_drop]; Stormshroud Pants (15057, -2.09 DPS) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 33.9 attack_power points (3.86 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS) [crafted]; Prowler's Leather Boots (252468, -0.26 DPS) [crafted]; Whisperwalk Boots (20255, -0.39 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 39.5 attack_power points (4.50 DPS) | yes | Masons Fraternity Ring (9533, -1.79 DPS) [quest]; Mark of Kern (2262, -2.22 DPS) [dungeon]; Assault Band (13095, -2.22 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 25.2 attack_power points (2.88 DPS) | yes | Masons Fraternity Ring (9533, -0.17 DPS) [quest]; Mark of Kern (2262, -0.60 DPS) [dungeon]; Assault Band (13095, -0.60 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.36 DPS, sim-verified) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hookfang Shanker (11635, -3.31 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -4.77 DPS) [quest]; Blade of Eternal Darkness (17780, -6.80 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (58.30 DPS) | yes | Lifeforce Dirk (10750, -3.21 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.76 DPS) [dungeon]; Guttbuster (13139, -0.82 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.94 DPS, sim-verified) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 707, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 342.3. Weights run: 6.3s. Verify run: 73.1s. 1667 eligible items had no known source.

Race sweep: gnome 342.8 ± 0.3, night-elf 342.3 ± 0.3, human 341.4 ± 0.3, high-order-skyborne 341.1 ± 0.3, dwarf 340.9 ± 0.3 (dps); gnome beats the guide race night-elf by 0.5 (0.14%) ± 0.5, beyond error.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.465 ± 0.026, crit=1.224 ± 0.046 per rating point (14 rating = 1%, 17.140 per %), hit=1.326 ± 0.050 per rating point (10 rating = 1%, 13.264 per %), melee_haste=10.805 ± 1.604

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 56.5 attack_power points (7.83 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Ebon Mask (19984, -0.46 DPS) [quest] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 43.1 attack_power points (5.98 DPS) | yes | Will of the Martyr (17044, -1.82 DPS) [quest]; Imperial Jewel (11933, -2.01 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -2.65 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -23.91 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.09 DPS) [rep]; Howler's Furs (272414, -0.10 DPS) [vendor]; Blackveil Cape (11626, -2.14 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cadaverous Armor (14637, +0.00 DPS) [dungeon]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -19.62 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 38.8 attack_power points (5.39 DPS) | yes | Marshal's Leather Armsplints (16460, -0.14 DPS) [pvp]; Deepfury Bracers (13120, -1.78 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -2.34 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.54 DPS) [pvp]; Marshal's Leather Handgrips (231544, -0.57 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.75 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 57.3 attack_power points (7.95 DPS) | yes | Marshal's Leather Cinch (16458, -1.29 DPS) [pvp]; Might of the Timbermaw (19044, -2.19 DPS) [crafted]; Highlander's Leather Girdle (20045, -3.01 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -23.91 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Pads of the Dread Wolf (13210, +0.00 DPS) [dungeon]; Highlander's Leather Boots (20052, -0.16 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.54 DPS) [dungeon]; Painweaver Band (13098, -1.56 DPS) [dungeon]; Naglering (11669, -9.40 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.07 DPS) [dungeon]; Painweaver Band (13098, -0.09 DPS) [dungeon]; Protector's Band (19514, -0.78 DPS) [rep] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heart of Wyrmthalak (22321, -1.63 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Finkle's Skinner (12709, -9.43 DPS, sim-verified) [dungeon] |
| off_hand | Electrified Dagger (19100) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Riphook (12653, +0.00 DPS) [dungeon]; Malgen's Long Bow (22318, -0.07 DPS) [dungeon]; Bloodseeker (19107, -0.31 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Frozen Heart of the Mountain; main_hand: Felstriker; off_hand: Electrified Dagger

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00530311251051031-3023030202012200000-000000000000000000)

Set DPS (verified): 923.2. Weights run: 6.8s. Verify run: 74.0s. 1667 eligible items had no known source.

Race sweep: gnome 927.2 ± 0.8, night-elf 923.2 ± 0.8, high-order-skyborne 922.6 ± 0.8, human 921.1 ± 0.8, dwarf 917.5 ± 0.8 (dps); gnome beats the guide race night-elf by 4.0 (0.43%) ± 1.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.957 ± 0.039, crit=1.753 ± 0.067 per rating point (14 rating = 1%, 24.537 per %), hit=1.799 ± 0.096 per rating point (10 rating = 1%, 17.986 per %), melee_haste=16.843 ± 2.326

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 77.9 attack_power points (20.44 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Eye of Rend (12587, -3.81 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 50.5 attack_power points (13.26 DPS) | yes | Imperial Jewel (11933, -4.86 DPS) [dungeon]; Skibi's Pendant (13089, -5.14 DPS) [world_drop]; Will of the Martyr (17044, -5.39 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -64.23 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of the Honor Guard (20073, -0.58 DPS) [rep]; Blackveil Cape (11626, -3.15 DPS) [dungeon]; Cape of the Black Baron (13340, -5.69 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -60.84 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 49.3 attack_power points (12.93 DPS) | yes | Marshal's Leather Armsplints (16460, -0.29 DPS) [pvp]; Deepfury Bracers (13120, -4.07 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -5.23 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271949, +0.00 DPS) [rep] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 77.1 attack_power points (20.22 DPS) | yes | Marshal's Leather Cinch (16458, -5.02 DPS) [pvp]; Might of the Timbermaw (19044, -6.97 DPS) [crafted]; Highlander's Leather Girdle (20045, -9.39 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (+64.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Sentinel's Leather Pants (22749, +0.00 DPS) [rep]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 46.5 attack_power points (12.20 DPS) | yes | Swiftwalker Boots (12553, -0.26 DPS) [dungeon]; Marshal's Leather Footguards (231546, -0.90 DPS) [pvp]; Swiftfoot Treads (15861, -1.51 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -4.20 DPS) [dungeon]; Blackstone Ring (17713, -4.86 DPS) [dungeon]; Naglering (11669, -15.91 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.94 DPS) [dungeon]; Protector's Band (19514, -2.80 DPS) [rep]; Painweaver Band (13098, -5.87 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Darrowspike (13984, -35.73 DPS, sim-verified) [quest] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Electrified Dagger (19100, -7.49 DPS, sim-verified) [rep] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.36 DPS) [quest]; Riphook (12653, -0.49 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -6.07 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Frozen Heart of the Mountain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00530200100000000-0000000000000000000-000000000000000000)

Set DPS (verified): 45.0. Weights run: 6.0s. Verify run: 14.1s. 240 eligible items had no known source.

Race sweep: undead 45.0 ± 0.0, troll 44.4 ± 0.0, windshaper-skyborne 44.3 ± 0.0, orc 43.8 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.733 ± 0.009, crit=0.443 ± 0.007 per rating point (14 rating = 1%, 6.197 per %), hit=0.611 ± 0.017 per rating point (10 rating = 1%, 6.111 per %), melee_haste=3.445 ± 0.272

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 13.9 attack_power points (1.09 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.4 attack_power points (0.82 DPS) | yes | Erudite's Amulet (277204, -0.49 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.7 attack_power points (0.68 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.07 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.7 attack_power points (0.68 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.27 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.4 attack_power points (1.13 DPS) | yes | Bristlebark Gloves (14572, -0.25 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.27 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Murloc Scale Belt (5780, -0.85 DPS) [crafted]; Deviate Scale Belt (6468, -3.84 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.03 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.12 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 10.9 attack_power points (0.86 DPS) | yes | Bounty Hunter's Ring (5351, -0.45 DPS) [quest]; Demon Band (12054, -0.54 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.74 DPS) | yes | Bounty Hunter's Ring (5351, -0.22 DPS, sim-verified) [quest]; Demon Band (12054, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.95 DPS) | yes | Tail Spike (6448, -2.04 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.36 DPS) [crafted]; Goblin Screwdriver (1936, -4.06 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.3 attack_power points (17.84 DPS) | yes | Tail Spike (6448, -0.19 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 6.9 attack_power points (0.54 DPS) | yes | Fine Longbow (11304, -0.22 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 240, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 30 (undead, 32530200230001000-0000000000000000000-000000000000000000)

Set DPS (verified): 76.9. Weights run: 6.2s. Verify run: 13.4s. 423 eligible items had no known source.

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

No-known-source sample (15 of 423, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow

### Band 40 (undead, 32530200251511001-0000000000000000000-000000000000000000)

Set DPS (verified): 131.4. Weights run: 5.8s. Verify run: 23.5s. 561 eligible items had no known source.

Race sweep: undead 131.4 ± 0.1, windshaper-skyborne 129.9 ± 0.1, troll 129.4 ± 0.1, orc 128.7 ± 0.1 (dps); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.654 ± 0.010, crit=0.978 ± 0.015 per rating point (14 rating = 1%, 13.689 per %), hit=1.612 ± 0.066 per rating point (10 rating = 1%, 16.122 per %), melee_haste=8.260 ± 1.020

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.2 attack_power points (2.77 DPS) | yes | Hawkeye's Helm (14591, -0.66 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.73 DPS) [world_drop]; Nightscape Headband (8176, -0.89 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Zealous Shadowshard Pendant (17772, +0.00 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.59 DPS) [world]; Ethereal Talisman (4430, -0.62 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.2 attack_power points (2.20 DPS) | yes | Flintrock Shoulders (7755, -0.16 DPS) [dungeon]; Nightscape Shoulders (8192, -0.47 DPS) [crafted]; Mantle of Thieves (2264, -0.63 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Quillward Harness (10583, +0.00 DPS) [dungeon]; Nightscape Tunic (8175, -0.31 DPS) [crafted]; Dusky Leather Armor (7374, -0.47 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Imperial Leather Bracers (4061, -0.64 DPS) [dungeon]; Dusky Bracers (7378, -0.64 DPS) [crafted]; Hawkeye's Bracers (14590, -0.72 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dog Training Gloves (7756, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.52 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.85 DPS) | yes | Defiler's Chain Girdle (20152, -0.54 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.58 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.87 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.7 attack_power points (3.30 DPS) | yes | Triprunner Dungarees (9624, -0.19 DPS) [quest]; Ferine Leggings (6690, -0.83 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.10 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.2 attack_power points (2.39 DPS) | yes | Imperial Leather Boots (6431, -0.19 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.31 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.2 attack_power points (2.02 DPS) | yes | Assault Band (13095, -0.12 DPS) [world_drop]; Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Falcon's Hook (7552, -0.32 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.90 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.14 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Toxic Revenger (9453, -9.55 DPS, sim-verified) [dungeon]; The Ziggler (8006, -9.86 DPS) [world_drop]; Sacrificial Kris (3187, -10.87 DPS) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (38.55 DPS) | yes | The Ziggler (8006, -2.65 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.23 DPS) [world_drop]; Monolithic Bow (9426, -0.29 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.32 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 32530200251551031-3000000000000000000-000000000000000000)

Set DPS (verified): 219.2. Weights run: 6.3s. Verify run: 24.7s. 706 eligible items had no known source.

Race sweep: undead 219.2 ± 0.2, windshaper-skyborne 215.6 ± 0.1, troll 215.5 ± 0.1, orc 214.8 ± 0.1 (dps); the guide race undead is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.418 ± 0.008, crit=0.896 ± 0.013 per rating point (14 rating = 1%, 12.542 per %), hit=2.270 ± 0.107 per rating point (10 rating = 1%, 22.699 per %), melee_haste=19.008 ± 1.009

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 48.5 attack_power points (3.06 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; White Bandit Mask (10008, -1.39 DPS) [crafted]; Embrace of the Lycan (9479, -1.85 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.4 attack_power points (1.48 DPS) | yes | Woven Ivy Necklace (19159, -0.29 DPS) [quest]; Scout's Medallion (19535, -0.41 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.02 DPS, sim-verified) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 24.6 attack_power points (1.55 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.03 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.05 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.27 DPS) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 53.2 attack_power points (3.36 DPS) | yes | Warbear Harness (15064, -1.06 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.11 DPS) [quest]; Quillward Harness (10583, -1.16 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.3 attack_power points (1.60 DPS) | yes | Skulker's Leather Bracers (252540, -0.35 DPS) [crafted]; Pridelord Bands (14672, -0.38 DPS) [world_drop]; Branded Leather Bracers (19508, -1.18 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -4.51 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -0.04 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 53.2 attack_power points (3.36 DPS) | yes | Basilisk Hide Pants (1718, -1.48 DPS) [world_drop]; Triprunner Dungarees (9624, -1.56 DPS) [quest]; Serpentskin Leggings (8262, -1.94 DPS, sim-verified) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 29.4 attack_power points (1.86 DPS) | yes | Prowler's Leather Boots (252468, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.07 DPS) [dungeon]; Shadefiend Boots (11675, -0.24 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 42.7 attack_power points (2.70 DPS) | yes | Mark of Kern (2262, -1.43 DPS) [dungeon]; Assault Band (13095, -1.43 DPS) [world_drop]; White Bone Band (11862, -4.26 DPS, sim-verified) [quest] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Mark of Kern (2262, -0.17 DPS) [dungeon]; Assault Band (13095, -0.17 DPS) [world_drop]; White Bone Band (11862, -0.27 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.01 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, -2.52 DPS, sim-verified) [world_drop]; Lifeforce Dirk (10750, -2.64 DPS) [quest]; Hookfang Shanker (11635, -2.84 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (32.29 DPS) | yes | Lifeforce Dirk (10750, -3.18 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.18 DPS) [dungeon]; Houndmaster's Bow (11628, -0.31 DPS) [dungeon]; Skull Splitting Crossbow (13039, -1.03 DPS, sim-verified) [world_drop] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 706, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 32530200251551031-3023030000000000000-011000000000000000)

Set DPS (verified): 350.9. Weights run: 6.3s. Verify run: 79.3s. 1671 eligible items had no known source.

Race sweep: undead 350.9 ± 0.3, windshaper-skyborne 345.9 ± 0.3, troll 345.2 ± 0.3, orc 343.9 ± 0.3 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.783 ± 0.020, crit=2.000 ± 0.036 per rating point (14 rating = 1%, 28.002 per %), hit=3.996 ± 0.180 per rating point (10 rating = 1%, 39.955 per %), melee_haste=33.441 ± 1.621

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -19.89 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.34 DPS) [dungeon]; Deathguard's Cloak (20068, -1.58 DPS) [rep]; Cloak of Flames (3475, -1.76 DPS) [world_drop] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Chestpiece (227060, +0.00 DPS) [pvp]; Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Dawn Armor (252483, -16.71 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.06 DPS) [pvp]; Deepfury Bracers (13120, -0.89 DPS) [world_drop]; Blackmist Armguards (12966, -2.31 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Gauntlets (15063, +0.00 DPS) [crafted]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -2.66 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 117.4 attack_power points (7.39 DPS) | yes | General's Leather Girdle (16557, -4.07 DPS) [pvp]; Defiler's Leather Girdle (20190, -4.22 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -21.16 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 49.4 attack_power points (3.11 DPS) | yes | Swiftwalker Boots (12553, -0.50 DPS) [dungeon]; Pads of the Dread Wolf (13210, -0.59 DPS) [dungeon]; General's Leather Treads (231552, -0.64 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.38 DPS) [dungeon]; Painweaver Band (13098, -2.39 DPS) [dungeon]; Naglering (11669, -6.41 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.31 DPS) [dungeon]; Painweaver Band (13098, -1.32 DPS) [dungeon]; Channeler's Ring (272406, -1.57 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -4.52 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.38 DPS) [crafted]; Heart of Wyrmthalak (22321, -1.44 DPS, sim-verified) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Finkle's Skinner (12709, -9.48 DPS, sim-verified) [dungeon] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Heartseeker (12783, -2.27 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.64 DPS, sim-verified) [world_drop]; Riphook (12653, -1.47 DPS) [dungeon]; Bloodseeker (19107, -1.56 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Rune of the Guard Captain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 00530311251051031-3023030202012200000-000000000000000000)

Set DPS (verified): 923.7. Weights run: 6.8s. Verify run: 71.8s. 1671 eligible items had no known source.

Race sweep: undead 923.7 ± 0.8, windshaper-skyborne 920.4 ± 0.8, troll 917.1 ± 0.8, orc 916.6 ± 0.8 (dps); the guide race undead is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.100 ± 0.002, agility=2.279 ± 0.029, crit=2.510 ± 0.048 per rating point (14 rating = 1%, 35.137 per %), hit=5.320 ± 0.239 per rating point (10 rating = 1%, 53.203 per %), melee_haste=41.108 ± 1.641

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -4.26 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wyrmhide Spaulders (12082, +0.00 DPS) [quest]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -66.47 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Flames (3475, -3.10 DPS) [world_drop]; Stalwart Cloak (272415, -3.10 DPS) [vendor]; Cape of the Black Baron (13340, -4.16 DPS, sim-verified) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Chestpiece (227060, +0.00 DPS) [pvp]; Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Dawn Armor (252483, -59.00 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.12 DPS) [pvp]; Deepfury Bracers (13120, -1.86 DPS) [world_drop]; Blackmist Armguards (12966, -9.24 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -1.16 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 154.3 attack_power points (17.05 DPS) | yes | General's Leather Girdle (16557, -10.11 DPS) [pvp]; Defiler's Leather Girdle (20193, -10.96 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -69.05 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 63.8 attack_power points (7.06 DPS) | yes | Swiftwalker Boots (12553, -1.28 DPS) [dungeon]; General's Leather Treads (231552, -1.51 DPS) [pvp]; Swiftfoot Treads (15861, -1.91 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.22 DPS) [dungeon]; Channeler's Ring (272406, -5.43 DPS) [vendor]; Naglering (11669, -19.17 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.32 DPS) [dungeon]; Channeler's Ring (272406, -3.53 DPS) [vendor]; Blessed Band of Light (272407, -3.53 DPS) [vendor] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.37 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -38.64 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | sim-verified (+11.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Heartseeker (12783, -11.70 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -3.11 DPS) [world_drop]; Ancient Bone Bow (18680, -3.87 DPS) [dungeon]; Bloodseeker (19107, -3.90 DPS) [quest] |

**New at 60:** head: Outlaw's Collar; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Rune of the Guard Captain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1671, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

