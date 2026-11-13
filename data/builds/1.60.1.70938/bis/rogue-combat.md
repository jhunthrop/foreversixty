# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-3113030000000000000-000000000000000000)

Set DPS (verified): 45.7. Weights run: 6.1s. Verify run: 16.2s. 275 eligible items had no known source.

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
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.1 attack_power points (0.83 DPS) | yes | Demon Band (12054, -0.53 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Ring of the Moon (12052, -0.76 DPS) [world_drop] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon]; Demon Band (12054, -0.55 DPS, sim-verified) [world_drop]; Ring of the Moon (12052, -0.63 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.25 DPS) | yes | Tail Spike (6448, -1.96 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.22 DPS) [crafted]; Goblin Screwdriver (1936, -3.89 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.4 attack_power points (17.16 DPS) | yes | Tail Spike (6448, -0.51 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.53 DPS) | yes | Fine Longbow (11304, -0.22 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (gnome, 00000000000000000-3153030000012300000-000000000000000000)

Set DPS (verified): 76.1. Weights run: 6.5s. Verify run: 12.5s. 460 eligible items had no known source.

Race sweep: night-elf 76.8 ± 0.1, gnome 76.1 ± 0.1, human 76.1 ± 0.1, high-order-skyborne 75.9 ± 0.1, dwarf 75.6 ± 0.1 (dps); night-elf beats the guide race gnome by 0.7 (0.90%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.644 ± 0.010, crit=0.629 ± 0.012 per rating point (14 rating = 1%, 8.812 per %), hit=0.875 ± 0.041 per rating point (10 rating = 1%, 8.749 per %), melee_haste=4.783 ± 0.541

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.4 attack_power points (1.44 DPS) | yes | Brawler's Leather Hood (252504, -0.30 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.39 DPS) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.2 attack_power points (1.15 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.80 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.1 attack_power points (2.02 DPS) | yes | Mantle of Thieves (2264, -0.59 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.87 DPS) [crafted]; Bristlebark Amice (14573, -0.90 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.12 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.34 DPS) [pvp]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.0 attack_power points (2.02 DPS) | yes | Brawler's Leather Tunic (252508, -0.36 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.43 DPS) [quest]; Brawler's Leather Armor (252490, -0.57 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (1.22 DPS) | yes | Barbaric Bracers (18948, -0.30 DPS, sim-verified) [crafted]; Cultist's Armguards (270032, -0.34 DPS) [quest]; Insignia Bracers (6410, -0.35 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.9 attack_power points (1.48 DPS) | yes | Heavy Earthen Gloves (7359, -0.08 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.18 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (2.10 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.28 DPS) | yes | Petrolspill Leggings (9509, -0.26 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.26 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.37 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.2 attack_power points (1.42 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.26 DPS) [world_drop]; Disjointed Shoes (277226, -0.36 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.9 attack_power points (1.39 DPS) | yes | Monkey Ring (6748, -0.38 DPS) [quest]; Pyrewood Signet Ring (277210, -0.51 DPS) [quest]; Ring of Precision (1491, -0.53 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.9 attack_power points (1.13 DPS) | yes | Monkey Ring (6748, -0.12 DPS) [quest]; Pyrewood Signet Ring (277210, -0.25 DPS) [quest]; Ring of Precision (1491, -0.27 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thornspike (6681, -0.29 DPS, sim-verified) [dungeon]; Vendetta (776, -0.58 DPS) [dungeon]; Torturing Poker (7682, -1.01 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (25.20 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.9 attack_power points (0.87 DPS) | yes | Gravespike Repeater (273033, -0.16 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.29 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (gnome, 00000000000000000-3153030000012321511-000000000000000000)

Set DPS (verified): 140.2. Weights run: 7.1s. Verify run: 26.8s. 662 eligible items had no known source.

Race sweep: night-elf 141.0 ± 0.1, gnome 140.2 ± 0.1, human 139.8 ± 0.1, high-order-skyborne 139.7 ± 0.1, dwarf 139.0 ± 0.1 (dps); night-elf beats the guide race gnome by 0.8 (0.59%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.705 ± 0.013, crit=0.938 ± 0.017 per rating point (14 rating = 1%, 13.131 per %), hit=1.306 ± 0.063 per rating point (10 rating = 1%, 13.060 per %), melee_haste=6.758 ± 0.775

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.8 attack_power points (3.06 DPS) | yes | Hawkeye's Helm (14591, -0.73 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.78 DPS) [world_drop]; Nightscape Headband (8176, -0.95 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 18.8 attack_power points (1.93 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.49 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.69 DPS) [world]; Kaleidoscope Chain (13084, -0.82 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 attack_power points (2.44 DPS) | yes | Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Nightscape Shoulders (8192, -0.51 DPS) [crafted]; Mantle of Thieves (2264, -0.69 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Imperial Cloak (6432, -0.27 DPS) [dungeon]; Parachute Cloak (10518, -0.27 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 40.4 attack_power points (4.15 DPS) | yes | Wolffear Harness (13110, -1.18 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.52 DPS) [crafted]; Dusky Leather Armor (7374, -1.70 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.05 DPS) | yes | Hawkeye's Bracers (14590, -0.61 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.65 DPS) [dungeon]; Dusky Bracers (7378, -0.65 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.85 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (3.08 DPS) | yes | Highlander's Chain Girdle (20090, -0.62 DPS) [rep]; Ogron's Sash (13117, -0.62 DPS, sim-verified) [world_drop]; Skulker's Leather Belt (252520, -0.89 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.8 attack_power points (3.68 DPS) | yes | Triprunner Dungarees (9624, -0.22 DPS) [quest]; Ferine Leggings (6690, -1.01 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.23 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.8 attack_power points (2.65 DPS) | yes | Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.29 DPS) [crafted]; Excelsior Boots (4109, -0.35 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.6 attack_power points (2.22 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Ring of the Underwood (2951, -0.27 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.05 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.17 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -5.15 DPS, sim-verified) [dungeon]; The Ziggler (8006, -10.66 DPS) [world_drop]; Sacrificial Kris (3187, -11.75 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (42.65 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.21 DPS) [world_drop]; Monolithic Bow (9426, -0.30 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 662, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (gnome, 32200000000000000-3153030000012321541-000000000000000000)

Set DPS (verified): 218.8. Weights run: 6.8s. Verify run: 29.1s. 851 eligible items had no known source.

Race sweep: night-elf 219.2 ± 0.2, gnome 218.8 ± 0.2, human 217.6 ± 0.2, high-order-skyborne 217.1 ± 0.1, dwarf 216.7 ± 0.2 (dps); night-elf beats the guide race gnome by 0.5 (0.22%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.543 ± 0.012, crit=0.979 ± 0.018 per rating point (14 rating = 1%, 13.707 per %), hit=1.201 ± 0.056 per rating point (10 rating = 1%, 12.009 per %), melee_haste=3.998 ± 0.684

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 49.7 attack_power points (5.66 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.91 DPS) [vendor]; Embrace of the Lycan (9479, -1.47 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.48 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.1 attack_power points (2.86 DPS) | yes | Sentinel's Medallion (19539, -0.80 DPS, sim-verified) [rep]; Zealous Shadowshard Pendant (17772, -1.26 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.0 attack_power points (2.96 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.06 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.12 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.51 DPS) [dungeon]; Dark Phantom Cape (13122, -0.51 DPS) [world_drop]; Duskbat Drape (19982, -0.68 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | PvP rank 6 · Knight · Alliance [vendor] | 43.7 attack_power points (4.98 DPS) | yes | Warbear Harness (15064, -0.56 DPS) [crafted]; Blazewind Breastplate (11193, -0.60 DPS) [quest]; Quillward Harness (10583, -0.73 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.1 attack_power points (3.09 DPS) | yes | Pridelord Bands (14672, -0.77 DPS) [world_drop]; Branded Leather Bracers (19508, -0.81 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.82 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -4.16 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (4.33 DPS) | yes | Highlander's Leather Girdle (20115, -0.49 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.50 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.62 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | PvP rank 6 · Knight · Alliance [vendor] | 43.7 attack_power points (4.98 DPS) | yes | Serpentskin Leggings (8262, -0.69 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.70 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.29 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 31.1 attack_power points (3.54 DPS) | yes | Albino Crocscale Boots (17728, -0.02 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.12 DPS) [crafted]; Shadefiend Boots (11675, -0.47 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 32.0 attack_power points (3.65 DPS) | yes | Masons Fraternity Ring (9533, -1.19 DPS) [quest]; Mark of Kern (2262, -1.37 DPS) [dungeon]; Assault Band (13095, -1.37 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 23.9 attack_power points (2.72 DPS) | yes | Masons Fraternity Ring (9533, -0.26 DPS) [quest]; Mark of Kern (2262, -0.44 DPS) [dungeon]; Assault Band (13095, -0.44 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.25 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.64 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -2.46 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -6.96 DPS) [quest]; Hookfang Shanker (11635, -7.31 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (58.29 DPS) | yes | Lifeforce Dirk (10750, -1.36 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.07 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.07 DPS) [world_drop]; Houndmaster's Rifle (11629, -0.12 DPS) [dungeon] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 851, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (gnome, 32530100000000000-3153030000012321541-030000000000000000)

Set DPS (verified): 314.0. Weights run: 7.8s. Verify run: 91.4s. 1921 eligible items had no known source.

Race sweep: gnome 314.0 ± 0.3, night-elf 313.4 ± 0.3, high-order-skyborne 313.2 ± 0.3, human 312.2 ± 0.3, dwarf 312.1 ± 0.3 (dps); the guide race gnome is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.470 ± 0.011, crit=1.071 ± 0.019 per rating point (14 rating = 1%, 15.000 per %), hit=1.251 ± 0.048 per rating point (10 rating = 1%, 12.510 per %), melee_haste=10.550 ± 1.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.0 attack_power points (7.65 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Eye of Rend (12587, -1.09 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 41.0 attack_power points (6.15 DPS) | yes | Imperial Jewel (11933, -1.37 DPS, sim-verified) [dungeon]; Will of the Martyr (17044, -1.65 DPS) [quest]; Skibi's Pendant (13089, -2.53 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, +0.00 DPS) [rep]; Highlander's Leather Shoulders (20059, -15.44 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.10 DPS) [rep]; Howler's Furs (272414, -0.23 DPS) [vendor]; Blackveil Cape (11626, -2.32 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -12.66 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 38.9 attack_power points (5.84 DPS) | yes | Marshal's Leather Armsplints (16460, -0.15 DPS) [pvp]; Deepfury Bracers (13120, -1.93 DPS) [world_drop]; Wristguards of Renown (22204, -2.28 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, -1.77 DPS, sim-verified) [rep] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 55.9 attack_power points (8.38 DPS) | yes | Marshal's Leather Cinch (16458, -1.18 DPS) [pvp]; Might of the Timbermaw (19044, -2.14 DPS) [crafted]; Highlander's Leather Girdle (20045, -2.50 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -15.51 DPS, sim-verified) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (6.00 DPS) | yes | Swiftwalker Boots (12553, -0.62 DPS) [dungeon]; Shadowcraft Boots (16711, -0.81 DPS) [dungeon]; Highlander's Leather Boots (20052, -0.95 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.35 DPS) [dungeon]; Protector's Band (19514, -2.00 DPS) [rep]; Naglering (11669, -7.10 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.09 DPS) [dungeon]; Protector's Band (19514, -0.74 DPS) [rep]; Painweaver Band (13098, -0.91 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.11 DPS) [crafted]; Heart of Wyrmthalak (22321, -0.86 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; The Lobotomizer (19324, -11.82 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (95.25 DPS) | yes | Electrified Dagger (19100, -6.13 DPS, sim-verified) [rep] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.56 DPS) [quest]; Malgen's Long Bow (22318, -0.60 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.63 DPS) [world_drop] |

**New at 60:** neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Riphook

No-known-source sample (15 of 1921, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (gnome, 00520300200000000-3043030002012321451-005000000000000000)

Set DPS (verified): 902.4. Weights run: 8.7s. Verify run: 94.9s. 1921 eligible items had no known source.

Race sweep: gnome 902.4 ± 0.7, high-order-skyborne 898.7 ± 0.7, night-elf 894.4 ± 0.7, human 893.7 ± 0.7, dwarf 892.7 ± 0.7 (dps); the guide race gnome is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.477 ± 0.026, crit=2.663 ± 0.042 per rating point (14 rating = 1%, 37.281 per %), hit=3.584 ± 0.209 per rating point (10 rating = 1%, 35.840 per %), melee_haste=16.702 ± 2.536

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 109.0 attack_power points (24.75 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Eye of Rend (12587, -4.23 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 63.3 attack_power points (14.37 DPS) | yes | Star of Mystaria (12103, -3.79 DPS) [dungeon]; Skibi's Pendant (13089, -5.81 DPS) [world_drop]; Emberfury Talisman (12929, -5.91 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -43.20 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.52 DPS) [dungeon]; Cloak of the Honor Guard (20073, -3.96 DPS) [rep]; Blackveil Cape (11626, -5.12 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -39.67 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 59.2 attack_power points (13.44 DPS) | yes | Marshal's Leather Armsplints (16460, -0.25 DPS) [pvp]; Deepfury Bracers (13120, -4.00 DPS) [world_drop]; Blackmist Armguards (12966, -4.05 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.11 DPS) [crafted]; Timbermaw Brawlers (19049, -4.82 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 123.7 attack_power points (28.10 DPS) | yes | Highlander's Leather Girdle (20045, -5.83 DPS, sim-verified) [rep]; Marshal's Leather Cinch (16458, -13.16 DPS) [pvp]; Might of the Timbermaw (19044, -14.97 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -43.20 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.8 attack_power points (14.26 DPS) | yes | Swiftwalker Boots (12553, -1.19 DPS) [dungeon]; Marshal's Leather Footguards (231546, -1.88 DPS) [pvp]; Albino Crocscale Boots (17728, -3.01 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -7.11 DPS) [dungeon]; Painweaver Band (13098, -8.60 DPS) [dungeon]; Naglering (11669, -12.72 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.34 DPS) [dungeon]; Painweaver Band (13098, -4.83 DPS) [dungeon]; Magni's Will (12548, -5.80 DPS) [quest] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (+9.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Smolderweb's Eye (13213, -3.02 DPS, sim-verified) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Shadowblade (2163, -25.25 DPS, sim-verified) [world_drop] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (144.33 DPS) | yes | Heartseeker (12783, -8.91 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.84 DPS) [world_drop]; Ancient Bone Bow (18680, -2.82 DPS) [dungeon]; Bloodseeker (19107, -3.08 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Molten Heart of the Mountain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1921, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00000000000000000-3113030000000000000-000000000000000000)

Set DPS (verified): 47.4. Weights run: 6.1s. Verify run: 15.8s. 273 eligible items had no known source.

Race sweep: undead 47.4 ± 0.0, troll 46.7 ± 0.0, windshaper-skyborne 46.6 ± 0.0, orc 46.2 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.767 ± 0.009, crit=0.458 ± 0.008 per rating point (14 rating = 1%, 6.410 per %), hit=0.574 ± 0.013 per rating point (10 rating = 1%, 5.738 per %), melee_haste=3.295 ± 0.296

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.1 attack_power points (1.07 DPS) | yes | Defender's Leather Hood (252447, -0.48 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.6 attack_power points (0.80 DPS) | yes | Erudite's Amulet (277204, -0.50 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.8 attack_power points (0.67 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Dark Leather Cloak (2316, -0.25 DPS) [crafted]; Bristlebark Cape (14571, -0.27 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.47 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.8 attack_power points (0.67 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.27 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.6 attack_power points (1.10 DPS) | yes | Bristlebark Gloves (14572, -0.25 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.27 DPS) [crafted]; Serpent Gloves (5970, -0.30 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.36 DPS) | yes | Brawler's Leather Belt (252428, -0.52 DPS) [crafted]; Murloc Scale Belt (5780, -0.81 DPS) [crafted]; Deviate Scale Belt (6468, -3.21 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.40 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.51 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.1 attack_power points (0.83 DPS) | yes | Bounty Hunter's Ring (5351, -0.44 DPS) [quest]; Demon Band (12054, -0.53 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.71 DPS) | yes | Bounty Hunter's Ring (5351, -0.38 DPS, sim-verified) [quest]; Demon Band (12054, -0.40 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackfang (2236, -1.15 DPS, sim-verified) [world_drop]; Tail Spike (6448, -1.86 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.13 DPS) [crafted] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.53 DPS) | yes | Fine Longbow (11304, -0.24 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.27 DPS) [crafted]; Light Bow (4576, -0.27 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Assassin's Blade; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 273, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (undead, 00000000000000000-3153030000012300000-000000000000000000)

Set DPS (verified): 76.8. Weights run: 6.5s. Verify run: 13.6s. 460 eligible items had no known source.

Race sweep: undead 76.8 ± 0.1, troll 76.1 ± 0.1, windshaper-skyborne 75.9 ± 0.1, orc 75.6 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.644 ± 0.010, crit=0.629 ± 0.012 per rating point (14 rating = 1%, 8.812 per %), hit=0.875 ± 0.041 per rating point (10 rating = 1%, 8.749 per %), melee_haste=4.783 ± 0.541

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.4 attack_power points (1.44 DPS) | yes | Brawler's Leather Hood (252504, -0.35 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.39 DPS) [crafted]; Humbert's Helm (4724, -0.43 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.2 attack_power points (1.15 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.10 DPS) [world]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.80 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.1 attack_power points (2.02 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.87 DPS) [crafted]; Bristlebark Amice (14573, -0.90 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.12 DPS) [world_drop]; Wolfmaster Cape (6314, -0.40 DPS) [dungeon]; Wildhunter Cloak (16658, -0.40 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.0 attack_power points (2.02 DPS) | yes | Brawler's Leather Tunic (252508, -0.34 DPS) [crafted]; Panther Armor (6670, -0.55 DPS) [quest]; Brawler's Leather Armor (252490, -0.57 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (1.22 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS) [crafted]; Cultist's Armguards (270032, -0.34 DPS) [quest]; Insignia Bracers (6410, -0.35 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.9 attack_power points (1.48 DPS) | yes | Heavy Earthen Gloves (7359, -0.08 DPS) [crafted]; Toughened Leather Gloves (4253, -0.09 DPS) [crafted]; Wolfclaw Gloves (1978, -0.18 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (2.10 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted]; Prowler's Leather Belt (252459, -0.45 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.28 DPS) | yes | Petrolspill Leggings (9509, -0.26 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.26 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.37 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.2 attack_power points (1.42 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.26 DPS) [world_drop]; Warsong Boots (16977, -0.26 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.9 attack_power points (1.39 DPS) | yes | Monkey Ring (6748, -0.38 DPS) [quest]; Band of the Fist (17694, -0.46 DPS) [quest]; Pyrewood Signet Ring (277210, -0.51 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.9 attack_power points (1.13 DPS) | yes | Monkey Ring (6748, -0.12 DPS) [quest]; Band of the Fist (17694, -0.21 DPS) [quest]; Pyrewood Signet Ring (277210, -0.25 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornspike (6681, -0.20 DPS, sim-verified) [dungeon]; Vendetta (776, -0.58 DPS) [dungeon]; Torturing Poker (7682, -1.01 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (25.20 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.9 attack_power points (0.87 DPS) | yes | Gravespike Repeater (273033, -0.16 DPS) [world]; Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Silver Star (3463, -0.29 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (undead, 00000000000000000-3153030000012321511-000000000000000000)

Set DPS (verified): 141.0. Weights run: 7.1s. Verify run: 27.3s. 660 eligible items had no known source.

Race sweep: undead 141.0 ± 0.1, windshaper-skyborne 139.4 ± 0.1, troll 139.3 ± 0.1, orc 138.8 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.705 ± 0.013, crit=0.938 ± 0.017 per rating point (14 rating = 1%, 13.131 per %), hit=1.306 ± 0.063 per rating point (10 rating = 1%, 13.060 per %), melee_haste=6.758 ± 0.775

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.8 attack_power points (3.06 DPS) | yes | Hawkeye's Helm (14591, -0.53 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.78 DPS) [world_drop]; Nightscape Headband (8176, -0.95 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 18.8 attack_power points (1.93 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.69 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -0.69 DPS) [world]; Ethereal Talisman (4430, -0.71 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 attack_power points (2.44 DPS) | yes | Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Nightscape Shoulders (8192, -0.51 DPS) [crafted]; Mantle of Thieves (2264, -0.69 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 40.4 attack_power points (4.15 DPS) | yes | Wolffear Harness (13110, -1.06 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.52 DPS) [crafted]; Dusky Leather Armor (7374, -1.70 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.05 DPS) | yes | Imperial Leather Bracers (4061, -0.65 DPS) [dungeon]; Dusky Bracers (7378, -0.65 DPS) [crafted]; Hawkeye's Bracers (14590, -0.70 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -4.37 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (3.08 DPS) | yes | Ogron's Sash (13117, -0.58 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.62 DPS) [rep]; Skulker's Leather Belt (252520, -0.89 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.8 attack_power points (3.68 DPS) | yes | Triprunner Dungarees (9624, -0.22 DPS) [quest]; Ferine Leggings (6690, -1.01 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.23 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.8 attack_power points (2.65 DPS) | yes | Imperial Leather Boots (6431, -0.21 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.29 DPS) [crafted]; Excelsior Boots (4109, -0.35 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.6 attack_power points (2.22 DPS) | yes | Assault Band (13095, -0.17 DPS) [world_drop]; Ring of the Underwood (2951, -0.27 DPS) [world_drop]; Falcon's Hook (7552, -0.34 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.05 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.10 DPS) [world_drop]; Falcon's Hook (7552, -0.17 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Ziggler (8006, -10.66 DPS) [world_drop]; Tok'kar's Murloc Shanker (9680, -11.73 DPS) [quest]; Toxic Revenger (9453, -12.78 DPS, sim-verified) [dungeon] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (41.69 DPS) | yes | The Ziggler (8006, -2.51 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.21 DPS) [world_drop]; Monolithic Bow (9426, -0.30 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.33 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 660, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 32200000000000000-3153030000012321541-000000000000000000)

Set DPS (verified): 222.4. Weights run: 6.8s. Verify run: 21.1s. 849 eligible items had no known source.

Race sweep: undead 222.4 ± 0.2, troll 218.4 ± 0.2, windshaper-skyborne 218.3 ± 0.2, orc 217.6 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.609 ± 0.013, crit=1.100 ± 0.021 per rating point (14 rating = 1%, 15.396 per %), hit=1.395 ± 0.062 per rating point (10 rating = 1%, 13.953 per %), melee_haste=5.566 ± 0.915

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.4 attack_power points (6.17 DPS) | yes | Blood Guard's Leather Headband (220851, -0.73 DPS) [vendor]; Embrace of the Lycan (9479, -1.59 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -2.73 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.9 attack_power points (3.11 DPS) | yes | Woven Ivy Necklace (19159, -0.65 DPS) [quest]; Scout's Medallion (19535, -0.79 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.7 attack_power points (3.20 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.15 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (222.4 DPS) | yes | Blisterbane Wrap (12552, -0.53 DPS) [dungeon]; Dark Phantom Cape (13122, -0.53 DPS) [world_drop]; Duskbat Drape (19982, -0.72 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 6 · Stone Guard · Horde [vendor] | 47.3 attack_power points (5.68 DPS) | yes | Warbear Harness (15064, -0.89 DPS) [crafted]; Quillward Harness (10583, -1.05 DPS) [dungeon]; Blazewind Breastplate (11193, -1.48 DPS, sim-verified) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 28.1 attack_power points (3.38 DPS) | yes | Skulker's Leather Bracers (252540, -0.79 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.85 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.94 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (222.4 DPS) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; First Sergeant's Leather Gauntlets (220857, -3.82 DPS, sim-verified) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (4.56 DPS) | yes | Defiler's Leather Girdle (20193, -0.31 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.42 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.56 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 6 · Stone Guard · Horde [vendor] | 47.3 attack_power points (5.68 DPS) | yes | Basilisk Hide Pants (1718, -1.63 DPS) [world_drop]; Serpentskin Leggings (8262, -1.66 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -1.85 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 32.2 attack_power points (3.86 DPS) | yes | Skulker's Leather Boots (252469, -0.03 DPS) [crafted]; Prowler's Leather Boots (252468, -0.18 DPS) [crafted]; Shadefiend Boots (11675, -0.54 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.0 attack_power points (4.08 DPS) | yes | White Bone Band (11862, -1.19 DPS) [quest]; Masons Fraternity Ring (9533, -1.37 DPS) [quest]; Mark of Kern (2262, -1.68 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.5 attack_power points (2.94 DPS) | yes | White Bone Band (11862, -0.06 DPS) [quest]; Masons Fraternity Ring (9533, -0.23 DPS) [quest]; Mark of Kern (2262, -0.54 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (222.4 DPS) | yes | Frozen Heart of the Mountain (249469, -1.35 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (222.4 DPS) | yes | Frozen Heart of the Mountain (249469, -1.21 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (222.4 DPS) | yes | Barman Shanker (12791, -2.77 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -7.60 DPS) [quest]; Hookfang Shanker (11635, -7.98 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (61.42 DPS) | yes | Lifeforce Dirk (10750, +0.00 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (222.4 DPS) | yes | Stinging Bow (10624, -0.24 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.24 DPS) [world_drop]; Houndmaster's Rifle (11629, -0.26 DPS) [dungeon] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 849, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 32530100000000000-3153030000012321541-030000000000000000)

Set DPS (verified): 318.0. Weights run: 7.8s. Verify run: 96.1s. 1924 eligible items had no known source.

Race sweep: undead 318.0 ± 0.3, windshaper-skyborne 313.5 ± 0.3, troll 313.5 ± 0.3, orc 312.7 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.470 ± 0.011, crit=1.071 ± 0.019 per rating point (14 rating = 1%, 15.000 per %), hit=1.251 ± 0.048 per rating point (10 rating = 1%, 12.510 per %), melee_haste=10.550 ± 1.076

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.0 attack_power points (7.65 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Eye of Rend (12587, -1.09 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 41.0 attack_power points (6.15 DPS) | yes | Imperial Jewel (11933, -1.14 DPS, sim-verified) [dungeon]; Will of the Martyr (17044, -1.65 DPS) [quest]; Skibi's Pendant (13089, -2.53 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Defiler's Leather Shoulders (20194, -13.55 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.10 DPS) [rep]; Howler's Furs (272414, -0.23 DPS) [vendor]; Blackveil Cape (11626, -2.32 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -11.30 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 38.9 attack_power points (5.84 DPS) | yes | General's Leather Armsplints (16559, -0.15 DPS) [pvp]; Deepfury Bracers (13120, -1.93 DPS) [world_drop]; Wristguards of Renown (22204, -2.28 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -1.73 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 55.9 attack_power points (8.38 DPS) | yes | General's Leather Girdle (16557, -1.18 DPS) [pvp]; Might of the Timbermaw (19044, -2.14 DPS) [crafted]; Defiler's Leather Girdle (20190, -2.25 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Devilsaur Leggings (15062, -14.86 DPS, sim-verified) [crafted] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (6.00 DPS) | yes | Swiftwalker Boots (12553, -0.62 DPS) [dungeon]; Shadowcraft Boots (16711, -0.81 DPS) [dungeon]; Defiler's Leather Boots (20186, -0.95 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.35 DPS) [dungeon]; Legionnaire's Band (19510, -2.00 DPS) [rep]; Naglering (11669, -6.85 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.09 DPS) [dungeon]; Legionnaire's Band (19510, -0.74 DPS) [rep]; Painweaver Band (13098, -0.91 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -1.05 DPS, sim-verified) [dungeon]; Hand of Justice (11815, -1.61 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.72 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (-0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -11.15 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (95.25 DPS) | yes | Glacial Blade (19099, -5.68 DPS, sim-verified) [rep] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.56 DPS) [quest]; Malgen's Long Bow (22318, -0.60 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.63 DPS) [world_drop] |

**New at 60:** neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Riphook

No-known-source sample (15 of 1924, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 00520300200000000-3043030002012321451-005000000000000000)

Set DPS (verified): 900.0. Weights run: 8.7s. Verify run: 91.0s. 1924 eligible items had no known source.

Race sweep: undead 900.0 ± 0.7, windshaper-skyborne 898.4 ± 0.7, orc 896.9 ± 0.7, troll 896.7 ± 0.7 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.477 ± 0.026, crit=2.663 ± 0.042 per rating point (14 rating = 1%, 37.281 per %), hit=3.584 ± 0.209 per rating point (10 rating = 1%, 35.840 per %), melee_haste=16.702 ± 2.536

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 109.0 attack_power points (24.75 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Eye of Rend (12587, -4.23 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 63.3 attack_power points (14.37 DPS) | yes | Skibi's Pendant (13089, -5.81 DPS) [world_drop]; Star of Mystaria (12103, -5.84 DPS, sim-verified) [dungeon]; Emberfury Talisman (12929, -5.91 DPS) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -46.23 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.52 DPS) [dungeon]; Deathguard's Cloak (20068, -3.96 DPS) [rep]; Blackveil Cape (11626, -5.12 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -38.22 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 59.2 attack_power points (13.44 DPS) | yes | General's Leather Armsplints (16559, -0.25 DPS) [pvp]; Deepfury Bracers (13120, -4.00 DPS) [world_drop]; Blackmist Armguards (12966, -4.05 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.11 DPS) [crafted]; Timbermaw Brawlers (19049, -4.82 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 123.7 attack_power points (28.10 DPS) | yes | Defiler's Leather Girdle (20190, -9.11 DPS, sim-verified) [rep]; General's Leather Girdle (16557, -13.16 DPS) [pvp]; Might of the Timbermaw (19044, -14.97 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Plaguehound Leggings (18736, -47.75 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.8 attack_power points (14.26 DPS) | yes | Swiftwalker Boots (12553, -1.19 DPS) [dungeon]; General's Leather Treads (231552, -1.88 DPS) [pvp]; Albino Crocscale Boots (17728, -3.01 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -7.11 DPS) [dungeon]; Painweaver Band (13098, -8.60 DPS) [dungeon]; Naglering (11669, -16.97 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.34 DPS) [dungeon]; Painweaver Band (13098, -4.83 DPS) [dungeon]; Legionnaire's Band (19510, -6.83 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -0.74 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.55 DPS) [crafted]; Cannonball Runner (13382, -4.81 DPS) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -27.03 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (144.33 DPS) | yes | Heartseeker (12783, -7.47 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.84 DPS) [world_drop]; Ancient Bone Bow (18680, -2.82 DPS) [dungeon]; Bloodseeker (19107, -3.08 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1924, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

