# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 53.1. Weights run: 5.7s. Verify run: 9.0s. 251 eligible items had no known source.

Race sweep: night-elf 53.1 ± 0.0, high-order-skyborne 52.6 ± 0.0 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.907 ± 0.010, crit=0.568 ± 0.009 per rating point (14 rating = 1%, 7.954 per %), hit=0.654 ± 0.036 per rating point (10 rating = 1%, 6.545 per %), melee_haste=2.420 ± 0.455

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.4 attack_power points (0.79 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.21 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.5 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.12 DPS) [crafted]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.40 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 11.7 attack_power points (0.81 DPS) | yes | Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Forest Leather Bracers (3202, -0.15 DPS) [world_drop]; Wolf Bracers (4794, -0.28 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.55 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.24 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Murloc Scale Belt (5780, -0.57 DPS) [crafted]; Brawler's Leather Belt (252428, -3.22 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.65 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.3 attack_power points (1.47 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Blackened Defias Boots (10402, -0.68 DPS) [dungeon]; Footpads of the Fang (10411, -0.68 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.6 attack_power points (1.08 DPS) | yes | Lavishly Jeweled Ring (1156, -0.82 DPS) [dungeon]; Ring of the Moon (12052, -0.94 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.55 DPS) | yes | Lavishly Jeweled Ring (1156, -0.28 DPS, sim-verified) [dungeon]; Ring of the Moon (12052, -0.41 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.6 attack_power points (21.14 DPS) | yes | Twisted Chanter's Staff (890, -1.51 DPS, sim-verified) [world_drop]; Rakzur Club (12983, -2.89 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -3.29 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 113.3. Weights run: 6.0s. Verify run: 7.3s. 424 eligible items had no known source.

Race sweep: night-elf 113.3 ± 0.1, high-order-skyborne 112.1 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.869 ± 0.011, crit=0.697 ± 0.011 per rating point (14 rating = 1%, 9.763 per %), hit=0.744 ± 0.046 per rating point (10 rating = 1%, 7.439 per %), melee_haste=2.752 ± 0.611

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (2.62 DPS) | yes | Brawler's Leather Helm (252512, -0.59 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.65 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.65 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.5 attack_power points (1.69 DPS) | yes | Sentinel's Medallion (19541, -0.06 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.38 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.6 attack_power points (3.33 DPS) | yes | Barbaric Shoulders (5964, -1.24 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -1.30 DPS) [dungeon]; Bristlebark Amice (14573, -1.46 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16315, -0.38 DPS, sim-verified) [pvp]; Tigerstrike Mantle (13108, -0.45 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.86 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 26.9 attack_power points (2.94 DPS) | yes | Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.41 DPS) [crafted]; Brawler's Leather Armor (252490, -0.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.2 attack_power points (2.10 DPS) | yes | Barbaric Bracers (18948, -0.41 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.63 DPS) [quest]; Bands of Serra'kis (6902, -0.79 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.2 attack_power points (2.75 DPS) | yes | Toughened Leather Gloves (4253, -0.22 DPS) [crafted]; Wolfclaw Gloves (1978, -0.44 DPS) [dungeon]; Gloves of the Fang (10413, -0.65 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.2 attack_power points (3.19 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.57 DPS) [rep]; Highlander's Leather Girdle (20117, -0.57 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 30.8 attack_power points (3.36 DPS) | yes | Brawler's Leather Pants (252500, -0.22 DPS) [crafted]; Trapper's Leather Pants (252501, -0.22 DPS) [crafted]; Barbaric Leggings (5963, -0.41 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.9 attack_power points (2.29 DPS) | yes | Brawler's Leather Boots (252439, -0.18 DPS) [crafted]; Insignia Boots (4055, -0.65 DPS) [world_drop]; Disjointed Shoes (277226, -0.98 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 23.2 attack_power points (2.53 DPS) | yes | Seal of Wrynn (2933, -1.27 DPS) [quest]; Ring of Precision (1491, -1.31 DPS) [dungeon]; Tiger Band (6749, -1.44 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.6 attack_power points (2.36 DPS) | yes | Ring of Precision (1491, -1.13 DPS) [dungeon]; Seal of Wrynn (2933, -1.14 DPS, sim-verified) [quest]; Tiger Band (6749, -1.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Viscous Hammer (13045, -2.34 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -4.66 DPS) [dungeon]; Gnarled Ash Staff (791, -9.67 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 178.3. Weights run: 7.5s. Verify run: 16.8s. 607 eligible items had no known source.

Race sweep: night-elf 178.3 ± 0.1, high-order-skyborne 176.2 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.064 ± 0.024, crit=1.037 ± 0.030 per rating point (14 rating = 1%, 14.516 per %), hit=0.806 ± 0.072 per rating point (10 rating = 1%, 8.057 per %), melee_haste=4.542 ± 0.966

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.7 attack_power points (5.87 DPS) | yes | Hawkeye's Helm (14591, -1.87 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -2.35 DPS) [world_drop]; Nightscape Headband (8176, -2.62 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 22.7 attack_power points (2.98 DPS) | yes | Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.14 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.41 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.7 attack_power points (4.30 DPS) | yes | Flintrock Shoulders (7755, -0.27 DPS) [dungeon]; Nightscape Shoulders (8192, -1.31 DPS) [crafted]; Fleshhide Shoulders (10774, -1.36 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.78 DPS, sim-verified) [world_drop]; Yeti Fur Cloak (2805, -0.79 DPS) [quest]; Parachute Cloak (10518, -1.03 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 55.2 attack_power points (7.26 DPS) | yes | Wolffear Harness (13110, -2.63 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -3.19 DPS) [crafted]; Barbaric Harness (5739, -3.25 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (2.68 DPS) | yes | Branded Leather Bracers (19508, -0.05 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.51 DPS) [dungeon]; Dusky Bracers (7378, -0.51 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Prowler's Leather Gloves (252524, -2.89 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.6 attack_power points (4.81 DPS) | yes | Skulker's Leather Belt (252520, -0.74 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -0.81 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.86 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 43.4 attack_power points (5.70 DPS) | yes | Triprunner Dungarees (9624, -0.94 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -1.42 DPS) [crafted]; Brawler's Leather Pants (252500, -1.68 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 36.7 attack_power points (4.82 DPS) | yes | Prowler's Leather Shoes (252465, -0.03 DPS) [crafted]; Imperial Leather Boots (6431, -0.53 DPS) [dungeon]; Excelsior Boots (4109, -0.54 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 32.5 attack_power points (4.27 DPS) | yes | Falcon's Hook (7552, -1.04 DPS) [dungeon]; Thunderbrow Ring (13097, -1.36 DPS) [world_drop]; Mark of Kern (2262, -1.64 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.6 attack_power points (3.24 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Mark of Kern (2262, -0.61 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Thornstone Sledgehammer (1722) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.62 DPS, sim-verified) [dungeon]; The Jackhammer (9423, -0.82 DPS) [dungeon]; Ironshod Bludgeon (9408, -3.13 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Thornstone Sledgehammer

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0500220000000000-50212032123012123001-0100000000000000)

Set DPS (verified): 267.7. Weights run: 7.9s. Verify run: 19.5s. 817 eligible items had no known source.

Race sweep: night-elf 267.7 ± 0.2, high-order-skyborne 264.5 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.040 ± 0.002, agility=2.080 ± 0.029, crit=1.258 ± 0.038 per rating point (14 rating = 1%, 17.606 per %), hit=1.005 ± 0.090 per rating point (10 rating = 1%, 10.052 per %), melee_haste=not significant (4.533 ± 1.221)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.3 attack_power points (6.97 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.67 DPS) [vendor]; White Bandit Mask (10008, -1.23 DPS, sim-verified) [crafted]; Helm of Fire (8348, -1.87 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (5.37 DPS) | yes | Sentinel's Medallion (19540, -2.07 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 41.2 attack_power points (5.95 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.01 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.81 DPS) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Phantom Cape (13122, -1.42 DPS, sim-verified) [world_drop]; Blisterbane Wrap (12552, -1.47 DPS) [dungeon]; Blackmetal Cape (9512, -2.37 DPS) [dungeon] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 68.0 attack_power points (9.81 DPS) | yes | Warbear Harness (15064, -1.10 DPS, sim-verified) [crafted]; Mixologist's Tunic (12793, -1.51 DPS) [dungeon]; Quillward Harness (10583, -1.75 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.4 attack_power points (5.68 DPS) | yes | Skulker's Leather Bracers (252540, -0.92 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.93 DPS) [crafted]; Pridelord Bands (14672, -1.21 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -2.79 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 53.6 attack_power points (7.73 DPS) | yes | Prowler's Leather Waistguard (252473, -0.01 DPS) [crafted]; Girdle of Beastial Fury (11686, -1.05 DPS) [dungeon]; Highlander's Leather Girdle (20115, -2.31 DPS) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 61.8 attack_power points (8.92 DPS) | yes | Serpentskin Leggings (8262, -0.29 DPS) [world_drop]; Knight's Leather Pants (220858, -2.33 DPS) [vendor]; Basilisk Hide Pants (1718, -2.61 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 49.5 attack_power points (7.14 DPS) | yes | Prowler's Leather Boots (252468, -0.01 DPS) [crafted]; Shadefiend Boots (11675, -0.89 DPS) [dungeon]; Albino Crocscale Boots (17728, -1.14 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 39.1 attack_power points (5.64 DPS) | yes | Masons Fraternity Ring (9533, -1.44 DPS) [quest]; Ring of the Underwood (2951, -2.05 DPS) [world_drop]; Falcon's Hook (7552, -2.06 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 30.1 attack_power points (4.34 DPS) | yes | Masons Fraternity Ring (9533, -0.13 DPS) [quest]; Ring of the Underwood (2951, -0.75 DPS) [world_drop]; Falcon's Hook (7552, -0.75 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.94 DPS, sim-verified) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.38 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -3.67 DPS, sim-verified) [world_drop]; Dark Iron Pulverizer (11608, -6.04 DPS) [crafted]; Soulkeeper (1607, -7.76 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 817, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0500220000000000-50212032123012123001-0550100000000000)

Set DPS (verified): 457.6. Weights run: 7.7s. Verify run: 31.8s. 1988 eligible items had no known source.

Race sweep: night-elf 457.6 ± 0.3, high-order-skyborne 453.9 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.211 ± 0.039, crit=1.730 ± 0.055 per rating point (14 rating = 1%, 24.218 per %), hit=1.281 ± 0.131 per rating point (10 rating = 1%, 12.814 per %), melee_haste=not significant (3.883 ± 1.720)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 75.7 attack_power points (13.33 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 50.2 attack_power points (8.84 DPS) | yes | Skibi's Pendant (13089, -1.84 DPS) [world_drop]; Sentinel's Medallion (19538, -3.00 DPS) [rep]; Imperial Jewel (11933, -3.21 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 83.9 attack_power points (14.77 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -0.42 DPS) [vendor]; Highlander's Leather Shoulders (20059, -2.43 DPS, sim-verified) [rep]; Lieutenant Commander's Dragonhide Shoulders (227172, -3.53 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.43 DPS) [rep]; Blackveil Cape (11626, -1.59 DPS) [dungeon]; Howler's Furs (272414, -2.17 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 136.6 attack_power points (24.05 DPS) | yes | Dawn Armor (252483, -3.23 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -4.34 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -4.66 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 66.2 attack_power points (11.66 DPS) | yes | Deepfury Bracers (13120, -4.27 DPS) [world_drop]; Wristguards of Renown (22204, -4.28 DPS) [dungeon]; Skulker's Leather Bracers (252540, -5.44 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, -0.80 DPS) [vendor]; Timbermaw Brawlers (19049, -1.19 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -3.59 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 83.9 attack_power points (14.77 DPS) | yes | Might of the Timbermaw (19044, -1.18 DPS) [crafted]; Cloudrunner Girdle (13252, -4.28 DPS) [dungeon]; Highlander's Leather Girdle (20045, -4.51 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 108.1 attack_power points (19.04 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Warbear Woolies (15065, -1.19 DPS) [crafted]; Knight-Captain's Dragonhide Leggings (227178, -3.20 DPS) [vendor] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 57.4 attack_power points (10.11 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.79 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.28 DPS) [vendor]; Magma Forged Band (22255, -1.60 DPS) [dungeon]; Naglering (11669, -8.15 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.22 DPS) [vendor]; Magma Forged Band (22255, -1.55 DPS) [dungeon]; Myrmidon's Signet (2246, -1.94 DPS) [world_drop] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Second Wind (11819, -4.16 DPS, sim-verified) [dungeon] |
| trinket2 | Heart of Wyrmthalak (22321) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -11.78 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Ferocity (22397, -11.50 DPS, sim-verified) [dungeon] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Blackhand's Breadth; trinket2: Heart of Wyrmthalak; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0500220000000000-50230032023012125001-0550000000000000)

Set DPS (verified): 1100.7. Weights run: 7.6s. Verify run: 56.2s. 1988 eligible items had no known source.

Race sweep: night-elf 1100.7 ± 0.7, high-order-skyborne 1093.4 ± 0.7 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.489 ± 0.035, crit=1.800 ± 0.047 per rating point (14 rating = 1%, 25.198 per %), hit=1.855 ± 0.157 per rating point (10 rating = 1%, 18.553 per %), melee_haste=7.750 ± 1.917

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 83.4 attack_power points (25.49 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 51.2 attack_power points (15.65 DPS) | yes | Skibi's Pendant (13089, -2.06 DPS) [world_drop]; Sentinel's Medallion (19538, -4.24 DPS) [rep]; Imperial Jewel (11933, -5.87 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 92.4 attack_power points (28.25 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -1.06 DPS) [vendor]; Highlander's Leather Shoulders (20059, -5.67 DPS, sim-verified) [rep]; Lieutenant Commander's Dragonhide Shoulders (227172, -7.10 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -2.44 DPS) [dungeon]; Howler's Furs (272414, -3.30 DPS) [vendor]; Cloak of the Honor Guard (20073, -3.33 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 149.9 attack_power points (45.82 DPS) | yes | Dawn Armor (252483, -3.00 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -6.99 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -8.96 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 73.9 attack_power points (22.60 DPS) | yes | Deepfury Bracers (13120, -8.22 DPS) [world_drop]; Wristguards of Renown (22204, -8.33 DPS) [dungeon]; Skulker's Leather Bracers (252540, -10.57 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, -2.28 DPS) [vendor]; Timbermaw Brawlers (19049, -2.50 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -5.79 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 93.6 attack_power points (28.62 DPS) | yes | Might of the Timbermaw (19044, -2.43 DPS) [crafted]; Cloudrunner Girdle (13252, -8.35 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -9.09 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 117.6 attack_power points (35.96 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -4.57 DPS) [vendor]; Warbear Woolies (15065, -4.61 DPS, sim-verified) [crafted] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 64.4 attack_power points (19.68 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -1.65 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.33 DPS) [vendor]; Magma Forged Band (22255, -3.34 DPS) [dungeon]; Naglering (11669, -19.54 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.92 DPS) [vendor]; Magma Forged Band (22255, -2.94 DPS) [dungeon]; Myrmidon's Signet (2246, -3.78 DPS) [world_drop] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -6.31 DPS, sim-verified) [crafted] |
| trinket2 | Smolderweb's Eye (13213) | Blackrock Spire: Mother Smolderweb [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.57 DPS) [crafted]; Hand of Justice (11815, -2.00 DPS) [dungeon]; Cannonball Runner (13382, -2.84 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Slavedriver's Cane (13372, -22.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Ferocity (22397, -18.69 DPS, sim-verified) [dungeon] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Blackhand's Breadth; trinket2: Smolderweb's Eye; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 52.2. Weights run: 5.7s. Verify run: 8.8s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 52.5 ± 0.0, tauren 52.2 ± 0.0 (dps); windshaper-skyborne beats the guide race tauren by 0.3 (0.59%) ± 0.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.907 ± 0.010, crit=0.568 ± 0.009 per rating point (14 rating = 1%, 7.954 per %), hit=0.654 ± 0.036 per rating point (10 rating = 1%, 6.545 per %), melee_haste=2.420 ± 0.455

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.4 attack_power points (0.79 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.5 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Leather Cloak (2316, -0.12 DPS) [crafted]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.32 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 9.7 attack_power points (0.67 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.28 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.46 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.24 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Murloc Scale Belt (5780, -0.57 DPS) [crafted]; Brawler's Leather Belt (252428, -3.12 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.59 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.3 attack_power points (1.47 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Blackened Defias Boots (10402, -0.68 DPS) [dungeon]; Footpads of the Fang (10411, -0.68 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.6 attack_power points (1.08 DPS) | yes | Loop of Sacrifice (281673, -0.67 DPS) [quest]; Bounty Hunter's Ring (5351, -0.69 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.82 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.55 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Bounty Hunter's Ring (5351, -0.16 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.6 attack_power points (21.14 DPS) | yes | Hammerbone (270018, -0.56 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -1.35 DPS) [world_drop]; Rakzur Club (12983, -2.89 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 111.6. Weights run: 6.0s. Verify run: 7.2s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 112.1 ± 0.1, tauren 111.6 ± 0.1 (dps); windshaper-skyborne beats the guide race tauren by 0.5 (0.45%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.869 ± 0.011, crit=0.697 ± 0.011 per rating point (14 rating = 1%, 9.763 per %), hit=0.744 ± 0.046 per rating point (10 rating = 1%, 7.439 per %), melee_haste=2.752 ± 0.611

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (2.62 DPS) | yes | Brawler's Leather Helm (252512, -0.58 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.65 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.65 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.5 attack_power points (1.69 DPS) | yes | Scout's Medallion (19537, -0.06 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.38 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 30.6 attack_power points (3.33 DPS) | yes | Barbaric Shoulders (5964, -1.25 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -1.30 DPS) [dungeon]; Bristlebark Amice (14573, -1.46 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.46 DPS, sim-verified) [world_drop]; Glowing Lizardscale Cloak (6449, -0.86 DPS) [dungeon]; Swiftrunner Cape (6745, -0.86 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 26.9 attack_power points (2.94 DPS) | yes | Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Defender's Leather Tunic (252450, -0.41 DPS) [crafted]; Brawler's Leather Armor (252490, -0.42 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.2 attack_power points (2.10 DPS) | yes | Barbaric Bracers (18948, -0.42 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.79 DPS) [dungeon]; Insignia Bracers (6410, -0.87 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.2 attack_power points (2.75 DPS) | yes | Toughened Leather Gloves (4253, -0.22 DPS) [crafted]; Wolfclaw Gloves (1978, -0.44 DPS) [dungeon]; Gloves of the Fang (10413, -0.65 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.2 attack_power points (3.19 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.57 DPS) [rep]; Defiler's Leather Girdle (20191, -0.57 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 30.8 attack_power points (3.36 DPS) | yes | Brawler's Leather Pants (252500, -0.22 DPS) [crafted]; Trapper's Leather Pants (252501, -0.22 DPS) [crafted]; Barbaric Leggings (5963, -0.41 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 20.9 attack_power points (2.29 DPS) | yes | Brawler's Leather Boots (252439, -0.18 DPS) [crafted]; Stomping Boots (3741, -0.60 DPS) [quest]; Insignia Boots (4055, -0.65 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 23.2 attack_power points (2.53 DPS) | yes | Band of the Fist (17694, -0.84 DPS) [quest]; Ring of Precision (1491, -1.31 DPS) [dungeon]; Tiger Band (6749, -1.44 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 21.6 attack_power points (2.36 DPS) | yes | Band of the Fist (17694, -0.68 DPS, sim-verified) [quest]; Ring of Precision (1491, -1.13 DPS) [dungeon]; Tiger Band (6749, -1.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Viscous Hammer (13045, -2.08 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -4.66 DPS) [dungeon]; Gnarled Ash Staff (791, -9.67 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 174.1. Weights run: 7.5s. Verify run: 19.2s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 175.1 ± 0.1, tauren 174.1 ± 0.1 (dps); windshaper-skyborne beats the guide race tauren by 0.9 (0.53%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.064 ± 0.024, crit=1.037 ± 0.030 per rating point (14 rating = 1%, 14.516 per %), hit=0.806 ± 0.072 per rating point (10 rating = 1%, 8.057 per %), melee_haste=4.542 ± 0.966

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.7 attack_power points (5.87 DPS) | yes | Hawkeye's Helm (14591, -1.82 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -2.35 DPS) [world_drop]; Nightscape Headband (8176, -2.62 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 22.7 attack_power points (2.98 DPS) | yes | Ethereal Talisman (4430, -0.58 DPS) [quest]; Kaleidoscope Chain (13084, -0.85 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.14 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.7 attack_power points (4.30 DPS) | yes | Flintrock Shoulders (7755, -0.27 DPS) [dungeon]; Nightscape Shoulders (8192, -1.31 DPS) [crafted]; Fleshhide Shoulders (10774, -1.36 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.52 DPS) [dungeon]; Parachute Cloak (10518, -0.52 DPS) [crafted]; Tigerstrike Mantle (13108, -0.52 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 55.2 attack_power points (7.26 DPS) | yes | Wolffear Harness (13110, -2.48 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -3.19 DPS) [crafted]; Barbaric Harness (5739, -3.25 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (2.68 DPS) | yes | Branded Leather Bracers (19508, -0.05 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.51 DPS) [dungeon]; Dusky Bracers (7378, -0.51 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Prowler's Leather Gloves (252524, -1.97 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.6 attack_power points (4.81 DPS) | yes | Skulker's Leather Belt (252520, -0.79 DPS) [crafted]; Prowler's Leather Belt (252459, -0.81 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.86 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 43.4 attack_power points (5.70 DPS) | yes | Triprunner Dungarees (9624, -0.03 DPS) [quest]; Brawler's Leather Legguards (252516, -1.42 DPS) [crafted]; Brawler's Leather Pants (252500, -1.68 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Skulker's Leather Shoes (252531, -0.36 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.49 DPS) [dungeon]; Excelsior Boots (4109, -0.51 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 32.5 attack_power points (4.27 DPS) | yes | Falcon's Hook (7552, -1.04 DPS) [dungeon]; Thunderbrow Ring (13097, -1.36 DPS) [world_drop]; Mark of Kern (2262, -1.64 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 24.6 attack_power points (3.24 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.32 DPS) [world_drop]; Mark of Kern (2262, -0.61 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Thornstone Sledgehammer (1722) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.14 DPS) [dungeon]; The Jackhammer (9423, -0.82 DPS) [dungeon]; Ironshod Bludgeon (9408, -3.13 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Quillward Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Thornstone Sledgehammer

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0500220000000000-50212032123012123001-0100000000000000)

Set DPS (verified): 263.5. Weights run: 7.9s. Verify run: 18.2s. 816 eligible items had no known source.

Race sweep: windshaper-skyborne 264.6 ± 0.2, tauren 263.5 ± 0.2 (dps); windshaper-skyborne beats the guide race tauren by 1.0 (0.40%) ± 0.3, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.040 ± 0.002, agility=2.080 ± 0.029, crit=1.258 ± 0.038 per rating point (14 rating = 1%, 17.606 per %), hit=1.005 ± 0.090 per rating point (10 rating = 1%, 10.052 per %), melee_haste=not significant (4.533 ± 1.221)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.3 attack_power points (6.97 DPS) | yes | Blood Guard's Leather Headband (220851, -0.67 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.73 DPS) [quest]; White Bandit Mask (10008, -1.06 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 37.2 attack_power points (5.37 DPS) | yes | Scout's Medallion (19535, -1.77 DPS) [rep]; Woven Ivy Necklace (19159, -1.89 DPS, sim-verified) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 41.2 attack_power points (5.95 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.01 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.81 DPS) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.47 DPS) [dungeon]; Dark Phantom Cape (13122, -1.82 DPS, sim-verified) [world_drop]; Blackmetal Cape (9512, -2.37 DPS) [dungeon] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 68.0 attack_power points (9.81 DPS) | yes | Warbear Harness (15064, -1.14 DPS, sim-verified) [crafted]; Mixologist's Tunic (12793, -1.51 DPS) [dungeon]; Quillward Harness (10583, -1.75 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 39.4 attack_power points (5.68 DPS) | yes | Prowler's Leather Bracers (252539, -0.93 DPS) [crafted]; Pridelord Bands (14672, -1.21 DPS) [world_drop]; Skulker's Leather Bracers (252540, -1.30 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -2.56 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 53.6 attack_power points (7.73 DPS) | yes | Prowler's Leather Waistguard (252473, -0.01 DPS) [crafted]; Girdle of Beastial Fury (11686, -1.05 DPS) [dungeon]; Defiler's Leather Girdle (20193, -2.31 DPS) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 59.8 attack_power points (8.63 DPS) | yes | Stone Guard's Leather Pants (220859, -2.26 DPS, sim-verified) [vendor]; Basilisk Hide Pants (1718, -2.33 DPS) [world_drop]; Triprunner Dungarees (9624, -2.34 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 49.5 attack_power points (7.14 DPS) | yes | Prowler's Leather Boots (252468, -0.01 DPS) [crafted]; Shadefiend Boots (11675, -0.89 DPS) [dungeon]; Albino Crocscale Boots (17728, -1.14 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 39.1 attack_power points (5.64 DPS) | yes | Masons Fraternity Ring (9533, -1.44 DPS) [quest]; Ring of the Underwood (2951, -2.05 DPS) [world_drop]; Falcon's Hook (7552, -2.06 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 30.1 attack_power points (4.34 DPS) | yes | Ring of the Underwood (2951, -0.75 DPS) [world_drop]; Falcon's Hook (7552, -0.75 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.93 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.73 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Mark of the Chosen (17774, -2.83 DPS, sim-verified) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -3.81 DPS, sim-verified) [world_drop]; Dark Iron Pulverizer (11608, -6.04 DPS) [crafted]; Soulkeeper (1607, -7.76 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 816, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0500220000000000-50212032123012123001-0550100000000000)

Set DPS (verified): 451.9. Weights run: 7.7s. Verify run: 30.8s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 454.2 ± 0.3, tauren 451.9 ± 0.3 (dps); windshaper-skyborne beats the guide race tauren by 2.3 (0.51%) ± 0.4, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.211 ± 0.039, crit=1.730 ± 0.055 per rating point (14 rating = 1%, 24.218 per %), hit=1.281 ± 0.131 per rating point (10 rating = 1%, 12.814 per %), melee_haste=not significant (3.883 ± 1.720)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 75.7 attack_power points (13.33 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 50.2 attack_power points (8.84 DPS) | yes | Skibi's Pendant (13089, -1.44 DPS, sim-verified) [world_drop]; Scout's Medallion (19534, -3.00 DPS) [rep]; Woven Ivy Necklace (19159, -3.01 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 83.9 attack_power points (14.77 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -0.42 DPS) [vendor]; Defiler's Leather Shoulders (20194, -3.23 DPS, sim-verified) [rep]; Champion's Dragonhide Shoulders (227175, -3.53 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -1.59 DPS) [dungeon]; Deathguard's Cloak (20068, -1.61 DPS, sim-verified) [rep]; Howler's Furs (272414, -2.17 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 136.6 attack_power points (24.05 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, -4.34 DPS) [vendor]; Dawn Armor (252483, -4.42 DPS, sim-verified) [crafted]; Warlord's Dragonhide Hauberk (16549, -4.66 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 66.2 attack_power points (11.66 DPS) | yes | Deepfury Bracers (13120, -4.27 DPS) [world_drop]; Wristguards of Renown (22204, -4.28 DPS) [dungeon]; Skulker's Leather Bracers (252540, -5.44 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, -0.80 DPS) [vendor]; Timbermaw Brawlers (19049, -1.19 DPS) [crafted]; Blood Guard's Dragonhide Grips (227180, -3.59 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 83.9 attack_power points (14.77 DPS) | yes | Might of the Timbermaw (19044, -1.18 DPS) [crafted]; Cloudrunner Girdle (13252, -4.28 DPS) [dungeon]; Defiler's Leather Girdle (20190, -4.51 DPS) [rep] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 101.4 attack_power points (17.85 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -2.01 DPS) [pvp]; Outrider's Leather Pants (22740, -2.55 DPS, sim-verified) [rep] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 57.4 attack_power points (10.11 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.79 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.28 DPS) [vendor]; Magma Forged Band (22255, -1.60 DPS) [dungeon]; Naglering (11669, -10.19 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magma Forged Band (22255, -1.55 DPS) [dungeon]; Cutthroat's Signet (272408, -1.76 DPS, sim-verified) [vendor]; Myrmidon's Signet (2246, -1.94 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -1.79 DPS) [dungeon]; Hand of Justice (11815, -1.93 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.01 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Molten Heart of the Mountain (249470, -3.55 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -7.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Ferocity (22397, -10.42 DPS, sim-verified) [dungeon] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Warbear Woolies; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0500220000000000-50230032023012125001-0550000000000000)

Set DPS (verified): 1085.5. Weights run: 7.6s. Verify run: 54.1s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 1091.1 ± 0.7, tauren 1085.5 ± 0.7 (dps); windshaper-skyborne beats the guide race tauren by 5.6 (0.52%) ± 1.0, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.489 ± 0.035, crit=1.800 ± 0.047 per rating point (14 rating = 1%, 25.198 per %), hit=1.855 ± 0.157 per rating point (10 rating = 1%, 18.553 per %), melee_haste=7.750 ± 1.917

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 83.4 attack_power points (25.49 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 51.2 attack_power points (15.65 DPS) | yes | Skibi's Pendant (13089, -2.06 DPS) [world_drop]; Scout's Medallion (19534, -4.24 DPS) [rep]; Woven Ivy Necklace (19159, -4.36 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 92.4 attack_power points (28.25 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -1.06 DPS) [vendor]; Defiler's Leather Shoulders (20194, -6.04 DPS, sim-verified) [rep]; Champion's Dragonhide Shoulders (227175, -7.10 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -2.44 DPS) [dungeon]; Howler's Furs (272414, -3.30 DPS) [vendor]; Deathguard's Cloak (20068, -3.33 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 149.9 attack_power points (45.82 DPS) | yes | Dawn Armor (252483, -4.28 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -6.99 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -8.96 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 73.9 attack_power points (22.60 DPS) | yes | Deepfury Bracers (13120, -8.22 DPS) [world_drop]; Wristguards of Renown (22204, -8.33 DPS) [dungeon]; Skulker's Leather Bracers (252540, -10.57 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, -2.28 DPS) [vendor]; Timbermaw Brawlers (19049, -3.13 DPS, sim-verified) [crafted]; Blood Guard's Dragonhide Grips (227180, -5.79 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 93.6 attack_power points (28.62 DPS) | yes | Might of the Timbermaw (19044, -3.03 DPS, sim-verified) [crafted]; Cloudrunner Girdle (13252, -8.35 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -9.09 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 112.6 attack_power points (34.41 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -3.03 DPS) [pvp]; Outrider's Leather Pants (22740, -5.40 DPS) [rep] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 64.4 attack_power points (19.68 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -1.65 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.33 DPS) [vendor]; Magma Forged Band (22255, -3.34 DPS) [dungeon]; Naglering (11669, -19.64 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.92 DPS) [vendor]; Magma Forged Band (22255, -2.94 DPS) [dungeon]; Myrmidon's Signet (2246, -3.78 DPS) [world_drop] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -10.22 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -2.58 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.15 DPS) [crafted]; Hand of Justice (11815, -6.27 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -21.61 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of Ferocity (22397, -19.45 DPS, sim-verified) [dungeon] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Warbear Woolies; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Blackhand's Breadth; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

