# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-0000000000000000000-035110001000000000)

Set DPS (verified): 41.2. Weights run: 4.9s. Verify run: 5.8s. 234 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.748 ± 0.009, crit=0.434 ± 0.008 per rating point (14 rating = 1%, 6.079 per %), hit=0.492 ± 0.026 per rating point (10 rating = 1%, 4.923 per %), melee_haste=2.819 ± 0.380

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.0 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.5 attack_power points (0.83 DPS) | yes | Erudite's Amulet (277204, -0.58 DPS, sim-verified) [quest]; Tarnished Locket (279870, -1.07 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.7 attack_power points (0.69 DPS) | yes | Slime-encrusted Pads (6461, -0.71 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-verified (41.2 DPS) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.26 DPS) [crafted]; Bristlebark Cape (14571, -0.28 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.2 attack_power points (1.52 DPS) | yes | Brawler's Leather Armor (252490, -0.18 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.53 DPS) [world_drop]; Defender's Leather Armor (252434, -0.55 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.7 attack_power points (0.69 DPS) | yes | Bravo's Armbands (270015, -0.04 DPS) [quest]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.5 attack_power points (1.15 DPS) | yes | Brawler's Leather Gloves (252494, -0.28 DPS) [crafted]; Bristlebark Gloves (14572, -0.29 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.32 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.42 DPS) | yes | Deviate Scale Belt (6468, -0.50 DPS, sim-verified) [crafted]; Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Murloc Scale Belt (5780, -0.85 DPS) [crafted] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 21.7 attack_power points (1.72 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.08 DPS) [dungeon]; Blackened Defias Leggings (10400, -0.41 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.0 attack_power points (1.34 DPS) | yes | Brawler's Leather Boots (252439, -0.28 DPS, sim-verified) [crafted]; Blackened Defias Boots (10402, -0.51 DPS) [dungeon]; Footpads of the Fang (10411, -0.51 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.0 attack_power points (0.87 DPS) | yes | Demon Band (12054, -0.55 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.0 attack_power points (0.71 DPS) | yes | Demon Band (12054, -0.27 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; The 1 Ring (8350, -0.49 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Diamond Hammer (2194) | World drop [world_drop] | 229.8 attack_power points (18.18 DPS) | yes | Blackfang (2236, -0.09 DPS) [world_drop]; Assassin's Blade (1935, -0.19 DPS) [dungeon]; Buzz Saw (1937, -0.91 DPS) [dungeon] |
| off_hand | Cruel Barb (5191) | The Deadmines: Edwin VanCleef [dungeon] | 229.5 attack_power points (18.16 DPS) | yes | Blackfang (2236, -0.46 DPS, sim-verified) [world_drop] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.0 attack_power points (0.55 DPS) | yes | Fine Longbow (11304, -0.26 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.28 DPS) [crafted]; Light Bow (4576, -0.28 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Diamond Hammer; off_hand: Cruel Barb; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 234, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 00000000000000000-0000000000000000000-035131011311000100)

Set DPS (verified): 66.4. Weights run: 5.9s. Verify run: 15.2s. 411 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.693 ± 0.009, crit=0.660 ± 0.009 per rating point (14 rating = 1%, 9.236 per %), hit=0.875 ± 0.033 per rating point (10 rating = 1%, 8.750 per %), melee_haste=4.480 ± 0.451

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.9 attack_power points (1.41 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.41 DPS) [crafted]; Humbert's Helm (4724, -0.42 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 13.5 attack_power points (1.13 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS) [world]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.79 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.6 attack_power points (1.96 DPS) | yes | Mantle of Thieves (2264, -0.59 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.84 DPS) [crafted]; Bristlebark Amice (14573, -0.87 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.34 DPS) [pvp]; Glowing Lizardscale Cloak (6449, -0.39 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.7 attack_power points (1.97 DPS) | yes | Brawler's Leather Tunic (252508, -0.33 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.42 DPS) [quest]; Brawler's Leather Armor (252490, -0.57 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.2 attack_power points (1.18 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.33 DPS) [world_drop]; Cultist's Armguards (270032, -0.35 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.2 attack_power points (1.43 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.10 DPS) [crafted]; Cartilage Shapers (273054, -0.16 DPS) [world] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.99 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.23 DPS) [crafted]; Prowler's Leather Belt (252459, -0.40 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.16 DPS) | yes | Troll's Bane Leggings (13114, -0.19 DPS) [world_drop]; Petrolspill Leggings (9509, -0.25 DPS, sim-verified) [dungeon]; Brawler's Leather Legguards (252516, -0.31 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.5 attack_power points (1.37 DPS) | yes | Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.26 DPS, sim-verified) [world_drop]; Disjointed Shoes (277226, -0.38 DPS) [quest] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 16.2 attack_power points (1.34 DPS) | yes | Pyrewood Signet Ring (277210, -0.49 DPS) [quest]; Ring of Precision (1491, -0.50 DPS) [dungeon]; Theramore Signet (271911, -0.59 DPS) [rep] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (1.09 DPS) | yes | Ring of Precision (1491, -0.24 DPS) [dungeon]; Theramore Signet (271911, -0.34 DPS) [rep]; Pyrewood Signet Ring (277210, -0.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloody Brass Knuckles (7683, +0.00 DPS) [dungeon]; Zealot Blade (13033, +0.00 DPS) [world_drop]; Claw of the Shadowmancer (2912, -0.52 DPS, sim-verified) [world_drop] |
| off_hand | Thornspike (6681) | Razorfen Kraul: Aggem Thorncurse [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Zealot Blade (13033, -0.49 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 attack_power points (0.84 DPS) | yes | Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Gravespike Repeater (273033, -0.18 DPS) [world]; Silver Star (3463, -0.28 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Thornspike; ranged: Golemsight Long Gun

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-0000000000000000000-035131031321001151)

Set DPS (verified): 100.2. Weights run: 6.6s. Verify run: 16.2s. 544 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.738 ± 0.011, crit=0.974 ± 0.014 per rating point (14 rating = 1%, 13.643 per %), hit=1.216 ± 0.053 per rating point (10 rating = 1%, 12.162 per %), melee_haste=5.208 ± 0.687

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.1 attack_power points (2.56 DPS) | yes | Hawkeye's Helm (14591, -0.60 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.64 DPS) [world_drop]; Nightscape Headband (8176, -0.79 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 19.1 attack_power points (1.62 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Kaleidoscope Chain (13084, -0.69 DPS) [world_drop]; River Pride Choker (13087, -1.28 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (2.05 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Nightscape Shoulders (8192, -0.42 DPS) [crafted]; Mantle of Thieves (2264, -0.57 DPS) [dungeon] |
| back | Sergeant Major's Cape (16336) | PvP rank 5 · Sergeant Major · Alliance [pvp] | sim-verified (100.2 DPS) | yes | Hawkeye's Cloak (14593, -0.11 DPS) [world_drop]; Imperial Cloak (6432, -0.21 DPS) [dungeon]; Parachute Cloak (10518, -0.21 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 29.5 attack_power points (2.51 DPS) | yes | Nightscape Tunic (8175, -0.30 DPS) [crafted]; Dusky Leather Armor (7374, -0.44 DPS) [crafted]; Barbaric Harness (5739, -0.73 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.70 DPS) | yes | Hawkeye's Bracers (14590, -0.48 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.52 DPS) [dungeon]; Dusky Bracers (7378, -0.52 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -3.21 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.55 DPS) | yes | Ogron's Sash (13117, -0.48 DPS, sim-verified) [world_drop]; Highlander's Chain Girdle (20090, -0.51 DPS) [rep]; Skulker's Leather Belt (252520, -0.71 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 36.5 attack_power points (3.10 DPS) | yes | Triprunner Dungarees (9624, -0.19 DPS) [quest]; Ferine Leggings (6690, -0.89 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.03 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 26.1 attack_power points (2.22 DPS) | yes | Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.30 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 21.9 attack_power points (1.86 DPS) | yes | Assault Band (13095, -0.16 DPS) [world_drop]; Ring of the Underwood (2951, -0.21 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.70 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Ardent Custodian (868, -1.23 DPS) [world_drop]; Southsea Lamp (9359, -2.78 DPS) [world_drop]; Coldrage Dagger (10761, -6.41 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 attack_power points (37.52 DPS) | yes | Southsea Lamp (9359, -0.78 DPS, sim-verified) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Swiftwind (13038, -0.16 DPS) [world_drop]; Monolithic Bow (9426, -0.24 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.26 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 544, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 00530200000000000-0000000000000000000-035131031321001151)

Set DPS (verified): 168.2. Weights run: 6.9s. Verify run: 21.0s. 681 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.593 ± 0.011, crit=1.069 ± 0.016 per rating point (14 rating = 1%, 14.967 per %), hit=1.608 ± 0.072 per rating point (10 rating = 1%, 16.079 per %), melee_haste=5.796 ± 1.060

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.5 attack_power points (3.17 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Helm of Fire (8348, -0.16 DPS) [crafted]; Lordrec Helmet (10741, -0.34 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (2.86 DPS) | yes | Sentinel's Medallion (19539, -0.76 DPS, sim-verified) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (2.95 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.13 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (168.2 DPS) | yes | Blackflame Cape (13109, -0.53 DPS) [world_drop]; Serpentskin Cloak (8259, -0.71 DPS) [world_drop]; Pridelord Cape (14673, -0.73 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 49.0 attack_power points (5.45 DPS) | yes | Blazewind Breastplate (11193, -1.04 DPS) [quest]; Warbear Harness (15064, -1.73 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -2.26 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.9 attack_power points (3.10 DPS) | yes | Skulker's Leather Bracers (252540, -0.74 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.77 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.86 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (168.2 DPS) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -4.19 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 35.0 attack_power points (3.88 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.21 DPS) [crafted]; Highlander's Chain Girdle (20090, -1.22 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 49.0 attack_power points (5.45 DPS) | yes | Serpentskin Leggings (8262, -1.17 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.73 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.87 DPS, sim-verified) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.9 attack_power points (3.54 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Prowler's Leather Boots (252468, -0.15 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.82 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.1 attack_power points (4.01 DPS) | yes | Masons Fraternity Ring (9533, -1.53 DPS) [quest]; Mark of Kern (2262, -1.79 DPS) [dungeon]; Assault Band (13095, -1.79 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.3 attack_power points (2.70 DPS) | yes | Masons Fraternity Ring (9533, -0.23 DPS) [quest]; Mark of Kern (2262, -0.48 DPS) [dungeon]; Assault Band (13095, -0.48 DPS) [world_drop] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (168.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-verified (168.2 DPS) | yes | Gut Ripper (2164, -0.63 DPS, sim-verified) [world_drop]; Julie's Dagger (6660, -4.64 DPS) [world_drop]; Hanzo Sword (8190, -4.96 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (58.47 DPS) | yes | Julie's Dagger (6660, +0.00 DPS) [world_drop]; Claw of Celebras (17738, -4.91 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (168.2 DPS) | yes | Skull Splitting Crossbow (13039, -0.59 DPS) [world_drop]; Guttbuster (13139, -0.73 DPS) [world_drop]; Houndmaster's Bow (11628, -0.81 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 681, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 00530201100000000-3200030000000000000-035131031321001151)

Set DPS (verified): 287.0. Weights run: 6.9s. Verify run: 31.7s. 1337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.575 ± 0.011, crit=1.270 ± 0.020 per rating point (14 rating = 1%, 17.777 per %), hit=1.452 ± 0.084 per rating point (10 rating = 1%, 14.517 per %), melee_haste=6.620 ± 1.411

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 59.7 attack_power points (8.64 DPS) | yes | Mask of the Unforgiven (13404, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (3.68 DPS) | yes | Sentinel's Medallion (19538, -0.27 DPS) [rep]; Sunprism Pendant (20649, -0.95 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 60.3 attack_power points (8.72 DPS) | yes | Highlander's Leather Shoulders (20059, -0.28 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, -0.87 DPS) [pvp]; Highlander's Lizardhide Shoulders (20060, -1.65 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -0.25 DPS) [rep]; Stormpike Soldier's Cloak (19084, -2.84 DPS) [rep]; Dark Phantom Cape (13122, -2.89 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 90.6 attack_power points (13.10 DPS) | yes | Timbermaw Tunic (252484, -2.39 DPS, sim-verified) [crafted]; Field Marshal's Leather Chestpiece (231543, -2.51 DPS) [pvp]; Knight-Captain's Leather Chestpiece (227058, -3.51 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 40.9 attack_power points (5.92 DPS) | yes | Marshal's Leather Armsplints (16460, -0.14 DPS) [pvp]; Deepfury Bracers (13120, -1.92 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -2.50 DPS) [pvp] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Theramore Gloves (271948, -2.12 DPS, sim-verified) [rep] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 51.8 attack_power points (7.49 DPS) | yes | Marshal's Leather Cinch (16458, -0.31 DPS) [pvp]; Might of the Timbermaw (19044, -1.27 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -2.56 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 78.1 attack_power points (11.29 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -1.70 DPS) [vendor]; Devilsaur Leggings (15062, -2.14 DPS, sim-verified) [crafted] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 37.4 attack_power points (5.41 DPS) | yes | Highlander's Leather Boots (20052, -0.37 DPS) [rep]; Marshal's Leather Footguards (231546, -0.40 DPS) [pvp]; Albino Crocscale Boots (17728, -0.86 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 46.3 attack_power points (6.70 DPS) | yes | Protector's Band (19514, -2.45 DPS) [rep]; Masons Fraternity Ring (9533, -3.51 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.5 attack_power points (4.99 DPS) | yes | Protector's Band (19516, -1.50 DPS) [rep]; Masons Fraternity Ring (9533, -1.80 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Electrified Dagger (19100, -4.32 DPS, sim-verified) [rep] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -0.55 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -27.49 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodseeker (19107, -0.44 DPS) [quest]; Blackcrow (12651, -0.45 DPS, sim-verified) [dungeon]; Skull Splitting Crossbow (13039, -0.73 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; waist: Highlander's Leather Girdle; legs: Sentinel's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket2: Thunderbrew's Boot Flask; main_hand: Teebu's Blazing Longsword; off_hand: Persuader

No-known-source sample (15 of 1337, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00530201100000000-3200030000000000000-035131031321001151)

Set DPS (verified): 677.4. Weights run: 7.3s. Verify run: 51.7s. 1337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=1.927 ± 0.015, crit=1.596 ± 0.024 per rating point (14 rating = 1%, 22.340 per %), hit=2.029 ± 0.139 per rating point (10 rating = 1%, 20.286 per %), melee_haste=10.102 ± 1.955

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 74.0 attack_power points (17.97 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -2.69 DPS) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.6 attack_power points (7.42 DPS) | yes | Sentinel's Medallion (19538, -0.40 DPS) [rep]; Sunprism Pendant (20649, -1.80 DPS) [quest] |
| shoulder | Theramore Pauldrons (271955) | Theramore Expeditionary Force [rep] | 74.4 attack_power points (18.07 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, -2.37 DPS) [pvp]; Highlander's Leather Shoulders (20059, -2.57 DPS, sim-verified) [rep]; Field Marshal's Leather Epaulets (231547, -3.31 DPS) [pvp] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of the Honor Guard (20073, -1.28 DPS) [rep]; Dark Phantom Cape (13122, -4.86 DPS) [world_drop]; Sergeant Major's Cape (16337, -6.00 DPS) [pvp] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 113.9 attack_power points (27.68 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -5.15 DPS) [pvp]; Timbermaw Tunic (252484, -6.69 DPS, sim-verified) [crafted]; Knight-Captain's Leather Chestpiece (227058, -9.06 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 48.7 attack_power points (11.84 DPS) | yes | Marshal's Leather Armsplints (16460, -0.27 DPS) [pvp]; Deepfury Bracers (13120, -3.74 DPS) [world_drop]; Knight-Captain's Leather Bracers (16394, -4.81 DPS) [pvp] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, -0.79 DPS) [pvp]; Knight-Lieutenant's Leather Gauntlets (16396, -3.13 DPS) [pvp]; Devilsaur Gauntlets (15063, -3.35 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20045) | The League of Arathor [rep] | 56.3 attack_power points (13.69 DPS) | yes | Marshal's Leather Cinch (16458, +0.00 DPS) [pvp]; Might of the Timbermaw (19044, -1.52 DPS) [crafted]; Skulker's Leather Waistguard (252474, -3.93 DPS) [crafted] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 96.7 attack_power points (23.50 DPS) | yes | Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Knight-Captain's Leather Legguards (23299, -4.88 DPS) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 46.6 attack_power points (11.31 DPS) | yes | Marshal's Leather Footguards (231546, -1.01 DPS) [pvp]; Highlander's Leather Boots (20052, -1.81 DPS) [rep]; Albino Crocscale Boots (17728, -1.95 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 56.6 attack_power points (13.76 DPS) | yes | Protector's Band (19514, -5.40 DPS) [rep]; Masons Fraternity Ring (9533, -7.20 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 40.3 attack_power points (9.79 DPS) | yes | Protector's Band (19516, -2.90 DPS) [rep]; Masons Fraternity Ring (9533, -3.23 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+8.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Swiftblade (234579, +0.00 DPS) [pvp]; Grand Marshal's Bonecracker (235480, +0.00 DPS) [vendor]; Greenhammer (279261, -11.79 DPS, sim-verified) [crafted] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [pvp]; Ironwood Blade (279259, -1.42 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -48.68 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Bloodseeker (19107, -0.92 DPS) [quest]; Precisely Calibrated Boomstick (2100, -1.04 DPS, sim-verified) [world_drop]; Eaglehorn Long Bow (13023, -1.65 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Theramore Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Highlander's Leather Girdle; legs: Sentinel's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Teebu's Blazing Longsword; off_hand: Persuader; ranged: Blackcrow

No-known-source sample (15 of 1337, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-035110001000000000)

Set DPS (verified): 41.3. Weights run: 4.9s. Verify run: 8.5s. 233 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.748 ± 0.009, crit=0.434 ± 0.008 per rating point (14 rating = 1%, 6.079 per %), hit=0.492 ± 0.026 per rating point (10 rating = 1%, 4.923 per %), melee_haste=2.819 ± 0.380

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.0 attack_power points (1.11 DPS) | yes | Defender's Leather Hood (252447, -0.47 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.5 attack_power points (0.83 DPS) | yes | Erudite's Amulet (277204, -0.53 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.7 attack_power points (0.69 DPS) | yes | Slime-encrusted Pads (6461, -0.67 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cape of the Brotherhood (5193, -0.14 DPS) [dungeon]; Dark Leather Cloak (2316, -0.26 DPS) [crafted]; Bristlebark Cape (14571, -0.28 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Defender's Leather Armor (252434, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -1.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.7 attack_power points (0.69 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.14 DPS) [vendor]; Spare Part Bindings (279875, -0.28 DPS) [quest] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.5 attack_power points (1.15 DPS) | yes | Bristlebark Gloves (14572, -0.27 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.32 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.42 DPS) | yes | Brawler's Leather Belt (252428, -0.55 DPS) [crafted]; Murloc Scale Belt (5780, -0.85 DPS) [crafted]; Deviate Scale Belt (6468, -2.43 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -1.55 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -1.46 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.0 attack_power points (0.87 DPS) | yes | Bounty Hunter's Ring (5351, -0.45 DPS) [quest]; Demon Band (12054, -0.55 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.59 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.0 attack_power points (0.71 DPS) | yes | Bounty Hunter's Ring (5351, -0.29 DPS) [quest]; Demon Band (12054, -0.39 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Diamond Hammer (2194) | World drop [world_drop] | 229.8 attack_power points (18.18 DPS) | yes | Wingblade (6504, -0.12 DPS) [quest]; Assassin's Blade (1935, -0.19 DPS) [dungeon]; Cruel Barb (5191, -3.75 DPS, sim-verified) [dungeon] |
| off_hand | Blackfang (2236) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Cruel Barb (5191, -0.45 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.0 attack_power points (0.55 DPS) | yes | Fine Longbow (11304, -0.23 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.28 DPS) [crafted]; Light Bow (4576, -0.28 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Diamond Hammer; off_hand: Blackfang; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 233, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-0000000000000000000-035131011311000100)

Set DPS (verified): 66.4. Weights run: 5.9s. Verify run: 14.2s. 412 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.693 ± 0.009, crit=0.660 ± 0.009 per rating point (14 rating = 1%, 9.236 per %), hit=0.875 ± 0.033 per rating point (10 rating = 1%, 8.750 per %), melee_haste=4.480 ± 0.451

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.9 attack_power points (1.41 DPS) | yes | Brawler's Leather Hood (252504, -0.29 DPS, sim-verified) [crafted]; Defender's Leather Helm (252455, -0.41 DPS) [crafted]; Humbert's Helm (4724, -0.42 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 13.5 attack_power points (1.13 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.13 DPS) [world]; Kaleidoscope Chain (13084, -0.23 DPS) [world_drop]; River Pride Choker (13087, -0.79 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 23.6 attack_power points (1.96 DPS) | yes | Mantle of Thieves (2264, -0.58 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.84 DPS) [crafted]; Bristlebark Amice (14573, -0.87 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -0.11 DPS) [world_drop]; Glowing Lizardscale Cloak (6449, -0.39 DPS) [dungeon]; Swiftrunner Cape (6745, -0.39 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 23.7 attack_power points (1.97 DPS) | yes | Brawler's Leather Tunic (252508, -0.32 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.54 DPS) [quest]; Brawler's Leather Armor (252490, -0.57 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 14.2 attack_power points (1.18 DPS) | yes | Barbaric Bracers (18948, -0.29 DPS, sim-verified) [crafted]; Insignia Bracers (6410, -0.33 DPS) [world_drop]; Cultist's Armguards (270032, -0.35 DPS) [quest] |
| hands | Insignia Gloves (6408) | World drop [world_drop] | 17.2 attack_power points (1.43 DPS) | yes | Toughened Leather Gloves (4253, -0.08 DPS) [crafted]; Heavy Earthen Gloves (7359, -0.10 DPS) [crafted]; Cartilage Shapers (273054, -0.16 DPS) [world] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.99 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, -0.23 DPS) [crafted]; Prowler's Leather Belt (252459, -0.40 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (2.16 DPS) | yes | Petrolspill Leggings (9509, -0.19 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.19 DPS) [world_drop]; Brawler's Leather Legguards (252516, -0.31 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 16.5 attack_power points (1.37 DPS) | yes | Warsong Boots (16977, -0.25 DPS) [quest]; Brawler's Leather Boots (252439, -0.26 DPS) [crafted]; Insignia Boots (4055, -0.26 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 16.2 attack_power points (1.34 DPS) | yes | Band of the Fist (17694, -0.45 DPS) [quest]; Pyrewood Signet Ring (277210, -0.49 DPS) [quest]; Ring of Precision (1491, -0.50 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 13.1 attack_power points (1.09 DPS) | yes | Band of the Fist (17694, -0.19 DPS) [quest]; Pyrewood Signet Ring (277210, -0.23 DPS) [quest]; Ring of Precision (1491, -0.24 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Toxic Revenger (9453) | Gnomeregan: Viscous Fallout [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloody Brass Knuckles (7683, +0.00 DPS) [dungeon]; Zealot Blade (13033, +0.00 DPS) [world_drop]; Claw of the Shadowmancer (2912, -0.33 DPS, sim-verified) [world_drop] |
| off_hand | Outlaw Sabre (16886) | Baron Aquanis [quest] | sim-verified (66.4 DPS) | yes | Zealot Blade (13033, -1.24 DPS, sim-verified) [world_drop] |
| ranged | Golemsight Long Gun (273029) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 10.2 attack_power points (0.84 DPS) | yes | Double-barreled Shotgun (2098, -0.17 DPS) [world_drop]; Gravespike Repeater (273033, -0.18 DPS) [world]; Silver Star (3463, -0.28 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Insignia Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Toxic Revenger; off_hand: Outlaw Sabre; ranged: Golemsight Long Gun

No-known-source sample (15 of 412, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 00000000000000000-0000000000000000000-035131031321001151)

Set DPS (verified): 98.8. Weights run: 6.6s. Verify run: 16.4s. 543 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.738 ± 0.011, crit=0.974 ± 0.014 per rating point (14 rating = 1%, 13.643 per %), hit=1.216 ± 0.053 per rating point (10 rating = 1%, 12.162 per %), melee_haste=5.208 ± 0.687

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 30.1 attack_power points (2.56 DPS) | yes | Hawkeye's Helm (14591, -0.60 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.64 DPS) [world_drop]; Nightscape Headband (8176, -0.79 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 19.1 attack_power points (1.62 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.60 DPS) [world]; Ethereal Talisman (4430, -0.61 DPS) [quest]; Kaleidoscope Chain (13084, -0.69 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.1 attack_power points (2.05 DPS) | yes | Flintrock Shoulders (7755, -0.15 DPS) [dungeon]; Nightscape Shoulders (8192, -0.42 DPS) [crafted]; Mantle of Thieves (2264, -0.57 DPS) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | sim-verified (98.8 DPS) | yes | Imperial Cloak (6432, -0.11 DPS) [dungeon]; Parachute Cloak (10518, -0.11 DPS) [crafted]; Tigerstrike Mantle (13108, -0.11 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 29.5 attack_power points (2.51 DPS) | yes | Nightscape Tunic (8175, -0.30 DPS) [crafted]; Dusky Leather Armor (7374, -0.44 DPS) [crafted]; Barbaric Harness (5739, -0.73 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.70 DPS) | yes | Hawkeye's Bracers (14590, -0.51 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.52 DPS) [dungeon]; Dusky Bracers (7378, -0.52 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-verified (98.8 DPS) | yes | Imperial Leather Gloves (4063, +0.00 DPS) [dungeon]; Prowler's Leather Gloves (252524, +0.00 DPS) [crafted]; Skulker's Leather Gloves (252525, -2.99 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.55 DPS) | yes | Ogron's Sash (13117, -0.49 DPS, sim-verified) [world_drop]; Defiler's Chain Girdle (20152, -0.51 DPS) [rep]; Skulker's Leather Belt (252520, -0.71 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 36.5 attack_power points (3.10 DPS) | yes | Triprunner Dungarees (9624, -0.19 DPS) [quest]; Ferine Leggings (6690, -0.89 DPS) [dungeon]; Hawkeye's Breeches (14595, -1.03 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 26.1 attack_power points (2.22 DPS) | yes | Imperial Leather Boots (6431, -0.17 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.25 DPS) [crafted]; Excelsior Boots (4109, -0.30 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 21.9 attack_power points (1.86 DPS) | yes | Assault Band (13095, -0.16 DPS) [world_drop]; Ring of the Underwood (2951, -0.21 DPS) [world_drop]; Falcon's Hook (7552, -0.28 DPS) [dungeon] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.70 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ring of the Underwood (2951, -0.05 DPS) [world_drop]; Falcon's Hook (7552, -0.12 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (98.8 DPS) | yes | Ardent Custodian (868, -1.23 DPS) [world_drop]; Southsea Lamp (9359, -2.78 DPS) [world_drop]; Coldrage Dagger (10761, -6.47 DPS, sim-verified) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 attack_power points (37.52 DPS) | yes | Southsea Lamp (9359, +0.00 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (98.8 DPS) | yes | Swiftwind (13038, -0.16 DPS) [world_drop]; Monolithic Bow (9426, -0.24 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.26 DPS) [quest] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; trinket1: Ankh of Life; main_hand: Gut Ripper; off_hand: Vanquisher's Sword; ranged: The Silencer

No-known-source sample (15 of 543, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 00530200000000000-0000000000000000000-035131031321001151)

Set DPS (verified): 168.4. Weights run: 6.9s. Verify run: 21.9s. 680 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.593 ± 0.011, crit=1.069 ± 0.016 per rating point (14 rating = 1%, 14.967 per %), hit=1.608 ± 0.072 per rating point (10 rating = 1%, 16.079 per %), melee_haste=5.796 ± 1.060

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.5 attack_power points (3.17 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.11 DPS) [quest]; Helm of Fire (8348, -0.16 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.7 attack_power points (2.86 DPS) | yes | Woven Ivy Necklace (19159, -0.60 DPS) [quest]; Scout's Medallion (19535, -0.73 DPS) [rep] |
| shoulder | Skulker's Leather Shoulder (252535) | Leatherworking [crafted] | 26.5 attack_power points (2.95 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Failed Flying Experiment (9647, -0.07 DPS) [quest]; Prowler's Leather Shoulders (252534, -0.13 DPS) [crafted] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackflame Cape (13109, -0.53 DPS) [world_drop]; Serpentskin Cloak (8259, -0.71 DPS) [world_drop]; Pridelord Cape (14673, -0.73 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 49.0 attack_power points (5.45 DPS) | yes | Warbear Harness (15064, -0.99 DPS, sim-verified) [crafted]; Blazewind Breastplate (11193, -1.04 DPS) [quest]; Charred Leather Tunic (19127, -2.26 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.9 attack_power points (3.10 DPS) | yes | Skulker's Leather Bracers (252540, -0.75 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.77 DPS) [world_drop]; Prowler's Leather Bracers (252539, -0.86 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serpentskin Gloves (8260, +0.00 DPS) [world_drop]; First Sergeant's Leather Gauntlets (220857, +0.00 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -3.76 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 35.0 attack_power points (3.88 DPS) | yes | Skulker's Leather Waistguard (252474, -0.07 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.21 DPS) [crafted]; Defiler's Chain Girdle (20152, -1.22 DPS) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 49.0 attack_power points (5.45 DPS) | yes | Serpentskin Leggings (8262, -1.13 DPS, sim-verified) [world_drop]; Basilisk Hide Pants (1718, -1.73 DPS) [world_drop]; Triprunner Dungarees (9624, -1.93 DPS) [quest] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 31.9 attack_power points (3.54 DPS) | yes | Skulker's Leather Boots (252469, -0.02 DPS) [crafted]; Prowler's Leather Boots (252468, -0.15 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.82 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.1 attack_power points (4.01 DPS) | yes | White Bone Band (11862, -1.34 DPS) [quest]; Masons Fraternity Ring (9533, -1.53 DPS) [quest]; Mark of Kern (2262, -1.79 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.3 attack_power points (2.70 DPS) | yes | White Bone Band (11862, -0.04 DPS) [quest]; Masons Fraternity Ring (9533, -0.23 DPS) [quest]; Mark of Kern (2262, -0.48 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.20 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.31 DPS, sim-verified) [crafted] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, -1.18 DPS, sim-verified) [world_drop]; Julie's Dagger (6660, -4.64 DPS) [world_drop]; Hanzo Sword (8190, -4.96 DPS) [world_drop] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 526.4 attack_power points (58.47 DPS) | yes | Julie's Dagger (6660, -0.66 DPS, sim-verified) [world_drop]; Claw of Celebras (17738, -4.91 DPS) [dungeon]; White Bone Shredder (11863, -8.47 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skull Splitting Crossbow (13039, -0.68 DPS, sim-verified) [world_drop]; Guttbuster (13139, -0.73 DPS) [world_drop]; Houndmaster's Bow (11628, -0.81 DPS) [dungeon] |

**New at 50:** neck: Skibi's Pendant; shoulder: Skulker's Leather Shoulder; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; waist: Defiler's Leather Girdle; legs: Stone Guard's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 680, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (troll, 00530201100000000-3200030000000000000-035131031321001151)

Set DPS (verified): 285.9. Weights run: 6.9s. Verify run: 33.8s. 1342 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.575 ± 0.011, crit=1.270 ± 0.020 per rating point (14 rating = 1%, 17.777 per %), hit=1.452 ± 0.084 per rating point (10 rating = 1%, 14.517 per %), melee_haste=6.620 ± 1.411

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 59.7 attack_power points (8.64 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -2.97 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.5 attack_power points (3.68 DPS) | yes | Scout's Medallion (19534, -0.27 DPS) [rep]; Woven Ivy Necklace (19159, -0.77 DPS) [quest]; Sunprism Pendant (20649, -0.95 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 60.3 attack_power points (8.72 DPS) | yes | Defiler's Leather Shoulders (20194, -0.28 DPS) [rep]; Champion's Leather Shoulders (227056, -0.87 DPS) [pvp]; Defiler's Lizardhide Shoulders (20175, -1.65 DPS) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-verified (285.9 DPS) | yes | Deathguard's Cloak (20068, -0.25 DPS) [rep]; Frostwolf Legionnaire's Cloak (19083, -2.84 DPS) [rep]; Dark Phantom Cape (13122, -2.89 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 90.6 attack_power points (13.10 DPS) | yes | Timbermaw Tunic (252484, -1.79 DPS, sim-verified) [crafted]; Warlord's Leather Breastplate (231549, -2.51 DPS) [pvp]; Legionnaire's Leather Chestpiece (227060, -3.51 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 40.9 attack_power points (5.92 DPS) | yes | General's Leather Armsplints (16559, -0.14 DPS) [pvp]; Deepfury Bracers (13120, -1.92 DPS) [world_drop]; Shadowcraft Bracers (16710, -2.50 DPS) [dungeon] |
| hands | Devilsaur Gauntlets (15063) | Leatherworking [crafted] | sim-verified (285.9 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, +0.00 DPS) [vendor]; Gloves of Holy Might (867, -2.84 DPS, sim-verified) [world_drop] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 51.8 attack_power points (7.49 DPS) | yes | General's Leather Girdle (16557, -0.31 DPS) [pvp]; Might of the Timbermaw (19044, -1.28 DPS, sim-verified) [crafted]; Skulker's Leather Waistguard (252474, -2.56 DPS) [crafted] |
| legs | Devilsaur Leggings (15062) | Leatherworking [crafted] | 63.8 attack_power points (9.22 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Outrider's Leather Pants (22740, -4.68 DPS, sim-verified) [rep] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 37.4 attack_power points (5.41 DPS) | yes | Defiler's Leather Boots (20186, -0.37 DPS) [rep]; General's Leather Treads (231552, -0.40 DPS) [pvp]; Albino Crocscale Boots (17728, -0.86 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 46.3 attack_power points (6.70 DPS) | yes | Legionnaire's Band (19510, -2.45 DPS) [rep]; White Bone Band (11862, -3.22 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 34.5 attack_power points (4.99 DPS) | yes | Legionnaire's Band (19510, -0.75 DPS) [rep]; White Bone Band (11862, -1.52 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (285.9 DPS) | yes | Frozen Heart of the Mountain (249469, -1.61 DPS) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (285.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-verified (285.9 DPS) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Glacial Blade (19099, -6.49 DPS, sim-verified) [rep] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (285.9 DPS) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -1.64 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -27.49 DPS) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (285.9 DPS) | yes | Blackcrow (12651, -0.41 DPS) [dungeon]; Bloodseeker (19107, -0.44 DPS) [quest]; Skull Splitting Crossbow (13039, -0.73 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Devilsaur Gauntlets; waist: Defiler's Leather Girdle; legs: Devilsaur Leggings; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; main_hand: Teebu's Blazing Longsword; off_hand: Persuader

No-known-source sample (15 of 1342, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (troll, 00530201100000000-3200030000000000000-035131031321001151)

Set DPS (verified): 672.4. Weights run: 7.3s. Verify run: 50.6s. 1342 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.102 ± 0.017, crit=1.933 ± 0.028 per rating point (14 rating = 1%, 27.057 per %), hit=2.856 ± 0.178 per rating point (10 rating = 1%, 28.559 per %), melee_haste=14.063 ± 2.196

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 84.7 attack_power points (17.41 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Mask of the Unforgiven (13404, -4.52 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 32.8 attack_power points (6.74 DPS) | yes | Scout's Medallion (19534, -0.27 DPS) [rep]; Woven Ivy Necklace (19159, -1.50 DPS) [quest]; Sunprism Pendant (20649, -1.56 DPS) [quest] |
| shoulder | Darkspear Pauldrons (272105) | Creeg Bothunk [vendor] | 83.8 attack_power points (17.22 DPS) | yes | Champion's Leather Shoulders (227056, -1.27 DPS) [pvp]; Warlord's Leather Spaulders (231551, -2.28 DPS) [pvp]; Defiler's Leather Shoulders (20194, -2.93 DPS, sim-verified) [rep] |
| back | Cape of the Black Baron (13340) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -1.44 DPS) [rep]; Dark Phantom Cape (13122, -4.11 DPS) [world_drop]; Cloak of Flames (3475, -4.72 DPS) [world_drop] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 133.4 attack_power points (27.40 DPS) | yes | Warlord's Leather Breastplate (231549, -4.75 DPS) [pvp]; Timbermaw Tunic (252484, -5.18 DPS, sim-verified) [crafted]; Legionnaire's Leather Chestpiece (227060, -8.99 DPS) [pvp] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 52.0 attack_power points (10.69 DPS) | yes | General's Leather Armsplints (16559, -0.23 DPS) [pvp]; Deepfury Bracers (13120, -3.31 DPS) [world_drop]; Shadowcraft Bracers (16710, -4.21 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Holy Might (867, +0.00 DPS) [world_drop]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -2.83 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20190) | The Defilers [rep] | 61.1 attack_power points (12.54 DPS) | yes | General's Leather Girdle (16557, -0.19 DPS) [pvp]; Might of the Timbermaw (19044, -1.75 DPS) [crafted]; Skulker's Leather Waistguard (252474, -3.79 DPS) [crafted] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 85.9 attack_power points (17.65 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Stone Guard's Leather Pants (220859, -2.91 DPS, sim-verified) [vendor] |
| feet | Shadowcraft Boots (16711) | Scholomance: Rattlegore [dungeon] | 52.7 attack_power points (10.83 DPS) | yes | General's Leather Treads (231552, -1.33 DPS) [pvp]; Albino Crocscale Boots (17728, -2.19 DPS) [dungeon]; Defiler's Leather Boots (20186, -2.36 DPS) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 69.6 attack_power points (14.30 DPS) | yes | Legionnaire's Band (19510, -6.84 DPS) [rep]; Masons Fraternity Ring (9533, -8.26 DPS) [quest] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 48.6 attack_power points (9.98 DPS) | yes | Legionnaire's Band (19510, -2.51 DPS) [rep]; Masons Fraternity Ring (9533, -3.93 DPS) [quest] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -1.70 DPS) [crafted]; Mark of the Chosen (17774, -6.69 DPS, sim-verified) [quest] |
| main_hand | Teebu's Blazing Longsword (1728) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Quickblade (234553, +0.00 DPS) [pvp]; High Warlord's Bonecracker (235477, +0.00 DPS) [vendor]; Glacial Blade (19099, -11.36 DPS, sim-verified) [rep] |
| off_hand | Persuader (22384) | Blacksmithing [crafted] | sim-verified (+4.3 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Left Claw (234558, +0.00 DPS) [pvp]; Ironwood Blade (279259, -4.26 DPS, sim-verified) [crafted]; Claw of Celebras (17738, -43.84 DPS) [dungeon] |
| ranged | Blackcrow (12651) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Precisely Calibrated Boomstick (2100, -1.12 DPS) [world_drop]; Bloodseeker (19107, -2.33 DPS) [quest]; Eaglehorn Long Bow (13023, -2.84 DPS) [world_drop] |

**New at 60:** head: Outlaw's Collar; shoulder: Darkspear Pauldrons; back: Cape of the Black Baron; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Defiler's Leather Girdle; legs: Outrider's Leather Pants; feet: Shadowcraft Boots; finger1: Don Julio's Band; finger2: Blackstone Ring; trinket1: Molten Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Teebu's Blazing Longsword; off_hand: Persuader; ranged: Blackcrow

No-known-source sample (15 of 1342, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

