# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 37.5. Weights run: 2.3s. Verify run: 4.0s. 198 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.790 ± 0.013, crit=0.510 ± 0.012 per rating point (14 rating = 1%, 7.141 per %), hit=0.659 ± 0.029 per rating point (10 rating = 1%, 6.591 per %), melee_haste=3.047 ± 0.329

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.3 attack_power points (0.92 DPS) | yes | Defender's Leather Hood (252447, -0.36 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 10.7 attack_power points (0.69 DPS) | yes | Erudite's Amulet (277204, -0.46 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.57 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 10.7 attack_power points (0.69 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.22 DPS) [crafted]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 19.7 attack_power points (1.27 DPS) | yes | Brawler's Leather Armor (252490, -0.14 DPS) [crafted]; Prospector's Chestpiece (14562, -0.45 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.46 DPS) [crafted] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 9.2 attack_power points (0.59 DPS) | yes | Forest Leather Bracers (3202, -0.01 DPS) [world_drop]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Wolf Bracers (4794, -0.13 DPS) [vendor] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.5 DPS) | yes | Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Brawler's Leather Gloves (252494, +0.00 DPS) [crafted]; Gloves of the Fang (10413, -2.79 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.44 DPS) [crafted]; Dusty Belt (279897, -0.58 DPS) [quest]; Deviate Scale Belt (6468, -3.80 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.5 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.01 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (37.5 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.91 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 11.2 attack_power points (0.72 DPS) | yes | Demon Band (12054, -0.46 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon]; The 1 Ring (8350, -0.54 DPS) [world] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.8 attack_power points (0.63 DPS) | yes | Demon Band (12054, -0.19 DPS, sim-verified) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon]; The 1 Ring (8350, -0.45 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.71 DPS) | yes | Evocator's Blade (2567, -0.92 DPS) [dungeon]; Buzzer Blade (2169, -2.10 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.74 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.6 attack_power points (14.64 DPS) | yes | Evocator's Blade (2567, -4.06 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.2 attack_power points (0.46 DPS) | yes | Fine Longbow (11304, -0.19 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.23 DPS) [crafted]; Light Bow (4576, -0.23 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Bravo's Armbands; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Protector's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 198, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5821 Darkstalker Boots; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18849 Insignia of the Horde

### Band 30 (night-elf, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 58.8. Weights run: 2.5s. Verify run: 4.4s. 366 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.452 ± 0.011, crit=0.479 ± 0.012 per rating point (14 rating = 1%, 6.710 per %), hit=0.911 ± 0.035 per rating point (10 rating = 1%, 9.108 per %), melee_haste=7.054 ± 0.404

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 14.5 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world]; Defender's Leather Helm (252455, -0.30 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.78 DPS) | yes | Sentinel's Medallion (19541, -0.02 DPS) [rep]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.44 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.57 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.2 attack_power points (0.85 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon]; Sergeant Major's Cape (16315, -0.22 DPS) [pvp] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 20.3 attack_power points (1.32 DPS) | yes | Brawler's Leather Tunic (252508, -0.25 DPS, sim-verified) [crafted]; Raptorbane Armor (3566, -0.28 DPS) [quest]; Tunic of Westfall (2041, -0.28 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.7 attack_power points (0.83 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Demonhide Bracers (270033, -0.25 DPS) [quest]; Cultist's Armguards (270032, -0.25 DPS, sim-verified) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.8 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -2.93 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (58.8 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.20 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.8 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.90 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (58.8 DPS) | yes | Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Disjointed Shoes (277226, +0.00 DPS) [quest]; Feet of the Lynx (1121, -3.17 DPS, sim-verified) [world_drop] |
| finger1 | Protector's Band (19517) | Silverwing Sentinels [rep] | 14.7 attack_power points (0.96 DPS) | yes | Pyrewood Signet Ring (277210, -0.34 DPS) [quest]; Insurgent's Band (272067, -0.37 DPS) [vendor]; Ring of Precision (1491, -0.39 DPS) [dungeon] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.4 attack_power points (0.80 DPS) | yes | Insurgent's Band (272067, -0.22 DPS) [vendor]; Ring of Precision (1491, -0.24 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.37 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.91 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -0.73 DPS, sim-verified) [vendor]; Thornspike (6681, -2.19 DPS) [dungeon]; Toxic Revenger (9453, -2.25 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (58.8 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -4.62 DPS, sim-verified) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.58 DPS) | yes | Golemsight Long Gun (273029, -0.02 DPS) [dungeon]; Gravespike Repeater (273033, -0.06 DPS) [world]; Double-barreled Shotgun (2098, -0.11 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Protector's Band; finger2: Thunderbrow Ring; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 99.5. Weights run: 2.7s. Verify run: 2.6s. 493 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.287 ± 0.009, crit=0.369 ± 0.011 per rating point (14 rating = 1%, 5.165 per %), hit=0.968 ± 0.037 per rating point (10 rating = 1%, 9.679 per %), melee_haste=4.905 ± 0.444

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.2 attack_power points (1.86 DPS) | yes | Hawkeye's Helm (14591, -0.60 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.62 DPS) [world_drop]; Nightscape Headband (8176, -0.72 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 14.2 attack_power points (1.04 DPS) | yes | Kaleidoscope Chain (13084, -0.37 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.48 DPS, sim-verified) [world]; River Pride Choker (13087, -0.75 DPS) [world_drop] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.2 attack_power points (1.93 DPS) | yes | Forest Tracker Epaulets (2278, -0.60 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.61 DPS) [dungeon]; Nightscape Shoulders (8192, -0.89 DPS) [crafted] |
| back | Sergeant Major's Cape (16336) | PvP rank 9 · Sergeant Major · Alliance [pvp] | 13.7 attack_power points (1.01 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS) [world_drop]; Yeti Fur Cloak (2805, -0.22 DPS) [quest]; Parachute Cloak (10518, -0.25 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 21.9 attack_power points (1.61 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.34 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Hawkeye's Bracers (14590, -0.54 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.72 DPS) [dungeon]; Dusky Bracers (7378, -0.72 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 25.2 attack_power points (1.86 DPS) | yes | Skulker's Leather Gloves (252525, -0.24 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.26 DPS) [crafted]; Dog Training Gloves (7756, -0.60 DPS, sim-verified) [dungeon] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.21 DPS) | yes | Highlander's Chain Girdle (20090, -0.51 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.70 DPS) [world_drop]; Blackened Defias Belt (10403, -0.89 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.0 attack_power points (1.99 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.62 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.2 attack_power points (1.56 DPS) | yes | Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.48 DPS) | yes | Protector's Band (19515, -0.13 DPS) [rep]; Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (99.5 DPS) | yes | Black Menace (6831, -4.36 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -5.13 DPS) [vendor]; Coldrage Dagger (10761, -8.64 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 441.2 attack_power points (32.57 DPS) | yes | Black Menace (6831, -1.48 DPS, sim-verified) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (99.5 DPS) | yes | Monolithic Bow (9426, -0.31 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.36 DPS) [quest]; Bow of Searing Arrows (2825, -1.20 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Sergeant Major's Cape; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 493, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 50 (night-elf, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 135.2. Weights run: 3.5s. Verify run: 3.1s. 616 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.417 ± 0.007, crit=0.701 ± 0.011 per rating point (14 rating = 1%, 9.811 per %), hit=1.584 ± 0.067 per rating point (10 rating = 1%, 15.837 per %), melee_haste=5.880 ± 0.944

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | 41.6 attack_power points (3.53 DPS) | yes | Helm of Fire (8348, -1.49 DPS) [crafted]; Scorpashi Skullcap (14658, -1.59 DPS) [world_drop]; White Bandit Mask (10008, -1.61 DPS, sim-verified) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.4 attack_power points (1.98 DPS) | yes | Sentinel's Medallion (19539, -0.65 DPS, sim-verified) [rep]; Archmagister's Faceted Pendant (273049, -0.97 DPS) [world] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 33.6 attack_power points (2.85 DPS) | yes | Sunburn Spaulders (274751, -0.72 DPS, sim-verified) [vendor]; Skulker's Leather Shoulder (252535, -0.77 DPS) [crafted]; Failed Flying Experiment (9647, -0.80 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 21.2 attack_power points (1.80 DPS) | yes | Blackflame Cape (13109, -0.36 DPS) [world_drop]; Pridelord Cape (14673, -0.47 DPS) [world_drop]; Serpentskin Cloak (8259, -0.48 DPS) [world_drop] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Blazewind Breastplate (11193, -0.68 DPS) [quest]; Warbear Harness (15064, -0.78 DPS, sim-verified) [crafted]; Charred Leather Tunic (19127, -1.54 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.2 attack_power points (2.14 DPS) | yes | Skulker's Leather Bracers (252540, -0.47 DPS) [crafted]; Pridelord Bands (14672, -0.52 DPS) [world_drop]; Branded Leather Bracers (19508, -0.65 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.2 attack_power points (3.49 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.96 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.04 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 31.8 attack_power points (2.70 DPS) | yes | Prowler's Leather Waistguard (252473, -0.07 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.16 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Serpentskin Leggings (8262, -0.68 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -0.84 DPS, sim-verified) [quest]; Basilisk Hide Pants (1718, -1.18 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 29.4 attack_power points (2.49 DPS) | yes | Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.09 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.58 DPS) [crafted] |
| finger1 | Protector's Band (19516) | Silverwing Sentinels [rep] | 22.7 attack_power points (1.93 DPS) | yes | Mark of Kern (2262, -0.23 DPS) [dungeon]; Assault Band (13095, -0.23 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.25 DPS) [quest] |
| finger2 | Blackstone Ring (17713) (or Mark of Kern (2262), Assault Band (13095)) | Maraudon: Princess Theradras [dungeon] | 20.0 attack_power points (1.69 DPS) | yes | Mark of Kern (2262, +0.00 DPS) [dungeon]; Assault Band (13095, +0.00 DPS) [world_drop]; Masons Fraternity Ring (9533, -0.01 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (135.2 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (135.2 DPS) | yes | Mark of the Chosen (17774, -0.49 DPS, sim-verified) [quest] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (135.2 DPS) | yes | Charstone Dirk (17710, +0.00 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, +0.00 DPS) [vendor]; Julie's Dagger (6660, -6.29 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.60 DPS) | yes | Charstone Dirk (17710, -26.86 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (135.2 DPS) | yes | Skull Splitting Crossbow (13039, -0.24 DPS) [world_drop]; Houndmaster's Bow (11628, -0.41 DPS) [dungeon]; Dark Iron Rifle (16004, -2.12 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Skulker's Leather Boots; finger1: Protector's Band; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 616, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60 (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 255.8. Weights run: 3.4s. Verify run: 9.0s. 1171 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.575 ± 0.012, crit=1.203 ± 0.020 per rating point (14 rating = 1%, 16.841 per %), hit=1.913 ± 0.086 per rating point (10 rating = 1%, 19.127 per %), melee_haste=8.267 ± 1.157

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 58.3 attack_power points (5.89 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.21 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 43.1 attack_power points (4.36 DPS) | yes | Amulet of the Darkmoon (19491, -0.32 DPS) [quest]; Zandalarian Shadow Talisman (19616, -1.36 DPS) [quest]; Medallion of the Dawn (22659, -1.55 DPS, sim-verified) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (255.8 DPS) | yes | Highlander's Leather Shoulders (20059, +0.00 DPS) [rep]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -13.07 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 47.1 attack_power points (4.76 DPS) | yes | Cloak of the Honor Guard (20073, -0.53 DPS) [rep]; Cape of the Black Baron (13340, -1.51 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -2.05 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 94.3 attack_power points (9.52 DPS) | yes | Darkmantle Tunic (226825, -1.69 DPS) [quest]; Field Marshal's Leather Chestpiece (231543, -1.75 DPS) [pvp]; Timbermaw Tunic (252484, -2.72 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 40.9 attack_power points (4.13 DPS) | yes | Marshal's Leather Armsplints (16460, -0.10 DPS) [pvp]; Bracers of Subterfuge (22668, -0.74 DPS) [quest]; Deepfury Bracers (13120, -1.34 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (255.8 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, +0.00 DPS) [pvp]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, -15.61 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 71.3 attack_power points (7.20 DPS) | yes | Ferocity of the Timbermaw (227805, -1.87 DPS) [vendor]; Highlander's Leather Girdle (20045, -2.07 DPS) [rep]; Belt of Preserved Heads (20216, -2.72 DPS, sim-verified) [quest] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (255.8 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -13.07 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 51.2 attack_power points (5.17 DPS) | yes | Dunestalker's Boots (20715, -1.24 DPS) [quest]; Shadowcraft Boots (16711, -1.25 DPS) [dungeon]; Drudge Boots (21532, -2.08 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 56.9 attack_power points (5.75 DPS) | yes | Cutthroat's Signet (272408, -1.59 DPS) [vendor]; Protector's Band (19514, -2.79 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 50.0 attack_power points (5.05 DPS) | yes | Cutthroat's Signet (272408, -0.94 DPS, sim-verified) [vendor]; Protector's Band (19514, -2.08 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (255.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (255.8 DPS) | yes | Molten Heart of the Mountain (249470, -2.43 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (255.8 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -17.40 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (66.76 DPS) | yes | Electrified Dagger (19100, +0.00 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (255.8 DPS) | yes | Bloodseeker (19107, -0.31 DPS) [quest]; Skull Splitting Crossbow (13039, -0.51 DPS) [world_drop]; Dark Iron Rifle (16004, -2.62 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1171, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

### Band 60, raid preset (night-elf, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 606.5. Weights run: 3.7s. Verify run: 9.5s. 1171 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.009 ± 0.017, crit=1.706 ± 0.028 per rating point (14 rating = 1%, 23.888 per %), hit=2.826 ± 0.148 per rating point (10 rating = 1%, 28.260 per %), melee_haste=13.252 ± 1.654

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 78.0 attack_power points (14.30 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.72 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 52.3 attack_power points (9.58 DPS) | yes | Amulet of the Darkmoon (19491, -0.57 DPS) [quest]; Medallion of the Dawn (22659, -0.80 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.85 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (606.5 DPS) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Field Marshal's Leather Epaulets (231547, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -30.64 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 56.3 attack_power points (10.31 DPS) | yes | Cloak of the Honor Guard (20073, -2.24 DPS) [rep]; Cape of the Black Baron (13340, -2.44 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -4.05 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 126.5 attack_power points (23.19 DPS) | yes | Darkmantle Tunic (226825, -2.75 DPS, sim-verified) [quest]; Field Marshal's Leather Chestpiece (231543, -4.05 DPS) [pvp]; Timbermaw Tunic (252484, -5.52 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 50.3 attack_power points (9.22 DPS) | yes | Marshal's Leather Armsplints (16460, -0.20 DPS) [pvp]; Bracers of Subterfuge (22668, -1.67 DPS) [quest]; Deepfury Bracers (13120, -2.88 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (606.5 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Raider Gloves (272099, -38.99 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 98.7 attack_power points (18.10 DPS) | yes | Belt of Preserved Heads (20216, -5.97 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -6.59 DPS) [vendor]; Marshal's Leather Cinch (16458, -7.33 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (606.5 DPS) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -30.64 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 68.0 attack_power points (12.47 DPS) | yes | Fine Dawn Treaders (227815, -3.23 DPS) [vendor]; Drudge Boots (21532, -3.58 DPS) [quest]; Shadowcraft Boots (16711, -3.76 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 76.5 attack_power points (14.02 DPS) | yes | Cutthroat's Signet (272408, -3.68 DPS) [vendor]; Protector's Band (19514, -7.55 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 66.1 attack_power points (12.13 DPS) | yes | Cutthroat's Signet (272408, -1.79 DPS) [vendor]; Protector's Band (19514, -5.65 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (606.5 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (606.5 DPS) | yes | Darkmoon Card: Heroism (19287, -2.76 DPS, sim-verified) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (606.5 DPS) | yes | Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor]; Electrified Dagger (19100, -30.19 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Stormpike Guard [rep] | 661.1 attack_power points (121.19 DPS) | yes | Electrified Dagger (19100, +0.00 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (606.5 DPS) | yes | The Purifier (22656, -0.78 DPS) [quest]; Bloodseeker (19107, -0.96 DPS) [quest]; Dark Iron Rifle (16004, -5.93 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1171, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 37.1. Weights run: 2.3s. Verify run: 4.1s. 191 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.790 ± 0.013, crit=0.510 ± 0.012 per rating point (14 rating = 1%, 7.141 per %), hit=0.659 ± 0.029 per rating point (10 rating = 1%, 6.591 per %), melee_haste=3.047 ± 0.329

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 14.3 attack_power points (0.92 DPS) | yes | Defender's Leather Hood (252447, -0.37 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 10.7 attack_power points (0.69 DPS) | yes | Erudite's Amulet (277204, -0.45 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 8.9 attack_power points (0.58 DPS) | yes | Slime-encrusted Pads (6461, -0.56 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 10.7 attack_power points (0.69 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Dark Leather Cloak (2316, -0.22 DPS) [crafted]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (37.1 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.88 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 8.9 attack_power points (0.58 DPS) | yes | Bristlebark Bindings (14569, -0.10 DPS) [world_drop]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Ratchet Wristwraps (274742, -0.23 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 14.7 attack_power points (0.95 DPS) | yes | Bristlebark Gloves (14572, -0.23 DPS, sim-verified) [world_drop]; Brawler's Leather Gloves (252494, -0.23 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.16 DPS) | yes | Brawler's Leather Belt (252428, -0.44 DPS) [crafted]; Murloc Scale Belt (5780, -0.68 DPS) [crafted]; Deviate Scale Belt (6468, -3.85 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (37.1 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.07 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (37.1 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.98 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 11.2 attack_power points (0.72 DPS) | yes | Bounty Hunter's Ring (5351, -0.37 DPS) [quest]; Demon Band (12054, -0.46 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.49 DPS) [dungeon] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 9.8 attack_power points (0.63 DPS) | yes | Bounty Hunter's Ring (5351, -0.28 DPS) [quest]; Demon Band (12054, -0.37 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (14.71 DPS) | yes | Scout's Blade (20441, -0.48 DPS) [pvp]; Evocator's Blade (2567, -0.92 DPS) [dungeon]; Buzzer Blade (2169, -2.10 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep [dungeon] | 227.6 attack_power points (14.64 DPS) | yes | Evocator's Blade (2567, -3.94 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 7.2 attack_power points (0.46 DPS) | yes | Fine Longbow (11304, -0.19 DPS, sim-verified) [vendor]; Deadly Blunderbuss (4369, -0.23 DPS) [crafted]; Light Bow (4576, -0.23 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Legionnaire's Band; finger2: Pyrewood Signet Ring; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword; 248007 Militia Shortblade; 254779 A'sharahm, the Roiling Tempest; 263006 Scout Ranger's Tunic

### Band 30 (troll, 00000000000000000-00000000000000000-5323220310000000000)

Set DPS (verified): 58.3. Weights run: 2.5s. Verify run: 4.4s. 354 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.452 ± 0.011, crit=0.479 ± 0.012 per rating point (14 rating = 1%, 6.710 per %), hit=0.911 ± 0.035 per rating point (10 rating = 1%, 9.108 per %), melee_haste=7.054 ± 0.404

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 14.5 attack_power points (0.94 DPS) | yes | Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Humbert's Helm (4724, -0.28 DPS) [world]; Defender's Leather Helm (252455, -0.31 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 12.0 attack_power points (0.78 DPS) | yes | Scout's Medallion (19537, -0.02 DPS) [rep]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 21.0 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.43 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.57 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 13.2 attack_power points (0.85 DPS) | yes | Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Wolfmaster Cape (6314, -0.21 DPS) [dungeon]; Wildhunter Cloak (16658, -0.21 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 20.3 attack_power points (1.32 DPS) | yes | Brawler's Leather Tunic (252508, -0.27 DPS, sim-verified) [crafted]; Brawler's Leather Armor (252490, -0.34 DPS) [crafted]; Panther Armor (6670, -0.34 DPS) [quest] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 12.7 attack_power points (0.83 DPS) | yes | Barbaric Bracers (18948, -0.19 DPS) [crafted]; Insignia Bracers (6410, -0.26 DPS) [world_drop]; Cultist's Armguards (270032, -0.27 DPS, sim-verified) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.3 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Insignia Gloves (6408, +0.00 DPS) [world_drop]; Heavy Earthen Gloves (7359, -2.90 DPS, sim-verified) [crafted] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (58.3 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -3.16 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (58.3 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.86 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (58.3 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.14 DPS, sim-verified) [world_drop] |
| finger1 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 14.7 attack_power points (0.96 DPS) | yes | Band of the Fist (17694, -0.32 DPS) [quest]; Pyrewood Signet Ring (277210, -0.34 DPS) [quest]; Insurgent's Band (272067, -0.37 DPS) [vendor] |
| finger2 | Thunderbrow Ring (13097) | World drop [world_drop] | 12.4 attack_power points (0.80 DPS) | yes | Band of the Fist (17694, -0.17 DPS) [quest]; Pyrewood Signet Ring (277210, -0.19 DPS) [quest]; Insurgent's Band (272067, -0.22 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (20.91 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -0.78 DPS, sim-verified) [vendor]; Scout's Blade (19545, -0.99 DPS) [pvp]; Thornspike (6681, -2.19 DPS) [dungeon] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (58.3 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -4.57 DPS, sim-verified) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.58 DPS) | yes | Golemsight Long Gun (273029, -0.02 DPS) [dungeon]; Gravespike Repeater (273033, -0.06 DPS) [world]; Double-barreled Shotgun (2098, -0.11 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Legionnaire's Band; finger2: Thunderbrow Ring; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 354, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 40 (troll, 00000000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 98.3. Weights run: 2.7s. Verify run: 2.7s. 474 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.287 ± 0.009, crit=0.369 ± 0.011 per rating point (14 rating = 1%, 5.165 per %), hit=0.968 ± 0.037 per rating point (10 rating = 1%, 9.679 per %), melee_haste=4.905 ± 0.444

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 25.2 attack_power points (1.86 DPS) | yes | Hawkeye's Helm (14591, -0.59 DPS, sim-verified) [world_drop]; Warden's Wizard Hat (14604, -0.62 DPS) [world_drop]; Nightscape Headband (8176, -0.72 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 14.2 attack_power points (1.04 DPS) | yes | Ethereal Talisman (4430, -0.30 DPS) [quest]; Kaleidoscope Chain (13084, -0.37 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.54 DPS, sim-verified) [world] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 26.2 attack_power points (1.93 DPS) | yes | Forest Tracker Epaulets (2278, -0.59 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.61 DPS) [dungeon]; Nightscape Shoulders (8192, -0.89 DPS) [crafted] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 12.0 attack_power points (0.89 DPS) | yes | Imperial Cloak (6432, -0.13 DPS) [dungeon]; Parachute Cloak (10518, -0.13 DPS) [crafted]; Tigerstrike Mantle (13108, -0.13 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 21.9 attack_power points (1.61 DPS) | yes | Nightscape Tunic (8175, -0.19 DPS) [crafted]; Dusky Leather Armor (7374, -0.28 DPS) [crafted]; Barbaric Harness (5739, -0.34 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.54 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.72 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 25.2 attack_power points (1.86 DPS) | yes | Skulker's Leather Gloves (252525, -0.24 DPS) [crafted]; Prowler's Leather Gloves (252524, -0.26 DPS) [crafted]; Dog Training Gloves (7756, -0.64 DPS, sim-verified) [dungeon] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.21 DPS) | yes | Defiler's Chain Girdle (20152, -0.51 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.70 DPS) [world_drop]; Blackened Defias Belt (10403, -0.89 DPS) [dungeon] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.0 attack_power points (1.99 DPS) | yes | Triprunner Dungarees (9624, -0.06 DPS) [quest]; Ferine Leggings (6690, -0.08 DPS) [dungeon]; Brawler's Leather Legguards (252516, -0.62 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 21.2 attack_power points (1.56 DPS) | yes | Prowler's Leather Shoes (252465, -0.08 DPS) [crafted]; Imperial Leather Boots (6431, -0.15 DPS) [dungeon]; Excelsior Boots (4109, -0.19 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery [dungeon] | 20.0 attack_power points (1.48 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.48 DPS) | yes | Legionnaire's Band (19512, -0.13 DPS) [rep]; Ring of the Underwood (2951, -0.38 DPS) [world_drop]; Falcon's Hook (7552, -0.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (98.3 DPS) | yes | Scout's Blade (19544, -4.17 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -5.13 DPS) [vendor]; Coldrage Dagger (10761, -8.30 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 441.2 attack_power points (32.57 DPS) | yes | Coldrage Dagger (10761, +0.00 DPS) [dungeon] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (98.3 DPS) | yes | Monolithic Bow (9426, -0.31 DPS) [dungeon]; Master Hunter's Rifle (17687, -0.36 DPS) [quest]; Bow of Searing Arrows (2825, -1.19 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Scout's Medallion; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 474, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 00532000000000000-00000000000000000-5323220310013011031)

Set DPS (verified): 137.8. Weights run: 3.5s. Verify run: 3.2s. 592 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.417 ± 0.007, crit=0.701 ± 0.011 per rating point (14 rating = 1%, 9.811 per %), hit=1.584 ± 0.067 per rating point (10 rating = 1%, 15.837 per %), melee_haste=5.880 ± 0.944

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 26.6 attack_power points (2.25 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Undercity Reservist's Cap (20643, -0.08 DPS) [quest]; Helm of Fire (8348, -0.21 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 23.4 attack_power points (1.98 DPS) | yes | Woven Ivy Necklace (19159, -0.49 DPS, sim-verified) [quest]; Scout's Medallion (19535, -0.54 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 27.6 attack_power points (2.34 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.29 DPS) [quest] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 21.2 attack_power points (1.80 DPS) | yes | Blackflame Cape (13109, -0.36 DPS) [world_drop]; Pridelord Cape (14673, -0.47 DPS) [world_drop]; Serpentskin Cloak (8259, -0.48 DPS) [world_drop] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Blazewind Breastplate (11193, -0.68 DPS) [quest]; Charred Leather Tunic (19127, -1.54 DPS) [quest]; Warbear Harness (15064, -1.61 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 25.2 attack_power points (2.14 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.62 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 41.2 attack_power points (3.49 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.96 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.04 DPS) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | 31.8 attack_power points (2.70 DPS) | yes | Prowler's Leather Waistguard (252473, -0.07 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.16 DPS) [rep] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 43.6 attack_power points (3.70 DPS) | yes | Basilisk Hide Pants (1718, -1.18 DPS) [world_drop]; Triprunner Dungarees (9624, -1.28 DPS) [quest]; Serpentskin Leggings (8262, -1.70 DPS, sim-verified) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 29.4 attack_power points (2.49 DPS) | yes | Prowler's Leather Boots (252468, -0.07 DPS) [crafted]; Albino Crocscale Boots (17728, -0.09 DPS) [dungeon]; Skulker's Leather Shoes (252531, -0.58 DPS) [crafted] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.03 DPS) | yes | Mark of Kern (2262, -0.34 DPS) [dungeon]; Assault Band (13095, -0.34 DPS) [world_drop]; Blackstone Ring (17713, -0.34 DPS) [dungeon] |
| finger2 | Legionnaire's Band (19511) | Warsong Outriders [rep] | 22.7 attack_power points (1.93 DPS) | yes | Mark of Kern (2262, -0.23 DPS) [dungeon]; Assault Band (13095, -0.23 DPS) [world_drop]; Blackstone Ring (17713, -0.23 DPS) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (137.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (137.8 DPS) | yes | Molten Heart of the Mountain (249470, -1.76 DPS, sim-verified) [crafted] |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (137.8 DPS) | yes | Charstone Dirk (17710, +0.00 DPS) [dungeon]; Scout's Blade (19543, +0.00 DPS) [pvp]; Julie's Dagger (6660, -7.01 DPS, sim-verified) [world_drop] |
| off_hand | Lifeforce Dirk (10750) | The God Hakkar [quest] | 503.2 attack_power points (42.60 DPS) | yes | Charstone Dirk (17710, -26.88 DPS, sim-verified) [dungeon] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (137.8 DPS) | yes | Skull Splitting Crossbow (13039, -0.24 DPS) [world_drop]; Houndmaster's Bow (11628, -0.41 DPS) [dungeon]; Dark Iron Rifle (16004, -2.04 DPS, sim-verified) [crafted] |

**New at 50:** neck: Skibi's Pendant; back: Dark Phantom Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Skulker's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Legionnaire's Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; off_hand: Lifeforce Dirk; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 259.5. Weights run: 3.4s. Verify run: 8.9s. 1165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.575 ± 0.012, crit=1.203 ± 0.020 per rating point (14 rating = 1%, 16.841 per %), hit=1.913 ± 0.086 per rating point (10 rating = 1%, 19.127 per %), melee_haste=8.267 ± 1.157

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 58.3 attack_power points (5.89 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.21 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 43.1 attack_power points (4.36 DPS) | yes | Amulet of the Darkmoon (19491, -0.32 DPS) [quest]; Zandalarian Shadow Talisman (19616, -1.36 DPS) [quest]; Medallion of the Dawn (22659, -2.23 DPS, sim-verified) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (259.5 DPS) | yes | Defiler's Leather Shoulders (20194, +0.00 DPS) [rep]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -13.62 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 47.1 attack_power points (4.76 DPS) | yes | Deathguard's Cloak (20068, -0.53 DPS) [rep]; Shifting Cloak (18511, -2.05 DPS) [crafted]; Cape of the Black Baron (13340, -2.28 DPS, sim-verified) [dungeon] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 94.3 attack_power points (9.52 DPS) | yes | Darkmantle Tunic (226825, -1.69 DPS) [quest]; Warlord's Leather Breastplate (231549, -1.75 DPS) [pvp]; Timbermaw Tunic (252484, -3.43 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 40.9 attack_power points (4.13 DPS) | yes | General's Leather Armsplints (16559, -0.10 DPS) [pvp]; Bracers of Subterfuge (22668, -0.74 DPS) [quest]; Deepfury Bracers (13120, -1.34 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (259.5 DPS) | yes | Blood Guard's Leather Vices (16499, +0.00 DPS) [pvp]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -16.26 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 71.3 attack_power points (7.20 DPS) | yes | Ferocity of the Timbermaw (227805, -1.87 DPS) [vendor]; Defiler's Leather Girdle (20190, -2.07 DPS) [rep]; Belt of Preserved Heads (20216, -3.41 DPS, sim-verified) [quest] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (259.5 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -13.62 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 51.2 attack_power points (5.17 DPS) | yes | Dunestalker's Boots (20715, -1.24 DPS) [quest]; Shadowcraft Boots (16711, -1.25 DPS) [dungeon]; Drudge Boots (21532, -2.10 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 56.9 attack_power points (5.75 DPS) | yes | Cutthroat's Signet (272408, -1.59 DPS) [vendor]; Legionnaire's Band (19510, -2.79 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 50.0 attack_power points (5.05 DPS) | yes | Cutthroat's Signet (272408, -0.96 DPS, sim-verified) [vendor]; Legionnaire's Band (19510, -2.08 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (259.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (259.5 DPS) | yes | Frozen Heart of the Mountain (249469, -5.38 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (259.5 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -16.43 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (66.76 DPS) | yes | Glacial Blade (19099, +0.00 DPS) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (259.5 DPS) | yes | Bloodseeker (19107, -0.31 DPS) [quest]; Skull Splitting Crossbow (13039, -0.51 DPS) [world_drop]; Dark Iron Rifle (16004, -2.62 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (troll, 00532310101400000-00000000000000000-5323220310013011031)

Set DPS (verified): 612.9. Weights run: 3.7s. Verify run: 9.5s. 1165 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.009 ± 0.017, crit=1.706 ± 0.028 per rating point (14 rating = 1%, 23.888 per %), hit=2.826 ± 0.148 per rating point (10 rating = 1%, 28.260 per %), melee_haste=13.252 ± 1.654

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 78.0 attack_power points (14.30 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Darkmantle Cap (226829, -0.72 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 52.3 attack_power points (9.58 DPS) | yes | Amulet of the Darkmoon (19491, -0.57 DPS) [quest]; Medallion of the Dawn (22659, -0.80 DPS) [quest]; Zandalarian Shadow Talisman (19616, -2.85 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (612.9 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Warlord's Leather Spaulders (231551, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -28.64 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 56.3 attack_power points (10.31 DPS) | yes | Cape of the Black Baron (13340, -1.12 DPS) [dungeon]; Deathguard's Cloak (20068, -2.24 DPS) [rep]; Shifting Cloak (18511, -4.05 DPS) [crafted] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | 126.5 attack_power points (23.19 DPS) | yes | Darkmantle Tunic (226825, -4.02 DPS, sim-verified) [quest]; Warlord's Leather Breastplate (231549, -4.05 DPS) [pvp]; Timbermaw Tunic (252484, -5.52 DPS) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 50.3 attack_power points (9.22 DPS) | yes | General's Leather Armsplints (16559, -0.20 DPS) [pvp]; Bracers of Subterfuge (22668, -1.67 DPS) [quest]; Deepfury Bracers (13120, -2.88 DPS) [world_drop] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (612.9 DPS) | yes | Darkmantle Gloves (22006, +0.00 DPS) [quest]; General's Leather Mitts (231555, +0.00 DPS) [pvp]; Raider Gloves (272099, -34.60 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 98.7 attack_power points (18.10 DPS) | yes | Belt of Preserved Heads (20216, -4.19 DPS, sim-verified) [quest]; Ferocity of the Timbermaw (227805, -6.59 DPS) [vendor]; General's Leather Girdle (16557, -7.33 DPS) [pvp] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (612.9 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -28.64 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 68.0 attack_power points (12.47 DPS) | yes | Shadowcraft Boots (16711, -3.18 DPS) [dungeon]; Fine Dawn Treaders (227815, -3.23 DPS) [vendor]; Drudge Boots (21532, -3.58 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234202) | Anachronos [vendor] | 76.5 attack_power points (14.02 DPS) | yes | Cutthroat's Signet (272408, -3.68 DPS) [vendor]; Legionnaire's Band (19510, -7.55 DPS) [rep] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 66.1 attack_power points (12.13 DPS) | yes | Cutthroat's Signet (272408, -1.79 DPS) [vendor]; Legionnaire's Band (19510, -5.65 DPS) [rep] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (612.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (612.9 DPS) | yes | Frozen Heart of the Mountain (249469, -7.46 DPS, sim-verified) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (612.9 DPS) | yes | High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor]; Glacial Blade (19099, -28.62 DPS, sim-verified) [rep] |
| off_hand | The Lobotomizer (19324) | Frostwolf Clan [rep] | 661.1 attack_power points (121.19 DPS) | yes | Glacial Blade (19099, +0.00 DPS, sim-verified) [rep] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (612.9 DPS) | yes | The Purifier (22656, -0.78 DPS) [quest]; Bloodseeker (19107, -0.96 DPS) [quest]; Dark Iron Rifle (16004, -6.05 DPS, sim-verified) [crafted] |

**New at 60:** head: Outlaw's Collar; neck: Beads of Ogre Might; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: The Lobotomizer

No-known-source sample (15 of 1165, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

