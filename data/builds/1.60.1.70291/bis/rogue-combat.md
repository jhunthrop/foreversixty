# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.9. Weights run: 1.3s. Verify run: 4.2s. 191 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.933 ± 0.020, crit=0.581 ± 0.018 per rating point (14 rating = 1%, 8.132 per %), hit=0.743 ± 0.044 per rating point (10 rating = 1%, 7.426 per %), melee_haste=3.791 ± 0.494

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.5 attack_power points (0.99 DPS) | yes | Defender's Leather Hood (252447, -0.50 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 11.6 attack_power points (0.74 DPS) | yes | Erudite's Amulet (277204, -0.52 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.7 attack_power points (0.62 DPS) | yes | Slime-encrusted Pads (6461, -0.64 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.6 attack_power points (0.74 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 21.3 attack_power points (1.36 DPS) | yes | Brawler's Leather Armor (252490, -0.17 DPS) [crafted]; Prospector's Chestpiece (14562, -0.49 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.49 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.7 attack_power points (0.62 DPS) | yes | Forest Leather Bracers (3202, -0.00 DPS) [world_drop]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.9 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.92 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.15 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Dusty Belt (279897, -0.53 DPS) [quest]; Deviate Scale Belt (6468, -3.97 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.18 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.11 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.7 attack_power points (0.75 DPS) | yes | Pyrewood Signet Ring (277210, -0.07 DPS) [quest]; Demon Band (12054, -0.49 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 11.6 attack_power points (0.74 DPS) | yes | Pyrewood Signet Ring (277210, -0.27 DPS, sim-verified) [quest]; Demon Band (12054, -0.48 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.58 DPS) | yes | Evocator's Blade (2567, -0.91 DPS) [dungeon]; Buzzer Blade (2169, -2.08 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.68 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 228.2 attack_power points (14.55 DPS) | yes | Evocator's Blade (2567, -4.86 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.7 attack_power points (0.49 DPS) | yes | Fine Longbow (11304, -0.24 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 70.9. Weights run: 1.2s. Verify run: 4.2s. 347 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.547 ± 0.018, crit=0.567 ± 0.019 per rating point (14 rating = 1%, 7.936 per %), hit=0.905 ± 0.051 per rating point (10 rating = 1%, 9.049 per %), melee_haste=7.066 ± 0.620

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.5 attack_power points (0.97 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.22 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.87 DPS) | yes | Sentinel's Medallion (19541, -0.10 DPS) [rep]; Kaleidoscope Chain (13084, -0.24 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.0 attack_power points (1.38 DPS) | yes | Mantle of Thieves (2264, -0.49 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS) [crafted]; Bristlebark Amice (14573, -0.61 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.8 attack_power points (0.86 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.7 attack_power points (1.35 DPS) | yes | Tunic of Westfall (2041, -0.29 DPS) [quest]; Brawler's Leather Tunic (252508, -0.33 DPS, sim-verified) [crafted]; Raptorbane Armor (3566, -0.35 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.3 attack_power points (0.83 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.21 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.9 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.10 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (70.9 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.46 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (70.9 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.23 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (70.9 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.44 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 17.9 attack_power points (1.12 DPS) | yes | Thunderbrow Ring (13097, -0.33 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Pyrewood Signet Ring (277210, -0.51 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.3 attack_power points (0.95 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.28 DPS) [quest]; Pyrewood Signet Ring (277210, -0.34 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.11 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.10 DPS, sim-verified) [vendor]; Thornspike (6681, -2.11 DPS) [dungeon]; Toxic Revenger (9453, -2.16 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (70.9 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -7.20 DPS, sim-verified) [vendor] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.75 DPS) | yes | Golemsight Long Gun (273029, -0.17 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.19 DPS) [vendor]; Silver Star (3463, -0.27 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Glass Shooter

No-known-source sample (15 of 347, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 116.4. Weights run: 1.2s. Verify run: 2.5s. 466 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.393 ± 0.019, crit=0.513 ± 0.026 per rating point (14 rating = 1%, 7.183 per %), hit=1.048 ± 0.054 per rating point (10 rating = 1%, 10.475 per %), melee_haste=8.345 ± 0.768

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.3 attack_power points (1.79 DPS) | yes | Warden's Wizard Hat (14604, -0.56 DPS) [world_drop]; Hawkeye's Helm (14591, -0.58 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.65 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 15.3 attack_power points (1.04 DPS) | yes | Ghostshard Talisman (7731, -0.09 DPS) [dungeon]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Fallen Guard's Pendant (279837, -0.63 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.3 attack_power points (1.85 DPS) | yes | Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.58 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.9 attack_power points (1.22 DPS) | yes | Sergeant Major's Cape (16336, -0.24 DPS) [pvp]; Hawkeye's Cloak (14593, -0.35 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.45 DPS) [quest] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 23.7 attack_power points (1.61 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.38 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Hawkeye's Bracers (14590, -0.48 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.60 DPS) [dungeon]; Dusky Bracers (7378, -0.60 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 27.2 attack_power points (1.84 DPS) | yes | Prowler's Leather Gloves (252524, -0.32 DPS) [crafted]; Imperial Leather Gloves (4063, -0.38 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.52 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.04 DPS) | yes | Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.57 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.78 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 29.3 attack_power points (1.99 DPS) | yes | Triprunner Dungarees (9624, -0.08 DPS) [quest]; Ferine Leggings (6690, -0.22 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.66 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.3 attack_power points (1.51 DPS) | yes | Prowler's Leather Shoes (252465, -0.11 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Protector's Band (19515, -0.06 DPS) [rep]; Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Ring of the Underwood (2951, -0.28 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.36 DPS) | yes | Protector's Band (19515, -0.06 DPS) [rep]; Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Ring of the Underwood (2951, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (116.4 DPS) | yes | Black Menace (6831, -4.01 DPS) [quest]; Coldrage Dagger (10761, -4.66 DPS, sim-verified) [dungeon]; Darkspear Insurgent's Spellblade (272085, -4.71 DPS) [vendor] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 442.2 attack_power points (30.00 DPS) | yes | Black Menace (6831, -0.85 DPS, sim-verified) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (116.4 DPS) | yes | Glass Shooter (9456, -0.14 DPS) [dungeon]; Monolithic Bow (9426, -0.26 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.15 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 466, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 161.0. Weights run: 1.3s. Verify run: 2.9s. 586 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.562 ± 0.020, crit=0.942 ± 0.030 per rating point (14 rating = 1%, 13.183 per %), hit=1.301 ± 0.086 per rating point (10 rating = 1%, 13.012 per %), melee_haste=5.339 ± 1.244

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 42.2 attack_power points (3.59 DPS) | yes | White Bandit Mask (10008, -1.30 DPS, sim-verified) [crafted]; Helm of Fire (8348, -1.33 DPS) [crafted]; Lordrec Helmet (10741, -1.46 DPS) [quest] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.3 attack_power points (2.15 DPS) | yes | Sentinel's Medallion (19539, -0.60 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.96 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 34.2 attack_power points (2.91 DPS) | yes | Sunburn Spaulders (274751, -0.43 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.68 DPS) [crafted]; Failed Flying Experiment (9647, -0.73 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.4 attack_power points (1.99 DPS) | yes | Dark Hooded Cape (5257, -0.32 DPS) [world]; Blackflame Cape (13109, -0.40 DPS) [world_drop]; Serpentskin Cloak (8259, -0.53 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 44.2 attack_power points (3.76 DPS) | yes | Warbear Harness (15064, -0.43 DPS) [crafted]; Blazewind Breastplate (11193, -0.45 DPS) [quest]; Charred Leather Tunic (19127, -1.37 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.4 attack_power points (2.33 DPS) | yes | Skulker's Leather Bracers (252540, -0.57 DPS, sim-verified) [crafted]; Pridelord Bands (14672, -0.58 DPS) [world_drop]; Branded Leather Bracers (19508, -0.63 DPS) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 44.4 attack_power points (3.78 DPS) | yes | Darkmantle Grips (226828, -0.09 DPS) [vendor]; Gloves of Holy Might (867, -0.95 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.15 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 33.9 attack_power points (2.88 DPS) | yes | Highlander's Leather Girdle (20115, -0.06 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 44.2 attack_power points (3.76 DPS) | yes | Gryphon Rider's Leggings (9652, -0.49 DPS) [quest]; Serpentskin Leggings (8262, -0.53 DPS) [world_drop]; Basilisk Hide Pants (1718, -0.97 DPS) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 32.5 attack_power points (2.77 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Albino Crocscale Boots (17728, -0.11 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.20 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 24.1 attack_power points (2.05 DPS) | yes | Mark of Kern (2262, -0.35 DPS) [dungeon]; Assault Band (13095, -0.35 DPS) [world_drop]; Blackstone Ring (17713, -0.35 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 21.9 attack_power points (1.86 DPS) | yes | Mark of Kern (2262, -0.16 DPS) [dungeon]; Assault Band (13095, -0.16 DPS) [world_drop]; Blackstone Ring (17713, -0.16 DPS) [dungeon] |
| trinket1 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (161.0 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (161.0 DPS) | yes | Mark of the Chosen (17774, -1.09 DPS, sim-verified) [quest] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (161.0 DPS) | yes | Charstone Dirk (17710, -0.71 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, -1.84 DPS) [vendor]; Gut Ripper (2164, -3.16 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.83 DPS) | yes | Charstone Dirk (17710, -29.67 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (161.0 DPS) | yes | Skull Splitting Crossbow (13039, -0.41 DPS) [world_drop]; Guttbuster (13139, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.05 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Protector's Band; finger2: Masons Fraternity Ring; trinket1: Molten Heart of the Mountain; trinket2: Frozen Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 586, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 258.8. Weights run: 1.4s. Verify run: 8.4s. 1136 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.772 ± 0.028, crit=1.595 ± 0.049 per rating point (14 rating = 1%, 22.334 per %), hit=2.223 ± 0.129 per rating point (10 rating = 1%, 22.226 per %), melee_haste=14.292 ± 1.920

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 70.7 attack_power points (5.71 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.33 DPS) [quest] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 46.3 attack_power points (3.74 DPS) | yes | Beads of Ogre Might (22150, -0.01 DPS) [quest]; Amulet of the Darkmoon (19491, -0.22 DPS) [quest]; Zandalarian Shadow Talisman (19616, -1.11 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (258.8 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -9.45 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 50.2 attack_power points (4.06 DPS) | yes | Cape of the Black Baron (13340, -0.29 DPS) [dungeon]; Cloak of the Honor Guard (20073, -0.59 DPS) [rep]; Shifting Cloak (18511, -1.62 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 110.1 attack_power points (8.89 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -1.57 DPS) [pvp]; Timbermaw Tunic (252484, -1.85 DPS) [crafted]; Darkmantle Tunic (226825, -1.93 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 44.7 attack_power points (3.61 DPS) | yes | Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Bracers of Subterfuge (22668, -0.65 DPS) [quest]; Deepfury Bracers (13120, -1.14 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (258.8 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, -13.06 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 81.7 attack_power points (6.59 DPS) | yes | Belt of Preserved Heads (20216, -1.52 DPS) [quest]; Highlander's Leather Girdle (20045, -2.04 DPS) [rep]; Ferocity of the Timbermaw (227805, -2.06 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (258.8 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -9.45 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 58.1 attack_power points (4.69 DPS) | yes | Shadowcraft Boots (16711, -1.15 DPS) [dungeon]; Drudge Boots (21532, -1.21 DPS) [quest]; Dunestalker's Boots (20715, -1.24 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 64.7 attack_power points (5.23 DPS) | yes | Cutthroat's Signet (272408, -1.43 DPS) [vendor]; Protector's Band (19514, -2.69 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 60.6 attack_power points (4.89 DPS) | yes | Cutthroat's Signet (272408, -1.09 DPS, sim-verified) [vendor]; Protector's Band (19514, -2.35 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (258.8 DPS) | yes | Frozen Heart of the Mountain (249469, -5.69 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (258.8 DPS) | yes | Darkmoon Card: Blue Dragon (19288, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (258.8 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -14.85 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (53.38 DPS) | yes | Electrified Dagger (19100, -1.12 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (258.8 DPS) | yes | The Purifier (22656, -0.20 DPS) [quest]; Bloodseeker (19107, -0.36 DPS) [quest]; Dark Iron Rifle (16004, -2.21 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Medallion of the Dawn; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Molten Heart of the Mountain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1136, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 706.9. Weights run: 1.4s. Verify run: 8.6s. 1136 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.100 ± 0.002, agility=2.080 ± 0.033, crit=1.898 ± 0.057 per rating point (14 rating = 1%, 26.572 per %), hit=3.009 ± 0.181 per rating point (10 rating = 1%, 30.093 per %), melee_haste=23.498 ± 2.451

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 83.5 attack_power points (14.20 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.84 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 54.1 attack_power points (9.20 DPS) | yes | Medallion of the Dawn (22659, -0.60 DPS) [quest]; Amulet of the Darkmoon (19491, -0.61 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.77 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (706.9 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -37.16 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 58.1 attack_power points (9.88 DPS) | yes | Cape of the Black Baron (13340, -1.17 DPS) [dungeon]; Cloak of the Honor Guard (20073, -2.33 DPS) [rep]; Shifting Cloak (18511, -3.87 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 133.6 attack_power points (22.72 DPS) | yes | Field Marshal's Leather Chestpiece (231543, -3.89 DPS) [pvp]; Darkmantle Tunic (226825, -5.40 DPS, sim-verified) [quest]; Timbermaw Tunic (252484, -5.66 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 51.6 attack_power points (8.77 DPS) | yes | Marshal's Leather Armsplints (16460, -0.19 DPS) [pvp]; Bracers of Subterfuge (22668, -1.60 DPS) [quest]; Deepfury Bracers (13120, -2.72 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (706.9 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, -42.56 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 103.9 attack_power points (17.66 DPS) | yes | Belt of Preserved Heads (20216, -4.32 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -6.78 DPS) [vendor]; Highlander's Leather Girdle (20045, -7.36 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (706.9 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -37.16 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 71.0 attack_power points (12.07 DPS) | yes | Shadowcraft Boots (16711, -3.11 DPS) [dungeon]; Fine Dawn Treaders (227815, -3.49 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -3.58 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 80.0 attack_power points (13.60 DPS) | yes | Cutthroat's Signet (272408, -3.54 DPS) [vendor]; Protector's Band (19514, -7.47 DPS) [rep]; Channeler's Ring (272406, -8.49 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 72.7 attack_power points (12.35 DPS) | yes | Cutthroat's Signet (272408, -2.29 DPS) [vendor]; Protector's Band (19514, -6.22 DPS) [rep]; Channeler's Ring (272406, -7.24 DPS) [vendor] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (706.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (706.9 DPS) | yes | Molten Heart of the Mountain (249470, -3.78 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (706.9 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -37.66 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (112.39 DPS) | yes | Electrified Dagger (19100, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (706.9 DPS) | yes | The Purifier (22656, -0.43 DPS) [quest]; Bloodseeker (19107, -0.98 DPS) [quest]; Dark Iron Rifle (16004, -6.25 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1136, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.2. Weights run: 1.3s. Verify run: 4.2s. 184 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.933 ± 0.020, crit=0.581 ± 0.018 per rating point (14 rating = 1%, 8.132 per %), hit=0.743 ± 0.044 per rating point (10 rating = 1%, 7.426 per %), melee_haste=3.791 ± 0.494

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 15.5 attack_power points (0.99 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 11.6 attack_power points (0.74 DPS) | yes | Erudite's Amulet (277204, -0.48 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 9.7 attack_power points (0.62 DPS) | yes | Slime-encrusted Pads (6461, -0.62 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 11.6 attack_power points (0.74 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.24 DPS) [crafted]; Bristlebark Cape (14571, -0.25 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (41.2 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.81 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 9.7 attack_power points (0.62 DPS) | yes | Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Ratchet Wristwraps (274742, -0.25 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 15.6 attack_power points (0.99 DPS) | yes | Brawler's Leather Gloves (252494, -0.25 DPS) [crafted]; Bristlebark Gloves (14572, -0.25 DPS, sim-verified) [world_drop]; Serpent Gloves (5970, -0.26 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.15 DPS) | yes | Brawler's Leather Belt (252428, -0.40 DPS) [crafted]; Murloc Scale Belt (5780, -0.65 DPS) [crafted]; Deviate Scale Belt (6468, -3.79 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.2 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.01 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.2 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.92 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.7 attack_power points (0.75 DPS) | yes | Pyrewood Signet Ring (277210, -0.07 DPS) [quest]; Bounty Hunter's Ring (5351, -0.38 DPS) [quest]; Demon Band (12054, -0.49 DPS) [world_drop] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 11.6 attack_power points (0.74 DPS) | yes | Pyrewood Signet Ring (277210, -0.25 DPS, sim-verified) [quest]; Bounty Hunter's Ring (5351, -0.37 DPS) [quest]; Demon Band (12054, -0.48 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.58 DPS) | yes | Scout's Blade (20441, -0.44 DPS) [pvp]; Evocator's Blade (2567, -0.91 DPS) [dungeon]; Buzzer Blade (2169, -2.08 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 228.2 attack_power points (14.55 DPS) | yes | Evocator's Blade (2567, -4.58 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.7 attack_power points (0.49 DPS) | yes | Fine Longbow (11304, -0.21 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.25 DPS) [crafted]; Light Bow (4576, -0.25 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 184, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 254779 A'sharahm, the Roiling Tempest; 263006 Scout Ranger's Tunic

### Band 30 (troll, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 71.1. Weights run: 1.2s. Verify run: 4.2s. 337 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.547 ± 0.018, crit=0.567 ± 0.019 per rating point (14 rating = 1%, 7.936 per %), hit=0.905 ± 0.051 per rating point (10 rating = 1%, 9.049 per %), melee_haste=7.066 ± 0.620

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 15.5 attack_power points (0.97 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.22 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.87 DPS) | yes | Scout's Medallion (19537, -0.10 DPS) [rep]; Kaleidoscope Chain (13084, -0.24 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.0 attack_power points (1.38 DPS) | yes | Mantle of Thieves (2264, -0.49 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS) [crafted]; Bristlebark Amice (14573, -0.61 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.8 attack_power points (0.86 DPS) | yes | Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon]; Wildhunter Cloak (16658, -0.24 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 21.7 attack_power points (1.35 DPS) | yes | Brawler's Leather Tunic (252508, -0.32 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.36 DPS) [quest]; Brawler's Leather Armor (252490, -0.36 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.3 attack_power points (0.83 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.21 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (71.1 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.45 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (71.1 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -2.87 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (71.1 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.61 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (71.1 DPS) | yes | Feet of the Lynx (1121, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Vorrel's Boots (7751, -2.77 DPS, sim-verified) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 17.9 attack_power points (1.12 DPS) | yes | Thunderbrow Ring (13097, -0.33 DPS) [world_drop]; Monkey Ring (6748, -0.44 DPS) [quest]; Band of the Fist (17694, -0.48 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.3 attack_power points (0.95 DPS) | yes | Thunderbrow Ring (13097, -0.16 DPS) [world_drop]; Monkey Ring (6748, -0.28 DPS) [quest]; Band of the Fist (17694, -0.32 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.11 DPS) | yes | Scout's Blade (19545, -0.91 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.37 DPS) [vendor]; Torturing Poker (7682, -1.59 DPS) [dungeon] |
| off_hand | Serrated Raptor Claw (280805) | Changing Tastes [quest] | 319.6 attack_power points (19.96 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -7.78 DPS, sim-verified) [vendor] |
| ranged | Glass Shooter (9456) | Gnomeregan: Dark Iron Ambassador [dungeon] | 12.0 attack_power points (0.75 DPS) | yes | Golemsight Long Gun (273029, -0.17 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.19 DPS) [vendor]; Silver Star (3463, -0.27 DPS) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Serrated Raptor Claw; ranged: Glass Shooter

No-known-source sample (15 of 337, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 115.6. Weights run: 1.2s. Verify run: 2.5s. 449 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.393 ± 0.019, crit=0.513 ± 0.026 per rating point (14 rating = 1%, 7.183 per %), hit=1.048 ± 0.054 per rating point (10 rating = 1%, 10.475 per %), melee_haste=8.345 ± 0.768

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.3 attack_power points (1.79 DPS) | yes | Warden's Wizard Hat (14604, -0.56 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.65 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 15.3 attack_power points (1.04 DPS) | yes | Ethereal Talisman (4430, -0.32 DPS) [quest]; Kaleidoscope Chain (13084, -0.39 DPS) [world_drop]; Ghostshard Talisman (7731, -0.49 DPS, sim-verified) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.3 attack_power points (1.85 DPS) | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Nightscape Shoulders (8192, -0.81 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 17.9 attack_power points (1.22 DPS) | yes | Imperial Cloak (6432, -0.46 DPS) [dungeon]; Parachute Cloak (10518, -0.46 DPS) [crafted]; Hawkeye's Cloak (14593, -0.53 DPS, sim-verified) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 23.7 attack_power points (1.61 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.38 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.52 DPS) [world_drop]; Dusky Bracers (7378, -0.60 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 27.2 attack_power points (1.84 DPS) | yes | Prowler's Leather Gloves (252524, -0.32 DPS) [crafted]; Imperial Leather Gloves (4063, -0.38 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.63 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.04 DPS) | yes | Defiler's Chain Girdle (20152, -0.49 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.57 DPS) [world_drop]; Skulker's Leather Belt (252520, -0.78 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 29.3 attack_power points (1.99 DPS) | yes | Triprunner Dungarees (9624, -0.08 DPS) [quest]; Ferine Leggings (6690, -0.22 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.66 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 22.3 attack_power points (1.51 DPS) | yes | Prowler's Leather Shoes (252465, -0.11 DPS) [crafted]; Imperial Leather Boots (6431, -0.14 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.36 DPS) | yes | Legionnaire's Band (19512, -0.06 DPS) [rep]; Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Ring of the Underwood (2951, -0.28 DPS) [world_drop] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.36 DPS) | yes | Legionnaire's Band (19512, -0.06 DPS) [rep]; Ironspine's Eye (7686, -0.23 DPS) [dungeon]; Ring of the Underwood (2951, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (115.6 DPS) | yes | Scout's Blade (19544, -3.77 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -4.71 DPS) [vendor]; Coldrage Dagger (10761, -5.07 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 442.2 attack_power points (30.00 DPS) | yes | Coldrage Dagger (10761, -0.61 DPS, sim-verified) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (115.6 DPS) | yes | Glass Shooter (9456, -0.14 DPS) [dungeon]; Monolithic Bow (9426, -0.26 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.14 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 449, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 162.5. Weights run: 1.3s. Verify run: 2.9s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.562 ± 0.020, crit=0.942 ± 0.030 per rating point (14 rating = 1%, 13.183 per %), hit=1.301 ± 0.086 per rating point (10 rating = 1%, 13.012 per %), melee_haste=5.339 ± 1.244

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.2 attack_power points (2.40 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.09 DPS) [quest]; Helm of Fire (8348, -0.14 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 25.3 attack_power points (2.15 DPS) | yes | Woven Ivy Necklace (19159, -0.45 DPS) [quest]; Scout's Medallion (19535, -0.56 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.2 attack_power points (2.48 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.26 DPS) [crafted]; Failed Flying Experiment (9647, -0.30 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 23.4 attack_power points (1.99 DPS) | yes | Dark Hooded Cape (5257, -0.32 DPS) [world]; Blackflame Cape (13109, -0.40 DPS) [world_drop]; Serpentskin Cloak (8259, -0.53 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 44.2 attack_power points (3.76 DPS) | yes | Blazewind Breastplate (11193, -0.45 DPS) [quest]; Warbear Harness (15064, -0.98 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -1.37 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.4 attack_power points (2.33 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.59 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 44.4 attack_power points (3.78 DPS) | yes | Darkmantle Grips (226828, -0.09 DPS) [vendor]; Gloves of Holy Might (867, -0.95 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.15 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 33.9 attack_power points (2.88 DPS) | yes | Defiler's Leather Girdle (20193, -0.06 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.10 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 44.2 attack_power points (3.76 DPS) | yes | Basilisk Hide Pants (1718, -0.97 DPS) [world_drop]; Serpentskin Leggings (8262, -1.10 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -1.11 DPS) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 32.5 attack_power points (2.77 DPS) | yes | Skulker's Leather Boots (252469, -0.11 DPS) [crafted]; Albino Crocscale Boots (17728, -0.11 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.20 DPS) [crafted] |
| finger1 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 24.1 attack_power points (2.05 DPS) | yes | Masons Fraternity Ring (9533, -0.19 DPS) [quest]; Mark of Kern (2262, -0.35 DPS) [dungeon]; Blackstone Ring (17713, -0.35 DPS) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.04 DPS) | yes | Masons Fraternity Ring (9533, -0.18 DPS) [quest]; Mark of Kern (2262, -0.34 DPS) [dungeon]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (162.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (162.5 DPS) | yes | Molten Heart of the Mountain (249470, -0.88 DPS, sim-verified) [crafted] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | sim-verified (162.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Scout's Blade (19543, -0.37 DPS) [pvp]; Charstone Dirk (17710, -0.71 DPS) [dungeon] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.83 DPS) | yes | Charstone Dirk (17710, -29.12 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (162.5 DPS) | yes | Skull Splitting Crossbow (13039, -0.41 DPS) [world_drop]; Guttbuster (13139, -0.54 DPS) [world_drop]; Dark Iron Rifle (16004, -2.10 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 263.1. Weights run: 1.4s. Verify run: 8.4s. 1132 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.000 ± 0.002, agility=1.772 ± 0.028, crit=1.595 ± 0.049 per rating point (14 rating = 1%, 22.334 per %), hit=2.223 ± 0.129 per rating point (10 rating = 1%, 22.226 per %), melee_haste=14.292 ± 1.920

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 70.7 attack_power points (5.71 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.33 DPS) [quest] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 46.3 attack_power points (3.74 DPS) | yes | Beads of Ogre Might (22150, -0.01 DPS) [quest]; Amulet of the Darkmoon (19491, -0.22 DPS) [quest]; Zandalarian Shadow Talisman (19616, -1.11 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (263.1 DPS) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -10.11 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 50.2 attack_power points (4.06 DPS) | yes | Cape of the Black Baron (13340, -0.29 DPS) [dungeon]; Deathguard's Cloak (20068, -0.59 DPS) [rep]; Shifting Cloak (18511, -1.62 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 110.1 attack_power points (8.89 DPS) | yes | Warlord's Leather Breastplate (231549, -1.57 DPS) [pvp]; Timbermaw Tunic (252484, -1.85 DPS) [crafted]; Darkmantle Tunic (226825, -1.85 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 44.7 attack_power points (3.61 DPS) | yes | General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Bracers of Subterfuge (22668, -0.65 DPS) [quest]; Deepfury Bracers (13120, -1.14 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (263.1 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -13.72 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 81.7 attack_power points (6.59 DPS) | yes | Belt of Preserved Heads (20216, -1.96 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -2.04 DPS) [rep]; Ferocity of the Timbermaw (227805, -2.06 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (263.1 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -10.11 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 58.1 attack_power points (4.69 DPS) | yes | Drudge Boots (21532, -1.21 DPS) [quest]; Dunestalker's Boots (20715, -1.24 DPS) [quest]; Shadowcraft Boots (16711, -1.64 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 64.7 attack_power points (5.23 DPS) | yes | Cutthroat's Signet (272408, -1.43 DPS) [vendor]; Legionnaire's Band (19510, -2.69 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 60.6 attack_power points (4.89 DPS) | yes | Cutthroat's Signet (272408, -1.10 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -2.35 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (263.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (263.1 DPS) | yes | Frozen Heart of the Mountain (249469, -3.03 DPS) [crafted]; Molten Heart of the Mountain (249470, -5.52 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (263.1 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -16.39 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (53.38 DPS) | yes | Glacial Blade (19099, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (263.1 DPS) | yes | The Purifier (22656, -0.20 DPS) [quest]; Bloodseeker (19107, -0.36 DPS) [quest]; Dark Iron Rifle (16004, -2.27 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Medallion of the Dawn; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 707.6. Weights run: 1.4s. Verify run: 8.7s. 1132 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.002, strength=1.100 ± 0.002, agility=2.080 ± 0.033, crit=1.898 ± 0.057 per rating point (14 rating = 1%, 26.572 per %), hit=3.009 ± 0.181 per rating point (10 rating = 1%, 30.093 per %), melee_haste=23.498 ± 2.451

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 83.5 attack_power points (14.20 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.84 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 54.1 attack_power points (9.20 DPS) | yes | Medallion of the Dawn (22659, +0.00 DPS, sim-verified) [quest]; Amulet of the Darkmoon (19491, -0.61 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.77 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (707.6 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -36.60 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 58.1 attack_power points (9.88 DPS) | yes | Cape of the Black Baron (13340, +0.00 DPS, sim-verified) [dungeon]; Deathguard's Cloak (20068, -2.33 DPS) [rep]; Shifting Cloak (18511, -3.87 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 133.6 attack_power points (22.72 DPS) | yes | Darkmantle Tunic (226825, -3.22 DPS, sim-verified) [quest]; Warlord's Leather Breastplate (231549, -3.89 DPS) [pvp]; Timbermaw Tunic (252484, -5.66 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 51.6 attack_power points (8.77 DPS) | yes | General's Leather Armsplints (16559, -0.19 DPS) [pvp]; Bracers of Subterfuge (22668, -1.60 DPS) [quest]; Deepfury Bracers (13120, -2.72 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (707.6 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -42.06 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 103.9 attack_power points (17.66 DPS) | yes | Belt of Preserved Heads (20216, -4.62 DPS) [quest]; Ferocity of the Timbermaw (227805, -6.78 DPS) [vendor]; Defiler's Leather Girdle (20190, -7.36 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (707.6 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -36.60 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 71.0 attack_power points (12.07 DPS) | yes | Fine Dawn Treaders (227815, -2.99 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -3.11 DPS) [dungeon]; Darkmantle Boots (22003, -3.58 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 80.0 attack_power points (13.60 DPS) | yes | Cutthroat's Signet (272408, -3.54 DPS) [vendor]; Legionnaire's Band (19510, -7.47 DPS) [rep]; Channeler's Ring (272406, -8.49 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 72.7 attack_power points (12.35 DPS) | yes | Cutthroat's Signet (272408, -2.29 DPS) [vendor]; Legionnaire's Band (19510, -6.22 DPS) [rep]; Channeler's Ring (272406, -7.24 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (707.6 DPS) | yes | Frozen Heart of the Mountain (249469, -6.55 DPS, sim-verified) [crafted] |
| trinket2 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (707.6 DPS) | yes | Frozen Heart of the Mountain (249469, -5.75 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (707.6 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -35.21 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (112.39 DPS) | yes | Glacial Blade (19099, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (707.6 DPS) | yes | The Purifier (22656, -0.43 DPS) [quest]; Bloodseeker (19107, -0.98 DPS) [quest]; Dark Iron Rifle (16004, -6.09 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket2: Darkmoon Card: Maelstrom; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1132, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

