# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.9. Weights run: 1.9s. Verify run: 3.3s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.042 ± 0.014, crit=0.656 ± 0.013 per rating point (14 rating = 1%, 9.190 per %), hit=0.821 ± 0.027 per rating point (10 rating = 1%, 8.206 per %), melee_haste=5.169 ± 0.269

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.3 attack_power points (0.94 DPS) | yes | Defender's Leather Hood (252447, -0.49 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 12.3 attack_power points (0.70 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.2 attack_power points (0.59 DPS) | yes | Slime-encrusted Pads (6461, -0.65 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.3 attack_power points (0.70 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Hide of Lupos (3018, -0.23 DPS) [world]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 22.5 attack_power points (1.29 DPS) | yes | Brawler's Leather Armor (252490, -0.18 DPS) [crafted]; Trapper's Leather Armor (252491, -0.47 DPS) [crafted]; Prospector's Chestpiece (14562, -0.47 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.2 attack_power points (0.59 DPS) | yes | Bravo's Armbands (270015, -0.00 DPS) [quest]; Wolf Bracers (4794, -0.12 DPS) [vendor]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.9 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Bristlebark Gloves (14572, +0.00 DPS) [world_drop]; Gloves of the Fang (10413, -2.92 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.34 DPS) [crafted]; Dusty Belt (279897, -0.45 DPS) [quest]; Deviate Scale Belt (6468, -3.96 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.9 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.18 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.9 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.10 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.3 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Demon Band (12054, -0.47 DPS) [world_drop] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 12.2 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.30 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Demon Band (12054, -0.47 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (13.14 DPS) | yes | Evocator's Blade (2567, -0.82 DPS) [dungeon]; Buzzer Blade (2169, -1.88 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.39 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 228.6 attack_power points (13.14 DPS) | yes | Evocator's Blade (2567, -9.45 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.2 attack_power points (0.47 DPS) | yes | Light Bow (4576, -0.23 DPS) [world_drop]; Owlsight Rifle (15205, -0.23 DPS) [quest]; Deadly Blunderbuss (4369, -0.25 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Blackened Defias Gloves; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Signet of the Zhevra; finger2: Protector's Band; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 65.9. Weights run: 2.0s. Verify run: 3.4s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.596 ± 0.013, crit=0.590 ± 0.014 per rating point (14 rating = 1%, 8.262 per %), hit=1.069 ± 0.033 per rating point (10 rating = 1%, 10.686 per %), melee_haste=7.480 ± 0.355

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.0 attack_power points (0.96 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.24 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.85 DPS) | yes | Sentinel's Medallion (19541, -0.07 DPS) [rep]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.6 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.2 attack_power points (0.86 DPS) | yes | Tigerstrike Mantle (13108, -0.08 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.23 DPS) [pvp]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.3 attack_power points (1.35 DPS) | yes | Brawler's Leather Tunic (252508, -0.29 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.29 DPS) [quest]; Brawler's Leather Armor (252490, -0.37 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.6 attack_power points (0.82 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.22 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.9 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -2.76 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (65.9 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Highlander's Chain Girdle (20090, -3.13 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.9 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -2.84 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (65.9 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.11 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.4 attack_power points (1.11 DPS) | yes | Thunderbrow Ring (13097, -0.34 DPS) [world_drop]; Monkey Ring (6748, -0.43 DPS) [quest]; Pyrewood Signet Ring (277210, -0.47 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 15.6 attack_power points (0.94 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Monkey Ring (6748, -0.27 DPS) [quest]; Pyrewood Signet Ring (277210, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (19.45 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.33 DPS) [vendor]; Torturing Poker (7682, -1.54 DPS) [dungeon]; Thornspike (6681, -2.04 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (19.33 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -7.98 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -19.23 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.54 DPS) | yes | Silver Star (3463, -0.06 DPS) [quest]; Double-barreled Shotgun (2098, -0.07 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.16 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 117.6. Weights run: 2.2s. Verify run: 1.9s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.582 ± 0.014, crit=0.747 ± 0.019 per rating point (14 rating = 1%, 10.462 per %), hit=1.126 ± 0.039 per rating point (10 rating = 1%, 11.262 per %), melee_haste=8.039 ± 0.489

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.4 attack_power points (1.90 DPS) | yes | Warden's Wizard Hat (14604, -0.52 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.63 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.34 DPS) | yes | Sentinel's Medallion (19540, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.4 attack_power points (1.97 DPS) | yes | Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.80 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.8 attack_power points (1.33 DPS) | yes | Sergeant Major's Cape (16336, -0.29 DPS) [pvp]; Hawkeye's Cloak (14593, -0.38 DPS) [world_drop]; Parachute Cloak (10518, -0.48 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 38.1 attack_power points (2.55 DPS) | yes | Wolffear Harness (13110, -0.92 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.96 DPS) [crafted]; Dusky Leather Armor (7374, -1.06 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.34 DPS) | yes | Hawkeye's Bracers (14590, -0.48 DPS, sim-verified) [world_drop]; Imperial Leather Bracers (4061, -0.49 DPS) [dungeon]; Dusky Bracers (7378, -0.49 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.5 attack_power points (2.04 DPS) | yes | Prowler's Leather Gloves (252524, -0.42 DPS) [crafted]; Imperial Leather Gloves (4063, -0.48 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.53 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (2.01 DPS) | yes | Ogron's Sash (13117, -0.45 DPS) [world_drop]; Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.65 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 33.2 attack_power points (2.22 DPS) | yes | Triprunner Dungarees (9624, -0.12 DPS) [quest]; Ferine Leggings (6690, -0.48 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.74 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.4 attack_power points (1.63 DPS) | yes | Imperial Leather Boots (6431, -0.13 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.16 DPS) [crafted]; Excelsior Boots (4109, -0.21 DPS) [quest] |
| finger1 | Protector's Band (19515) | Silverwing Sentinels [rep] | 20.7 attack_power points (1.38 DPS) | yes | Assault Band (13095, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Field Researcher's Loop (281634, -0.17 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.34 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon]; Field Researcher's Loop (281634, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (117.6 DPS) | yes | Black Menace (6831, -3.95 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -4.65 DPS) [vendor]; Coldrage Dagger (10761, -4.66 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 443.9 attack_power points (29.71 DPS) | yes | Black Menace (6831, -0.86 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -29.51 DPS) [world_drop]; Satyr's Rod (15962, -29.61 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (117.6 DPS) | yes | Swiftwind (13038, -0.20 DPS) [world_drop]; Monolithic Bow (9426, -0.22 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.15 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Protector's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 169.7. Weights run: 2.3s. Verify run: 2.2s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.533 ± 0.013, crit=0.893 ± 0.020 per rating point (14 rating = 1%, 12.502 per %), hit=1.596 ± 0.058 per rating point (10 rating = 1%, 15.959 per %), melee_haste=10.565 ± 0.857

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 48.5 attack_power points (4.07 DPS) | yes | Embrace of the Lycan (9479, -0.71 DPS) [dungeon]; Knight-Lieutenant's Leather Headband (220850, -1.26 DPS, sim-verified) [vendor]; White Bandit Mask (10008, -1.73 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.9 attack_power points (2.09 DPS) | yes | Zealous Shadowshard Pendant (17772, -0.41 DPS) [quest]; Sentinel's Medallion (19539, -0.55 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | 36.5 attack_power points (3.06 DPS) | yes | Skulker's Leather Shoulder (252535, -0.89 DPS) [crafted]; Failed Flying Experiment (9647, -0.93 DPS) [quest]; Sunburn Spaulders (274751, -1.43 DPS, sim-verified) [vendor] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 27.5 attack_power points (2.30 DPS) | yes | Blisterbane Wrap (12552, -0.37 DPS) [dungeon]; Dark Phantom Cape (13122, -0.37 DPS) [world_drop]; Duskbat Drape (19982, -0.50 DPS) [quest] |
| chest | Knight's Leather Armor (220854) | Captain Dirgehammer [vendor] | 46.5 attack_power points (3.89 DPS) | yes | Fungus Shroud Armor (17742, -0.68 DPS) [dungeon]; Blazewind Breastplate (11193, -0.69 DPS) [quest]; Warbear Harness (15064, -1.38 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.0 attack_power points (2.26 DPS) | yes | Pridelord Bands (14672, -0.56 DPS) [world_drop]; Branded Leather Bracers (19508, -0.59 DPS) [dungeon]; Skulker's Leather Bracers (252540, -0.63 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 43.7 attack_power points (3.66 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.94 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.11 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.18 DPS) | yes | Skulker's Leather Waistguard (252474, -0.38 DPS) [crafted]; Highlander's Leather Girdle (20115, -0.46 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.47 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 46.5 attack_power points (3.89 DPS) | yes | Serpentskin Leggings (8262, -0.75 DPS) [world_drop]; Basilisk Hide Pants (1718, -1.20 DPS) [world_drop]; Gryphon Rider's Leggings (9652, -1.49 DPS, sim-verified) [quest] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 32.1 attack_power points (2.69 DPS) | yes | Skulker's Leather Boots (252469, -0.09 DPS) [crafted]; Albino Crocscale Boots (17728, -0.12 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.18 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.0 attack_power points (3.01 DPS) | yes | Masons Fraternity Ring (9533, -1.22 DPS) [quest]; Mark of Kern (2262, -1.34 DPS) [dungeon]; Assault Band (13095, -1.34 DPS) [world_drop] |
| finger2 | Protector's Band (19516) | Silverwing Sentinels [rep] | 23.8 attack_power points (1.99 DPS) | yes | Masons Fraternity Ring (9533, -0.20 DPS) [quest]; Mark of Kern (2262, -0.32 DPS) [dungeon]; Assault Band (13095, -0.32 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (169.7 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (169.7 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (169.7 DPS) | yes | Barman Shanker (12791, -2.11 DPS, sim-verified) [dungeon]; Lifeforce Dirk (10750, -3.51 DPS) [quest]; Charstone Dirk (17710, -3.51 DPS) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (42.88 DPS) | yes | Barman Shanker (12791, -20.71 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -42.49 DPS) [quest]; Stonecloth Branch (15963, -42.62 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (169.7 DPS) | yes | Stinging Bow (10624, -0.37 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.37 DPS) [world_drop]; Dark Iron Rifle (16004, -2.13 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Blackveil Cape; chest: Knight's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Knight's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: Protector's Band; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 263.9. Weights run: 2.0s. Verify run: 6.7s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.083 ± 0.025, crit=2.239 ± 0.044 per rating point (14 rating = 1%, 31.353 per %), hit=3.314 ± 0.159 per rating point (10 rating = 1%, 33.141 per %), melee_haste=12.233 ± 2.173

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 97.6 attack_power points (7.95 DPS) | yes | Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.4 attack_power points (5.24 DPS) | yes | Mark of Fordring (15411, -0.57 DPS) [quest]; Beads of Ogre Might (22150, -0.59 DPS) [quest]; Medallion of the Dawn (22659, -0.74 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -11.80 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 61.1 attack_power points (4.98 DPS) | yes | Cape of the Black Baron (13340, -0.81 DPS) [dungeon]; Cloak of the Honor Guard (20073, -1.36 DPS) [rep]; Windshear Cape (20691, -1.78 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, -1.87 DPS) [pvp]; Darkmantle Tunic (226825, -2.22 DPS) [quest]; Tunic of Undead Slaying (23089, -9.85 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.08 DPS) [pvp]; Bracers of the Eclipse (18375, -0.47 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.36 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.42 DPS) [crafted]; Raider Gloves (272099, -13.40 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 110.0 attack_power points (8.96 DPS) | yes | Belt of Preserved Heads (20216, -1.34 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -3.64 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.96 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Leather Legguards (23299, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -10.94 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 73.2 attack_power points (5.96 DPS) | yes | Shadowcraft Boots (16711, -1.59 DPS) [dungeon]; Fine Dawn Treaders (227815, -1.74 DPS, sim-verified) [vendor]; Darkmantle Boots (22003, -1.89 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.31 DPS) [dungeon]; Cutthroat's Signet (272408, -1.48 DPS) [vendor]; Naglering (11669, -5.28 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -0.51 DPS) [dungeon]; Cutthroat's Signet (272408, -0.68 DPS) [vendor]; Naglering (11669, -4.80 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Hand of Justice (11815, -5.88 DPS, sim-verified) [dungeon] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.68 DPS) [crafted]; Counterattack Lodestone (18537, -3.31 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+9.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (263.9 DPS) | yes | The Lobotomizer (19324, -2.65 DPS, sim-verified) [rep]; Distracting Dagger (18392, -5.50 DPS) [dungeon]; Scepter of Interminable Focus (22329, -46.73 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.27 DPS) [dungeon]; The Purifier (22656, -0.65 DPS) [quest]; Dark Iron Rifle (16004, -2.30 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60, raid preset (night-elf, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 715.6. Weights run: 2.1s. Verify run: 6.7s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.504 ± 0.030, crit=2.686 ± 0.050 per rating point (14 rating = 1%, 37.604 per %), hit=3.804 ± 0.235 per rating point (10 rating = 1%, 38.044 per %), melee_haste=25.013 ± 3.160

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 113.7 attack_power points (19.06 DPS) | yes | Field Marshal's Leather Mask (231545, +0.00 DPS) [pvp]; Lieutenant Commander's Leather Helm (227055, -0.34 DPS) [pvp]; Outlaw's Collar (279253, -0.79 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 75.6 attack_power points (12.67 DPS) | yes | Mark of Fordring (15411, -2.01 DPS) [quest]; Beads of Ogre Might (22150, -2.27 DPS) [quest]; Medallion of the Dawn (22659, -2.35 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Lieutenant Commander's Leather Shoulders (227054, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -41.42 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.0 attack_power points (11.07 DPS) | yes | Cape of the Black Baron (13340, -1.42 DPS) [dungeon]; Cloak of the Honor Guard (20073, -3.27 DPS) [rep]; Windshear Cape (20691, -3.30 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Leather Chestpiece (231543, -4.62 DPS) [pvp]; Darkmantle Tunic (226825, -5.63 DPS) [quest]; Tunic of Undead Slaying (23089, -24.77 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Armsplints (16460, -0.18 DPS) [pvp]; Bracers of the Eclipse (18375, -1.78 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.19 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Leather Handgrips (231544, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.23 DPS) [quest]; Raider Gloves (272099, -46.00 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 128.7 attack_power points (21.57 DPS) | yes | Belt of Preserved Heads (20216, -3.06 DPS, sim-verified) [quest]; Highlander's Leather Girdle (20045, -9.57 DPS) [rep]; Ferocity of the Timbermaw (227805, -9.64 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; Marshal's Leather Leggings (231548, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -41.42 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 86.7 attack_power points (14.54 DPS) | yes | Shadowcraft Boots (16711, -3.81 DPS) [dungeon]; Darkmantle Boots (22003, -4.46 DPS) [quest]; Fine Dawn Treaders (227815, -4.75 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.69 DPS) [dungeon]; Cutthroat's Signet (272408, -3.11 DPS) [vendor]; Naglering (11669, -12.84 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.26 DPS) [dungeon]; Cutthroat's Signet (272408, -1.68 DPS) [vendor]; Naglering (11669, -11.91 DPS, sim-verified) [dungeon] |
| trinket1 | Hand of Justice (11815) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | sim-verified (+16.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Blackhand's Breadth (13965) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Darkmoon Card: Maelstrom (19289, -3.37 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -6.87 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+21.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (715.6 DPS) | yes | The Lobotomizer (19324, -8.14 DPS, sim-verified) [rep]; Distracting Dagger (18392, -11.31 DPS) [dungeon]; Scepter of Interminable Focus (22329, -94.30 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.64 DPS) [dungeon]; The Purifier (22656, -1.33 DPS) [quest]; Dark Iron Rifle (16004, -5.29 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Hand of Justice; trinket2: Blackhand's Breadth; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 41.2. Weights run: 1.9s. Verify run: 3.2s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.042 ± 0.014, crit=0.656 ± 0.013 per rating point (14 rating = 1%, 9.190 per %), hit=0.821 ± 0.027 per rating point (10 rating = 1%, 8.206 per %), melee_haste=5.169 ± 0.269

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 16.3 attack_power points (0.94 DPS) | yes | Defender's Leather Hood (252447, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 12.3 attack_power points (0.70 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.2 attack_power points (0.59 DPS) | yes | Slime-encrusted Pads (6461, -0.62 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 12.3 attack_power points (0.70 DPS) | yes | Cape of the Brotherhood (5193, -0.12 DPS) [dungeon]; Hide of Lupos (3018, -0.23 DPS) [world]; Bristlebark Cape (14571, -0.23 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | sim-verified (41.2 DPS) | yes | Prospector's Chestpiece (14562, +0.00 DPS) [world_drop]; Trapper's Leather Armor (252491, +0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -2.81 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.2 attack_power points (0.59 DPS) | yes | Wolf Bracers (4794, -0.12 DPS) [vendor]; Bristlebark Bindings (14569, -0.12 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.23 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 16.3 attack_power points (0.93 DPS) | yes | Bristlebark Gloves (14572, -0.23 DPS) [world_drop]; Brawler's Leather Gloves (252494, -0.23 DPS) [crafted]; Serpent Gloves (5970, -0.27 DPS, sim-verified) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (1.03 DPS) | yes | Brawler's Leather Belt (252428, -0.34 DPS) [crafted]; Dark Leather Belt (4249, -0.57 DPS) [crafted]; Deviate Scale Belt (6468, -3.79 DPS, sim-verified) [crafted] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (41.2 DPS) | yes | Leggings of the Fang (10410, +0.00 DPS) [dungeon]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Brawler's Leather Pants (252500, -3.01 DPS, sim-verified) [crafted] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (41.2 DPS) | yes | Footpads of the Fang (10411, +0.00 DPS) [dungeon]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -2.92 DPS, sim-verified) [world_drop] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 12.3 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.05 DPS) [quest]; Bounty Hunter's Ring (5351, -0.35 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 12.2 attack_power points (0.70 DPS) | yes | Pyrewood Signet Ring (277210, -0.28 DPS, sim-verified) [quest]; Bounty Hunter's Ring (5351, -0.35 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (13.14 DPS) | yes | Edward's Knife (251485, -0.10 DPS) [quest]; Scout's Blade (20441, -0.37 DPS) [pvp]; Evocator's Blade (2567, -0.82 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 228.6 attack_power points (13.14 DPS) | yes | Edward's Knife (251485, +0.00 DPS, sim-verified) [quest]; Tork Wrench (11855, -13.02 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 8.2 attack_power points (0.47 DPS) | yes | Light Bow (4576, -0.23 DPS) [world_drop]; Privateer Musket (5309, -0.23 DPS) [quest]; Deadly Blunderbuss (4369, -0.25 DPS, sim-verified) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Blackened Defias Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Blackened Defias Leggings; feet: Blackened Defias Boots; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-32530300001400000-0000000000000000000)

Set DPS (verified): 65.8. Weights run: 2.0s. Verify run: 3.3s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.596 ± 0.013, crit=0.590 ± 0.014 per rating point (14 rating = 1%, 8.262 per %), hit=1.069 ± 0.033 per rating point (10 rating = 1%, 10.686 per %), melee_haste=7.480 ± 0.355

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 16.0 attack_power points (0.96 DPS) | yes | Tribal Worg Helm (6204, -0.19 DPS) [world]; Brawler's Leather Hood (252504, -0.19 DPS) [crafted]; Defender's Leather Helm (252455, -0.24 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.85 DPS) | yes | Scout's Medallion (19537, -0.07 DPS) [rep]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 22.6 attack_power points (1.36 DPS) | yes | Mantle of Thieves (2264, -0.47 DPS, sim-verified) [dungeon]; Barbaric Shoulders (5964, -0.58 DPS) [crafted]; Bristlebark Amice (14573, -0.60 DPS) [world_drop] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 14.2 attack_power points (0.86 DPS) | yes | Tigerstrike Mantle (13108, -0.08 DPS) [world_drop]; Wolfmaster Cape (6314, -0.25 DPS) [dungeon]; Wildhunter Cloak (16658, -0.25 DPS) [quest] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 22.3 attack_power points (1.35 DPS) | yes | Brawler's Leather Tunic (252508, -0.28 DPS, sim-verified) [crafted]; Panther Armor (6670, -0.36 DPS) [quest]; Brawler's Leather Armor (252490, -0.37 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 13.6 attack_power points (0.82 DPS) | yes | Jurassic Wristguards (6198, -0.12 DPS) [world]; Barbaric Bracers (18948, -0.19 DPS) [crafted]; Cultist's Armguards (270032, -0.22 DPS) [quest] |
| hands | Blackened Defias Gloves (10401) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.8 DPS) | yes | Toughened Leather Gloves (4253, +0.00 DPS) [crafted]; Heavy Earthen Gloves (7359, +0.00 DPS) [crafted]; Insignia Gloves (6408, -3.15 DPS, sim-verified) [world_drop] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | sim-verified (65.8 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Skulker's Leather Belt (252520, +0.00 DPS) [crafted]; Defiler's Chain Girdle (20152, -3.48 DPS, sim-verified) [rep] |
| legs | Blackened Defias Leggings (10400) | The Deadmines: Defias Overseer [dungeon] | sim-verified (65.8 DPS) | yes | Petrolspill Leggings (9509, +0.00 DPS) [dungeon]; Troll's Bane Leggings (13114, +0.00 DPS) [world_drop]; Ferine Leggings (6690, -3.23 DPS, sim-verified) [dungeon] |
| feet | Blackened Defias Boots (10402) | The Deadmines: Defias Strip Miner [dungeon] | sim-verified (65.8 DPS) | yes | Insignia Boots (4055, +0.00 DPS) [world_drop]; Brawler's Leather Boots (252439, +0.00 DPS) [crafted]; Feet of the Lynx (1121, -3.48 DPS, sim-verified) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 18.4 attack_power points (1.11 DPS) | yes | Thunderbrow Ring (13097, -0.34 DPS) [world_drop]; Monkey Ring (6748, -0.43 DPS) [quest]; Pyrewood Signet Ring (277210, -0.47 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 15.6 attack_power points (0.94 DPS) | yes | Thunderbrow Ring (13097, -0.17 DPS) [world_drop]; Monkey Ring (6748, -0.27 DPS) [quest]; Pyrewood Signet Ring (277210, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (19.45 DPS) | yes | Scout's Blade (19545, -0.86 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.33 DPS) [vendor]; Torturing Poker (7682, -1.54 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (19.33 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -8.20 DPS, sim-verified) [vendor]; Tork Wrench (11855, -19.21 DPS) [quest]; Satyr's Rod (15962, -19.23 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.54 DPS) | yes | Silver Star (3463, -0.06 DPS) [quest]; Double-barreled Shotgun (2098, -0.07 DPS) [world_drop]; BKP "Sparrow" Smallbore (3042, -0.16 DPS) [world_drop] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Blackened Defias Gloves; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 116.8. Weights run: 2.2s. Verify run: 1.9s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.582 ± 0.014, crit=0.747 ± 0.019 per rating point (14 rating = 1%, 10.462 per %), hit=1.126 ± 0.039 per rating point (10 rating = 1%, 11.262 per %), melee_haste=8.039 ± 0.489

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 28.4 attack_power points (1.90 DPS) | yes | Warden's Wizard Hat (14604, -0.52 DPS) [world_drop]; Hawkeye's Helm (14591, -0.57 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.63 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.34 DPS) | yes | Scout's Medallion (19536, -0.17 DPS) [rep]; Ghostshard Talisman (7731, -0.40 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 29.4 attack_power points (1.97 DPS) | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Flintrock Shoulders (7755, -0.57 DPS) [dungeon]; Nightscape Shoulders (8192, -0.80 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 19.8 attack_power points (1.33 DPS) | yes | First Sergeant's Cloak (16340, -0.29 DPS) [pvp]; Hawkeye's Cloak (14593, -0.38 DPS) [world_drop]; Parachute Cloak (10518, -0.48 DPS) [crafted] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 38.1 attack_power points (2.55 DPS) | yes | Nightscape Tunic (8175, -0.96 DPS) [crafted]; Wolffear Harness (13110, -0.96 DPS, sim-verified) [world_drop]; Dusky Leather Armor (7374, -1.06 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.34 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Hawkeye's Bracers (14590, -0.44 DPS) [world_drop]; Dusky Bracers (7378, -0.49 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 30.5 attack_power points (2.04 DPS) | yes | Prowler's Leather Gloves (252524, -0.42 DPS) [crafted]; Imperial Leather Gloves (4063, -0.48 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.65 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (2.01 DPS) | yes | Ogron's Sash (13117, -0.45 DPS) [world_drop]; Defiler's Chain Girdle (20152, -0.49 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.65 DPS) [crafted] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 33.2 attack_power points (2.22 DPS) | yes | Triprunner Dungarees (9624, -0.12 DPS) [quest]; Ferine Leggings (6690, -0.48 DPS) [dungeon]; Hawkeye's Breeches (14595, -0.74 DPS) [world_drop] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 24.4 attack_power points (1.63 DPS) | yes | Imperial Leather Boots (6431, -0.13 DPS) [dungeon]; Prowler's Leather Shoes (252465, -0.16 DPS) [crafted]; Excelsior Boots (4109, -0.21 DPS) [quest] |
| finger1 | Legionnaire's Band (19512) | Warsong Outriders [rep] | 20.7 attack_power points (1.38 DPS) | yes | Assault Band (13095, -0.04 DPS) [world_drop]; Ironspine's Eye (7686, -0.16 DPS) [dungeon]; Field Researcher's Loop (281634, -0.17 DPS) [quest] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.34 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Ironspine's Eye (7686, -0.12 DPS) [dungeon]; Field Researcher's Loop (281634, -0.13 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (116.8 DPS) | yes | Scout's Blade (19544, -3.62 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -4.65 DPS) [vendor]; Coldrage Dagger (10761, -5.06 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 443.9 attack_power points (29.71 DPS) | yes | Coldrage Dagger (10761, -0.60 DPS, sim-verified) [dungeon]; Stonecloth Branch (15963, -29.51 DPS) [world_drop]; Tork Wrench (11855, -29.58 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (116.8 DPS) | yes | Swiftwind (13038, -0.20 DPS) [world_drop]; Monolithic Bow (9426, -0.22 DPS) [dungeon]; Bow of Searing Arrows (2825, -1.14 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; legs: Basilisk Hide Pants; feet: Skulker's Leather Shoes; finger1: Legionnaire's Band; finger2: Mark of Kern; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 32500000000000000-32530300001515201-0000000000000000000)

Set DPS (verified): 171.5. Weights run: 2.3s. Verify run: 2.2s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.533 ± 0.013, crit=0.893 ± 0.020 per rating point (14 rating = 1%, 12.502 per %), hit=1.596 ± 0.058 per rating point (10 rating = 1%, 15.959 per %), melee_haste=10.565 ± 0.857

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Ebon Mask (19984) | The Azure Key [quest] | 48.5 attack_power points (4.07 DPS) | yes | Blood Guard's Leather Headband (220851, -0.34 DPS) [vendor]; Embrace of the Lycan (9479, -1.15 DPS, sim-verified) [dungeon]; White Bandit Mask (10008, -1.73 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 24.9 attack_power points (2.09 DPS) | yes | Woven Ivy Necklace (19159, -0.43 DPS) [quest]; Scout's Medallion (19535, -0.55 DPS) [rep]; Zealous Shadowshard Pendant (17772, -0.64 DPS, sim-verified) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 28.9 attack_power points (2.42 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Skulker's Leather Shoulder (252535, -0.25 DPS) [crafted]; Failed Flying Experiment (9647, -0.30 DPS) [quest] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 27.5 attack_power points (2.30 DPS) | yes | Blisterbane Wrap (12552, -0.37 DPS) [dungeon]; Dark Phantom Cape (13122, -0.37 DPS) [world_drop]; Duskbat Drape (19982, -0.50 DPS) [quest] |
| chest | Stone Guard's Leather Armor (220855) | Lady Palanseer [vendor] | 46.5 attack_power points (3.89 DPS) | yes | Fungus Shroud Armor (17742, -0.68 DPS) [dungeon]; Blazewind Breastplate (11193, -0.69 DPS) [quest]; Warbear Harness (15064, -1.31 DPS, sim-verified) [crafted] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 27.0 attack_power points (2.26 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Skulker's Leather Bracers (252540, -0.68 DPS, sim-verified) [crafted] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 43.7 attack_power points (3.66 DPS) | yes | Darkmantle Grips (226828, -0.08 DPS) [vendor]; Gloves of Holy Might (867, -0.94 DPS) [world_drop]; Skulker's Leather Gauntlets (252548, -1.11 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (3.18 DPS) | yes | Skulker's Leather Waistguard (252474, -0.38 DPS) [crafted]; Defiler's Leather Girdle (20193, -0.46 DPS) [rep]; Prowler's Leather Waistguard (252473, -0.47 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 46.5 attack_power points (3.89 DPS) | yes | Basilisk Hide Pants (1718, -1.20 DPS) [world_drop]; Triprunner Dungarees (9624, -1.33 DPS) [quest]; Serpentskin Leggings (8262, -1.42 DPS, sim-verified) [world_drop] |
| feet | Sandstalker Ankleguards (12470) | Zul'Farrak: Zerillis [dungeon] | 32.1 attack_power points (2.69 DPS) | yes | Skulker's Leather Boots (252469, -0.09 DPS) [crafted]; Albino Crocscale Boots (17728, -0.12 DPS) [dungeon]; Prowler's Leather Boots (252468, -0.18 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 36.0 attack_power points (3.01 DPS) | yes | Legionnaire's Band (19511, -1.02 DPS) [rep]; Masons Fraternity Ring (9533, -1.22 DPS) [quest] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (2.01 DPS) | yes | Legionnaire's Band (19511, -0.02 DPS) [rep]; Masons Fraternity Ring (9533, -0.21 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (171.5 DPS) | yes | Frozen Heart of the Mountain (249469, -3.92 DPS, sim-verified) [crafted] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (171.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| main_hand | Shadowblade (2163) | World drop [world_drop] | sim-verified (171.5 DPS) | yes | Barman Shanker (12791, -2.85 DPS) [dungeon]; Scout's Blade (19543, -3.20 DPS) [pvp]; Hookfang Shanker (11635, -3.94 DPS, sim-verified) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (42.88 DPS) | yes | Barman Shanker (12791, -20.89 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -42.49 DPS) [quest]; Stonecloth Branch (15963, -42.62 DPS) [world_drop] |
| ranged | Precisely Calibrated Boomstick (2100) | World drop [world_drop] | sim-verified (171.5 DPS) | yes | Stinging Bow (10624, -0.37 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.37 DPS) [world_drop]; Dark Iron Rifle (16004, -2.14 DPS, sim-verified) [crafted] |

**New at 50:** head: Ebon Mask; neck: Skibi's Pendant; back: Blackveil Cape; chest: Stone Guard's Leather Armor; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Stone Guard's Leather Pants; feet: Sandstalker Ankleguards; finger1: Blackstone Ring; finger2: White Bone Band; trinket1: Rune of the Guard Captain; trinket2: Molten Heart of the Mountain; main_hand: Shadowblade; off_hand: Julie's Dagger; ranged: Precisely Calibrated Boomstick

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 32531000000000000-32530300001515201-5100000000000000000)

Set DPS (verified): 265.3. Weights run: 2.0s. Verify run: 6.5s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=2.083 ± 0.025, crit=2.239 ± 0.044 per rating point (14 rating = 1%, 31.353 per %), hit=3.314 ± 0.159 per rating point (10 rating = 1%, 33.141 per %), melee_haste=12.233 ± 2.173

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 97.6 attack_power points (7.95 DPS) | yes | Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Outlaw's Collar (279253, +0.00 DPS, sim-verified) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 64.4 attack_power points (5.24 DPS) | yes | Mark of Fordring (15411, -0.57 DPS) [quest]; Beads of Ogre Might (22150, -0.59 DPS) [quest]; Medallion of the Dawn (22659, -0.74 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-verified (265.3 DPS) | yes | Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, +0.00 DPS) [vendor]; Truestrike Shoulders (12927, -11.92 DPS, sim-verified) [dungeon] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 61.1 attack_power points (4.98 DPS) | yes | Cape of the Black Baron (13340, -0.81 DPS) [dungeon]; Deathguard's Cloak (20068, -1.36 DPS) [rep]; Windshear Cape (20691, -1.78 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-verified (265.3 DPS) | yes | Warlord's Leather Breastplate (231549, -1.87 DPS) [pvp]; Darkmantle Tunic (226825, -2.22 DPS) [quest]; Tunic of Undead Slaying (23089, -9.76 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-verified (265.3 DPS) | yes | General's Leather Armsplints (16559, -0.08 DPS) [pvp]; Bracers of the Eclipse (18375, -0.47 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.31 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-verified (265.3 DPS) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Devilsaur Gauntlets (15063, -0.42 DPS) [crafted]; Raider Gloves (272099, -14.06 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 110.0 attack_power points (8.96 DPS) | yes | Belt of Preserved Heads (20216, -1.31 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -3.64 DPS) [rep]; Ferocity of the Timbermaw (227805, -3.96 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-verified (265.3 DPS) | yes | Legionnaire's Leather Legguards (227059, +0.00 DPS) [pvp]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -10.93 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 73.2 attack_power points (5.96 DPS) | yes | Fine Dawn Treaders (227815, -1.27 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -1.59 DPS) [dungeon]; Darkmantle Boots (22003, -1.89 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (265.3 DPS) | yes | Tarnished Elven Ring (18500, -1.31 DPS) [dungeon]; Cutthroat's Signet (272408, -1.48 DPS) [vendor]; Naglering (11669, -5.18 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-verified (265.3 DPS) | yes | Tarnished Elven Ring (18500, -0.51 DPS) [dungeon]; Cutthroat's Signet (272408, -0.68 DPS) [vendor]; Naglering (11669, -4.72 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (265.3 DPS) | yes | Blackhand's Breadth (13965, +0.00 DPS) [quest]; Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (265.3 DPS) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Blackhand's Breadth (13965, -2.27 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -2.88 DPS) [crafted] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (265.3 DPS) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (265.3 DPS) | yes | The Lobotomizer (19324, -3.63 DPS, sim-verified) [rep]; Distracting Dagger (18392, -5.50 DPS) [dungeon]; Scepter of Interminable Focus (22329, -46.73 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-verified (265.3 DPS) | yes | Blackcrow (12651, -0.27 DPS) [dungeon]; The Purifier (22656, -0.65 DPS) [quest]; Dark Iron Rifle (16004, -2.32 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60, raid preset (troll, 00532010500000000-31530300001515231-0020000000000000000)

Set DPS (verified): 712.2. Weights run: 2.1s. Verify run: 6.7s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.100 ± 0.001, agility=2.504 ± 0.030, crit=2.686 ± 0.050 per rating point (14 rating = 1%, 37.604 per %), hit=3.804 ± 0.235 per rating point (10 rating = 1%, 38.044 per %), melee_haste=25.013 ± 3.160

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) | Stratholme: The Unforgiven [dungeon] | 113.7 attack_power points (19.06 DPS) | yes | Warlord's Leather Helm (231553, +0.00 DPS) [pvp]; Champion's Leather Helm (227057, -0.34 DPS) [pvp]; Outlaw's Collar (279253, -0.79 DPS) [crafted] |
| neck | Pendant of Celerity (22340) | Blackrock Spire: Lord Valthalak [dungeon] | 75.6 attack_power points (12.67 DPS) | yes | Mark of Fordring (15411, -2.01 DPS) [quest]; Beads of Ogre Might (22150, -2.27 DPS) [quest]; Medallion of the Dawn (22659, -2.35 DPS) [quest] |
| shoulder | Stormshroud Shoulders (15058) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Truestrike Shoulders (12927, +0.00 DPS) [dungeon]; Champion's Leather Shoulders (227056, +0.00 DPS) [pvp]; Darkspear Pauldrons (272105, -36.76 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | 66.0 attack_power points (11.07 DPS) | yes | Cape of the Black Baron (13340, -1.42 DPS) [dungeon]; Deathguard's Cloak (20068, -3.27 DPS) [rep]; Windshear Cape (20691, -3.30 DPS) [world] |
| chest | Dawn Armor (252483) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Leather Breastplate (231549, -4.62 DPS) [pvp]; Darkmantle Tunic (226825, -5.63 DPS) [quest]; Tunic of Undead Slaying (23089, -27.28 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Armsplints (16559, -0.18 DPS) [pvp]; Bracers of the Eclipse (18375, -1.78 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -11.21 DPS, sim-verified) [world] |
| hands | Stormshroud Gloves (21278) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Leather Mitts (231555, +0.00 DPS) [pvp]; Darkmantle Gloves (22006, -1.23 DPS) [quest]; Raider Gloves (272099, -42.56 DPS, sim-verified) [vendor] |
| waist | Assassin's Waistguard (272395) | Pix Xizzix [vendor] | 128.7 attack_power points (21.57 DPS) | yes | Belt of Preserved Heads (20216, -5.05 DPS, sim-verified) [quest]; Defiler's Leather Girdle (20190, -9.57 DPS) [rep]; Ferocity of the Timbermaw (227805, -9.64 DPS) [vendor] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguehound Leggings (18736, +0.00 DPS) [dungeon]; General's Leather Legguards (231554, +0.00 DPS) [pvp]; Sentinel's Leather Pants (237818, -36.76 DPS, sim-verified) [vendor] |
| feet | Darkmantle Footpads (226831) | Mokvar [vendor] | 86.7 attack_power points (14.54 DPS) | yes | Fine Dawn Treaders (227815, -3.63 DPS, sim-verified) [vendor]; Shadowcraft Boots (16711, -3.81 DPS) [dungeon]; Darkmantle Boots (22003, -4.46 DPS) [quest] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -2.69 DPS) [dungeon]; Cutthroat's Signet (272408, -3.11 DPS) [vendor]; Naglering (11669, -14.45 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tarnished Elven Ring (18500, -1.26 DPS) [dungeon]; Cutthroat's Signet (272408, -1.68 DPS) [vendor]; Naglering (11669, -13.72 DPS, sim-verified) [dungeon] |
| trinket1 | Blackhand's Breadth (13965) | For The Horde! [quest] | sim-verified (+10.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Royal Seal of Eldre'Thalas (18465, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -6.87 DPS) [crafted]; Counterattack Lodestone (18537, -8.92 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, -5.29 DPS, sim-verified) [quest]; Frozen Heart of the Mountain (249469, -5.76 DPS) [crafted]; Counterattack Lodestone (18537, -7.82 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+22.0 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (712.2 DPS) | yes | The Lobotomizer (19324, -7.75 DPS, sim-verified) [rep]; Distracting Dagger (18392, -11.31 DPS) [dungeon]; Scepter of Interminable Focus (22329, -94.30 DPS) [dungeon] |
| ranged | Satyr's Bow (18323) | Dire Maul: Zevrim Thornhoof [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackcrow (12651, -0.64 DPS) [dungeon]; The Purifier (22656, -1.33 DPS) [quest]; Dark Iron Rifle (16004, -7.08 DPS, sim-verified) [crafted] |

**New at 60:** head: Mask of the Unforgiven; neck: Pendant of Celerity; shoulder: Stormshroud Shoulders; back: Howler's Furs; chest: Dawn Armor; wrist: Forest Stalker's Bracers; hands: Stormshroud Gloves; waist: Assassin's Waistguard; legs: Stormshroud Pants; feet: Darkmantle Footpads; finger1: Don Julio's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Blackhand's Breadth; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Satyr's Bow

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

