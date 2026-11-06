# Leveling BiS: Survival

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 55.1. Weights run: 5.7s. Verify run: 7.3s. 268 eligible items had no known source.

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

### Band 30 (night-elf, 0000000000000000-0000000000000000-500230231040000100)

Set DPS (verified): 102.7. Weights run: 5.9s. Verify run: 11.8s. 463 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.381 ± 0.013, strength=1.210 ± 0.002, crit=1.353 ± 0.019 per rating point (14 rating = 1%, 18.935 per %), hit=1.825 ± 0.052 per rating point (10 rating = 1%, 18.249 per %), melee_haste=7.448 ± 1.020, intellect=1.188 ± 0.048

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.8 attack_power points (1.63 DPS) | yes | Enduring Cap (3020, -0.33 DPS) [world_drop]; Brawler's Leather Hood (252504, -0.34 DPS, sim-verified) [crafted]; Dokebi Hat (14584, -0.40 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 19.1 attack_power points (1.31 DPS) | yes | Kaleidoscope Chain (13084, -0.00 DPS) [world_drop]; Sentinel's Medallion (19541, -0.01 DPS) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.2 attack_power points (2.21 DPS) | yes | Mantle of Thieves (2264, -0.64 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.98 DPS) [crafted]; Bristlebark Amice (14573, -0.98 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.41 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.61 DPS, sim-verified) [pvp] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 39.3 attack_power points (2.69 DPS) | yes | Green Leather Armor (4255, -0.73 DPS) [crafted]; Dusky Leather Armor (7374, -0.86 DPS, sim-verified) [crafted]; Brawler's Leather Tunic (252508, -0.89 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.1 attack_power points (1.31 DPS) | yes | Insignia Bracers (6410, -0.33 DPS) [world_drop]; Barbaric Bracers (18948, -0.34 DPS, sim-verified) [crafted]; Demonhide Bracers (270033, -0.41 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.8 attack_power points (1.56 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Cartilage Shapers (273054, -0.09 DPS) [world]; Wolfclaw Gloves (1978, -0.17 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.9 attack_power points (2.05 DPS) | yes | Stalker's Leather Belt (252521, -0.09 DPS) [crafted]; Skulker's Leather Belt (252520, -0.32 DPS, sim-verified) [crafted]; Highlander's Chain Girdle (20090, -0.41 DPS) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 40.3 attack_power points (2.76 DPS) | yes | Petrolspill Leggings (9509, -0.47 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.47 DPS) [world_drop]; Trapper's Leather Pants (252501, -1.59 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.7 attack_power points (1.55 DPS) | yes | Insignia Boots (4055, -0.25 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.32 DPS) [crafted]; Trapper's Leather Boots (252440, -0.41 DPS) [crafted] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 21.5 attack_power points (1.48 DPS) | yes | Seal of Wrynn (2933, -0.41 DPS) [quest]; Ring of Precision (1491, -0.50 DPS) [dungeon]; Monkey Ring (6748, -0.66 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.8 attack_power points (1.15 DPS) | yes | Ring of Precision (1491, -0.17 DPS) [dungeon]; Monkey Ring (6748, -0.34 DPS) [quest]; Seal of Wrynn (2933, -0.46 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Claw of the Shadowmancer (2912, -4.82 DPS, sim-verified) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (21.79 DPS) | yes | Zealot Blade (13033, -1.30 DPS, sim-verified) [world_drop]; Shoni's Disarming Tool (9608, -5.90 DPS) [quest] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.3 attack_power points (0.98 DPS) | yes | Double-barreled Shotgun (2098, -0.24 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.33 DPS) [world_drop]; Crystalpine Stinger (13037, -0.51 DPS, sim-verified) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 463, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 40 (night-elf, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 174.4. Weights run: 5.9s. Verify run: 15.5s. 718 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.617 ± 0.015, strength=1.210 ± 0.002, crit=2.003 ± 0.025 per rating point (14 rating = 1%, 28.042 per %), hit=2.431 ± 0.096 per rating point (10 rating = 1%, 24.315 per %), melee_haste=9.540 ± 1.242, intellect=0.044 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 43.8 attack_power points (3.14 DPS) | yes | White Bandit Mask (10008, -0.12 DPS) [crafted]; Barbaric Iron Helm (7915, -0.67 DPS) [crafted]; Warden's Wizard Hat (14604, -0.69 DPS) [world_drop] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 28.8 attack_power points (2.06 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.63 DPS) [quest]; Kaleidoscope Chain (13084, -0.95 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.18 DPS) [world] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.8 attack_power points (2.50 DPS) | yes | Flintrock Shoulders (7755, -0.19 DPS) [dungeon]; Nightscape Shoulders (8192, -0.43 DPS) [crafted]; Mantle of Thieves (2264, -0.62 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hawkeye's Cloak (14593, -0.09 DPS) [world_drop]; Imperial Cloak (6432, -0.16 DPS) [dungeon]; Parachute Cloak (10518, -0.16 DPS) [crafted] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 87.1 attack_power points (6.24 DPS) | yes | Quillward Harness (10583, -2.16 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.05 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -3.41 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 20.9 attack_power points (1.50 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Branded Leather Bracers (19508, -0.07 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.07 DPS) [crafted]; Dragonscale Gauntlets (8347, -3.25 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 36.0 attack_power points (2.58 DPS) | yes | Ogron's Sash (13117, -0.11 DPS) [world_drop]; Scorpashi Sash (14652, -0.36 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.37 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 54.9 attack_power points (3.94 DPS) | yes | Triprunner Dungarees (9624, -0.30 DPS) [quest]; Petrolspill Leggings (9509, -1.31 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.31 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 37.3 attack_power points (2.67 DPS) | yes | Blackforge Greaves (6423, -0.01 DPS) [dungeon]; Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Worn Running Boots (9398, -0.23 DPS) [dungeon] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 30.6 attack_power points (2.19 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.69 DPS) [vendor]; Mark of Kern (2262, -0.76 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 28.6 attack_power points (2.05 DPS) | yes | Falcon's Hook (7552, -0.10 DPS) [dungeon]; Disengagement Ring (276202, -0.55 DPS) [vendor]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Black Menace (6831, -6.95 DPS, sim-verified) [quest] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (31.66 DPS) | yes | Nordic Longshank (9401, +0.00 DPS) [dungeon]; Shoni's Disarming Tool (9608, -15.04 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Bow (17686, -0.19 DPS) [quest]; Golemsight Long Gun (273029, -0.19 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.22 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 718, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (night-elf, 5000000000000000-0010000000000000-500230231051222151)

Set DPS (verified): 270.8. Weights run: 6.5s. Verify run: 24.6s. 900 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.645 ± 0.013, strength=1.210 ± 0.002, crit=2.173 ± 0.024 per rating point (14 rating = 1%, 30.423 per %), hit=3.123 ± 0.135 per rating point (10 rating = 1%, 31.229 per %), melee_haste=13.101 ± 1.600, intellect=0.047 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.2 attack_power points (4.10 DPS) | yes | Knight-Lieutenant's Chain Helmet (220822, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.11 DPS) [crafted]; Knight-Lieutenant's Mail Helmet (223075, -0.19 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.4 attack_power points (3.60 DPS) | yes | Sentinel's Medallion (19539, -0.77 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | sim-verified (270.8 DPS) | yes | Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Knight-Lieutenant's Chain Epaulets (220825, +0.00 DPS) [vendor]; Knight-Lieutenant's Mail Epaulets (223073, +0.00 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.71 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.7 attack_power points (8.42 DPS) | yes | Knight's Chain Armor (220828, -2.68 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -2.69 DPS) [quest]; Warbear Harness (15064, -3.00 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 44.5 attack_power points (3.96 DPS) | yes | Skulker's Leather Bracers (252540, -1.09 DPS) [crafted]; Pridelord Bands (14672, -1.23 DPS, sim-verified) [world_drop]; Arena Wristguards (18709, -1.25 DPS) [world] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -5.82 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 51.5 attack_power points (4.58 DPS) | yes | Highlander's Chain Girdle (20088, -0.10 DPS) [rep]; Highlander's Leather Girdle (20115, -0.10 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.26 DPS) [crafted] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | 67.9 attack_power points (6.04 DPS) | yes | Serpentskin Leggings (8262, -0.88 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.90 DPS) [quest]; Stormshroud Pants (15057, -1.24 DPS, sim-verified) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 52.9 attack_power points (4.70 DPS) | yes | Fleetfoot Greaves (11627, -0.21 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.46 DPS) [crafted]; Elven Chain Boots (13125, -0.47 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 51.2 attack_power points (4.56 DPS) | yes | Protector's Band (19516, -1.36 DPS) [rep]; Ring of the Underwood (2951, -1.99 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 37.0 attack_power points (3.29 DPS) | yes | Protector's Band (19516, -0.10 DPS) [rep]; Ring of the Underwood (2951, -0.73 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Molten Heart of the Mountain (249470, -1.02 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Flurry Axe (871, -7.65 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 540.4 attack_power points (48.06 DPS) | yes | Thorium Cestus (250614, -3.80 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.18 DPS) [dungeon]; Shoni's Disarming Tool (9608, -27.42 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.28 DPS, sim-verified) [world_drop]; Gryphonwing Long Bow (13022, -1.41 DPS) [dungeon]; Swiftwind (13038, -1.65 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Frozen Heart of the Mountain; main_hand: Bloodrazor; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 900, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (night-elf, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 401.0. Weights run: 4.5s. Verify run: 53.2s. 2028 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.235 ± 0.015, strength=1.210 ± 0.002, crit=2.007 ± 0.030 per rating point (14 rating = 1%, 28.099 per %), hit=1.728 ± 0.133 per rating point (10 rating = 1%, 17.281 per %), melee_haste=17.770 ± 1.700, intellect=1.321 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 125.2 attack_power points (16.87 DPS) | yes | Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Field Marshal's Chain Helm (231580, -0.71 DPS) [pvp]; Outlaw's Collar (279253, -5.57 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Sentinel's Medallion (19538, -0.21 DPS) [rep]; Imperial Jewel (11933, -0.42 DPS) [dungeon] |
| shoulder | Theramore Pauldrons (271955) (or Theramore Epaulets (271956)) | Theramore Expeditionary Force [rep] | 88.4 attack_power points (11.91 DPS) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Theramore Epaulets (271956, +0.00 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.11 DPS) [vendor]; Cloak of the Honor Guard (20073, -1.12 DPS) [rep]; Blackveil Cape (11626, -2.02 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 128.1 attack_power points (17.25 DPS) | yes | Field Marshal's Chain Armor (231563, +0.00 DPS) [vendor]; Field Marshal's Chain Breastplate (16466, -1.10 DPS) [vendor]; Timbermaw Tunic (252484, -5.46 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 55.8 attack_power points (7.51 DPS) | yes | Marshal's Chain Bracers (16461, -0.25 DPS) [pvp]; Windtalker's Wristguards (19582, -0.97 DPS) [rep]; Knight-Captain's Chain Armguards (16402, -1.58 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 72.1 attack_power points (9.72 DPS) | yes | Marshal's Chain Girdle (16464, -0.28 DPS) [pvp]; Marksman's Girdle (22232, -1.07 DPS) [dungeon]; Highlander's Chain Girdle (20043, -1.35 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 151.7 attack_power points (20.44 DPS) | yes | Marshal's Chain Legplates (231558, -1.61 DPS) [vendor]; Marshal's Chain Legguards (231577, -4.28 DPS) [pvp]; Sentinel's Leather Pants (237818, -4.74 DPS) [vendor] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.06 DPS) [dungeon]; Protector's Band (19514, -2.73 DPS) [rep]; Naglering (11669, -11.17 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Protector's Band (19514, -1.28 DPS) [rep]; Blackstone Ring (17713, -1.52 DPS) [dungeon]; Painweaver Band (13098, -3.28 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Molten Heart of the Mountain (249470, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -5.96 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (85.98 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -4.55 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -7.92 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.80 DPS) [quest]; Ancient Bone Bow (18680, -0.90 DPS) [dungeon]; Blackcrow (12651, -0.98 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Hand of Justice; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60, raid preset (night-elf, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 966.6. Weights run: 4.1s. Verify run: 53.7s. 2028 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.733 ± 0.027, strength=1.210 ± 0.002, crit=3.377 ± 0.052 per rating point (14 rating = 1%, 47.281 per %), hit=5.221 ± 0.392 per rating point (10 rating = 1%, 52.208 per %), melee_haste=22.997 ± 4.460, intellect=1.283 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 153.4 attack_power points (33.28 DPS) | yes | Lieutenant Commander's Chain Greathelm (227086, +0.00 DPS) [vendor]; Field Marshal's Chain Greathelm (231562, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -18.41 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 73.3 attack_power points (15.90 DPS) | yes | Diana's Pearl Necklace (22403, -2.35 DPS) [dungeon]; Emberfury Talisman (12929, -5.64 DPS) [dungeon]; Star of Mystaria (12103, -8.53 DPS, sim-verified) [dungeon] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Chain Spaulders (16468, +0.00 DPS) [vendor]; Field Marshal's Chain Pauldrons (231557, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -4.81 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanoweave Cloak (272411, -3.85 DPS) [vendor]; Cape of the Black Baron (13340, -4.17 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -4.30 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 200.6 attack_power points (43.52 DPS) | yes | Field Marshal's Chain Armor (231563, -7.31 DPS) [vendor]; Savage Gladiator Chain (11726, -8.53 DPS, sim-verified) [dungeon]; Field Marshal's Chain Breastplate (16466, -9.77 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (+5.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Bracers (16461, -0.35 DPS) [pvp]; Blackmist Armguards (12966, -1.51 DPS) [dungeon]; Slashclaw Bracers (13211, -5.59 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+7.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Chain Grips (231560, +0.00 DPS) [pvp]; Marshal's Chain Vices (231578, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -7.13 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Highlander's Chain Girdle (20043, -0.49 DPS) [rep]; Highlander's Leather Girdle (20045, -0.49 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 242.4 attack_power points (52.60 DPS) | yes | Sentinel's Leather Pants (237818, -16.07 DPS) [vendor]; Marshal's Chain Legplates (231558, -16.38 DPS) [vendor]; Marshal's Chain Legguards (231577, -18.84 DPS) [pvp] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Chain Boots (16462, +0.00 DPS) [vendor]; Marshal's Chain Sabatons (231561, +0.00 DPS) [vendor]; Marshal's Chain Greaves (231579, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.96 DPS) [dungeon]; Channeler's Ring (272406, -9.40 DPS) [vendor]; Naglering (11669, -16.21 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.96 DPS) [dungeon]; Channeler's Ring (272406, -4.41 DPS) [vendor]; Painweaver Band (13098, -5.90 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -5.86 DPS) [dungeon] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Claymore (234565, +0.00 DPS) [pvp]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Ironwood Blade (279259, -22.10 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (138.47 DPS) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -12.76 DPS) [dungeon]; Heartseeker (12783, -25.22 DPS, sim-verified) [crafted] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Blackcrow (12651, -1.65 DPS, sim-verified) [dungeon]; Ancient Bone Bow (18680, -1.78 DPS) [dungeon]; Bloodseeker (19107, -2.05 DPS) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Theramore Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Ranger's Belt; legs: Sentinel's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2028, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

## Horde

### Band 20 (orc, 0000000000000000-0000000000000000-500230100000000000)

Set DPS (verified): 56.1. Weights run: 5.7s. Verify run: 9.3s. 264 eligible items had no known source.

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

### Band 30 (orc, 0000000000000000-0000000000000000-500230231040000100)

Set DPS (verified): 102.1. Weights run: 5.9s. Verify run: 11.5s. 462 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.381 ± 0.013, strength=1.210 ± 0.002, crit=1.353 ± 0.019 per rating point (14 rating = 1%, 18.935 per %), hit=1.825 ± 0.052 per rating point (10 rating = 1%, 18.249 per %), melee_haste=7.448 ± 1.020, intellect=1.188 ± 0.048

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.8 attack_power points (1.63 DPS) | yes | Enduring Cap (3020, -0.33 DPS) [world_drop]; Brawler's Leather Hood (252504, -0.33 DPS, sim-verified) [crafted]; Dokebi Hat (14584, -0.40 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 19.1 attack_power points (1.31 DPS) | yes | Scout's Medallion (19537, -0.01 DPS) [rep]; Kaleidoscope Chain (13084, -0.91 DPS, sim-verified) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 32.2 attack_power points (2.21 DPS) | yes | Mantle of Thieves (2264, -0.97 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.98 DPS) [crafted]; Bristlebark Amice (14573, -0.98 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.41 DPS) [dungeon]; Swiftrunner Cape (6745, -0.41 DPS) [quest] |
| chest | Garb of Florid Feathers (273026) | Excavation Site: Wetlands: Shadetooth [dungeon] | 39.3 attack_power points (2.69 DPS) | yes | Dusky Leather Armor (7374, -0.73 DPS, sim-verified) [crafted]; Green Leather Armor (4255, -0.73 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.89 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 19.1 attack_power points (1.31 DPS) | yes | Insignia Bracers (6410, -0.33 DPS) [world_drop]; Barbaric Bracers (18948, -0.33 DPS, sim-verified) [crafted]; Glowing Magical Bracelets (13106, -0.41 DPS) [world_drop] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 22.8 attack_power points (1.56 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Cartilage Shapers (273054, -0.09 DPS) [world]; Wolfclaw Gloves (1978, -0.17 DPS) [dungeon] |
| waist | Prowler's Leather Belt (252459) | Leatherworking [crafted] | 29.9 attack_power points (2.05 DPS) | yes | Stalker's Leather Belt (252521, -0.09 DPS) [crafted]; Defiler's Chain Girdle (20152, -0.41 DPS) [rep]; Skulker's Leather Belt (252520, -1.34 DPS, sim-verified) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 40.3 attack_power points (2.76 DPS) | yes | Petrolspill Leggings (9509, -0.47 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.47 DPS) [world_drop]; Trapper's Leather Pants (252501, -1.81 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 22.7 attack_power points (1.55 DPS) | yes | Insignia Boots (4055, -0.25 DPS) [world_drop]; Warsong Boots (16977, -0.25 DPS) [quest]; Brawler's Leather Boots (252439, -0.32 DPS) [crafted] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 21.5 attack_power points (1.48 DPS) | yes | Band of the Fist (17694, -0.49 DPS) [quest]; Ring of Precision (1491, -0.50 DPS) [dungeon]; Monkey Ring (6748, -0.66 DPS) [quest] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 16.8 attack_power points (1.15 DPS) | yes | Band of the Fist (17694, -0.17 DPS) [quest]; Ring of Precision (1491, -0.17 DPS) [dungeon]; Monkey Ring (6748, -0.34 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Pronged Reaver (6692) | Razorfen Kraul: Charlga Razorflank [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Corpsemaker (6687, +0.00 DPS) [dungeon]; Felbough Arm (273037, +0.00 DPS) [world]; Claw of the Shadowmancer (2912, -4.36 DPS, sim-verified) [world_drop] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 318.2 attack_power points (21.79 DPS) | yes | Zealot Blade (13033, -1.09 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 14.3 attack_power points (0.98 DPS) | yes | Crystalpine Stinger (13037, -0.00 DPS) [world_drop]; Double-barreled Shotgun (2098, -0.24 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.33 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Garb of Florid Feathers; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Prowler's Leather Belt; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Pronged Reaver; off_hand: Bloody Brass Knuckles; ranged: Golemsight Long Gun

No-known-source sample (15 of 462, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band

### Band 40 (orc, 0000000000000000-0000000000000000-500230231051222111)

Set DPS (verified): 172.6. Weights run: 5.9s. Verify run: 15.4s. 708 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.617 ± 0.015, strength=1.210 ± 0.002, crit=2.003 ± 0.025 per rating point (14 rating = 1%, 28.042 per %), hit=2.431 ± 0.096 per rating point (10 rating = 1%, 24.315 per %), melee_haste=9.540 ± 1.242, intellect=0.044 ± 0.002

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 43.8 attack_power points (3.14 DPS) | yes | White Bandit Mask (10008, -0.12 DPS) [crafted]; Barbaric Iron Helm (7915, -0.67 DPS) [crafted]; Warden's Wizard Hat (14604, -0.69 DPS) [world_drop] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 28.8 attack_power points (2.06 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.63 DPS) [quest]; Ethereal Talisman (4430, -0.88 DPS) [quest]; Kaleidoscope Chain (13084, -0.95 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 34.8 attack_power points (2.50 DPS) | yes | Flintrock Shoulders (7755, -0.19 DPS) [dungeon]; Nightscape Shoulders (8192, -0.43 DPS) [crafted]; Mantle of Thieves (2264, -0.62 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Imperial Cloak (6432, -0.07 DPS) [dungeon]; Parachute Cloak (10518, -0.07 DPS) [crafted]; Tigerstrike Mantle (13108, -0.07 DPS) [world_drop] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 87.1 attack_power points (6.24 DPS) | yes | Quillward Harness (10583, -2.43 DPS, sim-verified) [dungeon]; Wolffear Harness (13110, -3.05 DPS) [world_drop]; Tough Scorpid Breastplate (8203, -3.41 DPS) [crafted] |
| wrist | Dusky Bracers (7378) (or Imperial Leather Bracers (4061)) | Leatherworking [crafted] | 20.9 attack_power points (1.50 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS) [dungeon]; Hawkeye's Bracers (14590, -0.03 DPS) [world_drop]; Branded Leather Bracers (19508, -0.07 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scarlet Gauntlets (10331, +0.00 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.07 DPS) [crafted]; Dragonscale Gauntlets (8347, -2.94 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 36.0 attack_power points (2.58 DPS) | yes | Ogron's Sash (13117, -0.11 DPS) [world_drop]; Scorpashi Sash (14652, -0.36 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.37 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 54.9 attack_power points (3.94 DPS) | yes | Triprunner Dungarees (9624, -0.30 DPS) [quest]; Petrolspill Leggings (9509, -1.31 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.31 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 37.3 attack_power points (2.67 DPS) | yes | Blackforge Greaves (6423, -0.01 DPS) [dungeon]; Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Worn Running Boots (9398, -0.23 DPS) [dungeon] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 30.6 attack_power points (2.19 DPS) | yes | Falcon's Hook (7552, -0.25 DPS) [dungeon]; Disengagement Ring (276202, -0.69 DPS) [vendor]; Mark of Kern (2262, -0.76 DPS) [dungeon] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 28.6 attack_power points (2.05 DPS) | yes | Falcon's Hook (7552, -0.10 DPS) [dungeon]; Disengagement Ring (276202, -0.55 DPS) [vendor]; Mark of Kern (2262, -0.62 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightblade (1982, +0.00 DPS) [world_drop]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Coldrage Dagger (10761, -6.71 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (31.66 DPS) | yes | Sword of Omen (6802, +0.00 DPS) [quest] |
| ranged | Swiftwind (13038) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Master Hunter's Bow (17686, -0.19 DPS) [quest]; Golemsight Long Gun (273029, -0.19 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.22 DPS) [quest] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; chest: Icemail Jerkin; wrist: Dusky Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: Swiftwind

No-known-source sample (15 of 708, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 50 (orc, 5000000000000000-0010000000000000-500230231051222151)

Set DPS (verified): 270.8. Weights run: 6.5s. Verify run: 24.9s. 891 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.645 ± 0.013, strength=1.210 ± 0.002, crit=2.173 ± 0.024 per rating point (14 rating = 1%, 30.423 per %), hit=3.123 ± 0.135 per rating point (10 rating = 1%, 31.229 per %), melee_haste=13.101 ± 1.600, intellect=0.047 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 46.2 attack_power points (4.10 DPS) | yes | Blood Guard's Chain Helmet (220821, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.11 DPS) [crafted]; Blood Guard's Mail Helmet (220820, -0.19 DPS) [vendor] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 40.4 attack_power points (3.60 DPS) | yes | Scout's Medallion (19535, -0.77 DPS) [rep]; Woven Ivy Necklace (19159, -0.83 DPS) [quest] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Blood Guard's Mail Epaulets (220823, +0.00 DPS) [vendor]; Blood Guard's Chain Epaulets (220824, +0.00 DPS) [vendor] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackveil Cape (11626, +0.00 DPS) [dungeon]; Blisterbane Wrap (12552, +0.00 DPS) [dungeon]; Blackmetal Cape (9512, -0.71 DPS) [dungeon] |
| chest | Icemail Jerkin (1981) | World drop [world_drop] | 94.7 attack_power points (8.42 DPS) | yes | Stone Guard's Chain Armor (220827, -2.15 DPS) [vendor]; Blazewind Breastplate (11193, -2.69 DPS) [quest]; Warbear Harness (15064, -3.00 DPS) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 44.5 attack_power points (3.96 DPS) | yes | Skulker's Leather Bracers (252540, -1.09 DPS) [crafted]; Arena Wristguards (18709, -1.25 DPS) [world]; Pridelord Bands (14672, -1.31 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -5.93 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 51.5 attack_power points (4.58 DPS) | yes | Defiler's Chain Girdle (20151, -0.10 DPS) [rep]; Defiler's Leather Girdle (20193, -0.10 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.26 DPS) [crafted] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | 67.9 attack_power points (6.04 DPS) | yes | Serpentskin Leggings (8262, -0.88 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.10 DPS) [world_drop]; Stormshroud Pants (15057, -1.51 DPS, sim-verified) [crafted] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 52.9 attack_power points (4.70 DPS) | yes | Fleetfoot Greaves (11627, -0.21 DPS) [dungeon]; Skulker's Leather Boots (252469, -0.46 DPS) [crafted]; Elven Chain Boots (13125, -0.47 DPS) [world_drop] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 51.2 attack_power points (4.56 DPS) | yes | Legionnaire's Band (19511, -1.36 DPS) [rep]; Ring of the Underwood (2951, -1.99 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 37.0 attack_power points (3.29 DPS) | yes | Legionnaire's Band (19511, -0.10 DPS) [rep]; Ring of the Underwood (2951, -0.73 DPS) [world_drop] |
| trinket1 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+5.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.10 DPS, sim-verified) [crafted] |
| main_hand | Bloodrazor (809) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warmonger (13052, +0.00 DPS) [world_drop]; Stoneraven (13059, +0.00 DPS) [world_drop]; Flurry Axe (871, -4.92 DPS, sim-verified) [world_drop] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | 540.4 attack_power points (48.06 DPS) | yes | Thorium Cestus (250614, -2.95 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -5.18 DPS) [dungeon]; White Bone Shredder (11863, -7.30 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guttbuster (13139, -1.35 DPS, sim-verified) [world_drop]; Gryphonwing Long Bow (13022, -1.41 DPS) [dungeon]; Swiftwind (13038, -1.65 DPS) [world_drop] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; wrist: Deepfury Bracers; waist: Skulker's Leather Waistguard; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Devilsaur Eye; trinket2: Rune of the Guard Captain; main_hand: Bloodrazor; off_hand: Dawn's Edge; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 891, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60 (orc, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 392.7. Weights run: 4.5s. Verify run: 53.8s. 2027 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.235 ± 0.015, strength=1.210 ± 0.002, crit=2.007 ± 0.030 per rating point (14 rating = 1%, 28.099 per %), hit=1.728 ± 0.133 per rating point (10 rating = 1%, 17.281 per %), melee_haste=17.770 ± 1.700, intellect=1.321 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 125.2 attack_power points (16.87 DPS) | yes | Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Warlord's Chain Helmet (16566, -0.71 DPS) [vendor]; Outlaw's Collar (279253, -5.65 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of Fordring (15411, +0.00 DPS) [quest]; Scout's Medallion (19534, -0.21 DPS) [rep]; Imperial Jewel (11933, -0.42 DPS) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) (or Darkspear Epaulets (272106)) | Creeg Bothunk [vendor] | 88.4 attack_power points (11.91 DPS) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Darkspear Epaulets (272106, +0.00 DPS) [vendor] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Howler's Furs (272414, -1.11 DPS) [vendor]; Deathguard's Cloak (20068, -1.12 DPS) [rep]; Blackveil Cape (11626, -2.02 DPS) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 128.1 attack_power points (17.25 DPS) | yes | Warlord's Chain Armor (231566, +0.00 DPS) [vendor]; Warlord's Chain Chestpiece (16565, -1.10 DPS) [vendor]; Timbermaw Tunic (252484, -4.89 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 55.8 attack_power points (7.51 DPS) | yes | General's Chain Wristguards (16570, -0.25 DPS) [pvp]; Windtalker's Wristguards (19582, -0.97 DPS) [rep]; Legionnaire's Chain Bracers (16517, -1.58 DPS) [pvp] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Gloves of Holy Might (867, -2.18 DPS, sim-verified) [world_drop] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | 72.1 attack_power points (9.72 DPS) | yes | General's Chain Girdle (16572, -0.28 DPS) [pvp]; Marksman's Girdle (22232, -1.07 DPS) [dungeon]; Defiler's Chain Girdle (20150, -1.35 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 123.6 attack_power points (16.65 DPS) | yes | General's Chain Legplates (231567, +0.00 DPS) [vendor]; General's Chain Legguards (231574, -0.49 DPS) [pvp]; Outrider's Leather Pants (22740, -6.72 DPS, sim-verified) [rep] |
| feet | Scalegut Treaders (275618) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Painweaver Band (13098, -2.06 DPS) [dungeon]; Legionnaire's Band (19510, -2.73 DPS) [rep]; Naglering (11669, -10.57 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Band (19510, -1.28 DPS) [rep]; Blackstone Ring (17713, -1.52 DPS) [dungeon]; Painweaver Band (13098, -2.61 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, -0.82 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.42 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -5.87 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (85.98 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Alcor's Sunrazor (14555, -3.30 DPS, sim-verified) [world_drop]; Dal'Rend's Tribal Guardian (12939, -7.92 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.80 DPS) [quest]; Ancient Bone Bow (18680, -0.90 DPS) [dungeon]; Blackcrow (12651, -0.98 DPS) [dungeon] |

**New at 60:** head: Black Dragonscale Helm; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Scalegut Treaders; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

### Band 60, raid preset (orc, 5000000000000000-0050050010000000-500230231051222151)

Set DPS (verified): 967.1. Weights run: 4.1s. Verify run: 53.2s. 2027 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, agility=2.733 ± 0.027, strength=1.210 ± 0.002, crit=3.377 ± 0.052 per rating point (14 rating = 1%, 47.281 per %), hit=5.221 ± 0.392 per rating point (10 rating = 1%, 52.208 per %), melee_haste=22.997 ± 4.460, intellect=1.283 ± 0.003

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 153.4 attack_power points (33.28 DPS) | yes | Champion's Chain Greathelm (227080, +0.00 DPS) [vendor]; Warlord's Chain Greathelm (231568, +0.00 DPS) [vendor]; Mask of the Unforgiven (13404, -17.33 DPS, sim-verified) [dungeon] |
| neck | Mark of Fordring (15411) | In Dreams [quest] | 73.3 attack_power points (15.90 DPS) | yes | Diana's Pearl Necklace (22403, -2.35 DPS) [dungeon]; Emberfury Talisman (12929, -5.64 DPS) [dungeon]; Star of Mystaria (12103, -6.99 DPS, sim-verified) [dungeon] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Chain Pauldrons (231565, +0.00 DPS) [vendor]; Warlord's Chain Shoulders (231572, +0.00 DPS) [pvp]; Truestrike Shoulders (12927, -4.96 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanoweave Cloak (272411, -3.85 DPS) [vendor]; Cape of the Black Baron (13340, -4.17 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -4.27 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 200.6 attack_power points (43.52 DPS) | yes | Warlord's Chain Armor (231566, -7.31 DPS) [vendor]; Savage Gladiator Chain (11726, -7.76 DPS, sim-verified) [dungeon]; Warlord's Chain Chestpiece (16565, -9.77 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Wristguards (16570, -0.35 DPS) [pvp]; Blackmist Armguards (12966, -1.51 DPS) [dungeon]; Slashclaw Bracers (13211, -5.20 DPS, sim-verified) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+6.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Chain Gloves (16571, +0.00 DPS) [vendor]; General's Chain Grips (231569, +0.00 DPS) [vendor]; Voone's Vice Grips (13963, -6.76 DPS, sim-verified) [quest] |
| waist | Ranger's Belt (272397) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marksman's Girdle (22232, +0.00 DPS) [dungeon]; Defiler's Chain Girdle (20150, -0.49 DPS) [rep]; Defiler's Leather Girdle (20190, -0.49 DPS) [rep] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 195.1 attack_power points (42.34 DPS) | yes | General's Chain Legplates (231567, -6.12 DPS) [vendor]; General's Chain Legguards (231574, -8.58 DPS) [pvp]; Plaguehound Leggings (18736, -11.49 DPS, sim-verified) [dungeon] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windreaver Greaves (13967, +0.00 DPS) [dungeon]; General's Chain Sabatons (231564, +0.00 DPS) [pvp]; General's Chain Greaves (231570, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -8.96 DPS) [dungeon]; Channeler's Ring (272406, -9.40 DPS) [vendor]; Naglering (11669, -15.78 DPS, sim-verified) [dungeon] |
| finger2 | Cutthroat's Signet (272408) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -3.96 DPS) [dungeon]; Channeler's Ring (272406, -4.41 DPS) [vendor]; Painweaver Band (13098, -5.90 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -0.77 DPS) [crafted]; Hand of Justice (11815, -6.63 DPS) [dungeon] |
| trinket2 | Devilsaur Eye (19991) | The Green Drake [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Greatsword (234542, +0.00 DPS) [pvp]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Ironwood Blade (279259, -20.84 DPS, sim-verified) [crafted] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | 638.2 attack_power points (138.47 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Dal'Rend's Tribal Guardian (12939, -12.76 DPS) [dungeon]; Heartseeker (12783, -25.23 DPS, sim-verified) [crafted] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -4.80 DPS) [world_drop]; Ancient Bone Bow (18680, -6.58 DPS) [dungeon]; Bloodseeker (19107, -6.86 DPS) [quest] |

**New at 60:** head: Black Dragonscale Helm; neck: Mark of Fordring; shoulder: Darkspear Pauldrons; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Ranger's Belt; legs: Outrider's Chain Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Cutthroat's Signet; trinket1: Rune of the Guard Captain; trinket2: Devilsaur Eye; main_hand: Teebu's Blazing Longsword; off_hand: Felstriker; ranged: Blackcrow

No-known-source sample (15 of 2027, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2314 Toughened Leather Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod

