# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 66.1. Weights run: 4.1s. Verify run: 4.7s. 193 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.001, agility=1.482 ± 0.009, crit=0.689 ± 0.012 per rating point (14 rating = 1%, 9.649 per %), hit=0.868 ± 0.030 per rating point (10 rating = 1%, 8.682 per %), melee_haste=5.607 ± 0.314

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.98 DPS) | yes | Brawler's Leather Hood (252504, -0.32 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.9 attack_power points (0.49 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.4 attack_power points (0.41 DPS) | yes | Slime-encrusted Pads (6461, -0.44 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.6 attack_power points (0.53 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (66.1 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.21 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.3 attack_power points (0.57 DPS) | yes | Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Forest Leather Bracers (3202, -0.16 DPS) [world_drop]; Wolf Bracers (4794, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.1 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -3.41 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Deviate Scale Belt (6468, -0.22 DPS) [crafted]; Ruffian Belt (5975, -0.27 DPS) [world]; Brawler's Leather Belt (252428, -3.86 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.1 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.43 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.5 attack_power points (1.03 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Defender's Leather Boots (252441, -0.41 DPS) [crafted]; Totemic Leather Boots (252442, -0.41 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 14.7 attack_power points (0.82 DPS) | yes | Demon Band (12054, -0.33 DPS) [world_drop]; The 1 Ring (8350, -0.61 DPS) [world]; Lavishly Jeweled Ring (1156, -0.65 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.9 attack_power points (0.49 DPS) | yes | Demon Band (12054, -0.01 DPS) [world_drop]; The 1 Ring (8350, -0.29 DPS) [world]; Lavishly Jeweled Ring (1156, -0.33 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 306.1 attack_power points (17.02 DPS) | yes | Staff of Westfall (2042, -1.06 DPS) [quest]; Twisted Chanter's Staff (890, -1.11 DPS) [world_drop]; Living Root (6631, -1.73 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 120.0. Weights run: 4.4s. Verify run: 2.9s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.001, agility=1.551 ± 0.011, crit=0.788 ± 0.016 per rating point (14 rating = 1%, 11.031 per %), hit=0.888 ± 0.032 per rating point (10 rating = 1%, 8.883 per %), melee_haste=5.289 ± 0.440

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.75 DPS) | yes | Azure Gustwoven Hood (277050, -0.44 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.57 DPS, sim-verified) [crafted]; Defender's Leather Hood (252447, -0.58 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Sentinel's Medallion (19541, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.30 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.1 attack_power points (1.86 DPS) | yes | Bristlebark Amice (14573, -0.81 DPS) [world_drop]; Mantle of Thieves (2264, -0.83 DPS) [dungeon]; Barbaric Shoulders (5964, -0.85 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (1.16 DPS) | yes | Sergeant Major's Cape (16315, -0.16 DPS) [pvp]; Tigerstrike Mantle (13108, -0.34 DPS) [world_drop]; Wolfmaster Cape (6314, -0.50 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.6 attack_power points (1.70 DPS) | yes | Brawler's Leather Armor (252490, -0.25 DPS) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted]; Defender's Leather Tunic (252450, -0.28 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.1 attack_power points (1.20 DPS) | yes | Barbaric Bracers (18948, -0.28 DPS, sim-verified) [crafted]; Jurassic Wristguards (6198, -0.29 DPS) [world]; Bands of Serra'kis (6902, -0.33 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 24.7 attack_power points (1.64 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.1 attack_power points (1.93 DPS) | yes | Skulker's Leather Belt (252520, -0.31 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.34 DPS) [rep]; Highlander's Leather Girdle (20117, -0.34 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.4 attack_power points (1.95 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.21 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 19.0 attack_power points (1.26 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Insignia Boots (4055, -0.44 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.44 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.8 attack_power points (1.51 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Tiger Band (6749, -0.63 DPS) [quest]; Seal of Wrynn (2933, -0.76 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.5 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Tiger Band (6749, -0.62 DPS) [quest]; Seal of Wrynn (2933, -0.75 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (120.0 DPS) | yes | Cobalt Crusher (7730, -2.93 DPS) [dungeon]; Wind Spirit Staff (6689, -5.02 DPS) [dungeon]; Viscous Hammer (13045, -22.57 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 169.0. Weights run: 4.9s. Verify run: 3.1s. 438 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.615 ± 0.019, crit=0.878 ± 0.027 per rating point (14 rating = 1%, 12.296 per %), hit=0.970 ± 0.040 per rating point (10 rating = 1%, 9.698 per %), melee_haste=5.322 ± 0.588

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.0 attack_power points (3.19 DPS) | yes | Defender's Leather Helm (252455, -1.18 DPS) [crafted]; Hawkeye's Helm (14591, -1.53 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.60 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.52 DPS) | yes | Sentinel's Medallion (19540, -0.17 DPS) [rep]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop]; Ghostshard Talisman (7731, -0.46 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.8 attack_power points (2.26 DPS) | yes | Forest Tracker Epaulets (2278, -0.08 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.69 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 25.0 attack_power points (1.90 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.54 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.66 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 48.3 attack_power points (3.67 DPS) | yes | Wolffear Harness (13110, -1.58 DPS) [world_drop]; Brawler's Leather Tunic (252508, -1.69 DPS) [crafted]; Barbaric Harness (5739, -2.11 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.52 DPS) | yes | Hawkeye's Bracers (14590, -0.11 DPS) [world_drop]; Barbaric Bracers (18948, -0.36 DPS) [crafted]; Jurassic Wristguards (6198, -0.45 DPS) [world] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 36.5 attack_power points (2.78 DPS) | yes | Skulker's Leather Gloves (252525, -0.04 DPS) [crafted]; Imperial Leather Gloves (4063, -0.17 DPS) [dungeon]; Gloves of Holy Might (867, -0.32 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.3 attack_power points (2.61 DPS) | yes | Prowler's Leather Belt (252459, -0.37 DPS) [crafted]; Skulker's Leather Belt (252520, -0.50 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.56 DPS, sim-verified) [rep] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 35.7 attack_power points (2.71 DPS) | yes | Basilisk Hide Pants (1718, -0.13 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.44 DPS) [crafted]; Brawler's Leather Pants (252500, -0.60 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 35.5 attack_power points (2.70 DPS) | yes | Skulker's Leather Shoes (252531, -0.18 DPS) [crafted]; Excelsior Boots (4109, -0.42 DPS) [quest]; Imperial Leather Boots (6431, -0.51 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.5 attack_power points (2.32 DPS) | yes | Thunderbrow Ring (13097, -0.61 DPS) [world_drop]; Falcon's Hook (7552, -0.71 DPS) [dungeon]; Ring of the Underwood (2951, -0.76 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.3 attack_power points (1.78 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Falcon's Hook (7552, -0.17 DPS) [dungeon]; Ring of the Underwood (2951, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (169.0 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; The Jackhammer (9423, -28.87 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 195.0. Weights run: 4.9s. Verify run: 3.7s. 577 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.693 ± 0.022, crit=0.990 ± 0.031 per rating point (14 rating = 1%, 13.866 per %), hit=1.106 ± 0.052 per rating point (10 rating = 1%, 11.060 per %), melee_haste=5.503 ± 0.769

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (4.07 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.71 DPS) [vendor]; White Bandit Mask (10008, -1.43 DPS, sim-verified) [crafted]; Scorpashi Skullcap (14658, -1.58 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.0 attack_power points (2.71 DPS) | yes | Zealous Shadowshard Pendant (17772, -1.07 DPS) [quest]; Sentinel's Medallion (19540, -1.18 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 39.4 attack_power points (3.24 DPS) | yes | Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.53 DPS) [vendor]; Failed Flying Experiment (9647, -1.16 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 36.9 attack_power points (3.03 DPS) | yes | Blisterbane Wrap (12552, -0.95 DPS) [dungeon]; Dark Phantom Cape (13122, -0.95 DPS) [world_drop]; Dark Hooded Cape (5257, -1.11 DPS, sim-verified) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 64.5 attack_power points (5.30 DPS) | yes | Mixologist's Tunic (12793, -0.68 DPS, sim-verified) [dungeon]; Warbear Harness (15064, -0.81 DPS) [crafted]; Quillward Harness (10583, -1.21 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.2 attack_power points (2.81 DPS) | yes | Prowler's Leather Bracers (252539, -0.21 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.29 DPS) [crafted]; Pridelord Bands (14672, -0.51 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 59.3 attack_power points (4.87 DPS) | yes | Feralheart Fists (226793, -1.07 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -1.09 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.29 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 51.1 attack_power points (4.20 DPS) | yes | Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.29 DPS) [dungeon]; Ogron's Sash (13117, -1.32 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 58.4 attack_power points (4.80 DPS) | yes | Serpentskin Leggings (8262, -0.22 DPS) [world_drop]; Knight's Leather Pants (220858, -1.27 DPS) [vendor]; Triprunner Dungarees (9624, -1.75 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.2 attack_power points (3.88 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.43 DPS) [dungeon]; Shadefiend Boots (11675, -0.54 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 37.2 attack_power points (3.06 DPS) | yes | Ironspine's Eye (7686, -1.08 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.11 DPS) [quest]; Thunderbrow Ring (13097, -1.20 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 31.1 attack_power points (2.55 DPS) | yes | Masons Fraternity Ring (9533, -0.60 DPS) [quest]; Thunderbrow Ring (13097, -0.69 DPS) [world_drop]; Ironspine's Eye (7686, -1.08 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (195.0 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (195.0 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (195.0 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; The Jackhammer (9423, -2.24 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 302.9. Weights run: 4.8s. Verify run: 11.9s. 1449 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.767 ± 0.029, crit=1.096 ± 0.041 per rating point (14 rating = 1%, 15.341 per %), hit=1.237 ± 0.071 per rating point (10 rating = 1%, 12.374 per %), melee_haste=6.472 ± 1.043

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 81.0 attack_power points (7.21 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -11.09 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 55.6 attack_power points (4.95 DPS) | yes | Medallion of the Dawn (22659, -1.45 DPS) [quest]; Pendant of Celerity (22340, -1.49 DPS) [dungeon]; Mark of Fordring (15411, -1.51 DPS, sim-verified) [quest] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (302.9 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -7.56 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.5 attack_power points (4.14 DPS) | yes | Windshear Cape (20691, -0.21 DPS) [world]; Cloak of the Honor Guard (20073, -0.33 DPS) [rep]; Howler's Furs (272414, -0.55 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (302.9 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -2.07 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.20 DPS) [vendor]; Tunic of Undead Slaying (23089, -12.29 DPS, sim-verified) [world] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (302.9 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 71.9 attack_power points (6.40 DPS) | yes | Timbermaw Brawlers (19049, -0.09 DPS) [crafted]; Marshal's Dragonhide Grips (231694, -0.53 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.91 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 87.2 attack_power points (7.76 DPS) | yes | Might of the Timbermaw (19044, -1.45 DPS) [crafted]; Belt of Preserved Heads (20216, -1.56 DPS) [quest]; Shifter's Belt (272396, -1.59 DPS, sim-verified) [vendor] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 93.4 attack_power points (8.31 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Sentinel's Leather Pants (237818, -1.52 DPS, sim-verified) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.61 DPS) [vendor] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (302.9 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Drudge Boots (21532, -8.25 DPS, sim-verified) [quest] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (302.9 DPS) | yes | Don Julio's Band (19325, -0.19 DPS) [rep]; Tarnished Elven Ring (18500, -0.62 DPS) [dungeon]; Naglering (11669, -4.92 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (302.9 DPS) | yes | Don Julio's Band (19325, -0.04 DPS) [rep]; Tarnished Elven Ring (18500, -0.47 DPS) [dungeon]; Naglering (11669, -3.53 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (302.9 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (302.9 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -1.98 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (302.9 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Felstriker (12590, -4.71 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (302.9 DPS) | yes | Howling Idol (272427, -2.33 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Feralheart Walkers; finger1: Protector's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 625.7. Weights run: 4.5s. Verify run: 11.7s. 1449 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.842 ± 0.027, crit=1.203 ± 0.038 per rating point (14 rating = 1%, 16.843 per %), hit=1.327 ± 0.078 per rating point (10 rating = 1%, 13.268 per %), melee_haste=11.159 ± 1.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 83.7 attack_power points (13.23 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -7.55 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 57.0 attack_power points (9.01 DPS) | yes | Mark of Fordring (15411, -2.36 DPS, sim-verified) [quest]; Pendant of Celerity (22340, -2.54 DPS) [dungeon]; Medallion of the Dawn (22659, -2.55 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (625.7 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -27.53 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.6 attack_power points (7.53 DPS) | yes | Windshear Cape (20691, -0.38 DPS) [world]; Cloak of the Honor Guard (20073, -0.70 DPS) [rep]; Howler's Furs (272414, -1.00 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (625.7 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -3.59 DPS) [vendor]; Dawn Armor (252483, -3.85 DPS) [crafted]; Tunic of Undead Slaying (23089, -24.39 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (625.7 DPS) | yes | Bracers of Subterfuge (22668, -1.51 DPS) [quest]; Bracers of the Eclipse (18375, -2.65 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.37 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (625.7 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, -28.73 DPS, sim-verified) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 88.5 attack_power points (13.99 DPS) | yes | Shifter's Belt (272396, -2.36 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -2.61 DPS) [crafted]; Belt of Preserved Heads (20216, -2.66 DPS) [quest] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (625.7 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Sentinel's Leather Pants (237818, +0.00 DPS) [vendor]; Warbear Woolies (15065, -26.32 DPS, sim-verified) [crafted] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | 59.9 attack_power points (9.47 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Knight-Lieutenant's Dragonhide Treads (227182, -0.36 DPS) [pvp]; Boots of Ferocity (22472, -1.16 DPS) [dungeon] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (625.7 DPS) | yes | Don Julio's Band (19325, -0.09 DPS) [rep]; Tarnished Elven Ring (18500, -0.91 DPS) [dungeon]; Naglering (11669, -9.79 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (625.7 DPS) | yes | Don Julio's Band (19325, -0.05 DPS) [rep]; Tarnished Elven Ring (18500, -0.87 DPS) [dungeon]; Naglering (11669, -10.48 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (625.7 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -2.65 DPS, sim-verified) [quest] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (625.7 DPS) | yes | Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, -3.24 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (625.7 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Felstriker (12590, -9.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (625.7 DPS) | yes | Howling Idol (272427, -8.64 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Ferocity of the Timbermaw; legs: Stormshroud Pants; feet: Drudge Boots; finger1: Protector's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 64.7. Weights run: 4.1s. Verify run: 4.7s. 183 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.001, agility=1.482 ± 0.009, crit=0.689 ± 0.012 per rating point (14 rating = 1%, 9.649 per %), hit=0.868 ± 0.030 per rating point (10 rating = 1%, 8.682 per %), melee_haste=5.607 ± 0.314

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.98 DPS) | yes | Brawler's Leather Hood (252504, -0.27 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.9 attack_power points (0.49 DPS) | yes | Erudite's Amulet (277204, -0.16 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.4 attack_power points (0.41 DPS) | yes | Slime-encrusted Pads (6461, -0.41 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 9.6 attack_power points (0.53 DPS) | yes | Glowing Lizardscale Cloak (6449, -0.04 DPS) [dungeon]; Dark Leather Cloak (2316, -0.04 DPS) [crafted]; Lambent Scale Cloak (4706, -0.04 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (64.7 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.16 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.8 attack_power points (0.49 DPS) | yes | Forest Leather Bracers (3202, -0.08 DPS) [world_drop]; Wolf Bracers (4794, -0.16 DPS) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (64.7 DPS) | yes | Gold-flecked Gloves (5195, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -3.35 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Deviate Scale Belt (6468, -0.22 DPS) [crafted]; Ruffian Belt (5975, -0.27 DPS) [world]; Brawler's Leather Belt (252428, -3.80 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (64.7 DPS) | yes | Defender's Leather Pants (252445, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.36 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 18.5 attack_power points (1.03 DPS) | yes | Brawler's Leather Boots (252439, -0.00 DPS) [crafted]; Defender's Leather Boots (252441, -0.41 DPS) [crafted]; Totemic Leather Boots (252442, -0.41 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 14.7 attack_power points (0.82 DPS) | yes | Demon Band (12054, -0.33 DPS) [world_drop]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; Bounty Hunter's Ring (5351, -0.57 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 8.9 attack_power points (0.49 DPS) | yes | Demon Band (12054, -0.01 DPS) [world_drop]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.25 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 308.2 attack_power points (17.13 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS, sim-verified) [dungeon]; Living Root (6631, -0.62 DPS) [dungeon]; Crescent Staff (6505, -0.72 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Hammerbone

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 118.2. Weights run: 4.4s. Verify run: 2.9s. 315 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.001, agility=1.551 ± 0.011, crit=0.788 ± 0.016 per rating point (14 rating = 1%, 11.031 per %), hit=0.888 ± 0.032 per rating point (10 rating = 1%, 8.883 per %), melee_haste=5.289 ± 0.440

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.75 DPS) | yes | Azure Gustwoven Hood (277050, -0.44 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.56 DPS, sim-verified) [crafted]; Defender's Leather Hood (252447, -0.58 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Scout's Medallion (19537, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.36 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.1 attack_power points (1.86 DPS) | yes | Bristlebark Amice (14573, -0.81 DPS) [world_drop]; Mantle of Thieves (2264, -0.83 DPS) [dungeon]; Barbaric Shoulders (5964, -0.90 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.5 attack_power points (1.16 DPS) | yes | Tigerstrike Mantle (13108, -0.43 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.50 DPS) [dungeon]; Wildhunter Cloak (16658, -0.50 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.6 attack_power points (1.70 DPS) | yes | Brawler's Leather Armor (252490, -0.25 DPS) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted]; Defender's Leather Tunic (252450, -0.31 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.1 attack_power points (1.20 DPS) | yes | Jurassic Wristguards (6198, -0.29 DPS) [world]; Barbaric Bracers (18948, -0.31 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.33 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 24.7 attack_power points (1.64 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.1 attack_power points (1.93 DPS) | yes | Skulker's Leather Belt (252520, -0.31 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.34 DPS) [rep]; Defiler's Leather Girdle (20191, -0.34 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.4 attack_power points (1.95 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.21 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 19.0 attack_power points (1.26 DPS) | yes | Brawler's Leather Boots (252439, -0.02 DPS) [crafted]; Stomping Boots (3741, -0.27 DPS) [quest]; Insignia Boots (4055, -0.44 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 22.8 attack_power points (1.51 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Band of the Fist (17694, -0.51 DPS) [quest]; Tiger Band (6749, -0.63 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.5 attack_power points (1.49 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Band of the Fist (17694, -0.50 DPS) [quest]; Tiger Band (6749, -0.62 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (118.2 DPS) | yes | Cobalt Crusher (7730, -2.93 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -4.00 DPS) [pvp]; Viscous Hammer (13045, -22.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 166.7. Weights run: 4.9s. Verify run: 3.0s. 426 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.615 ± 0.019, crit=0.878 ± 0.027 per rating point (14 rating = 1%, 12.296 per %), hit=0.970 ± 0.040 per rating point (10 rating = 1%, 9.698 per %), melee_haste=5.322 ± 0.588

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.0 attack_power points (3.19 DPS) | yes | Defender's Leather Helm (252455, -1.18 DPS) [crafted]; Hawkeye's Helm (14591, -1.50 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.60 DPS) [world_drop] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.52 DPS) | yes | Scout's Medallion (19536, -0.17 DPS) [rep]; Ethereal Talisman (4430, -0.19 DPS) [quest]; Kaleidoscope Chain (13084, -0.36 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.8 attack_power points (2.26 DPS) | yes | Forest Tracker Epaulets (2278, -0.08 DPS) [world_drop]; Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.69 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 25.0 attack_power points (1.90 DPS) | yes | First Sergeant's Cloak (16340, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.54 DPS) [world_drop]; Parachute Cloak (10518, -0.92 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 48.3 attack_power points (3.67 DPS) | yes | Wolffear Harness (13110, -1.58 DPS) [world_drop]; Brawler's Leather Tunic (252508, -1.69 DPS) [crafted]; Barbaric Harness (5739, -2.03 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.52 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.11 DPS) [world_drop]; Barbaric Bracers (18948, -0.36 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 36.5 attack_power points (2.78 DPS) | yes | Skulker's Leather Gloves (252525, -0.04 DPS) [crafted]; Imperial Leather Gloves (4063, -0.17 DPS) [dungeon]; Gloves of Holy Might (867, -0.32 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.3 attack_power points (2.61 DPS) | yes | Prowler's Leather Belt (252459, -0.37 DPS) [crafted]; Tharg's Shoelace (9705, -0.44 DPS) [quest]; Defiler's Leather Girdle (20192, -0.53 DPS, sim-verified) [rep] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 35.7 attack_power points (2.71 DPS) | yes | Basilisk Hide Pants (1718, -0.13 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.44 DPS) [crafted]; Brawler's Leather Pants (252500, -0.60 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 35.5 attack_power points (2.70 DPS) | yes | Excelsior Boots (4109, -0.42 DPS) [quest]; Skulker's Leather Shoes (252531, -0.46 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.51 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.5 attack_power points (2.32 DPS) | yes | Thunderbrow Ring (13097, -0.61 DPS) [world_drop]; Falcon's Hook (7552, -0.71 DPS) [dungeon]; Ring of the Underwood (2951, -0.76 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.3 attack_power points (1.78 DPS) | yes | Thunderbrow Ring (13097, -0.07 DPS) [world_drop]; Falcon's Hook (7552, -0.17 DPS) [dungeon]; Ring of the Underwood (2951, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (166.7 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; The Jackhammer (9423, -27.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 199.4. Weights run: 4.9s. Verify run: 3.5s. 561 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.693 ± 0.022, crit=0.990 ± 0.031 per rating point (14 rating = 1%, 13.866 per %), hit=1.106 ± 0.052 per rating point (10 rating = 1%, 11.060 per %), melee_haste=5.503 ± 0.769

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 49.6 attack_power points (4.07 DPS) | yes | Blood Guard's Leather Headband (220851, -0.71 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.74 DPS) [quest]; White Bandit Mask (10008, -1.44 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.0 attack_power points (2.71 DPS) | yes | Scout's Medallion (19535, -1.04 DPS) [rep]; Zealous Shadowshard Pendant (17772, -1.07 DPS) [quest]; Woven Ivy Necklace (19159, -1.08 DPS, sim-verified) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 39.4 attack_power points (3.24 DPS) | yes | Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.53 DPS) [vendor]; Failed Flying Experiment (9647, -0.56 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 36.9 attack_power points (3.03 DPS) | yes | Dark Hooded Cape (5257, -0.93 DPS, sim-verified) [world]; Blisterbane Wrap (12552, -0.95 DPS) [dungeon]; Dark Phantom Cape (13122, -0.95 DPS) [world_drop] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 64.5 attack_power points (5.30 DPS) | yes | Mixologist's Tunic (12793, -0.51 DPS) [dungeon]; Warbear Harness (15064, -0.81 DPS) [crafted]; Quillward Harness (10583, -1.21 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.2 attack_power points (2.81 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Prowler's Leather Bracers (252539, -0.21 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 59.3 attack_power points (4.87 DPS) | yes | Feralheart Fists (226793, -0.75 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -1.09 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.29 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 51.1 attack_power points (4.20 DPS) | yes | Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.29 DPS) [dungeon]; Ogron's Sash (13117, -1.32 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 55.7 attack_power points (4.57 DPS) | yes | Stone Guard's Leather Pants (220859, -1.29 DPS, sim-verified) [vendor]; Triprunner Dungarees (9624, -1.53 DPS) [quest]; Basilisk Hide Pants (1718, -1.65 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.2 attack_power points (3.88 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.43 DPS) [dungeon]; Shadefiend Boots (11675, -0.54 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 37.2 attack_power points (3.06 DPS) | yes | Ironspine's Eye (7686, -1.08 DPS) [dungeon]; White Bone Band (11862, -1.09 DPS) [quest]; Masons Fraternity Ring (9533, -1.11 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 31.1 attack_power points (2.55 DPS) | yes | White Bone Band (11862, -0.58 DPS) [quest]; Masons Fraternity Ring (9533, -0.60 DPS) [quest]; Ironspine's Eye (7686, -1.23 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (199.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (199.4 DPS) | yes | Molten Heart of the Mountain (249470, -1.02 DPS, sim-verified) [crafted] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (199.4 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; The Jackhammer (9423, -2.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 304.6. Weights run: 4.8s. Verify run: 11.4s. 1446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.767 ± 0.029, crit=1.096 ± 0.041 per rating point (14 rating = 1%, 15.341 per %), hit=1.237 ± 0.071 per rating point (10 rating = 1%, 12.374 per %), melee_haste=6.472 ± 1.043

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 81.0 attack_power points (7.21 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -11.48 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 55.6 attack_power points (4.95 DPS) | yes | Mark of Fordring (15411, -1.44 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -1.45 DPS) [quest]; Pendant of Celerity (22340, -1.49 DPS) [dungeon] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (304.6 DPS) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -7.88 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.5 attack_power points (4.14 DPS) | yes | Windshear Cape (20691, -0.21 DPS) [world]; Deathguard's Cloak (20068, -0.33 DPS) [rep]; Howler's Furs (272414, -0.55 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (304.6 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, -2.07 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.20 DPS) [vendor]; Tunic of Undead Slaying (23089, -12.75 DPS, sim-verified) [world] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (304.6 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 71.9 attack_power points (6.40 DPS) | yes | Timbermaw Brawlers (19049, -0.09 DPS) [crafted]; General's Dragonhide Grips (231688, -0.53 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.91 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 87.2 attack_power points (7.76 DPS) | yes | Shifter's Belt (272396, -1.25 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -1.45 DPS) [crafted]; Belt of Preserved Heads (20216, -1.56 DPS) [quest] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 93.4 attack_power points (8.31 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -1.43 DPS, sim-verified) [vendor]; Legionnaire's Dragonhide Leggings (227177, -1.61 DPS) [pvp] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (304.6 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Drudge Boots (21532, -9.14 DPS, sim-verified) [quest] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (304.6 DPS) | yes | Don Julio's Band (19325, -0.19 DPS) [rep]; Tarnished Elven Ring (18500, -0.62 DPS) [dungeon]; Naglering (11669, -4.99 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (304.6 DPS) | yes | Don Julio's Band (19325, -0.04 DPS) [rep]; Tarnished Elven Ring (18500, -0.47 DPS) [dungeon]; Naglering (11669, -6.19 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (304.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Hand of Justice (11815, -4.31 DPS, sim-verified) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (304.6 DPS) | yes | Eye of the Beast (13968, -1.78 DPS) [quest]; Counterattack Lodestone (18537, -2.55 DPS) [dungeon]; Blackhand's Breadth (13965, -2.60 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (304.6 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -4.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (304.6 DPS) | yes | Howling Idol (272427, -2.24 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Feralheart Walkers; finger1: Legionnaire's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 627.3. Weights run: 4.5s. Verify run: 11.2s. 1446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.842 ± 0.027, crit=1.203 ± 0.038 per rating point (14 rating = 1%, 16.843 per %), hit=1.327 ± 0.078 per rating point (10 rating = 1%, 13.268 per %), melee_haste=11.159 ± 1.112

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 83.7 attack_power points (13.23 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -5.59 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 57.0 attack_power points (9.01 DPS) | yes | Mark of Fordring (15411, -2.54 DPS, sim-verified) [quest]; Pendant of Celerity (22340, -2.54 DPS) [dungeon]; Medallion of the Dawn (22659, -2.55 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (627.3 DPS) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -26.36 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.6 attack_power points (7.53 DPS) | yes | Windshear Cape (20691, -0.38 DPS) [world]; Deathguard's Cloak (20068, -0.70 DPS) [rep]; Howler's Furs (272414, -1.00 DPS) [vendor] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (627.3 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, -3.59 DPS) [vendor]; Dawn Armor (252483, -3.85 DPS) [crafted]; Tunic of Undead Slaying (23089, -22.54 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (627.3 DPS) | yes | Bracers of Subterfuge (22668, -1.51 DPS) [quest]; Bracers of the Eclipse (18375, -2.65 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.41 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (627.3 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, -27.44 DPS, sim-verified) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 88.5 attack_power points (13.99 DPS) | yes | Shifter's Belt (272396, -2.24 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -2.61 DPS) [crafted]; Belt of Preserved Heads (20216, -2.66 DPS) [quest] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (627.3 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, +0.00 DPS) [vendor]; Warbear Woolies (15065, -26.37 DPS, sim-verified) [crafted] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | 59.9 attack_power points (9.47 DPS) | yes | General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Blood Guard's Dragonhide Treads (227181, -0.36 DPS) [pvp]; Boots of Ferocity (22472, -1.16 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (627.3 DPS) | yes | Don Julio's Band (19325, -0.09 DPS) [rep]; Tarnished Elven Ring (18500, -0.91 DPS) [dungeon]; Naglering (11669, -9.14 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (627.3 DPS) | yes | Don Julio's Band (19325, -0.05 DPS) [rep]; Tarnished Elven Ring (18500, -0.87 DPS) [dungeon]; Naglering (11669, -9.90 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (627.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (627.3 DPS) | yes | Blackhand's Breadth (13965, -2.78 DPS) [quest]; Eye of the Beast (13968, -2.78 DPS) [quest]; Hand of Justice (11815, -3.39 DPS, sim-verified) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (627.3 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -5.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (627.3 DPS) | yes | Howling Idol (272427, -8.11 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Stormshroud Shoulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Ferocity of the Timbermaw; legs: Stormshroud Pants; feet: Drudge Boots; finger1: Legionnaire's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

