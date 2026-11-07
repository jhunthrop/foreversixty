# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 55.1. Weights run: 6.9s. Verify run: 10.9s. 268 eligible items had no known source.

Race sweep: night-elf 55.1 ± 0.0, human 54.5 ± 0.0, high-order-skyborne 54.3 ± 0.0, dwarf 53.4 ± 0.0 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.790 ± 0.013, strength=1.210 ± 0.002, crit=1.078 ± 0.015 per rating point (14 rating = 1%, 15.098 per %), hit=1.498 ± 0.037 per rating point (10 rating = 1%, 14.978 per %), melee_haste=4.721 ± 0.821, intellect=0.143 ± 0.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 22.3 attack_power points (1.04 DPS) | yes | Defender's Leather Hood (252447, -0.60 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -0.98 DPS) [crafted]; Shadow Goggles (4373, -1.00 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 16.7 attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest]; Tarnished Locket (279870, -0.95 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 13.9 attack_power points (0.65 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.62 DPS) [crafted]; Magician's Mantle (12998, -0.64 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 30.7 attack_power points (1.43 DPS) | yes | Brawler's Leather Armor (252490, -0.25 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.50 DPS) [crafted]; Prospector's Chestpiece (14562, -0.54 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 13.9 attack_power points (0.65 DPS) | yes | Bravo's Armbands (270015, -0.09 DPS) [quest]; Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 21.6 attack_power points (1.00 DPS) | yes | Serpent Gloves (5970, -0.22 DPS, sim-verified) [dungeon]; Bristlebark Gloves (14572, -0.26 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.26 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Brawler's Leather Belt (252428, -0.09 DPS) [crafted]; Deviate Scale Belt (6468, -0.19 DPS, sim-verified) [crafted]; Trapper's Leather Belt (252429, -0.30 DPS) [crafted] |
| legs | Trapper's Leather Pants (252501) | Leatherworking [crafted] | 32.9 attack_power points (1.53 DPS) | yes | Brawler's Leather Pants (252500, -0.03 DPS) [crafted]; Leggings of the Fang (10410, -0.08 DPS) [dungeon]; Blackened Defias Leggings (10400, -0.42 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 25.9 attack_power points (1.21 DPS) | yes | Brawler's Leather Boots (252439, -0.28 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.43 DPS) [dungeon]; Footpads of the Fang (10411, -0.43 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 16.0 attack_power points (0.74 DPS) | yes | Demon Band (12054, -0.52 DPS) [world_drop]; The 1 Ring (8350, -0.55 DPS) [world]; Ring of the Moon (12052, -0.69 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Demon Band (12054, -0.07 DPS) [world_drop]; The 1 Ring (8350, -0.11 DPS) [world]; Ring of the Moon (12052, -0.24 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.1 attack_power points (11.07 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Assassin's Blade (1935, -5.44 DPS, sim-verified) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Assassin's Blade (1935, -0.07 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.2 attack_power points (0.52 DPS) | yes | Deadly Blunderbuss (4369, -0.25 DPS, sim-verified) [crafted]; Light Bow (4576, -0.26 DPS) [world_drop]; Owlsight Rifle (15205, -0.26 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Butcher's Cleaver; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 268, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 0000000000000000-0000000000000000-500230230050000100)

Set DPS (verified): 104.2. Weights run: 7.2s. Verify run: 22.6s. 463 eligible items had no known source.

Race sweep: night-elf 104.2 ± 0.1, human 103.2 ± 0.1, dwarf 102.9 ± 0.1, high-order-skyborne 102.6 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.377 ± 0.013, strength=1.210 ± 0.002, crit=1.354 ± 0.019 per rating point (14 rating = 1%, 18.957 per %), hit=1.811 ± 0.052 per rating point (10 rating = 1%, 18.108 per %), melee_haste=7.467 ± 1.009, intellect=1.141 ± 0.048

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.8 attack_power points (1.67 DPS) | yes | Brawler's Leather Hood (252504, -0.33 DPS) [crafted]; Enduring Cap (3020, -0.39 DPS) [world_drop]; Dokebi Hat (14584, -0.45 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Sentinel's Medallion (19541, +0.00 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.2 attack_power points (2.26 DPS) | yes | Mantle of Thieves (2264, -0.59 DPS) [dungeon]; Barbaric Shoulders (5964, -1.00 DPS) [crafted]; Bristlebark Amice (14573, -1.00 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.10 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.42 DPS) [dungeon] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 38.5 attack_power points (2.70 DPS) | yes | Dusky Leather Armor (7374, -0.37 DPS) [crafted]; Green Leather Armor (4255, -0.73 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.86 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.1 attack_power points (1.34 DPS) | yes | Barbaric Bracers (18948, -0.33 DPS) [crafted]; Insignia Bracers (6410, -0.34 DPS) [world_drop]; Demonhide Bracers (270033, -0.42 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.7 attack_power points (1.60 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Cartilage Shapers (273054, -0.09 DPS) [world]; Wolfclaw Gloves (1978, -0.17 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.7 attack_power points (2.09 DPS) | yes | Skulker's Leather Belt (252520, -0.07 DPS) [crafted]; Stalker's Leather Belt (252521, -0.10 DPS) [crafted]; Highlander's Chain Girdle (20090, -0.40 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 39.7 attack_power points (2.79 DPS) | yes | Petrolspill Leggings (9509, -0.45 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.45 DPS) [world_drop]; Trapper's Leather Pants (252501, -0.45 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.6 attack_power points (1.59 DPS) | yes | Insignia Boots (4055, -0.25 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.33 DPS) [crafted]; Trapper's Leather Boots (252440, -0.44 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.5 attack_power points (1.51 DPS) | yes | Seal of Wrynn (2933, -0.44 DPS) [quest]; Ring of Precision (1491, -0.51 DPS) [dungeon]; Monkey Ring (6748, -0.68 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Seal of Wrynn (2933, -0.10 DPS) [quest]; Ring of Precision (1491, -0.18 DPS) [dungeon]; Monkey Ring (6748, -0.35 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Combatant Claymore (2877, +0.00 DPS) [world_drop]; Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (22.34 DPS) | yes | Shoni's Disarming Tool (9608, -6.05 DPS) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalpine Stinger (13037, -0.01 DPS) [world_drop]; Double-barreled Shotgun (2098, -0.25 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.33 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 463, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0000000000000000-500230230050222131)

Set DPS (verified): 176.7. Weights run: 7.1s. Verify run: 24.1s. 718 eligible items had no known source.

Race sweep: night-elf 176.7 ± 0.1, human 175.3 ± 0.1, high-order-skyborne 174.9 ± 0.1, dwarf 174.6 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.764 ± 0.016, strength=1.210 ± 0.002, crit=1.994 ± 0.025 per rating point (14 rating = 1%, 27.914 per %), hit=2.402 ± 0.097 per rating point (10 rating = 1%, 24.020 per %), melee_haste=9.440 ± 1.281, intellect=0.044 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.7 attack_power points (3.15 DPS) | yes | Raging Berserker's Helm (7719, -0.00 DPS) [dungeon]; Warden's Wizard Hat (14604, -0.55 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.57 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.4 attack_power points (2.19 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.75 DPS) [quest]; Kaleidoscope Chain (13084, -1.03 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.31 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.5 attack_power points (2.63 DPS) | yes | Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.08 DPS) [world_drop]; Imperial Cloak (6432, -0.14 DPS) [dungeon]; Parachute Cloak (10518, -0.14 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 90.4 attack_power points (6.51 DPS) | yes | Quillward Harness (10583, -2.29 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.12 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -3.50 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.1 attack_power points (1.59 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Branded Leather Bracers (19508, -0.15 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Dragonscale Gauntlets (8347, -3.19 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 35.9 attack_power points (2.59 DPS) | yes | Ogron's Sash (13117, -0.01 DPS) [world_drop]; Scorpashi Sash (14652, -0.25 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.27 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.0 attack_power points (4.18 DPS) | yes | Triprunner Dungarees (9624, -0.34 DPS) [quest]; Petrolspill Leggings (9509, -1.39 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.39 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 38.9 attack_power points (2.80 DPS) | yes | Blackforge Greaves (6423, -0.02 DPS) [dungeon]; Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Worn Running Boots (9398, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 31.8 attack_power points (2.29 DPS) | yes | Falcon's Hook (7552, -0.24 DPS) [dungeon]; Disengagement Ring (276202, -0.70 DPS) [vendor]; Mark of Kern (2262, -0.85 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.1 attack_power points (2.17 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Disengagement Ring (276202, -0.57 DPS) [vendor]; Mark of Kern (2262, -0.72 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Black Menace (6831, -6.99 DPS, sim-verified) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (31.85 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.12 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Bow (17686, -0.20 DPS) [quest]; Golemsight Long Gun (273029, -0.20 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.25 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5000000000000000-0030000000000000-500230230050222151)

Set DPS (verified): 279.8. Weights run: 7.4s. Verify run: 26.8s. 900 eligible items had no known source.

Race sweep: night-elf 279.8 ± 0.2, human 277.8 ± 0.2, high-order-skyborne 277.5 ± 0.2, dwarf 277.4 ± 0.2 (dps); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.819 ± 0.028, strength=1.210 ± 0.002, crit=2.549 ± 0.051 per rating point (14 rating = 1%, 35.682 per %), hit=2.651 ± 0.149 per rating point (10 rating = 1%, 26.512 per %), melee_haste=13.932 ± 1.771, intellect=2.470 ± 0.053

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 82.6 attack_power points (6.94 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Exile (11124, -0.60 DPS) [quest]; Tough Scorpid Helm (8208, -0.72 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.7 attack_power points (3.59 DPS) | yes | Sentinel's Medallion (19539, -0.74 DPS) [rep]; Horizon Choker (13085, -4.02 DPS, sim-verified) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 53.5 attack_power points (4.49 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS) [vendor]; Tough Scorpid Shoulders (8207, -1.07 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.28 DPS) [pvp] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight's Mail Armor (223078, -1.20 DPS) [vendor]; Blazewind Breastplate (11193, -2.16 DPS) [quest]; Knight's Chain Armor (220828, -3.49 DPS, sim-verified) [vendor] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 47.1 attack_power points (3.96 DPS) | yes | Bracers of the Stone Princess (17714, -0.57 DPS) [dungeon]; Stalker's Mail Bracers (252570, -0.58 DPS) [crafted]; Prowler's Leather Bracers (252539, -1.54 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Sergeant Major's Mail Gauntlets (223076, +0.00 DPS) [vendor]; Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Stalker's Mail Gauntlets (252584, -6.04 DPS, sim-verified) [crafted] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 71.9 attack_power points (6.04 DPS) | yes | Stalker's Mail Belt (252588, -0.65 DPS) [crafted]; Highlander's Lizardhide Girdle (20103, -0.97 DPS) [rep]; Highlander's Mail Girdle (20118, -0.97 DPS) [vendor] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 99.9 attack_power points (8.39 DPS) | yes | Imperial Red Pants (8251, -1.72 DPS) [world_drop]; Knight's Mail Legplates (223074, -2.27 DPS, sim-verified) [vendor]; Tough Scorpid Leggings (8206, -2.29 DPS) [crafted] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.4 attack_power points (5.74 DPS) | yes | Stalker's Mail Sabatons (252576, -0.80 DPS) [crafted]; Albino Crocscale Boots (17728, -1.01 DPS) [dungeon]; Rancor Boots (11865, -1.88 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.5 attack_power points (3.91 DPS) | yes | Protector's Band (19516, -0.76 DPS) [rep]; Brainlash (6440, -0.79 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.5 attack_power points (3.32 DPS) | yes | Protector's Band (19516, -0.17 DPS) [rep]; Brainlash (6440, -0.20 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, -3.39 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -3.18 DPS, sim-verified) [quest] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -2.94 DPS, sim-verified) [crafted] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (45.78 DPS) | yes | Claw of Celebras (17738, -5.27 DPS) [dungeon]; Thorium Cestus (250614, -7.31 DPS, sim-verified) [crafted]; Shoni's Disarming Tool (9608, -26.28 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.12 DPS) [world_drop]; Gryphonwing Long Bow (13022, -1.42 DPS) [dungeon]; Skystriker Bow (13020, -2.42 DPS, sim-verified) [world_drop] |

**New at 50:** head: Eye of Theradras; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Sagebrush Girdle; legs: Knight's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 900, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 403.9. Weights run: 5.2s. Verify run: 60.1s. 2028 eligible items had no known source.

Race sweep: night-elf 403.9 ± 0.4, high-order-skyborne 402.8 ± 0.4, human 401.4 ± 0.4, dwarf 400.2 ± 0.4 (dps); the guide race night-elf is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.241 ± 0.016, strength=1.210 ± 0.002, crit=2.017 ± 0.030 per rating point (14 rating = 1%, 28.243 per %), hit=1.807 ± 0.134 per rating point (10 rating = 1%, 18.070 per %), melee_haste=17.923 ± 1.667, intellect=1.316 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 125.3 attack_power points (16.81 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, -0.68 DPS) [pvp]; Outlaw's Collar (279253, -6.06 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Sentinel's Medallion (19538, -0.21 DPS) [rep]; Imperial Jewel (11933, -0.43 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 88.7 attack_power points (11.91 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.01 DPS) [vendor]; Cloak of the Honor Guard (20073, -1.13 DPS) [rep]; Blackveil Cape (11626, -2.01 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 129.2 attack_power points (17.34 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -1.21 DPS) [vendor]; Timbermaw Tunic (252484, -5.52 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 55.9 attack_power points (7.50 DPS) | yes | Marshal's Chain Bracers (16461, -0.25 DPS) [pvp]; Windtalker's Wristguards (19582, -0.99 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -1.58 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 72.2 attack_power points (9.69 DPS) | yes | Marshal's Chain Girdle (16464, -0.29 DPS) [pvp]; Marksman's Girdle (22232, -0.95 DPS) [dungeon]; Highlander's Chain Girdle (20043, -1.34 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 153.0 attack_power points (20.52 DPS) | yes | Marshal's Chain Legplates (231558, -1.75 DPS) [vendor]; Marshal's Chain Legguards (231577, -4.39 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.83 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.16 DPS) [dungeon]; Protector's Band (19514, -2.84 DPS) [rep]; Naglering (11669, -11.28 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -1.38 DPS) [rep]; Blackstone Ring (17713, -1.53 DPS) [dungeon]; Painweaver Band (13098, -3.39 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -6.03 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (85.63 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -4.91 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -7.89 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.81 DPS) [quest]; Blackcrow (12651, -0.88 DPS) [dungeon]; Ancient Bone Bow (18680, -0.90 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 968.5. Weights run: 5.0s. Verify run: 56.5s. 2028 eligible items had no known source.

Race sweep: high-order-skyborne 968.7 ± 1.0, night-elf 968.5 ± 1.0, human 965.5 ± 1.0, dwarf 965.0 ± 0.9 (dps); high-order-skyborne beats the guide race night-elf by 0.2 (0.02%) ± 1.4, inside error.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.214 ± 0.016, strength=1.210 ± 0.002, crit=2.085 ± 0.030 per rating point (14 rating = 1%, 29.188 per %), hit=1.656 ± 0.208 per rating point (10 rating = 1%, 16.563 per %), melee_haste=17.890 ± 2.631, intellect=1.267 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 123.8 attack_power points (35.20 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -1.07 DPS) [vendor]; Outlaw's Collar (279253, -9.42 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Sentinel's Medallion (19538, -0.46 DPS) [rep]; Imperial Jewel (11933, -0.80 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 89.0 attack_power points (25.29 DPS) | yes | Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep]; Field Marshal's Chain Spaulders (16468, -0.62 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -2.31 DPS) [rep]; Howler's Furs (272414, -2.46 DPS) [vendor]; Blackveil Cape (11626, -4.25 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 127.7 attack_power points (36.29 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -2.28 DPS) [vendor]; Timbermaw Tunic (252484, -11.60 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 55.4 attack_power points (15.74 DPS) | yes | Marshal's Chain Bracers (16461, -0.63 DPS) [pvp]; Windtalker's Wristguards (19582, -2.06 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -3.32 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Mourning Handguards (279257, -2.52 DPS) [crafted] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 70.9 attack_power points (20.15 DPS) | yes | Marshal's Chain Girdle (16464, -0.62 DPS) [pvp]; Highlander's Chain Girdle (20043, -2.18 DPS) [rep]; Highlander's Leather Girdle (20045, -2.18 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 152.4 attack_power points (43.33 DPS) | yes | Marshal's Chain Legplates (231558, -3.59 DPS) [vendor]; Marshal's Chain Legguards (231577, -9.31 DPS) [pvp]; Sentinel's Leather Pants (237818, -9.74 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -4.14 DPS) [dungeon]; Protector's Band (19514, -5.93 DPS) [rep]; Naglering (11669, -23.12 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -2.47 DPS) [rep]; Blackstone Ring (17713, -3.13 DPS) [dungeon]; Painweaver Band (13098, -6.92 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -19.12 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (181.42 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -10.94 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -16.71 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -1.65 DPS) [quest]; Ancient Bone Bow (18680, -1.89 DPS) [dungeon]; Blackcrow (12651, -2.21 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Devilsaur Eye; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (orc, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 56.1. Weights run: 6.9s. Verify run: 13.1s. 264 eligible items had no known source.

Race sweep: orc 56.1 ± 0.0, windshaper-skyborne 55.4 ± 0.0, troll 55.3 ± 0.0, tauren 54.4 ± 0.0 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.790 ± 0.013, strength=1.210 ± 0.002, crit=1.078 ± 0.015 per rating point (14 rating = 1%, 15.098 per %), hit=1.498 ± 0.037 per rating point (10 rating = 1%, 14.978 per %), melee_haste=4.721 ± 0.821, intellect=0.143 ± 0.015

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 22.3 attack_power points (1.04 DPS) | yes | Defender's Leather Hood (252447, -0.60 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -0.98 DPS) [crafted]; Shadow Goggles (4373, -1.00 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 16.7 attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 13.9 attack_power points (0.65 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.62 DPS) [crafted]; Magician's Mantle (12998, -0.79 DPS, sim-verified) [world_drop] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.13 DPS) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Bristlebark Cape (14571, -0.26 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.73 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 13.9 attack_power points (0.65 DPS) | yes | Wolf Bracers (4794, -0.13 DPS) [vendor]; Bristlebark Bindings (14569, -0.15 DPS) [world_drop]; Spare Part Bindings (279875, -0.26 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 21.6 attack_power points (1.00 DPS) | yes | Serpent Gloves (5970, -0.23 DPS, sim-verified) [dungeon]; Bristlebark Gloves (14572, -0.26 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.26 DPS) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.84 DPS) | yes | Brawler's Leather Belt (252428, -0.09 DPS) [crafted]; Trapper's Leather Belt (252429, -0.30 DPS) [crafted]; Deviate Scale Belt (6468, -2.22 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Brawler's Leather Pants (252500, +0.00 DPS) [crafted]; Trapper's Leather Pants (252501, -1.65 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.61 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 16.0 attack_power points (0.74 DPS) | yes | Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon]; Demon Band (12054, -0.52 DPS) [world_drop]; Loop of Sacrifice (281673, -0.54 DPS) [quest] |
| finger2 | Bounty Hunter's Ring (5351) | Hezrul Bloodmark [quest] | 8.4 attack_power points (0.39 DPS) | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Demon Band (12054, -0.16 DPS) [world_drop]; Loop of Sacrifice (281673, -0.19 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Wingblade (6504) | Leaders of the Fang [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Forsaken Greataxe (251533, +0.00 DPS) [quest]; Butcher's Cleaver (1292, -0.45 DPS, sim-verified) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 232.2 attack_power points (10.80 DPS) | yes | Cruel Barb (5191, -0.22 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.2 attack_power points (0.52 DPS) | yes | Light Bow (4576, -0.26 DPS) [world_drop]; Privateer Musket (5309, -0.26 DPS) [quest]; Deadly Blunderbuss (4369, -0.27 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Bounty Hunter's Ring; main_hand: Wingblade; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 264, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (orc, 0000000000000000-0000000000000000-500230230050000100)

Set DPS (verified): 103.6. Weights run: 7.2s. Verify run: 20.8s. 462 eligible items had no known source.

Race sweep: orc 103.6 ± 0.1, windshaper-skyborne 102.6 ± 0.1, troll 102.2 ± 0.1, tauren 100.7 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.685 ± 0.007, strength=1.210 ± 0.001, crit=0.385 ± 0.010 per rating point (14 rating = 1%, 5.395 per %), hit=1.864 ± 0.031 per rating point (10 rating = 1%, 18.642 per %), melee_haste=9.305 ± 0.860, intellect=0.324 ± 0.040

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.9 attack_power points (1.08 DPS) | yes | Defender's Leather Helm (252455, -0.15 DPS) [crafted]; Brawler's Leather Hood (252504, -0.22 DPS) [crafted]; Humbert's Helm (4724, -0.32 DPS) [world] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scout's Medallion (19537, -0.03 DPS) [rep]; Kaleidoscope Chain (13084, -0.07 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.6 attack_power points (1.58 DPS) | yes | Mantle of Thieves (2264, -0.50 DPS) [dungeon]; Barbaric Shoulders (5964, -0.65 DPS) [crafted]; Bristlebark Amice (14573, -0.70 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.12 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.34 DPS) [dungeon]; Swiftrunner Cape (6745, -0.34 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dusky Leather Armor (7374, +0.00 DPS) [crafted]; Brawler's Leather Tunic (252508, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.14 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.0 attack_power points (0.96 DPS) | yes | Barbaric Bracers (18948, -0.22 DPS) [crafted]; Insignia Bracers (6410, -0.31 DPS) [world_drop]; Cultist's Armguards (270032, -0.32 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 18.6 attack_power points (1.19 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.17 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Chain Girdle (20152, +0.00 DPS) [rep]; Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 29.9 attack_power points (1.92 DPS) | yes | Trapper's Leather Pants (252501, -0.40 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.40 DPS) [crafted]; Petrolspill Leggings (9509, -0.40 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.1 attack_power points (1.10 DPS) | yes | Brawler's Leather Boots (252439, -0.17 DPS) [crafted]; Insignia Boots (4055, -0.23 DPS) [world_drop]; Warsong Boots (16977, -0.23 DPS) [quest] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.4 attack_power points (1.11 DPS) | yes | Band of the Fist (17694, -0.37 DPS) [quest]; Ring of Precision (1491, -0.47 DPS) [dungeon]; Insurgent's Band (272067, -0.54 DPS) [vendor] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 14.7 attack_power points (0.95 DPS) | yes | Band of the Fist (17694, -0.20 DPS) [quest]; Ring of Precision (1491, -0.30 DPS) [dungeon]; Insurgent's Band (272067, -0.37 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Combatant Claymore (2877, +0.00 DPS) [world_drop]; Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (20.42 DPS) | yes | - |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.1 attack_power points (0.65 DPS) | yes | Double-barreled Shotgun (2098, -0.09 DPS) [world_drop]; Crystalpine Stinger (13037, -0.13 DPS) [world_drop]; Gravespike Repeater (273033, -0.14 DPS) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (orc, 0000000000000000-0000000000000000-500230230050222131)

Set DPS (verified): 174.8. Weights run: 7.1s. Verify run: 23.2s. 708 eligible items had no known source.

Race sweep: orc 174.8 ± 0.1, windshaper-skyborne 174.3 ± 0.1, troll 174.1 ± 0.1, tauren 171.1 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.764 ± 0.016, strength=1.210 ± 0.002, crit=1.994 ± 0.025 per rating point (14 rating = 1%, 27.914 per %), hit=2.402 ± 0.097 per rating point (10 rating = 1%, 24.020 per %), melee_haste=9.440 ± 1.281, intellect=0.044 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 43.7 attack_power points (3.15 DPS) | yes | Raging Berserker's Helm (7719, -0.00 DPS) [dungeon]; Warden's Wizard Hat (14604, -0.55 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.57 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.4 attack_power points (2.19 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.75 DPS) [quest]; Ethereal Talisman (4430, -0.96 DPS) [quest]; Kaleidoscope Chain (13084, -1.03 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.5 attack_power points (2.63 DPS) | yes | Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.06 DPS) [dungeon]; Parachute Cloak (10518, -0.06 DPS) [crafted]; Tigerstrike Mantle (13108, -0.06 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 90.4 attack_power points (6.51 DPS) | yes | Quillward Harness (10583, -2.59 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.12 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -3.50 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.1 attack_power points (1.59 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Branded Leather Bracers (19508, -0.15 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Dragonscale Gauntlets (8347, -2.84 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 35.9 attack_power points (2.59 DPS) | yes | Ogron's Sash (13117, -0.01 DPS) [world_drop]; Scorpashi Sash (14652, -0.25 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.27 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.0 attack_power points (4.18 DPS) | yes | Triprunner Dungarees (9624, -0.34 DPS) [quest]; Petrolspill Leggings (9509, -1.39 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.39 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 38.9 attack_power points (2.80 DPS) | yes | Blackforge Greaves (6423, -0.02 DPS) [dungeon]; Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Worn Running Boots (9398, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 31.8 attack_power points (2.29 DPS) | yes | Falcon's Hook (7552, -0.24 DPS) [dungeon]; Disengagement Ring (276202, -0.70 DPS) [vendor]; Mark of Kern (2262, -0.85 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.1 attack_power points (2.17 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Disengagement Ring (276202, -0.57 DPS) [vendor]; Mark of Kern (2262, -0.72 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.77 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (31.85 DPS) | yes | Sword of Omen (6802, +0.00 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Bow (17686, -0.20 DPS) [quest]; Golemsight Long Gun (273029, -0.20 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.25 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 708, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (orc, 5000000000000000-0030000000000000-500230230050222151)

Set DPS (verified): 280.3. Weights run: 7.4s. Verify run: 28.4s. 891 eligible items had no known source.

Race sweep: orc 280.3 ± 0.2, windshaper-skyborne 278.6 ± 0.2, troll 278.1 ± 0.2, tauren 275.1 ± 0.2 (dps); the guide race orc is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.819 ± 0.028, strength=1.210 ± 0.002, crit=2.549 ± 0.051 per rating point (14 rating = 1%, 35.682 per %), hit=2.651 ± 0.149 per rating point (10 rating = 1%, 26.512 per %), melee_haste=13.932 ± 1.771, intellect=2.470 ± 0.053

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 82.6 attack_power points (6.94 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Exile (11124, -0.60 DPS) [quest]; Tough Scorpid Helm (8208, -0.72 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 42.7 attack_power points (3.59 DPS) | yes | Scout's Medallion (19535, -0.74 DPS) [rep]; Woven Ivy Necklace (19159, -0.85 DPS) [quest]; Horizon Choker (13085, -4.16 DPS, sim-verified) [world_drop] |
| shoulder | Prowler's Leather Shoulders (252534) | Leatherworking [crafted] | 53.5 attack_power points (4.49 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Tough Scorpid Shoulders (8207, -1.11 DPS, sim-verified) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Darkspear Raider's Cloak (272076, -0.65 DPS) [vendor] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Stone Guard's Mail Armor (220826, -1.20 DPS) [vendor]; Stone Guard's Chain Armor (220827, -2.08 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.16 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 47.1 attack_power points (3.96 DPS) | yes | Bracers of the Stone Princess (17714, -0.57 DPS) [dungeon]; Stalker's Mail Bracers (252570, -0.58 DPS) [crafted]; Prowler's Leather Bracers (252539, -1.54 DPS, sim-verified) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Stalker's Mail Gauntlets (252584, +0.00 DPS) [crafted]; First Sergeant's Mail Gauntlets (220831, -5.53 DPS, sim-verified) [vendor] |
| waist | Sagebrush Girdle (17778) | Vyletongue Corruption [quest] | 71.9 attack_power points (6.04 DPS) | yes | Stalker's Mail Belt (252588, -0.65 DPS) [crafted]; Defiler's Lizardhide Girdle (20174, -0.97 DPS) [rep]; Defiler's Mail Girdle (20196, -0.97 DPS) [rep] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 99.9 attack_power points (8.39 DPS) | yes | Imperial Red Pants (8251, -1.72 DPS) [world_drop]; Tough Scorpid Leggings (8206, -2.29 DPS) [crafted]; Stone Guard's Mail Legplates (220834, -2.33 DPS, sim-verified) [vendor] |
| feet | Fleetfoot Greaves (11627) | Blackrock Depths: Houndmaster Grebmar [dungeon] | 68.4 attack_power points (5.74 DPS) | yes | Stalker's Mail Sabatons (252576, -0.80 DPS) [crafted]; Albino Crocscale Boots (17728, -1.01 DPS) [dungeon]; Rancor Boots (11865, -1.93 DPS, sim-verified) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 46.5 attack_power points (3.91 DPS) | yes | Legionnaire's Band (19511, -0.76 DPS) [rep]; Brainlash (6440, -0.79 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 39.5 attack_power points (3.32 DPS) | yes | Legionnaire's Band (19511, -0.17 DPS) [rep]; Brainlash (6440, -0.20 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.56 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Dawn's Edge (12774, -3.12 DPS, sim-verified) [crafted] |
| off_hand | Shadowblade (2163) | World drop [world_drop] | 545.0 attack_power points (45.78 DPS) | yes | Claw of Celebras (17738, -5.27 DPS) [dungeon]; White Bone Shredder (11863, -7.18 DPS) [quest]; Thorium Cestus (250614, -7.58 DPS, sim-verified) [crafted] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.12 DPS) [world_drop]; Gryphonwing Long Bow (13022, -1.42 DPS) [dungeon]; Skystriker Bow (13020, -2.46 DPS, sim-verified) [world_drop] |

**New at 50:** head: Eye of Theradras; neck: Skibi's Pendant; shoulder: Prowler's Leather Shoulders; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Sagebrush Girdle; legs: Stone Guard's Chain Legplates; feet: Fleetfoot Greaves; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Bloodrazor; off_hand: Shadowblade; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (orc, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 395.4. Weights run: 5.2s. Verify run: 60.1s. 2027 eligible items had no known source.

Race sweep: windshaper-skyborne 397.8 ± 0.4, troll 397.0 ± 0.4, orc 395.4 ± 0.4, tauren 391.0 ± 0.4 (dps); windshaper-skyborne beats the guide race orc by 2.4 (0.60%) ± 0.6, beyond error.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.241 ± 0.016, strength=1.210 ± 0.002, crit=2.017 ± 0.030 per rating point (14 rating = 1%, 28.243 per %), hit=1.807 ± 0.134 per rating point (10 rating = 1%, 18.070 per %), melee_haste=17.923 ± 1.667, intellect=1.316 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 125.3 attack_power points (16.81 DPS) | yes | Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Warlord's Chain Helmet (16566, -0.68 DPS) [vendor]; Outlaw's Collar (279253, -6.35 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Scout's Medallion (19534, -0.21 DPS) [rep]; Imperial Jewel (11933, -0.43 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 88.7 attack_power points (11.91 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.01 DPS) [vendor]; Deathguard's Cloak (20068, -1.13 DPS) [rep]; Blackveil Cape (11626, -2.01 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 129.2 attack_power points (17.34 DPS) | yes | Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -1.21 DPS) [vendor]; Timbermaw Tunic (252484, -5.81 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 55.9 attack_power points (7.50 DPS) | yes | General's Chain Wristguards (16570, -0.25 DPS) [pvp]; Windtalker's Wristguards (19582, -0.99 DPS) [rep]; Legionnaire's Chain Bracers (16517, -1.58 DPS) [pvp] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 72.2 attack_power points (9.69 DPS) | yes | General's Chain Girdle (16572, -0.29 DPS) [pvp]; Marksman's Girdle (22232, -0.95 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.34 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 124.7 attack_power points (16.73 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, -0.61 DPS) [pvp]; Outrider's Leather Pants (22740, -7.65 DPS, sim-verified) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.16 DPS) [dungeon]; Legionnaire's Band (19510, -2.84 DPS) [rep]; Naglering (11669, -11.68 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -1.38 DPS) [rep]; Blackstone Ring (17713, -1.53 DPS) [dungeon]; Painweaver Band (13098, -3.47 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -0.89 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.39 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -5.90 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (85.63 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -5.83 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -7.89 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.81 DPS) [quest]; Blackcrow (12651, -0.88 DPS) [dungeon]; Ancient Bone Bow (18680, -0.90 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (orc, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 966.9. Weights run: 5.0s. Verify run: 61.0s. 2027 eligible items had no known source.

Race sweep: troll 970.0 ± 0.9, windshaper-skyborne 967.5 ± 1.0, orc 966.9 ± 0.9, tauren 953.8 ± 0.9 (dps); troll beats the guide race orc by 3.1 (0.32%) ± 1.3, beyond error.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.735 ± 0.027, strength=1.210 ± 0.002, crit=3.383 ± 0.052 per rating point (14 rating = 1%, 47.365 per %), hit=5.122 ± 0.387 per rating point (10 rating = 1%, 51.218 per %), melee_haste=23.632 ± 4.437, intellect=1.282 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 153.5 attack_power points (33.60 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -16.68 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Star of Mystaria (12103, +0.00 DPS) [dungeon]; Mark of Fordring (15411, +0.00 DPS) [quest]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -5.18 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanoweave Cloak (272411, -3.88 DPS) [vendor]; Cape of the Black Baron (13340, -3.98 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -4.31 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 199.8 attack_power points (43.72 DPS) | yes | Warlord's Chain Armor (231566, -7.17 DPS) [vendor]; Savage Gladiator Chain (11726, -7.54 DPS, sim-verified) [dungeon]; Warlord's Chain Chestpiece (16565, -9.64 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Wristguards (16570, -0.35 DPS) [pvp]; Blackmist Armguards (12966, -1.75 DPS) [dungeon]; Slashclaw Bracers (13211, -5.33 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+6.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -6.73 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.48 DPS) [rep]; Defiler's Leather Girdle (20190, -0.48 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 194.3 attack_power points (42.52 DPS) | yes | General's Chain Legplates (231567, -5.97 DPS) [vendor]; General's Chain Legguards (231574, -8.44 DPS) [pvp]; Plaguehound Leggings (18736, -11.63 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -9.05 DPS) [dungeon]; Channeler's Ring (272406, -9.50 DPS) [vendor]; Naglering (11669, -15.63 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -4.00 DPS) [dungeon]; Channeler's Ring (272406, -4.45 DPS) [vendor]; Painweaver Band (13098, -5.72 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.82 DPS) [crafted]; Hand of Justice (11815, -6.53 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -20.85 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (139.69 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -12.87 DPS) [dungeon]; Heartseeker (12783, -25.96 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -4.63 DPS) [world_drop]; Ancient Bone Bow (18680, -6.42 DPS) [dungeon]; Bloodseeker (19107, -6.70 DPS) [quest] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker; ranged: Blackcrow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

