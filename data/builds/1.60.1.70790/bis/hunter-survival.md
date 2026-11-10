# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 55.1. Weights run: 5.6s. Verify run: 9.7s. 298 eligible items had no known source.

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
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 16.0 attack_power points (0.74 DPS) | yes | Demon Band (12054, -0.52 DPS) [world_drop]; Ring of the Moon (12052, -0.69 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.72 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 6.4 attack_power points (0.30 DPS) | yes | Demon Band (12054, -0.07 DPS) [world_drop]; Ring of the Moon (12052, -0.24 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Butcher's Cleaver (1292) | Shadowfang Keep: Razorclaw the Butcher [dungeon] | 238.1 attack_power points (11.07 DPS) | yes | Duskbringer (2205, +0.00 DPS) [dungeon]; Impaling Harpoon (5200, +0.00 DPS) [dungeon]; Assassin's Blade (1935, -5.44 DPS, sim-verified) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Assassin's Blade (1935, -0.07 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.2 attack_power points (0.52 DPS) | yes | Deadly Blunderbuss (4369, -0.25 DPS, sim-verified) [crafted]; Light Bow (4576, -0.26 DPS) [world_drop]; Owlsight Rifle (15205, -0.26 DPS) [quest] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Trapper's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Lavishly Jeweled Ring; main_hand: Butcher's Cleaver; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 30 (night-elf, 0000000000000000-0000000000000000-500230230050000100)

Set DPS (verified): 104.2. Weights run: 5.8s. Verify run: 19.4s. 494 eligible items had no known source.

Race sweep: night-elf 104.2 ± 0.1, human 103.2 ± 0.1, dwarf 102.9 ± 0.1, high-order-skyborne 102.6 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.762 ± 0.010, strength=1.210 ± 0.001, crit=0.495 ± 0.013 per rating point (14 rating = 1%, 6.932 per %), hit=1.761 ± 0.037 per rating point (10 rating = 1%, 17.611 per %), melee_haste=7.151 ± 0.889, intellect=0.922 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.6 attack_power points (1.18 DPS) | yes | Enduring Cap (3020, -0.19 DPS) [world_drop]; Defender's Leather Helm (252455, -0.21 DPS) [crafted]; Dokebi Hat (14584, -1.46 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 17.5 attack_power points (1.17 DPS) | yes | Kaleidoscope Chain (13084, -0.17 DPS, sim-verified) [world_drop]; Sentinel's Medallion (19541, -0.23 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 25.4 attack_power points (1.70 DPS) | yes | Mantle of Thieves (2264, -0.48 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.71 DPS) [crafted]; Bristlebark Amice (14573, -0.75 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.13 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.36 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.59 DPS, sim-verified) [pvp] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 29.7 attack_power points (1.99 DPS) | yes | Green Leather Armor (4255, -0.55 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.56 DPS) [crafted]; Dusky Leather Armor (7374, -1.09 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.4 attack_power points (1.03 DPS) | yes | Glowing Magical Bracelets (13106, -0.30 DPS) [world_drop]; Demonhide Bracers (270033, -0.32 DPS) [quest]; Barbaric Bracers (18948, -0.34 DPS, sim-verified) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.0 attack_power points (1.28 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.20 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 25.1 attack_power points (1.69 DPS) | yes | Highlander's Chain Girdle (20090, -0.08 DPS) [rep]; Highlander's Leather Girdle (20117, -0.08 DPS) [rep]; Skulker's Leather Belt (252520, -0.14 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 37.1 attack_power points (2.48 DPS) | yes | Trapper's Leather Pants (252501, -0.34 DPS, sim-verified) [crafted]; Petrolspill Leggings (9509, -0.83 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.83 DPS) [world_drop] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.7 attack_power points (1.19 DPS) | yes | Insignia Boots (4055, -0.24 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.34 DPS, sim-verified) [crafted]; Trapper's Leather Boots (252440, -0.35 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 17.8 attack_power points (1.20 DPS) | yes | Seal of Wrynn (2933, -0.35 DPS) [quest]; Ring of Precision (1491, -0.49 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Ring of Precision (1491, -0.29 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; Seal of Wrynn (2933, -0.47 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Toxic Revenger (9453, -4.12 DPS, sim-verified) [dungeon] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (21.33 DPS) | yes | Zealot Blade (13033, -0.91 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -5.77 DPS) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Gravespike Repeater (273033, -0.17 DPS) [world]; Crystalpine Stinger (13037, -0.30 DPS, sim-verified) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 494, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0000000000000000-500230230050222131)

Set DPS (verified): 176.8. Weights run: 5.7s. Verify run: 18.1s. 827 eligible items had no known source.

Race sweep: night-elf 176.8 ± 0.1, human 175.4 ± 0.1, high-order-skyborne 175.0 ± 0.1, dwarf 174.7 ± 0.1 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.771 ± 0.016, strength=1.210 ± 0.002, crit=2.007 ± 0.025 per rating point (14 rating = 1%, 28.098 per %), hit=2.420 ± 0.098 per rating point (10 rating = 1%, 24.199 per %), melee_haste=9.528 ± 1.293, intellect=0.044 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 43.8 attack_power points (3.16 DPS) | yes | White Bandit Mask (10008, -0.00 DPS) [crafted]; Warden's Wizard Hat (14604, -0.55 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.58 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 30.5 attack_power points (2.20 DPS) | yes | Kaleidoscope Chain (13084, -1.04 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.19 DPS) [quest]; Archmagister's Faceted Pendant (273049, -1.31 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.5 attack_power points (2.63 DPS) | yes | Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.08 DPS) [world_drop]; Imperial Cloak (6432, -0.14 DPS) [dungeon]; Parachute Cloak (10518, -0.14 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 90.7 attack_power points (6.54 DPS) | yes | Quillward Harness (10583, -2.31 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.14 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -3.52 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.2 attack_power points (1.60 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Dragonscale Gauntlets (8347, -3.22 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 36.1 attack_power points (2.60 DPS) | yes | Ogron's Sash (13117, -0.02 DPS) [world_drop]; Scorpashi Sash (14652, -0.26 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.2 attack_power points (4.19 DPS) | yes | Triprunner Dungarees (9624, -0.34 DPS) [quest]; Petrolspill Leggings (9509, -1.40 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.40 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 39.0 attack_power points (2.81 DPS) | yes | Blackforge Greaves (6423, -0.03 DPS) [dungeon]; Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Worn Running Boots (9398, -0.21 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 31.9 attack_power points (2.30 DPS) | yes | Falcon's Hook (7552, -0.24 DPS) [dungeon]; Disengagement Ring (276202, -0.70 DPS) [vendor]; Mark of Kern (2262, -0.85 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.1 attack_power points (2.17 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Disengagement Ring (276202, -0.57 DPS) [vendor]; Mark of Kern (2262, -0.73 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Black Menace (6831, -6.98 DPS, sim-verified) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (31.85 DPS) | yes | Digmaster 5000 (9465, +0.00 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.12 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Bow (17686, -0.20 DPS) [quest]; Golemsight Long Gun (273029, -0.20 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.25 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 827, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5000000000000000-0030000000000000-500230230050222151)

Set DPS (verified): 283.9. Weights run: 6.0s. Verify run: 26.6s. 1062 eligible items had no known source.

Race sweep: night-elf 283.9 ± 0.2, human 282.0 ± 0.2, high-order-skyborne 281.1 ± 0.2, dwarf 280.0 ± 0.2 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.630 ± 0.014, strength=1.210 ± 0.002, crit=2.106 ± 0.025 per rating point (14 rating = 1%, 29.484 per %), hit=3.181 ± 0.142 per rating point (10 rating = 1%, 31.814 per %), melee_haste=13.418 ± 1.654, intellect=0.048 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 45.2 attack_power points (4.26 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.05 DPS) [crafted]; Knight-Lieutenant's Mail Helmet (223075, -0.20 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.2 attack_power points (3.79 DPS) | yes | Sentinel's Medallion (19539, -0.96 DPS, sim-verified) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 39.8 attack_power points (3.75 DPS) | yes | Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.03 DPS) [dungeon] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.44 DPS) [dungeon]; Dark Phantom Cape (13122, -0.44 DPS) [world_drop]; Blackmetal Cape (9512, -1.18 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.9 attack_power points (8.93 DPS) | yes | Knight's Chain Armor (220828, -2.40 DPS) [vendor]; Blazewind Breastplate (11193, -2.90 DPS) [quest]; Warbear Harness (15064, -3.23 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 44.3 attack_power points (4.17 DPS) | yes | Skulker's Leather Bracers (252540, -1.14 DPS) [crafted]; Pridelord Bands (14672, -1.25 DPS, sim-verified) [world_drop]; Prowler's Leather Bracers (252539, -1.39 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -4.30 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 51.3 attack_power points (4.83 DPS) | yes | Highlander's Chain Girdle (20088, -0.17 DPS) [rep]; Highlander's Leather Girdle (20115, -0.17 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.27 DPS) [crafted] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 84.8 attack_power points (7.98 DPS) | yes | Knight's Chain Legplates (220832, -1.70 DPS) [vendor]; Stormshroud Pants (15057, -2.43 DPS) [crafted]; Serpentskin Leggings (8262, -2.54 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 52.6 attack_power points (4.95 DPS) | yes | Fleetfoot Greaves (11627, -0.22 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.48 DPS) [crafted]; Elven Chain Boots (13125, -0.50 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 51.8 attack_power points (4.88 DPS) | yes | Protector's Band (19516, -1.51 DPS) [rep]; Ring of the Underwood (2951, -2.17 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 36.8 attack_power points (3.47 DPS) | yes | Protector's Band (19516, -0.10 DPS) [rep]; Ring of the Underwood (2951, -0.76 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Molten Heart of the Mountain (249470, -4.48 DPS, sim-verified) [crafted] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Devilsaur Tooth (19992, -0.96 DPS) [quest] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Shadowblade (2163, -8.39 DPS, sim-verified) [world_drop] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dawn's Edge (12774, -1.38 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -4.23 DPS) [dungeon]; Shoni's Disarming Tool (9608, -27.77 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.05 DPS, sim-verified) [world_drop]; Gryphonwing Long Bow (13022, -1.14 DPS) [dungeon]; Houndmaster's Rifle (11629, -1.25 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Infernal Trickster Leggings; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Doomforged Straightedge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1062, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 418.2. Weights run: 4.5s. Verify run: 55.8s. 2287 eligible items had no known source.

Race sweep: night-elf 418.2 ± 0.4, high-order-skyborne 417.7 ± 0.4, human 415.5 ± 0.4, dwarf 414.4 ± 0.4 (dps); the guide race night-elf is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.190 ± 0.016, strength=1.210 ± 0.002, crit=1.956 ± 0.032 per rating point (14 rating = 1%, 27.387 per %), hit=1.641 ± 0.142 per rating point (10 rating = 1%, 16.409 per %), melee_haste=16.836 ± 1.800, intellect=1.321 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 123.6 attack_power points (16.74 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, -0.79 DPS) [pvp]; Eye of Rend (12587, -7.97 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 53.4 attack_power points (7.23 DPS) | yes | Sentinel's Medallion (19538, -2.78 DPS) [rep]; Skibi's Pendant (13089, -2.88 DPS, sim-verified) [world_drop]; Imperial Jewel (11933, -2.90 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 86.5 attack_power points (11.72 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.07 DPS) [rep]; Howler's Furs (272414, -1.14 DPS) [vendor]; Blackveil Cape (11626, -2.02 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.8 attack_power points (16.91 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -0.97 DPS) [vendor]; Timbermaw Tunic (252484, -4.07 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 54.9 attack_power points (7.44 DPS) | yes | Marshal's Chain Bracers (16461, -0.25 DPS) [pvp]; Windtalker's Wristguards (19582, -0.86 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -1.58 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, -1.73 DPS, sim-verified) [rep] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 71.1 attack_power points (9.63 DPS) | yes | Marshal's Chain Girdle (16464, -0.25 DPS) [pvp]; Marksman's Girdle (22232, -1.17 DPS) [dungeon]; Highlander's Chain Girdle (20043, -1.31 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 147.8 attack_power points (20.03 DPS) | yes | Marshal's Chain Legplates (231558, -1.30 DPS) [vendor]; Marshal's Chain Legguards (231577, -4.08 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.60 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 65.5 attack_power points (8.87 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.95 DPS) [dungeon]; Protector's Band (19514, -2.60 DPS) [rep]; Naglering (11669, -9.74 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -1.15 DPS) [rep]; Blackstone Ring (17713, -1.44 DPS) [dungeon]; Painweaver Band (13098, -2.00 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, -0.09 DPS) [quest]; Hand of Justice (11815, -3.23 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Devilsaur Tooth (19992, +0.00 DPS) [quest]; Hand of Justice (11815, -1.52 DPS, sim-verified) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -8.82 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (86.46 DPS) | yes | Alcor's Sunrazor (14555, +0.00 DPS) [world_drop]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -7.96 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.54 DPS) [quest]; Ancient Bone Bow (18680, -0.67 DPS) [dungeon]; Riphook (12653, -0.95 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2287, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 996.5. Weights run: 4.2s. Verify run: 54.8s. 2287 eligible items had no known source.

Race sweep: high-order-skyborne 998.3 ± 0.9, night-elf 996.5 ± 1.0, human 994.9 ± 1.0, dwarf 993.4 ± 1.0 (dps); high-order-skyborne beats the guide race night-elf by 1.8 (0.18%) ± 1.3, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.214 ± 0.017, strength=1.210 ± 0.002, crit=2.111 ± 0.033 per rating point (14 rating = 1%, 29.549 per %), hit=2.276 ± 0.222 per rating point (10 rating = 1%, 22.755 per %), melee_haste=33.403 ± 3.047, intellect=1.269 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 124.3 attack_power points (35.94 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -1.01 DPS) [vendor]; Mask of the Unforgiven (13404, -7.15 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 55.5 attack_power points (16.06 DPS) | yes | Star of Mystaria (12103, -4.94 DPS) [dungeon]; Skibi's Pendant (13089, -5.99 DPS) [world_drop]; Sentinel's Medallion (19538, -6.46 DPS) [rep] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 89.3 attack_power points (25.83 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Cloak of the Honor Guard (20073, -1.64 DPS) [rep]; Cape of the Black Baron (13340, -2.71 DPS, sim-verified) [dungeon]; Blackveil Cape (11626, -3.61 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.2 attack_power points (38.82 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -4.10 DPS) [vendor]; Timbermaw Tunic (252484, -13.31 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 55.4 attack_power points (16.01 DPS) | yes | Marshal's Chain Bracers (16461, -0.64 DPS) [pvp]; Windtalker's Wristguards (19582, -2.09 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -3.38 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Theramore Gloves (271948, -2.15 DPS, sim-verified) [rep] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (+2.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Girdle (16464, -0.15 DPS) [pvp]; Sash of the Grand Hunt (22207, -1.64 DPS) [dungeon]; Ranger's Belt (272397, -2.85 DPS, sim-verified) [vendor] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 159.4 attack_power points (46.08 DPS) | yes | Marshal's Chain Legplates (231558, -5.54 DPS) [vendor]; Marshal's Chain Legguards (231577, -11.36 DPS) [pvp]; Sentinel's Leather Pants (237818, -11.70 DPS) [vendor] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 72.1 attack_power points (20.85 DPS) | yes | Bloodmail Boots (14616, +0.00 DPS) [dungeon]; Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -6.00 DPS) [dungeon]; Blackstone Ring (17713, -6.81 DPS) [dungeon]; Naglering (11669, -24.79 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.18 DPS) [dungeon]; Channeler's Ring (272406, -3.83 DPS) [vendor]; Painweaver Band (13098, -9.01 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -11.35 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Annihilator (12798, -49.48 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (184.56 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -10.84 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -17.00 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.56 DPS) [quest]; Ancient Bone Bow (18680, -0.80 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -2.35 DPS, sim-verified) [world_drop] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Marksman's Girdle; legs: Sentinel's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker; ranged: Blackcrow

No-known-source sample (15 of 2287, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (orc, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 56.2. Weights run: 5.6s. Verify run: 11.7s. 294 eligible items had no known source.

Race sweep: orc 56.2 ± 0.0, windshaper-skyborne 55.6 ± 0.0, troll 55.5 ± 0.0, tauren 54.5 ± 0.0 (dps); the guide race orc is best.

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
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Assassin's Blade (1935, -0.13 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 11.2 attack_power points (0.52 DPS) | yes | Light Bow (4576, -0.26 DPS) [world_drop]; Privateer Musket (5309, -0.26 DPS) [quest]; Deadly Blunderbuss (4369, -0.27 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Bounty Hunter's Ring; main_hand: Wingblade; off_hand: Edward's Knife; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (orc, 0000000000000000-0000000000000000-500230230050000100)

Set DPS (verified): 103.6. Weights run: 5.8s. Verify run: 17.8s. 493 eligible items had no known source.

Race sweep: orc 103.6 ± 0.1, windshaper-skyborne 102.6 ± 0.1, troll 102.2 ± 0.1, tauren 100.7 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=1.762 ± 0.010, strength=1.210 ± 0.001, crit=0.495 ± 0.013 per rating point (14 rating = 1%, 6.932 per %), hit=1.761 ± 0.037 per rating point (10 rating = 1%, 17.611 per %), melee_haste=7.151 ± 0.889, intellect=0.922 ± 0.044

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 17.6 attack_power points (1.18 DPS) | yes | Enduring Cap (3020, -0.19 DPS) [world_drop]; Defender's Leather Helm (252455, -0.21 DPS) [crafted]; Dokebi Hat (14584, -1.84 DPS, sim-verified) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 17.5 attack_power points (1.17 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Scout's Medallion (19537, -0.23 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 25.4 attack_power points (1.70 DPS) | yes | Barbaric Shoulders (5964, -0.71 DPS) [crafted]; Bristlebark Amice (14573, -0.75 DPS) [world_drop]; Mantle of Thieves (2264, -0.82 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.13 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.36 DPS) [dungeon]; Swiftrunner Cape (6745, -0.36 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 29.7 attack_power points (1.99 DPS) | yes | Green Leather Armor (4255, -0.55 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.56 DPS) [crafted]; Dusky Leather Armor (7374, -0.76 DPS, sim-verified) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 15.4 attack_power points (1.03 DPS) | yes | Glowing Magical Bracelets (13106, -0.30 DPS) [world_drop]; Insignia Bracers (6410, -0.32 DPS) [world_drop]; Barbaric Bracers (18948, -0.33 DPS, sim-verified) [crafted] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 19.0 attack_power points (1.28 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Wolfclaw Gloves (1978, -0.16 DPS) [dungeon]; Heavy Earthen Gloves (7359, -0.20 DPS) [crafted] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 25.1 attack_power points (1.69 DPS) | yes | Defiler's Chain Girdle (20152, -0.08 DPS) [rep]; Defiler's Leather Girdle (20191, -0.08 DPS) [rep]; Skulker's Leather Belt (252520, -0.14 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 37.1 attack_power points (2.48 DPS) | yes | Petrolspill Leggings (9509, -0.83 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.83 DPS) [world_drop]; Trapper's Leather Pants (252501, -1.67 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.7 attack_power points (1.19 DPS) | yes | Insignia Boots (4055, -0.24 DPS) [world_drop]; Warsong Boots (16977, -0.24 DPS) [quest]; Brawler's Leather Boots (252439, -0.33 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 17.8 attack_power points (1.20 DPS) | yes | Band of the Fist (17694, -0.40 DPS) [quest]; Ring of Precision (1491, -0.49 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 15.0 attack_power points (1.00 DPS) | yes | Ring of Precision (1491, -0.29 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; Band of the Fist (17694, -0.62 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Claw of the Shadowmancer (2912, -4.36 DPS, sim-verified) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (21.33 DPS) | yes | Zealot Blade (13033, -1.15 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Double-barreled Shotgun (2098, -0.11 DPS) [world_drop]; Crystalpine Stinger (13037, -0.14 DPS, sim-verified) [world_drop]; Gravespike Repeater (273033, -0.17 DPS) [world] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 493, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (orc, 0000000000000000-0000000000000000-500230230050222131)

Set DPS (verified): 175.0. Weights run: 5.7s. Verify run: 18.2s. 817 eligible items had no known source.

Race sweep: orc 175.0 ± 0.1, windshaper-skyborne 174.5 ± 0.1, troll 174.2 ± 0.1, tauren 171.4 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.771 ± 0.016, strength=1.210 ± 0.002, crit=2.007 ± 0.025 per rating point (14 rating = 1%, 28.098 per %), hit=2.420 ± 0.098 per rating point (10 rating = 1%, 24.199 per %), melee_haste=9.528 ± 1.293, intellect=0.044 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 43.8 attack_power points (3.16 DPS) | yes | White Bandit Mask (10008, -0.00 DPS) [crafted]; Warden's Wizard Hat (14604, -0.55 DPS) [world_drop]; Barbaric Iron Helm (7915, -0.58 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 30.5 attack_power points (2.20 DPS) | yes | Ethereal Talisman (4430, -0.96 DPS) [quest]; Kaleidoscope Chain (13084, -1.04 DPS) [world_drop]; Zealous Shadowshard Pendant (17772, -1.19 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 36.5 attack_power points (2.63 DPS) | yes | Flintrock Shoulders (7755, -0.20 DPS) [dungeon]; Nightscape Shoulders (8192, -0.44 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.06 DPS) [dungeon]; Parachute Cloak (10518, -0.06 DPS) [crafted]; Tigerstrike Mantle (13108, -0.06 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 90.7 attack_power points (6.54 DPS) | yes | Quillward Harness (10583, -2.51 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.14 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -3.52 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 22.2 attack_power points (1.60 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.05 DPS) [world_drop]; Branded Leather Bracers (19508, -0.16 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, +0.00 DPS) [crafted]; Dragonscale Gauntlets (8347, -2.88 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 36.1 attack_power points (2.60 DPS) | yes | Ogron's Sash (13117, -0.02 DPS) [world_drop]; Scorpashi Sash (14652, -0.26 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.28 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 58.2 attack_power points (4.19 DPS) | yes | Triprunner Dungarees (9624, -0.34 DPS) [quest]; Petrolspill Leggings (9509, -1.40 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.40 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 39.0 attack_power points (2.81 DPS) | yes | Blackforge Greaves (6423, -0.03 DPS) [dungeon]; Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Worn Running Boots (9398, -0.21 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 31.9 attack_power points (2.30 DPS) | yes | Falcon's Hook (7552, -0.24 DPS) [dungeon]; Disengagement Ring (276202, -0.70 DPS) [vendor]; Mark of Kern (2262, -0.85 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 30.1 attack_power points (2.17 DPS) | yes | Falcon's Hook (7552, -0.11 DPS) [dungeon]; Disengagement Ring (276202, -0.57 DPS) [vendor]; Mark of Kern (2262, -0.73 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.60 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (31.85 DPS) | yes | Sword of Omen (6802, +0.00 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Bow (17686, -0.20 DPS) [quest]; Golemsight Long Gun (273029, -0.20 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.25 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 817, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (orc, 5000000000000000-0030000000000000-500230230050222151)

Set DPS (verified): 283.9. Weights run: 6.0s. Verify run: 25.8s. 1053 eligible items had no known source.

Race sweep: orc 283.9 ± 0.2, windshaper-skyborne 282.1 ± 0.2, troll 281.8 ± 0.2, tauren 278.3 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.630 ± 0.014, strength=1.210 ± 0.002, crit=2.106 ± 0.025 per rating point (14 rating = 1%, 29.484 per %), hit=3.181 ± 0.142 per rating point (10 rating = 1%, 31.814 per %), melee_haste=13.418 ± 1.654, intellect=0.048 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 45.2 attack_power points (4.26 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.05 DPS) [crafted]; Blood Guard's Mail Helmet (220820, -0.20 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.2 attack_power points (3.79 DPS) | yes | Scout's Medallion (19535, -0.82 DPS) [rep]; Woven Ivy Necklace (19159, -0.88 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 39.8 attack_power points (3.75 DPS) | yes | Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -0.03 DPS) [dungeon] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blisterbane Wrap (12552, -0.44 DPS) [dungeon]; Dark Phantom Cape (13122, -0.44 DPS) [world_drop]; Blackmetal Cape (9512, -1.18 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.9 attack_power points (8.93 DPS) | yes | Stone Guard's Chain Armor (220827, -2.40 DPS) [vendor]; Blazewind Breastplate (11193, -2.90 DPS) [quest]; Warbear Harness (15064, -3.23 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 44.3 attack_power points (4.17 DPS) | yes | Skulker's Leather Bracers (252540, -1.14 DPS) [crafted]; Pridelord Bands (14672, -1.31 DPS, sim-verified) [world_drop]; Prowler's Leather Bracers (252539, -1.39 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -4.78 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 51.3 attack_power points (4.83 DPS) | yes | Defiler's Chain Girdle (20151, -0.17 DPS) [rep]; Defiler's Leather Girdle (20193, -0.17 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.27 DPS) [crafted] |
| legs | Infernal Trickster Leggings (17754) | Maraudon: Lord Vyletongue [dungeon] | 84.8 attack_power points (7.98 DPS) | yes | Stone Guard's Chain Legplates (220833, -1.70 DPS) [vendor]; Stormshroud Pants (15057, -2.43 DPS) [crafted]; Serpentskin Leggings (8262, -2.54 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 52.6 attack_power points (4.95 DPS) | yes | Fleetfoot Greaves (11627, -0.22 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.48 DPS) [crafted]; Elven Chain Boots (13125, -0.50 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 51.8 attack_power points (4.88 DPS) | yes | Legionnaire's Band (19511, -1.51 DPS) [rep]; Ring of the Underwood (2951, -2.17 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 36.8 attack_power points (3.47 DPS) | yes | Legionnaire's Band (19511, -0.10 DPS) [rep]; Ring of the Underwood (2951, -0.76 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Devilsaur Tooth (19992, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.43 DPS, sim-verified) [crafted]; Devilsaur Tooth (19992, -1.68 DPS) [quest] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Shadowblade (2163, -9.26 DPS, sim-verified) [world_drop] |
| off_hand | Doomforged Straightedge (12535) | Blackrock Depths [dungeon] | sim-verified (283.9 DPS) | yes | Dawn's Edge (12774, -1.24 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -4.23 DPS) [dungeon]; White Bone Shredder (11863, -6.48 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gryphonwing Long Bow (13022, -1.14 DPS) [dungeon]; Guttbuster (13139, -1.16 DPS, sim-verified) [world_drop]; Houndmaster's Rifle (11629, -1.25 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Blackveil Cape; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Infernal Trickster Leggings; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Bloodrazor; off_hand: Doomforged Straightedge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 1053, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (orc, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 411.1. Weights run: 4.5s. Verify run: 56.0s. 2286 eligible items had no known source.

Race sweep: windshaper-skyborne 413.8 ± 0.4, troll 412.1 ± 0.4, orc 411.1 ± 0.4, tauren 406.2 ± 0.4 (dps); windshaper-skyborne beats the guide race orc by 2.7 (0.66%) ± 0.6, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.190 ± 0.016, strength=1.210 ± 0.002, crit=1.956 ± 0.032 per rating point (14 rating = 1%, 27.387 per %), hit=1.641 ± 0.142 per rating point (10 rating = 1%, 16.409 per %), melee_haste=16.836 ± 1.800, intellect=1.321 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 123.6 attack_power points (16.74 DPS) | yes | Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Warlord's Chain Helmet (16566, -0.79 DPS) [vendor]; Eye of Rend (12587, -7.92 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 53.4 attack_power points (7.23 DPS) | yes | Scout's Medallion (19534, -2.78 DPS) [rep]; Skibi's Pendant (13089, -2.84 DPS, sim-verified) [world_drop]; Imperial Jewel (11933, -2.90 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 86.5 attack_power points (11.72 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -1.07 DPS) [rep]; Howler's Furs (272414, -1.14 DPS) [vendor]; Blackveil Cape (11626, -2.02 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 124.8 attack_power points (16.91 DPS) | yes | Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -0.97 DPS) [vendor]; Timbermaw Tunic (252484, -3.75 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 54.9 attack_power points (7.44 DPS) | yes | General's Chain Wristguards (16570, -0.25 DPS) [pvp]; Windtalker's Wristguards (19582, -0.86 DPS) [rep]; Legionnaire's Chain Bracers (16517, -1.58 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, -1.68 DPS, sim-verified) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 71.1 attack_power points (9.63 DPS) | yes | General's Chain Girdle (16572, -0.25 DPS) [pvp]; Marksman's Girdle (22232, -1.17 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.31 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 120.5 attack_power points (16.32 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, -0.37 DPS) [pvp]; Outrider's Leather Pants (22740, -5.75 DPS, sim-verified) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 65.5 attack_power points (8.87 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -1.01 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -1.95 DPS) [dungeon]; Legionnaire's Band (19510, -2.60 DPS) [rep]; Naglering (11669, -9.35 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -0.50 DPS) [dungeon]; Legionnaire's Band (19510, -1.15 DPS) [rep]; Blackstone Ring (17713, -1.44 DPS) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Hand of Justice (11815, -3.24 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Devilsaur Tooth (19992, +0.00 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -8.62 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (86.46 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -1.90 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -7.96 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.54 DPS) [quest]; Ancient Bone Bow (18680, -0.67 DPS) [dungeon]; Riphook (12653, -0.95 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2286, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (orc, 5000000000000000-0050052010000000-500230230050222151)

Set DPS (verified): 989.8. Weights run: 4.2s. Verify run: 54.3s. 2286 eligible items had no known source.

Race sweep: troll 993.3 ± 1.0, windshaper-skyborne 990.8 ± 1.0, orc 989.8 ± 1.0, tauren 976.1 ± 1.0 (dps); troll beats the guide race orc by 3.6 (0.36%) ± 1.4, beyond error.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.214 ± 0.017, strength=1.210 ± 0.002, crit=2.111 ± 0.033 per rating point (14 rating = 1%, 29.549 per %), hit=2.276 ± 0.222 per rating point (10 rating = 1%, 22.755 per %), melee_haste=33.403 ± 3.047, intellect=1.269 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 124.3 attack_power points (35.94 DPS) | yes | Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Champion's Chain Greathelm (227080, -1.01 DPS) [vendor]; Mask of the Unforgiven (13404, -8.97 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 55.5 attack_power points (16.06 DPS) | yes | Star of Mystaria (12103, -4.94 DPS) [dungeon]; Skibi's Pendant (13089, -5.99 DPS) [world_drop]; Scout's Medallion (19534, -6.46 DPS) [rep] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 89.3 attack_power points (25.83 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Deathguard's Cloak (20068, -1.64 DPS) [rep]; Cape of the Black Baron (13340, -1.93 DPS, sim-verified) [dungeon]; Blackveil Cape (11626, -3.61 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 134.2 attack_power points (38.82 DPS) | yes | Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -4.10 DPS) [vendor]; Timbermaw Tunic (252484, -13.42 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 55.4 attack_power points (16.01 DPS) | yes | General's Chain Wristguards (16570, -0.64 DPS) [pvp]; Windtalker's Wristguards (19582, -2.09 DPS) [rep]; Legionnaire's Chain Bracers (16517, -3.38 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Raider Gloves (272099, -2.41 DPS, sim-verified) [vendor] |
| waist | Marksman's Girdle (22232) | Blackrock Spire: Urok Doomhowl [dungeon] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Girdle (16572, -0.15 DPS) [pvp]; Sash of the Grand Hunt (22207, -1.64 DPS) [dungeon]; Ranger's Belt (272397, -2.30 DPS, sim-verified) [vendor] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 129.8 attack_power points (37.54 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, -2.82 DPS) [pvp]; Outrider's Leather Pants (22740, -16.63 DPS, sim-verified) [rep] |
| feet | Windreaver Greaves (13967) | Scholomance: Kirtonos the Herald [dungeon] | 72.1 attack_power points (20.85 DPS) | yes | General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor]; Bloodmail Boots (14616, -2.33 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -6.00 DPS) [dungeon]; Blackstone Ring (17713, -6.81 DPS) [dungeon]; Naglering (11669, -25.52 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.18 DPS) [dungeon]; Channeler's Ring (272406, -3.83 DPS) [vendor]; Painweaver Band (13098, -7.58 DPS, sim-verified) [dungeon] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Hand of Justice (11815, -11.21 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Annihilator (12798, -48.66 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (184.56 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -12.42 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -17.00 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.91 DPS) [dungeon]; Bloodseeker (19107, -1.47 DPS) [quest]; Ancient Bone Bow (18680, -1.71 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Marksman's Girdle; legs: Outrider's Chain Leggings; feet: Windreaver Greaves; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket2: Blackhand's Breadth; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2286, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

