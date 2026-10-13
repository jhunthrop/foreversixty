# Leveling BiS: Feral

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 67.3. Weights run: 2.3s. Verify run: 4.8s. 253 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.086 ± 0.018, crit=0.676 ± 0.017 per rating point (14 rating = 1%, 9.469 per %), hit=0.873 ± 0.041 per rating point (10 rating = 1%, 8.731 per %), melee_haste=5.840 ± 0.501

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.96 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.5 attack_power points (0.68 DPS) | yes | Erudite's Amulet (277204, -0.51 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.4 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.62 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.5 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.09 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (67.3 DPS) | yes | Tunic of Westfall (2041, +0.00 DPS) [quest]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.03 DPS, sim-verified) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 12.7 attack_power points (0.69 DPS) | yes | Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Forest Leather Bracers (3202, -0.13 DPS) [world_drop]; Wolf Bracers (4794, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.3 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.18 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.98 DPS) | yes | Deviate Scale Belt (6468, -0.05 DPS) [crafted]; Murloc Scale Belt (5780, -0.40 DPS) [crafted]; Brawler's Leather Belt (252428, -3.69 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (67.3 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.29 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.3 attack_power points (1.27 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.59 DPS) [dungeon]; Footpads of the Fang (10411, -0.59 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 17.1 attack_power points (0.93 DPS) | yes | The 1 Ring (8350, -0.70 DPS) [world]; Lavishly Jeweled Ring (1156, -0.70 DPS) [dungeon]; Ring of the Moon (12052, -0.81 DPS) [world_drop] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.8 attack_power points (0.48 DPS) | yes | The 1 Ring (8350, -0.20 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Ring of the Moon (12052, -0.36 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.5 attack_power points (16.76 DPS) | yes | Staff of Westfall (2042, -1.17 DPS) [quest]; Twisted Chanter's Staff (890, -1.21 DPS) [world_drop]; Living Root (6631, -1.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 253, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 3000 Brood Mother Carapace; 3018 Hide of Lupos

### Band 30 (night-elf, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 121.9. Weights run: 1.8s. Verify run: 3.0s. 424 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.996 ± 0.022, crit=0.810 ± 0.023 per rating point (14 rating = 1%, 11.335 per %), hit=0.939 ± 0.044 per rating point (10 rating = 1%, 9.385 per %), melee_haste=4.983 ± 0.639

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.72 DPS) | yes | Cloudy Gustwoven Hood (277042, -0.43 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.43 DPS) [crafted]; Brawler's Leather Helm (252512, -0.45 DPS, sim-verified) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.8 attack_power points (1.09 DPS) | yes | Sentinel's Medallion (19541, -0.05 DPS) [rep]; Fallen Guard's Pendant (279837, -0.23 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.0 attack_power points (2.14 DPS) | yes | Mantle of Thieves (2264, -0.84 DPS) [dungeon]; Bristlebark Amice (14573, -0.93 DPS) [world_drop]; Barbaric Shoulders (5964, -1.16 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.6 attack_power points (1.34 DPS) | yes | Sergeant Major's Cape (16315, -0.25 DPS) [pvp]; Tigerstrike Mantle (13108, -0.30 DPS) [world_drop]; Fenrus' Hide (6340, -0.56 DPS) [dungeon] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.2 attack_power points (1.90 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.8 attack_power points (1.35 DPS) | yes | Barbaric Bracers (18948, -0.37 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.40 DPS) [quest]; Bands of Serra'kis (6902, -0.49 DPS) [dungeon] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.4 attack_power points (1.78 DPS) | yes | Toughened Leather Gloves (4253, -0.14 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.43 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 31.8 attack_power points (2.07 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Warden's Leather Belt (252460, -0.43 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.51 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 33.4 attack_power points (2.17 DPS) | yes | Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Trapper's Leather Pants (252501, -0.14 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.6 attack_power points (1.47 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Insignia Boots (4055, -0.43 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.43 DPS) [vendor] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 25.2 attack_power points (1.64 DPS) | yes | Tiger Band (6749, -0.64 DPS) [quest]; Monkey Ring (6748, -0.73 DPS) [quest]; Seal of Wrynn (2933, -0.82 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.6 attack_power points (1.53 DPS) | yes | Monkey Ring (6748, -0.63 DPS) [quest]; Seal of Wrynn (2933, -0.72 DPS) [quest]; Tiger Band (6749, -0.76 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (121.9 DPS) | yes | Cobalt Crusher (7730, -3.02 DPS) [dungeon]; Gnarled Ash Staff (791, -6.01 DPS) [world_drop]; Viscous Hammer (13045, -24.15 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2944 Cursed Eye of Paleth

### Band 40 (night-elf, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 166.5. Weights run: 2.5s. Verify run: 3.2s. 539 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.917 ± 0.029, crit=0.922 ± 0.034 per rating point (14 rating = 1%, 12.906 per %), hit=0.919 ± 0.058 per rating point (10 rating = 1%, 9.193 per %), melee_haste=6.120 ± 0.774

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.3 attack_power points (3.36 DPS) | yes | Defender's Leather Helm (252455, -1.40 DPS) [crafted]; Warden's Wizard Hat (14604, -1.51 DPS) [world_drop]; Hawkeye's Helm (14591, -1.55 DPS, sim-verified) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 21.1 attack_power points (1.57 DPS) | yes | Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.59 DPS) [quest]; River Pride Choker (13087, -0.91 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 33.1 attack_power points (2.46 DPS) | yes | Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop]; Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.79 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 24.7 attack_power points (1.83 DPS) | yes | Yeti Fur Cloak (2805, -0.49 DPS) [quest]; Parachute Cloak (10518, -0.70 DPS) [crafted]; Hawkeye's Cloak (14593, -1.13 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 32.6 attack_power points (2.42 DPS) | yes | Nightscape Tunic (8175, -0.28 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.30 DPS) [crafted]; Barbaric Harness (5739, -0.55 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.3 attack_power points (1.51 DPS) | yes | Branded Leather Bracers (19508, -0.02 DPS) [dungeon]; Barbaric Bracers (18948, -0.28 DPS) [crafted]; Dusky Bracers (7378, -0.37 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 39.3 attack_power points (2.91 DPS) | yes | Skulker's Leather Gloves (252525, -0.02 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.47 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.1 attack_power points (2.75 DPS) | yes | Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.52 DPS) [rep] |
| legs | Triprunner Dungarees (9624) | The Grand Betrayal [quest] | 41.1 attack_power points (3.05 DPS) | yes | Basilisk Hide Pants (1718, -0.06 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.63 DPS) [crafted]; Brawler's Leather Pants (252500, -0.79 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.6 attack_power points (2.79 DPS) | yes | Skulker's Leather Shoes (252531, -0.08 DPS) [crafted]; Excelsior Boots (4109, -0.37 DPS) [quest]; Imperial Leather Boots (6431, -0.41 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 32.9 attack_power points (2.44 DPS) | yes | Ring of the Underwood (2951, -0.70 DPS) [world_drop]; Thunderbrow Ring (13097, -0.71 DPS) [world_drop]; Mark of Kern (2262, -0.96 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.9 attack_power points (1.77 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (166.5 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Gut Ripper (2164, -31.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Wolffear Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Protector's Band; finger2: Falcon's Hook

No-known-source sample (15 of 539, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

### Band 50 (night-elf, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 216.8. Weights run: 2.1s. Verify run: 3.6s. 677 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.864 ± 0.035, crit=1.058 ± 0.046 per rating point (14 rating = 1%, 14.814 per %), hit=1.017 ± 0.074 per rating point (10 rating = 1%, 10.166 per %), melee_haste=5.627 ± 1.110

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.7 attack_power points (3.59 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -0.30 DPS) [vendor]; Scorpashi Skullcap (14658, -0.96 DPS) [world_drop]; Helm of Fire (8348, -1.05 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (2.83 DPS) | yes | Sentinel's Medallion (19540, -1.18 DPS) [rep]; Kaleidoscope Chain (13084, -1.52 DPS) [world_drop] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 41.0 attack_power points (3.29 DPS) | yes | Skulker's Leather Shoulder (252535, -0.05 DPS) [crafted]; Knight-Lieutenant's Leather Shoulders (220852, -0.64 DPS) [vendor]; Failed Flying Experiment (9647, -0.95 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 28.0 attack_power points (2.25 DPS) | yes | Sergeant Major's Cape (16336, -0.29 DPS) [pvp]; Pridelord Cape (14673, -0.37 DPS) [world_drop]; Blackflame Cape (13109, -0.45 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.8 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.87 DPS, sim-verified) [quest]; Knight's Leather Armor (220854, -1.19 DPS) [vendor]; Charred Leather Tunic (19127, -1.94 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.8 attack_power points (2.95 DPS) | yes | Prowler's Leather Bracers (252539, -0.31 DPS) [crafted]; Skulker's Leather Bracers (252540, -0.37 DPS) [crafted]; Pridelord Bands (14672, -0.57 DPS) [world_drop] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 63.0 attack_power points (5.06 DPS) | yes | Feralheart Fists (226793, -0.94 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -1.24 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.37 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.2 attack_power points (4.27 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Ogron's Sash (13117, -1.33 DPS) [world_drop]; Highlander's Leather Girdle (20115, -1.47 DPS) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 61.0 attack_power points (4.90 DPS) | yes | Serpentskin Leggings (8262, -0.20 DPS) [world_drop]; Knight's Leather Pants (220858, -1.44 DPS) [vendor]; Triprunner Dungarees (9624, -1.67 DPS) [quest] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.1 attack_power points (3.94 DPS) | yes | Skulker's Leather Boots (252469, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.95 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.95 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 38.8 attack_power points (3.11 DPS) | yes | Falcon's Hook (7552, -1.24 DPS) [dungeon]; Thunderbrow Ring (13097, -1.25 DPS) [world_drop]; Ring of the Underwood (2951, -1.26 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.1 attack_power points (2.10 DPS) | yes | Falcon's Hook (7552, -0.22 DPS) [dungeon]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop]; Ring of the Underwood (2951, -0.25 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (216.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (216.8 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (216.8 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -36.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Gryphon Rider's Leggings; feet: Prowler's Leather Boots; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 677, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

### Band 60 (night-elf, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 286.2. Weights run: 1.9s. Verify run: 11.5s. 1319 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.791 ± 0.041, crit=1.130 ± 0.058 per rating point (14 rating = 1%, 15.823 per %), hit=1.103 ± 0.093 per rating point (10 rating = 1%, 11.030 per %), melee_haste=not significant (6.176 ± 1.572)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 81.9 attack_power points (6.99 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -6.16 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.0 attack_power points (4.78 DPS) | yes | Medallion of the Dawn (22659, -1.56 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -1.79 DPS) [quest]; Skibi's Pendant (13089, -1.86 DPS) [world_drop] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (286.2 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Stormshroud Shoulders (15058, -7.50 DPS, sim-verified) [crafted] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.9 attack_power points (4.00 DPS) | yes | Cloak of the Honor Guard (20073, -0.33 DPS) [rep]; Howler's Furs (272414, -0.67 DPS) [vendor]; Sergeant Major's Cape (16337, -1.28 DPS) [pvp] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 120.7 attack_power points (10.30 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, -2.11 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -2.11 DPS) [vendor]; Dawn Armor (252483, -2.28 DPS, sim-verified) [crafted] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (286.2 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Feralheart Fists (226793) | Mokvar [vendor] | sim-verified (286.2 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Stormshroud Gloves (21278, -6.23 DPS, sim-verified) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 87.7 attack_power points (7.48 DPS) | yes | Shifter's Belt (272396, -1.30 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -1.40 DPS) [crafted]; Belt of Preserved Heads (20216, -1.62 DPS) [quest] |
| legs | Feralheart Trousers (226791) | Mokvar [vendor] | sim-verified (286.2 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -6.38 DPS, sim-verified) [crafted] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (286.2 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Drudge Boots (21532, -3.24 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 63.8 attack_power points (5.45 DPS) | yes | Don Julio's Band (19325, -1.79 DPS) [rep]; Cutthroat's Signet (272408, -2.37 DPS) [vendor]; Masons Fraternity Ring (9533, -3.31 DPS) [quest] |
| finger2 | Protector's Band (19514) | Silverwing Sentinels [rep] | 46.1 attack_power points (3.94 DPS) | yes | Don Julio's Band (19325, -0.28 DPS) [rep]; Cutthroat's Signet (272408, -0.85 DPS) [vendor]; Masons Fraternity Ring (9533, -1.79 DPS) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (286.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (286.2 DPS) | yes | Darkmoon Card: Heroism (19287, -1.91 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (286.2 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Electrified Dagger (19100, -6.74 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (286.2 DPS) | yes | Howling Idol (272427, -7.29 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Feralheart Fists; waist: Ferocity of the Timbermaw; legs: Feralheart Trousers; feet: Feralheart Walkers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Protector's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1319, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

### Band 60, raid preset (night-elf, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 658.7. Weights run: 2.2s. Verify run: 12.9s. 1319 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.003, agility=2.290 ± 0.049, crit=1.532 ± 0.066 per rating point (14 rating = 1%, 21.450 per %), hit=1.821 ± 0.117 per rating point (10 rating = 1%, 18.209 per %), melee_haste=13.757 ± 1.827

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 99.2 attack_power points (15.23 DPS) | yes | Field Marshal's Dragonhide Helmet (16451, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -22.44 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 67.7 attack_power points (10.39 DPS) | yes | Medallion of the Dawn (22659, -3.83 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -3.91 DPS) [quest]; Skibi's Pendant (13089, -3.97 DPS) [world_drop] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (658.7 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -14.93 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.3 attack_power points (8.34 DPS) | yes | Howler's Furs (272414, -1.25 DPS) [vendor]; Cloak of the Honor Guard (20073, -1.37 DPS) [rep]; Shifting Cloak (18511, -2.37 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 142.5 attack_power points (21.87 DPS) | yes | Dawn Armor (252483, -2.78 DPS) [crafted]; Field Marshal's Dragonhide Chestpiece (231690, -3.44 DPS) [vendor]; Field Marshal's Dragonhide Breastplate (16452, -4.38 DPS) [vendor] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (658.7 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 88.4 attack_power points (13.57 DPS) | yes | Timbermaw Brawlers (19049, -0.85 DPS) [crafted]; Marshal's Dragonhide Grips (231694, -1.20 DPS) [vendor]; Knight-Lieutenant's Dragonhide Grips (227183, -2.44 DPS) [pvp] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 101.8 attack_power points (15.63 DPS) | yes | Shifter's Belt (272396, -1.96 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -2.36 DPS) [quest]; Might of the Timbermaw (19044, -2.91 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.0 attack_power points (16.72 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Sentinel's Leather Pants (237818, -0.65 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.96 DPS) [vendor] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (658.7 DPS) | yes | Knight-Lieutenant's Dragonhide Treads (227182, +0.00 DPS) [pvp]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Drudge Boots (21532, -16.72 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 76.3 attack_power points (11.71 DPS) | yes | Protector's Band (19514, -3.39 DPS) [rep]; Cutthroat's Signet (272408, -3.99 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 55.7 attack_power points (8.54 DPS) | yes | Protector's Band (19514, -0.22 DPS) [rep]; Cutthroat's Signet (272408, -0.83 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (658.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (658.7 DPS) | yes | Molten Heart of the Mountain (249470, -2.45 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (658.7 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Electrified Dagger (19100, -12.11 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (658.7 DPS) | yes | Idol of the Moon (23197, -5.03 DPS, sim-verified) [world_drop] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Feralheart Walkers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1319, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 2622 Nimar's Tribal Headdress

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 66.0. Weights run: 2.3s. Verify run: 4.6s. 243 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=2.086 ± 0.018, crit=0.676 ± 0.017 per rating point (14 rating = 1%, 9.469 per %), hit=0.873 ± 0.041 per rating point (10 rating = 1%, 8.731 per %), melee_haste=5.840 ± 0.501

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Hood (252447) | Leatherworking [crafted] | 17.6 attack_power points (0.96 DPS) | yes | Brawler's Leather Hood (252504, -0.05 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.5 attack_power points (0.68 DPS) | yes | Erudite's Amulet (277204, -0.50 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.4 attack_power points (0.57 DPS) | yes | Slime-encrusted Pads (6461, -0.63 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.5 attack_power points (0.68 DPS) | yes | Grave Shroud (279865, -0.09 DPS) [quest]; Dark Leather Cloak (2316, -0.10 DPS) [crafted]; Cape of the Brotherhood (5193, -0.11 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (66.0 DPS) | yes | Murloc Scale Breastplate (5781, +0.00 DPS) [crafted]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -3.09 DPS, sim-verified) [crafted] |
| wrist | Bristlebark Bindings (14569) | World drop [world_drop] | 10.7 attack_power points (0.58 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Ratchet Wristwraps (274742, -0.24 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.0 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -3.21 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.98 DPS) | yes | Deviate Scale Belt (6468, -0.05 DPS) [crafted]; Murloc Scale Belt (5780, -0.40 DPS) [crafted]; Brawler's Leather Belt (252428, -3.72 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (66.0 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.34 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 23.3 attack_power points (1.27 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Blackened Defias Boots (10402, -0.59 DPS) [dungeon]; Footpads of the Fang (10411, -0.59 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 17.1 attack_power points (0.93 DPS) | yes | Loop of Sacrifice (281673, -0.57 DPS) [quest]; Bounty Hunter's Ring (5351, -0.59 DPS) [quest]; The 1 Ring (8350, -0.70 DPS) [world] |
| finger2 | Demon Band (12054) | World drop [world_drop] | 8.8 attack_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; The 1 Ring (8350, -0.25 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Smite's Mighty Hammer (7230) | The Deadmines: Mr. Smite [dungeon] | 308.5 attack_power points (16.76 DPS) | yes | Crescent Staff (6505, -0.50 DPS) [quest]; Hammerbone (270018, -0.62 DPS, sim-verified) [quest]; Living Root (6631, -0.63 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Defender's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Bristlebark Bindings; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Demon Band; main_hand: Smite's Mighty Hammer

No-known-source sample (15 of 243, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer; 3000 Brood Mother Carapace

### Band 30 (tauren, 0000000000000000-55232031000000000000-0000000000000000)

Set DPS (verified): 119.7. Weights run: 1.8s. Verify run: 3.1s. 416 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.002, agility=1.996 ± 0.022, crit=0.810 ± 0.023 per rating point (14 rating = 1%, 11.335 per %), hit=0.939 ± 0.044 per rating point (10 rating = 1%, 9.385 per %), melee_haste=4.983 ± 0.639

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 26.4 attack_power points (1.72 DPS) | yes | Brawler's Leather Helm (252512, -0.40 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.43 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.43 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 16.8 attack_power points (1.09 DPS) | yes | Scout's Medallion (19537, -0.05 DPS) [rep]; River Pride Choker (13087, -0.52 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 33.0 attack_power points (2.14 DPS) | yes | Mantle of Thieves (2264, -0.84 DPS) [dungeon]; Bristlebark Amice (14573, -0.93 DPS) [world_drop]; Barbaric Shoulders (5964, -1.16 DPS, sim-verified) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.6 attack_power points (1.34 DPS) | yes | Construct Cloak (279848, -0.34 DPS) [quest]; Tigerstrike Mantle (13108, -0.37 DPS, sim-verified) [world_drop]; Swiftrunner Cape (6745, -0.56 DPS) [quest] |
| chest | Brawler's Leather Tunic (252508) | Leatherworking [crafted] | 29.2 attack_power points (1.90 DPS) | yes | Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Defender's Leather Tunic (252450, -0.26 DPS) [crafted]; Brawler's Leather Armor (252490, -0.27 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.8 attack_power points (1.35 DPS) | yes | Barbaric Bracers (18948, -0.39 DPS, sim-verified) [crafted]; Bands of Serra'kis (6902, -0.49 DPS) [dungeon]; Insignia Bracers (6410, -0.57 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 27.4 attack_power points (1.78 DPS) | yes | Toughened Leather Gloves (4253, -0.14 DPS) [crafted]; Wolfclaw Gloves (1978, -0.29 DPS) [dungeon]; Gloves of the Fang (10413, -0.43 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 31.8 attack_power points (2.07 DPS) | yes | Skulker's Leather Belt (252520, -0.04 DPS) [crafted]; Warden's Leather Belt (252460, -0.43 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.51 DPS) [rep] |
| legs | Brawler's Leather Legguards (252516) | Leatherworking [crafted] | 33.4 attack_power points (2.17 DPS) | yes | Brawler's Leather Pants (252500, -0.14 DPS) [crafted]; Trapper's Leather Pants (252501, -0.14 DPS) [crafted]; Barbaric Leggings (5963, -0.26 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.6 attack_power points (1.47 DPS) | yes | Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Vorrel's Boots (7751, -0.17 DPS) [quest]; Stomping Boots (3741, -0.38 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 25.2 attack_power points (1.64 DPS) | yes | Band of the Fist (17694, -0.55 DPS) [quest]; Tiger Band (6749, -0.64 DPS) [quest]; Monkey Ring (6748, -0.73 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 23.6 attack_power points (1.53 DPS) | yes | Tiger Band (6749, -0.53 DPS) [quest]; Band of the Fist (17694, -0.56 DPS, sim-verified) [quest]; Monkey Ring (6748, -0.63 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (119.7 DPS) | yes | Cobalt Crusher (7730, -3.02 DPS) [dungeon]; Advisor's Gnarled Staff (19569, -4.07 DPS) [pvp]; Viscous Hammer (13045, -23.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Defender's Leather Helm; neck: Kaleidoscope Chain; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Brawler's Leather Tunic; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Brawler's Leather Legguards; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 40 (tauren, 0000000000000000-55232032121032000000-0000000000000000)

Set DPS (verified): 162.9. Weights run: 2.5s. Verify run: 3.0s. 525 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.917 ± 0.029, crit=0.922 ± 0.034 per rating point (14 rating = 1%, 12.906 per %), hit=0.919 ± 0.058 per rating point (10 rating = 1%, 9.193 per %), melee_haste=6.120 ± 0.774

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 45.3 attack_power points (3.36 DPS) | yes | Defender's Leather Helm (252455, -1.40 DPS) [crafted]; Warden's Wizard Hat (14604, -1.51 DPS) [world_drop]; Hawkeye's Helm (14591, -1.52 DPS, sim-verified) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 21.1 attack_power points (1.57 DPS) | yes | Ethereal Talisman (4430, -0.18 DPS) [quest]; Kaleidoscope Chain (13084, -0.34 DPS) [world_drop]; River Pride Choker (13087, -0.91 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 33.1 attack_power points (2.46 DPS) | yes | Forest Tracker Epaulets (2278, -0.07 DPS) [world_drop]; Flintrock Shoulders (7755, -0.22 DPS) [dungeon]; Fleshhide Shoulders (10774, -0.79 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 20.0 attack_power points (1.49 DPS) | yes | Imperial Cloak (6432, -0.35 DPS) [dungeon]; Parachute Cloak (10518, -0.35 DPS) [crafted]; Construct Cloak (279848, -0.55 DPS, sim-verified) [quest] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 32.6 attack_power points (2.42 DPS) | yes | Nightscape Tunic (8175, -0.28 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.30 DPS) [crafted]; Barbaric Harness (5739, -0.59 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 20.3 attack_power points (1.51 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.02 DPS) [dungeon]; Barbaric Bracers (18948, -0.28 DPS) [crafted] |
| hands | Prowler's Leather Gloves (252524) | Leatherworking [crafted] | 39.3 attack_power points (2.91 DPS) | yes | Skulker's Leather Gloves (252525, -0.02 DPS) [crafted]; Imperial Leather Gloves (4063, -0.16 DPS) [dungeon]; Gloves of Holy Might (867, -0.47 DPS) [world_drop] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 37.1 attack_power points (2.75 DPS) | yes | Prowler's Leather Belt (252459, -0.43 DPS) [crafted]; Skulker's Leather Belt (252520, -0.49 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.52 DPS) [rep] |
| legs | Triprunner Dungarees (9624) | Rig Wars [quest] | 41.1 attack_power points (3.05 DPS) | yes | Basilisk Hide Pants (1718, -0.06 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.63 DPS) [crafted]; Brawler's Leather Pants (252500, -0.79 DPS) [crafted] |
| feet | Prowler's Leather Shoes (252465) | Leatherworking [crafted] | 37.6 attack_power points (2.79 DPS) | yes | Excelsior Boots (4109, -0.37 DPS) [quest]; Imperial Leather Boots (6431, -0.41 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.52 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 32.9 attack_power points (2.44 DPS) | yes | Ring of the Underwood (2951, -0.70 DPS) [world_drop]; Thunderbrow Ring (13097, -0.71 DPS) [world_drop]; Mark of Kern (2262, -0.96 DPS) [dungeon] |
| finger2 | Falcon's Hook (7552) | Uldaman: Ancient Treasure [dungeon] | 23.9 attack_power points (1.77 DPS) | yes | Ring of the Underwood (2951, -0.02 DPS) [world_drop]; Thunderbrow Ring (13097, -0.04 DPS) [world_drop]; Mark of Kern (2262, -0.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (162.9 DPS) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Thornstone Sledgehammer (1722, +0.00 DPS) [world_drop]; Gut Ripper (2164, -30.56 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Wolffear Harness; hands: Prowler's Leather Gloves; waist: Ogron's Sash; legs: Triprunner Dungarees; feet: Prowler's Leather Shoes; finger1: Legionnaire's Band; finger2: Falcon's Hook

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 50 (tauren, 0100000000000000-55232032121032012001-5000000000000000)

Set DPS (verified): 220.7. Weights run: 2.1s. Verify run: 3.6s. 659 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.864 ± 0.035, crit=1.058 ± 0.046 per rating point (14 rating = 1%, 14.814 per %), hit=1.017 ± 0.074 per rating point (10 rating = 1%, 10.166 per %), melee_haste=5.627 ± 1.110

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 44.7 attack_power points (3.59 DPS) | yes | Undercity Reservist's Cap (20643, -0.18 DPS) [quest]; Blood Guard's Leather Headband (220851, -0.30 DPS) [vendor]; Scorpashi Skullcap (14658, -0.96 DPS) [world_drop] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 35.2 attack_power points (2.83 DPS) | yes | Woven Ivy Necklace (19159, -0.65 DPS, sim-verified) [quest]; Scout's Medallion (19535, -1.03 DPS) [rep] |
| shoulder | Prowler's Leather Shoulder (252534) | Leatherworking [crafted] | 41.0 attack_power points (3.29 DPS) | yes | Skulker's Leather Shoulder (252535, -0.05 DPS) [crafted]; Blood Guard's Leather Shoulders (220853, -0.64 DPS) [vendor]; Failed Flying Experiment (9647, -1.15 DPS, sim-verified) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 28.0 attack_power points (2.25 DPS) | yes | Blackflame Cape (13109, -0.45 DPS) [world_drop]; Serpentskin Cloak (8259, -0.60 DPS) [world_drop]; Pridelord Cape (14673, -0.79 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.8 attack_power points (4.64 DPS) | yes | Blazewind Breastplate (11193, -0.92 DPS, sim-verified) [quest]; Stone Guard's Leather Armor (220855, -1.19 DPS) [vendor]; Charred Leather Tunic (19127, -1.94 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 36.8 attack_power points (2.95 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Prowler's Leather Bracers (252539, -0.31 DPS) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 63.0 attack_power points (5.06 DPS) | yes | Feralheart Fists (226793, -1.17 DPS, sim-verified) [vendor]; Prowler's Leather Gauntlets (252547, -1.24 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -1.37 DPS) [crafted] |
| waist | Prowler's Leather Waistguard (252473) | Leatherworking [crafted] | 53.2 attack_power points (4.27 DPS) | yes | Skulker's Leather Waistguard (252474, -0.05 DPS) [crafted]; Ogron's Sash (13117, -1.33 DPS) [world_drop]; Defiler's Leather Girdle (20193, -1.47 DPS) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 58.4 attack_power points (4.69 DPS) | yes | Stone Guard's Leather Pants (220859, -0.90 DPS, sim-verified) [vendor]; Triprunner Dungarees (9624, -1.47 DPS) [quest]; Basilisk Hide Pants (1718, -1.55 DPS) [world_drop] |
| feet | Prowler's Leather Boots (252468) | Leatherworking [crafted] | 49.1 attack_power points (3.94 DPS) | yes | Skulker's Leather Boots (252469, -0.05 DPS) [crafted]; Albino Crocscale Boots (17728, -0.95 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.95 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 38.8 attack_power points (3.11 DPS) | yes | White Bone Band (11862, -1.19 DPS) [quest]; Falcon's Hook (7552, -1.24 DPS) [dungeon]; Thunderbrow Ring (13097, -1.25 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 26.1 attack_power points (2.10 DPS) | yes | White Bone Band (11862, -0.17 DPS) [quest]; Falcon's Hook (7552, -0.22 DPS) [dungeon]; Thunderbrow Ring (13097, -0.23 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (220.7 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (220.7 DPS) | yes | Molten Heart of the Mountain (249470, -0.92 DPS, sim-verified) [crafted] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (220.7 DPS) | yes | Dark Iron Pulverizer (11608, +0.00 DPS) [crafted]; Thorium Greatmace (250613, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -36.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulder; back: Dark Phantom Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Prowler's Leather Waistguard; legs: Serpentskin Leggings; feet: Prowler's Leather Boots; finger1: Legionnaire's Band; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain

No-known-source sample (15 of 659, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 60 (tauren, 0100000000000000-55232032121032012001-5053200000000000)

Set DPS (verified): 290.2. Weights run: 1.9s. Verify run: 12.0s. 1315 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.200 ± 0.003, agility=1.791 ± 0.041, crit=1.130 ± 0.058 per rating point (14 rating = 1%, 15.823 per %), hit=1.103 ± 0.093 per rating point (10 rating = 1%, 11.030 per %), melee_haste=not significant (6.176 ± 1.572)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 81.9 attack_power points (6.99 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Blue Suede Hat (252482, -6.33 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 56.0 attack_power points (4.78 DPS) | yes | Medallion of the Dawn (22659, -1.59 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -1.79 DPS) [quest]; Skibi's Pendant (13089, -1.86 DPS) [world_drop] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (290.2 DPS) | yes | Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Stormshroud Shoulders (15058, -7.44 DPS, sim-verified) [crafted] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 46.9 attack_power points (4.00 DPS) | yes | Deathguard's Cloak (20068, -0.33 DPS) [rep]; Howler's Furs (272414, -0.67 DPS) [vendor]; Shifting Cloak (18511, -1.40 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 120.7 attack_power points (10.30 DPS) | yes | Dawn Armor (252483, -1.70 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -2.11 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -2.11 DPS) [vendor] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (290.2 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Feralheart Fists (226793) | Mokvar [vendor] | sim-verified (290.2 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Raider Gloves (272099, +0.00 DPS) [vendor]; Stormshroud Gloves (21278, -5.59 DPS, sim-verified) [crafted] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 87.7 attack_power points (7.48 DPS) | yes | Might of the Timbermaw (19044, -1.40 DPS) [crafted]; Shifter's Belt (272396, -1.43 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -1.62 DPS) [quest] |
| legs | Feralheart Trousers (226791) | Mokvar [vendor] | sim-verified (290.2 DPS) | yes | Warbear Woolies (15065, +0.00 DPS) [crafted]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Stormshroud Pants (15057, -6.75 DPS, sim-verified) [crafted] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (290.2 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Drudge Boots (21532, -3.26 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 63.8 attack_power points (5.45 DPS) | yes | Don Julio's Band (19325, -1.79 DPS) [rep]; Cutthroat's Signet (272408, -2.37 DPS) [vendor]; Masons Fraternity Ring (9533, -3.31 DPS) [quest] |
| finger2 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 46.1 attack_power points (3.94 DPS) | yes | Don Julio's Band (19325, -0.28 DPS) [rep]; Cutthroat's Signet (272408, -0.85 DPS) [vendor]; Masons Fraternity Ring (9533, -1.79 DPS) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (290.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (290.2 DPS) | yes | Frozen Heart of the Mountain (249469, -3.65 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (290.2 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Glacial Blade (19099, -6.59 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (290.2 DPS) | yes | Howling Idol (272427, -7.14 DPS, sim-verified) [vendor] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Feralheart Fists; waist: Ferocity of the Timbermaw; legs: Feralheart Trousers; feet: Feralheart Walkers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Legionnaire's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1315, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

### Band 60, raid preset (tauren, 0500200000000000-45211031021032212001-5053000000000000)

Set DPS (verified): 666.8. Weights run: 2.2s. Verify run: 12.1s. 1315 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, feral_attack_power=1.000 ± 0.001, strength=2.420 ± 0.003, agility=2.290 ± 0.049, crit=1.532 ± 0.066 per rating point (14 rating = 1%, 21.450 per %), hit=1.821 ± 0.117 per rating point (10 rating = 1%, 18.209 per %), melee_haste=13.757 ± 1.827

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cap (226792) | Mokvar [vendor] | 99.2 attack_power points (15.23 DPS) | yes | Warlord's Dragonhide Helmet (16550, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -24.55 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 67.7 attack_power points (10.39 DPS) | yes | Medallion of the Dawn (22659, -3.65 DPS, sim-verified) [quest]; Beads of Ogre Might (22150, -3.91 DPS) [quest]; Skibi's Pendant (13089, -3.97 DPS) [world_drop] |
| shoulder | Feralheart Epaulets (226790) | Mokvar [vendor] | sim-verified (666.8 DPS) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -15.78 DPS, sim-verified) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | 54.3 attack_power points (8.34 DPS) | yes | Howler's Furs (272414, -1.25 DPS) [vendor]; Deathguard's Cloak (20068, -1.37 DPS) [rep]; Shifting Cloak (18511, -2.37 DPS) [crafted] |
| chest | Timbermaw Tunic (252484) | Leatherworking [crafted] | 142.5 attack_power points (21.87 DPS) | yes | Dawn Armor (252483, -2.42 DPS, sim-verified) [crafted]; Warlord's Dragonhide Chestpiece (231686, -3.44 DPS) [vendor]; Warlord's Dragonhide Hauberk (16549, -4.38 DPS) [vendor] |
| wrist | Feralheart Bands (226788) | Mokvar [vendor] | sim-verified (666.8 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, +0.00 DPS) [quest] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | 88.4 attack_power points (13.57 DPS) | yes | Timbermaw Brawlers (19049, -0.85 DPS) [crafted]; General's Dragonhide Grips (231688, -1.20 DPS) [vendor]; Blood Guard's Dragonhide Grips (227180, -2.44 DPS) [pvp] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 101.8 attack_power points (15.63 DPS) | yes | Belt of Preserved Heads (20216, -2.36 DPS) [quest]; Shifter's Belt (272396, -2.38 DPS, sim-verified) [vendor]; Might of the Timbermaw (19044, -2.91 DPS) [crafted] |
| legs | Warbear Woolies (15065) | Leatherworking [crafted] | 109.0 attack_power points (16.72 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -0.65 DPS) [vendor]; Legionnaire's Dragonhide Leggings (227177, -1.96 DPS) [pvp] |
| feet | Feralheart Walkers (226794) | Mokvar [vendor] | sim-verified (666.8 DPS) | yes | Blood Guard's Dragonhide Treads (227181, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Drudge Boots (21532, -18.77 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234034) | Anachronos [vendor] | 76.3 attack_power points (11.71 DPS) | yes | Legionnaire's Band (19510, -3.39 DPS) [rep]; Cutthroat's Signet (272408, -3.99 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 55.7 attack_power points (8.54 DPS) | yes | Cutthroat's Signet (272408, -0.83 DPS) [vendor]; Legionnaire's Band (19511, -1.67 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (666.8 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (666.8 DPS) | yes | Frozen Heart of the Mountain (249469, -6.04 DPS, sim-verified) [crafted] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (666.8 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -18.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Dream (220606) | The Temple of Atal'Hakkar: Hazzas [dungeon] | sim-verified (666.8 DPS) | yes | Idol of the Moon (23197, -5.69 DPS, sim-verified) [world_drop] |

**New at 60:** head: Feralheart Cap; neck: Amulet of the Darkmoon; shoulder: Feralheart Epaulets; back: Cape of the Black Baron; chest: Timbermaw Tunic; wrist: Feralheart Bands; hands: Raider Gloves; waist: Ferocity of the Timbermaw; legs: Warbear Woolies; feet: Feralheart Walkers; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: The Unstoppable Force; ranged: Idol of the Dream

No-known-source sample (15 of 1315, see the JSON for more): 1189 Overseer's Ring; 1215 Support Girdle; 1314 Ghoul Fingers; 1355 Buckskin Cape; 1539 Gnarled Hermit's Staff; 1832 Lucky Trousers; 1913 Studded Blackjack; 1917 Jeweled Dagger; 1938 Block Mallet; 1965 White Wolf Gloves; 2058 Kazon's Maul; 2069 Black Bear Hide Vest; 2235 Brackclaw; 2241 Desperado Cape; 2254 Icepane Warhammer

