# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 31.3. Weights run: 2.5s. Verify run: 1.6s. 197 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.005 ± 0.001, crit=0.012 ± 0.001 per rating point (14 rating = 1%, 0.164 per %), hit=0.031 ± 0.001 per rating point (10 rating = 1%, 0.306 per %), melee_haste=0.709 ± 0.071

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.0 attack_power points (0.40 DPS) | yes | Defender's Leather Hood (252447, -0.00 DPS) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.0 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 attack_power points (0.25 DPS) | yes | Slime-encrusted Pads (6461, -0.31 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.0 attack_power points (0.30 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.0 attack_power points (0.61 DPS) | yes | Tunic of Westfall (2041, -0.05 DPS) [quest]; Defender's Leather Armor (252434, -0.10 DPS) [crafted]; Prospector's Chestpiece (14562, -0.20 DPS) [world_drop] |
| wrist | Bravo's Armbands (270015) | Underground Assault [quest] | 6.0 attack_power points (0.30 DPS) | yes | Forest Leather Bracers (3202, -0.05 DPS) [world_drop]; Bristlebark Bindings (14569, -0.05 DPS) [world_drop]; Wolf Bracers (4794, -0.10 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.0 attack_power points (0.50 DPS) | yes | Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Bristlebark Gloves (14572, -0.12 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.91 DPS) | yes | Brawler's Leather Belt (252428, -0.50 DPS) [crafted]; Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.60 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.0 attack_power points (0.76 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.10 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.25 DPS) [dungeon]; Footpads of the Fang (10411, -0.25 DPS) [dungeon] |
| finger1 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.0 attack_power points (0.40 DPS) | yes | Pyrewood Signet Ring (277210, -0.20 DPS) [quest]; Demon Band (12054, -0.20 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.0 attack_power points (0.30 DPS) | yes | Pyrewood Signet Ring (277210, -0.10 DPS, sim-verified) [quest]; Demon Band (12054, -0.10 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (11.51 DPS) | yes | Evocator's Blade (2567, -0.72 DPS) [dungeon]; Buzzer Blade (2169, -1.64 DPS) [dungeon]; Deadly Bronze Poniard (3490, -2.30 DPS) [crafted] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.4 attack_power points (11.30 DPS) | yes | Evocator's Blade (2567, -5.74 DPS, sim-verified) [dungeon] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Bravo's Armbands; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Protector's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 197, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 45.5. Weights run: 2.8s. Verify run: 1.7s. 331 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.008 ± 0.001, crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.237 per %), hit=0.041 ± 0.001 per rating point (10 rating = 1%, 0.413 per %), melee_haste=0.958 ± 0.143

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.63 DPS) | yes | Brawler's Leather Helm (252512, -0.10 DPS) [crafted]; Cloudy Gustwoven Hood (277042, -0.16 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.74 DPS) | yes | Kaleidoscope Chain (13084, -0.31 DPS) [world_drop]; Sentinel's Medallion (19541, -0.36 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.85 DPS) | yes | Barbaric Shoulders (5964, -0.32 DPS) [crafted]; Bristlebark Amice (14573, -0.37 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.53 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Tigerstrike Mantle (13108, -0.10 DPS) [world_drop]; Sergeant Major's Cape (16315, -0.11 DPS) [pvp] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 attack_power points (0.84 DPS) | yes | Dusky Leather Armor (7374, -0.10 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.10 DPS) [crafted]; Brawler's Leather Armor (252490, -0.21 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.0 attack_power points (0.53 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.11 DPS) [world]; Barbaric Bracers (18948, -0.11 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.84 DPS) | yes | Insignia Gloves (6408, -0.17 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -0.21 DPS) [crafted]; Wolfclaw Gloves (1978, -0.26 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 attack_power points (1.26 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.32 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.47 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.37 DPS) | yes | Brawler's Leather Pants (252500, -0.57 DPS) [crafted]; Trapper's Leather Pants (252501, -0.57 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.61 DPS, sim-verified) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 attack_power points (0.63 DPS) | yes | Feet of the Lynx (1121, -0.05 DPS) [world_drop]; Brawler's Leather Boots (252439, -0.10 DPS) [crafted]; Insignia Boots (4055, -0.21 DPS) [world_drop] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.69 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Insurgent's Band (272067, -0.21 DPS) [vendor]; Monkey Ring (6748, -0.32 DPS) [quest] |
| finger2 | Protector's Band (19517) | Silverwing Sentinels [rep] | 12.0 attack_power points (0.63 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.16 DPS) [vendor]; Monkey Ring (6748, -0.26 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (16.91 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -1.15 DPS) [vendor]; Torturing Poker (7682, -1.34 DPS) [dungeon]; Thornspike (6681, -1.77 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (16.81 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.27 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -16.76 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.47 DPS) | yes | Double-barreled Shotgun (2098, -0.18 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.21 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Raptorbane Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Protector's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 331, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 88.2. Weights run: 3.0s. Verify run: 1.8s. 459 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.011 ± 0.001, crit=0.024 ± 0.002 per rating point (14 rating = 1%, 0.330 per %), hit=0.035 ± 0.002 per rating point (10 rating = 1%, 0.347 per %), melee_haste=not significant (1.177 ± 0.596)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.1 attack_power points (1.27 DPS) | yes | Warden's Wizard Hat (14604, -0.51 DPS) [world_drop]; Hawkeye's Helm (14591, -0.56 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.57 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.14 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS, sim-verified) [dungeon]; Sentinel's Medallion (19540, -0.51 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.32 DPS) | yes | Flintrock Shoulders (7755, -0.46 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.69 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.1 attack_power points (0.81 DPS) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Wolfmaster Cape (6314, -0.24 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.30 DPS, sim-verified) [pvp] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.2 attack_power points (1.56 DPS) | yes | Raptorbane Armor (3566, -0.64 DPS) [quest]; Nightscape Tunic (8175, -0.69 DPS) [crafted]; Wolffear Harness (13110, -0.87 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.14 DPS) | yes | Cultist's Armguards (270032, -0.57 DPS) [quest]; Hawkeye's Bracers (14590, -0.62 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.68 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.3 attack_power points (1.16 DPS) | yes | Prowler's Leather Gloves (252524, -0.07 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.70 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 attack_power points (1.72 DPS) | yes | Highlander's Chain Girdle (20090, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.68 DPS) [world_drop]; Blackened Defias Belt (10403, -0.69 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.49 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.28 DPS) [quest]; Brawler's Leather Legguards (252516, -0.57 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.1 attack_power points (1.04 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.11 DPS) [dungeon]; Excelsior Boots (4109, -0.12 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.14 DPS) | yes | Protector's Band (19515, -0.22 DPS) [rep]; Field Researcher's Loop (281634, -0.34 DPS) [quest]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.14 DPS) | yes | Protector's Band (19515, -0.22 DPS) [rep]; Field Researcher's Loop (281634, -0.34 DPS) [quest]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (88.2 DPS) | yes | Black Menace (6831, -3.38 DPS) [quest]; Darkspear Insurgent's Spellblade (272085, -3.97 DPS) [vendor]; Coldrage Dagger (10761, -8.62 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 438.7 attack_power points (25.11 DPS) | yes | Black Menace (6831, -1.11 DPS, sim-verified) [quest]; Stonecloth Branch (15963, -24.94 DPS) [world_drop]; Satyr's Rod (15962, -25.06 DPS) [world_drop] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (88.2 DPS) | yes | Monolithic Bow (9426, -0.28 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.29 DPS) [vendor]; Bow of Searing Arrows (2825, -1.13 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 121.5. Weights run: 3.0s. Verify run: 2.1s. 584 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.108 ± 0.005, crit=0.218 ± 0.008 per rating point (14 rating = 1%, 3.046 per %), hit=0.051 ± 0.004 per rating point (10 rating = 1%, 0.505 per %), melee_haste=not significant (1.118 ± 0.969)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.42 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -1.02 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -1.24 DPS) [vendor] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.21 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Ghostshard Talisman (7731, -0.36 DPS) [dungeon]; Sentinel's Medallion (19539, -0.41 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (1.47 DPS) | yes | Skulker's Leather Shoulder (252535, -0.18 DPS) [crafted]; Failed Flying Experiment (9647, -0.19 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.19 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.5 attack_power points (1.30 DPS) | yes | Blisterbane Wrap (12552, -0.30 DPS) [dungeon]; Duskbat Drape (19982, -0.36 DPS) [quest]; Dark Phantom Cape (13122, -0.42 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 30.9 attack_power points (1.87 DPS) | yes | Mixologist's Tunic (12793, -0.05 DPS) [dungeon]; Quillward Harness (10583, -0.11 DPS) [dungeon]; Blazewind Breastplate (11193, -0.15 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 20.6 attack_power points (1.25 DPS) | yes | Skulker's Leather Bracers (252540, -0.22 DPS) [crafted]; Prowler's Leather Bracers (252539, -0.23 DPS) [crafted]; Branded Leather Bracers (19508, -0.58 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 34.4 attack_power points (2.08 DPS) | yes | Darkmantle Grips (226828, -0.06 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -0.60 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.63 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.30 DPS) | yes | Skulker's Leather Waistguard (252474, -0.64 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.65 DPS) [crafted]; Highlander's Leather Girdle (20116, -0.73 DPS, sim-verified) [rep] |
| legs | Gryphon Rider's Leggings (9652) | Saving Sharpbeak [quest] | 31.6 attack_power points (1.92 DPS) | yes | Serpentskin Leggings (8262, -0.05 DPS) [world_drop]; Ferine Leggings (6690, -0.34 DPS) [dungeon]; Basilisk Hide Pants (1718, -0.51 DPS) [world_drop] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 25.4 attack_power points (1.54 DPS) | yes | Prowler's Leather Boots (252468, -0.01 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.03 DPS) [dungeon]; Shadefiend Boots (11675, -0.19 DPS) [dungeon] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.5 attack_power points (1.24 DPS) | yes | Assault Band (13095, -0.03 DPS) [world_drop]; Protector's Band (19516, -0.03 DPS) [rep] |
| finger2 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.21 DPS) | yes | Assault Band (13095, +0.00 DPS) [world_drop]; Protector's Band (19516, -0.00 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (121.5 DPS) | yes | - |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (121.5 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (121.5 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Lifeforce Dirk (10750, -0.48 DPS) [quest]; Searing Needle (12531, -5.82 DPS, sim-verified) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (30.99 DPS) | yes | Thermotastic Egg Timer (9644, -30.79 DPS) [quest]; Stonecloth Branch (15963, -30.81 DPS) [world_drop]; Windchaser Orb (15965, -30.93 DPS) [world_drop] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (121.5 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.09 DPS) [world_drop]; Dark Iron Rifle (16004, -1.56 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Gryphon Rider's Leggings; feet: Skulker's Leather Boots; finger1: Blackstone Ring; finger2: Mark of Kern; trinket1: Frozen Heart of the Mountain; trinket2: Molten Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 584, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 188.6. Weights run: 3.0s. Verify run: 2.2s. 1370 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.136 ± 0.007, crit=0.285 ± 0.011 per rating point (14 rating = 1%, 3.996 per %), hit=0.063 ± 0.006 per rating point (10 rating = 1%, 0.627 per %), melee_haste=not significant (1.050 ± 1.577)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.39 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Lieutenant Commander's Leather Helm (227055, +0.00 DPS) [pvp]; Field Marshal's Leather Mask (231545, -0.28 DPS) [pvp] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.91 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Will of the Martyr (17044, -0.12 DPS) [quest]; Medallion of the Dawn (22659, -0.24 DPS) [quest] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 50.4 attack_power points (3.01 DPS) | yes | Highlander's Lizardhide Shoulders (20060, -0.88 DPS, sim-verified) [rep]; Darkspear Pauldrons (272105, -0.94 DPS) [vendor]; Dark Warder's Pauldrons (22241, -1.20 DPS) [dungeon] |
| back | Cloak of the Honor Guard (20073) | The League of Arathor [rep] | 39.7 attack_power points (2.37 DPS) | yes | Cape of the Black Baron (13340, -0.16 DPS) [dungeon]; Howler's Furs (272414, -0.66 DPS) [vendor]; Windshear Cape (20691, -0.87 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.17 DPS) [crafted]; Nightbrace Tunic (12603, -1.28 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.91 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.17 DPS) [rep]; Marshal's Leather Armsplints (16460, -0.23 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.54 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (2.63 DPS) | yes | Knight-Lieutenant's Leather Gauntlets (16396, -0.04 DPS) [pvp]; Raider Gloves (272099, -0.14 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.24 DPS) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 45.3 attack_power points (2.70 DPS) | yes | Marshal's Leather Cinch (16458, -0.14 DPS) [pvp]; Girdle of Beastial Fury (11686, -0.44 DPS) [dungeon]; Cadaverous Belt (14636, -1.22 DPS, sim-verified) [dungeon] |
| legs | Cadaverous Leggings (14638) | Scholomance: Lady Illucia Barov [dungeon] | 52.0 attack_power points (3.10 DPS) | yes | Devilsaur Leggings (15062, -0.12 DPS) [crafted]; Warbear Woolies (15065, -0.21 DPS) [crafted]; Sentinel's Leather Pants (237818, -0.80 DPS) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.39 DPS) | yes | Drudge Boots (21532, -0.46 DPS) [quest]; Dunestalker's Boots (20715, -0.56 DPS) [quest]; Highlander's Leather Boots (20052, -0.62 DPS) [rep] |
| finger1 | Protector's Band (19514) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.23 DPS) [dungeon]; Don Julio's Band (19325, -0.23 DPS) [rep]; Naglering (11669, -2.65 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21201) | The Path of the Conqueror [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blackstone Ring (17713, -0.03 DPS) [dungeon]; Don Julio's Band (19325, -0.03 DPS) [rep]; Naglering (11669, -4.26 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Royal Seal of Eldre'Thalas (18465) | Garona: A Study on Stealth and Treachery [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+5.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Mageblade (234574, +0.00 DPS) [pvp]; Grand Marshal's Dirk (234582, +0.00 DPS) [pvp]; Grand Marshal's Shiv (235479, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (188.6 DPS) | yes | Distracting Dagger (18392, -4.03 DPS) [dungeon]; The Lobotomizer (19324, -4.52 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -37.08 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.12 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.30 DPS) [world_drop]; Dark Iron Rifle (16004, -1.91 DPS, sim-verified) [crafted] |

**New at 60:** neck: Imperial Jewel; shoulder: Highlander's Leather Shoulders; back: Cloak of the Honor Guard; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Cadaverous Gloves; waist: Ferocity of the Timbermaw; legs: Cadaverous Leggings; feet: Pads of the Dread Wolf; finger1: Protector's Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Darkmoon Card: Maelstrom; trinket2: Royal Seal of Eldre'Thalas; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1370, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4765 Enamelled Broadsword; 4797 Fiery Cloak; 4798 Heavy Runed Cloak

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.9. Weights run: 2.5s. Verify run: 1.6s. 190 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.005 ± 0.001, crit=0.012 ± 0.001 per rating point (14 rating = 1%, 0.164 per %), hit=0.031 ± 0.001 per rating point (10 rating = 1%, 0.306 per %), melee_haste=0.709 ± 0.071

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.0 attack_power points (0.40 DPS) | yes | Defender's Leather Hood (252447, -0.00 DPS) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.0 attack_power points (0.30 DPS) | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.0 attack_power points (0.25 DPS) | yes | Slime-encrusted Pads (6461, -0.30 DPS, sim-verified) [dungeon] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.0 attack_power points (0.30 DPS) | yes | Catacomb Cloak (279899, -0.00 DPS) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Dark Leather Cloak (2316, -0.05 DPS) [crafted] |
| chest | Brawler's Leather Armor (252490) | Leatherworking [crafted] | 12.0 attack_power points (0.61 DPS) | yes | Defender's Leather Armor (252434, -0.12 DPS, sim-verified) [crafted]; Prospector's Chestpiece (14562, -0.20 DPS) [world_drop]; Murloc Scale Breastplate (5781, -0.20 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.0 attack_power points (0.25 DPS) | yes | Bristlebark Bindings (14569, -0.00 DPS) [world_drop]; Wolf Bracers (4794, -0.05 DPS) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Gloves of the Fang (10413) | Wailing Caverns: Druid of the Fang [dungeon] | 10.0 attack_power points (0.50 DPS) | yes | Brawler's Leather Gloves (252494, -0.10 DPS) [crafted]; Bristlebark Gloves (14572, -0.12 DPS, sim-verified) [world_drop]; Gold-flecked Gloves (5195, -0.15 DPS) [dungeon] |
| waist | Blackened Defias Belt (10403) | The Deadmines: Captain Greenskin [dungeon] | 18.0 attack_power points (0.91 DPS) | yes | Brawler's Leather Belt (252428, -0.50 DPS) [crafted]; Deviate Scale Belt (6468, -0.59 DPS, sim-verified) [crafted]; Ruffian Belt (5975, -0.60 DPS) [world] |
| legs | Brawler's Leather Pants (252500) (or Trapper's Leather Pants (252501)) | Leatherworking [crafted] | 15.0 attack_power points (0.76 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Leggings of the Fang (10410, -0.05 DPS) [dungeon]; Defender's Leather Pants (252445, -0.10 DPS) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.0 attack_power points (0.56 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Blackened Defias Boots (10402, -0.25 DPS) [dungeon]; Footpads of the Fang (10411, -0.25 DPS) [dungeon] |
| finger1 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.0 attack_power points (0.40 DPS) | yes | Pyrewood Signet Ring (277210, -0.20 DPS) [quest]; Demon Band (12054, -0.20 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.25 DPS) [quest] |
| finger2 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.0 attack_power points (0.30 DPS) | yes | Pyrewood Signet Ring (277210, -0.09 DPS, sim-verified) [quest]; Demon Band (12054, -0.10 DPS) [world_drop]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 attack_power points (11.51 DPS) | yes | Edward's Knife (251485, -0.30 DPS) [quest]; Scout's Blade (20441, -0.53 DPS) [pvp]; Evocator's Blade (2567, -0.72 DPS) [dungeon] |
| off_hand | Assassin's Blade (1935) | Shadowfang Keep: Son of Arugal [dungeon] | 224.4 attack_power points (11.30 DPS) | yes | Edward's Knife (251485, +0.00 DPS) [quest]; Tork Wrench (11855, -11.20 DPS) [quest] |
| ranged | Lil Timmy's Peashooter (13136) | World drop [world_drop] | 4.0 attack_power points (0.20 DPS) | yes | Fine Longbow (11304, -0.00 DPS) [vendor]; Deadly Blunderbuss (4369, -0.10 DPS) [crafted]; Light Bow (4576, -0.10 DPS) [world_drop] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Gloves of the Fang; waist: Blackened Defias Belt; legs: Brawler's Leather Pants; feet: Feet of the Lynx; finger1: Legionnaire's Band; finger2: Signet of the Zhevra; main_hand: Blackfang; off_hand: Assassin's Blade; ranged: Lil Timmy's Peashooter

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18857 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 44.9. Weights run: 2.8s. Verify run: 1.7s. 322 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.008 ± 0.001, crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.237 per %), hit=0.041 ± 0.001 per rating point (10 rating = 1%, 0.413 per %), melee_haste=0.958 ± 0.143

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Defender's Leather Helm (252455) | Leatherworking [crafted] | 12.0 attack_power points (0.63 DPS) | yes | Brawler's Leather Helm (252512, -0.12 DPS, sim-verified) [crafted]; Cloudy Gustwoven Hood (277042, -0.16 DPS) [crafted]; Azure Gustwoven Hood (277050, -0.16 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 attack_power points (0.74 DPS) | yes | Kaleidoscope Chain (13084, -0.31 DPS) [world_drop]; Scout's Medallion (19537, -0.37 DPS, sim-verified) [rep] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 16.1 attack_power points (0.85 DPS) | yes | Barbaric Shoulders (5964, -0.32 DPS) [crafted]; Bristlebark Amice (14573, -0.37 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Hawkeye's Cloak (14593) | World drop [world_drop] | 10.1 attack_power points (0.53 DPS) | yes | Wolfmaster Cape (6314, -0.00 DPS) [dungeon]; Wildhunter Cloak (16658, -0.00 DPS) [quest]; Tigerstrike Mantle (13108, -0.10 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 attack_power points (0.74 DPS) | yes | Brawler's Leather Tunic (252508, -0.00 DPS) [crafted]; Brawler's Leather Armor (252490, -0.11 DPS) [crafted]; Defender's Leather Tunic (252450, -0.11 DPS) [crafted] |
| wrist | Hawkeye's Bracers (14590) | World drop [world_drop] | 10.0 attack_power points (0.53 DPS) | yes | Cultist's Armguards (270032, -0.00 DPS) [quest]; Jurassic Wristguards (6198, -0.11 DPS) [world]; Barbaric Bracers (18948, -0.11 DPS) [crafted] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 attack_power points (0.84 DPS) | yes | Insignia Gloves (6408, -0.18 DPS, sim-verified) [world_drop]; Toughened Leather Gloves (4253, -0.21 DPS) [crafted]; Wolfclaw Gloves (1978, -0.26 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 attack_power points (1.26 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS) [rep]; Blackened Defias Belt (10403, -0.32 DPS) [dungeon]; Deftkin Belt (16659, -0.42 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.37 DPS) | yes | Brawler's Leather Pants (252500, -0.57 DPS) [crafted]; Trapper's Leather Pants (252501, -0.57 DPS) [crafted]; Brawler's Leather Legguards (252516, -0.62 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 11.1 attack_power points (0.58 DPS) | yes | Brawler's Leather Boots (252439, -0.05 DPS) [crafted]; Insignia Boots (4055, -0.16 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.16 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 13.1 attack_power points (0.69 DPS) | yes | Thunderbrow Ring (13097, -0.11 DPS) [world_drop]; Insurgent's Band (272067, -0.21 DPS) [vendor]; Band of the Fist (17694, -0.26 DPS) [quest] |
| finger2 | Legionnaire's Band (19513) | Warsong Outriders [rep] | 12.0 attack_power points (0.63 DPS) | yes | Thunderbrow Ring (13097, -0.05 DPS) [world_drop]; Insurgent's Band (272067, -0.16 DPS) [vendor]; Band of the Fist (17694, -0.21 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 attack_power points (16.91 DPS) | yes | Scout's Blade (19545, -0.97 DPS) [pvp]; Darkspear Insurgent's Spellblade (272086, -1.15 DPS) [vendor]; Torturing Poker (7682, -1.34 DPS) [dungeon] |
| off_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | 320.0 attack_power points (16.81 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -5.17 DPS, sim-verified) [vendor]; Tork Wrench (11855, -16.70 DPS) [quest]; Satyr's Rod (15962, -16.76 DPS) [world_drop] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 9.0 attack_power points (0.47 DPS) | yes | Double-barreled Shotgun (2098, -0.19 DPS, sim-verified) [world_drop]; Silver Star (3463, -0.21 DPS) [quest]; BKP "Sparrow" Smallbore (3042, -0.26 DPS) [world_drop] |

**New at 30:** head: Defender's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Hawkeye's Cloak; chest: Dusky Leather Armor; wrist: Hawkeye's Bracers; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; finger1: Ironspine's Eye; finger2: Legionnaire's Band; main_hand: Swinetusk Shank; off_hand: Scorn's Focal Dagger; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5748 Centaur Longbow; 5821 Darkstalker Boots

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 86.7. Weights run: 3.0s. Verify run: 1.8s. 444 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.011 ± 0.001, crit=0.024 ± 0.002 per rating point (14 rating = 1%, 0.330 per %), hit=0.035 ± 0.002 per rating point (10 rating = 1%, 0.347 per %), melee_haste=not significant (1.177 ± 0.596)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | White Bandit Mask (10008) | Tailoring [crafted] | 22.1 attack_power points (1.27 DPS) | yes | Warden's Wizard Hat (14604, -0.51 DPS) [world_drop]; Hawkeye's Helm (14591, -0.56 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -0.57 DPS) [crafted] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.14 DPS) | yes | Ghostshard Talisman (7731, -0.48 DPS, sim-verified) [dungeon]; Scout's Medallion (19536, -0.51 DPS) [rep]; Ethereal Talisman (4430, -0.63 DPS) [quest] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 attack_power points (1.32 DPS) | yes | Flintrock Shoulders (7755, -0.46 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.69 DPS) [crafted] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 14.1 attack_power points (0.81 DPS) | yes | Hawkeye's Cloak (14593, -0.23 DPS) [world_drop]; Wildhunter Cloak (16658, -0.24 DPS) [quest]; First Sergeant's Cloak (16340, -0.27 DPS, sim-verified) [pvp] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 27.2 attack_power points (1.56 DPS) | yes | Nightscape Tunic (8175, -0.69 DPS) [crafted]; Barbaric Harness (5739, -0.69 DPS) [crafted]; Wolffear Harness (13110, -0.86 DPS, sim-verified) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 attack_power points (1.14 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Cultist's Armguards (270032, -0.57 DPS) [quest]; Hawkeye's Bracers (14590, -0.58 DPS, sim-verified) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 20.3 attack_power points (1.16 DPS) | yes | Prowler's Leather Gloves (252524, -0.07 DPS) [crafted]; Imperial Leather Gloves (4063, -0.13 DPS) [dungeon]; Skulker's Leather Gloves (252525, -0.67 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 attack_power points (1.72 DPS) | yes | Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Ogron's Sash (13117, -0.68 DPS) [world_drop]; Blackened Defias Belt (10403, -0.69 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 attack_power points (1.49 DPS) | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.28 DPS) [quest]; Brawler's Leather Legguards (252516, -0.57 DPS) [crafted] |
| feet | Skulker's Leather Shoes (252531) | Leatherworking [crafted] | 18.1 attack_power points (1.04 DPS) | yes | Prowler's Leather Shoes (252465, -0.00 DPS) [crafted]; Imperial Leather Boots (6431, -0.11 DPS) [dungeon]; Excelsior Boots (4109, -0.12 DPS) [quest] |
| finger1 | Mark of Kern (2262) (or Assault Band (13095)) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 20.0 attack_power points (1.14 DPS) | yes | Legionnaire's Band (19512, -0.22 DPS) [rep]; Field Researcher's Loop (281634, -0.34 DPS) [quest]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| finger2 | Assault Band (13095) | World drop [world_drop] | 20.0 attack_power points (1.14 DPS) | yes | Legionnaire's Band (19512, -0.22 DPS) [rep]; Field Researcher's Loop (281634, -0.34 DPS) [quest]; Ironspine's Eye (7686, -0.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gut Ripper (2164) | World drop [world_drop] | sim-verified (86.7 DPS) | yes | Scout's Blade (19544, -3.36 DPS) [pvp]; Darkspear Insurgent's Spellblade (272085, -3.97 DPS) [vendor]; Coldrage Dagger (10761, -8.19 DPS, sim-verified) [dungeon] |
| off_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 438.7 attack_power points (25.11 DPS) | yes | Coldrage Dagger (10761, -0.63 DPS, sim-verified) [dungeon]; Stonecloth Branch (15963, -24.94 DPS) [world_drop]; Tork Wrench (11855, -25.00 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (86.7 DPS) | yes | Monolithic Bow (9426, -0.28 DPS) [dungeon]; Booty Bay Bruiser's Buckshot (274748, -0.29 DPS) [vendor]; Bow of Searing Arrows (2825, -1.11 DPS, sim-verified) [world_drop] |

**New at 40:** head: White Bandit Mask; neck: Zealous Shadowshard Pendant; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Skulker's Leather Shoes; finger1: Mark of Kern; finger2: Assault Band; main_hand: Gut Ripper; off_hand: Jhordy's Misplaced Screwdriver; ranged: The Silencer

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 123.8. Weights run: 3.0s. Verify run: 2.2s. 564 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.108 ± 0.005, crit=0.218 ± 0.008 per rating point (14 rating = 1%, 3.046 per %), hit=0.051 ± 0.004 per rating point (10 rating = 1%, 0.505 per %), melee_haste=not significant (1.118 ± 0.969)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.42 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; White Bandit Mask (10008, -1.02 DPS) [crafted]; Undercity Reservist's Cap (20643, -1.08 DPS) [quest] |
| neck | Zealous Shadowshard Pendant (17772) | Shadowshard Fragments [quest] | 20.0 attack_power points (1.21 DPS) | yes | Skibi's Pendant (13089, +0.00 DPS, sim-verified) [world_drop]; Woven Ivy Necklace (19159, -0.24 DPS) [quest]; Ghostshard Talisman (7731, -0.36 DPS) [dungeon] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 attack_power points (1.47 DPS) | yes | Skulker's Leather Shoulder (252535, -0.18 DPS) [crafted]; Failed Flying Experiment (9647, -0.19 DPS) [quest]; Prowler's Leather Shoulder (252534, -0.19 DPS) [crafted] |
| back | Blackveil Cape (11626) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.5 attack_power points (1.30 DPS) | yes | Blisterbane Wrap (12552, -0.30 DPS) [dungeon]; Duskbat Drape (19982, -0.36 DPS) [quest]; Dark Phantom Cape (13122, -0.42 DPS, sim-verified) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 30.9 attack_power points (1.87 DPS) | yes | Mixologist's Tunic (12793, -0.05 DPS) [dungeon]; Quillward Harness (10583, -0.11 DPS) [dungeon]; Blazewind Breastplate (11193, -0.15 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 20.6 attack_power points (1.25 DPS) | yes | Forest Stalker's Bracers (19589, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -0.57 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272100) | Creeg Bothunk [vendor] | 34.4 attack_power points (2.08 DPS) | yes | Darkmantle Grips (226828, -0.06 DPS) [vendor]; Skulker's Leather Gauntlets (252548, -0.60 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -0.63 DPS) [crafted] |
| waist | Girdle of Beastial Fury (11686) | Blackrock Depths: Eviscerator [dungeon] | 38.0 attack_power points (2.30 DPS) | yes | Skulker's Leather Waistguard (252474, -0.64 DPS) [crafted]; Prowler's Leather Waistguard (252473, -0.65 DPS) [crafted]; Defiler's Leather Girdle (20192, -0.73 DPS, sim-verified) [rep] |
| legs | Serpentskin Leggings (8262) | World drop [world_drop] | 30.7 attack_power points (1.86 DPS) | yes | Basilisk Hide Pants (1718, -0.45 DPS) [world_drop]; Triprunner Dungarees (9624, -0.47 DPS) [quest]; Ferine Leggings (6690, -0.97 DPS, sim-verified) [dungeon] |
| feet | Skulker's Leather Boots (252469) | Leatherworking [crafted] | 25.4 attack_power points (1.54 DPS) | yes | Prowler's Leather Boots (252468, -0.01 DPS) [crafted]; Sandstalker Ankleguards (12470, -0.03 DPS) [dungeon]; Shadefiend Boots (11675, -0.19 DPS) [dungeon] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 attack_power points (1.45 DPS) | yes | Mark of Kern (2262, -0.24 DPS) [dungeon]; Assault Band (13095, -0.24 DPS) [world_drop]; Legionnaire's Band (19511, -0.24 DPS) [rep] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.5 attack_power points (1.24 DPS) | yes | Assault Band (13095, -0.03 DPS) [world_drop]; Legionnaire's Band (19511, -0.03 DPS) [rep]; Mark of Kern (2262, -0.93 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (123.8 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (123.8 DPS) | yes | Molten Heart of the Mountain (249470, -0.42 DPS, sim-verified) [crafted] |
| main_hand | Barman Shanker (12791) | Blackrock Depths: Plugger Spazzring [dungeon] | sim-verified (123.8 DPS) | yes | Shadowblade (2163, +0.00 DPS) [world_drop]; Lifeforce Dirk (10750, -0.48 DPS) [quest]; Searing Needle (12531, -5.42 DPS, sim-verified) [dungeon] |
| off_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 attack_power points (30.99 DPS) | yes | Thermotastic Egg Timer (9644, -30.79 DPS) [quest]; Stonecloth Branch (15963, -30.81 DPS) [world_drop]; Tork Wrench (11855, -30.87 DPS) [quest] |
| ranged | Skull Splitting Crossbow (13039) | World drop [world_drop] | sim-verified (123.8 DPS) | yes | Stinging Bow (10624, +0.00 DPS) [dungeon]; Precisely Calibrated Boomstick (2100, -0.09 DPS) [world_drop]; Dark Iron Rifle (16004, -1.54 DPS, sim-verified) [crafted] |

**New at 50:** head: Embrace of the Lycan; back: Blackveil Cape; chest: Warbear Harness; wrist: Deepfury Bracers; hands: Raider Gloves; waist: Girdle of Beastial Fury; legs: Serpentskin Leggings; feet: Skulker's Leather Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Barman Shanker; off_hand: Julie's Dagger; ranged: Skull Splitting Crossbow

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 184.5. Weights run: 3.0s. Verify run: 2.3s. 1367 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): attack_power=1.000 ± 0.001, strength=1.000 ± 0.001, agility=1.136 ± 0.007, crit=0.285 ± 0.011 per rating point (14 rating = 1%, 3.996 per %), hit=0.063 ± 0.006 per rating point (10 rating = 1%, 0.627 per %), melee_haste=not significant (1.050 ± 1.577)

| Slot | Item | Source | Score (attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 40.0 attack_power points (2.39 DPS) | yes | Ebon Mask (19984, +0.00 DPS, sim-verified) [quest]; Champion's Leather Helm (227057, +0.00 DPS) [pvp]; Warlord's Leather Helm (231553, -0.28 DPS) [pvp] |
| neck | Imperial Jewel (11933) | Blackrock Depths: Emperor Dagran Thaurissan [dungeon] | 32.0 attack_power points (1.91 DPS) | yes | Amulet of the Darkmoon (19491, +0.00 DPS, sim-verified) [quest]; Will of the Martyr (17044, -0.12 DPS) [quest]; Medallion of the Dawn (22659, -0.24 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 50.4 attack_power points (3.01 DPS) | yes | Defiler's Lizardhide Shoulders (20175, -0.86 DPS, sim-verified) [rep]; Darkspear Pauldrons (272105, -0.94 DPS) [vendor]; Dark Warder's Pauldrons (22241, -1.20 DPS) [dungeon] |
| back | Deathguard's Cloak (20068) | The Defilers [rep] | 39.7 attack_power points (2.37 DPS) | yes | Cape of the Black Baron (13340, -0.16 DPS) [dungeon]; Howler's Furs (272414, -0.66 DPS) [vendor]; Windshear Cape (20691, -0.87 DPS) [world] |
| chest | Cadaverous Armor (14637) | Scholomance: Lady Illucia Barov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Timbermaw Tunic (252484, -1.17 DPS) [crafted]; Nightbrace Tunic (12603, -1.28 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.59 DPS, sim-verified) [world] |
| wrist | Bracers of the Eclipse (18375) | Dire Maul: Prince Tortheldrin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -0.17 DPS) [rep]; General's Leather Armsplints (16559, -0.23 DPS) [pvp]; Wristwraps of Undead Slaying (23093, -3.42 DPS, sim-verified) [world] |
| hands | Cadaverous Gloves (14640) | Scholomance: Lady Illucia Barov [dungeon] | 44.0 attack_power points (2.63 DPS) | yes | Blood Guard's Leather Vices (16499, -0.04 DPS) [pvp]; Raider Gloves (272099, -0.14 DPS) [vendor]; Skul's Fingerbone Claws (13395, -0.24 DPS) [dungeon] |
| waist | Ferocity of the Timbermaw (227805) | Meilosh [vendor] | 45.3 attack_power points (2.70 DPS) | yes | General's Leather Girdle (16557, -0.14 DPS) [pvp]; Girdle of Beastial Fury (11686, -0.44 DPS) [dungeon]; Cadaverous Belt (14636, -1.27 DPS, sim-verified) [dungeon] |
| legs | Cadaverous Leggings (14638) | Scholomance: Lady Illucia Barov [dungeon] | 52.0 attack_power points (3.10 DPS) | yes | Devilsaur Leggings (15062, -0.12 DPS) [crafted]; Warbear Woolies (15065, -0.21 DPS) [crafted]; Sentinel's Leather Pants (237818, -0.80 DPS) [vendor] |
| feet | Pads of the Dread Wolf (13210) | Blackrock Spire: Halycon [dungeon] | 40.0 attack_power points (2.39 DPS) | yes | Drudge Boots (21532, -0.46 DPS) [quest]; Dunestalker's Boots (20715, -0.56 DPS) [quest]; Defiler's Leather Boots (20186, -0.62 DPS) [rep] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.20 DPS) [quest]; Blackstone Ring (17713, -0.23 DPS) [dungeon]; Naglering (11669, -2.58 DPS, sim-verified) [dungeon] |
| finger2 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Signet Ring of the Bronze Dragonflight (21201, -0.17 DPS) [quest]; Blackstone Ring (17713, -0.20 DPS) [dungeon]; Naglering (11669, -1.99 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Hand of Justice (11815, +0.00 DPS) [dungeon]; Blackhand's Breadth (13965, +0.00 DPS) [quest]; Counterattack Lodestone (18537, +0.00 DPS) [dungeon] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18465, -0.86 DPS, sim-verified) [quest]; Counterattack Lodestone (18537, -1.22 DPS) [dungeon]; Hand of Justice (11815, -1.34 DPS) [dungeon] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (+5.2 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Spellblade (234550, +0.00 DPS) [pvp]; High Warlord's Razor (234556, +0.00 DPS) [pvp]; High Warlord's Shiv (235478, +0.00 DPS) [vendor] |
| off_hand | Felstriker (12590) | Blackrock Spire: Warchief Rend Blackhand [dungeon] | sim-verified (184.5 DPS) | yes | Distracting Dagger (18392, -4.03 DPS) [dungeon]; The Lobotomizer (19324, -4.38 DPS, sim-verified) [rep]; Tome of Knowledge (13385, -37.08 DPS) [dungeon] |
| ranged | Riphook (12653) | Blackrock Spire: Shadow Hunter Vosh'gajin [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malgen's Long Bow (22318, -0.12 DPS) [dungeon]; Skull Splitting Crossbow (13039, -0.30 DPS) [world_drop]; Dark Iron Rifle (16004, -1.83 DPS, sim-verified) [crafted] |

**New at 60:** neck: Imperial Jewel; shoulder: Defiler's Leather Shoulders; back: Deathguard's Cloak; chest: Cadaverous Armor; wrist: Bracers of the Eclipse; hands: Cadaverous Gloves; waist: Ferocity of the Timbermaw; legs: Cadaverous Leggings; feet: Pads of the Dread Wolf; finger1: Legionnaire's Band; finger2: White Bone Band; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: Shadowsong's Sorrow; off_hand: Felstriker; ranged: Riphook

No-known-source sample (15 of 1367, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

