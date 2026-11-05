# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 47.0. Weights run: 7.4s. Verify run: 10.6s. 222 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.895 ± 0.010, crit=0.560 ± 0.009 per rating point (14 rating = 1%, 7.847 per %), hit=0.665 ± 0.033 per rating point (10 rating = 1%, 6.646 per %), melee_haste=2.873 ± 0.406

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.4 attack_power points (0.71 DPS) | yes | Erudite's Amulet (277204, -0.48 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.09 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.5 attack_power points (0.59 DPS) | yes | Slime-encrusted Pads (6461, -0.60 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (47.0 DPS) | yes | Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.21 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (47.0 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.47 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 11.7 attack_power points (0.73 DPS) | yes | Bristlebark Bindings (14569, -0.13 DPS) [world_drop]; Forest Leather Bracers (3202, -0.14 DPS) [world_drop]; Wolf Bracers (4794, -0.26 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (47.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.61 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.13 DPS) | yes | Deviate Scale Belt (6468, -0.16 DPS) [crafted]; Murloc Scale Belt (5780, -0.52 DPS) [crafted]; Brawler's Leather Belt (252428, -3.21 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (47.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.70 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.2 attack_power points (1.32 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Blackened Defias Boots (10402, -0.61 DPS) [dungeon]; Footpads of the Fang (10411, -0.61 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 15.6 attack_power points (0.97 DPS) | yes | The 1 Ring (8350, -0.73 DPS) [world]; Lavishly Jeweled Ring (1156, -0.74 DPS) [dungeon]; Ring of the Moon (12052, -0.85 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.50 DPS) | yes | The 1 Ring (8350, -0.26 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Ring of the Moon (12052, -0.38 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.6 attack_power points (19.12 DPS) | yes | Twisted Chanter's Staff (890, -1.35 DPS, sim-verified) [world_drop]; Rakzur Club (12983, -2.62 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -2.97 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 222, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 85.1. Weights run: 7.9s. Verify run: 16.2s. 390 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.987 ± 0.013, crit=0.790 ± 0.013 per rating point (14 rating = 1%, 11.065 per %), hit=0.937 ± 0.053 per rating point (10 rating = 1%, 9.365 per %), melee_haste=4.396 ± 0.641

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.77 DPS) | yes | Brawler's Leather Helm (252512, -0.34 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.44 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.9 attack_power points (1.18 DPS) | yes | Sentinel's Medallion (19541, -0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.9 attack_power points (2.35 DPS) | yes | Mantle of Thieves (2264, -0.89 DPS) [dungeon]; Barbaric Shoulders (5964, -0.90 DPS, sim-verified) [crafted]; Bristlebark Amice (14573, -1.03 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (85.1 DPS) | yes | Tigerstrike Mantle (13108, -0.30 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.30 DPS, sim-verified) [pvp]; Glowing Lizardscale Cloak (6449, -0.59 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.9 attack_power points (2.06 DPS) | yes | Dusky Leather Armor (7374, -0.01 DPS) [crafted]; Defender's Leather Tunic (252450, -0.29 DPS) [crafted]; Brawler's Leather Armor (252490, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.9 attack_power points (1.47 DPS) | yes | Barbaric Bracers (18948, -0.30 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.44 DPS) [quest]; Bands of Serra'kis (6902, -0.59 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.9 attack_power points (1.92 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.30 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.9 attack_power points (2.21 DPS) | yes | Skulker's Leather Belt (252520, -0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.44 DPS) [rep]; Highlander's Leather Girdle (20117, -0.44 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 31.9 attack_power points (2.36 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.29 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.9 attack_power points (1.62 DPS) | yes | Brawler's Leather Boots (252439, -0.14 DPS) [crafted]; Insignia Boots (4055, -0.44 DPS) [world_drop]; Disjointed Shoes (277226, -0.73 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 23.9 attack_power points (1.77 DPS) | yes | Seal of Wrynn (2933, -0.88 DPS) [quest]; Ring of Precision (1491, -0.89 DPS) [dungeon]; Tiger Band (6749, -1.03 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.0 attack_power points (1.62 DPS) | yes | Ring of Precision (1491, -0.74 DPS) [dungeon]; Seal of Wrynn (2933, -0.78 DPS, sim-verified) [quest]; Tiger Band (6749, -0.88 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (85.1 DPS) | yes | Viscous Hammer (13045, -2.92 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -3.20 DPS) [dungeon]; Gnarled Ash Staff (791, -6.59 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 125.5. Weights run: 10.2s. Verify run: 10.1s. 509 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.115 ± 0.028, crit=1.094 ± 0.033 per rating point (14 rating = 1%, 15.316 per %), hit=0.946 ± 0.078 per rating point (10 rating = 1%, 9.464 per %), melee_haste=6.318 ± 1.039

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.3 attack_power points (3.91 DPS) | yes | Hawkeye's Helm (14591, -1.22 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.53 DPS) [world_drop]; Nightscape Headband (8176, -1.72 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 23.3 attack_power points (2.01 DPS) | yes | Kaleidoscope Chain (13084, -0.59 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.97 DPS) [world]; River Pride Choker (13087, -1.32 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 attack_power points (2.87 DPS) | yes | Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Nightscape Shoulders (8192, -0.86 DPS) [crafted]; Fleshhide Shoulders (10774, -0.91 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (125.5 DPS) | yes | Yeti Fur Cloak (2805, -0.52 DPS) [quest]; Parachute Cloak (10518, -0.67 DPS) [crafted]; Hawkeye's Cloak (14593, -1.28 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 36.0 attack_power points (3.10 DPS) | yes | Barbaric Harness (5739, -0.43 DPS) [crafted]; Dusky Leather Armor (7374, -0.55 DPS) [crafted]; Nightscape Tunic (8175, -0.81 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.7 attack_power points (1.79 DPS) | yes | Branded Leather Bracers (19508, -0.06 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.33 DPS) [dungeon]; Dusky Bracers (7378, -0.33 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (125.5 DPS) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Prowler's Leather Gloves (252524, -1.73 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.0 attack_power points (3.20 DPS) | yes | Prowler's Leather Belt (252459, -0.55 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.61 DPS) [rep]; Skulker's Leather Belt (252520, -0.73 DPS, sim-verified) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.4 attack_power points (3.83 DPS) | yes | Triprunner Dungarees (9624, -0.45 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -0.98 DPS) [crafted]; Brawler's Leather Pants (252500, -1.15 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 37.3 attack_power points (3.22 DPS) | yes | Imperial Leather Boots (6431, -0.35 DPS) [dungeon]; Excelsior Boots (4109, -0.36 DPS) [quest]; Prowler's Leather Shoes (252465, -0.58 DPS, sim-verified) [crafted] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 32.9 attack_power points (2.84 DPS) | yes | Falcon's Hook (7552, -0.68 DPS) [dungeon]; Thunderbrow Ring (13097, -0.91 DPS) [world_drop]; Mark of Kern (2262, -1.11 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.2 attack_power points (2.17 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Mark of Kern (2262, -0.44 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Thornstone Sledgehammer (1722) | World drop [world_drop] | sim-verified (125.5 DPS) | yes | Mograine's Might (7723, -0.09 DPS) [dungeon]; Ironshod Bludgeon (9408, -2.06 DPS) [dungeon]; Mithril Warhammer (250608, -4.92 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Thornstone Sledgehammer

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 0100000000000000-50212032123012123001-5310000000000000)

Set DPS (verified): 176.7. Weights run: 10.1s. Verify run: 32.6s. 665 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.120 ± 0.002, agility=2.107 ± 0.030, crit=1.337 ± 0.039 per rating point (14 rating = 1%, 18.714 per %), hit=1.225 ± 0.083 per rating point (10 rating = 1%, 12.248 per %), melee_haste=6.360 ± 1.054

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 46.5 attack_power points (3.68 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.84 DPS) [crafted]; Scorpashi Skullcap (14658, -0.92 DPS, sim-verified) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 38.0 attack_power points (3.00 DPS) | yes | Sentinel's Medallion (19539, -1.00 DPS) [rep]; Kaleidoscope Chain (13084, -1.67 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 42.3 attack_power points (3.34 DPS) | yes | Failed Flying Experiment (9647, -0.00 DPS) [quest]; Skulker's Leather Shoulder (252535, -0.00 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.26 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant Major's Cape (16336, -0.49 DPS) [pvp]; Pridelord Cape (14673, -0.50 DPS) [world_drop]; Blackflame Cape (13109, -0.50 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 61.2 attack_power points (4.84 DPS) | yes | Blazewind Breastplate (11193, -0.51 DPS) [quest]; Knight's Leather Armor (220854, -0.97 DPS) [vendor]; Charred Leather Tunic (19127, -1.84 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 40.1 attack_power points (3.17 DPS) | yes | Prowler's Leather Bracers (252539, -0.49 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.50 DPS) [crafted]; Pridelord Bands (14672, -0.67 DPS) [world_drop] |
| hands | Prowler's Leather Gauntlets (252547) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -0.01 DPS) [crafted]; Serpentskin Gloves (8260, -0.17 DPS) [world_drop]; Heraldic Gloves (8121, -0.50 DPS) [world_drop] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 55.0 attack_power points (4.35 DPS) | yes | Skulker's Leather Waistguard (252474, -0.00 DPS) [crafted]; Highlander's Leather Girdle (20115, -1.28 DPS) [rep]; Ogron's Sash (13117, -1.34 DPS) [world_drop] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 63.4 attack_power points (5.01 DPS) | yes | Serpentskin Leggings (8262, -0.17 DPS) [world_drop]; Knight's Leather Pants (220858, -1.14 DPS) [vendor]; Triprunner Dungarees (9624, -1.51 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 50.7 attack_power points (4.01 DPS) | yes | Skulker's Leather Boots (252469, -0.00 DPS) [crafted]; Albino Crocscale Boots (17728, -0.68 DPS) [dungeon]; Prowler's Leather Shoes (252465, -1.00 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 40.2 attack_power points (3.18 DPS) | yes | Blackstone Ring (17713, -0.63 DPS) [dungeon]; Falcon's Hook (7552, -1.17 DPS) [dungeon]; Ring of the Underwood (2951, -1.17 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Falcon's Hook (7552, -0.33 DPS) [dungeon]; Ring of the Underwood (2951, -0.33 DPS) [world_drop]; Blackstone Ring (17713, -1.08 DPS, sim-verified) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.92 DPS, sim-verified) [crafted] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Iron Pulverizer (11608, -3.54 DPS, sim-verified) [crafted]; Soulkeeper (1607, -4.39 DPS) [world_drop]; Princess Theradras' Scepter (17766, -4.51 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Prowler's Leather Gauntlets; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 665, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 0100000000000000-50212032123012123001-5350330000000000)

Set DPS (verified): 294.2. Weights run: 10.3s. Verify run: 70.2s. 1443 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.120 ± 0.002, agility=2.548 ± 0.049, crit=2.260 ± 0.070 per rating point (14 rating = 1%, 31.642 per %), hit=2.005 ± 0.161 per rating point (10 rating = 1%, 20.052 per %), melee_haste=9.691 ± 1.927

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 101.0 attack_power points (8.62 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.34 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 43.7 attack_power points (3.73 DPS) | yes | Sentinel's Medallion (19538, -0.47 DPS) [rep]; Sunprism Pendant (20649, -1.12 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 100.4 attack_power points (8.57 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -0.73 DPS) [vendor]; Highlander's Leather Shoulders (20059, -1.60 DPS, sim-verified) [rep]; Lieutenant Commander's Dragonhide Shoulders (227172, -2.39 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (294.2 DPS) | yes | Cloak of the Honor Guard (20073, -0.98 DPS) [rep]; Dark Phantom Cape (13122, -1.71 DPS) [world_drop]; Sergeant Major's Cape (16337, -1.78 DPS) [pvp] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.5 attack_power points (12.58 DPS) | yes | Dawn Armor (252483, -0.13 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -1.51 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.32 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 71.7 attack_power points (6.12 DPS) | yes | Deepfury Bracers (13120, -2.14 DPS) [world_drop]; Shadowcraft Bracers (16710, -2.86 DPS) [dungeon]; Skulker's Leather Bracers (252540, -2.90 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (294.2 DPS) | yes | Marshal's Dragonhide Grips (231694, -0.27 DPS) [vendor]; Timbermaw Brawlers (19049, -1.02 DPS) [crafted]; Knight-Lieutenant's Dragonhide Grips (227183, -1.62 DPS) [pvp] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 80.2 attack_power points (6.84 DPS) | yes | Skulker's Leather Waistguard (252474, -1.63 DPS) [crafted]; Prowler's Leather Waistguard (252473, -1.70 DPS) [crafted]; Highlander's Leather Girdle (20045, -1.88 DPS, sim-verified) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 132.1 attack_power points (11.27 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Warbear Woolies (15065, -1.05 DPS, sim-verified) [crafted]; Knight-Captain's Dragonhide Leggings (227178, -2.08 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 59.5 attack_power points (5.08 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.26 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 65.7 attack_power points (5.61 DPS) | yes | Blackstone Ring (17713, -2.19 DPS) [dungeon]; Masons Fraternity Ring (9533, -2.56 DPS) [quest]; Band of the Penitent (13217, -2.91 DPS) [quest] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 53.5 attack_power points (4.56 DPS) | yes | Blackstone Ring (17713, -1.14 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.52 DPS) [quest]; Band of the Penitent (13217, -1.86 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (294.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (294.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (294.2 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -8.73 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (294.2 DPS) | yes | Idol of the Moon (23197, -9.19 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Might of the Timbermaw; legs: Sentinel's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Protector's Band; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 0500200000000000-50230032023011125001-0550300000000000)

Set DPS (verified): 724.7. Weights run: 10.7s. Verify run: 61.3s. 1443 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.425 ± 0.048, crit=1.717 ± 0.066 per rating point (14 rating = 1%, 24.042 per %), hit=2.473 ± 0.169 per rating point (10 rating = 1%, 24.727 per %), melee_haste=29.199 ± 1.927

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 87.0 attack_power points (17.73 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -3.06 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 43.6 attack_power points (8.89 DPS) | yes | Sentinel's Medallion (19538, -1.48 DPS) [rep]; Sunprism Pendant (20649, -2.96 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 89.5 attack_power points (18.25 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -0.51 DPS) [vendor]; Highlander's Leather Shoulders (20059, -3.47 DPS, sim-verified) [rep]; Dark Warder's Pauldrons (22241, -4.42 DPS) [dungeon] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (724.7 DPS) | yes | Cloak of the Honor Guard (20073, -2.09 DPS) [rep]; Sergeant Major's Cape (16337, -3.59 DPS) [pvp]; Dark Phantom Cape (13122, -4.08 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.5 attack_power points (30.08 DPS) | yes | Dawn Armor (252483, -1.85 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.35 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -5.92 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 72.7 attack_power points (14.82 DPS) | yes | Deepfury Bracers (13120, -5.43 DPS) [world_drop]; Skulker's Leather Bracers (252540, -6.92 DPS) [crafted]; Prowler's Leather Bracers (252539, -6.92 DPS) [crafted] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (724.7 DPS) | yes | Timbermaw Brawlers (19049, -1.49 DPS) [crafted]; Marshal's Dragonhide Grips (231694, -1.53 DPS) [vendor]; Knight-Lieutenant's Dragonhide Grips (227183, -2.38 DPS) [pvp] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 84.8 attack_power points (17.28 DPS) | yes | Prowler's Leather Waistguard (252473, -4.44 DPS) [crafted]; Skulker's Leather Waistguard (252474, -4.47 DPS, sim-verified) [crafted]; Highlander's Leather Girdle (20045, -5.45 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 113.5 attack_power points (23.15 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.35 DPS) [vendor]; Warbear Woolies (15065, -2.69 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.3 attack_power points (11.89 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.04 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 62.8 attack_power points (12.80 DPS) | yes | Blackstone Ring (17713, -3.68 DPS) [dungeon]; Masons Fraternity Ring (9533, -5.88 DPS) [quest]; Ring of the Underwood (2951, -6.87 DPS) [world_drop] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 55.7 attack_power points (11.36 DPS) | yes | Blackstone Ring (17713, -2.24 DPS) [dungeon]; Masons Fraternity Ring (9533, -4.44 DPS) [quest]; Ring of the Underwood (2951, -5.43 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (724.7 DPS) | yes | Thunderbrew's Boot Flask (744, -4.76 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (724.7 DPS) | yes | Thunderbrew's Boot Flask (744, -2.83 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (724.7 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -19.38 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (724.7 DPS) | yes | Idol of the Moon (23197, -17.32 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Might of the Timbermaw; legs: Sentinel's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Protector's Band; trinket2: Frozen Heart of the Mountain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 0000000000000000-50212010000000000000-0000000000000000)

Set DPS (verified): 46.1. Weights run: 7.4s. Verify run: 10.7s. 218 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.895 ± 0.010, crit=0.560 ± 0.009 per rating point (14 rating = 1%, 7.847 per %), hit=0.665 ± 0.033 per rating point (10 rating = 1%, 6.646 per %), melee_haste=2.873 ± 0.406

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 16.0 attack_power points (1.00 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.4 attack_power points (0.71 DPS) | yes | Erudite's Amulet (277204, -0.48 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.5 attack_power points (0.59 DPS) | yes | Slime-encrusted Pads (6461, -0.59 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (46.1 DPS) | yes | Dark Leather Cloak (2316, -0.11 DPS) [crafted]; Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Lambent Scale Cloak (4706, -0.21 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (46.1 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.39 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 9.7 attack_power points (0.61 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Spare Part Bindings (279875, -0.25 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.1 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.52 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.13 DPS) | yes | Deviate Scale Belt (6468, -0.16 DPS) [crafted]; Murloc Scale Belt (5780, -0.52 DPS) [crafted]; Brawler's Leather Belt (252428, -3.11 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (46.1 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -2.63 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.2 attack_power points (1.32 DPS) | yes | Brawler's Leather Boots (252439, -0.11 DPS) [crafted]; Blackened Defias Boots (10402, -0.61 DPS) [dungeon]; Footpads of the Fang (10411, -0.61 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 15.6 attack_power points (0.97 DPS) | yes | Loop of Sacrifice (281673, -0.60 DPS) [quest]; Bounty Hunter's Ring (5351, -0.62 DPS) [quest]; The 1 Ring (8350, -0.73 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.0 attack_power points (0.50 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; The 1 Ring (8350, -0.26 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 305.6 attack_power points (19.12 DPS) | yes | Hammerbone (270018, -0.49 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -1.21 DPS) [world_drop]; Rakzur Club (12983, -2.62 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer

### Band 30 (tauren, 0000000000000000-50212032123000000000-0000000000000000)

Set DPS (verified): 83.7. Weights run: 7.9s. Verify run: 19.6s. 391 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=1.987 ± 0.013, crit=0.790 ± 0.013 per rating point (14 rating = 1%, 11.065 per %), hit=0.937 ± 0.053 per rating point (10 rating = 1%, 9.365 per %), melee_haste=4.396 ± 0.641

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 24.0 attack_power points (1.77 DPS) | yes | Brawler's Leather Helm (252512, -0.30 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.44 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.44 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 15.9 attack_power points (1.18 DPS) | yes | Scout's Medallion (19537, -0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.29 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 31.9 attack_power points (2.35 DPS) | yes | Mantle of Thieves (2264, -0.89 DPS) [dungeon]; Barbaric Shoulders (5964, -0.95 DPS, sim-verified) [crafted]; Bristlebark Amice (14573, -1.03 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (83.7 DPS) | yes | Tigerstrike Mantle (13108, -0.31 DPS, sim-verified) [world_drop]; Glowing Lizardscale Cloak (6449, -0.59 DPS) [dungeon]; Swiftrunner Cape (6745, -0.59 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 27.9 attack_power points (2.06 DPS) | yes | Dusky Leather Armor (7374, -0.01 DPS) [crafted]; Defender's Leather Tunic (252450, -0.29 DPS) [crafted]; Brawler's Leather Armor (252490, -0.29 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.9 attack_power points (1.47 DPS) | yes | Barbaric Bracers (18948, -0.32 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.59 DPS) [dungeon]; Insignia Bracers (6410, -0.59 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 25.9 attack_power points (1.92 DPS) | yes | Toughened Leather Gloves (4253, -0.15 DPS) [crafted]; Wolfclaw Gloves (1978, -0.30 DPS) [dungeon]; Gloves of the Fang (10413, -0.44 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.9 attack_power points (2.21 DPS) | yes | Skulker's Leather Belt (252520, -0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.44 DPS) [rep]; Defiler's Leather Girdle (20191, -0.44 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 31.9 attack_power points (2.36 DPS) | yes | Brawler's Leather Pants (252500, -0.15 DPS) [crafted]; Trapper's Leather Pants (252501, -0.15 DPS) [crafted]; Barbaric Leggings (5963, -0.29 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 21.9 attack_power points (1.62 DPS) | yes | Brawler's Leather Boots (252439, -0.14 DPS) [crafted]; Stomping Boots (3741, -0.44 DPS) [quest]; Insignia Boots (4055, -0.44 DPS) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 23.9 attack_power points (1.77 DPS) | yes | Band of the Fist (17694, -0.59 DPS) [quest]; Ring of Precision (1491, -0.89 DPS) [dungeon]; Tiger Band (6749, -1.03 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 22.0 attack_power points (1.62 DPS) | yes | Band of the Fist (17694, -0.46 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.74 DPS) [dungeon]; Tiger Band (6749, -0.88 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (83.7 DPS) | yes | Viscous Hammer (13045, -2.71 DPS, sim-verified) [world_drop]; Cobalt Crusher (7730, -3.20 DPS) [dungeon]; Gnarled Ash Staff (791, -6.59 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-50212032123012123001-0000000000000000)

Set DPS (verified): 122.1. Weights run: 10.2s. Verify run: 12.2s. 508 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.000 ± 0.002, agility=2.115 ± 0.028, crit=1.094 ± 0.033 per rating point (14 rating = 1%, 15.316 per %), hit=0.946 ± 0.078 per rating point (10 rating = 1%, 9.464 per %), melee_haste=6.318 ± 1.039

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.3 attack_power points (3.91 DPS) | yes | Hawkeye's Helm (14591, -1.18 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -1.53 DPS) [world_drop]; Nightscape Headband (8176, -1.72 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 23.3 attack_power points (2.01 DPS) | yes | Ethereal Talisman (4430, -0.46 DPS, sim-verified) [quest]; Kaleidoscope Chain (13084, -0.59 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.97 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.3 attack_power points (2.87 DPS) | yes | Flintrock Shoulders (7755, -0.18 DPS) [dungeon]; Nightscape Shoulders (8192, -0.86 DPS) [crafted]; Fleshhide Shoulders (10774, -0.91 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.34 DPS) [dungeon]; Parachute Cloak (10518, -0.34 DPS) [crafted]; Tigerstrike Mantle (13108, -0.34 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 36.0 attack_power points (3.10 DPS) | yes | Barbaric Harness (5739, -0.43 DPS) [crafted]; Dusky Leather Armor (7374, -0.55 DPS) [crafted]; Nightscape Tunic (8175, -0.63 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.7 attack_power points (1.79 DPS) | yes | Branded Leather Bracers (19508, -0.06 DPS) [dungeon]; Imperial Leather Bracers (4061, -0.33 DPS) [dungeon]; Dusky Bracers (7378, -0.33 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -0.87 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.0 attack_power points (3.20 DPS) | yes | Skulker's Leather Belt (252520, -0.52 DPS) [crafted]; Prowler's Leather Belt (252459, -0.55 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.61 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 44.4 attack_power points (3.83 DPS) | yes | Triprunner Dungarees (9624, -0.03 DPS) [quest]; Brawler's Leather Legguards (252516, -0.98 DPS) [crafted]; Brawler's Leather Pants (252500, -1.15 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 37.3 attack_power points (3.22 DPS) | yes | Prowler's Leather Shoes (252465, -0.04 DPS) [crafted]; Imperial Leather Boots (6431, -0.35 DPS) [dungeon]; Excelsior Boots (4109, -0.36 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 32.9 attack_power points (2.84 DPS) | yes | Falcon's Hook (7552, -0.68 DPS) [dungeon]; Thunderbrow Ring (13097, -0.91 DPS) [world_drop]; Mark of Kern (2262, -1.11 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 25.2 attack_power points (2.17 DPS) | yes | Falcon's Hook (7552, -0.01 DPS) [dungeon]; Thunderbrow Ring (13097, -0.24 DPS) [world_drop]; Mark of Kern (2262, -0.44 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Mograine's Might (7723) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Thornstone Sledgehammer (1722, -0.18 DPS, sim-verified) [world_drop]; Ironshod Bludgeon (9408, -1.96 DPS) [dungeon]; Cragwood Maul (11265, -3.41 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; hands: Gloves of Holy Might; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Mograine's Might

No-known-source sample (15 of 508, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 0100000000000000-50212032123012123001-5310000000000000)

Set DPS (verified): 173.2. Weights run: 10.1s. Verify run: 30.3s. 664 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.120 ± 0.002, agility=2.329 ± 0.036, crit=1.587 ± 0.046 per rating point (14 rating = 1%, 22.217 per %), hit=1.494 ± 0.101 per rating point (10 rating = 1%, 14.938 per %), melee_haste=6.040 ± 1.318

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 48.9 attack_power points (4.26 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.18 DPS) [quest]; Helm of Fire (8348, -0.81 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.9 attack_power points (3.56 DPS) | yes | Woven Ivy Necklace (19159, -1.07 DPS, sim-verified) [quest]; Scout's Medallion (19535, -1.13 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 44.7 attack_power points (3.89 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Prowler's Leather Shoulders (252534, -0.04 DPS) [crafted]; Failed Flying Experiment (9647, -0.69 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (173.2 DPS) | yes | Pridelord Cape (14673, -0.66 DPS) [world_drop]; Serpentskin Cloak (8259, -0.81 DPS) [world_drop]; Blackflame Cape (13109, -1.32 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 65.2 attack_power points (5.68 DPS) | yes | Blazewind Breastplate (11193, -0.58 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -0.88 DPS) [vendor]; Charred Leather Tunic (19127, -2.03 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 43.4 attack_power points (3.78 DPS) | yes | Prowler's Leather Bracers (252539, -0.70 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.78 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.83 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (173.2 DPS) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 58.0 attack_power points (5.06 DPS) | yes | Prowler's Leather Waistguard (252473, -0.04 DPS) [crafted]; Defiler's Leather Girdle (20193, -1.38 DPS) [rep]; Ogron's Sash (13117, -1.57 DPS) [world_drop] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 64.8 attack_power points (5.65 DPS) | yes | Stone Guard's Leather Pants (220859, -0.81 DPS, sim-verified) [vendor]; Basilisk Hide Pants (1718, -1.39 DPS) [world_drop]; Triprunner Dungarees (9624, -1.44 DPS) [quest] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 53.6 attack_power points (4.67 DPS) | yes | Prowler's Leather Boots (252468, -0.04 DPS) [crafted]; Albino Crocscale Boots (17728, -0.61 DPS) [dungeon]; Skulker's Leather Shoes (252531, -1.14 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 42.2 attack_power points (3.67 DPS) | yes | Masons Fraternity Ring (9533, -0.83 DPS) [quest]; Ring of the Underwood (2951, -1.27 DPS) [world_drop]; Falcon's Hook (7552, -1.29 DPS) [dungeon] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.9 attack_power points (3.04 DPS) | yes | Masons Fraternity Ring (9533, -0.20 DPS) [quest]; Ring of the Underwood (2951, -0.65 DPS) [world_drop]; Falcon's Hook (7552, -0.66 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (173.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (173.2 DPS) | yes | Frozen Heart of the Mountain (249469, -0.96 DPS) [crafted]; Ankh of Life (1713, -1.82 DPS, sim-verified) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (173.2 DPS) | yes | Dark Iron Pulverizer (11608, -3.82 DPS, sim-verified) [crafted]; The Rockpounder (9413, -4.77 DPS) [dungeon]; Soulkeeper (1607, -4.83 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: Legionnaire's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Thorium Greatmace

No-known-source sample (15 of 664, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 0100000000000000-50212032123012123001-5350330000000000)

Set DPS (verified): 289.1. Weights run: 10.3s. Verify run: 66.7s. 1448 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.120 ± 0.002, agility=2.548 ± 0.049, crit=2.260 ± 0.070 per rating point (14 rating = 1%, 31.642 per %), hit=2.005 ± 0.161 per rating point (10 rating = 1%, 20.052 per %), melee_haste=9.691 ± 1.927

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 101.0 attack_power points (8.62 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -1.81 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 43.7 attack_power points (3.73 DPS) | yes | Scout's Medallion (19534, -0.47 DPS) [rep]; Woven Ivy Necklace (19159, -0.69 DPS) [quest]; Sunprism Pendant (20649, -1.12 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 100.4 attack_power points (8.57 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -0.73 DPS) [vendor]; Defiler's Leather Shoulders (20194, -2.05 DPS, sim-verified) [rep]; Champion's Dragonhide Shoulders (227175, -2.39 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (289.1 DPS) | yes | Deathguard's Cloak (20068, -1.21 DPS, sim-verified) [rep]; Dark Phantom Cape (13122, -1.71 DPS) [world_drop]; Fel Cape (279269, -2.27 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.5 attack_power points (12.58 DPS) | yes | Dawn Armor (252483, -0.13 DPS) [crafted]; Warlord's Dragonhide Chestpiece (231686, -1.51 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.32 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 71.7 attack_power points (6.12 DPS) | yes | Deepfury Bracers (13120, -2.14 DPS) [world_drop]; Shadowcraft Bracers (16710, -2.86 DPS) [dungeon]; Skulker's Leather Bracers (252540, -2.90 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (289.1 DPS) | yes | General's Dragonhide Grips (231688, -0.27 DPS) [vendor]; Timbermaw Brawlers (19049, -1.02 DPS) [crafted]; Gloves of Holy Might (867, -3.09 DPS, sim-verified) [world_drop] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 80.2 attack_power points (6.84 DPS) | yes | Defiler's Leather Girdle (20190, -1.44 DPS, sim-verified) [rep]; Skulker's Leather Waistguard (252474, -1.63 DPS) [crafted]; Prowler's Leather Waistguard (252473, -1.70 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 105.2 attack_power points (8.98 DPS) | yes | Legionnaire's Dragonhide Leggings (227177, +0.00 DPS) [pvp]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -0.19 DPS) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 59.5 attack_power points (5.08 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.26 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 65.7 attack_power points (5.61 DPS) | yes | Blackstone Ring (17713, -2.19 DPS) [dungeon]; Masons Fraternity Ring (9533, -2.56 DPS) [quest]; Band of the Penitent (13217, -2.91 DPS) [quest] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 53.5 attack_power points (4.56 DPS) | yes | Blackstone Ring (17713, -1.14 DPS) [dungeon]; Masons Fraternity Ring (9533, -1.52 DPS) [quest]; Band of the Penitent (13217, -1.86 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (289.1 DPS) | yes | Frozen Heart of the Mountain (249469, -0.85 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (289.1 DPS) | yes | Frozen Heart of the Mountain (249469, -1.88 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (289.1 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -8.84 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (289.1 DPS) | yes | Idol of the Moon (23197, -12.56 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 0500200000000000-50230032023011125001-0550300000000000)

Set DPS (verified): 715.7. Weights run: 10.7s. Verify run: 54.6s. 1448 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.002, agility=2.425 ± 0.048, crit=1.717 ± 0.066 per rating point (14 rating = 1%, 24.042 per %), hit=2.473 ± 0.169 per rating point (10 rating = 1%, 24.727 per %), melee_haste=29.199 ± 1.927

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 87.0 attack_power points (17.73 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -5.35 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 43.6 attack_power points (8.89 DPS) | yes | Scout's Medallion (19534, -1.48 DPS) [rep]; Woven Ivy Necklace (19159, -1.48 DPS) [quest]; Sunprism Pendant (20649, -2.96 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 89.5 attack_power points (18.25 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -0.51 DPS) [vendor]; Dark Warder's Pauldrons (22241, -4.42 DPS) [dungeon]; Defiler's Leather Shoulders (20194, -6.38 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (715.7 DPS) | yes | Deathguard's Cloak (20068, -3.81 DPS, sim-verified) [rep]; Dark Phantom Cape (13122, -4.08 DPS) [world_drop]; Blackflame Cape (13109, -5.56 DPS) [world_drop] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 147.5 attack_power points (30.08 DPS) | yes | Dawn Armor (252483, -1.65 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.35 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -5.92 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 72.7 attack_power points (14.82 DPS) | yes | Deepfury Bracers (13120, -5.43 DPS) [world_drop]; Skulker's Leather Bracers (252540, -6.92 DPS) [crafted]; Prowler's Leather Bracers (252539, -6.92 DPS) [crafted] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (715.7 DPS) | yes | General's Dragonhide Grips (231688, -1.53 DPS) [vendor]; Blood Guard's Dragonhide Grips (227180, -2.38 DPS) [pvp]; Timbermaw Brawlers (19049, -3.79 DPS, sim-verified) [crafted] |
| waist | Might of the Timbermaw (19044) | Leatherworking [crafted] | 84.8 attack_power points (17.28 DPS) | yes | Skulker's Leather Waistguard (252474, -4.32 DPS, sim-verified) [crafted]; Prowler's Leather Waistguard (252473, -4.44 DPS) [crafted]; Defiler's Leather Girdle (20190, -5.45 DPS) [rep] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 111.4 attack_power points (22.71 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Legionnaire's Dragonhide Leggings (227177, -0.92 DPS) [pvp]; Outrider's Leather Pants (22740, -2.17 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 58.3 attack_power points (11.89 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Skulker's Leather Boots (252469, -0.04 DPS) [crafted] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 62.8 attack_power points (12.80 DPS) | yes | Blackstone Ring (17713, -3.68 DPS) [dungeon]; Masons Fraternity Ring (9533, -5.88 DPS) [quest]; Ring of the Underwood (2951, -6.87 DPS) [world_drop] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 55.7 attack_power points (11.36 DPS) | yes | Blackstone Ring (17713, -2.24 DPS) [dungeon]; Masons Fraternity Ring (9533, -4.44 DPS) [quest]; Ring of the Underwood (2951, -5.43 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (715.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (715.7 DPS) | yes | Frozen Heart of the Mountain (249469, -1.85 DPS) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (715.7 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hammer of the Titans (12796, -19.33 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (715.7 DPS) | yes | Idol of the Moon (23197, -14.68 DPS, sim-verified) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Might of the Timbermaw; legs: Warbear Woolies; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Legionnaire's Band; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

