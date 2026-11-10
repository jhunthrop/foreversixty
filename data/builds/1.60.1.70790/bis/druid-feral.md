# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 53.1. Weights run: 5.3s. Verify run: 9.0s. 251 eligible items had no known source.

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

Set DPS (verified): 97.5. Weights run: 5.7s. Verify run: 12.6s. 424 eligible items had no known source.

Race sweep: night-elf 97.5 ± 0.1, high-order-skyborne 96.7 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.995 ± 0.013, crit=0.796 ± 0.014 per rating point (14 rating = 1%, 11.145 per %), hit=0.929 ± 0.057 per rating point (10 rating = 1%, 9.285 per %), melee_haste=3.761 ± 0.713

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.99 DPS) | yes | Brawler's Leather Helm (252512, -0.37 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.0 attack_power points (1.33 DPS) | yes | Sentinel's Medallion (19541, -0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.33 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.9 attack_power points (2.65 DPS) | yes | Mantle of Thieves (2264, -1.00 DPS) [dungeon]; Barbaric Shoulders (5964, -1.01 DPS, sim-verified) [crafted]; Bristlebark Amice (14573, -1.16 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.33 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.33 DPS, sim-verified) [pvp]; Glowing Lizardscale Cloak (6449, -0.66 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 28.0 attack_power points (2.32 DPS) | yes | Dusky Leather Armor (7374, -0.00 DPS) [crafted]; Defender's Leather Tunic (252450, -0.33 DPS) [crafted]; Brawler's Leather Armor (252490, -0.33 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.0 attack_power points (1.66 DPS) | yes | Barbaric Bracers (18948, -0.34 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.50 DPS) [quest]; Bands of Serra'kis (6902, -0.66 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.0 attack_power points (2.16 DPS) | yes | Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.33 DPS) [dungeon]; Gloves of the Fang (10413, -0.50 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 30.0 attack_power points (2.49 DPS) | yes | Skulker's Leather Belt (252520, -0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.50 DPS) [rep]; Highlander's Leather Girdle (20117, -0.50 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 32.0 attack_power points (2.65 DPS) | yes | Brawler's Leather Pants (252500, -0.17 DPS) [crafted]; Trapper's Leather Pants (252501, -0.17 DPS) [crafted]; Barbaric Leggings (5963, -0.33 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.0 attack_power points (1.82 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Insignia Boots (4055, -0.50 DPS) [world_drop]; Disjointed Shoes (277226, -0.83 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 24.0 attack_power points (1.99 DPS) | yes | Seal of Wrynn (2933, -0.99 DPS) [quest]; Ring of Precision (1491, -1.00 DPS) [dungeon]; Tiger Band (6749, -1.16 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.0 attack_power points (1.82 DPS) | yes | Ring of Precision (1491, -0.83 DPS) [dungeon]; Seal of Wrynn (2933, -0.87 DPS, sim-verified) [quest]; Tiger Band (6749, -0.99 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Viscous Hammer (13045, -2.77 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -3.59 DPS) [dungeon]; Gnarled Ash Staff (791, -7.41 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 149.5. Weights run: 7.0s. Verify run: 9.2s. 607 eligible items had no known source.

Race sweep: night-elf 149.5 ± 0.1, high-order-skyborne 147.7 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.229 ± 0.032, crit=1.257 ± 0.039 per rating point (14 rating = 1%, 17.597 per %), hit=1.110 ± 0.088 per rating point (10 rating = 1%, 11.098 per %), melee_haste=not significant (4.364 ± 1.102)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.5 attack_power points (4.61 DPS) | yes | Hawkeye's Helm (14591, -1.40 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.74 DPS) [world_drop]; Nightscape Headband (8176, -1.96 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.5 attack_power points (2.43 DPS) | yes | Kaleidoscope Chain (13084, -0.75 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.04 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.24 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.5 attack_power points (3.42 DPS) | yes | Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Nightscape Shoulders (8192, -0.99 DPS) [crafted]; Fleshhide Shoulders (10774, -1.10 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Yeti Fur Cloak (2805, -0.59 DPS) [quest]; Parachute Cloak (10518, -0.75 DPS) [crafted]; Hawkeye's Cloak (14593, -1.78 DPS, sim-verified) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 58.4 attack_power points (5.78 DPS) | yes | Wolffear Harness (13110, -1.79 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.47 DPS) [crafted]; Barbaric Harness (5739, -2.63 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.4 attack_power points (2.12 DPS) | yes | Branded Leather Bracers (19508, -0.14 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.35 DPS) [dungeon]; Dusky Bracers (7378, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Prowler's Leather Gloves (252524, -1.88 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 38.1 attack_power points (3.77 DPS) | yes | Skulker's Leather Belt (252520, -0.59 DPS) [crafted]; Prowler's Leather Belt (252459, -0.66 DPS) [crafted]; Scorpashi Sash (14652, -0.77 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.8 attack_power points (4.64 DPS) | yes | Triprunner Dungarees (9624, -0.90 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -1.26 DPS) [crafted]; Brawler's Leather Pants (252500, -1.46 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 38.5 attack_power points (3.82 DPS) | yes | Imperial Leather Boots (6431, -0.40 DPS) [dungeon]; Excelsior Boots (4109, -0.44 DPS) [quest]; Prowler's Leather Shoes (252465, -0.66 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 33.8 attack_power points (3.35 DPS) | yes | Falcon's Hook (7552, -0.77 DPS) [dungeon]; Thunderbrow Ring (13097, -1.10 DPS) [world_drop]; Mark of Kern (2262, -1.37 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 26.3 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.02 DPS) [dungeon]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Thornstone Sledgehammer (1722) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.11 DPS) [dungeon]; The Jackhammer (9423, -0.62 DPS) [dungeon]; Ironshod Bludgeon (9408, -2.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Thornstone Sledgehammer

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0500220000000000-50212032123012123001-0100000000000000)

Set DPS (verified): 224.2. Weights run: 7.0s. Verify run: 24.5s. 815 eligible items had no known source.

Race sweep: night-elf 224.2 ± 0.2, high-order-skyborne 221.9 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.040 ± 0.002, agility=2.379 ± 0.039, crit=1.652 ± 0.052 per rating point (14 rating = 1%, 23.123 per %), hit=1.327 ± 0.113 per rating point (10 rating = 1%, 13.265 per %), melee_haste=7.301 ± 1.445

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; White Bandit Mask (10008, -0.64 DPS, sim-verified) [crafted]; Helm of Fire (8348, -0.84 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 41.1 attack_power points (4.40 DPS) | yes | Sentinel's Medallion (19539, -0.97 DPS, sim-verified) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 44.5 attack_power points (4.77 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -0.02 DPS) [vendor]; Failed Flying Experiment (9647, -0.04 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.07 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -1.06 DPS) [dungeon]; Dark Phantom Cape (13122, -1.06 DPS) [world_drop]; Blackmetal Cape (9512, -1.82 DPS) [dungeon] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 72.8 attack_power points (7.79 DPS) | yes | Warbear Harness (15064, -0.80 DPS) [crafted]; Quillward Harness (10583, -1.20 DPS) [dungeon]; Blazewind Breastplate (11193, -1.27 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.9 attack_power points (4.69 DPS) | yes | Skulker's Leather Bracers (252540, -0.87 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.95 DPS) [crafted]; Pridelord Bands (14672, -1.05 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -0.55 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 57.8 attack_power points (6.19 DPS) | yes | Prowler's Leather Waistguard (252473, -0.07 DPS) [crafted]; Girdle of Beastial Fury (11686, -1.23 DPS) [dungeon]; Highlander's Leather Girdle (20115, -1.57 DPS) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 66.3 attack_power points (7.09 DPS) | yes | Serpentskin Leggings (8262, -0.18 DPS) [world_drop]; Knight's Leather Pants (220858, -1.27 DPS) [vendor]; Basilisk Hide Pants (1718, -1.75 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 53.4 attack_power points (5.71 DPS) | yes | Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.62 DPS) [dungeon]; Shadefiend Boots (11675, -0.73 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 41.8 attack_power points (4.48 DPS) | yes | Masons Fraternity Ring (9533, -0.91 DPS) [quest]; Ring of the Underwood (2951, -1.49 DPS) [world_drop]; Falcon's Hook (7552, -1.53 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of the Underwood (2951, -0.58 DPS) [world_drop]; Falcon's Hook (7552, -0.61 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.69 DPS, sim-verified) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -1.57 DPS, sim-verified) [world_drop]; Dark Iron Pulverizer (11608, -4.48 DPS) [crafted]; The Rockpounder (9413, -5.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 815, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0500220000000000-50212032123012123001-0550100000000000)

Set DPS (verified): 389.6. Weights run: 7.0s. Verify run: 52.1s. 1969 eligible items had no known source.

Race sweep: night-elf 389.6 ± 0.3, high-order-skyborne 386.9 ± 0.3 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.577 ± 0.050, crit=2.252 ± 0.072 per rating point (14 rating = 1%, 31.527 per %), hit=1.959 ± 0.149 per rating point (10 rating = 1%, 19.585 per %), melee_haste=not significant (4.979 ± 1.899)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 90.1 attack_power points (9.71 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, +0.00 DPS) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 57.5 attack_power points (6.20 DPS) | yes | Skibi's Pendant (13089, -1.41 DPS) [world_drop]; Sentinel's Medallion (19538, -2.04 DPS) [rep]; Imperial Jewel (11933, -2.75 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 101.1 attack_power points (10.90 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -0.85 DPS) [vendor]; Highlander's Leather Shoulders (20059, -1.85 DPS, sim-verified) [rep]; Lieutenant Commander's Dragonhide Shoulders (227172, -2.99 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -1.01 DPS) [dungeon]; Howler's Furs (272414, -1.19 DPS) [vendor]; Cloak of the Honor Guard (20073, -1.27 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 150.5 attack_power points (16.23 DPS) | yes | Dawn Armor (252483, -0.44 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -2.08 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -3.01 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 73.2 attack_power points (7.89 DPS) | yes | Deepfury Bracers (13120, -2.77 DPS) [world_drop]; Wristguards of Renown (22204, -2.98 DPS) [dungeon]; Shadowcraft Bracers (16710, -3.72 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, -0.37 DPS) [vendor]; Timbermaw Brawlers (19049, -1.24 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -2.14 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 92.6 attack_power points (9.99 DPS) | yes | Might of the Timbermaw (19044, -1.12 DPS) [crafted]; Highlander's Leather Girdle (20045, -2.92 DPS) [rep]; Cloudrunner Girdle (13252, -3.02 DPS) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 132.6 attack_power points (14.30 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -2.61 DPS) [vendor]; Warbear Woolies (15065, -2.66 DPS) [crafted] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 65.1 attack_power points (7.02 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -0.66 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -1.12 DPS) [rep]; Painweaver Band (13098, -1.90 DPS) [dungeon]; Naglering (11669, -7.49 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -0.10 DPS) [rep]; Painweaver Band (13098, -0.88 DPS) [dungeon]; Magni's Will (12548, -0.94 DPS) [quest] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Thunderbrew's Boot Flask (744, -3.13 DPS, sim-verified) [quest] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -11.24 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -14.95 DPS, sim-verified) [world_drop] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Molten Heart of the Mountain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1969, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0500220000000000-50230032023012125001-0550000000000000)

Set DPS (verified): 952.3. Weights run: 7.2s. Verify run: 29.0s. 1969 eligible items had no known source.

Race sweep: night-elf 952.3 ± 0.6, high-order-skyborne 950.3 ± 0.6 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.635 ± 0.042, crit=1.944 ± 0.056 per rating point (14 rating = 1%, 27.211 per %), hit=2.271 ± 0.172 per rating point (10 rating = 1%, 22.714 per %), melee_haste=26.367 ± 2.087

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 88.5 attack_power points (21.87 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -3.03 DPS, sim-verified) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 53.2 attack_power points (13.15 DPS) | yes | Skibi's Pendant (13089, -1.70 DPS) [world_drop]; Sentinel's Medallion (19538, -3.38 DPS) [rep]; Imperial Jewel (11933, -5.24 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 98.3 attack_power points (24.31 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -1.39 DPS) [vendor]; Lieutenant Commander's Dragonhide Shoulders (227172, -6.50 DPS) [pvp]; Highlander's Leather Shoulders (20059, -6.73 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -2.01 DPS) [dungeon]; Howler's Furs (272414, -2.18 DPS) [vendor]; Cloak of the Honor Guard (20073, -3.05 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 154.5 attack_power points (38.18 DPS) | yes | Dawn Armor (252483, -3.61 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -4.77 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -7.39 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 76.7 attack_power points (18.95 DPS) | yes | Deepfury Bracers (13120, -6.79 DPS) [world_drop]; Wristguards of Renown (22204, -7.06 DPS) [dungeon]; Skulker's Leather Bracers (252540, -8.90 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, -1.95 DPS) [vendor]; Timbermaw Brawlers (19049, -2.48 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -4.26 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 97.1 attack_power points (24.00 DPS) | yes | Might of the Timbermaw (19044, -2.33 DPS) [crafted]; Cloudrunner Girdle (13252, -7.11 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -7.71 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 125.6 attack_power points (31.03 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -3.70 DPS) [vendor]; Warbear Woolies (15065, -4.32 DPS, sim-verified) [crafted] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 67.4 attack_power points (16.67 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -3.29 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -1.46 DPS) [rep]; Magma Forged Band (22255, -3.80 DPS) [dungeon]; Naglering (11669, -21.62 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Magma Forged Band (22255, -2.73 DPS) [dungeon]; Protector's Band (19516, -2.89 DPS) [rep] |
| trinket1 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Hand of Justice (11815, -7.80 DPS, sim-verified) [dungeon] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Hand of Justice (11815, -2.09 DPS) [dungeon]; Cannonball Runner (13382, -2.25 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Slavedriver's Cane (13372, -18.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -23.38 DPS, sim-verified) [vendor] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; trinket2: Frozen Heart of the Mountain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1969, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 52.2. Weights run: 5.3s. Verify run: 8.9s. 247 eligible items had no known source.

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

Set DPS (verified): 95.9. Weights run: 5.7s. Verify run: 14.0s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 96.7 ± 0.1, tauren 95.9 ± 0.1 (dps); windshaper-skyborne beats the guide race tauren by 0.9 (0.90%) ± 0.1, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.995 ± 0.013, crit=0.796 ± 0.014 per rating point (14 rating = 1%, 11.145 per %), hit=0.929 ± 0.057 per rating point (10 rating = 1%, 9.285 per %), melee_haste=3.761 ± 0.713

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.99 DPS) | yes | Brawler's Leather Helm (252512, -0.32 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.0 attack_power points (1.33 DPS) | yes | Scout's Medallion (19537, -0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.33 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.9 attack_power points (2.65 DPS) | yes | Mantle of Thieves (2264, -1.00 DPS) [dungeon]; Barbaric Shoulders (5964, -1.06 DPS, sim-verified) [crafted]; Bristlebark Amice (14573, -1.16 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.35 DPS, sim-verified) [world_drop]; Glowing Lizardscale Cloak (6449, -0.66 DPS) [dungeon]; Swiftrunner Cape (6745, -0.66 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 28.0 attack_power points (2.32 DPS) | yes | Dusky Leather Armor (7374, -0.00 DPS) [crafted]; Defender's Leather Tunic (252450, -0.33 DPS) [crafted]; Brawler's Leather Armor (252490, -0.33 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.0 attack_power points (1.66 DPS) | yes | Barbaric Bracers (18948, -0.36 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.66 DPS) [dungeon]; Insignia Bracers (6410, -0.66 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.0 attack_power points (2.16 DPS) | yes | Toughened Leather Gloves (4253, -0.17 DPS) [crafted]; Wolfclaw Gloves (1978, -0.33 DPS) [dungeon]; Gloves of the Fang (10413, -0.50 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 30.0 attack_power points (2.49 DPS) | yes | Skulker's Leather Belt (252520, -0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.50 DPS) [rep]; Defiler's Leather Girdle (20191, -0.50 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 32.0 attack_power points (2.65 DPS) | yes | Brawler's Leather Pants (252500, -0.17 DPS) [crafted]; Trapper's Leather Pants (252501, -0.17 DPS) [crafted]; Barbaric Leggings (5963, -0.33 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.0 attack_power points (1.82 DPS) | yes | Brawler's Leather Boots (252439, -0.16 DPS) [crafted]; Stomping Boots (3741, -0.50 DPS) [quest]; Insignia Boots (4055, -0.50 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 24.0 attack_power points (1.99 DPS) | yes | Band of the Fist (17694, -0.66 DPS) [quest]; Ring of Precision (1491, -1.00 DPS) [dungeon]; Tiger Band (6749, -1.16 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.0 attack_power points (1.82 DPS) | yes | Band of the Fist (17694, -0.52 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.83 DPS) [dungeon]; Tiger Band (6749, -0.99 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Viscous Hammer (13045, -2.57 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -3.59 DPS) [dungeon]; Gnarled Ash Staff (791, -7.41 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 145.8. Weights run: 7.0s. Verify run: 10.8s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 147.2 ± 0.1, tauren 145.8 ± 0.1 (dps); windshaper-skyborne beats the guide race tauren by 1.3 (0.92%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.229 ± 0.032, crit=1.257 ± 0.039 per rating point (14 rating = 1%, 17.597 per %), hit=1.110 ± 0.088 per rating point (10 rating = 1%, 11.098 per %), melee_haste=not significant (4.364 ± 1.102)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.5 attack_power points (4.61 DPS) | yes | Hawkeye's Helm (14591, -1.36 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.74 DPS) [world_drop]; Nightscape Headband (8176, -1.96 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.5 attack_power points (2.43 DPS) | yes | Ethereal Talisman (4430, -0.68 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.75 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.04 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.5 attack_power points (3.42 DPS) | yes | Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Nightscape Shoulders (8192, -0.99 DPS) [crafted]; Fleshhide Shoulders (10774, -1.10 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.37 DPS) [dungeon]; Parachute Cloak (10518, -0.37 DPS) [crafted]; Tigerstrike Mantle (13108, -0.37 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 58.4 attack_power points (5.78 DPS) | yes | Wolffear Harness (13110, -1.72 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.47 DPS) [crafted]; Barbaric Harness (5739, -2.63 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.4 attack_power points (2.12 DPS) | yes | Branded Leather Bracers (19508, -0.14 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.35 DPS) [dungeon]; Dusky Bracers (7378, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -1.17 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 38.1 attack_power points (3.77 DPS) | yes | Skulker's Leather Belt (252520, -0.59 DPS) [crafted]; Prowler's Leather Belt (252459, -0.66 DPS) [crafted]; Scorpashi Sash (14652, -0.77 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.8 attack_power points (4.64 DPS) | yes | Triprunner Dungarees (9624, -0.07 DPS) [quest]; Brawler's Leather Legguards (252516, -1.26 DPS) [crafted]; Brawler's Leather Pants (252500, -1.46 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 38.5 attack_power points (3.82 DPS) | yes | Prowler's Leather Shoes (252465, -0.09 DPS) [crafted]; Imperial Leather Boots (6431, -0.40 DPS) [dungeon]; Excelsior Boots (4109, -0.44 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 33.8 attack_power points (3.35 DPS) | yes | Falcon's Hook (7552, -0.77 DPS) [dungeon]; Thunderbrow Ring (13097, -1.10 DPS) [world_drop]; Mark of Kern (2262, -1.37 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 26.3 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.02 DPS) [dungeon]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Mograine's Might (7723) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornstone Sledgehammer (1722, -0.47 DPS, sim-verified) [world_drop]; The Jackhammer (9423, -0.51 DPS) [dungeon]; Ironshod Bludgeon (9408, -2.25 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Quillward Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Mograine's Might

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0500220000000000-50212032123012123001-0100000000000000)

Set DPS (verified): 220.8. Weights run: 7.0s. Verify run: 20.7s. 814 eligible items had no known source.

Race sweep: windshaper-skyborne 221.9 ± 0.2, tauren 220.8 ± 0.2 (dps); windshaper-skyborne beats the guide race tauren by 1.1 (0.52%) ± 0.2, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.040 ± 0.002, agility=2.136 ± 0.033, crit=1.365 ± 0.044 per rating point (14 rating = 1%, 19.108 per %), hit=1.200 ± 0.089 per rating point (10 rating = 1%, 11.998 per %), melee_haste=not significant (3.754 ± 1.116)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 48.3 attack_power points (4.69 DPS) | yes | Blood Guard's Leather Headband (220851, -0.12 DPS) [vendor]; White Bandit Mask (10008, -0.23 DPS) [crafted]; Undercity Reservist's Cap (20643, -0.43 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 38.0 attack_power points (3.69 DPS) | yes | Scout's Medallion (19535, -1.20 DPS) [rep]; Woven Ivy Necklace (19159, -1.24 DPS, sim-verified) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 41.9 attack_power points (4.07 DPS) | yes | Failed Flying Experiment (9647, -0.01 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.02 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.27 DPS) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.98 DPS) [dungeon]; Dark Phantom Cape (13122, -0.98 DPS) [world_drop]; Blackmetal Cape (9512, -1.60 DPS) [dungeon] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 68.8 attack_power points (6.69 DPS) | yes | Warbear Harness (15064, -0.77 DPS) [crafted]; Mixologist's Tunic (12793, -1.05 DPS) [dungeon]; Quillward Harness (10583, -1.16 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 40.2 attack_power points (3.90 DPS) | yes | Skulker's Leather Bracers (252540, -0.53 DPS, sim-verified) [crafted]; Prowler's Leather Bracers (252539, -0.67 DPS) [crafted]; Pridelord Bands (14672, -0.84 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 54.4 attack_power points (5.28 DPS) | yes | Prowler's Leather Waistguard (252473, -0.02 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.78 DPS) [dungeon]; Defiler's Leather Girdle (20193, -1.48 DPS) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 60.7 attack_power points (5.89 DPS) | yes | Stone Guard's Leather Pants (220859, -1.27 DPS, sim-verified) [vendor]; Basilisk Hide Pants (1718, -1.54 DPS) [world_drop]; Triprunner Dungarees (9624, -1.57 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 50.2 attack_power points (4.88 DPS) | yes | Prowler's Leather Boots (252468, -0.02 DPS) [crafted]; Shadefiend Boots (11675, -0.61 DPS) [dungeon]; Albino Crocscale Boots (17728, -0.73 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 39.6 attack_power points (3.85 DPS) | yes | Masons Fraternity Ring (9533, -0.94 DPS) [quest]; Ring of the Underwood (2951, -1.38 DPS) [world_drop]; Falcon's Hook (7552, -1.39 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 32.0 attack_power points (3.11 DPS) | yes | Masons Fraternity Ring (9533, -0.39 DPS, sim-verified) [quest]; Ring of the Underwood (2951, -0.64 DPS) [world_drop]; Falcon's Hook (7552, -0.65 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.13 DPS) [crafted]; Ankh of Life (1713, -1.85 DPS, sim-verified) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -1.38 DPS, sim-verified) [world_drop]; Dark Iron Pulverizer (11608, -4.06 DPS) [crafted]; Soulkeeper (1607, -5.22 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Thorium Greatmace

No-known-source sample (15 of 814, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0500220000000000-50212032123012123001-0550100000000000)

Set DPS (verified): 381.9. Weights run: 7.0s. Verify run: 54.3s. 1973 eligible items had no known source.

Race sweep: windshaper-skyborne 385.3 ± 0.3, tauren 381.9 ± 0.3 (dps); windshaper-skyborne beats the guide race tauren by 3.4 (0.88%) ± 0.4, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.439 ± 0.044, crit=2.055 ± 0.063 per rating point (14 rating = 1%, 28.774 per %), hit=1.542 ± 0.156 per rating point (10 rating = 1%, 15.418 per %), melee_haste=not significant (7.391 ± 1.903)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 83.9 attack_power points (10.98 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.09 DPS, sim-verified) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 54.8 attack_power points (7.17 DPS) | yes | Skibi's Pendant (13089, -0.82 DPS, sim-verified) [world_drop]; Scout's Medallion (19534, -2.38 DPS) [rep]; Woven Ivy Necklace (19159, -2.57 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 94.6 attack_power points (12.39 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -0.76 DPS) [vendor]; Defiler's Leather Shoulders (20194, -1.82 DPS, sim-verified) [rep]; Champion's Dragonhide Shoulders (227175, -3.25 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -0.51 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -1.36 DPS) [rep]; Howler's Furs (272414, -1.72 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 145.3 attack_power points (19.02 DPS) | yes | Dawn Armor (252483, -1.42 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.00 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -3.58 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 70.5 attack_power points (9.24 DPS) | yes | Deepfury Bracers (13120, -3.29 DPS) [world_drop]; Wristguards of Renown (22204, -3.45 DPS) [dungeon]; Skulker's Leather Bracers (252540, -4.35 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, -0.51 DPS) [vendor]; Timbermaw Brawlers (19049, -1.27 DPS) [crafted]; Blood Guard's Dragonhide Grips (227180, -2.83 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 89.3 attack_power points (11.70 DPS) | yes | Might of the Timbermaw (19044, -1.18 DPS) [crafted]; Defiler's Leather Girdle (20190, -3.48 DPS) [rep]; Cloudrunner Girdle (13252, -3.48 DPS) [dungeon] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 105.5 attack_power points (13.81 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -0.74 DPS) [pvp]; Outrider's Leather Pants (22740, -1.11 DPS) [rep] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 62.2 attack_power points (8.15 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.83 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.13 DPS) [vendor]; Painweaver Band (13098, -1.76 DPS) [dungeon]; Naglering (11669, -5.50 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.48 DPS) [vendor]; Painweaver Band (13098, -1.11 DPS) [dungeon]; Magma Forged Band (22255, -1.12 DPS) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -1.23 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.43 DPS) [crafted]; Hand of Justice (11815, -1.68 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -10.52 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -16.50 DPS, sim-verified) [vendor] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Warbear Woolies; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1973, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0500220000000000-50230032023012125001-0550000000000000)

Set DPS (verified): 937.8. Weights run: 7.2s. Verify run: 29.0s. 1973 eligible items had no known source.

Race sweep: windshaper-skyborne 947.9 ± 0.6, tauren 937.8 ± 0.6 (dps); windshaper-skyborne beats the guide race tauren by 10.1 (1.08%) ± 0.9, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.635 ± 0.042, crit=1.944 ± 0.056 per rating point (14 rating = 1%, 27.211 per %), hit=2.271 ± 0.172 per rating point (10 rating = 1%, 22.714 per %), melee_haste=26.367 ± 2.087

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Rend (12587) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 88.5 attack_power points (21.87 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -4.59 DPS, sim-verified) [crafted] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 53.2 attack_power points (13.15 DPS) | yes | Skibi's Pendant (13089, -2.89 DPS, sim-verified) [world_drop]; Scout's Medallion (19534, -3.38 DPS) [rep]; Woven Ivy Necklace (19159, -3.70 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 98.3 attack_power points (24.31 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -1.39 DPS) [vendor]; Champion's Dragonhide Shoulders (227175, -6.50 DPS) [pvp]; Defiler's Leather Shoulders (20194, -8.97 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -2.01 DPS) [dungeon]; Howler's Furs (272414, -2.18 DPS) [vendor]; Deathguard's Cloak (20068, -3.05 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 154.5 attack_power points (38.18 DPS) | yes | Dawn Armor (252483, -1.75 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -4.77 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -7.39 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 76.7 attack_power points (18.95 DPS) | yes | Deepfury Bracers (13120, -6.79 DPS) [world_drop]; Wristguards of Renown (22204, -7.06 DPS) [dungeon]; Skulker's Leather Bracers (252540, -8.90 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, -1.95 DPS) [vendor]; Blood Guard's Dragonhide Grips (227180, -4.26 DPS) [pvp]; Timbermaw Brawlers (19049, -5.51 DPS, sim-verified) [crafted] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 97.1 attack_power points (24.00 DPS) | yes | Might of the Timbermaw (19044, -4.28 DPS, sim-verified) [crafted]; Cloudrunner Girdle (13252, -7.11 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -7.71 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 115.2 attack_power points (28.47 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.14 DPS) [pvp]; Plaguehound Leggings (18736, -2.77 DPS, sim-verified) [dungeon] |
| feet | Swiftwalker Boots (12553) | Blackrock Depths: Princess Moira Bronzebeard  [dungeon] | 67.4 attack_power points (16.67 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -1.31 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -1.46 DPS) [rep]; Magma Forged Band (22255, -3.80 DPS) [dungeon]; Naglering (11669, -19.57 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Magma Forged Band (22255, -2.73 DPS) [dungeon]; Legionnaire's Band (19511, -2.89 DPS) [rep] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -6.50 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -1.78 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -2.34 DPS) [crafted]; Hand of Justice (11815, -4.42 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Slavedriver's Cane (13372, -22.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -17.68 DPS, sim-verified) [vendor] |

**New at 60:** head: Eye of Rend; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Warbear Woolies; feet: Swiftwalker Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1973, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

