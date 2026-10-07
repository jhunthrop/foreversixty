# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 64.6. Weights run: 2.4s. Verify run: 1.2s. 193 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.483 ± 0.009, crit=0.691 ± 0.012 per rating point (14 rating = 1%, 9.673 per %), hit=0.176 ± 0.005 per rating point (10 rating = 1%, 1.755 per %), melee_haste=5.249 ± 0.310

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 19.2 attack_power points (1.07 DPS) | yes | Brawler's Leather Hood (252504, -0.36 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 8.9 attack_power points (0.50 DPS) | yes | Erudite's Amulet (277204, -0.18 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.4 attack_power points (0.41 DPS) | yes | Slime-encrusted Pads (6461, -0.44 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Abominable Creatures [quest] | 10.2 attack_power points (0.57 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 22.4 attack_power points (1.25 DPS) | yes | Defender's Leather Armor (252434, -0.06 DPS) [crafted]; Totemic Leather Armor (252435, -0.31 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.33 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 10.7 attack_power points (0.60 DPS) | yes | Bristlebark Bindings (14569, -0.08 DPS) [world_drop]; Forest Leather Bracers (3202, -0.18 DPS) [world_drop]; Wolf Bracers (4794, -0.27 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 18.5 attack_power points (1.03 DPS) | yes | Gold-flecked Gloves (5195, -0.09 DPS) [dungeon]; Bristlebark Gloves (14572, -0.17 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.17 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.14 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Ruffian Belt (5975, -0.20 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 27.8 attack_power points (1.55 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.4 attack_power points (1.08 DPS) | yes | Feet of the Lynx (1121, -0.02 DPS) [world_drop]; Defender's Leather Boots (252441, -0.41 DPS) [crafted]; Totemic Leather Boots (252442, -0.41 DPS) [crafted] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.5 attack_power points (0.87 DPS) | yes | Signet of the Zhevra (285330, -0.37 DPS) [world]; The 1 Ring (8350, -0.65 DPS) [world]; Lavishly Jeweled Ring (1156, -0.70 DPS) [dungeon] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.6 attack_power points (0.54 DPS) | yes | Signet of the Zhevra (285330, -0.04 DPS) [world]; The 1 Ring (8350, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.37 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.3 attack_power points (17.18 DPS) | yes | Staff of Westfall (2042, -1.19 DPS) [quest]; Twisted Chanter's Staff (890, -1.23 DPS) [world_drop]; Living Root (6631, -1.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-55232220000000000000-0000000000000000)

Set DPS (verified): 111.9. Weights run: 2.8s. Verify run: 1.4s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.537 ± 0.011, crit=0.767 ± 0.015 per rating point (14 rating = 1%, 10.738 per %), hit=0.218 ± 0.006 per rating point (10 rating = 1%, 2.182 per %), melee_haste=6.889 ± 0.540

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 28.8 attack_power points (1.62 DPS) | yes | Azure Gustwoven Hood (277050, -0.40 DPS) [crafted]; Defender's Leather Hood (252447, -0.54 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.55 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 attack_power points (0.89 DPS) | yes | Sentinel's Medallion (19541, -0.19 DPS) [rep]; Ghostshard Talisman (7731, -0.34 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -0.35 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.9 attack_power points (1.62 DPS) | yes | Bristlebark Amice (14573, -0.70 DPS) [world_drop]; Barbaric Shoulders (5964, -0.76 DPS, sim-verified) [crafted]; Mantle of Thieves (2264, -0.76 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 18.0 attack_power points (1.01 DPS) | yes | Sergeant Major's Cape (16315, -0.12 DPS) [pvp]; Tigerstrike Mantle (13108, -0.32 DPS) [world_drop]; Grave Shroud (279865, -0.43 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 26.7 attack_power points (1.50 DPS) | yes | Brawler's Leather Armor (252490, -0.22 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS, sim-verified) [crafted]; Thick Murloc Armor (5782, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.8 attack_power points (1.06 DPS) | yes | Bands of Serra'kis (6902, -0.25 DPS) [dungeon]; Barbaric Bracers (18948, -0.26 DPS, sim-verified) [crafted]; Jurassic Wristguards (6198, -0.27 DPS) [world] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.0 attack_power points (1.46 DPS) | yes | Toughened Leather Gloves (4253, -0.13 DPS) [crafted]; Wolfclaw Gloves (1978, -0.27 DPS) [dungeon]; Brawler Gloves (720, -0.38 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 30.8 attack_power points (1.73 DPS) | yes | Skulker's Leather Belt (252520, -0.35 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.38 DPS) [rep]; Highlander's Leather Girdle (20117, -0.38 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 30.6 attack_power points (1.72 DPS) | yes | Brawler's Leather Pants (252500, -0.13 DPS) [crafted]; Trapper's Leather Pants (252501, -0.13 DPS) [crafted]; Defender's Leather Pants (252445, -0.16 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.7 attack_power points (1.11 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Insignia Boots (4055, -0.42 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.42 DPS) [vendor] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.8 attack_power points (1.34 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Tiger Band (6749, -0.53 DPS) [quest]; Silverlaine's Family Seal (6321, -0.66 DPS) [dungeon] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 23.6 attack_power points (1.33 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Tiger Band (6749, -0.52 DPS) [quest]; Silverlaine's Family Seal (6321, -0.65 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (111.9 DPS) | yes | Cobalt Crusher (7730, -2.66 DPS) [dungeon]; Wind Spirit Staff (6689, -4.42 DPS) [dungeon]; Viscous Hammer (13045, -23.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Protector's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-55232232121030000000-0000000000000000)

Set DPS (verified): 161.1. Weights run: 3.0s. Verify run: 1.4s. 438 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.608 ± 0.012, crit=0.868 ± 0.017 per rating point (14 rating = 1%, 12.152 per %), hit=0.219 ± 0.007 per rating point (10 rating = 1%, 2.190 per %), melee_haste=6.258 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.1 attack_power points (2.92 DPS) | yes | Hawkeye's Helm (14591, -1.11 DPS) [world_drop]; Cloudy Gustwoven Hood (277042, -1.49 DPS) [crafted]; Defender's Leather Helm (252455, -1.51 DPS, sim-verified) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.33 DPS) | yes | Sentinel's Medallion (19540, -0.15 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.7 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.00 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.53 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 25.7 attack_power points (1.70 DPS) | yes | Sergeant Major's Cape (16336, +0.00 DPS, sim-verified) [pvp]; Hawkeye's Cloak (14593, -0.48 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.59 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 49.7 attack_power points (3.30 DPS) | yes | Wolffear Harness (13110, -1.49 DPS) [world_drop]; Brawler's Leather Tunic (252508, -1.49 DPS) [crafted]; Barbaric Harness (5739, -1.93 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.33 DPS) | yes | Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS) [crafted]; Jurassic Wristguards (6198, -0.37 DPS) [world] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 38.5 attack_power points (2.55 DPS) | yes | Skulker's Leather Gloves (252525, -0.05 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.42 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.1 attack_power points (2.39 DPS) | yes | Prowler's Leather Belt (252459, -0.32 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.40 DPS) [rep]; Skulker's Leather Belt (252520, -0.48 DPS) [crafted] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 36.1 attack_power points (2.40 DPS) | yes | Basilisk Hide Pants (1718, -0.16 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.32 DPS) [crafted]; Brawler's Leather Pants (252500, -0.48 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.7 attack_power points (2.50 DPS) | yes | Skulker's Leather Shoes (252531, -0.21 DPS) [crafted]; Excelsior Boots (4109, -0.42 DPS) [quest]; Imperial Leather Boots (6431, -0.53 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 32.1 attack_power points (2.13 DPS) | yes | Thunderbrow Ring (13097, -0.53 DPS) [world_drop]; Falcon's Hook (7552, -0.69 DPS) [dungeon]; Ring of the Underwood (2951, -0.74 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.1 attack_power points (1.60 DPS) | yes | Thunderbrow Ring (13097, -0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (161.1 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Thornstone Sledgehammer (1722, -28.44 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 0100000000000000-55232232121032012001-3000000000000000)

Set DPS (verified): 183.6. Weights run: 3.0s. Verify run: 1.6s. 577 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.678 ± 0.023, crit=0.968 ± 0.033 per rating point (14 rating = 1%, 13.554 per %), hit=0.250 ± 0.008 per rating point (10 rating = 1%, 2.498 per %), melee_haste=6.986 ± 1.037

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 51.2 attack_power points (3.65 DPS) | yes | White Bandit Mask (10008, -0.45 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -1.37 DPS) [vendor]; Scorpashi Skullcap (14658, -1.46 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.8 attack_power points (2.41 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.99 DPS) [quest]; Sentinel's Medallion (19540, -1.10 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 41.5 attack_power points (2.96 DPS) | yes | Failed Flying Experiment (9647, -0.05 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.10 DPS) [crafted]; Warden's Leather Shoulder (252536, -0.58 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 37.9 attack_power points (2.70 DPS) | yes | Blisterbane Wrap (12552, -0.91 DPS) [dungeon]; Dark Phantom Cape (13122, -0.91 DPS) [world_drop]; Dark Hooded Cape (5257, -1.00 DPS, sim-verified) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 67.6 attack_power points (4.83 DPS) | yes | Mixologist's Tunic (12793, -0.60 DPS, sim-verified) [dungeon]; Warbear Harness (15064, -0.79 DPS) [crafted]; Quillward Harness (10583, -1.18 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.8 attack_power points (2.48 DPS) | yes | Prowler's Leather Bracers (252539, -0.10 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.20 DPS) [crafted]; Pridelord Bands (14672, -0.43 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 60.9 attack_power points (4.35 DPS) | yes | Feralheart Fists (226793, -0.85 DPS) [vendor]; Prowler's Leather Gauntlets (252547, -0.87 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.13 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.7 attack_power points (3.83 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.32 DPS) [dungeon]; Ogron's Sash (13117, -1.22 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 61.2 attack_power points (4.36 DPS) | yes | Serpentskin Leggings (8262, -0.22 DPS) [world_drop]; Triprunner Dungarees (9624, -1.70 DPS) [quest]; Dragonflight Leggings (10742, -1.78 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.7 attack_power points (3.54 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.48 DPS) [dungeon]; Shadefiend Boots (11675, -0.51 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 39.1 attack_power points (2.79 DPS) | yes | Thunderbrow Ring (13097, -1.06 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.11 DPS) [quest]; Blackstone Ring (17713, -1.18 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.7 attack_power points (1.76 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.09 DPS) [quest]; Blackstone Ring (17713, -0.16 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (183.6 DPS) | yes | Taran Icebreaker (2915, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -2.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 0100000000000000-55232232121032012001-5053000000000000)

Set DPS (verified): 281.6. Weights run: 3.0s. Verify run: 1.6s. 1449 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.849 ± 0.032, crit=1.213 ± 0.046 per rating point (14 rating = 1%, 16.987 per %), hit=0.299 ± 0.010 per rating point (10 rating = 1%, 2.989 per %), melee_haste=7.090 ± 1.511

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 87.4 attack_power points (6.42 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -3.25 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 59.1 attack_power points (4.34 DPS) | yes | Medallion of the Dawn (22659, -1.57 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -1.70 DPS) [world_drop]; Imperial Jewel (11933, -1.99 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 66.9 attack_power points (4.92 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Highlander's Leather Shoulders (20059, -0.27 DPS) [rep]; Dark Warder's Pauldrons (22241, -0.67 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.7 attack_power points (3.51 DPS) | yes | Windshear Cape (20691, -0.06 DPS) [world]; Cloak of the Honor Guard (20073, -0.33 DPS) [rep]; Cloak of Revanchion (23127, -0.39 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (281.6 DPS) | yes | Field Marshal's Dragonhide Breastplate (16452, -1.95 DPS) [vendor]; Cadaverous Armor (14637, -2.61 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.10 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (281.6 DPS) | yes | Bracers of Subterfuge (22668, -0.72 DPS) [quest]; Bracers of the Eclipse (18375, -1.40 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -5.16 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 76.3 attack_power points (5.61 DPS) | yes | Timbermaw Brawlers (19049, -0.00 DPS) [crafted]; Marshal's Dragonhide Grips (231694, -0.36 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.72 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 93.8 attack_power points (6.89 DPS) | yes | Might of the Timbermaw (19044, -1.29 DPS) [crafted]; Shifter's Belt (272396, -1.49 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -2.17 DPS) [quest] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 100.5 attack_power points (7.38 DPS) | yes | Marshal's Dragonhide Leggings (231691, -0.35 DPS) [vendor]; Sentinel's Leather Pants (237818, -1.30 DPS, sim-verified) [vendor]; Traveler's Leggings (8300, -2.08 DPS) [dungeon] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | 62.6 attack_power points (4.60 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Knight-Lieutenant's Dragonhide Treads (227182, -0.19 DPS) [pvp]; Boots of Ferocity (22472, -0.54 DPS) [dungeon] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-verified (281.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.95 DPS) [quest]; Don Julio's Band (19325, -0.97 DPS) [rep]; Naglering (11669, -4.02 DPS, sim-verified) [dungeon] |
| finger2 | Myrmidon's Signet (2246) | World drop [world_drop] | sim-verified (281.6 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.05 DPS) [quest]; Don Julio's Band (19325, -0.07 DPS) [rep]; Naglering (11669, -3.01 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (281.6 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (281.6 DPS) | yes | Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.27 DPS) [crafted]; Blackhand's Breadth (13965, -2.72 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (281.6 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Felstriker (12590, -4.54 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (281.6 DPS) | yes | Howling Idol (272427, -4.47 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Drudge Boots; finger1: Protector's Band; finger2: Myrmidon's Signet; trinket1: Darkmoon Card: Maelstrom; trinket2: Hand of Justice; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 63.2. Weights run: 2.4s. Verify run: 1.2s. 183 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.483 ± 0.009, crit=0.691 ± 0.012 per rating point (14 rating = 1%, 9.673 per %), hit=0.176 ± 0.005 per rating point (10 rating = 1%, 1.755 per %), melee_haste=5.249 ± 0.310

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 19.2 attack_power points (1.07 DPS) | yes | Brawler's Leather Hood (252504, -0.36 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 8.9 attack_power points (0.50 DPS) | yes | Erudite's Amulet (277204, -0.17 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 7.4 attack_power points (0.41 DPS) | yes | Slime-encrusted Pads (6461, -0.44 DPS, sim-verified) [dungeon] |
| back | Grave Shroud (279865) | Unending Torment [quest] | 10.2 attack_power points (0.57 DPS) | yes | Lambent Scale Cloak (4706, -0.03 DPS) [world_drop]; Dark Leather Cloak (2316, -0.05 DPS) [crafted]; Glowing Lizardscale Cloak (6449, -0.07 DPS) [dungeon] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 22.4 attack_power points (1.25 DPS) | yes | Defender's Leather Armor (252434, -0.06 DPS) [crafted]; Totemic Leather Armor (252435, -0.31 DPS) [crafted]; Murloc Scale Breastplate (5781, -0.33 DPS) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 9.3 attack_power points (0.52 DPS) | yes | Forest Leather Bracers (3202, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.18 DPS) [vendor]; Ratchet Wristwraps (274742, -0.27 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 18.5 attack_power points (1.03 DPS) | yes | Gold-flecked Gloves (5195, -0.09 DPS) [dungeon]; Bristlebark Gloves (14572, -0.17 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.17 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Belt (252428, -0.14 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS) [crafted]; Ruffian Belt (5975, -0.20 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 27.8 attack_power points (1.55 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS, sim-verified) [crafted]; Defender's Leather Pants (252445, -0.01 DPS) [crafted]; Leggings of the Fang (10410, -0.13 DPS) [dungeon] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.4 attack_power points (1.08 DPS) | yes | Feet of the Lynx (1121, -0.02 DPS) [world_drop]; Defender's Leather Boots (252441, -0.41 DPS) [crafted]; Totemic Leather Boots (252442, -0.41 DPS) [crafted] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.5 attack_power points (0.87 DPS) | yes | Signet of the Zhevra (285330, -0.37 DPS) [world]; Loop of Sacrifice (281673, -0.46 DPS) [quest]; Bounty Hunter's Ring (5351, -0.62 DPS) [quest] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 9.6 attack_power points (0.54 DPS) | yes | Signet of the Zhevra (285330, -0.04 DPS) [world]; Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hammerbone (270018) | Leaders of the Fang [quest] | 310.2 attack_power points (17.29 DPS) | yes | Smite's Mighty Hammer (7230, +0.00 DPS, sim-verified) [dungeon]; Living Root (6631, -0.74 DPS) [dungeon]; Crescent Staff (6505, -0.84 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Grave Shroud; chest: Brawler's Leather Armor; wrist: Bristlebark Bindings; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Brawler's Leather Boots; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Hammerbone

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-55232220000000000000-0000000000000000)

Set DPS (verified): 110.5. Weights run: 2.8s. Verify run: 1.4s. 315 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.537 ± 0.011, crit=0.767 ± 0.015 per rating point (14 rating = 1%, 10.738 per %), hit=0.218 ± 0.006 per rating point (10 rating = 1%, 2.182 per %), melee_haste=6.889 ± 0.540

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 28.8 attack_power points (1.62 DPS) | yes | Azure Gustwoven Hood (277050, -0.40 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.54 DPS, sim-verified) [crafted]; Defender's Leather Hood (252447, -0.54 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.7 attack_power points (0.89 DPS) | yes | Scout's Medallion (19537, -0.19 DPS) [rep]; River Pride Choker (13087, -0.35 DPS) [world_drop]; Ghostshard Talisman (7731, -0.40 DPS, sim-verified) [dungeon] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 28.9 attack_power points (1.62 DPS) | yes | Bristlebark Amice (14573, -0.70 DPS) [world_drop]; Mantle of Thieves (2264, -0.76 DPS) [dungeon]; Barbaric Shoulders (5964, -0.79 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 18.0 attack_power points (1.01 DPS) | yes | Tigerstrike Mantle (13108, -0.42 DPS, sim-verified) [world_drop]; Grave Shroud (279865, -0.43 DPS) [quest]; Wildhunter Cloak (16658, -0.45 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 26.7 attack_power points (1.50 DPS) | yes | Brawler's Leather Armor (252490, -0.22 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS, sim-verified) [crafted]; Thick Murloc Armor (5782, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 18.8 attack_power points (1.06 DPS) | yes | Bands of Serra'kis (6902, -0.25 DPS) [dungeon]; Barbaric Bracers (18948, -0.26 DPS, sim-verified) [crafted]; Jurassic Wristguards (6198, -0.27 DPS) [world] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 26.0 attack_power points (1.46 DPS) | yes | Toughened Leather Gloves (4253, -0.13 DPS) [crafted]; Wolfclaw Gloves (1978, -0.27 DPS) [dungeon]; Brawler Gloves (720, -0.38 DPS) [world_drop] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 30.8 attack_power points (1.73 DPS) | yes | Skulker's Leather Belt (252520, -0.34 DPS, sim-verified) [crafted]; Defiler's Chain Girdle (20152, -0.38 DPS) [rep]; Defiler's Leather Girdle (20191, -0.38 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 30.6 attack_power points (1.72 DPS) | yes | Brawler's Leather Pants (252500, -0.13 DPS) [crafted]; Trapper's Leather Pants (252501, -0.13 DPS) [crafted]; Defender's Leather Pants (252445, -0.16 DPS) [crafted] |
| feet | Brawler's Leather Boots (252439) | Leatherworking [crafted] | 19.7 attack_power points (1.11 DPS) | yes | Feet of the Lynx (1121, -0.01 DPS) [world_drop]; Stomping Boots (3741, -0.22 DPS) [quest]; Insignia Boots (4055, -0.42 DPS) [world_drop] |
| finger1 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.8 attack_power points (1.34 DPS) | yes | Ironspine's Eye (7686, -0.02 DPS) [dungeon]; Band of the Fist (17694, -0.45 DPS) [quest]; Tiger Band (6749, -0.53 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 23.6 attack_power points (1.33 DPS) | yes | Ironspine's Eye (7686, -0.01 DPS) [dungeon]; Band of the Fist (17694, -0.44 DPS) [quest]; Tiger Band (6749, -0.52 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (110.5 DPS) | yes | Cobalt Crusher (7730, -2.66 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -3.57 DPS) [pvp]; Viscous Hammer (13045, -22.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Thunderbrow Ring; finger2: Legionnaire's Band; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-55232232121030000000-0000000000000000)

Set DPS (verified): 159.1. Weights run: 3.0s. Verify run: 1.4s. 426 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.608 ± 0.012, crit=0.868 ± 0.017 per rating point (14 rating = 1%, 12.152 per %), hit=0.219 ± 0.007 per rating point (10 rating = 1%, 2.190 per %), melee_haste=6.258 ± 0.736

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.1 attack_power points (2.92 DPS) | yes | Hawkeye's Helm (14591, -1.11 DPS) [world_drop]; Defender's Leather Helm (252455, -1.47 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -1.49 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.33 DPS) | yes | Ethereal Talisman (4430, -0.10 DPS) [quest]; Scout's Medallion (19536, -0.15 DPS) [rep]; Kaleidoscope Chain (13084, -0.26 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.7 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.00 DPS) [world_drop]; Flintrock Shoulders (7755, -0.11 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.53 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 25.7 attack_power points (1.70 DPS) | yes | First Sergeant's Cloak (16340, -0.11 DPS) [pvp]; Hawkeye's Cloak (14593, -0.48 DPS) [world_drop]; Parachute Cloak (10518, -0.85 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 49.7 attack_power points (3.30 DPS) | yes | Wolffear Harness (13110, -1.49 DPS) [world_drop]; Brawler's Leather Tunic (252508, -1.49 DPS) [crafted]; Barbaric Harness (5739, -1.89 DPS, sim-verified) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.33 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Barbaric Bracers (18948, -0.26 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 38.5 attack_power points (2.55 DPS) | yes | Skulker's Leather Gloves (252525, -0.05 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.42 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 36.1 attack_power points (2.39 DPS) | yes | Prowler's Leather Belt (252459, -0.32 DPS) [crafted]; Tharg's Shoelace (9705, -0.32 DPS) [quest]; Defiler's Leather Girdle (20192, -0.40 DPS) [rep] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 36.1 attack_power points (2.40 DPS) | yes | Basilisk Hide Pants (1718, -0.16 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.32 DPS) [crafted]; Brawler's Leather Pants (252500, -0.48 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.7 attack_power points (2.50 DPS) | yes | Excelsior Boots (4109, -0.42 DPS) [quest]; Imperial Leather Boots (6431, -0.53 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.69 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 32.1 attack_power points (2.13 DPS) | yes | Thunderbrow Ring (13097, -0.53 DPS) [world_drop]; Falcon's Hook (7552, -0.69 DPS) [dungeon]; Ring of the Underwood (2951, -0.74 DPS) [world_drop] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.1 attack_power points (1.60 DPS) | yes | Thunderbrow Ring (13097, -0.00 DPS) [world_drop]; Falcon's Hook (7552, -0.16 DPS) [dungeon]; Ring of the Underwood (2951, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (159.1 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Thornstone Sledgehammer (1722, -27.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Ironspine's Eye

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0100000000000000-55232232121032012001-3000000000000000)

Set DPS (verified): 188.0. Weights run: 3.0s. Verify run: 1.5s. 561 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.678 ± 0.023, crit=0.968 ± 0.033 per rating point (14 rating = 1%, 13.554 per %), hit=0.250 ± 0.008 per rating point (10 rating = 1%, 2.498 per %), melee_haste=6.986 ± 1.037

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 51.2 attack_power points (3.65 DPS) | yes | Undercity Reservist's Cap (20643, -0.62 DPS) [quest]; White Bandit Mask (10008, -0.66 DPS, sim-verified) [crafted]; Blood Guard's Leather Headband (220851, -1.37 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 33.8 attack_power points (2.41 DPS) | yes | Woven Ivy Necklace (19159, -0.31 DPS) [quest]; Scout's Medallion (19535, -0.98 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.99 DPS) [quest] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 41.5 attack_power points (2.96 DPS) | yes | Skulker's Leather Shoulder (252535, -0.10 DPS) [crafted]; Warden's Leather Shoulder (252536, -0.58 DPS) [crafted]; Failed Flying Experiment (9647, -0.71 DPS, sim-verified) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 37.9 attack_power points (2.70 DPS) | yes | Blisterbane Wrap (12552, -0.91 DPS) [dungeon]; Dark Phantom Cape (13122, -0.91 DPS) [world_drop]; Dark Hooded Cape (5257, -0.92 DPS, sim-verified) [world] |
| chest | Grizzled Pelt (22274) | A Better Ingredient [quest] | 67.6 attack_power points (4.83 DPS) | yes | Mixologist's Tunic (12793, -0.43 DPS) [dungeon]; Warbear Harness (15064, -0.79 DPS) [crafted]; Quillward Harness (10583, -1.18 DPS) [dungeon] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.8 attack_power points (2.48 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Prowler's Leather Bracers (252539, -0.10 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 60.9 attack_power points (4.35 DPS) | yes | Feralheart Fists (226793, -0.64 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -0.87 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.13 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.7 attack_power points (3.83 DPS) | yes | Skulker's Leather Waistguard (252474, -0.10 DPS) [crafted]; Girdle of Beastial Fury (11686, -0.32 DPS) [dungeon]; Ogron's Sash (13117, -1.22 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 58.0 attack_power points (4.14 DPS) | yes | Triprunner Dungarees (9624, -1.49 DPS, sim-verified) [quest]; Dragonflight Leggings (10742, -1.56 DPS) [quest]; Basilisk Hide Pants (1718, -1.63 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.7 attack_power points (3.54 DPS) | yes | Skulker's Leather Boots (252469, -0.10 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.48 DPS) [dungeon]; Shadefiend Boots (11675, -0.51 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 39.1 attack_power points (2.79 DPS) | yes | Thunderbrow Ring (13097, -1.06 DPS) [world_drop]; White Bone Band (11862, -1.08 DPS) [quest]; Masons Fraternity Ring (9533, -1.11 DPS) [quest] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 24.7 attack_power points (1.76 DPS) | yes | Thunderbrow Ring (13097, -0.03 DPS) [world_drop]; White Bone Band (11862, -0.05 DPS) [quest]; Masons Fraternity Ring (9533, -0.09 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (188.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (188.0 DPS) | yes | Molten Heart of the Mountain (249470, -0.94 DPS, sim-verified) [crafted] |
| main_hand | Ragehammer (10626) | Sunken Temple: Atal'ai Warrior [dungeon] | sim-verified (188.0 DPS) | yes | Glowing Brightwood Staff (812, +0.00 DPS) [world_drop]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; The Jackhammer (9423, -1.19 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Blackveil Cape; chest: Grizzled Pelt; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Ragehammer

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0100000000000000-55232232121032012001-5053000000000000)

Set DPS (verified): 281.0. Weights run: 3.0s. Verify run: 1.6s. 1446 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.400 ± 0.002, agility=1.849 ± 0.032, crit=1.213 ± 0.046 per rating point (14 rating = 1%, 16.987 per %), hit=0.299 ± 0.010 per rating point (10 rating = 1%, 2.989 per %), melee_haste=7.090 ± 1.511

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 87.4 attack_power points (6.42 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -3.53 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 59.1 attack_power points (4.34 DPS) | yes | Medallion of the Dawn (22659, -1.58 DPS, sim-verified) [quest]; Skibi's Pendant (13089, -1.70 DPS) [world_drop]; Imperial Jewel (11933, -1.99 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 66.9 attack_power points (4.92 DPS) | yes | Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Defiler's Leather Shoulders (20194, -0.27 DPS) [rep]; Warlord's Dragonhide Epaulets (16551, -0.55 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 47.7 attack_power points (3.51 DPS) | yes | Windshear Cape (20691, -0.06 DPS) [world]; Deathguard's Cloak (20068, -0.33 DPS) [rep]; Cloak of Revanchion (23127, -0.39 DPS) [dungeon] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | sim-verified (281.0 DPS) | yes | Warlord's Dragonhide Hauberk (16549, -1.95 DPS) [vendor]; Cadaverous Armor (14637, -2.61 DPS) [dungeon]; Tunic of Undead Slaying (23089, -10.75 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (281.0 DPS) | yes | Bracers of Subterfuge (22668, -0.72 DPS) [quest]; Bracers of the Eclipse (18375, -1.40 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.90 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 76.3 attack_power points (5.61 DPS) | yes | Timbermaw Brawlers (19049, -0.00 DPS) [crafted]; General's Dragonhide Grips (231688, -0.36 DPS) [vendor]; Studded Timbermaw Brawlers (227809, -0.72 DPS) [vendor] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 93.8 attack_power points (6.89 DPS) | yes | Might of the Timbermaw (19044, -1.29 DPS) [crafted]; Shifter's Belt (272396, -1.36 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -2.17 DPS) [quest] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 100.5 attack_power points (7.38 DPS) | yes | General's Dragonhide Leggings (231685, -0.35 DPS) [pvp]; Traveler's Leggings (8300, -2.08 DPS) [dungeon]; Sentinel's Leather Pants (237818, -2.21 DPS, sim-verified) [vendor] |
| feet | Drudge Boots (21532) | The Nightmare Manifests [quest] | 62.6 attack_power points (4.60 DPS) | yes | General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Blood Guard's Dragonhide Treads (227181, -0.19 DPS) [pvp]; Boots of Ferocity (22472, -0.54 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-verified (281.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.95 DPS) [quest]; Don Julio's Band (19325, -0.97 DPS) [rep]; Naglering (11669, -3.95 DPS, sim-verified) [dungeon] |
| finger2 | Myrmidon's Signet (2246) | World drop [world_drop] | sim-verified (281.0 DPS) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.05 DPS) [quest]; Don Julio's Band (19325, -0.07 DPS) [rep]; Naglering (11669, -2.90 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (281.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (281.0 DPS) | yes | Counterattack Lodestone (18537, -1.62 DPS) [dungeon]; Hand of Justice (11815, -1.77 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.04 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (281.0 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Felstriker (12590, -3.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (281.0 DPS) | yes | Howling Idol (272427, -4.61 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Drudge Boots; finger1: Legionnaire's Band; finger2: Myrmidon's Signet; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

