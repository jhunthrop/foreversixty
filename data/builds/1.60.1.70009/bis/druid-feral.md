# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 62.9. Weights run: 1.4s. Verify run: 1.1s. 189 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=0.640 ± 0.016 per rating point (14 rating = 1%, 8.959 per %), hit=0.174 ± 0.007 per rating point (10 rating = 1%, 1.745 per %), melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 attack_power points (1.03 DPS) | yes | Brawler's Leather Hood (252504, -0.37 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.6 attack_power points (0.48 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 attack_power points (0.40 DPS) | yes | Slime-encrusted Pads (6461, -0.45 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 9.8 attack_power points (0.55 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 attack_power points (1.20 DPS) | yes | Defender's Leather Armor (252434, -0.06 DPS) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.4 attack_power points (0.58 DPS) | yes | Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Forest Leather Bracers (3202, -0.18 DPS) [world_drop]; Wolf Bracers (4794, -0.26 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.9 attack_power points (0.99 DPS) | yes | Gold-flecked Gloves (5195, -0.09 DPS) [dungeon]; Bristlebark Gloves (14572, -0.16 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.16 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.16 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 attack_power points (1.49 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 attack_power points (1.04 DPS) | yes | Feet of the Lynx (1121, -0.02 DPS) [world_drop]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.0 attack_power points (0.83 DPS) | yes | Signet of the Zhevra (285330, -0.36 DPS) [world]; The 1 Ring (8350, -0.62 DPS) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.3 attack_power points (0.51 DPS) | yes | Signet of the Zhevra (285330, -0.04 DPS) [world]; The 1 Ring (8350, -0.31 DPS) [world]; Lavishly Jeweled Ring (1156, -0.36 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | Westfall: Mr. Smite [dungeon] | 307.2 attack_power points (17.03 DPS) | yes | Staff of Westfall (2042, -1.12 DPS) [quest]; Twisted Chanter's Staff (890, -1.17 DPS) [world_drop]; Living Root (6631, -1.81 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 189, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14148 Crystalline Cuffs

### Band 30 (night-elf, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 102.6. Weights run: 1.6s. Verify run: 1.4s. 318 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=0.729 ± 0.021 per rating point (14 rating = 1%, 10.211 per %), hit=0.205 ± 0.008 per rating point (10 rating = 1%, 2.054 per %), melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 attack_power points (1.45 DPS) | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.50 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 attack_power points (0.80 DPS) | yes | Sentinel's Medallion (19541, -0.17 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.32 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.46 DPS) | yes | Bristlebark Amice (14573, -0.63 DPS) [world_drop]; Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.76 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.4 attack_power points (0.91 DPS) | yes | Sergeant Major's Cape (16315, -0.11 DPS) [pvp]; Tigerstrike Mantle (13108, -0.29 DPS) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 attack_power points (1.35 DPS) | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.2 attack_power points (0.95 DPS) | yes | Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Jurassic Wristguards (6198, -0.24 DPS) [world]; Barbaric Bracers (18948, -0.26 DPS, sim-verified) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.2 attack_power points (1.32 DPS) | yes | Toughened Leather Gloves (4253, -0.12 DPS) [crafted]; Wolfclaw Gloves (1978, -0.24 DPS) [dungeon]; Brawler Gloves (720, -0.35 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 attack_power points (1.56 DPS) | yes | Skulker's Leather Belt (252520, -0.30 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.31 DPS) [rep]; Highlander's Leather Girdle (20117, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 attack_power points (1.55 DPS) | yes | Brawler's Leather Pants (252500, -0.12 DPS) [crafted]; Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 attack_power points (1.00 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Disjointed Shoes (277226, -0.37 DPS) [quest]; Insignia Boots (4055, -0.37 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Tiger Band (6749, -0.48 DPS) [quest]; Silverlaine's Family Seal (6321, -0.60 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 22.9 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Tiger Band (6749, -0.47 DPS) [quest]; Silverlaine's Family Seal (6321, -0.59 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (102.6 DPS) | yes | Cobalt Crusher (7730, -2.40 DPS) [dungeon]; Wind Spirit Staff (6689, -4.03 DPS) [dungeon]; Viscous Hammer (13045, -22.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 318, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 133.7. Weights run: 1.7s. Verify run: 1.5s. 434 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=0.831 ± 0.024 per rating point (14 rating = 1%, 11.632 per %), hit=0.227 ± 0.010 per rating point (10 rating = 1%, 2.271 per %), melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 attack_power points (2.35 DPS) | yes | Hawkeye's Helm (14591, -0.90 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted]; Defender's Leather Helm (252455, -1.30 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.11 DPS) | yes | Sentinel's Medallion (19540, -0.17 DPS) [rep]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop]; Ghostshard Talisman (7731, -0.33 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (1.60 DPS) | yes | Forest Tracker Epaulets (2278, -0.02 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Imperial Leather Spaulders (4737, -0.44 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 24.6 attack_power points (1.36 DPS) | yes | Sergeant Major's Cape (16336, -0.08 DPS) [pvp]; Hawkeye's Cloak (14593, -0.38 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.47 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 47.7 attack_power points (2.64 DPS) | yes | Brawler's Leather Tunic (252508, -1.19 DPS) [crafted]; Wolffear Harness (13110, -1.20 DPS) [world_drop]; Barbaric Harness (5739, -1.66 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Hawkeye's Bracers (14590, -0.08 DPS) [world_drop]; Barbaric Bracers (18948, -0.25 DPS) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 37.0 attack_power points (2.05 DPS) | yes | Skulker's Leather Gloves (252525, -0.04 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Gloves of Holy Might (867, -0.30 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 attack_power points (1.92 DPS) | yes | Prowler's Leather Belt (252459, -0.25 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.26 DPS) [rep]; Skulker's Leather Belt (252520, -0.39 DPS) [crafted] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 34.6 attack_power points (1.91 DPS) | yes | Basilisk Hide Pants (1718, -0.13 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 attack_power points (2.01 DPS) | yes | Skulker's Leather Shoes (252531, -0.34 DPS, sim-verified) [crafted]; Excelsior Boots (4109, -0.34 DPS) [quest]; Imperial Leather Boots (6431, -0.43 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.8 attack_power points (1.71 DPS) | yes | Ironspine's Eye (7686, -0.43 DPS) [dungeon]; Falcon's Hook (7552, -0.56 DPS) [dungeon]; Mark of Kern (2262, -0.60 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.2 attack_power points (1.28 DPS) | yes | Ironspine's Eye (7686, -0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [dungeon]; Mark of Kern (2262, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (133.7 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Thornstone Sledgehammer (1722, -27.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Thunderbrow Ring

No-known-source sample (15 of 434, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (night-elf, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 134.5. Weights run: 1.7s. Verify run: 1.6s. 573 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=0.922 ± 0.026 per rating point (14 rating = 1%, 12.908 per %), hit=0.256 ± 0.011 per rating point (10 rating = 1%, 2.557 per %), melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 50.6 attack_power points (2.80 DPS) | yes | White Bandit Mask (10008, -0.41 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -1.06 DPS) [vendor]; Scorpashi Skullcap (14658, -1.18 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.4 attack_power points (1.79 DPS) | yes | Sentinel's Medallion (19539, -0.73 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.95 DPS, sim-verified) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 39.9 attack_power points (2.21 DPS) | yes | Failed Flying Experiment (9647, -0.04 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Warden's Leather Shoulder (252536, -0.43 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 36.3 attack_power points (2.01 DPS) | yes | Blisterbane Wrap (12552, -0.68 DPS) [dungeon]; Dark Phantom Cape (13122, -0.68 DPS) [world_drop]; Dark Hooded Cape (5257, -0.71 DPS, sim-verified) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 65.0 attack_power points (3.60 DPS) | yes | Mixologist's Tunic (12793, -0.40 DPS, sim-verified) [dungeon]; Warbear Harness (15064, -0.59 DPS) [crafted]; Quillward Harness (10583, -0.89 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 attack_power points (1.84 DPS) | yes | Prowler's Leather Bracers (252539, -0.07 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 58.3 attack_power points (3.23 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.64 DPS) [crafted]; Feralheart Fists (226793, -0.70 DPS, sim-verified) [vendor]; Skulker's Leather Gauntlets (252548, -0.84 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 51.6 attack_power points (2.86 DPS) | yes | Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.17 DPS) [dungeon]; Ogron's Sash (13117, -0.91 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 58.8 attack_power points (3.26 DPS) | yes | Serpentskin Leggings (8262, -0.17 DPS) [world_drop]; Triprunner Dungarees (9624, -1.28 DPS) [quest]; Dragonflight Leggings (10742, -1.33 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 attack_power points (2.64 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Shadefiend Boots (11675, -0.39 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 37.6 attack_power points (2.08 DPS) | yes | Thunderbrow Ring (13097, -0.79 DPS) [world_drop]; Blackstone Ring (17713, -0.83 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.84 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 23.7 attack_power points (1.31 DPS) | yes | Thunderbrow Ring (13097, -0.02 DPS) [world_drop]; Blackstone Ring (17713, -0.06 DPS) [dungeon]; Masons Fraternity Ring (9533, -0.07 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (134.5 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; The Jackhammer (9423, -0.51 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Ironspine's Eye; main_hand: Ragehammer

No-known-source sample (15 of 573, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (night-elf, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 192.5. Weights run: 1.7s. Verify run: 1.5s. 1403 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=0.980 ± 0.029 per rating point (14 rating = 1%, 13.720 per %), hit=0.305 ± 0.013 per rating point (10 rating = 1%, 3.054 per %), melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Field Marshal's Dragonhide Headguard (231689) | Captain Dirgehammer [vendor] | 93.6 attack_power points (5.08 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Headguard (227173, -0.79 DPS) [pvp]; Feralheart Cap (226792, -0.80 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.6 attack_power points (3.07 DPS) | yes | Medallion of the Dawn (22659, -1.10 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -1.20 DPS) [world_drop]; Imperial Jewel (11933, -1.34 DPS) [dungeon] |
| shoulder | Field Marshal's Dragonhide Shoulders (231693) | Captain Dirgehammer [vendor] | 67.3 attack_power points (3.65 DPS) | yes | Darkspear Pauldrons (272105, -0.33 DPS) [vendor]; Dark Warder's Pauldrons (22241, -0.64 DPS) [dungeon]; Highlander's Leather Shoulders (20059, -0.79 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.4 attack_power points (2.52 DPS) | yes | Windshear Cape (20691, -0.08 DPS) [world]; Cloak of the Honor Guard (20073, -0.20 DPS) [rep]; Cloak of Revanchion (23127, -0.30 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (192.5 DPS) | yes | Field Marshal's Dragonhide Breastplate (16452, -1.39 DPS) [vendor]; Cadaverous Armor (14637, -1.59 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.06 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (192.5 DPS) | yes | Bracers of Subterfuge (22668, -0.51 DPS) [quest]; Bracers of the Eclipse (18375, -0.94 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -3.85 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 73.4 attack_power points (3.98 DPS) | yes | Raider Gloves (272099, -0.02 DPS) [vendor]; Marshal's Dragonhide Grips (231694, -0.39 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.51 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 90.2 attack_power points (4.90 DPS) | yes | Might of the Timbermaw (19044, -0.92 DPS) [crafted]; Shifter's Belt (272396, -0.93 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -1.54 DPS) [quest] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 96.6 attack_power points (5.24 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -1.18 DPS) [vendor]; Traveler's Leggings (8300, -1.48 DPS) [dungeon] |
| feet | Marshal's Dragonhide Treads (231692) | Captain Dirgehammer [vendor] | 63.8 attack_power points (3.46 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, -0.47 DPS) [pvp]; Boots of Ferocity (22472, -0.58 DPS) [dungeon]; Drudge Boots (21532, -0.63 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (192.5 DPS) | yes | Myrmidon's Signet (2246, -1.26 DPS) [world_drop]; Don Julio's Band (19325, -1.41 DPS) [rep]; Naglering (11669, -3.68 DPS, sim-verified) [dungeon] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (192.5 DPS) | yes | Myrmidon's Signet (2246, -0.63 DPS) [world_drop]; Don Julio's Band (19325, -0.78 DPS) [rep]; Naglering (11669, -3.05 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (192.5 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -3.53 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (192.5 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (192.5 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Seeping Willow (12969, -3.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), and 12 more) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS) [vendor] |

**New at 60:** head: Field Marshal's Dragonhide Headguard; neck: Amulet of the Darkmoon; shoulder: Field Marshal's Dragonhide Shoulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Timbermaw Brawlers; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Marshal's Dragonhide Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: The Unstoppable Force; ranged: Idol of the Moon

No-known-source sample (15 of 1403, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (tauren, 0000000000000000-5420000000000000000-0000000000000000)

Set DPS (verified): 61.5. Weights run: 1.4s. Verify run: 1.1s. 185 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.436 ± 0.049, crit=0.640 ± 0.016 per rating point (14 rating = 1%, 8.959 per %), hit=0.174 ± 0.007 per rating point (10 rating = 1%, 1.745 per %), melee_haste=5.281 ± 0.416

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 18.6 attack_power points (1.03 DPS) | yes | Brawler's Leather Hood (252504, -0.34 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.6 attack_power points (0.48 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.2 attack_power points (0.40 DPS) | yes | Slime-encrusted Pads (6461, -0.43 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 9.8 attack_power points (0.55 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 21.7 attack_power points (1.20 DPS) | yes | Defender's Leather Armor (252434, -0.06 DPS) [crafted]; Totemic Leather Armor (252435, -0.30 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.32 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 8.9 attack_power points (0.50 DPS) | yes | Forest Leather Bracers (3202, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.18 DPS) [vendor]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 17.9 attack_power points (0.99 DPS) | yes | Gold-flecked Gloves (5195, -0.14 DPS, sim-verified) [dungeon]; Bristlebark Gloves (14572, -0.16 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.16 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.15 DPS, sim-verified) [crafted]; Deviate Scale Belt (6468, -0.21 DPS) [crafted]; Ruffian Belt (5975, -0.23 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 26.8 attack_power points (1.49 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 18.8 attack_power points (1.04 DPS) | yes | Feet of the Lynx (1121, -0.02 DPS) [world_drop]; Defender's Leather Boots (252441, -0.40 DPS) [crafted]; Totemic Leather Boots (252442, -0.40 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.0 attack_power points (0.83 DPS) | yes | Signet of the Zhevra (285330, -0.36 DPS) [world]; Loop of Sacrifice (281673, -0.45 DPS) [quest]; Bounty Hunter's Ring (5351, -0.59 DPS) [quest] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.3 attack_power points (0.51 DPS) | yes | Signet of the Zhevra (285330, -0.04 DPS) [world]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.28 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 309.4 attack_power points (17.15 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS, sim-verified) [dungeon]; Living Root (6631, -0.69 DPS) [dungeon]; Crescent Staff (6505, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bristlebark Bindings; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 185, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (tauren, 0000000000000000-5423222100000000000-0000000000000000)

Set DPS (verified): 101.2. Weights run: 1.6s. Verify run: 1.4s. 319 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.002, agility=1.494 ± 0.064, crit=0.729 ± 0.021 per rating point (14 rating = 1%, 10.211 per %), hit=0.205 ± 0.008 per rating point (10 rating = 1%, 2.054 per %), melee_haste=6.340 ± 0.693

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 27.8 attack_power points (1.45 DPS) | yes | Azure Gustwoven Hood (277050, -0.36 DPS) [crafted]; Defender's Leather Hood (252447, -0.48 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.49 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.3 attack_power points (0.80 DPS) | yes | Scout's Medallion (19537, -0.17 DPS) [rep]; River Pride Choker (13087, -0.31 DPS) [world_drop]; Ghostshard Talisman (7731, -0.33 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.0 attack_power points (1.46 DPS) | yes | Bristlebark Amice (14573, -0.63 DPS) [world_drop]; Mantle of Thieves (2264, -0.68 DPS) [dungeon]; Barbaric Shoulders (5964, -0.72 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 17.4 attack_power points (0.91 DPS) | yes | Tigerstrike Mantle (13108, -0.35 DPS, sim-verified) [world_drop]; Wolfmaster Cape (6314, -0.39 DPS) [dungeon]; Wildhunter Cloak (16658, -0.39 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 25.9 attack_power points (1.35 DPS) | yes | Brawler's Leather Armor (252490, -0.20 DPS) [crafted]; Defender's Leather Tunic (252450, -0.24 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.26 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.2 attack_power points (0.95 DPS) | yes | Bands of Serra'kis (6902, -0.23 DPS) [dungeon]; Barbaric Bracers (18948, -0.24 DPS, sim-verified) [crafted]; Jurassic Wristguards (6198, -0.24 DPS) [world] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.2 attack_power points (1.32 DPS) | yes | Toughened Leather Gloves (4253, -0.12 DPS) [crafted]; Wolfclaw Gloves (1978, -0.24 DPS) [dungeon]; Brawler Gloves (720, -0.35 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.8 attack_power points (1.56 DPS) | yes | Skulker's Leather Belt (252520, -0.29 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.31 DPS) [rep]; Defiler's Leather Girdle (20191, -0.31 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 29.7 attack_power points (1.55 DPS) | yes | Brawler's Leather Pants (252500, -0.12 DPS) [crafted]; Trapper's Leather Pants (252501, -0.12 DPS) [crafted]; Defender's Leather Pants (252445, -0.15 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.1 attack_power points (1.00 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Stomping Boots (3741, -0.20 DPS) [quest]; Insignia Boots (4055, -0.37 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.0 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Band of the Fist (17694, -0.41 DPS) [quest]; Tiger Band (6749, -0.48 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 22.9 attack_power points (1.20 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Band of the Fist (17694, -0.40 DPS) [quest]; Tiger Band (6749, -0.47 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (101.2 DPS) | yes | Cobalt Crusher (7730, -2.40 DPS) [dungeon]; Wind Spirit Staff (6689, -4.03 DPS) [dungeon]; Viscous Hammer (13045, -21.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-5423222121032010001-0000000000000000)

Set DPS (verified): 132.1. Weights run: 1.7s. Verify run: 1.5s. 435 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.534 ± 0.071, crit=0.831 ± 0.024 per rating point (14 rating = 1%, 11.632 per %), hit=0.227 ± 0.010 per rating point (10 rating = 1%, 2.271 per %), melee_haste=7.650 ± 0.975

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 42.4 attack_power points (2.35 DPS) | yes | Hawkeye's Helm (14591, -0.90 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.19 DPS) [crafted]; Defender's Leather Helm (252455, -1.33 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.11 DPS) | yes | Ethereal Talisman (4430, -0.13 DPS) [quest]; Scout's Medallion (19536, -0.17 DPS) [rep]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (1.60 DPS) | yes | Forest Tracker Epaulets (2278, -0.02 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Imperial Leather Spaulders (4737, -0.44 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 24.6 attack_power points (1.36 DPS) | yes | Hawkeye's Cloak (14593, -0.59 DPS, sim-verified) [world_drop]; Scorpashi Cape (14656, -0.68 DPS) [world_drop]; Parachute Cloak (10518, -0.68 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 47.7 attack_power points (2.64 DPS) | yes | Brawler's Leather Tunic (252508, -1.19 DPS) [crafted]; Wolffear Harness (13110, -1.20 DPS) [world_drop]; Barbaric Harness (5739, -1.68 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.11 DPS) | yes | Hawkeye's Bracers (14590, -0.08 DPS) [world_drop]; Barbaric Bracers (18948, -0.25 DPS) [crafted]; Bands of Serra'kis (6902, -0.34 DPS) [dungeon] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 37.0 attack_power points (2.05 DPS) | yes | Skulker's Leather Gloves (252525, -0.04 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Gloves of Holy Might (867, -0.30 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 34.7 attack_power points (1.92 DPS) | yes | Prowler's Leather Belt (252459, -0.25 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.26 DPS) [rep]; Tharg's Shoelace (9705, -0.54 DPS, sim-verified) [quest] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 34.6 attack_power points (1.91 DPS) | yes | Basilisk Hide Pants (1718, -0.13 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.25 DPS) [crafted]; Brawler's Leather Pants (252500, -0.38 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 36.3 attack_power points (2.01 DPS) | yes | Excelsior Boots (4109, -0.34 DPS) [quest]; Skulker's Leather Shoes (252531, -0.35 DPS, sim-verified) [crafted]; Imperial Leather Boots (6431, -0.43 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.8 attack_power points (1.71 DPS) | yes | Ironspine's Eye (7686, -0.43 DPS) [dungeon]; Falcon's Hook (7552, -0.56 DPS) [dungeon]; Mark of Kern (2262, -0.60 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.2 attack_power points (1.28 DPS) | yes | Ironspine's Eye (7686, -0.00 DPS) [dungeon]; Falcon's Hook (7552, -0.13 DPS) [dungeon]; Mark of Kern (2262, -0.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (132.1 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Thornstone Sledgehammer (1722, -27.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Thunderbrow Ring

No-known-source sample (15 of 435, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0000000000000000-5423222121032010001-5500000000000000)

Set DPS (verified): 136.6. Weights run: 1.7s. Verify run: 1.6s. 574 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.597 ± 0.077, crit=0.922 ± 0.026 per rating point (14 rating = 1%, 12.908 per %), hit=0.256 ± 0.011 per rating point (10 rating = 1%, 2.557 per %), melee_haste=7.662 ± 1.230

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 50.6 attack_power points (2.80 DPS) | yes | White Bandit Mask (10008, -0.41 DPS) [crafted]; Undercity Reservist's Cap (20643, -0.54 DPS) [quest]; Blood Guard's Leather Headband (220851, -1.06 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.4 attack_power points (1.79 DPS) | yes | Woven Ivy Necklace (19159, -0.23 DPS) [quest]; Zealous Shadowshard Pendant (17772, -0.69 DPS) [quest]; Scout's Medallion (19535, -0.73 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 39.9 attack_power points (2.21 DPS) | yes | Failed Flying Experiment (9647, -0.04 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.08 DPS) [crafted]; Warden's Leather Shoulder (252536, -0.43 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 36.3 attack_power points (2.01 DPS) | yes | Blisterbane Wrap (12552, -0.68 DPS) [dungeon]; Dark Phantom Cape (13122, -0.68 DPS) [world_drop]; Dark Hooded Cape (5257, -0.69 DPS, sim-verified) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 65.0 attack_power points (3.60 DPS) | yes | Mixologist's Tunic (12793, -0.39 DPS, sim-verified) [dungeon]; Warbear Harness (15064, -0.59 DPS) [crafted]; Quillward Harness (10583, -0.89 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 33.2 attack_power points (1.84 DPS) | yes | Prowler's Leather Bracers (252539, -0.07 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.15 DPS) [crafted]; Pridelord Bands (14672, -0.31 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 58.3 attack_power points (3.23 DPS) | yes | Prowler's Leather Gauntlets (252547, -0.64 DPS) [crafted]; Feralheart Fists (226793, -0.68 DPS, sim-verified) [vendor]; Skulker's Leather Gauntlets (252548, -0.84 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 51.6 attack_power points (2.86 DPS) | yes | Skulker's Leather Waistguard (252474, -0.08 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.17 DPS) [dungeon]; Ogron's Sash (13117, -0.91 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 55.7 attack_power points (3.09 DPS) | yes | Dragonflight Leggings (10742, -1.17 DPS) [quest]; Triprunner Dungarees (9624, -1.17 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.23 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 47.7 attack_power points (2.64 DPS) | yes | Skulker's Leather Boots (252469, -0.08 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.37 DPS) [dungeon]; Shadefiend Boots (11675, -0.39 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 37.6 attack_power points (2.08 DPS) | yes | Ironspine's Eye (7686, -0.77 DPS) [dungeon]; Thunderbrow Ring (13097, -0.79 DPS) [world_drop]; Blackstone Ring (17713, -0.83 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.33 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Blackstone Ring (17713, -0.08 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (136.6 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (136.6 DPS) | yes | Molten Heart of the Mountain (249470, -0.56 DPS, sim-verified) [crafted] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (136.6 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; The Jackhammer (9423, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, +0.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 574, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0000000000000000-5423222121032010001-5553200000000000)

Set DPS (verified): 193.7. Weights run: 1.7s. Verify run: 1.6s. 1405 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.320 ± 0.003, agility=1.760 ± 0.098, crit=0.980 ± 0.029 per rating point (14 rating = 1%, 13.720 per %), hit=0.305 ± 0.013 per rating point (10 rating = 1%, 3.054 per %), melee_haste=not significant (7.156 ± 1.916)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warlord's Dragonhide Headguard (231687) | Lady Palanseer [vendor] | 93.6 attack_power points (5.08 DPS) | yes | Warlord's Dragonhide Helmet (16550, -0.00 DPS) [vendor]; Feralheart Cap (226792, -0.79 DPS, sim-verified) [vendor]; Champion's Dragonhide Headguard (227174, -0.79 DPS) [pvp] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.6 attack_power points (3.07 DPS) | yes | Medallion of the Dawn (22659, -1.08 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -1.20 DPS) [world_drop]; Imperial Jewel (11933, -1.34 DPS) [dungeon] |
| shoulder | Warlord's Dragonhide Shoulders (231684) | Lady Palanseer [vendor] | 67.3 attack_power points (3.65 DPS) | yes | Darkspear Pauldrons (272105, -0.33 DPS) [vendor]; Warlord's Dragonhide Epaulets (16551, -0.55 DPS) [vendor]; Defiler's Leather Shoulders (20194, -0.79 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.4 attack_power points (2.52 DPS) | yes | Windshear Cape (20691, -0.08 DPS) [world]; Deathguard's Cloak (20068, -0.20 DPS) [rep]; Cloak of Revanchion (23127, -0.30 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (193.7 DPS) | yes | Warlord's Dragonhide Hauberk (16549, -1.39 DPS) [vendor]; Cadaverous Armor (14637, -1.59 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.98 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (193.7 DPS) | yes | Bracers of Subterfuge (22668, -0.51 DPS) [quest]; Bracers of the Eclipse (18375, -0.94 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -3.80 DPS, sim-verified) [world] |
| hands | Timbermaw Brawlers (19049) | Leatherworking [crafted] | 73.4 attack_power points (3.98 DPS) | yes | Raider Gloves (272099, -0.02 DPS) [vendor]; General's Dragonhide Grips (231688, -0.39 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.51 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 90.2 attack_power points (4.90 DPS) | yes | Might of the Timbermaw (19044, -0.92 DPS) [crafted]; Shifter's Belt (272396, -0.93 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -1.54 DPS) [quest] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 96.6 attack_power points (5.24 DPS) | yes | General's Dragonhide Leggings (231685, -0.38 DPS) [pvp]; Sentinel's Leather Pants (237818, -0.58 DPS, sim-verified) [vendor]; Traveler's Leggings (8300, -1.48 DPS) [dungeon] |
| feet | General's Dragonhide Treads (231683) | Lady Palanseer [vendor] | 63.8 attack_power points (3.46 DPS) | yes | Blood Guard's Dragonhide Treads (227181, -0.47 DPS) [pvp]; Boots of Ferocity (22472, -0.58 DPS) [dungeon]; Drudge Boots (21532, -0.66 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | sim-verified (193.7 DPS) | yes | Myrmidon's Signet (2246, -1.26 DPS) [world_drop]; Don Julio's Band (19325, -1.41 DPS) [rep]; Naglering (11669, -5.29 DPS, sim-verified) [dungeon] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (193.7 DPS) | yes | Myrmidon's Signet (2246, -0.63 DPS) [world_drop]; Don Julio's Band (19325, -0.78 DPS) [rep]; Naglering (11669, -3.00 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (193.7 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (193.7 DPS) | yes | Counterattack Lodestone (18537, -1.20 DPS) [dungeon]; Hand of Justice (11815, -1.22 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, -2.25 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (193.7 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -3.56 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), and 12 more) | World drop [world_drop] | 0.0 attack_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS) [vendor] |

**New at 60:** head: Warlord's Dragonhide Headguard; neck: Amulet of the Darkmoon; shoulder: Warlord's Dragonhide Shoulders; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Timbermaw Brawlers; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: General's Dragonhide Treads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Moon

No-known-source sample (15 of 1405, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

