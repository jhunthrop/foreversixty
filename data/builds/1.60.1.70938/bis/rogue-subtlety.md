# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-0000000000000000000-035200001000000000)

Set DPS (verified): 45.7. Weights run: 5.6s. Verify run: 14.7s. 275 eligible items had no known source.

Race sweep: night-elf 45.7 ± 0.0, human 45.0 ± 0.0, high-order-skyborne 45.0 ± 0.0, gnome 44.8 ± 0.0, dwarf 44.3 ± 0.0 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.774 ± 0.009, crit=0.451 ± 0.008 per rating point (14 rating = 1%, 6.308 per %), hit=0.568 ± 0.018 per rating point (10 rating = 1%, 5.680 per %), melee_haste=2.468 ± 0.295

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.2 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.6 attack_power points (0.84 DPS) | yes | Erudite's Amulet (277204, -0.57 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.07 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.70 DPS) | yes | Slime-encrusted Pads (6461, -0.72 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.26 DPS) [crafted]; Bristlebark Cape (14571, -0.28 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.5 attack_power points (1.53 DPS) | yes | Brawler's Leather Armor (252490, -0.18 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.54 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.56 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.9 attack_power points (0.70 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.27 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.54 DPS) [crafted]; Murloc Scale Belt (5780, -0.84 DPS) [crafted]; Deviate Scale Belt (6468, -3.42 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.53 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.43 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.1 attack_power points (0.87 DPS) | yes | Demon Band (12054, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; Ring of the Moon (12052, -0.79 DPS) [world_drop] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.74 DPS) | yes | Demon Band (12054, -0.28 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Ring of the Moon (12052, -0.66 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (17.96 DPS) | yes | Tail Spike (6448, -2.03 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.35 DPS) [crafted]; Goblin Screwdriver (1936, -4.05 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.5 attack_power points (17.87 DPS) | yes | Tail Spike (6448, -0.43 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.56 DPS) | yes | Fine Longbow (11304, -0.26 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.28 DPS) [crafted]; Light Bow (4576, -0.28 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 275, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish

### Band 30 (night-elf, 00000000000000000-0000000000000000000-035220031301000100)

Set DPS (verified): 90.4. Weights run: 6.3s. Verify run: 13.2s. 460 eligible items had no known source.

Race sweep: night-elf 90.4 ± 0.1, human 89.6 ± 0.1, high-order-skyborne 89.2 ± 0.1, gnome 88.8 ± 0.1, dwarf 88.2 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.652 ± 0.008, crit=0.624 ± 0.009 per rating point (14 rating = 1%, 8.734 per %), hit=0.930 ± 0.059 per rating point (10 rating = 1%, 9.298 per %), melee_haste=not significant (3.257 ± 0.921)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.5 attack_power points (1.84 DPS) | yes | Brawler's Leather Hood (252504, -0.39 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.50 DPS) [crafted]; Humbert's Helm (4724, -0.55 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.2 attack_power points (1.47 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; River Pride Choker (13087, -1.03 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.2 attack_power points (2.58 DPS) | yes | Mantle of Thieves (2264, -0.77 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -1.10 DPS) [crafted]; Bristlebark Amice (14573, -1.14 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.15 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.44 DPS) [pvp]; Wolfmaster Cape (6314, -0.51 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.1 attack_power points (2.58 DPS) | yes | Brawler's Leather Tunic (252508, -0.45 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.55 DPS) [quest]; Brawler's Leather Armor (252490, -0.73 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (1.55 DPS) | yes | Barbaric Bracers (18948, -0.39 DPS, sim-verified) [crafted]; Cultist's Armguards (270032, -0.44 DPS) [quest]; Insignia Bracers (6410, -0.45 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.9 attack_power points (1.88 DPS) | yes | Heavy Earthen Gloves (7359, -0.10 DPS) [crafted]; Toughened Leather Gloves (4253, -0.11 DPS) [crafted]; Wolfclaw Gloves (1978, -0.22 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (2.67 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.35 DPS) [crafted]; Prowler's Leather Belt (252459, -0.57 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.90 DPS) | yes | Troll's Bane Leggings (13114, -0.32 DPS) [world_drop]; Petrolspill Leggings (9509, -0.40 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.46 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.2 attack_power points (1.81 DPS) | yes | Insignia Boots (4055, -0.33 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.36 DPS, sim-verified) [crafted]; Disjointed Shoes (277226, -0.47 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.9 attack_power points (1.77 DPS) | yes | Pyrewood Signet Ring (277210, -0.62 DPS) [quest]; Ring of Precision (1491, -0.67 DPS) [dungeon]; Theramore Signet (271911, -0.77 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.0 attack_power points (1.44 DPS) | yes | Ring of Precision (1491, -0.34 DPS) [dungeon]; Theramore Signet (271911, -0.44 DPS) [rep]; Pyrewood Signet Ring (277210, -0.60 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Thornspike (6681) | Razorfen Kraul: Aggem Thorncurse [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Claw of the Shadowmancer (2912, -0.10 DPS) [world_drop]; Vendetta (776, -0.83 DPS) [dungeon]; Torturing Poker (7682, -1.38 DPS) [dungeon] |
| off_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Claw of the Shadowmancer (2912, -0.28 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.9 attack_power points (1.10 DPS) | yes | Gravespike Repeater (273033, -0.21 DPS) [world]; Double-barreled Shotgun (2098, -0.22 DPS) [world_drop]; Silver Star (3463, -0.37 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Thornspike; off_hand: Toxic Revenger; ranged: Golemsight Long Gun

No-known-source sample (15 of 460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 154.9. Weights run: 6.4s. Verify run: 20.0s. 662 eligible items had no known source.

Race sweep: night-elf 154.9 ± 0.1, human 153.7 ± 0.1, high-order-skyborne 153.5 ± 0.1, gnome 153.1 ± 0.1, dwarf 151.8 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.711 ± 0.010, crit=0.944 ± 0.013 per rating point (14 rating = 1%, 13.219 per %), hit=1.567 ± 0.090 per rating point (10 rating = 1%, 15.665 per %), melee_haste=not significant (5.835 ± 1.535)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.8 attack_power points (3.64 DPS) | yes | Hawkeye's Helm (14591, -0.87 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.92 DPS) [world_drop]; Nightscape Headband (8176, -1.13 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 18.8 attack_power points (2.29 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.59 DPS) [quest]; Archmagister's Faceted Pendant (273049, -0.83 DPS) [world]; Kaleidoscope Chain (13084, -0.97 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 attack_power points (2.90 DPS) | yes | Flintrock Shoulders (7755, -0.21 DPS) [dungeon]; Nightscape Shoulders (8192, -0.61 DPS) [crafted]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.16 DPS) [world_drop]; Imperial Cloak (6432, -0.31 DPS) [dungeon]; Parachute Cloak (10518, -0.31 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 40.5 attack_power points (4.94 DPS) | yes | Wolffear Harness (13110, -1.39 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.81 DPS) [crafted]; Dusky Leather Armor (7374, -2.02 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.44 DPS) | yes | Hawkeye's Bracers (14590, -0.72 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.77 DPS) [dungeon]; Dusky Bracers (7378, -0.77 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -4.11 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (3.66 DPS) | yes | Ogron's Sash (13117, -0.70 DPS, sim-verified) [world_drop]; Highlander's Chain Girdle (20090, -0.73 DPS) [rep]; Skulker's Leather Belt (252520, -1.05 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.9 attack_power points (4.38 DPS) | yes | Triprunner Dungarees (9624, -0.26 DPS) [quest]; Ferine Leggings (6690, -1.21 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.46 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.8 attack_power points (3.15 DPS) | yes | Imperial Leather Boots (6431, -0.24 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.35 DPS) [crafted]; Excelsior Boots (4109, -0.42 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.7 attack_power points (2.64 DPS) | yes | Assault Band (13095, -0.21 DPS) [world_drop]; Ring of the Underwood (2951, -0.31 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.44 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Coldrage Dagger (10761, -6.61 DPS, sim-verified) [dungeon]; The Ziggler (8006, -12.66 DPS) [world_drop]; Sacrificial Kris (3187, -13.95 DPS) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (50.63 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.25 DPS) [world_drop]; Monolithic Bow (9426, -0.35 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.38 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 662, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00530200000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 237.0. Weights run: 6.7s. Verify run: 19.2s. 851 eligible items had no known source.

Race sweep: night-elf 237.0 ± 0.2, dwarf 235.8 ± 0.2, human 235.7 ± 0.2, high-order-skyborne 235.7 ± 0.2, gnome 234.6 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.570 ± 0.009, crit=1.045 ± 0.013 per rating point (14 rating = 1%, 14.634 per %), hit=1.850 ± 0.095 per rating point (10 rating = 1%, 18.502 per %), melee_haste=not significant (5.460 ± 1.511)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 50.6 attack_power points (7.72 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.23 DPS) [vendor]; Embrace of the Lycan (9479, -1.58 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -3.41 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.4 attack_power points (3.87 DPS) | yes | Sentinel's Medallion (19539, -1.00 DPS, sim-verified) [rep]; Zealous Shadowshard Pendant (17772, -1.74 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.3 attack_power points (4.01 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.09 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.17 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.68 DPS) [dungeon]; Dark Phantom Cape (13122, -0.68 DPS) [world_drop]; Duskbat Drape (19982, -0.91 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | PvP rank 6 · Knight · Alliance [vendor] | 51.1 attack_power points (7.80 DPS) | yes | Warbear Harness (15064, -1.83 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.84 DPS) [quest]; Quillward Harness (10583, -2.03 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.5 attack_power points (4.20 DPS) | yes | Skulker's Leather Bracers (252540, -0.96 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -1.04 DPS) [world_drop]; Branded Leather Bracers (19508, -1.15 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -2.92 DPS, sim-verified) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (5.79 DPS) | yes | Highlander's Leather Girdle (20115, -0.51 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.61 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.79 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | PvP rank 6 · Knight · Alliance [vendor] | 51.1 attack_power points (7.80 DPS) | yes | Gryphon Rider's Leggings (9652, -1.93 DPS, sim-verified) [quest]; Serpentskin Leggings (8262, -1.99 DPS) [world_drop]; Basilisk Hide Pants (1718, -2.77 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 31.4 attack_power points (4.79 DPS) | yes | Albino Crocscale Boots (17728, -0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Shadefiend Boots (11675, -0.63 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 38.5 attack_power points (5.87 DPS) | yes | Masons Fraternity Ring (9533, -2.52 DPS) [quest]; Mark of Kern (2262, -2.82 DPS) [dungeon]; Assault Band (13095, -2.82 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.1 attack_power points (3.68 DPS) | yes | Masons Fraternity Ring (9533, -0.33 DPS) [quest]; Mark of Kern (2262, -0.63 DPS) [dungeon]; Assault Band (13095, -0.63 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, -3.11 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -1.03 DPS, sim-verified) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -4.18 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -10.32 DPS) [quest]; Hookfang Shanker (11635, -10.79 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (78.00 DPS) | yes | Lifeforce Dirk (10750, -1.72 DPS, sim-verified) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.20 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.20 DPS) [world_drop]; Houndmaster's Rifle (11629, -0.24 DPS) [dungeon] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 851, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 407.4. Weights run: 6.9s. Verify run: 63.5s. 1922 eligible items had no known source.

Race sweep: night-elf 407.4 ± 0.3, human 406.3 ± 0.3, high-order-skyborne 406.1 ± 0.3, dwarf 406.0 ± 0.3, gnome 404.3 ± 0.4 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.496 ± 0.008, crit=1.122 ± 0.015 per rating point (14 rating = 1%, 15.714 per %), hit=0.913 ± 0.037 per rating point (10 rating = 1%, 9.131 per %), melee_haste=6.635 ± 1.303

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.7 attack_power points (10.45 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Eye of Rend (12587, -1.73 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 41.7 attack_power points (8.43 DPS) | yes | Imperial Jewel (11933, -1.96 DPS) [dungeon]; Will of the Martyr (17044, -2.37 DPS) [quest]; Skibi's Pendant (13089, -3.49 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Highlander's Lizardhide Shoulders (20060, +0.00 DPS) [rep]; Theramore Pauldrons (271955, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.19 DPS) [rep]; Howler's Furs (272414, -1.07 DPS) [vendor]; Blackveil Cape (11626, -3.13 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 39.4 attack_power points (7.97 DPS) | yes | Marshal's Leather Armsplints (16460, -0.20 DPS) [pvp]; Deepfury Bracers (13120, -2.62 DPS) [world_drop]; Wristguards of Renown (22204, -3.12 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.90 DPS) [pvp]; Marshal's Leather Handgrips (231544, -1.16 DPS) [pvp]; Cadaverous Gloves (14640, -1.49 DPS) [dungeon] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Highlander's Leather Girdle (20045, +0.00 DPS) [rep]; Marshal's Leather Cinch (16458, -0.25 DPS) [pvp]; Might of the Timbermaw (19044, -1.56 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Leggings (15062, +0.00 DPS) [crafted]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, +0.00 DPS) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (8.08 DPS) | yes | Swiftwalker Boots (12553, -0.72 DPS) [dungeon]; Shadowcraft Boots (16711, -1.18 DPS) [dungeon]; Highlander's Leather Boots (20052, -1.22 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.96 DPS) [dungeon]; Protector's Band (19514, -2.10 DPS) [rep]; Painweaver Band (13098, -2.25 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.19 DPS) [dungeon]; Protector's Band (19514, -0.33 DPS) [rep]; Painweaver Band (13098, -0.48 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -5.71 DPS) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (128.40 DPS) | yes | - |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.71 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.74 DPS) [world_drop]; Malgen's Long Bow (22318, -0.81 DPS) [dungeon] |

**New at 60:** neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Molten Heart of the Mountain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Riphook

No-known-source sample (15 of 1922, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 891.0. Weights run: 7.4s. Verify run: 77.1s. 1922 eligible items had no known source.

Race sweep: high-order-skyborne 892.4 ± 0.6, night-elf 891.0 ± 0.6, human 890.4 ± 0.6, gnome 889.6 ± 0.6, dwarf 886.2 ± 0.6 (dps); high-order-skyborne beats the guide race night-elf by 1.4 (0.16%) ± 0.9, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.457 ± 0.020, crit=2.606 ± 0.033 per rating point (14 rating = 1%, 36.489 per %), hit=3.551 ± 0.223 per rating point (10 rating = 1%, 35.508 per %), melee_haste=18.770 ± 2.922

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 107.5 attack_power points (24.05 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Eye of Rend (12587, -4.14 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 62.5 attack_power points (13.98 DPS) | yes | Skibi's Pendant (13089, -5.60 DPS) [world_drop]; Sentinel's Medallion (19538, -5.73 DPS) [rep]; Star of Mystaria (12103, -9.30 DPS, sim-verified) [dungeon] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -45.77 DPS, sim-verified) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Black Baron (13340, -1.49 DPS) [dungeon]; Cloak of the Honor Guard (20073, -3.85 DPS) [rep]; Blackveil Cape (11626, -5.04 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -41.64 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 58.8 attack_power points (13.15 DPS) | yes | Marshal's Leather Armsplints (16460, -0.25 DPS) [pvp]; Deepfury Bracers (13120, -3.92 DPS) [world_drop]; Blackmist Armguards (12966, -3.98 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.12 DPS) [crafted]; Timbermaw Brawlers (19049, -4.68 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 122.6 attack_power points (27.43 DPS) | yes | Highlander's Leather Girdle (20045, -5.05 DPS, sim-verified) [rep]; Marshal's Leather Cinch (16458, -12.79 DPS) [pvp]; Might of the Timbermaw (19044, -14.57 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -45.77 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.2 attack_power points (13.93 DPS) | yes | Swiftwalker Boots (12553, -1.15 DPS) [dungeon]; Marshal's Leather Footguards (231546, -1.83 DPS) [pvp]; Albino Crocscale Boots (17728, -2.93 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.82 DPS) [dungeon]; Painweaver Band (13098, -8.39 DPS) [dungeon]; Naglering (11669, -15.61 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.22 DPS) [dungeon]; Painweaver Band (13098, -4.79 DPS) [dungeon]; Magni's Will (12548, -5.75 DPS) [quest] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (+8.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; The Lobotomizer (19324, -35.85 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (142.14 DPS) | yes | Heartseeker (12783, -8.42 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.87 DPS) [world_drop]; Ancient Bone Bow (18680, -2.75 DPS) [dungeon]; Bloodseeker (19107, -2.98 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Molten Heart of the Mountain; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1922, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (undead, 00000000000000000-0000000000000000000-035200001000000000)

Set DPS (verified): 46.6. Weights run: 5.6s. Verify run: 14.4s. 273 eligible items had no known source.

Race sweep: undead 46.6 ± 0.0, troll 45.8 ± 0.0, windshaper-skyborne 45.8 ± 0.0, orc 45.5 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.774 ± 0.009, crit=0.451 ± 0.008 per rating point (14 rating = 1%, 6.308 per %), hit=0.568 ± 0.018 per rating point (10 rating = 1%, 5.680 per %), melee_haste=2.468 ± 0.295

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.2 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.46 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.6 attack_power points (0.84 DPS) | yes | Erudite's Amulet (277204, -0.52 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.70 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.26 DPS) [crafted]; Bristlebark Cape (14571, -0.28 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.54 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.9 attack_power points (0.70 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.28 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.6 attack_power points (1.15 DPS) | yes | Brawler's Leather Gloves (252494, -0.28 DPS) [crafted]; Bristlebark Gloves (14572, -0.28 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.31 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.41 DPS) | yes | Brawler's Leather Belt (252428, -0.54 DPS) [crafted]; Murloc Scale Belt (5780, -0.84 DPS) [crafted]; Deviate Scale Belt (6468, -3.32 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.50 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.59 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.1 attack_power points (0.87 DPS) | yes | Bounty Hunter's Ring (5351, -0.45 DPS) [quest]; Demon Band (12054, -0.56 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.4 attack_power points (0.74 DPS) | yes | Bounty Hunter's Ring (5351, -0.32 DPS) [quest]; Demon Band (12054, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackfang (2236, -1.03 DPS, sim-verified) [world_drop]; Tail Spike (6448, -1.94 DPS) [dungeon]; Deadly Bronze Poniard (3490, -3.26 DPS) [crafted] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.1 attack_power points (0.56 DPS) | yes | Fine Longbow (11304, -0.24 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.28 DPS) [crafted]; Light Bow (4576, -0.28 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Assassin's Blade; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 273, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (undead, 00000000000000000-0000000000000000000-035220031301000100)

Set DPS (verified): 90.6. Weights run: 6.3s. Verify run: 13.2s. 460 eligible items had no known source.

Race sweep: undead 90.6 ± 0.1, troll 90.0 ± 0.1, windshaper-skyborne 89.9 ± 0.1, orc 89.4 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.652 ± 0.008, crit=0.624 ± 0.009 per rating point (14 rating = 1%, 8.734 per %), hit=0.930 ± 0.059 per rating point (10 rating = 1%, 9.298 per %), melee_haste=not significant (3.257 ± 0.921)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.5 attack_power points (1.84 DPS) | yes | Brawler's Leather Hood (252504, -0.43 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.50 DPS) [crafted]; Humbert's Helm (4724, -0.55 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.2 attack_power points (1.47 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.14 DPS) [world]; Kaleidoscope Chain (13084, -0.29 DPS) [world_drop]; River Pride Choker (13087, -1.03 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.2 attack_power points (2.58 DPS) | yes | Mantle of Thieves (2264, -0.64 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -1.10 DPS) [crafted]; Bristlebark Amice (14573, -1.14 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.15 DPS) [world_drop]; Wolfmaster Cape (6314, -0.51 DPS) [dungeon]; Wildhunter Cloak (16658, -0.51 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.1 attack_power points (2.58 DPS) | yes | Brawler's Leather Tunic (252508, -0.44 DPS) [crafted]; Panther Armor (6670, -0.70 DPS) [quest]; Brawler's Leather Armor (252490, -0.73 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.9 attack_power points (1.55 DPS) | yes | Barbaric Bracers (18948, -0.37 DPS) [crafted]; Cultist's Armguards (270032, -0.44 DPS) [quest]; Insignia Bracers (6410, -0.45 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.9 attack_power points (1.88 DPS) | yes | Heavy Earthen Gloves (7359, -0.10 DPS) [crafted]; Toughened Leather Gloves (4253, -0.11 DPS) [crafted]; Wolfclaw Gloves (1978, -0.22 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (2.67 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.35 DPS) [crafted]; Prowler's Leather Belt (252459, -0.57 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.90 DPS) | yes | Petrolspill Leggings (9509, -0.32 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.32 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.46 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.2 attack_power points (1.81 DPS) | yes | Brawler's Leather Boots (252439, -0.33 DPS) [crafted]; Insignia Boots (4055, -0.33 DPS) [world_drop]; Warsong Boots (16977, -0.33 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.9 attack_power points (1.77 DPS) | yes | Band of the Fist (17694, -0.59 DPS) [quest]; Pyrewood Signet Ring (277210, -0.62 DPS) [quest]; Ring of Precision (1491, -0.67 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.0 attack_power points (1.44 DPS) | yes | Band of the Fist (17694, -0.26 DPS) [quest]; Pyrewood Signet Ring (277210, -0.29 DPS) [quest]; Ring of Precision (1491, -0.34 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornspike (6681, -0.57 DPS, sim-verified) [dungeon]; Vendetta (776, -0.73 DPS) [dungeon]; Torturing Poker (7682, -1.28 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (32.03 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.9 attack_power points (1.10 DPS) | yes | Gravespike Repeater (273033, -0.21 DPS) [world]; Double-barreled Shotgun (2098, -0.22 DPS) [world_drop]; Silver Star (3463, -0.37 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk

### Band 40 (undead, 00000000000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 154.8. Weights run: 6.4s. Verify run: 22.4s. 660 eligible items had no known source.

Race sweep: undead 154.8 ± 0.1, windshaper-skyborne 153.0 ± 0.1, troll 153.0 ± 0.1, orc 152.7 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.711 ± 0.010, crit=0.944 ± 0.013 per rating point (14 rating = 1%, 13.219 per %), hit=1.567 ± 0.090 per rating point (10 rating = 1%, 15.665 per %), melee_haste=not significant (5.835 ± 1.535)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 29.8 attack_power points (3.64 DPS) | yes | Hawkeye's Helm (14591, -0.85 DPS) [world_drop]; Warden's Wizard Hat (14604, -0.92 DPS) [world_drop]; Nightscape Headband (8176, -1.13 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 18.8 attack_power points (2.29 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.76 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -0.83 DPS) [world]; Ethereal Talisman (4430, -0.85 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.8 attack_power points (2.90 DPS) | yes | Flintrock Shoulders (7755, -0.21 DPS) [dungeon]; Nightscape Shoulders (8192, -0.61 DPS) [crafted]; Mantle of Thieves (2264, -0.82 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.16 DPS) [dungeon]; Parachute Cloak (10518, -0.16 DPS) [crafted]; Tigerstrike Mantle (13108, -0.16 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 40.5 attack_power points (4.94 DPS) | yes | Wolffear Harness (13110, -1.23 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -1.81 DPS) [crafted]; Dusky Leather Armor (7374, -2.02 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (2.44 DPS) | yes | Imperial Leather Bracers (4061, -0.77 DPS) [dungeon]; Dusky Bracers (7378, -0.77 DPS) [crafted]; Hawkeye's Bracers (14590, -0.87 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.68 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (3.66 DPS) | yes | Ogron's Sash (13117, -0.68 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.73 DPS) [rep]; Skulker's Leather Belt (252520, -1.05 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 35.9 attack_power points (4.38 DPS) | yes | Triprunner Dungarees (9624, -0.26 DPS) [quest]; Ferine Leggings (6690, -1.21 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.46 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 25.8 attack_power points (3.15 DPS) | yes | Imperial Leather Boots (6431, -0.24 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.35 DPS) [crafted]; Excelsior Boots (4109, -0.42 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.7 attack_power points (2.64 DPS) | yes | Assault Band (13095, -0.21 DPS) [world_drop]; Ring of the Underwood (2951, -0.31 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (2.44 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.11 DPS) [world_drop]; Falcon's Hook (7552, -0.20 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Ziggler (8006, -12.66 DPS) [world_drop]; Toxic Revenger (9453, -12.73 DPS, sim-verified) [dungeon]; Tok'kar's Murloc Shanker (9680, -13.92 DPS) [quest] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (49.49 DPS) | yes | The Ziggler (8006, -2.29 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwind (13038, -0.25 DPS) [world_drop]; Monolithic Bow (9426, -0.35 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.38 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 660, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 50 (undead, 00530200000000000-0000000000000000000-035230031321001151)

Set DPS (verified): 240.0. Weights run: 6.7s. Verify run: 15.9s. 849 eligible items had no known source.

Race sweep: undead 240.0 ± 0.2, windshaper-skyborne 237.2 ± 0.2, troll 236.9 ± 0.2, orc 236.3 ± 0.2 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.570 ± 0.009, crit=1.045 ± 0.013 per rating point (14 rating = 1%, 14.634 per %), hit=1.850 ± 0.095 per rating point (10 rating = 1%, 18.502 per %), melee_haste=not significant (5.460 ± 1.511)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 50.6 attack_power points (7.72 DPS) | yes | Blood Guard's Leather Headband (220851, -0.23 DPS) [vendor]; Embrace of the Lycan (9479, -1.46 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -3.41 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.4 attack_power points (3.87 DPS) | yes | Woven Ivy Necklace (19159, -0.80 DPS) [quest]; Scout's Medallion (19535, -1.00 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.3 attack_power points (4.01 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.09 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.17 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.68 DPS) [dungeon]; Dark Phantom Cape (13122, -0.68 DPS) [world_drop]; Duskbat Drape (19982, -0.91 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 6 · Stone Guard · Horde [vendor] | 51.1 attack_power points (7.80 DPS) | yes | Warbear Harness (15064, -1.64 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.84 DPS) [quest]; Quillward Harness (10583, -2.03 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.5 attack_power points (4.20 DPS) | yes | Skulker's Leather Bracers (252540, -0.88 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -1.04 DPS) [world_drop]; Branded Leather Bracers (19508, -1.15 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; First Sergeant's Leather Gauntlets (220857, -3.29 DPS, sim-verified) [vendor] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (5.79 DPS) | yes | Defiler's Leather Girdle (20193, -0.51 DPS) [rep]; Skulker's Leather Waistguard (252474, -0.61 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.79 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 6 · Stone Guard · Horde [vendor] | 51.1 attack_power points (7.80 DPS) | yes | Serpentskin Leggings (8262, -2.40 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -2.77 DPS) [world_drop]; Triprunner Dungarees (9624, -3.03 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 31.4 attack_power points (4.79 DPS) | yes | Albino Crocscale Boots (17728, -0.00 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.17 DPS) [crafted]; Shadefiend Boots (11675, -0.63 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 38.5 attack_power points (5.87 DPS) | yes | White Bone Band (11862, -2.21 DPS) [quest]; Masons Fraternity Ring (9533, -2.52 DPS) [quest]; Mark of Kern (2262, -2.82 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.1 attack_power points (3.68 DPS) | yes | White Bone Band (11862, -0.02 DPS) [quest]; Masons Fraternity Ring (9533, -0.33 DPS) [quest]; Mark of Kern (2262, -0.63 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.57 DPS) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Barman Shanker (12791, -2.99 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -10.32 DPS) [quest]; Hookfang Shanker (11635, -10.79 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.5 attack_power points (78.00 DPS) | yes | Lifeforce Dirk (10750, +0.00 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stinging Bow (10624, -0.20 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.20 DPS) [world_drop]; Houndmaster's Rifle (11629, -0.24 DPS) [dungeon] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 849, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60 (undead, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 410.3. Weights run: 6.9s. Verify run: 73.5s. 1925 eligible items had no known source.

Race sweep: undead 410.3 ± 0.3, windshaper-skyborne 407.3 ± 0.3, troll 406.6 ± 0.3, orc 406.5 ± 0.4 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.496 ± 0.008, crit=1.122 ± 0.015 per rating point (14 rating = 1%, 15.714 per %), hit=0.913 ± 0.037 per rating point (10 rating = 1%, 9.131 per %), melee_haste=6.635 ± 1.303

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 51.7 attack_power points (10.45 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Eye of Rend (12587, -1.73 DPS) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 41.7 attack_power points (8.43 DPS) | yes | Imperial Jewel (11933, -1.96 DPS) [dungeon]; Will of the Martyr (17044, -2.37 DPS) [quest]; Skibi's Pendant (13089, -3.49 DPS) [world_drop] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Lizardhide Shoulders (20175, +0.00 DPS) [rep]; Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -0.19 DPS) [rep]; Howler's Furs (272414, -1.07 DPS) [vendor]; Blackveil Cape (11626, -3.13 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Dawn Armor (252483, +0.00 DPS) [crafted]; Timbermaw Tunic (252484, +0.00 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 39.4 attack_power points (7.97 DPS) | yes | General's Leather Armsplints (16559, -0.20 DPS) [pvp]; Deepfury Bracers (13120, -2.62 DPS) [world_drop]; Wristguards of Renown (22204, -3.12 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blood Guard's Leather Vices (16499, -0.90 DPS) [pvp]; General's Leather Mitts (231555, -1.16 DPS) [pvp]; Cadaverous Gloves (14640, -1.49 DPS) [dungeon] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Leather Girdle (20190, +0.00 DPS) [rep]; General's Leather Girdle (16557, -0.25 DPS) [pvp]; Might of the Timbermaw (19044, -1.56 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Leggings (15062, +0.00 DPS) [crafted]; Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (8.08 DPS) | yes | Swiftwalker Boots (12553, -0.72 DPS) [dungeon]; Shadowcraft Boots (16711, -1.18 DPS) [dungeon]; Defiler's Leather Boots (20186, -1.22 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -1.96 DPS) [dungeon]; Legionnaire's Band (19510, -2.10 DPS) [rep]; Painweaver Band (13098, -2.25 DPS) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.19 DPS) [dungeon]; Legionnaire's Band (19510, -0.33 DPS) [rep]; Painweaver Band (13098, -0.48 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (410.3 DPS) | yes | Hand of Justice (11815, -1.70 DPS) [dungeon]; Smolderweb's Eye (13213, -2.28 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.46 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (410.3 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -5.71 DPS) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (128.40 DPS) | yes | - |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.71 DPS) [quest]; Precisely Calibrated Boomstick (2100, -0.74 DPS) [world_drop]; Malgen's Long Bow (22318, -0.81 DPS) [dungeon] |

**New at 60:** neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Pads of the Dread Wolf; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Riphook

No-known-source sample (15 of 1925, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

### Band 60, raid preset (undead, 00530201100000000-3200030000000000000-035230031321001151)

Set DPS (verified): 896.6. Weights run: 7.4s. Verify run: 82.9s. 1925 eligible items had no known source.

Race sweep: undead 896.6 ± 0.7, windshaper-skyborne 893.4 ± 0.6, orc 893.2 ± 0.6, troll 891.2 ± 0.6 (dps); the guide race undead is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.457 ± 0.020, crit=2.606 ± 0.033 per rating point (14 rating = 1%, 36.489 per %), hit=3.551 ± 0.223 per rating point (10 rating = 1%, 35.508 per %), melee_haste=18.770 ± 2.922

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Eye of Rend (12587, -3.75 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 62.5 attack_power points (13.98 DPS) | yes | Star of Mystaria (12103, -4.42 DPS, sim-verified) [dungeon]; Skibi's Pendant (13089, -5.60 DPS) [world_drop]; Scout's Medallion (19534, -5.73 DPS) [rep] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -49.05 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Cape of the Black Baron (13340, -2.30 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -3.85 DPS) [rep]; Blackveil Cape (11626, -5.04 DPS) [dungeon] |
| chest | Stormshroud Armor (15056) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, +0.00 DPS) [pvp]; Timbermaw Tunic (252484, +0.00 DPS) [crafted]; Dawn Armor (252483, -38.71 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 58.8 attack_power points (13.15 DPS) | yes | General's Leather Armsplints (16559, -0.25 DPS) [pvp]; Deepfury Bracers (13120, -3.92 DPS) [world_drop]; Blackmist Armguards (12966, -3.98 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.12 DPS) [crafted] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 122.6 attack_power points (27.43 DPS) | yes | General's Leather Girdle (16557, -12.79 DPS) [pvp]; Defiler's Leather Girdle (20190, -12.89 DPS, sim-verified) [rep]; Might of the Timbermaw (19044, -14.57 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -51.63 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 62.2 attack_power points (13.93 DPS) | yes | Swiftwalker Boots (12553, -1.15 DPS) [dungeon]; General's Leather Treads (231552, -1.83 DPS) [pvp]; Albino Crocscale Boots (17728, -2.93 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -6.82 DPS) [dungeon]; Painweaver Band (13098, -8.39 DPS) [dungeon]; Naglering (11669, -16.06 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.22 DPS) [dungeon]; Painweaver Band (13098, -4.79 DPS) [dungeon]; Legionnaire's Band (19510, -6.64 DPS) [rep] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -9.06 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -0.75 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.54 DPS) [crafted]; Hand of Justice (11815, -3.24 DPS, sim-verified) [dungeon] |
| main_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; The Lobotomizer (19324, -34.78 DPS, sim-verified) [rep] |
| off_hand | Alcor's Sunrazor (14555) | World drop [world_drop] | 635.4 attack_power points (142.14 DPS) | yes | Heartseeker (12783, -12.21 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Precisely Calibrated Boomstick (2100, -1.91 DPS, sim-verified) [world_drop]; Ancient Bone Bow (18680, -2.75 DPS) [dungeon]; Bloodseeker (19107, -2.98 DPS) [quest] |

**New at 60:** head: Mask of the Unforgiven; neck: Mark of Fordring; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Stormshroud Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; main_hand: Felstriker; off_hand: Alcor's Sunrazor; ranged: Blackcrow

No-known-source sample (15 of 1925, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band

