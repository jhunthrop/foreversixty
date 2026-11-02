# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5320100000000000000)

Set DPS (verified): 38.4. Weights run: 1.8s. Verify run: 4.5s. 242 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.848 ± 0.013, crit=0.541 ± 0.012 per rating point (14 rating = 1%, 7.577 per %), hit=0.630 ± 0.030 per rating point (10 rating = 1%, 6.303 per %), melee_haste=3.650 ± 0.304

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.8 attack_power points (0.95 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.7 attack_power points (0.24 DPS) | yes | Scholarly Pendant (277203, -0.25 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.43 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.2 attack_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.62 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.1 attack_power points (0.72 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.23 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 20.3 attack_power points (1.31 DPS) | yes | Brawler's Leather Armor (252490, -0.15 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.48 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.2 attack_power points (0.60 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.4 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.32 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.43 DPS) [crafted]; Murloc Scale Belt (5780, -0.67 DPS) [crafted]; Deviate Scale Belt (6468, -4.35 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (38.4 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.54 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (38.4 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.46 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.9 attack_power points (0.64 DPS) | yes | Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; The 1 Ring (8350, -0.46 DPS) [world]; Ring of the Moon (12052, -0.58 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 4.0 attack_power points (0.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.02 DPS) [dungeon]; The 1 Ring (8350, -0.07 DPS) [world]; Ring of the Moon (12052, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (14.76 DPS) | yes | Tail Spike (6448, -1.66 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.74 DPS) [crafted]; Goblin Screwdriver (1936, -3.31 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.8 attack_power points (14.70 DPS) | yes | Tail Spike (6448, -0.45 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.4 attack_power points (0.48 DPS) | yes | Fine Longbow (11304, -0.21 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.24 DPS) [crafted]; Light Bow (4576, -0.24 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Demon Band; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 242, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 00000000000000000-00000000000000000-5320201310013000000)

Set DPS (verified): 60.1. Weights run: 2.0s. Verify run: 4.5s. 436 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.535 ± 0.011, crit=0.557 ± 0.012 per rating point (14 rating = 1%, 7.795 per %), hit=0.795 ± 0.031 per rating point (10 rating = 1%, 7.951 per %), melee_haste=3.809 ± 0.324

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.3 attack_power points (1.07 DPS) | yes | Defender's Leather Helm (252455, -0.23 DPS) [crafted]; Brawler's Leather Hood (252504, -0.25 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.32 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.84 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; River Pride Choker (13087, -0.56 DPS) [world_drop]; Erudite's Amulet (277204, -0.62 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.9 attack_power points (1.53 DPS) | yes | Mantle of Thieves (2264, -0.50 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.64 DPS) [crafted]; Bristlebark Amice (14573, -0.67 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.7 attack_power points (0.96 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Wolfmaster Cape (6314, -0.26 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.5 attack_power points (1.50 DPS) | yes | Tunic of Westfall (2041, -0.32 DPS) [quest]; Brawler's Leather Tunic (252508, -0.33 DPS, sim-verified) [crafted]; Raptorbane Armor (3566, -0.38 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.2 attack_power points (0.92 DPS) | yes | Cultist's Armguards (270032, -0.22 DPS) [quest]; Barbaric Bracers (18948, -0.25 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.28 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.55 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Belt (252459, +0.00 DPS) [crafted]; Warden's Leather Belt (252460, -0.19 DPS) [crafted]; Skulker's Leather Belt (252520, -3.16 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.62 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.90 DPS, sim-verified) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.6 attack_power points (0.88 DPS) | yes | Ring of Precision (1491, -0.24 DPS) [dungeon]; Monkey Ring (6748, -0.34 DPS) [quest]; Seal of Wrynn (2933, -0.35 DPS) [quest] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.3 attack_power points (0.65 DPS) | yes | Ring of Precision (1491, -0.01 DPS) [dungeon]; Monkey Ring (6748, -0.11 DPS) [quest]; Seal of Wrynn (2933, -0.12 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (60.1 DPS) | yes | Vendetta (776, -0.52 DPS) [dungeon]; Thornspike (6681, -0.61 DPS, sim-verified) [dungeon]; Torturing Poker (7682, -0.80 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (20.04 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.2 attack_power points (0.64 DPS) | yes | Gravespike Repeater (273033, -0.08 DPS) [world]; Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.21 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Thunderbrow Ring; finger2: Pyrewood Signet Ring; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 436, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-00000000000000000-5320201310013211501)

Set DPS (verified): 95.7. Weights run: 2.1s. Verify run: 3.1s. 585 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.356 ± 0.009, crit=0.462 ± 0.012 per rating point (14 rating = 1%, 6.464 per %), hit=0.838 ± 0.032 per rating point (10 rating = 1%, 8.381 per %), melee_haste=3.977 ± 0.382

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.9 attack_power points (1.92 DPS) | yes | Hawkeye's Helm (14591, -0.61 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.61 DPS) [world_drop]; Nightscape Headband (8176, -0.71 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.89 DPS) | yes | Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; River Pride Choker (13087, -0.59 DPS) [world_drop]; Gazlowe's Charm (13088, -0.59 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 19.9 attack_power points (1.47 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Nightscape Shoulders (8192, -0.37 DPS) [crafted]; Mantle of Thieves (2264, -0.47 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 14.1 attack_power points (1.05 DPS) | yes | Hawkeye's Cloak (14593, -0.12 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.22 DPS) [quest]; Parachute Cloak (10518, -0.24 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 23.1 attack_power points (1.71 DPS) | yes | Dusky Leather Armor (7374, -0.30 DPS) [crafted]; Nightscape Tunic (8175, -0.32 DPS, sim-verified) [crafted]; Barbaric Harness (5739, -0.38 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Hawkeye's Bracers (14590, -0.51 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.68 DPS) [dungeon]; Dusky Bracers (7378, -0.68 DPS) [crafted] |
| hands | Dog Training Gloves (7756) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 24.0 attack_power points (1.78 DPS) | yes | Skulker's Leather Gloves (252525, -0.11 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.13 DPS) [crafted]; Imperial Leather Gloves (4063, -0.21 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 21.2 attack_power points (1.57 DPS) | yes | Skulker's Leather Belt (252520, -0.22 DPS) [crafted]; Blackened Defias Belt (10403, -0.24 DPS) [dungeon]; Scorpashi Sash (14652, -0.27 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.5 attack_power points (2.11 DPS) | yes | Triprunner Dungarees (9624, -0.08 DPS) [quest]; Ferine Leggings (6690, -0.18 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.69 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.9 attack_power points (1.62 DPS) | yes | Prowler's Leather Shoes (252465, -0.11 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Excelsior Boots (4109, -0.20 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Falcon's Hook (7552, -0.35 DPS) [dungeon]; Thunderbrow Ring (13097, -0.59 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Falcon's Hook (7552, -0.35 DPS) [dungeon]; Thunderbrow Ring (13097, -0.59 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (95.7 DPS) | yes | The Ziggler (8006, -7.68 DPS) [world_drop]; Sacrificial Kris (3187, -8.46 DPS) [dungeon]; Coldrage Dagger (10761, -8.91 DPS, sim-verified) [dungeon] |
| off_hand | Black Menace (6831) | In the Name of the Light [quest] | 415.3 attack_power points (30.72 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (95.7 DPS) | yes | Monolithic Bow (9426, -0.29 DPS) [dungeon]; Swiftwind (13038, -0.33 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.34 DPS) [quest] |

**New at 40:** head: White Bandit Mask; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Dog Training Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 00532000000000000-00000000000000000-5320201310013211501)

Set DPS (verified): 135.0. Weights run: 2.7s. Verify run: 3.5s. 736 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.519 ± 0.008, crit=0.868 ± 0.012 per rating point (14 rating = 1%, 12.152 per %), hit=1.451 ± 0.056 per rating point (10 rating = 1%, 14.512 per %), melee_haste=5.583 ± 0.793

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 27.7 attack_power points (2.39 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.16 DPS) [crafted]; Lordrec Helmet (10741, -0.29 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.7 attack_power points (2.14 DPS) | yes | Kaleidoscope Chain (13084, -1.27 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.39 DPS, sim-verified) [world]; River Pride Choker (13087, -1.79 DPS) [world_drop] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 25.7 attack_power points (2.22 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.09 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 22.8 attack_power points (1.97 DPS) | yes | Blackflame Cape (13109, -0.48 DPS, sim-verified) [world_drop]; Serpentskin Cloak (8259, -0.52 DPS) [world_drop]; Pridelord Cape (14673, -0.53 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 44.7 attack_power points (3.86 DPS) | yes | Warbear Harness (15064, -0.55 DPS) [crafted]; Blazewind Breastplate (11193, -0.58 DPS) [quest]; Charred Leather Tunic (19127, -1.50 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 26.8 attack_power points (2.31 DPS) | yes | Pridelord Bands (14672, -0.57 DPS) [world_drop]; Branded Leather Bracers (19508, -0.59 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.67 DPS, sim-verified) [crafted] |
| hands | Skulker's Leather Gauntlets (252548) | Leatherworking [crafted] | 30.3 attack_power points (2.62 DPS) | yes | Serpentskin Gloves (8260, -0.22 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.22 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, -0.36 DPS) [vendor] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 33.3 attack_power points (2.87 DPS) | yes | Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted]; Sagebrush Girdle (17778, -0.91 DPS) [quest]; Ogron's Sash (13117, -0.92 DPS) [world_drop] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 44.7 attack_power points (3.86 DPS) | yes | Gryphon Rider's Leggings (9652, -0.59 DPS) [quest]; Serpentskin Leggings (8262, -0.64 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.10 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 30.7 attack_power points (2.66 DPS) | yes | Albino Crocscale Boots (17728, -0.03 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.61 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.5 attack_power points (2.98 DPS) | yes | Mark of Kern (2262, -1.25 DPS) [dungeon]; Assault Band (13095, -1.25 DPS) [world_drop]; Ring of the Underwood (2951, -1.50 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.3 attack_power points (1.84 DPS) | yes | Mark of Kern (2262, -0.11 DPS) [dungeon]; Assault Band (13095, -0.11 DPS) [world_drop]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (135.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (135.0 DPS) | yes | Mark of the Chosen (17774, -1.03 DPS, sim-verified) [quest] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (135.0 DPS) | yes | Widowmaker (4091, -1.14 DPS) [world_drop]; Gahz'rilla Fang (9467, -2.02 DPS) [dungeon]; Julie's Dagger (6660, -6.48 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.1 attack_power points (43.48 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (135.0 DPS) | yes | Guttbuster (13139, -0.53 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.53 DPS, sim-verified) [world_drop]; Houndmaster's Bow (11628, -0.54 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Skulker's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 736, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 00532310101400000-00000000000000000-5320201310013211501)

Set DPS (verified): 219.6. Weights run: 2.6s. Verify run: 5.7s. 1495 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.711 ± 0.013, crit=1.476 ± 0.022 per rating point (14 rating = 1%, 20.664 per %), hit=1.980 ± 0.082 per rating point (10 rating = 1%, 19.803 per %), melee_haste=8.070 ± 1.095

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 66.9 attack_power points (7.00 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -1.12 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.2 attack_power points (2.85 DPS) | yes | Sunprism Pendant (20649, -0.79 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -1.59 DPS) [world]; Kaleidoscope Chain (13084, -1.71 DPS) [world_drop] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 41.6 attack_power points (4.35 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.7 attack_power points (4.77 DPS) | yes | Dark Phantom Cape (13122, -2.45 DPS, sim-verified) [world_drop]; Sergeant Major's Cape (16337, -2.51 DPS) [pvp]; Fel Cape (279269, -2.61 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 103.8 attack_power points (10.85 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -1.97 DPS) [pvp]; Timbermaw Tunic (252484, -2.14 DPS, sim-verified) [crafted]; Knight-Captain's Leather Chestpiece (227058, -3.07 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (219.6 DPS) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -3.18 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 48.7 attack_power points (5.09 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Timbermaw Brawlers (19049, -4.44 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 45.0 attack_power points (4.70 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Knight-Captain's Leather Belt (16398, -1.08 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.10 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 66.7 attack_power points (6.97 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Warbear Woolies (15065, -5.02 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 41.9 attack_power points (4.38 DPS) | yes | Marshal's Leather Footguards (231546, -0.44 DPS) [pvp]; Skulker's Leather Boots (252469, -0.90 DPS) [crafted]; Albino Crocscale Boots (17728, -0.94 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 39.8 attack_power points (4.16 DPS) | yes | Band of the Penitent (13217, -2.00 DPS) [quest]; Mark of Kern (2262, -2.07 DPS) [dungeon]; Assault Band (13095, -2.07 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 24.0 attack_power points (2.50 DPS) | yes | Mark of Kern (2262, -0.41 DPS) [dungeon]; Assault Band (13095, -0.41 DPS) [world_drop]; Band of the Penitent (13217, -0.69 DPS, sim-verified) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.06 DPS, sim-verified) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -10.26 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 595.8 attack_power points (62.29 DPS) | yes | Skilled Fighting Blade (12062, -0.75 DPS, sim-verified) [quest] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.10 DPS) [world_drop]; Eaglehorn Long Bow (13023, -0.82 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.83 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; main_hand: Heartseeker; off_hand: Scarlet Kris; ranged: Blackcrow

No-known-source sample (15 of 1495, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00532310101400000-00000000000000000-5320201310013211501)

Set DPS (verified): 559.1. Weights run: 2.8s. Verify run: 6.0s. 1495 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.034 ± 0.017, crit=1.773 ± 0.027 per rating point (14 rating = 1%, 24.823 per %), hit=2.949 ± 0.144 per rating point (10 rating = 1%, 29.487 per %), melee_haste=14.754 ± 1.752

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 83.8 attack_power points (16.52 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.76 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.9 attack_power points (6.30 DPS) | yes | Sunprism Pendant (20649, -1.49 DPS) [quest]; Kaleidoscope Chain (13084, -3.83 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -3.93 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 50.6 attack_power points (9.99 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.5 attack_power points (9.96 DPS) | yes | Cloak of Flames (3475, -4.15 DPS) [world_drop]; Dark Phantom Cape (13122, -4.56 DPS, sim-verified) [world_drop]; Sergeant Major's Cape (16337, -5.02 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 129.6 attack_power points (25.55 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -4.41 DPS) [pvp]; Timbermaw Tunic (252484, -6.67 DPS, sim-verified) [crafted]; Knight-Captain's Leather Chestpiece (227058, -8.14 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Leather Bracers (16394, +0.00 DPS) [pvp]; Marshal's Leather Armsplints (16460, +0.00 DPS) [pvp]; Deepfury Bracers (13120, -5.84 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 52.8 attack_power points (10.42 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Knight-Lieutenant's Leather Gauntlets (16396, -0.10 DPS) [pvp]; Timbermaw Brawlers (19049, -0.25 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 51.6 attack_power points (10.17 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Knight-Captain's Leather Belt (16398, -2.20 DPS) [pvp]; Skulker's Leather Waistguard (252474, -2.26 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (+12.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight's Leather Pants (220858, -12.37 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 51.6 attack_power points (10.17 DPS) | yes | Marshal's Leather Footguards (231546, -1.34 DPS) [pvp]; Hivethrasher's Boots (275615, -2.55 DPS) [crafted]; Albino Crocscale Boots (17728, -2.64 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 49.5 attack_power points (9.76 DPS) | yes | Band of the Penitent (13217, -4.86 DPS) [quest]; Ring of the Underwood (2951, -5.31 DPS) [world_drop]; Falcon's Hook (7552, -5.50 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.5 attack_power points (5.62 DPS) | yes | Band of the Penitent (13217, -0.72 DPS) [quest]; Ring of the Underwood (2951, -1.17 DPS) [world_drop]; Falcon's Hook (7552, -1.35 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -5.58 DPS, sim-verified) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Gut Ripper (2164, -26.34 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 599.0 attack_power points (118.12 DPS) | yes | Skilled Fighting Blade (12062, -3.53 DPS, sim-verified) [quest] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.89 DPS, sim-verified) [world_drop]; Eaglehorn Long Bow (13023, -3.01 DPS) [world_drop]; Guttbuster (13139, -3.16 DPS) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; main_hand: Heartseeker; off_hand: Scarlet Kris; ranged: Blackcrow

No-known-source sample (15 of 1495, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5320100000000000000)

Set DPS (verified): 37.9. Weights run: 1.8s. Verify run: 4.1s. 241 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.848 ± 0.013, crit=0.541 ± 0.012 per rating point (14 rating = 1%, 7.577 per %), hit=0.630 ± 0.030 per rating point (10 rating = 1%, 6.303 per %), melee_haste=3.650 ± 0.304

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.8 attack_power points (0.95 DPS) | yes | Defender's Leather Hood (252447, -0.47 DPS, sim-verified) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.7 attack_power points (0.24 DPS) | yes | Roadwatcher's Confidence (281265, -0.24 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.2 attack_power points (0.60 DPS) | yes | Slime-encrusted Pads (6461, -0.63 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.1 attack_power points (0.72 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.23 DPS) [crafted]; Bristlebark Cape (14571, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.9 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.11 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.2 attack_power points (0.60 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Spare Part Bindings (279875, -0.24 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.1 attack_power points (0.97 DPS) | yes | Brawler's Leather Gloves (252494, -0.24 DPS) [crafted]; Bristlebark Gloves (14572, -0.24 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.26 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.43 DPS) [crafted]; Murloc Scale Belt (5780, -0.67 DPS) [crafted]; Deviate Scale Belt (6468, -4.10 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.30 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (37.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.22 DPS, sim-verified) [world_drop] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.9 attack_power points (0.64 DPS) | yes | Demon Band (12054, -0.38 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; Loop of Sacrifice (281673, -0.45 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 5.5 attack_power points (0.36 DPS) | yes | Demon Band (12054, -0.10 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Loop of Sacrifice (281673, -0.16 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.7 attack_power points (14.76 DPS) | yes | Tail Spike (6448, -1.66 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.74 DPS) [crafted]; Goblin Screwdriver (1936, -3.31 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.8 attack_power points (14.70 DPS) | yes | Tail Spike (6448, -0.47 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.4 attack_power points (0.48 DPS) | yes | Fine Longbow (11304, -0.23 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.24 DPS) [crafted]; Light Bow (4576, -0.24 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Pyrewood Signet Ring; finger2: Bounty Hunter's Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 241, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade

### Band 30 (troll, 00000000000000000-00000000000000000-5320201310013000000)

Set DPS (verified): 58.8. Weights run: 2.0s. Verify run: 2.8s. 435 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.535 ± 0.011, crit=0.557 ± 0.012 per rating point (14 rating = 1%, 7.795 per %), hit=0.795 ± 0.031 per rating point (10 rating = 1%, 7.951 per %), melee_haste=3.809 ± 0.324

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.3 attack_power points (1.07 DPS) | yes | Defender's Leather Helm (252455, -0.23 DPS) [crafted]; Brawler's Leather Hood (252504, -0.27 DPS, sim-verified) [crafted]; Humbert's Helm (4724, -0.32 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.84 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; River Pride Choker (13087, -0.56 DPS) [world_drop]; Erudite's Amulet (277204, -0.62 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.9 attack_power points (1.53 DPS) | yes | Mantle of Thieves (2264, -0.51 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.64 DPS) [crafted]; Bristlebark Amice (14573, -0.67 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.7 attack_power points (0.96 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.26 DPS) [dungeon]; Wildhunter Cloak (16658, -0.26 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.5 attack_power points (1.50 DPS) | yes | Brawler's Leather Tunic (252508, -0.34 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.40 DPS) [quest]; Brawler's Leather Armor (252490, -0.40 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.2 attack_power points (0.92 DPS) | yes | Cultist's Armguards (270032, -0.22 DPS) [quest]; Barbaric Bracers (18948, -0.27 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.28 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 16.2 attack_power points (1.13 DPS) | yes | Heavy Earthen Gloves (7359, -0.01 DPS) [crafted]; Toughened Leather Gloves (4253, -0.07 DPS) [crafted]; Wolfclaw Gloves (1978, -0.14 DPS) [dungeon] |
| waist | Skulker's Leather Belt (252520) | Leatherworking [crafted] | 19.8 attack_power points (1.38 DPS) | yes | Prowler's Leather Belt (252459, -0.11 DPS) [crafted]; Deftkin Belt (16659, -0.12 DPS) [quest]; Blackened Defias Belt (10403, -0.13 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.81 DPS) | yes | Petrolspill Leggings (9509, -0.31 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.31 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.36 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 15.3 attack_power points (1.07 DPS) | yes | Insignia Boots (4055, -0.21 DPS) [world_drop]; Warsong Boots (16977, -0.21 DPS) [quest]; Brawler's Leather Boots (252439, -0.25 DPS, sim-verified) [crafted] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.6 attack_power points (0.88 DPS) | yes | Pyrewood Signet Ring (277210, -0.23 DPS) [quest]; Ring of Precision (1491, -0.24 DPS) [dungeon]; Monkey Ring (6748, -0.34 DPS) [quest] |
| finger2 | Band of the Fist (17694) | Allegiance to the Old Gods [quest] | 10.1 attack_power points (0.71 DPS) | yes | Ring of Precision (1491, -0.06 DPS) [dungeon]; Monkey Ring (6748, -0.17 DPS) [quest]; Pyrewood Signet Ring (277210, -0.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-verified (58.8 DPS) | yes | Vendetta (776, -0.52 DPS) [dungeon]; Thornspike (6681, -0.75 DPS, sim-verified) [dungeon]; Torturing Poker (7682, -0.80 DPS) [dungeon] |
| off_hand | Claw of the Shadowmancer (2912) (or Toxic Revenger (9453)) | World drop [world_drop] | 287.4 attack_power points (20.04 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 9.2 attack_power points (0.64 DPS) | yes | Gravespike Repeater (273033, -0.08 DPS) [world]; Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Silver Star (3463, -0.21 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Skulker's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Thunderbrow Ring; finger2: Band of the Fist; main_hand: Toxic Revenger; off_hand: Claw of the Shadowmancer; ranged: Golemsight Long Gun

No-known-source sample (15 of 435, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 00000000000000000-00000000000000000-5320201310013211501)

Set DPS (verified): 94.7. Weights run: 2.1s. Verify run: 3.2s. 582 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.356 ± 0.009, crit=0.462 ± 0.012 per rating point (14 rating = 1%, 6.464 per %), hit=0.838 ± 0.032 per rating point (10 rating = 1%, 8.381 per %), melee_haste=3.977 ± 0.382

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.9 attack_power points (1.92 DPS) | yes | Hawkeye's Helm (14591, -0.60 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.61 DPS) [world_drop]; Nightscape Headband (8176, -0.71 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.89 DPS) | yes | Ethereal Talisman (4430, -0.12 DPS) [quest]; Kaleidoscope Chain (13084, -0.19 DPS) [world_drop]; River Pride Choker (13087, -0.59 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 19.9 attack_power points (1.47 DPS) | yes | Flintrock Shoulders (7755, -0.10 DPS) [dungeon]; Nightscape Shoulders (8192, -0.37 DPS) [crafted]; Mantle of Thieves (2264, -0.47 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.5 attack_power points (0.92 DPS) | yes | Imperial Cloak (6432, -0.12 DPS) [dungeon]; Parachute Cloak (10518, -0.12 DPS) [crafted]; Tigerstrike Mantle (13108, -0.12 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 23.1 attack_power points (1.71 DPS) | yes | Dusky Leather Armor (7374, -0.30 DPS) [crafted]; Nightscape Tunic (8175, -0.31 DPS, sim-verified) [crafted]; Barbaric Harness (5739, -0.38 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Hawkeye's Bracers (14590, -0.53 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.68 DPS) [dungeon]; Dusky Bracers (7378, -0.68 DPS) [crafted] |
| hands | Dog Training Gloves (7756) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 24.0 attack_power points (1.78 DPS) | yes | Skulker's Leather Gloves (252525, -0.11 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.13 DPS) [crafted]; Imperial Leather Gloves (4063, -0.21 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 21.2 attack_power points (1.57 DPS) | yes | Skulker's Leather Belt (252520, -0.22 DPS) [crafted]; Blackened Defias Belt (10403, -0.24 DPS) [dungeon]; Scorpashi Sash (14652, -0.27 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 28.5 attack_power points (2.11 DPS) | yes | Triprunner Dungarees (9624, -0.08 DPS) [quest]; Ferine Leggings (6690, -0.18 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.69 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.9 attack_power points (1.62 DPS) | yes | Prowler's Leather Shoes (252465, -0.11 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Excelsior Boots (4109, -0.20 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Falcon's Hook (7552, -0.35 DPS) [dungeon]; Thunderbrow Ring (13097, -0.59 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.48 DPS) | yes | Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Falcon's Hook (7552, -0.35 DPS) [dungeon]; Thunderbrow Ring (13097, -0.59 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | 474.4 attack_power points (35.09 DPS) | yes | The Ziggler (8006, -7.68 DPS) [world_drop]; Sacrificial Kris (3187, -8.46 DPS) [dungeon]; Tok'kar's Murloc Shanker (9680, -8.60 DPS) [quest] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 406.0 attack_power points (30.03 DPS) | yes | The Ziggler (8006, -1.21 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (94.7 DPS) | yes | Monolithic Bow (9426, -0.29 DPS) [dungeon]; Swiftwind (13038, -0.33 DPS) [world_drop]; Master Hunter's Rifle (17687, -0.34 DPS) [quest] |

**New at 40:** head: White Bandit Mask; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Dog Training Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 582, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 00532000000000000-00000000000000000-5320201310013211501)

Set DPS (verified): 138.3. Weights run: 2.7s. Verify run: 3.6s. 733 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.519 ± 0.008, crit=0.868 ± 0.012 per rating point (14 rating = 1%, 12.152 per %), hit=1.451 ± 0.056 per rating point (10 rating = 1%, 14.512 per %), melee_haste=5.583 ± 0.793

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 27.7 attack_power points (2.39 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.16 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.7 attack_power points (2.14 DPS) | yes | Woven Ivy Necklace (19159, -0.54 DPS, sim-verified) [quest]; Archmagister's Faceted Pendant (273049, -1.10 DPS) [world]; Ethereal Talisman (4430, -1.18 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 25.7 attack_power points (2.22 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.09 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 22.8 attack_power points (1.97 DPS) | yes | Blackflame Cape (13109, -0.48 DPS, sim-verified) [world_drop]; Serpentskin Cloak (8259, -0.52 DPS) [world_drop]; Pridelord Cape (14673, -0.53 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 10 · Stone Guard · Horde [vendor] | 44.7 attack_power points (3.86 DPS) | yes | Blazewind Breastplate (11193, -0.58 DPS) [quest]; Warbear Harness (15064, -1.22 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -1.50 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 26.8 attack_power points (2.31 DPS) | yes | Pridelord Bands (14672, -0.57 DPS) [world_drop]; Branded Leather Bracers (19508, -0.59 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.66 DPS, sim-verified) [crafted] |
| hands | Skulker's Leather Gauntlets (252548) | Leatherworking [crafted] | 30.3 attack_power points (2.62 DPS) | yes | Serpentskin Gloves (8260, -0.22 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.22 DPS) [crafted]; First Sergeant's Leather Gauntlets (220857, -0.36 DPS) [vendor] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 33.3 attack_power points (2.87 DPS) | yes | Prowler's Leather Waistguard (252473, -0.09 DPS) [crafted]; Sagebrush Girdle (17778, -0.91 DPS) [quest]; Ogron's Sash (13117, -0.92 DPS) [world_drop] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 44.7 attack_power points (3.86 DPS) | yes | Basilisk Hide Pants (1718, -1.10 DPS) [world_drop]; Triprunner Dungarees (9624, -1.24 DPS) [quest]; Serpentskin Leggings (8262, -1.33 DPS, sim-verified) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 30.7 attack_power points (2.66 DPS) | yes | Albino Crocscale Boots (17728, -0.03 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.09 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.61 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.5 attack_power points (2.98 DPS) | yes | Masons Fraternity Ring (9533, -1.15 DPS) [quest]; Mark of Kern (2262, -1.25 DPS) [dungeon]; Assault Band (13095, -1.25 DPS) [world_drop] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.07 DPS) | yes | Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Mark of Kern (2262, -0.35 DPS) [dungeon]; Assault Band (13095, -0.35 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (138.3 DPS) | yes | Frozen Heart of the Mountain (249469, -3.38 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (138.3 DPS) | yes | Frozen Heart of the Mountain (249469, -0.48 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (138.3 DPS) | yes | Widowmaker (4091, -1.14 DPS) [world_drop]; Gahz'rilla Fang (9467, -2.02 DPS) [dungeon]; Julie's Dagger (6660, -7.20 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.1 attack_power points (43.48 DPS) | yes | - |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (138.3 DPS) | yes | Guttbuster (13139, -0.53 DPS) [world_drop]; Houndmaster's Bow (11628, -0.54 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.58 DPS, sim-verified) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Skulker's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 733, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 00532310101400000-00000000000000000-5320201310013211501)

Set DPS (verified): 221.7. Weights run: 2.6s. Verify run: 5.8s. 1494 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.711 ± 0.013, crit=1.476 ± 0.022 per rating point (14 rating = 1%, 20.664 per %), hit=1.980 ± 0.082 per rating point (10 rating = 1%, 19.803 per %), melee_haste=8.070 ± 1.095

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 66.9 attack_power points (7.00 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -0.70 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 27.2 attack_power points (2.85 DPS) | yes | Woven Ivy Necklace (19159, -0.69 DPS, sim-verified) [quest]; Sunprism Pendant (20649, -0.70 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.59 DPS) [world] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 41.6 attack_power points (4.35 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 45.7 attack_power points (4.77 DPS) | yes | Dark Phantom Cape (13122, -2.43 DPS, sim-verified) [world_drop]; Fel Cape (279269, -2.61 DPS) [crafted]; Blackflame Cape (13109, -2.63 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 103.8 attack_power points (10.85 DPS) | yes | Timbermaw Tunic (252484, -1.91 DPS, sim-verified) [crafted]; Warlord's Leather Breastplate (231549, -1.97 DPS) [pvp]; Legionnaire's Leather Chestpiece (227060, -3.07 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (221.7 DPS) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.34 DPS) [crafted]; Deepfury Bracers (13120, -3.15 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 48.7 attack_power points (5.09 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Timbermaw Brawlers (19049, -3.78 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 45.0 attack_power points (4.70 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Legionnaire's Leather Girdle (16500, -1.08 DPS) [pvp]; Skulker's Leather Waistguard (252474, -1.09 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 66.7 attack_power points (6.97 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Warbear Woolies (15065, -4.31 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 41.9 attack_power points (4.38 DPS) | yes | General's Leather Treads (231552, -0.44 DPS) [pvp]; Skulker's Leather Boots (252469, -0.90 DPS) [crafted]; Albino Crocscale Boots (17728, -0.92 DPS, sim-verified) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 39.8 attack_power points (4.16 DPS) | yes | Masons Fraternity Ring (9533, -1.66 DPS) [quest]; Band of the Penitent (13217, -2.00 DPS) [quest]; Mark of Kern (2262, -2.07 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.51 DPS) | yes | Masons Fraternity Ring (9533, -0.00 DPS) [quest]; Band of the Penitent (13217, -0.35 DPS) [quest]; Mark of Kern (2262, -0.42 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.72 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -10.70 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 595.8 attack_power points (62.29 DPS) | yes | The Shadowfoot Stabber (24222, -1.96 DPS, sim-verified) [world_drop] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -0.10 DPS) [world_drop]; Eaglehorn Long Bow (13023, -0.82 DPS) [world_drop]; Skull Splitting Crossbow (13039, -0.83 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; trinket2: Frozen Heart of the Mountain; main_hand: Heartseeker; off_hand: Scarlet Kris; ranged: Blackcrow

No-known-source sample (15 of 1494, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 00532310101400000-00000000000000000-5320201310013211501)

Set DPS (verified): 564.6. Weights run: 2.8s. Verify run: 6.2s. 1494 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.034 ± 0.017, crit=1.773 ± 0.027 per rating point (14 rating = 1%, 24.823 per %), hit=2.949 ± 0.144 per rating point (10 rating = 1%, 29.487 per %), melee_haste=14.754 ± 1.752

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 83.8 attack_power points (16.52 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, -0.76 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 31.9 attack_power points (6.30 DPS) | yes | Woven Ivy Necklace (19159, -1.39 DPS) [quest]; Sunprism Pendant (20649, -1.49 DPS) [quest]; Ethereal Talisman (4430, -3.61 DPS) [quest] |
| shoulder | Shadowcraft Spaulders (16708) | Stratholme: Cannon Master Willey [dungeon] | 50.6 attack_power points (9.99 DPS) | yes | Dark Warder's Pauldrons (22241, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 50.5 attack_power points (9.96 DPS) | yes | Cloak of Flames (3475, -4.15 DPS) [world_drop]; Dark Phantom Cape (13122, -4.55 DPS, sim-verified) [world_drop]; Fel Cape (279269, -5.07 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 129.6 attack_power points (25.55 DPS) | yes | Warlord's Leather Breastplate (231549, -4.41 DPS) [pvp]; Timbermaw Tunic (252484, -4.65 DPS, sim-verified) [crafted]; Legionnaire's Leather Chestpiece (227060, -8.14 DPS) [pvp] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Leather Armsplints (16559, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.89 DPS) [crafted]; Deepfury Bracers (13120, -5.83 DPS, sim-verified) [world_drop] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | 52.8 attack_power points (10.42 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Blood Guard's Leather Vices (16499, -0.10 DPS) [pvp]; Timbermaw Brawlers (19049, -0.25 DPS) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 51.6 attack_power points (10.17 DPS) | yes | General's Leather Girdle (16557, +0.00 DPS) [pvp]; Legionnaire's Leather Girdle (16500, -2.20 DPS) [pvp]; Skulker's Leather Waistguard (252474, -2.25 DPS, sim-verified) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | sim-verified (+13.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Stone Guard's Leather Pants (220859, -13.91 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 51.6 attack_power points (10.17 DPS) | yes | General's Leather Treads (231552, -1.34 DPS) [pvp]; Albino Crocscale Boots (17728, -2.35 DPS, sim-verified) [dungeon]; Hivethrasher's Boots (275615, -2.55 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 49.5 attack_power points (9.76 DPS) | yes | Band of the Penitent (13217, -4.86 DPS) [quest]; White Bone Band (11862, -5.03 DPS) [quest]; Ring of the Underwood (2951, -5.31 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 28.5 attack_power points (5.62 DPS) | yes | Band of the Penitent (13217, -0.72 DPS) [quest]; White Bone Band (11862, -0.88 DPS) [quest]; Ring of the Underwood (2951, -1.17 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+12.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Gut Ripper (2164, -24.80 DPS, sim-verified) [world_drop] |
| off_hand | Scarlet Kris (5267) | World drop [world_drop] | 599.0 attack_power points (118.12 DPS) | yes | The Shadowfoot Stabber (24222, -4.16 DPS, sim-verified) [world_drop] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.40 DPS) [world_drop]; Eaglehorn Long Bow (13023, -3.01 DPS) [world_drop]; Guttbuster (13139, -3.16 DPS) [world_drop] |

**New at 60:** head: Mask of the Unforgiven; shoulder: Shadowcraft Spaulders; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Shadowcraft Bracers; hands: Devilsaur Gauntlets; waist: Might of the Timbermaw; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger2: Masons Fraternity Ring; trinket2: Frozen Heart of the Mountain; main_hand: Heartseeker; off_hand: Scarlet Kris; ranged: Blackcrow

No-known-source sample (15 of 1494, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

