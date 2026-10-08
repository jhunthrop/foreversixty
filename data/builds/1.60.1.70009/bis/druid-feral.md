# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 67.6. Weights run: 2.6s. Verify run: 3.3s. 193 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.100 ± 0.013, crit=0.686 ± 0.012 per rating point (14 rating = 1%, 9.603 per %), hit=0.958 ± 0.030 per rating point (10 rating = 1%, 9.576 per %), melee_haste=5.988 ± 0.357

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.95 DPS) | yes | Brawler's Leather Hood (252504, -0.04 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.6 attack_power points (0.68 DPS) | yes | Erudite's Amulet (277204, -0.24 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.5 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.62 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.6 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (67.6 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.05 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.8 attack_power points (0.69 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Forest Leather Bracers (3202, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.6 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.19 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.97 DPS) | yes | Deviate Scale Belt (6468, -0.05 DPS) [crafted]; Ruffian Belt (5975, -0.26 DPS) [world]; Brawler's Leather Belt (252428, -3.71 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.6 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.30 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.4 attack_power points (1.26 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 17.2 attack_power points (0.93 DPS) | yes | Demon Band (12054, -0.45 DPS) [world_drop]; The 1 Ring (8350, -0.70 DPS) [world]; Lavishly Jeweled Ring (1156, -0.70 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.6 attack_power points (0.68 DPS) | yes | Demon Band (12054, -0.25 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.45 DPS) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.5 attack_power points (16.65 DPS) | yes | Staff of Westfall (2042, -1.16 DPS) [quest]; Twisted Chanter's Staff (890, -1.21 DPS) [world_drop]; Living Root (6631, -1.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 122.3. Weights run: 2.8s. Verify run: 1.9s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.985 ± 0.015, crit=0.781 ± 0.016 per rating point (14 rating = 1%, 10.940 per %), hit=0.954 ± 0.032 per rating point (10 rating = 1%, 9.540 per %), melee_haste=5.539 ± 0.450

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.71 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.43 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.43 DPS) [crafted]; Brawler's Leather Helm (252512, -0.49 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.7 attack_power points (1.08 DPS) | yes | Sentinel's Medallion (19541, -0.06 DPS) [rep]; Ghostshard Talisman (7731, -0.18 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (2.13 DPS) | yes | Mantle of Thieves (2264, -0.84 DPS) [dungeon]; Bristlebark Amice (14573, -0.93 DPS) [world_drop]; Barbaric Shoulders (5964, -1.12 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.33 DPS) | yes | Sergeant Major's Cape (16315, -0.24 DPS) [pvp]; Tigerstrike Mantle (13108, -0.30 DPS) [world_drop]; Cloak of Night (4447, -0.56 DPS) [world] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.1 attack_power points (1.88 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.7 attack_power points (1.34 DPS) | yes | Jurassic Wristguards (6198, -0.28 DPS) [world]; Barbaric Bracers (18948, -0.37 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.40 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.3 attack_power points (1.77 DPS) | yes | Toughened Leather Gloves (4253, -0.14 DPS) [crafted]; Wolfclaw Gloves (1978, -0.28 DPS) [dungeon]; Gloves of the Fang (10413, -0.43 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 31.7 attack_power points (2.05 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Warden's Leather Belt (252460, -0.43 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.50 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 33.3 attack_power points (2.15 DPS) | yes | Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Trapper's Leather Pants (252501, -0.14 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.5 attack_power points (1.46 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Insignia Boots (4055, -0.43 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.43 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.7 attack_power points (1.73 DPS) | yes | Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Monkey Ring (6748, -0.83 DPS) [quest]; Tiger Band (6749, -0.87 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 25.1 attack_power points (1.63 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Monkey Ring (6748, -0.73 DPS) [quest]; Tiger Band (6749, -0.77 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (122.3 DPS) | yes | Cobalt Crusher (7730, -3.00 DPS) [dungeon]; Wind Spirit Staff (6689, -5.03 DPS) [dungeon]; Viscous Hammer (13045, -24.25 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 168.2. Weights run: 3.2s. Verify run: 2.0s. 438 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.936 ± 0.022, crit=0.953 ± 0.028 per rating point (14 rating = 1%, 13.348 per %), hit=0.997 ± 0.040 per rating point (10 rating = 1%, 9.969 per %), melee_haste=5.690 ± 0.585

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.5 attack_power points (3.35 DPS) | yes | Defender's Leather Helm (252455, -1.41 DPS) [crafted]; Warden's Wizard Hat (14604, -1.50 DPS) [world_drop]; Hawkeye's Helm (14591, -1.55 DPS, sim-verified) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 21.3 attack_power points (1.57 DPS) | yes | Kaleidoscope Chain (13084, -0.35 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.50 DPS, sim-verified) [quest]; Ghostshard Talisman (7731, -0.54 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 33.3 attack_power points (2.45 DPS) | yes | Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop]; Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.79 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 28.2 attack_power points (2.07 DPS) | yes | Sergeant Major's Cape (16336, -0.25 DPS) [pvp]; Hawkeye's Cloak (14593, -0.59 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.73 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 54.4 attack_power points (4.00 DPS) | yes | Barbaric Harness (5739, -1.73 DPS) [crafted]; Nightscape Tunic (8175, -1.87 DPS) [crafted]; Wolffear Harness (13110, -2.03 DPS, sim-verified) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.50 DPS) | yes | Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Barbaric Bracers (18948, -0.28 DPS) [crafted]; Jurassic Wristguards (6198, -0.32 DPS) [world] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 39.4 attack_power points (2.90 DPS) | yes | Skulker's Leather Gloves (252525, -0.02 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.45 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.2 attack_power points (2.74 DPS) | yes | Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.53 DPS) [rep] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 41.4 attack_power points (3.05 DPS) | yes | Basilisk Hide Pants (1718, -0.06 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.63 DPS) [crafted]; Brawler's Leather Pants (252500, -0.80 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.7 attack_power points (2.78 DPS) | yes | Skulker's Leather Shoes (252531, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.36 DPS) [quest]; Imperial Leather Boots (6431, -0.40 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 33.1 attack_power points (2.44 DPS) | yes | Falcon's Hook (7552, -0.67 DPS) [dungeon]; Ring of the Underwood (2951, -0.69 DPS) [world_drop]; Thunderbrow Ring (13097, -0.71 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.2 attack_power points (1.93 DPS) | yes | Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.18 DPS) [world_drop]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (168.2 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; The Jackhammer (9423, -30.34 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 185.9. Weights run: 3.1s. Verify run: 2.5s. 577 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.852 ± 0.024, crit=1.044 ± 0.032 per rating point (14 rating = 1%, 14.612 per %), hit=1.099 ± 0.053 per rating point (10 rating = 1%, 10.987 per %), melee_haste=5.619 ± 0.755

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (3.92 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.63 DPS) [vendor]; White Bandit Mask (10008, -1.02 DPS, sim-verified) [crafted]; Scorpashi Skullcap (14658, -1.35 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.1 attack_power points (2.77 DPS) | yes | Sentinel's Medallion (19540, -1.16 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.19 DPS) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 40.9 attack_power points (3.23 DPS) | yes | Skulker's Leather Shoulder (252535, -0.06 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.57 DPS) [vendor]; Failed Flying Experiment (9647, -1.03 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 39.1 attack_power points (3.09 DPS) | yes | Dark Phantom Cape (13122, -0.72 DPS, sim-verified) [world_drop]; Blisterbane Wrap (12552, -0.90 DPS) [dungeon]; Dark Hooded Cape (5257, -0.93 DPS) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 67.0 attack_power points (5.30 DPS) | yes | Mixologist's Tunic (12793, -0.56 DPS) [dungeon]; Warbear Harness (15064, -0.75 DPS) [crafted]; Quillward Harness (10583, -1.13 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.6 attack_power points (2.89 DPS) | yes | Prowler's Leather Bracers (252539, -0.30 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.56 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 62.7 attack_power points (4.96 DPS) | yes | Feralheart Fists (226793, -1.05 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -1.21 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.35 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.0 attack_power points (4.19 DPS) | yes | Skulker's Leather Waistguard (252474, -0.06 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.43 DPS) [dungeon]; Ogron's Sash (13117, -1.31 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 60.8 attack_power points (4.81 DPS) | yes | Serpentskin Leggings (8262, -0.20 DPS) [world_drop]; Knight's Leather Pants (220858, -1.36 DPS) [vendor]; Triprunner Dungarees (9624, -1.65 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.0 attack_power points (3.87 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.34 DPS) [dungeon]; Shadefiend Boots (11675, -0.52 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 38.7 attack_power points (3.06 DPS) | yes | Masons Fraternity Ring (9533, -1.01 DPS) [quest]; Ironspine's Eye (7686, -1.04 DPS) [dungeon]; Falcon's Hook (7552, -1.22 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 31.0 attack_power points (2.45 DPS) | yes | Masons Fraternity Ring (9533, -0.40 DPS) [quest]; Ironspine's Eye (7686, -0.44 DPS) [dungeon]; Falcon's Hook (7552, -0.61 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (185.9 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (185.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (185.9 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; The Jackhammer (9423, -2.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 287.0. Weights run: 2.9s. Verify run: 8.1s. 1449 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.813 ± 0.029, crit=1.161 ± 0.041 per rating point (14 rating = 1%, 16.249 per %), hit=1.277 ± 0.068 per rating point (10 rating = 1%, 12.766 per %), melee_haste=6.349 ± 0.994

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 82.7 attack_power points (7.05 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -12.44 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.4 attack_power points (4.81 DPS) | yes | Medallion of the Dawn (22659, -1.38 DPS) [quest]; Pendant of Celerity (22340, -1.41 DPS) [dungeon]; Mark of Fordring (15411, -1.42 DPS, sim-verified) [quest] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (287.0 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -8.68 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.2 attack_power points (4.02 DPS) | yes | Windshear Cape (20691, -0.20 DPS) [world]; Cloak of the Honor Guard (20073, -0.35 DPS) [rep]; Howler's Furs (272414, -0.55 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (287.0 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -1.97 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.12 DPS) [vendor]; Tunic of Undead Slaying (23089, -11.77 DPS, sim-verified) [world] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (287.0 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 73.1 attack_power points (6.24 DPS) | yes | Timbermaw Brawlers (19049, -0.13 DPS) [crafted]; Marshal's Dragonhide Grips (231694, -0.49 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.94 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 88.0 attack_power points (7.50 DPS) | yes | Shifter's Belt (272396, -1.33 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -1.40 DPS) [crafted]; Belt of Preserved Heads (20216, -1.47 DPS) [quest] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 94.2 attack_power points (8.03 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Sentinel's Leather Pants (237818, -1.09 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.45 DPS) [vendor] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (287.0 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Drudge Boots (21532, -9.60 DPS, sim-verified) [quest] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (287.0 DPS) | yes | Don Julio's Band (19325, -0.11 DPS) [rep]; Tarnished Elven Ring (18500, -0.54 DPS) [dungeon]; Naglering (11669, -4.57 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (287.0 DPS) | yes | Don Julio's Band (19325, -0.03 DPS) [rep]; Tarnished Elven Ring (18500, -0.46 DPS) [dungeon]; Naglering (11669, -3.42 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (287.0 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Hand of Justice (11815, -3.61 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (287.0 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, -0.89 DPS) [dungeon]; Hand of Justice (11815, -1.07 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (287.0 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Felstriker (12590, -3.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (287.0 DPS) | yes | Howling Idol (272427, -8.47 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Feralheart Walkers; finger1: Protector's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 654.0. Weights run: 3.0s. Verify run: 8.0s. 1449 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.301 ± 0.034, crit=1.563 ± 0.047 per rating point (14 rating = 1%, 21.877 per %), hit=1.729 ± 0.090 per rating point (10 rating = 1%, 17.290 per %), melee_haste=11.365 ± 1.216

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 99.8 attack_power points (15.30 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -4.56 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 67.9 attack_power points (10.41 DPS) | yes | Pendant of Celerity (22340, -3.04 DPS, sim-verified) [dungeon]; Mark of Fordring (15411, -3.07 DPS) [quest]; Medallion of the Dawn (22659, -3.38 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (654.0 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -22.07 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.5 attack_power points (8.35 DPS) | yes | Windshear Cape (20691, +0.00 DPS, sim-verified) [world]; Cloak of Revanchion (23127, -1.12 DPS) [dungeon]; Blackveil Cape (11626, -1.19 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (654.0 DPS) | yes | Dawn Armor (252483, -2.89 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.58 DPS) [vendor]; Tunic of Undead Slaying (23089, -25.47 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (654.0 DPS) | yes | Bracers of Subterfuge (22668, -1.78 DPS) [quest]; Bracers of the Eclipse (18375, -3.58 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.39 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (654.0 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, -25.27 DPS, sim-verified) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 102.0 attack_power points (15.64 DPS) | yes | Shifter's Belt (272396, -2.19 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -2.51 DPS) [quest]; Might of the Timbermaw (19044, -2.91 DPS) [crafted] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (654.0 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Sentinel's Leather Pants (237818, +0.00 DPS) [vendor]; Warbear Woolies (15065, -22.25 DPS, sim-verified) [crafted] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | 70.6 attack_power points (10.82 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Knight-Lieutenant's Dragonhide Treads (227182, -0.21 DPS) [pvp]; Boots of Ferocity (22472, -1.41 DPS) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (654.0 DPS) | yes | Protector's Band (19514, -0.67 DPS) [rep]; Tarnished Elven Ring (18500, -1.06 DPS) [dungeon]; Naglering (11669, -14.22 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (654.0 DPS) | yes | Protector's Band (19514, -0.12 DPS) [rep]; Tarnished Elven Ring (18500, -0.51 DPS) [dungeon]; Naglering (11669, -13.61 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (654.0 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -5.62 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (654.0 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -3.46 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (654.0 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Felstriker (12590, -12.11 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (654.0 DPS) | yes | Howling Idol (272427, -5.99 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Ferocity of the Timbermaw; legs: Stormshroud Pants; feet: Drudge Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 66.2. Weights run: 2.6s. Verify run: 3.2s. 183 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.100 ± 0.013, crit=0.686 ± 0.012 per rating point (14 rating = 1%, 9.603 per %), hit=0.958 ± 0.030 per rating point (10 rating = 1%, 9.576 per %), melee_haste=5.988 ± 0.357

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.95 DPS) | yes | Brawler's Leather Hood (252504, -0.04 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.6 attack_power points (0.68 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.5 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.6 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.10 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (66.2 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.11 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 10.7 attack_power points (0.58 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.2 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.23 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.97 DPS) | yes | Deviate Scale Belt (6468, -0.05 DPS) [crafted]; Ruffian Belt (5975, -0.26 DPS) [world]; Brawler's Leather Belt (252428, -3.75 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.2 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.34 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.4 attack_power points (1.26 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.58 DPS) [dungeon]; Footpads of the Fang (10411, -0.58 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 17.2 attack_power points (0.93 DPS) | yes | Demon Band (12054, -0.45 DPS) [world_drop]; Loop of Sacrifice (281673, -0.57 DPS) [quest]; Bounty Hunter's Ring (5351, -0.59 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.6 attack_power points (0.68 DPS) | yes | Demon Band (12054, -0.29 DPS, sim-verified) [world_drop]; Loop of Sacrifice (281673, -0.32 DPS) [quest]; Bounty Hunter's Ring (5351, -0.34 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.5 attack_power points (16.65 DPS) | yes | Crescent Staff (6505, -0.49 DPS) [quest]; Living Root (6631, -0.63 DPS) [dungeon]; Hammerbone (270018, -0.63 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 120.1. Weights run: 2.8s. Verify run: 2.0s. 315 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.985 ± 0.015, crit=0.781 ± 0.016 per rating point (14 rating = 1%, 10.940 per %), hit=0.954 ± 0.032 per rating point (10 rating = 1%, 9.540 per %), melee_haste=5.539 ± 0.450

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.71 DPS) | yes | Brawler's Leather Helm (252512, -0.40 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.43 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.43 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.7 attack_power points (1.08 DPS) | yes | Scout's Medallion (19537, -0.06 DPS) [rep]; Ghostshard Talisman (7731, -0.18 DPS) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.8 attack_power points (2.13 DPS) | yes | Mantle of Thieves (2264, -0.84 DPS) [dungeon]; Bristlebark Amice (14573, -0.93 DPS) [world_drop]; Barbaric Shoulders (5964, -1.11 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.5 attack_power points (1.33 DPS) | yes | Tigerstrike Mantle (13108, -0.37 DPS, sim-verified) [world_drop]; Cloak of Night (4447, -0.56 DPS) [world]; Swiftrunner Cape (6745, -0.56 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.1 attack_power points (1.88 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.7 attack_power points (1.34 DPS) | yes | Jurassic Wristguards (6198, -0.28 DPS) [world]; Barbaric Bracers (18948, -0.36 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.49 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.3 attack_power points (1.77 DPS) | yes | Toughened Leather Gloves (4253, -0.14 DPS) [crafted]; Wolfclaw Gloves (1978, -0.28 DPS) [dungeon]; Gloves of the Fang (10413, -0.43 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 31.7 attack_power points (2.05 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Warden's Leather Belt (252460, -0.43 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.50 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 33.3 attack_power points (2.15 DPS) | yes | Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Trapper's Leather Pants (252501, -0.14 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.5 attack_power points (1.46 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Stomping Boots (3741, -0.37 DPS) [quest]; Insignia Boots (4055, -0.43 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.7 attack_power points (1.73 DPS) | yes | Thunderbrow Ring (13097, -0.20 DPS) [world_drop]; Band of the Fist (17694, -0.64 DPS) [quest]; Monkey Ring (6748, -0.83 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 25.1 attack_power points (1.63 DPS) | yes | Thunderbrow Ring (13097, -0.10 DPS) [world_drop]; Band of the Fist (17694, -0.54 DPS) [quest]; Monkey Ring (6748, -0.73 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (120.1 DPS) | yes | Cobalt Crusher (7730, -3.00 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -4.04 DPS) [pvp]; Viscous Hammer (13045, -23.67 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 165.8. Weights run: 3.2s. Verify run: 2.0s. 426 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.936 ± 0.022, crit=0.953 ± 0.028 per rating point (14 rating = 1%, 13.348 per %), hit=0.997 ± 0.040 per rating point (10 rating = 1%, 9.969 per %), melee_haste=5.690 ± 0.585

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.5 attack_power points (3.35 DPS) | yes | Defender's Leather Helm (252455, -1.41 DPS) [crafted]; Warden's Wizard Hat (14604, -1.50 DPS) [world_drop]; Hawkeye's Helm (14591, -1.52 DPS, sim-verified) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 21.3 attack_power points (1.57 DPS) | yes | Ethereal Talisman (4430, -0.19 DPS) [quest]; Kaleidoscope Chain (13084, -0.35 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -0.64 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 33.3 attack_power points (2.45 DPS) | yes | Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop]; Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.79 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 28.2 attack_power points (2.07 DPS) | yes | First Sergeant's Cloak (16340, -0.25 DPS) [pvp]; Hawkeye's Cloak (14593, -0.59 DPS) [world_drop]; Parachute Cloak (10518, -0.93 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 54.4 attack_power points (4.00 DPS) | yes | Barbaric Harness (5739, -1.73 DPS) [crafted]; Nightscape Tunic (8175, -1.87 DPS) [crafted]; Wolffear Harness (13110, -2.00 DPS, sim-verified) [world_drop] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.4 attack_power points (1.50 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.03 DPS) [dungeon]; Barbaric Bracers (18948, -0.28 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 39.4 attack_power points (2.90 DPS) | yes | Skulker's Leather Gloves (252525, -0.02 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.45 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.2 attack_power points (2.74 DPS) | yes | Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.53 DPS) [rep] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 41.4 attack_power points (3.05 DPS) | yes | Basilisk Hide Pants (1718, -0.06 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.63 DPS) [crafted]; Brawler's Leather Pants (252500, -0.80 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.7 attack_power points (2.78 DPS) | yes | Excelsior Boots (4109, -0.36 DPS) [quest]; Imperial Leather Boots (6431, -0.40 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 33.1 attack_power points (2.44 DPS) | yes | Falcon's Hook (7552, -0.67 DPS) [dungeon]; Ring of the Underwood (2951, -0.69 DPS) [world_drop]; Thunderbrow Ring (13097, -0.71 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 26.2 attack_power points (1.93 DPS) | yes | Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.18 DPS) [world_drop]; Thunderbrow Ring (13097, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (165.8 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; The Jackhammer (9423, -29.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 190.2. Weights run: 3.1s. Verify run: 2.3s. 561 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.852 ± 0.024, crit=1.044 ± 0.032 per rating point (14 rating = 1%, 14.612 per %), hit=1.099 ± 0.053 per rating point (10 rating = 1%, 10.987 per %), melee_haste=5.619 ± 0.755

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (3.92 DPS) | yes | Undercity Reservist's Cap (20643, -0.57 DPS) [quest]; Blood Guard's Leather Headband (220851, -0.63 DPS) [vendor]; White Bandit Mask (10008, -1.26 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.1 attack_power points (2.77 DPS) | yes | Woven Ivy Necklace (19159, -0.89 DPS, sim-verified) [quest]; Scout's Medallion (19535, -1.02 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 40.9 attack_power points (3.23 DPS) | yes | Skulker's Leather Shoulder (252535, -0.06 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.57 DPS) [vendor]; Failed Flying Experiment (9647, -0.60 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 39.1 attack_power points (3.09 DPS) | yes | Blisterbane Wrap (12552, -0.90 DPS) [dungeon]; Dark Phantom Cape (13122, -0.92 DPS, sim-verified) [world_drop]; Dark Hooded Cape (5257, -0.93 DPS) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 67.0 attack_power points (5.30 DPS) | yes | Mixologist's Tunic (12793, -0.62 DPS, sim-verified) [dungeon]; Warbear Harness (15064, -0.75 DPS) [crafted]; Quillward Harness (10583, -1.13 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.6 attack_power points (2.89 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Prowler's Leather Bracers (252539, -0.72 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 62.7 attack_power points (4.96 DPS) | yes | Feralheart Fists (226793, -1.17 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -1.21 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.35 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.0 attack_power points (4.19 DPS) | yes | Skulker's Leather Waistguard (252474, -0.06 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.43 DPS) [dungeon]; Ogron's Sash (13117, -1.31 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 58.2 attack_power points (4.61 DPS) | yes | Stone Guard's Leather Pants (220859, -1.32 DPS, sim-verified) [vendor]; Triprunner Dungarees (9624, -1.45 DPS) [quest]; Basilisk Hide Pants (1718, -1.53 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.0 attack_power points (3.87 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.34 DPS) [dungeon]; Shadefiend Boots (11675, -0.52 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 38.7 attack_power points (3.06 DPS) | yes | Masons Fraternity Ring (9533, -1.01 DPS) [quest]; Ironspine's Eye (7686, -1.04 DPS) [dungeon]; White Bone Band (11862, -1.16 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 31.0 attack_power points (2.45 DPS) | yes | Ironspine's Eye (7686, -0.44 DPS) [dungeon]; White Bone Band (11862, -0.55 DPS) [quest]; Masons Fraternity Ring (9533, -0.64 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (190.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (190.2 DPS) | yes | Molten Heart of the Mountain (249470, -0.96 DPS, sim-verified) [crafted] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (190.2 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; The Jackhammer (9423, -2.01 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 287.5. Weights run: 2.9s. Verify run: 8.0s. 1446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.813 ± 0.029, crit=1.161 ± 0.041 per rating point (14 rating = 1%, 16.249 per %), hit=1.277 ± 0.068 per rating point (10 rating = 1%, 12.766 per %), melee_haste=6.349 ± 0.994

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 82.7 attack_power points (7.05 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -13.88 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.4 attack_power points (4.81 DPS) | yes | Mark of Fordring (15411, -1.37 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -1.38 DPS) [quest]; Pendant of Celerity (22340, -1.41 DPS) [dungeon] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (287.5 DPS) | yes | Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Stormshroud Shoulders (15058, -14.38 DPS, sim-verified) [crafted] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.2 attack_power points (4.02 DPS) | yes | Windshear Cape (20691, -0.20 DPS) [world]; Deathguard's Cloak (20068, -0.35 DPS) [rep]; Howler's Furs (272414, -0.55 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (287.5 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, -1.97 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.12 DPS) [vendor]; Tunic of Undead Slaying (23089, -12.66 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (287.5 DPS) | yes | Bracers of Subterfuge (22668, -0.81 DPS) [quest]; Bracers of the Eclipse (18375, -1.41 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.45 DPS, sim-verified) [world] |
| hands | Feralheart Fists (226793) | Mokvar [vendor] | sim-verified (287.5 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Stormshroud Gloves (21278, -13.58 DPS, sim-verified) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 88.0 attack_power points (7.50 DPS) | yes | Shifter's Belt (272396, -0.98 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -1.40 DPS) [crafted]; Belt of Preserved Heads (20216, -1.47 DPS) [quest] |
| legs | Feralheart Trousers (226791) | Mokvar [vendor] | sim-verified (287.5 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Stormshroud Pants (15057, -13.61 DPS, sim-verified) [crafted] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | 59.4 attack_power points (5.06 DPS) | yes | General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Blood Guard's Dragonhide Treads (227181, -0.23 DPS) [pvp]; Boots of Ferocity (22472, -0.92 DPS, sim-verified) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (287.5 DPS) | yes | Don Julio's Band (19325, -0.11 DPS) [rep]; Tarnished Elven Ring (18500, -0.54 DPS) [dungeon]; Naglering (11669, -4.77 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (287.5 DPS) | yes | Don Julio's Band (19325, -0.03 DPS) [rep]; Tarnished Elven Ring (18500, -0.46 DPS) [dungeon]; Naglering (11669, -6.39 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (287.5 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Hand of Justice (11815, -4.82 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (287.5 DPS) | yes | Eye of the Beast (13968, -1.57 DPS) [quest]; Blackhand's Breadth (13965, -2.31 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, -2.47 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (287.5 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -4.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (287.5 DPS) | yes | Howling Idol (272427, -8.64 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Feralheart Fists; waist: Ferocity of the Timbermaw; legs: Feralheart Trousers; feet: Drudge Boots; finger1: Legionnaire's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 663.8. Weights run: 3.0s. Verify run: 7.7s. 1446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.301 ± 0.034, crit=1.563 ± 0.047 per rating point (14 rating = 1%, 21.877 per %), hit=1.729 ± 0.090 per rating point (10 rating = 1%, 17.290 per %), melee_haste=11.365 ± 1.216

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 99.8 attack_power points (15.30 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -25.65 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 67.9 attack_power points (10.41 DPS) | yes | Pendant of Celerity (22340, -1.88 DPS, sim-verified) [dungeon]; Mark of Fordring (15411, -3.07 DPS) [quest]; Medallion of the Dawn (22659, -3.38 DPS) [quest] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (663.8 DPS) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -17.83 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.5 attack_power points (8.35 DPS) | yes | Windshear Cape (20691, +0.00 DPS, sim-verified) [world]; Cloak of Revanchion (23127, -1.12 DPS) [dungeon]; Blackveil Cape (11626, -1.19 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (663.8 DPS) | yes | Dawn Armor (252483, -2.89 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.58 DPS) [vendor]; Tunic of Undead Slaying (23089, -24.89 DPS, sim-verified) [world] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (663.8 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 88.8 attack_power points (13.60 DPS) | yes | Timbermaw Brawlers (19049, -0.88 DPS) [crafted]; General's Dragonhide Grips (231688, -1.16 DPS) [vendor]; Blood Guard's Dragonhide Grips (227180, -2.60 DPS) [pvp] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 102.0 attack_power points (15.64 DPS) | yes | Shifter's Belt (272396, -2.05 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -2.51 DPS) [quest]; Might of the Timbermaw (19044, -2.91 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.2 attack_power points (16.73 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -0.50 DPS) [vendor]; Legionnaire's Dragonhide Leggings (227177, -2.05 DPS) [pvp] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (663.8 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Drudge Boots (21532, -19.25 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (663.8 DPS) | yes | Legionnaire's Band (19510, -0.67 DPS) [rep]; Tarnished Elven Ring (18500, -1.06 DPS) [dungeon]; Naglering (11669, -10.95 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (663.8 DPS) | yes | Legionnaire's Band (19510, -0.12 DPS) [rep]; Tarnished Elven Ring (18500, -0.51 DPS) [dungeon]; Naglering (11669, -10.29 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (663.8 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -3.46 DPS, sim-verified) [quest] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (663.8 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -2.56 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (663.8 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -9.00 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (663.8 DPS) | yes | Howling Idol (272427, -4.92 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Feralheart Walkers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

