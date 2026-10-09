# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 37.7. Weights run: 2.4s. Verify run: 4.0s. 198 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.832 ± 0.011, crit=0.539 ± 0.011 per rating point (14 rating = 1%, 7.542 per %), hit=0.754 ± 0.027 per rating point (10 rating = 1%, 7.544 per %), melee_haste=4.688 ± 0.293

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.7 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.36 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.0 attack_power points (0.63 DPS) | yes | Erudite's Amulet (277204, -0.45 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.2 attack_power points (0.53 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.0 attack_power points (0.63 DPS) | yes | Cape of the Brotherhood (5193, -0.11 DPS) [dungeon]; Dark Leather Cloak (2316, -0.20 DPS) [crafted]; Bristlebark Cape (14571, -0.21 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 20.2 attack_power points (1.16 DPS) | yes | Brawler's Leather Armor (252490, -0.13 DPS) [crafted]; Prospector's Chestpiece (14562, -0.41 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.42 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.3 attack_power points (0.54 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Bristlebark Bindings (14569, -0.11 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.7 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.82 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.38 DPS) [crafted]; Dusty Belt (279897, -0.51 DPS) [quest]; Deviate Scale Belt (6468, -3.82 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.7 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.03 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (37.7 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.94 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.3 attack_power points (0.65 DPS) | yes | Pyrewood Signet Ring (277210, -0.06 DPS) [quest]; Demon Band (12054, -0.42 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.44 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 11.0 attack_power points (0.63 DPS) | yes | Pyrewood Signet Ring (277210, -0.22 DPS, sim-verified) [quest]; Demon Band (12054, -0.40 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.42 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (13.13 DPS) | yes | Evocator's Blade (2567, -0.82 DPS) [dungeon]; Buzzer Blade (2169, -1.87 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.43 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 227.7 attack_power points (13.08 DPS) | yes | Evocator's Blade (2567, -4.08 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.3 attack_power points (0.42 DPS) | yes | Fine Longbow (11304, -0.18 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.21 DPS) [crafted]; Light Bow (4576, -0.21 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 60.4. Weights run: 2.7s. Verify run: 4.5s. 358 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.455 ± 0.010, crit=0.486 ± 0.011 per rating point (14 rating = 1%, 6.809 per %), hit=1.016 ± 0.034 per rating point (10 rating = 1%, 10.164 per %), melee_haste=7.908 ± 0.371

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 14.5 attack_power points (0.86 DPS) | yes | Tribal Worg Helm (6204, -0.17 DPS) [world]; Brawler's Leather Hood (252504, -0.17 DPS) [crafted]; Defender's Leather Helm (252455, -0.27 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Sentinel's Medallion (19541, -0.14 DPS) [rep]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (1.24 DPS) | yes | Mantle of Thieves (2264, -0.44 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Bristlebark Amice (14573, -0.55 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.2 attack_power points (0.78 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.19 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.20 DPS) [pvp] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 20.4 attack_power points (1.20 DPS) | yes | Brawler's Leather Tunic (252508, -0.16 DPS) [crafted]; Tunic of Westfall (2041, -0.26 DPS) [quest]; Raptorbane Armor (3566, -0.26 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.7 attack_power points (0.75 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Cultist's Armguards (270032, -0.16 DPS) [quest]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (60.4 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -3.27 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (60.4 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.54 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (60.4 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.25 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (60.4 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Disjointed Shoes (277226, +0.00 DPS) [quest]; Feet of the Lynx (1121, -3.51 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 17.1 attack_power points (1.01 DPS) | yes | Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Monkey Ring (6748, -0.41 DPS) [quest]; Pyrewood Signet Ring (277210, -0.43 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.7 attack_power points (0.87 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Monkey Ring (6748, -0.27 DPS) [quest]; Pyrewood Signet Ring (277210, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (18.99 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.30 DPS) [vendor]; Torturing Poker (7682, -1.50 DPS) [dungeon]; Thornspike (6681, -1.99 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (18.87 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.65 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -18.78 DPS) [world_drop] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.71 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.18 DPS) [vendor]; Golemsight Long Gun (273029, -0.19 DPS) [dungeon]; Double-barreled Shotgun (2098, -0.27 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Glass Shooter

No-known-source sample (15 of 358, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 101.1. Weights run: 3.0s. Verify run: 2.6s. 483 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.303 ± 0.008, crit=0.386 ± 0.011 per rating point (14 rating = 1%, 5.401 per %), hit=0.990 ± 0.032 per rating point (10 rating = 1%, 9.897 per %), melee_haste=5.242 ± 0.373

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.3 attack_power points (1.79 DPS) | yes | Warden's Wizard Hat (14604, -0.59 DPS) [world_drop]; Hawkeye's Helm (14591, -0.60 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.68 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.41 DPS) | yes | Sentinel's Medallion (19540, -0.40 DPS) [rep]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.3 attack_power points (1.86 DPS) | yes | Flintrock Shoulders (7755, -0.59 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.60 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.0 attack_power points (1.20 DPS) | yes | Hawkeye's Cloak (14593, -0.35 DPS) [world_drop]; Sergeant Major's Cape (16336, -0.38 DPS, sim-verified) [pvp]; Yeti Fur Cloak (2805, -0.44 DPS) [quest] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 32.7 attack_power points (2.31 DPS) | yes | Nightscape Tunic (8175, -0.93 DPS) [crafted]; Wolffear Harness (13110, -0.97 DPS, sim-verified) [world_drop]; Dusky Leather Armor (7374, -1.02 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Hawkeye's Bracers (14590, -0.56 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.68 DPS) [dungeon]; Dusky Bracers (7378, -0.68 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 25.4 attack_power points (1.79 DPS) | yes | Prowler's Leather Gloves (252524, -0.26 DPS) [crafted]; Imperial Leather Gloves (4063, -0.33 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.51 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.12 DPS) | yes | Highlander's Chain Girdle (20090, -0.51 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.65 DPS) [world_drop]; Blackened Defias Belt (10403, -0.85 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.4 attack_power points (1.93 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.10 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.61 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.3 attack_power points (1.51 DPS) | yes | Prowler's Leather Shoes (252465, -0.09 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.18 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Protector's Band (19515, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.41 DPS) | yes | Protector's Band (19515, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (101.1 DPS) | yes | Black Menace (6831, -4.17 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -4.90 DPS) [vendor]; Coldrage Dagger (10761, -8.75 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 441.4 attack_power points (31.14 DPS) | yes | Black Menace (6831, -1.51 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -30.93 DPS) [world_drop]; Satyr's Rod (15962, -31.05 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (101.1 DPS) | yes | Glass Shooter (9456, -0.14 DPS) [dungeon]; Monolithic Bow (9426, -0.29 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.20 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 483, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 142.7. Weights run: 3.6s. Verify run: 3.2s. 610 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.438 ± 0.007, crit=0.739 ± 0.011 per rating point (14 rating = 1%, 10.339 per %), hit=1.515 ± 0.056 per rating point (10 rating = 1%, 15.149 per %), melee_haste=5.320 ± 0.803

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 46.3 attack_power points (3.93 DPS) | yes | Embrace of the Lycan (9479, -0.54 DPS) [dungeon]; Knight-Lieutenant's Leather Headband (220850, -0.54 DPS, sim-verified) [vendor]; White Bandit Mask (10008, -1.66 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.7 attack_power points (2.01 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.31 DPS) [quest]; Sentinel's Medallion (19539, -0.55 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 33.5 attack_power points (2.84 DPS) | yes | Skulker's Leather Shoulder (252535, -0.74 DPS) [crafted]; Failed Flying Experiment (9647, -0.77 DPS) [quest]; Sunburn Spaulders (274751, -0.79 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.1 attack_power points (2.22 DPS) | yes | Blisterbane Wrap (12552, -0.39 DPS) [dungeon]; Dark Phantom Cape (13122, -0.39 DPS) [world_drop]; Duskbat Drape (19982, -0.51 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 43.5 attack_power points (3.69 DPS) | yes | Blazewind Breastplate (11193, -0.63 DPS) [quest]; Fungus Shroud Armor (17742, -0.64 DPS) [dungeon]; Warbear Harness (15064, -0.84 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.6 attack_power points (2.17 DPS) | yes | Skulker's Leather Bracers (252540, -0.48 DPS) [crafted]; Pridelord Bands (14672, -0.52 DPS) [world_drop]; Branded Leather Bracers (19508, -0.66 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.6 attack_power points (3.53 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.96 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.06 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.22 DPS) | yes | Skulker's Leather Waistguard (252474, -0.57 DPS, sim-verified) [crafted]; Prowler's Leather Waistguard (252473, -0.57 DPS) [crafted]; Highlander's Leather Girdle (20115, -0.65 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 43.5 attack_power points (3.69 DPS) | yes | Serpentskin Leggings (8262, -0.63 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.87 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.13 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.4 attack_power points (2.58 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Prowler's Leather Boots (252468, -0.14 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.9 attack_power points (1.95 DPS) | yes | Mark of Kern (2262, -0.25 DPS) [dungeon]; Assault Band (13095, -0.25 DPS) [world_drop]; Blackstone Ring (17713, -0.25 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 20.1 attack_power points (1.71 DPS) | yes | Mark of Kern (2262, -0.01 DPS) [dungeon]; Assault Band (13095, -0.01 DPS) [world_drop]; Blackstone Ring (17713, -0.01 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (142.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (142.7 DPS) | yes | Mark of the Chosen (17774, -0.51 DPS, sim-verified) [quest] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (142.7 DPS) | yes | Lifeforce Dirk (10750, -0.66 DPS) [quest]; Charstone Dirk (17710, -0.66 DPS) [dungeon]; Shadowblade (2163, -4.55 DPS, sim-verified) [world_drop] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (43.38 DPS) | yes | Thermotastic Egg Timer (9644, -43.02 DPS) [quest]; Stonecloth Branch (15963, -43.13 DPS) [world_drop]; Windchaser Orb (15965, -43.26 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (142.7 DPS) | yes | Stinging Bow (10624, -0.27 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.27 DPS) [world_drop]; Dark Iron Rifle (16004, -2.15 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 610, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 263.0. Weights run: 3.5s. Verify run: 9.6s. 1378 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.875 ± 0.018, crit=1.812 ± 0.032 per rating point (14 rating = 1%, 25.366 per %), hit=2.906 ± 0.132 per rating point (10 rating = 1%, 29.060 per %), melee_haste=10.925 ± 1.906

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 77.4 attack_power points (7.74 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.52 DPS) [quest] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 57.2 attack_power points (5.72 DPS) | yes | Beads of Ogre Might (22150, -0.41 DPS) [quest]; Mark of Fordring (15411, -0.58 DPS) [quest]; Medallion of the Dawn (22659, -0.78 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (263.0 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -13.45 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 57.1 attack_power points (5.71 DPS) | yes | Cape of the Black Baron (13340, -0.89 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.37 DPS) [rep]; Windshear Cape (20691, -2.09 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (263.0 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -2.06 DPS) [pvp]; Darkmantle Tunic (226825, -2.21 DPS) [quest]; Tunic of Undead Slaying (23089, -10.77 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (263.0 DPS) | yes | Marshal's Leather Armsplints (16460, -0.10 DPS) [pvp]; Bracers of the Eclipse (18375, -0.39 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.97 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (263.0 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.11 DPS) [crafted]; Raider Gloves (272099, -15.76 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 97.5 attack_power points (9.75 DPS) | yes | Belt of Preserved Heads (20216, -1.68 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -3.81 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.96 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (263.0 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -13.37 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 65.3 attack_power points (6.54 DPS) | yes | Fine Dawn Treaders (227815, -1.67 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.73 DPS) [dungeon]; Darkmantle Boots (22003, -2.03 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-verified (263.0 DPS) | yes | Tarnished Elven Ring (18500, -1.69 DPS) [dungeon]; Cutthroat's Signet (272408, -1.88 DPS) [vendor]; Naglering (11669, -6.83 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (263.0 DPS) | yes | Tarnished Elven Ring (18500, -1.32 DPS) [dungeon]; Cutthroat's Signet (272408, -1.51 DPS) [vendor]; Naglering (11669, -6.09 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (263.0 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -6.34 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-verified (263.0 DPS) | yes | Hand of Justice (11815, +0.00 DPS, sim-verified) [dungeon]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.46 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (263.0 DPS) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (263.0 DPS) | yes | The Lobotomizer (19324, -4.46 DPS, sim-verified) [rep]; Distracting Dagger (18392, -6.75 DPS) [dungeon]; Scepter of Interminable Focus (22329, -58.40 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (263.0 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.84 DPS) [world_drop]; The Purifier (22656, -0.93 DPS) [quest]; Dark Iron Rifle (16004, -2.93 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1378, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 626.4. Weights run: 3.7s. Verify run: 10.3s. 1378 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.309 ± 0.024, crit=2.291 ± 0.040 per rating point (14 rating = 1%, 32.072 per %), hit=4.258 ± 0.245 per rating point (10 rating = 1%, 42.584 per %), melee_haste=17.990 ± 3.114

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 96.6 attack_power points (17.68 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Knight-Lieutenant's Leather Headband (220850, -1.09 DPS) [vendor] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 77.2 attack_power points (14.14 DPS) | yes | Beads of Ogre Might (22150, -1.95 DPS) [quest]; Mark of Fordring (15411, -3.51 DPS) [quest]; Medallion of the Dawn (22659, -3.87 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (626.4 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -31.05 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.6 attack_power points (12.92 DPS) | yes | Cape of the Black Baron (13340, -2.92 DPS) [dungeon]; Cloak of the Honor Guard (20073, -4.58 DPS) [rep]; Windshear Cape (20691, -4.97 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (626.4 DPS) | yes | Darkmantle Tunic (226825, -4.51 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -4.65 DPS) [pvp]; Tunic of Undead Slaying (23089, -24.62 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-verified (626.4 DPS) | yes | Marshal's Leather Armsplints (16460, -0.20 DPS) [pvp]; Blackmist Armguards (12966, -1.44 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -10.78 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 74.7 attack_power points (13.67 DPS) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.95 DPS) [quest]; Raider Gloves (272099, -41.47 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 133.7 attack_power points (24.47 DPS) | yes | Belt of Preserved Heads (20216, -4.05 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -12.05 DPS) [vendor]; Highlander's Leather Girdle (20045, -12.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (626.4 DPS) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -33.86 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 85.2 attack_power points (15.60 DPS) | yes | Fine Dawn Treaders (227815, -3.09 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -4.39 DPS) [dungeon]; Darkmantle Boots (22003, -5.46 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-verified (626.4 DPS) | yes | Tarnished Elven Ring (18500, -3.80 DPS) [dungeon]; Cutthroat's Signet (272408, -4.23 DPS) [vendor]; Naglering (11669, -15.07 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (626.4 DPS) | yes | Tarnished Elven Ring (18500, -2.46 DPS) [dungeon]; Cutthroat's Signet (272408, -2.88 DPS) [vendor]; Naglering (11669, -12.91 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (626.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-verified (626.4 DPS) | yes | Darkmoon Card: Maelstrom (19289, +0.00 DPS, sim-verified) [quest]; Blackhand's Breadth (13965, -3.85 DPS) [quest]; Frozen Heart of the Mountain (249469, -8.57 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (626.4 DPS) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (626.4 DPS) | yes | The Lobotomizer (19324, -9.59 DPS, sim-verified) [rep]; Distracting Dagger (18392, -12.35 DPS) [dungeon]; Scepter of Interminable Focus (22329, -103.17 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (626.4 DPS) | yes | Precisely Calibrated Boomstick (2100, -3.15 DPS) [world_drop]; The Purifier (22656, -3.19 DPS) [quest]; Dark Iron Rifle (16004, -6.31 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Hand of Justice; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1378, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 37.3. Weights run: 2.4s. Verify run: 4.1s. 191 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.832 ± 0.011, crit=0.539 ± 0.011 per rating point (14 rating = 1%, 7.542 per %), hit=0.754 ± 0.027 per rating point (10 rating = 1%, 7.544 per %), melee_haste=4.688 ± 0.293

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.7 attack_power points (0.84 DPS) | yes | Defender's Leather Hood (252447, -0.36 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.0 attack_power points (0.63 DPS) | yes | Erudite's Amulet (277204, -0.44 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.2 attack_power points (0.53 DPS) | yes | Slime-encrusted Pads (6461, -0.55 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.0 attack_power points (0.63 DPS) | yes | Cape of the Brotherhood (5193, -0.11 DPS) [dungeon]; Dark Leather Cloak (2316, -0.20 DPS) [crafted]; Bristlebark Cape (14571, -0.21 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.3 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.86 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.2 attack_power points (0.53 DPS) | yes | Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.11 DPS) [vendor]; Ratchet Wristwraps (274742, -0.21 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.0 attack_power points (0.86 DPS) | yes | Brawler's Leather Gloves (252494, -0.21 DPS) [crafted]; Bristlebark Gloves (14572, -0.21 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.23 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.38 DPS) [crafted]; Murloc Scale Belt (5780, -0.60 DPS) [crafted]; Deviate Scale Belt (6468, -3.83 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.3 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.04 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (37.3 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.94 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.3 attack_power points (0.65 DPS) | yes | Pyrewood Signet Ring (277210, -0.06 DPS) [quest]; Bounty Hunter's Ring (5351, -0.33 DPS) [quest]; Demon Band (12054, -0.42 DPS) [world_drop] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 11.0 attack_power points (0.63 DPS) | yes | Pyrewood Signet Ring (277210, -0.21 DPS, sim-verified) [quest]; Bounty Hunter's Ring (5351, -0.32 DPS) [quest]; Demon Band (12054, -0.40 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (13.13 DPS) | yes | Scout's Blade (20441, -0.42 DPS) [pvp]; Evocator's Blade (2567, -0.82 DPS) [dungeon]; Buzzer Blade (2169, -1.87 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 227.7 attack_power points (13.08 DPS) | yes | Evocator's Blade (2567, -3.96 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -12.96 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.3 attack_power points (0.42 DPS) | yes | Fine Longbow (11304, -0.18 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.21 DPS) [crafted]; Light Bow (4576, -0.21 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 60.4. Weights run: 2.7s. Verify run: 4.6s. 348 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.455 ± 0.010, crit=0.486 ± 0.011 per rating point (14 rating = 1%, 6.809 per %), hit=1.016 ± 0.034 per rating point (10 rating = 1%, 10.164 per %), melee_haste=7.908 ± 0.371

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 14.5 attack_power points (0.86 DPS) | yes | Tribal Worg Helm (6204, -0.17 DPS) [world]; Brawler's Leather Hood (252504, -0.17 DPS) [crafted]; Defender's Leather Helm (252455, -0.32 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.83 DPS) | yes | Scout's Medallion (19537, -0.14 DPS) [rep]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (1.24 DPS) | yes | Mantle of Thieves (2264, -0.43 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.51 DPS) [crafted]; Bristlebark Amice (14573, -0.55 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.2 attack_power points (0.78 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.19 DPS) [dungeon]; Wildhunter Cloak (16658, -0.19 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 20.4 attack_power points (1.20 DPS) | yes | Brawler's Leather Tunic (252508, -0.26 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.31 DPS) [crafted]; Panther Armor (6670, -0.31 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.7 attack_power points (0.75 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Cultist's Armguards (270032, -0.16 DPS) [quest]; Barbaric Bracers (18948, -0.17 DPS) [crafted] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -2.64 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -2.91 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.62 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Vorrel's Boots (7751, +0.00 DPS) [quest]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.87 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 17.1 attack_power points (1.01 DPS) | yes | Thunderbrow Ring (13097, -0.28 DPS) [world_drop]; Monkey Ring (6748, -0.41 DPS) [quest]; Pyrewood Signet Ring (277210, -0.43 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.7 attack_power points (0.87 DPS) | yes | Thunderbrow Ring (13097, -0.14 DPS) [world_drop]; Monkey Ring (6748, -0.27 DPS) [quest]; Pyrewood Signet Ring (277210, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (18.99 DPS) | yes | Scout's Blade (19545, -0.90 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.30 DPS) [vendor]; Scorn's Focal Dagger (23168, -4.32 DPS, sim-verified) [dungeon] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | sim-verified (60.4 DPS) | yes | Scorn's Focal Dagger (23168, -0.76 DPS, sim-verified) [dungeon]; Tork Wrench (11855, -18.70 DPS) [quest]; Satyr's Rod (15962, -18.73 DPS) [world_drop] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.71 DPS) | yes | Booty Bay Bruiser's Buckshot (274748, -0.18 DPS) [vendor]; Golemsight Long Gun (273029, -0.19 DPS) [dungeon]; Double-barreled Shotgun (2098, -0.27 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Serrated Raptor Claw; ranged: Glass Shooter

No-known-source sample (15 of 348, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 99.9. Weights run: 3.0s. Verify run: 2.7s. 465 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.303 ± 0.008, crit=0.386 ± 0.011 per rating point (14 rating = 1%, 5.401 per %), hit=0.990 ± 0.032 per rating point (10 rating = 1%, 9.897 per %), melee_haste=5.242 ± 0.373

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.3 attack_power points (1.79 DPS) | yes | Hawkeye's Helm (14591, -0.59 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.59 DPS) [world_drop]; Nightscape Headband (8176, -0.68 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.41 DPS) | yes | Scout's Medallion (19536, -0.40 DPS) [rep]; Ghostshard Talisman (7731, -0.42 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.3 attack_power points (1.86 DPS) | yes | Flintrock Shoulders (7755, -0.59 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.59 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.85 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.0 attack_power points (1.20 DPS) | yes | Hawkeye's Cloak (14593, -0.35 DPS) [world_drop]; First Sergeant's Cloak (16340, -0.41 DPS, sim-verified) [pvp]; Parachute Cloak (10518, -0.47 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 32.7 attack_power points (2.31 DPS) | yes | Nightscape Tunic (8175, -0.93 DPS) [crafted]; Wolffear Harness (13110, -0.97 DPS, sim-verified) [world_drop]; Dusky Leather Armor (7374, -1.02 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.52 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.68 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 25.4 attack_power points (1.79 DPS) | yes | Prowler's Leather Gloves (252524, -0.26 DPS) [crafted]; Imperial Leather Gloves (4063, -0.33 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.54 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.12 DPS) | yes | Defiler's Chain Girdle (20152, -0.51 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.65 DPS) [world_drop]; Blackened Defias Belt (10403, -0.85 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.4 attack_power points (1.93 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.10 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.61 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.3 attack_power points (1.51 DPS) | yes | Prowler's Leather Shoes (252465, -0.09 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.18 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.41 DPS) | yes | Legionnaire's Band (19512, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.41 DPS) | yes | Legionnaire's Band (19512, -0.11 DPS) [rep]; Ironspine's Eye (7686, -0.30 DPS) [dungeon]; Ring of the Underwood (2951, -0.35 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (99.9 DPS) | yes | Scout's Blade (19544, -3.98 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -4.90 DPS) [vendor]; Coldrage Dagger (10761, -8.39 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 441.4 attack_power points (31.14 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon]; Stonecloth Branch (15963, -30.93 DPS) [world_drop]; Tork Wrench (11855, -31.00 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (99.9 DPS) | yes | Glass Shooter (9456, -0.14 DPS) [dungeon]; Monolithic Bow (9426, -0.29 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.18 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 465, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 146.3. Weights run: 3.6s. Verify run: 3.2s. 587 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.438 ± 0.007, crit=0.739 ± 0.011 per rating point (14 rating = 1%, 10.339 per %), hit=1.515 ± 0.056 per rating point (10 rating = 1%, 15.149 per %), melee_haste=5.320 ± 0.803

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 46.3 attack_power points (3.93 DPS) | yes | Blood Guard's Leather Headband (220851, -0.41 DPS) [vendor]; Embrace of the Lycan (9479, -0.82 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.66 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.7 attack_power points (2.01 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.31 DPS) [quest]; Woven Ivy Necklace (19159, -0.40 DPS) [quest]; Scout's Medallion (19535, -0.55 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.8 attack_power points (2.36 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.29 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.1 attack_power points (2.22 DPS) | yes | Blisterbane Wrap (12552, -0.39 DPS) [dungeon]; Dark Phantom Cape (13122, -0.39 DPS) [world_drop]; Duskbat Drape (19982, -0.51 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 43.5 attack_power points (3.69 DPS) | yes | Blazewind Breastplate (11193, -0.63 DPS) [quest]; Fungus Shroud Armor (17742, -0.64 DPS) [dungeon]; Warbear Harness (15064, -1.29 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.6 attack_power points (2.17 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.66 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.6 attack_power points (3.53 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.96 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.06 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.22 DPS) | yes | Skulker's Leather Waistguard (252474, -0.50 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.57 DPS) [crafted]; Defiler's Leather Girdle (20193, -0.65 DPS) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 43.5 attack_power points (3.69 DPS) | yes | Basilisk Hide Pants (1718, -1.13 DPS) [world_drop]; Triprunner Dungarees (9624, -1.24 DPS) [quest]; Serpentskin Leggings (8262, -1.39 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 30.4 attack_power points (2.58 DPS) | yes | Skulker's Leather Boots (252469, -0.06 DPS) [crafted]; Prowler's Leather Boots (252468, -0.14 DPS) [crafted]; Albino Crocscale Boots (17728, -0.14 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.04 DPS) | yes | Masons Fraternity Ring (9533, -0.33 DPS) [quest]; Mark of Kern (2262, -0.34 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 22.9 attack_power points (1.95 DPS) | yes | Masons Fraternity Ring (9533, -0.24 DPS) [quest]; Mark of Kern (2262, -0.25 DPS) [dungeon]; Blackstone Ring (17713, -0.25 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (146.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (146.3 DPS) | yes | Molten Heart of the Mountain (249470, -1.09 DPS, sim-verified) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (146.3 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Scout's Blade (19543, -0.44 DPS) [pvp]; Searing Needle (12531, -5.45 DPS, sim-verified) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (43.38 DPS) | yes | Thermotastic Egg Timer (9644, -43.02 DPS) [quest]; Stonecloth Branch (15963, -43.13 DPS) [world_drop]; Tork Wrench (11855, -43.22 DPS) [quest] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (146.3 DPS) | yes | Stinging Bow (10624, -0.27 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.27 DPS) [world_drop]; Dark Iron Rifle (16004, -2.12 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: White Bone Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 587, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 265.4. Weights run: 3.5s. Verify run: 9.7s. 1372 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.875 ± 0.018, crit=1.812 ± 0.032 per rating point (14 rating = 1%, 25.366 per %), hit=2.906 ± 0.132 per rating point (10 rating = 1%, 29.060 per %), melee_haste=10.925 ± 1.906

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 77.4 attack_power points (7.74 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.52 DPS) [quest] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 57.2 attack_power points (5.72 DPS) | yes | Beads of Ogre Might (22150, -0.41 DPS) [quest]; Mark of Fordring (15411, -0.58 DPS) [quest]; Medallion of the Dawn (22659, -0.78 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (265.4 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -13.06 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 57.1 attack_power points (5.71 DPS) | yes | Cape of the Black Baron (13340, -0.89 DPS) [dungeon]; Deathguard's Cloak (20068, -1.37 DPS) [rep]; Windshear Cape (20691, -2.09 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (265.4 DPS) | yes | Warlord's Leather Breastplate (231549, -2.06 DPS) [pvp]; Darkmantle Tunic (226825, -2.21 DPS) [quest]; Tunic of Undead Slaying (23089, -10.34 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (265.4 DPS) | yes | General's Leather Armsplints (16559, -0.10 DPS) [pvp]; Bracers of the Eclipse (18375, -0.39 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.97 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (265.4 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.11 DPS) [crafted]; Raider Gloves (272099, -16.00 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 97.5 attack_power points (9.75 DPS) | yes | Belt of Preserved Heads (20216, -1.07 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -3.81 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.96 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (265.4 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -13.38 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 65.3 attack_power points (6.54 DPS) | yes | Shadowcraft Boots (16711, -1.73 DPS) [dungeon]; Darkmantle Boots (22003, -2.03 DPS) [quest]; Fine Dawn Treaders (227815, -2.17 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-verified (265.4 DPS) | yes | Tarnished Elven Ring (18500, -1.69 DPS) [dungeon]; Cutthroat's Signet (272408, -1.88 DPS) [vendor]; Naglering (11669, -6.21 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (265.4 DPS) | yes | Tarnished Elven Ring (18500, -1.32 DPS) [dungeon]; Cutthroat's Signet (272408, -1.51 DPS) [vendor]; Naglering (11669, -5.50 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (265.4 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (265.4 DPS) | yes | Royal Seal of Eldre'Thalas (18465, -0.42 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.62 DPS) [crafted]; Blackhand's Breadth (13965, -3.66 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (265.4 DPS) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (265.4 DPS) | yes | The Lobotomizer (19324, -6.17 DPS, sim-verified) [rep]; Distracting Dagger (18392, -6.75 DPS) [dungeon]; Scepter of Interminable Focus (22329, -58.40 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (265.4 DPS) | yes | Precisely Calibrated Boomstick (2100, -0.84 DPS) [world_drop]; The Purifier (22656, -0.93 DPS) [quest]; Dark Iron Rifle (16004, -2.31 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1372, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 629.2. Weights run: 3.7s. Verify run: 10.3s. 1372 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.309 ± 0.024, crit=2.291 ± 0.040 per rating point (14 rating = 1%, 32.072 per %), hit=4.258 ± 0.245 per rating point (10 rating = 1%, 42.584 per %), melee_haste=17.990 ± 3.114

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 96.6 attack_power points (17.68 DPS) | yes | Darkmantle Cap (226829, +0.00 DPS) [quest]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 77.2 attack_power points (14.14 DPS) | yes | Beads of Ogre Might (22150, -1.95 DPS) [quest]; Mark of Fordring (15411, -3.51 DPS) [quest]; Medallion of the Dawn (22659, -3.87 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (629.2 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -33.68 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 70.6 attack_power points (12.92 DPS) | yes | Cape of the Black Baron (13340, -2.92 DPS) [dungeon]; Deathguard's Cloak (20068, -4.58 DPS) [rep]; Windshear Cape (20691, -4.97 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (629.2 DPS) | yes | Darkmantle Tunic (226825, -4.51 DPS) [quest]; Warlord's Leather Breastplate (231549, -4.65 DPS) [pvp]; Tunic of Undead Slaying (23089, -23.73 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (629.2 DPS) | yes | General's Leather Armsplints (16559, -0.20 DPS) [pvp]; Blackmist Armguards (12966, -1.44 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.08 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | 74.7 attack_power points (13.67 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.95 DPS) [quest]; Raider Gloves (272099, -40.12 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 133.7 attack_power points (24.47 DPS) | yes | Belt of Preserved Heads (20216, -7.51 DPS) [quest]; Ferocity of the Timbermaw (227805, -12.05 DPS) [vendor]; Defiler's Leather Girdle (20190, -12.37 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (629.2 DPS) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -34.49 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 85.2 attack_power points (15.60 DPS) | yes | Shadowcraft Boots (16711, -4.39 DPS) [dungeon]; Fine Dawn Treaders (227815, -4.43 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -5.46 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | sim-verified (629.2 DPS) | yes | Tarnished Elven Ring (18500, -3.80 DPS) [dungeon]; Cutthroat's Signet (272408, -4.23 DPS) [vendor]; Naglering (11669, -13.56 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (629.2 DPS) | yes | Tarnished Elven Ring (18500, -2.46 DPS) [dungeon]; Cutthroat's Signet (272408, -2.88 DPS) [vendor]; Naglering (11669, -11.64 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (629.2 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (629.2 DPS) | yes | Blackhand's Breadth (13965, -1.40 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, -3.57 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -6.13 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (629.2 DPS) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (629.2 DPS) | yes | The Lobotomizer (19324, -12.21 DPS, sim-verified) [rep]; Distracting Dagger (18392, -12.35 DPS) [dungeon]; Scepter of Interminable Focus (22329, -103.17 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (629.2 DPS) | yes | Precisely Calibrated Boomstick (2100, -3.15 DPS) [world_drop]; The Purifier (22656, -3.19 DPS) [quest]; Dark Iron Rifle (16004, -4.43 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1372, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

