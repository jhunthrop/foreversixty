# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 53.1. Weights run: 8.1s. Verify run: 11.1s. 228 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.907 ± 0.010, crit=0.568 ± 0.009 per rating point (14 rating = 1%, 7.954 per %), hit=0.654 ± 0.036 per rating point (10 rating = 1%, 6.545 per %), melee_haste=2.420 ± 0.455

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.4 attack_power points (0.79 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.21 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.5 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (53.1 DPS) | yes | Dark Leather Cloak (2316, -0.12 DPS) [crafted]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (53.1 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.40 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 11.7 attack_power points (0.81 DPS) | yes | Bristlebark Bindings (14569, -0.14 DPS) [world_drop]; Forest Leather Bracers (3202, -0.15 DPS) [world_drop]; Wolf Bracers (4794, -0.28 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (53.1 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.55 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.24 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Murloc Scale Belt (5780, -0.57 DPS) [crafted]; Brawler's Leather Belt (252428, -3.22 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (53.1 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.65 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.3 attack_power points (1.47 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Blackened Defias Boots (10402, -0.68 DPS) [dungeon]; Footpads of the Fang (10411, -0.68 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.6 attack_power points (1.08 DPS) | yes | The 1 Ring (8350, -0.81 DPS) [world]; Lavishly Jeweled Ring (1156, -0.82 DPS) [dungeon]; Ring of the Moon (12052, -0.94 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.55 DPS) | yes | The 1 Ring (8350, -0.28 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Ring of the Moon (12052, -0.41 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.6 attack_power points (21.14 DPS) | yes | Twisted Chanter's Staff (890, -1.51 DPS, sim-verified) [world_drop]; Rakzur Club (12983, -2.89 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -3.29 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 228, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 97.5. Weights run: 8.6s. Verify run: 16.7s. 400 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.995 ± 0.013, crit=0.796 ± 0.014 per rating point (14 rating = 1%, 11.145 per %), hit=0.929 ± 0.057 per rating point (10 rating = 1%, 9.285 per %), melee_haste=3.761 ± 0.713

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.99 DPS) | yes | Brawler's Leather Helm (252512, -0.37 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.0 attack_power points (1.33 DPS) | yes | Sentinel's Medallion (19541, -0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.33 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.9 attack_power points (2.65 DPS) | yes | Mantle of Thieves (2264, -1.00 DPS) [dungeon]; Barbaric Shoulders (5964, -1.01 DPS, sim-verified) [crafted]; Bristlebark Amice (14573, -1.16 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (97.5 DPS) | yes | Tigerstrike Mantle (13108, -0.33 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.33 DPS, sim-verified) [pvp]; Glowing Lizardscale Cloak (6449, -0.66 DPS) [dungeon] |
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
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (97.5 DPS) | yes | Viscous Hammer (13045, -2.77 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -3.59 DPS) [dungeon]; Gnarled Ash Staff (791, -7.41 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 149.5. Weights run: 11.0s. Verify run: 13.1s. 526 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.229 ± 0.032, crit=1.257 ± 0.039 per rating point (14 rating = 1%, 17.597 per %), hit=1.110 ± 0.088 per rating point (10 rating = 1%, 11.098 per %), melee_haste=not significant (4.364 ± 1.102)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.5 attack_power points (4.61 DPS) | yes | Hawkeye's Helm (14591, -1.40 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.74 DPS) [world_drop]; Nightscape Headband (8176, -1.96 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.5 attack_power points (2.43 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.45 DPS) [quest]; Kaleidoscope Chain (13084, -0.75 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.24 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.5 attack_power points (3.42 DPS) | yes | Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Nightscape Shoulders (8192, -0.99 DPS) [crafted]; Fleshhide Shoulders (10774, -1.10 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (149.5 DPS) | yes | Yeti Fur Cloak (2805, -0.59 DPS) [quest]; Parachute Cloak (10518, -0.75 DPS) [crafted]; Hawkeye's Cloak (14593, -1.78 DPS, sim-verified) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs [dungeon] | 58.4 attack_power points (5.78 DPS) | yes | Wolffear Harness (13110, -1.79 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -2.47 DPS) [crafted]; Barbaric Harness (5739, -2.63 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 21.4 attack_power points (2.12 DPS) | yes | Branded Leather Bracers (19508, -0.14 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.35 DPS) [dungeon]; Dusky Bracers (7378, -0.35 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (149.5 DPS) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Prowler's Leather Gloves (252524, -1.88 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 38.1 attack_power points (3.77 DPS) | yes | Skulker's Leather Belt (252520, -0.59 DPS) [crafted]; Prowler's Leather Belt (252459, -0.66 DPS) [crafted]; Scorpashi Sash (14652, -0.77 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.8 attack_power points (4.64 DPS) | yes | Triprunner Dungarees (9624, -0.90 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -1.26 DPS) [crafted]; Brawler's Leather Pants (252500, -1.46 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 38.5 attack_power points (3.82 DPS) | yes | Imperial Leather Boots (6431, -0.40 DPS) [dungeon]; Excelsior Boots (4109, -0.44 DPS) [quest]; Prowler's Leather Shoes (252465, -0.66 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 33.8 attack_power points (3.35 DPS) | yes | Falcon's Hook (7552, -0.77 DPS) [dungeon]; Thunderbrow Ring (13097, -1.10 DPS) [world_drop]; Mark of Kern (2262, -1.37 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 26.3 attack_power points (2.61 DPS) | yes | Falcon's Hook (7552, -0.02 DPS) [dungeon]; Thunderbrow Ring (13097, -0.36 DPS) [world_drop]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Thornstone Sledgehammer (1722) | World drop [world_drop] | sim-verified (149.5 DPS) | yes | Mograine's Might (7723, -0.11 DPS) [dungeon]; The Jackhammer (9423, -0.62 DPS) [dungeon]; Ironshod Bludgeon (9408, -2.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Quillward Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Thornstone Sledgehammer

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0500220000000000-50212032123012123001-0100000000000000)

Set DPS (verified): 221.9. Weights run: 11.0s. Verify run: 19.6s. 690 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.040 ± 0.002, agility=2.287 ± 0.037, crit=1.549 ± 0.050 per rating point (14 rating = 1%, 21.692 per %), hit=1.462 ± 0.115 per rating point (10 rating = 1%, 14.620 per %), melee_haste=7.944 ± 1.470

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.93 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.9 attack_power points (4.27 DPS) | yes | Sentinel's Medallion (19540, -1.58 DPS) [rep]; Zealous Shadowshard Pendant (17772, -2.13 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 43.5 attack_power points (4.65 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Prowler's Leather Shoulders (252534, -0.05 DPS) [crafted]; Failed Flying Experiment (9647, -0.48 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.73 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-verified (221.9 DPS) | yes | Grizzled Pelt (22274, +0.00 DPS) [quest]; Mixologist's Tunic (12793, -0.19 DPS) [dungeon]; Quillward Harness (10583, -0.41 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 42.5 attack_power points (4.54 DPS) | yes | Skulker's Leather Bracers (252540, -0.79 DPS, sim-verified) [crafted]; Prowler's Leather Bracers (252539, -0.87 DPS) [crafted]; Pridelord Bands (14672, -1.00 DPS) [world_drop] |
| hands | Skulker's Leather Gauntlets (252548) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.13 DPS) [crafted]; Serpentskin Gloves (8260, -0.30 DPS) [world_drop] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 56.5 attack_power points (6.04 DPS) | yes | Prowler's Leather Waistguard (252473, -0.05 DPS) [crafted]; Girdle of Beastial Fury (11686, -1.09 DPS) [dungeon]; Highlander's Leather Girdle (20115, -1.58 DPS) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 64.9 attack_power points (6.94 DPS) | yes | Serpentskin Leggings (8262, -0.19 DPS) [world_drop]; Knight's Leather Pants (220858, -1.13 DPS) [vendor]; Basilisk Hide Pants (1718, -1.80 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 52.2 attack_power points (5.58 DPS) | yes | Prowler's Leather Boots (252468, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.69 DPS) [dungeon]; Shadefiend Boots (11675, -0.71 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 41.0 attack_power points (4.38 DPS) | yes | Masons Fraternity Ring (9533, -0.96 DPS) [quest]; Ring of the Underwood (2951, -1.50 DPS) [world_drop]; Falcon's Hook (7552, -1.53 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.6 attack_power points (3.70 DPS) | yes | Masons Fraternity Ring (9533, -0.28 DPS) [quest]; Ring of the Underwood (2951, -0.82 DPS) [world_drop]; Falcon's Hook (7552, -0.85 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -4.14 DPS, sim-verified) [world_drop]; Dark Iron Pulverizer (11608, -4.47 DPS) [crafted]; Soulkeeper (1607, -5.75 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Skulker's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0500220000000000-50212032123012123001-0550100000000000)

Set DPS (verified): 385.3. Weights run: 11.4s. Verify run: 84.9s. 1755 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.378 ± 0.043, crit=1.969 ± 0.061 per rating point (14 rating = 1%, 27.566 per %), hit=1.605 ± 0.148 per rating point (10 rating = 1%, 16.049 per %), melee_haste=not significant (7.291 ± 1.858)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 91.3 attack_power points (12.03 DPS) | yes | Eye of Rend (12587, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Sentinel's Medallion (19538, -0.82 DPS) [rep]; Imperial Jewel (11933, -1.31 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 91.8 attack_power points (12.10 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -0.64 DPS) [vendor]; Lieutenant Commander's Dragonhide Shoulders (227172, -3.10 DPS) [pvp]; Highlander's Leather Shoulders (20059, -3.23 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.29 DPS) [rep]; Blackveil Cape (11626, -1.42 DPS, sim-verified) [dungeon]; Howler's Furs (272414, -1.53 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 143.0 attack_power points (18.85 DPS) | yes | Dawn Armor (252483, -2.29 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -2.91 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -3.57 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 69.4 attack_power points (9.15 DPS) | yes | Deepfury Bracers (13120, -3.28 DPS) [world_drop]; Skulker's Leather Bracers (252540, -4.29 DPS) [crafted]; Prowler's Leather Bracers (252539, -4.34 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, -0.54 DPS) [vendor]; Timbermaw Brawlers (19049, -1.18 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -2.63 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 87.9 attack_power points (11.58 DPS) | yes | Might of the Timbermaw (19044, -1.11 DPS) [crafted]; Cloudrunner Girdle (13252, -2.82 DPS) [dungeon]; Highlander's Leather Girdle (20045, -3.47 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 119.3 attack_power points (15.73 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -2.74 DPS) [vendor]; Warbear Woolies (15065, -2.95 DPS, sim-verified) [crafted] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | sim-verified (385.3 DPS) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.09 DPS) [vendor]; Painweaver Band (13098, -1.85 DPS) [dungeon]; Naglering (11669, -8.11 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.42 DPS) [vendor]; Painweaver Band (13098, -1.19 DPS) [dungeon]; Magni's Will (12548, -1.55 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Heart of Wyrmthalak (22321, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Heart of Wyrmthalak (22321, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -10.14 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -17.56 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; finger1: Don Julio's Band; finger2: Protector's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Molten Heart of the Mountain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0500220000000000-50230032023012125001-0550000000000000)

Set DPS (verified): 938.9. Weights run: 11.6s. Verify run: 71.2s. 1755 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.578 ± 0.041, crit=1.943 ± 0.056 per rating point (14 rating = 1%, 27.201 per %), hit=2.284 ± 0.164 per rating point (10 rating = 1%, 22.845 per %), melee_haste=24.460 ± 1.977

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 94.9 attack_power points (23.58 DPS) | yes | Eye of Rend (12587, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Sentinel's Medallion (19538, -1.72 DPS) [rep]; Talisman of Evasion (13177, -3.01 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 96.8 attack_power points (24.05 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -1.19 DPS) [vendor]; Lieutenant Commander's Dragonhide Shoulders (227172, -6.24 DPS) [pvp]; Highlander's Leather Shoulders (20059, -9.08 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -2.00 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.93 DPS) [rep]; Howler's Furs (272414, -5.21 DPS, sim-verified) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 153.5 attack_power points (38.12 DPS) | yes | Dawn Armor (252483, -3.66 DPS, sim-verified) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -4.70 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -7.37 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 75.6 attack_power points (18.78 DPS) | yes | Deepfury Bracers (13120, -6.77 DPS) [world_drop]; Skulker's Leather Bracers (252540, -8.81 DPS) [crafted]; Prowler's Leather Bracers (252539, -8.89 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, -1.73 DPS) [vendor]; Timbermaw Brawlers (19049, -3.27 DPS, sim-verified) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -4.01 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 95.8 attack_power points (23.79 DPS) | yes | Might of the Timbermaw (19044, -2.20 DPS) [crafted]; Cloudrunner Girdle (13252, -5.76 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -7.61 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 124.0 attack_power points (30.80 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -3.48 DPS) [vendor]; Warbear Woolies (15065, -7.12 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (938.9 DPS) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -1.65 DPS) [rep]; Naglering (11669, -20.68 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Protector's Band (19516, -2.87 DPS) [rep]; Painweaver Band (13098, -3.91 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+12.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -5.03 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -27.26 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -23.47 DPS, sim-verified) [vendor] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Molten Heart of the Mountain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 52.1. Weights run: 8.1s. Verify run: 11.4s. 224 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.907 ± 0.010, crit=0.568 ± 0.009 per rating point (14 rating = 1%, 7.954 per %), hit=0.654 ± 0.036 per rating point (10 rating = 1%, 6.545 per %), melee_haste=2.420 ± 0.455

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (1.11 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.4 attack_power points (0.79 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.5 attack_power points (0.66 DPS) | yes | Slime-encrusted Pads (6461, -0.66 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (52.1 DPS) | yes | Dark Leather Cloak (2316, -0.12 DPS) [crafted]; Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.24 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (52.1 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.32 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 9.7 attack_power points (0.67 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.28 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (52.1 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.46 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.24 DPS) | yes | Deviate Scale Belt (6468, -0.17 DPS) [crafted]; Murloc Scale Belt (5780, -0.57 DPS) [crafted]; Brawler's Leather Belt (252428, -3.12 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (52.1 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.59 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.3 attack_power points (1.47 DPS) | yes | Brawler's Leather Boots (252439, -0.12 DPS) [crafted]; Blackened Defias Boots (10402, -0.68 DPS) [dungeon]; Footpads of the Fang (10411, -0.68 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.6 attack_power points (1.08 DPS) | yes | Loop of Sacrifice (281673, -0.67 DPS) [quest]; Bounty Hunter's Ring (5351, -0.69 DPS) [quest]; The 1 Ring (8350, -0.81 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.55 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Bounty Hunter's Ring (5351, -0.16 DPS) [quest]; The 1 Ring (8350, -0.28 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.6 attack_power points (21.14 DPS) | yes | Hammerbone (270018, -0.56 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -1.35 DPS) [world_drop]; Rakzur Club (12983, -2.89 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 96.0. Weights run: 8.6s. Verify run: 19.2s. 401 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.995 ± 0.013, crit=0.796 ± 0.014 per rating point (14 rating = 1%, 11.145 per %), hit=0.929 ± 0.057 per rating point (10 rating = 1%, 9.285 per %), melee_haste=3.761 ± 0.713

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.99 DPS) | yes | Brawler's Leather Helm (252512, -0.32 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.50 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.0 attack_power points (1.33 DPS) | yes | Scout's Medallion (19537, -0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.33 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.9 attack_power points (2.65 DPS) | yes | Mantle of Thieves (2264, -1.00 DPS) [dungeon]; Barbaric Shoulders (5964, -1.06 DPS, sim-verified) [crafted]; Bristlebark Amice (14573, -1.16 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (96.0 DPS) | yes | Tigerstrike Mantle (13108, -0.35 DPS, sim-verified) [world_drop]; Glowing Lizardscale Cloak (6449, -0.66 DPS) [dungeon]; Swiftrunner Cape (6745, -0.66 DPS) [quest] |
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
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (96.0 DPS) | yes | Viscous Hammer (13045, -2.57 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -3.59 DPS) [dungeon]; Gnarled Ash Staff (791, -7.41 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 145.8. Weights run: 11.0s. Verify run: 15.3s. 525 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.229 ± 0.032, crit=1.257 ± 0.039 per rating point (14 rating = 1%, 17.597 per %), hit=1.110 ± 0.088 per rating point (10 rating = 1%, 11.098 per %), melee_haste=not significant (4.364 ± 1.102)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.5 attack_power points (4.61 DPS) | yes | Hawkeye's Helm (14591, -1.36 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.74 DPS) [world_drop]; Nightscape Headband (8176, -1.96 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.5 attack_power points (2.43 DPS) | yes | Ethereal Talisman (4430, -0.56 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.62 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.75 DPS) [world_drop] |
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

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0500220000000000-50212032123012123001-0100000000000000)

Set DPS (verified): 217.2. Weights run: 11.0s. Verify run: 19.0s. 689 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.040 ± 0.002, agility=2.287 ± 0.037, crit=1.549 ± 0.050 per rating point (14 rating = 1%, 21.692 per %), hit=1.462 ± 0.115 per rating point (10 rating = 1%, 14.620 per %), melee_haste=7.944 ± 1.470

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Embrace of the Lycan (9479, +0.00 DPS) [dungeon]; Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.22 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 39.9 attack_power points (4.27 DPS) | yes | Woven Ivy Necklace (19159, -1.05 DPS, sim-verified) [quest]; Scout's Medallion (19535, -1.34 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 43.5 attack_power points (4.65 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.03 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.05 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.73 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-verified (217.2 DPS) | yes | Grizzled Pelt (22274, +0.00 DPS) [quest]; Mixologist's Tunic (12793, -0.19 DPS) [dungeon]; Quillward Harness (10583, -0.41 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 42.5 attack_power points (4.54 DPS) | yes | Skulker's Leather Bracers (252540, -0.81 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.87 DPS) [crafted]; Pridelord Bands (14672, -1.00 DPS) [world_drop] |
| hands | Skulker's Leather Gauntlets (252548) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, -0.13 DPS) [crafted]; Serpentskin Gloves (8260, -0.30 DPS) [world_drop] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 56.5 attack_power points (6.04 DPS) | yes | Prowler's Leather Waistguard (252473, -0.05 DPS) [crafted]; Girdle of Beastial Fury (11686, -1.09 DPS) [dungeon]; Defiler's Leather Girdle (20193, -1.58 DPS) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 63.1 attack_power points (6.75 DPS) | yes | Stone Guard's Leather Pants (220859, -0.94 DPS) [vendor]; Basilisk Hide Pants (1718, -1.61 DPS) [world_drop]; Triprunner Dungarees (9624, -1.69 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 52.2 attack_power points (5.58 DPS) | yes | Prowler's Leather Boots (252468, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.69 DPS) [dungeon]; Shadefiend Boots (11675, -0.71 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 41.0 attack_power points (4.38 DPS) | yes | Masons Fraternity Ring (9533, -0.96 DPS) [quest]; Ring of the Underwood (2951, -1.50 DPS) [world_drop]; Falcon's Hook (7552, -1.53 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.6 attack_power points (3.70 DPS) | yes | Masons Fraternity Ring (9533, -0.28 DPS) [quest]; Ring of the Underwood (2951, -0.82 DPS) [world_drop]; Falcon's Hook (7552, -0.85 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.18 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -1.85 DPS, sim-verified) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -1.28 DPS) [world_drop]; Dark Iron Pulverizer (11608, -4.47 DPS) [crafted]; Ragehammer (10626, -4.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Skulker's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0500220000000000-50212032123012123001-0550100000000000)

Set DPS (verified): 377.5. Weights run: 11.4s. Verify run: 92.0s. 1759 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.378 ± 0.043, crit=1.969 ± 0.061 per rating point (14 rating = 1%, 27.566 per %), hit=1.605 ± 0.148 per rating point (10 rating = 1%, 16.049 per %), melee_haste=not significant (7.291 ± 1.858)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 91.3 attack_power points (12.03 DPS) | yes | Eye of Rend (12587, +0.00 DPS) [dungeon]; Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Scout's Medallion (19534, -0.82 DPS) [rep]; Woven Ivy Necklace (19159, -0.96 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 91.8 attack_power points (12.10 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -0.64 DPS) [vendor]; Defiler's Leather Shoulders (20194, -2.86 DPS, sim-verified) [rep]; Champion's Dragonhide Shoulders (227175, -3.10 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, -1.21 DPS) [dungeon]; Deathguard's Cloak (20068, -1.29 DPS) [rep]; Howler's Furs (272414, -1.53 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 143.0 attack_power points (18.85 DPS) | yes | Dawn Armor (252483, -1.50 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -2.91 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -3.57 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 69.4 attack_power points (9.15 DPS) | yes | Deepfury Bracers (13120, -3.28 DPS) [world_drop]; Skulker's Leather Bracers (252540, -4.29 DPS) [crafted]; Prowler's Leather Bracers (252539, -4.34 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, -0.54 DPS) [vendor]; Timbermaw Brawlers (19049, -1.18 DPS) [crafted]; Blood Guard's Dragonhide Grips (227180, -2.63 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 87.9 attack_power points (11.58 DPS) | yes | Might of the Timbermaw (19044, -1.38 DPS, sim-verified) [crafted]; Cloudrunner Girdle (13252, -2.82 DPS) [dungeon]; Defiler's Leather Girdle (20190, -3.47 DPS) [rep] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 104.4 attack_power points (13.76 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -0.77 DPS) [pvp]; Outrider's Leather Pants (22740, -1.35 DPS) [rep] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -1.09 DPS) [vendor]; Painweaver Band (13098, -1.85 DPS) [dungeon]; Naglering (11669, -7.25 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cutthroat's Signet (272408, -0.42 DPS) [vendor]; Painweaver Band (13098, -1.19 DPS) [dungeon]; Myrmidon's Signet (2246, -1.83 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Heart of Wyrmthalak (22321, -3.32 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (377.5 DPS) | yes | Hand of Justice (11815, -0.69 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.42 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -8.55 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -16.16 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Warbear Woolies; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0500220000000000-50230032023012125001-0550000000000000)

Set DPS (verified): 917.3. Weights run: 11.6s. Verify run: 75.1s. 1759 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.578 ± 0.041, crit=1.943 ± 0.056 per rating point (14 rating = 1%, 27.201 per %), hit=2.284 ± 0.164 per rating point (10 rating = 1%, 22.845 per %), melee_haste=24.460 ± 1.977

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 94.9 attack_power points (23.58 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Eye of Rend (12587, -3.28 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Scout's Medallion (19534, -1.72 DPS) [rep]; Woven Ivy Necklace (19159, -1.96 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 96.8 attack_power points (24.05 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -1.19 DPS) [vendor]; Champion's Dragonhide Shoulders (227175, -6.24 DPS) [pvp]; Defiler's Leather Shoulders (20194, -6.43 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.95 DPS) [vendor]; Blackveil Cape (11626, -2.00 DPS) [dungeon]; Deathguard's Cloak (20068, -2.93 DPS) [rep] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 153.5 attack_power points (38.12 DPS) | yes | Dawn Armor (252483, -1.60 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -4.70 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -7.37 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 75.6 attack_power points (18.78 DPS) | yes | Deepfury Bracers (13120, -6.77 DPS) [world_drop]; Skulker's Leather Bracers (252540, -8.81 DPS) [crafted]; Prowler's Leather Bracers (252539, -8.89 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, -1.73 DPS) [vendor]; Timbermaw Brawlers (19049, -3.24 DPS, sim-verified) [crafted]; Blood Guard's Dragonhide Grips (227180, -4.01 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 95.8 attack_power points (23.79 DPS) | yes | Might of the Timbermaw (19044, -3.69 DPS, sim-verified) [crafted]; Cloudrunner Girdle (13252, -5.76 DPS) [dungeon]; Skulker's Leather Waistguard (252474, -7.61 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 114.2 attack_power points (28.36 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -1.03 DPS) [pvp]; Plaguehound Leggings (18736, -3.47 DPS) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-verified (917.3 DPS) | yes | Swiftwalker Boots (12553, +0.00 DPS) [dungeon]; Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -1.65 DPS) [rep]; Naglering (11669, -14.48 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -1.60 DPS, sim-verified) [rep]; Painweaver Band (13098, -3.91 DPS) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -2.34 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -26.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howling Idol (272427, -16.65 DPS, sim-verified) [vendor] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Hand of Justice; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

